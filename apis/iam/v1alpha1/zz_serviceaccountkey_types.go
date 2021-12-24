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

type ServiceAccountKeyObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	EncryptedPrivateKey *string `json:"encryptedPrivateKey,omitempty" tf:"encrypted_private_key,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	KeyFingerprint *string `json:"keyFingerprint,omitempty" tf:"key_fingerprint,omitempty"`

	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`
}

type ServiceAccountKeyParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Format *string `json:"format,omitempty" tf:"format,omitempty"`

	// +kubebuilder:validation:Optional
	KeyAlgorithm *string `json:"keyAlgorithm,omitempty" tf:"key_algorithm,omitempty"`

	// +kubebuilder:validation:Optional
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`

	// +crossplane:generate:reference:type=ServiceAccount
	// +kubebuilder:validation:Optional
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceAccountIDRef *v1.Reference `json:"serviceAccountIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceAccountIDSelector *v1.Selector `json:"serviceAccountIdSelector,omitempty" tf:"-"`
}

// ServiceAccountKeySpec defines the desired state of ServiceAccountKey
type ServiceAccountKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceAccountKeyParameters `json:"forProvider"`
}

// ServiceAccountKeyStatus defines the observed state of ServiceAccountKey.
type ServiceAccountKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceAccountKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountKey is the Schema for the ServiceAccountKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type ServiceAccountKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceAccountKeySpec   `json:"spec"`
	Status            ServiceAccountKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountKeyList contains a list of ServiceAccountKeys
type ServiceAccountKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceAccountKey `json:"items"`
}

// Repository type metadata.
var (
	ServiceAccountKey_Kind             = "ServiceAccountKey"
	ServiceAccountKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceAccountKey_Kind}.String()
	ServiceAccountKey_KindAPIVersion   = ServiceAccountKey_Kind + "." + CRDGroupVersion.String()
	ServiceAccountKey_GroupVersionKind = CRDGroupVersion.WithKind(ServiceAccountKey_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceAccountKey{}, &ServiceAccountKeyList{})
}
