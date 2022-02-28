// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: internal/pb/protobuf-spec/data_type.proto

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

type DataType int32

const (
	DataType_STRING        DataType = 0
	DataType_LIST          DataType = 1
	DataType_SET           DataType = 2
	DataType_SORTED_SET    DataType = 3
	DataType_BLOOM_FILTER  DataType = 4
	DataType_HYPER_LOG_LOG DataType = 5
	DataType_BITSET        DataType = 6
	DataType_MAP           DataType = 7
	DataType_GEO_HASH      DataType = 8
)

// Enum value maps for DataType.
var (
	DataType_name = map[int32]string{
		0: "STRING",
		1: "LIST",
		2: "SET",
		3: "SORTED_SET",
		4: "BLOOM_FILTER",
		5: "HYPER_LOG_LOG",
		6: "BITSET",
		7: "MAP",
		8: "GEO_HASH",
	}
	DataType_value = map[string]int32{
		"STRING":        0,
		"LIST":          1,
		"SET":           2,
		"SORTED_SET":    3,
		"BLOOM_FILTER":  4,
		"HYPER_LOG_LOG": 5,
		"BITSET":        6,
		"MAP":           7,
		"GEO_HASH":      8,
	}
)

func (x DataType) Enum() *DataType {
	p := new(DataType)
	*p = x
	return p
}

func (x DataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataType) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_pb_protobuf_spec_data_type_proto_enumTypes[0].Descriptor()
}

func (DataType) Type() protoreflect.EnumType {
	return &file_internal_pb_protobuf_spec_data_type_proto_enumTypes[0]
}

func (x DataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataType.Descriptor instead.
func (DataType) EnumDescriptor() ([]byte, []int) {
	return file_internal_pb_protobuf_spec_data_type_proto_rawDescGZIP(), []int{0}
}

var File_internal_pb_protobuf_spec_data_type_proto protoreflect.FileDescriptor

var file_internal_pb_protobuf_spec_data_type_proto_rawDesc = []byte{
	0x0a, 0x29, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2a, 0x81, 0x01, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4c,
	0x49, 0x53, 0x54, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x45, 0x54, 0x10, 0x02, 0x12, 0x0e,
	0x0a, 0x0a, 0x53, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x03, 0x12, 0x10,
	0x0a, 0x0c, 0x42, 0x4c, 0x4f, 0x4f, 0x4d, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x10, 0x04,
	0x12, 0x11, 0x0a, 0x0d, 0x48, 0x59, 0x50, 0x45, 0x52, 0x5f, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x4f,
	0x47, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x49, 0x54, 0x53, 0x45, 0x54, 0x10, 0x06, 0x12,
	0x07, 0x0a, 0x03, 0x4d, 0x41, 0x50, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x47, 0x45, 0x4f, 0x5f,
	0x48, 0x41, 0x53, 0x48, 0x10, 0x08, 0x42, 0x0d, 0x5a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pb_protobuf_spec_data_type_proto_rawDescOnce sync.Once
	file_internal_pb_protobuf_spec_data_type_proto_rawDescData = file_internal_pb_protobuf_spec_data_type_proto_rawDesc
)

func file_internal_pb_protobuf_spec_data_type_proto_rawDescGZIP() []byte {
	file_internal_pb_protobuf_spec_data_type_proto_rawDescOnce.Do(func() {
		file_internal_pb_protobuf_spec_data_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pb_protobuf_spec_data_type_proto_rawDescData)
	})
	return file_internal_pb_protobuf_spec_data_type_proto_rawDescData
}

var file_internal_pb_protobuf_spec_data_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_pb_protobuf_spec_data_type_proto_goTypes = []interface{}{
	(DataType)(0), // 0: proto.DataType
}
var file_internal_pb_protobuf_spec_data_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_pb_protobuf_spec_data_type_proto_init() }
func file_internal_pb_protobuf_spec_data_type_proto_init() {
	if File_internal_pb_protobuf_spec_data_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_pb_protobuf_spec_data_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_pb_protobuf_spec_data_type_proto_goTypes,
		DependencyIndexes: file_internal_pb_protobuf_spec_data_type_proto_depIdxs,
		EnumInfos:         file_internal_pb_protobuf_spec_data_type_proto_enumTypes,
	}.Build()
	File_internal_pb_protobuf_spec_data_type_proto = out.File
	file_internal_pb_protobuf_spec_data_type_proto_rawDesc = nil
	file_internal_pb_protobuf_spec_data_type_proto_goTypes = nil
	file_internal_pb_protobuf_spec_data_type_proto_depIdxs = nil
}