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

type FolderIamBindingObservation struct {
}

type FolderIamBindingParameters struct {

	// +kubebuilder:validation:Required
	FolderID *string `json:"folderId" tf:"folder_id,omitempty"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	SleepAfter *int64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

// FolderIamBindingSpec defines the desired state of FolderIamBinding
type FolderIamBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FolderIamBindingParameters `json:"forProvider"`
}

// FolderIamBindingStatus defines the observed state of FolderIamBinding.
type FolderIamBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FolderIamBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FolderIamBinding is the Schema for the FolderIamBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type FolderIamBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FolderIamBindingSpec   `json:"spec"`
	Status            FolderIamBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FolderIamBindingList contains a list of FolderIamBindings
type FolderIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FolderIamBinding `json:"items"`
}

// Repository type metadata.
var (
	FolderIamBinding_Kind             = "FolderIamBinding"
	FolderIamBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FolderIamBinding_Kind}.String()
	FolderIamBinding_KindAPIVersion   = FolderIamBinding_Kind + "." + CRDGroupVersion.String()
	FolderIamBinding_GroupVersionKind = CRDGroupVersion.WithKind(FolderIamBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&FolderIamBinding{}, &FolderIamBindingList{})
}
