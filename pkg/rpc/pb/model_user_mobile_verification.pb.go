// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: models/model_user_mobile_verification.proto

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

// 手机号码认证
type UserMobileVerification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                 // ID
	Mobile     string `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`          // 手机号码
	UserId     int64  `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`         // 用户ID
	Code       string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`              // 代号
	CreatedAt  int64  `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`   // 创建时间
	IsSent     bool   `protobuf:"varint,6,opt,name=isSent,proto3" json:"isSent,omitempty"`         // 已发送
	IsVerified bool   `protobuf:"varint,7,opt,name=isVerified,proto3" json:"isVerified,omitempty"` // 已激活
	ExpiresAt  int64  `protobuf:"varint,8,opt,name=expiresAt,proto3" json:"expiresAt,omitempty"`   // 过期时间，动态计算而来
}

func (x *UserMobileVerification) Reset() {
	*x = UserMobileVerification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_user_mobile_verification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMobileVerification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMobileVerification) ProtoMessage() {}

func (x *UserMobileVerification) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_user_mobile_verification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMobileVerification.ProtoReflect.Descriptor instead.
func (*UserMobileVerification) Descriptor() ([]byte, []int) {
	return file_models_model_user_mobile_verification_proto_rawDescGZIP(), []int{0}
}

func (x *UserMobileVerification) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserMobileVerification) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *UserMobileVerification) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserMobileVerification) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UserMobileVerification) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *UserMobileVerification) GetIsSent() bool {
	if x != nil {
		return x.IsSent
	}
	return false
}

func (x *UserMobileVerification) GetIsVerified() bool {
	if x != nil {
		return x.IsVerified
	}
	return false
}

func (x *UserMobileVerification) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

var File_models_model_user_mobile_verification_proto protoreflect.FileDescriptor

var file_models_model_user_mobile_verification_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0xe0, 0x01, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x69, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x73, 0x41, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_user_mobile_verification_proto_rawDescOnce sync.Once
	file_models_model_user_mobile_verification_proto_rawDescData = file_models_model_user_mobile_verification_proto_rawDesc
)

func file_models_model_user_mobile_verification_proto_rawDescGZIP() []byte {
	file_models_model_user_mobile_verification_proto_rawDescOnce.Do(func() {
		file_models_model_user_mobile_verification_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_user_mobile_verification_proto_rawDescData)
	})
	return file_models_model_user_mobile_verification_proto_rawDescData
}

var file_models_model_user_mobile_verification_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_user_mobile_verification_proto_goTypes = []interface{}{
	(*UserMobileVerification)(nil), // 0: pb.UserMobileVerification
}
var file_models_model_user_mobile_verification_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_user_mobile_verification_proto_init() }
func file_models_model_user_mobile_verification_proto_init() {
	if File_models_model_user_mobile_verification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_user_mobile_verification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMobileVerification); i {
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
			RawDescriptor: file_models_model_user_mobile_verification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_user_mobile_verification_proto_goTypes,
		DependencyIndexes: file_models_model_user_mobile_verification_proto_depIdxs,
		MessageInfos:      file_models_model_user_mobile_verification_proto_msgTypes,
	}.Build()
	File_models_model_user_mobile_verification_proto = out.File
	file_models_model_user_mobile_verification_proto_rawDesc = nil
	file_models_model_user_mobile_verification_proto_goTypes = nil
	file_models_model_user_mobile_verification_proto_depIdxs = nil
}
