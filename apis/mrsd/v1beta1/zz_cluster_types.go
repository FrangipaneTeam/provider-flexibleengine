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

type AddJobsObservation struct {
}

type AddJobsParameters struct {

	// Key parameter for program execution The parameter
	// is specified by the function of the user's program. MRS is only responsible
	// for loading the parameter. The parameter contains a maximum of 2047 characters,
	// excluding special characters such as ;|&>'<$, and can be empty.
	// +kubebuilder:validation:Optional
	Arguments *string `json:"arguments,omitempty" tf:"arguments,omitempty"`

	// Data import and export import export
	// +kubebuilder:validation:Optional
	FileAction *string `json:"fileAction,omitempty" tf:"file_action,omitempty"`

	// SQL program path This parameter is needed
	// by Spark Script and Hive Script jobs only and must meet the following requirements:
	// Contains a maximum of 1023 characters, excluding special characters such as
	// ;|&><'$. The address cannot be empty or full of spaces. Starts with / or s3a://.
	// Ends with .sql. sql is case-insensitive.
	// +kubebuilder:validation:Optional
	HiveScriptPath *string `json:"hiveScriptPath,omitempty" tf:"hive_script_path,omitempty"`

	// HiveQL statement
	// +kubebuilder:validation:Optional
	Hql *string `json:"hql,omitempty" tf:"hql,omitempty"`

	// Path for inputting data, which must start with / or s3a://.
	// A correct OBS path is required. The parameter contains a maximum of 1023 characters,
	// excluding special characters such as ;|&>'<$, and can be empty.
	// +kubebuilder:validation:Optional
	Input *string `json:"input,omitempty" tf:"input,omitempty"`

	// Path of the .jar file or .sql file for program execution
	// The parameter must meet the following requirements: Contains a maximum of 1023
	// characters, excluding special characters such as ;|&><'$. The address cannot
	// be empty or full of spaces. Starts with / or s3a://. Spark Script must end with
	// .sql; while MapReduce and Spark Jar must end with .jar. sql and jar are case-insensitive.
	// +kubebuilder:validation:Required
	JarPath *string `json:"jarPath" tf:"jar_path,omitempty"`

	// Path for storing job logs that record job running status.
	// This path must start with / or s3a://. A correct OBS path is required. The parameter
	// contains a maximum of 1023 characters, excluding special characters such as
	// ;|&>'<$, and can be empty.
	// +kubebuilder:validation:Optional
	JobLog *string `json:"jobLog,omitempty" tf:"job_log,omitempty"`

	// Job name It contains only 1 to 64 letters, digits,
	// hyphens (-), and underscores (_). NOTE: Identical job names are allowed but
	// not recommended.
	// +kubebuilder:validation:Required
	JobName *string `json:"jobName" tf:"job_name,omitempty"`

	// Job type. 1: MapReduce 2: Spark 3: Hive Script 4: HiveQL
	// (not supported currently) 5: DistCp, importing and exporting data (not supported
	// in this API currently). 6: Spark Script 7: Spark SQL, submitting Spark SQL statements
	// (not supported in this API currently). NOTE: Spark and Hive jobs can be added
	// to only clusters including Spark and Hive components.
	// +kubebuilder:validation:Required
	JobType *float64 `json:"jobType" tf:"job_type,omitempty"`

	// Path for outputting data, which must start with / or
	// s3a://. A correct OBS path is required. If the path does not exist, the system
	// automatically creates it. The parameter contains a maximum of 1023 characters,
	// excluding special characters such as ;|&>'<$, and can be empty.
	// +kubebuilder:validation:Optional
	Output *string `json:"output,omitempty" tf:"output,omitempty"`

	// Whether to delete the cluster after the jobs
	// are complete true: Yes false: No
	// +kubebuilder:validation:Optional
	ShutdownCluster *bool `json:"shutdownCluster,omitempty" tf:"shutdown_cluster,omitempty"`

	// true: A job is submitted when a
	// cluster is created. false: A job is submitted separately. The parameter is set
	// to true in this example.
	// +kubebuilder:validation:Required
	SubmitJobOnceClusterRun *bool `json:"submitJobOnceClusterRun" tf:"submit_job_once_cluster_run,omitempty"`
}

type ClusterObservation struct {

	// Name of an availability zone.
	AvailableZoneName *string `json:"availableZoneName,omitempty" tf:"available_zone_name,omitempty"`

	// Time when charging starts.
	ChargingStartTime *string `json:"chargingStartTime,omitempty" tf:"charging_start_time,omitempty"`

	// Cluster ID.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Cluster status. Valid values include: starting, running, terminated, failed, abnormal,
	// terminating, frozen, scaling-out and scaling-in.
	ClusterState *string `json:"clusterState,omitempty" tf:"cluster_state,omitempty"`

	// Service component list. The object structure is documented below.
	// +kubebuilder:validation:Required
	ComponentList []ComponentListObservation `json:"componentList,omitempty" tf:"component_list,omitempty"`

	// Product ID of a Core node.
	CoreNodeProductID *string `json:"coreNodeProductId,omitempty" tf:"core_node_product_id,omitempty"`

	// Specification ID of a Core node.
	CoreNodeSpecID *string `json:"coreNodeSpecId,omitempty" tf:"core_node_spec_id,omitempty"`

	// Cluster creation time.
	CreateAt *string `json:"createAt,omitempty" tf:"create_at,omitempty"`

	// Deployment ID of a cluster.
	DeploymentID *string `json:"deploymentId,omitempty" tf:"deployment_id,omitempty"`

	// Cluster subscription duration.
	Duration *string `json:"duration,omitempty" tf:"duration,omitempty"`

	// Error information.
	ErrorInfo *string `json:"errorInfo,omitempty" tf:"error_info,omitempty"`

	// Backup external IP address.
	ExternalAlternateIP *string `json:"externalAlternateIp,omitempty" tf:"external_alternate_ip,omitempty"`

	// External IP address.
	ExternalIP *string `json:"externalIp,omitempty" tf:"external_ip,omitempty"`

	// Cluster creation fee, which is automatically calculated.
	Fee *string `json:"fee,omitempty" tf:"fee,omitempty"`

	// Hadoop version.
	HadoopVersion *string `json:"hadoopVersion,omitempty" tf:"hadoop_version,omitempty"`

	// The resource ID in UUID format.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Instance ID.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	InternalIP *string `json:"internalIp,omitempty" tf:"internal_ip,omitempty"`

	// IP address of a Master node.
	MasterNodeIP *string `json:"masterNodeIp,omitempty" tf:"master_node_ip,omitempty"`

	// Product ID of a Master node.
	MasterNodeProductID *string `json:"masterNodeProductId,omitempty" tf:"master_node_product_id,omitempty"`

	// Specification ID of a Master node.
	MasterNodeSpecID *string `json:"masterNodeSpecId,omitempty" tf:"master_node_spec_id,omitempty"`

	// Order ID for creating clusters.
	OrderID *string `json:"orderId,omitempty" tf:"order_id,omitempty"`

	// Primary private IP address.
	PrivateIPFirst *string `json:"privateIpFirst,omitempty" tf:"private_ip_first,omitempty"`

	// Remarks of a cluster.
	Remark *string `json:"remark,omitempty" tf:"remark,omitempty"`

	// Security group ID.
	SecurityGroupsID *string `json:"securityGroupsId,omitempty" tf:"security_groups_id,omitempty"`

	// Standby security group ID.
	SlaveSecurityGroupsID *string `json:"slaveSecurityGroupsId,omitempty" tf:"slave_security_groups_id,omitempty"`

	// Project ID.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// Cluster update time.
	UpdateAt *string `json:"updateAt,omitempty" tf:"update_at,omitempty"`

	// URI address for remote login of the elastic cloud server.
	Vnc *string `json:"vnc,omitempty" tf:"vnc,omitempty"`
}

type ClusterParameters struct {

	// You can submit a job when you create a cluster to save time and use MRS easily.
	// Only one job can be added. The object structure is documented below.
	// +kubebuilder:validation:Optional
	AddJobs []AddJobsParameters `json:"addJobs,omitempty" tf:"add_jobs,omitempty"`

	// ID or Name of an available zone. Obtain the value
	// from Regions and Endpoints.
	// +kubebuilder:validation:Required
	AvailableZoneID *string `json:"availableZoneId" tf:"available_zone_id,omitempty"`

	// +kubebuilder:validation:Optional
	BillingType *float64 `json:"billingType,omitempty" tf:"billing_type,omitempty"`

	// Indicates the password of the MRS Manager
	// administrator.
	// +kubebuilder:validation:Optional
	ClusterAdminSecretSecretRef *v1.SecretKeySelector `json:"clusterAdminSecretSecretRef,omitempty" tf:"-"`

	// Cluster name, which is globally unique and contains
	// only 1 to 64 letters, digits, hyphens (-), and underscores (_).
	// +kubebuilder:validation:Required
	ClusterName *string `json:"clusterName" tf:"cluster_name,omitempty"`

	// Type of clusters. 0: analysis cluster; 1: streaming cluster.
	// The default value is 0.
	// +kubebuilder:validation:Optional
	ClusterType *float64 `json:"clusterType,omitempty" tf:"cluster_type,omitempty"`

	// Version of the clusters. Possible values are as follows:
	// MRS 1.8.9, MRS 2.0.1, MRS 2.1.0 and MRS 3.1.0-LTS.1. The latest version of MRS is used by default.
	// +kubebuilder:validation:Optional
	ClusterVersion *string `json:"clusterVersion,omitempty" tf:"cluster_version,omitempty"`

	// Service component list. The object structure is documented below.
	// +kubebuilder:validation:Required
	ComponentList []ComponentListParameters `json:"componentList" tf:"component_list,omitempty"`

	// Number of Core nodes Value range: 3 to 500. A
	// maximum of 500 Core nodes are supported by default. If more than 500 Core nodes
	// are required, contact technical support engineers or invoke background APIs
	// to modify the database.
	// +kubebuilder:validation:Required
	CoreNodeNum *float64 `json:"coreNodeNum" tf:"core_node_num,omitempty"`

	// Instance specification of a Core node Configuration
	// method of this parameter is identical to that of master_node_size.
	// +kubebuilder:validation:Required
	CoreNodeSize *string `json:"coreNodeSize" tf:"core_node_size,omitempty"`

	// Indicates whether logs are collected when cluster
	// installation fails. 0: not collected 1: collected The default value is 0. If
	// log_collection is set to 1, OBS buckets will be created to collect the MRS logs.
	// These buckets will be charged.
	// +kubebuilder:validation:Optional
	LogCollection *float64 `json:"logCollection,omitempty" tf:"log_collection,omitempty"`

	// Number of Master nodes The value is 2.
	// +kubebuilder:validation:Required
	MasterNodeNum *float64 `json:"masterNodeNum" tf:"master_node_num,omitempty"`

	// Best match based on several years of commissioning
	// experience. MRS supports specifications of hosts, and host specifications are
	// determined by CPUs, memory, and disks space.
	// +kubebuilder:validation:Required
	MasterNodeSize *string `json:"masterNodeSize" tf:"master_node_size,omitempty"`

	// Name of a key pair You can use a key
	// to log in to the Master node in the cluster.
	// +kubebuilder:validation:Required
	NodePublicCertName *string `json:"nodePublicCertName" tf:"node_public_cert_name,omitempty"`

	// Cluster region information. Obtain the value from
	// Regions and Endpoints.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// MRS cluster running mode.
	// +kubebuilder:validation:Required
	SafeMode *float64 `json:"safeMode" tf:"safe_mode,omitempty"`

	// Subnet ID Obtain the subnet ID from the management
	// console as follows: Register an account and log in to the management console.
	// Click Virtual Private Cloud and select Virtual Private Cloud from the left list.
	// On the Virtual Private Cloud page, obtain the subnet ID from the list.
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

	// ID of the VPC where the subnet locates Obtain the VPC
	// ID from the management console as follows: Register an account and log in to
	// the management console. Click Virtual Private Cloud and select Virtual Private
	// Cloud from the left list. On the Virtual Private Cloud page, obtain the VPC
	// ID from the list.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// Data disk storage space of a Core node Users can
	// add disks to expand storage capacity when creating a cluster. There are the
	// following scenarios: Separation of data storage and computing: Data is stored
	// in the OBS system. Costs of clusters are relatively low but computing performance
	// is poor. The clusters can be deleted at any time. It is recommended when data
	// computing is not frequently performed. Integration of data storage and computing:
	// Data is stored in the HDFS system. Costs of clusters are relatively high but
	// computing performance is good. The clusters cannot be deleted in a short term.
	// It is recommended when data computing is frequently performed. Value range:
	// 100 GB to 32000 GB
	// +kubebuilder:validation:Required
	VolumeSize *float64 `json:"volumeSize" tf:"volume_size,omitempty"`

	// Type of disks SATA and SSD are supported. SATA:
	// common I/O SSD: super high-speed I/O
	// +kubebuilder:validation:Required
	VolumeType *string `json:"volumeType" tf:"volume_type,omitempty"`
}

type ComponentListObservation struct {

	// Component description.
	ComponentDesc *string `json:"componentDesc,omitempty" tf:"component_desc,omitempty"`

	// Component ID. For example, component_id of Hadoop is MRS 3.1.0-LTS.1_001, MRS 2.1.0_001,
	// MRS 2.0.1_001, and MRS 1.8.9_001.
	ComponentID *string `json:"componentId,omitempty" tf:"component_id,omitempty"`

	// Component version.
	ComponentVersion *string `json:"componentVersion,omitempty" tf:"component_version,omitempty"`
}

type ComponentListParameters struct {

	// the Component name.
	// +kubebuilder:validation:Required
	ComponentName *string `json:"componentName" tf:"component_name,omitempty"`
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

// Cluster is the Schema for the Clusters API. ""page_title: "flexibleengine_mrs_cluster_v1"
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
