// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/oses.proto

package crimson

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

// An operating system in the database.
type OS struct {
	// The name of this operating system. Uniquely identifies this operating system.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A description of this operating system.
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OS) Reset()         { *m = OS{} }
func (m *OS) String() string { return proto.CompactTextString(m) }
func (*OS) ProtoMessage()    {}
func (*OS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccc648486e27b263, []int{0}
}

func (m *OS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OS.Unmarshal(m, b)
}
func (m *OS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OS.Marshal(b, m, deterministic)
}
func (m *OS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OS.Merge(m, src)
}
func (m *OS) XXX_Size() int {
	return xxx_messageInfo_OS.Size(m)
}
func (m *OS) XXX_DiscardUnknown() {
	xxx_messageInfo_OS.DiscardUnknown(m)
}

var xxx_messageInfo_OS proto.InternalMessageInfo

func (m *OS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OS) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// A request to list operating systems in the database.
type ListOSesRequest struct {
	// The names of operating systems to retrieve.
	Names                []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOSesRequest) Reset()         { *m = ListOSesRequest{} }
func (m *ListOSesRequest) String() string { return proto.CompactTextString(m) }
func (*ListOSesRequest) ProtoMessage()    {}
func (*ListOSesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccc648486e27b263, []int{1}
}

func (m *ListOSesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOSesRequest.Unmarshal(m, b)
}
func (m *ListOSesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOSesRequest.Marshal(b, m, deterministic)
}
func (m *ListOSesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOSesRequest.Merge(m, src)
}
func (m *ListOSesRequest) XXX_Size() int {
	return xxx_messageInfo_ListOSesRequest.Size(m)
}
func (m *ListOSesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOSesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListOSesRequest proto.InternalMessageInfo

func (m *ListOSesRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

// A response containing a list of operating systems in the database.
type ListOSesResponse struct {
	// The operating systems matching the request.
	Oses                 []*OS    `protobuf:"bytes,1,rep,name=oses,proto3" json:"oses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOSesResponse) Reset()         { *m = ListOSesResponse{} }
func (m *ListOSesResponse) String() string { return proto.CompactTextString(m) }
func (*ListOSesResponse) ProtoMessage()    {}
func (*ListOSesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccc648486e27b263, []int{2}
}

func (m *ListOSesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOSesResponse.Unmarshal(m, b)
}
func (m *ListOSesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOSesResponse.Marshal(b, m, deterministic)
}
func (m *ListOSesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOSesResponse.Merge(m, src)
}
func (m *ListOSesResponse) XXX_Size() int {
	return xxx_messageInfo_ListOSesResponse.Size(m)
}
func (m *ListOSesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOSesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListOSesResponse proto.InternalMessageInfo

func (m *ListOSesResponse) GetOses() []*OS {
	if m != nil {
		return m.Oses
	}
	return nil
}

func init() {
	proto.RegisterType((*OS)(nil), "crimson.OS")
	proto.RegisterType((*ListOSesRequest)(nil), "crimson.ListOSesRequest")
	proto.RegisterType((*ListOSesResponse)(nil), "crimson.ListOSesResponse")
}

func init() {
	proto.RegisterFile("github.com/TriggerMail/luci-go/machine-db/api/crimson/v1/oses.proto", fileDescriptor_ccc648486e27b263)
}

var fileDescriptor_ccc648486e27b263 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xe9, 0xba, 0x2a, 0x3b, 0x3d, 0x28, 0xc1, 0x43, 0x6f, 0x96, 0x5e, 0xdc, 0x8b, 0x09,
	0xba, 0x27, 0x7d, 0x06, 0xa1, 0x90, 0x3e, 0x41, 0x37, 0x3b, 0xb4, 0x03, 0x26, 0x13, 0x33, 0xa9,
	0xcf, 0x2f, 0x8d, 0x05, 0xbd, 0xcd, 0xfc, 0xff, 0xf7, 0xc3, 0x07, 0x6f, 0x13, 0x6b, 0x37, 0x27,
	0xf6, 0xb4, 0x78, 0xcd, 0x69, 0x32, 0x9f, 0x8b, 0x23, 0xe3, 0x47, 0x37, 0x53, 0xc0, 0xe7, 0xcb,
	0xd9, 0x8c, 0x91, 0x8c, 0x4b, 0xe4, 0x85, 0x83, 0xf9, 0x7e, 0x31, 0x2c, 0x28, 0x3a, 0x26, 0xce,
	0xac, 0x6e, 0xb7, 0xb8, 0x7b, 0x87, 0x5d, 0x3f, 0x28, 0x05, 0xfb, 0x30, 0x7a, 0x6c, 0xaa, 0xb6,
	0x3a, 0x1e, 0x6c, 0xb9, 0x55, 0x0b, 0xf5, 0x05, 0xc5, 0x25, 0x8a, 0x99, 0x38, 0x34, 0xbb, 0x52,
	0xfd, 0x8f, 0xba, 0x27, 0xb8, 0xfb, 0x20, 0xc9, 0xfd, 0x80, 0x62, 0xf1, 0x6b, 0x41, 0xc9, 0xea,
	0x01, 0xae, 0xd7, 0xb1, 0x34, 0x55, 0x7b, 0x75, 0x3c, 0xd8, 0xdf, 0xa7, 0x3b, 0xc1, 0xfd, 0x1f,
	0x28, 0x91, 0x83, 0xa0, 0x7a, 0x84, 0xfd, 0xea, 0x53, 0xc0, 0xfa, 0xb5, 0xd6, 0x9b, 0x90, 0xee,
	0x07, 0x5b, 0x8a, 0xf3, 0x4d, 0x31, 0x3d, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0xee, 0x58, 0x42,
	0x23, 0xe6, 0x00, 0x00, 0x00,
}
