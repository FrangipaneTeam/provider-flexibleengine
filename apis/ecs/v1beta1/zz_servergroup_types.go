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

type ServerGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The instances that are part of this server group.
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`
}

type ServerGroupParameters struct {

	// A unique name for the server group. Changing this creates
	// a new server group.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The set of policies for the server group. Only anti-affinity
	// policy is supported right now, which menas all servers in this group must be
	// deployed on different hosts. Changing this creates a new server group.
	// +kubebuilder:validation:Optional
	Policies []*string `json:"policies,omitempty" tf:"policies,omitempty"`

	// The region in which to obtain the V2 Compute client.
	// If omitted, the region argument of the provider is used. Changing
	// this creates a new server group.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Map of additional options.
	// +kubebuilder:validation:Optional
	ValueSpecs map[string]*string `json:"valueSpecs,omitempty" tf:"value_specs,omitempty"`
}

// ServerGroupSpec defines the desired state of ServerGroup
type ServerGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerGroupParameters `json:"forProvider"`
}

// ServerGroupStatus defines the observed state of ServerGroup.
type ServerGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServerGroup is the Schema for the ServerGroups API. ""page_title: "flexibleengine_compute_servergroup_v2"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type ServerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServerGroupSpec   `json:"spec"`
	Status            ServerGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerGroupList contains a list of ServerGroups
type ServerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServerGroup `json:"items"`
}

// Repository type metadata.
var (
	ServerGroup_Kind             = "ServerGroup"
	ServerGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServerGroup_Kind}.String()
	ServerGroup_KindAPIVersion   = ServerGroup_Kind + "." + CRDGroupVersion.String()
	ServerGroup_GroupVersionKind = CRDGroupVersion.WithKind(ServerGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ServerGroup{}, &ServerGroupList{})
}
