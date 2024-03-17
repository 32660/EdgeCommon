// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: service_server_client_browser_monthly_stat.proto

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

// 查找前N个浏览器
type FindTopServerClientBrowserMonthlyStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId int64  `protobuf:"varint,1,opt,name=serverId,proto3" json:"serverId,omitempty"`
	Month    string `protobuf:"bytes,2,opt,name=month,proto3" json:"month,omitempty"`
	Offset   int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Size     int64  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *FindTopServerClientBrowserMonthlyStatsRequest) Reset() {
	*x = FindTopServerClientBrowserMonthlyStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_client_browser_monthly_stat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindTopServerClientBrowserMonthlyStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindTopServerClientBrowserMonthlyStatsRequest) ProtoMessage() {}

func (x *FindTopServerClientBrowserMonthlyStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_client_browser_monthly_stat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindTopServerClientBrowserMonthlyStatsRequest.ProtoReflect.Descriptor instead.
func (*FindTopServerClientBrowserMonthlyStatsRequest) Descriptor() ([]byte, []int) {
	return file_service_server_client_browser_monthly_stat_proto_rawDescGZIP(), []int{0}
}

func (x *FindTopServerClientBrowserMonthlyStatsRequest) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

func (x *FindTopServerClientBrowserMonthlyStatsRequest) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *FindTopServerClientBrowserMonthlyStatsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *FindTopServerClientBrowserMonthlyStatsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type FindTopServerClientBrowserMonthlyStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stats []*FindTopServerClientBrowserMonthlyStatsResponse_Stat `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (x *FindTopServerClientBrowserMonthlyStatsResponse) Reset() {
	*x = FindTopServerClientBrowserMonthlyStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_client_browser_monthly_stat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindTopServerClientBrowserMonthlyStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindTopServerClientBrowserMonthlyStatsResponse) ProtoMessage() {}

func (x *FindTopServerClientBrowserMonthlyStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_client_browser_monthly_stat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindTopServerClientBrowserMonthlyStatsResponse.ProtoReflect.Descriptor instead.
func (*FindTopServerClientBrowserMonthlyStatsResponse) Descriptor() ([]byte, []int) {
	return file_service_server_client_browser_monthly_stat_proto_rawDescGZIP(), []int{1}
}

func (x *FindTopServerClientBrowserMonthlyStatsResponse) GetStats() []*FindTopServerClientBrowserMonthlyStatsResponse_Stat {
	if x != nil {
		return x.Stats
	}
	return nil
}

type FindTopServerClientBrowserMonthlyStatsResponse_Stat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientBrowser *ClientBrowser `protobuf:"bytes,1,opt,name=clientBrowser,proto3" json:"clientBrowser,omitempty"`
	Version       string         `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Count         int64          `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *FindTopServerClientBrowserMonthlyStatsResponse_Stat) Reset() {
	*x = FindTopServerClientBrowserMonthlyStatsResponse_Stat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_client_browser_monthly_stat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindTopServerClientBrowserMonthlyStatsResponse_Stat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindTopServerClientBrowserMonthlyStatsResponse_Stat) ProtoMessage() {}

func (x *FindTopServerClientBrowserMonthlyStatsResponse_Stat) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_client_browser_monthly_stat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindTopServerClientBrowserMonthlyStatsResponse_Stat.ProtoReflect.Descriptor instead.
func (*FindTopServerClientBrowserMonthlyStatsResponse_Stat) Descriptor() ([]byte, []int) {
	return file_service_server_client_browser_monthly_stat_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FindTopServerClientBrowserMonthlyStatsResponse_Stat) GetClientBrowser() *ClientBrowser {
	if x != nil {
		return x.ClientBrowser
	}
	return nil
}

func (x *FindTopServerClientBrowserMonthlyStatsResponse_Stat) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *FindTopServerClientBrowserMonthlyStatsResponse_Stat) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_service_server_client_browser_monthly_stat_proto protoreflect.FileDescriptor

var file_service_server_client_browser_monthly_stat_proto_rawDesc = []byte{
	0x0a, 0x30, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x5f,
	0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x21, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x72, 0x6f, 0x77,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x2d, 0x46, 0x69,
	0x6e, 0x64, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0xf0, 0x01, 0x0a, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68,
	0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x1a, 0x6f, 0x0a, 0x04, 0x53,
	0x74, 0x61, 0x74, 0x12, 0x37, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x72, 0x6f,
	0x77, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x52, 0x0d, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xb9, 0x01, 0x0a,
	0x25, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x72, 0x6f,
	0x77, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8f, 0x01, 0x0a, 0x26, 0x66, 0x69, 0x6e, 0x64, 0x54,
	0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x72,
	0x6f, 0x77, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x31, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x6f, 0x70, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65,
	0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x6f,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x72, 0x6f,
	0x77, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_server_client_browser_monthly_stat_proto_rawDescOnce sync.Once
	file_service_server_client_browser_monthly_stat_proto_rawDescData = file_service_server_client_browser_monthly_stat_proto_rawDesc
)

func file_service_server_client_browser_monthly_stat_proto_rawDescGZIP() []byte {
	file_service_server_client_browser_monthly_stat_proto_rawDescOnce.Do(func() {
		file_service_server_client_browser_monthly_stat_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_server_client_browser_monthly_stat_proto_rawDescData)
	})
	return file_service_server_client_browser_monthly_stat_proto_rawDescData
}

var file_service_server_client_browser_monthly_stat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_service_server_client_browser_monthly_stat_proto_goTypes = []interface{}{
	(*FindTopServerClientBrowserMonthlyStatsRequest)(nil),       // 0: pb.FindTopServerClientBrowserMonthlyStatsRequest
	(*FindTopServerClientBrowserMonthlyStatsResponse)(nil),      // 1: pb.FindTopServerClientBrowserMonthlyStatsResponse
	(*FindTopServerClientBrowserMonthlyStatsResponse_Stat)(nil), // 2: pb.FindTopServerClientBrowserMonthlyStatsResponse.Stat
	(*ClientBrowser)(nil), // 3: pb.ClientBrowser
}
var file_service_server_client_browser_monthly_stat_proto_depIdxs = []int32{
	2, // 0: pb.FindTopServerClientBrowserMonthlyStatsResponse.stats:type_name -> pb.FindTopServerClientBrowserMonthlyStatsResponse.Stat
	3, // 1: pb.FindTopServerClientBrowserMonthlyStatsResponse.Stat.clientBrowser:type_name -> pb.ClientBrowser
	0, // 2: pb.ServerClientBrowserMonthlyStatService.findTopServerClientBrowserMonthlyStats:input_type -> pb.FindTopServerClientBrowserMonthlyStatsRequest
	1, // 3: pb.ServerClientBrowserMonthlyStatService.findTopServerClientBrowserMonthlyStats:output_type -> pb.FindTopServerClientBrowserMonthlyStatsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_service_server_client_browser_monthly_stat_proto_init() }
func file_service_server_client_browser_monthly_stat_proto_init() {
	if File_service_server_client_browser_monthly_stat_proto != nil {
		return
	}
	file_models_model_client_browser_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_server_client_browser_monthly_stat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindTopServerClientBrowserMonthlyStatsRequest); i {
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
		file_service_server_client_browser_monthly_stat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindTopServerClientBrowserMonthlyStatsResponse); i {
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
		file_service_server_client_browser_monthly_stat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindTopServerClientBrowserMonthlyStatsResponse_Stat); i {
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
			RawDescriptor: file_service_server_client_browser_monthly_stat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_server_client_browser_monthly_stat_proto_goTypes,
		DependencyIndexes: file_service_server_client_browser_monthly_stat_proto_depIdxs,
		MessageInfos:      file_service_server_client_browser_monthly_stat_proto_msgTypes,
	}.Build()
	File_service_server_client_browser_monthly_stat_proto = out.File
	file_service_server_client_browser_monthly_stat_proto_rawDesc = nil
	file_service_server_client_browser_monthly_stat_proto_goTypes = nil
	file_service_server_client_browser_monthly_stat_proto_depIdxs = nil
}
