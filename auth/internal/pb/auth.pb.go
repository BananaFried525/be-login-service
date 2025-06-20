// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: auth.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetAnonymousTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ApiKey        string                 `protobuf:"bytes,1,opt,name=apiKey,proto3" json:"apiKey,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAnonymousTokenRequest) Reset() {
	*x = GetAnonymousTokenRequest{}
	mi := &file_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAnonymousTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnonymousTokenRequest) ProtoMessage() {}

func (x *GetAnonymousTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnonymousTokenRequest.ProtoReflect.Descriptor instead.
func (*GetAnonymousTokenRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *GetAnonymousTokenRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

type GetAnonymousTokenResponse struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Token           string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	TokenExpiration string                 `protobuf:"bytes,2,opt,name=tokenExpiration,proto3" json:"tokenExpiration,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *GetAnonymousTokenResponse) Reset() {
	*x = GetAnonymousTokenResponse{}
	mi := &file_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAnonymousTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnonymousTokenResponse) ProtoMessage() {}

func (x *GetAnonymousTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnonymousTokenResponse.ProtoReflect.Descriptor instead.
func (*GetAnonymousTokenResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *GetAnonymousTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetAnonymousTokenResponse) GetTokenExpiration() string {
	if x != nil {
		return x.TokenExpiration
	}
	return ""
}

var File_auth_proto protoreflect.FileDescriptor

const file_auth_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"auth.proto\x12\x04grpc\"2\n" +
	"\x18GetAnonymousTokenRequest\x12\x16\n" +
	"\x06apiKey\x18\x01 \x01(\tR\x06apiKey\"[\n" +
	"\x19GetAnonymousTokenResponse\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\x12(\n" +
	"\x0ftokenExpiration\x18\x02 \x01(\tR\x0ftokenExpiration2c\n" +
	"\vAuthService\x12T\n" +
	"\x11GetAnonymousToken\x12\x1e.grpc.GetAnonymousTokenRequest\x1a\x1f.grpc.GetAnonymousTokenResponseB\x1bZ\x19go-grpc-clean/internal/pbb\x06proto3"

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData []byte
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_auth_proto_rawDesc), len(file_auth_proto_rawDesc)))
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auth_proto_goTypes = []any{
	(*GetAnonymousTokenRequest)(nil),  // 0: grpc.GetAnonymousTokenRequest
	(*GetAnonymousTokenResponse)(nil), // 1: grpc.GetAnonymousTokenResponse
}
var file_auth_proto_depIdxs = []int32{
	0, // 0: grpc.AuthService.GetAnonymousToken:input_type -> grpc.GetAnonymousTokenRequest
	1, // 1: grpc.AuthService.GetAnonymousToken:output_type -> grpc.GetAnonymousTokenResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_auth_proto_rawDesc), len(file_auth_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}
