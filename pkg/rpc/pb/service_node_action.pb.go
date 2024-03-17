// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: service_node_action.proto

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

// 添加动作
type CreateNodeActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId       int64  `protobuf:"varint,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`            // 节点ID
	Role         string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`                 // 节点角色
	CondsJSON    []byte `protobuf:"bytes,3,opt,name=condsJSON,proto3" json:"condsJSON,omitempty"`       // 条件设置
	ActionJSON   []byte `protobuf:"bytes,4,opt,name=actionJSON,proto3" json:"actionJSON,omitempty"`     // 动作设置
	DurationJSON []byte `protobuf:"bytes,5,opt,name=durationJSON,proto3" json:"durationJSON,omitempty"` // 持续时间
}

func (x *CreateNodeActionRequest) Reset() {
	*x = CreateNodeActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNodeActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNodeActionRequest) ProtoMessage() {}

func (x *CreateNodeActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNodeActionRequest.ProtoReflect.Descriptor instead.
func (*CreateNodeActionRequest) Descriptor() ([]byte, []int) {
	return file_service_node_action_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNodeActionRequest) GetNodeId() int64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *CreateNodeActionRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *CreateNodeActionRequest) GetCondsJSON() []byte {
	if x != nil {
		return x.CondsJSON
	}
	return nil
}

func (x *CreateNodeActionRequest) GetActionJSON() []byte {
	if x != nil {
		return x.ActionJSON
	}
	return nil
}

func (x *CreateNodeActionRequest) GetDurationJSON() []byte {
	if x != nil {
		return x.DurationJSON
	}
	return nil
}

type CreateNodeActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeActionId int64 `protobuf:"varint,1,opt,name=nodeActionId,proto3" json:"nodeActionId,omitempty"`
}

func (x *CreateNodeActionResponse) Reset() {
	*x = CreateNodeActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_action_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNodeActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNodeActionResponse) ProtoMessage() {}

func (x *CreateNodeActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_action_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNodeActionResponse.ProtoReflect.Descriptor instead.
func (*CreateNodeActionResponse) Descriptor() ([]byte, []int) {
	return file_service_node_action_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNodeActionResponse) GetNodeActionId() int64 {
	if x != nil {
		return x.NodeActionId
	}
	return 0
}

// 删除动作
type DeleteNodeActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeActionId int64 `protobuf:"varint,1,opt,name=nodeActionId,proto3" json:"nodeActionId,omitempty"`
}

func (x *DeleteNodeActionRequest) Reset() {
	*x = DeleteNodeActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_action_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNodeActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNodeActionRequest) ProtoMessage() {}

func (x *DeleteNodeActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_action_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNodeActionRequest.ProtoReflect.Descriptor instead.
func (*DeleteNodeActionRequest) Descriptor() ([]byte, []int) {
	return file_service_node_action_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteNodeActionRequest) GetNodeActionId() int64 {
	if x != nil {
		return x.NodeActionId
	}
	return 0
}

// 修改动作
type UpdateNodeActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeActionId int64  `protobuf:"varint,1,opt,name=nodeActionId,proto3" json:"nodeActionId,omitempty"` // 动作ID
	CondsJSON    []byte `protobuf:"bytes,2,opt,name=condsJSON,proto3" json:"condsJSON,omitempty"`
	ActionJSON   []byte `protobuf:"bytes,3,opt,name=actionJSON,proto3" json:"actionJSON,omitempty"`
	DurationJSON []byte `protobuf:"bytes,4,opt,name=durationJSON,proto3" json:"durationJSON,omitempty"` // 持续时间
	IsOn         bool   `protobuf:"varint,5,opt,name=isOn,proto3" json:"isOn,omitempty"`                // 是否启用
}

func (x *UpdateNodeActionRequest) Reset() {
	*x = UpdateNodeActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_action_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNodeActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeActionRequest) ProtoMessage() {}

func (x *UpdateNodeActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_action_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeActionRequest.ProtoReflect.Descriptor instead.
func (*UpdateNodeActionRequest) Descriptor() ([]byte, []int) {
	return file_service_node_action_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateNodeActionRequest) GetNodeActionId() int64 {
	if x != nil {
		return x.NodeActionId
	}
	return 0
}

func (x *UpdateNodeActionRequest) GetCondsJSON() []byte {
	if x != nil {
		return x.CondsJSON
	}
	return nil
}

func (x *UpdateNodeActionRequest) GetActionJSON() []byte {
	if x != nil {
		return x.ActionJSON
	}
	return nil
}

func (x *UpdateNodeActionRequest) GetDurationJSON() []byte {
	if x != nil {
		return x.DurationJSON
	}
	return nil
}

func (x *UpdateNodeActionRequest) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

// 列出某个节点的所有动作
type FindAllNodeActionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId int64  `protobuf:"varint,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"` // 节点ID
	Role   string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`      // 节点角色
}

func (x *FindAllNodeActionsRequest) Reset() {
	*x = FindAllNodeActionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_action_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllNodeActionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllNodeActionsRequest) ProtoMessage() {}

func (x *FindAllNodeActionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_action_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllNodeActionsRequest.ProtoReflect.Descriptor instead.
func (*FindAllNodeActionsRequest) Descriptor() ([]byte, []int) {
	return file_service_node_action_proto_rawDescGZIP(), []int{4}
}

func (x *FindAllNodeActionsRequest) GetNodeId() int64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *FindAllNodeActionsRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type FindAllNodeActionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeActions []*NodeAction `protobuf:"bytes,1,rep,name=nodeActions,proto3" json:"nodeActions,omitempty"` // 动作列表
}

func (x *FindAllNodeActionsResponse) Reset() {
	*x = FindAllNodeActionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_action_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllNodeActionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllNodeActionsResponse) ProtoMessage() {}

func (x *FindAllNodeActionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_action_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllNodeActionsResponse.ProtoReflect.Descriptor instead.
func (*FindAllNodeActionsResponse) Descriptor() ([]byte, []int) {
	return file_service_node_action_proto_rawDescGZIP(), []int{5}
}

func (x *FindAllNodeActionsResponse) GetNodeActions() []*NodeAction {
	if x != nil {
		return x.NodeActions
	}
	return nil
}

// 查找单个节点动作
type FindNodeActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeActionId int64 `protobuf:"varint,1,opt,name=nodeActionId,proto3" json:"nodeActionId,omitempty"` // 动作ID
}

func (x *FindNodeActionRequest) Reset() {
	*x = FindNodeActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_action_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindNodeActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindNodeActionRequest) ProtoMessage() {}

func (x *FindNodeActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_action_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindNodeActionRequest.ProtoReflect.Descriptor instead.
func (*FindNodeActionRequest) Descriptor() ([]byte, []int) {
	return file_service_node_action_proto_rawDescGZIP(), []int{6}
}

func (x *FindNodeActionRequest) GetNodeActionId() int64 {
	if x != nil {
		return x.NodeActionId
	}
	return 0
}

type FindNodeActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeAction *NodeAction `protobuf:"bytes,1,opt,name=nodeAction,proto3" json:"nodeAction,omitempty"`
}

func (x *FindNodeActionResponse) Reset() {
	*x = FindNodeActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_action_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindNodeActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindNodeActionResponse) ProtoMessage() {}

func (x *FindNodeActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_action_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindNodeActionResponse.ProtoReflect.Descriptor instead.
func (*FindNodeActionResponse) Descriptor() ([]byte, []int) {
	return file_service_node_action_proto_rawDescGZIP(), []int{7}
}

func (x *FindNodeActionResponse) GetNodeAction() *NodeAction {
	if x != nil {
		return x.NodeAction
	}
	return nil
}

// 设置节点动作排序
type UpdateNodeActionOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeActionIds []int64 `protobuf:"varint,1,rep,packed,name=nodeActionIds,proto3" json:"nodeActionIds,omitempty"` // 节点动作ID列表
}

func (x *UpdateNodeActionOrdersRequest) Reset() {
	*x = UpdateNodeActionOrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_node_action_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNodeActionOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNodeActionOrdersRequest) ProtoMessage() {}

func (x *UpdateNodeActionOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_node_action_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNodeActionOrdersRequest.ProtoReflect.Descriptor instead.
func (*UpdateNodeActionOrdersRequest) Descriptor() ([]byte, []int) {
	return file_service_node_action_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateNodeActionOrdersRequest) GetNodeActionIds() []int64 {
	if x != nil {
		return x.NodeActionIds
	}
	return nil
}

var File_service_node_action_proto protoreflect.FileDescriptor

var file_service_node_action_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x1e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x17, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x4a, 0x53, 0x4f, 0x4e,
	0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x53, 0x4f, 0x4e,
	0x12, 0x22, 0x0a, 0x0c, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x53, 0x4f, 0x4e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4a, 0x53, 0x4f, 0x4e, 0x22, 0x3e, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0xb3, 0x01, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x4a, 0x53, 0x4f, 0x4e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x4a, 0x53, 0x4f,
	0x4e, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x53, 0x4f, 0x4e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x53, 0x4f,
	0x4e, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x53, 0x4f,
	0x4e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x22, 0x47, 0x0a, 0x19, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x22, 0x4e, 0x0a, 0x1a, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x64,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x30, 0x0a, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x3b, 0x0a, 0x15, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x6e,
	0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x48, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0a, 0x6e, 0x6f, 0x64,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6e,
	0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x45, 0x0a, 0x1d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x6f,
	0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73,
	0x32, 0xcf, 0x03, 0x0a, 0x11, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x3f, 0x0a, 0x10, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x53, 0x0a, 0x12, 0x66, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0e,
	0x66, 0x69, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19,
	0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x16, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x21, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_service_node_action_proto_rawDescOnce sync.Once
	file_service_node_action_proto_rawDescData = file_service_node_action_proto_rawDesc
)

func file_service_node_action_proto_rawDescGZIP() []byte {
	file_service_node_action_proto_rawDescOnce.Do(func() {
		file_service_node_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_node_action_proto_rawDescData)
	})
	return file_service_node_action_proto_rawDescData
}

var file_service_node_action_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_service_node_action_proto_goTypes = []interface{}{
	(*CreateNodeActionRequest)(nil),       // 0: pb.CreateNodeActionRequest
	(*CreateNodeActionResponse)(nil),      // 1: pb.CreateNodeActionResponse
	(*DeleteNodeActionRequest)(nil),       // 2: pb.DeleteNodeActionRequest
	(*UpdateNodeActionRequest)(nil),       // 3: pb.UpdateNodeActionRequest
	(*FindAllNodeActionsRequest)(nil),     // 4: pb.FindAllNodeActionsRequest
	(*FindAllNodeActionsResponse)(nil),    // 5: pb.FindAllNodeActionsResponse
	(*FindNodeActionRequest)(nil),         // 6: pb.FindNodeActionRequest
	(*FindNodeActionResponse)(nil),        // 7: pb.FindNodeActionResponse
	(*UpdateNodeActionOrdersRequest)(nil), // 8: pb.UpdateNodeActionOrdersRequest
	(*NodeAction)(nil),                    // 9: pb.NodeAction
	(*RPCSuccess)(nil),                    // 10: pb.RPCSuccess
}
var file_service_node_action_proto_depIdxs = []int32{
	9,  // 0: pb.FindAllNodeActionsResponse.nodeActions:type_name -> pb.NodeAction
	9,  // 1: pb.FindNodeActionResponse.nodeAction:type_name -> pb.NodeAction
	0,  // 2: pb.NodeActionService.createNodeAction:input_type -> pb.CreateNodeActionRequest
	2,  // 3: pb.NodeActionService.deleteNodeAction:input_type -> pb.DeleteNodeActionRequest
	3,  // 4: pb.NodeActionService.updateNodeAction:input_type -> pb.UpdateNodeActionRequest
	4,  // 5: pb.NodeActionService.findAllNodeActions:input_type -> pb.FindAllNodeActionsRequest
	6,  // 6: pb.NodeActionService.findNodeAction:input_type -> pb.FindNodeActionRequest
	8,  // 7: pb.NodeActionService.updateNodeActionOrders:input_type -> pb.UpdateNodeActionOrdersRequest
	1,  // 8: pb.NodeActionService.createNodeAction:output_type -> pb.CreateNodeActionResponse
	10, // 9: pb.NodeActionService.deleteNodeAction:output_type -> pb.RPCSuccess
	10, // 10: pb.NodeActionService.updateNodeAction:output_type -> pb.RPCSuccess
	5,  // 11: pb.NodeActionService.findAllNodeActions:output_type -> pb.FindAllNodeActionsResponse
	7,  // 12: pb.NodeActionService.findNodeAction:output_type -> pb.FindNodeActionResponse
	10, // 13: pb.NodeActionService.updateNodeActionOrders:output_type -> pb.RPCSuccess
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_service_node_action_proto_init() }
func file_service_node_action_proto_init() {
	if File_service_node_action_proto != nil {
		return
	}
	file_models_model_node_action_proto_init()
	file_models_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_node_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNodeActionRequest); i {
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
		file_service_node_action_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNodeActionResponse); i {
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
		file_service_node_action_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNodeActionRequest); i {
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
		file_service_node_action_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNodeActionRequest); i {
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
		file_service_node_action_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllNodeActionsRequest); i {
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
		file_service_node_action_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllNodeActionsResponse); i {
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
		file_service_node_action_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindNodeActionRequest); i {
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
		file_service_node_action_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindNodeActionResponse); i {
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
		file_service_node_action_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNodeActionOrdersRequest); i {
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
			RawDescriptor: file_service_node_action_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_node_action_proto_goTypes,
		DependencyIndexes: file_service_node_action_proto_depIdxs,
		MessageInfos:      file_service_node_action_proto_msgTypes,
	}.Build()
	File_service_node_action_proto = out.File
	file_service_node_action_proto_rawDesc = nil
	file_service_node_action_proto_goTypes = nil
	file_service_node_action_proto_depIdxs = nil
}
