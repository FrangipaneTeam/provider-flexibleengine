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

type ApprovalObservation struct {

	// An array of VPC endpoints connect to the VPC endpoint service. Structure is documented below.
	Connections []ConnectionsObservation `json:"connections,omitempty" tf:"connections,omitempty"`

	// The unique ID in UUID format which equals to the ID of the VPC endpoint service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ApprovalParameters struct {

	// Specifies the list of VPC endpoint IDs which accepted to connect to VPC endpoint service.
	// The VPC endpoints will be rejected when the resource was destroyed.
	// +crossplane:generate:reference:type=Endpoint
	// +kubebuilder:validation:Optional
	Endpoints []*string `json:"endpoints,omitempty" tf:"endpoints,omitempty"`

	// References to Endpoint to populate endpoints.
	// +kubebuilder:validation:Optional
	EndpointsRefs []v1.Reference `json:"endpointsRefs,omitempty" tf:"-"`

	// Selector for a list of Endpoint to populate endpoints.
	// +kubebuilder:validation:Optional
	EndpointsSelector *v1.Selector `json:"endpointsSelector,omitempty" tf:"-"`

	// The region in which to obtain the VPC endpoint service.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the ID of the VPC endpoint service. Changing this creates a new resource.
	// +crossplane:generate:reference:type=Service
	// +kubebuilder:validation:Optional
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// Reference to a Service to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDRef *v1.Reference `json:"serviceIdRef,omitempty" tf:"-"`

	// Selector for a Service to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDSelector *v1.Selector `json:"serviceIdSelector,omitempty" tf:"-"`
}

type ConnectionsObservation struct {

	// The user's domain ID.
	DomainID *string `json:"domainId,omitempty" tf:"domain_id,omitempty"`

	// The unique ID of the VPC endpoint.
	EndpointID *string `json:"endpointId,omitempty" tf:"endpoint_id,omitempty"`

	// The packet ID of the VPC endpoint.
	PacketID *float64 `json:"packetId,omitempty" tf:"packet_id,omitempty"`

	// The connection status of the VPC endpoint.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ConnectionsParameters struct {
}

// ApprovalSpec defines the desired state of Approval
type ApprovalSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApprovalParameters `json:"forProvider"`
}

// ApprovalStatus defines the observed state of Approval.
type ApprovalStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApprovalObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Approval is the Schema for the Approvals API. ""page_title: "flexibleengine_vpcep_approval"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Approval struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApprovalSpec   `json:"spec"`
	Status            ApprovalStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApprovalList contains a list of Approvals
type ApprovalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Approval `json:"items"`
}

// Repository type metadata.
var (
	Approval_Kind             = "Approval"
	Approval_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Approval_Kind}.String()
	Approval_KindAPIVersion   = Approval_Kind + "." + CRDGroupVersion.String()
	Approval_GroupVersionKind = CRDGroupVersion.WithKind(Approval_Kind)
)

func init() {
	SchemeBuilder.Register(&Approval{}, &ApprovalList{})
}