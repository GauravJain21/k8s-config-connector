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

type ComputerouternatLogConfig struct {
	/* Indicates whether or not to export logs. */
	Enable bool `json:"enable,omitempty"`
	/* Specifies the desired filtering of logs on this NAT. Possible values: ["ERRORS_ONLY", "TRANSLATIONS_ONLY", "ALL"] */
	Filter string `json:"filter,omitempty"`
}

type Subnetwork struct {
	/* List of the secondary ranges of the subnetwork that are allowed
	to use NAT. This can be populated only if
	'LIST_OF_SECONDARY_IP_RANGES' is one of the values in
	sourceIpRangesToNat */
	SecondaryIpRangeNames []string `json:"secondaryIpRangeNames,omitempty"`
	/* List of options for which source IPs in the subnetwork
	should have NAT enabled. Supported values include:
	'ALL_IP_RANGES', 'LIST_OF_SECONDARY_IP_RANGES',
	'PRIMARY_IP_RANGE'. */
	SourceIpRangesToNat []string `json:"sourceIpRangesToNat,omitempty"`
	/* The subnetwork to NAT. */
	SubnetworkRef v1alpha1.ResourceRef `json:"subnetworkRef,omitempty"`
}

type ComputeRouterNATSpec struct {
	/*  */
	DrainNatIps []v1alpha1.ResourceRef `json:"drainNatIps,omitempty"`
	/* Specifies if endpoint independent mapping is enabled. This is enabled by default. For more information
	see the [official documentation](https://cloud.google.com/nat/docs/overview#specs-rfcs). */
	EnableEndpointIndependentMapping bool `json:"enableEndpointIndependentMapping,omitempty"`
	/* Timeout (in seconds) for ICMP connections. Defaults to 30s if not set. */
	IcmpIdleTimeoutSec int `json:"icmpIdleTimeoutSec,omitempty"`
	/* Configuration for logging on NAT */
	LogConfig ComputerouternatLogConfig `json:"logConfig,omitempty"`
	/* Minimum number of ports allocated to a VM from this NAT. */
	MinPortsPerVm int `json:"minPortsPerVm,omitempty"`
	/* How external IPs should be allocated for this NAT. Valid values are
	'AUTO_ONLY' for only allowing NAT IPs allocated by Google Cloud
	Platform, or 'MANUAL_ONLY' for only user-allocated NAT IP addresses. Possible values: ["MANUAL_ONLY", "AUTO_ONLY"] */
	NatIpAllocateOption string `json:"natIpAllocateOption,omitempty"`
	/*  */
	NatIps []v1alpha1.ResourceRef `json:"natIps,omitempty"`
	/* Immutable. Region where the router and NAT reside. */
	Region string `json:"region,omitempty"`
	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	ResourceID string `json:"resourceID,omitempty"`
	/* The Cloud Router in which this NAT will be configured. */
	RouterRef v1alpha1.ResourceRef `json:"routerRef,omitempty"`
	/* How NAT should be configured per Subnetwork.
	If 'ALL_SUBNETWORKS_ALL_IP_RANGES', all of the
	IP ranges in every Subnetwork are allowed to Nat.
	If 'ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES', all of the primary IP
	ranges in every Subnetwork are allowed to Nat.
	'LIST_OF_SUBNETWORKS': A list of Subnetworks are allowed to Nat
	(specified in the field subnetwork below). Note that if this field
	contains ALL_SUBNETWORKS_ALL_IP_RANGES or
	ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES, then there should not be any
	other RouterNat section in any Router for this network in this region. Possible values: ["ALL_SUBNETWORKS_ALL_IP_RANGES", "ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES", "LIST_OF_SUBNETWORKS"] */
	SourceSubnetworkIpRangesToNat string `json:"sourceSubnetworkIpRangesToNat,omitempty"`
	/* One or more subnetwork NAT configurations. Only used if
	'source_subnetwork_ip_ranges_to_nat' is set to 'LIST_OF_SUBNETWORKS' */
	Subnetwork []Subnetwork `json:"subnetwork,omitempty"`
	/* Timeout (in seconds) for TCP established connections.
	Defaults to 1200s if not set. */
	TcpEstablishedIdleTimeoutSec int `json:"tcpEstablishedIdleTimeoutSec,omitempty"`
	/* Timeout (in seconds) for TCP transitory connections.
	Defaults to 30s if not set. */
	TcpTransitoryIdleTimeoutSec int `json:"tcpTransitoryIdleTimeoutSec,omitempty"`
	/* Timeout (in seconds) for UDP connections. Defaults to 30s if not set. */
	UdpIdleTimeoutSec int `json:"udpIdleTimeoutSec,omitempty"`
}

type ComputeRouterNATStatus struct {
	/* Conditions represents the latest available observations of the
	   ComputeRouterNAT's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeRouterNAT is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeRouterNAT struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeRouterNATSpec   `json:"spec,omitempty"`
	Status ComputeRouterNATStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeRouterNATList contains a list of ComputeRouterNAT
type ComputeRouterNATList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []ComputeRouterNAT `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeRouterNAT{}, &ComputeRouterNATList{})
}
