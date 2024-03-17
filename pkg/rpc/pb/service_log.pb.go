// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: service_log.proto

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

// 创建日志
type CreateLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level               string `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`                             // 级别：info, warn, error
	Description         string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`                 // 描述
	Action              string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`                           // 可选项，发生操作所在的页面URL
	Ip                  string `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`                                   // 可选项，操作用户IP
	LangMessageCode     string `protobuf:"bytes,5,opt,name=langMessageCode,proto3" json:"langMessageCode,omitempty"`         // 多语言消息
	LangMessageArgsJSON []byte `protobuf:"bytes,6,opt,name=langMessageArgsJSON,proto3" json:"langMessageArgsJSON,omitempty"` // 多语言消息中的参数项，格式：[arg1, arg2, ...]
}

func (x *CreateLogRequest) Reset() {
	*x = CreateLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLogRequest) ProtoMessage() {}

func (x *CreateLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLogRequest.ProtoReflect.Descriptor instead.
func (*CreateLogRequest) Descriptor() ([]byte, []int) {
	return file_service_log_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLogRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *CreateLogRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateLogRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *CreateLogRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *CreateLogRequest) GetLangMessageCode() string {
	if x != nil {
		return x.LangMessageCode
	}
	return ""
}

func (x *CreateLogRequest) GetLangMessageArgsJSON() []byte {
	if x != nil {
		return x.LangMessageArgsJSON
	}
	return nil
}

type CreateLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateLogResponse) Reset() {
	*x = CreateLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLogResponse) ProtoMessage() {}

func (x *CreateLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLogResponse.ProtoReflect.Descriptor instead.
func (*CreateLogResponse) Descriptor() ([]byte, []int) {
	return file_service_log_proto_rawDescGZIP(), []int{1}
}

// 计算日志数量
type CountLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DayFrom  string `protobuf:"bytes,1,opt,name=dayFrom,proto3" json:"dayFrom,omitempty"`   // 可选项，开始日期
	DayTo    string `protobuf:"bytes,2,opt,name=dayTo,proto3" json:"dayTo,omitempty"`       // 可选项，结束日期
	Keyword  string `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`   // 可选项，关键词
	UserType string `protobuf:"bytes,4,opt,name=userType,proto3" json:"userType,omitempty"` // 可选项，用户类型：admin|user；用户端固定为user
	Level    string `protobuf:"bytes,5,opt,name=level,proto3" json:"level,omitempty"`       // 可选项，错误级别：info, warn, error
}

func (x *CountLogRequest) Reset() {
	*x = CountLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountLogRequest) ProtoMessage() {}

func (x *CountLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountLogRequest.ProtoReflect.Descriptor instead.
func (*CountLogRequest) Descriptor() ([]byte, []int) {
	return file_service_log_proto_rawDescGZIP(), []int{2}
}

func (x *CountLogRequest) GetDayFrom() string {
	if x != nil {
		return x.DayFrom
	}
	return ""
}

func (x *CountLogRequest) GetDayTo() string {
	if x != nil {
		return x.DayTo
	}
	return ""
}

func (x *CountLogRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *CountLogRequest) GetUserType() string {
	if x != nil {
		return x.UserType
	}
	return ""
}

func (x *CountLogRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

// 列出单页日志
type ListLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset   int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`    // 读取位置，从0开始
	Size     int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`        // 读取数量
	DayFrom  string `protobuf:"bytes,3,opt,name=dayFrom,proto3" json:"dayFrom,omitempty"`   // 可选项，开始日期
	DayTo    string `protobuf:"bytes,4,opt,name=dayTo,proto3" json:"dayTo,omitempty"`       // 可选项，结束日期
	Keyword  string `protobuf:"bytes,5,opt,name=keyword,proto3" json:"keyword,omitempty"`   // 可选项，关键词
	UserType string `protobuf:"bytes,6,opt,name=userType,proto3" json:"userType,omitempty"` // 可选项，用户类型：admin|user；用户端固定为user
	Level    string `protobuf:"bytes,7,opt,name=level,proto3" json:"level,omitempty"`       // 可选项，错误级别：info, warn, error
}

func (x *ListLogsRequest) Reset() {
	*x = ListLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLogsRequest) ProtoMessage() {}

func (x *ListLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLogsRequest.ProtoReflect.Descriptor instead.
func (*ListLogsRequest) Descriptor() ([]byte, []int) {
	return file_service_log_proto_rawDescGZIP(), []int{3}
}

func (x *ListLogsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListLogsRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListLogsRequest) GetDayFrom() string {
	if x != nil {
		return x.DayFrom
	}
	return ""
}

func (x *ListLogsRequest) GetDayTo() string {
	if x != nil {
		return x.DayTo
	}
	return ""
}

func (x *ListLogsRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *ListLogsRequest) GetUserType() string {
	if x != nil {
		return x.UserType
	}
	return ""
}

func (x *ListLogsRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

type ListLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs []*Log `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *ListLogsResponse) Reset() {
	*x = ListLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLogsResponse) ProtoMessage() {}

func (x *ListLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLogsResponse.ProtoReflect.Descriptor instead.
func (*ListLogsResponse) Descriptor() ([]byte, []int) {
	return file_service_log_proto_rawDescGZIP(), []int{4}
}

func (x *ListLogsResponse) GetLogs() []*Log {
	if x != nil {
		return x.Logs
	}
	return nil
}

// 删除单条
type DeleteLogPermanentlyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogId int64 `protobuf:"varint,1,opt,name=logId,proto3" json:"logId,omitempty"`
}

func (x *DeleteLogPermanentlyRequest) Reset() {
	*x = DeleteLogPermanentlyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_log_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLogPermanentlyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLogPermanentlyRequest) ProtoMessage() {}

func (x *DeleteLogPermanentlyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_log_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLogPermanentlyRequest.ProtoReflect.Descriptor instead.
func (*DeleteLogPermanentlyRequest) Descriptor() ([]byte, []int) {
	return file_service_log_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteLogPermanentlyRequest) GetLogId() int64 {
	if x != nil {
		return x.LogId
	}
	return 0
}

// 批量删除
type DeleteLogsPermanentlyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogIds []int64 `protobuf:"varint,1,rep,packed,name=logIds,proto3" json:"logIds,omitempty"`
}

func (x *DeleteLogsPermanentlyRequest) Reset() {
	*x = DeleteLogsPermanentlyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_log_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLogsPermanentlyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLogsPermanentlyRequest) ProtoMessage() {}

func (x *DeleteLogsPermanentlyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_log_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLogsPermanentlyRequest.ProtoReflect.Descriptor instead.
func (*DeleteLogsPermanentlyRequest) Descriptor() ([]byte, []int) {
	return file_service_log_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteLogsPermanentlyRequest) GetLogIds() []int64 {
	if x != nil {
		return x.LogIds
	}
	return nil
}

// 清理
type CleanLogsPermanentlyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Days     int32 `protobuf:"varint,1,opt,name=days,proto3" json:"days,omitempty"`
	ClearAll bool  `protobuf:"varint,2,opt,name=clearAll,proto3" json:"clearAll,omitempty"`
}

func (x *CleanLogsPermanentlyRequest) Reset() {
	*x = CleanLogsPermanentlyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_log_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CleanLogsPermanentlyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CleanLogsPermanentlyRequest) ProtoMessage() {}

func (x *CleanLogsPermanentlyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_log_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CleanLogsPermanentlyRequest.ProtoReflect.Descriptor instead.
func (*CleanLogsPermanentlyRequest) Descriptor() ([]byte, []int) {
	return file_service_log_proto_rawDescGZIP(), []int{7}
}

func (x *CleanLogsPermanentlyRequest) GetDays() int32 {
	if x != nil {
		return x.Days
	}
	return 0
}

func (x *CleanLogsPermanentlyRequest) GetClearAll() bool {
	if x != nil {
		return x.ClearAll
	}
	return false
}

// 计算日志容量大小
type SumLogsSizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SumLogsSizeRequest) Reset() {
	*x = SumLogsSizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_log_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumLogsSizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumLogsSizeRequest) ProtoMessage() {}

func (x *SumLogsSizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_log_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumLogsSizeRequest.ProtoReflect.Descriptor instead.
func (*SumLogsSizeRequest) Descriptor() ([]byte, []int) {
	return file_service_log_proto_rawDescGZIP(), []int{8}
}

type SumLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SizeBytes int64 `protobuf:"varint,1,opt,name=sizeBytes,proto3" json:"sizeBytes,omitempty"`
}

func (x *SumLogsResponse) Reset() {
	*x = SumLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_log_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumLogsResponse) ProtoMessage() {}

func (x *SumLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_log_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumLogsResponse.ProtoReflect.Descriptor instead.
func (*SumLogsResponse) Descriptor() ([]byte, []int) {
	return file_service_log_proto_rawDescGZIP(), []int{9}
}

func (x *SumLogsResponse) GetSizeBytes() int64 {
	if x != nil {
		return x.SizeBytes
	}
	return 0
}

var File_service_log_proto protoreflect.FileDescriptor

var file_service_log_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x16, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x01, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12,
	0x28, 0x0a, 0x0f, 0x6c, 0x61, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x61, 0x6e, 0x67, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x6c, 0x61, 0x6e,
	0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x41, 0x72, 0x67, 0x73, 0x4a, 0x53, 0x4f, 0x4e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x6c, 0x61, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x41, 0x72, 0x67, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x22, 0x13, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x8d, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x61, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x61, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x14,
	0x0a, 0x05, 0x64, 0x61, 0x79, 0x54, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64,
	0x61, 0x79, 0x54, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x22, 0xb9, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x64, 0x61, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x64, 0x61, 0x79, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x61,
	0x79, 0x54, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x61, 0x79, 0x54, 0x6f,
	0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x2f, 0x0a, 0x10,
	0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x33, 0x0a,
	0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x50, 0x65, 0x72, 0x6d, 0x61, 0x6e,
	0x65, 0x6e, 0x74, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x6f, 0x67,
	0x49, 0x64, 0x22, 0x36, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73,
	0x50, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x73, 0x22, 0x4d, 0x0a, 0x1b, 0x43, 0x6c,
	0x65, 0x61, 0x6e, 0x4c, 0x6f, 0x67, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74,
	0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x79,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x79, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x41, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x63, 0x6c, 0x65, 0x61, 0x72, 0x41, 0x6c, 0x6c, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x75, 0x6d,
	0x4c, 0x6f, 0x67, 0x73, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x2f, 0x0a, 0x0f, 0x53, 0x75, 0x6d, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x32, 0xce, 0x03, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x14, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x50, 0x43, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x13, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x14, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4c, 0x6f, 0x67, 0x50, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x6c, 0x79,
	0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x50,
	0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x49, 0x0a, 0x15, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x50,
	0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x61, 0x6e,
	0x65, 0x6e, 0x74, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x47, 0x0a, 0x14,
	0x63, 0x6c, 0x65, 0x61, 0x6e, 0x4c, 0x6f, 0x67, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65,
	0x6e, 0x74, 0x6c, 0x79, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x4c,
	0x6f, 0x67, 0x73, 0x50, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x3a, 0x0a, 0x0b, 0x73, 0x75, 0x6d, 0x4c, 0x6f, 0x67, 0x73,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x75, 0x6d, 0x4c, 0x6f, 0x67,
	0x73, 0x53, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x75, 0x6d, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_service_log_proto_rawDescOnce sync.Once
	file_service_log_proto_rawDescData = file_service_log_proto_rawDesc
)

func file_service_log_proto_rawDescGZIP() []byte {
	file_service_log_proto_rawDescOnce.Do(func() {
		file_service_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_log_proto_rawDescData)
	})
	return file_service_log_proto_rawDescData
}

var file_service_log_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_service_log_proto_goTypes = []interface{}{
	(*CreateLogRequest)(nil),             // 0: pb.CreateLogRequest
	(*CreateLogResponse)(nil),            // 1: pb.CreateLogResponse
	(*CountLogRequest)(nil),              // 2: pb.CountLogRequest
	(*ListLogsRequest)(nil),              // 3: pb.ListLogsRequest
	(*ListLogsResponse)(nil),             // 4: pb.ListLogsResponse
	(*DeleteLogPermanentlyRequest)(nil),  // 5: pb.DeleteLogPermanentlyRequest
	(*DeleteLogsPermanentlyRequest)(nil), // 6: pb.DeleteLogsPermanentlyRequest
	(*CleanLogsPermanentlyRequest)(nil),  // 7: pb.CleanLogsPermanentlyRequest
	(*SumLogsSizeRequest)(nil),           // 8: pb.SumLogsSizeRequest
	(*SumLogsResponse)(nil),              // 9: pb.SumLogsResponse
	(*Log)(nil),                          // 10: pb.Log
	(*RPCCountResponse)(nil),             // 11: pb.RPCCountResponse
	(*RPCSuccess)(nil),                   // 12: pb.RPCSuccess
}
var file_service_log_proto_depIdxs = []int32{
	10, // 0: pb.ListLogsResponse.logs:type_name -> pb.Log
	0,  // 1: pb.LogService.createLog:input_type -> pb.CreateLogRequest
	2,  // 2: pb.LogService.countLogs:input_type -> pb.CountLogRequest
	3,  // 3: pb.LogService.listLogs:input_type -> pb.ListLogsRequest
	5,  // 4: pb.LogService.deleteLogPermanently:input_type -> pb.DeleteLogPermanentlyRequest
	6,  // 5: pb.LogService.deleteLogsPermanently:input_type -> pb.DeleteLogsPermanentlyRequest
	7,  // 6: pb.LogService.cleanLogsPermanently:input_type -> pb.CleanLogsPermanentlyRequest
	8,  // 7: pb.LogService.sumLogsSize:input_type -> pb.SumLogsSizeRequest
	1,  // 8: pb.LogService.createLog:output_type -> pb.CreateLogResponse
	11, // 9: pb.LogService.countLogs:output_type -> pb.RPCCountResponse
	4,  // 10: pb.LogService.listLogs:output_type -> pb.ListLogsResponse
	12, // 11: pb.LogService.deleteLogPermanently:output_type -> pb.RPCSuccess
	12, // 12: pb.LogService.deleteLogsPermanently:output_type -> pb.RPCSuccess
	12, // 13: pb.LogService.cleanLogsPermanently:output_type -> pb.RPCSuccess
	9,  // 14: pb.LogService.sumLogsSize:output_type -> pb.SumLogsResponse
	8,  // [8:15] is the sub-list for method output_type
	1,  // [1:8] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_service_log_proto_init() }
func file_service_log_proto_init() {
	if File_service_log_proto != nil {
		return
	}
	file_models_rpc_messages_proto_init()
	file_models_model_log_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLogRequest); i {
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
		file_service_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLogResponse); i {
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
		file_service_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountLogRequest); i {
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
		file_service_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLogsRequest); i {
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
		file_service_log_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLogsResponse); i {
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
		file_service_log_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLogPermanentlyRequest); i {
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
		file_service_log_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLogsPermanentlyRequest); i {
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
		file_service_log_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CleanLogsPermanentlyRequest); i {
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
		file_service_log_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumLogsSizeRequest); i {
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
		file_service_log_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumLogsResponse); i {
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
			RawDescriptor: file_service_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_log_proto_goTypes,
		DependencyIndexes: file_service_log_proto_depIdxs,
		MessageInfos:      file_service_log_proto_msgTypes,
	}.Build()
	File_service_log_proto = out.File
	file_service_log_proto_rawDesc = nil
	file_service_log_proto_goTypes = nil
	file_service_log_proto_depIdxs = nil
}
