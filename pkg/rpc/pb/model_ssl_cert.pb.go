// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: models/model_ssl_cert.proto

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

type SSLCert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsOn          bool     `protobuf:"varint,2,opt,name=isOn,proto3" json:"isOn,omitempty"`
	Name          string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	TimeBeginAt   int64    `protobuf:"varint,4,opt,name=timeBeginAt,proto3" json:"timeBeginAt,omitempty"`
	TimeEndAt     int64    `protobuf:"varint,5,opt,name=timeEndAt,proto3" json:"timeEndAt,omitempty"`
	DnsNames      []string `protobuf:"bytes,6,rep,name=dnsNames,proto3" json:"dnsNames,omitempty"`
	CommonNames   []string `protobuf:"bytes,7,rep,name=commonNames,proto3" json:"commonNames,omitempty"`
	IsACME        bool     `protobuf:"varint,8,opt,name=isACME,proto3" json:"isACME,omitempty"`
	AcmeTaskId    int64    `protobuf:"varint,17,opt,name=acmeTaskId,proto3" json:"acmeTaskId,omitempty"`
	Ocsp          []byte   `protobuf:"bytes,9,opt,name=ocsp,proto3" json:"ocsp,omitempty"`
	OcspIsUpdated bool     `protobuf:"varint,10,opt,name=ocspIsUpdated,proto3" json:"ocspIsUpdated,omitempty"`
	OcspError     string   `protobuf:"bytes,11,opt,name=ocspError,proto3" json:"ocspError,omitempty"`
	Description   string   `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	IsCA          bool     `protobuf:"varint,13,opt,name=isCA,proto3" json:"isCA,omitempty"`
	ServerName    string   `protobuf:"bytes,14,opt,name=serverName,proto3" json:"serverName,omitempty"`
	CreatedAt     int64    `protobuf:"varint,15,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt     int64    `protobuf:"varint,16,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *SSLCert) Reset() {
	*x = SSLCert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_ssl_cert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSLCert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSLCert) ProtoMessage() {}

func (x *SSLCert) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_ssl_cert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSLCert.ProtoReflect.Descriptor instead.
func (*SSLCert) Descriptor() ([]byte, []int) {
	return file_models_model_ssl_cert_proto_rawDescGZIP(), []int{0}
}

func (x *SSLCert) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SSLCert) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *SSLCert) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SSLCert) GetTimeBeginAt() int64 {
	if x != nil {
		return x.TimeBeginAt
	}
	return 0
}

func (x *SSLCert) GetTimeEndAt() int64 {
	if x != nil {
		return x.TimeEndAt
	}
	return 0
}

func (x *SSLCert) GetDnsNames() []string {
	if x != nil {
		return x.DnsNames
	}
	return nil
}

func (x *SSLCert) GetCommonNames() []string {
	if x != nil {
		return x.CommonNames
	}
	return nil
}

func (x *SSLCert) GetIsACME() bool {
	if x != nil {
		return x.IsACME
	}
	return false
}

func (x *SSLCert) GetAcmeTaskId() int64 {
	if x != nil {
		return x.AcmeTaskId
	}
	return 0
}

func (x *SSLCert) GetOcsp() []byte {
	if x != nil {
		return x.Ocsp
	}
	return nil
}

func (x *SSLCert) GetOcspIsUpdated() bool {
	if x != nil {
		return x.OcspIsUpdated
	}
	return false
}

func (x *SSLCert) GetOcspError() string {
	if x != nil {
		return x.OcspError
	}
	return ""
}

func (x *SSLCert) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SSLCert) GetIsCA() bool {
	if x != nil {
		return x.IsCA
	}
	return false
}

func (x *SSLCert) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *SSLCert) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *SSLCert) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_models_model_ssl_cert_proto protoreflect.FileDescriptor

var file_models_model_ssl_cert_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73,
	0x73, 0x6c, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0xe1, 0x03, 0x0a, 0x07, 0x53, 0x53, 0x4c, 0x43, 0x65, 0x72, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x65, 0x67,
	0x69, 0x6e, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65,
	0x42, 0x65, 0x67, 0x69, 0x6e, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x45,
	0x6e, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x45, 0x6e, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x6e, 0x73, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x64, 0x6e, 0x73, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x41, 0x43, 0x4d, 0x45, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x41, 0x43, 0x4d, 0x45, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x63, 0x6d, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x61, 0x63, 0x6d, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6f,
	0x63, 0x73, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6f, 0x63, 0x73, 0x70, 0x12,
	0x24, 0x0a, 0x0d, 0x6f, 0x63, 0x73, 0x70, 0x49, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6f, 0x63, 0x73, 0x70, 0x49, 0x73, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x63, 0x73, 0x70, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x63, 0x73, 0x70, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73, 0x43, 0x41, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x43, 0x41, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_ssl_cert_proto_rawDescOnce sync.Once
	file_models_model_ssl_cert_proto_rawDescData = file_models_model_ssl_cert_proto_rawDesc
)

func file_models_model_ssl_cert_proto_rawDescGZIP() []byte {
	file_models_model_ssl_cert_proto_rawDescOnce.Do(func() {
		file_models_model_ssl_cert_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_ssl_cert_proto_rawDescData)
	})
	return file_models_model_ssl_cert_proto_rawDescData
}

var file_models_model_ssl_cert_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_ssl_cert_proto_goTypes = []interface{}{
	(*SSLCert)(nil), // 0: pb.SSLCert
}
var file_models_model_ssl_cert_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_ssl_cert_proto_init() }
func file_models_model_ssl_cert_proto_init() {
	if File_models_model_ssl_cert_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_ssl_cert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSLCert); i {
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
			RawDescriptor: file_models_model_ssl_cert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_ssl_cert_proto_goTypes,
		DependencyIndexes: file_models_model_ssl_cert_proto_depIdxs,
		MessageInfos:      file_models_model_ssl_cert_proto_msgTypes,
	}.Build()
	File_models_model_ssl_cert_proto = out.File
	file_models_model_ssl_cert_proto_rawDesc = nil
	file_models_model_ssl_cert_proto_goTypes = nil
	file_models_model_ssl_cert_proto_depIdxs = nil
}
