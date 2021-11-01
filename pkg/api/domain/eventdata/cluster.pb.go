// Copyright 2021 Monoskope Authors
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: api/domain/eventdata/cluster.proto

package eventdata

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClusterCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Display name of the cluster
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Unique name of the cluster, to be utilized for generating unique labels
	// and symbols, e.g. with metrics.
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	// Address of the KubeApiServer
	ApiServerAddress string `protobuf:"bytes,3,opt,name=api_server_address,json=apiServerAddress,proto3" json:"api_server_address,omitempty"`
	// CA certificate of the K8s cluster
	CaCertificateBundle []byte `protobuf:"bytes,4,opt,name=ca_certificate_bundle,json=caCertificateBundle,proto3" json:"ca_certificate_bundle,omitempty"`
}

func (x *ClusterCreated) Reset() {
	*x = ClusterCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_domain_eventdata_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterCreated) ProtoMessage() {}

func (x *ClusterCreated) ProtoReflect() protoreflect.Message {
	mi := &file_api_domain_eventdata_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterCreated.ProtoReflect.Descriptor instead.
func (*ClusterCreated) Descriptor() ([]byte, []int) {
	return file_api_domain_eventdata_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterCreated) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClusterCreated) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *ClusterCreated) GetApiServerAddress() string {
	if x != nil {
		return x.ApiServerAddress
	}
	return ""
}

func (x *ClusterCreated) GetCaCertificateBundle() []byte {
	if x != nil {
		return x.CaCertificateBundle
	}
	return nil
}

type ClusterCreatedV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique name of the cluster, to be utilized for generating unique labels
	// and symbols, e.g. with metrics.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Display name of the cluster
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Address of the KubeApiServer
	ApiServerAddress string `protobuf:"bytes,3,opt,name=api_server_address,json=apiServerAddress,proto3" json:"api_server_address,omitempty"`
	// CA certificate of the K8s cluster
	CaCertificateBundle []byte `protobuf:"bytes,4,opt,name=ca_certificate_bundle,json=caCertificateBundle,proto3" json:"ca_certificate_bundle,omitempty"`
}

func (x *ClusterCreatedV2) Reset() {
	*x = ClusterCreatedV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_domain_eventdata_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterCreatedV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterCreatedV2) ProtoMessage() {}

func (x *ClusterCreatedV2) ProtoReflect() protoreflect.Message {
	mi := &file_api_domain_eventdata_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterCreatedV2.ProtoReflect.Descriptor instead.
func (*ClusterCreatedV2) Descriptor() ([]byte, []int) {
	return file_api_domain_eventdata_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterCreatedV2) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClusterCreatedV2) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *ClusterCreatedV2) GetApiServerAddress() string {
	if x != nil {
		return x.ApiServerAddress
	}
	return ""
}

func (x *ClusterCreatedV2) GetCaCertificateBundle() []byte {
	if x != nil {
		return x.CaCertificateBundle
	}
	return nil
}

type ClusterUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Display name of the cluster
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Address of the KubeApiServer
	ApiServerAddress string `protobuf:"bytes,2,opt,name=api_server_address,json=apiServerAddress,proto3" json:"api_server_address,omitempty"`
	// CA certificate of the K8s cluster
	CaCertificateBundle []byte `protobuf:"bytes,3,opt,name=ca_certificate_bundle,json=caCertificateBundle,proto3" json:"ca_certificate_bundle,omitempty"`
}

func (x *ClusterUpdated) Reset() {
	*x = ClusterUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_domain_eventdata_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterUpdated) ProtoMessage() {}

func (x *ClusterUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_api_domain_eventdata_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterUpdated.ProtoReflect.Descriptor instead.
func (*ClusterUpdated) Descriptor() ([]byte, []int) {
	return file_api_domain_eventdata_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *ClusterUpdated) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *ClusterUpdated) GetApiServerAddress() string {
	if x != nil {
		return x.ApiServerAddress
	}
	return ""
}

func (x *ClusterUpdated) GetCaCertificateBundle() []byte {
	if x != nil {
		return x.CaCertificateBundle
	}
	return nil
}

type ClusterBootstrapTokenCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// JsonWebToken for cluster bootstrap
	Jwt string `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

func (x *ClusterBootstrapTokenCreated) Reset() {
	*x = ClusterBootstrapTokenCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_domain_eventdata_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterBootstrapTokenCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterBootstrapTokenCreated) ProtoMessage() {}

func (x *ClusterBootstrapTokenCreated) ProtoReflect() protoreflect.Message {
	mi := &file_api_domain_eventdata_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterBootstrapTokenCreated.ProtoReflect.Descriptor instead.
func (*ClusterBootstrapTokenCreated) Descriptor() ([]byte, []int) {
	return file_api_domain_eventdata_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *ClusterBootstrapTokenCreated) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

var File_api_domain_eventdata_cluster_proto protoreflect.FileDescriptor

var file_api_domain_eventdata_cluster_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x9c, 0x01, 0x0a, 0x0e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x2c, 0x0a, 0x12,
	0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x70, 0x69, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x61,
	0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x63, 0x61, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0xab,
	0x01, 0x0a, 0x10, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x56, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x70,
	0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x61, 0x5f, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x95, 0x01, 0x0a,
	0x0e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x61, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x32, 0x0a, 0x15, 0x63, 0x61, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x5f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x13, 0x63, 0x61, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x22, 0x30, 0x0a, 0x1c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x42,
	0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6e, 0x6c, 0x65, 0x61, 0x70, 0x2d, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2f, 0x6d, 0x6f, 0x6e, 0x6f, 0x73, 0x6b, 0x6f, 0x70, 0x65, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_domain_eventdata_cluster_proto_rawDescOnce sync.Once
	file_api_domain_eventdata_cluster_proto_rawDescData = file_api_domain_eventdata_cluster_proto_rawDesc
)

func file_api_domain_eventdata_cluster_proto_rawDescGZIP() []byte {
	file_api_domain_eventdata_cluster_proto_rawDescOnce.Do(func() {
		file_api_domain_eventdata_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_domain_eventdata_cluster_proto_rawDescData)
	})
	return file_api_domain_eventdata_cluster_proto_rawDescData
}

var file_api_domain_eventdata_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_domain_eventdata_cluster_proto_goTypes = []interface{}{
	(*ClusterCreated)(nil),               // 0: eventdata.ClusterCreated
	(*ClusterCreatedV2)(nil),             // 1: eventdata.ClusterCreatedV2
	(*ClusterUpdated)(nil),               // 2: eventdata.ClusterUpdated
	(*ClusterBootstrapTokenCreated)(nil), // 3: eventdata.ClusterBootstrapTokenCreated
}
var file_api_domain_eventdata_cluster_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_domain_eventdata_cluster_proto_init() }
func file_api_domain_eventdata_cluster_proto_init() {
	if File_api_domain_eventdata_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_domain_eventdata_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterCreated); i {
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
		file_api_domain_eventdata_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterCreatedV2); i {
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
		file_api_domain_eventdata_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterUpdated); i {
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
		file_api_domain_eventdata_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterBootstrapTokenCreated); i {
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
			RawDescriptor: file_api_domain_eventdata_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_domain_eventdata_cluster_proto_goTypes,
		DependencyIndexes: file_api_domain_eventdata_cluster_proto_depIdxs,
		MessageInfos:      file_api_domain_eventdata_cluster_proto_msgTypes,
	}.Build()
	File_api_domain_eventdata_cluster_proto = out.File
	file_api_domain_eventdata_cluster_proto_rawDesc = nil
	file_api_domain_eventdata_cluster_proto_goTypes = nil
	file_api_domain_eventdata_cluster_proto_depIdxs = nil
}
