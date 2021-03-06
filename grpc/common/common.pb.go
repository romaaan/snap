/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2016 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go.
// source: github.com/intelsdi-x/snap/grpc/common/common.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	github.com/intelsdi-x/snap/grpc/common/common.proto

It has these top-level messages:
	Time
	Empty
	SnapError
	Label
	Metric
	NamespaceElement
	SubscribedPlugin
	ConfigMap
	Plugin
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Time struct {
	Sec  int64 `protobuf:"varint,1,opt,name=sec" json:"sec,omitempty"`
	Nsec int64 `protobuf:"varint,2,opt,name=nsec" json:"nsec,omitempty"`
}

func (m *Time) Reset()                    { *m = Time{} }
func (m *Time) String() string            { return proto.CompactTextString(m) }
func (*Time) ProtoMessage()               {}
func (*Time) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type SnapError struct {
	ErrorString string            `protobuf:"bytes,1,opt,name=error_string,json=errorString" json:"error_string,omitempty"`
	ErrorFields map[string]string `protobuf:"bytes,2,rep,name=error_fields,json=errorFields" json:"error_fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SnapError) Reset()                    { *m = SnapError{} }
func (m *SnapError) String() string            { return proto.CompactTextString(m) }
func (*SnapError) ProtoMessage()               {}
func (*SnapError) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SnapError) GetErrorFields() map[string]string {
	if m != nil {
		return m.ErrorFields
	}
	return nil
}

type Label struct {
	Index uint64 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Label) Reset()                    { *m = Label{} }
func (m *Label) String() string            { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()               {}
func (*Label) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// core.Metric
type Metric struct {
	Namespace          []*NamespaceElement `protobuf:"bytes,1,rep,name=Namespace,json=namespace" json:"Namespace,omitempty"`
	Version            int64               `protobuf:"varint,2,opt,name=Version,json=version" json:"Version,omitempty"`
	Config             *ConfigMap          `protobuf:"bytes,3,opt,name=Config,json=config" json:"Config,omitempty"`
	LastAdvertisedTime *Time               `protobuf:"bytes,4,opt,name=LastAdvertisedTime,json=lastAdvertisedTime" json:"LastAdvertisedTime,omitempty"`
	Tags               map[string]string   `protobuf:"bytes,5,rep,name=Tags,json=tags" json:"Tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Timestamp          *Time               `protobuf:"bytes,6,opt,name=Timestamp,json=timestamp" json:"Timestamp,omitempty"`
	Unit               string              `protobuf:"bytes,7,opt,name=Unit,json=unit" json:"Unit,omitempty"`
	Description        string              `protobuf:"bytes,8,opt,name=Description,json=description" json:"Description,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*Metric_StringData
	//	*Metric_Float32Data
	//	*Metric_Float64Data
	//	*Metric_Int32Data
	//	*Metric_Int64Data
	//	*Metric_BytesData
	//	*Metric_BoolData
	Data isMetric_Data `protobuf_oneof:"data"`
}

func (m *Metric) Reset()                    { *m = Metric{} }
func (m *Metric) String() string            { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()               {}
func (*Metric) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type isMetric_Data interface {
	isMetric_Data()
}

type Metric_StringData struct {
	StringData string `protobuf:"bytes,9,opt,name=string_data,json=stringData,oneof"`
}
type Metric_Float32Data struct {
	Float32Data float32 `protobuf:"fixed32,10,opt,name=float32_data,json=float32Data,oneof"`
}
type Metric_Float64Data struct {
	Float64Data float64 `protobuf:"fixed64,11,opt,name=float64_data,json=float64Data,oneof"`
}
type Metric_Int32Data struct {
	Int32Data int32 `protobuf:"varint,12,opt,name=int32_data,json=int32Data,oneof"`
}
type Metric_Int64Data struct {
	Int64Data int64 `protobuf:"varint,13,opt,name=int64_data,json=int64Data,oneof"`
}
type Metric_BytesData struct {
	BytesData []byte `protobuf:"bytes,14,opt,name=bytes_data,json=bytesData,proto3,oneof"`
}
type Metric_BoolData struct {
	BoolData bool `protobuf:"varint,15,opt,name=bool_data,json=boolData,oneof"`
}

func (*Metric_StringData) isMetric_Data()  {}
func (*Metric_Float32Data) isMetric_Data() {}
func (*Metric_Float64Data) isMetric_Data() {}
func (*Metric_Int32Data) isMetric_Data()   {}
func (*Metric_Int64Data) isMetric_Data()   {}
func (*Metric_BytesData) isMetric_Data()   {}
func (*Metric_BoolData) isMetric_Data()    {}

func (m *Metric) GetData() isMetric_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Metric) GetNamespace() []*NamespaceElement {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *Metric) GetConfig() *ConfigMap {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Metric) GetLastAdvertisedTime() *Time {
	if m != nil {
		return m.LastAdvertisedTime
	}
	return nil
}

func (m *Metric) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Metric) GetTimestamp() *Time {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Metric) GetStringData() string {
	if x, ok := m.GetData().(*Metric_StringData); ok {
		return x.StringData
	}
	return ""
}

func (m *Metric) GetFloat32Data() float32 {
	if x, ok := m.GetData().(*Metric_Float32Data); ok {
		return x.Float32Data
	}
	return 0
}

func (m *Metric) GetFloat64Data() float64 {
	if x, ok := m.GetData().(*Metric_Float64Data); ok {
		return x.Float64Data
	}
	return 0
}

func (m *Metric) GetInt32Data() int32 {
	if x, ok := m.GetData().(*Metric_Int32Data); ok {
		return x.Int32Data
	}
	return 0
}

func (m *Metric) GetInt64Data() int64 {
	if x, ok := m.GetData().(*Metric_Int64Data); ok {
		return x.Int64Data
	}
	return 0
}

func (m *Metric) GetBytesData() []byte {
	if x, ok := m.GetData().(*Metric_BytesData); ok {
		return x.BytesData
	}
	return nil
}

func (m *Metric) GetBoolData() bool {
	if x, ok := m.GetData().(*Metric_BoolData); ok {
		return x.BoolData
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Metric) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Metric_OneofMarshaler, _Metric_OneofUnmarshaler, _Metric_OneofSizer, []interface{}{
		(*Metric_StringData)(nil),
		(*Metric_Float32Data)(nil),
		(*Metric_Float64Data)(nil),
		(*Metric_Int32Data)(nil),
		(*Metric_Int64Data)(nil),
		(*Metric_BytesData)(nil),
		(*Metric_BoolData)(nil),
	}
}

func _Metric_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Metric)
	// data
	switch x := m.Data.(type) {
	case *Metric_StringData:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringData)
	case *Metric_Float32Data:
		b.EncodeVarint(10<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.Float32Data)))
	case *Metric_Float64Data:
		b.EncodeVarint(11<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.Float64Data))
	case *Metric_Int32Data:
		b.EncodeVarint(12<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Int32Data))
	case *Metric_Int64Data:
		b.EncodeVarint(13<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Int64Data))
	case *Metric_BytesData:
		b.EncodeVarint(14<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.BytesData)
	case *Metric_BoolData:
		t := uint64(0)
		if x.BoolData {
			t = 1
		}
		b.EncodeVarint(15<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("Metric.Data has unexpected type %T", x)
	}
	return nil
}

func _Metric_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Metric)
	switch tag {
	case 9: // data.string_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Data = &Metric_StringData{x}
		return true, err
	case 10: // data.float32_data
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Data = &Metric_Float32Data{math.Float32frombits(uint32(x))}
		return true, err
	case 11: // data.float64_data
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Data = &Metric_Float64Data{math.Float64frombits(x)}
		return true, err
	case 12: // data.int32_data
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Data = &Metric_Int32Data{int32(x)}
		return true, err
	case 13: // data.int64_data
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Data = &Metric_Int64Data{int64(x)}
		return true, err
	case 14: // data.bytes_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Data = &Metric_BytesData{x}
		return true, err
	case 15: // data.bool_data
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Data = &Metric_BoolData{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _Metric_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Metric)
	// data
	switch x := m.Data.(type) {
	case *Metric_StringData:
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.StringData)))
		n += len(x.StringData)
	case *Metric_Float32Data:
		n += proto.SizeVarint(10<<3 | proto.WireFixed32)
		n += 4
	case *Metric_Float64Data:
		n += proto.SizeVarint(11<<3 | proto.WireFixed64)
		n += 8
	case *Metric_Int32Data:
		n += proto.SizeVarint(12<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Int32Data))
	case *Metric_Int64Data:
		n += proto.SizeVarint(13<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Int64Data))
	case *Metric_BytesData:
		n += proto.SizeVarint(14<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.BytesData)))
		n += len(x.BytesData)
	case *Metric_BoolData:
		n += proto.SizeVarint(15<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type NamespaceElement struct {
	Value       string `protobuf:"bytes,1,opt,name=Value,json=value" json:"Value,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=Description,json=description" json:"Description,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=Name,json=name" json:"Name,omitempty"`
}

func (m *NamespaceElement) Reset()                    { *m = NamespaceElement{} }
func (m *NamespaceElement) String() string            { return proto.CompactTextString(m) }
func (*NamespaceElement) ProtoMessage()               {}
func (*NamespaceElement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// core.SubscribedPlugin
type SubscribedPlugin struct {
	TypeName string     `protobuf:"bytes,1,opt,name=TypeName,json=typeName" json:"TypeName,omitempty"`
	Name     string     `protobuf:"bytes,2,opt,name=Name,json=name" json:"Name,omitempty"`
	Version  int64      `protobuf:"varint,3,opt,name=Version,json=version" json:"Version,omitempty"`
	Config   *ConfigMap `protobuf:"bytes,4,opt,name=Config,json=config" json:"Config,omitempty"`
}

func (m *SubscribedPlugin) Reset()                    { *m = SubscribedPlugin{} }
func (m *SubscribedPlugin) String() string            { return proto.CompactTextString(m) }
func (*SubscribedPlugin) ProtoMessage()               {}
func (*SubscribedPlugin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SubscribedPlugin) GetConfig() *ConfigMap {
	if m != nil {
		return m.Config
	}
	return nil
}

type ConfigMap struct {
	IntMap    map[string]int64  `protobuf:"bytes,1,rep,name=IntMap,json=intMap" json:"IntMap,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	StringMap map[string]string `protobuf:"bytes,2,rep,name=StringMap,json=stringMap" json:"StringMap,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// double is float64
	FloatMap map[string]float64 `protobuf:"bytes,3,rep,name=FloatMap,json=floatMap" json:"FloatMap,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
	BoolMap  map[string]bool    `protobuf:"bytes,4,rep,name=BoolMap,json=boolMap" json:"BoolMap,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *ConfigMap) Reset()                    { *m = ConfigMap{} }
func (m *ConfigMap) String() string            { return proto.CompactTextString(m) }
func (*ConfigMap) ProtoMessage()               {}
func (*ConfigMap) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ConfigMap) GetIntMap() map[string]int64 {
	if m != nil {
		return m.IntMap
	}
	return nil
}

func (m *ConfigMap) GetStringMap() map[string]string {
	if m != nil {
		return m.StringMap
	}
	return nil
}

func (m *ConfigMap) GetFloatMap() map[string]float64 {
	if m != nil {
		return m.FloatMap
	}
	return nil
}

func (m *ConfigMap) GetBoolMap() map[string]bool {
	if m != nil {
		return m.BoolMap
	}
	return nil
}

// core.Plugin
type Plugin struct {
	TypeName string `protobuf:"bytes,1,opt,name=TypeName,json=typeName" json:"TypeName,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,json=name" json:"Name,omitempty"`
	Version  int64  `protobuf:"varint,3,opt,name=Version,json=version" json:"Version,omitempty"`
}

func (m *Plugin) Reset()                    { *m = Plugin{} }
func (m *Plugin) String() string            { return proto.CompactTextString(m) }
func (*Plugin) ProtoMessage()               {}
func (*Plugin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*Time)(nil), "common.Time")
	proto.RegisterType((*Empty)(nil), "common.Empty")
	proto.RegisterType((*SnapError)(nil), "common.SnapError")
	proto.RegisterType((*Label)(nil), "common.Label")
	proto.RegisterType((*Metric)(nil), "common.Metric")
	proto.RegisterType((*NamespaceElement)(nil), "common.NamespaceElement")
	proto.RegisterType((*SubscribedPlugin)(nil), "common.SubscribedPlugin")
	proto.RegisterType((*ConfigMap)(nil), "common.ConfigMap")
	proto.RegisterType((*Plugin)(nil), "common.Plugin")
}

func init() {
	proto.RegisterFile("github.com/intelsdi-x/snap/grpc/common/common.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 764 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0x51, 0x6f, 0xe3, 0x44,
	0x10, 0x8e, 0x63, 0xc7, 0xb1, 0xc7, 0xb9, 0x23, 0xac, 0x78, 0xb0, 0x22, 0x1d, 0xe7, 0x33, 0x2f,
	0x06, 0x41, 0x22, 0xda, 0xa3, 0x1c, 0x6d, 0x55, 0x89, 0xd2, 0x54, 0x41, 0x6a, 0x11, 0x72, 0x4b,
	0x1f, 0xa9, 0xd6, 0xf6, 0x26, 0xac, 0xb0, 0xd7, 0x96, 0xbd, 0xa9, 0x9a, 0x67, 0x1e, 0xf8, 0x45,
	0xfc, 0x2a, 0xfe, 0x04, 0xda, 0x5d, 0xdb, 0x75, 0x1a, 0x50, 0x54, 0x89, 0x97, 0x76, 0xe6, 0x9b,
	0xef, 0x1b, 0xef, 0x8c, 0x3f, 0x6f, 0xe0, 0x70, 0x45, 0xf9, 0x6f, 0xeb, 0x68, 0x1a, 0xe7, 0xd9,
	0x8c, 0x32, 0x4e, 0xd2, 0x2a, 0xa1, 0x5f, 0x3d, 0xce, 0x2a, 0x86, 0x8b, 0xd9, 0xaa, 0x2c, 0xe2,
	0x59, 0x9c, 0x67, 0x59, 0xce, 0xea, 0x7f, 0xd3, 0xa2, 0xcc, 0x79, 0x8e, 0x4c, 0x95, 0xf9, 0x5f,
	0x82, 0x71, 0x4b, 0x33, 0x82, 0xc6, 0xa0, 0x57, 0x24, 0x76, 0x35, 0x4f, 0x0b, 0xf4, 0x50, 0x84,
	0x08, 0x81, 0xc1, 0x04, 0xd4, 0x97, 0x90, 0x8c, 0xfd, 0x21, 0x0c, 0xe6, 0x59, 0xc1, 0x37, 0xfe,
	0x5f, 0x1a, 0xd8, 0x37, 0x0c, 0x17, 0xf3, 0xb2, 0xcc, 0x4b, 0xf4, 0x0e, 0x46, 0x44, 0x04, 0xf7,
	0x15, 0x2f, 0x29, 0x5b, 0xc9, 0x2e, 0x76, 0xe8, 0x48, 0xec, 0x46, 0x42, 0x68, 0xde, 0x50, 0x96,
	0x94, 0xa4, 0x49, 0xe5, 0xf6, 0x3d, 0x3d, 0x70, 0x0e, 0xfc, 0x69, 0x7d, 0xa8, 0xb6, 0xd7, 0x54,
	0xfe, 0xbd, 0x94, 0xa4, 0x39, 0xe3, 0xe5, 0xa6, 0x6e, 0xa3, 0x90, 0xc9, 0x19, 0x8c, 0x9f, 0x13,
	0xc4, 0xd1, 0x7f, 0x27, 0x9b, 0xfa, 0xa1, 0x22, 0x44, 0x9f, 0xc0, 0xe0, 0x01, 0xa7, 0x6b, 0x22,
	0xcf, 0x6e, 0x87, 0x2a, 0x39, 0xee, 0x7f, 0xd0, 0xfc, 0xaf, 0x61, 0x70, 0x85, 0x23, 0x92, 0x0a,
	0x0a, 0x65, 0x09, 0x79, 0x94, 0x32, 0x23, 0x54, 0x89, 0x9c, 0x19, 0x67, 0x8d, 0x4e, 0xc6, 0xfe,
	0x1f, 0x03, 0x30, 0xaf, 0x09, 0x2f, 0x69, 0x8c, 0x8e, 0xc0, 0xfe, 0x09, 0x67, 0xa4, 0x2a, 0x70,
	0x4c, 0x5c, 0x4d, 0x4e, 0xe0, 0x36, 0x13, 0xb4, 0x85, 0x79, 0x4a, 0x32, 0xc2, 0x78, 0x68, 0xb3,
	0x06, 0x41, 0x2e, 0x0c, 0xef, 0x48, 0x59, 0xd1, 0x9c, 0xd5, 0xdb, 0x1c, 0x3e, 0xa8, 0x14, 0x7d,
	0x0e, 0xe6, 0x0f, 0x39, 0x5b, 0xd2, 0x95, 0xab, 0x7b, 0x5a, 0xe0, 0x1c, 0x7c, 0xdc, 0xb4, 0x53,
	0xe8, 0x35, 0x2e, 0x42, 0x33, 0x96, 0x21, 0x3a, 0x05, 0x74, 0x85, 0x2b, 0xfe, 0x7d, 0xf2, 0x40,
	0x4a, 0x4e, 0x2b, 0x92, 0x88, 0xf7, 0xe6, 0x1a, 0x52, 0x36, 0x6a, 0x64, 0x02, 0x0b, 0x51, 0xba,
	0xc3, 0x43, 0xe2, 0x3d, 0xe3, 0x55, 0xe5, 0x0e, 0xb6, 0x4f, 0xad, 0x06, 0x9b, 0x8a, 0x92, 0xda,
	0xb6, 0xc1, 0xf1, 0xaa, 0x42, 0x5f, 0x80, 0x2d, 0x54, 0x15, 0xc7, 0x59, 0xe1, 0x9a, 0xff, 0xf2,
	0x08, 0x9b, 0x37, 0x65, 0xb1, 0xb3, 0x5f, 0x18, 0xe5, 0xee, 0x50, 0xed, 0x6c, 0xcd, 0x28, 0x47,
	0x1e, 0x38, 0x17, 0xa4, 0x8a, 0x4b, 0x5a, 0x70, 0x31, 0xb4, 0xa5, 0xfc, 0x90, 0x3c, 0x41, 0xe8,
	0x1d, 0x38, 0xca, 0x2c, 0xf7, 0x09, 0xe6, 0xd8, 0xb5, 0x05, 0x63, 0xd1, 0x0b, 0x41, 0x81, 0x17,
	0x98, 0x63, 0xf4, 0x19, 0x8c, 0x96, 0x69, 0x8e, 0xf9, 0xe1, 0x81, 0xe2, 0x80, 0xa7, 0x05, 0xfd,
	0x45, 0x2f, 0x74, 0x6a, 0x74, 0x8b, 0x74, 0xf4, 0x5e, 0x91, 0x1c, 0x4f, 0x0b, 0xb4, 0x96, 0x74,
	0xf4, 0x5e, 0x92, 0xde, 0x02, 0x50, 0xd6, 0xf6, 0x19, 0x79, 0x5a, 0x30, 0x58, 0xf4, 0x42, 0x5b,
	0x62, 0x1d, 0x42, 0xd3, 0xe3, 0x95, 0x78, 0x47, 0x35, 0xe1, 0xa9, 0x43, 0xb4, 0xe1, 0xa4, 0x52,
	0x84, 0xd7, 0x9e, 0x16, 0x8c, 0x04, 0x41, 0x62, 0x92, 0xf0, 0x06, 0xec, 0x28, 0xcf, 0x53, 0x55,
	0xff, 0xc8, 0xd3, 0x02, 0x6b, 0xd1, 0x0b, 0x2d, 0x01, 0x89, 0xf2, 0xe4, 0x5b, 0xb0, 0xdb, 0x1d,
	0xbf, 0xc4, 0xb0, 0xe7, 0x26, 0x18, 0xa2, 0xa5, 0xff, 0x2b, 0x8c, 0x9f, 0x3b, 0x4c, 0xa8, 0xee,
	0xa4, 0x4a, 0xeb, 0xa8, 0x9e, 0xef, 0xbe, 0xbf, 0xbb, 0x7b, 0x04, 0x86, 0xe8, 0x25, 0x2d, 0xd7,
	0xb8, 0xfc, 0x4f, 0x0d, 0xc6, 0x37, 0xeb, 0x48, 0x90, 0x22, 0x92, 0xfc, 0x9c, 0xae, 0x57, 0x94,
	0xa1, 0x09, 0x58, 0xb7, 0x9b, 0x82, 0x48, 0xb2, 0x7a, 0x86, 0xc5, 0xeb, 0xbc, 0x6d, 0xd2, 0xf9,
	0x54, 0xba, 0x3e, 0xd7, 0xff, 0xcb, 0xe7, 0xc6, 0x1e, 0x9f, 0xfb, 0x7f, 0xeb, 0x60, 0xb7, 0x28,
	0xfa, 0x06, 0xcc, 0x1f, 0x19, 0xbf, 0xc6, 0x45, 0xfd, 0xbd, 0xbd, 0xd9, 0x11, 0x4e, 0x55, 0x5d,
	0xd9, 0xd7, 0xa4, 0x32, 0x41, 0x67, 0x60, 0xab, 0x8b, 0x47, 0x28, 0xd5, 0x5d, 0xe3, 0xed, 0x2a,
	0x5b, 0x8a, 0x12, 0xdb, 0x55, 0x93, 0xa3, 0x13, 0xb0, 0x2e, 0x85, 0x81, 0x84, 0x5c, 0x97, 0xf2,
	0xb7, 0xbb, 0xf2, 0x86, 0xa1, 0xd4, 0xd6, 0xb2, 0x4e, 0xd1, 0x07, 0x18, 0x9e, 0xe7, 0x79, 0x2a,
	0xb4, 0x86, 0xd4, 0x7e, 0xba, 0xab, 0xad, 0x09, 0x4a, 0x3a, 0x8c, 0x54, 0x36, 0xf9, 0x0e, 0x9c,
	0xce, 0x34, 0xfb, 0x8c, 0xa2, 0x77, 0x8c, 0x32, 0x39, 0x85, 0xd7, 0xdb, 0xe3, 0xbc, 0xc4, 0x66,
	0x93, 0x13, 0x78, 0xb5, 0x35, 0xcd, 0x3e, 0xb1, 0xd6, 0x15, 0x1f, 0xc3, 0xa8, 0x3b, 0xce, 0x3e,
	0xad, 0xd5, 0xbd, 0x90, 0x43, 0x30, 0xff, 0x6f, 0xb3, 0x45, 0xa6, 0xfc, 0x89, 0x3b, 0xfc, 0x27,
	0x00, 0x00, 0xff, 0xff, 0xef, 0x9e, 0x96, 0x24, 0x19, 0x07, 0x00, 0x00,
}
