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

type SoftwareConfigObservation struct {

	// The id of the software config.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SoftwareConfigParameters struct {

	// The software configuration code.
	// +kubebuilder:validation:Optional
	Config *string `json:"config,omitempty" tf:"config,omitempty"`

	// The namespace that groups this software configuration by when it is delivered to a server.
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// A list of software configuration inputs.
	// +kubebuilder:validation:Optional
	InputValues []map[string]*string `json:"inputValues,omitempty" tf:"input_values,omitempty"`

	// The name of the software configuration.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The software configuration options.
	// +kubebuilder:validation:Optional
	Options map[string]*string `json:"options,omitempty" tf:"options,omitempty"`

	// A list of software configuration outputs.
	// +kubebuilder:validation:Optional
	OutputValues []map[string]*string `json:"outputValues,omitempty" tf:"output_values,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// SoftwareConfigSpec defines the desired state of SoftwareConfig
type SoftwareConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SoftwareConfigParameters `json:"forProvider"`
}

// SoftwareConfigStatus defines the observed state of SoftwareConfig.
type SoftwareConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SoftwareConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SoftwareConfig is the Schema for the SoftwareConfigs API. ""page_title: "flexibleengine_rts_software_config_v1"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type SoftwareConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SoftwareConfigSpec   `json:"spec"`
	Status            SoftwareConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SoftwareConfigList contains a list of SoftwareConfigs
type SoftwareConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SoftwareConfig `json:"items"`
}

// Repository type metadata.
var (
	SoftwareConfig_Kind             = "SoftwareConfig"
	SoftwareConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SoftwareConfig_Kind}.String()
	SoftwareConfig_KindAPIVersion   = SoftwareConfig_Kind + "." + CRDGroupVersion.String()
	SoftwareConfig_GroupVersionKind = CRDGroupVersion.WithKind(SoftwareConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&SoftwareConfig{}, &SoftwareConfigList{})
}
