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

type Fields struct {
	/* Immutable. Indicates that this field supports operations on arrayValues. Only one of 'order' and 'arrayConfig' can
	be specified. Possible values: ["CONTAINS"] */
	ArrayConfig string `json:"arrayConfig,omitempty"`
	/* Immutable. Name of the field. */
	FieldPath string `json:"fieldPath,omitempty"`
	/* Immutable. Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=.
	Only one of 'order' and 'arrayConfig' can be specified. Possible values: ["ASCENDING", "DESCENDING"] */
	Order string `json:"order,omitempty"`
}

type FirestoreIndexSpec struct {
	/* Immutable. The collection being indexed. */
	Collection string `json:"collection,omitempty"`
	/* Immutable. The Firestore database id. Defaults to '"(default)"'. */
	Database string `json:"database,omitempty"`
	/* Immutable. The fields supported by this index. The last field entry is always for
	the field path '__name__'. If, on creation, '__name__' was not
	specified as the last field, it will be added automatically with the
	same direction as that of the last field defined. If the final field
	in a composite index is not directional, the '__name__' will be
	ordered '"ASCENDING"' (unless explicitly specified otherwise). */
	Fields []Fields `json:"fields,omitempty"`
	/* Immutable. The scope at which a query is run. Default value: "COLLECTION" Possible values: ["COLLECTION", "COLLECTION_GROUP"] */
	QueryScope string `json:"queryScope,omitempty"`
}

type FirestoreIndexStatus struct {
	/* Conditions represents the latest available observations of the
	   FirestoreIndex's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* A server defined name for this index. Format:
	'projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}' */
	Name string `json:"name,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FirestoreIndex is the Schema for the firestore API
// +k8s:openapi-gen=true
type FirestoreIndex struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirestoreIndexSpec   `json:"spec,omitempty"`
	Status FirestoreIndexStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FirestoreIndexList contains a list of FirestoreIndex
type FirestoreIndexList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []FirestoreIndex `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FirestoreIndex{}, &FirestoreIndexList{})
}
