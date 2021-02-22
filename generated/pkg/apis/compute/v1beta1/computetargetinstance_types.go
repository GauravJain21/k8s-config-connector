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

type ComputeTargetInstanceSpec struct {
	/* Immutable. An optional description of this resource. */
	Description string `json:"description,omitempty"`
	/* The ComputeInstance handling traffic for this target instance. */
	InstanceRef v1alpha1.ResourceRef `json:"instanceRef,omitempty"`
	/* Immutable. NAT option controlling how IPs are NAT'ed to the instance.
	Currently only NO_NAT (default value) is supported. Default value: "NO_NAT" Possible values: ["NO_NAT"] */
	NatPolicy string `json:"natPolicy,omitempty"`
	/* The network this target instance uses to forward
	traffic. If not specified, the traffic will be forwarded to the network
	that the default network interface belongs to. */
	NetworkRef v1alpha1.ResourceRef `json:"networkRef,omitempty"`
	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	ResourceID string `json:"resourceID,omitempty"`
	/* Immutable. URL of the zone where the target instance resides. */
	Zone string `json:"zone,omitempty"`
}

type ComputeTargetInstanceStatus struct {
	/* Conditions represents the latest available observations of the
	   ComputeTargetInstance's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Creation timestamp in RFC3339 text format. */
	CreationTimestamp string `json:"creationTimestamp,omitempty"`
	/*  */
	SelfLink string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeTargetInstance is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeTargetInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeTargetInstanceSpec   `json:"spec,omitempty"`
	Status ComputeTargetInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeTargetInstanceList contains a list of ComputeTargetInstance
type ComputeTargetInstanceList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []ComputeTargetInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeTargetInstance{}, &ComputeTargetInstanceList{})
}
