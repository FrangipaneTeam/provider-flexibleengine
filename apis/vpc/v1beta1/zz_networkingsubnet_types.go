/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AllocationPoolsObservation struct {
}

type AllocationPoolsParameters struct {

	// The ending addresss.
	// +kubebuilder:validation:Required
	End *string `json:"end" tf:"end,omitempty"`

	// The starting address.
	// +kubebuilder:validation:Required
	Start *string `json:"start" tf:"start,omitempty"`
}

type HostRoutesObservation struct {
}

type HostRoutesParameters struct {

	// The destination CIDR.
	// +kubebuilder:validation:Required
	DestinationCidr *string `json:"destinationCidr" tf:"destination_cidr,omitempty"`

	// The next hop in the route.
	// +kubebuilder:validation:Required
	NextHop *string `json:"nextHop" tf:"next_hop,omitempty"`
}

type NetworkingSubnetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NetworkingSubnetParameters struct {

	// An array of sub-ranges of CIDR available for
	// dynamic allocation to ports. The allocation_pool object structure is
	// documented below. Changing this creates a new subnet.
	// +kubebuilder:validation:Optional
	AllocationPools []AllocationPoolsParameters `json:"allocationPools,omitempty" tf:"allocation_pools,omitempty"`

	// CIDR representing IP range for this subnet, based on IP
	// version. Changing this creates a new subnet.
	// +kubebuilder:validation:Required
	Cidr *string `json:"cidr" tf:"cidr,omitempty"`

	// An array of DNS name server names used by hosts
	// in this subnet. Changing this updates the DNS name servers for the existing
	// subnet.
	// +kubebuilder:validation:Optional
	DNSNameservers []*string `json:"dnsNameservers,omitempty" tf:"dns_nameservers,omitempty"`

	// The administrative state of the network.
	// Acceptable values are "true" and "false". Changing this value enables or
	// disables the DHCP capabilities of the existing subnet. Defaults to true.
	// +kubebuilder:validation:Optional
	EnableDHCP *bool `json:"enableDhcp,omitempty" tf:"enable_dhcp,omitempty"`

	// Default gateway used by devices in this subnet.
	// Leaving this blank and not setting no_gateway will cause a default
	// gateway of .1 to be used. Changing this updates the gateway IP of the
	// existing subnet.
	// +kubebuilder:validation:Optional
	GatewayIP *string `json:"gatewayIp,omitempty" tf:"gateway_ip,omitempty"`

	// An array of routes that should be used by devices
	// with IPs from this subnet (not including local subnet route). The host_route
	// object structure is documented below. Changing this updates the host routes
	// for the existing subnet.
	// +kubebuilder:validation:Optional
	HostRoutes []HostRoutesParameters `json:"hostRoutes,omitempty" tf:"host_routes,omitempty"`

	// IP version, either 4 (default) or 6. Changing this creates a
	// new subnet.
	// +kubebuilder:validation:Optional
	IPVersion *float64 `json:"ipVersion,omitempty" tf:"ip_version,omitempty"`

	// The name of the subnet. Changing this updates the name of
	// the existing subnet.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The UUID of the parent network. Changing this
	// creates a new subnet.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.VPCSubnet
	// +crossplane:generate:reference:extractor=github.com/FrangipaneTeam/provider-flexibleengine/config/common.IDExtractor()
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a VPCSubnet in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a VPCSubnet in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// Do not set a gateway IP on this subnet. Changing
	// this removes or adds a default gateway IP of the existing subnet.
	// +kubebuilder:validation:Optional
	NoGateway *bool `json:"noGateway,omitempty" tf:"no_gateway,omitempty"`

	// The region in which to obtain the V2 Networking client.
	// A Networking client is needed to create a Neutron subnet. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// subnet.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The owner of the subnet. Required if admin wants to
	// create a subnet for another tenant. Changing this creates a new subnet.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/iam/v1beta1.Project
	// +kubebuilder:validation:Optional
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Reference to a Project in iam to populate tenantId.
	// +kubebuilder:validation:Optional
	TenantIDRef *v1.Reference `json:"tenantIdRef,omitempty" tf:"-"`

	// Selector for a Project in iam to populate tenantId.
	// +kubebuilder:validation:Optional
	TenantIDSelector *v1.Selector `json:"tenantIdSelector,omitempty" tf:"-"`

	// Map of additional options.
	// +kubebuilder:validation:Optional
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

// NetworkingSubnetSpec defines the desired state of NetworkingSubnet
type NetworkingSubnetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NetworkingSubnetParameters `json:"forProvider"`
}

// NetworkingSubnetStatus defines the observed state of NetworkingSubnet.
type NetworkingSubnetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NetworkingSubnetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkingSubnet is the Schema for the NetworkingSubnets API. ""page_title: "flexibleengine_networking_subnet_v2"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type NetworkingSubnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkingSubnetSpec   `json:"spec"`
	Status            NetworkingSubnetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NetworkingSubnetList contains a list of NetworkingSubnets
type NetworkingSubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkingSubnet `json:"items"`
}

// Repository type metadata.
var (
	NetworkingSubnet_Kind             = "NetworkingSubnet"
	NetworkingSubnet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NetworkingSubnet_Kind}.String()
	NetworkingSubnet_KindAPIVersion   = NetworkingSubnet_Kind + "." + CRDGroupVersion.String()
	NetworkingSubnet_GroupVersionKind = CRDGroupVersion.WithKind(NetworkingSubnet_Kind)
)

func init() {
	SchemeBuilder.Register(&NetworkingSubnet{}, &NetworkingSubnetList{})
}
