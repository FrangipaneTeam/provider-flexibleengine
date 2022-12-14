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

type S3BucketObjectObservation struct {

	// the key of the resource supplied above
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A unique version ID value for the object, if bucket versioning
	// is enabled.
	VersionID *string `json:"versionId,omitempty" tf:"version_id,omitempty"`
}

type S3BucketObjectParameters struct {

	// The canned ACL to apply.
	// Defaults to "private".
	// +kubebuilder:validation:Optional
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// The name of the bucket to put the file in.
	// +crossplane:generate:reference:type=S3Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a S3Bucket to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a S3Bucket to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// Specifies caching behavior along the request/reply chain Read w3c cache_control
	// for further details.
	// +kubebuilder:validation:Optional
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// The literal content being uploaded to the bucket.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Specifies presentational information for the object. Read wc3 content_disposition
	// for further information.
	// +kubebuilder:validation:Optional
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// Specifies what content encodings have been applied to the object and thus what decoding
	// mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.
	// Read w3c content encoding for further information.
	// +kubebuilder:validation:Optional
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// The language the content is in e.g. en-US or en-GB.
	// +kubebuilder:validation:Optional
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`

	// A standard MIME type describing the format of the object data, e.g. application/octet-stream.
	// All Valid MIME Types are valid for this input.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Used to trigger updates. The only meaningful value is ${md5(file("path/to/file"))}.
	// This attribute is not compatible with kms_key_id.
	// +kubebuilder:validation:Optional
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// The name of the object once it is in the bucket.
	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// Specifies server-side encryption of the object in S3.
	// Valid values are "AES256" and "aws:kms".
	// +kubebuilder:validation:Optional
	ServerSideEncryption *string `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`

	// The path to the source file being uploaded to the bucket.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Specifies a target URL for website redirect.
	// +kubebuilder:validation:Optional
	WebsiteRedirect *string `json:"websiteRedirect,omitempty" tf:"website_redirect,omitempty"`
}

// S3BucketObjectSpec defines the desired state of S3BucketObject
type S3BucketObjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     S3BucketObjectParameters `json:"forProvider"`
}

// S3BucketObjectStatus defines the observed state of S3BucketObject.
type S3BucketObjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        S3BucketObjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketObject is the Schema for the S3BucketObjects API. ""page_title: "flexibleengine_s3_bucket_object"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type S3BucketObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3BucketObjectSpec   `json:"spec"`
	Status            S3BucketObjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3BucketObjectList contains a list of S3BucketObjects
type S3BucketObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3BucketObject `json:"items"`
}

// Repository type metadata.
var (
	S3BucketObject_Kind             = "S3BucketObject"
	S3BucketObject_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: S3BucketObject_Kind}.String()
	S3BucketObject_KindAPIVersion   = S3BucketObject_Kind + "." + CRDGroupVersion.String()
	S3BucketObject_GroupVersionKind = CRDGroupVersion.WithKind(S3BucketObject_Kind)
)

func init() {
	SchemeBuilder.Register(&S3BucketObject{}, &S3BucketObjectList{})
}
