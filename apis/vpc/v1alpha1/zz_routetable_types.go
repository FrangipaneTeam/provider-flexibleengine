/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RouteTableObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RouteTableParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Optional
	Route []RouteTableRouteParameters `json:"route,omitempty" tf:"route,omitempty"`

	// +crossplane:generate:reference:type=Subnet
	// +kubebuilder:validation:Optional
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// References to Subnet to populate subnets.
	// +kubebuilder:validation:Optional
	SubnetsRefs []v1.Reference `json:"subnetsRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet to populate subnets.
	// +kubebuilder:validation:Optional
	SubnetsSelector *v1.Selector `json:"subnetsSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/gaetanars/provider-flexibleengine/apis/vpc/v1alpha1.Vpc
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a Vpc in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a Vpc in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type RouteTableRouteObservation struct {
}

type RouteTableRouteParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Destination *string `json:"destination" tf:"destination,omitempty"`

	// +kubebuilder:validation:Required
	Nexthop *string `json:"nexthop" tf:"nexthop,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// RouteTableSpec defines the desired state of RouteTable
type RouteTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteTableParameters `json:"forProvider"`
}

// RouteTableStatus defines the observed state of RouteTable.
type RouteTableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTable is the Schema for the RouteTables API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type RouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTableSpec   `json:"spec"`
	Status            RouteTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTableList contains a list of RouteTables
type RouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTable `json:"items"`
}

// Repository type metadata.
var (
	RouteTable_Kind             = "RouteTable"
	RouteTable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouteTable_Kind}.String()
	RouteTable_KindAPIVersion   = RouteTable_Kind + "." + CRDGroupVersion.String()
	RouteTable_GroupVersionKind = CRDGroupVersion.WithKind(RouteTable_Kind)
)

func init() {
	SchemeBuilder.Register(&RouteTable{}, &RouteTableList{})
}
