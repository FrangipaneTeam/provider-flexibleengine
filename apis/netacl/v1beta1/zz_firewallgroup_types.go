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

type FirewallGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type FirewallGroupParameters struct {

	// Administrative up/down status for the firewall group
	// (must be "true" or "false" if provided - defaults to "true").
	// Changing this updates the admin_state_up of an existing firewall group.
	// +kubebuilder:validation:Optional
	AdminStateUp *bool `json:"adminStateUp,omitempty" tf:"admin_state_up,omitempty"`

	// A description for the firewall group. Changing this
	// updates the description of an existing firewall group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The egress policy resource id for the firewall group. Changing
	// this updates the egress_policy_id of an existing firewall group.
	// +kubebuilder:validation:Optional
	EgressPolicyID *string `json:"egressPolicyId,omitempty" tf:"egress_policy_id,omitempty"`

	// The ingress policy resource id for the firewall group. Changing
	// this updates the ingress_policy_id of an existing firewall group.
	// +kubebuilder:validation:Optional
	IngressPolicyID *string `json:"ingressPolicyId,omitempty" tf:"ingress_policy_id,omitempty"`

	// A name for the firewall group. Changing this
	// updates the name of an existing firewall group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Port(s) to associate this firewall group instance
	// with. Must be a list of strings. Changing this updates the associated routers
	// of an existing firewall group.
	// +kubebuilder:validation:Optional
	Ports []*string `json:"ports,omitempty" tf:"ports,omitempty"`

	// The region in which to obtain the v2 networking client.
	// A networking client is needed to create a firewall group. If omitted, the
	// region argument of the provider is used. Changing this creates a new
	// firewall group.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The owner of the floating IP. Required if admin wants
	// to create a firewall group for another tenant. Changing this creates a new
	// firewall group.
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

// FirewallGroupSpec defines the desired state of FirewallGroup
type FirewallGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallGroupParameters `json:"forProvider"`
}

// FirewallGroupStatus defines the observed state of FirewallGroup.
type FirewallGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallGroup is the Schema for the FirewallGroups API. ""page_title: "flexibleengine_fw_firewall_group_v2"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type FirewallGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallGroupSpec   `json:"spec"`
	Status            FirewallGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallGroupList contains a list of FirewallGroups
type FirewallGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallGroup `json:"items"`
}

// Repository type metadata.
var (
	FirewallGroup_Kind             = "FirewallGroup"
	FirewallGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallGroup_Kind}.String()
	FirewallGroup_KindAPIVersion   = FirewallGroup_Kind + "." + CRDGroupVersion.String()
	FirewallGroup_GroupVersionKind = CRDGroupVersion.WithKind(FirewallGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallGroup{}, &FirewallGroupList{})
}
