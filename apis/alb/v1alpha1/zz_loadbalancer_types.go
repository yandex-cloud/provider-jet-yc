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

type AddressObservation struct {

	// External IPv4 address. The structure is documented below.
	ExternalIPv4Address []ExternalIPv4AddressObservation `json:"externalIpv4Address,omitempty" tf:"external_ipv4_address,omitempty"`

	// External IPv6 address. The structure is documented below.
	ExternalIPv6Address []ExternalIPv6AddressObservation `json:"externalIpv6Address,omitempty" tf:"external_ipv6_address,omitempty"`

	// Internal IPv4 address. The structure is documented below.
	InternalIPv4Address []InternalIPv4AddressObservation `json:"internalIpv4Address,omitempty" tf:"internal_ipv4_address,omitempty"`
}

type AddressParameters struct {

	// External IPv4 address. The structure is documented below.
	// +kubebuilder:validation:Optional
	ExternalIPv4Address []ExternalIPv4AddressParameters `json:"externalIpv4Address,omitempty" tf:"external_ipv4_address,omitempty"`

	// External IPv6 address. The structure is documented below.
	// +kubebuilder:validation:Optional
	ExternalIPv6Address []ExternalIPv6AddressParameters `json:"externalIpv6Address,omitempty" tf:"external_ipv6_address,omitempty"`

	// Internal IPv4 address. The structure is documented below.
	// +kubebuilder:validation:Optional
	InternalIPv4Address []InternalIPv4AddressParameters `json:"internalIpv4Address,omitempty" tf:"internal_ipv4_address,omitempty"`
}

type AllocationPolicyObservation struct {

	// Unique set of locations. The structure is documented below.
	Location []LocationObservation `json:"location,omitempty" tf:"location,omitempty"`
}

type AllocationPolicyParameters struct {

	// Unique set of locations. The structure is documented below.
	// +kubebuilder:validation:Required
	Location []LocationParameters `json:"location" tf:"location,omitempty"`
}

type DefaultHandlerObservation struct {

	// Certificate IDs in the Certificate Manager. Multiple TLS certificates can be associated
	// with the same context to allow both RSA and ECDSA certificates. Only the first certificate of each type will be used.
	CertificateIds []*string `json:"certificateIds,omitempty" tf:"certificate_ids,omitempty"`

	// HTTP handler resource. The structure is documented below.
	HTTPHandler []HTTPHandlerObservation `json:"httpHandler,omitempty" tf:"http_handler,omitempty"`

	// Stream handler resource. The structure is documented below.
	StreamHandler []DefaultHandlerStreamHandlerObservation `json:"streamHandler,omitempty" tf:"stream_handler,omitempty"`
}

type DefaultHandlerParameters struct {

	// Certificate IDs in the Certificate Manager. Multiple TLS certificates can be associated
	// with the same context to allow both RSA and ECDSA certificates. Only the first certificate of each type will be used.
	// +kubebuilder:validation:Required
	CertificateIds []*string `json:"certificateIds" tf:"certificate_ids,omitempty"`

	// HTTP handler resource. The structure is documented below.
	// +kubebuilder:validation:Optional
	HTTPHandler []HTTPHandlerParameters `json:"httpHandler,omitempty" tf:"http_handler,omitempty"`

	// Stream handler resource. The structure is documented below.
	// +kubebuilder:validation:Optional
	StreamHandler []DefaultHandlerStreamHandlerParameters `json:"streamHandler,omitempty" tf:"stream_handler,omitempty"`
}

type DefaultHandlerStreamHandlerObservation struct {

	// Backend group id.
	BackendGroupID *string `json:"backendGroupId,omitempty" tf:"backend_group_id,omitempty"`
}

type DefaultHandlerStreamHandlerParameters struct {

	// Backend group id.
	// +kubebuilder:validation:Optional
	BackendGroupID *string `json:"backendGroupId,omitempty" tf:"backend_group_id,omitempty"`
}

type EndpointObservation struct {

	// Provided by the client or computed automatically.
	Address []AddressObservation `json:"address,omitempty" tf:"address,omitempty"`

	// One or more ports to listen on.
	Ports []*float64 `json:"ports,omitempty" tf:"ports,omitempty"`
}

type EndpointParameters struct {

	// Provided by the client or computed automatically.
	// +kubebuilder:validation:Required
	Address []AddressParameters `json:"address" tf:"address,omitempty"`

	// One or more ports to listen on.
	// +kubebuilder:validation:Required
	Ports []*float64 `json:"ports" tf:"ports,omitempty"`
}

type ExternalIPv4AddressObservation struct {

	// Provided by the client or computed automatically.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`
}

type ExternalIPv4AddressParameters struct {

	// Provided by the client or computed automatically.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`
}

type ExternalIPv6AddressObservation struct {

	// Provided by the client or computed automatically.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`
}

type ExternalIPv6AddressParameters struct {

	// Provided by the client or computed automatically.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`
}

type HTTPHandlerHttp2OptionsObservation struct {

	// Maximum number of concurrent streams.
	MaxConcurrentStreams *float64 `json:"maxConcurrentStreams,omitempty" tf:"max_concurrent_streams,omitempty"`
}

type HTTPHandlerHttp2OptionsParameters struct {

	// Maximum number of concurrent streams.
	// +kubebuilder:validation:Optional
	MaxConcurrentStreams *float64 `json:"maxConcurrentStreams,omitempty" tf:"max_concurrent_streams,omitempty"`
}

type HTTPHandlerObservation struct {

	// If set, will enable only HTTP1 protocol with HTTP1.0 support.
	AllowHttp10 *bool `json:"allowHttp10,omitempty" tf:"allow_http10,omitempty"`

	// HTTP router id.
	HTTPRouterID *string `json:"httpRouterId,omitempty" tf:"http_router_id,omitempty"`

	// If set, will enable HTTP2 protocol for the handler. The structure is documented below.
	Http2Options []HTTPHandlerHttp2OptionsObservation `json:"http2Options,omitempty" tf:"http2_options,omitempty"`
}

type HTTPHandlerParameters struct {

	// If set, will enable only HTTP1 protocol with HTTP1.0 support.
	// +kubebuilder:validation:Optional
	AllowHttp10 *bool `json:"allowHttp10,omitempty" tf:"allow_http10,omitempty"`

	// HTTP router id.
	// +kubebuilder:validation:Optional
	HTTPRouterID *string `json:"httpRouterId,omitempty" tf:"http_router_id,omitempty"`

	// If set, will enable HTTP2 protocol for the handler. The structure is documented below.
	// +kubebuilder:validation:Optional
	Http2Options []HTTPHandlerHttp2OptionsParameters `json:"http2Options,omitempty" tf:"http2_options,omitempty"`
}

type HTTPObservation struct {

	// Stream handler that sets plaintext Stream backend group. The structure is documented below.
	Handler []HandlerObservation `json:"handler,omitempty" tf:"handler,omitempty"`

	// Shortcut for adding http -> https redirects. The structure is documented below.
	Redirects []RedirectsObservation `json:"redirects,omitempty" tf:"redirects,omitempty"`
}

type HTTPParameters struct {

	// Stream handler that sets plaintext Stream backend group. The structure is documented below.
	// +kubebuilder:validation:Optional
	Handler []HandlerParameters `json:"handler,omitempty" tf:"handler,omitempty"`

	// Shortcut for adding http -> https redirects. The structure is documented below.
	// +kubebuilder:validation:Optional
	Redirects []RedirectsParameters `json:"redirects,omitempty" tf:"redirects,omitempty"`
}

type HandlerHTTPHandlerHttp2OptionsObservation struct {

	// Maximum number of concurrent streams.
	MaxConcurrentStreams *float64 `json:"maxConcurrentStreams,omitempty" tf:"max_concurrent_streams,omitempty"`
}

type HandlerHTTPHandlerHttp2OptionsParameters struct {

	// Maximum number of concurrent streams.
	// +kubebuilder:validation:Optional
	MaxConcurrentStreams *float64 `json:"maxConcurrentStreams,omitempty" tf:"max_concurrent_streams,omitempty"`
}

type HandlerHTTPHandlerObservation struct {

	// If set, will enable only HTTP1 protocol with HTTP1.0 support.
	AllowHttp10 *bool `json:"allowHttp10,omitempty" tf:"allow_http10,omitempty"`

	// HTTP router id.
	HTTPRouterID *string `json:"httpRouterId,omitempty" tf:"http_router_id,omitempty"`

	// If set, will enable HTTP2 protocol for the handler. The structure is documented below.
	Http2Options []HandlerHTTPHandlerHttp2OptionsObservation `json:"http2Options,omitempty" tf:"http2_options,omitempty"`
}

type HandlerHTTPHandlerParameters struct {

	// If set, will enable only HTTP1 protocol with HTTP1.0 support.
	// +kubebuilder:validation:Optional
	AllowHttp10 *bool `json:"allowHttp10,omitempty" tf:"allow_http10,omitempty"`

	// HTTP router id.
	// +kubebuilder:validation:Optional
	HTTPRouterID *string `json:"httpRouterId,omitempty" tf:"http_router_id,omitempty"`

	// If set, will enable HTTP2 protocol for the handler. The structure is documented below.
	// +kubebuilder:validation:Optional
	Http2Options []HandlerHTTPHandlerHttp2OptionsParameters `json:"http2Options,omitempty" tf:"http2_options,omitempty"`
}

type HandlerObservation struct {

	// If set, will enable only HTTP1 protocol with HTTP1.0 support.
	AllowHttp10 *bool `json:"allowHttp10,omitempty" tf:"allow_http10,omitempty"`

	// HTTP router id.
	HTTPRouterID *string `json:"httpRouterId,omitempty" tf:"http_router_id,omitempty"`

	// If set, will enable HTTP2 protocol for the handler. The structure is documented below.
	Http2Options []Http2OptionsObservation `json:"http2Options,omitempty" tf:"http2_options,omitempty"`
}

type HandlerParameters struct {

	// If set, will enable only HTTP1 protocol with HTTP1.0 support.
	// +kubebuilder:validation:Optional
	AllowHttp10 *bool `json:"allowHttp10,omitempty" tf:"allow_http10,omitempty"`

	// HTTP router id.
	// +kubebuilder:validation:Optional
	HTTPRouterID *string `json:"httpRouterId,omitempty" tf:"http_router_id,omitempty"`

	// If set, will enable HTTP2 protocol for the handler. The structure is documented below.
	// +kubebuilder:validation:Optional
	Http2Options []Http2OptionsParameters `json:"http2Options,omitempty" tf:"http2_options,omitempty"`
}

type HandlerStreamHandlerObservation struct {

	// Backend group id.
	BackendGroupID *string `json:"backendGroupId,omitempty" tf:"backend_group_id,omitempty"`
}

type HandlerStreamHandlerParameters struct {

	// Backend group id.
	// +kubebuilder:validation:Optional
	BackendGroupID *string `json:"backendGroupId,omitempty" tf:"backend_group_id,omitempty"`
}

type Http2OptionsObservation struct {

	// Maximum number of concurrent streams.
	MaxConcurrentStreams *float64 `json:"maxConcurrentStreams,omitempty" tf:"max_concurrent_streams,omitempty"`
}

type Http2OptionsParameters struct {

	// Maximum number of concurrent streams.
	// +kubebuilder:validation:Optional
	MaxConcurrentStreams *float64 `json:"maxConcurrentStreams,omitempty" tf:"max_concurrent_streams,omitempty"`
}

type InternalIPv4AddressObservation struct {

	// Provided by the client or computed automatically.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Provided by the client or computed automatically.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type InternalIPv4AddressParameters struct {

	// Provided by the client or computed automatically.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Provided by the client or computed automatically.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type ListenerObservation struct {

	// Network endpoints (addresses and ports) of the listener. The structure is documented below.
	Endpoint []EndpointObservation `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// HTTP listener resource. The structure is documented below.
	HTTP []HTTPObservation `json:"http,omitempty" tf:"http,omitempty"`

	// name of the listener.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Stream listener resource. The structure is documented below.
	Stream []StreamObservation `json:"stream,omitempty" tf:"stream,omitempty"`

	// TLS listener resource. The structure is documented below.
	TLS []ListenerTLSObservation `json:"tls,omitempty" tf:"tls,omitempty"`
}

type ListenerParameters struct {

	// Network endpoints (addresses and ports) of the listener. The structure is documented below.
	// +kubebuilder:validation:Optional
	Endpoint []EndpointParameters `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// HTTP listener resource. The structure is documented below.
	// +kubebuilder:validation:Optional
	HTTP []HTTPParameters `json:"http,omitempty" tf:"http,omitempty"`

	// name of the listener.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Stream listener resource. The structure is documented below.
	// +kubebuilder:validation:Optional
	Stream []StreamParameters `json:"stream,omitempty" tf:"stream,omitempty"`

	// TLS listener resource. The structure is documented below.
	// +kubebuilder:validation:Optional
	TLS []ListenerTLSParameters `json:"tls,omitempty" tf:"tls,omitempty"`
}

type ListenerTLSObservation struct {

	// TLS handler resource. The structure is documented below.
	DefaultHandler []DefaultHandlerObservation `json:"defaultHandler,omitempty" tf:"default_handler,omitempty"`

	// SNI match resource. The structure is documented below.
	SniHandler []SniHandlerObservation `json:"sniHandler,omitempty" tf:"sni_handler,omitempty"`
}

type ListenerTLSParameters struct {

	// TLS handler resource. The structure is documented below.
	// +kubebuilder:validation:Required
	DefaultHandler []DefaultHandlerParameters `json:"defaultHandler" tf:"default_handler,omitempty"`

	// SNI match resource. The structure is documented below.
	// +kubebuilder:validation:Optional
	SniHandler []SniHandlerParameters `json:"sniHandler,omitempty" tf:"sni_handler,omitempty"`
}

type LoadBalancerObservation struct {

	// Allocation zones for the Load Balancer instance. The structure is documented below.
	AllocationPolicy []AllocationPolicyObservation `json:"allocationPolicy,omitempty" tf:"allocation_policy,omitempty"`

	// The Load Balancer creation timestamp.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// An optional description of the Load Balancer.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder to which the resource belongs. If omitted, the provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// The ID of the Load Balancer.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Labels to assign to this Load Balancer. A list of key/value pairs.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// List of listeners for the Load Balancer. The structure is documented below.
	Listener []ListenerObservation `json:"listener,omitempty" tf:"listener,omitempty"`

	// Cloud log group used by the Load Balancer to store access logs.
	LogGroupID *string `json:"logGroupId,omitempty" tf:"log_group_id,omitempty"`

	// Name of the Load Balancer. Provided by the client when the Load Balancer is created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network that the Load Balancer is located at.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// ID of the region that the Load Balancer is located at.
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// A list of ID's of security groups attached to the Load Balancer.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// Status of the Load Balancer.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type LoadBalancerParameters struct {

	// Allocation zones for the Load Balancer instance. The structure is documented below.
	// +kubebuilder:validation:Optional
	AllocationPolicy []AllocationPolicyParameters `json:"allocationPolicy,omitempty" tf:"allocation_policy,omitempty"`

	// An optional description of the Load Balancer.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder to which the resource belongs. If omitted, the provider folder is used.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// Labels to assign to this Load Balancer. A list of key/value pairs.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// List of listeners for the Load Balancer. The structure is documented below.
	// +kubebuilder:validation:Optional
	Listener []ListenerParameters `json:"listener,omitempty" tf:"listener,omitempty"`

	// Name of the Load Balancer. Provided by the client when the Load Balancer is created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network that the Load Balancer is located at.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// ID of the region that the Load Balancer is located at.
	// +kubebuilder:validation:Optional
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// A list of ID's of security groups attached to the Load Balancer.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to SecurityGroup in vpc to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsRefs []v1.Reference `json:"securityGroupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in vpc to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsSelector *v1.Selector `json:"securityGroupIdsSelector,omitempty" tf:"-"`
}

type LocationObservation struct {

	// If set, will disable all L7 instances in the zone for request handling.
	DisableTraffic *bool `json:"disableTraffic,omitempty" tf:"disable_traffic,omitempty"`

	// ID of the subnet that location is located at.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// ID of the zone that location is located at.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type LocationParameters struct {

	// If set, will disable all L7 instances in the zone for request handling.
	// +kubebuilder:validation:Optional
	DisableTraffic *bool `json:"disableTraffic,omitempty" tf:"disable_traffic,omitempty"`

	// ID of the subnet that location is located at.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// ID of the zone that location is located at.
	// +kubebuilder:validation:Required
	ZoneID *string `json:"zoneId" tf:"zone_id,omitempty"`
}

type RedirectsObservation struct {

	// If set redirects all unencrypted HTTP requests to the same URI with scheme changed to https.
	HTTPToHTTPS *bool `json:"httpToHttps,omitempty" tf:"http_to_https,omitempty"`
}

type RedirectsParameters struct {

	// If set redirects all unencrypted HTTP requests to the same URI with scheme changed to https.
	// +kubebuilder:validation:Optional
	HTTPToHTTPS *bool `json:"httpToHttps,omitempty" tf:"http_to_https,omitempty"`
}

type SniHandlerHandlerObservation struct {

	// Certificate IDs in the Certificate Manager. Multiple TLS certificates can be associated
	// with the same context to allow both RSA and ECDSA certificates. Only the first certificate of each type will be used.
	CertificateIds []*string `json:"certificateIds,omitempty" tf:"certificate_ids,omitempty"`

	// HTTP handler resource. The structure is documented below.
	HTTPHandler []HandlerHTTPHandlerObservation `json:"httpHandler,omitempty" tf:"http_handler,omitempty"`

	// Stream handler resource. The structure is documented below.
	StreamHandler []HandlerStreamHandlerObservation `json:"streamHandler,omitempty" tf:"stream_handler,omitempty"`
}

type SniHandlerHandlerParameters struct {

	// Certificate IDs in the Certificate Manager. Multiple TLS certificates can be associated
	// with the same context to allow both RSA and ECDSA certificates. Only the first certificate of each type will be used.
	// +kubebuilder:validation:Required
	CertificateIds []*string `json:"certificateIds" tf:"certificate_ids,omitempty"`

	// HTTP handler resource. The structure is documented below.
	// +kubebuilder:validation:Optional
	HTTPHandler []HandlerHTTPHandlerParameters `json:"httpHandler,omitempty" tf:"http_handler,omitempty"`

	// Stream handler resource. The structure is documented below.
	// +kubebuilder:validation:Optional
	StreamHandler []HandlerStreamHandlerParameters `json:"streamHandler,omitempty" tf:"stream_handler,omitempty"`
}

type SniHandlerObservation struct {

	// Stream handler that sets plaintext Stream backend group. The structure is documented below.
	Handler []SniHandlerHandlerObservation `json:"handler,omitempty" tf:"handler,omitempty"`

	// name of SNI match.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A set of server names.
	ServerNames []*string `json:"serverNames,omitempty" tf:"server_names,omitempty"`
}

type SniHandlerParameters struct {

	// Stream handler that sets plaintext Stream backend group. The structure is documented below.
	// +kubebuilder:validation:Required
	Handler []SniHandlerHandlerParameters `json:"handler" tf:"handler,omitempty"`

	// name of SNI match.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A set of server names.
	// +kubebuilder:validation:Required
	ServerNames []*string `json:"serverNames" tf:"server_names,omitempty"`
}

type StreamHandlerObservation struct {

	// Backend group id.
	BackendGroupID *string `json:"backendGroupId,omitempty" tf:"backend_group_id,omitempty"`
}

type StreamHandlerParameters struct {

	// Backend group id.
	// +kubebuilder:validation:Optional
	BackendGroupID *string `json:"backendGroupId,omitempty" tf:"backend_group_id,omitempty"`
}

type StreamObservation struct {

	// Stream handler that sets plaintext Stream backend group. The structure is documented below.
	Handler []StreamHandlerObservation `json:"handler,omitempty" tf:"handler,omitempty"`
}

type StreamParameters struct {

	// Stream handler that sets plaintext Stream backend group. The structure is documented below.
	// +kubebuilder:validation:Optional
	Handler []StreamHandlerParameters `json:"handler,omitempty" tf:"handler,omitempty"`
}

// LoadBalancerSpec defines the desired state of LoadBalancer
type LoadBalancerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadBalancerParameters `json:"forProvider"`
}

// LoadBalancerStatus defines the observed state of LoadBalancer.
type LoadBalancerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadBalancerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancer is the Schema for the LoadBalancers API. A Load Balancer is used for receiving incoming traffic and transmitting it to the backend endpoints specified in the ALB Target Groups.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type LoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.allocationPolicy)",message="allocationPolicy is a required parameter"
	Spec   LoadBalancerSpec   `json:"spec"`
	Status LoadBalancerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerList contains a list of LoadBalancers
type LoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancer `json:"items"`
}

// Repository type metadata.
var (
	LoadBalancer_Kind             = "LoadBalancer"
	LoadBalancer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadBalancer_Kind}.String()
	LoadBalancer_KindAPIVersion   = LoadBalancer_Kind + "." + CRDGroupVersion.String()
	LoadBalancer_GroupVersionKind = CRDGroupVersion.WithKind(LoadBalancer_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadBalancer{}, &LoadBalancerList{})
}
