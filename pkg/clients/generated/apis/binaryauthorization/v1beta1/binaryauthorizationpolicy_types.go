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

type PolicyAdmissionWhitelistPatterns struct {
	/* An image name pattern to allowlist, in the form `registry/path/to/image`. This supports a trailing `*` as a wildcard, but this is allowed only in text after the `registry/` part. */
	// +optional
	NamePattern *string `json:"namePattern,omitempty"`
}

type PolicyClusterAdmissionRules struct {
	/* Required. The action when a pod creation is denied by the admission rule. Possible values: ENFORCEMENT_MODE_UNSPECIFIED, ENFORCED_BLOCK_AND_AUDIT_LOG, DRYRUN_AUDIT_LOG_ONLY */
	EnforcementMode string `json:"enforcementMode"`

	/* Required. How this admission rule will be evaluated. Possible values: ALWAYS_ALLOW, ALWAYS_DENY, REQUIRE_ATTESTATION */
	EvaluationMode string `json:"evaluationMode"`

	/*  */
	// +optional
	RequireAttestationsBy []v1alpha1.ResourceRef `json:"requireAttestationsBy,omitempty"`
}

type PolicyDefaultAdmissionRule struct {
	/* Required. The action when a pod creation is denied by the admission rule. Possible values: ENFORCEMENT_MODE_UNSPECIFIED, ENFORCED_BLOCK_AND_AUDIT_LOG, DRYRUN_AUDIT_LOG_ONLY */
	EnforcementMode string `json:"enforcementMode"`

	/* Required. How this admission rule will be evaluated. Possible values: ALWAYS_ALLOW, ALWAYS_DENY, REQUIRE_ATTESTATION */
	EvaluationMode string `json:"evaluationMode"`

	/*  */
	// +optional
	RequireAttestationsBy []v1alpha1.ResourceRef `json:"requireAttestationsBy,omitempty"`
}

type PolicyIstioServiceIdentityAdmissionRules struct {
	/* Required. The action when a pod creation is denied by the admission rule. Possible values: ENFORCEMENT_MODE_UNSPECIFIED, ENFORCED_BLOCK_AND_AUDIT_LOG, DRYRUN_AUDIT_LOG_ONLY */
	EnforcementMode string `json:"enforcementMode"`

	/* Required. How this admission rule will be evaluated. Possible values: ALWAYS_ALLOW, ALWAYS_DENY, REQUIRE_ATTESTATION */
	EvaluationMode string `json:"evaluationMode"`

	/*  */
	// +optional
	RequireAttestationsBy []v1alpha1.ResourceRef `json:"requireAttestationsBy,omitempty"`
}

type PolicyKubernetesNamespaceAdmissionRules struct {
	/* Required. The action when a pod creation is denied by the admission rule. Possible values: ENFORCEMENT_MODE_UNSPECIFIED, ENFORCED_BLOCK_AND_AUDIT_LOG, DRYRUN_AUDIT_LOG_ONLY */
	EnforcementMode string `json:"enforcementMode"`

	/* Required. How this admission rule will be evaluated. Possible values: ALWAYS_ALLOW, ALWAYS_DENY, REQUIRE_ATTESTATION */
	EvaluationMode string `json:"evaluationMode"`

	/*  */
	// +optional
	RequireAttestationsBy []v1alpha1.ResourceRef `json:"requireAttestationsBy,omitempty"`
}

type PolicyKubernetesServiceAccountAdmissionRules struct {
	/* Required. The action when a pod creation is denied by the admission rule. Possible values: ENFORCEMENT_MODE_UNSPECIFIED, ENFORCED_BLOCK_AND_AUDIT_LOG, DRYRUN_AUDIT_LOG_ONLY */
	EnforcementMode string `json:"enforcementMode"`

	/* Required. How this admission rule will be evaluated. Possible values: ALWAYS_ALLOW, ALWAYS_DENY, REQUIRE_ATTESTATION */
	EvaluationMode string `json:"evaluationMode"`

	/*  */
	// +optional
	RequireAttestationsBy []v1alpha1.ResourceRef `json:"requireAttestationsBy,omitempty"`
}

type BinaryAuthorizationPolicySpec struct {
	/* Optional. Admission policy allowlisting. A matching admission request will always be permitted. This feature is typically used to exclude Google or third-party infrastructure images from Binary Authorization policies. */
	// +optional
	AdmissionWhitelistPatterns []PolicyAdmissionWhitelistPatterns `json:"admissionWhitelistPatterns,omitempty"`

	/* Optional. Per-cluster admission rules. Cluster spec format: location.clusterId. There can be at most one admission rule per cluster spec. A location is either a compute zone (e.g. us-central1-a) or a region (e.g. us-central1). For clusterId syntax restrictions see https://cloud.google.com/container-engine/reference/rest/v1/projects.zones.clusters. */
	// +optional
	ClusterAdmissionRules map[string]PolicyClusterAdmissionRules `json:"clusterAdmissionRules,omitempty"`

	/* Required. Default admission rule for a cluster without a per-cluster, per-kubernetes-service-account, or per-istio-service-identity admission rule. */
	DefaultAdmissionRule PolicyDefaultAdmissionRule `json:"defaultAdmissionRule"`

	/* Optional. A descriptive comment. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Optional. Controls the evaluation of a Google-maintained global admission policy for common system-level images. Images not covered by the global policy will be subject to the project admission policy. This setting has no effect when specified inside a global admission policy. Possible values: GLOBAL_POLICY_EVALUATION_MODE_UNSPECIFIED, ENABLE, DISABLE */
	// +optional
	GlobalPolicyEvaluationMode *string `json:"globalPolicyEvaluationMode,omitempty"`

	/* Optional. Per-istio-service-identity admission rules. Istio service identity spec format: spiffe:///ns//sa/ or /ns//sa/ e.g. spiffe://example.com/ns/test-ns/sa/default */
	// +optional
	IstioServiceIdentityAdmissionRules map[string]PolicyIstioServiceIdentityAdmissionRules `json:"istioServiceIdentityAdmissionRules,omitempty"`

	/* Optional. Per-kubernetes-namespace admission rules. K8s namespace spec format: [a-z.-]+, e.g. 'some-namespace' */
	// +optional
	KubernetesNamespaceAdmissionRules map[string]PolicyKubernetesNamespaceAdmissionRules `json:"kubernetesNamespaceAdmissionRules,omitempty"`

	/* Optional. Per-kubernetes-service-account admission rules. Service account spec format: namespace:serviceaccount. e.g. 'test-ns:default' */
	// +optional
	KubernetesServiceAccountAdmissionRules map[string]PolicyKubernetesServiceAccountAdmissionRules `json:"kubernetesServiceAccountAdmissionRules,omitempty"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`
}

type BinaryAuthorizationPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   BinaryAuthorizationPolicy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* Output only. The resource name, in the format `projects/* /policy`. There is at most one policy per project. */
	SelfLink string `json:"selfLink,omitempty"`
	/* Output only. Time when the policy was last updated. */
	UpdateTime string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BinaryAuthorizationPolicy is the Schema for the binaryauthorization API
// +k8s:openapi-gen=true
type BinaryAuthorizationPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BinaryAuthorizationPolicySpec   `json:"spec,omitempty"`
	Status BinaryAuthorizationPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BinaryAuthorizationPolicyList contains a list of BinaryAuthorizationPolicy
type BinaryAuthorizationPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BinaryAuthorizationPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BinaryAuthorizationPolicy{}, &BinaryAuthorizationPolicyList{})
}
