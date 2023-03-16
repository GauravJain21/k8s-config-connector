// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package scheme

import (
	accesscontextmanagerv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/accesscontextmanager/v1alpha1"
	accesscontextmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/accesscontextmanager/v1beta1"
	apigeev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/apigee/v1beta1"
	artifactregistryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/artifactregistry/v1beta1"
	bigqueryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/bigquery/v1beta1"
	bigtablev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/bigtable/v1beta1"
	billingbudgetsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/billingbudgets/v1beta1"
	binaryauthorizationv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/binaryauthorization/v1beta1"
	cloudbuildv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/cloudbuild/v1beta1"
	cloudfunctionsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/cloudfunctions/v1beta1"
	cloudidentityv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/cloudidentity/v1beta1"
	cloudschedulerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/cloudscheduler/v1beta1"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1beta1"
	configcontrollerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/configcontroller/v1beta1"
	containerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/container/v1beta1"
	containeranalysisv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/containeranalysis/v1beta1"
	datacatalogv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/datacatalog/v1beta1"
	dataflowv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/dataflow/v1beta1"
	datafusionv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/datafusion/v1beta1"
	dataprocv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/dataproc/v1beta1"
	dlpv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/dlp/v1beta1"
	dnsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/dns/v1beta1"
	eventarcv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/eventarc/v1beta1"
	filestorev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/filestore/v1beta1"
	firestorev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/firestore/v1beta1"
	gkehubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/gkehub/v1beta1"
	iamv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/iam/v1beta1"
	iapv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/iap/v1beta1"
	identityplatformv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/identityplatform/v1beta1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/kms/v1beta1"
	loggingv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/logging/v1beta1"
	memcachev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/memcache/v1beta1"
	monitoringv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/monitoring/v1beta1"
	networkconnectivityv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/networkconnectivity/v1beta1"
	networksecurityv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/networksecurity/v1beta1"
	networkservicesv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/networkservices/v1beta1"
	osconfigv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/osconfig/v1beta1"
	privatecav1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/privateca/v1beta1"
	pubsubv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/pubsub/v1beta1"
	pubsublitev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/pubsublite/v1beta1"
	recaptchaenterprisev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/recaptchaenterprise/v1beta1"
	redisv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/redis/v1beta1"
	resourcemanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/resourcemanager/v1beta1"
	runv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/run/v1beta1"
	secretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/secretmanager/v1beta1"
	servicedirectoryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/servicedirectory/v1beta1"
	servicenetworkingv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/servicenetworking/v1beta1"
	serviceusagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/serviceusage/v1beta1"
	sourcerepov1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/sourcerepo/v1beta1"
	spannerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/spanner/v1beta1"
	sqlv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/sql/v1beta1"
	storagev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/storage/v1beta1"
	storagetransferv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/storagetransfer/v1beta1"
	tagsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/tags/v1beta1"
	vpcaccessv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/vpcaccess/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)
var ParameterCodec = runtime.NewParameterCodec(Scheme)
var localSchemeBuilder = runtime.SchemeBuilder{
	accesscontextmanagerv1beta1.AddToScheme,
	accesscontextmanagerv1alpha1.AddToScheme,
	apigeev1beta1.AddToScheme,
	artifactregistryv1beta1.AddToScheme,
	bigqueryv1beta1.AddToScheme,
	bigtablev1beta1.AddToScheme,
	billingbudgetsv1beta1.AddToScheme,
	binaryauthorizationv1beta1.AddToScheme,
	cloudbuildv1beta1.AddToScheme,
	cloudfunctionsv1beta1.AddToScheme,
	cloudidentityv1beta1.AddToScheme,
	cloudschedulerv1beta1.AddToScheme,
	computev1beta1.AddToScheme,
	configcontrollerv1beta1.AddToScheme,
	containerv1beta1.AddToScheme,
	containeranalysisv1beta1.AddToScheme,
	datacatalogv1beta1.AddToScheme,
	dataflowv1beta1.AddToScheme,
	datafusionv1beta1.AddToScheme,
	dataprocv1beta1.AddToScheme,
	dlpv1beta1.AddToScheme,
	dnsv1beta1.AddToScheme,
	eventarcv1beta1.AddToScheme,
	filestorev1beta1.AddToScheme,
	firestorev1beta1.AddToScheme,
	gkehubv1beta1.AddToScheme,
	iamv1beta1.AddToScheme,
	iapv1beta1.AddToScheme,
	identityplatformv1beta1.AddToScheme,
	k8sv1alpha1.AddToScheme,
	kmsv1beta1.AddToScheme,
	loggingv1beta1.AddToScheme,
	memcachev1beta1.AddToScheme,
	monitoringv1beta1.AddToScheme,
	networkconnectivityv1beta1.AddToScheme,
	networksecurityv1beta1.AddToScheme,
	networkservicesv1beta1.AddToScheme,
	osconfigv1beta1.AddToScheme,
	privatecav1beta1.AddToScheme,
	pubsubv1beta1.AddToScheme,
	pubsublitev1beta1.AddToScheme,
	recaptchaenterprisev1beta1.AddToScheme,
	redisv1beta1.AddToScheme,
	resourcemanagerv1beta1.AddToScheme,
	runv1beta1.AddToScheme,
	secretmanagerv1beta1.AddToScheme,
	servicedirectoryv1beta1.AddToScheme,
	servicenetworkingv1beta1.AddToScheme,
	serviceusagev1beta1.AddToScheme,
	sourcerepov1beta1.AddToScheme,
	spannerv1beta1.AddToScheme,
	sqlv1beta1.AddToScheme,
	storagev1beta1.AddToScheme,
	storagetransferv1beta1.AddToScheme,
	tagsv1beta1.AddToScheme,
	vpcaccessv1beta1.AddToScheme,
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//	import (
//	  "k8s.io/client-go/kubernetes"
//	  clientsetscheme "k8s.io/client-go/kubernetes/scheme"
//	  aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//	)
//
//	kclientset, _ := kubernetes.NewForConfig(c)
//	_ = aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	v1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(Scheme))
}
