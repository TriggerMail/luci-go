// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/TriggerMail/luci-go/scheduler/appengine/internal/cursors.proto

package internal

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

// InvocationsCursor is used to paginate results of GetInvocations RPC call.
//
// It is serialized in base64 and sent to the clients. There's no integrity
// protection: we assume broken cursors are rejected down the call stack.
//
// The internal structure of the cursor is implementation detail and clients
// must not depend on it.
type InvocationsCursor struct {
	// ID of the last scanned invocation (active or finished).
	//
	// The query will return all IDs that are larger than this one.
	LastScanned          int64    `protobuf:"varint,2,opt,name=last_scanned,json=lastScanned,proto3" json:"last_scanned,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvocationsCursor) Reset()         { *m = InvocationsCursor{} }
func (m *InvocationsCursor) String() string { return proto.CompactTextString(m) }
func (*InvocationsCursor) ProtoMessage()    {}
func (*InvocationsCursor) Descriptor() ([]byte, []int) {
	return fileDescriptor_dded6d90ada69499, []int{0}
}

func (m *InvocationsCursor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvocationsCursor.Unmarshal(m, b)
}
func (m *InvocationsCursor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvocationsCursor.Marshal(b, m, deterministic)
}
func (m *InvocationsCursor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvocationsCursor.Merge(m, src)
}
func (m *InvocationsCursor) XXX_Size() int {
	return xxx_messageInfo_InvocationsCursor.Size(m)
}
func (m *InvocationsCursor) XXX_DiscardUnknown() {
	xxx_messageInfo_InvocationsCursor.DiscardUnknown(m)
}

var xxx_messageInfo_InvocationsCursor proto.InternalMessageInfo

func (m *InvocationsCursor) GetLastScanned() int64 {
	if m != nil {
		return m.LastScanned
	}
	return 0
}

func init() {
	proto.RegisterType((*InvocationsCursor)(nil), "internal.cursors.InvocationsCursor")
}

func init() {
	proto.RegisterFile("github.com/TriggerMail/luci-go/scheduler/appengine/internal/cursors.proto", fileDescriptor_dded6d90ada69499)
}

var fileDescriptor_dded6d90ada69499 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xcd, 0x31, 0x0b, 0xc2, 0x30,
	0x10, 0x40, 0x61, 0x5a, 0x45, 0x4a, 0x74, 0xa8, 0x9d, 0x1c, 0xd5, 0xc9, 0x29, 0x19, 0x5c, 0x05,
	0x41, 0x27, 0x1d, 0x75, 0x73, 0x91, 0x98, 0x1e, 0x6d, 0x20, 0xbd, 0x2b, 0x77, 0x89, 0xbf, 0x5f,
	0xac, 0x76, 0xfd, 0x86, 0xf7, 0xd4, 0xb1, 0x21, 0xed, 0x5a, 0xa6, 0xce, 0xa7, 0x4e, 0x13, 0x37,
	0x26, 0x24, 0xe7, 0x8d, 0xb8, 0x16, 0xea, 0x14, 0x80, 0x8d, 0xed, 0x7b, 0xc0, 0xc6, 0x23, 0x18,
	0x8f, 0x11, 0x18, 0x6d, 0x30, 0x2e, 0xb1, 0x10, 0x8b, 0xee, 0x99, 0x22, 0x55, 0xe5, 0xe8, 0xfa,
	0xef, 0xdb, 0x83, 0x5a, 0x5e, 0xf0, 0x4d, 0xce, 0x46, 0x4f, 0x28, 0xe7, 0x41, 0xab, 0x8d, 0x5a,
	0x04, 0x2b, 0xf1, 0x29, 0xce, 0x22, 0x42, 0xbd, 0xca, 0xd7, 0xd9, 0x6e, 0x72, 0x9b, 0x7f, 0xed,
	0xfe, 0xa3, 0xeb, 0xb4, 0xc8, 0xca, 0xfc, 0xa4, 0x1e, 0xc5, 0x58, 0x7c, 0xcd, 0x86, 0xc5, 0xfe,
	0x13, 0x00, 0x00, 0xff, 0xff, 0x97, 0x4f, 0xda, 0x97, 0xa5, 0x00, 0x00, 0x00,
}
