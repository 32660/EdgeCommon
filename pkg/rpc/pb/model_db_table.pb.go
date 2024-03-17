// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: models/model_db_table.proto

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

// 数据表信息
type DBTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Schema      string `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	Type        string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Engine      string `protobuf:"bytes,4,opt,name=engine,proto3" json:"engine,omitempty"`
	Rows        int64  `protobuf:"varint,5,opt,name=rows,proto3" json:"rows,omitempty"`
	DataLength  int64  `protobuf:"varint,6,opt,name=dataLength,proto3" json:"dataLength,omitempty"`
	IndexLength int64  `protobuf:"varint,7,opt,name=indexLength,proto3" json:"indexLength,omitempty"`
	Comment     string `protobuf:"bytes,8,opt,name=comment,proto3" json:"comment,omitempty"`
	Collation   string `protobuf:"bytes,9,opt,name=collation,proto3" json:"collation,omitempty"`
	IsBaseTable bool   `protobuf:"varint,10,opt,name=isBaseTable,proto3" json:"isBaseTable,omitempty"`
	CanClean    bool   `protobuf:"varint,11,opt,name=canClean,proto3" json:"canClean,omitempty"`
	CanDelete   bool   `protobuf:"varint,12,opt,name=canDelete,proto3" json:"canDelete,omitempty"`
}

func (x *DBTable) Reset() {
	*x = DBTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_model_db_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBTable) ProtoMessage() {}

func (x *DBTable) ProtoReflect() protoreflect.Message {
	mi := &file_models_model_db_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBTable.ProtoReflect.Descriptor instead.
func (*DBTable) Descriptor() ([]byte, []int) {
	return file_models_model_db_table_proto_rawDescGZIP(), []int{0}
}

func (x *DBTable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DBTable) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *DBTable) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DBTable) GetEngine() string {
	if x != nil {
		return x.Engine
	}
	return ""
}

func (x *DBTable) GetRows() int64 {
	if x != nil {
		return x.Rows
	}
	return 0
}

func (x *DBTable) GetDataLength() int64 {
	if x != nil {
		return x.DataLength
	}
	return 0
}

func (x *DBTable) GetIndexLength() int64 {
	if x != nil {
		return x.IndexLength
	}
	return 0
}

func (x *DBTable) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *DBTable) GetCollation() string {
	if x != nil {
		return x.Collation
	}
	return ""
}

func (x *DBTable) GetIsBaseTable() bool {
	if x != nil {
		return x.IsBaseTable
	}
	return false
}

func (x *DBTable) GetCanClean() bool {
	if x != nil {
		return x.CanClean
	}
	return false
}

func (x *DBTable) GetCanDelete() bool {
	if x != nil {
		return x.CanDelete
	}
	return false
}

var File_models_model_db_table_proto protoreflect.FileDescriptor

var file_models_model_db_table_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x64,
	0x62, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0xcb, 0x02, 0x0a, 0x07, 0x44, 0x42, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x74,
	0x61, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64,
	0x61, 0x74, 0x61, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x42, 0x61, 0x73, 0x65, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x42, 0x61, 0x73, 0x65,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6e, 0x43, 0x6c, 0x65, 0x61,
	0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x61, 0x6e, 0x43, 0x6c, 0x65, 0x61,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x61, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_model_db_table_proto_rawDescOnce sync.Once
	file_models_model_db_table_proto_rawDescData = file_models_model_db_table_proto_rawDesc
)

func file_models_model_db_table_proto_rawDescGZIP() []byte {
	file_models_model_db_table_proto_rawDescOnce.Do(func() {
		file_models_model_db_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_model_db_table_proto_rawDescData)
	})
	return file_models_model_db_table_proto_rawDescData
}

var file_models_model_db_table_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_models_model_db_table_proto_goTypes = []interface{}{
	(*DBTable)(nil), // 0: pb.DBTable
}
var file_models_model_db_table_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_models_model_db_table_proto_init() }
func file_models_model_db_table_proto_init() {
	if File_models_model_db_table_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_model_db_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBTable); i {
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
			RawDescriptor: file_models_model_db_table_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_models_model_db_table_proto_goTypes,
		DependencyIndexes: file_models_model_db_table_proto_depIdxs,
		MessageInfos:      file_models_model_db_table_proto_msgTypes,
	}.Build()
	File_models_model_db_table_proto = out.File
	file_models_model_db_table_proto_rawDesc = nil
	file_models_model_db_table_proto_goTypes = nil
	file_models_model_db_table_proto_depIdxs = nil
}
