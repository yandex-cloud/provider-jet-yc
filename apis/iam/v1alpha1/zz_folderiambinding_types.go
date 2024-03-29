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

type FolderIAMBindingObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FolderIAMBindingParameters struct {

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	// (Required) ID of the folder to attach a policy to.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/yandex-cloud/provider-jet-yc/config/iam.ServiceAccountRefValue()
	// +crossplane:generate:reference:refFieldName=ServiceAccountsRef
	// +crossplane:generate:reference:selectorFieldName=ServiceAccountsSelector
	// +kubebuilder:validation:Optional
	// (Required) An array of identities that will be granted the privilege that is specified in the `role` field.
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) The role that should be assigned. Only one
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceAccountsRef []v1.Reference `json:"serviceAccountsRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceAccountsSelector *v1.Selector `json:"serviceAccountsSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

// FolderIAMBindingSpec defines the desired state of FolderIAMBinding
type FolderIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FolderIAMBindingParameters `json:"forProvider"`
}

// FolderIAMBindingStatus defines the observed state of FolderIAMBinding.
type FolderIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FolderIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FolderIAMBinding is the Schema for the FolderIAMBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type FolderIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FolderIAMBindingSpec   `json:"spec"`
	Status            FolderIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FolderIAMBindingList contains a list of FolderIAMBindings
type FolderIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FolderIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	FolderIAMBinding_Kind             = "FolderIAMBinding"
	FolderIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FolderIAMBinding_Kind}.String()
	FolderIAMBinding_KindAPIVersion   = FolderIAMBinding_Kind + "." + CRDGroupVersion.String()
	FolderIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(FolderIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&FolderIAMBinding{}, &FolderIAMBindingList{})
}
