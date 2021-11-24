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

type DataNodeObservation struct {
}

type DataNodeParameters struct {

	// +kubebuilder:validation:Required
	Resources []DataNodeResourcesParameters `json:"resources" tf:"resources,omitempty"`
}

type DataNodeResourcesObservation struct {
}

type DataNodeResourcesParameters struct {

	// +kubebuilder:validation:Required
	DiskSize *int64 `json:"diskSize" tf:"disk_size,omitempty"`

	// +kubebuilder:validation:Required
	DiskTypeID *string `json:"diskTypeId" tf:"disk_type_id,omitempty"`

	// +kubebuilder:validation:Required
	ResourcePresetID *string `json:"resourcePresetId" tf:"resource_preset_id,omitempty"`
}

type ElasticsearchClusterConfigObservation struct {
}

type ElasticsearchClusterConfigParameters struct {

	// +kubebuilder:validation:Required
	AdminPasswordSecretRef v1.SecretKeySelector `json:"adminPasswordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	DataNode []DataNodeParameters `json:"dataNode" tf:"data_node,omitempty"`

	// +kubebuilder:validation:Optional
	Edition *string `json:"edition,omitempty" tf:"edition,omitempty"`

	// +kubebuilder:validation:Optional
	MasterNode []MasterNodeParameters `json:"masterNode,omitempty" tf:"master_node,omitempty"`

	// +kubebuilder:validation:Optional
	Plugins []*string `json:"plugins,omitempty" tf:"plugins,omitempty"`

	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ElasticsearchClusterHostObservation struct {
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
}

type ElasticsearchClusterHostParameters struct {

	// +kubebuilder:validation:Optional
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Required
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type ElasticsearchClusterObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	Health *string `json:"health,omitempty" tf:"health,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ElasticsearchClusterParameters struct {

	// +kubebuilder:validation:Required
	Config []ElasticsearchClusterConfigParameters `json:"config" tf:"config,omitempty"`

	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Environment *string `json:"environment" tf:"environment,omitempty"`

	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// +kubebuilder:validation:Optional
	Host []ElasticsearchClusterHostParameters `json:"host,omitempty" tf:"host,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Required
	NetworkID *string `json:"networkId" tf:"network_id,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`
}

type MasterNodeObservation struct {
}

type MasterNodeParameters struct {

	// +kubebuilder:validation:Required
	Resources []MasterNodeResourcesParameters `json:"resources" tf:"resources,omitempty"`
}

type MasterNodeResourcesObservation struct {
}

type MasterNodeResourcesParameters struct {

	// +kubebuilder:validation:Required
	DiskSize *int64 `json:"diskSize" tf:"disk_size,omitempty"`

	// +kubebuilder:validation:Required
	DiskTypeID *string `json:"diskTypeId" tf:"disk_type_id,omitempty"`

	// +kubebuilder:validation:Required
	ResourcePresetID *string `json:"resourcePresetId" tf:"resource_preset_id,omitempty"`
}

// ElasticsearchClusterSpec defines the desired state of ElasticsearchCluster
type ElasticsearchClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ElasticsearchClusterParameters `json:"forProvider"`
}

// ElasticsearchClusterStatus defines the observed state of ElasticsearchCluster.
type ElasticsearchClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ElasticsearchClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticsearchCluster is the Schema for the ElasticsearchClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type ElasticsearchCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticsearchClusterSpec   `json:"spec"`
	Status            ElasticsearchClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticsearchClusterList contains a list of ElasticsearchClusters
type ElasticsearchClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticsearchCluster `json:"items"`
}

// Repository type metadata.
var (
	ElasticsearchCluster_Kind             = "ElasticsearchCluster"
	ElasticsearchCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ElasticsearchCluster_Kind}.String()
	ElasticsearchCluster_KindAPIVersion   = ElasticsearchCluster_Kind + "." + CRDGroupVersion.String()
	ElasticsearchCluster_GroupVersionKind = CRDGroupVersion.WithKind(ElasticsearchCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&ElasticsearchCluster{}, &ElasticsearchClusterList{})
}
