// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.24.4
// source: hermes.proto

package hermes

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

type HermesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Headers       []string               `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	Body          string                 `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HermesRequest) Reset() {
	*x = HermesRequest{}
	mi := &file_hermes_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HermesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HermesRequest) ProtoMessage() {}

func (x *HermesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hermes_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HermesRequest.ProtoReflect.Descriptor instead.
func (*HermesRequest) Descriptor() ([]byte, []int) {
	return file_hermes_proto_rawDescGZIP(), []int{0}
}

func (x *HermesRequest) GetHeaders() []string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *HermesRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type HermesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        string                 `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HermesResponse) Reset() {
	*x = HermesResponse{}
	mi := &file_hermes_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HermesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HermesResponse) ProtoMessage() {}

func (x *HermesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hermes_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HermesResponse.ProtoReflect.Descriptor instead.
func (*HermesResponse) Descriptor() ([]byte, []int) {
	return file_hermes_proto_rawDescGZIP(), []int{1}
}

func (x *HermesResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_hermes_proto protoreflect.FileDescriptor

const file_hermes_proto_rawDesc = "" +
	"\n" +
	"\fhermes.proto\x12\x06hermes\"=\n" +
	"\rHermesRequest\x12\x18\n" +
	"\aheaders\x18\x01 \x03(\tR\aheaders\x12\x12\n" +
	"\x04body\x18\x02 \x01(\tR\x04body\"(\n" +
	"\x0eHermesResponse\x12\x16\n" +
	"\x06result\x18\x01 \x01(\tR\x06result2H\n" +
	"\rHermesHandler\x127\n" +
	"\x06Handle\x12\x15.hermes.HermesRequest\x1a\x16.hermes.HermesResponseB!\xca\x02\x1eApp\\Infrastructure\\Grpc\\Hermesb\x06proto3"

var (
	file_hermes_proto_rawDescOnce sync.Once
	file_hermes_proto_rawDescData []byte
)

func file_hermes_proto_rawDescGZIP() []byte {
	file_hermes_proto_rawDescOnce.Do(func() {
		file_hermes_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_hermes_proto_rawDesc), len(file_hermes_proto_rawDesc)))
	})
	return file_hermes_proto_rawDescData
}

var file_hermes_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_hermes_proto_goTypes = []any{
	(*HermesRequest)(nil),  // 0: hermes.HermesRequest
	(*HermesResponse)(nil), // 1: hermes.HermesResponse
}
var file_hermes_proto_depIdxs = []int32{
	0, // 0: hermes.HermesHandler.Handle:input_type -> hermes.HermesRequest
	1, // 1: hermes.HermesHandler.Handle:output_type -> hermes.HermesResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hermes_proto_init() }
func file_hermes_proto_init() {
	if File_hermes_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_hermes_proto_rawDesc), len(file_hermes_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hermes_proto_goTypes,
		DependencyIndexes: file_hermes_proto_depIdxs,
		MessageInfos:      file_hermes_proto_msgTypes,
	}.Build()
	File_hermes_proto = out.File
	file_hermes_proto_goTypes = nil
	file_hermes_proto_depIdxs = nil
}
