// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto3_proto/test.proto

package proto3_proto

import (
	fmt "fmt"
	math "math"

	proto2_proto "github.com/golang/protobuf/internal/testprotos/proto2_proto"
	proto "github.com/golang/protobuf/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message_Humour int32

const (
	Message_UNKNOWN     Message_Humour = 0
	Message_PUNS        Message_Humour = 1
	Message_SLAPSTICK   Message_Humour = 2
	Message_BILL_BAILEY Message_Humour = 3
)

var Message_Humour_name = map[int32]string{
	0: "UNKNOWN",
	1: "PUNS",
	2: "SLAPSTICK",
	3: "BILL_BAILEY",
}

var Message_Humour_value = map[string]int32{
	"UNKNOWN":     0,
	"PUNS":        1,
	"SLAPSTICK":   2,
	"BILL_BAILEY": 3,
}

func (x Message_Humour) String() string {
	return proto.EnumName(Message_Humour_name, int32(x))
}

func (Message_Humour) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ff83f0b8d2b92afa, []int{0, 0}
}

type Message struct {
	Name                 string                               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Hilarity             Message_Humour                       `protobuf:"varint,2,opt,name=hilarity,proto3,enum=proto3_test.Message_Humour" json:"hilarity,omitempty"`
	HeightInCm           uint32                               `protobuf:"varint,3,opt,name=height_in_cm,json=heightInCm,proto3" json:"height_in_cm,omitempty"`
	Data                 []byte                               `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	ResultCount          int64                                `protobuf:"varint,7,opt,name=result_count,json=resultCount,proto3" json:"result_count,omitempty"`
	TrueScotsman         bool                                 `protobuf:"varint,8,opt,name=true_scotsman,json=trueScotsman,proto3" json:"true_scotsman,omitempty"`
	Score                float32                              `protobuf:"fixed32,9,opt,name=score,proto3" json:"score,omitempty"`
	Key                  []uint64                             `protobuf:"varint,5,rep,packed,name=key,proto3" json:"key,omitempty"`
	ShortKey             []int32                              `protobuf:"varint,19,rep,packed,name=short_key,json=shortKey,proto3" json:"short_key,omitempty"`
	Nested               *Nested                              `protobuf:"bytes,6,opt,name=nested,proto3" json:"nested,omitempty"`
	RFunny               []Message_Humour                     `protobuf:"varint,16,rep,packed,name=r_funny,json=rFunny,proto3,enum=proto3_test.Message_Humour" json:"r_funny,omitempty"`
	Terrain              map[string]*Nested                   `protobuf:"bytes,10,rep,name=terrain,proto3" json:"terrain,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Proto2Field          *proto2_proto.SubDefaults            `protobuf:"bytes,11,opt,name=proto2_field,json=proto2Field,proto3" json:"proto2_field,omitempty"`
	Proto2Value          map[string]*proto2_proto.SubDefaults `protobuf:"bytes,13,rep,name=proto2_value,json=proto2Value,proto3" json:"proto2_value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Anything             *anypb.Any                           `protobuf:"bytes,14,opt,name=anything,proto3" json:"anything,omitempty"`
	ManyThings           []*anypb.Any                         `protobuf:"bytes,15,rep,name=many_things,json=manyThings,proto3" json:"many_things,omitempty"`
	Submessage           *Message                             `protobuf:"bytes,17,opt,name=submessage,proto3" json:"submessage,omitempty"`
	Children             []*Message                           `protobuf:"bytes,18,rep,name=children,proto3" json:"children,omitempty"`
	StringMap            map[string]string                    `protobuf:"bytes,20,rep,name=string_map,json=stringMap,proto3" json:"string_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff83f0b8d2b92afa, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Message) GetHilarity() Message_Humour {
	if m != nil {
		return m.Hilarity
	}
	return Message_UNKNOWN
}

func (m *Message) GetHeightInCm() uint32 {
	if m != nil {
		return m.HeightInCm
	}
	return 0
}

func (m *Message) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Message) GetResultCount() int64 {
	if m != nil {
		return m.ResultCount
	}
	return 0
}

func (m *Message) GetTrueScotsman() bool {
	if m != nil {
		return m.TrueScotsman
	}
	return false
}

func (m *Message) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Message) GetKey() []uint64 {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Message) GetShortKey() []int32 {
	if m != nil {
		return m.ShortKey
	}
	return nil
}

func (m *Message) GetNested() *Nested {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *Message) GetRFunny() []Message_Humour {
	if m != nil {
		return m.RFunny
	}
	return nil
}

func (m *Message) GetTerrain() map[string]*Nested {
	if m != nil {
		return m.Terrain
	}
	return nil
}

func (m *Message) GetProto2Field() *proto2_proto.SubDefaults {
	if m != nil {
		return m.Proto2Field
	}
	return nil
}

func (m *Message) GetProto2Value() map[string]*proto2_proto.SubDefaults {
	if m != nil {
		return m.Proto2Value
	}
	return nil
}

func (m *Message) GetAnything() *anypb.Any {
	if m != nil {
		return m.Anything
	}
	return nil
}

func (m *Message) GetManyThings() []*anypb.Any {
	if m != nil {
		return m.ManyThings
	}
	return nil
}

func (m *Message) GetSubmessage() *Message {
	if m != nil {
		return m.Submessage
	}
	return nil
}

func (m *Message) GetChildren() []*Message {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *Message) GetStringMap() map[string]string {
	if m != nil {
		return m.StringMap
	}
	return nil
}

type Nested struct {
	Bunny                string   `protobuf:"bytes,1,opt,name=bunny,proto3" json:"bunny,omitempty"`
	Cute                 bool     `protobuf:"varint,2,opt,name=cute,proto3" json:"cute,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nested) Reset()         { *m = Nested{} }
func (m *Nested) String() string { return proto.CompactTextString(m) }
func (*Nested) ProtoMessage()    {}
func (*Nested) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff83f0b8d2b92afa, []int{1}
}

func (m *Nested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nested.Unmarshal(m, b)
}
func (m *Nested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nested.Marshal(b, m, deterministic)
}
func (m *Nested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nested.Merge(m, src)
}
func (m *Nested) XXX_Size() int {
	return xxx_messageInfo_Nested.Size(m)
}
func (m *Nested) XXX_DiscardUnknown() {
	xxx_messageInfo_Nested.DiscardUnknown(m)
}

var xxx_messageInfo_Nested proto.InternalMessageInfo

func (m *Nested) GetBunny() string {
	if m != nil {
		return m.Bunny
	}
	return ""
}

func (m *Nested) GetCute() bool {
	if m != nil {
		return m.Cute
	}
	return false
}

type MessageWithMap struct {
	ByteMapping          map[bool][]byte `protobuf:"bytes,1,rep,name=byte_mapping,json=byteMapping,proto3" json:"byte_mapping,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MessageWithMap) Reset()         { *m = MessageWithMap{} }
func (m *MessageWithMap) String() string { return proto.CompactTextString(m) }
func (*MessageWithMap) ProtoMessage()    {}
func (*MessageWithMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff83f0b8d2b92afa, []int{2}
}

func (m *MessageWithMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithMap.Unmarshal(m, b)
}
func (m *MessageWithMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithMap.Marshal(b, m, deterministic)
}
func (m *MessageWithMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithMap.Merge(m, src)
}
func (m *MessageWithMap) XXX_Size() int {
	return xxx_messageInfo_MessageWithMap.Size(m)
}
func (m *MessageWithMap) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithMap.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithMap proto.InternalMessageInfo

func (m *MessageWithMap) GetByteMapping() map[bool][]byte {
	if m != nil {
		return m.ByteMapping
	}
	return nil
}

type IntMap struct {
	Rtt                  map[int32]int32 `protobuf:"bytes,1,rep,name=rtt,proto3" json:"rtt,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IntMap) Reset()         { *m = IntMap{} }
func (m *IntMap) String() string { return proto.CompactTextString(m) }
func (*IntMap) ProtoMessage()    {}
func (*IntMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff83f0b8d2b92afa, []int{3}
}

func (m *IntMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntMap.Unmarshal(m, b)
}
func (m *IntMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntMap.Marshal(b, m, deterministic)
}
func (m *IntMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntMap.Merge(m, src)
}
func (m *IntMap) XXX_Size() int {
	return xxx_messageInfo_IntMap.Size(m)
}
func (m *IntMap) XXX_DiscardUnknown() {
	xxx_messageInfo_IntMap.DiscardUnknown(m)
}

var xxx_messageInfo_IntMap proto.InternalMessageInfo

func (m *IntMap) GetRtt() map[int32]int32 {
	if m != nil {
		return m.Rtt
	}
	return nil
}

type IntMaps struct {
	Maps                 []*IntMap `protobuf:"bytes,1,rep,name=maps,proto3" json:"maps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *IntMaps) Reset()         { *m = IntMaps{} }
func (m *IntMaps) String() string { return proto.CompactTextString(m) }
func (*IntMaps) ProtoMessage()    {}
func (*IntMaps) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff83f0b8d2b92afa, []int{4}
}

func (m *IntMaps) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntMaps.Unmarshal(m, b)
}
func (m *IntMaps) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntMaps.Marshal(b, m, deterministic)
}
func (m *IntMaps) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntMaps.Merge(m, src)
}
func (m *IntMaps) XXX_Size() int {
	return xxx_messageInfo_IntMaps.Size(m)
}
func (m *IntMaps) XXX_DiscardUnknown() {
	xxx_messageInfo_IntMaps.DiscardUnknown(m)
}

var xxx_messageInfo_IntMaps proto.InternalMessageInfo

func (m *IntMaps) GetMaps() []*IntMap {
	if m != nil {
		return m.Maps
	}
	return nil
}

type TestUTF8 struct {
	Scalar string   `protobuf:"bytes,1,opt,name=scalar,proto3" json:"scalar,omitempty"`
	Vector []string `protobuf:"bytes,2,rep,name=vector,proto3" json:"vector,omitempty"`
	// Types that are valid to be assigned to Oneof:
	//	*TestUTF8_Field
	Oneof                isTestUTF8_Oneof `protobuf_oneof:"oneof"`
	MapKey               map[string]int64 `protobuf:"bytes,4,rep,name=map_key,json=mapKey,proto3" json:"map_key,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MapValue             map[int64]string `protobuf:"bytes,5,rep,name=map_value,json=mapValue,proto3" json:"map_value,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TestUTF8) Reset()         { *m = TestUTF8{} }
func (m *TestUTF8) String() string { return proto.CompactTextString(m) }
func (*TestUTF8) ProtoMessage()    {}
func (*TestUTF8) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff83f0b8d2b92afa, []int{5}
}

func (m *TestUTF8) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestUTF8.Unmarshal(m, b)
}
func (m *TestUTF8) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestUTF8.Marshal(b, m, deterministic)
}
func (m *TestUTF8) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestUTF8.Merge(m, src)
}
func (m *TestUTF8) XXX_Size() int {
	return xxx_messageInfo_TestUTF8.Size(m)
}
func (m *TestUTF8) XXX_DiscardUnknown() {
	xxx_messageInfo_TestUTF8.DiscardUnknown(m)
}

var xxx_messageInfo_TestUTF8 proto.InternalMessageInfo

func (m *TestUTF8) GetScalar() string {
	if m != nil {
		return m.Scalar
	}
	return ""
}

func (m *TestUTF8) GetVector() []string {
	if m != nil {
		return m.Vector
	}
	return nil
}

type isTestUTF8_Oneof interface {
	isTestUTF8_Oneof()
}

type TestUTF8_Field struct {
	Field string `protobuf:"bytes,3,opt,name=field,proto3,oneof"`
}

func (*TestUTF8_Field) isTestUTF8_Oneof() {}

func (m *TestUTF8) GetOneof() isTestUTF8_Oneof {
	if m != nil {
		return m.Oneof
	}
	return nil
}

func (m *TestUTF8) GetField() string {
	if x, ok := m.GetOneof().(*TestUTF8_Field); ok {
		return x.Field
	}
	return ""
}

func (m *TestUTF8) GetMapKey() map[string]int64 {
	if m != nil {
		return m.MapKey
	}
	return nil
}

func (m *TestUTF8) GetMapValue() map[int64]string {
	if m != nil {
		return m.MapValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TestUTF8) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestUTF8_Field)(nil),
	}
}

func init() {
	proto.RegisterEnum("proto3_test.Message_Humour", Message_Humour_name, Message_Humour_value)
	proto.RegisterType((*Message)(nil), "proto3_test.Message")
	proto.RegisterMapType((map[string]*proto2_proto.SubDefaults)(nil), "proto3_test.Message.Proto2ValueEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto3_test.Message.StringMapEntry")
	proto.RegisterMapType((map[string]*Nested)(nil), "proto3_test.Message.TerrainEntry")
	proto.RegisterType((*Nested)(nil), "proto3_test.Nested")
	proto.RegisterType((*MessageWithMap)(nil), "proto3_test.MessageWithMap")
	proto.RegisterMapType((map[bool][]byte)(nil), "proto3_test.MessageWithMap.ByteMappingEntry")
	proto.RegisterType((*IntMap)(nil), "proto3_test.IntMap")
	proto.RegisterMapType((map[int32]int32)(nil), "proto3_test.IntMap.RttEntry")
	proto.RegisterType((*IntMaps)(nil), "proto3_test.IntMaps")
	proto.RegisterType((*TestUTF8)(nil), "proto3_test.TestUTF8")
	proto.RegisterMapType((map[string]int64)(nil), "proto3_test.TestUTF8.MapKeyEntry")
	proto.RegisterMapType((map[int64]string)(nil), "proto3_test.TestUTF8.MapValueEntry")
}

func init() { proto.RegisterFile("proto3_proto/test.proto", fileDescriptor_ff83f0b8d2b92afa) }

var fileDescriptor_ff83f0b8d2b92afa = []byte{
	// 926 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0x6d, 0x6f, 0xdb, 0x36,
	0x10, 0xae, 0x2c, 0xbf, 0xc8, 0x67, 0x3b, 0xf5, 0x98, 0xa0, 0xe3, 0xdc, 0x7d, 0x50, 0x5d, 0x0c,
	0xd3, 0xb0, 0x41, 0x1e, 0xbc, 0x6e, 0xeb, 0x9a, 0xbd, 0xc5, 0x59, 0x83, 0x18, 0x89, 0x9d, 0x40,
	0x76, 0xd6, 0x6d, 0x5f, 0x04, 0xda, 0xa1, 0x6d, 0x61, 0x12, 0x65, 0x88, 0x54, 0x01, 0xfd, 0x9c,
	0xfd, 0xa4, 0xfd, 0xa2, 0x0d, 0x24, 0xe5, 0x54, 0x6e, 0x15, 0xe4, 0x93, 0x79, 0x8f, 0x9f, 0xbb,
	0xe7, 0x78, 0x77, 0x3c, 0xc1, 0xc7, 0xdb, 0x24, 0x16, 0xf1, 0x37, 0xbe, 0xfa, 0x19, 0x08, 0xca,
	0x85, 0xab, 0x8e, 0xa8, 0x95, 0xff, 0x21, 0xa1, 0xde, 0x27, 0xeb, 0x38, 0x5e, 0x87, 0x74, 0xa0,
	0xb0, 0x45, 0xba, 0x1a, 0x10, 0x96, 0x69, 0x5e, 0x4f, 0x07, 0x18, 0x7e, 0x10, 0xa0, 0xff, 0x9f,
	0x05, 0x8d, 0x09, 0xe5, 0x9c, 0xac, 0x29, 0x42, 0x50, 0x65, 0x24, 0xa2, 0xd8, 0xb0, 0x0d, 0xa7,
	0xe9, 0xa9, 0x33, 0xfa, 0x1e, 0xac, 0x4d, 0x10, 0x92, 0x24, 0x10, 0x19, 0xae, 0xd8, 0x86, 0x73,
	0x30, 0x7c, 0xea, 0x16, 0x34, 0xdd, 0xdc, 0xd7, 0x3d, 0x4f, 0xa3, 0x38, 0x4d, 0xbc, 0x3b, 0x32,
	0xb2, 0xa1, 0xbd, 0xa1, 0xc1, 0x7a, 0x23, 0xfc, 0x80, 0xf9, 0xcb, 0x08, 0x9b, 0xb6, 0xe1, 0x74,
	0x3c, 0xd0, 0xd8, 0x98, 0x9d, 0x46, 0x52, 0xee, 0x96, 0x08, 0x82, 0xab, 0xb6, 0xe1, 0xb4, 0x3d,
	0x75, 0x46, 0xcf, 0xa0, 0x9d, 0x50, 0x9e, 0x86, 0xc2, 0x5f, 0xc6, 0x29, 0x13, 0xb8, 0x61, 0x1b,
	0x8e, 0xe9, 0xb5, 0x34, 0x76, 0x2a, 0x21, 0xf4, 0x1c, 0x3a, 0x22, 0x49, 0xa9, 0xcf, 0x97, 0xb1,
	0xe0, 0x11, 0x61, 0xd8, 0xb2, 0x0d, 0xc7, 0xf2, 0xda, 0x12, 0x9c, 0xe5, 0x18, 0x3a, 0x82, 0x1a,
	0x5f, 0xc6, 0x09, 0xc5, 0x4d, 0xdb, 0x70, 0x2a, 0x9e, 0x36, 0x50, 0x17, 0xcc, 0xbf, 0x69, 0x86,
	0x6b, 0xb6, 0xe9, 0x54, 0x3d, 0x79, 0x44, 0x4f, 0xa1, 0xc9, 0x37, 0x71, 0x22, 0x7c, 0x89, 0x1f,
	0xda, 0xa6, 0x53, 0xf3, 0x2c, 0x05, 0x5c, 0xd0, 0x0c, 0x7d, 0x09, 0x75, 0x46, 0xb9, 0xa0, 0xb7,
	0xb8, 0x6e, 0x1b, 0x4e, 0x6b, 0x78, 0xb8, 0x77, 0xf3, 0xa9, 0xfa, 0xcb, 0xcb, 0x29, 0xe8, 0x05,
	0x34, 0x12, 0x7f, 0x95, 0x32, 0x96, 0xe1, 0xae, 0x6d, 0x3e, 0x54, 0xa7, 0x7a, 0x72, 0x26, 0xa9,
	0xe8, 0x18, 0x1a, 0x82, 0x26, 0x09, 0x09, 0x18, 0x06, 0xdb, 0x74, 0x5a, 0xc3, 0x67, 0xa5, 0x5e,
	0x73, 0xcd, 0x79, 0xcd, 0x44, 0x92, 0x79, 0x3b, 0x0f, 0x74, 0x0c, 0xed, 0xbc, 0xad, 0xab, 0x80,
	0x86, 0xb7, 0xb8, 0xa5, 0xb2, 0xc4, 0x6e, 0x0e, 0xaa, 0x08, 0xb3, 0x74, 0xf1, 0x1b, 0x5d, 0x91,
	0x34, 0x14, 0xdc, 0xd3, 0xc3, 0x32, 0x3c, 0x93, 0x64, 0x74, 0x7e, 0xe7, 0xfc, 0x96, 0x84, 0x29,
	0xc5, 0x1d, 0x25, 0xff, 0x59, 0xa9, 0xfc, 0xb5, 0x22, 0xfe, 0x2e, 0x79, 0x3a, 0x85, 0x3c, 0x92,
	0x42, 0xd0, 0xd7, 0x60, 0x11, 0x96, 0x89, 0x4d, 0xc0, 0xd6, 0xf8, 0x40, 0xa5, 0x70, 0xe4, 0xea,
	0x49, 0x74, 0x77, 0x93, 0xe8, 0x9e, 0xb0, 0xcc, 0xbb, 0x63, 0xa1, 0x6f, 0xa1, 0x15, 0x11, 0x96,
	0xf9, 0xca, 0xe2, 0xf8, 0xb1, 0x92, 0x2e, 0x77, 0x02, 0x49, 0x9c, 0x2b, 0x1e, 0x7a, 0x01, 0xc0,
	0xd3, 0x45, 0xa4, 0x93, 0xc2, 0x1f, 0xe5, 0x52, 0x25, 0x09, 0x7b, 0x05, 0x9e, 0x4c, 0x6f, 0xb9,
	0x09, 0xc2, 0xdb, 0x84, 0x32, 0x8c, 0x72, 0xa5, 0x32, 0x9f, 0x3b, 0x16, 0x1a, 0x01, 0x70, 0x91,
	0x04, 0x6c, 0xed, 0x47, 0x64, 0x8b, 0x8f, 0x94, 0xcf, 0xf3, 0xd2, 0xc2, 0xcc, 0x14, 0x6d, 0x42,
	0xb6, 0xba, 0x2c, 0x4d, 0xbe, 0xb3, 0x7b, 0x57, 0xd0, 0x2e, 0x36, 0x6d, 0x37, 0x7a, 0xfa, 0x69,
	0xa9, 0xd1, 0xfb, 0x02, 0x6a, 0xba, 0xf2, 0x95, 0xfb, 0x87, 0x4b, 0x33, 0x5e, 0x55, 0x5e, 0x1a,
	0xbd, 0x3f, 0xa0, 0xfb, 0x7e, 0x1b, 0x4a, 0x82, 0xba, 0xfb, 0x41, 0xef, 0x9f, 0x85, 0x42, 0xe4,
	0x1f, 0xe1, 0x60, 0xff, 0x1e, 0x25, 0x71, 0x8f, 0x8a, 0x71, 0x9b, 0x05, 0xef, 0xfe, 0x2f, 0x50,
	0xd7, 0x33, 0x8d, 0x5a, 0xd0, 0xb8, 0x99, 0x5e, 0x4c, 0xaf, 0xde, 0x4c, 0xbb, 0x8f, 0x90, 0x05,
	0xd5, 0xeb, 0x9b, 0xe9, 0xac, 0x6b, 0xa0, 0x0e, 0x34, 0x67, 0x97, 0x27, 0xd7, 0xb3, 0xf9, 0xf8,
	0xf4, 0xa2, 0x5b, 0x41, 0x8f, 0xa1, 0x35, 0x1a, 0x5f, 0x5e, 0xfa, 0xa3, 0x93, 0xf1, 0xe5, 0xeb,
	0x3f, 0xbb, 0x66, 0x7f, 0x08, 0x75, 0x7d, 0x5b, 0x29, 0xb2, 0x50, 0x0f, 0x48, 0x0b, 0x6b, 0x43,
	0xae, 0x89, 0x65, 0x2a, 0xb4, 0xb2, 0xe5, 0xa9, 0x73, 0xff, 0x1f, 0x03, 0x0e, 0xf2, 0x1e, 0xbc,
	0x09, 0xc4, 0x66, 0x42, 0xb6, 0xe8, 0x0a, 0xda, 0x8b, 0x4c, 0x50, 0xd9, 0xb2, 0xad, 0x9c, 0x44,
	0x43, 0xb5, 0xed, 0xab, 0xb2, 0xb6, 0xe5, 0x2e, 0xee, 0x28, 0x13, 0x74, 0xa2, 0xe9, 0xf9, 0x58,
	0x2f, 0xde, 0x21, 0xbd, 0x9f, 0xa1, 0xfb, 0x3e, 0xa1, 0x58, 0x18, 0xab, 0xa4, 0x30, 0xed, 0x62,
	0x61, 0xb6, 0x50, 0x1f, 0x33, 0x21, 0x53, 0x73, 0xc1, 0x4c, 0x84, 0xc8, 0x33, 0xfa, 0x74, 0x2f,
	0x23, 0xcd, 0x70, 0x3d, 0x21, 0x74, 0x06, 0x92, 0xd8, 0xfb, 0x0e, 0xac, 0x1d, 0x50, 0x54, 0xac,
	0x95, 0x28, 0xd6, 0x8a, 0x8a, 0x43, 0x68, 0xe8, 0x78, 0x1c, 0x7d, 0x0e, 0xd5, 0x88, 0x6c, 0x79,
	0xae, 0x79, 0x58, 0xa2, 0xe9, 0x29, 0x42, 0xff, 0xdf, 0x0a, 0x58, 0x73, 0xca, 0xc5, 0xcd, 0xfc,
	0xec, 0x25, 0x7a, 0x02, 0x75, 0xbe, 0x24, 0x21, 0x49, 0xf2, 0x0e, 0xe4, 0x96, 0xc4, 0xdf, 0xd2,
	0xa5, 0x88, 0x13, 0x5c, 0xb1, 0x4d, 0x89, 0x6b, 0x0b, 0x3d, 0x81, 0x9a, 0xde, 0x3c, 0x72, 0xb9,
	0x37, 0xcf, 0x1f, 0x79, 0xda, 0x44, 0xaf, 0xa0, 0x11, 0x91, 0xad, 0xda, 0xa9, 0xd5, 0x92, 0xad,
	0xb6, 0xd3, 0x73, 0x27, 0x64, 0x7b, 0x41, 0x33, 0x7d, 0xf3, 0x7a, 0xa4, 0x0c, 0xf4, 0x2b, 0x34,
	0xa5, 0xaf, 0xbe, 0x62, 0xad, 0xe4, 0xed, 0x15, 0xbd, 0x0b, 0x2b, 0xc9, 0x8a, 0x72, 0xb3, 0xf7,
	0x03, 0xb4, 0x0a, 0x81, 0x1f, 0x1a, 0x66, 0xb3, 0xf8, 0x14, 0x8e, 0xa1, 0xb3, 0x17, 0xb5, 0xe8,
	0x6c, 0x3e, 0xf0, 0x12, 0x46, 0x0d, 0xa8, 0xc5, 0x8c, 0xc6, 0xab, 0xd1, 0x4f, 0x7f, 0x1d, 0xaf,
	0x03, 0xb1, 0x49, 0x17, 0xee, 0x32, 0x8e, 0x06, 0xeb, 0x38, 0x24, 0x6c, 0xfd, 0xee, 0xa3, 0x1c,
	0x30, 0x41, 0x13, 0x46, 0x42, 0xf5, 0x15, 0x56, 0x28, 0x1f, 0x14, 0x3f, 0xef, 0x8b, 0xba, 0xb6,
	0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x10, 0x03, 0x1b, 0x06, 0xf5, 0x07, 0x00, 0x00,
}
