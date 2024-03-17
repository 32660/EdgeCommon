// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: models/model_user_bandwidth_stat.proto

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

// 带宽统计数据
type UserBandwidthStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Day    string `protobuf:"bytes,3,opt,name=day,proto3" json:"day,omitempty"`
	TimeAt string `protobuf:"bytes,4,opt,name=timeAt,proto3" json:"timeAt,omitempty"`
	Bytes  int64  `protobuf:"varint,5,opt,name=bytes,proto3" json:"bytes,omitempty"` // 峰值带宽字节
}

func (x *UserBandwidthStat) Reset() {
	*x = UserBandwidthStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_user_bandwidth_stat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBandwidthStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBandwidthStat) ProtoMessage() {}

func (x *UserBandwidthStat) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_user_bandwidth_stat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBandwidthStat.ProtoReflect.Descriptor instead.
func (*UserBandwidthStat) Descriptor() ([]byte, []int) {
	return file_models_model_user_bandwidth_stat_proto_rawDescGZIP(), []int{0}
}

func (x *UserBandwidthStat) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserBandwidthStat) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserBandwidthStat) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *UserBandwidthStat) GetTimeAt() string {
	if x != nil {
		return x.TimeAt
	}
	return ""
}

func (x *UserBandwidthStat) GetBytes() int64 {
	if x != nil {
		return x.Bytes
	}
	return 0
}

var File_models_model_user_bandwidth_stat_proto protoreflect.FileDescriptor

var file_models_model_user_bandwidth_stat_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x7b, 0x0a, 0x11,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x74,
	0x69, 0x6d, 0x65, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x6d,
	0x65, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_user_bandwidth_stat_proto_rawDescOnce sync.Once
	file_models_model_user_bandwidth_stat_proto_rawDescData = file_models_model_user_bandwidth_stat_proto_rawDesc
)

func file_models_model_user_bandwidth_stat_proto_rawDescGZIP() []byte {
	file_models_model_user_bandwidth_stat_proto_rawDescOnce.Do(func() {
		file_models_model_user_bandwidth_stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_user_bandwidth_stat_proto_rawDescData)
	})
	return file_models_model_user_bandwidth_stat_proto_rawDescData
}

var file_models_model_user_bandwidth_stat_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_user_bandwidth_stat_proto_goTypes = []interface{}{
	(*UserBandwidthStat)(nil), // 0: pb.UserBandwidthStat
}
var file_models_model_user_bandwidth_stat_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_user_bandwidth_stat_proto_init() }
func file_models_model_user_bandwidth_stat_proto_init() {
	if File_models_model_user_bandwidth_stat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_user_bandwidth_stat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBandwidthStat); i {
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
			RawDescriptor: file_models_model_user_bandwidth_stat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_user_bandwidth_stat_proto_goTypes,
		DependencyIndexes: file_models_model_user_bandwidth_stat_proto_depIdxs,
		MessageInfos:      file_models_model_user_bandwidth_stat_proto_msgTypes,
	}.Build()
	File_models_model_user_bandwidth_stat_proto = out.File
	file_models_model_user_bandwidth_stat_proto_rawDesc = nil
	file_models_model_user_bandwidth_stat_proto_goTypes = nil
	file_models_model_user_bandwidth_stat_proto_depIdxs = nil
}
