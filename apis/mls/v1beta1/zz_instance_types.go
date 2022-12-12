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

type InstanceObservation struct {

	// Specifies the ID of the MRS cluster. Changing this creates a new instance.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type InstanceParameters struct {

	// Specifies the agency name. This parameter is mandatory only
	// when you bind an instance to an elastic IP address (EIP). An instance must be
	// bound to an EIP to grant MLS rights to abtain a tenant's token. NOTE: The tenant
	// must create an agency on the Identity and Access Management (IAM) interface in
	// advance. Changing this creates a new instance.
	// +kubebuilder:validation:Optional
	Agency *string `json:"agency,omitempty" tf:"agency,omitempty"`

	// Indicates the creation time in the following format: yyyy-mm-dd Thh:mm:ssZ.
	// +kubebuilder:validation:Optional
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	// Specifies the instance flavor, only mls.c2.2xlarge.common
	// is supported now. Changing this creates a new instance.
	// +kubebuilder:validation:Required
	Flavor *string `json:"flavor" tf:"flavor,omitempty"`

	// Indicates the URL for accessing the instance. Only machines in the same
	// VPC and subnet as the instance can access the URL.
	// +kubebuilder:validation:Optional
	InnerEndpoint *string `json:"innerEndpoint,omitempty" tf:"inner_endpoint,omitempty"`

	// Specifies the MRS cluster information which the instance
	// is associated. The structure is described below. NOTE: The current MRS instance
	// requires an MRS cluster whose version is 1.3.0 and that is configured with the
	// Spark component. MRS clusters whose version is not 1.3.0 or that are not configured
	// with the Spark component cannot be selected. Changing this creates a new instance.
	// +kubebuilder:validation:Required
	MrsCluster []MrsClusterParameters `json:"mrsCluster" tf:"mrs_cluster,omitempty"`

	// Specifies the MLS instance name. The DB instance name of
	// the same type is unique in the same tenant. Changing this creates a new instance.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the instance network information. The structure
	// is described below. Changing this creates a new instance.
	// +kubebuilder:validation:Required
	Network []NetworkParameters `json:"network" tf:"network,omitempty"`

	// Indicates the URL for accessing the instance. The URL can be accessed
	// from the Internet. The URL is created only after the instance is bound to an EIP.
	// +kubebuilder:validation:Optional
	PublicEndpoint *string `json:"publicEndpoint,omitempty" tf:"public_endpoint,omitempty"`

	// The region in which to create the MLS instance. If
	// omitted, the region argument of the provider is used. Changing this
	// creates a new instance.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Indicates the MLS instance status.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Indicates the update time in the following format: yyyy-mm-dd Thh:mm:ssZ.
	// +kubebuilder:validation:Optional
	Updated *string `json:"updated,omitempty" tf:"updated,omitempty"`

	// Specifies MLS Software version, only 1.2.0 is supported
	// now. Changing this creates a new instance.
	// +kubebuilder:validation:Required
	Version *string `json:"version" tf:"version,omitempty"`
}

type MrsClusterObservation struct {
}

type MrsClusterParameters struct {

	// Specifies the ID of the MRS cluster. Changing this creates a new instance.
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// Specifies the MRS cluster username. This parameter is mandatory
	// only when the MRS cluster is in the security mode. Changing this creates a new instance.
	// +kubebuilder:validation:Optional
	UserName *string `json:"userName,omitempty" tf:"user_name,omitempty"`

	// Specifies the password of the MRS cluster user. The password
	// and username work in a pair. Changing this creates a new instance.
	// +kubebuilder:validation:Optional
	UserPasswordSecretRef *v1.SecretKeySelector `json:"userPasswordSecretRef,omitempty" tf:"-"`
}

type NetworkObservation struct {
}

type NetworkParameters struct {

	// Specifies the AZ of the instance.
	// Changing this creates a new instance.
	// +kubebuilder:validation:Required
	AvailableZone *string `json:"availableZone" tf:"available_zone,omitempty"`

	// Specifies the IP address of the instance. The structure is
	// described below. Changing this creates a new instance.
	// +kubebuilder:validation:Required
	PublicIP []PublicIPParameters `json:"publicIp" tf:"public_ip,omitempty"`

	// Specifies the ID of the security group of the instance.
	// Changing this creates a new instance.
	// +kubebuilder:validation:Optional
	SecurityGroup *string `json:"securityGroup,omitempty" tf:"security_group,omitempty"`

	// Specifies the ID of the subnet where the instance resides.
	// Changing this creates a new instance.
	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`

	// Specifies the ID of the virtual private cloud (VPC) where the
	// instance resides. Changing this creates a new instance.
	// +kubebuilder:validation:Required
	VPCID *string `json:"vpcId" tf:"vpc_id,omitempty"`
}

type PublicIPObservation struct {
}

type PublicIPParameters struct {

	// Specifies the bind type. Possible values: auto_assign and
	// not_use. Changing this creates a new instance.
	// +kubebuilder:validation:Required
	BindType *string `json:"bindType" tf:"bind_type,omitempty"`

	// +kubebuilder:validation:Optional
	EIPID *string `json:"eipId,omitempty" tf:"eip_id,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API. ""page_title: "flexibleengine_mls_instance_v1"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}