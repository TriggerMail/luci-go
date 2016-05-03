// Code generated by protoc-gen-go.
// source: service_accounts.proto
// DO NOT EDIT!

package admin

import prpccommon "github.com/luci/luci-go/common/prpc"
import prpc "github.com/luci/luci-go/server/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tokenserver "github.com/luci/luci-go/common/api/tokenserver"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// CreateServiceAccountRequest is parameters for CreateServiceAccount call.
type CreateServiceAccountRequest struct {
	Ca    string `protobuf:"bytes,1,opt,name=ca" json:"ca,omitempty"`
	Fqdn  string `protobuf:"bytes,2,opt,name=fqdn" json:"fqdn,omitempty"`
	Force bool   `protobuf:"varint,3,opt,name=force" json:"force,omitempty"`
}

func (m *CreateServiceAccountRequest) Reset()                    { *m = CreateServiceAccountRequest{} }
func (m *CreateServiceAccountRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateServiceAccountRequest) ProtoMessage()               {}
func (*CreateServiceAccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// CreateServiceAccountResponse is returned by CreateServiceAccount call.
type CreateServiceAccountResponse struct {
	ServiceAccount *tokenserver.ServiceAccount `protobuf:"bytes,1,opt,name=service_account,json=serviceAccount" json:"service_account,omitempty"`
}

func (m *CreateServiceAccountResponse) Reset()                    { *m = CreateServiceAccountResponse{} }
func (m *CreateServiceAccountResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateServiceAccountResponse) ProtoMessage()               {}
func (*CreateServiceAccountResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *CreateServiceAccountResponse) GetServiceAccount() *tokenserver.ServiceAccount {
	if m != nil {
		return m.ServiceAccount
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateServiceAccountRequest)(nil), "tokenserver.admin.CreateServiceAccountRequest")
	proto.RegisterType((*CreateServiceAccountResponse)(nil), "tokenserver.admin.CreateServiceAccountResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for ServiceAccounts service

type ServiceAccountsClient interface {
	// CreateServiceAccount creates Google Cloud IAM service account associated
	// with given host.
	//
	// It uses token server configuration to pick a cloud project and to derive
	// service account ID. See documentation for CertificateAuthorityConfig proto
	// message for more info.
	//
	// This operation is idempotent.
	CreateServiceAccount(ctx context.Context, in *CreateServiceAccountRequest, opts ...grpc.CallOption) (*CreateServiceAccountResponse, error)
}
type serviceAccountsPRPCClient struct {
	client *prpccommon.Client
}

func NewServiceAccountsPRPCClient(client *prpccommon.Client) ServiceAccountsClient {
	return &serviceAccountsPRPCClient{client}
}

func (c *serviceAccountsPRPCClient) CreateServiceAccount(ctx context.Context, in *CreateServiceAccountRequest, opts ...grpc.CallOption) (*CreateServiceAccountResponse, error) {
	out := new(CreateServiceAccountResponse)
	err := c.client.Call(ctx, "tokenserver.admin.ServiceAccounts", "CreateServiceAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type serviceAccountsClient struct {
	cc *grpc.ClientConn
}

func NewServiceAccountsClient(cc *grpc.ClientConn) ServiceAccountsClient {
	return &serviceAccountsClient{cc}
}

func (c *serviceAccountsClient) CreateServiceAccount(ctx context.Context, in *CreateServiceAccountRequest, opts ...grpc.CallOption) (*CreateServiceAccountResponse, error) {
	out := new(CreateServiceAccountResponse)
	err := grpc.Invoke(ctx, "/tokenserver.admin.ServiceAccounts/CreateServiceAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ServiceAccounts service

type ServiceAccountsServer interface {
	// CreateServiceAccount creates Google Cloud IAM service account associated
	// with given host.
	//
	// It uses token server configuration to pick a cloud project and to derive
	// service account ID. See documentation for CertificateAuthorityConfig proto
	// message for more info.
	//
	// This operation is idempotent.
	CreateServiceAccount(context.Context, *CreateServiceAccountRequest) (*CreateServiceAccountResponse, error)
}

func RegisterServiceAccountsServer(s prpc.Registrar, srv ServiceAccountsServer) {
	s.RegisterService(&_ServiceAccounts_serviceDesc, srv)
}

func _ServiceAccounts_CreateServiceAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAccountsServer).CreateServiceAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenserver.admin.ServiceAccounts/CreateServiceAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAccountsServer).CreateServiceAccount(ctx, req.(*CreateServiceAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceAccounts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tokenserver.admin.ServiceAccounts",
	HandlerType: (*ServiceAccountsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateServiceAccount",
			Handler:    _ServiceAccounts_CreateServiceAccount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor2 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x8d, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x2c, 0xc9, 0xcf, 0x4e, 0xcd, 0x03, 0x49, 0xa6, 0x16, 0xe9, 0x25, 0xa6,
	0xe4, 0x66, 0xe6, 0x49, 0xb9, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea,
	0xe7, 0x94, 0x26, 0x67, 0x82, 0x09, 0xdd, 0xf4, 0x7c, 0x7d, 0xa0, 0x40, 0x6e, 0x7e, 0x9e, 0x7e,
	0x62, 0x41, 0xa6, 0x3e, 0x92, 0x2e, 0x7d, 0x34, 0x93, 0x21, 0x06, 0x2b, 0x85, 0x73, 0x49, 0x3b,
	0x17, 0xa5, 0x26, 0x96, 0xa4, 0x06, 0x43, 0xa4, 0x1d, 0x21, 0xb2, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x42, 0x7c, 0x5c, 0x4c, 0xc9, 0x89, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x40,
	0x96, 0x90, 0x10, 0x17, 0x4b, 0x5a, 0x61, 0x4a, 0x9e, 0x04, 0x13, 0x58, 0x04, 0xcc, 0x16, 0x12,
	0xe1, 0x62, 0x4d, 0xcb, 0x2f, 0x4a, 0x4e, 0x95, 0x60, 0x06, 0x0a, 0x72, 0x04, 0x41, 0x38, 0x4a,
	0x29, 0x5c, 0x32, 0xd8, 0x0d, 0x2e, 0x2e, 0xc8, 0x07, 0x3a, 0x49, 0xc8, 0x85, 0x8b, 0x1f, 0xcd,
	0x45, 0x60, 0x6b, 0xb8, 0x8d, 0xa4, 0xf5, 0x90, 0xfd, 0x8a, 0xa6, 0x9b, 0xaf, 0x18, 0x85, 0x6f,
	0xd4, 0xc5, 0xc8, 0xc5, 0x8f, 0xaa, 0xa4, 0x58, 0xa8, 0x9c, 0x4b, 0x04, 0x9b, 0xcd, 0x42, 0x7a,
	0x7a, 0x18, 0x81, 0xa8, 0x87, 0xc7, 0xef, 0x52, 0xfa, 0x44, 0xab, 0x87, 0x78, 0x29, 0x89, 0x0d,
	0x1c, 0xa4, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xf5, 0x3b, 0xad, 0xc5, 0x01, 0x00,
	0x00,
}
