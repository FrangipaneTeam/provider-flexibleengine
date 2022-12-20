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

type BackupObservation struct {

	// Specifies backup record ID.
	BackupRecordID *string `json:"backupRecordId,omitempty" tf:"backup_record_id,omitempty"`

	// Specifies Cinder backup ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// It specifies the status of backup.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// block supports the following arguments:
	VMMetadata []VMMetadataObservation `json:"vmMetadata,omitempty" tf:"vm_metadata,omitempty"`

	// block supports the following arguments:
	VolumeBackups []VolumeBackupsObservation `json:"volumeBackups,omitempty" tf:"volume_backups,omitempty"`
}

type BackupParameters struct {

	// Name for the backup. The value consists of 1 to 255 characters and can contain
	// only letters, digits, underscores (_), and hyphens (-). Changing backup_name creates a new backup.
	// +kubebuilder:validation:Optional
	BackupName *string `json:"backupName,omitempty" tf:"backup_name,omitempty"`

	// Backup description. The value consists of 0 to 255 characters and must not contain
	// a greater-than sign (>) or less-than sign (<). Changing description creates a new backup.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// ID of the target to which the backup is restored. Changing this creates a new backup.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/ecs/v1beta1.Instance
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceId,omitempty" tf:"resource_id,omitempty"`

	// Reference to a Instance in ecs to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDRef *v1.Reference `json:"resourceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in ecs to populate resourceId.
	// +kubebuilder:validation:Optional
	ResourceIDSelector *v1.Selector `json:"resourceIdSelector,omitempty" tf:"-"`

	// Type of the target to which the backup is restored. The default value is
	// OS::Nova::Server for an ECS. Changing this creates a new backup.
	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`
}

type VMMetadataObservation struct {

	// Specifies ECS type.
	CloudServiceType *string `json:"cloudServiceType,omitempty" tf:"cloud_service_type,omitempty"`

	// Shows system disk size corresponding to the ECS specifications.
	Disk *float64 `json:"disk,omitempty" tf:"disk,omitempty"`

	// Specifies elastic IP address of the ECS.
	EIP *string `json:"eip,omitempty" tf:"eip,omitempty"`

	// Specifies image type.
	ImageType *string `json:"imageType,omitempty" tf:"image_type,omitempty"`

	// Name of backup data.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// It specifies internal IP address of the ECS.
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// Specifies memory size of the ECS, in MB.
	RAM *float64 `json:"ram,omitempty" tf:"ram,omitempty"`

	// Specifies CPU cores corresponding to the ECS.
	Vcpus *float64 `json:"vcpus,omitempty" tf:"vcpus,omitempty"`
}

type VMMetadataParameters struct {
}

type VolumeBackupsObservation struct {

	// Specifies the average speed.
	AverageSpeed *float64 `json:"averageSpeed,omitempty" tf:"average_speed,omitempty"`

	// Specifies whether the disk is bootable.
	Bootable *bool `json:"bootable,omitempty" tf:"bootable,omitempty"`

	// Specifies Cinder backup ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// It specifies backup. The default value is backup.
	ImageType *string `json:"imageType,omitempty" tf:"image_type,omitempty"`

	// Shows whether incremental backup is used.
	Incremental *bool `json:"incremental,omitempty" tf:"incremental,omitempty"`

	// It gives EVS disk backup name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies accumulated size (MB) of backups.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// ID of snapshot.
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// It specifies source volume ID.
	SourceVolumeID *string `json:"sourceVolumeId,omitempty" tf:"source_volume_id,omitempty"`

	// Specifies source volume name.
	SourceVolumeName *string `json:"sourceVolumeName,omitempty" tf:"source_volume_name,omitempty"`

	// Shows source volume size in GB.
	SourceVolumeSize *float64 `json:"sourceVolumeSize,omitempty" tf:"source_volume_size,omitempty"`

	// Specifies space saving rate.
	SpaceSavingRatio *float64 `json:"spaceSavingRatio,omitempty" tf:"space_saving_ratio,omitempty"`

	// Status of backup Volume.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type VolumeBackupsParameters struct {
}

// BackupSpec defines the desired state of Backup
type BackupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupParameters `json:"forProvider"`
}

// BackupStatus defines the observed state of Backup.
type BackupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Backup is the Schema for the Backups API. ""page_title: "flexibleengine_csbs_backup_v1"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Backup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupSpec   `json:"spec"`
	Status            BackupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupList contains a list of Backups
type BackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Backup `json:"items"`
}

// Repository type metadata.
var (
	Backup_Kind             = "Backup"
	Backup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Backup_Kind}.String()
	Backup_KindAPIVersion   = Backup_Kind + "." + CRDGroupVersion.String()
	Backup_GroupVersionKind = CRDGroupVersion.WithKind(Backup_Kind)
)

func init() {
	SchemeBuilder.Register(&Backup{}, &BackupList{})
}
