//
//Copyright 2021 The Dapr Authors
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//http://www.apache.org/licenses/LICENSE-2.0
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: dapr/proto/operator/v1/operator.proto

package operator

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ListComponentsRequest is the request to get components for a sidecar in namespace.
type ListComponentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	PodName   string `protobuf:"bytes,2,opt,name=podName,proto3" json:"podName,omitempty"`
}

func (x *ListComponentsRequest) Reset() {
	*x = ListComponentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dapr_proto_operator_v1_operator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListComponentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListComponentsRequest) ProtoMessage() {}

func (x *ListComponentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dapr_proto_operator_v1_operator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListComponentsRequest.ProtoReflect.Descriptor instead.
func (*ListComponentsRequest) Descriptor() ([]byte, []int) {
	return file_dapr_proto_operator_v1_operator_proto_rawDescGZIP(), []int{0}
}

func (x *ListComponentsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ListComponentsRequest) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

// ComponentUpdateRequest is the request to get updates about new components for a given namespace.
type ComponentUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	PodName   string `protobuf:"bytes,2,opt,name=podName,proto3" json:"podName,omitempty"`
}

func (x *ComponentUpdateRequest) Reset() {
	*x = ComponentUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dapr_proto_operator_v1_operator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentUpdateRequest) ProtoMessage() {}

func (x *ComponentUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dapr_proto_operator_v1_operator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentUpdateRequest.ProtoReflect.Descriptor instead.
func (*ComponentUpdateRequest) Descriptor() ([]byte, []int) {
	return file_dapr_proto_operator_v1_operator_proto_rawDescGZIP(), []int{1}
}

func (x *ComponentUpdateRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ComponentUpdateRequest) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

// ComponentUpdateEvent includes the updated component event.
type ComponentUpdateEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Component []byte `protobuf:"bytes,1,opt,name=component,proto3" json:"component,omitempty"`
}

func (x *ComponentUpdateEvent) Reset() {
	*x = ComponentUpdateEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dapr_proto_operator_v1_operator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComponentUpdateEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComponentUpdateEvent) ProtoMessage() {}

func (x *ComponentUpdateEvent) ProtoReflect() protoreflect.Message {
	mi := &file_dapr_proto_operator_v1_operator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComponentUpdateEvent.ProtoReflect.Descriptor instead.
func (*ComponentUpdateEvent) Descriptor() ([]byte, []int) {
	return file_dapr_proto_operator_v1_operator_proto_rawDescGZIP(), []int{2}
}

func (x *ComponentUpdateEvent) GetComponent() []byte {
	if x != nil {
		return x.Component
	}
	return nil
}

// ListComponentResponse includes the list of available components.
type ListComponentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Components [][]byte `protobuf:"bytes,1,rep,name=components,proto3" json:"components,omitempty"`
}

func (x *ListComponentResponse) Reset() {
	*x = ListComponentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dapr_proto_operator_v1_operator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListComponentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListComponentResponse) ProtoMessage() {}

func (x *ListComponentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dapr_proto_operator_v1_operator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListComponentResponse.ProtoReflect.Descriptor instead.
func (*ListComponentResponse) Descriptor() ([]byte, []int) {
	return file_dapr_proto_operator_v1_operator_proto_rawDescGZIP(), []int{3}
}

func (x *ListComponentResponse) GetComponents() [][]byte {
	if x != nil {
		return x.Components
	}
	return nil
}

// GetConfigurationRequest is the request message to get the configuration.
type GetConfigurationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	PodName   string `protobuf:"bytes,3,opt,name=podName,proto3" json:"podName,omitempty"`
}

func (x *GetConfigurationRequest) Reset() {
	*x = GetConfigurationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dapr_proto_operator_v1_operator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigurationRequest) ProtoMessage() {}

func (x *GetConfigurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dapr_proto_operator_v1_operator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigurationRequest.ProtoReflect.Descriptor instead.
func (*GetConfigurationRequest) Descriptor() ([]byte, []int) {
	return file_dapr_proto_operator_v1_operator_proto_rawDescGZIP(), []int{4}
}

func (x *GetConfigurationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetConfigurationRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GetConfigurationRequest) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

// GetConfigurationResponse includes the requested configuration.
type GetConfigurationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configuration []byte `protobuf:"bytes,1,opt,name=configuration,proto3" json:"configuration,omitempty"`
}

func (x *GetConfigurationResponse) Reset() {
	*x = GetConfigurationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dapr_proto_operator_v1_operator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigurationResponse) ProtoMessage() {}

func (x *GetConfigurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dapr_proto_operator_v1_operator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigurationResponse.ProtoReflect.Descriptor instead.
func (*GetConfigurationResponse) Descriptor() ([]byte, []int) {
	return file_dapr_proto_operator_v1_operator_proto_rawDescGZIP(), []int{5}
}

func (x *GetConfigurationResponse) GetConfiguration() []byte {
	if x != nil {
		return x.Configuration
	}
	return nil
}

// ListSubscriptionsResponse includes pub/sub subscriptions.
type ListSubscriptionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscriptions [][]byte `protobuf:"bytes,1,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
}

func (x *ListSubscriptionsResponse) Reset() {
	*x = ListSubscriptionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dapr_proto_operator_v1_operator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubscriptionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubscriptionsResponse) ProtoMessage() {}

func (x *ListSubscriptionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dapr_proto_operator_v1_operator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubscriptionsResponse.ProtoReflect.Descriptor instead.
func (*ListSubscriptionsResponse) Descriptor() ([]byte, []int) {
	return file_dapr_proto_operator_v1_operator_proto_rawDescGZIP(), []int{6}
}

func (x *ListSubscriptionsResponse) GetSubscriptions() [][]byte {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

type ListSubscriptionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodName   string `protobuf:"bytes,1,opt,name=podName,proto3" json:"podName,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *ListSubscriptionsRequest) Reset() {
	*x = ListSubscriptionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dapr_proto_operator_v1_operator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubscriptionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubscriptionsRequest) ProtoMessage() {}

func (x *ListSubscriptionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dapr_proto_operator_v1_operator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubscriptionsRequest.ProtoReflect.Descriptor instead.
func (*ListSubscriptionsRequest) Descriptor() ([]byte, []int) {
	return file_dapr_proto_operator_v1_operator_proto_rawDescGZIP(), []int{7}
}

func (x *ListSubscriptionsRequest) GetPodName() string {
	if x != nil {
		return x.PodName
	}
	return ""
}

func (x *ListSubscriptionsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

var File_dapr_proto_operator_v1_operator_proto protoreflect.FileDescriptor

var file_dapr_proto_operator_v1_operator_proto_rawDesc = []byte{
	0x0a, 0x25, 0x64, 0x61, 0x70, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x15,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x50, 0x0a,
	0x16, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x34, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x22, 0x37, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x65,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x40, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x41, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0d, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x52, 0x0a, 0x18, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x32, 0xca,
	0x04, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x73, 0x0a, 0x0f, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2e,
	0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x70, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x2d, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2d, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x77, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x31, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7c, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x56, 0x32, 0x12, 0x30, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x64, 0x61, 0x70, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x70, 0x72, 0x2f, 0x64,
	0x61, 0x70, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dapr_proto_operator_v1_operator_proto_rawDescOnce sync.Once
	file_dapr_proto_operator_v1_operator_proto_rawDescData = file_dapr_proto_operator_v1_operator_proto_rawDesc
)

func file_dapr_proto_operator_v1_operator_proto_rawDescGZIP() []byte {
	file_dapr_proto_operator_v1_operator_proto_rawDescOnce.Do(func() {
		file_dapr_proto_operator_v1_operator_proto_rawDescData = protoimpl.X.CompressGZIP(file_dapr_proto_operator_v1_operator_proto_rawDescData)
	})
	return file_dapr_proto_operator_v1_operator_proto_rawDescData
}

var file_dapr_proto_operator_v1_operator_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_dapr_proto_operator_v1_operator_proto_goTypes = []interface{}{
	(*ListComponentsRequest)(nil),     // 0: dapr.proto.operator.v1.ListComponentsRequest
	(*ComponentUpdateRequest)(nil),    // 1: dapr.proto.operator.v1.ComponentUpdateRequest
	(*ComponentUpdateEvent)(nil),      // 2: dapr.proto.operator.v1.ComponentUpdateEvent
	(*ListComponentResponse)(nil),     // 3: dapr.proto.operator.v1.ListComponentResponse
	(*GetConfigurationRequest)(nil),   // 4: dapr.proto.operator.v1.GetConfigurationRequest
	(*GetConfigurationResponse)(nil),  // 5: dapr.proto.operator.v1.GetConfigurationResponse
	(*ListSubscriptionsResponse)(nil), // 6: dapr.proto.operator.v1.ListSubscriptionsResponse
	(*ListSubscriptionsRequest)(nil),  // 7: dapr.proto.operator.v1.ListSubscriptionsRequest
	(*emptypb.Empty)(nil),             // 8: google.protobuf.Empty
}
var file_dapr_proto_operator_v1_operator_proto_depIdxs = []int32{
	1, // 0: dapr.proto.operator.v1.Operator.ComponentUpdate:input_type -> dapr.proto.operator.v1.ComponentUpdateRequest
	0, // 1: dapr.proto.operator.v1.Operator.ListComponents:input_type -> dapr.proto.operator.v1.ListComponentsRequest
	4, // 2: dapr.proto.operator.v1.Operator.GetConfiguration:input_type -> dapr.proto.operator.v1.GetConfigurationRequest
	8, // 3: dapr.proto.operator.v1.Operator.ListSubscriptions:input_type -> google.protobuf.Empty
	7, // 4: dapr.proto.operator.v1.Operator.ListSubscriptionsV2:input_type -> dapr.proto.operator.v1.ListSubscriptionsRequest
	2, // 5: dapr.proto.operator.v1.Operator.ComponentUpdate:output_type -> dapr.proto.operator.v1.ComponentUpdateEvent
	3, // 6: dapr.proto.operator.v1.Operator.ListComponents:output_type -> dapr.proto.operator.v1.ListComponentResponse
	5, // 7: dapr.proto.operator.v1.Operator.GetConfiguration:output_type -> dapr.proto.operator.v1.GetConfigurationResponse
	6, // 8: dapr.proto.operator.v1.Operator.ListSubscriptions:output_type -> dapr.proto.operator.v1.ListSubscriptionsResponse
	6, // 9: dapr.proto.operator.v1.Operator.ListSubscriptionsV2:output_type -> dapr.proto.operator.v1.ListSubscriptionsResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dapr_proto_operator_v1_operator_proto_init() }
func file_dapr_proto_operator_v1_operator_proto_init() {
	if File_dapr_proto_operator_v1_operator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dapr_proto_operator_v1_operator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListComponentsRequest); i {
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
		file_dapr_proto_operator_v1_operator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComponentUpdateRequest); i {
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
		file_dapr_proto_operator_v1_operator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComponentUpdateEvent); i {
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
		file_dapr_proto_operator_v1_operator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListComponentResponse); i {
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
		file_dapr_proto_operator_v1_operator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigurationRequest); i {
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
		file_dapr_proto_operator_v1_operator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigurationResponse); i {
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
		file_dapr_proto_operator_v1_operator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSubscriptionsResponse); i {
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
		file_dapr_proto_operator_v1_operator_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSubscriptionsRequest); i {
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
			RawDescriptor: file_dapr_proto_operator_v1_operator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dapr_proto_operator_v1_operator_proto_goTypes,
		DependencyIndexes: file_dapr_proto_operator_v1_operator_proto_depIdxs,
		MessageInfos:      file_dapr_proto_operator_v1_operator_proto_msgTypes,
	}.Build()
	File_dapr_proto_operator_v1_operator_proto = out.File
	file_dapr_proto_operator_v1_operator_proto_rawDesc = nil
	file_dapr_proto_operator_v1_operator_proto_goTypes = nil
	file_dapr_proto_operator_v1_operator_proto_depIdxs = nil
}
