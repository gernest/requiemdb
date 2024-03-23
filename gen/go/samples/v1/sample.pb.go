// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: samples/v1/sample.proto

package v1

import (
	v11 "go.opentelemetry.io/proto/otlp/logs/v1"
	v1 "go.opentelemetry.io/proto/otlp/metrics/v1"
	v12 "go.opentelemetry.io/proto/otlp/trace/v1"
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
	PREFIX_METRICS_NAME        PREFIX = 6
	PREFIX_METRICS_ATTRIBUTES  PREFIX = 7
	PREFIX_SPAN_TRACE_ID       PREFIX = 8
	PREFIX_SPAN_SPAN_ID        PREFIX = 9
	PREFIX_SPAN_NAME           PREFIX = 10
	PREFIX_SPAN_ATTRIBUTES     PREFIX = 11
	PREFIX_SPAN_PARENT_SPAN_ID PREFIX = 13
	PREFIX_LOGS_TRACE_ID       PREFIX = 14
	PREFIX_LOGS_SPAN_ID        PREFIX = 15
	PREFIX_LOGS_LEVEL          PREFIX = 16
	PREFIX_LOGS_ATTRIBUTES     PREFIX = 17
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
		6:  "METRICS_NAME",
		7:  "METRICS_ATTRIBUTES",
		8:  "SPAN_TRACE_ID",
		9:  "SPAN_SPAN_ID",
		10: "SPAN_NAME",
		11: "SPAN_ATTRIBUTES",
		13: "SPAN_PARENT_SPAN_ID",
		14: "LOGS_TRACE_ID",
		15: "LOGS_SPAN_ID",
		16: "LOGS_LEVEL",
		17: "LOGS_ATTRIBUTES",
	}
	PREFIX_value = map[string]int32{
		"RESOURCE_SCHEMA":     0,
		"RESOURCE_ATTRIBUTES": 1,
		"SCOPE_SCHEMA":        2,
		"SCOPE_NAME":          3,
		"SCOPE_VERSION":       4,
		"SCOPE_ATTRIBUTES":    5,
		"METRICS_NAME":        6,
		"METRICS_ATTRIBUTES":  7,
		"SPAN_TRACE_ID":       8,
		"SPAN_SPAN_ID":        9,
		"SPAN_NAME":           10,
		"SPAN_ATTRIBUTES":     11,
		"SPAN_PARENT_SPAN_ID": 13,
		"LOGS_TRACE_ID":       14,
		"LOGS_SPAN_ID":        15,
		"LOGS_LEVEL":          16,
		"LOGS_ATTRIBUTES":     17,
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
	return file_samples_v1_sample_proto_enumTypes[0].Descriptor()
}

func (PREFIX) Type() protoreflect.EnumType {
	return &file_samples_v1_sample_proto_enumTypes[0]
}

func (x PREFIX) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PREFIX.Descriptor instead.
func (PREFIX) EnumDescriptor() ([]byte, []int) {
	return file_samples_v1_sample_proto_rawDescGZIP(), []int{0}
}

type SampleKind int32

const (
	SampleKind_METRICS SampleKind = 0
	SampleKind_TRACES  SampleKind = 2
	SampleKind_LOGS    SampleKind = 3
)

// Enum value maps for SampleKind.
var (
	SampleKind_name = map[int32]string{
		0: "METRICS",
		2: "TRACES",
		3: "LOGS",
	}
	SampleKind_value = map[string]int32{
		"METRICS": 0,
		"TRACES":  2,
		"LOGS":    3,
	}
)

func (x SampleKind) Enum() *SampleKind {
	p := new(SampleKind)
	*p = x
	return p
}

func (x SampleKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SampleKind) Descriptor() protoreflect.EnumDescriptor {
	return file_samples_v1_sample_proto_enumTypes[1].Descriptor()
}

func (SampleKind) Type() protoreflect.EnumType {
	return &file_samples_v1_sample_proto_enumTypes[1]
}

func (x SampleKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SampleKind.Descriptor instead.
func (SampleKind) EnumDescriptor() ([]byte, []int) {
	return file_samples_v1_sample_proto_rawDescGZIP(), []int{1}
}

type Sample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Serialized Data object, compressed with zstd. We use bytes here because we
	// automatically sore Sample as a arrow.Record.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Minimum timetamp observed in this sample in milliseconds
	MinTs uint64 `protobuf:"varint,2,opt,name=min_ts,json=minTs,proto3" json:"min_ts,omitempty"`
	// Maximum timestamp observed in this sample in milliseconds
	MaxTs uint64 `protobuf:"varint,3,opt,name=max_ts,json=maxTs,proto3" json:"max_ts,omitempty"`
	// Date in nillisecond in which the sample was taken
	Date uint64 `protobuf:"varint,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Sample) Reset() {
	*x = Sample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samples_v1_sample_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sample) ProtoMessage() {}

func (x *Sample) ProtoReflect() protoreflect.Message {
	mi := &file_samples_v1_sample_proto_msgTypes[0]
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
	return file_samples_v1_sample_proto_rawDescGZIP(), []int{0}
}

func (x *Sample) GetData() []byte {
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

func (x *Sample) GetDate() uint64 {
	if x != nil {
		return x.Date
	}
	return 0
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*Data_Metrics
	//	*Data_Logs
	//	*Data_Traces
	Data isData_Data `protobuf_oneof:"data"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_samples_v1_sample_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_samples_v1_sample_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_samples_v1_sample_proto_rawDescGZIP(), []int{1}
}

func (m *Data) GetData() isData_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Data) GetMetrics() *v1.MetricsData {
	if x, ok := x.GetData().(*Data_Metrics); ok {
		return x.Metrics
	}
	return nil
}

func (x *Data) GetLogs() *v11.LogsData {
	if x, ok := x.GetData().(*Data_Logs); ok {
		return x.Logs
	}
	return nil
}

func (x *Data) GetTraces() *v12.TracesData {
	if x, ok := x.GetData().(*Data_Traces); ok {
		return x.Traces
	}
	return nil
}

type isData_Data interface {
	isData_Data()
}

type Data_Metrics struct {
	Metrics *v1.MetricsData `protobuf:"bytes,1,opt,name=metrics,proto3,oneof"`
}

type Data_Logs struct {
	Logs *v11.LogsData `protobuf:"bytes,2,opt,name=logs,proto3,oneof"`
}

type Data_Traces struct {
	Traces *v12.TracesData `protobuf:"bytes,3,opt,name=traces,proto3,oneof"`
}

func (*Data_Metrics) isData_Data() {}

func (*Data_Logs) isData_Data() {}

func (*Data_Traces) isData_Data() {}

var File_samples_v1_sample_proto protoreflect.FileDescriptor

var file_samples_v1_sample_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x2c, 0x6f,
	0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x6f, 0x70, 0x65,
	0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a,
	0x06, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x15, 0x0a, 0x06, 0x6d,
	0x69, 0x6e, 0x5f, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x69, 0x6e,
	0x54, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x6d, 0x61, 0x78, 0x54, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0xd8, 0x01,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12,
	0x3b, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x42, 0x0a, 0x06,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x06, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73,
	0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0xd7, 0x02, 0x0a, 0x06, 0x50, 0x52, 0x45,
	0x46, 0x49, 0x58, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f,
	0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x53, 0x10,
	0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x4d,
	0x41, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x4e, 0x41, 0x4d,
	0x45, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x56, 0x45, 0x52,
	0x53, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f,
	0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x53, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c,
	0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x06, 0x12, 0x16,
	0x0a, 0x12, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42,
	0x55, 0x54, 0x45, 0x53, 0x10, 0x07, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x50, 0x41, 0x4e, 0x5f, 0x54,
	0x52, 0x41, 0x43, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x50, 0x41,
	0x4e, 0x5f, 0x53, 0x50, 0x41, 0x4e, 0x5f, 0x49, 0x44, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09, 0x53,
	0x50, 0x41, 0x4e, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x0a, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x50,
	0x41, 0x4e, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x53, 0x10, 0x0b, 0x12,
	0x17, 0x0a, 0x13, 0x53, 0x50, 0x41, 0x4e, 0x5f, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x53,
	0x50, 0x41, 0x4e, 0x5f, 0x49, 0x44, 0x10, 0x0d, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x4f, 0x47, 0x53,
	0x5f, 0x54, 0x52, 0x41, 0x43, 0x45, 0x5f, 0x49, 0x44, 0x10, 0x0e, 0x12, 0x10, 0x0a, 0x0c, 0x4c,
	0x4f, 0x47, 0x53, 0x5f, 0x53, 0x50, 0x41, 0x4e, 0x5f, 0x49, 0x44, 0x10, 0x0f, 0x12, 0x0e, 0x0a,
	0x0a, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x10, 0x10, 0x12, 0x13, 0x0a,
	0x0f, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x53,
	0x10, 0x11, 0x2a, 0x2f, 0x0a, 0x0a, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4b, 0x69, 0x6e, 0x64,
	0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x54, 0x52, 0x41, 0x43, 0x45, 0x53, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x4f, 0x47,
	0x53, 0x10, 0x03, 0x42, 0x6f, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x65, 0x6d,
	0x64, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x65, 0x6d, 0x64, 0x62, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02,
	0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_samples_v1_sample_proto_rawDescOnce sync.Once
	file_samples_v1_sample_proto_rawDescData = file_samples_v1_sample_proto_rawDesc
)

func file_samples_v1_sample_proto_rawDescGZIP() []byte {
	file_samples_v1_sample_proto_rawDescOnce.Do(func() {
		file_samples_v1_sample_proto_rawDescData = protoimpl.X.CompressGZIP(file_samples_v1_sample_proto_rawDescData)
	})
	return file_samples_v1_sample_proto_rawDescData
}

var file_samples_v1_sample_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_samples_v1_sample_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_samples_v1_sample_proto_goTypes = []interface{}{
	(PREFIX)(0),            // 0: v1.PREFIX
	(SampleKind)(0),        // 1: v1.SampleKind
	(*Sample)(nil),         // 2: v1.Sample
	(*Data)(nil),           // 3: v1.Data
	(*v1.MetricsData)(nil), // 4: opentelemetry.proto.metrics.v1.MetricsData
	(*v11.LogsData)(nil),   // 5: opentelemetry.proto.logs.v1.LogsData
	(*v12.TracesData)(nil), // 6: opentelemetry.proto.trace.v1.TracesData
}
var file_samples_v1_sample_proto_depIdxs = []int32{
	4, // 0: v1.Data.metrics:type_name -> opentelemetry.proto.metrics.v1.MetricsData
	5, // 1: v1.Data.logs:type_name -> opentelemetry.proto.logs.v1.LogsData
	6, // 2: v1.Data.traces:type_name -> opentelemetry.proto.trace.v1.TracesData
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_samples_v1_sample_proto_init() }
func file_samples_v1_sample_proto_init() {
	if File_samples_v1_sample_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_samples_v1_sample_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_samples_v1_sample_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
	file_samples_v1_sample_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Data_Metrics)(nil),
		(*Data_Logs)(nil),
		(*Data_Traces)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_samples_v1_sample_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_samples_v1_sample_proto_goTypes,
		DependencyIndexes: file_samples_v1_sample_proto_depIdxs,
		EnumInfos:         file_samples_v1_sample_proto_enumTypes,
		MessageInfos:      file_samples_v1_sample_proto_msgTypes,
	}.Build()
	File_samples_v1_sample_proto = out.File
	file_samples_v1_sample_proto_rawDesc = nil
	file_samples_v1_sample_proto_goTypes = nil
	file_samples_v1_sample_proto_depIdxs = nil
}
