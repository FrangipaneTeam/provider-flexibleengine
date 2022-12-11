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

type AgencyObservation struct {

	// The time when the agency was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The expiration time of agency.
	ExpireTime *string `json:"expireTime,omitempty" tf:"expire_time,omitempty"`

	// The agency ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AgencyParameters struct {

	// Specifies the name of delegated user domain.
	// This parameter and delegated_service_name are alternative.
	// +kubebuilder:validation:Optional
	DelegatedDomainName *string `json:"delegatedDomainName,omitempty" tf:"delegated_domain_name,omitempty"`

	// Specifies the name of delegated cloud service.
	// This parameter and delegated_domain_name are alternative.
	// +kubebuilder:validation:Optional
	DelegatedServiceName *string `json:"delegatedServiceName,omitempty" tf:"delegated_service_name,omitempty"`

	// Specifies the supplementary information about the agency.
	// The value is a string of 0 to 255 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies an array of one or more role names which stand for the permissionis to
	// be granted to agency on domain.
	// +kubebuilder:validation:Optional
	DomainRoles []*string `json:"domainRoles,omitempty" tf:"domain_roles,omitempty"`

	// Specifies the validity period of an agency.
	// The valid value are ONEDAY and FOREVER, defaults to FOREVER.
	// +kubebuilder:validation:Optional
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Specifies the name of agency. The name is a string of 1 to 64 characters.
	// Changing this will create a new agency.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies an array of one or more roles and projects which are used to grant
	// permissions to agency on project. The structure is documented below.
	// +kubebuilder:validation:Optional
	ProjectRole []ProjectRoleParameters `json:"projectRole,omitempty" tf:"project_role,omitempty"`
}

type ProjectRoleObservation struct {
}

type ProjectRoleParameters struct {

	// Specifies the name of project.
	// +kubebuilder:validation:Required
	Project *string `json:"project" tf:"project,omitempty"`

	// Specifies an array of role names.
	// +kubebuilder:validation:Required
	Roles []*string `json:"roles" tf:"roles,omitempty"`
}

// AgencySpec defines the desired state of Agency
type AgencySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AgencyParameters `json:"forProvider"`
}

// AgencyStatus defines the observed state of Agency.
type AgencyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AgencyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Agency is the Schema for the Agencys API. ""page_title: "flexibleengine_identity_agency_v3"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Agency struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AgencySpec   `json:"spec"`
	Status            AgencyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AgencyList contains a list of Agencys
type AgencyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Agency `json:"items"`
}

// Repository type metadata.
var (
	Agency_Kind             = "Agency"
	Agency_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Agency_Kind}.String()
	Agency_KindAPIVersion   = Agency_Kind + "." + CRDGroupVersion.String()
	Agency_GroupVersionKind = CRDGroupVersion.WithKind(Agency_Kind)
)

func init() {
	SchemeBuilder.Register(&Agency{}, &AgencyList{})
}