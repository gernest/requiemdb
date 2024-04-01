// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: rq/v1/sample.proto

package v1

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

type PREFIX int32

const (
	PREFIX_RESOURCE_SCHEMA     PREFIX = 0
	PREFIX_RESOURCE_ATTRIBUTES PREFIX = 1
	PREFIX_SCOPE_SCHEMA        PREFIX = 2
	PREFIX_SCOPE_NAME          PREFIX = 3
	PREFIX_SCOPE_VERSION       PREFIX = 4
	PREFIX_SCOPE_ATTRIBUTES    PREFIX = 5
	PREFIX_NAME                PREFIX = 6
	PREFIX_ATTRIBUTES          PREFIX = 7
	PREFIX_TRACE_ID            PREFIX = 8
	PREFIX_SPAN_ID             PREFIX = 9
	PREFIX_PARENT_SPAN_ID      PREFIX = 10
	PREFIX_LOGS_LEVEL          PREFIX = 11
)

// Enum value maps for PREFIX.
var (
	PREFIX_name = map[int32]string{
		0:  "RESOURCE_SCHEMA",
		1:  "RESOURCE_ATTRIBUTES",
		2:  "SCOPE_SCHEMA",
		3:  "SCOPE_NAME",
		4:  "SCOPE_VERSION",
		5:  "SCOPE_ATTRIBUTES",
		6:  "NAME",
		7:  "ATTRIBUTES",
		8:  "TRACE_ID",
		9:  "SPAN_ID",
		10: "PARENT_SPAN_ID",
		11: "LOGS_LEVEL",
	}
	PREFIX_value = map[string]int32{
		"RESOURCE_SCHEMA":     0,
		"RESOURCE_ATTRIBUTES": 1,
		"SCOPE_SCHEMA":        2,
		"SCOPE_NAME":          3,
		"SCOPE_VERSION":       4,
		"SCOPE_ATTRIBUTES":    5,
		"NAME":                6,
		"ATTRIBUTES":          7,
		"TRACE_ID":            8,
		"SPAN_ID":             9,
		"PARENT_SPAN_ID":      10,
		"LOGS_LEVEL":          11,
	}
)

func (x PREFIX) Enum() *PREFIX {
	p := new(PREFIX)
	*p = x
	return p
}

func (x PREFIX) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PREFIX) Descriptor() protoreflect.EnumDescriptor {
	return file_rq_v1_sample_proto_enumTypes[0].Descriptor()
}

func (PREFIX) Type() protoreflect.EnumType {
	return &file_rq_v1_sample_proto_enumTypes[0]
}

func (x PREFIX) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PREFIX.Descriptor instead.
func (PREFIX) EnumDescriptor() ([]byte, []int) {
	return file_rq_v1_sample_proto_rawDescGZIP(), []int{0}
}

type Sample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Serialized Data object, compressed with zstd. We use bytes here because we
	// automatically sore Sample as a arrow.Record.
	Data *Data `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// Minimum timetamp observed in this sample in milliseconds
	MinTs uint64 `protobuf:"varint,3,opt,name=min_ts,json=minTs,proto3" json:"min_ts,omitempty"`
	// Maximum timestamp observed in this sample in milliseconds
	MaxTs uint64 `protobuf:"varint,4,opt,name=max_ts,json=maxTs,proto3" json:"max_ts,omitempty"`
}

func (x *Sample) Reset() {
	*x = Sample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rq_v1_sample_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sample) ProtoMessage() {}

func (x *Sample) ProtoReflect() protoreflect.Message {
	mi := &file_rq_v1_sample_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sample.ProtoReflect.Descriptor instead.
func (*Sample) Descriptor() ([]byte, []int) {
	return file_rq_v1_sample_proto_rawDescGZIP(), []int{0}
}

func (x *Sample) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Sample) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Sample) GetMinTs() uint64 {
	if x != nil {
		return x.MinTs
	}
	return 0
}

func (x *Sample) GetMaxTs() uint64 {
	if x != nil {
		return x.MaxTs
	}
	return 0
}

// Meta stores sample metaddata.
type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Minimum timetamp observed in this sample in milliseconds
	MinTs uint64 `protobuf:"varint,2,opt,name=min_ts,json=minTs,proto3" json:"min_ts,omitempty"`
	// Maximum timestamp observed in this sample in milliseconds
	MaxTs    uint64 `protobuf:"varint,3,opt,name=max_ts,json=maxTs,proto3" json:"max_ts,omitempty"`
	Resource uint64 `protobuf:"varint,4,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rq_v1_sample_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_rq_v1_sample_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_rq_v1_sample_proto_rawDescGZIP(), []int{1}
}

func (x *Meta) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Meta) GetMinTs() uint64 {
	if x != nil {
		return x.MinTs
	}
	return 0
}

func (x *Meta) GetMaxTs() uint64 {
	if x != nil {
		return x.MaxTs
	}
	return 0
}

func (x *Meta) GetResource() uint64 {
	if x != nil {
		return x.Resource
	}
	return 0
}

var File_rq_v1_sample_proto protoreflect.FileDescriptor

var file_rq_v1_sample_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x71, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x10, 0x72, 0x71, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x63, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x06, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x69, 0x6e, 0x5f, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x54, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x61, 0x78,
	0x5f, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x61, 0x78, 0x54, 0x73,
	0x22, 0x60, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x69, 0x6e, 0x5f,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x54, 0x73, 0x12,
	0x15, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x6d, 0x61, 0x78, 0x54, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2a, 0xda, 0x01, 0x0a, 0x06, 0x50, 0x52, 0x45, 0x46, 0x49, 0x58, 0x12, 0x13, 0x0a,
	0x0f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41,
	0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x41,
	0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53,
	0x43, 0x4f, 0x50, 0x45, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x10, 0x02, 0x12, 0x0e, 0x0a,
	0x0a, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x03, 0x12, 0x11, 0x0a,
	0x0d, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x04,
	0x12, 0x14, 0x0a, 0x10, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42,
	0x55, 0x54, 0x45, 0x53, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x06,
	0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x53, 0x10, 0x07,
	0x12, 0x0c, 0x0a, 0x08, 0x54, 0x52, 0x41, 0x43, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x08, 0x12, 0x0b,
	0x0a, 0x07, 0x53, 0x50, 0x41, 0x4e, 0x5f, 0x49, 0x44, 0x10, 0x09, 0x12, 0x12, 0x0a, 0x0e, 0x50,
	0x41, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x50, 0x41, 0x4e, 0x5f, 0x49, 0x44, 0x10, 0x0a, 0x12,
	0x0e, 0x0a, 0x0a, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x10, 0x0b, 0x42,
	0x6a, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x65, 0x6d, 0x64, 0x62, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x65, 0x6d, 0x64, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f,
	0x72, 0x71, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31,
	0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rq_v1_sample_proto_rawDescOnce sync.Once
	file_rq_v1_sample_proto_rawDescData = file_rq_v1_sample_proto_rawDesc
)

func file_rq_v1_sample_proto_rawDescGZIP() []byte {
	file_rq_v1_sample_proto_rawDescOnce.Do(func() {
		file_rq_v1_sample_proto_rawDescData = protoimpl.X.CompressGZIP(file_rq_v1_sample_proto_rawDescData)
	})
	return file_rq_v1_sample_proto_rawDescData
}

var file_rq_v1_sample_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rq_v1_sample_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rq_v1_sample_proto_goTypes = []interface{}{
	(PREFIX)(0),    // 0: v1.PREFIX
	(*Sample)(nil), // 1: v1.Sample
	(*Meta)(nil),   // 2: v1.Meta
	(*Data)(nil),   // 3: v1.Data
}
var file_rq_v1_sample_proto_depIdxs = []int32{
	3, // 0: v1.Sample.data:type_name -> v1.Data
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rq_v1_sample_proto_init() }
func file_rq_v1_sample_proto_init() {
	if File_rq_v1_sample_proto != nil {
		return
	}
	file_rq_v1_scan_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rq_v1_sample_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sample); i {
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
		file_rq_v1_sample_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
			RawDescriptor: file_rq_v1_sample_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rq_v1_sample_proto_goTypes,
		DependencyIndexes: file_rq_v1_sample_proto_depIdxs,
		EnumInfos:         file_rq_v1_sample_proto_enumTypes,
		MessageInfos:      file_rq_v1_sample_proto_msgTypes,
	}.Build()
	File_rq_v1_sample_proto = out.File
	file_rq_v1_sample_proto_rawDesc = nil
	file_rq_v1_sample_proto_goTypes = nil
	file_rq_v1_sample_proto_depIdxs = nil
}
