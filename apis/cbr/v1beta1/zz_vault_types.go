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

type ResourcesObservation struct {
}

type ResourcesParameters struct {

	// +kubebuilder:validation:Optional
	Excludes []*string `json:"excludes,omitempty" tf:"excludes,omitempty"`

	// Specifies the array of disk or SFS file system IDs which will be included in the backup.
	// Only disk and turbo vault support this parameter.
	// +kubebuilder:validation:Optional
	Includes []*string `json:"includes,omitempty" tf:"includes,omitempty"`

	// Specifies the ID of the ECS instance to be backed up.
	// +kubebuilder:validation:Optional
	ServerID *string `json:"serverId,omitempty" tf:"server_id,omitempty"`
}

type VaultObservation struct {

	// The allocated capacity of the vault, in GB.
	Allocated *float64 `json:"allocated,omitempty" tf:"allocated,omitempty"`

	// A resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The specification code.
	SpecCode *string `json:"specCode,omitempty" tf:"spec_code,omitempty"`

	// The vault status.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The name of the bucket for the vault.
	Storage *string `json:"storage,omitempty" tf:"storage,omitempty"`

	// The used capacity, in GB.
	Used *float64 `json:"used,omitempty" tf:"used,omitempty"`
}

type VaultParameters struct {

	// +kubebuilder:validation:Optional
	AutoBind *bool `json:"autoBind,omitempty" tf:"auto_bind,omitempty"`

	// Specifies to enable auto capacity expansion for the backup protection type vault.
	// Defaults to false.
	// +kubebuilder:validation:Optional
	AutoExpand *bool `json:"autoExpand,omitempty" tf:"auto_expand,omitempty"`

	// +kubebuilder:validation:Optional
	AutoPay *string `json:"autoPay,omitempty" tf:"auto_pay,omitempty"`

	// +kubebuilder:validation:Optional
	AutoRenew *string `json:"autoRenew,omitempty" tf:"auto_renew,omitempty"`

	// +kubebuilder:validation:Optional
	BindRules map[string]*string `json:"bindRules,omitempty" tf:"bind_rules,omitempty"`

	// +kubebuilder:validation:Optional
	ChargingMode *string `json:"chargingMode,omitempty" tf:"charging_mode,omitempty"`

	// Specifies the backup specifications.
	// Currently, Only server type vaults support application consistent and only crash_consistent is valid.
	// Changing this will create a new vault.
	// +kubebuilder:validation:Optional
	ConsistentLevel *string `json:"consistentLevel,omitempty" tf:"consistent_level,omitempty"`

	// A resource ID in UUID format.
	// +kubebuilder:validation:Optional
	EnterpriseProjectID *string `json:"enterpriseProjectId,omitempty" tf:"enterprise_project_id,omitempty"`

	// Specifies a unique name of the CBR vault. This parameter can contain a maximum of 64
	// characters, which may consist of letters, digits, underscores(_) and hyphens (-).
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Period *float64 `json:"period,omitempty" tf:"period,omitempty"`

	// +kubebuilder:validation:Optional
	PeriodUnit *string `json:"periodUnit,omitempty" tf:"period_unit,omitempty"`

	// Specifies a policy to associate with the CBR vault.
	// policy_id cannot be used with the vault of replicate protection type.
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Specifies the protection type of the CBR vault.
	// The valid values are backup and replication. Vaults of type disk don't support replication.
	// Changing this will create a new vault.
	// +kubebuilder:validation:Required
	ProtectionType *string `json:"protectionType" tf:"protection_type,omitempty"`

	// Specifies the region in which to create the CBR vault. If omitted, the
	// provider-level region will be used. Changing this will create a new vault.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies an array of one or more resources to attach to the CBR vault.
	// The object structure is documented below.
	// +kubebuilder:validation:Optional
	Resources []ResourcesParameters `json:"resources,omitempty" tf:"resources,omitempty"`

	// Specifies the vault sapacity, in GB. The valid value range is 1 to 10,485,760.
	// +kubebuilder:validation:Required
	Size *float64 `json:"size" tf:"size,omitempty"`

	// Specifies the key/value pairs to associate with the CBR vault.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the object type of the CBR vault.
	// Changing this will create a new vault. Vaild values are as follows:
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// VaultSpec defines the desired state of Vault
type VaultSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VaultParameters `json:"forProvider"`
}

// VaultStatus defines the observed state of Vault.
type VaultStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VaultObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Vault is the Schema for the Vaults API. ""page_title: "flexibleengine_cbr_vault"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Vault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VaultSpec   `json:"spec"`
	Status            VaultStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultList contains a list of Vaults
type VaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vault `json:"items"`
}

// Repository type metadata.
var (
	Vault_Kind             = "Vault"
	Vault_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Vault_Kind}.String()
	Vault_KindAPIVersion   = Vault_Kind + "." + CRDGroupVersion.String()
	Vault_GroupVersionKind = CRDGroupVersion.WithKind(Vault_Kind)
)

func init() {
	SchemeBuilder.Register(&Vault{}, &VaultList{})
}
