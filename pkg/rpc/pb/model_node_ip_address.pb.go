// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.4
// source: models/model_node_ip_address.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 节点地址
type NodeIPAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NodeId      int64  `protobuf:"varint,2,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Ip          string `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	State       int64  `protobuf:"varint,6,opt,name=state,proto3" json:"state,omitempty"`
	Order       int64  `protobuf:"varint,7,opt,name=order,proto3" json:"order,omitempty"`
	CanAccess   bool   `protobuf:"varint,8,opt,name=canAccess,proto3" json:"canAccess,omitempty"`
	IsOn        bool   `protobuf:"varint,9,opt,name=isOn,proto3" json:"isOn,omitempty"`
	IsUp        bool   `protobuf:"varint,10,opt,name=isUp,proto3" json:"isUp,omitempty"`
	Role        string `protobuf:"bytes,12,opt,name=role,proto3" json:"role,omitempty"`
	BackupIP    string `protobuf:"bytes,13,opt,name=backupIP,proto3" json:"backupIP,omitempty"`
	IsHealthy   bool   `protobuf:"varint,14,opt,name=isHealthy,proto3" json:"isHealthy,omitempty"`
}

func (x *NodeIPAddress) Reset() {
	*x = NodeIPAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_node_ip_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeIPAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeIPAddress) ProtoMessage() {}

func (x *NodeIPAddress) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_node_ip_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeIPAddress.ProtoReflect.Descriptor instead.
func (*NodeIPAddress) Descriptor() ([]byte, []int) {
	return file_models_model_node_ip_address_proto_rawDescGZIP(), []int{0}
}

func (x *NodeIPAddress) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NodeIPAddress) GetNodeId() int64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *NodeIPAddress) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeIPAddress) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *NodeIPAddress) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NodeIPAddress) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *NodeIPAddress) GetOrder() int64 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *NodeIPAddress) GetCanAccess() bool {
	if x != nil {
		return x.CanAccess
	}
	return false
}

func (x *NodeIPAddress) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *NodeIPAddress) GetIsUp() bool {
	if x != nil {
		return x.IsUp
	}
	return false
}

func (x *NodeIPAddress) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *NodeIPAddress) GetBackupIP() string {
	if x != nil {
		return x.BackupIP
	}
	return ""
}

func (x *NodeIPAddress) GetIsHealthy() bool {
	if x != nil {
		return x.IsHealthy
	}
	return false
}

var File_models_model_node_ip_address_proto protoreflect.FileDescriptor

var file_models_model_node_ip_address_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xbd, 0x02, 0x0a, 0x0d, 0x4e, 0x6f, 0x64,
	0x65, 0x49, 0x50, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x61, 0x6e, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x55, 0x70, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x55, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x50, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x49, 0x50, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_node_ip_address_proto_rawDescOnce sync.Once
	file_models_model_node_ip_address_proto_rawDescData = file_models_model_node_ip_address_proto_rawDesc
)

func file_models_model_node_ip_address_proto_rawDescGZIP() []byte {
	file_models_model_node_ip_address_proto_rawDescOnce.Do(func() {
		file_models_model_node_ip_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_node_ip_address_proto_rawDescData)
	})
	return file_models_model_node_ip_address_proto_rawDescData
}

var file_models_model_node_ip_address_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_node_ip_address_proto_goTypes = []interface{}{
	(*NodeIPAddress)(nil), // 0: pb.NodeIPAddress
}
var file_models_model_node_ip_address_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_node_ip_address_proto_init() }
func file_models_model_node_ip_address_proto_init() {
	if File_models_model_node_ip_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_node_ip_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeIPAddress); i {
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
			RawDescriptor: file_models_model_node_ip_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_node_ip_address_proto_goTypes,
		DependencyIndexes: file_models_model_node_ip_address_proto_depIdxs,
		MessageInfos:      file_models_model_node_ip_address_proto_msgTypes,
	}.Build()
	File_models_model_node_ip_address_proto = out.File
	file_models_model_node_ip_address_proto_rawDesc = nil
	file_models_model_node_ip_address_proto_goTypes = nil
	file_models_model_node_ip_address_proto_depIdxs = nil
}
