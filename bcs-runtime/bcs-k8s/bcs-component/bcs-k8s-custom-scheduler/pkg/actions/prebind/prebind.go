/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package prebind action for prebind
package prebind

import (
	"fmt"
	"time"

	"github.com/emicklei/go-restful"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"

	"github.com/Tencent/bk-bcs/bcs-common/common/blog"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-k8s-custom-scheduler/pkg/actions"
	v3 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-k8s-custom-scheduler/pkg/ipscheduler/v3"
	"github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/bcs-component/bcs-k8s-custom-scheduler/pkg/metrics"
)

const (
	// PreBindPrefix prebind prefix
	PreBindPrefix = "prebind"
)

func init() {
	actions.RegisterAction(actions.Action{
		Verb:    "POST",
		Path:    actions.BcsCustomSchedulerPrefix + "ipscheduler/" + "{version}/" + PreBindPrefix,
		Params:  nil,
		Handler: handleIpSchedulerPreBind})
}

func handleIpSchedulerPreBind(req *restful.Request, resp *restful.Response) {

	ipSchedulerVersion := req.PathParameter("version")

	var extenderBindingArgs schedulerapi.ExtenderBindingArgs
	var extenderBindingResult *schedulerapi.ExtenderBindingResult
	var metricsArgs = &metrics.RecordConfig{
		Version: ipSchedulerVersion,
		Handler: "handleIpSchedulerPreBind",
		Method:  "POST",
		Status:  "",
		Started: time.Now(),
	}

	defer func() {
		metrics.ReportK8sCustomSchedulerAPIMetrics(metricsArgs.Version, metricsArgs.Handler, metricsArgs.Method,
			metricsArgs.Status, metricsArgs.Started)
	}()

	err := req.ReadEntity(&extenderBindingArgs)
	if err != nil {
		blog.Errorf("error when read request: %s", err.Error())
		extenderBindingResult = &schedulerapi.ExtenderBindingResult{
			Error: err.Error(),
		}

		metricsArgs.Status = metrics.ErrStatus
		resp.WriteEntity(extenderBindingResult)
		return
	}

	if ipSchedulerVersion == actions.IpSchedulerV3 {
		err = v3.HandleIpSchedulerPreBind(extenderBindingArgs)
	} else {
		err = fmt.Errorf("invalid IpScheduler version")
	}
	if err != nil {
		blog.Errorf("error handling extender prebind: %s", err.Error())
		extenderBindingResult = &schedulerapi.ExtenderBindingResult{
			Error: err.Error(),
		}

		metricsArgs.Status = metrics.ErrStatus
		resp.WriteEntity(extenderBindingResult)
		return
	}

	extenderBindingResult = &schedulerapi.ExtenderBindingResult{
		Error: "",
	}

	blog.Info("prebind finished")
	metricsArgs.Status = metrics.ErrStatus
	resp.WriteEntity(extenderBindingResult)
	return
}
