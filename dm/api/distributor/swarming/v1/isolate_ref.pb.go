// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/TriggerMail/luci-go/dm/api/distributor/swarming/v1/isolate_ref.proto

package swarmingV1

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

type IsolatedRef struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Server               string   `protobuf:"bytes,2,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsolatedRef) Reset()         { *m = IsolatedRef{} }
func (m *IsolatedRef) String() string { return proto.CompactTextString(m) }
func (*IsolatedRef) ProtoMessage()    {}
func (*IsolatedRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ca7d5f1aabd0c83, []int{0}
}

func (m *IsolatedRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsolatedRef.Unmarshal(m, b)
}
func (m *IsolatedRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsolatedRef.Marshal(b, m, deterministic)
}
func (m *IsolatedRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsolatedRef.Merge(m, src)
}
func (m *IsolatedRef) XXX_Size() int {
	return xxx_messageInfo_IsolatedRef.Size(m)
}
func (m *IsolatedRef) XXX_DiscardUnknown() {
	xxx_messageInfo_IsolatedRef.DiscardUnknown(m)
}

var xxx_messageInfo_IsolatedRef proto.InternalMessageInfo

func (m *IsolatedRef) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *IsolatedRef) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func init() {
	proto.RegisterType((*IsolatedRef)(nil), "swarmingV1.IsolatedRef")
}

func init() {
	proto.RegisterFile("github.com/TriggerMail/luci-go/dm/api/distributor/swarming/v1/isolate_ref.proto", fileDescriptor_0ca7d5f1aabd0c83)
}

var fileDescriptor_0ca7d5f1aabd0c83 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4d, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0x4f, 0xc9, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xc9, 0x2c, 0x2e, 0x29, 0xca, 0x4c, 0x2a,
	0x2d, 0xc9, 0x2f, 0xd2, 0x2f, 0x2e, 0x4f, 0x2c, 0xca, 0xcd, 0xcc, 0x4b, 0xd7, 0x2f, 0x33, 0xd4,
	0xcf, 0x2c, 0xce, 0xcf, 0x49, 0x2c, 0x49, 0x8d, 0x2f, 0x4a, 0x4d, 0xd3, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x82, 0x49, 0x87, 0x19, 0x2a, 0x99, 0x72, 0x71, 0x7b, 0x42, 0x14, 0xa4, 0x04,
	0xa5, 0xa6, 0x09, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31,
	0x65, 0xa6, 0x08, 0x89, 0x71, 0xb1, 0x15, 0xa7, 0x16, 0x95, 0xa5, 0x16, 0x49, 0x30, 0x81, 0xc5,
	0xa0, 0xbc, 0x24, 0x36, 0xb0, 0x49, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x75, 0xf9, 0x75,
	0x31, 0x92, 0x00, 0x00, 0x00,
}
