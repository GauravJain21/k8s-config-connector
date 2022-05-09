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

type UptimecheckconfigAuthInfo struct {
	/*  */
	Password UptimecheckconfigPassword `json:"password"`

	/*  */
	Username string `json:"username"`
}

type UptimecheckconfigContentMatchers struct {
	/*  */
	Content string `json:"content"`

	/*  Possible values: CONTENT_MATCHER_OPTION_UNSPECIFIED, CONTAINS_STRING, NOT_CONTAINS_STRING, MATCHES_REGEX, NOT_MATCHES_REGEX */
	// +optional
	Matcher *string `json:"matcher,omitempty"`
}

type UptimecheckconfigHttpCheck struct {
	/* The authentication information. Optional when creating an HTTP check; defaults to empty. */
	// +optional
	AuthInfo *UptimecheckconfigAuthInfo `json:"authInfo,omitempty"`

	/* The request body associated with the HTTP POST request. If `content_type` is `URL_ENCODED`, the body passed in must be URL-encoded. Users can provide a `Content-Length` header via the `headers` field or the API will do so. If the `request_method` is `GET` and `body` is not empty, the API will return an error. The maximum byte size is 1 megabyte. Note: As with all `bytes` fields JSON representations are base64 encoded. e.g.: "foo=bar" in URL-encoded form is "foo%3Dbar" and in base64 encoding is "Zm9vJTI1M0RiYXI=". */
	// +optional
	Body *string `json:"body,omitempty"`

	/* Immutable. The content type to use for the check.  Possible values: TYPE_UNSPECIFIED, URL_ENCODED */
	// +optional
	ContentType *string `json:"contentType,omitempty"`

	/* The list of headers to send as part of the Uptime check request. If two headers have the same key and different values, they should be entered as a single header, with the value being a comma-separated list of all the desired values as described at https://www.w3.org/Protocols/rfc2616/rfc2616.txt (page 31). Entering two separate headers with the same key in a Create call will cause the first to be overwritten by the second. The maximum number of headers allowed is 100. */
	// +optional
	Headers map[string]string `json:"headers,omitempty"`

	/* Immutable. Boolean specifying whether to encrypt the header information. Encryption should be specified for any headers related to authentication that you do not wish to be seen when retrieving the configuration. The server will be responsible for encrypting the headers. On Get/List calls, if `mask_headers` is set to `true` then the headers will be obscured with `******.` */
	// +optional
	MaskHeaders *bool `json:"maskHeaders,omitempty"`

	/* Optional (defaults to "/"). The path to the page against which to run the check. Will be combined with the `host` (specified within the `monitored_resource`) and `port` to construct the full URL. If the provided path does not begin with "/", a "/" will be prepended automatically. */
	// +optional
	Path *string `json:"path,omitempty"`

	/* Optional (defaults to 80 when `use_ssl` is `false`, and 443 when `use_ssl` is `true`). The TCP port on the HTTP server against which to run the check. Will be combined with host (specified within the `monitored_resource`) and `path` to construct the full URL. */
	// +optional
	Port *int `json:"port,omitempty"`

	/* Immutable. The HTTP request method to use for the check. If set to `METHOD_UNSPECIFIED` then `request_method` defaults to `GET`. */
	// +optional
	RequestMethod *string `json:"requestMethod,omitempty"`

	/* If `true`, use HTTPS instead of HTTP to run the check. */
	// +optional
	UseSsl *bool `json:"useSsl,omitempty"`

	/* Boolean specifying whether to include SSL certificate validation as a part of the Uptime check. Only applies to checks where `monitored_resource` is set to `uptime_url`. If `use_ssl` is `false`, setting `validate_ssl` to `true` has no effect. */
	// +optional
	ValidateSsl *bool `json:"validateSsl,omitempty"`
}

type UptimecheckconfigMonitoredResource struct {
	/* Immutable. */
	FilterLabels map[string]string `json:"filterLabels"`

	/* Immutable. */
	Type string `json:"type"`
}

type UptimecheckconfigPassword struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *UptimecheckconfigValueFrom `json:"valueFrom,omitempty"`
}

type UptimecheckconfigResourceGroup struct {
	/* Immutable. */
	// +optional
	GroupRef *v1alpha1.ResourceRef `json:"groupRef,omitempty"`

	/* Immutable. The resource type of the group members. Possible values: RESOURCE_TYPE_UNSPECIFIED, INSTANCE, AWS_ELB_LOAD_BALANCER */
	// +optional
	ResourceType *string `json:"resourceType,omitempty"`
}

type UptimecheckconfigTcpCheck struct {
	/* The TCP port on the server against which to run the check. Will be combined with host (specified within the `monitored_resource`) to construct the full URL. Required. */
	Port int `json:"port"`
}

type UptimecheckconfigValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	// +optional
	SecretKeyRef *v1alpha1.ResourceRef `json:"secretKeyRef,omitempty"`
}

type MonitoringUptimeCheckConfigSpec struct {
	/* The content that is expected to appear in the data returned by the target server against which the check is run.  Currently, only the first entry in the `content_matchers` list is supported, and additional entries will be ignored. This field is optional and should only be specified if a content match is required as part of the/ Uptime check. */
	// +optional
	ContentMatchers []UptimecheckconfigContentMatchers `json:"contentMatchers,omitempty"`

	/* A human-friendly name for the Uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced. Required. */
	DisplayName string `json:"displayName"`

	/* Contains information needed to make an HTTP or HTTPS check. */
	// +optional
	HttpCheck *UptimecheckconfigHttpCheck `json:"httpCheck,omitempty"`

	/* Immutable. The [monitored resource](https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for Uptime checks:   `uptime_url`,   `gce_instance`,   `gae_app`,   `aws_ec2_instance`,   `aws_elb_load_balancer` */
	// +optional
	MonitoredResource *UptimecheckconfigMonitoredResource `json:"monitoredResource,omitempty"`

	/* How often, in seconds, the Uptime check is performed. Currently, the only supported values are `60s` (1 minute), `300s` (5 minutes), `600s` (10 minutes), and `900s` (15 minutes). Optional, defaults to `60s`. */
	// +optional
	Period *string `json:"period,omitempty"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. The group resource associated with the configuration. */
	// +optional
	ResourceGroup *UptimecheckconfigResourceGroup `json:"resourceGroup,omitempty"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions must be provided to include a minimum of 3 locations.  Not specifying this field will result in Uptime checks running from all available regions. */
	// +optional
	SelectedRegions []string `json:"selectedRegions,omitempty"`

	/* Contains information needed to make a TCP check. */
	// +optional
	TcpCheck *UptimecheckconfigTcpCheck `json:"tcpCheck,omitempty"`

	/* The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Required. */
	Timeout string `json:"timeout"`
}

type MonitoringUptimeCheckConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   MonitoringUptimeCheckConfig's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MonitoringUptimeCheckConfig is the Schema for the monitoring API
// +k8s:openapi-gen=true
type MonitoringUptimeCheckConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitoringUptimeCheckConfigSpec   `json:"spec,omitempty"`
	Status MonitoringUptimeCheckConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MonitoringUptimeCheckConfigList contains a list of MonitoringUptimeCheckConfig
type MonitoringUptimeCheckConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitoringUptimeCheckConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MonitoringUptimeCheckConfig{}, &MonitoringUptimeCheckConfigList{})
}
