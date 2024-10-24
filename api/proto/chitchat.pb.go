// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: api/proto/chitchat.proto

package proto

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

// The request message containing the user's name.
type FromClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Body string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *FromClient) Reset() {
	*x = FromClient{}
	mi := &file_api_proto_chitchat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FromClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FromClient) ProtoMessage() {}

func (x *FromClient) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_chitchat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FromClient.ProtoReflect.Descriptor instead.
func (*FromClient) Descriptor() ([]byte, []int) {
	return file_api_proto_chitchat_proto_rawDescGZIP(), []int{0}
}

func (x *FromClient) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FromClient) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

// The response message containing the greetings.
type FromServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Body string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *FromServer) Reset() {
	*x = FromServer{}
	mi := &file_api_proto_chitchat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FromServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FromServer) ProtoMessage() {}

func (x *FromServer) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_chitchat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FromServer.ProtoReflect.Descriptor instead.
func (*FromServer) Descriptor() ([]byte, []int) {
	return file_api_proto_chitchat_proto_rawDescGZIP(), []int{1}
}

func (x *FromServer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FromServer) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_api_proto_chitchat_proto protoreflect.FileDescriptor

var file_api_proto_chitchat_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x69, 0x74,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x68, 0x69, 0x74,
	0x63, 0x68, 0x61, 0x74, 0x22, 0x34, 0x0a, 0x0a, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x34, 0x0a, 0x0a, 0x46, 0x72,
	0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x32, 0x4b, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x0b,
	0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x14, 0x2e, 0x63, 0x68,
	0x69, 0x74, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x1a, 0x14, 0x2e, 0x63, 0x68, 0x69, 0x74, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x46, 0x72, 0x6f,
	0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x42, 0x5a,
	0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x69, 0x6b, 0x74,
	0x6f, 0x72, 0x45, 0x6d, 0x69, 0x6c, 0x32, 0x30, 0x30, 0x30, 0x2f, 0x44, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x73, 0x2d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2d, 0x43, 0x61, 0x74, 0x61, 0x6e, 0x45,
	0x6e, 0x6a, 0x6f, 0x79, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_chitchat_proto_rawDescOnce sync.Once
	file_api_proto_chitchat_proto_rawDescData = file_api_proto_chitchat_proto_rawDesc
)

func file_api_proto_chitchat_proto_rawDescGZIP() []byte {
	file_api_proto_chitchat_proto_rawDescOnce.Do(func() {
		file_api_proto_chitchat_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_chitchat_proto_rawDescData)
	})
	return file_api_proto_chitchat_proto_rawDescData
}

var file_api_proto_chitchat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_chitchat_proto_goTypes = []any{
	(*FromClient)(nil), // 0: chitchat.FromClient
	(*FromServer)(nil), // 1: chitchat.FromServer
}
var file_api_proto_chitchat_proto_depIdxs = []int32{
	0, // 0: chitchat.Services.ChatService:input_type -> chitchat.FromClient
	1, // 1: chitchat.Services.ChatService:output_type -> chitchat.FromServer
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_chitchat_proto_init() }
func file_api_proto_chitchat_proto_init() {
	if File_api_proto_chitchat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_chitchat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_chitchat_proto_goTypes,
		DependencyIndexes: file_api_proto_chitchat_proto_depIdxs,
		MessageInfos:      file_api_proto_chitchat_proto_msgTypes,
	}.Build()
	File_api_proto_chitchat_proto = out.File
	file_api_proto_chitchat_proto_rawDesc = nil
	file_api_proto_chitchat_proto_goTypes = nil
	file_api_proto_chitchat_proto_depIdxs = nil
}
