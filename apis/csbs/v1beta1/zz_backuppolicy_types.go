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

type BackupPolicyObservation struct {

	// Backup Policy ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// block supports the following arguments:
	// +kubebuilder:validation:Required
	ScheduledOperation []ScheduledOperationObservation `json:"scheduledOperation,omitempty" tf:"scheduled_operation,omitempty"`

	// Status of Backup Policy.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type BackupPolicyParameters struct {

	// General backup policy parameters, which are blank by default.
	// +kubebuilder:validation:Optional
	Common map[string]*string `json:"common,omitempty" tf:"common,omitempty"`

	// Backup policy description. The value consists of 0 to 255 characters and must not contain a greater-than sign (>) or less-than sign (<).
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the name of backup policy. The value consists of 1 to 255 characters and can contain only letters, digits, underscores (_), and hyphens (-).
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies backup provider ID. Default value is fc4d5750-22e7-4798-8a46-f48f62c4c1da
	// +kubebuilder:validation:Optional
	ProviderID *string `json:"providerId,omitempty" tf:"provider_id,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// block supports the following arguments:
	// +kubebuilder:validation:Required
	Resource []ResourceParameters `json:"resource" tf:"resource,omitempty"`

	// block supports the following arguments:
	// +kubebuilder:validation:Required
	ScheduledOperation []ScheduledOperationParameters `json:"scheduledOperation" tf:"scheduled_operation,omitempty"`
}

type ResourceObservation struct {
}

type ResourceParameters struct {

	// Specifies the ID of the object to be backed up.
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// Specifies backup object name.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Entity object type of the backup object. If the type is VMs, the value is OS::Nova::Server.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ScheduledOperationObservation struct {

	// Specifies the ID of the object to be backed up.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies Scheduler ID.
	TriggerID *string `json:"triggerId,omitempty" tf:"trigger_id,omitempty"`

	// Specifies Scheduler name.
	TriggerName *string `json:"triggerName,omitempty" tf:"trigger_name,omitempty"`

	// Specifies Scheduler type.
	TriggerType *string `json:"triggerType,omitempty" tf:"trigger_type,omitempty"`
}

type ScheduledOperationParameters struct {

	// Specifies Scheduling period description.The value consists of 0 to 255 characters and must not contain a greater-than sign (>) or less-than sign (<).
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies whether the scheduling period is enabled. Default value is true
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies maximum number of backups that can be automatically created for a backup object.
	// +kubebuilder:validation:Optional
	MaxBackups *float64 `json:"maxBackups,omitempty" tf:"max_backups,omitempty"`

	// Specifies Scheduling period name.The value consists of 1 to 255 characters and can contain only letters, digits, underscores (_), and hyphens (-).
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies Operation type, which can be backup.
	// +kubebuilder:validation:Required
	OperationType *string `json:"operationType" tf:"operation_type,omitempty"`

	// Specifies whether backups are permanently retained.
	// +kubebuilder:validation:Optional
	Permanent *bool `json:"permanent,omitempty" tf:"permanent,omitempty"`

	// Specifies duration of retaining a backup, in days.
	// +kubebuilder:validation:Optional
	RetentionDurationDays *float64 `json:"retentionDurationDays,omitempty" tf:"retention_duration_days,omitempty"`

	// Specifies Scheduling policy of the scheduler.
	// +kubebuilder:validation:Required
	TriggerPattern *string `json:"triggerPattern" tf:"trigger_pattern,omitempty"`
}

// BackupPolicySpec defines the desired state of BackupPolicy
type BackupPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupPolicyParameters `json:"forProvider"`
}

// BackupPolicyStatus defines the observed state of BackupPolicy.
type BackupPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicy is the Schema for the BackupPolicys API. ""page_title: "flexibleengine_csbs_backup_policy_v1"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type BackupPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupPolicySpec   `json:"spec"`
	Status            BackupPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupPolicyList contains a list of BackupPolicys
type BackupPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupPolicy `json:"items"`
}

// Repository type metadata.
var (
	BackupPolicy_Kind             = "BackupPolicy"
	BackupPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupPolicy_Kind}.String()
	BackupPolicy_KindAPIVersion   = BackupPolicy_Kind + "." + CRDGroupVersion.String()
	BackupPolicy_GroupVersionKind = CRDGroupVersion.WithKind(BackupPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupPolicy{}, &BackupPolicyList{})
}
