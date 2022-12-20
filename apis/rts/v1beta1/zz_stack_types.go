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

type StackObservation struct {

	// List of stack capabilities for stack.
	Capabilities []*string `json:"capabilities,omitempty" tf:"capabilities,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of notification topics for stack.
	NotificationTopics []*string `json:"notificationTopics,omitempty" tf:"notification_topics,omitempty"`

	// A map of outputs from the stack.
	Outputs map[string]*string `json:"outputs,omitempty" tf:"outputs,omitempty"`

	// Specifies the stack status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	StatusReason *string `json:"statusReason,omitempty" tf:"status_reason,omitempty"`
}

type StackParameters struct {

	// Set to true to disable rollback of the stack if stack creation failed.
	// +kubebuilder:validation:Optional
	DisableRollback *bool `json:"disableRollback,omitempty" tf:"disable_rollback,omitempty"`

	// Tthe environment information about the stack.
	// +kubebuilder:validation:Optional
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// Files used in the environment.
	// +kubebuilder:validation:Optional
	Files map[string]*string `json:"files,omitempty" tf:"files,omitempty"`

	// A list of Parameter structures that specify input parameters for the stack.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Structure containing the template body.
	// The template content must use the yaml syntax.
	// +kubebuilder:validation:Optional
	TemplateBody *string `json:"templateBody,omitempty" tf:"template_body,omitempty"`

	// Location of a file containing the template body.
	// +kubebuilder:validation:Optional
	TemplateURL *string `json:"templateUrl,omitempty" tf:"template_url,omitempty"`

	// Specifies the timeout duration.
	// +kubebuilder:validation:Optional
	TimeoutMins *float64 `json:"timeoutMins,omitempty" tf:"timeout_mins,omitempty"`
}

// StackSpec defines the desired state of Stack
type StackSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StackParameters `json:"forProvider"`
}

// StackStatus defines the observed state of Stack.
type StackStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StackObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Stack is the Schema for the Stacks API. ""page_title: "flexibleengine_rts_stack_v1"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Stack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StackSpec   `json:"spec"`
	Status            StackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StackList contains a list of Stacks
type StackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Stack `json:"items"`
}

// Repository type metadata.
var (
	Stack_Kind             = "Stack"
	Stack_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Stack_Kind}.String()
	Stack_KindAPIVersion   = Stack_Kind + "." + CRDGroupVersion.String()
	Stack_GroupVersionKind = CRDGroupVersion.WithKind(Stack_Kind)
)

func init() {
	SchemeBuilder.Register(&Stack{}, &StackList{})
}
