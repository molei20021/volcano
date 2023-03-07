/*
Copyright 2019 The Volcano Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main // note!!! package must be named main

import (
	"k8s.io/klog"

	"volcano.sh/volcano/pkg/scheduler/api"
	"volcano.sh/volcano/pkg/scheduler/framework"
	"volcano.sh/volcano/pkg/scheduler/plugins/util"
)

const PluginName = "prioritycustom"

type priorityCustomPlugin struct{}

func (mp *priorityCustomPlugin) Name() string {
	return PluginName
}

// New is a PluginBuilder, remove the comment when used.
func New(arguments framework.Arguments) framework.Plugin {
	return &priorityCustomPlugin{}
}

func (mp *priorityCustomPlugin) OnSessionOpen(ssn *framework.Session) {
	klog.V(4).Info("Enter priorityCustomPlugin plugin ...")

	ssn.AddJobEnqueueableFn(mp.Name(), func(obj interface{}) int {
		job := obj.(*api.JobInfo)
		queueID := job.Queue
		queue := ssn.Queues[queueID]
		if queue.Weight < 10 {
			return util.Reject
		}
		return util.Permit
	})
}

func (mp *priorityCustomPlugin) OnSessionClose(ssn *framework.Session) {}
