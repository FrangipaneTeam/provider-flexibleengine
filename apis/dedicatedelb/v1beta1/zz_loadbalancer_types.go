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

type LoadBalancerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ipv4 eip address of the Load Balancer.
	IPv4EIP *string `json:"ipv4Eip,omitempty" tf:"ipv4_eip,omitempty"`

	// The ipv6 address of the Load Balancer.
	IPv6Address *string `json:"ipv6Address,omitempty" tf:"ipv6_address,omitempty"`

	// The ipv6 eip address of the Load Balancer.
	IPv6EIP *string `json:"ipv6Eip,omitempty" tf:"ipv6_eip,omitempty"`

	// The ipv6 eip id of the Load Balancer.
	IPv6EIPID *string `json:"ipv6EipId,omitempty" tf:"ipv6_eip_id,omitempty"`
}

type LoadBalancerParameters struct {

	// +kubebuilder:validation:Optional
	AutoPay *string `json:"autoPay,omitempty" tf:"auto_pay,omitempty"`

	// +kubebuilder:validation:Optional
	AutoRenew *string `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// +kubebuilder:validation:Optional
	AutoscalingEnabled *bool `json:"autoscalingEnabled,omitempty" tf:"autoscaling_enabled,omitempty"`

	// Specifies the list of AZ names. Changing this parameter will create a
	// new resource.
	// +kubebuilder:validation:Required
	AvailabilityZone []*string `json:"availabilityZone" tf:"availability_zone,omitempty"`

	// Bandwidth billing type. Changing this parameter will create a
	// new resource.
	// +kubebuilder:validation:Optional
	BandwidthChargeMode *string `json:"bandwidthChargeMode,omitempty" tf:"bandwidth_charge_mode,omitempty"`

	// Bandwidth size. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	BandwidthSize *float64 `json:"bandwidthSize,omitempty" tf:"bandwidth_size,omitempty"`

	// +kubebuilder:validation:Optional
	ChargingMode *string `json:"chargingMode,omitempty" tf:"charging_mode,omitempty"`

	// Enable this if you want to associate the IP addresses of backend servers with
	// your load balancer. Can only be true when updating.
	// +kubebuilder:validation:Optional
	CrossVPCBackend *bool `json:"crossVpcBackend,omitempty" tf:"cross_vpc_backend,omitempty"`

	// Human-readable description for the loadbalancer.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// The ipv4 address of the load balancer.
	// +kubebuilder:validation:Optional
	IPv4Address *string `json:"ipv4Address,omitempty" tf:"ipv4_address,omitempty"`

	// The ID of the EIP. Changing this parameter will create a new resource.
	// +crossplane:generate:reference:type=github.com/gaetanars/provider-flexibleengine/apis/eip/v1beta1.EIP
	// +kubebuilder:validation:Optional
	IPv4EIPID *string `json:"ipv4EipId,omitempty" tf:"ipv4_eip_id,omitempty"`

	// Reference to a EIP in eip to populate ipv4EipId.
	// +kubebuilder:validation:Optional
	IPv4EIPIDRef *v1.Reference `json:"ipv4EipIdRef,omitempty" tf:"-"`

	// Selector for a EIP in eip to populate ipv4EipId.
	// +kubebuilder:validation:Optional
	IPv4EIPIDSelector *v1.Selector `json:"ipv4EipIdSelector,omitempty" tf:"-"`

	// The subnet on which to allocate the loadbalancer's ipv4 address.
	// the IPv4 subnet ID of the subnet where the load balancer resides
	// +crossplane:generate:reference:type=github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1.VPCSubnet
	// +kubebuilder:validation:Optional
	IPv4SubnetID *string `json:"ipv4SubnetId,omitempty" tf:"ipv4_subnet_id,omitempty"`

	// Reference to a VPCSubnet in vpc to populate ipv4SubnetId.
	// +kubebuilder:validation:Optional
	IPv4SubnetIDRef *v1.Reference `json:"ipv4SubnetIdRef,omitempty" tf:"-"`

	// Selector for a VPCSubnet in vpc to populate ipv4SubnetId.
	// +kubebuilder:validation:Optional
	IPv4SubnetIDSelector *v1.Selector `json:"ipv4SubnetIdSelector,omitempty" tf:"-"`

	// The ipv6 bandwidth id. Only support shared bandwidth.
	// +kubebuilder:validation:Optional
	IPv6BandwidthID *string `json:"ipv6BandwidthId,omitempty" tf:"ipv6_bandwidth_id,omitempty"`

	// The network on which to allocate the loadbalancer's ipv6 address.
	// the ID of the subnet where the load balancer resides
	// +crossplane:generate:reference:type=github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1.Network
	// +kubebuilder:validation:Optional
	IPv6NetworkID *string `json:"ipv6NetworkId,omitempty" tf:"ipv6_network_id,omitempty"`

	// Reference to a Network in vpc to populate ipv6NetworkId.
	// +kubebuilder:validation:Optional
	IPv6NetworkIDRef *v1.Reference `json:"ipv6NetworkIdRef,omitempty" tf:"-"`

	// Selector for a Network in vpc to populate ipv6NetworkId.
	// +kubebuilder:validation:Optional
	IPv6NetworkIDSelector *v1.Selector `json:"ipv6NetworkIdSelector,omitempty" tf:"-"`

	// Elastic IP type. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Iptype *string `json:"iptype,omitempty" tf:"iptype,omitempty"`

	// The L4 flavor id of the load balancer.
	// +kubebuilder:validation:Optional
	L4FlavorID *string `json:"l4FlavorId,omitempty" tf:"l4_flavor_id,omitempty"`

	// The L7 flavor id of the load balancer.
	// +kubebuilder:validation:Optional
	L7FlavorID *string `json:"l7FlavorId,omitempty" tf:"l7_flavor_id,omitempty"`

	// The L7 flavor id of the load balancer.
	// +kubebuilder:validation:Optional
	MinL7FlavorID *string `json:"minL7FlavorId,omitempty" tf:"min_l7_flavor_id,omitempty"`

	// Human-readable name for the loadbalancer.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// +kubebuilder:validation:Optional
	PeriodUnit *string `json:"periodUnit,omitempty" tf:"period_unit,omitempty"`

	// The region in which to create the loadbalancer resource. If omitted, the
	// provider-level region will be used. Changing this creates a new loadbalancer.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Bandwidth sharing type. Changing this parameter will create a new resource.
	// +kubebuilder:validation:Optional
	Sharetype *string `json:"sharetype,omitempty" tf:"sharetype,omitempty"`

	// The key/value pairs to associate with the loadbalancer.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The vpc on which to create the loadbalancer. Changing this creates a new
	// loadbalancer.
	// +crossplane:generate:reference:type=github.com/gaetanars/provider-flexibleengine/apis/vpc/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
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

// LoadBalancer is the Schema for the LoadBalancers API. ""page_title: "flexibleengine_lb_loadbalancer_v3"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
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