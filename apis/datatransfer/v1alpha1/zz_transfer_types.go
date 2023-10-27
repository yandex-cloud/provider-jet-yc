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

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TransferObservation struct {

	// Arbitrary description text for the transfer.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the folder to create the transfer in. If it is not provided, the default provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// (Computed) Identifier of a new Data Transfer transfer.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to the Data Transfer transfer.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the transfer.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the source endpoint for the transfer.
	SourceID *string `json:"sourceId,omitempty" tf:"source_id,omitempty"`

	// ID of the target endpoint for the transfer.
	TargetID *string `json:"targetId,omitempty" tf:"target_id,omitempty"`

	// Type of the transfer. One of "SNAPSHOT_ONLY", "INCREMENT_ONLY", "SNAPSHOT_AND_INCREMENT".
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (Computed) Error description if transfer has any errors.
	Warning *string `json:"warning,omitempty" tf:"warning,omitempty"`
}

type TransferParameters struct {

	// Arbitrary description text for the transfer.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the folder to create the transfer in. If it is not provided, the default provider folder is used.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the Data Transfer transfer.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the transfer.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the source endpoint for the transfer.
	// +crossplane:generate:reference:type=Endpoint
	// +kubebuilder:validation:Optional
	SourceID *string `json:"sourceId,omitempty" tf:"source_id,omitempty"`

	// Reference to a Endpoint to populate sourceId.
	// +kubebuilder:validation:Optional
	SourceIDRef *v1.Reference `json:"sourceIdRef,omitempty" tf:"-"`

	// Selector for a Endpoint to populate sourceId.
	// +kubebuilder:validation:Optional
	SourceIDSelector *v1.Selector `json:"sourceIdSelector,omitempty" tf:"-"`

	// ID of the target endpoint for the transfer.
	// +crossplane:generate:reference:type=Endpoint
	// +kubebuilder:validation:Optional
	TargetID *string `json:"targetId,omitempty" tf:"target_id,omitempty"`

	// Reference to a Endpoint to populate targetId.
	// +kubebuilder:validation:Optional
	TargetIDRef *v1.Reference `json:"targetIdRef,omitempty" tf:"-"`

	// Selector for a Endpoint to populate targetId.
	// +kubebuilder:validation:Optional
	TargetIDSelector *v1.Selector `json:"targetIdSelector,omitempty" tf:"-"`

	// Type of the transfer. One of "SNAPSHOT_ONLY", "INCREMENT_ONLY", "SNAPSHOT_AND_INCREMENT".
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// TransferSpec defines the desired state of Transfer
type TransferSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransferParameters `json:"forProvider"`
}

// TransferStatus defines the observed state of Transfer.
type TransferStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransferObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Transfer is the Schema for the Transfers API. Manages a Data Transfer transfer within Yandex.Cloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type Transfer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransferSpec   `json:"spec"`
	Status            TransferStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransferList contains a list of Transfers
type TransferList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Transfer `json:"items"`
}

// Repository type metadata.
var (
	Transfer_Kind             = "Transfer"
	Transfer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Transfer_Kind}.String()
	Transfer_KindAPIVersion   = Transfer_Kind + "." + CRDGroupVersion.String()
	Transfer_GroupVersionKind = CRDGroupVersion.WithKind(Transfer_Kind)
)

func init() {
	SchemeBuilder.Register(&Transfer{}, &TransferList{})
}
