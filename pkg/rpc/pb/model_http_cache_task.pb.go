// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: models/model_http_cache_task.proto

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

type HTTPCacheTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // 任务ID
	UserId            int64               `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Type              string              `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	KeyType           string              `protobuf:"bytes,4,opt,name=keyType,proto3" json:"keyType,omitempty"`
	CreatedAt         int64               `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	DoneAt            int64               `protobuf:"varint,6,opt,name=doneAt,proto3" json:"doneAt,omitempty"`
	IsDone            bool                `protobuf:"varint,7,opt,name=isDone,proto3" json:"isDone,omitempty"`
	IsOk              bool                `protobuf:"varint,8,opt,name=isOk,proto3" json:"isOk,omitempty"`
	Description       string              `protobuf:"bytes,9,opt,name=description,proto3" json:"description,omitempty"`
	User              *User               `protobuf:"bytes,30,opt,name=user,proto3" json:"user,omitempty"`                           // 所属用户
	HttpCacheTaskKeys []*HTTPCacheTaskKey `protobuf:"bytes,31,rep,name=httpCacheTaskKeys,proto3" json:"httpCacheTaskKeys,omitempty"` // 包含的Key
}

func (x *HTTPCacheTask) Reset() {
	*x = HTTPCacheTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_http_cache_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPCacheTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPCacheTask) ProtoMessage() {}

func (x *HTTPCacheTask) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_http_cache_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPCacheTask.ProtoReflect.Descriptor instead.
func (*HTTPCacheTask) Descriptor() ([]byte, []int) {
	return file_models_model_http_cache_task_proto_rawDescGZIP(), []int{0}
}

func (x *HTTPCacheTask) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HTTPCacheTask) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *HTTPCacheTask) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *HTTPCacheTask) GetKeyType() string {
	if x != nil {
		return x.KeyType
	}
	return ""
}

func (x *HTTPCacheTask) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *HTTPCacheTask) GetDoneAt() int64 {
	if x != nil {
		return x.DoneAt
	}
	return 0
}

func (x *HTTPCacheTask) GetIsDone() bool {
	if x != nil {
		return x.IsDone
	}
	return false
}

func (x *HTTPCacheTask) GetIsOk() bool {
	if x != nil {
		return x.IsOk
	}
	return false
}

func (x *HTTPCacheTask) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *HTTPCacheTask) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *HTTPCacheTask) GetHttpCacheTaskKeys() []*HTTPCacheTaskKey {
	if x != nil {
		return x.HttpCacheTaskKeys
	}
	return nil
}

var File_models_model_http_cache_task_proto protoreflect.FileDescriptor

var file_models_model_http_cache_task_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x26, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x02, 0x0a, 0x0d, 0x48, 0x54,
	0x54, 0x50, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6e, 0x65, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x64, 0x6f, 0x6e, 0x65, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x44, 0x6f, 0x6e,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69,
	0x73, 0x4f, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x11, 0x68, 0x74, 0x74, 0x70, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x4b, 0x65, 0x79, 0x73, 0x18, 0x1f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x70, 0x62, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x4b, 0x65, 0x79, 0x52, 0x11, 0x68, 0x74, 0x74, 0x70, 0x43, 0x61, 0x63, 0x68, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x4b, 0x65, 0x79, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_http_cache_task_proto_rawDescOnce sync.Once
	file_models_model_http_cache_task_proto_rawDescData = file_models_model_http_cache_task_proto_rawDesc
)

func file_models_model_http_cache_task_proto_rawDescGZIP() []byte {
	file_models_model_http_cache_task_proto_rawDescOnce.Do(func() {
		file_models_model_http_cache_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_http_cache_task_proto_rawDescData)
	})
	return file_models_model_http_cache_task_proto_rawDescData
}

var file_models_model_http_cache_task_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_http_cache_task_proto_goTypes = []interface{}{
	(*HTTPCacheTask)(nil),    // 0: pb.HTTPCacheTask
	(*User)(nil),             // 1: pb.User
	(*HTTPCacheTaskKey)(nil), // 2: pb.HTTPCacheTaskKey
}
var file_models_model_http_cache_task_proto_depIdxs = []int32{
	1, // 0: pb.HTTPCacheTask.user:type_name -> pb.User
	2, // 1: pb.HTTPCacheTask.httpCacheTaskKeys:type_name -> pb.HTTPCacheTaskKey
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_models_model_http_cache_task_proto_init() }
func file_models_model_http_cache_task_proto_init() {
	if File_models_model_http_cache_task_proto != nil {
		return
	}
	file_models_model_user_proto_init()
	file_models_model_http_cache_task_key_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_http_cache_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPCacheTask); i {
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
			RawDescriptor: file_models_model_http_cache_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_http_cache_task_proto_goTypes,
		DependencyIndexes: file_models_model_http_cache_task_proto_depIdxs,
		MessageInfos:      file_models_model_http_cache_task_proto_msgTypes,
	}.Build()
	File_models_model_http_cache_task_proto = out.File
	file_models_model_http_cache_task_proto_rawDesc = nil
	file_models_model_http_cache_task_proto_goTypes = nil
	file_models_model_http_cache_task_proto_depIdxs = nil
}
