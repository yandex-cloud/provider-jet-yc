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

type AddressObservation struct {
}

type AddressParameters struct {

	// +kubebuilder:validation:Optional
	ExternalIPv4Address []ExternalIPv4AddressParameters `json:"externalIpv4Address,omitempty" tf:"external_ipv4_address,omitempty"`

	// +kubebuilder:validation:Optional
	ExternalIPv6Address []ExternalIPv6AddressParameters `json:"externalIpv6Address,omitempty" tf:"external_ipv6_address,omitempty"`

	// +kubebuilder:validation:Optional
	InternalIPv4Address []InternalIPv4AddressParameters `json:"internalIpv4Address,omitempty" tf:"internal_ipv4_address,omitempty"`
}

type AllocationPolicyObservation struct {
}

type AllocationPolicyParameters struct {

	// +kubebuilder:validation:Required
	Location []LocationParameters `json:"location" tf:"location,omitempty"`
}

type DefaultHandlerObservation struct {
}

type DefaultHandlerParameters struct {

	// +kubebuilder:validation:Required
	CertificateIds []*string `json:"certificateIds" tf:"certificate_ids,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPHandler []HTTPHandlerParameters `json:"httpHandler,omitempty" tf:"http_handler,omitempty"`
}

type EndpointObservation struct {
}

type EndpointParameters struct {

	// +kubebuilder:validation:Required
	Address []AddressParameters `json:"address" tf:"address,omitempty"`

	// +kubebuilder:validation:Required
	Ports []*int64 `json:"ports" tf:"ports,omitempty"`
}

type ExternalIPv4AddressObservation struct {
}

type ExternalIPv4AddressParameters struct {

	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`
}

type ExternalIPv6AddressObservation struct {
}

type ExternalIPv6AddressParameters struct {

	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`
}

type HTTPHandlerHttp2OptionsObservation struct {
}

type HTTPHandlerHttp2OptionsParameters struct {

	// +kubebuilder:validation:Optional
	MaxConcurrentStreams *int64 `json:"maxConcurrentStreams,omitempty" tf:"max_concurrent_streams,omitempty"`
}

type HTTPHandlerObservation struct {
}

type HTTPHandlerParameters struct {

	// +kubebuilder:validation:Optional
	AllowHttp10 *bool `json:"allowHttp10,omitempty" tf:"allow_http10,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPRouterID *string `json:"httpRouterId,omitempty" tf:"http_router_id,omitempty"`

	// +kubebuilder:validation:Optional
	Http2Options []HTTPHandlerHttp2OptionsParameters `json:"http2Options,omitempty" tf:"http2_options,omitempty"`
}

type HTTPObservation struct {
}

type HTTPParameters struct {

	// +kubebuilder:validation:Optional
	Handler []HandlerParameters `json:"handler,omitempty" tf:"handler,omitempty"`

	// +kubebuilder:validation:Optional
	Redirects []RedirectsParameters `json:"redirects,omitempty" tf:"redirects,omitempty"`
}

type HandlerHTTPHandlerHttp2OptionsObservation struct {
}

type HandlerHTTPHandlerHttp2OptionsParameters struct {

	// +kubebuilder:validation:Optional
	MaxConcurrentStreams *int64 `json:"maxConcurrentStreams,omitempty" tf:"max_concurrent_streams,omitempty"`
}

type HandlerHTTPHandlerObservation struct {
}

type HandlerHTTPHandlerParameters struct {

	// +kubebuilder:validation:Optional
	AllowHttp10 *bool `json:"allowHttp10,omitempty" tf:"allow_http10,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPRouterID *string `json:"httpRouterId,omitempty" tf:"http_router_id,omitempty"`

	// +kubebuilder:validation:Optional
	Http2Options []HandlerHTTPHandlerHttp2OptionsParameters `json:"http2Options,omitempty" tf:"http2_options,omitempty"`
}

type HandlerObservation struct {
}

type HandlerParameters struct {

	// +kubebuilder:validation:Optional
	AllowHttp10 *bool `json:"allowHttp10,omitempty" tf:"allow_http10,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPRouterID *string `json:"httpRouterId,omitempty" tf:"http_router_id,omitempty"`

	// +kubebuilder:validation:Optional
	Http2Options []Http2OptionsParameters `json:"http2Options,omitempty" tf:"http2_options,omitempty"`
}

type Http2OptionsObservation struct {
}

type Http2OptionsParameters struct {

	// +kubebuilder:validation:Optional
	MaxConcurrentStreams *int64 `json:"maxConcurrentStreams,omitempty" tf:"max_concurrent_streams,omitempty"`
}

type InternalIPv4AddressObservation struct {
}

type InternalIPv4AddressParameters struct {

	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type ListenerObservation struct {
}

type ListenerParameters struct {

	// +kubebuilder:validation:Optional
	Endpoint []EndpointParameters `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	HTTP []HTTPParameters `json:"http,omitempty" tf:"http,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	TLS []ListenerTLSParameters `json:"tls,omitempty" tf:"tls,omitempty"`
}

type ListenerTLSObservation struct {
}

type ListenerTLSParameters struct {

	// +kubebuilder:validation:Required
	DefaultHandler []DefaultHandlerParameters `json:"defaultHandler" tf:"default_handler,omitempty"`

	// +kubebuilder:validation:Optional
	SniHandler []SniHandlerParameters `json:"sniHandler,omitempty" tf:"sni_handler,omitempty"`
}

type LoadBalancerObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	LogGroupID *string `json:"logGroupId,omitempty" tf:"log_group_id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type LoadBalancerParameters struct {

	// +kubebuilder:validation:Required
	AllocationPolicy []AllocationPolicyParameters `json:"allocationPolicy" tf:"allocation_policy,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	Listener []ListenerParameters `json:"listener,omitempty" tf:"listener,omitempty"`

	// +kubebuilder:validation:Required
	NetworkID *string `json:"networkId" tf:"network_id,omitempty"`

	// +kubebuilder:validation:Optional
	RegionID *string `json:"regionId,omitempty" tf:"region_id,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`
}

type LocationObservation struct {
}

type LocationParameters struct {

	// +kubebuilder:validation:Optional
	DisableTraffic *bool `json:"disableTraffic,omitempty" tf:"disable_traffic,omitempty"`

	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Required
	ZoneID *string `json:"zoneId" tf:"zone_id,omitempty"`
}

type RedirectsObservation struct {
}

type RedirectsParameters struct {

	// +kubebuilder:validation:Optional
	HTTPToHTTPS *bool `json:"httpToHttps,omitempty" tf:"http_to_https,omitempty"`
}

type SniHandlerHandlerObservation struct {
}

type SniHandlerHandlerParameters struct {

	// +kubebuilder:validation:Required
	CertificateIds []*string `json:"certificateIds" tf:"certificate_ids,omitempty"`

	// +kubebuilder:validation:Optional
	HTTPHandler []HandlerHTTPHandlerParameters `json:"httpHandler,omitempty" tf:"http_handler,omitempty"`
}

type SniHandlerObservation struct {
}

type SniHandlerParameters struct {

	// +kubebuilder:validation:Required
	Handler []SniHandlerHandlerParameters `json:"handler" tf:"handler,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ServerNames []*string `json:"serverNames" tf:"server_names,omitempty"`
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

// LoadBalancer is the Schema for the LoadBalancers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type LoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoadBalancerSpec   `json:"spec"`
	Status            LoadBalancerStatus `json:"status,omitempty"`
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
