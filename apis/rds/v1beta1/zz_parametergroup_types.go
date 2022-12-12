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

type ConfigurationParametersObservation struct {

	// The parameter group description. It contains a maximum of 256 characters and cannot contain the following special characters:>!<"&'= the value is left blank by default.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The parameter group name. It contains a maximum of 64 characters.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Indicates whether the parameter is read-only.
	Readonly *bool `json:"readonly,omitempty" tf:"readonly,omitempty"`

	// Indicates whether a restart is required.
	RestartRequired *bool `json:"restartRequired,omitempty" tf:"restart_required,omitempty"`

	// Indicates the parameter type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Indicates the parameter value.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// Indicates the parameter value range.
	ValueRange *string `json:"valueRange,omitempty" tf:"value_range,omitempty"`
}

type ConfigurationParametersParameters struct {
}

type DatastoreObservation struct {
}

type DatastoreParameters struct {

	// The DB engine. Currently, MySQL, PostgreSQL, and Microsoft SQL Server are supported. The value is case-insensitive and can be mysql, postgresql, or sqlserver.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Specifies the database version.
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type ParameterGroupObservation struct {

	// Indicates the parameter configuration defined by users based on the default parameters groups.
	ConfigurationParameters []ConfigurationParametersObservation `json:"configurationParameters,omitempty" tf:"configuration_parameters,omitempty"`

	// ID of the parameter group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ParameterGroupParameters struct {

	// Database object. The database object structure is documented below. Changing this creates a new parameter group.
	// +kubebuilder:validation:Required
	Datastore []DatastoreParameters `json:"datastore" tf:"datastore,omitempty"`

	// The parameter group description. It contains a maximum of 256 characters and cannot contain the following special characters:>!<"&'= the value is left blank by default.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The parameter group name. It contains a maximum of 64 characters.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Parameter group values key/value pairs defined by users based on the default parameter groups.
	// +kubebuilder:validation:Optional
	Values map[string]*string `json:"values,omitempty" tf:"values,omitempty"`
}

// ParameterGroupSpec defines the desired state of ParameterGroup
type ParameterGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ParameterGroupParameters `json:"forProvider"`
}

// ParameterGroupStatus defines the observed state of ParameterGroup.
type ParameterGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ParameterGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ParameterGroup is the Schema for the ParameterGroups API. ""page_title: "flexibleengine_rds_parametergroup_v3"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type ParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ParameterGroupSpec   `json:"spec"`
	Status            ParameterGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ParameterGroupList contains a list of ParameterGroups
type ParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ParameterGroup `json:"items"`
}

// Repository type metadata.
var (
	ParameterGroup_Kind             = "ParameterGroup"
	ParameterGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ParameterGroup_Kind}.String()
	ParameterGroup_KindAPIVersion   = ParameterGroup_Kind + "." + CRDGroupVersion.String()
	ParameterGroup_GroupVersionKind = CRDGroupVersion.WithKind(ParameterGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&ParameterGroup{}, &ParameterGroupList{})
}