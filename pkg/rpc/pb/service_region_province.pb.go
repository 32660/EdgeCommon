// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: service_region_province.proto

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

// 根据国家|地区ID查找所有省份
type FindAllEnabledRegionProvincesWithCountryIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionCountryId int64 `protobuf:"varint,1,opt,name=regionCountryId,proto3" json:"regionCountryId,omitempty"` // 国家|地区ID
}

func (x *FindAllEnabledRegionProvincesWithCountryIdRequest) Reset() {
	*x = FindAllEnabledRegionProvincesWithCountryIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_province_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledRegionProvincesWithCountryIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledRegionProvincesWithCountryIdRequest) ProtoMessage() {}

func (x *FindAllEnabledRegionProvincesWithCountryIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_province_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledRegionProvincesWithCountryIdRequest.ProtoReflect.Descriptor instead.
func (*FindAllEnabledRegionProvincesWithCountryIdRequest) Descriptor() ([]byte, []int) {
	return file_service_region_province_proto_rawDescGZIP(), []int{0}
}

func (x *FindAllEnabledRegionProvincesWithCountryIdRequest) GetRegionCountryId() int64 {
	if x != nil {
		return x.RegionCountryId
	}
	return 0
}

type FindAllEnabledRegionProvincesWithCountryIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionProvinces []*RegionProvince `protobuf:"bytes,1,rep,name=regionProvinces,proto3" json:"regionProvinces,omitempty"` // 省份列表
}

func (x *FindAllEnabledRegionProvincesWithCountryIdResponse) Reset() {
	*x = FindAllEnabledRegionProvincesWithCountryIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_province_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllEnabledRegionProvincesWithCountryIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllEnabledRegionProvincesWithCountryIdResponse) ProtoMessage() {}

func (x *FindAllEnabledRegionProvincesWithCountryIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_province_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllEnabledRegionProvincesWithCountryIdResponse.ProtoReflect.Descriptor instead.
func (*FindAllEnabledRegionProvincesWithCountryIdResponse) Descriptor() ([]byte, []int) {
	return file_service_region_province_proto_rawDescGZIP(), []int{1}
}

func (x *FindAllEnabledRegionProvincesWithCountryIdResponse) GetRegionProvinces() []*RegionProvince {
	if x != nil {
		return x.RegionProvinces
	}
	return nil
}

// 查找单个省份信息
type FindEnabledRegionProvinceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionProvinceId int64 `protobuf:"varint,1,opt,name=regionProvinceId,proto3" json:"regionProvinceId,omitempty"` // 省份ID
}

func (x *FindEnabledRegionProvinceRequest) Reset() {
	*x = FindEnabledRegionProvinceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_province_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledRegionProvinceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledRegionProvinceRequest) ProtoMessage() {}

func (x *FindEnabledRegionProvinceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_province_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledRegionProvinceRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledRegionProvinceRequest) Descriptor() ([]byte, []int) {
	return file_service_region_province_proto_rawDescGZIP(), []int{2}
}

func (x *FindEnabledRegionProvinceRequest) GetRegionProvinceId() int64 {
	if x != nil {
		return x.RegionProvinceId
	}
	return 0
}

type FindEnabledRegionProvinceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionProvince *RegionProvince `protobuf:"bytes,1,opt,name=regionProvince,proto3" json:"regionProvince,omitempty"` // 省份信息
}

func (x *FindEnabledRegionProvinceResponse) Reset() {
	*x = FindEnabledRegionProvinceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_province_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledRegionProvinceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledRegionProvinceResponse) ProtoMessage() {}

func (x *FindEnabledRegionProvinceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_province_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledRegionProvinceResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledRegionProvinceResponse) Descriptor() ([]byte, []int) {
	return file_service_region_province_proto_rawDescGZIP(), []int{3}
}

func (x *FindEnabledRegionProvinceResponse) GetRegionProvince() *RegionProvince {
	if x != nil {
		return x.RegionProvince
	}
	return nil
}

// 根据国家|地区ID查找所有省份
type FindAllRegionProvincesWithRegionCountryIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionCountryId int64 `protobuf:"varint,1,opt,name=regionCountryId,proto3" json:"regionCountryId,omitempty"` // 国家|地区ID
}

func (x *FindAllRegionProvincesWithRegionCountryIdRequest) Reset() {
	*x = FindAllRegionProvincesWithRegionCountryIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_province_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllRegionProvincesWithRegionCountryIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllRegionProvincesWithRegionCountryIdRequest) ProtoMessage() {}

func (x *FindAllRegionProvincesWithRegionCountryIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_province_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllRegionProvincesWithRegionCountryIdRequest.ProtoReflect.Descriptor instead.
func (*FindAllRegionProvincesWithRegionCountryIdRequest) Descriptor() ([]byte, []int) {
	return file_service_region_province_proto_rawDescGZIP(), []int{4}
}

func (x *FindAllRegionProvincesWithRegionCountryIdRequest) GetRegionCountryId() int64 {
	if x != nil {
		return x.RegionCountryId
	}
	return 0
}

type FindAllRegionProvincesWithRegionCountryIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionProvinces []*RegionProvince `protobuf:"bytes,1,rep,name=regionProvinces,proto3" json:"regionProvinces,omitempty"` // 省份列表
}

func (x *FindAllRegionProvincesWithRegionCountryIdResponse) Reset() {
	*x = FindAllRegionProvincesWithRegionCountryIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_province_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllRegionProvincesWithRegionCountryIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllRegionProvincesWithRegionCountryIdResponse) ProtoMessage() {}

func (x *FindAllRegionProvincesWithRegionCountryIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_province_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllRegionProvincesWithRegionCountryIdResponse.ProtoReflect.Descriptor instead.
func (*FindAllRegionProvincesWithRegionCountryIdResponse) Descriptor() ([]byte, []int) {
	return file_service_region_province_proto_rawDescGZIP(), []int{5}
}

func (x *FindAllRegionProvincesWithRegionCountryIdResponse) GetRegionProvinces() []*RegionProvince {
	if x != nil {
		return x.RegionProvinces
	}
	return nil
}

// 查找所有国家|地区的所有省份
type FindAllRegionProvincesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllRegionProvincesRequest) Reset() {
	*x = FindAllRegionProvincesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_province_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllRegionProvincesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllRegionProvincesRequest) ProtoMessage() {}

func (x *FindAllRegionProvincesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_province_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllRegionProvincesRequest.ProtoReflect.Descriptor instead.
func (*FindAllRegionProvincesRequest) Descriptor() ([]byte, []int) {
	return file_service_region_province_proto_rawDescGZIP(), []int{6}
}

type FindAllRegionProvincesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionProvinces []*RegionProvince `protobuf:"bytes,1,rep,name=regionProvinces,proto3" json:"regionProvinces,omitempty"` // 省份列表
}

func (x *FindAllRegionProvincesResponse) Reset() {
	*x = FindAllRegionProvincesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_province_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllRegionProvincesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllRegionProvincesResponse) ProtoMessage() {}

func (x *FindAllRegionProvincesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_province_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllRegionProvincesResponse.ProtoReflect.Descriptor instead.
func (*FindAllRegionProvincesResponse) Descriptor() ([]byte, []int) {
	return file_service_region_province_proto_rawDescGZIP(), []int{7}
}

func (x *FindAllRegionProvincesResponse) GetRegionProvinces() []*RegionProvince {
	if x != nil {
		return x.RegionProvinces
	}
	return nil
}

// 查找单个省份信息
type FindRegionProvinceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionProvinceId int64 `protobuf:"varint,1,opt,name=regionProvinceId,proto3" json:"regionProvinceId,omitempty"` // 省份ID
}

func (x *FindRegionProvinceRequest) Reset() {
	*x = FindRegionProvinceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_province_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRegionProvinceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRegionProvinceRequest) ProtoMessage() {}

func (x *FindRegionProvinceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_province_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRegionProvinceRequest.ProtoReflect.Descriptor instead.
func (*FindRegionProvinceRequest) Descriptor() ([]byte, []int) {
	return file_service_region_province_proto_rawDescGZIP(), []int{8}
}

func (x *FindRegionProvinceRequest) GetRegionProvinceId() int64 {
	if x != nil {
		return x.RegionProvinceId
	}
	return 0
}

type FindRegionProvinceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionProvince *RegionProvince `protobuf:"bytes,1,opt,name=regionProvince,proto3" json:"regionProvince,omitempty"` // 省份信息
}

func (x *FindRegionProvinceResponse) Reset() {
	*x = FindRegionProvinceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_province_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRegionProvinceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRegionProvinceResponse) ProtoMessage() {}

func (x *FindRegionProvinceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_province_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRegionProvinceResponse.ProtoReflect.Descriptor instead.
func (*FindRegionProvinceResponse) Descriptor() ([]byte, []int) {
	return file_service_region_province_proto_rawDescGZIP(), []int{9}
}

func (x *FindRegionProvinceResponse) GetRegionProvince() *RegionProvince {
	if x != nil {
		return x.RegionProvince
	}
	return nil
}

// 修改省份定制信息
type UpdateRegionProvinceCustomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionProvinceId int64    `protobuf:"varint,1,opt,name=regionProvinceId,proto3" json:"regionProvinceId,omitempty"` // 省份ID
	CustomName       string   `protobuf:"bytes,2,opt,name=customName,proto3" json:"customName,omitempty"`              // 自定义名称
	CustomCodes      []string `protobuf:"bytes,3,rep,name=customCodes,proto3" json:"customCodes,omitempty"`            // 自定义代号
}

func (x *UpdateRegionProvinceCustomRequest) Reset() {
	*x = UpdateRegionProvinceCustomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_region_province_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRegionProvinceCustomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRegionProvinceCustomRequest) ProtoMessage() {}

func (x *UpdateRegionProvinceCustomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_region_province_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRegionProvinceCustomRequest.ProtoReflect.Descriptor instead.
func (*UpdateRegionProvinceCustomRequest) Descriptor() ([]byte, []int) {
	return file_service_region_province_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateRegionProvinceCustomRequest) GetRegionProvinceId() int64 {
	if x != nil {
		return x.RegionProvinceId
	}
	return 0
}

func (x *UpdateRegionProvinceCustomRequest) GetCustomName() string {
	if x != nil {
		return x.CustomName
	}
	return ""
}

func (x *UpdateRegionProvinceCustomRequest) GetCustomCodes() []string {
	if x != nil {
		return x.CustomCodes
	}
	return nil
}

var File_service_region_province_proto protoreflect.FileDescriptor

var file_service_region_province_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x22, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x31, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49,
	0x64, 0x22, 0x72, 0x0a, 0x32, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x4e, 0x0a, 0x20, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x10, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x22, 0x5f, 0x0a, 0x21, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x22, 0x5c, 0x0a, 0x30, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73,
	0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x49, 0x64, 0x22, 0x71, 0x0a, 0x31, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x57, 0x69,
	0x74, 0x68, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x1f, 0x0a, 0x1d, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5e, 0x0a, 0x1e, 0x46, 0x69, 0x6e, 0x64,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0f, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x22, 0x47, 0x0a, 0x19, 0x46, 0x69, 0x6e, 0x64,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x10, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49,
	0x64, 0x22, 0x58, 0x0a, 0x1a, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3a, 0x0a, 0x0e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52, 0x0e, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x21,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x32,
	0xcf, 0x05, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa0, 0x01, 0x0a, 0x2a, 0x66, 0x69,
	0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x35, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69,
	0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x36, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x88, 0x02, 0x01, 0x12, 0x6d, 0x0a, 0x19,
	0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x24, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x88, 0x02, 0x01, 0x12, 0x98, 0x01, 0x0a, 0x29,
	0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x34, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x35, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x57, 0x69, 0x74, 0x68, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x16, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73,
	0x12, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x12, 0x66, 0x69, 0x6e, 0x64, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x1d, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x1a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x12, 0x25, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x6e, 0x63, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_service_region_province_proto_rawDescOnce sync.Once
	file_service_region_province_proto_rawDescData = file_service_region_province_proto_rawDesc
)

func file_service_region_province_proto_rawDescGZIP() []byte {
	file_service_region_province_proto_rawDescOnce.Do(func() {
		file_service_region_province_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_region_province_proto_rawDescData)
	})
	return file_service_region_province_proto_rawDescData
}

var file_service_region_province_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_service_region_province_proto_goTypes = []interface{}{
	(*FindAllEnabledRegionProvincesWithCountryIdRequest)(nil),  // 0: pb.FindAllEnabledRegionProvincesWithCountryIdRequest
	(*FindAllEnabledRegionProvincesWithCountryIdResponse)(nil), // 1: pb.FindAllEnabledRegionProvincesWithCountryIdResponse
	(*FindEnabledRegionProvinceRequest)(nil),                   // 2: pb.FindEnabledRegionProvinceRequest
	(*FindEnabledRegionProvinceResponse)(nil),                  // 3: pb.FindEnabledRegionProvinceResponse
	(*FindAllRegionProvincesWithRegionCountryIdRequest)(nil),   // 4: pb.FindAllRegionProvincesWithRegionCountryIdRequest
	(*FindAllRegionProvincesWithRegionCountryIdResponse)(nil),  // 5: pb.FindAllRegionProvincesWithRegionCountryIdResponse
	(*FindAllRegionProvincesRequest)(nil),                      // 6: pb.FindAllRegionProvincesRequest
	(*FindAllRegionProvincesResponse)(nil),                     // 7: pb.FindAllRegionProvincesResponse
	(*FindRegionProvinceRequest)(nil),                          // 8: pb.FindRegionProvinceRequest
	(*FindRegionProvinceResponse)(nil),                         // 9: pb.FindRegionProvinceResponse
	(*UpdateRegionProvinceCustomRequest)(nil),                  // 10: pb.UpdateRegionProvinceCustomRequest
	(*RegionProvince)(nil),                                     // 11: pb.RegionProvince
	(*RPCSuccess)(nil),                                         // 12: pb.RPCSuccess
}
var file_service_region_province_proto_depIdxs = []int32{
	11, // 0: pb.FindAllEnabledRegionProvincesWithCountryIdResponse.regionProvinces:type_name -> pb.RegionProvince
	11, // 1: pb.FindEnabledRegionProvinceResponse.regionProvince:type_name -> pb.RegionProvince
	11, // 2: pb.FindAllRegionProvincesWithRegionCountryIdResponse.regionProvinces:type_name -> pb.RegionProvince
	11, // 3: pb.FindAllRegionProvincesResponse.regionProvinces:type_name -> pb.RegionProvince
	11, // 4: pb.FindRegionProvinceResponse.regionProvince:type_name -> pb.RegionProvince
	0,  // 5: pb.RegionProvinceService.findAllEnabledRegionProvincesWithCountryId:input_type -> pb.FindAllEnabledRegionProvincesWithCountryIdRequest
	2,  // 6: pb.RegionProvinceService.findEnabledRegionProvince:input_type -> pb.FindEnabledRegionProvinceRequest
	4,  // 7: pb.RegionProvinceService.findAllRegionProvincesWithRegionCountryId:input_type -> pb.FindAllRegionProvincesWithRegionCountryIdRequest
	6,  // 8: pb.RegionProvinceService.findAllRegionProvinces:input_type -> pb.FindAllRegionProvincesRequest
	8,  // 9: pb.RegionProvinceService.findRegionProvince:input_type -> pb.FindRegionProvinceRequest
	10, // 10: pb.RegionProvinceService.updateRegionProvinceCustom:input_type -> pb.UpdateRegionProvinceCustomRequest
	1,  // 11: pb.RegionProvinceService.findAllEnabledRegionProvincesWithCountryId:output_type -> pb.FindAllEnabledRegionProvincesWithCountryIdResponse
	3,  // 12: pb.RegionProvinceService.findEnabledRegionProvince:output_type -> pb.FindEnabledRegionProvinceResponse
	5,  // 13: pb.RegionProvinceService.findAllRegionProvincesWithRegionCountryId:output_type -> pb.FindAllRegionProvincesWithRegionCountryIdResponse
	7,  // 14: pb.RegionProvinceService.findAllRegionProvinces:output_type -> pb.FindAllRegionProvincesResponse
	9,  // 15: pb.RegionProvinceService.findRegionProvince:output_type -> pb.FindRegionProvinceResponse
	12, // 16: pb.RegionProvinceService.updateRegionProvinceCustom:output_type -> pb.RPCSuccess
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_service_region_province_proto_init() }
func file_service_region_province_proto_init() {
	if File_service_region_province_proto != nil {
		return
	}
	file_models_model_region_province_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_region_province_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledRegionProvincesWithCountryIdRequest); i {
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
		file_service_region_province_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllEnabledRegionProvincesWithCountryIdResponse); i {
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
		file_service_region_province_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledRegionProvinceRequest); i {
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
		file_service_region_province_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledRegionProvinceResponse); i {
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
		file_service_region_province_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllRegionProvincesWithRegionCountryIdRequest); i {
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
		file_service_region_province_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllRegionProvincesWithRegionCountryIdResponse); i {
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
		file_service_region_province_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllRegionProvincesRequest); i {
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
		file_service_region_province_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllRegionProvincesResponse); i {
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
		file_service_region_province_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRegionProvinceRequest); i {
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
		file_service_region_province_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindRegionProvinceResponse); i {
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
		file_service_region_province_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRegionProvinceCustomRequest); i {
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
			RawDescriptor: file_service_region_province_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_region_province_proto_goTypes,
		DependencyIndexes: file_service_region_province_proto_depIdxs,
		MessageInfos:      file_service_region_province_proto_msgTypes,
	}.Build()
	File_service_region_province_proto = out.File
	file_service_region_province_proto_rawDesc = nil
	file_service_region_province_proto_goTypes = nil
	file_service_region_province_proto_depIdxs = nil
}
