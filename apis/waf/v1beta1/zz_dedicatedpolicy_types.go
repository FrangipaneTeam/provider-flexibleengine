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

type DedicatedPolicyObservation struct {

	// The detection mode in Precise Protection.
	FullDetection *bool `json:"fullDetection,omitempty" tf:"full_detection,omitempty"`

	// The policy ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The protection switches. The options object structure is documented below.
	Options []OptionsObservation `json:"options,omitempty" tf:"options,omitempty"`
}

type DedicatedPolicyParameters struct {

	// Specifies the protection level. Defaults to 2. Valid values are:
	// +kubebuilder:validation:Optional
	Level *float64 `json:"level,omitempty" tf:"level,omitempty"`

	// Specifies the policy name. The maximum length is 256 characters. Only digits, letters,
	// underscores(_), and hyphens(-) are allowed.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the protective action after a rule is matched. Defaults to log.
	// Valid values are:
	// +kubebuilder:validation:Optional
	ProtectionMode *string `json:"protectionMode,omitempty" tf:"protection_mode,omitempty"`

	// The region in which to create the WAF policy resource. If omitted, the
	// provider-level region will be used. Changing this setting will push a new certificate.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type OptionsObservation struct {

	// Indicates whether Basic Web Protection is enabled.
	BasicWebProtection *bool `json:"basicWebProtection,omitempty" tf:"basic_web_protection,omitempty"`

	// Indicates whether Blacklist and Whitelist is enabled.
	Blacklist *bool `json:"blacklist,omitempty" tf:"blacklist,omitempty"`

	// Indicates whether CC Attack Protection is enabled.
	CcAttackProtection *bool `json:"ccAttackProtection,omitempty" tf:"cc_attack_protection,omitempty"`

	// Indicates whether the master crawler detection switch in Basic Web Protection is enabled.
	Crawler *bool `json:"crawler,omitempty" tf:"crawler,omitempty"`

	// Indicates whether the Search Engine switch in Basic Web Protection is enabled.
	CrawlerEngine *bool `json:"crawlerEngine,omitempty" tf:"crawler_engine,omitempty"`

	// Indicates whether detection of other crawlers in Basic Web Protection is enabled.
	CrawlerOther *bool `json:"crawlerOther,omitempty" tf:"crawler_other,omitempty"`

	// Indicates whether the Scanner switch in Basic Web Protection is enabled.
	CrawlerScanner *bool `json:"crawlerScanner,omitempty" tf:"crawler_scanner,omitempty"`

	// Indicates whether the Script Tool switch in Basic Web Protection is enabled.
	CrawlerScript *bool `json:"crawlerScript,omitempty" tf:"crawler_script,omitempty"`

	// Indicates whether Data Masking is enabled.
	DataMasking *bool `json:"dataMasking,omitempty" tf:"data_masking,omitempty"`

	// Indicates whether False Alarm Masking is enabled.
	FalseAlarmMasking *bool `json:"falseAlarmMasking,omitempty" tf:"false_alarm_masking,omitempty"`

	// Indicates whether General Check in Basic Web Protection is enabled.
	GeneralCheck *bool `json:"generalCheck,omitempty" tf:"general_check,omitempty"`

	// Indicates whether Precise Protection is enabled.
	PreciseProtection *bool `json:"preciseProtection,omitempty" tf:"precise_protection,omitempty"`

	// Indicates whether Web Tamper Protection is enabled.
	WebTamperProtection *bool `json:"webTamperProtection,omitempty" tf:"web_tamper_protection,omitempty"`

	// Indicates whether webshell detection in Basic Web Protection is enabled.
	Webshell *bool `json:"webshell,omitempty" tf:"webshell,omitempty"`
}

type OptionsParameters struct {
}

// DedicatedPolicySpec defines the desired state of DedicatedPolicy
type DedicatedPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DedicatedPolicyParameters `json:"forProvider"`
}

// DedicatedPolicyStatus defines the observed state of DedicatedPolicy.
type DedicatedPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DedicatedPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedPolicy is the Schema for the DedicatedPolicys API. ""page_title: "flexibleengine_waf_dedicated_policy"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type DedicatedPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DedicatedPolicySpec   `json:"spec"`
	Status            DedicatedPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedPolicyList contains a list of DedicatedPolicys
type DedicatedPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DedicatedPolicy `json:"items"`
}

// Repository type metadata.
var (
	DedicatedPolicy_Kind             = "DedicatedPolicy"
	DedicatedPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DedicatedPolicy_Kind}.String()
	DedicatedPolicy_KindAPIVersion   = DedicatedPolicy_Kind + "." + CRDGroupVersion.String()
	DedicatedPolicy_GroupVersionKind = CRDGroupVersion.WithKind(DedicatedPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&DedicatedPolicy{}, &DedicatedPolicyList{})
}
