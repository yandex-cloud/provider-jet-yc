/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SymmetricKeyIamBindingObservation struct {
}

type SymmetricKeyIamBindingParameters struct {

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	SleepAfter *int64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`

	// +kubebuilder:validation:Required
	SymmetricKeyID *string `json:"symmetricKeyId" tf:"symmetric_key_id,omitempty"`
}

// SymmetricKeyIamBindingSpec defines the desired state of SymmetricKeyIamBinding
type SymmetricKeyIamBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SymmetricKeyIamBindingParameters `json:"forProvider"`
}

// SymmetricKeyIamBindingStatus defines the observed state of SymmetricKeyIamBinding.
type SymmetricKeyIamBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SymmetricKeyIamBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SymmetricKeyIamBinding is the Schema for the SymmetricKeyIamBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type SymmetricKeyIamBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SymmetricKeyIamBindingSpec   `json:"spec"`
	Status            SymmetricKeyIamBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SymmetricKeyIamBindingList contains a list of SymmetricKeyIamBindings
type SymmetricKeyIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SymmetricKeyIamBinding `json:"items"`
}

// Repository type metadata.
var (
	SymmetricKeyIamBinding_Kind             = "SymmetricKeyIamBinding"
	SymmetricKeyIamBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SymmetricKeyIamBinding_Kind}.String()
	SymmetricKeyIamBinding_KindAPIVersion   = SymmetricKeyIamBinding_Kind + "." + CRDGroupVersion.String()
	SymmetricKeyIamBinding_GroupVersionKind = CRDGroupVersion.WithKind(SymmetricKeyIamBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&SymmetricKeyIamBinding{}, &SymmetricKeyIamBindingList{})
}
