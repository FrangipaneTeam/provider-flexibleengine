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

type InterfaceAttachObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type InterfaceAttachParameters struct {

	// +kubebuilder:validation:Optional
	FixedIP *string `json:"fixedIp,omitempty" tf:"fixed_ip,omitempty"`

	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/gaetanars/provider-flexibleengine/apis/vpc/v1alpha1.Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/gaetanars/provider-flexibleengine/apis/vpc/v1alpha1.NetworkPort
	// +kubebuilder:validation:Optional
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// Reference to a NetworkPort in vpc to populate portId.
	// +kubebuilder:validation:Optional
	PortIDRef *v1.Reference `json:"portIdRef,omitempty" tf:"-"`

	// Selector for a NetworkPort in vpc to populate portId.
	// +kubebuilder:validation:Optional
	PortIDSelector *v1.Selector `json:"portIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// InterfaceAttachSpec defines the desired state of InterfaceAttach
type InterfaceAttachSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InterfaceAttachParameters `json:"forProvider"`
}

// InterfaceAttachStatus defines the observed state of InterfaceAttach.
type InterfaceAttachStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InterfaceAttachObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InterfaceAttach is the Schema for the InterfaceAttachs API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type InterfaceAttach struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InterfaceAttachSpec   `json:"spec"`
	Status            InterfaceAttachStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InterfaceAttachList contains a list of InterfaceAttachs
type InterfaceAttachList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InterfaceAttach `json:"items"`
}

// Repository type metadata.
var (
	InterfaceAttach_Kind             = "InterfaceAttach"
	InterfaceAttach_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InterfaceAttach_Kind}.String()
	InterfaceAttach_KindAPIVersion   = InterfaceAttach_Kind + "." + CRDGroupVersion.String()
	InterfaceAttach_GroupVersionKind = CRDGroupVersion.WithKind(InterfaceAttach_Kind)
)

func init() {
	SchemeBuilder.Register(&InterfaceAttach{}, &InterfaceAttachList{})
}