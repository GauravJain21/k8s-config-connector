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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Config struct {
	/* Set of IP addresses or ranges (IPV4 or IPV6) in CIDR notation to match against inbound traffic. There is a limit of 10 IP ranges per rule. A value of '*' matches all IPs (can be used to override the default behavior). */
	SrcIpRanges []string `json:"srcIpRanges,omitempty"`
}

type Expr struct {
	/* Textual representation of an expression in Common Expression Language syntax. The application context of the containing message determines which well-known feature set of CEL is supported. */
	Expression string `json:"expression,omitempty"`
}

type Match struct {
	/* The configuration options available when specifying versioned_expr. This field must be specified if versioned_expr is specified and cannot be specified if versioned_expr is not specified. */
	Config Config `json:"config,omitempty"`
	/* User defined CEVAL expression. A CEVAL expression is used to specify match criteria such as origin.ip, source.region_code and contents in the request header. */
	Expr Expr `json:"expr,omitempty"`
	/* Predefined rule expression. If this field is specified, config must also be specified. Available options:   SRC_IPS_V1: Must specify the corresponding src_ip_ranges field in config. */
	VersionedExpr string `json:"versionedExpr,omitempty"`
}

type Rule struct {
	/* Action to take when match matches the request. Valid values:   "allow" : allow access to target, "deny(status)" : deny access to target, returns the HTTP response code specified (valid values are 403, 404 and 502) */
	Action string `json:"action,omitempty"`
	/* An optional description of this rule. Max size is 64. */
	Description string `json:"description,omitempty"`
	/* A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding action is enforced. */
	Match Match `json:"match,omitempty"`
	/* When set to true, the action specified above is not enforced. Stackdriver logs for requests that trigger a preview action are annotated as such. */
	Preview bool `json:"preview,omitempty"`
	/* An unique positive integer indicating the priority of evaluation for a rule. Rules are evaluated from highest priority (lowest numerically) to lowest priority (highest numerically) in order. */
	Priority int `json:"priority,omitempty"`
}

type ComputeSecurityPolicySpec struct {
	/* An optional description of this security policy. Max size is 2048. */
	Description string `json:"description,omitempty"`
	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	ResourceID string `json:"resourceID,omitempty"`
	/* The set of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added. */
	Rule []Rule `json:"rule,omitempty"`
}

type ComputeSecurityPolicyStatus struct {
	/* Conditions represents the latest available observations of the
	   ComputeSecurityPolicy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Fingerprint of this resource. */
	Fingerprint string `json:"fingerprint,omitempty"`
	/* The URI of the created resource. */
	SelfLink string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeSecurityPolicy is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeSecurityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeSecurityPolicySpec   `json:"spec,omitempty"`
	Status ComputeSecurityPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeSecurityPolicyList contains a list of ComputeSecurityPolicy
type ComputeSecurityPolicyList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []ComputeSecurityPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeSecurityPolicy{}, &ComputeSecurityPolicyList{})
}
