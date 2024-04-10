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

type InboundsamlconfigIdpCertificates struct {
	/* The IdP's x509 certificate. */
	// +optional
	X509Certificate *string `json:"x509Certificate,omitempty"`
}

type InboundsamlconfigIdpConfig struct {
	/* The IdP's certificate data to verify the signature in the SAMLResponse issued by the IDP. */
	IdpCertificates []InboundsamlconfigIdpCertificates `json:"idpCertificates"`

	/* Unique identifier for all SAML entities. */
	IdpEntityId string `json:"idpEntityId"`

	/* Indicates if outbounding SAMLRequest should be signed. */
	// +optional
	SignRequest *bool `json:"signRequest,omitempty"`

	/* URL to send Authentication request to. */
	SsoUrl string `json:"ssoUrl"`
}

type InboundsamlconfigSpCertificates struct {
	/* The x509 certificate. */
	// +optional
	X509Certificate *string `json:"x509Certificate,omitempty"`
}

type InboundsamlconfigSpConfig struct {
	/* Callback URI where responses from IDP are handled. Must start with 'https://'. */
	// +optional
	CallbackUri *string `json:"callbackUri,omitempty"`

	/* The IDP's certificate data to verify the signature in the SAMLResponse issued by the IDP. */
	// +optional
	SpCertificates []InboundsamlconfigSpCertificates `json:"spCertificates,omitempty"`

	/* Unique identifier for all SAML entities. */
	// +optional
	SpEntityId *string `json:"spEntityId,omitempty"`
}

type IdentityPlatformInboundSAMLConfigSpec struct {
	/* Human friendly display name. */
	DisplayName string `json:"displayName"`

	/* If this config allows users to sign in with the provider. */
	// +optional
	Enabled *bool `json:"enabled,omitempty"`

	/* SAML IdP configuration when the project acts as the relying party. */
	IdpConfig InboundsamlconfigIdpConfig `json:"idpConfig"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* SAML SP (Service Provider) configuration when the project acts as the relying party to receive
	and accept an authentication assertion issued by a SAML identity provider. */
	SpConfig InboundsamlconfigSpConfig `json:"spConfig"`
}

type IdentityPlatformInboundSAMLConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   IdentityPlatformInboundSAMLConfig's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpidentityplatforminboundsamlconfig;gcpidentityplatforminboundsamlconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// IdentityPlatformInboundSAMLConfig is the Schema for the identityplatform API
// +k8s:openapi-gen=true
type IdentityPlatformInboundSAMLConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IdentityPlatformInboundSAMLConfigSpec   `json:"spec,omitempty"`
	Status IdentityPlatformInboundSAMLConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IdentityPlatformInboundSAMLConfigList contains a list of IdentityPlatformInboundSAMLConfig
type IdentityPlatformInboundSAMLConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityPlatformInboundSAMLConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IdentityPlatformInboundSAMLConfig{}, &IdentityPlatformInboundSAMLConfigList{})
}
