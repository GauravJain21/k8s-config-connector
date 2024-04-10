// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BackendbucketBypassCacheOnRequestHeaders struct {
	/* The header field name to match on when bypassing cache. Values are case-insensitive. */
	// +optional
	HeaderName *string `json:"headerName,omitempty"`
}

type BackendbucketCacheKeyPolicy struct {
	/* Allows HTTP request headers (by name) to be used in the
	cache key. */
	// +optional
	IncludeHttpHeaders []string `json:"includeHttpHeaders,omitempty"`

	/* Names of query string parameters to include in cache keys.
	Default parameters are always included. '&' and '=' will
	be percent encoded and not treated as delimiters. */
	// +optional
	QueryStringWhitelist []string `json:"queryStringWhitelist,omitempty"`
}

type BackendbucketCdnPolicy struct {
	/* Bypass the cache when the specified request headers are matched - e.g. Pragma or Authorization headers. Up to 5 headers can be specified. The cache is bypassed for all cdnPolicy.cacheMode settings. */
	// +optional
	BypassCacheOnRequestHeaders []BackendbucketBypassCacheOnRequestHeaders `json:"bypassCacheOnRequestHeaders,omitempty"`

	/* The CacheKeyPolicy for this CdnPolicy. */
	// +optional
	CacheKeyPolicy *BackendbucketCacheKeyPolicy `json:"cacheKeyPolicy,omitempty"`

	/* Specifies the cache setting for all responses from this backend.
	The possible values are: USE_ORIGIN_HEADERS, FORCE_CACHE_ALL and CACHE_ALL_STATIC Possible values: ["USE_ORIGIN_HEADERS", "FORCE_CACHE_ALL", "CACHE_ALL_STATIC"]. */
	// +optional
	CacheMode *string `json:"cacheMode,omitempty"`

	/* Specifies the maximum allowed TTL for cached content served by this origin. */
	// +optional
	ClientTtl *int `json:"clientTtl,omitempty"`

	/* Specifies the default TTL for cached content served by this origin for responses
	that do not have an existing valid TTL (max-age or s-max-age). */
	// +optional
	DefaultTtl *int `json:"defaultTtl,omitempty"`

	/* Specifies the maximum allowed TTL for cached content served by this origin. */
	// +optional
	MaxTtl *int `json:"maxTtl,omitempty"`

	/* Negative caching allows per-status code TTLs to be set, in order to apply fine-grained caching for common errors or redirects. */
	// +optional
	NegativeCaching *bool `json:"negativeCaching,omitempty"`

	/* Sets a cache TTL for the specified HTTP status code. negativeCaching must be enabled to configure negativeCachingPolicy.
	Omitting the policy and leaving negativeCaching enabled will use Cloud CDN's default cache TTLs. */
	// +optional
	NegativeCachingPolicy []BackendbucketNegativeCachingPolicy `json:"negativeCachingPolicy,omitempty"`

	/* If true then Cloud CDN will combine multiple concurrent cache fill requests into a small number of requests to the origin. */
	// +optional
	RequestCoalescing *bool `json:"requestCoalescing,omitempty"`

	/* Serve existing content from the cache (if available) when revalidating content with the origin, or when an error is encountered when refreshing the cache. */
	// +optional
	ServeWhileStale *int `json:"serveWhileStale,omitempty"`

	/* Maximum number of seconds the response to a signed URL request will
	be considered fresh. After this time period,
	the response will be revalidated before being served.
	When serving responses to signed URL requests,
	Cloud CDN will internally behave as though
	all responses from this backend had a "Cache-Control: public,
	max-age=[TTL]" header, regardless of any existing Cache-Control
	header. The actual headers served in responses will not be altered. */
	// +optional
	SignedUrlCacheMaxAgeSec *int `json:"signedUrlCacheMaxAgeSec,omitempty"`
}

type BackendbucketNegativeCachingPolicy struct {
	/* The HTTP status code to define a TTL against. Only HTTP status codes 300, 301, 308, 404, 405, 410, 421, 451 and 501
	can be specified as values, and you cannot specify a status code more than once. */
	// +optional
	Code *int `json:"code,omitempty"`

	/* The TTL (in seconds) for which to cache responses with the corresponding status code. The maximum allowed value is 1800s
	(30 minutes), noting that infrequently accessed objects may be evicted from the cache before the defined TTL. */
	// +optional
	Ttl *int `json:"ttl,omitempty"`
}

type ComputeBackendBucketSpec struct {
	/* Reference to the bucket. */
	BucketRef v1alpha1.ResourceRef `json:"bucketRef"`

	/* Cloud CDN configuration for this Backend Bucket. */
	// +optional
	CdnPolicy *BackendbucketCdnPolicy `json:"cdnPolicy,omitempty"`

	/* Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header. Possible values: ["AUTOMATIC", "DISABLED"]. */
	// +optional
	CompressionMode *string `json:"compressionMode,omitempty"`

	/* Headers that the HTTP/S load balancer should add to proxied responses. */
	// +optional
	CustomResponseHeaders []string `json:"customResponseHeaders,omitempty"`

	/* An optional textual description of the resource; provided by the
	client when the resource is created. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The security policy associated with this backend bucket. */
	// +optional
	EdgeSecurityPolicy *string `json:"edgeSecurityPolicy,omitempty"`

	/* If true, enable Cloud CDN for this BackendBucket. */
	// +optional
	EnableCdn *bool `json:"enableCdn,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type ComputeBackendBucketStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeBackendBucket's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Creation timestamp in RFC3339 text format. */
	// +optional
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputebackendbucket;gcpcomputebackendbuckets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeBackendBucket is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeBackendBucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeBackendBucketSpec   `json:"spec,omitempty"`
	Status ComputeBackendBucketStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeBackendBucketList contains a list of ComputeBackendBucket
type ComputeBackendBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeBackendBucket `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeBackendBucket{}, &ComputeBackendBucketList{})
}
