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

type DiskEncryptionKey struct {
	/* The encryption key used to encrypt the disk. Your project's Compute
	Engine System service account
	('service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com')
	must have 'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this
	feature. See
	https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys */
	KmsKeyRef v1alpha1.ResourceRef `json:"kmsKeyRef,omitempty"`
	/* The service account used for the encryption request for the given KMS key.
	If absent, the Compute Engine Service Agent service account is used. */
	KmsKeyServiceAccountRef v1alpha1.ResourceRef `json:"kmsKeyServiceAccountRef,omitempty"`
	/* Immutable. Specifies a 256-bit customer-supplied encryption key, encoded in
	RFC 4648 base64 to either encrypt or decrypt this resource. */
	RawKey RawKey `json:"rawKey,omitempty"`
	/* The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
	encryption key that protects this resource. */
	Sha256 string `json:"sha256,omitempty"`
}

type RawKey struct {
	/* Value of the field. Cannot be used if 'valueFrom' is specified. */
	Value string `json:"value,omitempty"`
	/* Source for the field's value. Cannot be used if 'value' is specified. */
	ValueFrom ValueFrom `json:"valueFrom,omitempty"`
}

type SourceImageEncryptionKey struct {
	/* The encryption key used to encrypt the disk. Your project's Compute
	Engine System service account
	('service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com')
	must have 'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this
	feature. See
	https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys */
	KmsKeyRef v1alpha1.ResourceRef `json:"kmsKeyRef,omitempty"`
	/* The service account used for the encryption request for the given KMS key.
	If absent, the Compute Engine Service Agent service account is used. */
	KmsKeyServiceAccountRef v1alpha1.ResourceRef `json:"kmsKeyServiceAccountRef,omitempty"`
	/* Immutable. Specifies a 256-bit customer-supplied encryption key, encoded in
	RFC 4648 base64 to either encrypt or decrypt this resource. */
	RawKey string `json:"rawKey,omitempty"`
	/* The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
	encryption key that protects this resource. */
	Sha256 string `json:"sha256,omitempty"`
}

type SourceSnapshotEncryptionKey struct {
	/* The encryption key used to encrypt the disk. Your project's Compute
	Engine System service account
	('service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com')
	must have 'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this
	feature. See
	https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys */
	KmsKeyRef v1alpha1.ResourceRef `json:"kmsKeyRef,omitempty"`
	/* The service account used for the encryption request for the given KMS key.
	If absent, the Compute Engine Service Agent service account is used. */
	KmsKeyServiceAccountRef v1alpha1.ResourceRef `json:"kmsKeyServiceAccountRef,omitempty"`
	/* Immutable. Specifies a 256-bit customer-supplied encryption key, encoded in
	RFC 4648 base64 to either encrypt or decrypt this resource. */
	RawKey string `json:"rawKey,omitempty"`
	/* The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
	encryption key that protects this resource. */
	Sha256 string `json:"sha256,omitempty"`
}

type ValueFrom struct {
	/* Reference to a value with the given key in the given Secret in the resource's namespace. */
	SecretKeyRef v1alpha1.ResourceRef `json:"secretKeyRef,omitempty"`
}

type ComputeDiskSpec struct {
	/* Immutable. An optional description of this resource. Provide this property when
	you create the resource. */
	Description string `json:"description,omitempty"`
	/* Immutable. Encrypts the disk using a customer-supplied encryption key.

	After you encrypt a disk with a customer-supplied key, you must
	provide the same key if you use the disk later (e.g. to create a disk
	snapshot or an image, or to attach the disk to a virtual machine).

	Customer-supplied encryption keys do not protect access to metadata of
	the disk.

	If you do not provide an encryption key when creating the disk, then
	the disk will be encrypted using an automatically generated key and
	you do not need to provide a key to use the disk later. */
	DiskEncryptionKey DiskEncryptionKey `json:"diskEncryptionKey,omitempty"`
	/* The image from which to initialize this disk. */
	ImageRef v1alpha1.ResourceRef `json:"imageRef,omitempty"`
	/* Immutable. Specifies the disk interface to use for attaching this disk, which is either SCSI or NVME. The default is SCSI. Default value: "SCSI" Possible values: ["SCSI", "NVME"] */
	Interface string `json:"interface,omitempty"`
	/* Location represents the geographical location of the ComputeDisk. Specify a region name or a zone name. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/) */
	Location string `json:"location,omitempty"`
	/* Immutable. Indicates whether or not the disk can be read/write attached to more than one instance. */
	MultiWriter bool `json:"multiWriter,omitempty"`
	/* Immutable. Physical block size of the persistent disk, in bytes. If not present
	in a request, a default value is used. Currently supported sizes
	are 4096 and 16384, other sizes may be added in the future.
	If an unsupported value is requested, the error message will list
	the supported values for the caller's project. */
	PhysicalBlockSizeBytes int `json:"physicalBlockSizeBytes,omitempty"`
	/* Immutable. URLs of the zones where the disk should be replicated to. */
	ReplicaZones []string `json:"replicaZones,omitempty"`
	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	ResourceID string `json:"resourceID,omitempty"`
	/*  */
	ResourcePolicies []v1alpha1.ResourceRef `json:"resourcePolicies,omitempty"`
	/* Size of the persistent disk, specified in GB. You can specify this
	field when creating a persistent disk using the 'image' or
	'snapshot' parameter, or specify it alone to create an empty
	persistent disk.

	If you specify this field along with 'image' or 'snapshot',
	the value must not be less than the size of the image
	or the size of the snapshot. */
	Size int `json:"size,omitempty"`
	/* The source snapshot used to create this disk. */
	SnapshotRef v1alpha1.ResourceRef `json:"snapshotRef,omitempty"`
	/* Immutable. The customer-supplied encryption key of the source image. Required if
	the source image is protected by a customer-supplied encryption key. */
	SourceImageEncryptionKey SourceImageEncryptionKey `json:"sourceImageEncryptionKey,omitempty"`
	/* Immutable. The customer-supplied encryption key of the source snapshot. Required
	if the source snapshot is protected by a customer-supplied encryption
	key. */
	SourceSnapshotEncryptionKey SourceSnapshotEncryptionKey `json:"sourceSnapshotEncryptionKey,omitempty"`
	/* Immutable. URL of the disk type resource describing which disk type to use to
	create the disk. Provide this when creating the disk. */
	Type string `json:"type,omitempty"`
}

type ComputeDiskStatus struct {
	/* Conditions represents the latest available observations of the
	   ComputeDisk's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Creation timestamp in RFC3339 text format. */
	CreationTimestamp string `json:"creationTimestamp,omitempty"`
	/* The fingerprint used for optimistic locking of this resource.  Used
	internally during updates. */
	LabelFingerprint string `json:"labelFingerprint,omitempty"`
	/* Last attach timestamp in RFC3339 text format. */
	LastAttachTimestamp string `json:"lastAttachTimestamp,omitempty"`
	/* Last detach timestamp in RFC3339 text format. */
	LastDetachTimestamp string `json:"lastDetachTimestamp,omitempty"`
	/*  */
	SelfLink string `json:"selfLink,omitempty"`
	/* The ID value of the image used to create this disk. This value
	identifies the exact image that was used to create this persistent
	disk. For example, if you created the persistent disk from an image
	that was later deleted and recreated under the same name, the source
	image ID would identify the exact version of the image that was used. */
	SourceImageId string `json:"sourceImageId,omitempty"`
	/* The unique ID of the snapshot used to create this disk. This value
	identifies the exact snapshot that was used to create this persistent
	disk. For example, if you created the persistent disk from a snapshot
	that was later deleted and recreated under the same name, the source
	snapshot ID would identify the exact version of the snapshot that was
	used. */
	SourceSnapshotId string `json:"sourceSnapshotId,omitempty"`
	/* Links to the users of the disk (attached instances) in form:
	project/zones/zone/instances/instance */
	Users []string `json:"users,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeDisk is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeDisk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeDiskSpec   `json:"spec,omitempty"`
	Status ComputeDiskStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeDiskList contains a list of ComputeDisk
type ComputeDiskList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []ComputeDisk `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeDisk{}, &ComputeDiskList{})
}
