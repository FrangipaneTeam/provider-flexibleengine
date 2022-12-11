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

type MonitorObservation struct {

	// The unique ID for the monitor.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MonitorParameters struct {

	// Specifies the Domain Name of the Monitor.
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// Specifies the time, in seconds, between sending probes to members.
	// +kubebuilder:validation:Required
	Interval *float64 `json:"interval" tf:"interval,omitempty"`

	// Specifies the number of permissible ping failures before changing the member's
	// status to INACTIVE. Must be a number between 1 and 10.
	// +kubebuilder:validation:Required
	MaxRetries *float64 `json:"maxRetries" tf:"max_retries,omitempty"`

	// Specifies the id of the pool that this monitor will be assigned to.
	// +crossplane:generate:reference:type=Pool
	// +kubebuilder:validation:Optional
	PoolID *string `json:"poolId,omitempty" tf:"pool_id,omitempty"`

	// Reference to a Pool to populate poolId.
	// +kubebuilder:validation:Optional
	PoolIDRef *v1.Reference `json:"poolIdRef,omitempty" tf:"-"`

	// Selector for a Pool to populate poolId.
	// +kubebuilder:validation:Optional
	PoolIDSelector *v1.Selector `json:"poolIdSelector,omitempty" tf:"-"`

	// Specifies the health check port. The value ranges from 1 to 65535.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Specifies the type of probe, which is TCP, HTTP, or HTTPS, that is
	// sent by the load balancer to verify the member state. Changing this creates a new monitor.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Specifies the region in which to create the ELB monitor resource.
	// If omitted, the provider-level region will be used. Changing this creates a new monitor.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the Maximum number of seconds for a monitor to wait for a ping reply before
	// it times out. The value must be less than the delay value.
	// +kubebuilder:validation:Required
	Timeout *float64 `json:"timeout" tf:"timeout,omitempty"`

	// Specifies the required for HTTP(S) types. URI path that will be accessed if monitor
	// type is HTTP or HTTPS.
	// +kubebuilder:validation:Optional
	URLPath *string `json:"urlPath,omitempty" tf:"url_path,omitempty"`
}

// MonitorSpec defines the desired state of Monitor
type MonitorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorParameters `json:"forProvider"`
}

// MonitorStatus defines the observed state of Monitor.
type MonitorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Monitor is the Schema for the Monitors API. ""page_title: "flexibleengine_lb_monitor_v3"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Monitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorSpec   `json:"spec"`
	Status            MonitorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorList contains a list of Monitors
type MonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Monitor `json:"items"`
}

// Repository type metadata.
var (
	Monitor_Kind             = "Monitor"
	Monitor_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Monitor_Kind}.String()
	Monitor_KindAPIVersion   = Monitor_Kind + "." + CRDGroupVersion.String()
	Monitor_GroupVersionKind = CRDGroupVersion.WithKind(Monitor_Kind)
)

func init() {
	SchemeBuilder.Register(&Monitor{}, &MonitorList{})
}