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

type RecordsetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RecordsetParameters struct {

	// +kubebuilder:validation:Required
	// (Optional) The string data for the records in this record set.
	Data []*string `json:"data" tf:"data,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) The DNS name this record set will apply to.
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	// (Optional) The time-to-live of this record set (seconds).
	TTL *float64 `json:"ttl" tf:"ttl,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) The DNS record set type.
	Type *string `json:"type" tf:"type,omitempty"`

	// +crossplane:generate:reference:type=Zone
	// +kubebuilder:validation:Optional
	// (Required) The id of the zone in which this record set will reside.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

// RecordsetSpec defines the desired state of Recordset
type RecordsetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RecordsetParameters `json:"forProvider"`
}

// RecordsetStatus defines the observed state of Recordset.
type RecordsetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RecordsetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Recordset is the Schema for the Recordsets API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type Recordset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RecordsetSpec   `json:"spec"`
	Status            RecordsetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecordsetList contains a list of Recordsets
type RecordsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Recordset `json:"items"`
}

// Repository type metadata.
var (
	Recordset_Kind             = "Recordset"
	Recordset_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Recordset_Kind}.String()
	Recordset_KindAPIVersion   = Recordset_Kind + "." + CRDGroupVersion.String()
	Recordset_GroupVersionKind = CRDGroupVersion.WithKind(Recordset_Kind)
)

func init() {
	SchemeBuilder.Register(&Recordset{}, &RecordsetList{})
}
