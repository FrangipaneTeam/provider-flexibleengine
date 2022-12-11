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

type ConditionsObservation struct {
}

type ConditionsParameters struct {

	// Specifies the content matching the condition.
	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// Specifies the matched field. The value can be path, user-agent, ip,
	// params, cookie, referer, and header.
	// +kubebuilder:validation:Required
	Field *string `json:"field" tf:"field,omitempty"`

	// Specifies the logic relationship. The value can be contain, not_contain,
	// equal, not_equal, prefix, not_prefix, suffix, and not_suffix.
	// If field is set to ip, logic can only be equal or not_equal.
	// +kubebuilder:validation:Required
	Logic *string `json:"logic" tf:"logic,omitempty"`

	// Specifies the matched subfield.
	// +kubebuilder:validation:Optional
	Subfield *string `json:"subfield,omitempty" tf:"subfield,omitempty"`
}

type RulePreciseProtectionObservation struct {

	// The rule ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RulePreciseProtectionParameters struct {

	// Specifies the protective action after the precise protection rule is matched.
	// The value can be block or pass. Defaults to block.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action,omitempty"`

	// Specifies the condition parameters. The object structure is documented below.
	// +kubebuilder:validation:Required
	Conditions []ConditionsParameters `json:"conditions" tf:"conditions,omitempty"`

	// Specifies the UTC time when the precise protection rule expires.
	// The time must be in "yyyy-MM-dd HH:mm:ss" format.
	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// Specifies the name of a precise protection rule. The maximum length is
	// 256 characters. Only digits, letters, underscores (_), and hyphens (-) are allowed.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the WAF policy ID. Changing this creates a new rule.
	// +crossplane:generate:reference:type=Policy
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Reference to a Policy to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDRef *v1.Reference `json:"policyIdRef,omitempty" tf:"-"`

	// Selector for a Policy to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDSelector *v1.Selector `json:"policyIdSelector,omitempty" tf:"-"`

	// Specifies the priority of a rule being executed. Smaller values correspond to higher priorities.
	// If two rules are assigned with the same priority, the rule added earlier has higher priority, the rule added earlier
	// has higher priority. The value ranges from 0 to 65535.
	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`

	// Specifies the UTC time when the precise protection rule takes effect.
	// The time must be in "yyyy-MM-dd HH:mm:ss" format.
	// If not specified, the rule takes effect immediately.
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

// RulePreciseProtectionSpec defines the desired state of RulePreciseProtection
type RulePreciseProtectionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RulePreciseProtectionParameters `json:"forProvider"`
}

// RulePreciseProtectionStatus defines the observed state of RulePreciseProtection.
type RulePreciseProtectionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RulePreciseProtectionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RulePreciseProtection is the Schema for the RulePreciseProtections API. ""page_title: "flexibleengine_waf_rule_precise_protection"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type RulePreciseProtection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RulePreciseProtectionSpec   `json:"spec"`
	Status            RulePreciseProtectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RulePreciseProtectionList contains a list of RulePreciseProtections
type RulePreciseProtectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RulePreciseProtection `json:"items"`
}

// Repository type metadata.
var (
	RulePreciseProtection_Kind             = "RulePreciseProtection"
	RulePreciseProtection_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RulePreciseProtection_Kind}.String()
	RulePreciseProtection_KindAPIVersion   = RulePreciseProtection_Kind + "." + CRDGroupVersion.String()
	RulePreciseProtection_GroupVersionKind = CRDGroupVersion.WithKind(RulePreciseProtection_Kind)
)

func init() {
	SchemeBuilder.Register(&RulePreciseProtection{}, &RulePreciseProtectionList{})
}