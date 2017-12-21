// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/config/v1/vlans.proto

package config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// VLAN describes a virtual LAN.
type VLAN struct {
	// The ID of this virtual LAN.
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// An alias for this virtual LAN.
	Alias string `protobuf:"bytes,2,opt,name=alias" json:"alias,omitempty"`
	// The blocks of IPv4 addresses belonging to this virtual LAN.
	CidrBlock []string `protobuf:"bytes,3,rep,name=cidr_block,json=cidrBlock" json:"cidr_block,omitempty"`
}

func (m *VLAN) Reset()                    { *m = VLAN{} }
func (m *VLAN) String() string            { return proto.CompactTextString(m) }
func (*VLAN) ProtoMessage()               {}
func (*VLAN) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *VLAN) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *VLAN) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *VLAN) GetCidrBlock() []string {
	if m != nil {
		return m.CidrBlock
	}
	return nil
}

// VLANs enumerates virtual LANs.
type VLANs struct {
	// A list of virtual LANs.
	Vlan []*VLAN `protobuf:"bytes,1,rep,name=vlan" json:"vlan,omitempty"`
}

func (m *VLANs) Reset()                    { *m = VLANs{} }
func (m *VLANs) String() string            { return proto.CompactTextString(m) }
func (*VLANs) ProtoMessage()               {}
func (*VLANs) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *VLANs) GetVlan() []*VLAN {
	if m != nil {
		return m.Vlan
	}
	return nil
}

func init() {
	proto.RegisterType((*VLAN)(nil), "config.VLAN")
	proto.RegisterType((*VLANs)(nil), "config.VLANs")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/config/v1/vlans.proto", fileDescriptor3)
}

var fileDescriptor3 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x24, 0x8e, 0xb1, 0x4f, 0x84, 0x30,
	0x14, 0x87, 0x53, 0x0a, 0x24, 0x3c, 0x8d, 0x43, 0xe3, 0xd0, 0xc5, 0xa4, 0x61, 0xaa, 0x83, 0x6d,
	0xd4, 0xc9, 0x51, 0x57, 0x8d, 0x43, 0x07, 0x57, 0x53, 0xda, 0x3b, 0x78, 0xb9, 0x42, 0x49, 0x39,
	0xf8, 0xfb, 0x2f, 0x85, 0xf1, 0x7d, 0x5f, 0xf2, 0xbe, 0x1f, 0x7c, 0xf4, 0x51, 0xb9, 0x21, 0xc5,
	0x11, 0xd7, 0x51, 0xc5, 0xd4, 0xeb, 0xb0, 0x3a, 0xd4, 0xa3, 0x75, 0x03, 0x4e, 0xa7, 0x17, 0xdf,
	0x69, 0x3b, 0xa3, 0x76, 0x71, 0x3a, 0x63, 0xaf, 0xb7, 0x57, 0xbd, 0x05, 0x3b, 0x2d, 0x6a, 0x4e,
	0xf1, 0x1a, 0x59, 0x7d, 0xe0, 0xf6, 0x1b, 0xca, 0xbf, 0x9f, 0xcf, 0x5f, 0xf6, 0x00, 0x05, 0x7a,
	0x4e, 0x04, 0x91, 0xd4, 0x14, 0xe8, 0xd9, 0x23, 0x54, 0x36, 0xa0, 0x5d, 0x78, 0x21, 0x88, 0x6c,
	0xcc, 0x71, 0xb0, 0x27, 0x00, 0x87, 0x3e, 0xfd, 0x77, 0x21, 0xba, 0x0b, 0xa7, 0x82, 0xca, 0xc6,
	0x34, 0x99, 0x7c, 0x65, 0xd0, 0x3e, 0x43, 0x95, 0x9f, 0x2d, 0x4c, 0x40, 0x99, 0x63, 0x9c, 0x08,
	0x2a, 0xef, 0xde, 0xee, 0xd5, 0x11, 0x53, 0x59, 0x9a, 0xdd, 0x74, 0xf5, 0x3e, 0xe3, 0xfd, 0x16,
	0x00, 0x00, 0xff, 0xff, 0xa3, 0x3c, 0x7e, 0x37, 0xc3, 0x00, 0x00, 0x00,
}
