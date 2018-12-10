// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/examples/appengine/helloworld_standard/proto/helloworld.proto

package helloworld

import prpc "github.com/TriggerMail/luci-go/grpc/prpc"

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a45dd7ca0973b79, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a45dd7ca0973b79, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

func init() {
	proto.RegisterFile("github.com/TriggerMail/luci-go/examples/appengine/helloworld_standard/proto/helloworld.proto", fileDescriptor_2a45dd7ca0973b79)
}

var fileDescriptor_2a45dd7ca0973b79 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8e, 0x3d, 0x0b, 0xc2, 0x30,
	0x10, 0x86, 0x2d, 0x88, 0xd5, 0xc3, 0x29, 0x83, 0x14, 0x27, 0xc9, 0x20, 0x4e, 0x09, 0xe8, 0xee,
	0xaa, 0xae, 0xf5, 0x07, 0x48, 0x6c, 0x8f, 0xb4, 0x90, 0x2f, 0x93, 0x14, 0xed, 0xbf, 0x17, 0x83,
	0xc5, 0x0e, 0x6e, 0xf7, 0x7e, 0x1c, 0xef, 0x03, 0xa5, 0xb4, 0xac, 0x6a, 0xbc, 0xd5, 0x6d, 0xa7,
	0x99, 0xf5, 0x92, 0xab, 0xae, 0x6a, 0x39, 0xbe, 0x84, 0x76, 0x0a, 0x03, 0x17, 0xce, 0xa1, 0x91,
	0xad, 0x41, 0xde, 0xa0, 0x52, 0xf6, 0x69, 0xbd, 0xaa, 0x6f, 0x21, 0x0a, 0x53, 0x0b, 0x5f, 0x73,
	0xe7, 0x6d, 0xb4, 0xa3, 0x84, 0x25, 0x83, 0xc0, 0xcf, 0xa1, 0x14, 0x96, 0xe7, 0x8f, 0x2a, 0xf1,
	0xd1, 0x61, 0x88, 0x84, 0xc0, 0xd4, 0x08, 0x8d, 0x45, 0xb6, 0xc9, 0x76, 0x8b, 0x32, 0xdd, 0x74,
	0x0b, 0xf0, 0xed, 0x38, 0xd5, 0x93, 0x02, 0x72, 0x8d, 0x21, 0x08, 0x39, 0x94, 0x06, 0xb9, 0xbf,
	0x40, 0x7e, 0xf2, 0x88, 0x11, 0x3d, 0x39, 0xc2, 0xfc, 0x2a, 0xfa, 0xf4, 0x45, 0x0a, 0x36, 0x22,
	0x18, 0x8f, 0xad, 0x57, 0x7f, 0x12, 0xa7, 0x7a, 0x3a, 0xb9, 0xcf, 0x12, 0xe9, 0xe1, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x46, 0xcf, 0xff, 0x05, 0xff, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}
type greeterPRPCClient struct {
	client *prpc.Client
}

func NewGreeterPRPCClient(client *prpc.Client) GreeterClient {
	return &greeterPRPCClient{client}
}

func (c *greeterPRPCClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.client.Call(ctx, "helloworld.Greeter", "SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func RegisterGreeterServer(s prpc.Registrar, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/TriggerMail/luci-go/examples/appengine/helloworld_standard/proto/helloworld.proto",
}
