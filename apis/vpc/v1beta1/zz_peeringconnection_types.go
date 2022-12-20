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

type PeeringConnectionObservation struct {

	// The VPC peering connection ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The VPC peering connection status. The value can be PENDING_ACCEPTANCE, REJECTED, EXPIRED, DELETED, or ACTIVE.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type PeeringConnectionParameters struct {

	// Specifies the name of the VPC peering connection. The value can contain 1 to 64 characters.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specified the Tenant Id of the accepter tenant. Changing this creates a new VPC peering connection.
	// +kubebuilder:validation:Optional
	PeerTenantID *string `json:"peerTenantId,omitempty" tf:"peer_tenant_id,omitempty"`

	// Specifies the VPC ID of the accepter tenant. Changing this creates a new VPC peering connection.
	// +crossplane:generate:reference:type=VPC
	// +kubebuilder:validation:Optional
	PeerVPCID *string `json:"peerVpcId,omitempty" tf:"peer_vpc_id,omitempty"`

	// Reference to a VPC to populate peerVpcId.
	// +kubebuilder:validation:Optional
	PeerVPCIDRef *v1.Reference `json:"peerVpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC to populate peerVpcId.
	// +kubebuilder:validation:Optional
	PeerVPCIDSelector *v1.Selector `json:"peerVpcIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the ID of a VPC involved in a VPC peering connection.
	// Changing this creates a new VPC peering connection.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// PeeringConnectionSpec defines the desired state of PeeringConnection
type PeeringConnectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PeeringConnectionParameters `json:"forProvider"`
}

// PeeringConnectionStatus defines the observed state of PeeringConnection.
type PeeringConnectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PeeringConnectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PeeringConnection is the Schema for the PeeringConnections API. ""page_title: "flexibleengine_vpc_peering_connection_v2"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type PeeringConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PeeringConnectionSpec   `json:"spec"`
	Status            PeeringConnectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PeeringConnectionList contains a list of PeeringConnections
type PeeringConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PeeringConnection `json:"items"`
}

// Repository type metadata.
var (
	PeeringConnection_Kind             = "PeeringConnection"
	PeeringConnection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PeeringConnection_Kind}.String()
	PeeringConnection_KindAPIVersion   = PeeringConnection_Kind + "." + CRDGroupVersion.String()
	PeeringConnection_GroupVersionKind = CRDGroupVersion.WithKind(PeeringConnection_Kind)
)

func init() {
	SchemeBuilder.Register(&PeeringConnection{}, &PeeringConnectionList{})
}
