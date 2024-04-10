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

type ComputeRouteSpec struct {
	/* Immutable. An optional description of this resource. Provide this property
	when you create the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. The destination range of outgoing packets that this route applies to.
	Only IPv4 is supported. */
	DestRange string `json:"destRange"`

	/* The network that this route applies to. */
	NetworkRef v1alpha1.ResourceRef `json:"networkRef"`

	/* Immutable. URL to a gateway that should handle matching packets.
	Currently, you can only specify the internet gateway, using a full or
	partial valid URL:
	* 'https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway'
	* 'projects/project/global/gateways/default-internet-gateway'
	* 'global/gateways/default-internet-gateway'
	* The string 'default-internet-gateway'. */
	// +optional
	NextHopGateway *string `json:"nextHopGateway,omitempty"`

	/* A forwarding rule of type loadBalancingScheme=INTERNAL that should
	handle matching packets.  Note that this can only be used when the
	destinationRange is a public (non-RFC 1918) IP CIDR range. */
	// +optional
	NextHopILBRef *v1alpha1.ResourceRef `json:"nextHopILBRef,omitempty"`

	/* Instance that should handle matching packets. */
	// +optional
	NextHopInstanceRef *v1alpha1.ResourceRef `json:"nextHopInstanceRef,omitempty"`

	/* Immutable. Network IP address of an instance that should handle matching packets. */
	// +optional
	NextHopIp *string `json:"nextHopIp,omitempty"`

	/* The ComputeVPNTunnel that should handle matching packets */
	// +optional
	NextHopVPNTunnelRef *v1alpha1.ResourceRef `json:"nextHopVPNTunnelRef,omitempty"`

	/* Immutable. The priority of this route. Priority is used to break ties in cases
	where there is more than one matching route of equal prefix length.

	In the case of two routes with equal prefix length, the one with the
	lowest-numbered priority value wins.

	Default value is 1000. Valid range is 0 through 65535. */
	// +optional
	Priority *int `json:"priority,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. A list of instance tags to which this route applies. */
	// +optional
	Tags []string `json:"tags,omitempty"`
}

type ComputeRouteStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeRoute's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* URL to a Network that should handle matching packets. */
	// +optional
	NextHopNetwork *string `json:"nextHopNetwork,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeroute;gcpcomputeroutes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeRoute is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeRouteSpec   `json:"spec,omitempty"`
	Status ComputeRouteStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeRouteList contains a list of ComputeRoute
type ComputeRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeRoute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeRoute{}, &ComputeRouteList{})
}
