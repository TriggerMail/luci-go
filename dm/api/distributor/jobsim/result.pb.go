// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/dm/api/distributor/jobsim/result.proto

package jobsim

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

// Result is the jobsim distributor result object used for executions.
type Result struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Value                int64    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_f53ecbabb0d591b3, []int{0}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Result) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Result)(nil), "jobsim.Result")
}

func init() {
	proto.RegisterFile("github.com/TriggerMail/luci-go/dm/api/distributor/jobsim/result.proto", fileDescriptor_f53ecbabb0d591b3)
}

var fileDescriptor_f53ecbabb0d591b3 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4e, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0x4f, 0xc9, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xc9, 0x2c, 0x2e, 0x29, 0xca, 0x4c, 0x2a,
	0x2d, 0xc9, 0x2f, 0xd2, 0xcf, 0xca, 0x4f, 0x2a, 0xce, 0xcc, 0xd5, 0x2f, 0x4a, 0x2d, 0x2e, 0xcd,
	0x29, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0x08, 0x2a, 0x59, 0x70, 0xb1, 0x05,
	0x81, 0xc5, 0x85, 0x24, 0xb8, 0xd8, 0x8b, 0x4b, 0x93, 0x93, 0x53, 0x8b, 0x8b, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x82, 0x60, 0x5c, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09,
	0x26, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x08, 0x27, 0x89, 0x0d, 0x6c, 0x90, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x0d, 0x6e, 0x99, 0x4e, 0x87, 0x00, 0x00, 0x00,
}
