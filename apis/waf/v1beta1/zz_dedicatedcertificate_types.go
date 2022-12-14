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

type DedicatedCertificateObservation struct {

	// Indicates the time when the certificate expires.
	Expiration *string `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// The certificate ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DedicatedCertificateParameters struct {

	// Specifies the certificate content. Changing this creates a new
	// certificate.
	// +kubebuilder:validation:Required
	Certificate *string `json:"certificate" tf:"certificate,omitempty"`

	// Specifies the certificate name. The maximum length is 256 characters. Only digits,
	// letters, underscores(_), and hyphens(-) are allowed.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the private key. Changing this creates a new certificate.
	// +kubebuilder:validation:Required
	PrivateKey *string `json:"privateKey" tf:"private_key,omitempty"`

	// The region in which to create the WAF certificate resource. If omitted, the
	// provider-level region will be used. Changing this setting will push a new certificate.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

// DedicatedCertificateSpec defines the desired state of DedicatedCertificate
type DedicatedCertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DedicatedCertificateParameters `json:"forProvider"`
}

// DedicatedCertificateStatus defines the observed state of DedicatedCertificate.
type DedicatedCertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DedicatedCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedCertificate is the Schema for the DedicatedCertificates API. ""page_title: "flexibleengine_waf_dedicated_certificate"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type DedicatedCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DedicatedCertificateSpec   `json:"spec"`
	Status            DedicatedCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DedicatedCertificateList contains a list of DedicatedCertificates
type DedicatedCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DedicatedCertificate `json:"items"`
}

// Repository type metadata.
var (
	DedicatedCertificate_Kind             = "DedicatedCertificate"
	DedicatedCertificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DedicatedCertificate_Kind}.String()
	DedicatedCertificate_KindAPIVersion   = DedicatedCertificate_Kind + "." + CRDGroupVersion.String()
	DedicatedCertificate_GroupVersionKind = CRDGroupVersion.WithKind(DedicatedCertificate_Kind)
)

func init() {
	SchemeBuilder.Register(&DedicatedCertificate{}, &DedicatedCertificateList{})
}
