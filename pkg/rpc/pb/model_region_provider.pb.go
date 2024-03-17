// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: models/model_region_provider.proto

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

type RegionProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Codes       []string `protobuf:"bytes,3,rep,name=codes,proto3" json:"codes,omitempty"`
	CustomName  string   `protobuf:"bytes,4,opt,name=customName,proto3" json:"customName,omitempty"`
	CustomCodes []string `protobuf:"bytes,5,rep,name=customCodes,proto3" json:"customCodes,omitempty"`
	DisplayName string   `protobuf:"bytes,6,opt,name=displayName,proto3" json:"displayName,omitempty"`
}

func (x *RegionProvider) Reset() {
	*x = RegionProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_region_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionProvider) ProtoMessage() {}

func (x *RegionProvider) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_region_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionProvider.ProtoReflect.Descriptor instead.
func (*RegionProvider) Descriptor() ([]byte, []int) {
	return file_models_model_region_provider_proto_rawDescGZIP(), []int{0}
}

func (x *RegionProvider) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RegionProvider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegionProvider) GetCodes() []string {
	if x != nil {
		return x.Codes
	}
	return nil
}

func (x *RegionProvider) GetCustomName() string {
	if x != nil {
		return x.CustomName
	}
	return ""
}

func (x *RegionProvider) GetCustomCodes() []string {
	if x != nil {
		return x.CustomCodes
	}
	return nil
}

func (x *RegionProvider) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

var File_models_model_region_provider_proto protoreflect.FileDescriptor

var file_models_model_region_provider_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xae, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x63, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43,
	0x6f, 0x64, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_region_provider_proto_rawDescOnce sync.Once
	file_models_model_region_provider_proto_rawDescData = file_models_model_region_provider_proto_rawDesc
)

func file_models_model_region_provider_proto_rawDescGZIP() []byte {
	file_models_model_region_provider_proto_rawDescOnce.Do(func() {
		file_models_model_region_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_region_provider_proto_rawDescData)
	})
	return file_models_model_region_provider_proto_rawDescData
}

var file_models_model_region_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_region_provider_proto_goTypes = []interface{}{
	(*RegionProvider)(nil), // 0: pb.RegionProvider
}
var file_models_model_region_provider_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_region_provider_proto_init() }
func file_models_model_region_provider_proto_init() {
	if File_models_model_region_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_region_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionProvider); i {
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
			RawDescriptor: file_models_model_region_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_region_provider_proto_goTypes,
		DependencyIndexes: file_models_model_region_provider_proto_depIdxs,
		MessageInfos:      file_models_model_region_provider_proto_msgTypes,
	}.Build()
	File_models_model_region_provider_proto = out.File
	file_models_model_region_provider_proto_rawDesc = nil
	file_models_model_region_provider_proto_goTypes = nil
	file_models_model_region_provider_proto_depIdxs = nil
}
