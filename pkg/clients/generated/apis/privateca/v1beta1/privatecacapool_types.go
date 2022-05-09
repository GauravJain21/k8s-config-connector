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

type CapoolAdditionalExtensions struct {
	/* Required. The parts of an OID path. The most significant parts of the path come first. */
	ObjectIdPath []int `json:"objectIdPath"`
}

type CapoolAllowedIssuanceModes struct {
	/* Optional. When true, allows callers to create Certificates by specifying a CertificateConfig. */
	// +optional
	AllowConfigBasedIssuance *bool `json:"allowConfigBasedIssuance,omitempty"`

	/* Optional. When true, allows callers to create Certificates by specifying a CSR. */
	// +optional
	AllowCsrBasedIssuance *bool `json:"allowCsrBasedIssuance,omitempty"`
}

type CapoolAllowedKeyTypes struct {
	/* Represents an allowed Elliptic Curve key type. */
	// +optional
	EllipticCurve *CapoolEllipticCurve `json:"ellipticCurve,omitempty"`

	/* Represents an allowed RSA key type. */
	// +optional
	Rsa *CapoolRsa `json:"rsa,omitempty"`
}

type CapoolBaseKeyUsage struct {
	/* The key may be used to sign certificates. */
	// +optional
	CertSign *bool `json:"certSign,omitempty"`

	/* The key may be used for cryptographic commitments. Note that this may also be referred to as "non-repudiation". */
	// +optional
	ContentCommitment *bool `json:"contentCommitment,omitempty"`

	/* The key may be used sign certificate revocation lists. */
	// +optional
	CrlSign *bool `json:"crlSign,omitempty"`

	/* The key may be used to encipher data. */
	// +optional
	DataEncipherment *bool `json:"dataEncipherment,omitempty"`

	/* The key may be used to decipher only. */
	// +optional
	DecipherOnly *bool `json:"decipherOnly,omitempty"`

	/* The key may be used for digital signatures. */
	// +optional
	DigitalSignature *bool `json:"digitalSignature,omitempty"`

	/* The key may be used to encipher only. */
	// +optional
	EncipherOnly *bool `json:"encipherOnly,omitempty"`

	/* The key may be used in a key agreement protocol. */
	// +optional
	KeyAgreement *bool `json:"keyAgreement,omitempty"`

	/* The key may be used to encipher other keys. */
	// +optional
	KeyEncipherment *bool `json:"keyEncipherment,omitempty"`
}

type CapoolBaselineValues struct {
	/* Optional. Describes custom X.509 extensions. */
	// +optional
	AdditionalExtensions []CapoolAdditionalExtensions `json:"additionalExtensions,omitempty"`

	/* Optional. Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the "Authority Information Access" extension in the certificate. */
	// +optional
	AiaOcspServers []string `json:"aiaOcspServers,omitempty"`

	/* Optional. Describes options in this X509Parameters that are relevant in a CA certificate. */
	// +optional
	CaOptions *CapoolCaOptions `json:"caOptions,omitempty"`

	/* Optional. Indicates the intended use for keys that correspond to a certificate. */
	// +optional
	KeyUsage *CapoolKeyUsage `json:"keyUsage,omitempty"`

	/* Optional. Describes the X.509 certificate policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4. */
	// +optional
	PolicyIds []CapoolPolicyIds `json:"policyIds,omitempty"`
}

type CapoolCaOptions struct {
	/* Optional. Refers to the "CA" X.509 extension, which is a boolean value. When this value is missing, the extension will be omitted from the CA certificate. */
	// +optional
	IsCa *bool `json:"isCa,omitempty"`

	/* Optional. Refers to the path length restriction X.509 extension. For a CA certificate, this value describes the depth of subordinate CA certificates that are allowed. If this value is less than 0, the request will fail. If this value is missing, the max path length will be omitted from the CA certificate. */
	// +optional
	MaxIssuerPathLength *int `json:"maxIssuerPathLength,omitempty"`
}

type CapoolCelExpression struct {
	/* Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Textual representation of an expression in Common Expression Language syntax. */
	// +optional
	Expression *string `json:"expression,omitempty"`

	/* Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file. */
	// +optional
	Location *string `json:"location,omitempty"`

	/* Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression. */
	// +optional
	Title *string `json:"title,omitempty"`
}

type CapoolEllipticCurve struct {
	/* Optional. A signature algorithm that must be used. If this is omitted, any EC-based signature algorithm will be allowed. Possible values: EC_SIGNATURE_ALGORITHM_UNSPECIFIED, ECDSA_P256, ECDSA_P384, EDDSA_25519 */
	// +optional
	SignatureAlgorithm *string `json:"signatureAlgorithm,omitempty"`
}

type CapoolExtendedKeyUsage struct {
	/* Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as "TLS WWW client authentication", though regularly used for non-WWW TLS. */
	// +optional
	ClientAuth *bool `json:"clientAuth,omitempty"`

	/* Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as "Signing of downloadable executable code client authentication". */
	// +optional
	CodeSigning *bool `json:"codeSigning,omitempty"`

	/* Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as "Email protection". */
	// +optional
	EmailProtection *bool `json:"emailProtection,omitempty"`

	/* Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as "Signing OCSP responses". */
	// +optional
	OcspSigning *bool `json:"ocspSigning,omitempty"`

	/* Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as "TLS WWW server authentication", though regularly used for non-WWW TLS. */
	// +optional
	ServerAuth *bool `json:"serverAuth,omitempty"`

	/* Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as "Binding the hash of an object to a time". */
	// +optional
	TimeStamping *bool `json:"timeStamping,omitempty"`
}

type CapoolIdentityConstraints struct {
	/* Required. If this is true, the SubjectAltNames extension may be copied from a certificate request into the signed certificate. Otherwise, the requested SubjectAltNames will be discarded. */
	AllowSubjectAltNamesPassthrough bool `json:"allowSubjectAltNamesPassthrough"`

	/* Required. If this is true, the Subject field may be copied from a certificate request into the signed certificate. Otherwise, the requested Subject will be discarded. */
	AllowSubjectPassthrough bool `json:"allowSubjectPassthrough"`

	/* Optional. A CEL expression that may be used to validate the resolved X.509 Subject and/or Subject Alternative Name before a certificate is signed. To see the full allowed syntax and some examples, see https://cloud.google.com/certificate-authority-service/docs/using-cel */
	// +optional
	CelExpression *CapoolCelExpression `json:"celExpression,omitempty"`
}

type CapoolIssuancePolicy struct {
	/* Optional. If specified, then only methods allowed in the IssuanceModes may be used to issue Certificates. */
	// +optional
	AllowedIssuanceModes *CapoolAllowedIssuanceModes `json:"allowedIssuanceModes,omitempty"`

	/* Optional. If any AllowedKeyType is specified, then the certificate request's public key must match one of the key types listed here. Otherwise, any key may be used. */
	// +optional
	AllowedKeyTypes []CapoolAllowedKeyTypes `json:"allowedKeyTypes,omitempty"`

	/* Optional. A set of X.509 values that will be applied to all certificates issued through this CaPool. If a certificate request includes conflicting values for the same properties, they will be overwritten by the values defined here. If a certificate request uses a CertificateTemplate that defines conflicting predefined_values for the same properties, the certificate issuance request will fail. */
	// +optional
	BaselineValues *CapoolBaselineValues `json:"baselineValues,omitempty"`

	/* Optional. Describes constraints on identities that may appear in Certificates issued through this CaPool. If this is omitted, then this CaPool will not add restrictions on a certificate's identity. */
	// +optional
	IdentityConstraints *CapoolIdentityConstraints `json:"identityConstraints,omitempty"`

	/* Optional. The maximum lifetime allowed for issued Certificates. Note that if the issuing CertificateAuthority expires before a Certificate's requested maximum_lifetime, the effective lifetime will be explicitly truncated to match it. */
	// +optional
	MaximumLifetime *string `json:"maximumLifetime,omitempty"`

	/* Optional. Describes the set of X.509 extensions that may appear in a Certificate issued through this CaPool. If a certificate request sets extensions that don't appear in the passthrough_extensions, those extensions will be dropped. If a certificate request uses a CertificateTemplate with predefined_values that don't appear here, the certificate issuance request will fail. If this is omitted, then this CaPool will not add restrictions on a certificate's X.509 extensions. These constraints do not apply to X.509 extensions set in this CaPool's baseline_values. */
	// +optional
	PassthroughExtensions *CapoolPassthroughExtensions `json:"passthroughExtensions,omitempty"`
}

type CapoolKeyUsage struct {
	/* Describes high-level ways in which a key may be used. */
	// +optional
	BaseKeyUsage *CapoolBaseKeyUsage `json:"baseKeyUsage,omitempty"`

	/* Detailed scenarios in which a key may be used. */
	// +optional
	ExtendedKeyUsage *CapoolExtendedKeyUsage `json:"extendedKeyUsage,omitempty"`

	/* Used to describe extended key usages that are not listed in the KeyUsage.ExtendedKeyUsageOptions message. */
	// +optional
	UnknownExtendedKeyUsages []CapoolUnknownExtendedKeyUsages `json:"unknownExtendedKeyUsages,omitempty"`
}

type CapoolObjectId struct {
	/* Required. The parts of an OID path. The most significant parts of the path come first. */
	ObjectIdPath []int `json:"objectIdPath"`
}

type CapoolPassthroughExtensions struct {
	/* Optional. A set of ObjectIds identifying custom X.509 extensions. Will be combined with known_extensions to determine the full set of X.509 extensions. */
	// +optional
	AdditionalExtensions []CapoolAdditionalExtensions `json:"additionalExtensions,omitempty"`

	/* Optional. A set of named X.509 extensions. Will be combined with additional_extensions to determine the full set of X.509 extensions. */
	// +optional
	KnownExtensions []string `json:"knownExtensions,omitempty"`
}

type CapoolPolicyIds struct {
	/* Required. The parts of an OID path. The most significant parts of the path come first. */
	ObjectIdPath []int `json:"objectIdPath"`
}

type CapoolPublishingOptions struct {
	/* Optional. When true, publishes each CertificateAuthority's CA certificate and includes its URL in the "Authority Information Access" X.509 extension in all issued Certificates. If this is false, the CA certificate will not be published and the corresponding X.509 extension will not be written in issued certificates. */
	// +optional
	PublishCaCert *bool `json:"publishCaCert,omitempty"`

	/* Optional. When true, publishes each CertificateAuthority's CRL and includes its URL in the "CRL Distribution Points" X.509 extension in all issued Certificates. If this is false, CRLs will not be published and the corresponding X.509 extension will not be written in issued certificates. CRLs will expire 7 days from their creation. However, we will rebuild daily. CRLs are also rebuilt shortly after a certificate is revoked. */
	// +optional
	PublishCrl *bool `json:"publishCrl,omitempty"`
}

type CapoolRsa struct {
	/* Optional. The maximum allowed RSA modulus size, in bits. If this is not set, or if set to zero, the service will not enforce an explicit upper bound on RSA modulus sizes. */
	// +optional
	MaxModulusSize *int `json:"maxModulusSize,omitempty"`

	/* Optional. The minimum allowed RSA modulus size, in bits. If this is not set, or if set to zero, the service-level min RSA modulus size will continue to apply. */
	// +optional
	MinModulusSize *int `json:"minModulusSize,omitempty"`
}

type CapoolUnknownExtendedKeyUsages struct {
	/* Required. The parts of an OID path. The most significant parts of the path come first. */
	ObjectIdPath []int `json:"objectIdPath"`
}

type PrivateCACAPoolSpec struct {
	/* Optional. The IssuancePolicy to control how Certificates will be issued from this CaPool. */
	// +optional
	IssuancePolicy *CapoolIssuancePolicy `json:"issuancePolicy,omitempty"`

	/* Immutable. The location for the resource */
	Location string `json:"location"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Optional. The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool. */
	// +optional
	PublishingOptions *CapoolPublishingOptions `json:"publishingOptions,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. Required. Immutable. The Tier of this CaPool. Possible values: TIER_UNSPECIFIED, ENTERPRISE, DEVOPS */
	Tier string `json:"tier"`
}

type PrivateCACAPoolStatus struct {
	/* Conditions represent the latest available observations of the
	   PrivateCACAPool's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PrivateCACAPool is the Schema for the privateca API
// +k8s:openapi-gen=true
type PrivateCACAPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PrivateCACAPoolSpec   `json:"spec,omitempty"`
	Status PrivateCACAPoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PrivateCACAPoolList contains a list of PrivateCACAPool
type PrivateCACAPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateCACAPool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PrivateCACAPool{}, &PrivateCACAPoolList{})
}
