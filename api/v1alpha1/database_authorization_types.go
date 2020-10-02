/*

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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const Finalizer = "databaseauthorization.containers.kubernetes.ovhcloud.com/finalizer"

// DatabaseAuthorizationSpec defines the desired state of DatabaseAuthorization
type DatabaseAuthorizationSpec struct {
	// Represents the list of services names to synchronize authorized connections
	OvhServices []string `json:"ovhServices"`

	// Represents the API credentiels stored in secrets
	OvhCredentials OvhCredentialsSpec `json:"ovhCredentials"`

	// Represents the node selector to apply for authorization sync
	LabelSelector metav1.LabelSelector `json:"labelSelector"`
}

// OvhCredentialsSpec defined the API credentials stored in secrets
type OvhCredentialsSpec struct {
	Token       v1.SecretReference `json:"token"`
	Application v1.SecretReference `json:"application"`
}

// DatabaseAuthorizationStatus defines the observed state of DatabaseAuthorization
type DatabaseAuthorizationStatus struct {
	// Represents the latest generation value of the object (is incremented for all changed except for metadata and statuses)
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	// Represents the latest available observations of a nodepool current state.
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []DatabaseAuthorizationCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,6,rep,name=conditions"`
}

// DatabaseAuthorizationCondition defined the last status update condition
type DatabaseAuthorizationCondition struct {
	// Type of nodepool condition.
	Type string `json:"type"`

	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`

	// The last time this condition was updated.
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`

	// Last time the condition transitioned from one status to another.
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`

	// The reason for the condition's last transition.
	Reason string `json:"reason"`

	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseAuthorization is the Schema for the DatabaseAuthorization API
type DatabaseAuthorization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatabaseAuthorizationSpec   `json:"spec,omitempty"`
	Status DatabaseAuthorizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseAuthorizationList contains a list of DatabaseAuthorization
type DatabaseAuthorizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []DatabaseAuthorization `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatabaseAuthorization{}, &DatabaseAuthorizationList{})
}
