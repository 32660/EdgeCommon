// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.4
// source: models/model_ns_cluster.proto

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

// 域名服务集群
type NSCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsOn            bool     `protobuf:"varint,2,opt,name=isOn,proto3" json:"isOn,omitempty"`
	Name            string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	InstallDir      string   `protobuf:"bytes,4,opt,name=installDir,proto3" json:"installDir,omitempty"`
	TcpJSON         []byte   `protobuf:"bytes,5,opt,name=tcpJSON,proto3" json:"tcpJSON,omitempty"`  // TCP设置
	TlsJSON         []byte   `protobuf:"bytes,6,opt,name=tlsJSON,proto3" json:"tlsJSON,omitempty"`  // TLS设置
	UdpJSON         []byte   `protobuf:"bytes,7,opt,name=udpJSON,proto3" json:"udpJSON,omitempty"`  // UDP设置
	DohJSON         []byte   `protobuf:"bytes,16,opt,name=dohJSON,proto3" json:"dohJSON,omitempty"` // DoH设置
	Hosts           []string `protobuf:"bytes,8,rep,name=hosts,proto3" json:"hosts,omitempty"`
	SoaJSON         []byte   `protobuf:"bytes,12,opt,name=soaJSON,proto3" json:"soaJSON,omitempty"`
	Email           string   `protobuf:"bytes,13,opt,name=email,proto3" json:"email,omitempty"`
	AutoRemoteStart bool     `protobuf:"varint,9,opt,name=autoRemoteStart,proto3" json:"autoRemoteStart,omitempty"`
	TimeZone        string   `protobuf:"bytes,10,opt,name=timeZone,proto3" json:"timeZone,omitempty"`
	AnswerJSON      []byte   `protobuf:"bytes,11,opt,name=answerJSON,proto3" json:"answerJSON,omitempty"`
	DetectAgents    bool     `protobuf:"varint,14,opt,name=detectAgents,proto3" json:"detectAgents,omitempty"`
	CheckingPorts   bool     `protobuf:"varint,15,opt,name=checkingPorts,proto3" json:"checkingPorts,omitempty"` // 检查端口连通性
}

func (x *NSCluster) Reset() {
	*x = NSCluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_ns_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NSCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NSCluster) ProtoMessage() {}

func (x *NSCluster) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_ns_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NSCluster.ProtoReflect.Descriptor instead.
func (*NSCluster) Descriptor() ([]byte, []int) {
	return file_models_model_ns_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *NSCluster) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NSCluster) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *NSCluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NSCluster) GetInstallDir() string {
	if x != nil {
		return x.InstallDir
	}
	return ""
}

func (x *NSCluster) GetTcpJSON() []byte {
	if x != nil {
		return x.TcpJSON
	}
	return nil
}

func (x *NSCluster) GetTlsJSON() []byte {
	if x != nil {
		return x.TlsJSON
	}
	return nil
}

func (x *NSCluster) GetUdpJSON() []byte {
	if x != nil {
		return x.UdpJSON
	}
	return nil
}

func (x *NSCluster) GetDohJSON() []byte {
	if x != nil {
		return x.DohJSON
	}
	return nil
}

func (x *NSCluster) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *NSCluster) GetSoaJSON() []byte {
	if x != nil {
		return x.SoaJSON
	}
	return nil
}

func (x *NSCluster) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *NSCluster) GetAutoRemoteStart() bool {
	if x != nil {
		return x.AutoRemoteStart
	}
	return false
}

func (x *NSCluster) GetTimeZone() string {
	if x != nil {
		return x.TimeZone
	}
	return ""
}

func (x *NSCluster) GetAnswerJSON() []byte {
	if x != nil {
		return x.AnswerJSON
	}
	return nil
}

func (x *NSCluster) GetDetectAgents() bool {
	if x != nil {
		return x.DetectAgents
	}
	return false
}

func (x *NSCluster) GetCheckingPorts() bool {
	if x != nil {
		return x.CheckingPorts
	}
	return false
}

var File_models_model_ns_cluster_proto protoreflect.FileDescriptor

var file_models_model_ns_cluster_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e,
	0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0xc1, 0x03, 0x0a, 0x09, 0x4e, 0x53, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x69, 0x73, 0x4f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6c, 0x6c, 0x44, 0x69, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x44, 0x69, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x63, 0x70,
	0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x63, 0x70, 0x4a,
	0x53, 0x4f, 0x4e, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x6c, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x6c, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x18, 0x0a,
	0x07, 0x75, 0x64, 0x70, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x75, 0x64, 0x70, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6f, 0x68, 0x4a, 0x53,
	0x4f, 0x4e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x64, 0x6f, 0x68, 0x4a, 0x53, 0x4f,
	0x4e, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f, 0x61, 0x4a, 0x53,
	0x4f, 0x4e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x73, 0x6f, 0x61, 0x4a, 0x53, 0x4f,
	0x4e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0a, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x22, 0x0a,
	0x0c, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72,
	0x74, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_ns_cluster_proto_rawDescOnce sync.Once
	file_models_model_ns_cluster_proto_rawDescData = file_models_model_ns_cluster_proto_rawDesc
)

func file_models_model_ns_cluster_proto_rawDescGZIP() []byte {
	file_models_model_ns_cluster_proto_rawDescOnce.Do(func() {
		file_models_model_ns_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_ns_cluster_proto_rawDescData)
	})
	return file_models_model_ns_cluster_proto_rawDescData
}

var file_models_model_ns_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_ns_cluster_proto_goTypes = []interface{}{
	(*NSCluster)(nil), // 0: pb.NSCluster
}
var file_models_model_ns_cluster_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_ns_cluster_proto_init() }
func file_models_model_ns_cluster_proto_init() {
	if File_models_model_ns_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_ns_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NSCluster); i {
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
			RawDescriptor: file_models_model_ns_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_ns_cluster_proto_goTypes,
		DependencyIndexes: file_models_model_ns_cluster_proto_depIdxs,
		MessageInfos:      file_models_model_ns_cluster_proto_msgTypes,
	}.Build()
	File_models_model_ns_cluster_proto = out.File
	file_models_model_ns_cluster_proto_rawDesc = nil
	file_models_model_ns_cluster_proto_goTypes = nil
	file_models_model_ns_cluster_proto_depIdxs = nil
}
