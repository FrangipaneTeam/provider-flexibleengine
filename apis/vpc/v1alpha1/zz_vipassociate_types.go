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

type VipAssociateObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	VipIPAddress *string `json:"vipIpAddress,omitempty" tf:"vip_ip_address,omitempty"`

	VipSubnetID *string `json:"vipSubnetId,omitempty" tf:"vip_subnet_id,omitempty"`
}

type VipAssociateParameters struct {

	// +crossplane:generate:reference:type=NetworkPort
	// +kubebuilder:validation:Optional
	PortIds []*string `json:"portIds,omitempty" tf:"port_ids,omitempty"`

	// References to NetworkPort to populate portIds.
	// +kubebuilder:validation:Optional
	PortIdsRefs []v1.Reference `json:"portIdsRefs,omitempty" tf:"-"`

	// Selector for a list of NetworkPort to populate portIds.
	// +kubebuilder:validation:Optional
	PortIdsSelector *v1.Selector `json:"portIdsSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=Vip
	// +kubebuilder:validation:Optional
	VipID *string `json:"vipId,omitempty" tf:"vip_id,omitempty"`

	// Reference to a Vip to populate vipId.
	// +kubebuilder:validation:Optional
	VipIDRef *v1.Reference `json:"vipIdRef,omitempty" tf:"-"`

	// Selector for a Vip to populate vipId.
	// +kubebuilder:validation:Optional
	VipIDSelector *v1.Selector `json:"vipIdSelector,omitempty" tf:"-"`
}

// VipAssociateSpec defines the desired state of VipAssociate
type VipAssociateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VipAssociateParameters `json:"forProvider"`
}

// VipAssociateStatus defines the observed state of VipAssociate.
type VipAssociateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VipAssociateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VipAssociate is the Schema for the VipAssociates API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type VipAssociate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VipAssociateSpec   `json:"spec"`
	Status            VipAssociateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VipAssociateList contains a list of VipAssociates
type VipAssociateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VipAssociate `json:"items"`
}

// Repository type metadata.
var (
	VipAssociate_Kind             = "VipAssociate"
	VipAssociate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VipAssociate_Kind}.String()
	VipAssociate_KindAPIVersion   = VipAssociate_Kind + "." + CRDGroupVersion.String()
	VipAssociate_GroupVersionKind = CRDGroupVersion.WithKind(VipAssociate_Kind)
)

func init() {
	SchemeBuilder.Register(&VipAssociate{}, &VipAssociateList{})
}