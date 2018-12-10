// Code generated by protoc-gen-go.
// source: go.chromium.org/luci/grpc/internal/svctool/testdata/test.proto
// DO NOT EDIT!

/*
Package test is a generated protocol buffer package.

It is generated from these files:
	go.chromium.org/luci/grpc/internal/svctool/testdata/test.proto

It has these top-level messages:
	M1
	M2
	Void
*/
package test

import prpc "github.com/TriggerMail/luci-go/grpc/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import sub "github.com/TriggerMail/luci-go/grpc/internal/svctool/testdata/sub"
import google_protobuf "github.com/TriggerMail/luci-go/common/proto/google"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// The request message containing the user's name.
type M1 struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *M1) Reset()                    { *m = M1{} }
func (m *M1) String() string            { return proto.CompactTextString(m) }
func (*M1) ProtoMessage()               {}
func (*M1) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The response message containing the greetings
type M2 struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *M2) Reset()                    { *m = M2{} }
func (m *M2) String() string            { return proto.CompactTextString(m) }
func (*M2) ProtoMessage()               {}
func (*M2) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Void struct {
}

func (m *Void) Reset()                    { *m = Void{} }
func (m *Void) String() string            { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()               {}
func (*Void) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*M1)(nil), "test.M1")
	proto.RegisterType((*M2)(nil), "test.M2")
	proto.RegisterType((*Void)(nil), "test.Void")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for S1 service

type S1Client interface {
	M(ctx context.Context, in *M1, opts ...grpc.CallOption) (*M2, error)
}
type s1PRPCClient struct {
	client *prpc.Client
}

func NewS1PRPCClient(client *prpc.Client) S1Client {
	return &s1PRPCClient{client}
}

func (c *s1PRPCClient) M(ctx context.Context, in *M1, opts ...grpc.CallOption) (*M2, error) {
	out := new(M2)
	err := c.client.Call(ctx, "test.S1", "M", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type s1Client struct {
	cc *grpc.ClientConn
}

func NewS1Client(cc *grpc.ClientConn) S1Client {
	return &s1Client{cc}
}

func (c *s1Client) M(ctx context.Context, in *M1, opts ...grpc.CallOption) (*M2, error) {
	out := new(M2)
	err := grpc.Invoke(ctx, "/test.S1/M", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for S1 service

type S1Server interface {
	M(context.Context, *M1) (*M2, error)
}

func RegisterS1Server(s prpc.Registrar, srv S1Server) {
	s.RegisterService(&_S1_serviceDesc, srv)
}

func _S1_M_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(M1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S1Server).M(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.S1/M",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S1Server).M(ctx, req.(*M1))
	}
	return interceptor(ctx, in, info, handler)
}

var _S1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.S1",
	HandlerType: (*S1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "M",
			Handler:    _S1_M_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

// Client API for S2 service

type S2Client interface {
	Get(ctx context.Context, in *Void, opts ...grpc.CallOption) (*M1, error)
	Set(ctx context.Context, in *M1, opts ...grpc.CallOption) (*Void, error)
	Imp(ctx context.Context, in *sub.Sub, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}
type s2PRPCClient struct {
	client *prpc.Client
}

func NewS2PRPCClient(client *prpc.Client) S2Client {
	return &s2PRPCClient{client}
}

func (c *s2PRPCClient) Get(ctx context.Context, in *Void, opts ...grpc.CallOption) (*M1, error) {
	out := new(M1)
	err := c.client.Call(ctx, "test.S2", "Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s2PRPCClient) Set(ctx context.Context, in *M1, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.client.Call(ctx, "test.S2", "Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s2PRPCClient) Imp(ctx context.Context, in *sub.Sub, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := c.client.Call(ctx, "test.S2", "Imp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type s2Client struct {
	cc *grpc.ClientConn
}

func NewS2Client(cc *grpc.ClientConn) S2Client {
	return &s2Client{cc}
}

func (c *s2Client) Get(ctx context.Context, in *Void, opts ...grpc.CallOption) (*M1, error) {
	out := new(M1)
	err := grpc.Invoke(ctx, "/test.S2/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s2Client) Set(ctx context.Context, in *M1, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := grpc.Invoke(ctx, "/test.S2/Set", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s2Client) Imp(ctx context.Context, in *sub.Sub, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/test.S2/Imp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for S2 service

type S2Server interface {
	Get(context.Context, *Void) (*M1, error)
	Set(context.Context, *M1) (*Void, error)
	Imp(context.Context, *sub.Sub) (*google_protobuf.Empty, error)
}

func RegisterS2Server(s prpc.Registrar, srv S2Server) {
	s.RegisterService(&_S2_serviceDesc, srv)
}

func _S2_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S2Server).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.S2/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S2Server).Get(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _S2_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(M1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S2Server).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.S2/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S2Server).Set(ctx, req.(*M1))
	}
	return interceptor(ctx, in, info, handler)
}

func _S2_Imp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sub.Sub)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S2Server).Imp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.S2/Imp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S2Server).Imp(ctx, req.(*sub.Sub))
	}
	return interceptor(ctx, in, info, handler)
}

var _S2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.S2",
	HandlerType: (*S2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _S2_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _S2_Set_Handler,
		},
		{
			MethodName: "Imp",
			Handler:    _S2_Imp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() {
	proto.RegisterFile("github.com/TriggerMail/luci-go/grpc/internal/svctool/testdata/test.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x8f, 0x3d, 0x4f, 0xc5, 0x20,
	0x14, 0x86, 0xdb, 0x7b, 0x9b, 0xab, 0x9e, 0x91, 0xc1, 0xd4, 0x1a, 0x3f, 0xc2, 0xa4, 0x83, 0x90,
	0xe2, 0x2f, 0x30, 0xd1, 0x18, 0x87, 0x2e, 0x36, 0x71, 0x87, 0x16, 0xb1, 0x49, 0x29, 0x4d, 0x01,
	0x13, 0xff, 0xbd, 0x40, 0x53, 0x75, 0x76, 0x80, 0xbc, 0x87, 0xe7, 0xf0, 0x26, 0x0f, 0x3c, 0xa8,
	0xc1, 0x7d, 0x78, 0x41, 0x3a, 0xa3, 0xe9, 0xe8, 0xbb, 0x21, 0x5d, 0x77, 0xca, 0x50, 0xb5, 0xcc,
	0x1d, 0x1d, 0x26, 0x27, 0x97, 0x89, 0x8f, 0xd4, 0x7e, 0x76, 0xce, 0x98, 0x91, 0x3a, 0x69, 0x5d,
	0xcf, 0x1d, 0x4f, 0x81, 0xcc, 0x8b, 0x71, 0x06, 0x15, 0x31, 0x57, 0x8f, 0xff, 0x2c, 0xb2, 0x5e,
	0xc4, 0xb3, 0x76, 0x55, 0xe7, 0xca, 0x18, 0x35, 0x4a, 0x9a, 0x26, 0xe1, 0xdf, 0xa9, 0xd4, 0xb3,
	0xfb, 0x5a, 0x21, 0x2e, 0x61, 0xd7, 0xd4, 0x08, 0x41, 0x31, 0x71, 0x2d, 0xcb, 0xfc, 0x3a, 0xbf,
	0x39, 0x79, 0x4d, 0x19, 0x5f, 0x06, 0xc2, 0x50, 0x09, 0x47, 0x5a, 0x5a, 0xcb, 0xd5, 0x06, 0xb7,
	0x11, 0x1f, 0xa0, 0x78, 0x33, 0x43, 0xcf, 0xae, 0x60, 0xd7, 0xd6, 0xe8, 0x0c, 0xf2, 0x06, 0x1d,
	0x93, 0xa4, 0xd0, 0xd4, 0xd5, 0x96, 0x18, 0xce, 0x98, 0x0e, 0x0b, 0x0c, 0x5d, 0xc0, 0xfe, 0x59,
	0x3a, 0x04, 0x2b, 0x88, 0x3f, 0x7f, 0x96, 0x6a, 0x9c, 0x45, 0xdc, 0x06, 0xfc, 0xdb, 0xf0, 0x67,
	0x31, 0xe0, 0x5b, 0xd8, 0xbf, 0xe8, 0x39, 0xe0, 0xa8, 0xd5, 0x7a, 0x51, 0x9d, 0x92, 0xd5, 0x8a,
	0x6c, 0x56, 0xe4, 0x29, 0x5a, 0xe1, 0x4c, 0x1c, 0xd2, 0xcb, 0xfd, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x17, 0xd5, 0x22, 0xc1, 0x86, 0x01, 0x00, 0x00,
}
