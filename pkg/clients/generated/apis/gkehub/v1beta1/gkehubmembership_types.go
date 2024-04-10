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

type MembershipAuthority struct {
	/* Optional. A JSON Web Token (JWT) issuer URI. `issuer` must start with `https://` and be a valid URL with length <2000 characters. If set, then Google will allow valid OIDC tokens from this issuer to authenticate within the workload_identity_pool. OIDC discovery will be performed on this URI to validate tokens from the issuer. Clearing `issuer` disables Workload Identity. `issuer` cannot be directly modified; it must be cleared (and Workload Identity disabled) before using a new issuer (and re-enabling Workload Identity). */
	// +optional
	Issuer *string `json:"issuer,omitempty"`
}

type MembershipEndpoint struct {
	/* Optional. GKE-specific information. Only present if this Membership is a GKE cluster. */
	// +optional
	GkeCluster *MembershipGkeCluster `json:"gkeCluster,omitempty"`

	/* Optional. The in-cluster Kubernetes Resources that should be applied for a correctly registered cluster, in the steady state. These resources: * Ensure that the cluster is exclusively registered to one and only one Hub Membership. * Propagate Workload Pool Information available in the Membership Authority field. * Ensure proper initial configuration of default Hub Features. */
	// +optional
	KubernetesResource *MembershipKubernetesResource `json:"kubernetesResource,omitempty"`
}

type MembershipGkeCluster struct {
	// +optional
	ResourceRef *v1alpha1.ResourceRef `json:"resourceRef,omitempty"`
}

type MembershipKubernetesResource struct {
	/* Input only. The YAML representation of the Membership CR. This field is ignored for GKE clusters where Hub can read the CR directly. Callers should provide the CR that is currently present in the cluster during CreateMembership or UpdateMembership, or leave this field empty if none exists. The CR manifest is used to validate the cluster has not been registered with another Membership. */
	// +optional
	MembershipCrManifest *string `json:"membershipCrManifest,omitempty"`

	/* Optional. Options for Kubernetes resource generation. */
	// +optional
	ResourceOptions *MembershipResourceOptions `json:"resourceOptions,omitempty"`
}

type MembershipResourceOptions struct {
	/* Optional. The Connect agent version to use for connect_resources. Defaults to the latest GKE Connect version. The version must be a currently supported version, obsolete versions will be rejected. */
	// +optional
	ConnectVersion *string `json:"connectVersion,omitempty"`

	/* Optional. Use `apiextensions/v1beta1` instead of `apiextensions/v1` for CustomResourceDefinition resources. This option should be set for clusters with Kubernetes apiserver versions <1.16. */
	// +optional
	V1beta1Crd *bool `json:"v1beta1Crd,omitempty"`
}

type GKEHubMembershipSpec struct {
	/* Optional. How to identify workloads from this Membership. See the documentation on Workload Identity for more details: https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity */
	// +optional
	Authority *MembershipAuthority `json:"authority,omitempty"`

	/* Description of this membership, limited to 63 characters. Must match the regex: `*` This field is present for legacy purposes. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Optional. Endpoint information to reach this member. */
	// +optional
	Endpoint *MembershipEndpoint `json:"endpoint,omitempty"`

	/* Optional. An externally-generated and managed ID for this Membership. This ID may be modified after creation, but this is not recommended. The ID must match the regex: `*` If this Membership represents a Kubernetes cluster, this value should be set to the UID of the `kube-system` namespace object. */
	// +optional
	ExternalId *string `json:"externalId,omitempty"`

	/* Optional. The infrastructure type this Membership is running on. Possible values: INFRASTRUCTURE_TYPE_UNSPECIFIED, ON_PREM, MULTI_CLOUD */
	// +optional
	InfrastructureType *string `json:"infrastructureType,omitempty"`

	/* Immutable. The location for the resource */
	Location string `json:"location"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type MembershipAuthorityStatus struct {
	/* Output only. An identity provider that reflects the `issuer` in the workload identity pool. */
	// +optional
	IdentityProvider *string `json:"identityProvider,omitempty"`

	/* Output only. The name of the workload identity pool in which `issuer` will be recognized. There is a single Workload Identity Pool per Hub that is shared between all Memberships that belong to that Hub. For a Hub hosted in: {PROJECT_ID}, the workload pool format is `{PROJECT_ID}.hub.id.goog`, although this is subject to change in newer versions of this API. */
	// +optional
	WorkloadIdentityPool *string `json:"workloadIdentityPool,omitempty"`
}

type MembershipConnectResourcesStatus struct {
	/* Whether the resource provided in the manifest is `cluster_scoped`. If unset, the manifest is assumed to be namespace scoped. This field is used for REST mapping when applying the resource in a cluster. */
	// +optional
	ClusterScoped *bool `json:"clusterScoped,omitempty"`

	/* YAML manifest of the resource. */
	// +optional
	Manifest *string `json:"manifest,omitempty"`
}

type MembershipEndpointStatus struct {
	/* Output only. Useful Kubernetes-specific metadata. */
	// +optional
	KubernetesMetadata *MembershipKubernetesMetadataStatus `json:"kubernetesMetadata,omitempty"`

	// +optional
	KubernetesResource *MembershipKubernetesResourceStatus `json:"kubernetesResource,omitempty"`
}

type MembershipKubernetesMetadataStatus struct {
	/* Output only. Kubernetes API server version string as reported by `/version`. */
	// +optional
	KubernetesApiServerVersion *string `json:"kubernetesApiServerVersion,omitempty"`

	/* Output only. The total memory capacity as reported by the sum of all Kubernetes nodes resources, defined in MB. */
	// +optional
	MemoryMb *int `json:"memoryMb,omitempty"`

	/* Output only. Node count as reported by Kubernetes nodes resources. */
	// +optional
	NodeCount *int `json:"nodeCount,omitempty"`

	/* Output only. Node providerID as reported by the first node in the list of nodes on the Kubernetes endpoint. On Kubernetes platforms that support zero-node clusters (like GKE-on-GCP), the node_count will be zero and the node_provider_id will be empty. */
	// +optional
	NodeProviderId *string `json:"nodeProviderId,omitempty"`

	/* Output only. The time at which these details were last updated. This update_time is different from the Membership-level update_time since EndpointDetails are updated internally for API consumers. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`

	/* Output only. vCPU count as reported by Kubernetes nodes resources. */
	// +optional
	VcpuCount *int `json:"vcpuCount,omitempty"`
}

type MembershipKubernetesResourceStatus struct {
	/* Output only. The Kubernetes resources for installing the GKE Connect agent This field is only populated in the Membership returned from a successful long-running operation from CreateMembership or UpdateMembership. It is not populated during normal GetMembership or ListMemberships requests. To get the resource manifest after the initial registration, the caller should make a UpdateMembership call with an empty field mask. */
	// +optional
	ConnectResources []MembershipConnectResourcesStatus `json:"connectResources,omitempty"`

	/* Output only. Additional Kubernetes resources that need to be applied to the cluster after Membership creation, and after every update. This field is only populated in the Membership returned from a successful long-running operation from CreateMembership or UpdateMembership. It is not populated during normal GetMembership or ListMemberships requests. To get the resource manifest after the initial registration, the caller should make a UpdateMembership call with an empty field mask. */
	// +optional
	MembershipResources []MembershipMembershipResourcesStatus `json:"membershipResources,omitempty"`
}

type MembershipMembershipResourcesStatus struct {
	/* Whether the resource provided in the manifest is `cluster_scoped`. If unset, the manifest is assumed to be namespace scoped. This field is used for REST mapping when applying the resource in a cluster. */
	// +optional
	ClusterScoped *bool `json:"clusterScoped,omitempty"`

	/* YAML manifest of the resource. */
	// +optional
	Manifest *string `json:"manifest,omitempty"`
}

type MembershipStateStatus struct {
	/* Output only. The current state of the Membership resource. Possible values: CODE_UNSPECIFIED, CREATING, READY, DELETING, UPDATING, SERVICE_UPDATING */
	// +optional
	Code *string `json:"code,omitempty"`
}

type GKEHubMembershipStatus struct {
	/* Conditions represent the latest available observations of the
	   GKEHubMembership's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	// +optional
	Authority *MembershipAuthorityStatus `json:"authority,omitempty"`

	/* Output only. When the Membership was created. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Output only. When the Membership was deleted. */
	// +optional
	DeleteTime *string `json:"deleteTime,omitempty"`

	// +optional
	Endpoint *MembershipEndpointStatus `json:"endpoint,omitempty"`

	/* Output only. For clusters using Connect, the timestamp of the most recent connection established with Google Cloud. This time is updated every several minutes, not continuously. For clusters that do not use GKE Connect, or that have never connected successfully, this field will be unset. */
	// +optional
	LastConnectionTime *string `json:"lastConnectionTime,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* Output only. State of the Membership resource. */
	// +optional
	State *MembershipStateStatus `json:"state,omitempty"`

	/* Output only. Google-generated UUID for this resource. This is unique across all Membership resources. If a Membership resource is deleted and another resource with the same name is created, it gets a different unique_id. */
	// +optional
	UniqueId *string `json:"uniqueId,omitempty"`

	/* Output only. When the Membership was last updated. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgkehubmembership;gcpgkehubmemberships
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GKEHubMembership is the Schema for the gkehub API
// +k8s:openapi-gen=true
type GKEHubMembership struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GKEHubMembershipSpec   `json:"spec,omitempty"`
	Status GKEHubMembershipStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GKEHubMembershipList contains a list of GKEHubMembership
type GKEHubMembershipList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GKEHubMembership `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GKEHubMembership{}, &GKEHubMembershipList{})
}
