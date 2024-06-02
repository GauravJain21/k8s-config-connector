// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The definition of a Expander service.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: proto/expander.proto

package expander

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status int32

const (
	Status_SUCCESS          Status = 0
	Status_VALIDATE_FAILED  Status = 1
	Status_EVALUATE_FAILED  Status = 2
	Status_EVALUATE_WAIT    Status = 3
	Status_UNEXPECTED_ERROR Status = 4
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "SUCCESS",
		1: "VALIDATE_FAILED",
		2: "EVALUATE_FAILED",
		3: "EVALUATE_WAIT",
		4: "UNEXPECTED_ERROR",
	}
	Status_value = map[string]int32{
		"SUCCESS":          0,
		"VALIDATE_FAILED":  1,
		"EVALUATE_FAILED":  2,
		"EVALUATE_WAIT":    3,
		"UNEXPECTED_ERROR": 4,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_expander_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_proto_expander_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_expander_proto_rawDescGZIP(), []int{0}
}

type ResultType int32

const (
	ResultType_UNKNOWN   ResultType = 0
	ResultType_MANIFESTS ResultType = 1
	ResultType_VALUES    ResultType = 2
)

// Enum value maps for ResultType.
var (
	ResultType_name = map[int32]string{
		0: "UNKNOWN",
		1: "MANIFESTS",
		2: "VALUES",
	}
	ResultType_value = map[string]int32{
		"UNKNOWN":   0,
		"MANIFESTS": 1,
		"VALUES":    2,
	}
)

func (x ResultType) Enum() *ResultType {
	p := new(ResultType)
	*p = x
	return p
}

func (x ResultType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResultType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_expander_proto_enumTypes[1].Descriptor()
}

func (ResultType) Type() protoreflect.EnumType {
	return &file_proto_expander_proto_enumTypes[1]
}

func (x ResultType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResultType.Descriptor instead.
func (ResultType) EnumDescriptor() ([]byte, []int) {
	return file_proto_expander_proto_rawDescGZIP(), []int{1}
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_expander_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_proto_expander_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_proto_expander_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ValidateResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=expander_grpc.Status" json:"status,omitempty"`
	Error  *Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ValidateResult) Reset() {
	*x = ValidateResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_expander_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateResult) ProtoMessage() {}

func (x *ValidateResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_expander_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateResult.ProtoReflect.Descriptor instead.
func (*ValidateResult) Descriptor() ([]byte, []int) {
	return file_proto_expander_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateResult) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_SUCCESS
}

func (x *ValidateResult) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type EvaluateResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    Status     `protobuf:"varint,1,opt,name=status,proto3,enum=expander_grpc.Status" json:"status,omitempty"`
	Error     *Error     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Type      ResultType `protobuf:"varint,3,opt,name=type,proto3,enum=expander_grpc.ResultType" json:"type,omitempty"`
	Manifests []byte     `protobuf:"bytes,4,opt,name=manifests,proto3" json:"manifests,omitempty"`
	Values    []byte     `protobuf:"bytes,5,opt,name=values,proto3" json:"values,omitempty"`
}

func (x *EvaluateResult) Reset() {
	*x = EvaluateResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_expander_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluateResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluateResult) ProtoMessage() {}

func (x *EvaluateResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_expander_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluateResult.ProtoReflect.Descriptor instead.
func (*EvaluateResult) Descriptor() ([]byte, []int) {
	return file_proto_expander_proto_rawDescGZIP(), []int{2}
}

func (x *EvaluateResult) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_SUCCESS
}

func (x *EvaluateResult) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *EvaluateResult) GetType() ResultType {
	if x != nil {
		return x.Type
	}
	return ResultType_UNKNOWN
}

func (x *EvaluateResult) GetManifests() []byte {
	if x != nil {
		return x.Manifests
	}
	return nil
}

func (x *EvaluateResult) GetValues() []byte {
	if x != nil {
		return x.Values
	}
	return nil
}

type EvaluateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config   []byte `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	Context  []byte `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
	Facade   []byte `protobuf:"bytes,3,opt,name=facade,proto3" json:"facade,omitempty"`
	Value    []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Resource string `protobuf:"bytes,5,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *EvaluateRequest) Reset() {
	*x = EvaluateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_expander_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluateRequest) ProtoMessage() {}

func (x *EvaluateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_expander_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluateRequest.ProtoReflect.Descriptor instead.
func (*EvaluateRequest) Descriptor() ([]byte, []int) {
	return file_proto_expander_proto_rawDescGZIP(), []int{3}
}

func (x *EvaluateRequest) GetConfig() []byte {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *EvaluateRequest) GetContext() []byte {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *EvaluateRequest) GetFacade() []byte {
	if x != nil {
		return x.Facade
	}
	return nil
}

func (x *EvaluateRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *EvaluateRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

type ValidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config   []byte `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	Context  []byte `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
	Facade   []byte `protobuf:"bytes,3,opt,name=facade,proto3" json:"facade,omitempty"`
	Value    []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Resource string `protobuf:"bytes,5,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *ValidateRequest) Reset() {
	*x = ValidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_expander_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRequest) ProtoMessage() {}

func (x *ValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_expander_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRequest.ProtoReflect.Descriptor instead.
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return file_proto_expander_proto_rawDescGZIP(), []int{4}
}

func (x *ValidateRequest) GetConfig() []byte {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *ValidateRequest) GetContext() []byte {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *ValidateRequest) GetFacade() []byte {
	if x != nil {
		return x.Facade
	}
	return nil
}

func (x *ValidateRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ValidateRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

var File_proto_expander_proto protoreflect.FileDescriptor

var file_proto_expander_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x72,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x22, 0x21, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x6b, 0x0a, 0x0e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x65, 0x78, 0x70,
	0x61, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x78, 0x70, 0x61, 0x6e,
	0x64, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xd0, 0x01, 0x0a, 0x0e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x65, 0x78, 0x70, 0x61, 0x6e,
	0x64, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65,
	0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x19, 0x2e, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x8d, 0x01, 0x0a, 0x0f, 0x45, 0x76, 0x61,
	0x6c, 0x75, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x61, 0x63, 0x61, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x66, 0x61, 0x63, 0x61, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x0f, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x61, 0x63, 0x61, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x66, 0x61, 0x63, 0x61, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2a, 0x68, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12,
	0x13, 0x0a, 0x0f, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x45, 0x56, 0x41, 0x4c, 0x55, 0x41, 0x54, 0x45,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x56, 0x41,
	0x4c, 0x55, 0x41, 0x54, 0x45, 0x5f, 0x57, 0x41, 0x49, 0x54, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10,
	0x55, 0x4e, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x04, 0x2a, 0x34, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0d, 0x0a,
	0x09, 0x4d, 0x41, 0x4e, 0x49, 0x46, 0x45, 0x53, 0x54, 0x53, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x56, 0x41, 0x4c, 0x55, 0x45, 0x53, 0x10, 0x02, 0x32, 0xa4, 0x01, 0x0a, 0x08, 0x45, 0x78, 0x70,
	0x61, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x1e, 0x2e, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x00, 0x12, 0x4b, 0x0a, 0x08, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x12, 0x1e,
	0x2e, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42,
	0x27, 0x5a, 0x25, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_expander_proto_rawDescOnce sync.Once
	file_proto_expander_proto_rawDescData = file_proto_expander_proto_rawDesc
)

func file_proto_expander_proto_rawDescGZIP() []byte {
	file_proto_expander_proto_rawDescOnce.Do(func() {
		file_proto_expander_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_expander_proto_rawDescData)
	})
	return file_proto_expander_proto_rawDescData
}

var file_proto_expander_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_expander_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_expander_proto_goTypes = []interface{}{
	(Status)(0),             // 0: expander_grpc.Status
	(ResultType)(0),         // 1: expander_grpc.ResultType
	(*Error)(nil),           // 2: expander_grpc.Error
	(*ValidateResult)(nil),  // 3: expander_grpc.ValidateResult
	(*EvaluateResult)(nil),  // 4: expander_grpc.EvaluateResult
	(*EvaluateRequest)(nil), // 5: expander_grpc.EvaluateRequest
	(*ValidateRequest)(nil), // 6: expander_grpc.ValidateRequest
}
var file_proto_expander_proto_depIdxs = []int32{
	0, // 0: expander_grpc.ValidateResult.status:type_name -> expander_grpc.Status
	2, // 1: expander_grpc.ValidateResult.error:type_name -> expander_grpc.Error
	0, // 2: expander_grpc.EvaluateResult.status:type_name -> expander_grpc.Status
	2, // 3: expander_grpc.EvaluateResult.error:type_name -> expander_grpc.Error
	1, // 4: expander_grpc.EvaluateResult.type:type_name -> expander_grpc.ResultType
	6, // 5: expander_grpc.Expander.Validate:input_type -> expander_grpc.ValidateRequest
	5, // 6: expander_grpc.Expander.Evaluate:input_type -> expander_grpc.EvaluateRequest
	3, // 7: expander_grpc.Expander.Validate:output_type -> expander_grpc.ValidateResult
	4, // 8: expander_grpc.Expander.Evaluate:output_type -> expander_grpc.EvaluateResult
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_expander_proto_init() }
func file_proto_expander_proto_init() {
	if File_proto_expander_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_expander_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_expander_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_expander_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvaluateResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_expander_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvaluateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_expander_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_expander_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_expander_proto_goTypes,
		DependencyIndexes: file_proto_expander_proto_depIdxs,
		EnumInfos:         file_proto_expander_proto_enumTypes,
		MessageInfos:      file_proto_expander_proto_msgTypes,
	}.Build()
	File_proto_expander_proto = out.File
	file_proto_expander_proto_rawDesc = nil
	file_proto_expander_proto_goTypes = nil
	file_proto_expander_proto_depIdxs = nil
}
