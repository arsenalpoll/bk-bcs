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

package run

import "time"

// Limiter request number limiter.
type Limiter struct {
	limit chan struct{}
}

// NewStreamLimiter new stream limiter.
func NewStreamLimiter(concurrent int) *Limiter {
	if 0 == concurrent {
		concurrent = 100
	}
	return &Limiter{
		limit: make(chan struct{}, concurrent),
	}
}

// IsEmpty test limiter if empty.
func (sl *Limiter) IsEmpty() bool {
	return len(sl.limit) == 0
}

// Require limiter require.
func (sl *Limiter) Require() {
	sl.limit <- struct{}{}
}

// Release limiter release.
func (sl *Limiter) Release() {
	<-sl.limit
}

// Execute please use this func carefully.
// for instance, when you use this func in a for loop, please
// be sure that the f variable is concurrent secure.
func (sl *Limiter) Execute(ch chan<- *Status, f func() error) {
	sl.Require()
	go func() {
		defer sl.Release()
		s := new(Status)
		start := time.Now()
		s.Error = f()
		s.CostDuration = time.Since(start)
		ch <- s
		return
	}()
	return
}
