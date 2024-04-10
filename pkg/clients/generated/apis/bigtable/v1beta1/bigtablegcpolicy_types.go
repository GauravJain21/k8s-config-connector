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

type GcpolicyMaxAge struct {
	/* DEPRECATED. Deprecated in favor of duration. Immutable. Number of days before applying GC policy. */
	// +optional
	Days *int `json:"days,omitempty"`

	/* Immutable. Duration before applying GC policy. */
	// +optional
	Duration *string `json:"duration,omitempty"`
}

type GcpolicyMaxVersion struct {
	/* Immutable. Number of version before applying the GC policy. */
	Number int `json:"number"`
}

type BigtableGCPolicySpec struct {
	/* Immutable. The name of the column family. */
	ColumnFamily string `json:"columnFamily"`

	/* The deletion policy for the GC policy. Setting ABANDON allows the resource
	to be abandoned rather than deleted. This is useful for GC policy as it cannot be deleted
	in a replicated instance. Possible values are: "ABANDON". */
	// +optional
	DeletionPolicy *string `json:"deletionPolicy,omitempty"`

	/* Serialized JSON string for garbage collection policy. Conflicts with "mode", "max_age" and "max_version". */
	// +optional
	GcRules *string `json:"gcRules,omitempty"`

	/* The name of the Bigtable instance. */
	InstanceRef v1alpha1.ResourceRef `json:"instanceRef"`

	/* Immutable. NOTE: 'gc_rules' is more flexible, and should be preferred over this field for new resources. This field may be deprecated in the future. GC policy that applies to all cells older than the given age. */
	// +optional
	MaxAge []GcpolicyMaxAge `json:"maxAge,omitempty"`

	/* Immutable. NOTE: 'gc_rules' is more flexible, and should be preferred over this field for new resources. This field may be deprecated in the future. GC policy that applies to all versions of a cell except for the most recent. */
	// +optional
	MaxVersion []GcpolicyMaxVersion `json:"maxVersion,omitempty"`

	/* Immutable. NOTE: 'gc_rules' is more flexible, and should be preferred over this field for new resources. This field may be deprecated in the future. If multiple policies are set, you should choose between UNION OR INTERSECTION. */
	// +optional
	Mode *string `json:"mode,omitempty"`

	/* The name of the table. */
	TableRef v1alpha1.ResourceRef `json:"tableRef"`
}

type BigtableGCPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   BigtableGCPolicy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpbigtablegcpolicy;gcpbigtablegcpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// BigtableGCPolicy is the Schema for the bigtable API
// +k8s:openapi-gen=true
type BigtableGCPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BigtableGCPolicySpec   `json:"spec,omitempty"`
	Status BigtableGCPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BigtableGCPolicyList contains a list of BigtableGCPolicy
type BigtableGCPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BigtableGCPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BigtableGCPolicy{}, &BigtableGCPolicyList{})
}
