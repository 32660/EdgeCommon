// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: service_ns_route_category.proto

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

// 创建线路分类
type CreateNSRouteCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateNSRouteCategoryRequest) Reset() {
	*x = CreateNSRouteCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_route_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNSRouteCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNSRouteCategoryRequest) ProtoMessage() {}

func (x *CreateNSRouteCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_route_category_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNSRouteCategoryRequest.ProtoReflect.Descriptor instead.
func (*CreateNSRouteCategoryRequest) Descriptor() ([]byte, []int) {
	return file_service_ns_route_category_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNSRouteCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateNSRouteCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsRouteCategoryId int64 `protobuf:"varint,1,opt,name=nsRouteCategoryId,proto3" json:"nsRouteCategoryId,omitempty"`
}

func (x *CreateNSRouteCategoryResponse) Reset() {
	*x = CreateNSRouteCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_route_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNSRouteCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNSRouteCategoryResponse) ProtoMessage() {}

func (x *CreateNSRouteCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_route_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNSRouteCategoryResponse.ProtoReflect.Descriptor instead.
func (*CreateNSRouteCategoryResponse) Descriptor() ([]byte, []int) {
	return file_service_ns_route_category_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNSRouteCategoryResponse) GetNsRouteCategoryId() int64 {
	if x != nil {
		return x.NsRouteCategoryId
	}
	return 0
}

// 修改线路分类
type UpdateNSRouteCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsRouteCategoryId int64  `protobuf:"varint,1,opt,name=nsRouteCategoryId,proto3" json:"nsRouteCategoryId,omitempty"`
	Name              string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsOn              bool   `protobuf:"varint,3,opt,name=isOn,proto3" json:"isOn,omitempty"`
}

func (x *UpdateNSRouteCategoryRequest) Reset() {
	*x = UpdateNSRouteCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_route_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNSRouteCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNSRouteCategoryRequest) ProtoMessage() {}

func (x *UpdateNSRouteCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_route_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNSRouteCategoryRequest.ProtoReflect.Descriptor instead.
func (*UpdateNSRouteCategoryRequest) Descriptor() ([]byte, []int) {
	return file_service_ns_route_category_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateNSRouteCategoryRequest) GetNsRouteCategoryId() int64 {
	if x != nil {
		return x.NsRouteCategoryId
	}
	return 0
}

func (x *UpdateNSRouteCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateNSRouteCategoryRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

// 删除线路分类
type DeleteNSRouteCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsRouteCategoryId int64 `protobuf:"varint,1,opt,name=nsRouteCategoryId,proto3" json:"nsRouteCategoryId,omitempty"`
}

func (x *DeleteNSRouteCategoryRequest) Reset() {
	*x = DeleteNSRouteCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_route_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNSRouteCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNSRouteCategoryRequest) ProtoMessage() {}

func (x *DeleteNSRouteCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_route_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNSRouteCategoryRequest.ProtoReflect.Descriptor instead.
func (*DeleteNSRouteCategoryRequest) Descriptor() ([]byte, []int) {
	return file_service_ns_route_category_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteNSRouteCategoryRequest) GetNsRouteCategoryId() int64 {
	if x != nil {
		return x.NsRouteCategoryId
	}
	return 0
}

// 列出所有线路分类
type FindAllNSRouteCategoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindAllNSRouteCategoriesRequest) Reset() {
	*x = FindAllNSRouteCategoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_route_category_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllNSRouteCategoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllNSRouteCategoriesRequest) ProtoMessage() {}

func (x *FindAllNSRouteCategoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_route_category_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllNSRouteCategoriesRequest.ProtoReflect.Descriptor instead.
func (*FindAllNSRouteCategoriesRequest) Descriptor() ([]byte, []int) {
	return file_service_ns_route_category_proto_rawDescGZIP(), []int{4}
}

type FindAllNSRouteCategoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsRouteCategories []*NSRouteCategory `protobuf:"bytes,1,rep,name=nsRouteCategories,proto3" json:"nsRouteCategories,omitempty"`
}

func (x *FindAllNSRouteCategoriesResponse) Reset() {
	*x = FindAllNSRouteCategoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_route_category_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllNSRouteCategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllNSRouteCategoriesResponse) ProtoMessage() {}

func (x *FindAllNSRouteCategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_route_category_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllNSRouteCategoriesResponse.ProtoReflect.Descriptor instead.
func (*FindAllNSRouteCategoriesResponse) Descriptor() ([]byte, []int) {
	return file_service_ns_route_category_proto_rawDescGZIP(), []int{5}
}

func (x *FindAllNSRouteCategoriesResponse) GetNsRouteCategories() []*NSRouteCategory {
	if x != nil {
		return x.NsRouteCategories
	}
	return nil
}

// 对线路分类进行排序
type UpdateNSRouteCategoryOrders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsRouteCategoryIds []int64 `protobuf:"varint,1,rep,packed,name=nsRouteCategoryIds,proto3" json:"nsRouteCategoryIds,omitempty"`
}

func (x *UpdateNSRouteCategoryOrders) Reset() {
	*x = UpdateNSRouteCategoryOrders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_route_category_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNSRouteCategoryOrders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNSRouteCategoryOrders) ProtoMessage() {}

func (x *UpdateNSRouteCategoryOrders) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_route_category_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNSRouteCategoryOrders.ProtoReflect.Descriptor instead.
func (*UpdateNSRouteCategoryOrders) Descriptor() ([]byte, []int) {
	return file_service_ns_route_category_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateNSRouteCategoryOrders) GetNsRouteCategoryIds() []int64 {
	if x != nil {
		return x.NsRouteCategoryIds
	}
	return nil
}

// 查找单个线路分类
type FindNSRouteCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsRouteCategoryId int64 `protobuf:"varint,1,opt,name=nsRouteCategoryId,proto3" json:"nsRouteCategoryId,omitempty"`
}

func (x *FindNSRouteCategoryRequest) Reset() {
	*x = FindNSRouteCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_route_category_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindNSRouteCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindNSRouteCategoryRequest) ProtoMessage() {}

func (x *FindNSRouteCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_route_category_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindNSRouteCategoryRequest.ProtoReflect.Descriptor instead.
func (*FindNSRouteCategoryRequest) Descriptor() ([]byte, []int) {
	return file_service_ns_route_category_proto_rawDescGZIP(), []int{7}
}

func (x *FindNSRouteCategoryRequest) GetNsRouteCategoryId() int64 {
	if x != nil {
		return x.NsRouteCategoryId
	}
	return 0
}

type FindNSRouteCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NsRouteCategory *NSRouteCategory `protobuf:"bytes,1,opt,name=nsRouteCategory,proto3" json:"nsRouteCategory,omitempty"`
}

func (x *FindNSRouteCategoryResponse) Reset() {
	*x = FindNSRouteCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ns_route_category_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindNSRouteCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindNSRouteCategoryResponse) ProtoMessage() {}

func (x *FindNSRouteCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ns_route_category_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindNSRouteCategoryResponse.ProtoReflect.Descriptor instead.
func (*FindNSRouteCategoryResponse) Descriptor() ([]byte, []int) {
	return file_service_ns_route_category_proto_rawDescGZIP(), []int{8}
}

func (x *FindNSRouteCategoryResponse) GetNsRouteCategory() *NSRouteCategory {
	if x != nil {
		return x.NsRouteCategory
	}
	return nil
}

var File_service_ns_route_category_proto protoreflect.FileDescriptor

var file_service_ns_route_category_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x73, 0x5f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x24, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x73, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4d, 0x0a, 0x1d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x6e,
	0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6e, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x74, 0x0a, 0x1c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x6e, 0x73, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6e, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69,
	0x73, 0x4f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x22,
	0x4c, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2c, 0x0a, 0x11, 0x6e, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6e, 0x73, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22, 0x21, 0x0a,
	0x1f, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x65, 0x0a, 0x20, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4e, 0x53, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x11, 0x6e, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x52, 0x11, 0x6e, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x4d, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x6e, 0x73, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x12, 0x6e, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x73, 0x22, 0x4a, 0x0a, 0x1a, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x53,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x6e, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x11, 0x6e, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x49, 0x64, 0x22, 0x5c, 0x0a, 0x1b, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3d, 0x0a, 0x0f, 0x6e, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e,
	0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x0f, 0x6e, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x32, 0x9b, 0x04, 0x0a, 0x16, 0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x15, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x15, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x53,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x49, 0x0a, 0x15, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x53,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x20, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x65, 0x0a, 0x18, 0x66, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4e, 0x53, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x1b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x56, 0x0a, 0x13, 0x66, 0x69, 0x6e, 0x64, 0x4e, 0x53,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1e, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_ns_route_category_proto_rawDescOnce sync.Once
	file_service_ns_route_category_proto_rawDescData = file_service_ns_route_category_proto_rawDesc
)

func file_service_ns_route_category_proto_rawDescGZIP() []byte {
	file_service_ns_route_category_proto_rawDescOnce.Do(func() {
		file_service_ns_route_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_ns_route_category_proto_rawDescData)
	})
	return file_service_ns_route_category_proto_rawDescData
}

var file_service_ns_route_category_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_service_ns_route_category_proto_goTypes = []interface{}{
	(*CreateNSRouteCategoryRequest)(nil),     // 0: pb.CreateNSRouteCategoryRequest
	(*CreateNSRouteCategoryResponse)(nil),    // 1: pb.CreateNSRouteCategoryResponse
	(*UpdateNSRouteCategoryRequest)(nil),     // 2: pb.UpdateNSRouteCategoryRequest
	(*DeleteNSRouteCategoryRequest)(nil),     // 3: pb.DeleteNSRouteCategoryRequest
	(*FindAllNSRouteCategoriesRequest)(nil),  // 4: pb.FindAllNSRouteCategoriesRequest
	(*FindAllNSRouteCategoriesResponse)(nil), // 5: pb.FindAllNSRouteCategoriesResponse
	(*UpdateNSRouteCategoryOrders)(nil),      // 6: pb.UpdateNSRouteCategoryOrders
	(*FindNSRouteCategoryRequest)(nil),       // 7: pb.FindNSRouteCategoryRequest
	(*FindNSRouteCategoryResponse)(nil),      // 8: pb.FindNSRouteCategoryResponse
	(*NSRouteCategory)(nil),                  // 9: pb.NSRouteCategory
	(*RPCSuccess)(nil),                       // 10: pb.RPCSuccess
}
var file_service_ns_route_category_proto_depIdxs = []int32{
	9,  // 0: pb.FindAllNSRouteCategoriesResponse.nsRouteCategories:type_name -> pb.NSRouteCategory
	9,  // 1: pb.FindNSRouteCategoryResponse.nsRouteCategory:type_name -> pb.NSRouteCategory
	0,  // 2: pb.NSRouteCategoryService.createNSRouteCategory:input_type -> pb.CreateNSRouteCategoryRequest
	2,  // 3: pb.NSRouteCategoryService.updateNSRouteCategory:input_type -> pb.UpdateNSRouteCategoryRequest
	3,  // 4: pb.NSRouteCategoryService.deleteNSRouteCategory:input_type -> pb.DeleteNSRouteCategoryRequest
	4,  // 5: pb.NSRouteCategoryService.findAllNSRouteCategories:input_type -> pb.FindAllNSRouteCategoriesRequest
	6,  // 6: pb.NSRouteCategoryService.updateNSRouteCategoryOrders:input_type -> pb.UpdateNSRouteCategoryOrders
	7,  // 7: pb.NSRouteCategoryService.findNSRouteCategory:input_type -> pb.FindNSRouteCategoryRequest
	1,  // 8: pb.NSRouteCategoryService.createNSRouteCategory:output_type -> pb.CreateNSRouteCategoryResponse
	10, // 9: pb.NSRouteCategoryService.updateNSRouteCategory:output_type -> pb.RPCSuccess
	10, // 10: pb.NSRouteCategoryService.deleteNSRouteCategory:output_type -> pb.RPCSuccess
	5,  // 11: pb.NSRouteCategoryService.findAllNSRouteCategories:output_type -> pb.FindAllNSRouteCategoriesResponse
	10, // 12: pb.NSRouteCategoryService.updateNSRouteCategoryOrders:output_type -> pb.RPCSuccess
	8,  // 13: pb.NSRouteCategoryService.findNSRouteCategory:output_type -> pb.FindNSRouteCategoryResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_service_ns_route_category_proto_init() }
func file_service_ns_route_category_proto_init() {
	if File_service_ns_route_category_proto != nil {
		return
	}
	file_models_model_ns_route_category_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_ns_route_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNSRouteCategoryRequest); i {
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
		file_service_ns_route_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNSRouteCategoryResponse); i {
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
		file_service_ns_route_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNSRouteCategoryRequest); i {
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
		file_service_ns_route_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNSRouteCategoryRequest); i {
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
		file_service_ns_route_category_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllNSRouteCategoriesRequest); i {
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
		file_service_ns_route_category_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllNSRouteCategoriesResponse); i {
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
		file_service_ns_route_category_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNSRouteCategoryOrders); i {
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
		file_service_ns_route_category_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindNSRouteCategoryRequest); i {
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
		file_service_ns_route_category_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindNSRouteCategoryResponse); i {
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
			RawDescriptor: file_service_ns_route_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_ns_route_category_proto_goTypes,
		DependencyIndexes: file_service_ns_route_category_proto_depIdxs,
		MessageInfos:      file_service_ns_route_category_proto_msgTypes,
	}.Build()
	File_service_ns_route_category_proto = out.File
	file_service_ns_route_category_proto_rawDesc = nil
	file_service_ns_route_category_proto_goTypes = nil
	file_service_ns_route_category_proto_depIdxs = nil
}
