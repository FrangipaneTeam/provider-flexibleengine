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

type RouterInterfaceObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RouterInterfaceParameters struct {

	// +crossplane:generate:reference:type=NetworkPort
	// +kubebuilder:validation:Optional
	PortID *string `json:"portId,omitempty" tf:"port_id,omitempty"`

	// Reference to a NetworkPort to populate portId.
	// +kubebuilder:validation:Optional
	PortIDRef *v1.Reference `json:"portIdRef,omitempty" tf:"-"`

	// Selector for a NetworkPort to populate portId.
	// +kubebuilder:validation:Optional
	PortIDSelector *v1.Selector `json:"portIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +crossplane:generate:reference:type=Router
	// +kubebuilder:validation:Optional
	RouterID *string `json:"routerId,omitempty" tf:"router_id,omitempty"`

	// Reference to a Router to populate routerId.
	// +kubebuilder:validation:Optional
	RouterIDRef *v1.Reference `json:"routerIdRef,omitempty" tf:"-"`

	// Selector for a Router to populate routerId.
	// +kubebuilder:validation:Optional
	RouterIDSelector *v1.Selector `json:"routerIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// RouterInterfaceSpec defines the desired state of RouterInterface
type RouterInterfaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouterInterfaceParameters `json:"forProvider"`
}

// RouterInterfaceStatus defines the observed state of RouterInterface.
type RouterInterfaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouterInterfaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RouterInterface is the Schema for the RouterInterfaces API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type RouterInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouterInterfaceSpec   `json:"spec"`
	Status            RouterInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouterInterfaceList contains a list of RouterInterfaces
type RouterInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouterInterface `json:"items"`
}

// Repository type metadata.
var (
	RouterInterface_Kind             = "RouterInterface"
	RouterInterface_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouterInterface_Kind}.String()
	RouterInterface_KindAPIVersion   = RouterInterface_Kind + "." + CRDGroupVersion.String()
	RouterInterface_GroupVersionKind = CRDGroupVersion.WithKind(RouterInterface_Kind)
)

func init() {
	SchemeBuilder.Register(&RouterInterface{}, &RouterInterfaceList{})
}
