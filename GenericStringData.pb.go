// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: GenericStringData.proto

package robot_data

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

type GenericStringData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/////////////////////////////////////// add by robot-data service ///////////////////////
	// @inject_tag: json:"timestamp"
	Publishtime string `protobuf:"bytes,1,opt,name=publishtime,proto3" json:"timestamp"`
	// @inject_tag: json:"-"
	Identity string `protobuf:"bytes,2,opt,name=identity,proto3" json:"-"`
	// @inject_tag: json:"-"
	Groups []string `protobuf:"bytes,3,rep,name=groups,proto3" json:"-"`
	// @inject_tag: json:"-"
	Paths []string `protobuf:"bytes,4,rep,name=paths,proto3" json:"-"`
	/////////////////////////////////////// add SpoonExportableData ///////////////////////
	// @inject_tag: json:"DataType"
	DataType string `protobuf:"bytes,5,opt,name=DataType,proto3" json:"DataType"`
	Data     string `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GenericStringData) Reset() {
	*x = GenericStringData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GenericStringData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericStringData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericStringData) ProtoMessage() {}

func (x *GenericStringData) ProtoReflect() protoreflect.Message {
	mi := &file_GenericStringData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericStringData.ProtoReflect.Descriptor instead.
func (*GenericStringData) Descriptor() ([]byte, []int) {
	return file_GenericStringData_proto_rawDescGZIP(), []int{0}
}

func (x *GenericStringData) GetPublishtime() string {
	if x != nil {
		return x.Publishtime
	}
	return ""
}

func (x *GenericStringData) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *GenericStringData) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *GenericStringData) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *GenericStringData) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *GenericStringData) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_GenericStringData_proto protoreflect.FileDescriptor

var file_GenericStringData_proto_rawDesc = []byte{
	0x0a, 0x17, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x0e, 0x62, 0x71, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x62, 0x71, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69,
	0x63, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x0b, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x10, 0xea, 0x3f, 0x0d, 0x08, 0x01, 0x12, 0x09, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41,
	0x4d, 0x50, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x05, 0xea, 0x3f, 0x02, 0x08, 0x01, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x74, 0x68, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73,
	0x12, 0x21, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x05, 0xea, 0x3f, 0x02, 0x08, 0x01, 0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x3a, 0x16, 0xea, 0x3f, 0x13, 0x0a, 0x11, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x69, 0x63, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x42,
	0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6c,
	0x61, 0x76, 0x61, 0x79, 0x73, 0x73, 0x69, 0x65, 0x72, 0x65, 0x2d, 0x73, 0x70, 0x6f, 0x6f, 0x6e,
	0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_GenericStringData_proto_rawDescOnce sync.Once
	file_GenericStringData_proto_rawDescData = file_GenericStringData_proto_rawDesc
)

func file_GenericStringData_proto_rawDescGZIP() []byte {
	file_GenericStringData_proto_rawDescOnce.Do(func() {
		file_GenericStringData_proto_rawDescData = protoimpl.X.CompressGZIP(file_GenericStringData_proto_rawDescData)
	})
	return file_GenericStringData_proto_rawDescData
}

var file_GenericStringData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GenericStringData_proto_goTypes = []interface{}{
	(*GenericStringData)(nil), // 0: robot_data.GenericStringData
}
var file_GenericStringData_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GenericStringData_proto_init() }
func file_GenericStringData_proto_init() {
	if File_GenericStringData_proto != nil {
		return
	}
	file_bq_table_proto_init()
	file_bq_field_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GenericStringData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericStringData); i {
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
			RawDescriptor: file_GenericStringData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GenericStringData_proto_goTypes,
		DependencyIndexes: file_GenericStringData_proto_depIdxs,
		MessageInfos:      file_GenericStringData_proto_msgTypes,
	}.Build()
	File_GenericStringData_proto = out.File
	file_GenericStringData_proto_rawDesc = nil
	file_GenericStringData_proto_goTypes = nil
	file_GenericStringData_proto_depIdxs = nil
}
