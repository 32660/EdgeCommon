// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: models/model_ns_record.proto

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

// 域名记录
type NSRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                    // 记录ID
	Description     string     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`   // 备注
	Name            string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`                 // 记录名
	Type            string     `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`                 // 记录类型
	Value           string     `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`               // 记录值
	MxPriority      int32      `protobuf:"varint,12,opt,name=mxPriority,proto3" json:"mxPriority,omitempty"`   // mx优先级
	SrvPriority     int32      `protobuf:"varint,13,opt,name=srvPriority,proto3" json:"srvPriority,omitempty"` // SRV优先级
	SrvWeight       int32      `protobuf:"varint,14,opt,name=srvWeight,proto3" json:"srvWeight,omitempty"`     // SRV权重
	SrvPort         int32      `protobuf:"varint,15,opt,name=srvPort,proto3" json:"srvPort,omitempty"`         // SRV端口
	CaaFlag         int32      `protobuf:"varint,16,opt,name=caaFlag,proto3" json:"caaFlag,omitempty"`         // CAA Flag
	CaaTag          string     `protobuf:"bytes,17,opt,name=caaTag,proto3" json:"caaTag,omitempty"`            // CAA TAG
	Ttl             int32      `protobuf:"varint,6,opt,name=ttl,proto3" json:"ttl,omitempty"`                  // TTL
	Weight          int32      `protobuf:"varint,7,opt,name=weight,proto3" json:"weight,omitempty"`            // 权重
	CreatedAt       int64      `protobuf:"varint,8,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	IsDeleted       bool       `protobuf:"varint,9,opt,name=isDeleted,proto3" json:"isDeleted,omitempty"`
	Version         int64      `protobuf:"varint,10,opt,name=version,proto3" json:"version,omitempty"`
	IsOn            bool       `protobuf:"varint,11,opt,name=isOn,proto3" json:"isOn,omitempty"`                      // 是否启用
	HealthCheckJSON []byte     `protobuf:"bytes,18,opt,name=healthCheckJSON,proto3" json:"healthCheckJSON,omitempty"` // 健康检查配置
	IsUp            bool       `protobuf:"varint,19,opt,name=isUp,proto3" json:"isUp,omitempty"`                      // 是否在线（根据健康检查结果）
	NsDomain        *NSDomain  `protobuf:"bytes,30,opt,name=nsDomain,proto3" json:"nsDomain,omitempty"`               // 所属域名
	NsRoutes        []*NSRoute `protobuf:"bytes,31,rep,name=nsRoutes,proto3" json:"nsRoutes,omitempty"`               // 线路
}

func (x *NSRecord) Reset() {
	*x = NSRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_ns_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NSRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NSRecord) ProtoMessage() {}

func (x *NSRecord) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_ns_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NSRecord.ProtoReflect.Descriptor instead.
func (*NSRecord) Descriptor() ([]byte, []int) {
	return file_models_model_ns_record_proto_rawDescGZIP(), []int{0}
}

func (x *NSRecord) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NSRecord) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NSRecord) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NSRecord) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *NSRecord) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *NSRecord) GetMxPriority() int32 {
	if x != nil {
		return x.MxPriority
	}
	return 0
}

func (x *NSRecord) GetSrvPriority() int32 {
	if x != nil {
		return x.SrvPriority
	}
	return 0
}

func (x *NSRecord) GetSrvWeight() int32 {
	if x != nil {
		return x.SrvWeight
	}
	return 0
}

func (x *NSRecord) GetSrvPort() int32 {
	if x != nil {
		return x.SrvPort
	}
	return 0
}

func (x *NSRecord) GetCaaFlag() int32 {
	if x != nil {
		return x.CaaFlag
	}
	return 0
}

func (x *NSRecord) GetCaaTag() string {
	if x != nil {
		return x.CaaTag
	}
	return ""
}

func (x *NSRecord) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *NSRecord) GetWeight() int32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *NSRecord) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *NSRecord) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

func (x *NSRecord) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *NSRecord) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *NSRecord) GetHealthCheckJSON() []byte {
	if x != nil {
		return x.HealthCheckJSON
	}
	return nil
}

func (x *NSRecord) GetIsUp() bool {
	if x != nil {
		return x.IsUp
	}
	return false
}

func (x *NSRecord) GetNsDomain() *NSDomain {
	if x != nil {
		return x.NsDomain
	}
	return nil
}

func (x *NSRecord) GetNsRoutes() []*NSRoute {
	if x != nil {
		return x.NsRoutes
	}
	return nil
}

var File_models_model_ns_record_proto protoreflect.FileDescriptor

var file_models_model_ns_record_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e,
	0x73, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x1c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x5f, 0x6e, 0x73, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e,
	0x73, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x04,
	0x0a, 0x08, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x78,
	0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x6d, 0x78, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x72,
	0x76, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x73, 0x72, 0x76, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x72, 0x76, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x73, 0x72, 0x76, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x72,
	0x76, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x72, 0x76,
	0x50, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x61, 0x46, 0x6c, 0x61, 0x67, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x61, 0x61, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x61, 0x61, 0x54, 0x61, 0x67, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x63, 0x61, 0x61, 0x54, 0x61, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x55, 0x70, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x55, 0x70, 0x12, 0x28, 0x0a, 0x08, 0x6e, 0x73, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x4e, 0x53, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x08, 0x6e, 0x73, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x27, 0x0a, 0x08, 0x6e, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x1f,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x53, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x52, 0x08, 0x6e, 0x73, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_ns_record_proto_rawDescOnce sync.Once
	file_models_model_ns_record_proto_rawDescData = file_models_model_ns_record_proto_rawDesc
)

func file_models_model_ns_record_proto_rawDescGZIP() []byte {
	file_models_model_ns_record_proto_rawDescOnce.Do(func() {
		file_models_model_ns_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_ns_record_proto_rawDescData)
	})
	return file_models_model_ns_record_proto_rawDescData
}

var file_models_model_ns_record_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_ns_record_proto_goTypes = []interface{}{
	(*NSRecord)(nil), // 0: pb.NSRecord
	(*NSDomain)(nil), // 1: pb.NSDomain
	(*NSRoute)(nil),  // 2: pb.NSRoute
}
var file_models_model_ns_record_proto_depIdxs = []int32{
	1, // 0: pb.NSRecord.nsDomain:type_name -> pb.NSDomain
	2, // 1: pb.NSRecord.nsRoutes:type_name -> pb.NSRoute
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_models_model_ns_record_proto_init() }
func file_models_model_ns_record_proto_init() {
	if File_models_model_ns_record_proto != nil {
		return
	}
	file_models_model_ns_domain_proto_init()
	file_models_model_ns_route_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_models_model_ns_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NSRecord); i {
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
			RawDescriptor: file_models_model_ns_record_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_ns_record_proto_goTypes,
		DependencyIndexes: file_models_model_ns_record_proto_depIdxs,
		MessageInfos:      file_models_model_ns_record_proto_msgTypes,
	}.Build()
	File_models_model_ns_record_proto = out.File
	file_models_model_ns_record_proto_rawDesc = nil
	file_models_model_ns_record_proto_goTypes = nil
	file_models_model_ns_record_proto_depIdxs = nil
}
