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

type HTTPRouterObservation struct {

	// The HTTP Router creation timestamp.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The ID of the HTTP Router.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type HTTPRouterParameters struct {

	// An optional description of the HTTP Router. Provide this property when
	// you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder to which the resource belongs.
	// If omitted, the provider folder is used.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// Labels to assign to this HTTP Router. A list of key/value pairs.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the HTTP Router. Provided by the client when the HTTP Router is created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// HTTPRouterSpec defines the desired state of HTTPRouter
type HTTPRouterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HTTPRouterParameters `json:"forProvider"`
}

// HTTPRouterStatus defines the observed state of HTTPRouter.
type HTTPRouterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HTTPRouterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HTTPRouter is the Schema for the HTTPRouters API. The HTTP router defines the routing rules for HTTP requests to backend groups.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type HTTPRouter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HTTPRouterSpec   `json:"spec"`
	Status            HTTPRouterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HTTPRouterList contains a list of HTTPRouters
type HTTPRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HTTPRouter `json:"items"`
}

// Repository type metadata.
var (
	HTTPRouter_Kind             = "HTTPRouter"
	HTTPRouter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HTTPRouter_Kind}.String()
	HTTPRouter_KindAPIVersion   = HTTPRouter_Kind + "." + CRDGroupVersion.String()
	HTTPRouter_GroupVersionKind = CRDGroupVersion.WithKind(HTTPRouter_Kind)
)

func init() {
	SchemeBuilder.Register(&HTTPRouter{}, &HTTPRouterList{})
}
