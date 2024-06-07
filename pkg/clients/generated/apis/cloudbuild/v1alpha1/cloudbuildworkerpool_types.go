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

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type WorkerpoolNetworkConfig struct {
	// +optional
	EgressOption *string `json:"egressOption,omitempty"`

	// +optional
	PeeredNetworkIpRange *string `json:"peeredNetworkIpRange,omitempty"`

	PeeredNetworkRef v1alpha1.ResourceRef `json:"peeredNetworkRef"`
}

type WorkerpoolPrivatePoolV1Config struct {
	// +optional
	NetworkConfig *WorkerpoolNetworkConfig `json:"networkConfig,omitempty"`

	WorkerConfig WorkerpoolWorkerConfig `json:"workerConfig"`
}

type WorkerpoolWorkerConfig struct {
	// +optional
	DiskSizeGb *int64 `json:"diskSizeGb,omitempty"`

	// +optional
	MachineType *string `json:"machineType,omitempty"`
}

type CloudBuildWorkerPoolSpec struct {
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	Location string `json:"location"`

	// +optional
	Name *string `json:"name,omitempty"`

	PrivatePoolV1Config WorkerpoolPrivatePoolV1Config `json:"privatePoolV1Config"`

	/* The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type WorkerpoolNetworkConfigStatus struct {
	// +optional
	EgressOption *string `json:"egressOption,omitempty"`

	// +optional
	PeeredNetwork *string `json:"peeredNetwork,omitempty"`

	// +optional
	PeeredNetworkIpRange *string `json:"peeredNetworkIpRange,omitempty"`
}

type WorkerpoolObservedStateStatus struct {
	/* The creation timestamp of the workerpool. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	// +optional
	NetworkConfig *WorkerpoolNetworkConfigStatus `json:"networkConfig,omitempty"`

	/* The last update timestamp of the workerpool. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`

	WorkerConfig WorkerpoolWorkerConfigStatus `json:"workerConfig"`
}

type WorkerpoolWorkerConfigStatus struct {
	// +optional
	DiskSizeGb *int64 `json:"diskSizeGb,omitempty"`

	// +optional
	MachineType *string `json:"machineType,omitempty"`
}

type CloudBuildWorkerPoolStatus struct {
	/* Conditions represent the latest available observations of the
	   CloudBuildWorkerPool's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* ObservedState is the state of the resource as most recently observed in GCP. */
	// +optional
	ObservedState *WorkerpoolObservedStateStatus `json:"observedState,omitempty"`

	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=,shortName=
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=beta"

// CloudBuildWorkerPool is the Schema for the cloudbuild API
// +k8s:openapi-gen=true
type CloudBuildWorkerPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudBuildWorkerPoolSpec   `json:"spec,omitempty"`
	Status CloudBuildWorkerPoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudBuildWorkerPoolList contains a list of CloudBuildWorkerPool
type CloudBuildWorkerPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudBuildWorkerPool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudBuildWorkerPool{}, &CloudBuildWorkerPoolList{})
}
