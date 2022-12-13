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

type EIPAssociateObservation struct {

	// The resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The MAC address of the private IP.
	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	// The status of EIP, should be BOUND.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type EIPAssociateParameters struct {

	// Specifies a private IP address to associate with the EIP.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	FixedIP *string `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	// Specifies the ID of the network to which the fixed_ip belongs.
	// It is mandatory when fixed_ip is set. Changing this creates a new resource.
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

	// Specifies an existing port ID to associate with the EIP.
	// This parameter and fixed_ip are alternative. Changing this creates a new resource.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.Port
	// +kubebuilder:validation:Optional
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// Reference to a Port in vpc to populate portId.
	// +kubebuilder:validation:Optional
	PortIDRef *v1.Reference `json:"portIdRef,omitempty" tf:"-"`

	// Selector for a Port in vpc to populate portId.
	// +kubebuilder:validation:Optional
	PortIDSelector *v1.Selector `json:"portIdSelector,omitempty" tf:"-"`

	// Specifies the EIP address to associate. Changing this creates a new resource.
	// +crossplane:generate:reference:type=EIP
	// +crossplane:generate:reference:extractor=github.com/FrangipaneTeam/provider-flexibleengine/config/common.AddressExtractor()
	// +kubebuilder:validation:Optional
	PublicIP *string `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// Reference to a EIP to populate publicIp.
	// +kubebuilder:validation:Optional
	PublicIPRef *v1.Reference `json:"publicIpRef,omitempty" tf:"-"`

	// Selector for a EIP to populate publicIp.
	// +kubebuilder:validation:Optional
	PublicIPSelector *v1.Selector `json:"publicIpSelector,omitempty" tf:"-"`

	// Specifies the region in which to associate the EIP. If omitted, the provider-level
	// region will be used. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// EIPAssociateSpec defines the desired state of EIPAssociate
type EIPAssociateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EIPAssociateParameters `json:"forProvider"`
}

// EIPAssociateStatus defines the observed state of EIPAssociate.
type EIPAssociateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EIPAssociateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EIPAssociate is the Schema for the EIPAssociates API. ""page_title: "flexibleengine_vpc_eip_associate"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type EIPAssociate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EIPAssociateSpec   `json:"spec"`
	Status            EIPAssociateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EIPAssociateList contains a list of EIPAssociates
type EIPAssociateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EIPAssociate `json:"items"`
}

// Repository type metadata.
var (
	EIPAssociate_Kind             = "EIPAssociate"
	EIPAssociate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EIPAssociate_Kind}.String()
	EIPAssociate_KindAPIVersion   = EIPAssociate_Kind + "." + CRDGroupVersion.String()
	EIPAssociate_GroupVersionKind = CRDGroupVersion.WithKind(EIPAssociate_Kind)
)

func init() {
	SchemeBuilder.Register(&EIPAssociate{}, &EIPAssociateList{})
}
