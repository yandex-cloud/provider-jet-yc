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

type KafkaUserObservation struct {
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the user.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Set of permissions granted to the user. The structure is documented below.
	Permission []KafkaUserPermissionObservation `json:"permission,omitempty" tf:"permission,omitempty"`
}

type KafkaUserParameters struct {

	// +crossplane:generate:reference:type=KafkaCluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a KafkaCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a KafkaCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// The name of the user.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The password of the user.
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Set of permissions granted to the user. The structure is documented below.
	// +kubebuilder:validation:Optional
	Permission []KafkaUserPermissionParameters `json:"permission,omitempty" tf:"permission,omitempty"`
}

type KafkaUserPermissionObservation struct {

	// Set of hosts, to which this permission grants access to.
	AllowHosts []*string `json:"allowHosts,omitempty" tf:"allow_hosts,omitempty"`

	// The role type to grant to the topic.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The name of the topic that the permission grants access to.
	TopicName *string `json:"topicName,omitempty" tf:"topic_name,omitempty"`
}

type KafkaUserPermissionParameters struct {

	// Set of hosts, to which this permission grants access to.
	// +kubebuilder:validation:Optional
	AllowHosts []*string `json:"allowHosts,omitempty" tf:"allow_hosts,omitempty"`

	// The role type to grant to the topic.
	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// The name of the topic that the permission grants access to.
	// +kubebuilder:validation:Required
	TopicName *string `json:"topicName" tf:"topic_name,omitempty"`
}

// KafkaUserSpec defines the desired state of KafkaUser
type KafkaUserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KafkaUserParameters `json:"forProvider"`
}

// KafkaUserStatus defines the observed state of KafkaUser.
type KafkaUserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KafkaUserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KafkaUser is the Schema for the KafkaUsers API. Manages a user of a Kafka cluster within Yandex.Cloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type KafkaUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.passwordSecretRef)",message="passwordSecretRef is a required parameter"
	Spec   KafkaUserSpec   `json:"spec"`
	Status KafkaUserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KafkaUserList contains a list of KafkaUsers
type KafkaUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaUser `json:"items"`
}

// Repository type metadata.
var (
	KafkaUser_Kind             = "KafkaUser"
	KafkaUser_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KafkaUser_Kind}.String()
	KafkaUser_KindAPIVersion   = KafkaUser_Kind + "." + CRDGroupVersion.String()
	KafkaUser_GroupVersionKind = CRDGroupVersion.WithKind(KafkaUser_Kind)
)

func init() {
	SchemeBuilder.Register(&KafkaUser{}, &KafkaUserList{})
}