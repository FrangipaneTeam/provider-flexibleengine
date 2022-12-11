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

type PolicyObservation struct {

	// The policy ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PolicyParameters struct {

	// An array of domain IDs.
	// +kubebuilder:validation:Optional
	Domains []*string `json:"domains,omitempty" tf:"domains,omitempty"`

	// Specifies the detection mode in Precise Protection. Valid values are:
	// +kubebuilder:validation:Optional
	FullDetection *bool `json:"fullDetection,omitempty" tf:"full_detection,omitempty"`

	// Specifies the protection level. Valid values are:
	// +kubebuilder:validation:Optional
	Level *float64 `json:"level,omitempty" tf:"level,omitempty"`

	// Specifies the policy name. The maximum length is 256 characters.
	// Only digits, letters, underscores(_), and hyphens(-) are allowed.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the protective action after a rule is matched. Valid values are:
	// +kubebuilder:validation:Optional
	ProtectionMode *string `json:"protectionMode,omitempty" tf:"protection_mode,omitempty"`

	// Specifies the protection switches. The object structure is documented below.
	// +kubebuilder:validation:Optional
	ProtectionStatus []ProtectionStatusParameters `json:"protectionStatus,omitempty" tf:"protection_status,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type ProtectionStatusObservation struct {
}

type ProtectionStatusParameters struct {

	// Specifies whether Basic Web Protection is enabled.
	// +kubebuilder:validation:Optional
	BasicWebProtection *bool `json:"basicWebProtection,omitempty" tf:"basic_web_protection,omitempty"`

	// Specifies whether Blacklist and Whitelist is enabled.
	// +kubebuilder:validation:Optional
	Blacklist *bool `json:"blacklist,omitempty" tf:"blacklist,omitempty"`

	// Specifies whether CC Attack Protection is enabled.
	// +kubebuilder:validation:Optional
	CcProtection *bool `json:"ccProtection,omitempty" tf:"cc_protection,omitempty"`

	// Specifies whether the Search Engine switch in Basic Web Protection is enabled.
	// +kubebuilder:validation:Optional
	CrawlerEngine *bool `json:"crawlerEngine,omitempty" tf:"crawler_engine,omitempty"`

	// Specifies whether detection of other crawlers in Basic Web Protection is enabled.
	// +kubebuilder:validation:Optional
	CrawlerOther *bool `json:"crawlerOther,omitempty" tf:"crawler_other,omitempty"`

	// Specifies whether the Scanner switch in Basic Web Protection is enabled.
	// +kubebuilder:validation:Optional
	CrawlerScanner *bool `json:"crawlerScanner,omitempty" tf:"crawler_scanner,omitempty"`

	// Specifies whether the Script Tool switch in Basic Web Protection is enabled.
	// +kubebuilder:validation:Optional
	CrawlerScript *bool `json:"crawlerScript,omitempty" tf:"crawler_script,omitempty"`

	// Specifies whether Data Masking is enabled.
	// +kubebuilder:validation:Optional
	DataMasking *bool `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Specifies whether False Alarm Masking is enabled.
	// +kubebuilder:validation:Optional
	FalseAlarmMasking *bool `json:"falseAlarmMasking,omitempty" tf:"false_alarm_masking,omitempty"`

	// Specifies whether General Check in Basic Web Protection is enabled.
	// +kubebuilder:validation:Optional
	GeneralCheck *bool `json:"generalCheck,omitempty" tf:"general_check,omitempty"`

	// Specifies whether Precise Protection is enabled.
	// +kubebuilder:validation:Optional
	PreciseProtection *bool `json:"preciseProtection,omitempty" tf:"precise_protection,omitempty"`

	// Specifies whether Web Tamper Protection is enabled.
	// +kubebuilder:validation:Optional
	WebTamperProtection *bool `json:"webTamperProtection,omitempty" tf:"web_tamper_protection,omitempty"`

	// Specifies whether webshell detection in Basic Web Protection is enabled.
	// +kubebuilder:validation:Optional
	Webshell *bool `json:"webshell,omitempty" tf:"webshell,omitempty"`
}

// PolicySpec defines the desired state of Policy
type PolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PolicyParameters `json:"forProvider"`
}

// PolicyStatus defines the observed state of Policy.
type PolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Policy is the Schema for the Policys API. ""page_title: "flexibleengine_waf_policy"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicySpec   `json:"spec"`
	Status            PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PolicyList contains a list of Policys
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

// Repository type metadata.
var (
	Policy_Kind             = "Policy"
	Policy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Policy_Kind}.String()
	Policy_KindAPIVersion   = Policy_Kind + "." + CRDGroupVersion.String()
	Policy_GroupVersionKind = CRDGroupVersion.WithKind(Policy_Kind)
)

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}