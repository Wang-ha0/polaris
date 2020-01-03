// Copyright 2019 FairwindsOps Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validator

import (
	"github.com/fairwindsops/polaris/pkg/config"
	corev1 "k8s.io/api/core/v1"
)

func ValidateContainer(conf *config.Configuration, basePod *corev1.PodSpec, container *corev1.Container, controllerName string, controllerKind config.SupportedController, isInit bool) ContainerResult {
	results, err := applyContainerSchemaChecks(conf, basePod, container, controllerName, controllerKind, isInit)
	// FIXME: don't panic
	if err != nil {
		panic(err)
	}

	cRes := ContainerResult{
		Name:    container.Name,
		Results: results,
	}

	return cRes
}

func ValidateContainers(conf *config.Configuration, basePod *corev1.PodSpec, containers []corev1.Container, controllerName string, controllerKind config.SupportedController, isInit bool) []ContainerResult {
	results := []ContainerResult{}
	for _, container := range containers {
		cRes := ValidateContainer(conf, basePod, &container, controllerName, controllerKind, isInit)
		results = append(results, cRes)
	}
	return results
}
