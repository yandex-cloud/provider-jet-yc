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

type AccessObservation struct {

	// Allow access for Yandex DataLens.
	DataLens *bool `json:"dataLens,omitempty" tf:"data_lens,omitempty"`
}

type AccessParameters struct {
}

type AuditLogObservation struct {

	// Configuration of the audit log filter in JSON format.
	// For more information see auditLog.filter
	// description in the official documentation. Available only in enterprise edition.
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// Specifies if a node allows runtime configuration of audit filters and the auditAuthorizationSuccess variable.
	// For more information see auditLog.runtimeConfiguration
	// description in the official documentation. Available only in enterprise edition.
	RuntimeConfiguration *bool `json:"runtimeConfiguration,omitempty" tf:"runtime_configuration,omitempty"`
}

type AuditLogParameters struct {

	// Configuration of the audit log filter in JSON format.
	// For more information see auditLog.filter
	// description in the official documentation. Available only in enterprise edition.
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// Specifies if a node allows runtime configuration of audit filters and the auditAuthorizationSuccess variable.
	// For more information see auditLog.runtimeConfiguration
	// description in the official documentation. Available only in enterprise edition.
	// +kubebuilder:validation:Optional
	RuntimeConfiguration *bool `json:"runtimeConfiguration,omitempty" tf:"runtime_configuration,omitempty"`
}

type BackupWindowStartObservation struct {

	// The hour at which backup will be started.
	Hours *float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// The minute at which backup will be started.
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`
}

type BackupWindowStartParameters struct {

	// The hour at which backup will be started.
	// +kubebuilder:validation:Optional
	Hours *float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// The minute at which backup will be started.
	// +kubebuilder:validation:Optional
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`
}

type ClusterConfigObservation struct {

	// Access policy to the MongoDB cluster. The structure is documented below.
	Access []AccessObservation `json:"access,omitempty" tf:"access,omitempty"`

	// Time to start the daily backup, in the UTC timezone. The structure is documented below.
	BackupWindowStart []BackupWindowStartObservation `json:"backupWindowStart,omitempty" tf:"backup_window_start,omitempty"`

	// Feature compatibility version of MongoDB. If not provided version is taken. Can be either 6.0, 5.0, 4.4 and 4.2.
	FeatureCompatibilityVersion *string `json:"featureCompatibilityVersion,omitempty" tf:"feature_compatibility_version,omitempty"`

	// Configuration of the mongod service. The structure is documented below.
	Mongod []MongodObservation `json:"mongod,omitempty" tf:"mongod,omitempty"`

	// Version of the MongoDB server software. Can be either 4.2, 4.4, 4.4-enterprise, 5.0, 5.0-enterprise, 6.0 and 6.0-enterprise.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ClusterConfigParameters struct {

	// Access policy to the MongoDB cluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	Access []AccessParameters `json:"access,omitempty" tf:"access,omitempty"`

	// Time to start the daily backup, in the UTC timezone. The structure is documented below.
	// +kubebuilder:validation:Optional
	BackupWindowStart []BackupWindowStartParameters `json:"backupWindowStart,omitempty" tf:"backup_window_start,omitempty"`

	// Feature compatibility version of MongoDB. If not provided version is taken. Can be either 6.0, 5.0, 4.4 and 4.2.
	// +kubebuilder:validation:Optional
	FeatureCompatibilityVersion *string `json:"featureCompatibilityVersion,omitempty" tf:"feature_compatibility_version,omitempty"`

	// Configuration of the mongod service. The structure is documented below.
	// +kubebuilder:validation:Optional
	Mongod []MongodParameters `json:"mongod,omitempty" tf:"mongod,omitempty"`

	// Version of the MongoDB server software. Can be either 4.2, 4.4, 4.4-enterprise, 5.0, 5.0-enterprise, 6.0 and 6.0-enterprise.
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type DatabaseObservation struct {

	// The name of the database.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type DatabaseParameters struct {

	// The name of the database.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type KmipObservation struct {

	// String containing the client certificate used for authenticating MongoDB to the KMIP server.
	// For more information see security.kmip.clientCertificateFile
	// description in the official documentation.
	ClientCertificate *string `json:"clientCertificate,omitempty" tf:"client_certificate,omitempty"`

	// Unique KMIP identifier for an existing key within the KMIP server.
	// For more information see security.kmip.keyIdentifier
	// description in the official documentation.
	KeyIdentifier *string `json:"keyIdentifier,omitempty" tf:"key_identifier,omitempty"`

	// Port number to use to communicate with the KMIP server. Default: 5696
	// For more information see security.kmip.port
	// description in the official documentation.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Path to CA File. Used for validating secure client connection to KMIP server.
	// For more information see security.kmip.serverCAFile
	// description in the official documentation.
	ServerCA *string `json:"serverCa,omitempty" tf:"server_ca,omitempty"`

	// Hostname or IP address of the KMIP server to connect to.
	// For more information see security.kmip.serverName
	// description in the official documentation.
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`
}

type KmipParameters struct {

	// String containing the client certificate used for authenticating MongoDB to the KMIP server.
	// For more information see security.kmip.clientCertificateFile
	// description in the official documentation.
	// +kubebuilder:validation:Optional
	ClientCertificate *string `json:"clientCertificate,omitempty" tf:"client_certificate,omitempty"`

	// Unique KMIP identifier for an existing key within the KMIP server.
	// For more information see security.kmip.keyIdentifier
	// description in the official documentation.
	// +kubebuilder:validation:Optional
	KeyIdentifier *string `json:"keyIdentifier,omitempty" tf:"key_identifier,omitempty"`

	// Port number to use to communicate with the KMIP server. Default: 5696
	// For more information see security.kmip.port
	// description in the official documentation.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Path to CA File. Used for validating secure client connection to KMIP server.
	// For more information see security.kmip.serverCAFile
	// description in the official documentation.
	// +kubebuilder:validation:Optional
	ServerCA *string `json:"serverCa,omitempty" tf:"server_ca,omitempty"`

	// Hostname or IP address of the KMIP server to connect to.
	// For more information see security.kmip.serverName
	// description in the official documentation.
	// +kubebuilder:validation:Optional
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`
}

type MongodObservation struct {

	// A set of audit log settings
	// (see the auditLog option).
	// The structure is documented below. Available only in enterprise edition.
	AuditLog []AuditLogObservation `json:"auditLog,omitempty" tf:"audit_log,omitempty"`

	// A set of MongoDB Security settings
	// (see the security option).
	// The structure is documented below. Available only in enterprise edition.
	Security []SecurityObservation `json:"security,omitempty" tf:"security,omitempty"`

	// A set of MongoDB Server Parameters
	// (see the setParameter option).
	// The structure is documented below.
	SetParameter []SetParameterObservation `json:"setParameter,omitempty" tf:"set_parameter,omitempty"`
}

type MongodParameters struct {

	// A set of audit log settings
	// (see the auditLog option).
	// The structure is documented below. Available only in enterprise edition.
	// +kubebuilder:validation:Optional
	AuditLog []AuditLogParameters `json:"auditLog,omitempty" tf:"audit_log,omitempty"`

	// A set of MongoDB Security settings
	// (see the security option).
	// The structure is documented below. Available only in enterprise edition.
	// +kubebuilder:validation:Optional
	Security []SecurityParameters `json:"security,omitempty" tf:"security,omitempty"`

	// A set of MongoDB Server Parameters
	// (see the setParameter option).
	// The structure is documented below.
	// +kubebuilder:validation:Optional
	SetParameter []SetParameterParameters `json:"setParameter,omitempty" tf:"set_parameter,omitempty"`
}

type MongodbClusterHostObservation struct {

	// Should this host have assigned public IP assigned. Can be either true or false.
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// (Computed) The health of the host.
	Health *string `json:"health,omitempty" tf:"health,omitempty"`

	// (Computed) The fully qualified domain name of the host. Computed on server side.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The role of the cluster (either PRIMARY or SECONDARY).
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The name of the shard to which the host belongs. Only for sharded cluster.
	ShardName *string `json:"shardName,omitempty" tf:"shard_name,omitempty"`

	// The ID of the subnet, to which the host belongs. The subnet must
	// be a part of the network to which the cluster belongs.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// type of mongo daemon which runs on this host (mongod, mongos, mongocfg, mongoinfra). Defaults to mongod.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The availability zone where the MongoDB host will be created.
	// For more information see the official documentation.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type MongodbClusterHostParameters struct {

	// Should this host have assigned public IP assigned. Can be either true or false.
	// +kubebuilder:validation:Optional
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// The role of the cluster (either PRIMARY or SECONDARY).
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The name of the shard to which the host belongs. Only for sharded cluster.
	// +kubebuilder:validation:Optional
	ShardName *string `json:"shardName,omitempty" tf:"shard_name,omitempty"`

	// The ID of the subnet, to which the host belongs. The subnet must
	// be a part of the network to which the cluster belongs.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// type of mongo daemon which runs on this host (mongod, mongos, mongocfg, mongoinfra). Defaults to mongod.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The availability zone where the MongoDB host will be created.
	// For more information see the official documentation.
	// +kubebuilder:validation:Required
	ZoneID *string `json:"zoneId" tf:"zone_id,omitempty"`
}

type MongodbClusterMaintenanceWindowObservation struct {

	// Day of week for maintenance window if window type is weekly. Possible values: MON, TUE, WED, THU, FRI, SAT, SUN.
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// Hour of day in UTC time zone (1-24) for maintenance window if window type is weekly.
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// Type of maintenance window. Can be either ANYTIME or WEEKLY. A day and hour of window need to be specified with weekly window.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type MongodbClusterMaintenanceWindowParameters struct {

	// Day of week for maintenance window if window type is weekly. Possible values: MON, TUE, WED, THU, FRI, SAT, SUN.
	// +kubebuilder:validation:Optional
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// Hour of day in UTC time zone (1-24) for maintenance window if window type is weekly.
	// +kubebuilder:validation:Optional
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// Type of maintenance window. Can be either ANYTIME or WEEKLY. A day and hour of window need to be specified with weekly window.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type MongodbClusterObservation struct {

	// Configuration of the MongoDB subcluster. The structure is documented below.
	ClusterConfig []ClusterConfigObservation `json:"clusterConfig,omitempty" tf:"cluster_config,omitempty"`

	// The ID of the cluster.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Creation timestamp of the key.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// A database of the MongoDB cluster. The structure is documented below.
	Database []DatabaseObservation `json:"database,omitempty" tf:"database,omitempty"`

	// Inhibits deletion of the cluster.  Can be either true or false.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Description of the MongoDB cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Deployment environment of the MongoDB cluster. Can be either PRESTABLE or PRODUCTION.
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Aggregated health of the cluster. Can be either ALIVE, DEGRADED, DEAD or HEALTH_UNKNOWN.
	// For more information see health field of JSON representation in the official documentation.
	Health *string `json:"health,omitempty" tf:"health,omitempty"`

	// A host of the MongoDB cluster. The structure is documented below.
	Host []MongodbClusterHostObservation `json:"host,omitempty" tf:"host,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to the MongoDB cluster.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Maintenance window settings of the MongoDB cluster. The structure is documented below.
	MaintenanceWindow []MongodbClusterMaintenanceWindowObservation `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// Name of the MongoDB cluster. Provided by the client when the cluster is created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network, to which the MongoDB cluster belongs.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// (DEPRECATED, use resources_* instead) Resources allocated to hosts of the MongoDB cluster. The structure is documented below.
	Resources []MongodbClusterResourcesObservation `json:"resources,omitempty" tf:"resources,omitempty"`

	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// MongoDB Cluster mode enabled/disabled.
	Sharded *bool `json:"sharded,omitempty" tf:"sharded,omitempty"`

	// Status of the cluster. Can be either CREATING, STARTING, RUNNING, UPDATING, STOPPING, STOPPED, ERROR or STATUS_UNKNOWN.
	// For more information see status field of JSON representation in the official documentation.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A user of the MongoDB cluster. The structure is documented below.
	User []UserObservation `json:"user,omitempty" tf:"user,omitempty"`
}

type MongodbClusterParameters struct {

	// Configuration of the MongoDB subcluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	ClusterConfig []ClusterConfigParameters `json:"clusterConfig,omitempty" tf:"cluster_config,omitempty"`

	// The ID of the cluster.
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// A database of the MongoDB cluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	Database []DatabaseParameters `json:"database,omitempty" tf:"database,omitempty"`

	// Inhibits deletion of the cluster.  Can be either true or false.
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Description of the MongoDB cluster.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Deployment environment of the MongoDB cluster. Can be either PRESTABLE or PRODUCTION.
	// +kubebuilder:validation:Optional
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A host of the MongoDB cluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	Host []MongodbClusterHostParameters `json:"host,omitempty" tf:"host,omitempty"`

	// A set of key/value label pairs to assign to the MongoDB cluster.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Maintenance window settings of the MongoDB cluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	MaintenanceWindow []MongodbClusterMaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// Name of the MongoDB cluster. Provided by the client when the cluster is created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network, to which the MongoDB cluster belongs.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// (DEPRECATED, use resources_* instead) Resources allocated to hosts of the MongoDB cluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	Resources []MongodbClusterResourcesParameters `json:"resources,omitempty" tf:"resources,omitempty"`

	// A set of ids of security groups assigned to hosts of the cluster.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to SecurityGroup in vpc to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsRefs []v1.Reference `json:"securityGroupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in vpc to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsSelector *v1.Selector `json:"securityGroupIdsSelector,omitempty" tf:"-"`

	// A user of the MongoDB cluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	User []UserParameters `json:"user,omitempty" tf:"user,omitempty"`
}

type MongodbClusterResourcesObservation struct {

	// Volume of the storage available to a MongoDB host, in gigabytes.
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// Type of the storage of MongoDB hosts.
	// For more information see the official documentation.
	DiskTypeID *string `json:"diskTypeId,omitempty" tf:"disk_type_id,omitempty"`

	ResourcePresetID *string `json:"resourcePresetId,omitempty" tf:"resource_preset_id,omitempty"`
}

type MongodbClusterResourcesParameters struct {

	// Volume of the storage available to a MongoDB host, in gigabytes.
	// +kubebuilder:validation:Required
	DiskSize *float64 `json:"diskSize" tf:"disk_size,omitempty"`

	// Type of the storage of MongoDB hosts.
	// For more information see the official documentation.
	// +kubebuilder:validation:Required
	DiskTypeID *string `json:"diskTypeId" tf:"disk_type_id,omitempty"`

	// +kubebuilder:validation:Required
	ResourcePresetID *string `json:"resourcePresetId" tf:"resource_preset_id,omitempty"`
}

type PermissionObservation struct {

	// The name of the database that the permission grants access to.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// The roles of the user in this database. For more information see the official documentation.
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

type PermissionParameters struct {

	// The name of the database that the permission grants access to.
	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// The roles of the user in this database. For more information see the official documentation.
	// +kubebuilder:validation:Optional
	Roles []*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

type SecurityObservation struct {

	// Enables the encryption for the WiredTiger storage engine. Can be either true or false.
	// For more information see security.enableEncryption
	// description in the official documentation. Available only in enterprise edition.
	EnableEncryption *bool `json:"enableEncryption,omitempty" tf:"enable_encryption,omitempty"`

	// Configuration of the third party key management appliance via the Key Management Interoperability Protocol (KMIP)
	// (see Encryption tutorial ). Requires enable_encryption to be true.
	// The structure is documented below. Available only in enterprise edition.
	Kmip []KmipObservation `json:"kmip,omitempty" tf:"kmip,omitempty"`
}

type SecurityParameters struct {

	// Enables the encryption for the WiredTiger storage engine. Can be either true or false.
	// For more information see security.enableEncryption
	// description in the official documentation. Available only in enterprise edition.
	// +kubebuilder:validation:Optional
	EnableEncryption *bool `json:"enableEncryption,omitempty" tf:"enable_encryption,omitempty"`

	// Configuration of the third party key management appliance via the Key Management Interoperability Protocol (KMIP)
	// (see Encryption tutorial ). Requires enable_encryption to be true.
	// The structure is documented below. Available only in enterprise edition.
	// +kubebuilder:validation:Optional
	Kmip []KmipParameters `json:"kmip,omitempty" tf:"kmip,omitempty"`
}

type SetParameterObservation struct {

	// Enables the auditing of authorization successes. Can be either true or false.
	// For more information, see the auditAuthorizationSuccess
	// description in the official documentation. Available only in enterprise edition.
	AuditAuthorizationSuccess *bool `json:"auditAuthorizationSuccess,omitempty" tf:"audit_authorization_success,omitempty"`
}

type SetParameterParameters struct {

	// Enables the auditing of authorization successes. Can be either true or false.
	// For more information, see the auditAuthorizationSuccess
	// description in the official documentation. Available only in enterprise edition.
	// +kubebuilder:validation:Optional
	AuditAuthorizationSuccess *bool `json:"auditAuthorizationSuccess,omitempty" tf:"audit_authorization_success,omitempty"`
}

type UserObservation struct {

	// The name of the user.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Set of permissions granted to the user. The structure is documented below.
	Permission []PermissionObservation `json:"permission,omitempty" tf:"permission,omitempty"`
}

type UserParameters struct {

	// The name of the user.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The password of the user.
	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// Set of permissions granted to the user. The structure is documented below.
	// +kubebuilder:validation:Optional
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

// MongodbCluster is the Schema for the MongodbClusters API. Manages a MongoDB cluster within Yandex.Cloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type MongodbCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.clusterConfig)",message="clusterConfig is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.database)",message="database is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.environment)",message="environment is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.host)",message="host is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.resources)",message="resources is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.user)",message="user is a required parameter"
	Spec   MongodbClusterSpec   `json:"spec"`
	Status MongodbClusterStatus `json:"status,omitempty"`
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
