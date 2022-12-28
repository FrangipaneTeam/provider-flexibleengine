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

type DependentPackagesObservation struct {
}

type DependentPackagesParameters struct {

	// Specifies the spark job name.
	// The value contains a maximum of 128 characters.
	// Changing this parameter will submit a new spark job.
	// +kubebuilder:validation:Required
	GroupName *string `json:"groupName" tf:"group_name,omitempty"`

	// +kubebuilder:validation:Required
	Packages []PackagesParameters `json:"packages" tf:"packages,omitempty"`
}

type PackagesObservation struct {
}

type PackagesParameters struct {

	// Specifies the resource name of the package.
	// Changing this parameter will submit a new spark job.
	// +kubebuilder:validation:Required
	PackageName *string `json:"packageName" tf:"package_name,omitempty"`

	// Specifies the resource type of the package.
	// Changing this parameter will submit a new spark job.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type SparkJobObservation struct {

	// Time of the DLI spark job submit.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// ID of the spark job.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The owner of the spark job.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`
}

type SparkJobParameters struct {

	// Specifies the name of the package that is of the JAR or python file type and
	// has been uploaded to the DLI resource management system.
	// The OBS paths are allowed, for example, obs://<bucket name>/<package name>.
	// Changing this parameter will submit a new spark job.
	// +crossplane:generate:reference:type=DLIPackage
	// +kubebuilder:validation:Optional
	AppName *string `json:"appName,omitempty" tf:"app_name,omitempty"`

	// Reference to a DLIPackage to populate appName.
	// +kubebuilder:validation:Optional
	AppNameRef *v1.Reference `json:"appNameRef,omitempty" tf:"-"`

	// Selector for a DLIPackage to populate appName.
	// +kubebuilder:validation:Optional
	AppNameSelector *v1.Selector `json:"appNameSelector,omitempty" tf:"-"`

	// Specifies the input parameters of the main class.
	// Changing this parameter will submit a new spark job.
	// +kubebuilder:validation:Optional
	AppParameters *string `json:"appParameters,omitempty" tf:"app_parameters,omitempty"`

	// Specifies the configuration items of the DLI spark.
	// Please following the document of Spark configurations for
	// this argument. If you want to enable the access metadata of DLI spark in Flexibleengine, please set
	// spark.dli.metaAccess.enable to true. Changing this parameter will submit a new spark job.
	// +kubebuilder:validation:Optional
	Configurations map[string]*string `json:"configurations,omitempty" tf:"configurations,omitempty"`

	// Specifies a list of package resource objects.
	// The object structure is documented below.
	// Changing this parameter will submit a new spark job.
	// +kubebuilder:validation:Optional
	DependentPackages []DependentPackagesParameters `json:"dependentPackages,omitempty" tf:"dependent_packages,omitempty"`

	// Specifies the number of CPU cores of the Spark application driver.
	// The default value of this value corresponds to the configuration of the selected specification.
	// If you set this value instead of the default value, specification will be invalid.
	// Changing this parameter will submit a new spark job.
	// +kubebuilder:validation:Optional
	DriverCores *float64 `json:"driverCores,omitempty" tf:"driver_cores,omitempty"`

	// Specifies the driver memory of the spark application.
	// The default value of this value corresponds to the configuration of the selected specification.
	// If you set this value instead of the default value, specification will be invalid.
	// Changing this parameter will submit a new spark job.
	// +kubebuilder:validation:Optional
	DriverMemory *string `json:"driverMemory,omitempty" tf:"driver_memory,omitempty"`

	// Specifies the number of CPU cores of each executor in the Spark
	// application. The default value of this value corresponds to the configuration of the selected specification.
	// If you set this value instead of the default value, specification will be invalid.
	// Changing this parameter will submit a new spark job.
	// +kubebuilder:validation:Optional
	ExecutorCores *float64 `json:"executorCores,omitempty" tf:"executor_cores,omitempty"`

	// Specifies the executor memory of the spark application.
	// application. The default value of this value corresponds to the configuration of the selected specification.
	// If you set this value instead of the default value, specification will be invalid.
	// Changing this parameter will submit a new spark job.
	// +kubebuilder:validation:Optional
	ExecutorMemory *string `json:"executorMemory,omitempty" tf:"executor_memory,omitempty"`

	// Specifies the number of executors in a spark application.
	// The default value of this value corresponds to the configuration of the selected specification.
	// If you set this value instead of the default value, specification will be invalid.
	// Changing this parameter will submit a new spark job.
	// +kubebuilder:validation:Optional
	Executors *float64 `json:"executors,omitempty" tf:"executors,omitempty"`

	// Specifies a list of the other dependencies name which has been uploaded to the
	// DLI resource management system. The OBS paths are allowed, for example, obs://<bucket name>/<dependent files>.
	// Changing this parameter will submit a new spark job.
	// +crossplane:generate:reference:type=DLIPackage
	// +kubebuilder:validation:Optional
	Files []*string `json:"files,omitempty" tf:"files,omitempty"`

	// References to DLIPackage to populate files.
	// +kubebuilder:validation:Optional
	FilesRefs []v1.Reference `json:"filesRefs,omitempty" tf:"-"`

	// Selector for a list of DLIPackage to populate files.
	// +kubebuilder:validation:Optional
	FilesSelector *v1.Selector `json:"filesSelector,omitempty" tf:"-"`

	// Specifies a list of the jar package name which has been uploaded to the DLI
	// resource management system. The OBS paths are allowed, for example, obs://<bucket name>/<package name>.
	// Changing this parameter will submit a new spark job.
	// +crossplane:generate:reference:type=DLIPackage
	// +kubebuilder:validation:Optional
	Jars []*string `json:"jars,omitempty" tf:"jars,omitempty"`

	// References to DLIPackage to populate jars.
	// +kubebuilder:validation:Optional
	JarsRefs []v1.Reference `json:"jarsRefs,omitempty" tf:"-"`

	// Selector for a list of DLIPackage to populate jars.
	// +kubebuilder:validation:Optional
	JarsSelector *v1.Selector `json:"jarsSelector,omitempty" tf:"-"`

	// Specifies the main class of the spark job.
	// Required if the app_name is the JAR type.
	// Changing this parameter will submit a new spark job.
	// +kubebuilder:validation:Optional
	MainClass *string `json:"mainClass,omitempty" tf:"main_class,omitempty"`

	// Specifies the maximum retry times.
	// The default value of this value corresponds to the configuration of the selected specification.
	// If you set this value instead of the default value, specification will be invalid.
	// Changing this parameter will submit a new spark job.
	// +kubebuilder:validation:Optional
	MaxRetries *float64 `json:"maxRetries,omitempty" tf:"max_retries,omitempty"`

	// Specifies a list of modules that depend on system resources.
	// The dependent modules and corresponding services are as follows.
	// Changing this parameter will submit a new spark job.
	// +kubebuilder:validation:Optional
	Modules []*string `json:"modules,omitempty" tf:"modules,omitempty"`

	// Specifies the spark job name.
	// The value contains a maximum of 128 characters.
	// Changing this parameter will submit a new spark job.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies a list of the python file name which has been uploaded to the
	// DLI resource management system. The OBS paths are allowed, for example, obs://<bucket name>/<python file name>.
	// Changing this parameter will submit a new spark job.
	// +crossplane:generate:reference:type=DLIPackage
	// +kubebuilder:validation:Optional
	PythonFiles []*string `json:"pythonFiles,omitempty" tf:"python_files,omitempty"`

	// References to DLIPackage to populate pythonFiles.
	// +kubebuilder:validation:Optional
	PythonFilesRefs []v1.Reference `json:"pythonFilesRefs,omitempty" tf:"-"`

	// Selector for a list of DLIPackage to populate pythonFiles.
	// +kubebuilder:validation:Optional
	PythonFilesSelector *v1.Selector `json:"pythonFilesSelector,omitempty" tf:"-"`

	// Specifies the DLI queue name.
	// Changing this parameter will submit a new spark job.
	// +crossplane:generate:reference:type=Queue
	// +kubebuilder:validation:Optional
	QueueName *string `json:"queueName,omitempty" tf:"queue_name,omitempty"`

	// Reference to a Queue to populate queueName.
	// +kubebuilder:validation:Optional
	QueueNameRef *v1.Reference `json:"queueNameRef,omitempty" tf:"-"`

	// Selector for a Queue to populate queueName.
	// +kubebuilder:validation:Optional
	QueueNameSelector *v1.Selector `json:"queueNameSelector,omitempty" tf:"-"`

	// Specifies the region in which to submit a spark job.
	// If omitted, the provider-level region will be used.
	// Changing this parameter will submit a new spark job.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the compute resource type for spark application.
	// The available types and related specifications are as follows, default to minimum configuration (type A).
	// Changing this parameter will submit a new spark job.
	// +kubebuilder:validation:Optional
	Specification *string `json:"specification,omitempty" tf:"specification,omitempty"`
}

// SparkJobSpec defines the desired state of SparkJob
type SparkJobSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SparkJobParameters `json:"forProvider"`
}

// SparkJobStatus defines the observed state of SparkJob.
type SparkJobStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SparkJobObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SparkJob is the Schema for the SparkJobs API. ""page_title: "flexibleengine_dli_spark_job"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type SparkJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SparkJobSpec   `json:"spec"`
	Status            SparkJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SparkJobList contains a list of SparkJobs
type SparkJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SparkJob `json:"items"`
}

// Repository type metadata.
var (
	SparkJob_Kind             = "SparkJob"
	SparkJob_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SparkJob_Kind}.String()
	SparkJob_KindAPIVersion   = SparkJob_Kind + "." + CRDGroupVersion.String()
	SparkJob_GroupVersionKind = CRDGroupVersion.WithKind(SparkJob_Kind)
)

func init() {
	SchemeBuilder.Register(&SparkJob{}, &SparkJobList{})
}
