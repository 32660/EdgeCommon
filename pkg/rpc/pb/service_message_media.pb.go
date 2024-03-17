// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: service_message_media.proto

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

// 获取所有支持的媒介
type FindAllMessageMediasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllMessageMediasRequest) Reset() {
	*x = FindAllMessageMediasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_media_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllMessageMediasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllMessageMediasRequest) ProtoMessage() {}

func (x *FindAllMessageMediasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_media_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllMessageMediasRequest.ProtoReflect.Descriptor instead.
func (*FindAllMessageMediasRequest) Descriptor() ([]byte, []int) {
	return file_service_message_media_proto_rawDescGZIP(), []int{0}
}

type FindAllMessageMediasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageMedias []*MessageMedia `protobuf:"bytes,1,rep,name=messageMedias,proto3" json:"messageMedias,omitempty"`
}

func (x *FindAllMessageMediasResponse) Reset() {
	*x = FindAllMessageMediasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_media_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllMessageMediasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllMessageMediasResponse) ProtoMessage() {}

func (x *FindAllMessageMediasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_media_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllMessageMediasResponse.ProtoReflect.Descriptor instead.
func (*FindAllMessageMediasResponse) Descriptor() ([]byte, []int) {
	return file_service_message_media_proto_rawDescGZIP(), []int{1}
}

func (x *FindAllMessageMediasResponse) GetMessageMedias() []*MessageMedia {
	if x != nil {
		return x.MessageMedias
	}
	return nil
}

// 发送媒介信息
type SendMediaMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MediaType   string `protobuf:"bytes,1,opt,name=mediaType,proto3" json:"mediaType,omitempty"`     // 媒介类型
	OptionsJSON []byte `protobuf:"bytes,2,opt,name=optionsJSON,proto3" json:"optionsJSON,omitempty"` // 媒介参数
	User        string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`               // 接收用户
	Subject     string `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`         // 标题
	Body        string `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`               // 内容
}

func (x *SendMediaMessageRequest) Reset() {
	*x = SendMediaMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_message_media_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMediaMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMediaMessageRequest) ProtoMessage() {}

func (x *SendMediaMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_message_media_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMediaMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMediaMessageRequest) Descriptor() ([]byte, []int) {
	return file_service_message_media_proto_rawDescGZIP(), []int{2}
}

func (x *SendMediaMessageRequest) GetMediaType() string {
	if x != nil {
		return x.MediaType
	}
	return ""
}

func (x *SendMediaMessageRequest) GetOptionsJSON() []byte {
	if x != nil {
		return x.OptionsJSON
	}
	return nil
}

func (x *SendMediaMessageRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *SendMediaMessageRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendMediaMessageRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_service_message_media_proto protoreflect.FileDescriptor

var file_service_message_media_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x20, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d,
	0x0a, 0x1b, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x56, 0x0a,
	0x1c, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a,
	0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x73, 0x22, 0x9b, 0x01, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4a, 0x53, 0x4f,
	0x4e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x32, 0xb1, 0x01, 0x0a, 0x13, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d,
	0x65, 0x64, 0x69, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x14, 0x66,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x10, 0x73, 0x65, 0x6e, 0x64, 0x4d, 0x65,
	0x64, 0x69, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_message_media_proto_rawDescOnce sync.Once
	file_service_message_media_proto_rawDescData = file_service_message_media_proto_rawDesc
)

func file_service_message_media_proto_rawDescGZIP() []byte {
	file_service_message_media_proto_rawDescOnce.Do(func() {
		file_service_message_media_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_message_media_proto_rawDescData)
	})
	return file_service_message_media_proto_rawDescData
}

var file_service_message_media_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_service_message_media_proto_goTypes = []interface{}{
	(*FindAllMessageMediasRequest)(nil),  // 0: pb.FindAllMessageMediasRequest
	(*FindAllMessageMediasResponse)(nil), // 1: pb.FindAllMessageMediasResponse
	(*SendMediaMessageRequest)(nil),      // 2: pb.SendMediaMessageRequest
	(*MessageMedia)(nil),                 // 3: pb.MessageMedia
	(*RPCSuccess)(nil),                   // 4: pb.RPCSuccess
}
var file_service_message_media_proto_depIdxs = []int32{
	3, // 0: pb.FindAllMessageMediasResponse.messageMedias:type_name -> pb.MessageMedia
	0, // 1: pb.MessageMediaService.findAllMessageMedias:input_type -> pb.FindAllMessageMediasRequest
	2, // 2: pb.MessageMediaService.sendMediaMessage:input_type -> pb.SendMediaMessageRequest
	1, // 3: pb.MessageMediaService.findAllMessageMedias:output_type -> pb.FindAllMessageMediasResponse
	4, // 4: pb.MessageMediaService.sendMediaMessage:output_type -> pb.RPCSuccess
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_message_media_proto_init() }
func file_service_message_media_proto_init() {
	if File_service_message_media_proto != nil {
		return
	}
	file_models_model_message_media_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_message_media_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllMessageMediasRequest); i {
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
		file_service_message_media_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllMessageMediasResponse); i {
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
		file_service_message_media_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMediaMessageRequest); i {
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
			RawDescriptor: file_service_message_media_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_message_media_proto_goTypes,
		DependencyIndexes: file_service_message_media_proto_depIdxs,
		MessageInfos:      file_service_message_media_proto_msgTypes,
	}.Build()
	File_service_message_media_proto = out.File
	file_service_message_media_proto_rawDesc = nil
	file_service_message_media_proto_goTypes = nil
	file_service_message_media_proto_depIdxs = nil
}
