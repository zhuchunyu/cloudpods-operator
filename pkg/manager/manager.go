// Copyright 2019 Yunion
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

package manager

import "yunion.io/x/onecloud-operator/pkg/apis/onecloud/v1alpha1"

// Manager implements the logic for syncing onecloud cluster.
type Manager interface {
	// Sync implements the logic for syncing onecloud cluster.
	Sync(cluster *v1alpha1.OnecloudCluster) error
	GetComponentType() v1alpha1.ComponentType
	IsDisabled(cluster *v1alpha1.OnecloudCluster) bool
}

type ServiceManager interface {
	Manager

	GetServiceName() string
}
