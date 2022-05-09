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

type CertificateauthorityAdditionalExtensions struct {
	/* Immutable. Optional. Indicates whether or not this extension is critical (i.e., if the client does not know how to handle this extension, the client should consider this to be an error). */
	// +optional
	Critical *bool `json:"critical,omitempty"`

	/* Immutable. Required. The OID for this X.509 extension. */
	ObjectId CertificateauthorityObjectId `json:"objectId"`

	/* Immutable. Required. The value of this X.509 extension. */
	Value string `json:"value"`
}

type CertificateauthorityBaseKeyUsage struct {
	/* Immutable. The key may be used to sign certificates. */
	// +optional
	CertSign *bool `json:"certSign,omitempty"`

	/* Immutable. The key may be used for cryptographic commitments. Note that this may also be referred to as "non-repudiation". */
	// +optional
	ContentCommitment *bool `json:"contentCommitment,omitempty"`

	/* Immutable. The key may be used sign certificate revocation lists. */
	// +optional
	CrlSign *bool `json:"crlSign,omitempty"`

	/* Immutable. The key may be used to encipher data. */
	// +optional
	DataEncipherment *bool `json:"dataEncipherment,omitempty"`

	/* Immutable. The key may be used to decipher only. */
	// +optional
	DecipherOnly *bool `json:"decipherOnly,omitempty"`

	/* Immutable. The key may be used for digital signatures. */
	// +optional
	DigitalSignature *bool `json:"digitalSignature,omitempty"`

	/* Immutable. The key may be used to encipher only. */
	// +optional
	EncipherOnly *bool `json:"encipherOnly,omitempty"`

	/* Immutable. The key may be used in a key agreement protocol. */
	// +optional
	KeyAgreement *bool `json:"keyAgreement,omitempty"`

	/* Immutable. The key may be used to encipher other keys. */
	// +optional
	KeyEncipherment *bool `json:"keyEncipherment,omitempty"`
}

type CertificateauthorityCaOptions struct {
	/* Immutable. Optional. Refers to the "CA" X.509 extension, which is a boolean value. When this value is missing, the extension will be omitted from the CA certificate. */
	// +optional
	IsCa *bool `json:"isCa,omitempty"`

	/* Immutable. Optional. Refers to the path length restriction X.509 extension. For a CA certificate, this value describes the depth of subordinate CA certificates that are allowed. If this value is less than 0, the request will fail. If this value is missing, the max path length will be omitted from the CA certificate. */
	// +optional
	MaxIssuerPathLength *int `json:"maxIssuerPathLength,omitempty"`
}

type CertificateauthorityConfig struct {
	/* Immutable. Required. Specifies some of the values in a certificate that are related to the subject. */
	SubjectConfig CertificateauthoritySubjectConfig `json:"subjectConfig"`

	/* Immutable. Required. Describes how some of the technical X.509 fields in a certificate should be populated. */
	X509Config CertificateauthorityX509Config `json:"x509Config"`
}

type CertificateauthorityCustomSans struct {
	/* Immutable. Optional. Indicates whether or not this extension is critical (i.e., if the client does not know how to handle this extension, the client should consider this to be an error). */
	// +optional
	Critical *bool `json:"critical,omitempty"`

	/* Immutable. Required. The OID for this X.509 extension. */
	ObjectId CertificateauthorityObjectId `json:"objectId"`

	/* Immutable. Required. The value of this X.509 extension. */
	Value string `json:"value"`
}

type CertificateauthorityExtendedKeyUsage struct {
	/* Immutable. Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as "TLS WWW client authentication", though regularly used for non-WWW TLS. */
	// +optional
	ClientAuth *bool `json:"clientAuth,omitempty"`

	/* Immutable. Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as "Signing of downloadable executable code client authentication". */
	// +optional
	CodeSigning *bool `json:"codeSigning,omitempty"`

	/* Immutable. Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as "Email protection". */
	// +optional
	EmailProtection *bool `json:"emailProtection,omitempty"`

	/* Immutable. Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as "Signing OCSP responses". */
	// +optional
	OcspSigning *bool `json:"ocspSigning,omitempty"`

	/* Immutable. Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as "TLS WWW server authentication", though regularly used for non-WWW TLS. */
	// +optional
	ServerAuth *bool `json:"serverAuth,omitempty"`

	/* Immutable. Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as "Binding the hash of an object to a time". */
	// +optional
	TimeStamping *bool `json:"timeStamping,omitempty"`
}

type CertificateauthorityKeySpec struct {
	/* Immutable. The algorithm to use for creating a managed Cloud KMS key for a for a simplified experience. All managed keys will be have their ProtectionLevel as `HSM`. Possible values: RSA_PSS_2048_SHA256, RSA_PSS_3072_SHA256, RSA_PSS_4096_SHA256, RSA_PKCS1_2048_SHA256, RSA_PKCS1_3072_SHA256, RSA_PKCS1_4096_SHA256, EC_P256_SHA256, EC_P384_SHA384 */
	// +optional
	Algorithm *string `json:"algorithm,omitempty"`

	/* Immutable. */
	// +optional
	CloudKmsKeyVersionRef *v1alpha1.ResourceRef `json:"cloudKmsKeyVersionRef,omitempty"`
}

type CertificateauthorityKeyUsage struct {
	/* Immutable. Describes high-level ways in which a key may be used. */
	// +optional
	BaseKeyUsage *CertificateauthorityBaseKeyUsage `json:"baseKeyUsage,omitempty"`

	/* Immutable. Detailed scenarios in which a key may be used. */
	// +optional
	ExtendedKeyUsage *CertificateauthorityExtendedKeyUsage `json:"extendedKeyUsage,omitempty"`

	/* Immutable. Used to describe extended key usages that are not listed in the KeyUsage.ExtendedKeyUsageOptions message. */
	// +optional
	UnknownExtendedKeyUsages []CertificateauthorityUnknownExtendedKeyUsages `json:"unknownExtendedKeyUsages,omitempty"`
}

type CertificateauthorityObjectId struct {
	/* Immutable. Required. The parts of an OID path. The most significant parts of the path come first. */
	ObjectIdPath []int `json:"objectIdPath"`
}

type CertificateauthorityPolicyIds struct {
	/* Immutable. Required. The parts of an OID path. The most significant parts of the path come first. */
	ObjectIdPath []int `json:"objectIdPath"`
}

type CertificateauthoritySubject struct {
	/* Immutable. The "common name" of the subject. */
	// +optional
	CommonName *string `json:"commonName,omitempty"`

	/* Immutable. The country code of the subject. */
	// +optional
	CountryCode *string `json:"countryCode,omitempty"`

	/* Immutable. The locality or city of the subject. */
	// +optional
	Locality *string `json:"locality,omitempty"`

	/* Immutable. The organization of the subject. */
	// +optional
	Organization *string `json:"organization,omitempty"`

	/* Immutable. The organizational_unit of the subject. */
	// +optional
	OrganizationalUnit *string `json:"organizationalUnit,omitempty"`

	/* Immutable. The postal code of the subject. */
	// +optional
	PostalCode *string `json:"postalCode,omitempty"`

	/* Immutable. The province, territory, or regional state of the subject. */
	// +optional
	Province *string `json:"province,omitempty"`

	/* Immutable. The street address of the subject. */
	// +optional
	StreetAddress *string `json:"streetAddress,omitempty"`
}

type CertificateauthoritySubjectAltName struct {
	/* Immutable. Contains additional subject alternative name values. */
	// +optional
	CustomSans []CertificateauthorityCustomSans `json:"customSans,omitempty"`

	/* Immutable. Contains only valid, fully-qualified host names. */
	// +optional
	DnsNames []string `json:"dnsNames,omitempty"`

	/* Immutable. Contains only valid RFC 2822 E-mail addresses. */
	// +optional
	EmailAddresses []string `json:"emailAddresses,omitempty"`

	/* Immutable. Contains only valid 32-bit IPv4 addresses or RFC 4291 IPv6 addresses. */
	// +optional
	IpAddresses []string `json:"ipAddresses,omitempty"`

	/* Immutable. Contains only valid RFC 3986 URIs. */
	// +optional
	Uris []string `json:"uris,omitempty"`
}

type CertificateauthoritySubjectConfig struct {
	/* Immutable. Required. Contains distinguished name fields such as the common name, location and organization. */
	Subject CertificateauthoritySubject `json:"subject"`

	/* Immutable. Optional. The subject alternative name fields. */
	// +optional
	SubjectAltName *CertificateauthoritySubjectAltName `json:"subjectAltName,omitempty"`
}

type CertificateauthorityUnknownExtendedKeyUsages struct {
	/* Immutable. Required. The parts of an OID path. The most significant parts of the path come first. */
	ObjectIdPath []int `json:"objectIdPath"`
}

type CertificateauthorityX509Config struct {
	/* Immutable. Optional. Describes custom X.509 extensions. */
	// +optional
	AdditionalExtensions []CertificateauthorityAdditionalExtensions `json:"additionalExtensions,omitempty"`

	/* Immutable. Optional. Describes options in this X509Parameters that are relevant in a CA certificate. */
	// +optional
	CaOptions *CertificateauthorityCaOptions `json:"caOptions,omitempty"`

	/* Immutable. Optional. Indicates the intended use for keys that correspond to a certificate. */
	// +optional
	KeyUsage *CertificateauthorityKeyUsage `json:"keyUsage,omitempty"`

	/* Immutable. Optional. Describes the X.509 certificate policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4. */
	// +optional
	PolicyIds []CertificateauthorityPolicyIds `json:"policyIds,omitempty"`
}

type PrivateCACertificateAuthoritySpec struct {
	/* Immutable. */
	CaPoolRef v1alpha1.ResourceRef `json:"caPoolRef"`

	/* Immutable. Required. Immutable. The config used to create a self-signed X.509 certificate or CSR. */
	Config CertificateauthorityConfig `json:"config"`

	/* Immutable. */
	// +optional
	GcsBucketRef *v1alpha1.ResourceRef `json:"gcsBucketRef,omitempty"`

	/* Immutable. Required. Immutable. Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA certificate. Otherwise, it is used to sign a CSR. */
	KeySpec CertificateauthorityKeySpec `json:"keySpec"`

	/* Immutable. Required. The desired lifetime of the CA certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate. */
	Lifetime string `json:"lifetime"`

	/* Immutable. The location for the resource */
	Location string `json:"location"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. Required. Immutable. The Type of this CertificateAuthority. Possible values: SELF_SIGNED, SUBORDINATE */
	Type string `json:"type"`
}

type CertificateauthorityAccessUrlsStatus struct {
	/* The URL where this CertificateAuthority's CA certificate is published. This will only be set for CAs that have been activated. */
	CaCertificateAccessUrl string `json:"caCertificateAccessUrl,omitempty"`

	/* The URLs where this CertificateAuthority's CRLs are published. This will only be set for CAs that have been activated. */
	CrlAccessUrls []string `json:"crlAccessUrls,omitempty"`
}

type CertificateauthorityAdditionalExtensionsStatus struct {
	/* Optional. Indicates whether or not this extension is critical (i.e., if the client does not know how to handle this extension, the client should consider this to be an error). */
	Critical bool `json:"critical,omitempty"`

	/* Required. The OID for this X.509 extension. */
	ObjectId CertificateauthorityObjectIdStatus `json:"objectId,omitempty"`

	/* Required. The value of this X.509 extension. */
	Value string `json:"value,omitempty"`
}

type CertificateauthorityAuthorityKeyIdStatus struct {
	/* Optional. The value of this KeyId encoded in lowercase hexadecimal. This is most likely the 160 bit SHA-1 hash of the public key. */
	KeyId string `json:"keyId,omitempty"`
}

type CertificateauthorityBaseKeyUsageStatus struct {
	/* The key may be used to sign certificates. */
	CertSign bool `json:"certSign,omitempty"`

	/* The key may be used for cryptographic commitments. Note that this may also be referred to as "non-repudiation". */
	ContentCommitment bool `json:"contentCommitment,omitempty"`

	/* The key may be used sign certificate revocation lists. */
	CrlSign bool `json:"crlSign,omitempty"`

	/* The key may be used to encipher data. */
	DataEncipherment bool `json:"dataEncipherment,omitempty"`

	/* The key may be used to decipher only. */
	DecipherOnly bool `json:"decipherOnly,omitempty"`

	/* The key may be used for digital signatures. */
	DigitalSignature bool `json:"digitalSignature,omitempty"`

	/* The key may be used to encipher only. */
	EncipherOnly bool `json:"encipherOnly,omitempty"`

	/* The key may be used in a key agreement protocol. */
	KeyAgreement bool `json:"keyAgreement,omitempty"`

	/* The key may be used to encipher other keys. */
	KeyEncipherment bool `json:"keyEncipherment,omitempty"`
}

type CertificateauthorityCaCertificateDescriptionsStatus struct {
	/* Describes lists of issuer CA certificate URLs that appear in the "Authority Information Access" extension in the certificate. */
	AiaIssuingCertificateUrls []string `json:"aiaIssuingCertificateUrls,omitempty"`

	/* Identifies the subject_key_id of the parent certificate, per https://tools.ietf.org/html/rfc5280#section-4.2.1.1 */
	AuthorityKeyId CertificateauthorityAuthorityKeyIdStatus `json:"authorityKeyId,omitempty"`

	/* The hash of the x.509 certificate. */
	CertFingerprint CertificateauthorityCertFingerprintStatus `json:"certFingerprint,omitempty"`

	/* Describes a list of locations to obtain CRL information, i.e. the DistributionPoint.fullName described by https://tools.ietf.org/html/rfc5280#section-4.2.1.13 */
	CrlDistributionPoints []string `json:"crlDistributionPoints,omitempty"`

	/* The public key that corresponds to an issued certificate. */
	PublicKey CertificateauthorityPublicKeyStatus `json:"publicKey,omitempty"`

	/* Describes some of the values in a certificate that are related to the subject and lifetime. */
	SubjectDescription CertificateauthoritySubjectDescriptionStatus `json:"subjectDescription,omitempty"`

	/* Provides a means of identifiying certificates that contain a particular public key, per https://tools.ietf.org/html/rfc5280#section-4.2.1.2. */
	SubjectKeyId CertificateauthoritySubjectKeyIdStatus `json:"subjectKeyId,omitempty"`

	/* Describes some of the technical X.509 fields in a certificate. */
	X509Description CertificateauthorityX509DescriptionStatus `json:"x509Description,omitempty"`
}

type CertificateauthorityCaOptionsStatus struct {
	/* Optional. Refers to the "CA" X.509 extension, which is a boolean value. When this value is missing, the extension will be omitted from the CA certificate. */
	IsCa bool `json:"isCa,omitempty"`

	/* Optional. Refers to the path length restriction X.509 extension. For a CA certificate, this value describes the depth of subordinate CA certificates that are allowed. If this value is less than 0, the request will fail. If this value is missing, the max path length will be omitted from the CA certificate. */
	MaxIssuerPathLength int `json:"maxIssuerPathLength,omitempty"`
}

type CertificateauthorityCertFingerprintStatus struct {
	/* The SHA 256 hash, encoded in hexadecimal, of the DER x509 certificate. */
	Sha256Hash string `json:"sha256Hash,omitempty"`
}

type CertificateauthorityConfigStatus struct {
	/* Optional. The public key that corresponds to this config. This is, for example, used when issuing Certificates, but not when creating a self-signed CertificateAuthority or CertificateAuthority CSR. */
	PublicKey CertificateauthorityPublicKeyStatus `json:"publicKey,omitempty"`

	/*  */
	X509Config CertificateauthorityX509ConfigStatus `json:"x509Config,omitempty"`
}

type CertificateauthorityCustomSansStatus struct {
	/* Optional. Indicates whether or not this extension is critical (i.e., if the client does not know how to handle this extension, the client should consider this to be an error). */
	Critical bool `json:"critical,omitempty"`

	/* Required. The OID for this X.509 extension. */
	ObjectId CertificateauthorityObjectIdStatus `json:"objectId,omitempty"`

	/* Required. The value of this X.509 extension. */
	Value string `json:"value,omitempty"`
}

type CertificateauthorityExtendedKeyUsageStatus struct {
	/* Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially described as "TLS WWW client authentication", though regularly used for non-WWW TLS. */
	ClientAuth bool `json:"clientAuth,omitempty"`

	/* Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially described as "Signing of downloadable executable code client authentication". */
	CodeSigning bool `json:"codeSigning,omitempty"`

	/* Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially described as "Email protection". */
	EmailProtection bool `json:"emailProtection,omitempty"`

	/* Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially described as "Signing OCSP responses". */
	OcspSigning bool `json:"ocspSigning,omitempty"`

	/* Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially described as "TLS WWW server authentication", though regularly used for non-WWW TLS. */
	ServerAuth bool `json:"serverAuth,omitempty"`

	/* Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially described as "Binding the hash of an object to a time". */
	TimeStamping bool `json:"timeStamping,omitempty"`
}

type CertificateauthorityKeyUsageStatus struct {
	/* Describes high-level ways in which a key may be used. */
	BaseKeyUsage CertificateauthorityBaseKeyUsageStatus `json:"baseKeyUsage,omitempty"`

	/* Detailed scenarios in which a key may be used. */
	ExtendedKeyUsage CertificateauthorityExtendedKeyUsageStatus `json:"extendedKeyUsage,omitempty"`

	/* Used to describe extended key usages that are not listed in the KeyUsage.ExtendedKeyUsageOptions message. */
	UnknownExtendedKeyUsages []CertificateauthorityUnknownExtendedKeyUsagesStatus `json:"unknownExtendedKeyUsages,omitempty"`
}

type CertificateauthorityObjectIdStatus struct {
	/* Required. The parts of an OID path. The most significant parts of the path come first. */
	ObjectIdPath []int `json:"objectIdPath,omitempty"`
}

type CertificateauthorityPemIssuerChainStatus struct {
	/* Required. Expected to be in leaf-to-root order according to RFC 5246. */
	PemCertificates []string `json:"pemCertificates,omitempty"`
}

type CertificateauthorityPolicyIdsStatus struct {
	/* Required. The parts of an OID path. The most significant parts of the path come first. */
	ObjectIdPath []int `json:"objectIdPath,omitempty"`
}

type CertificateauthorityPublicKeyStatus struct {
	/* Required. The format of the public key. Possible values: PEM */
	Format string `json:"format,omitempty"`

	/* Required. A public key. The padding and encoding must match with the `KeyFormat` value specified for the `format` field. */
	Key string `json:"key,omitempty"`
}

type CertificateauthoritySubjectAltNameStatus struct {
	/* Contains additional subject alternative name values. */
	CustomSans []CertificateauthorityCustomSansStatus `json:"customSans,omitempty"`

	/* Contains only valid, fully-qualified host names. */
	DnsNames []string `json:"dnsNames,omitempty"`

	/* Contains only valid RFC 2822 E-mail addresses. */
	EmailAddresses []string `json:"emailAddresses,omitempty"`

	/* Contains only valid 32-bit IPv4 addresses or RFC 4291 IPv6 addresses. */
	IpAddresses []string `json:"ipAddresses,omitempty"`

	/* Contains only valid RFC 3986 URIs. */
	Uris []string `json:"uris,omitempty"`
}

type CertificateauthoritySubjectDescriptionStatus struct {
	/* The serial number encoded in lowercase hexadecimal. */
	HexSerialNumber string `json:"hexSerialNumber,omitempty"`

	/* For convenience, the actual lifetime of an issued certificate. */
	Lifetime string `json:"lifetime,omitempty"`

	/* The time after which the certificate is expired. Per RFC 5280, the validity period for a certificate is the period of time from not_before_time through not_after_time, inclusive. Corresponds to 'not_before_time' + 'lifetime' - 1 second. */
	NotAfterTime string `json:"notAfterTime,omitempty"`

	/* The time at which the certificate becomes valid. */
	NotBeforeTime string `json:"notBeforeTime,omitempty"`

	/* Contains distinguished name fields such as the common name, location and organization. */
	Subject CertificateauthoritySubjectStatus `json:"subject,omitempty"`

	/* The subject alternative name fields. */
	SubjectAltName CertificateauthoritySubjectAltNameStatus `json:"subjectAltName,omitempty"`
}

type CertificateauthoritySubjectKeyIdStatus struct {
	/* Optional. The value of this KeyId encoded in lowercase hexadecimal. This is most likely the 160 bit SHA-1 hash of the public key. */
	KeyId string `json:"keyId,omitempty"`
}

type CertificateauthoritySubjectStatus struct {
	/* The "common name" of the subject. */
	CommonName string `json:"commonName,omitempty"`

	/* The country code of the subject. */
	CountryCode string `json:"countryCode,omitempty"`

	/* The locality or city of the subject. */
	Locality string `json:"locality,omitempty"`

	/* The organization of the subject. */
	Organization string `json:"organization,omitempty"`

	/* The organizational_unit of the subject. */
	OrganizationalUnit string `json:"organizationalUnit,omitempty"`

	/* The postal code of the subject. */
	PostalCode string `json:"postalCode,omitempty"`

	/* The province, territory, or regional state of the subject. */
	Province string `json:"province,omitempty"`

	/* The street address of the subject. */
	StreetAddress string `json:"streetAddress,omitempty"`
}

type CertificateauthoritySubordinateConfigStatus struct {
	/* Required. This can refer to a CertificateAuthority in the same project that was used to create a subordinate CertificateAuthority. This field is used for information and usability purposes only. The resource name is in the format `projects/* /locations/* /caPools/* /certificateAuthorities/*`. */
	CertificateAuthority string `json:"certificateAuthority,omitempty"`

	/* Required. Contains the PEM certificate chain for the issuers of this CertificateAuthority, but not pem certificate for this CA itself. */
	PemIssuerChain CertificateauthorityPemIssuerChainStatus `json:"pemIssuerChain,omitempty"`
}

type CertificateauthorityUnknownExtendedKeyUsagesStatus struct {
	/* Required. The parts of an OID path. The most significant parts of the path come first. */
	ObjectIdPath []int `json:"objectIdPath,omitempty"`
}

type CertificateauthorityX509ConfigStatus struct {
	/* Optional. Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the "Authority Information Access" extension in the certificate. */
	AiaOcspServers []string `json:"aiaOcspServers,omitempty"`
}

type CertificateauthorityX509DescriptionStatus struct {
	/* Optional. Describes custom X.509 extensions. */
	AdditionalExtensions []CertificateauthorityAdditionalExtensionsStatus `json:"additionalExtensions,omitempty"`

	/* Optional. Describes Online Certificate Status Protocol (OCSP) endpoint addresses that appear in the "Authority Information Access" extension in the certificate. */
	AiaOcspServers []string `json:"aiaOcspServers,omitempty"`

	/* Optional. Describes options in this X509Parameters that are relevant in a CA certificate. */
	CaOptions CertificateauthorityCaOptionsStatus `json:"caOptions,omitempty"`

	/* Optional. Indicates the intended use for keys that correspond to a certificate. */
	KeyUsage CertificateauthorityKeyUsageStatus `json:"keyUsage,omitempty"`

	/* Optional. Describes the X.509 certificate policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4. */
	PolicyIds []CertificateauthorityPolicyIdsStatus `json:"policyIds,omitempty"`
}

type PrivateCACertificateAuthorityStatus struct {
	/* Conditions represent the latest available observations of the
	   PrivateCACertificateAuthority's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Output only. URLs for accessing content published by this CA, such as the CA certificate and CRLs. */
	AccessUrls CertificateauthorityAccessUrlsStatus `json:"accessUrls,omitempty"`
	/* Output only. A structured description of this CertificateAuthority's CA certificate and its issuers. Ordered as self-to-root. */
	CaCertificateDescriptions []CertificateauthorityCaCertificateDescriptionsStatus `json:"caCertificateDescriptions,omitempty"`
	/*  */
	Config CertificateauthorityConfigStatus `json:"config,omitempty"`
	/* Output only. The time at which this CertificateAuthority was created. */
	CreateTime string `json:"createTime,omitempty"`
	/* Output only. The time at which this CertificateAuthority was soft deleted, if it is in the DELETED state. */
	DeleteTime string `json:"deleteTime,omitempty"`
	/* Output only. The time at which this CertificateAuthority will be permanently purged, if it is in the DELETED state. */
	ExpireTime string `json:"expireTime,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* Output only. This CertificateAuthority's certificate chain, including the current CertificateAuthority's certificate. Ordered such that the root issuer is the final element (consistent with RFC 5246). For a self-signed CA, this will only list the current CertificateAuthority's certificate. */
	PemCaCertificates []string `json:"pemCaCertificates,omitempty"`
	/* Output only. The State for this CertificateAuthority. Possible values: ENABLED, DISABLED, STAGED, AWAITING_USER_ACTIVATION, DELETED */
	State string `json:"state,omitempty"`
	/* Optional. If this is a subordinate CertificateAuthority, this field will be set with the subordinate configuration, which describes its issuers. This may be updated, but this CertificateAuthority must continue to validate. */
	SubordinateConfig CertificateauthoritySubordinateConfigStatus `json:"subordinateConfig,omitempty"`
	/* Output only. The CaPool.Tier of the CaPool that includes this CertificateAuthority. Possible values: ENTERPRISE, DEVOPS */
	Tier string `json:"tier,omitempty"`
	/* Output only. The time at which this CertificateAuthority was last updated. */
	UpdateTime string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PrivateCACertificateAuthority is the Schema for the privateca API
// +k8s:openapi-gen=true
type PrivateCACertificateAuthority struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PrivateCACertificateAuthoritySpec   `json:"spec,omitempty"`
	Status PrivateCACertificateAuthorityStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PrivateCACertificateAuthorityList contains a list of PrivateCACertificateAuthority
type PrivateCACertificateAuthorityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateCACertificateAuthority `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PrivateCACertificateAuthority{}, &PrivateCACertificateAuthorityList{})
}
