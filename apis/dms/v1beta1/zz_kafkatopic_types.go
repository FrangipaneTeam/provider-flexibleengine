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

type KafkaTopicObservation struct {

	// The resource ID which equals to the topic name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type KafkaTopicParameters struct {

	// Specifies the aging time in hours. The value ranges from 1 to 720 and defaults to 72.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	AgingTime *float64 `json:"agingTime,omitempty" tf:"aging_time,omitempty"`

	// Specifies the ID of the DMS Kafka instance to which the topic belongs.
	// Changing this creates a new resource.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/ecs/v1beta1.Instance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in ecs to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in ecs to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the name of the topic. The name starts with a letter,
	// consists of 4 to 64 characters, and supports only letters, digits, hyphens (-) and underscores (_).
	// Changing this creates a new resource.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the partition number. The value ranges from 1 to 50 and defaults to 3.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Partitions *float64 `json:"partitions,omitempty" tf:"partitions,omitempty"`

	// The region in which to create the DMS Kafka topic resource.
	// If omitted, the provider-level region will be used. Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the replica number. The value ranges from 1 to 3 and defaults to 3.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	Replicas *float64 `json:"replicas,omitempty" tf:"replicas,omitempty"`

	// Whether or not to enable synchronous flushing.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	SyncFlushing *bool `json:"syncFlushing,omitempty" tf:"sync_flushing,omitempty"`

	// Whether or not to enable synchronous replication.
	// Changing this creates a new resource.
	// +kubebuilder:validation:Optional
	SyncReplication *bool `json:"syncReplication,omitempty" tf:"sync_replication,omitempty"`
}

// KafkaTopicSpec defines the desired state of KafkaTopic
type KafkaTopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KafkaTopicParameters `json:"forProvider"`
}

// KafkaTopicStatus defines the observed state of KafkaTopic.
type KafkaTopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KafkaTopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// KafkaTopic is the Schema for the KafkaTopics API. ""page_title: "flexibleengine_dms_kafka_topic"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type KafkaTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KafkaTopicSpec   `json:"spec"`
	Status            KafkaTopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KafkaTopicList contains a list of KafkaTopics
type KafkaTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaTopic `json:"items"`
}

// Repository type metadata.
var (
	KafkaTopic_Kind             = "KafkaTopic"
	KafkaTopic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KafkaTopic_Kind}.String()
	KafkaTopic_KindAPIVersion   = KafkaTopic_Kind + "." + CRDGroupVersion.String()
	KafkaTopic_GroupVersionKind = CRDGroupVersion.WithKind(KafkaTopic_Kind)
)

func init() {
	SchemeBuilder.Register(&KafkaTopic{}, &KafkaTopicList{})
}
