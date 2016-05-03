// Code generated by protoc-gen-go.
// source: machine_token.proto
// DO NOT EDIT!

/*
Package tokenserver is a generated protocol buffer package.

It is generated from these files:
	machine_token.proto
	service_account.proto
	token_file.proto

It has these top-level messages:
	MachineTokenBody
	MachineTokenEnvelope
	ServiceAccount
	TokenFile
*/
package tokenserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// MachineTokenBody describes internal structure of the machine token.
//
// The token will be put in HTTP headers and its body shouldn't be too large.
// For that reason we use unix timestamps instead of google.protobuf.Timestamp
// (no need for microsecond precision), and assume certificate serial numbers
// are smallish uint64 integers (not random blobs).
type MachineTokenBody struct {
	// Machine identity this token conveys.
	//
	// It has format <fqdn>@<tokenserver>, where <fqdn> is FQDN of the
	// machine and <tokenserver> is URL of the server that generated the token.
	//
	// <fqdn> is Common Name of a certificate used as a basis for the token.
	//
	// For example: vm123-m4.golo.chromium.org@luci-token-server.appspot.com.
	MachineId string `protobuf:"bytes,1,opt,name=machine_id,json=machineId" json:"machine_id,omitempty"`
	// Unix timestamp in seconds when this token was issued. Required.
	IssuedAt uint64 `protobuf:"varint,2,opt,name=issued_at,json=issuedAt" json:"issued_at,omitempty"`
	// Number of seconds the token is considered valid.
	//
	// Usually 3600. Set by the token server. Required.
	Lifetime uint64 `protobuf:"varint,3,opt,name=lifetime" json:"lifetime,omitempty"`
	// Id of a CA that issued machine certificate used to make this token.
	//
	// These IDs are defined in token server config (via unique_id field).
	CaId int64 `protobuf:"varint,4,opt,name=ca_id,json=caId" json:"ca_id,omitempty"`
	// Serial number of the machine certificate used to make this token.
	//
	// ca_id and cert_sn together uniquely identify the certificate, and can be
	// used to check for certificate revocation (by asking token server whether
	// the given certificate is in CRL). Revocation checks are optional, most
	// callers can rely on expiration checks only.
	CertSn uint64 `protobuf:"varint,5,opt,name=cert_sn,json=certSn" json:"cert_sn,omitempty"`
}

func (m *MachineTokenBody) Reset()                    { *m = MachineTokenBody{} }
func (m *MachineTokenBody) String() string            { return proto.CompactTextString(m) }
func (*MachineTokenBody) ProtoMessage()               {}
func (*MachineTokenBody) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// MachineTokenEnvelope is what is actually being serialized and represented
// as a machine token (after being encoded using base64 standard raw encoding).
//
// Resulting token (including base64 encoding) is usually ~500 bytes long.
type MachineTokenEnvelope struct {
	TokenBody []byte `protobuf:"bytes,1,opt,name=token_body,json=tokenBody,proto3" json:"token_body,omitempty"`
	KeyId     string `protobuf:"bytes,2,opt,name=key_id,json=keyId" json:"key_id,omitempty"`
	RsaSha256 []byte `protobuf:"bytes,3,opt,name=rsa_sha256,json=rsaSha256,proto3" json:"rsa_sha256,omitempty"`
}

func (m *MachineTokenEnvelope) Reset()                    { *m = MachineTokenEnvelope{} }
func (m *MachineTokenEnvelope) String() string            { return proto.CompactTextString(m) }
func (*MachineTokenEnvelope) ProtoMessage()               {}
func (*MachineTokenEnvelope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*MachineTokenBody)(nil), "tokenserver.MachineTokenBody")
	proto.RegisterType((*MachineTokenEnvelope)(nil), "tokenserver.MachineTokenEnvelope")
}

var fileDescriptor0 = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0x36, 0x09, 0xcd, 0xc1, 0x80, 0x5c, 0x10, 0x11, 0x08, 0x09, 0x75, 0x62, 0x62,
	0x00, 0xc1, 0x0e, 0x12, 0x03, 0x03, 0x4b, 0xca, 0x6e, 0xb9, 0xf6, 0xa1, 0x5a, 0x69, 0xed, 0xca,
	0x36, 0x95, 0xfa, 0x4f, 0xf8, 0xb9, 0x9c, 0x2f, 0x04, 0x31, 0xbe, 0xf7, 0xee, 0xdd, 0x7d, 0x36,
	0xcc, 0xb7, 0x4a, 0xaf, 0xad, 0x43, 0x99, 0x7c, 0x8f, 0xee, 0x6e, 0x17, 0x7c, 0xf2, 0xe2, 0x98,
	0x45, 0xc4, 0xb0, 0xc7, 0xb0, 0xf8, 0x2e, 0xe0, 0xf4, 0x7d, 0x18, 0xfa, 0xc8, 0xf6, 0x8b, 0x37,
	0x07, 0x71, 0x0d, 0x30, 0x16, 0xad, 0x69, 0x8b, 0x9b, 0xe2, 0xb6, 0xe9, 0x9a, 0x5f, 0xe7, 0xcd,
	0x88, 0x2b, 0x68, 0x6c, 0x8c, 0x5f, 0x68, 0xa4, 0x4a, 0xed, 0x84, 0xd2, 0xb2, 0x9b, 0x0d, 0xc6,
	0x73, 0x12, 0x97, 0x30, 0xdb, 0xd8, 0x4f, 0x4c, 0x76, 0x8b, 0xed, 0x74, 0xc8, 0x46, 0x2d, 0xe6,
	0x50, 0x69, 0x95, 0x57, 0x96, 0x14, 0x4c, 0xbb, 0x52, 0x2b, 0xda, 0x76, 0x01, 0x47, 0x1a, 0x43,
	0x92, 0xd1, 0xb5, 0x15, 0xcf, 0xd7, 0x59, 0x2e, 0xdd, 0xa2, 0x87, 0xb3, 0xff, 0x64, 0xaf, 0x6e,
	0x8f, 0x1b, 0xbf, 0xc3, 0x4c, 0xc7, 0x2f, 0x90, 0x2b, 0x62, 0x65, 0xba, 0x93, 0xae, 0x49, 0x7f,
	0xf0, 0xe7, 0x50, 0xf7, 0x78, 0xc8, 0x57, 0x26, 0x0c, 0x5e, 0x91, 0xa2, 0x33, 0xd4, 0x0a, 0x51,
	0xc9, 0xb8, 0x56, 0xf7, 0x8f, 0x4f, 0x4c, 0x46, 0x2d, 0x72, 0x96, 0x6c, 0xac, 0x6a, 0xfe, 0x9b,
	0x87, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0x5e, 0x8a, 0x6a, 0x32, 0x01, 0x00, 0x00,
}
