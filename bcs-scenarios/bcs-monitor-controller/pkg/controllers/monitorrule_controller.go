/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package controllers

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"time"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/util/retry"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	monitorextensionv1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-monitor-controller/api/v1"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-monitor-controller/pkg/apiclient"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-monitor-controller/pkg/fileoperator"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-monitor-controller/pkg/utils"
)

// MonitorRuleReconciler reconciles a MonitorRule object
type MonitorRuleReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	Ctx           context.Context
	FileOp        *fileoperator.FileOperator
	MonitorApiCli apiclient.IMonitorApiClient
}

// +kubebuilder:rbac:groups=monitorextension.bkbcs.tencent.com,resources=monitorrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=monitorextension.bkbcs.tencent.com,resources=monitorrules/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=monitorextension.bkbcs.tencent.com,resources=monitorrules/finalizers,verbs=update

// Reconcile monitor rule
func (r *MonitorRuleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	blog.Infof("MonitorRule '%s' triggered", req.NamespacedName)

	monitorRule := &monitorextensionv1.MonitorRule{}
	if err := r.Get(context.Background(), req.NamespacedName, monitorRule); err != nil {
		if !k8serrors.IsNotFound(err) {
			blog.Errorf("Get MonitorRule '%s' failed, err: %s", req.NamespacedName.String(), err.Error())
			return ctrl.Result{}, err
		}

		blog.Infof("MonitorRule '%s' is deleted, skip...", req.NamespacedName.String())
		return ctrl.Result{}, nil
	}

	if monitorRule.DeletionTimestamp != nil {
		blog.Infof("found deleting monitor rule '%s'", req.NamespacedName)
		if err := r.processDelete(monitorRule); err != nil {
			return ctrl.Result{}, err
		}

		blog.Infof("delete monitor rule '%s' success", req.NamespacedName)
		return ctrl.Result{}, nil
	}

	if err := r.checkFinalizer(monitorRule); err != nil {
		return ctrl.Result{}, err
	}

	outputPath, err := r.FileOp.Compress(monitorRule)
	if err != nil {
		blog.Errorf("compress monitor rule '%s/%s' failed, err: %s", monitorRule.Namespace, monitorRule.Name,
			err.Error())
		if inErr := r.updateSyncStatus(monitorRule, monitorextensionv1.SyncStateFailed, err); inErr != nil {
			blog.Warnf("update monitorRule '%s/%s' sync status failed, err: %s", monitorRule.GetNamespace(),
				monitorRule.GetName(), inErr.Error())
		}
		return ctrl.Result{}, err
	}
	defer os.Remove(outputPath)

	if err = r.MonitorApiCli.UploadConfig(monitorRule.Spec.BizID, monitorRule.Spec.BizToken, outputPath,
		r.getAppName(monitorRule), monitorRule.Spec.Override); err != nil {
		blog.Errorf("upload config to monitor failed, err: %s", err.Error())
		if inErr := r.updateSyncStatus(monitorRule, monitorextensionv1.SyncStateFailed, err); inErr != nil {
			blog.Warnf("update monitorRule '%s/%s' sync status failed, err: %s", monitorRule.GetNamespace(),
				monitorRule.GetName(), inErr.Error())
		}
		return ctrl.Result{}, err
	}

	blog.Infof("sync monitorRule '%s' success", req.NamespacedName)
	if inErr := r.updateSyncStatus(monitorRule, monitorextensionv1.SyncStateCompleted, nil); inErr != nil {
		blog.Warnf("update monitorRule '%s/%s' sync status failed, err: %s", monitorRule.GetNamespace(),
			monitorRule.GetName(), inErr.Error())
	}

	return ctrl.Result{}, nil
}

func (r *MonitorRuleReconciler) eventPredicate() predicate.Predicate {
	return predicate.Funcs{
		CreateFunc: func(createEvent event.CreateEvent) bool {
			mr := createEvent.Object.(*monitorextensionv1.MonitorRule)
			if mr.DeletionTimestamp == nil && mr.Status.SyncStatus.State == monitorextensionv1.SyncStateCompleted &&
				mr.Spec.IgnoreChange == true {
				blog.V(3).Infof("monitor rule  '%s/%s' got create event, but is synced and ignore change",
					mr.GetNamespace(), mr.GetName())
				return false
			}
			return true
		},
		UpdateFunc: func(e event.UpdateEvent) bool {
			newMr, okNew := e.ObjectNew.(*monitorextensionv1.MonitorRule)
			oldMr, okOld := e.ObjectOld.(*monitorextensionv1.MonitorRule)
			if !okNew || !okOld {
				return true
			}
			if reflect.DeepEqual(newMr.Spec, oldMr.Spec) &&
				reflect.DeepEqual(newMr.Finalizers, oldMr.Finalizers) &&
				reflect.DeepEqual(newMr.DeletionTimestamp, oldMr.DeletionTimestamp) {
				blog.V(5).Infof("monitor rule %+v updated, but spec and finalizer and deletionTimestamp not change",
					newMr)
				return false
			}
			if newMr.DeletionTimestamp == nil && newMr.Status.SyncStatus.
				State == monitorextensionv1.SyncStateCompleted && newMr.Spec.IgnoreChange == true {
				blog.V(3).Infof("monitor rule '%s/%s' updated, but is synced and ignore change",
					newMr.GetNamespace(), newMr.GetName())
				return false
			}
			return true
		},
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *MonitorRuleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&monitorextensionv1.MonitorRule{}).
		WithEventFilter(r.eventPredicate()).
		Complete(r)
}

func (r *MonitorRuleReconciler) updateSyncStatus(monitorRule *monitorextensionv1.MonitorRule,
	state monitorextensionv1.SyncState, err error) error {
	blog.Infof("Update sync state of monitorRule (%s/%s) to %s", monitorRule.GetNamespace(), monitorRule.GetName(), state)
	monitorRule.Status.SyncStatus.State = state
	// err message
	if err != nil {
		monitorRule.Status.SyncStatus.Message = err.Error()
	} else {
		monitorRule.Status.SyncStatus.Message = ""
	}
	monitorRule.Status.SyncStatus.LastSyncTime = metav1.NewTime(time.Now())
	monitorRule.Status.SyncStatus.App = r.getAppName(monitorRule)
	if inErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		return r.Client.Status().Update(r.Ctx, monitorRule)
	}); inErr != nil {
		blog.Warnf("update monitorRule'%s/%s' failed, err: %s", monitorRule.GetNamespace(), monitorRule.GetName(), inErr.Error())
		return inErr
	}

	return nil
}

func (r *MonitorRuleReconciler) getAppName(monitorRule *monitorextensionv1.MonitorRule) string {
	return fmt.Sprintf("bcs-mr-%s-%s", monitorRule.Spec.Scenario, monitorRule.GetName())
}

// checkFinalizer add finalizer if not exist
func (r *MonitorRuleReconciler) checkFinalizer(monitorRule *monitorextensionv1.MonitorRule) error {
	if utils.ContainsString(monitorRule.Finalizers, FinalizerMonitorController) {
		return nil
	}

	monitorRule.Finalizers = append(monitorRule.Finalizers, FinalizerMonitorController)
	if err := r.Update(r.Ctx, monitorRule); err != nil {
		blog.Warnf("Update monitor rule'%s/%s' failed, err: %s", monitorRule.Namespace, monitorRule.Name,
			err.Error())
		return err
	}
	return nil
}

func (r *MonitorRuleReconciler) removeFinalizer(monitorRule *monitorextensionv1.MonitorRule) error {
	monitorRule.Finalizers = utils.RemoveString(monitorRule.Finalizers, FinalizerMonitorController)
	if err := r.Update(context.Background(), monitorRule, &client.UpdateOptions{}); err != nil {
		blog.Warnf("remove finalizer for monitorRule %s/%s failed, err %s", monitorRule.GetNamespace(), monitorRule.GetName(),
			err.Error())
		return fmt.Errorf("remove finalizer for monitorRule %s/%s failed, err %s", monitorRule.GetNamespace(),
			monitorRule.GetName(), err.Error())
	}
	blog.V(3).Infof("remove finalizer for monitorRule %s/%s successfully", monitorRule.GetNamespace(),
		monitorRule.GetName())
	return nil
}

func (r *MonitorRuleReconciler) processDelete(monitorRule *monitorextensionv1.MonitorRule) error {
	if err := r.MonitorApiCli.UploadConfig(monitorRule.Spec.BizID, monitorRule.Spec.BizToken, EmptyTARLocation,
		r.getAppName(monitorRule), monitorRule.Spec.Override); err != nil {
		blog.Errorf("upload config to monitor failed, err: %s", err.Error())
		if inErr := r.updateSyncStatus(monitorRule, monitorextensionv1.SyncStateFailed, err); inErr != nil {
			blog.Warnf("update monitorRule '%s/%s' sync status failed, err: %s", monitorRule.GetNamespace(),
				monitorRule.GetName(), inErr.Error())
		}
		return err
	}

	if err := r.removeFinalizer(monitorRule); err != nil {
		return err
	}

	blog.Infof("delete monitorRule '%s/%s' success", monitorRule.GetNamespace(), monitorRule.GetName())
	return nil
}
