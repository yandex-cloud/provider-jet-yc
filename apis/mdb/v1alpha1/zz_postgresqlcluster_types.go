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

type ConfigAccessObservation struct {

	// Allow access for Yandex DataLens.
	DataLens *bool `json:"dataLens,omitempty" tf:"data_lens,omitempty"`

	// Allow access for connection to managed databases from functions
	Serverless *bool `json:"serverless,omitempty" tf:"serverless,omitempty"`

	// Allow access for SQL queries in the management console
	WebSQL *bool `json:"webSql,omitempty" tf:"web_sql,omitempty"`
}

type ConfigAccessParameters struct {

	// Allow access for Yandex DataLens.
	// +kubebuilder:validation:Optional
	DataLens *bool `json:"dataLens,omitempty" tf:"data_lens,omitempty"`

	// Allow access for connection to managed databases from functions
	// +kubebuilder:validation:Optional
	Serverless *bool `json:"serverless,omitempty" tf:"serverless,omitempty"`

	// Allow access for SQL queries in the management console
	// +kubebuilder:validation:Optional
	WebSQL *bool `json:"webSql,omitempty" tf:"web_sql,omitempty"`
}

type ConfigBackupWindowStartObservation struct {

	// The hour at which backup will be started (UTC).
	Hours *float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// The minute at which backup will be started (UTC).
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`
}

type ConfigBackupWindowStartParameters struct {

	// The hour at which backup will be started (UTC).
	// +kubebuilder:validation:Optional
	Hours *float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// The minute at which backup will be started (UTC).
	// +kubebuilder:validation:Optional
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`
}

type ConfigPerformanceDiagnosticsObservation struct {

	// Enable performance diagnostics
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Interval (in seconds) for pg_stat_activity sampling Acceptable values are 1 to 86400, inclusive.
	SessionsSamplingInterval *float64 `json:"sessionsSamplingInterval,omitempty" tf:"sessions_sampling_interval,omitempty"`

	// Interval (in seconds) for pg_stat_statements sampling Acceptable values are 1 to 86400, inclusive.
	StatementsSamplingInterval *float64 `json:"statementsSamplingInterval,omitempty" tf:"statements_sampling_interval,omitempty"`
}

type ConfigPerformanceDiagnosticsParameters struct {

	// Enable performance diagnostics
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Interval (in seconds) for pg_stat_activity sampling Acceptable values are 1 to 86400, inclusive.
	// +kubebuilder:validation:Required
	SessionsSamplingInterval *float64 `json:"sessionsSamplingInterval" tf:"sessions_sampling_interval,omitempty"`

	// Interval (in seconds) for pg_stat_statements sampling Acceptable values are 1 to 86400, inclusive.
	// +kubebuilder:validation:Required
	StatementsSamplingInterval *float64 `json:"statementsSamplingInterval" tf:"statements_sampling_interval,omitempty"`
}

type ConfigResourcesObservation struct {

	// Volume of the storage available to a PostgreSQL host, in gigabytes.
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// Type of the storage of PostgreSQL hosts.
	DiskTypeID *string `json:"diskTypeId,omitempty" tf:"disk_type_id,omitempty"`

	ResourcePresetID *string `json:"resourcePresetId,omitempty" tf:"resource_preset_id,omitempty"`
}

type ConfigResourcesParameters struct {

	// Volume of the storage available to a PostgreSQL host, in gigabytes.
	// +kubebuilder:validation:Required
	DiskSize *float64 `json:"diskSize" tf:"disk_size,omitempty"`

	// Type of the storage of PostgreSQL hosts.
	// +kubebuilder:validation:Optional
	DiskTypeID *string `json:"diskTypeId,omitempty" tf:"disk_type_id,omitempty"`

	// +kubebuilder:validation:Required
	ResourcePresetID *string `json:"resourcePresetId" tf:"resource_preset_id,omitempty"`
}

type ExtensionObservation struct {

	// Name of the PostgreSQL cluster. Provided by the client when the cluster is created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Version of the PostgreSQL cluster. (allowed versions are: 10, 10-1c, 11, 11-1c, 12, 12-1c, 13, 14)
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ExtensionParameters struct {

	// Name of the PostgreSQL cluster. Provided by the client when the cluster is created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Version of the PostgreSQL cluster. (allowed versions are: 10, 10-1c, 11, 11-1c, 12, 12-1c, 13, 14)
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type PoolerConfigObservation struct {

	// Setting pool_discard parameter in Odyssey.
	PoolDiscard *bool `json:"poolDiscard,omitempty" tf:"pool_discard,omitempty"`

	// Mode that the connection pooler is working in. See descriptions of all modes in the [documentation for Odyssey](https://github.com/yandex/odyssey/blob/master/documentation/configuration.md#pool-string.
	PoolingMode *string `json:"poolingMode,omitempty" tf:"pooling_mode,omitempty"`
}

type PoolerConfigParameters struct {

	// Setting pool_discard parameter in Odyssey.
	// +kubebuilder:validation:Optional
	PoolDiscard *bool `json:"poolDiscard,omitempty" tf:"pool_discard,omitempty"`

	// Mode that the connection pooler is working in. See descriptions of all modes in the [documentation for Odyssey](https://github.com/yandex/odyssey/blob/master/documentation/configuration.md#pool-string.
	// +kubebuilder:validation:Optional
	PoolingMode *string `json:"poolingMode,omitempty" tf:"pooling_mode,omitempty"`
}

type PostgresqlClusterConfigObservation struct {

	// Access policy to the PostgreSQL cluster. The structure is documented below.
	Access []ConfigAccessObservation `json:"access,omitempty" tf:"access,omitempty"`

	// Configuration setting which enables/disables autofailover in cluster.
	Autofailover *bool `json:"autofailover,omitempty" tf:"autofailover,omitempty"`

	// The period in days during which backups are stored.
	BackupRetainPeriodDays *float64 `json:"backupRetainPeriodDays,omitempty" tf:"backup_retain_period_days,omitempty"`

	// Time to start the daily backup, in the UTC timezone. The structure is documented below.
	BackupWindowStart []ConfigBackupWindowStartObservation `json:"backupWindowStart,omitempty" tf:"backup_window_start,omitempty"`

	// Cluster performance diagnostics settings. The structure is documented below. YC Documentation
	PerformanceDiagnostics []ConfigPerformanceDiagnosticsObservation `json:"performanceDiagnostics,omitempty" tf:"performance_diagnostics,omitempty"`

	// Configuration of the connection pooler. The structure is documented below.
	PoolerConfig []PoolerConfigObservation `json:"poolerConfig,omitempty" tf:"pooler_config,omitempty"`

	// PostgreSQL cluster config. Detail info in "postresql config" section (documented below).
	PostgresqlConfig map[string]*string `json:"postgresqlConfig,omitempty" tf:"postgresql_config,omitempty"`

	// Resources allocated to hosts of the PostgreSQL cluster. The structure is documented below.
	Resources []ConfigResourcesObservation `json:"resources,omitempty" tf:"resources,omitempty"`

	// Version of the PostgreSQL cluster. (allowed versions are: 10, 10-1c, 11, 11-1c, 12, 12-1c, 13, 14)
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type PostgresqlClusterConfigParameters struct {

	// Access policy to the PostgreSQL cluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	Access []ConfigAccessParameters `json:"access,omitempty" tf:"access,omitempty"`

	// Configuration setting which enables/disables autofailover in cluster.
	// +kubebuilder:validation:Optional
	Autofailover *bool `json:"autofailover,omitempty" tf:"autofailover,omitempty"`

	// The period in days during which backups are stored.
	// +kubebuilder:validation:Optional
	BackupRetainPeriodDays *float64 `json:"backupRetainPeriodDays,omitempty" tf:"backup_retain_period_days,omitempty"`

	// Time to start the daily backup, in the UTC timezone. The structure is documented below.
	// +kubebuilder:validation:Optional
	BackupWindowStart []ConfigBackupWindowStartParameters `json:"backupWindowStart,omitempty" tf:"backup_window_start,omitempty"`

	// Cluster performance diagnostics settings. The structure is documented below. YC Documentation
	// +kubebuilder:validation:Optional
	PerformanceDiagnostics []ConfigPerformanceDiagnosticsParameters `json:"performanceDiagnostics,omitempty" tf:"performance_diagnostics,omitempty"`

	// Configuration of the connection pooler. The structure is documented below.
	// +kubebuilder:validation:Optional
	PoolerConfig []PoolerConfigParameters `json:"poolerConfig,omitempty" tf:"pooler_config,omitempty"`

	// PostgreSQL cluster config. Detail info in "postresql config" section (documented below).
	// +kubebuilder:validation:Optional
	PostgresqlConfig map[string]*string `json:"postgresqlConfig,omitempty" tf:"postgresql_config,omitempty"`

	// Resources allocated to hosts of the PostgreSQL cluster. The structure is documented below.
	// +kubebuilder:validation:Required
	Resources []ConfigResourcesParameters `json:"resources" tf:"resources,omitempty"`

	// Version of the PostgreSQL cluster. (allowed versions are: 10, 10-1c, 11, 11-1c, 12, 12-1c, 13, 14)
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type PostgresqlClusterDatabaseObservation struct {
	Extension []ExtensionObservation `json:"extension,omitempty" tf:"extension,omitempty"`

	LcCollate *string `json:"lcCollate,omitempty" tf:"lc_collate,omitempty"`

	LcType *string `json:"lcType,omitempty" tf:"lc_type,omitempty"`

	// Name of the PostgreSQL cluster. Provided by the client when the cluster is created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`
}

type PostgresqlClusterDatabaseParameters struct {

	// +kubebuilder:validation:Optional
	Extension []ExtensionParameters `json:"extension,omitempty" tf:"extension,omitempty"`

	// +kubebuilder:validation:Optional
	LcCollate *string `json:"lcCollate,omitempty" tf:"lc_collate,omitempty"`

	// +kubebuilder:validation:Optional
	LcType *string `json:"lcType,omitempty" tf:"lc_type,omitempty"`

	// Name of the PostgreSQL cluster. Provided by the client when the cluster is created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Owner *string `json:"owner" tf:"owner,omitempty"`
}

type PostgresqlClusterHostObservation struct {

	// Sets whether the host should get a public IP address on creation. It can be changed on the fly only when name is set.
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// (Computed) The fully qualified domain name of the host.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// Host state name. It should be set for all hosts or unset for all hosts. This field can be used by another host, to select which host will be its replication source. Please see replication_source_name parameter.
	// Also, this field is used to select which host will be selected as a master host. Please see host_master_name parameter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Host priority in HA group. It works only when name is set.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// (Computed) Host replication source (fqdn), when replication_source is empty then host is in HA group.
	ReplicationSource *string `json:"replicationSource,omitempty" tf:"replication_source,omitempty"`

	// Host replication source name points to host's name from which this host should replicate. When not set then host in HA group. It works only when name is set.
	ReplicationSourceName *string `json:"replicationSourceName,omitempty" tf:"replication_source_name,omitempty"`

	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// The availability zone where the PostgreSQL host will be created.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type PostgresqlClusterHostParameters struct {

	// Sets whether the host should get a public IP address on creation. It can be changed on the fly only when name is set.
	// +kubebuilder:validation:Optional
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// Host state name. It should be set for all hosts or unset for all hosts. This field can be used by another host, to select which host will be its replication source. Please see replication_source_name parameter.
	// Also, this field is used to select which host will be selected as a master host. Please see host_master_name parameter.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Host priority in HA group. It works only when name is set.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Host replication source name points to host's name from which this host should replicate. When not set then host in HA group. It works only when name is set.
	// +kubebuilder:validation:Optional
	ReplicationSourceName *string `json:"replicationSourceName,omitempty" tf:"replication_source_name,omitempty"`

	// The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// The availability zone where the PostgreSQL host will be created.
	// +kubebuilder:validation:Required
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type PostgresqlClusterMaintenanceWindowObservation struct {

	// Day of the week (in DDD format). Allowed values: "MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// Hour of the day in UTC (in HH format). Allowed value is between 1 and 24.
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// Type of maintenance window. Can be either ANYTIME or WEEKLY. A day and hour of window need to be specified with weekly window.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PostgresqlClusterMaintenanceWindowParameters struct {

	// Day of the week (in DDD format). Allowed values: "MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"
	// +kubebuilder:validation:Optional
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// Hour of the day in UTC (in HH format). Allowed value is between 1 and 24.
	// +kubebuilder:validation:Optional
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// Type of maintenance window. Can be either ANYTIME or WEEKLY. A day and hour of window need to be specified with weekly window.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type PostgresqlClusterObservation struct {

	// Configuration of the PostgreSQL cluster. The structure is documented below.
	Config []PostgresqlClusterConfigObservation `json:"config,omitempty" tf:"config,omitempty"`

	// Timestamp of cluster creation.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// (Deprecated) To manage databases, please switch to using a separate resource type yandex_mdb_postgresql_database.
	Database []PostgresqlClusterDatabaseObservation `json:"database,omitempty" tf:"database,omitempty"`

	// Inhibits deletion of the cluster.  Can be either true or false.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Description of the PostgreSQL cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Deployment environment of the PostgreSQL cluster.
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// The ID of the folder that the resource belongs to. If it is unset, the default provider folder_id is used for create.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Aggregated health of the cluster.
	Health *string `json:"health,omitempty" tf:"health,omitempty"`

	// A host of the PostgreSQL cluster. The structure is documented below.
	Host []PostgresqlClusterHostObservation `json:"host,omitempty" tf:"host,omitempty"`

	HostGroupIds []*string `json:"hostGroupIds,omitempty" tf:"host_group_ids,omitempty"`

	// It sets name of master host. It works only when host.name is set.
	HostMasterName *string `json:"hostMasterName,omitempty" tf:"host_master_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to the PostgreSQL cluster.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Maintenance policy of the PostgreSQL cluster. The structure is documented below.
	MaintenanceWindow []PostgresqlClusterMaintenanceWindowObservation `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// Name of the PostgreSQL cluster. Provided by the client when the cluster is created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network, to which the PostgreSQL cluster belongs.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// The cluster will be created from the specified backup. The structure is documented below.
	Restore []PostgresqlClusterRestoreObservation `json:"restore,omitempty" tf:"restore,omitempty"`

	// A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// Status of the cluster.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (Deprecated) To manage users, please switch to using a separate resource type yandex_mdb_postgresql_user.
	User []PostgresqlClusterUserObservation `json:"user,omitempty" tf:"user,omitempty"`
}

type PostgresqlClusterParameters struct {

	// Configuration of the PostgreSQL cluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	Config []PostgresqlClusterConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// (Deprecated) To manage databases, please switch to using a separate resource type yandex_mdb_postgresql_database.
	// +kubebuilder:validation:Optional
	Database []PostgresqlClusterDatabaseParameters `json:"database,omitempty" tf:"database,omitempty"`

	// Inhibits deletion of the cluster.  Can be either true or false.
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// Description of the PostgreSQL cluster.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Deployment environment of the PostgreSQL cluster.
	// +kubebuilder:validation:Optional
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// The ID of the folder that the resource belongs to. If it is unset, the default provider folder_id is used for create.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A host of the PostgreSQL cluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	Host []PostgresqlClusterHostParameters `json:"host,omitempty" tf:"host,omitempty"`

	// +kubebuilder:validation:Optional
	HostGroupIds []*string `json:"hostGroupIds,omitempty" tf:"host_group_ids,omitempty"`

	// It sets name of master host. It works only when host.name is set.
	// +kubebuilder:validation:Optional
	HostMasterName *string `json:"hostMasterName,omitempty" tf:"host_master_name,omitempty"`

	// A set of key/value label pairs to assign to the PostgreSQL cluster.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Maintenance policy of the PostgreSQL cluster. The structure is documented below.
	// +kubebuilder:validation:Optional
	MaintenanceWindow []PostgresqlClusterMaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// Name of the PostgreSQL cluster. Provided by the client when the cluster is created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network, to which the PostgreSQL cluster belongs.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// The cluster will be created from the specified backup. The structure is documented below.
	// +kubebuilder:validation:Optional
	Restore []PostgresqlClusterRestoreParameters `json:"restore,omitempty" tf:"restore,omitempty"`

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

	// (Deprecated) To manage users, please switch to using a separate resource type yandex_mdb_postgresql_user.
	// +kubebuilder:validation:Optional
	User []PostgresqlClusterUserParameters `json:"user,omitempty" tf:"user,omitempty"`
}

type PostgresqlClusterRestoreObservation struct {

	// Backup ID. The cluster will be created from the specified backup. How to get a list of PostgreSQL backups.
	BackupID *string `json:"backupId,omitempty" tf:"backup_id,omitempty"`

	// Timestamp of the moment to which the PostgreSQL cluster should be restored. (Format: "2006-01-02T15:04:05" - UTC). When not set, current time is used.
	Time *string `json:"time,omitempty" tf:"time,omitempty"`

	// Flag that indicates whether a database should be restored to the first backup point available just after the timestamp specified in the [time] field instead of just before.
	// Possible values:
	TimeInclusive *bool `json:"timeInclusive,omitempty" tf:"time_inclusive,omitempty"`
}

type PostgresqlClusterRestoreParameters struct {

	// Backup ID. The cluster will be created from the specified backup. How to get a list of PostgreSQL backups.
	// +kubebuilder:validation:Required
	BackupID *string `json:"backupId" tf:"backup_id,omitempty"`

	// Timestamp of the moment to which the PostgreSQL cluster should be restored. (Format: "2006-01-02T15:04:05" - UTC). When not set, current time is used.
	// +kubebuilder:validation:Optional
	Time *string `json:"time,omitempty" tf:"time,omitempty"`

	// Flag that indicates whether a database should be restored to the first backup point available just after the timestamp specified in the [time] field instead of just before.
	// Possible values:
	// +kubebuilder:validation:Optional
	TimeInclusive *bool `json:"timeInclusive,omitempty" tf:"time_inclusive,omitempty"`
}

type PostgresqlClusterUserObservation struct {
	ConnLimit *float64 `json:"connLimit,omitempty" tf:"conn_limit,omitempty"`

	Grants []*string `json:"grants,omitempty" tf:"grants,omitempty"`

	Login *bool `json:"login,omitempty" tf:"login,omitempty"`

	// Name of the PostgreSQL cluster. Provided by the client when the cluster is created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Permission []PostgresqlClusterUserPermissionObservation `json:"permission,omitempty" tf:"permission,omitempty"`

	Settings map[string]*string `json:"settings,omitempty" tf:"settings,omitempty"`
}

type PostgresqlClusterUserParameters struct {

	// +kubebuilder:validation:Optional
	ConnLimit *float64 `json:"connLimit,omitempty" tf:"conn_limit,omitempty"`

	// +kubebuilder:validation:Optional
	Grants []*string `json:"grants,omitempty" tf:"grants,omitempty"`

	// +kubebuilder:validation:Optional
	Login *bool `json:"login,omitempty" tf:"login,omitempty"`

	// Name of the PostgreSQL cluster. Provided by the client when the cluster is created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// +kubebuilder:validation:Optional
	Permission []PostgresqlClusterUserPermissionParameters `json:"permission,omitempty" tf:"permission,omitempty"`

	// +kubebuilder:validation:Optional
	Settings map[string]*string `json:"settings,omitempty" tf:"settings,omitempty"`
}

type PostgresqlClusterUserPermissionObservation struct {

	// Name of the PostgreSQL cluster. Provided by the client when the cluster is created.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`
}

type PostgresqlClusterUserPermissionParameters struct {

	// Name of the PostgreSQL cluster. Provided by the client when the cluster is created.
	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`
}

// PostgresqlClusterSpec defines the desired state of PostgresqlCluster
type PostgresqlClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PostgresqlClusterParameters `json:"forProvider"`
}

// PostgresqlClusterStatus defines the observed state of PostgresqlCluster.
type PostgresqlClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PostgresqlClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlCluster is the Schema for the PostgresqlClusters API. Manages a PostgreSQL cluster within Yandex.Cloud.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type PostgresqlCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.config)",message="config is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.environment)",message="environment is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.host)",message="host is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	Spec   PostgresqlClusterSpec   `json:"spec"`
	Status PostgresqlClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlClusterList contains a list of PostgresqlClusters
type PostgresqlClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresqlCluster `json:"items"`
}

// Repository type metadata.
var (
	PostgresqlCluster_Kind             = "PostgresqlCluster"
	PostgresqlCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PostgresqlCluster_Kind}.String()
	PostgresqlCluster_KindAPIVersion   = PostgresqlCluster_Kind + "." + CRDGroupVersion.String()
	PostgresqlCluster_GroupVersionKind = CRDGroupVersion.WithKind(PostgresqlCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&PostgresqlCluster{}, &PostgresqlClusterList{})
}
