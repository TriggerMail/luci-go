// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/TriggerMail/luci-go/cipd/appengine/impl/repo/tasks/tasks.proto

package tasks

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1 "github.com/TriggerMail/luci-go/cipd/api/cipd/v1"
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

// RunProcessors task runs a processing step on an uploaded package instance.
type RunProcessors struct {
	Instance             *v1.Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RunProcessors) Reset()         { *m = RunProcessors{} }
func (m *RunProcessors) String() string { return proto.CompactTextString(m) }
func (*RunProcessors) ProtoMessage()    {}
func (*RunProcessors) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2aa67c95ad0b65e, []int{0}
}

func (m *RunProcessors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunProcessors.Unmarshal(m, b)
}
func (m *RunProcessors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunProcessors.Marshal(b, m, deterministic)
}
func (m *RunProcessors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunProcessors.Merge(m, src)
}
func (m *RunProcessors) XXX_Size() int {
	return xxx_messageInfo_RunProcessors.Size(m)
}
func (m *RunProcessors) XXX_DiscardUnknown() {
	xxx_messageInfo_RunProcessors.DiscardUnknown(m)
}

var xxx_messageInfo_RunProcessors proto.InternalMessageInfo

func (m *RunProcessors) GetInstance() *v1.Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func init() {
	proto.RegisterType((*RunProcessors)(nil), "tasks.RunProcessors")
}

func init() {
	proto.RegisterFile("github.com/TriggerMail/luci-go/cipd/appengine/impl/repo/tasks/tasks.proto", fileDescriptor_d2aa67c95ad0b65e)
}

var fileDescriptor_d2aa67c95ad0b65e = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4f, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0x4f, 0xce, 0x2c, 0x48, 0xd1, 0x4f, 0x2c, 0x28, 0x48, 0xcd, 0x4b, 0xcf, 0xcc, 0x4b, 0xd5,
	0xcf, 0xcc, 0x2d, 0xc8, 0xd1, 0x2f, 0x4a, 0x2d, 0xc8, 0xd7, 0x2f, 0x49, 0x2c, 0xce, 0x2e, 0x86,
	0x90, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x8e, 0x94, 0x01, 0x3e, 0x73, 0xa0,
	0x8c, 0x32, 0x43, 0xb0, 0x21, 0x10, 0x8d, 0x4a, 0xd6, 0x5c, 0xbc, 0x41, 0xa5, 0x79, 0x01, 0x45,
	0xf9, 0xc9, 0xa9, 0xc5, 0xc5, 0xf9, 0x45, 0xc5, 0x42, 0x5a, 0x5c, 0x1c, 0x99, 0x79, 0xc5, 0x25,
	0x89, 0x79, 0xc9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x7c, 0x7a, 0x20, 0x7d, 0x7a,
	0x9e, 0x50, 0xd1, 0x20, 0xb8, 0x7c, 0x12, 0x1b, 0xd8, 0x0c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x83, 0xbc, 0x88, 0x8c, 0xbf, 0x00, 0x00, 0x00,
}
