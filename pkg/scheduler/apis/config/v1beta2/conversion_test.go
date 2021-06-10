/*
Copyright 2021 The Kubernetes Authors.

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

package v1beta2

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kube-scheduler/config/v1beta2"
	"k8s.io/kubernetes/pkg/scheduler/apis/config"
	"k8s.io/utils/pointer"
)

func TestV1beta2ToConfigKubeSchedulerConfigurationConversion(t *testing.T) {
	cases := []struct {
		name   string
		config v1beta2.KubeSchedulerConfiguration
		want   config.KubeSchedulerConfiguration
	}{
		{
			name:   "default conversion v1beta2 to config",
			config: v1beta2.KubeSchedulerConfiguration{},
			want: config.KubeSchedulerConfiguration{
				AlgorithmSource: config.SchedulerAlgorithmSource{
					Provider: pointer.StringPtr(v1beta2.SchedulerDefaultProviderName),
				},
			},
		},
	}

	scheme := runtime.NewScheme()
	if err := AddToScheme(scheme); err != nil {
		t.Fatal(err)
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var got config.KubeSchedulerConfiguration
			if err := scheme.Convert(&tc.config, &got, nil); err != nil {
				t.Errorf("failed to convert: %+v", err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("unexpected conversion (-want, +got):\n%s", diff)
			}
		})
	}
}
