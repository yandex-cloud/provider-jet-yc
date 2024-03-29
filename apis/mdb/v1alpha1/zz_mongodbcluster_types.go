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

type AccessObservation struct {
	// (Optional) Allow access for DataLens.
	DataLens *bool `json:"dataLens,omitempty" tf:"data_lens,omitempty"`
}

type AccessParameters struct {
}

type AuditLogObservation struct {
}

type AuditLogParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Configuration of the audit log filter in JSON format.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// +kubebuilder:validation:Optional
	RuntimeConfiguration *bool `json:"runtimeConfiguration,omitempty" tf:"runtime_configuration,omitempty"`
}

type BackupWindowStartObservation struct {
}

type BackupWindowStartParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) The hour at which backup will be started.
	Hours *float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) The minute at which backup will be started.
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`
}

type ClusterConfigObservation struct {
	// (Optional) Access policy to the MongoDB cluster. The structure is documented below.
	Access []AccessObservation `json:"access,omitempty" tf:"access,omitempty"`
}

type ClusterConfigParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Access policy to the MongoDB cluster. The structure is documented below.
	Access []AccessParameters `json:"access,omitempty" tf:"access,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Time to start the daily backup, in the UTC timezone. The structure is documented below.
	BackupWindowStart []BackupWindowStartParameters `json:"backupWindowStart,omitempty" tf:"backup_window_start,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Feature compatibility version of MongoDB. If not provided version is taken. Can be either `5.0`, `4.4`, `4.2` and `4.0`.
	FeatureCompatibilityVersion *string `json:"featureCompatibilityVersion,omitempty" tf:"feature_compatibility_version,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Configuration of the mongod service. The structure is documented below.
	Mongod []MongodParameters `json:"mongod,omitempty" tf:"mongod,omitempty"`

	// +kubebuilder:validation:Required
	// (Optional) Version of the MongoDB server software. Can be either `4.0`, `4.2`, `4.4`, `4.4-enterprise`, `5.0` and `5.0-enterprise`.
	Version *string `json:"version" tf:"version,omitempty"`
}

type DatabaseObservation struct {
}

type DatabaseParameters struct {

	// +kubebuilder:validation:Required
	// (Required) Name of the MongoDB cluster. Provided by the client when the cluster is created.
	Name *string `json:"name" tf:"name,omitempty"`
}

type KmipObservation struct {
}

type KmipParameters struct {

	// +kubebuilder:validation:Optional
	// (Required) String containing the client certificate used for authenticating MongoDB to the KMIP server.
	ClientCertificate *string `json:"clientCertificate,omitempty" tf:"client_certificate,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Unique KMIP identifier for an existing key within the KMIP server.
	KeyIdentifier *string `json:"keyIdentifier,omitempty" tf:"key_identifier,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Port number to use to communicate with the KMIP server. Default: 5696
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	// (Required) Path to CA File. Used for validating secure client connection to KMIP server.
	ServerCA *string `json:"serverCa,omitempty" tf:"server_ca,omitempty"`

	// +kubebuilder:validation:Optional
	// (Required) Hostname or IP address of the KMIP server to connect to.
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`
}

type MongodObservation struct {
}

type MongodParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) A set of audit log settings 
	AuditLog []AuditLogParameters `json:"auditLog,omitempty" tf:"audit_log,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) A set of MongoDB Security settings
	Security []SecurityParameters `json:"security,omitempty" tf:"security,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) A set of MongoDB Server Parameters 
	SetParameter []SetParameterParameters `json:"setParameter,omitempty" tf:"set_parameter,omitempty"`
}

type MongodbClusterHostObservation struct {
	// (Computed) The health of the host.
	Health *string `json:"health,omitempty" tf:"health,omitempty"`

	// (Required) Name of the MongoDB cluster. Provided by the client when the cluster is created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type MongodbClusterHostParameters struct {

	// +kubebuilder:validation:Optional
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) The role of the cluster (either PRIMARY or SECONDARY).
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) The name of the shard to which the host belongs.
	ShardName *string `json:"shardName,omitempty" tf:"shard_name,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	// (Required) The ID of the subnet, to which the host belongs. The subnet must
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	// (Optional) type of mongo daemon which runs on this host (mongod, mongos or monogcfg). Defaults to mongod.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) The availability zone where the MongoDB host will be created.
	ZoneID *string `json:"zoneId" tf:"zone_id,omitempty"`
}

type MongodbClusterMaintenanceWindowObservation struct {
}

type MongodbClusterMaintenanceWindowParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Day of week for maintenance window if window type is weekly. Possible values: `MON`, `TUE`, `WED`, `THU`, `FRI`, `SAT`, `SUN`.
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Hour of day in UTC time zone (1-24) for maintenance window if window type is weekly.
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// +kubebuilder:validation:Required
	// (Optional) type of mongo daemon which runs on this host (mongod, mongos or monogcfg). Defaults to mongod.
	Type *string `json:"type" tf:"type,omitempty"`
}

type MongodbClusterObservation struct {
	// (Required) Configuration of the MongoDB subcluster. The structure is documented below.
	ClusterConfig []ClusterConfigObservation `json:"clusterConfig,omitempty" tf:"cluster_config,omitempty"`

	// Creation timestamp of the key.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// (Computed) The health of the host.
	Health *string `json:"health,omitempty" tf:"health,omitempty"`

	// (Required) A host of the MongoDB cluster. The structure is documented below.
	Host []MongodbClusterHostObservation `json:"host,omitempty" tf:"host,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// MongoDB Cluster mode enabled/disabled.
	Sharded *bool `json:"sharded,omitempty" tf:"sharded,omitempty"`

	// Status of the cluster. Can be either `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or `STATUS_UNKNOWN`.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type MongodbClusterParameters struct {

	// +kubebuilder:validation:Required
	// (Required) Configuration of the MongoDB subcluster. The structure is documented below.
	ClusterConfig []ClusterConfigParameters `json:"clusterConfig" tf:"cluster_config,omitempty"`

	// +kubebuilder:validation:Optional
	// The ID of the cluster.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) A database of the MongoDB cluster. The structure is documented below.
	Database []DatabaseParameters `json:"database" tf:"database,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Description of the MongoDB cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Deployment environment of the MongoDB cluster. Can be either `PRESTABLE` or `PRODUCTION`.
	Environment *string `json:"environment" tf:"environment,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	// (Optional) The ID of the folder that the resource belongs to. If it
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	// (Required) A host of the MongoDB cluster. The structure is documented below.
	Host []MongodbClusterHostParameters `json:"host" tf:"host,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) A set of key/value label pairs to assign to the MongoDB cluster.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	MaintenanceWindow []MongodbClusterMaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Name of the MongoDB cluster. Provided by the client when the cluster is created.
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Network
	// +kubebuilder:validation:Optional
	// (Required) ID of the network, to which the MongoDB cluster belongs.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	// (Required) Resources allocated to hosts of the MongoDB cluster. The structure is documented below.
	Resources []MongodbClusterResourcesParameters `json:"resources" tf:"resources,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.SecurityGroup
	// +kubebuilder:validation:Optional
	// (Optional) A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupIdsRefs []v1.Reference `json:"securityGroupIdsRefs,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SecurityGroupIdsSelector *v1.Selector `json:"securityGroupIdsSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	// (Required) A user of the MongoDB cluster. The structure is documented below.
	User []UserParameters `json:"user" tf:"user,omitempty"`
}

type MongodbClusterResourcesObservation struct {
}

type MongodbClusterResourcesParameters struct {

	// +kubebuilder:validation:Required
	// (Required) Volume of the storage available to a MongoDB host, in gigabytes.
	DiskSize *float64 `json:"diskSize" tf:"disk_size,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Type of the storage of MongoDB hosts.
	DiskTypeID *string `json:"diskTypeId" tf:"disk_type_id,omitempty"`

	// +kubebuilder:validation:Required
	ResourcePresetID *string `json:"resourcePresetId" tf:"resource_preset_id,omitempty"`
}

type PermissionObservation struct {
}

type PermissionParameters struct {

	// +kubebuilder:validation:Required
	// (Required) The name of the database that the permission grants access to.
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) The roles of the user in this database. For more information see [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/concepts/users-and-roles).
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

type SecurityObservation struct {
}

type SecurityParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Enables the encryption for the WiredTiger storage engine. Can be either true or false.
	EnableEncryption *bool `json:"enableEncryption,omitempty" tf:"enable_encryption,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Configuration of the third party key management appliance via the Key Management Interoperability Protocol (KMIP)
	Kmip []KmipParameters `json:"kmip,omitempty" tf:"kmip,omitempty"`
}

type SetParameterObservation struct {
}

type SetParameterParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Enables the auditing of authorization successes. Can be either true or false.
	AuditAuthorizationSuccess *bool `json:"auditAuthorizationSuccess,omitempty" tf:"audit_authorization_success,omitempty"`
}

type UserObservation struct {
}

type UserParameters struct {

	// +kubebuilder:validation:Required
	// (Required) Name of the MongoDB cluster. Provided by the client when the cluster is created.
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	// (Optional) Set of permissions granted to the user. The structure is documented below.
	Permission []PermissionParameters `json:"permission,omitempty" tf:"permission,omitempty"`
}

// MongodbClusterSpec defines the desired state of MongodbCluster
type MongodbClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MongodbClusterParameters `json:"forProvider"`
}

// MongodbClusterStatus defines the observed state of MongodbCluster.
type MongodbClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MongodbClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MongodbCluster is the Schema for the MongodbClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type MongodbCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MongodbClusterSpec   `json:"spec"`
	Status            MongodbClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MongodbClusterList contains a list of MongodbClusters
type MongodbClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbCluster `json:"items"`
}

// Repository type metadata.
var (
	MongodbCluster_Kind             = "MongodbCluster"
	MongodbCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MongodbCluster_Kind}.String()
	MongodbCluster_KindAPIVersion   = MongodbCluster_Kind + "." + CRDGroupVersion.String()
	MongodbCluster_GroupVersionKind = CRDGroupVersion.WithKind(MongodbCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&MongodbCluster{}, &MongodbClusterList{})
}
