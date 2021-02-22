// +build !ignore_autogenerated

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
// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Access) DeepCopyInto(out *Access) {
	*out = *in
	out.View = in.View
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Access.
func (in *Access) DeepCopy() *Access {
	if in == nil {
		return nil
	}
	out := new(Access)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryDataset) DeepCopyInto(out *BigQueryDataset) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryDataset.
func (in *BigQueryDataset) DeepCopy() *BigQueryDataset {
	if in == nil {
		return nil
	}
	out := new(BigQueryDataset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryDataset) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryDatasetList) DeepCopyInto(out *BigQueryDatasetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BigQueryDataset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryDatasetList.
func (in *BigQueryDatasetList) DeepCopy() *BigQueryDatasetList {
	if in == nil {
		return nil
	}
	out := new(BigQueryDatasetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryDatasetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryDatasetSpec) DeepCopyInto(out *BigQueryDatasetSpec) {
	*out = *in
	if in.Access != nil {
		in, out := &in.Access, &out.Access
		*out = make([]Access, len(*in))
		copy(*out, *in)
	}
	out.DefaultEncryptionConfiguration = in.DefaultEncryptionConfiguration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryDatasetSpec.
func (in *BigQueryDatasetSpec) DeepCopy() *BigQueryDatasetSpec {
	if in == nil {
		return nil
	}
	out := new(BigQueryDatasetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryDatasetStatus) DeepCopyInto(out *BigQueryDatasetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryDatasetStatus.
func (in *BigQueryDatasetStatus) DeepCopy() *BigQueryDatasetStatus {
	if in == nil {
		return nil
	}
	out := new(BigQueryDatasetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryJob) DeepCopyInto(out *BigQueryJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryJob.
func (in *BigQueryJob) DeepCopy() *BigQueryJob {
	if in == nil {
		return nil
	}
	out := new(BigQueryJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryJobList) DeepCopyInto(out *BigQueryJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BigQueryJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryJobList.
func (in *BigQueryJobList) DeepCopy() *BigQueryJobList {
	if in == nil {
		return nil
	}
	out := new(BigQueryJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryJobSpec) DeepCopyInto(out *BigQueryJobSpec) {
	*out = *in
	in.Copy.DeepCopyInto(&out.Copy)
	in.Extract.DeepCopyInto(&out.Extract)
	in.Load.DeepCopyInto(&out.Load)
	in.Query.DeepCopyInto(&out.Query)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryJobSpec.
func (in *BigQueryJobSpec) DeepCopy() *BigQueryJobSpec {
	if in == nil {
		return nil
	}
	out := new(BigQueryJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryJobStatus) DeepCopyInto(out *BigQueryJobStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryJobStatus.
func (in *BigQueryJobStatus) DeepCopy() *BigQueryJobStatus {
	if in == nil {
		return nil
	}
	out := new(BigQueryJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryTable) DeepCopyInto(out *BigQueryTable) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryTable.
func (in *BigQueryTable) DeepCopy() *BigQueryTable {
	if in == nil {
		return nil
	}
	out := new(BigQueryTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryTable) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryTableList) DeepCopyInto(out *BigQueryTableList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BigQueryTable, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryTableList.
func (in *BigQueryTableList) DeepCopy() *BigQueryTableList {
	if in == nil {
		return nil
	}
	out := new(BigQueryTableList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BigQueryTableList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryTableSpec) DeepCopyInto(out *BigQueryTableSpec) {
	*out = *in
	if in.Clustering != nil {
		in, out := &in.Clustering, &out.Clustering
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.DatasetRef = in.DatasetRef
	out.EncryptionConfiguration = in.EncryptionConfiguration
	in.ExternalDataConfiguration.DeepCopyInto(&out.ExternalDataConfiguration)
	out.MaterializedView = in.MaterializedView
	out.RangePartitioning = in.RangePartitioning
	out.TimePartitioning = in.TimePartitioning
	out.View = in.View
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryTableSpec.
func (in *BigQueryTableSpec) DeepCopy() *BigQueryTableSpec {
	if in == nil {
		return nil
	}
	out := new(BigQueryTableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigQueryTableStatus) DeepCopyInto(out *BigQueryTableStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigQueryTableStatus.
func (in *BigQueryTableStatus) DeepCopy() *BigQueryTableStatus {
	if in == nil {
		return nil
	}
	out := new(BigQueryTableStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigquerydatasetView) DeepCopyInto(out *BigquerydatasetView) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigquerydatasetView.
func (in *BigquerydatasetView) DeepCopy() *BigquerydatasetView {
	if in == nil {
		return nil
	}
	out := new(BigquerydatasetView)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BigqueryjobTimePartitioning) DeepCopyInto(out *BigqueryjobTimePartitioning) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BigqueryjobTimePartitioning.
func (in *BigqueryjobTimePartitioning) DeepCopy() *BigqueryjobTimePartitioning {
	if in == nil {
		return nil
	}
	out := new(BigqueryjobTimePartitioning)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Copy) DeepCopyInto(out *Copy) {
	*out = *in
	out.DestinationEncryptionConfiguration = in.DestinationEncryptionConfiguration
	out.DestinationTable = in.DestinationTable
	if in.SourceTables != nil {
		in, out := &in.SourceTables, &out.SourceTables
		*out = make([]SourceTables, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Copy.
func (in *Copy) DeepCopy() *Copy {
	if in == nil {
		return nil
	}
	out := new(Copy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CsvOptions) DeepCopyInto(out *CsvOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CsvOptions.
func (in *CsvOptions) DeepCopy() *CsvOptions {
	if in == nil {
		return nil
	}
	out := new(CsvOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultDataset) DeepCopyInto(out *DefaultDataset) {
	*out = *in
	out.DatasetRef = in.DatasetRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultDataset.
func (in *DefaultDataset) DeepCopy() *DefaultDataset {
	if in == nil {
		return nil
	}
	out := new(DefaultDataset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultEncryptionConfiguration) DeepCopyInto(out *DefaultEncryptionConfiguration) {
	*out = *in
	out.KmsKeyRef = in.KmsKeyRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultEncryptionConfiguration.
func (in *DefaultEncryptionConfiguration) DeepCopy() *DefaultEncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := new(DefaultEncryptionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationEncryptionConfiguration) DeepCopyInto(out *DestinationEncryptionConfiguration) {
	*out = *in
	out.KmsKeyRef = in.KmsKeyRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationEncryptionConfiguration.
func (in *DestinationEncryptionConfiguration) DeepCopy() *DestinationEncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := new(DestinationEncryptionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationTable) DeepCopyInto(out *DestinationTable) {
	*out = *in
	out.TableRef = in.TableRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationTable.
func (in *DestinationTable) DeepCopy() *DestinationTable {
	if in == nil {
		return nil
	}
	out := new(DestinationTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionConfiguration) DeepCopyInto(out *EncryptionConfiguration) {
	*out = *in
	out.KmsKeyRef = in.KmsKeyRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionConfiguration.
func (in *EncryptionConfiguration) DeepCopy() *EncryptionConfiguration {
	if in == nil {
		return nil
	}
	out := new(EncryptionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalDataConfiguration) DeepCopyInto(out *ExternalDataConfiguration) {
	*out = *in
	out.CsvOptions = in.CsvOptions
	out.GoogleSheetsOptions = in.GoogleSheetsOptions
	out.HivePartitioningOptions = in.HivePartitioningOptions
	if in.SourceUris != nil {
		in, out := &in.SourceUris, &out.SourceUris
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalDataConfiguration.
func (in *ExternalDataConfiguration) DeepCopy() *ExternalDataConfiguration {
	if in == nil {
		return nil
	}
	out := new(ExternalDataConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Extract) DeepCopyInto(out *Extract) {
	*out = *in
	if in.DestinationUris != nil {
		in, out := &in.DestinationUris, &out.DestinationUris
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.SourceTable = in.SourceTable
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Extract.
func (in *Extract) DeepCopy() *Extract {
	if in == nil {
		return nil
	}
	out := new(Extract)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoogleSheetsOptions) DeepCopyInto(out *GoogleSheetsOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoogleSheetsOptions.
func (in *GoogleSheetsOptions) DeepCopy() *GoogleSheetsOptions {
	if in == nil {
		return nil
	}
	out := new(GoogleSheetsOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HivePartitioningOptions) DeepCopyInto(out *HivePartitioningOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HivePartitioningOptions.
func (in *HivePartitioningOptions) DeepCopy() *HivePartitioningOptions {
	if in == nil {
		return nil
	}
	out := new(HivePartitioningOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Load) DeepCopyInto(out *Load) {
	*out = *in
	out.DestinationEncryptionConfiguration = in.DestinationEncryptionConfiguration
	out.DestinationTable = in.DestinationTable
	if in.ProjectionFields != nil {
		in, out := &in.ProjectionFields, &out.ProjectionFields
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SchemaUpdateOptions != nil {
		in, out := &in.SchemaUpdateOptions, &out.SchemaUpdateOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SourceUris != nil {
		in, out := &in.SourceUris, &out.SourceUris
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.TimePartitioning = in.TimePartitioning
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Load.
func (in *Load) DeepCopy() *Load {
	if in == nil {
		return nil
	}
	out := new(Load)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaterializedView) DeepCopyInto(out *MaterializedView) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaterializedView.
func (in *MaterializedView) DeepCopy() *MaterializedView {
	if in == nil {
		return nil
	}
	out := new(MaterializedView)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Query) DeepCopyInto(out *Query) {
	*out = *in
	out.DefaultDataset = in.DefaultDataset
	out.DestinationEncryptionConfiguration = in.DestinationEncryptionConfiguration
	out.DestinationTable = in.DestinationTable
	if in.SchemaUpdateOptions != nil {
		in, out := &in.SchemaUpdateOptions, &out.SchemaUpdateOptions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.ScriptOptions = in.ScriptOptions
	if in.UserDefinedFunctionResources != nil {
		in, out := &in.UserDefinedFunctionResources, &out.UserDefinedFunctionResources
		*out = make([]UserDefinedFunctionResources, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Query.
func (in *Query) DeepCopy() *Query {
	if in == nil {
		return nil
	}
	out := new(Query)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Range) DeepCopyInto(out *Range) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Range.
func (in *Range) DeepCopy() *Range {
	if in == nil {
		return nil
	}
	out := new(Range)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RangePartitioning) DeepCopyInto(out *RangePartitioning) {
	*out = *in
	out.Range = in.Range
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RangePartitioning.
func (in *RangePartitioning) DeepCopy() *RangePartitioning {
	if in == nil {
		return nil
	}
	out := new(RangePartitioning)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScriptOptions) DeepCopyInto(out *ScriptOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScriptOptions.
func (in *ScriptOptions) DeepCopy() *ScriptOptions {
	if in == nil {
		return nil
	}
	out := new(ScriptOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceTable) DeepCopyInto(out *SourceTable) {
	*out = *in
	out.TableRef = in.TableRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceTable.
func (in *SourceTable) DeepCopy() *SourceTable {
	if in == nil {
		return nil
	}
	out := new(SourceTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceTables) DeepCopyInto(out *SourceTables) {
	*out = *in
	out.TableRef = in.TableRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceTables.
func (in *SourceTables) DeepCopy() *SourceTables {
	if in == nil {
		return nil
	}
	out := new(SourceTables)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimePartitioning) DeepCopyInto(out *TimePartitioning) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimePartitioning.
func (in *TimePartitioning) DeepCopy() *TimePartitioning {
	if in == nil {
		return nil
	}
	out := new(TimePartitioning)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserDefinedFunctionResources) DeepCopyInto(out *UserDefinedFunctionResources) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserDefinedFunctionResources.
func (in *UserDefinedFunctionResources) DeepCopy() *UserDefinedFunctionResources {
	if in == nil {
		return nil
	}
	out := new(UserDefinedFunctionResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *View) DeepCopyInto(out *View) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new View.
func (in *View) DeepCopy() *View {
	if in == nil {
		return nil
	}
	out := new(View)
	in.DeepCopyInto(out)
	return out
}
