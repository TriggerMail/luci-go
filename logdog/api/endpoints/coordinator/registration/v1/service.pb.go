// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/logdog/api/endpoints/coordinator/registration/v1/service.proto

package logdog

import prpc "github.com/TriggerMail/luci-go/grpc/prpc"

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// RegisterPrefixRequest registers a new Prefix with the Coordinator.
type RegisterPrefixRequest struct {
	// The log stream's project.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// The log stream prefix to register.
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// Optional information about the registering agent.
	SourceInfo []string `protobuf:"bytes,3,rep,name=source_info,json=sourceInfo,proto3" json:"source_info,omitempty"`
	// Optional nonce to allow retries of this RPC. ALL CLIENTS SHOULD PROVIDE
	// THIS. The client should generate the nonce once while preparing the request
	// message, and then re-use the same nonce for retries of the request.
	//
	// The nonce should be 32 bytes of random data.
	// The nonce must not be reused between different requests (only for retries
	//   of the same request).
	//
	// NOTE: This is currently optional, but once all clients have upgraded to
	// this scheme, it will become mandatory. During the transition if this is
	// omitted, then NO RETRIES will be allowed for this request, if the server
	// processes it correctly but the client fails to get the response from the
	// server.
	OpNonce []byte `protobuf:"bytes,4,opt,name=op_nonce,json=opNonce,proto3" json:"op_nonce,omitempty"`
	// The prefix expiration time. If <= 0, the project's default prefix
	// expiration period will be applied.
	//
	// The prefix will be closed by the Coordinator after its expiration period.
	// Once closed, new stream registration requests will no longer be accepted.
	//
	// If supplied, this value should exceed the timeout of the local task, else
	// some of the task's streams may be dropped due to failing registration.
	Expiration           *duration.Duration `protobuf:"bytes,10,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RegisterPrefixRequest) Reset()         { *m = RegisterPrefixRequest{} }
func (m *RegisterPrefixRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterPrefixRequest) ProtoMessage()    {}
func (*RegisterPrefixRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_317a8d00c4420d83, []int{0}
}

func (m *RegisterPrefixRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterPrefixRequest.Unmarshal(m, b)
}
func (m *RegisterPrefixRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterPrefixRequest.Marshal(b, m, deterministic)
}
func (m *RegisterPrefixRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterPrefixRequest.Merge(m, src)
}
func (m *RegisterPrefixRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterPrefixRequest.Size(m)
}
func (m *RegisterPrefixRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterPrefixRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterPrefixRequest proto.InternalMessageInfo

func (m *RegisterPrefixRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *RegisterPrefixRequest) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *RegisterPrefixRequest) GetSourceInfo() []string {
	if m != nil {
		return m.SourceInfo
	}
	return nil
}

func (m *RegisterPrefixRequest) GetOpNonce() []byte {
	if m != nil {
		return m.OpNonce
	}
	return nil
}

func (m *RegisterPrefixRequest) GetExpiration() *duration.Duration {
	if m != nil {
		return m.Expiration
	}
	return nil
}

// The response message for the RegisterPrefix RPC.
type RegisterPrefixResponse struct {
	// Secret is the prefix's secret. This must be included verbatim in Butler
	// bundles to assert ownership of this prefix.
	Secret []byte `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	// The name of the Pub/Sub topic to publish butlerproto-formatted Butler log
	// bundles to.
	LogBundleTopic       string   `protobuf:"bytes,2,opt,name=log_bundle_topic,json=logBundleTopic,proto3" json:"log_bundle_topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterPrefixResponse) Reset()         { *m = RegisterPrefixResponse{} }
func (m *RegisterPrefixResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterPrefixResponse) ProtoMessage()    {}
func (*RegisterPrefixResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_317a8d00c4420d83, []int{1}
}

func (m *RegisterPrefixResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterPrefixResponse.Unmarshal(m, b)
}
func (m *RegisterPrefixResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterPrefixResponse.Marshal(b, m, deterministic)
}
func (m *RegisterPrefixResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterPrefixResponse.Merge(m, src)
}
func (m *RegisterPrefixResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterPrefixResponse.Size(m)
}
func (m *RegisterPrefixResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterPrefixResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterPrefixResponse proto.InternalMessageInfo

func (m *RegisterPrefixResponse) GetSecret() []byte {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (m *RegisterPrefixResponse) GetLogBundleTopic() string {
	if m != nil {
		return m.LogBundleTopic
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterPrefixRequest)(nil), "logdog.RegisterPrefixRequest")
	proto.RegisterType((*RegisterPrefixResponse)(nil), "logdog.RegisterPrefixResponse")
}

func init() {
	proto.RegisterFile("github.com/TriggerMail/luci-go/logdog/api/endpoints/coordinator/registration/v1/service.proto", fileDescriptor_317a8d00c4420d83)
}

var fileDescriptor_317a8d00c4420d83 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4f, 0x2a, 0x31,
	0x14, 0x85, 0x33, 0x8f, 0x17, 0x78, 0x14, 0x42, 0x5e, 0x9a, 0x48, 0x06, 0x12, 0x71, 0xc2, 0x6a,
	0x56, 0x6d, 0xc4, 0x95, 0x5b, 0xe3, 0xc6, 0x8d, 0x9a, 0xea, 0xca, 0xcd, 0x04, 0x3a, 0x77, 0x6a,
	0xcd, 0xd0, 0x5b, 0xdb, 0x0e, 0xe1, 0xe7, 0xf9, 0xd3, 0x0c, 0xd3, 0x21, 0x41, 0xa3, 0xcb, 0x73,
	0xee, 0xbd, 0xed, 0x77, 0x0e, 0x79, 0x52, 0xc8, 0xe4, 0xab, 0xc3, 0xad, 0x6e, 0xb6, 0x0c, 0x9d,
	0xe2, 0x75, 0x23, 0x35, 0xaf, 0x51, 0x95, 0xa8, 0xf8, 0xda, 0x6a, 0x0e, 0xa6, 0xb4, 0xa8, 0x4d,
	0xf0, 0x5c, 0x22, 0xba, 0x52, 0x9b, 0x75, 0x40, 0xc7, 0x1d, 0x28, 0xed, 0x83, 0x5b, 0x07, 0x8d,
	0x86, 0xef, 0x2e, 0xb9, 0x07, 0xb7, 0xd3, 0x12, 0x98, 0x75, 0x18, 0x90, 0xf6, 0xe3, 0xfd, 0x7c,
	0xa1, 0x10, 0x55, 0x0d, 0xbc, 0x75, 0x37, 0x4d, 0xc5, 0xcb, 0x26, 0x9e, 0xc4, 0xbd, 0xe5, 0x47,
	0x42, 0xce, 0x44, 0xfb, 0x12, 0xb8, 0x47, 0x07, 0x95, 0xde, 0x0b, 0x78, 0x6f, 0xc0, 0x07, 0x9a,
	0x92, 0x81, 0x75, 0xf8, 0x06, 0x32, 0xa4, 0x49, 0x96, 0xe4, 0x43, 0x71, 0x94, 0x74, 0x4a, 0xfa,
	0xb6, 0x5d, 0x4d, 0xff, 0xb4, 0x83, 0x4e, 0xd1, 0x0b, 0x32, 0xf2, 0xd8, 0x38, 0x09, 0x85, 0x36,
	0x15, 0xa6, 0xbd, 0xac, 0x97, 0x0f, 0x05, 0x89, 0xd6, 0x9d, 0xa9, 0x90, 0xce, 0xc8, 0x3f, 0xb4,
	0x85, 0x41, 0x23, 0x21, 0xfd, 0x9b, 0x25, 0xf9, 0x58, 0x0c, 0xd0, 0xde, 0x1f, 0x24, 0xbd, 0x26,
	0x04, 0xf6, 0x56, 0x47, 0xb6, 0x94, 0x64, 0x49, 0x3e, 0x5a, 0xcd, 0x58, 0x84, 0x67, 0x47, 0x78,
	0x76, 0xdb, 0xc1, 0x8b, 0x93, 0xe5, 0xe5, 0x0b, 0x99, 0x7e, 0x4f, 0xe0, 0x2d, 0x1a, 0x0f, 0x07,
	0x50, 0x0f, 0xd2, 0x41, 0x4c, 0x30, 0x16, 0x9d, 0xa2, 0x39, 0xf9, 0x5f, 0xa3, 0x2a, 0x36, 0x8d,
	0x29, 0x6b, 0x28, 0x02, 0x5a, 0x2d, 0xbb, 0x28, 0x93, 0x1a, 0xd5, 0x4d, 0x6b, 0x3f, 0x1f, 0xdc,
	0x55, 0x41, 0xc6, 0xe2, 0xa4, 0x67, 0xfa, 0x40, 0x26, 0x5f, 0xff, 0xa2, 0xe7, 0x2c, 0x36, 0xcd,
	0x7e, 0x6c, 0x71, 0xbe, 0xf8, 0x6d, 0x1c, 0x11, 0x37, 0xfd, 0x36, 0xdb, 0xd5, 0x67, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x20, 0x5e, 0xf9, 0xf8, 0x05, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RegistrationClient is the client API for Registration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegistrationClient interface {
	// RegisterStream allows a Butler instance to register a log stream with the
	// Coordinator. Upon success, the Coordinator will return registration
	// information and streaming parameters to the Butler.
	//
	// This should be called by a Butler instance to gain the ability to publish
	// to a prefix space. The caller must have WRITE access to its project's
	// stream space. If WRITE access is not present, this will fail with the
	// "PermissionDenied" gRPC code.
	//
	// A stream prefix may be registered at most once. Additional registration
	// requests will fail with the "AlreadyExists" gRPC code.
	RegisterPrefix(ctx context.Context, in *RegisterPrefixRequest, opts ...grpc.CallOption) (*RegisterPrefixResponse, error)
}
type registrationPRPCClient struct {
	client *prpc.Client
}

func NewRegistrationPRPCClient(client *prpc.Client) RegistrationClient {
	return &registrationPRPCClient{client}
}

func (c *registrationPRPCClient) RegisterPrefix(ctx context.Context, in *RegisterPrefixRequest, opts ...grpc.CallOption) (*RegisterPrefixResponse, error) {
	out := new(RegisterPrefixResponse)
	err := c.client.Call(ctx, "logdog.Registration", "RegisterPrefix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type registrationClient struct {
	cc *grpc.ClientConn
}

func NewRegistrationClient(cc *grpc.ClientConn) RegistrationClient {
	return &registrationClient{cc}
}

func (c *registrationClient) RegisterPrefix(ctx context.Context, in *RegisterPrefixRequest, opts ...grpc.CallOption) (*RegisterPrefixResponse, error) {
	out := new(RegisterPrefixResponse)
	err := c.cc.Invoke(ctx, "/logdog.Registration/RegisterPrefix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistrationServer is the server API for Registration service.
type RegistrationServer interface {
	// RegisterStream allows a Butler instance to register a log stream with the
	// Coordinator. Upon success, the Coordinator will return registration
	// information and streaming parameters to the Butler.
	//
	// This should be called by a Butler instance to gain the ability to publish
	// to a prefix space. The caller must have WRITE access to its project's
	// stream space. If WRITE access is not present, this will fail with the
	// "PermissionDenied" gRPC code.
	//
	// A stream prefix may be registered at most once. Additional registration
	// requests will fail with the "AlreadyExists" gRPC code.
	RegisterPrefix(context.Context, *RegisterPrefixRequest) (*RegisterPrefixResponse, error)
}

func RegisterRegistrationServer(s prpc.Registrar, srv RegistrationServer) {
	s.RegisterService(&_Registration_serviceDesc, srv)
}

func _Registration_RegisterPrefix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterPrefixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).RegisterPrefix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logdog.Registration/RegisterPrefix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).RegisterPrefix(ctx, req.(*RegisterPrefixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Registration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logdog.Registration",
	HandlerType: (*RegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterPrefix",
			Handler:    _Registration_RegisterPrefix_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/TriggerMail/luci-go/logdog/api/endpoints/coordinator/registration/v1/service.proto",
}
