// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: models/model_http_firewall_rule_group.proto

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

type HTTPFirewallRuleGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsOn        bool   `protobuf:"varint,3,opt,name=isOn,proto3" json:"isOn,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Code        string `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *HTTPFirewallRuleGroup) Reset() {
	*x = HTTPFirewallRuleGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_http_firewall_rule_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HTTPFirewallRuleGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HTTPFirewallRuleGroup) ProtoMessage() {}

func (x *HTTPFirewallRuleGroup) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_http_firewall_rule_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HTTPFirewallRuleGroup.ProtoReflect.Descriptor instead.
func (*HTTPFirewallRuleGroup) Descriptor() ([]byte, []int) {
	return file_models_model_http_firewall_rule_group_proto_rawDescGZIP(), []int{0}
}

func (x *HTTPFirewallRuleGroup) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HTTPFirewallRuleGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HTTPFirewallRuleGroup) GetIsOn() bool {
	if x != nil {
		return x.IsOn
	}
	return false
}

func (x *HTTPFirewallRuleGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *HTTPFirewallRuleGroup) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_models_model_http_firewall_rule_group_proto protoreflect.FileDescriptor

var file_models_model_http_firewall_rule_group_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x68,
	0x74, 0x74, 0x70, 0x5f, 0x66, 0x69, 0x72, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x85, 0x01, 0x0a, 0x15, 0x48, 0x54, 0x54, 0x50, 0x46, 0x69, 0x72, 0x65, 0x77, 0x61,
	0x6c, 0x6c, 0x52, 0x75, 0x6c, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x69, 0x73, 0x4f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69,
	0x73, 0x4f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_http_firewall_rule_group_proto_rawDescOnce sync.Once
	file_models_model_http_firewall_rule_group_proto_rawDescData = file_models_model_http_firewall_rule_group_proto_rawDesc
)

func file_models_model_http_firewall_rule_group_proto_rawDescGZIP() []byte {
	file_models_model_http_firewall_rule_group_proto_rawDescOnce.Do(func() {
		file_models_model_http_firewall_rule_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_http_firewall_rule_group_proto_rawDescData)
	})
	return file_models_model_http_firewall_rule_group_proto_rawDescData
}

var file_models_model_http_firewall_rule_group_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_http_firewall_rule_group_proto_goTypes = []interface{}{
	(*HTTPFirewallRuleGroup)(nil), // 0: pb.HTTPFirewallRuleGroup
}
var file_models_model_http_firewall_rule_group_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_http_firewall_rule_group_proto_init() }
func file_models_model_http_firewall_rule_group_proto_init() {
	if File_models_model_http_firewall_rule_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_http_firewall_rule_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HTTPFirewallRuleGroup); i {
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
			RawDescriptor: file_models_model_http_firewall_rule_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_http_firewall_rule_group_proto_goTypes,
		DependencyIndexes: file_models_model_http_firewall_rule_group_proto_depIdxs,
		MessageInfos:      file_models_model_http_firewall_rule_group_proto_msgTypes,
	}.Build()
	File_models_model_http_firewall_rule_group_proto = out.File
	file_models_model_http_firewall_rule_group_proto_rawDesc = nil
	file_models_model_http_firewall_rule_group_proto_goTypes = nil
	file_models_model_http_firewall_rule_group_proto_depIdxs = nil
}
