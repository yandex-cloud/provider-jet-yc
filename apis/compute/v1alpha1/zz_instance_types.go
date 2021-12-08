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

type BootDiskObservation struct {
}

type BootDiskParameters struct {

	// +kubebuilder:validation:Optional
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// +kubebuilder:validation:Optional
	DiskID *string `json:"diskId,omitempty" tf:"disk_id,omitempty"`

	// +kubebuilder:validation:Optional
	InitializeParams []InitializeParamsParameters `json:"initializeParams,omitempty" tf:"initialize_params,omitempty"`

	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type DNSRecordObservation struct {
}

type DNSRecordParameters struct {

	// +kubebuilder:validation:Optional
	DNSZoneID *string `json:"dnsZoneId,omitempty" tf:"dns_zone_id,omitempty"`

	// +kubebuilder:validation:Required
	Fqdn *string `json:"fqdn" tf:"fqdn,omitempty"`

	// +kubebuilder:validation:Optional
	Ptr *bool `json:"ptr,omitempty" tf:"ptr,omitempty"`

	// +kubebuilder:validation:Optional
	TTL *int64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type IPv6DNSRecordObservation struct {
}

type IPv6DNSRecordParameters struct {

	// +kubebuilder:validation:Optional
	DNSZoneID *string `json:"dnsZoneId,omitempty" tf:"dns_zone_id,omitempty"`

	// +kubebuilder:validation:Required
	Fqdn *string `json:"fqdn" tf:"fqdn,omitempty"`

	// +kubebuilder:validation:Optional
	Ptr *bool `json:"ptr,omitempty" tf:"ptr,omitempty"`

	// +kubebuilder:validation:Optional
	TTL *int64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type InitializeParamsObservation struct {
}

type InitializeParamsParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Size *int64 `json:"size,omitempty" tf:"size,omitempty"`

	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type InstanceObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type InstanceParameters struct {

	// +kubebuilder:validation:Optional
	AllowStoppingForUpdate *bool `json:"allowStoppingForUpdate,omitempty" tf:"allow_stopping_for_update,omitempty"`

	// +kubebuilder:validation:Required
	BootDisk []BootDiskParameters `json:"bootDisk" tf:"boot_disk,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +crossplane:generate:reference:type=bb.yandex-team.ru/crossplane/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkAccelerationType *string `json:"networkAccelerationType,omitempty" tf:"network_acceleration_type,omitempty"`

	// +kubebuilder:validation:Required
	NetworkInterface []NetworkInterfaceParameters `json:"networkInterface" tf:"network_interface,omitempty"`

	// +kubebuilder:validation:Optional
	PlacementPolicy []PlacementPolicyParameters `json:"placementPolicy,omitempty" tf:"placement_policy,omitempty"`

	// +kubebuilder:validation:Optional
	PlatformID *string `json:"platformId,omitempty" tf:"platform_id,omitempty"`

	// +kubebuilder:validation:Required
	Resources []ResourcesParameters `json:"resources" tf:"resources,omitempty"`

	// +kubebuilder:validation:Optional
	SchedulingPolicy []SchedulingPolicyParameters `json:"schedulingPolicy,omitempty" tf:"scheduling_policy,omitempty"`

	// +kubebuilder:validation:Optional
	SecondaryDisk []SecondaryDiskParameters `json:"secondaryDisk,omitempty" tf:"secondary_disk,omitempty"`

	// +crossplane:generate:reference:type=bb.yandex-team.ru/crossplane/provider-jet-yc/apis/iam/v1alpha1.ServiceAccount
	// +kubebuilder:validation:Optional
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceAccountIDRef *v1.Reference `json:"serviceAccountIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceAccountIDSelector *v1.Selector `json:"serviceAccountIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type NatDNSRecordObservation struct {
}

type NatDNSRecordParameters struct {

	// +kubebuilder:validation:Optional
	DNSZoneID *string `json:"dnsZoneId,omitempty" tf:"dns_zone_id,omitempty"`

	// +kubebuilder:validation:Required
	Fqdn *string `json:"fqdn" tf:"fqdn,omitempty"`

	// +kubebuilder:validation:Optional
	Ptr *bool `json:"ptr,omitempty" tf:"ptr,omitempty"`

	// +kubebuilder:validation:Optional
	TTL *int64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type NetworkInterfaceObservation struct {
	Index *int64 `json:"index,omitempty" tf:"index,omitempty"`

	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	NatIPVersion *string `json:"natIpVersion,omitempty" tf:"nat_ip_version,omitempty"`
}

type NetworkInterfaceParameters struct {

	// +kubebuilder:validation:Optional
	DNSRecord []DNSRecordParameters `json:"dnsRecord,omitempty" tf:"dns_record,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	IPv4 *bool `json:"ipv4,omitempty" tf:"ipv4,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6 *bool `json:"ipv6,omitempty" tf:"ipv6,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6Address *string `json:"ipv6Address,omitempty" tf:"ipv6_address,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6DNSRecord []IPv6DNSRecordParameters `json:"ipv6DnsRecord,omitempty" tf:"ipv6_dns_record,omitempty"`

	// +kubebuilder:validation:Optional
	Nat *bool `json:"nat,omitempty" tf:"nat,omitempty"`

	// +kubebuilder:validation:Optional
	NatDNSRecord []NatDNSRecordParameters `json:"natDnsRecord,omitempty" tf:"nat_dns_record,omitempty"`

	// +kubebuilder:validation:Optional
	NatIPAddress *string `json:"natIpAddress,omitempty" tf:"nat_ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// +crossplane:generate:reference:type=bb.yandex-team.ru/crossplane/provider-jet-yc/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type PlacementPolicyObservation struct {
}

type PlacementPolicyParameters struct {

	// +kubebuilder:validation:Required
	PlacementGroupID *string `json:"placementGroupId" tf:"placement_group_id,omitempty"`
}

type ResourcesObservation struct {
}

type ResourcesParameters struct {

	// +kubebuilder:validation:Optional
	CoreFraction *int64 `json:"coreFraction,omitempty" tf:"core_fraction,omitempty"`

	// +kubebuilder:validation:Required
	Cores *int64 `json:"cores" tf:"cores,omitempty"`

	// +kubebuilder:validation:Optional
	Gpus *int64 `json:"gpus,omitempty" tf:"gpus,omitempty"`

	// +kubebuilder:validation:Required
	Memory *float64 `json:"memory" tf:"memory,omitempty"`
}

type SchedulingPolicyObservation struct {
}

type SchedulingPolicyParameters struct {

	// +kubebuilder:validation:Optional
	Preemptible *bool `json:"preemptible,omitempty" tf:"preemptible,omitempty"`
}

type SecondaryDiskObservation struct {
}

type SecondaryDiskParameters struct {

	// +kubebuilder:validation:Optional
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// +kubebuilder:validation:Required
	DiskID *string `json:"diskId" tf:"disk_id,omitempty"`

	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}