// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: service_server_stat_board_chart.proto

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

// 添加图表
type EnableServerStatBoardChartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerStatBoardId int64 `protobuf:"varint,1,opt,name=serverStatBoardId,proto3" json:"serverStatBoardId,omitempty"`
	MetricChartId     int64 `protobuf:"varint,2,opt,name=metricChartId,proto3" json:"metricChartId,omitempty"`
}

func (x *EnableServerStatBoardChartRequest) Reset() {
	*x = EnableServerStatBoardChartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_stat_board_chart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnableServerStatBoardChartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableServerStatBoardChartRequest) ProtoMessage() {}

func (x *EnableServerStatBoardChartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_stat_board_chart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableServerStatBoardChartRequest.ProtoReflect.Descriptor instead.
func (*EnableServerStatBoardChartRequest) Descriptor() ([]byte, []int) {
	return file_service_server_stat_board_chart_proto_rawDescGZIP(), []int{0}
}

func (x *EnableServerStatBoardChartRequest) GetServerStatBoardId() int64 {
	if x != nil {
		return x.ServerStatBoardId
	}
	return 0
}

func (x *EnableServerStatBoardChartRequest) GetMetricChartId() int64 {
	if x != nil {
		return x.MetricChartId
	}
	return 0
}

// 取消图表
type DisableServerStatBoardChartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerStatBoardId int64 `protobuf:"varint,1,opt,name=serverStatBoardId,proto3" json:"serverStatBoardId,omitempty"`
	MetricChartId     int64 `protobuf:"varint,2,opt,name=metricChartId,proto3" json:"metricChartId,omitempty"`
}

func (x *DisableServerStatBoardChartRequest) Reset() {
	*x = DisableServerStatBoardChartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_stat_board_chart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisableServerStatBoardChartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisableServerStatBoardChartRequest) ProtoMessage() {}

func (x *DisableServerStatBoardChartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_stat_board_chart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisableServerStatBoardChartRequest.ProtoReflect.Descriptor instead.
func (*DisableServerStatBoardChartRequest) Descriptor() ([]byte, []int) {
	return file_service_server_stat_board_chart_proto_rawDescGZIP(), []int{1}
}

func (x *DisableServerStatBoardChartRequest) GetServerStatBoardId() int64 {
	if x != nil {
		return x.ServerStatBoardId
	}
	return 0
}

func (x *DisableServerStatBoardChartRequest) GetMetricChartId() int64 {
	if x != nil {
		return x.MetricChartId
	}
	return 0
}

// 读取看板中的图表
type FindAllEnabledServerStatBoardChartsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerStatBoardId int64 `protobuf:"varint,1,opt,name=serverStatBoardId,proto3" json:"serverStatBoardId,omitempty"`
}

func (x *FindAllEnabledServerStatBoardChartsRequest) Reset() {
	*x = FindAllEnabledServerStatBoardChartsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_stat_board_chart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledServerStatBoardChartsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledServerStatBoardChartsRequest) ProtoMessage() {}

func (x *FindAllEnabledServerStatBoardChartsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_stat_board_chart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledServerStatBoardChartsRequest.ProtoReflect.Descriptor instead.
func (*FindAllEnabledServerStatBoardChartsRequest) Descriptor() ([]byte, []int) {
	return file_service_server_stat_board_chart_proto_rawDescGZIP(), []int{2}
}

func (x *FindAllEnabledServerStatBoardChartsRequest) GetServerStatBoardId() int64 {
	if x != nil {
		return x.ServerStatBoardId
	}
	return 0
}

type FindAllEnabledServerStatBoardChartsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerStatBoardCharts []*ServerStatBoardChart `protobuf:"bytes,1,rep,name=serverStatBoardCharts,proto3" json:"serverStatBoardCharts,omitempty"`
}

func (x *FindAllEnabledServerStatBoardChartsResponse) Reset() {
	*x = FindAllEnabledServerStatBoardChartsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_server_stat_board_chart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledServerStatBoardChartsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledServerStatBoardChartsResponse) ProtoMessage() {}

func (x *FindAllEnabledServerStatBoardChartsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_server_stat_board_chart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledServerStatBoardChartsResponse.ProtoReflect.Descriptor instead.
func (*FindAllEnabledServerStatBoardChartsResponse) Descriptor() ([]byte, []int) {
	return file_service_server_stat_board_chart_proto_rawDescGZIP(), []int{3}
}

func (x *FindAllEnabledServerStatBoardChartsResponse) GetServerStatBoardCharts() []*ServerStatBoardChart {
	if x != nil {
		return x.ServerStatBoardCharts
	}
	return nil
}

var File_service_server_stat_board_chart_proto protoreflect.FileDescriptor

var file_service_server_stat_board_chart_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x72,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x5f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x77, 0x0a, 0x21, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x68, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43,
	0x68, 0x61, 0x72, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x43, 0x68, 0x61, 0x72, 0x74, 0x49, 0x64, 0x22, 0x78, 0x0a, 0x22, 0x44,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x42, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12,
	0x24, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x68, 0x61, 0x72, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x68,
	0x61, 0x72, 0x74, 0x49, 0x64, 0x22, 0x5a, 0x0a, 0x2a, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x49,
	0x64, 0x22, 0x7d, 0x0a, 0x2b, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4e, 0x0a, 0x15, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6f,
	0x61, 0x72, 0x64, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x15, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73,
	0x32, 0xd2, 0x02, 0x0a, 0x1b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x42,
	0x6f, 0x61, 0x72, 0x64, 0x43, 0x68, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x53, 0x0a, 0x1a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x25,
	0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x55, 0x0a, 0x1b, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x43,
	0x68, 0x61, 0x72, 0x74, 0x12, 0x26, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64,
	0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x86, 0x01, 0x0a,
	0x23, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x68,
	0x61, 0x72, 0x74, 0x73, 0x12, 0x2e, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_server_stat_board_chart_proto_rawDescOnce sync.Once
	file_service_server_stat_board_chart_proto_rawDescData = file_service_server_stat_board_chart_proto_rawDesc
)

func file_service_server_stat_board_chart_proto_rawDescGZIP() []byte {
	file_service_server_stat_board_chart_proto_rawDescOnce.Do(func() {
		file_service_server_stat_board_chart_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_server_stat_board_chart_proto_rawDescData)
	})
	return file_service_server_stat_board_chart_proto_rawDescData
}

var file_service_server_stat_board_chart_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_server_stat_board_chart_proto_goTypes = []interface{}{
	(*EnableServerStatBoardChartRequest)(nil),           // 0: pb.EnableServerStatBoardChartRequest
	(*DisableServerStatBoardChartRequest)(nil),          // 1: pb.DisableServerStatBoardChartRequest
	(*FindAllEnabledServerStatBoardChartsRequest)(nil),  // 2: pb.FindAllEnabledServerStatBoardChartsRequest
	(*FindAllEnabledServerStatBoardChartsResponse)(nil), // 3: pb.FindAllEnabledServerStatBoardChartsResponse
	(*ServerStatBoardChart)(nil),                        // 4: pb.ServerStatBoardChart
	(*RPCSuccess)(nil),                                  // 5: pb.RPCSuccess
}
var file_service_server_stat_board_chart_proto_depIdxs = []int32{
	4, // 0: pb.FindAllEnabledServerStatBoardChartsResponse.serverStatBoardCharts:type_name -> pb.ServerStatBoardChart
	0, // 1: pb.ServerStatBoardChartService.enableServerStatBoardChart:input_type -> pb.EnableServerStatBoardChartRequest
	1, // 2: pb.ServerStatBoardChartService.disableServerStatBoardChart:input_type -> pb.DisableServerStatBoardChartRequest
	2, // 3: pb.ServerStatBoardChartService.findAllEnabledServerStatBoardCharts:input_type -> pb.FindAllEnabledServerStatBoardChartsRequest
	5, // 4: pb.ServerStatBoardChartService.enableServerStatBoardChart:output_type -> pb.RPCSuccess
	5, // 5: pb.ServerStatBoardChartService.disableServerStatBoardChart:output_type -> pb.RPCSuccess
	3, // 6: pb.ServerStatBoardChartService.findAllEnabledServerStatBoardCharts:output_type -> pb.FindAllEnabledServerStatBoardChartsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_server_stat_board_chart_proto_init() }
func file_service_server_stat_board_chart_proto_init() {
	if File_service_server_stat_board_chart_proto != nil {
		return
	}
	file_models_rpc_messages_proto_init()
	file_models_model_server_stat_board_chart_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_server_stat_board_chart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnableServerStatBoardChartRequest); i {
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
		file_service_server_stat_board_chart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisableServerStatBoardChartRequest); i {
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
		file_service_server_stat_board_chart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledServerStatBoardChartsRequest); i {
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
		file_service_server_stat_board_chart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledServerStatBoardChartsResponse); i {
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
			RawDescriptor: file_service_server_stat_board_chart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_server_stat_board_chart_proto_goTypes,
		DependencyIndexes: file_service_server_stat_board_chart_proto_depIdxs,
		MessageInfos:      file_service_server_stat_board_chart_proto_msgTypes,
	}.Build()
	File_service_server_stat_board_chart_proto = out.File
	file_service_server_stat_board_chart_proto_rawDesc = nil
	file_service_server_stat_board_chart_proto_goTypes = nil
	file_service_server_stat_board_chart_proto_depIdxs = nil
}
