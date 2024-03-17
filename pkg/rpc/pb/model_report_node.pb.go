// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: models/model_report_node.proto

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

type ReportNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64              `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UniqueId         string             `protobuf:"bytes,2,opt,name=uniqueId,proto3" json:"uniqueId,omitempty"`
	Secret           string             `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty"`
	IsOn             bool               `protobuf:"varint,4,opt,name=isOn,proto3" json:"isOn,omitempty"`
	Name             string             `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Location         string             `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	Isp              string             `protobuf:"bytes,7,opt,name=isp,proto3" json:"isp,omitempty"`
	IsActive         bool               `protobuf:"varint,8,opt,name=isActive,proto3" json:"isActive,omitempty"`
	StatusJSON       []byte             `protobuf:"bytes,9,opt,name=statusJSON,proto3" json:"statusJSON,omitempty"`
	AllowIPs         []string           `protobuf:"bytes,10,rep,name=allowIPs,proto3" json:"allowIPs,omitempty"`
	ReportNodeGroups []*ReportNodeGroup `protobuf:"bytes,11,rep,name=reportNodeGroups,proto3" json:"reportNodeGroups,omitempty"`
}

func (x *ReportNode) Reset() {
	*x = ReportNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_report_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportNode) ProtoMessage() {}

func (x *ReportNode) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_report_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportNode.ProtoReflect.Descriptor instead.
func (*ReportNode) Descriptor() ([]byte, []int) {
	return file_models_model_report_node_proto_rawDescGZIP(), []int{0}
}

func (x *ReportNode) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReportNode) GetUniqueId() string {
	if x != nil {
		return x.UniqueId
	}
	return ""
}

func (x *ReportNode) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *ReportNode) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *ReportNode) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReportNode) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *ReportNode) GetIsp() string {
	if x != nil {
		return x.Isp
	}
	return ""
}

func (x *ReportNode) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *ReportNode) GetStatusJSON() []byte {
	if x != nil {
		return x.StatusJSON
	}
	return nil
}

func (x *ReportNode) GetAllowIPs() []string {
	if x != nil {
		return x.AllowIPs
	}
	return nil
}

func (x *ReportNode) GetReportNodeGroups() []*ReportNodeGroup {
	if x != nil {
		return x.ReportNodeGroups
	}
	return nil
}

var File_models_model_report_node_proto protoreflect.FileDescriptor

var file_models_model_report_node_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x24, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x02, 0x0a, 0x0a, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x73, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x69, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x50, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x50, 0x73, 0x12, 0x3f, 0x0a, 0x10, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x10, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_report_node_proto_rawDescOnce sync.Once
	file_models_model_report_node_proto_rawDescData = file_models_model_report_node_proto_rawDesc
)

func file_models_model_report_node_proto_rawDescGZIP() []byte {
	file_models_model_report_node_proto_rawDescOnce.Do(func() {
		file_models_model_report_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_report_node_proto_rawDescData)
	})
	return file_models_model_report_node_proto_rawDescData
}

var file_models_model_report_node_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_report_node_proto_goTypes = []interface{}{
	(*ReportNode)(nil),      // 0: pb.ReportNode
	(*ReportNodeGroup)(nil), // 1: pb.ReportNodeGroup
}
var file_models_model_report_node_proto_depIdxs = []int32{
	1, // 0: pb.ReportNode.reportNodeGroups:type_name -> pb.ReportNodeGroup
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_models_model_report_node_proto_init() }
func file_models_model_report_node_proto_init() {
	if File_models_model_report_node_proto != nil {
		return
	}
	file_models_model_report_node_group_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_report_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportNode); i {
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
			RawDescriptor: file_models_model_report_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_report_node_proto_goTypes,
		DependencyIndexes: file_models_model_report_node_proto_depIdxs,
		MessageInfos:      file_models_model_report_node_proto_msgTypes,
	}.Build()
	File_models_model_report_node_proto = out.File
	file_models_model_report_node_proto_rawDesc = nil
	file_models_model_report_node_proto_goTypes = nil
	file_models_model_report_node_proto_depIdxs = nil
}
