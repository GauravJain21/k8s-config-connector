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

type DeviceCredentials struct {
	/* The time at which this credential becomes invalid. */
	// +optional
	ExpirationTime *string `json:"expirationTime,omitempty"`

	/* A public key used to verify the signature of JSON Web Tokens (JWTs). */
	PublicKey DevicePublicKey `json:"publicKey"`
}

type DeviceGatewayConfig struct {
	/* Indicates whether the device is a gateway. Possible values: ["ASSOCIATION_ONLY", "DEVICE_AUTH_TOKEN_ONLY", "ASSOCIATION_AND_DEVICE_AUTH_TOKEN"]. */
	// +optional
	GatewayAuthMethod *string `json:"gatewayAuthMethod,omitempty"`

	/* Immutable. Indicates whether the device is a gateway. Default value: "NON_GATEWAY" Possible values: ["GATEWAY", "NON_GATEWAY"]. */
	// +optional
	GatewayType *string `json:"gatewayType,omitempty"`

	/* The ID of the gateway the device accessed most recently. */
	// +optional
	LastAccessedGatewayId *string `json:"lastAccessedGatewayId,omitempty"`

	/* The most recent time at which the device accessed the gateway specified in last_accessed_gateway. */
	// +optional
	LastAccessedGatewayTime *string `json:"lastAccessedGatewayTime,omitempty"`
}

type DevicePublicKey struct {
	/* The format of the key. Possible values: ["RSA_PEM", "RSA_X509_PEM", "ES256_PEM", "ES256_X509_PEM"]. */
	Format string `json:"format"`

	/* The key data. */
	Key string `json:"key"`
}

type CloudIOTDeviceSpec struct {
	/* If a device is blocked, connections or requests from this device will fail. */
	// +optional
	Blocked *bool `json:"blocked,omitempty"`

	/* The credentials used to authenticate this device. */
	// +optional
	Credentials []DeviceCredentials `json:"credentials,omitempty"`

	/* Gateway-related configuration and state. */
	// +optional
	GatewayConfig *DeviceGatewayConfig `json:"gatewayConfig,omitempty"`

	/* The logging verbosity for device activity. Possible values: ["NONE", "ERROR", "INFO", "DEBUG"]. */
	// +optional
	LogLevel *string `json:"logLevel,omitempty"`

	/* The metadata key-value pairs assigned to the device. */
	// +optional
	Metadata map[string]string `json:"metadata,omitempty"`

	/* Immutable. The name of the device registry where this device should be created. */
	Registry string `json:"registry"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type DeviceConfigStatus struct {
	/* The device configuration data. */
	// +optional
	BinaryData *string `json:"binaryData,omitempty"`

	/* The time at which this configuration version was updated in Cloud IoT Core. */
	// +optional
	CloudUpdateTime *string `json:"cloudUpdateTime,omitempty"`

	/* The time at which Cloud IoT Core received the acknowledgment from the device,
	indicating that the device has received this configuration version. */
	// +optional
	DeviceAckTime *string `json:"deviceAckTime,omitempty"`

	/* The version of this update. */
	// +optional
	Version *string `json:"version,omitempty"`
}

type DeviceDetailsStatus struct {
}

type DeviceLastErrorStatusStatus struct {
	/* A list of messages that carry the error details. */
	// +optional
	Details []DeviceDetailsStatus `json:"details,omitempty"`

	/* A developer-facing error message, which should be in English. */
	// +optional
	Message *string `json:"message,omitempty"`

	/* The status code, which should be an enum value of google.rpc.Code. */
	// +optional
	Number *int `json:"number,omitempty"`
}

type DeviceStateStatus struct {
	/* The device state data. */
	// +optional
	BinaryData *string `json:"binaryData,omitempty"`

	/* The time at which this state version was updated in Cloud IoT Core. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

type CloudIOTDeviceStatus struct {
	/* Conditions represent the latest available observations of the
	   CloudIOTDevice's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The most recent device configuration, which is eventually sent from Cloud IoT Core to the device. */
	// +optional
	Config []DeviceConfigStatus `json:"config,omitempty"`

	/* The last time a cloud-to-device config version acknowledgment was received from the device. */
	// +optional
	LastConfigAckTime *string `json:"lastConfigAckTime,omitempty"`

	/* The last time a cloud-to-device config version was sent to the device. */
	// +optional
	LastConfigSendTime *string `json:"lastConfigSendTime,omitempty"`

	/* The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub. */
	// +optional
	LastErrorStatus []DeviceLastErrorStatusStatus `json:"lastErrorStatus,omitempty"`

	/* The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub. */
	// +optional
	LastErrorTime *string `json:"lastErrorTime,omitempty"`

	/* The last time a telemetry event was received. */
	// +optional
	LastEventTime *string `json:"lastEventTime,omitempty"`

	/* The last time an MQTT PINGREQ was received. */
	// +optional
	LastHeartbeatTime *string `json:"lastHeartbeatTime,omitempty"`

	/* The last time a state event was received. */
	// +optional
	LastStateTime *string `json:"lastStateTime,omitempty"`

	/* A server-defined unique numeric ID for the device.
	This is a more compact way to identify devices, and it is globally unique. */
	// +optional
	NumId *string `json:"numId,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The state most recently received from the device. */
	// +optional
	State []DeviceStateStatus `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudiotdevice;gcpcloudiotdevices
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudIOTDevice is the Schema for the cloudiot API
// +k8s:openapi-gen=true
type CloudIOTDevice struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CloudIOTDeviceSpec   `json:"spec,omitempty"`
	Status CloudIOTDeviceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CloudIOTDeviceList contains a list of CloudIOTDevice
type CloudIOTDeviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudIOTDevice `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudIOTDevice{}, &CloudIOTDeviceList{})
}
