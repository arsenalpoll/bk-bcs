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

// Package The main package serves as the entry point for the configuration server.
package main

import (
	"fmt"
	"os"

	"github.com/TencentBlueking/bk-bcs/bcs-services/bcs-bscp/cmd/config-server/app"
	"github.com/TencentBlueking/bk-bcs/bcs-services/bcs-bscp/cmd/config-server/options"
	"github.com/TencentBlueking/bk-bcs/bcs-services/bcs-bscp/pkg/cc"
	"github.com/TencentBlueking/bk-bcs/bcs-services/bcs-bscp/pkg/logs"
)

func main() {
	cc.InitService(cc.ConfigServerName)

	opts := options.InitOptions()
	if err := app.Run(opts); err != nil {
		fmt.Fprintf(os.Stderr, "start config server failed, err: %v", err)
		logs.CloseLogs()
		os.Exit(1)
	}
}
