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
 */

package remove

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/i18n"
	"k8s.io/kubectl/pkg/util/templates"

	nodegroup "github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/node_group"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-cli/bcs-cluster-manager/pkg/manager/types"
)

var (
	removeNodesFromGroupExample = templates.Examples(i18n.T(`
	kubectl-bcs-cluster-manager remove nodesFromGroup --clusterID xxx --nodeGroupID xxx --nodes xxx`))
)

func newRemoveNodesFromGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "nodesFromGroup",
		Short:   "remove nodes to group from bcs-cluster-manager",
		Example: removeNodesFromGroupExample,
		Run:     removeNodesFromGroup,
	}

	cmd.Flags().StringVarP(&clusterID, "clusterID", "c", "", "cluster ID")
	_ = cmd.MarkFlagRequired("clusterID")
	cmd.Flags().StringVarP(&nodeGroupID, "nodeGroupID", "n", "", "node group ID")
	_ = cmd.MarkFlagRequired("nodeGroupID")
	cmd.Flags().StringSliceVarP(&nodes, "nodes", "i", []string{},
		"node inner ip, for example: -i xxx.xxx.xxx.xxx -i xxx.xxx.xxx.xxx")
	_ = cmd.MarkFlagRequired("nodes")

	return cmd
}

func removeNodesFromGroup(cmd *cobra.Command, args []string) {
	resp, err := nodegroup.New(context.Background()).RemoveNodes(types.RemoveNodesFromGroupReq{
		ClusterID:   clusterID,
		NodeGroupID: nodeGroupID,
		Nodes:       nodes,
	})
	if err != nil {
		klog.Fatalf("remove nodes to group failed: %v", err)
	}

	fmt.Printf("remove nodes to group succeed: taskID: %v\n", resp.TaskID)
}
