// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package test

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

// This is a test comment shouldn't be included
type Msg1 struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Floatingvalue        float64  `protobuf:"fixed64,3,opt,name=floatingvalue,proto3" json:"floatingvalue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg1) Reset()         { *m = Msg1{} }
func (m *Msg1) String() string { return proto.CompactTextString(m) }
func (*Msg1) ProtoMessage()    {}
func (*Msg1) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_bf1b0ab336e7e7aa, []int{0}
}
func (m *Msg1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg1.Unmarshal(m, b)
}
func (m *Msg1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg1.Marshal(b, m, deterministic)
}
func (dst *Msg1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg1.Merge(dst, src)
}
func (m *Msg1) XXX_Size() int {
	return xxx_messageInfo_Msg1.Size(m)
}
func (m *Msg1) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg1.DiscardUnknown(m)
}

var xxx_messageInfo_Msg1 proto.InternalMessageInfo

func (m *Msg1) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Msg1) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Msg1) GetFloatingvalue() float64 {
	if m != nil {
		return m.Floatingvalue
	}
	return 0
}

type Msg2 struct {
	Liststring           []string         `protobuf:"bytes,2,rep,name=Liststring,proto3" json:"Liststring,omitempty"`
	Idtostringmap        map[int32]string `protobuf:"bytes,3,rep,name=Idtostringmap,proto3" json:"Idtostringmap,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Msg2) Reset()         { *m = Msg2{} }
func (m *Msg2) String() string { return proto.CompactTextString(m) }
func (*Msg2) ProtoMessage()    {}
func (*Msg2) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_bf1b0ab336e7e7aa, []int{1}
}
func (m *Msg2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg2.Unmarshal(m, b)
}
func (m *Msg2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg2.Marshal(b, m, deterministic)
}
func (dst *Msg2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg2.Merge(dst, src)
}
func (m *Msg2) XXX_Size() int {
	return xxx_messageInfo_Msg2.Size(m)
}
func (m *Msg2) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg2.DiscardUnknown(m)
}

var xxx_messageInfo_Msg2 proto.InternalMessageInfo

func (m *Msg2) GetListstring() []string {
	if m != nil {
		return m.Liststring
	}
	return nil
}

func (m *Msg2) GetIdtostringmap() map[int32]string {
	if m != nil {
		return m.Idtostringmap
	}
	return nil
}

type Msg3 struct {
	Nested               *Msg2    `protobuf:"bytes,1,opt,name=Nested,proto3" json:"Nested,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg3) Reset()         { *m = Msg3{} }
func (m *Msg3) String() string { return proto.CompactTextString(m) }
func (*Msg3) ProtoMessage()    {}
func (*Msg3) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_bf1b0ab336e7e7aa, []int{2}
}
func (m *Msg3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg3.Unmarshal(m, b)
}
func (m *Msg3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg3.Marshal(b, m, deterministic)
}
func (dst *Msg3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg3.Merge(dst, src)
}
func (m *Msg3) XXX_Size() int {
	return xxx_messageInfo_Msg3.Size(m)
}
func (m *Msg3) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg3.DiscardUnknown(m)
}

var xxx_messageInfo_Msg3 proto.InternalMessageInfo

func (m *Msg3) GetNested() *Msg2 {
	if m != nil {
		return m.Nested
	}
	return nil
}

type Msg4 struct {
	Params               []byte   `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg4) Reset()         { *m = Msg4{} }
func (m *Msg4) String() string { return proto.CompactTextString(m) }
func (*Msg4) ProtoMessage()    {}
func (*Msg4) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_bf1b0ab336e7e7aa, []int{3}
}
func (m *Msg4) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg4.Unmarshal(m, b)
}
func (m *Msg4) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg4.Marshal(b, m, deterministic)
}
func (dst *Msg4) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg4.Merge(dst, src)
}
func (m *Msg4) XXX_Size() int {
	return xxx_messageInfo_Msg4.Size(m)
}
func (m *Msg4) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg4.DiscardUnknown(m)
}

var xxx_messageInfo_Msg4 proto.InternalMessageInfo

func (m *Msg4) GetParams() []byte {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*Msg1)(nil), "Msg1")
	proto.RegisterType((*Msg2)(nil), "msg2")
	proto.RegisterMapType((map[int32]string)(nil), "msg2.IdtostringmapEntry")
	proto.RegisterType((*Msg3)(nil), "Msg3")
	proto.RegisterType((*Msg4)(nil), "Msg4")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_test_bf1b0ab336e7e7aa) }

var fileDescriptor_test_bf1b0ab336e7e7aa = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x51, 0x4b, 0xfb, 0x30,
	0x14, 0xc5, 0x97, 0x74, 0xd9, 0x9f, 0xdd, 0xfd, 0x27, 0x72, 0x11, 0x0d, 0x03, 0x47, 0x09, 0x0a,
	0x7d, 0x2a, 0xac, 0xed, 0x83, 0xf8, 0x20, 0xbe, 0xf8, 0x20, 0xe8, 0x90, 0xe0, 0x17, 0xc8, 0x6c,
	0x56, 0xaa, 0x6b, 0x3b, 0x9a, 0x4c, 0xd8, 0x37, 0xf2, 0x63, 0x4a, 0x93, 0xc9, 0x2c, 0xbe, 0xdd,
	0x7b, 0x7e, 0xe7, 0x94, 0x7b, 0x1a, 0x00, 0xab, 0x8d, 0x8d, 0xb7, 0x6d, 0x63, 0x1b, 0xf1, 0x02,
	0xc3, 0x67, 0x53, 0x2c, 0x10, 0x61, 0xb8, 0x54, 0x95, 0xe6, 0x24, 0x24, 0xd1, 0x58, 0xba, 0x19,
	0x4f, 0x80, 0x96, 0x39, 0xa7, 0x21, 0x89, 0x98, 0xa4, 0x65, 0x8e, 0x57, 0x30, 0x5d, 0x6f, 0x1a,
	0x65, 0xcb, 0xba, 0xf8, 0x54, 0x9b, 0x9d, 0xe6, 0x41, 0x48, 0x22, 0x22, 0xfb, 0xa2, 0xf8, 0x22,
	0x30, 0xac, 0x4c, 0x91, 0xe0, 0x1c, 0xe0, 0xa9, 0x34, 0xd6, 0xd8, 0xb6, 0xac, 0x0b, 0x4e, 0xc3,
	0x20, 0x1a, 0xcb, 0x5f, 0x0a, 0xde, 0xc1, 0xf4, 0x31, 0xb7, 0x8d, 0xdf, 0x2a, 0xb5, 0xe5, 0x41,
	0x18, 0x44, 0x93, 0x84, 0xc7, 0x5d, 0x3a, 0xee, 0xa1, 0x87, 0xda, 0xb6, 0x7b, 0xd9, 0xb7, 0xcf,
	0xee, 0x01, 0xff, 0x9a, 0xf0, 0x14, 0x82, 0x0f, 0xbd, 0x77, 0x3d, 0x98, 0xec, 0x46, 0x3c, 0x03,
	0xe6, 0xcf, 0xa5, 0xae, 0x9b, 0x5f, 0x6e, 0xe9, 0x0d, 0x11, 0xd7, 0xae, 0x7c, 0x8a, 0x97, 0x30,
	0x5a, 0x6a, 0x63, 0x75, 0xee, 0x62, 0x93, 0x84, 0xb9, 0x13, 0xe4, 0x41, 0x14, 0x73, 0x67, 0xcb,
	0xf0, 0x1c, 0x46, 0x5b, 0xd5, 0xaa, 0xca, 0x38, 0xdb, 0x7f, 0x79, 0xd8, 0x92, 0x0a, 0xfe, 0xbd,
	0x6a, 0x63, 0x9b, 0xd5, 0x3b, 0x5e, 0x00, 0x5b, 0xef, 0xea, 0xb7, 0x05, 0xb2, 0xb8, 0xfb, 0xad,
	0x33, 0xff, 0x25, 0x31, 0xf8, 0x01, 0x09, 0x7a, 0x65, 0xe6, 0x78, 0x7a, 0x04, 0xa9, 0x4f, 0xa4,
	0x1e, 0x2c, 0x8e, 0x20, 0xf3, 0x20, 0xf3, 0x20, 0x13, 0x83, 0xd5, 0xc8, 0xbd, 0x5c, 0xfa, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0xa4, 0xad, 0x21, 0x3b, 0xc7, 0x01, 0x00, 0x00,
}
