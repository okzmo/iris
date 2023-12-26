// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.2
// source: rpc_create_server.proto

package pb

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

type CreateServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *CreateServerRequest) Reset() {
	*x = CreateServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServerRequest) ProtoMessage() {}

func (x *CreateServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServerRequest.ProtoReflect.Descriptor instead.
func (*CreateServerRequest) Descriptor() ([]byte, []int) {
	return file_rpc_create_server_proto_rawDescGZIP(), []int{0}
}

func (x *CreateServerRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type CreateServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateServerResponse) Reset() {
	*x = CreateServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_create_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServerResponse) ProtoMessage() {}

func (x *CreateServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_create_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateServerResponse.ProtoReflect.Descriptor instead.
func (*CreateServerResponse) Descriptor() ([]byte, []int) {
	return file_rpc_create_server_proto_rawDescGZIP(), []int{1}
}

func (x *CreateServerResponse) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

var File_rpc_create_server_proto protoreflect.FileDescriptor

var file_rpc_create_server_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x31, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x2a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x42, 0x1c, 0x5a, 0x1a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6b, 0x7a, 0x6d, 0x6f,
	0x2f, 0x6e, 0x6f, 0x6e, 0x61, 0x6d, 0x65, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_rpc_create_server_proto_rawDescOnce sync.Once
	file_rpc_create_server_proto_rawDescData = file_rpc_create_server_proto_rawDesc
)

func file_rpc_create_server_proto_rawDescGZIP() []byte {
	file_rpc_create_server_proto_rawDescOnce.Do(func() {
		file_rpc_create_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_create_server_proto_rawDescData)
	})
	return file_rpc_create_server_proto_rawDescData
}

var file_rpc_create_server_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_create_server_proto_goTypes = []interface{}{
	(*CreateServerRequest)(nil),  // 0: pb.CreateServerRequest
	(*CreateServerResponse)(nil), // 1: pb.CreateServerResponse
}
var file_rpc_create_server_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rpc_create_server_proto_init() }
func file_rpc_create_server_proto_init() {
	if File_rpc_create_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_create_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServerRequest); i {
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
		file_rpc_create_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateServerResponse); i {
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
			RawDescriptor: file_rpc_create_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_create_server_proto_goTypes,
		DependencyIndexes: file_rpc_create_server_proto_depIdxs,
		MessageInfos:      file_rpc_create_server_proto_msgTypes,
	}.Build()
	File_rpc_create_server_proto = out.File
	file_rpc_create_server_proto_rawDesc = nil
	file_rpc_create_server_proto_goTypes = nil
	file_rpc_create_server_proto_depIdxs = nil
}