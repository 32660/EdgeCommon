// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: service_user_traffic_package.proto

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

// 创建用户流量包
type CreateUserTrafficPackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId                 int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	TrafficPackageId       int64 `protobuf:"varint,2,opt,name=trafficPackageId,proto3" json:"trafficPackageId,omitempty"`
	NodeRegionId           int64 `protobuf:"varint,3,opt,name=nodeRegionId,proto3" json:"nodeRegionId,omitempty"`
	TrafficPackagePeriodId int64 `protobuf:"varint,4,opt,name=trafficPackagePeriodId,proto3" json:"trafficPackagePeriodId,omitempty"`
	Count                  int32 `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CreateUserTrafficPackageRequest) Reset() {
	*x = CreateUserTrafficPackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_traffic_package_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserTrafficPackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserTrafficPackageRequest) ProtoMessage() {}

func (x *CreateUserTrafficPackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_traffic_package_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserTrafficPackageRequest.ProtoReflect.Descriptor instead.
func (*CreateUserTrafficPackageRequest) Descriptor() ([]byte, []int) {
	return file_service_user_traffic_package_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserTrafficPackageRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateUserTrafficPackageRequest) GetTrafficPackageId() int64 {
	if x != nil {
		return x.TrafficPackageId
	}
	return 0
}

func (x *CreateUserTrafficPackageRequest) GetNodeRegionId() int64 {
	if x != nil {
		return x.NodeRegionId
	}
	return 0
}

func (x *CreateUserTrafficPackageRequest) GetTrafficPackagePeriodId() int64 {
	if x != nil {
		return x.TrafficPackagePeriodId
	}
	return 0
}

func (x *CreateUserTrafficPackageRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CreateUserTrafficPackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTrafficPackageIds []int64 `protobuf:"varint,1,rep,packed,name=userTrafficPackageIds,proto3" json:"userTrafficPackageIds,omitempty"`
}

func (x *CreateUserTrafficPackageResponse) Reset() {
	*x = CreateUserTrafficPackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_traffic_package_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserTrafficPackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserTrafficPackageResponse) ProtoMessage() {}

func (x *CreateUserTrafficPackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_traffic_package_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserTrafficPackageResponse.ProtoReflect.Descriptor instead.
func (*CreateUserTrafficPackageResponse) Descriptor() ([]byte, []int) {
	return file_service_user_traffic_package_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserTrafficPackageResponse) GetUserTrafficPackageIds() []int64 {
	if x != nil {
		return x.UserTrafficPackageIds
	}
	return nil
}

// 购买用户流量包
type BuyUserTrafficPackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId                 int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	TrafficPackageId       int64 `protobuf:"varint,2,opt,name=trafficPackageId,proto3" json:"trafficPackageId,omitempty"`
	NodeRegionId           int64 `protobuf:"varint,3,opt,name=nodeRegionId,proto3" json:"nodeRegionId,omitempty"`
	TrafficPackagePeriodId int64 `protobuf:"varint,4,opt,name=trafficPackagePeriodId,proto3" json:"trafficPackagePeriodId,omitempty"`
	Count                  int32 `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *BuyUserTrafficPackageRequest) Reset() {
	*x = BuyUserTrafficPackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_traffic_package_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyUserTrafficPackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyUserTrafficPackageRequest) ProtoMessage() {}

func (x *BuyUserTrafficPackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_traffic_package_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyUserTrafficPackageRequest.ProtoReflect.Descriptor instead.
func (*BuyUserTrafficPackageRequest) Descriptor() ([]byte, []int) {
	return file_service_user_traffic_package_proto_rawDescGZIP(), []int{2}
}

func (x *BuyUserTrafficPackageRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *BuyUserTrafficPackageRequest) GetTrafficPackageId() int64 {
	if x != nil {
		return x.TrafficPackageId
	}
	return 0
}

func (x *BuyUserTrafficPackageRequest) GetNodeRegionId() int64 {
	if x != nil {
		return x.NodeRegionId
	}
	return 0
}

func (x *BuyUserTrafficPackageRequest) GetTrafficPackagePeriodId() int64 {
	if x != nil {
		return x.TrafficPackagePeriodId
	}
	return 0
}

func (x *BuyUserTrafficPackageRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type BuyUserTrafficPackageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTrafficPackageIds []int64 `protobuf:"varint,1,rep,packed,name=userTrafficPackageIds,proto3" json:"userTrafficPackageIds,omitempty"`
}

func (x *BuyUserTrafficPackageResponse) Reset() {
	*x = BuyUserTrafficPackageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_traffic_package_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyUserTrafficPackageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyUserTrafficPackageResponse) ProtoMessage() {}

func (x *BuyUserTrafficPackageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_traffic_package_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyUserTrafficPackageResponse.ProtoReflect.Descriptor instead.
func (*BuyUserTrafficPackageResponse) Descriptor() ([]byte, []int) {
	return file_service_user_traffic_package_proto_rawDescGZIP(), []int{3}
}

func (x *BuyUserTrafficPackageResponse) GetUserTrafficPackageIds() []int64 {
	if x != nil {
		return x.UserTrafficPackageIds
	}
	return nil
}

// 查询当前流量包数量
type CountUserTrafficPackagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrafficPackageId       int64  `protobuf:"varint,1,opt,name=trafficPackageId,proto3" json:"trafficPackageId,omitempty"`
	UserId                 int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	NodeRegionId           int64  `protobuf:"varint,3,opt,name=nodeRegionId,proto3" json:"nodeRegionId,omitempty"`
	TrafficPackagePeriodId int64  `protobuf:"varint,4,opt,name=trafficPackagePeriodId,proto3" json:"trafficPackagePeriodId,omitempty"`
	ExpiresDay             string `protobuf:"bytes,5,opt,name=expiresDay,proto3" json:"expiresDay,omitempty"`
	AvailableOnly          bool   `protobuf:"varint,6,opt,name=availableOnly,proto3" json:"availableOnly,omitempty"` // 是否只查询有效的流量包
}

func (x *CountUserTrafficPackagesRequest) Reset() {
	*x = CountUserTrafficPackagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_traffic_package_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountUserTrafficPackagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountUserTrafficPackagesRequest) ProtoMessage() {}

func (x *CountUserTrafficPackagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_traffic_package_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountUserTrafficPackagesRequest.ProtoReflect.Descriptor instead.
func (*CountUserTrafficPackagesRequest) Descriptor() ([]byte, []int) {
	return file_service_user_traffic_package_proto_rawDescGZIP(), []int{4}
}

func (x *CountUserTrafficPackagesRequest) GetTrafficPackageId() int64 {
	if x != nil {
		return x.TrafficPackageId
	}
	return 0
}

func (x *CountUserTrafficPackagesRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CountUserTrafficPackagesRequest) GetNodeRegionId() int64 {
	if x != nil {
		return x.NodeRegionId
	}
	return 0
}

func (x *CountUserTrafficPackagesRequest) GetTrafficPackagePeriodId() int64 {
	if x != nil {
		return x.TrafficPackagePeriodId
	}
	return 0
}

func (x *CountUserTrafficPackagesRequest) GetExpiresDay() string {
	if x != nil {
		return x.ExpiresDay
	}
	return ""
}

func (x *CountUserTrafficPackagesRequest) GetAvailableOnly() bool {
	if x != nil {
		return x.AvailableOnly
	}
	return false
}

// 列出单页流量包
type ListUserTrafficPackagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrafficPackageId       int64  `protobuf:"varint,1,opt,name=trafficPackageId,proto3" json:"trafficPackageId,omitempty"`
	UserId                 int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	NodeRegionId           int64  `protobuf:"varint,3,opt,name=nodeRegionId,proto3" json:"nodeRegionId,omitempty"`
	TrafficPackagePeriodId int64  `protobuf:"varint,4,opt,name=trafficPackagePeriodId,proto3" json:"trafficPackagePeriodId,omitempty"`
	ExpiresDay             string `protobuf:"bytes,5,opt,name=expiresDay,proto3" json:"expiresDay,omitempty"`
	AvailableOnly          bool   `protobuf:"varint,6,opt,name=availableOnly,proto3" json:"availableOnly,omitempty"` // 是否只查询有效的流量包
	Offset                 int64  `protobuf:"varint,7,opt,name=offset,proto3" json:"offset,omitempty"`
	Size                   int64  `protobuf:"varint,8,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListUserTrafficPackagesRequest) Reset() {
	*x = ListUserTrafficPackagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_traffic_package_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserTrafficPackagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserTrafficPackagesRequest) ProtoMessage() {}

func (x *ListUserTrafficPackagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_traffic_package_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserTrafficPackagesRequest.ProtoReflect.Descriptor instead.
func (*ListUserTrafficPackagesRequest) Descriptor() ([]byte, []int) {
	return file_service_user_traffic_package_proto_rawDescGZIP(), []int{5}
}

func (x *ListUserTrafficPackagesRequest) GetTrafficPackageId() int64 {
	if x != nil {
		return x.TrafficPackageId
	}
	return 0
}

func (x *ListUserTrafficPackagesRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ListUserTrafficPackagesRequest) GetNodeRegionId() int64 {
	if x != nil {
		return x.NodeRegionId
	}
	return 0
}

func (x *ListUserTrafficPackagesRequest) GetTrafficPackagePeriodId() int64 {
	if x != nil {
		return x.TrafficPackagePeriodId
	}
	return 0
}

func (x *ListUserTrafficPackagesRequest) GetExpiresDay() string {
	if x != nil {
		return x.ExpiresDay
	}
	return ""
}

func (x *ListUserTrafficPackagesRequest) GetAvailableOnly() bool {
	if x != nil {
		return x.AvailableOnly
	}
	return false
}

func (x *ListUserTrafficPackagesRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListUserTrafficPackagesRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListUserTrafficPackagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTrafficPackages []*UserTrafficPackage `protobuf:"bytes,1,rep,name=userTrafficPackages,proto3" json:"userTrafficPackages,omitempty"`
}

func (x *ListUserTrafficPackagesResponse) Reset() {
	*x = ListUserTrafficPackagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_traffic_package_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserTrafficPackagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserTrafficPackagesResponse) ProtoMessage() {}

func (x *ListUserTrafficPackagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_traffic_package_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserTrafficPackagesResponse.ProtoReflect.Descriptor instead.
func (*ListUserTrafficPackagesResponse) Descriptor() ([]byte, []int) {
	return file_service_user_traffic_package_proto_rawDescGZIP(), []int{6}
}

func (x *ListUserTrafficPackagesResponse) GetUserTrafficPackages() []*UserTrafficPackage {
	if x != nil {
		return x.UserTrafficPackages
	}
	return nil
}

// 删除流量包
type DeleteUserTrafficPackageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserTrafficPackageId int64 `protobuf:"varint,1,opt,name=userTrafficPackageId,proto3" json:"userTrafficPackageId,omitempty"`
}

func (x *DeleteUserTrafficPackageRequest) Reset() {
	*x = DeleteUserTrafficPackageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_traffic_package_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserTrafficPackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserTrafficPackageRequest) ProtoMessage() {}

func (x *DeleteUserTrafficPackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_traffic_package_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserTrafficPackageRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserTrafficPackageRequest) Descriptor() ([]byte, []int) {
	return file_service_user_traffic_package_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteUserTrafficPackageRequest) GetUserTrafficPackageId() int64 {
	if x != nil {
		return x.UserTrafficPackageId
	}
	return 0
}

var File_service_user_traffic_package_proto protoreflect.FileDescriptor

var file_service_user_traffic_package_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x27, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x01, 0x0a,
	0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x16, 0x74, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x58, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x75, 0x73,
	0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x15, 0x75, 0x73, 0x65, 0x72, 0x54,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x73,
	0x22, 0xd4, 0x01, 0x0a, 0x1c, 0x42, 0x75, 0x79, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6e, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x16, 0x74, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x74, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x55, 0x0a, 0x1d, 0x42, 0x75, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x75, 0x73, 0x65, 0x72,
	0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x15, 0x75, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x73, 0x22, 0x87,
	0x02, 0x0a, 0x1f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x74, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x16, 0x74, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x74, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x44, 0x61, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x44,
	0x61, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4f,
	0x6e, 0x6c, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x22, 0xb2, 0x02, 0x0a, 0x1e, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x74,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x16, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x16, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x44, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x44, 0x61, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x6e, 0x6c,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x6b, 0x0a,
	0x1f, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x48, 0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x13, 0x75, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x22, 0x55, 0x0a, 0x1f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a,
	0x14, 0x75, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x75, 0x73, 0x65,
	0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x49,
	0x64, 0x32, 0xec, 0x03, 0x0a, 0x19, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x65, 0x0a, 0x18, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x23, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x15, 0x62, 0x75, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12,
	0x20, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x75, 0x79, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66,
	0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x75, 0x79, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x18, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x23, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x17, 0x6c,
	0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x62, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4f, 0x0a, 0x18, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x23, 0x2e, 0x70, 0x62,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_user_traffic_package_proto_rawDescOnce sync.Once
	file_service_user_traffic_package_proto_rawDescData = file_service_user_traffic_package_proto_rawDesc
)

func file_service_user_traffic_package_proto_rawDescGZIP() []byte {
	file_service_user_traffic_package_proto_rawDescOnce.Do(func() {
		file_service_user_traffic_package_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_user_traffic_package_proto_rawDescData)
	})
	return file_service_user_traffic_package_proto_rawDescData
}

var file_service_user_traffic_package_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_service_user_traffic_package_proto_goTypes = []interface{}{
	(*CreateUserTrafficPackageRequest)(nil),  // 0: pb.CreateUserTrafficPackageRequest
	(*CreateUserTrafficPackageResponse)(nil), // 1: pb.CreateUserTrafficPackageResponse
	(*BuyUserTrafficPackageRequest)(nil),     // 2: pb.BuyUserTrafficPackageRequest
	(*BuyUserTrafficPackageResponse)(nil),    // 3: pb.BuyUserTrafficPackageResponse
	(*CountUserTrafficPackagesRequest)(nil),  // 4: pb.CountUserTrafficPackagesRequest
	(*ListUserTrafficPackagesRequest)(nil),   // 5: pb.ListUserTrafficPackagesRequest
	(*ListUserTrafficPackagesResponse)(nil),  // 6: pb.ListUserTrafficPackagesResponse
	(*DeleteUserTrafficPackageRequest)(nil),  // 7: pb.DeleteUserTrafficPackageRequest
	(*UserTrafficPackage)(nil),               // 8: pb.UserTrafficPackage
	(*RPCCountResponse)(nil),                 // 9: pb.RPCCountResponse
	(*RPCSuccess)(nil),                       // 10: pb.RPCSuccess
}
var file_service_user_traffic_package_proto_depIdxs = []int32{
	8,  // 0: pb.ListUserTrafficPackagesResponse.userTrafficPackages:type_name -> pb.UserTrafficPackage
	0,  // 1: pb.UserTrafficPackageService.createUserTrafficPackage:input_type -> pb.CreateUserTrafficPackageRequest
	2,  // 2: pb.UserTrafficPackageService.buyUserTrafficPackage:input_type -> pb.BuyUserTrafficPackageRequest
	4,  // 3: pb.UserTrafficPackageService.countUserTrafficPackages:input_type -> pb.CountUserTrafficPackagesRequest
	5,  // 4: pb.UserTrafficPackageService.listUserTrafficPackages:input_type -> pb.ListUserTrafficPackagesRequest
	7,  // 5: pb.UserTrafficPackageService.deleteUserTrafficPackage:input_type -> pb.DeleteUserTrafficPackageRequest
	1,  // 6: pb.UserTrafficPackageService.createUserTrafficPackage:output_type -> pb.CreateUserTrafficPackageResponse
	3,  // 7: pb.UserTrafficPackageService.buyUserTrafficPackage:output_type -> pb.BuyUserTrafficPackageResponse
	9,  // 8: pb.UserTrafficPackageService.countUserTrafficPackages:output_type -> pb.RPCCountResponse
	6,  // 9: pb.UserTrafficPackageService.listUserTrafficPackages:output_type -> pb.ListUserTrafficPackagesResponse
	10, // 10: pb.UserTrafficPackageService.deleteUserTrafficPackage:output_type -> pb.RPCSuccess
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_service_user_traffic_package_proto_init() }
func file_service_user_traffic_package_proto_init() {
	if File_service_user_traffic_package_proto != nil {
		return
	}
	file_models_model_user_traffic_package_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_user_traffic_package_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserTrafficPackageRequest); i {
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
		file_service_user_traffic_package_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserTrafficPackageResponse); i {
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
		file_service_user_traffic_package_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyUserTrafficPackageRequest); i {
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
		file_service_user_traffic_package_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyUserTrafficPackageResponse); i {
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
		file_service_user_traffic_package_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountUserTrafficPackagesRequest); i {
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
		file_service_user_traffic_package_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserTrafficPackagesRequest); i {
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
		file_service_user_traffic_package_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserTrafficPackagesResponse); i {
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
		file_service_user_traffic_package_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserTrafficPackageRequest); i {
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
			RawDescriptor: file_service_user_traffic_package_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_user_traffic_package_proto_goTypes,
		DependencyIndexes: file_service_user_traffic_package_proto_depIdxs,
		MessageInfos:      file_service_user_traffic_package_proto_msgTypes,
	}.Build()
	File_service_user_traffic_package_proto = out.File
	file_service_user_traffic_package_proto_rawDesc = nil
	file_service_user_traffic_package_proto_goTypes = nil
	file_service_user_traffic_package_proto_depIdxs = nil
}
