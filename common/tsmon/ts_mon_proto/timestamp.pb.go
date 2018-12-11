// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/TriggerMail/luci-go/common/tsmon/ts_mon_proto/timestamp.proto

package ts_mon_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Timestamp struct {
	Seconds              *int64   `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	Nanos                *int32   `protobuf:"varint,2,opt,name=nanos" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_7afd1967c97dbebc, []int{0}
}

func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamp.Unmarshal(m, b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return xxx_messageInfo_Timestamp.Size(m)
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetSeconds() int64 {
	if m != nil && m.Seconds != nil {
		return *m.Seconds
	}
	return 0
}

func (m *Timestamp) GetNanos() int32 {
	if m != nil && m.Nanos != nil {
		return *m.Nanos
	}
	return 0
}

func init() {
	proto.RegisterType((*Timestamp)(nil), "ts_mon.proto.Timestamp")
}

func init() {
	proto.RegisterFile("github.com/TriggerMail/luci-go/common/tsmon/ts_mon_proto/timestamp.proto", fileDescriptor_7afd1967c97dbebc)
}

var fileDescriptor_7afd1967c97dbebc = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4b, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2f, 0x29, 0x86, 0x90, 0xf1, 0xb9, 0xf9, 0x79,
	0xf1, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0xfa, 0x25, 0x99, 0xb9, 0xa9, 0xc5, 0x25, 0x89, 0xb9, 0x05,
	0x7a, 0x60, 0xbe, 0x10, 0x0f, 0x44, 0x16, 0xc2, 0x53, 0xb2, 0xe6, 0xe2, 0x0c, 0x81, 0x29, 0x10,
	0x92, 0xe0, 0x62, 0x2f, 0x4e, 0x4d, 0xce, 0xcf, 0x4b, 0x29, 0x96, 0x60, 0x54, 0x60, 0xd4, 0x60,
	0x0e, 0x82, 0x71, 0x85, 0x44, 0xb8, 0x58, 0xf3, 0x12, 0xf3, 0xf2, 0x8b, 0x25, 0x98, 0x14, 0x18,
	0x35, 0x58, 0x83, 0x20, 0x1c, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x83, 0x91, 0xed, 0x8b,
	0x00, 0x00, 0x00,
}
