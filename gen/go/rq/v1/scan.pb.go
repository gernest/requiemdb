// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: rq/v1/scan.proto

package v1

import (
	v11 "go.opentelemetry.io/proto/otlp/logs/v1"
	v1 "go.opentelemetry.io/proto/otlp/metrics/v1"
	v12 "go.opentelemetry.io/proto/otlp/trace/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Scan_SCOPE int32

const (
	Scan_METRICS Scan_SCOPE = 0
	Scan_TRACES  Scan_SCOPE = 2
	Scan_LOGS    Scan_SCOPE = 3
)

// Enum value maps for Scan_SCOPE.
var (
	Scan_SCOPE_name = map[int32]string{
		0: "METRICS",
		2: "TRACES",
		3: "LOGS",
	}
	Scan_SCOPE_value = map[string]int32{
		"METRICS": 0,
		"TRACES":  2,
		"LOGS":    3,
	}
)

func (x Scan_SCOPE) Enum() *Scan_SCOPE {
	p := new(Scan_SCOPE)
	*p = x
	return p
}

func (x Scan_SCOPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Scan_SCOPE) Descriptor() protoreflect.EnumDescriptor {
	return file_rq_v1_scan_proto_enumTypes[0].Descriptor()
}

func (Scan_SCOPE) Type() protoreflect.EnumType {
	return &file_rq_v1_scan_proto_enumTypes[0]
}

func (x Scan_SCOPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Scan_SCOPE.Descriptor instead.
func (Scan_SCOPE) EnumDescriptor() ([]byte, []int) {
	return file_rq_v1_scan_proto_rawDescGZIP(), []int{0, 0}
}

type Scan_BaseProp int32

const (
	Scan_RESOURCE_SCHEMA Scan_BaseProp = 0
	Scan_SCOPE_SCHEMA    Scan_BaseProp = 2
	Scan_SCOPE_NAME      Scan_BaseProp = 3
	Scan_SCOPE_VERSION   Scan_BaseProp = 4
	Scan_NAME            Scan_BaseProp = 6
	Scan_TRACE_ID        Scan_BaseProp = 8
	Scan_SPAN_ID         Scan_BaseProp = 9
	Scan_PARENT_SPAN_ID  Scan_BaseProp = 10
	Scan_LOGS_LEVEL      Scan_BaseProp = 11
)

// Enum value maps for Scan_BaseProp.
var (
	Scan_BaseProp_name = map[int32]string{
		0:  "RESOURCE_SCHEMA",
		2:  "SCOPE_SCHEMA",
		3:  "SCOPE_NAME",
		4:  "SCOPE_VERSION",
		6:  "NAME",
		8:  "TRACE_ID",
		9:  "SPAN_ID",
		10: "PARENT_SPAN_ID",
		11: "LOGS_LEVEL",
	}
	Scan_BaseProp_value = map[string]int32{
		"RESOURCE_SCHEMA": 0,
		"SCOPE_SCHEMA":    2,
		"SCOPE_NAME":      3,
		"SCOPE_VERSION":   4,
		"NAME":            6,
		"TRACE_ID":        8,
		"SPAN_ID":         9,
		"PARENT_SPAN_ID":  10,
		"LOGS_LEVEL":      11,
	}
)

func (x Scan_BaseProp) Enum() *Scan_BaseProp {
	p := new(Scan_BaseProp)
	*p = x
	return p
}

func (x Scan_BaseProp) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Scan_BaseProp) Descriptor() protoreflect.EnumDescriptor {
	return file_rq_v1_scan_proto_enumTypes[1].Descriptor()
}

func (Scan_BaseProp) Type() protoreflect.EnumType {
	return &file_rq_v1_scan_proto_enumTypes[1]
}

func (x Scan_BaseProp) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Scan_BaseProp.Descriptor instead.
func (Scan_BaseProp) EnumDescriptor() ([]byte, []int) {
	return file_rq_v1_scan_proto_rawDescGZIP(), []int{0, 1}
}

type Scan_AttributeProp int32

const (
	Scan_UNKOWN_ATTR         Scan_AttributeProp = 0
	Scan_RESOURCE_ATTRIBUTES Scan_AttributeProp = 1
	Scan_SCOPE_ATTRIBUTES    Scan_AttributeProp = 5
	Scan_ATTRIBUTES          Scan_AttributeProp = 7
)

// Enum value maps for Scan_AttributeProp.
var (
	Scan_AttributeProp_name = map[int32]string{
		0: "UNKOWN_ATTR",
		1: "RESOURCE_ATTRIBUTES",
		5: "SCOPE_ATTRIBUTES",
		7: "ATTRIBUTES",
	}
	Scan_AttributeProp_value = map[string]int32{
		"UNKOWN_ATTR":         0,
		"RESOURCE_ATTRIBUTES": 1,
		"SCOPE_ATTRIBUTES":    5,
		"ATTRIBUTES":          7,
	}
)

func (x Scan_AttributeProp) Enum() *Scan_AttributeProp {
	p := new(Scan_AttributeProp)
	*p = x
	return p
}

func (x Scan_AttributeProp) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Scan_AttributeProp) Descriptor() protoreflect.EnumDescriptor {
	return file_rq_v1_scan_proto_enumTypes[2].Descriptor()
}

func (Scan_AttributeProp) Type() protoreflect.EnumType {
	return &file_rq_v1_scan_proto_enumTypes[2]
}

func (x Scan_AttributeProp) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Scan_AttributeProp.Descriptor instead.
func (Scan_AttributeProp) EnumDescriptor() ([]byte, []int) {
	return file_rq_v1_scan_proto_rawDescGZIP(), []int{0, 2}
}

type Scan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scope Scan_SCOPE `protobuf:"varint,1,opt,name=scope,proto3,enum=v1.Scan_SCOPE" json:"scope,omitempty"`
	// Timestamps to bound scan. This is optional, if it is not set a time range
	// of the last 15 minutes since now.
	TimeRange *Scan_TimeRange `protobuf:"bytes,2,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	Filters   []*Scan_Filter  `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	// Number of samples to process. Defauluts to no limit.
	Limit uint64 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// Scans in reverse order, with latest samples comming first.  To get the
	// latest sample you can set reverse to true and limit 1.
	Reverse bool `protobuf:"varint,5,opt,name=reverse,proto3" json:"reverse,omitempty"`
	// Now is current scan evaluation time. This is optional, when not set current
	// system time is used.
	//
	// Useful for reprdocucible scanning by compining this with time_range a
	// script can ensure it will be processing the same samples.
	Now *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=now,proto3" json:"now,omitempty"`
	// Offset relative to current scanning time.
	Offset *durationpb.Duration `protobuf:"bytes,7,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Scan) Reset() {
	*x = Scan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rq_v1_scan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scan) ProtoMessage() {}

func (x *Scan) ProtoReflect() protoreflect.Message {
	mi := &file_rq_v1_scan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scan.ProtoReflect.Descriptor instead.
func (*Scan) Descriptor() ([]byte, []int) {
	return file_rq_v1_scan_proto_rawDescGZIP(), []int{0}
}

func (x *Scan) GetScope() Scan_SCOPE {
	if x != nil {
		return x.Scope
	}
	return Scan_METRICS
}

func (x *Scan) GetTimeRange() *Scan_TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return nil
}

func (x *Scan) GetFilters() []*Scan_Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *Scan) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Scan) GetReverse() bool {
	if x != nil {
		return x.Reverse
	}
	return false
}

func (x *Scan) GetNow() *timestamppb.Timestamp {
	if x != nil {
		return x.Now
	}
	return nil
}

func (x *Scan) GetOffset() *durationpb.Duration {
	if x != nil {
		return x.Offset
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*Data_Metrics
	//	*Data_Logs
	//	*Data_Trace
	Data isData_Data `protobuf_oneof:"data"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rq_v1_scan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_rq_v1_scan_proto_msgTypes[1]
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
	return file_rq_v1_scan_proto_rawDescGZIP(), []int{1}
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

func (x *Data) GetTrace() *v12.TracesData {
	if x, ok := x.GetData().(*Data_Trace); ok {
		return x.Trace
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

type Data_Trace struct {
	Trace *v12.TracesData `protobuf:"bytes,3,opt,name=trace,proto3,oneof"`
}

func (*Data_Metrics) isData_Data() {}

func (*Data_Logs) isData_Data() {}

func (*Data_Trace) isData_Data() {}

type Scan_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*Scan_Filter_Base
	//	*Scan_Filter_Attr
	Value isScan_Filter_Value `protobuf_oneof:"value"`
}

func (x *Scan_Filter) Reset() {
	*x = Scan_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rq_v1_scan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scan_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scan_Filter) ProtoMessage() {}

func (x *Scan_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_rq_v1_scan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scan_Filter.ProtoReflect.Descriptor instead.
func (*Scan_Filter) Descriptor() ([]byte, []int) {
	return file_rq_v1_scan_proto_rawDescGZIP(), []int{0, 0}
}

func (m *Scan_Filter) GetValue() isScan_Filter_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Scan_Filter) GetBase() *Scan_BaseFilter {
	if x, ok := x.GetValue().(*Scan_Filter_Base); ok {
		return x.Base
	}
	return nil
}

func (x *Scan_Filter) GetAttr() *Scan_AttrFilter {
	if x, ok := x.GetValue().(*Scan_Filter_Attr); ok {
		return x.Attr
	}
	return nil
}

type isScan_Filter_Value interface {
	isScan_Filter_Value()
}

type Scan_Filter_Base struct {
	Base *Scan_BaseFilter `protobuf:"bytes,1,opt,name=base,proto3,oneof"`
}

type Scan_Filter_Attr struct {
	Attr *Scan_AttrFilter `protobuf:"bytes,2,opt,name=attr,proto3,oneof"`
}

func (*Scan_Filter_Base) isScan_Filter_Value() {}

func (*Scan_Filter_Attr) isScan_Filter_Value() {}

type Scan_BaseFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prop  Scan_BaseProp `protobuf:"varint,1,opt,name=prop,proto3,enum=v1.Scan_BaseProp" json:"prop,omitempty"`
	Value string        `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Scan_BaseFilter) Reset() {
	*x = Scan_BaseFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rq_v1_scan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scan_BaseFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scan_BaseFilter) ProtoMessage() {}

func (x *Scan_BaseFilter) ProtoReflect() protoreflect.Message {
	mi := &file_rq_v1_scan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scan_BaseFilter.ProtoReflect.Descriptor instead.
func (*Scan_BaseFilter) Descriptor() ([]byte, []int) {
	return file_rq_v1_scan_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Scan_BaseFilter) GetProp() Scan_BaseProp {
	if x != nil {
		return x.Prop
	}
	return Scan_RESOURCE_SCHEMA
}

func (x *Scan_BaseFilter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Scan_AttrFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prop  Scan_AttributeProp `protobuf:"varint,1,opt,name=prop,proto3,enum=v1.Scan_AttributeProp" json:"prop,omitempty"`
	Key   string             `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value string             `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Scan_AttrFilter) Reset() {
	*x = Scan_AttrFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rq_v1_scan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scan_AttrFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scan_AttrFilter) ProtoMessage() {}

func (x *Scan_AttrFilter) ProtoReflect() protoreflect.Message {
	mi := &file_rq_v1_scan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scan_AttrFilter.ProtoReflect.Descriptor instead.
func (*Scan_AttrFilter) Descriptor() ([]byte, []int) {
	return file_rq_v1_scan_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Scan_AttrFilter) GetProp() Scan_AttributeProp {
	if x != nil {
		return x.Prop
	}
	return Scan_UNKOWN_ATTR
}

func (x *Scan_AttrFilter) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Scan_AttrFilter) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Scan_TimeRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *Scan_TimeRange) Reset() {
	*x = Scan_TimeRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rq_v1_scan_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scan_TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scan_TimeRange) ProtoMessage() {}

func (x *Scan_TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_rq_v1_scan_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scan_TimeRange.ProtoReflect.Descriptor instead.
func (*Scan_TimeRange) Descriptor() ([]byte, []int) {
	return file_rq_v1_scan_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Scan_TimeRange) GetStart() *timestamppb.Timestamp {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *Scan_TimeRange) GetEnd() *timestamppb.Timestamp {
	if x != nil {
		return x.End
	}
	return nil
}

var File_rq_v1_scan_proto protoreflect.FileDescriptor

var file_rq_v1_scan_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x71, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x6f,
	0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x07, 0x0a, 0x04, 0x53, 0x63, 0x61, 0x6e,
	0x12, 0x24, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x2e, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x52,
	0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x63, 0x61, 0x6e, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x63, 0x61, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x76, 0x65, 0x72, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x6e, 0x6f, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x03, 0x6e,
	0x6f, 0x77, 0x12, 0x31, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x1a, 0x67, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x29, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x48, 0x00, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x61, 0x74,
	0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63,
	0x61, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52,
	0x04, 0x61, 0x74, 0x74, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x49,
	0x0a, 0x0a, 0x42, 0x61, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x04,
	0x70, 0x72, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x63, 0x61, 0x6e, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x52, 0x04, 0x70,
	0x72, 0x6f, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x60, 0x0a, 0x0a, 0x41, 0x74, 0x74,
	0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x04, 0x70, 0x72, 0x6f, 0x70, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x52, 0x04, 0x70,
	0x72, 0x6f, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x6b, 0x0a, 0x09, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x2a, 0x0a, 0x05, 0x53, 0x43, 0x4f, 0x50,
	0x45, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x53, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x54, 0x52, 0x41, 0x43, 0x45, 0x53, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x4f,
	0x47, 0x53, 0x10, 0x03, 0x22, 0x9d, 0x01, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x50, 0x72, 0x6f,
	0x70, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x53, 0x43,
	0x48, 0x45, 0x4d, 0x41, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f,
	0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x43, 0x4f, 0x50,
	0x45, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x43, 0x4f, 0x50,
	0x45, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x41, 0x4d, 0x45, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x52, 0x41, 0x43, 0x45, 0x5f, 0x49,
	0x44, 0x10, 0x08, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x50, 0x41, 0x4e, 0x5f, 0x49, 0x44, 0x10, 0x09,
	0x12, 0x12, 0x0a, 0x0e, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x50, 0x41, 0x4e, 0x5f,
	0x49, 0x44, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x4c, 0x45, 0x56,
	0x45, 0x4c, 0x10, 0x0b, 0x22, 0x5f, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x70, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x4b, 0x4f, 0x57, 0x4e, 0x5f,
	0x41, 0x54, 0x54, 0x52, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x45, 0x53, 0x10, 0x01, 0x12,
	0x14, 0x0a, 0x10, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55,
	0x54, 0x45, 0x53, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55,
	0x54, 0x45, 0x53, 0x10, 0x07, 0x22, 0xd6, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x47,
	0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x07,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x3b, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x6f, 0x67, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04,
	0x6c, 0x6f, 0x67, 0x73, 0x12, 0x40, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x73, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52,
	0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x66,
	0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x53, 0x63, 0x61, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x65, 0x72, 0x6e, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x65,
	0x6d, 0x64, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x71, 0x2f, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31,
	0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rq_v1_scan_proto_rawDescOnce sync.Once
	file_rq_v1_scan_proto_rawDescData = file_rq_v1_scan_proto_rawDesc
)

func file_rq_v1_scan_proto_rawDescGZIP() []byte {
	file_rq_v1_scan_proto_rawDescOnce.Do(func() {
		file_rq_v1_scan_proto_rawDescData = protoimpl.X.CompressGZIP(file_rq_v1_scan_proto_rawDescData)
	})
	return file_rq_v1_scan_proto_rawDescData
}

var file_rq_v1_scan_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_rq_v1_scan_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_rq_v1_scan_proto_goTypes = []interface{}{
	(Scan_SCOPE)(0),               // 0: v1.Scan.SCOPE
	(Scan_BaseProp)(0),            // 1: v1.Scan.BaseProp
	(Scan_AttributeProp)(0),       // 2: v1.Scan.AttributeProp
	(*Scan)(nil),                  // 3: v1.Scan
	(*Data)(nil),                  // 4: v1.Data
	(*Scan_Filter)(nil),           // 5: v1.Scan.Filter
	(*Scan_BaseFilter)(nil),       // 6: v1.Scan.BaseFilter
	(*Scan_AttrFilter)(nil),       // 7: v1.Scan.AttrFilter
	(*Scan_TimeRange)(nil),        // 8: v1.Scan.TimeRange
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 10: google.protobuf.Duration
	(*v1.MetricsData)(nil),        // 11: opentelemetry.proto.metrics.v1.MetricsData
	(*v11.LogsData)(nil),          // 12: opentelemetry.proto.logs.v1.LogsData
	(*v12.TracesData)(nil),        // 13: opentelemetry.proto.trace.v1.TracesData
}
var file_rq_v1_scan_proto_depIdxs = []int32{
	0,  // 0: v1.Scan.scope:type_name -> v1.Scan.SCOPE
	8,  // 1: v1.Scan.time_range:type_name -> v1.Scan.TimeRange
	5,  // 2: v1.Scan.filters:type_name -> v1.Scan.Filter
	9,  // 3: v1.Scan.now:type_name -> google.protobuf.Timestamp
	10, // 4: v1.Scan.offset:type_name -> google.protobuf.Duration
	11, // 5: v1.Data.metrics:type_name -> opentelemetry.proto.metrics.v1.MetricsData
	12, // 6: v1.Data.logs:type_name -> opentelemetry.proto.logs.v1.LogsData
	13, // 7: v1.Data.trace:type_name -> opentelemetry.proto.trace.v1.TracesData
	6,  // 8: v1.Scan.Filter.base:type_name -> v1.Scan.BaseFilter
	7,  // 9: v1.Scan.Filter.attr:type_name -> v1.Scan.AttrFilter
	1,  // 10: v1.Scan.BaseFilter.prop:type_name -> v1.Scan.BaseProp
	2,  // 11: v1.Scan.AttrFilter.prop:type_name -> v1.Scan.AttributeProp
	9,  // 12: v1.Scan.TimeRange.start:type_name -> google.protobuf.Timestamp
	9,  // 13: v1.Scan.TimeRange.end:type_name -> google.protobuf.Timestamp
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_rq_v1_scan_proto_init() }
func file_rq_v1_scan_proto_init() {
	if File_rq_v1_scan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rq_v1_scan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scan); i {
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
		file_rq_v1_scan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_rq_v1_scan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scan_Filter); i {
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
		file_rq_v1_scan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scan_BaseFilter); i {
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
		file_rq_v1_scan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scan_AttrFilter); i {
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
		file_rq_v1_scan_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scan_TimeRange); i {
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
	file_rq_v1_scan_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Data_Metrics)(nil),
		(*Data_Logs)(nil),
		(*Data_Trace)(nil),
	}
	file_rq_v1_scan_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Scan_Filter_Base)(nil),
		(*Scan_Filter_Attr)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rq_v1_scan_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rq_v1_scan_proto_goTypes,
		DependencyIndexes: file_rq_v1_scan_proto_depIdxs,
		EnumInfos:         file_rq_v1_scan_proto_enumTypes,
		MessageInfos:      file_rq_v1_scan_proto_msgTypes,
	}.Build()
	File_rq_v1_scan_proto = out.File
	file_rq_v1_scan_proto_rawDesc = nil
	file_rq_v1_scan_proto_goTypes = nil
	file_rq_v1_scan_proto_depIdxs = nil
}
