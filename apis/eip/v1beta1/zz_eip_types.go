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

type BandwidthObservation struct {
}

type BandwidthParameters struct {

	// Specifies whether the bandwidth is billed by traffic or by bandwidth size.
	// Only traffic supported now. Changing this creates a new EIP.
	// +kubebuilder:validation:Optional
	ChargeMode *string `json:"chargeMode,omitempty" tf:"charge_mode,omitempty"`

	// The bandwidth name, which is a string of 1 to 64 characters
	// that contain letters, digits, underscores (_), and hyphens (-).
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the bandwidth type.
	// The value is PER, indicating that the bandwidth is dedicated.
	// Changing this creates a new EIP.
	// +kubebuilder:validation:Required
	ShareType *string `json:"shareType" tf:"share_type,omitempty"`

	// The bandwidth size. The value ranges from 1 to 1000 Mbit/s.
	// +kubebuilder:validation:Required
	Size *float64 `json:"size" tf:"size,omitempty"`
}

type EIPObservation struct {

	// The IP address of the EIP.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The status of EIP.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type EIPParameters struct {

	// The bandwidth object.
	// +kubebuilder:validation:Required
	Bandwidth []BandwidthParameters `json:"bandwidth" tf:"bandwidth,omitempty"`

	// The elastic IP address object.
	// +kubebuilder:validation:Required
	Publicip []PublicipParameters `json:"publicip" tf:"publicip,omitempty"`

	// The region in which to create the EIP. If omitted,
	// the region argument of the provider is used. Changing this creates a new EIP.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The key/value pairs to associate with the EIP.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PublicipObservation struct {
}

type PublicipParameters struct {

	// The value must be a valid IP address in the available IP address segment.
	// Changing this creates a new EIP.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The port id which this EIP will associate with. If the value
	// is not specified, the EIP will be in unbind state.
	// +kubebuilder:validation:Optional
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// The value must be a type supported by the system. Only 5_bgp supported now.
	// Changing this creates a new EIP.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// EIPSpec defines the desired state of EIP
type EIPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EIPParameters `json:"forProvider"`
}

// EIPStatus defines the observed state of EIP.
type EIPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EIPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EIP is the Schema for the EIPs API. ""page_title: "flexibleengine_vpc_eip"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type EIP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EIPSpec   `json:"spec"`
	Status            EIPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EIPList contains a list of EIPs
type EIPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EIP `json:"items"`
}

// Repository type metadata.
var (
	EIP_Kind             = "EIP"
	EIP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EIP_Kind}.String()
	EIP_KindAPIVersion   = EIP_Kind + "." + CRDGroupVersion.String()
	EIP_GroupVersionKind = CRDGroupVersion.WithKind(EIP_Kind)
)

func init() {
	SchemeBuilder.Register(&EIP{}, &EIPList{})
}