// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/TriggerMail/luci-go/common/tsmon/ts_mon_proto/acquisition_network_device.proto

package ts_mon_proto

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

type NetworkDevice_TypeId int32

const (
	NetworkDevice_MESSAGE_TYPE_ID NetworkDevice_TypeId = 34049749
)

var NetworkDevice_TypeId_name = map[int32]string{
	34049749: "MESSAGE_TYPE_ID",
}

var NetworkDevice_TypeId_value = map[string]int32{
	"MESSAGE_TYPE_ID": 34049749,
}

func (x NetworkDevice_TypeId) Enum() *NetworkDevice_TypeId {
	p := new(NetworkDevice_TypeId)
	*p = x
	return p
}

func (x NetworkDevice_TypeId) String() string {
	return proto.EnumName(NetworkDevice_TypeId_name, int32(x))
}

func (x *NetworkDevice_TypeId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NetworkDevice_TypeId_value, data, "NetworkDevice_TypeId")
	if err != nil {
		return err
	}
	*x = NetworkDevice_TypeId(value)
	return nil
}

func (NetworkDevice_TypeId) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_617c78492cfca97c, []int{0, 0}
}

type NetworkDevice struct {
	ProxyEnvironment     *string  `protobuf:"bytes,5,opt,name=proxy_environment,json=proxyEnvironment" json:"proxy_environment,omitempty"`
	AcquisitionName      *string  `protobuf:"bytes,10,opt,name=acquisition_name,json=acquisitionName" json:"acquisition_name,omitempty"`
	Pop                  *string  `protobuf:"bytes,30,opt,name=pop" json:"pop,omitempty"`
	Alertable            *bool    `protobuf:"varint,101,opt,name=alertable" json:"alertable,omitempty"`
	Realm                *string  `protobuf:"bytes,102,opt,name=realm" json:"realm,omitempty"`
	Asn                  *int64   `protobuf:"varint,103,opt,name=asn" json:"asn,omitempty"`
	Metro                *string  `protobuf:"bytes,104,opt,name=metro" json:"metro,omitempty"`
	Role                 *string  `protobuf:"bytes,105,opt,name=role" json:"role,omitempty"`
	Hostname             *string  `protobuf:"bytes,106,opt,name=hostname" json:"hostname,omitempty"`
	Vendor               *string  `protobuf:"bytes,70,opt,name=vendor" json:"vendor,omitempty"`
	Hostgroup            *string  `protobuf:"bytes,108,opt,name=hostgroup" json:"hostgroup,omitempty"`
	ProxyZone            *string  `protobuf:"bytes,100,opt,name=proxy_zone,json=proxyZone" json:"proxy_zone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkDevice) Reset()         { *m = NetworkDevice{} }
func (m *NetworkDevice) String() string { return proto.CompactTextString(m) }
func (*NetworkDevice) ProtoMessage()    {}
func (*NetworkDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_617c78492cfca97c, []int{0}
}

func (m *NetworkDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkDevice.Unmarshal(m, b)
}
func (m *NetworkDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkDevice.Marshal(b, m, deterministic)
}
func (m *NetworkDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkDevice.Merge(m, src)
}
func (m *NetworkDevice) XXX_Size() int {
	return xxx_messageInfo_NetworkDevice.Size(m)
}
func (m *NetworkDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkDevice.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkDevice proto.InternalMessageInfo

func (m *NetworkDevice) GetProxyEnvironment() string {
	if m != nil && m.ProxyEnvironment != nil {
		return *m.ProxyEnvironment
	}
	return ""
}

func (m *NetworkDevice) GetAcquisitionName() string {
	if m != nil && m.AcquisitionName != nil {
		return *m.AcquisitionName
	}
	return ""
}

func (m *NetworkDevice) GetPop() string {
	if m != nil && m.Pop != nil {
		return *m.Pop
	}
	return ""
}

func (m *NetworkDevice) GetAlertable() bool {
	if m != nil && m.Alertable != nil {
		return *m.Alertable
	}
	return false
}

func (m *NetworkDevice) GetRealm() string {
	if m != nil && m.Realm != nil {
		return *m.Realm
	}
	return ""
}

func (m *NetworkDevice) GetAsn() int64 {
	if m != nil && m.Asn != nil {
		return *m.Asn
	}
	return 0
}

func (m *NetworkDevice) GetMetro() string {
	if m != nil && m.Metro != nil {
		return *m.Metro
	}
	return ""
}

func (m *NetworkDevice) GetRole() string {
	if m != nil && m.Role != nil {
		return *m.Role
	}
	return ""
}

func (m *NetworkDevice) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *NetworkDevice) GetVendor() string {
	if m != nil && m.Vendor != nil {
		return *m.Vendor
	}
	return ""
}

func (m *NetworkDevice) GetHostgroup() string {
	if m != nil && m.Hostgroup != nil {
		return *m.Hostgroup
	}
	return ""
}

func (m *NetworkDevice) GetProxyZone() string {
	if m != nil && m.ProxyZone != nil {
		return *m.ProxyZone
	}
	return ""
}

func init() {
	proto.RegisterEnum("ts_mon.proto.NetworkDevice_TypeId", NetworkDevice_TypeId_name, NetworkDevice_TypeId_value)
	proto.RegisterType((*NetworkDevice)(nil), "ts_mon.proto.NetworkDevice")
}

func init() {
	proto.RegisterFile("github.com/TriggerMail/luci-go/common/tsmon/ts_mon_proto/acquisition_network_device.proto", fileDescriptor_617c78492cfca97c)
}

var fileDescriptor_617c78492cfca97c = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4e, 0xe3, 0x30,
	0x10, 0x86, 0xd5, 0xed, 0xb6, 0x6a, 0x47, 0xbb, 0x6a, 0xd6, 0x5a, 0x55, 0x16, 0x02, 0x14, 0xf5,
	0x54, 0x84, 0xd4, 0x3c, 0x03, 0x52, 0x03, 0xea, 0x81, 0x82, 0xda, 0x5e, 0xe0, 0x12, 0x99, 0x64,
	0x48, 0x0d, 0xb1, 0x27, 0x38, 0x4e, 0xa1, 0xbc, 0x0b, 0xbc, 0x15, 0x6f, 0xc3, 0x01, 0xc5, 0x46,
	0x14, 0x2e, 0xd1, 0xfc, 0xdf, 0xff, 0xc5, 0x1e, 0x19, 0x2e, 0x72, 0x9a, 0xa4, 0x6b, 0x43, 0x4a,
	0xd6, 0x6a, 0x42, 0x26, 0x8f, 0x8a, 0x3a, 0x95, 0x51, 0x4a, 0x4a, 0x91, 0x8e, 0x6c, 0xe5, 0xbf,
	0x89, 0x22, 0x9d, 0x94, 0x86, 0x2c, 0x45, 0x22, 0x7d, 0xa8, 0x65, 0x25, 0xad, 0x24, 0x9d, 0x68,
	0xb4, 0x8f, 0x64, 0xee, 0x93, 0x0c, 0x37, 0x32, 0xc5, 0x89, 0x13, 0xd8, 0x1f, 0xaf, 0xfb, 0x34,
	0x7a, 0xff, 0x05, 0x7f, 0xe7, 0x5e, 0x9b, 0x3a, 0x8b, 0x1d, 0xc3, 0xbf, 0xd2, 0xd0, 0xd3, 0x36,
	0x41, 0xbd, 0x91, 0x86, 0xb4, 0x42, 0x6d, 0x79, 0x27, 0x6c, 0x8d, 0xfb, 0x8b, 0xc0, 0x15, 0xf1,
	0x8e, 0xb3, 0x23, 0x08, 0x7e, 0x5c, 0x28, 0x14, 0x72, 0x70, 0xee, 0xe0, 0x1b, 0x9f, 0x0b, 0x85,
	0x2c, 0x80, 0x76, 0x49, 0x25, 0x3f, 0x74, 0x6d, 0x33, 0xb2, 0x7d, 0xe8, 0x8b, 0x02, 0x8d, 0x15,
	0x37, 0x05, 0x72, 0x0c, 0x5b, 0xe3, 0xde, 0x62, 0x07, 0xd8, 0x7f, 0xe8, 0x18, 0x14, 0x85, 0xe2,
	0xb7, 0xee, 0x0f, 0x1f, 0x9a, 0x53, 0x44, 0xa5, 0x79, 0x1e, 0xb6, 0xc6, 0xed, 0x45, 0x33, 0x36,
	0x9e, 0x42, 0x6b, 0x88, 0xaf, 0xbd, 0xe7, 0x02, 0x63, 0xf0, 0xdb, 0x50, 0x81, 0x5c, 0x3a, 0xe8,
	0x66, 0xb6, 0x07, 0xbd, 0x35, 0x55, 0xd6, 0x2d, 0x79, 0xe7, 0xf8, 0x57, 0x66, 0x43, 0xe8, 0x6e,
	0x50, 0x67, 0x64, 0xf8, 0xa9, 0x6b, 0x3e, 0x53, 0xb3, 0x63, 0xe3, 0xe4, 0x86, 0xea, 0x92, 0x17,
	0xae, 0xda, 0x01, 0x76, 0x00, 0xe0, 0xdf, 0xea, 0x99, 0x34, 0xf2, 0xcc, 0xd7, 0x8e, 0x5c, 0x93,
	0xc6, 0x51, 0x08, 0xdd, 0xd5, 0xb6, 0xc4, 0x59, 0xc6, 0x86, 0x30, 0x38, 0x8f, 0x97, 0xcb, 0x93,
	0xb3, 0x38, 0x59, 0x5d, 0x5d, 0xc6, 0xc9, 0x6c, 0x1a, 0xbc, 0xbd, 0xbc, 0x06, 0x1f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x97, 0x60, 0x65, 0x38, 0xde, 0x01, 0x00, 0x00,
}
