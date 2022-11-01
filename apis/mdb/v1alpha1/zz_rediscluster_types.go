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

type RedisClusterConfigObservation struct {
}

type RedisClusterConfigParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Normal clients output buffer limits.
	ClientOutputBufferLimitNormal *string `json:"clientOutputBufferLimitNormal,omitempty" tf:"client_output_buffer_limit_normal,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Pubsub clients output buffer limits.
	ClientOutputBufferLimitPubsub *string `json:"clientOutputBufferLimitPubsub,omitempty" tf:"client_output_buffer_limit_pubsub,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Number of databases (changing requires redis-server restart).
	Databases *float64 `json:"databases,omitempty" tf:"databases,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Redis key eviction policy for a dataset that reaches maximum memory.
	MaxmemoryPolicy *string `json:"maxmemoryPolicy,omitempty" tf:"maxmemory_policy,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Select the events that Redis will notify among a set of classes.
	NotifyKeyspaceEvents *string `json:"notifyKeyspaceEvents,omitempty" tf:"notify_keyspace_events,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	// (Optional) Log slow queries below this number in microseconds.
	SlowlogLogSlowerThan *float64 `json:"slowlogLogSlowerThan,omitempty" tf:"slowlog_log_slower_than,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Slow queries log length.
	SlowlogMaxLen *float64 `json:"slowlogMaxLen,omitempty" tf:"slowlog_max_len,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Close the connection after a client is idle for N seconds.
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Version of Redis (5.0, 6.0 or 6.2).
	Version *string `json:"version" tf:"version,omitempty"`
}

type RedisClusterHostObservation struct {
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
}

type RedisClusterHostParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Sets whether the host should get a public IP address or not.
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Replica priority of a current replica (usable for non-sharded only).
	ReplicaPriority *float64 `json:"replicaPriority,omitempty" tf:"replica_priority,omitempty"`

	// +kubebuilder:validation:Optional
	ShardName *string `json:"shardName,omitempty" tf:"shard_name,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	// (Required) The availability zone where the Redis host will be created.
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type RedisClusterMaintenanceWindowObservation struct {
}

type RedisClusterMaintenanceWindowParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Day of week for maintenance window if window type is weekly. Possible values: `MON`, `TUE`, `WED`, `THU`, `FRI`, `SAT`, `SUN`.
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Hour of day in UTC time zone (1-24) for maintenance window if window type is weekly.
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Type of maintenance window. Can be either `ANYTIME` or `WEEKLY`. A day and hour of window need to be specified with weekly window.
	Type *string `json:"type" tf:"type,omitempty"`
}

type RedisClusterObservation struct {
	// Creation timestamp of the key.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Aggregated health of the cluster. Can be either `ALIVE`, `DEGRADED`, `DEAD` or `HEALTH_UNKNOWN`.
	Health *string `json:"health,omitempty" tf:"health,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) A host of the Redis cluster. The structure is documented below.
	Host []RedisClusterHostObservation `json:"host,omitempty" tf:"host,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Status of the cluster. Can be either `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or `STATUS_UNKNOWN`.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type RedisClusterParameters struct {

	// +kubebuilder:validation:Required
	// (Required) Configuration of the Redis cluster. The structure is documented below.
	Config []RedisClusterConfigParameters `json:"config" tf:"config,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Description of the Redis cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Deployment environment of the Redis cluster. Can be either `PRESTABLE` or `PRODUCTION`.
	Environment *string `json:"environment" tf:"environment,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	// (Optional) The ID of the folder that the resource belongs to. If it
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	// (Required) A host of the Redis cluster. The structure is documented below.
	Host []RedisClusterHostParameters `json:"host" tf:"host,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) A set of key/value label pairs to assign to the Redis cluster.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	MaintenanceWindow []RedisClusterMaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Name of the Redis cluster. Provided by the client when the cluster is created.
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Network
	// +kubebuilder:validation:Optional
	// (Required) ID of the network, to which the Redis cluster belongs.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	// (Optional) Persistence mode.
	PersistenceMode *string `json:"persistenceMode,omitempty" tf:"persistence_mode,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Resources allocated to hosts of the Redis cluster. The structure is documented below.
	Resources []RedisClusterResourcesParameters `json:"resources" tf:"resources,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.SecurityGroup
	// +kubebuilder:validation:Optional
	// (Optional) A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to SecurityGroup in vpc to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsRefs []v1.Reference `json:"securityGroupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in vpc to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsSelector *v1.Selector `json:"securityGroupIdsSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	// (Optional) Redis Cluster mode enabled/disabled.
	Sharded *bool `json:"sharded,omitempty" tf:"sharded,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) TLS support mode enabled/disabled.
	TLSEnabled *bool `json:"tlsEnabled,omitempty" tf:"tls_enabled,omitempty"`
}

type RedisClusterResourcesObservation struct {
}

type RedisClusterResourcesParameters struct {

	// +kubebuilder:validation:Required
	// (Required) Volume of the storage available to a host, in gigabytes.
	DiskSize *float64 `json:"diskSize" tf:"disk_size,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Type of the storage of Redis hosts - environment default is used if missing.
	DiskTypeID *string `json:"diskTypeId,omitempty" tf:"disk_type_id,omitempty"`

	// +kubebuilder:validation:Required
	ResourcePresetID *string `json:"resourcePresetId" tf:"resource_preset_id,omitempty"`
}

// RedisClusterSpec defines the desired state of RedisCluster
type RedisClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RedisClusterParameters `json:"forProvider"`
}

// RedisClusterStatus defines the observed state of RedisCluster.
type RedisClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RedisClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RedisCluster is the Schema for the RedisClusters API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type RedisCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisClusterSpec   `json:"spec"`
	Status            RedisClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisClusterList contains a list of RedisClusters
type RedisClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisCluster `json:"items"`
}

// Repository type metadata.
var (
	RedisCluster_Kind             = "RedisCluster"
	RedisCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RedisCluster_Kind}.String()
	RedisCluster_KindAPIVersion   = RedisCluster_Kind + "." + CRDGroupVersion.String()
	RedisCluster_GroupVersionKind = CRDGroupVersion.WithKind(RedisCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&RedisCluster{}, &RedisClusterList{})
}
