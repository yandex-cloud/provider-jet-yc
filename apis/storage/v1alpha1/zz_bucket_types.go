/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AnonymousAccessFlagsObservation struct {
}

type AnonymousAccessFlagsParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Allows to list object in bucket anonymously.
	List *bool `json:"list,omitempty" tf:"list,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Allows to read objects in bucket anonymously.
	Read *bool `json:"read,omitempty" tf:"read,omitempty"`
}

type ApplyServerSideEncryptionByDefaultObservation struct {
}

type ApplyServerSideEncryptionByDefaultParameters struct {

	// +kubebuilder:validation:Required
	// (Optional) The KMS master key ID used for the SSE-KMS encryption.
	KMSMasterKeyID *string `json:"kmsMasterKeyId" tf:"kms_master_key_id,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) The server-side encryption algorithm to use. Single valid value is `aws:kms`
	SseAlgorithm *string `json:"sseAlgorithm" tf:"sse_algorithm,omitempty"`
}

type BucketObservation struct {
	// The bucket domain name.
	BucketDomainName *string `json:"bucketDomainName,omitempty" tf:"bucket_domain_name,omitempty"`

	// (Optional) Unique identifier for the rule. Must be less than or equal to 255 characters in length.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BucketParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) The [predefined ACL](https://cloud.yandex.com/docs/storage/concepts/acl#predefined_acls) to apply. Defaults to `private`. Conflicts with `grant`.
	ACL *string `json:"acl,omitempty" tf:"acl,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/iam/v1alpha1.ServiceAccountStaticAccessKey
	// +crossplane:generate:reference:extractor=github.com/yandex-cloud/provider-jet-yc/config/storage.ExtractAccessKey()
	// +kubebuilder:validation:Optional
	// (Optional) The access key to use when applying changes. If omitted, `storage_access_key` specified in provider config is used.
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// +kubebuilder:validation:Optional
	AccessKeyRef *v1.Reference `json:"accessKeyRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	AccessKeySelector *v1.Selector `json:"accessKeySelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	// (Optional) Provides various access to objects.
	AnonymousAccessFlags []AnonymousAccessFlagsParameters `json:"anonymousAccessFlags,omitempty" tf:"anonymous_access_flags,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional, Forces new resource) The name of the bucket. If omitted, Terraform will assign a random, unique name.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional, Forces new resource) Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) A rule of [Cross-Origin Resource Sharing](https://cloud.yandex.com/docs/storage/cors/) (documented below).
	CorsRule []CorsRuleParameters `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Storage class which is used for storing objects by default.
	DefaultStorageClass *string `json:"defaultStorageClass,omitempty" tf:"default_storage_class,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	// (Optional) Allow to create bucket in different folder.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	// (Optional, Default: `false`) A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) An [ACL policy grant](https://cloud.yandex.com/docs/storage/concepts/acl#permissions-types). Conflicts with `acl`.
	Grant []GrantParameters `json:"grant,omitempty" tf:"grant,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Manages https certificates for bucket. See [https](https://cloud.yandex.com/en-ru/docs/storage/operations/hosting/certificate) for more infomation.
	HTTPS []HTTPSParameters `json:"https,omitempty" tf:"https,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) A configuration of [object lifecycle management](https://cloud.yandex.com/docs/storage/concepts/lifecycles) (documented below).
	LifecycleRule []LifecycleRuleParameters `json:"lifecycleRule,omitempty" tf:"lifecycle_rule,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) A settings of [bucket logging](https://cloud.yandex.com/docs/storage/concepts/server-logs) (documented below).
	Logging []LoggingParameters `json:"logging,omitempty" tf:"logging,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) The size of bucket, in bytes. See [size limiting](https://cloud.yandex.com/en-ru/docs/storage/operations/buckets/limit-max-volume) for more information.
	MaxSize *float64 `json:"maxSize,omitempty" tf:"max_size,omitempty"`

	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// +kubebuilder:validation:Optional
	SecretKeySecretRef *v1.SecretKeySelector `json:"secretKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	// (Optional) A configuration of server-side encryption for the bucket (documented below)
	ServerSideEncryptionConfiguration []ServerSideEncryptionConfigurationParameters `json:"serverSideEncryptionConfiguration,omitempty" tf:"server_side_encryption_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) A state of [versioning](https://cloud.yandex.com/docs/storage/concepts/versioning) (documented below)
	Versioning []VersioningParameters `json:"versioning,omitempty" tf:"versioning,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) A [website object](https://cloud.yandex.com/docs/storage/concepts/hosting) (documented below).
	Website []WebsiteParameters `json:"website,omitempty" tf:"website,omitempty"`

	// +kubebuilder:validation:Optional
	// The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteDomain *string `json:"websiteDomain,omitempty" tf:"website_domain,omitempty"`

	// +kubebuilder:validation:Optional
	// The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
	WebsiteEndpoint *string `json:"websiteEndpoint,omitempty" tf:"website_endpoint,omitempty"`
}

type CorsRuleObservation struct {
}

type CorsRuleParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Specifies which headers are allowed.
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Specifies which methods are allowed. Can be `GET`, `PUT`, `POST`, `DELETE` or `HEAD`.
	AllowedMethods []*string `json:"allowedMethods" tf:"allowed_methods,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Specifies which origins are allowed.
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Specifies expose header in the response.
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Specifies time in seconds that browser can cache the response for a preflight request.
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type ExpirationObservation struct {
}

type ExpirationParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Specifies the date after which you want the corresponding action to take effect.
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Specifies the number of days after object creation when the specific rule action takes effect.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) On a versioned bucket (versioning-enabled or versioning-suspended bucket), you can add this element in the lifecycle configuration to direct Object Storage to delete expired object delete markers.
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type GrantObservation struct {
}

type GrantParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Unique identifier for the rule. Must be less than or equal to 255 characters in length.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Required
	Permissions []*string `json:"permissions" tf:"permissions,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type HTTPSObservation struct {
}

type HTTPSParameters struct {

	// +kubebuilder:validation:Required
	CertificateID *string `json:"certificateId" tf:"certificate_id,omitempty"`
}

type LifecycleRuleObservation struct {
}

type LifecycleRuleParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Specifies the number of days after initiating a multipart upload when the multipart upload must be completed.
	AbortIncompleteMultipartUploadDays *float64 `json:"abortIncompleteMultipartUploadDays,omitempty" tf:"abort_incomplete_multipart_upload_days,omitempty"`

	// +kubebuilder:validation:Required
	// (Optional) Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket.
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Specifies a period in the object's expire (documented below).
	Expiration []ExpirationParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Unique identifier for the rule. Must be less than or equal to 255 characters in length.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Specifies when noncurrent object versions expire (documented below).
	NoncurrentVersionExpiration []NoncurrentVersionExpirationParameters `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Specifies when noncurrent object versions transitions (documented below).
	NoncurrentVersionTransition []NoncurrentVersionTransitionParameters `json:"noncurrentVersionTransition,omitempty" tf:"noncurrent_version_transition,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Object key prefix identifying one or more objects to which the rule applies.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Specifies a period in the object's transitions (documented below).
	Transition []TransitionParameters `json:"transition,omitempty" tf:"transition,omitempty"`
}

type LoggingObservation struct {
}

type LoggingParameters struct {

	// +kubebuilder:validation:Required
	// (Required) The name of the bucket that will receive the log objects.
	TargetBucket *string `json:"targetBucket" tf:"target_bucket,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) To specify a key prefix for log objects.
	TargetPrefix *string `json:"targetPrefix,omitempty" tf:"target_prefix,omitempty"`
}

type NoncurrentVersionExpirationObservation struct {
}

type NoncurrentVersionExpirationParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Specifies the number of days after object creation when the specific rule action takes effect.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`
}

type NoncurrentVersionTransitionObservation struct {
}

type NoncurrentVersionTransitionParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Specifies the number of days after object creation when the specific rule action takes effect.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Specifies the storage class to which you want the object to transition. Can only be `COLD` or `STANDARD_IA`.
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`
}

type RuleObservation struct {
}

type RuleParameters struct {

	// +kubebuilder:validation:Required
	// (Required) A single object for setting server-side encryption by default. (documented below)
	ApplyServerSideEncryptionByDefault []ApplyServerSideEncryptionByDefaultParameters `json:"applyServerSideEncryptionByDefault" tf:"apply_server_side_encryption_by_default,omitempty"`
}

type ServerSideEncryptionConfigurationObservation struct {
}

type ServerSideEncryptionConfigurationParameters struct {

	// +kubebuilder:validation:Required
	// (Required) A single object for server-side encryption by default configuration. (documented below)
	Rule []RuleParameters `json:"rule" tf:"rule,omitempty"`
}

type TransitionObservation struct {
}

type TransitionParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Specifies the date after which you want the corresponding action to take effect.
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Specifies the number of days after object creation when the specific rule action takes effect.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Specifies the storage class to which you want the object to transition. Can only be `COLD` or `STANDARD_IA`.
	StorageClass *string `json:"storageClass" tf:"storage_class,omitempty"`
}

type VersioningObservation struct {
}

type VersioningParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type WebsiteObservation struct {
}

type WebsiteParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) An absolute path to the document to return in case of a 4XX error.
	ErrorDocument *string `json:"errorDocument,omitempty" tf:"error_document,omitempty"`

	// +kubebuilder:validation:Optional
	// (Required, unless using `redirect_all_requests_to`) Storage returns this index document when requests are made to the root domain or any of the subfolders.
	IndexDocument *string `json:"indexDocument,omitempty" tf:"index_document,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) A hostname to redirect all website requests for this bucket to. Hostname can optionally be prefixed with a protocol (`http://` or `https://`) to use when redirecting requests. The default is the protocol that is used in the original request.
	RedirectAllRequestsTo *string `json:"redirectAllRequestsTo,omitempty" tf:"redirect_all_requests_to,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) A json array containing [routing rules](https://cloud.yandex.com/docs/storage/s3/api-ref/hosting/upload#request-scheme) describing redirect behavior and when redirects are applied.
	RoutingRules *string `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketParameters `json:"forProvider"`
}

// BucketStatus defines the observed state of Bucket.
type BucketStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Bucket is the Schema for the Buckets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketSpec   `json:"spec"`
	Status            BucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketList contains a list of Buckets
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

// Repository type metadata.
var (
	Bucket_Kind             = "Bucket"
	Bucket_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Bucket_Kind}.String()
	Bucket_KindAPIVersion   = Bucket_Kind + "." + CRDGroupVersion.String()
	Bucket_GroupVersionKind = CRDGroupVersion.WithKind(Bucket_Kind)
)

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}
