/*
Copyright 2023 clementreiffers.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WorkerVersionSpec defines the desired state of WorkerVersion
type WorkerVersionSpec struct {
	Accounts string `json:"accounts"`
	Scripts  string `json:"scripts"`
	Url      string `json:"url"`
}

// WorkerVersionStatus defines the observed state of WorkerVersion
type WorkerVersionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// WorkerVersion is the Schema for the workerversions API
type WorkerVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkerVersionSpec   `json:"spec,omitempty"`
	Status WorkerVersionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// WorkerVersionList contains a list of WorkerVersion
type WorkerVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkerVersion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WorkerVersion{}, &WorkerVersionList{})
}
