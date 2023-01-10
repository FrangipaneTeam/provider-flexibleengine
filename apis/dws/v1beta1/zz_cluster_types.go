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

type ClusterObservation struct {

	// Cluster creation time. The format is
	// ISO8601:YYYY-MM-DDThh:mm:ssZ.
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	// The private network connection information about the cluster.
	// The object structure is documented below.
	Endpoints []EndpointsObservation `json:"endpoints,omitempty" tf:"endpoints,omitempty"`

	// Cluster ID
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The public network connection information about the cluster.
	// The object structure is documented below.
	PublicEndpoints []PublicEndpointsObservation `json:"publicEndpoints,omitempty" tf:"public_endpoints,omitempty"`

	// Public IP address. The object structure is documented below.
	// +kubebuilder:validation:Optional
	PublicIP []PublicIPObservation `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// Cluster status, which can be one of the following: CREATING, AVAILABLE, UNAVAILABLE and CREATION FAILED.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Sub-status of clusters in the AVAILABLE state.
	SubStatus *string `json:"subStatus,omitempty" tf:"sub_status,omitempty"`

	// Cluster management task.
	TaskStatus *string `json:"taskStatus,omitempty" tf:"task_status,omitempty"`

	// Last modification time of a cluster. The format is
	// ISO8601:YYYY-MM-DDThh:mm:ssZ.
	Updated *string `json:"updated,omitempty" tf:"updated,omitempty"`

	// Data warehouse version
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type ClusterParameters struct {

	// AZ in a cluster
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// Cluster name, which must be unique and contains 4 to 64
	// characters, which consist of letters, digits, hyphens (-), or underscores
	// (_) only and must start with a letter.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Node type.
	// +kubebuilder:validation:Required
	NodeType *string `json:"nodeType" tf:"node_type,omitempty"`

	// Number of nodes in a cluster. The value ranges
	// from 3 to 32.
	// +kubebuilder:validation:Required
	NumberOfNode *float64 `json:"numberOfNode" tf:"number_of_node,omitempty"`

	// Service port of a cluster (8000 to 10000). The default
	// value is 8000.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Public IP address. The object structure is documented below.
	// +kubebuilder:validation:Optional
	PublicIP []PublicIPParameters `json:"publicIp,omitempty" tf:"public_ip,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// ID of a security group. The ID is used for
	// configuring cluster network.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.SecurityGroup
	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// Reference to a SecurityGroup in vpc to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRef *v1.Reference `json:"securityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup in vpc to populate securityGroupId.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// Subnet ID, which is used for configuring cluster network.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.VPCSubnet
	// +crossplane:generate:reference:extractor=github.com/FrangipaneTeam/provider-flexibleengine/config/common.IDExtractor()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a VPCSubnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a VPCSubnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// Administrator username for logging in to a data
	// warehouse cluster The administrator username must:
	// +kubebuilder:validation:Required
	UserNameSecretRef v1.SecretKeySelector `json:"userNameSecretRef" tf:"-"`

	// Administrator password for logging in to a data
	// warehouse cluster. A password must conform to the following rules:
	// +kubebuilder:validation:Required
	UserPwdSecretRef v1.SecretKeySelector `json:"userPwdSecretRef" tf:"-"`

	// VPC ID, which is used for configuring cluster network.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

type EndpointsObservation struct {

	// Private network connection information
	ConnectInfo *string `json:"connectInfo,omitempty" tf:"connect_info,omitempty"`

	// JDBC URL. The following is the default format:
	// jdbc:postgresql://< connect_info>/<YOUR_DATABASE_NAME>
	JdbcURL *string `json:"jdbcUrl,omitempty" tf:"jdbc_url,omitempty"`
}

type EndpointsParameters struct {
}

type PublicEndpointsObservation struct {

	// JDBC URL. The following is the default format:
	// jdbc:postgresql://< public_connect_info>/<YOUR_DATABASE_NAME>
	JdbcURL *string `json:"jdbcUrl,omitempty" tf:"jdbc_url,omitempty"`

	// Public network connection information
	PublicConnectInfo *string `json:"publicConnectInfo,omitempty" tf:"public_connect_info,omitempty"`
}

type PublicEndpointsParameters struct {
}

type PublicIPObservation struct {

	// EIP ID
	EIPID *string `json:"eipId,omitempty" tf:"eip_id,omitempty"`

	// Binding type of an EIP. The value can be
	// either of the following: auto_assign, not_use and bind_existing.
	// The default value is not_use.
	PublicBindType *string `json:"publicBindType,omitempty" tf:"public_bind_type,omitempty"`
}

type PublicIPParameters struct {
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API. ""page_title: "flexibleengine_dws_cluster_v1"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
