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

type IamBindingObservation struct {
}

type IamBindingParameters struct {

	// +kubebuilder:validation:Required
	FunctionID *string `json:"functionId" tf:"function_id,omitempty"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	SleepAfter *int64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

// IamBindingSpec defines the desired state of IamBinding
type IamBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IamBindingParameters `json:"forProvider"`
}

// IamBindingStatus defines the observed state of IamBinding.
type IamBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IamBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IamBinding is the Schema for the IamBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type IamBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamBindingSpec   `json:"spec"`
	Status            IamBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamBindingList contains a list of IamBindings
type IamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamBinding `json:"items"`
}

// Repository type metadata.
var (
	IamBinding_Kind             = "IamBinding"
	IamBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IamBinding_Kind}.String()
	IamBinding_KindAPIVersion   = IamBinding_Kind + "." + CRDGroupVersion.String()
	IamBinding_GroupVersionKind = CRDGroupVersion.WithKind(IamBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&IamBinding{}, &IamBindingList{})
}
