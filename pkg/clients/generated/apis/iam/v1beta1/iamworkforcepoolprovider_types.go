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

type WorkforcepoolproviderClientSecret struct {
	/* The value of the client secret. */
	// +optional
	Value *WorkforcepoolproviderValue `json:"value,omitempty"`
}

type WorkforcepoolproviderOidc struct {
	/* Required. The client ID. Must match the audience claim of the JWT issued by the identity provider. */
	ClientId string `json:"clientId"`

	/* The optional client secret. Required to enable Authorization Code flow for web sign-in. */
	// +optional
	ClientSecret *WorkforcepoolproviderClientSecret `json:"clientSecret,omitempty"`

	/* Required. The OIDC issuer URI. Must be a valid URI using the 'https' scheme. */
	IssuerUri string `json:"issuerUri"`

	/* OIDC JWKs in JSON String format. For details on definition of a JWK, see https:tools.ietf.org/html/rfc7517. If not set, then we use the `jwks_uri` from the discovery document fetched from the .well-known path for the `issuer_uri`. Currently, RSA and EC asymmetric keys are supported. The JWK must use following format and include only the following fields: ```{"keys": [{"kty": "RSA/EC", "alg": "<algorithm>", "use": "sig", "kid": "<key-id>", "n": "", "e": "", "x": "", "y": "", "crv": ""}]}``` */
	// +optional
	JwksJson *string `json:"jwksJson,omitempty"`

	/* Required. Configuration for web single sign-on for the OIDC provider. Here, web sign-in refers to console sign-in and gcloud sign-in through the browser. */
	WebSsoConfig WorkforcepoolproviderWebSsoConfig `json:"webSsoConfig"`
}

type WorkforcepoolproviderPlainText struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	// +optional
	Value *string `json:"value,omitempty"`

	/* Source for the field's value. Cannot be used if 'value' is specified. */
	// +optional
	ValueFrom *WorkforcepoolproviderValueFrom `json:"valueFrom,omitempty"`
}

type WorkforcepoolproviderSaml struct {
	/* Required. SAML Identity provider configuration metadata xml doc. The xml document should comply with [SAML 2.0 specification](https://docs.oasis-open.org/security/saml/v2.0/saml-metadata-2.0-os.pdf). The max size of the acceptable xml document will be bounded to 128k characters. The metadata xml document should satisfy the following constraints: 1) Must contain an Identity Provider Entity ID. 2) Must contain at least one non-expired signing key certificate. 3) For each signing key: a) Valid from should be no more than 7 days from now. b) Valid to should be no more than 10 years in the future. 4) Up to 3 IdP signing keys are allowed in the metadata xml. When updating the provider's metadata xml, at least one non-expired signing key must overlap with the existing metadata. This requirement is skipped if there are no non-expired signing keys present in the existing metadata. */
	IdpMetadataXml string `json:"idpMetadataXml"`
}

type WorkforcepoolproviderValue struct {
	/* Input only. The plain text of the client secret value. */
	// +optional
	PlainText *WorkforcepoolproviderPlainText `json:"plainText,omitempty"`
}

type WorkforcepoolproviderValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	// +optional
	SecretKeyRef *v1alpha1.SecretKeyRef `json:"secretKeyRef,omitempty"`
}

type WorkforcepoolproviderWebSsoConfig struct {
	/* Additional scopes to request for in the OIDC authentication request on top of scopes requested by default. By default, the `openid`, `profile` and `email` scopes that are supported by the identity provider are requested. Each additional scope may be at most 256 characters. A maximum of 10 additional scopes may be configured. */
	// +optional
	AdditionalScopes []string `json:"additionalScopes,omitempty"`

	/* Required. The behavior for how OIDC Claims are included in the `assertion` object used for attribute mapping and attribute condition. Possible values: ASSERTION_CLAIMS_BEHAVIOR_UNSPECIFIED, MERGE_USER_INFO_OVER_ID_TOKEN_CLAIMS, ONLY_ID_TOKEN_CLAIMS */
	AssertionClaimsBehavior string `json:"assertionClaimsBehavior"`

	/* Required. The Response Type to request for in the OIDC Authorization Request for web sign-in. The `CODE` Response Type is recommended to avoid the Implicit Flow, for security reasons. Possible values: RESPONSE_TYPE_UNSPECIFIED, CODE, ID_TOKEN */
	ResponseType string `json:"responseType"`
}

type IAMWorkforcePoolProviderSpec struct {
	/* A [Common Expression Language](https://opensource.google/projects/cel) expression, in plain text, to restrict what otherwise valid authentication credentials issued by the provider should not be accepted. The expression must output a boolean representing whether to allow the federation. The following keywords may be referenced in the expressions: * `assertion`: JSON representing the authentication credential issued by the provider. * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`. `google.profile_photo` and `google.display_name` are not supported. * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`. The maximum length of the attribute condition expression is 4096 characters. If unspecified, all valid authentication credentials will be accepted. The following example shows how to only allow credentials with a mapped `google.groups` value of `admins`: ``` "'admins' in google.groups" ``` */
	// +optional
	AttributeCondition *string `json:"attributeCondition,omitempty"`

	/* Required. Maps attributes from the authentication credentials issued by an external identity provider to Google Cloud attributes, such as `subject` and `segment`. Each key must be a string specifying the Google Cloud IAM attribute to map to. The following keys are supported: * `google.subject`: The principal IAM is authenticating. You can reference this value in IAM bindings. This is also the subject that appears in Cloud Logging logs. This is a required field and the mapped subject cannot exceed 127 bytes. * `google.groups`: Groups the authenticating user belongs to. You can grant groups access to resources using an IAM `principalSet` binding; access applies to all members of the group. * `google.display_name`: The name of the authenticated user. This is an optional field and the mapped display name cannot exceed 100 bytes. If not set, `google.subject` will be displayed instead. This attribute cannot be referenced in IAM bindings. * `google.profile_photo`: The URL that specifies the authenticated user's thumbnail photo. This is an optional field. When set, the image will be visible as the user's profile picture. If not set, a generic user icon will be displayed instead. This attribute cannot be referenced in IAM bindings. You can also provide custom attributes by specifying `attribute.{custom_attribute}`, where {custom_attribute} is the name of the custom attribute to be mapped. You can define a maximum of 50 custom attributes. The maximum length of a mapped attribute key is 100 characters, and the key may only contain the characters [a-z0-9_]. You can reference these attributes in IAM policies to define fine-grained access for a workforce pool to Google Cloud resources. For example: */
	AttributeMapping map[string]string `json:"attributeMapping"`

	/* A user-specified description of the provider. Cannot exceed 256 characters. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Whether the provider is disabled. You cannot use a disabled provider to exchange tokens. However, existing tokens still grant access. */
	// +optional
	Disabled *bool `json:"disabled,omitempty"`

	/* A user-specified display name for the provider. Cannot exceed 32 characters. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* Immutable. The location for the resource */
	Location string `json:"location"`

	/* An OpenId Connect 1.0 identity provider configuration. */
	// +optional
	Oidc *WorkforcepoolproviderOidc `json:"oidc,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* A SAML identity provider configuration. */
	// +optional
	Saml *WorkforcepoolproviderSaml `json:"saml,omitempty"`

	/* Immutable. */
	WorkforcePoolRef v1alpha1.ResourceRef `json:"workforcePoolRef"`
}

type WorkforcepoolproviderClientSecretStatus struct {
	// +optional
	Value *WorkforcepoolproviderValueStatus `json:"value,omitempty"`
}

type WorkforcepoolproviderOidcStatus struct {
	// +optional
	ClientSecret *WorkforcepoolproviderClientSecretStatus `json:"clientSecret,omitempty"`
}

type WorkforcepoolproviderValueStatus struct {
	/* Output only. A thumbprint to represent the current client secret value. */
	// +optional
	Thumbprint *string `json:"thumbprint,omitempty"`
}

type IAMWorkforcePoolProviderStatus struct {
	/* Conditions represent the latest available observations of the
	   IAMWorkforcePoolProvider's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	// +optional
	Oidc *WorkforcepoolproviderOidcStatus `json:"oidc,omitempty"`

	/* Output only. The state of the provider. Possible values: STATE_UNSPECIFIED, ACTIVE, DELETED */
	// +optional
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpiamworkforcepoolprovider;gcpiamworkforcepoolproviders
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// IAMWorkforcePoolProvider is the Schema for the iam API
// +k8s:openapi-gen=true
type IAMWorkforcePoolProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMWorkforcePoolProviderSpec   `json:"spec,omitempty"`
	Status IAMWorkforcePoolProviderStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IAMWorkforcePoolProviderList contains a list of IAMWorkforcePoolProvider
type IAMWorkforcePoolProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IAMWorkforcePoolProvider `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IAMWorkforcePoolProvider{}, &IAMWorkforcePoolProviderList{})
}
