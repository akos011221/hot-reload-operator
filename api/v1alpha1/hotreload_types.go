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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HotReloadSpec defines the desired state of HotReload.
type HotReloadSpec struct {
	// SourceSync defines how source code changes are detected and synchronized
	SourceSync SourceSyncSpec `json:"sourceSync"`

	// Build defines the build configuration for the application
	Build BuildSpec `json:"build"`

	// Deployment defines the target deployment to manage
	Deployment DeploymentSpec `json:"deployment"`
}

// SourceSyncSpec defines the source code synchronization configuration
type SourceSyncSpec struct {
	// Webhook endpoint that will be notified when changes are detected
	Webhook string `json:"webhook"`

	// List of paths/patterns to ignore when watching for changes
	IgnorePaths []string `json:"ignorePaths"`
}

// BuildSpec defines the build configuration
type BuildSpec struct {
	// Path to the Dockerfile
	Dockerfile string `json:"dockerfile"`

	// Build context path
	Context string `json:"context"`

	// Target registry for the built image
	Registry string `json:"registry"`

	// Name of the image to build
	ImageName string `json:"imageName"`

	// Additional build arguments
	// +optional
	Args map[string]string `json:"args,omitempty"`
}

// DeploymentSpec defines the target deployment configuration
type DeploymentSpec struct {
	// Name of the Deployment to manage
	Name string `json:"name"`

	// Namespace of the Deployment
	// +optional
	// +kubebuilder:default=default
	Namespace string `json:"namespace,omitempty"`

	// Name of the container in the Deployment to update
	ContainerName string `json:"containerName"`
}

// HotReloadStatus defines the observed state of HotReload.
type HotReloadStatus struct {
	// LastSyncTime is when the last code sync occurred
	// +optional
	LastSyncTime *metav1.Time `json:"lastSyncTime,omitempty"`

	// LastBuildTime is when the last  build occured
	// +optional
	LastBuildTime *metav1.Time `json:"lastBuildTime,omitempty"`

	// LastDeploymentTime is when the last deployment update occured
	// +optional
	LastDeploymentTime *metav1.Time `json:"lastDeploymentTime,omitempty"`

	// Conditions represents the current state of the HotReload resource
	// +Optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
//+kubebuilder:resource:path=cleaners,scope=Cluster
// +kubebuilder:subresource:status

// HotReload is the Schema for the hotreloads API.
type HotReload struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HotReloadSpec   `json:"spec,omitempty"`
	Status HotReloadStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HotReloadList contains a list of HotReload.
type HotReloadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HotReload `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HotReload{}, &HotReloadList{})
}
