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

type RuleWebTamperProtectionObservation struct {

	// The rule ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RuleWebTamperProtectionParameters struct {

	// Specifies the domain name. Changing this creates a new rule.
	// +kubebuilder:validation:Required
	Domain *string `json:"domain" tf:"domain,omitempty"`

	// Specifies the URL protected by the web tamper protection rule,
	// excluding a domain name. Changing this creates a new rule.
	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`

	// Specifies the WAF policy ID. Changing this creates a new rule.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/waf/v1beta1.Policy
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Reference to a Policy in waf to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDRef *v1.Reference `json:"policyIdRef,omitempty" tf:"-"`

	// Selector for a Policy in waf to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDSelector *v1.Selector `json:"policyIdSelector,omitempty" tf:"-"`
}

// RuleWebTamperProtectionSpec defines the desired state of RuleWebTamperProtection
type RuleWebTamperProtectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RuleWebTamperProtectionParameters `json:"forProvider"`
}

// RuleWebTamperProtectionStatus defines the observed state of RuleWebTamperProtection.
type RuleWebTamperProtectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RuleWebTamperProtectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RuleWebTamperProtection is the Schema for the RuleWebTamperProtections API. ""page_title: "flexibleengine_waf_rule_web_tamper_protection"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type RuleWebTamperProtection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RuleWebTamperProtectionSpec   `json:"spec"`
	Status            RuleWebTamperProtectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RuleWebTamperProtectionList contains a list of RuleWebTamperProtections
type RuleWebTamperProtectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RuleWebTamperProtection `json:"items"`
}

// Repository type metadata.
var (
	RuleWebTamperProtection_Kind             = "RuleWebTamperProtection"
	RuleWebTamperProtection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RuleWebTamperProtection_Kind}.String()
	RuleWebTamperProtection_KindAPIVersion   = RuleWebTamperProtection_Kind + "." + CRDGroupVersion.String()
	RuleWebTamperProtection_GroupVersionKind = CRDGroupVersion.WithKind(RuleWebTamperProtection_Kind)
)

func init() {
	SchemeBuilder.Register(&RuleWebTamperProtection{}, &RuleWebTamperProtectionList{})
}
