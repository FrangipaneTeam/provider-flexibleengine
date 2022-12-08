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

type ConnectionsObservation struct {
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	EndpointID *string `json:"endpointId,omitempty" tf:"endpoint_id,omitempty"`

	PacketID *float64 `json:"packetId,omitempty" tf:"packet_id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ConnectionsParameters struct {
}

type EndpointApprovalObservation struct {
	Connections []ConnectionsObservation `json:"connections,omitempty" tf:"connections,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type EndpointApprovalParameters struct {

	// +kubebuilder:validation:Required
	Endpoints []*string `json:"endpoints" tf:"endpoints,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +crossplane:generate:reference:type=EndpointService
	// +kubebuilder:validation:Optional
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// Reference to a EndpointService to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDRef *v1.Reference `json:"serviceIdRef,omitempty" tf:"-"`

	// Selector for a EndpointService to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDSelector *v1.Selector `json:"serviceIdSelector,omitempty" tf:"-"`
}

// EndpointApprovalSpec defines the desired state of EndpointApproval
type EndpointApprovalSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointApprovalParameters `json:"forProvider"`
}

// EndpointApprovalStatus defines the observed state of EndpointApproval.
type EndpointApprovalStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointApprovalObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointApproval is the Schema for the EndpointApprovals API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type EndpointApproval struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointApprovalSpec   `json:"spec"`
	Status            EndpointApprovalStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointApprovalList contains a list of EndpointApprovals
type EndpointApprovalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EndpointApproval `json:"items"`
}

// Repository type metadata.
var (
	EndpointApproval_Kind             = "EndpointApproval"
	EndpointApproval_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EndpointApproval_Kind}.String()
	EndpointApproval_KindAPIVersion   = EndpointApproval_Kind + "." + CRDGroupVersion.String()
	EndpointApproval_GroupVersionKind = CRDGroupVersion.WithKind(EndpointApproval_Kind)
)

func init() {
	SchemeBuilder.Register(&EndpointApproval{}, &EndpointApprovalList{})
}
