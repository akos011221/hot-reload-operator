/*
Copyright 2025 Akos Orban.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HotReloadProjectSpec defines the desired state of HotReloadProject.
type HotReloadProjectSpec struct {
	// Path to the Dockerfile used for building the container image.
	// Example: "./Dockerfile" or "docker/my.Dockerfile"
	Dockerfile string `json:"dockerfile"`

	// Full reference to the container image (registry/repository:tag).
	// Example: "localhost:5000/my-api:v1" or "docker.io/username/app:latest"
	Image string `json:"image"`

	// Name of the Kubernetes Deployment that will be updated on changes.
	// Must exist in the same namespace as the HotReloadProject.
	// Example: "nginx-deployment"
	Deployment string `json:"deployment"`
}

// HotReloadProjectStatus defines the observed state of HotReloadProject.
type HotReloadProjectStatus struct {
	// LastSyncTime is when the last code sync occurred
	// +optional
	LastSyncTime *metav1.Time `json:"lastSyncTime,omitempty"`

	// LastBuildTime is when the last build occured
	// +optional
	LastBuildTime *metav1.Time `json:"lastBuildTime,omitempty"`

	// LastDeploymentTime is when the last deployment update occured
	// +optional
	LastDeploymentTime *metav1.Time `json:"lastDeploymentTime,omitempty"`

	// Human-readable message for the status
	// +optional
	Message string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// HotReloadProject is the Schema for the hotreloadprojects API.
type HotReloadProject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HotReloadProjectSpec   `json:"spec,omitempty"`
	Status HotReloadProjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HotReloadProjectList contains a list of HotReloadProject.
type HotReloadProjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HotReloadProject `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HotReloadProject{}, &HotReloadProjectList{})
}
