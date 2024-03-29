// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/dracs.proto

package crimson

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

// A DRAC in the database.
type DRAC struct {
	// The name of this DRAC on the network. Uniquely identifies this DRAC.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The machine this DRAC belongs to. Uniquely identifies this DRAC.
	Machine string `protobuf:"bytes,2,opt,name=machine,proto3" json:"machine,omitempty"`
	// The IPv4 address associated with this DRAC.
	Ipv4 string `protobuf:"bytes,3,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	// The VLAN this DRAC belongs to.
	// When creating a DRAC, omit this field. It will be inferred from the IPv4 address.
	Vlan int64 `protobuf:"varint,4,opt,name=vlan,proto3" json:"vlan,omitempty"`
	// The MAC address associated with this DRAC.
	MacAddress string `protobuf:"bytes,5,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	// The switch this DRAC is connected to.
	Switch string `protobuf:"bytes,6,opt,name=switch,proto3" json:"switch,omitempty"`
	// The switchport this DRAC is connected to.
	Switchport           int32    `protobuf:"varint,7,opt,name=switchport,proto3" json:"switchport,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DRAC) Reset()         { *m = DRAC{} }
func (m *DRAC) String() string { return proto.CompactTextString(m) }
func (*DRAC) ProtoMessage()    {}
func (*DRAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e7c5f5f1ce4adfa, []int{0}
}

func (m *DRAC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DRAC.Unmarshal(m, b)
}
func (m *DRAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DRAC.Marshal(b, m, deterministic)
}
func (m *DRAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DRAC.Merge(m, src)
}
func (m *DRAC) XXX_Size() int {
	return xxx_messageInfo_DRAC.Size(m)
}
func (m *DRAC) XXX_DiscardUnknown() {
	xxx_messageInfo_DRAC.DiscardUnknown(m)
}

var xxx_messageInfo_DRAC proto.InternalMessageInfo

func (m *DRAC) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DRAC) GetMachine() string {
	if m != nil {
		return m.Machine
	}
	return ""
}

func (m *DRAC) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *DRAC) GetVlan() int64 {
	if m != nil {
		return m.Vlan
	}
	return 0
}

func (m *DRAC) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *DRAC) GetSwitch() string {
	if m != nil {
		return m.Switch
	}
	return ""
}

func (m *DRAC) GetSwitchport() int32 {
	if m != nil {
		return m.Switchport
	}
	return 0
}

// A request to create a new DRAC in the database.
type CreateDRACRequest struct {
	// The DRAC to create in the database.
	Drac                 *DRAC    `protobuf:"bytes,1,opt,name=drac,proto3" json:"drac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDRACRequest) Reset()         { *m = CreateDRACRequest{} }
func (m *CreateDRACRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDRACRequest) ProtoMessage()    {}
func (*CreateDRACRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e7c5f5f1ce4adfa, []int{1}
}

func (m *CreateDRACRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDRACRequest.Unmarshal(m, b)
}
func (m *CreateDRACRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDRACRequest.Marshal(b, m, deterministic)
}
func (m *CreateDRACRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDRACRequest.Merge(m, src)
}
func (m *CreateDRACRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDRACRequest.Size(m)
}
func (m *CreateDRACRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDRACRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDRACRequest proto.InternalMessageInfo

func (m *CreateDRACRequest) GetDrac() *DRAC {
	if m != nil {
		return m.Drac
	}
	return nil
}

// A request to list DRACs in the database.
type ListDRACsRequest struct {
	// The names of DRACs to get.
	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	// The machines to filter retrieved DRACs on.
	Machines []string `protobuf:"bytes,2,rep,name=machines,proto3" json:"machines,omitempty"`
	// The IPv4 addresses to filter retrieved DRACs on.
	Ipv4S []string `protobuf:"bytes,3,rep,name=ipv4s,proto3" json:"ipv4s,omitempty"`
	// The VLANs to filter retrieved DRACs on.
	Vlans []int64 `protobuf:"varint,4,rep,packed,name=vlans,proto3" json:"vlans,omitempty"`
	// The MAC addresses to filter retrieved DRACs on.
	MacAddresses []string `protobuf:"bytes,5,rep,name=mac_addresses,json=macAddresses,proto3" json:"mac_addresses,omitempty"`
	// The switches to filter retrieved DRACs on.
	Switches             []string `protobuf:"bytes,6,rep,name=switches,proto3" json:"switches,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDRACsRequest) Reset()         { *m = ListDRACsRequest{} }
func (m *ListDRACsRequest) String() string { return proto.CompactTextString(m) }
func (*ListDRACsRequest) ProtoMessage()    {}
func (*ListDRACsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e7c5f5f1ce4adfa, []int{2}
}

func (m *ListDRACsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDRACsRequest.Unmarshal(m, b)
}
func (m *ListDRACsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDRACsRequest.Marshal(b, m, deterministic)
}
func (m *ListDRACsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDRACsRequest.Merge(m, src)
}
func (m *ListDRACsRequest) XXX_Size() int {
	return xxx_messageInfo_ListDRACsRequest.Size(m)
}
func (m *ListDRACsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDRACsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDRACsRequest proto.InternalMessageInfo

func (m *ListDRACsRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *ListDRACsRequest) GetMachines() []string {
	if m != nil {
		return m.Machines
	}
	return nil
}

func (m *ListDRACsRequest) GetIpv4S() []string {
	if m != nil {
		return m.Ipv4S
	}
	return nil
}

func (m *ListDRACsRequest) GetVlans() []int64 {
	if m != nil {
		return m.Vlans
	}
	return nil
}

func (m *ListDRACsRequest) GetMacAddresses() []string {
	if m != nil {
		return m.MacAddresses
	}
	return nil
}

func (m *ListDRACsRequest) GetSwitches() []string {
	if m != nil {
		return m.Switches
	}
	return nil
}

// A response containing a list of DRACs in the database.
type ListDRACsResponse struct {
	// The DRACs matching this request.
	Dracs                []*DRAC  `protobuf:"bytes,1,rep,name=dracs,proto3" json:"dracs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDRACsResponse) Reset()         { *m = ListDRACsResponse{} }
func (m *ListDRACsResponse) String() string { return proto.CompactTextString(m) }
func (*ListDRACsResponse) ProtoMessage()    {}
func (*ListDRACsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e7c5f5f1ce4adfa, []int{3}
}

func (m *ListDRACsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDRACsResponse.Unmarshal(m, b)
}
func (m *ListDRACsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDRACsResponse.Marshal(b, m, deterministic)
}
func (m *ListDRACsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDRACsResponse.Merge(m, src)
}
func (m *ListDRACsResponse) XXX_Size() int {
	return xxx_messageInfo_ListDRACsResponse.Size(m)
}
func (m *ListDRACsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDRACsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDRACsResponse proto.InternalMessageInfo

func (m *ListDRACsResponse) GetDracs() []*DRAC {
	if m != nil {
		return m.Dracs
	}
	return nil
}

// A request to update a DRAC in the database.
type UpdateDRACRequest struct {
	// The DRAC to update in the database.
	Drac *DRAC `protobuf:"bytes,1,opt,name=drac,proto3" json:"drac,omitempty"`
	// The fields to update in the DRAC.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateDRACRequest) Reset()         { *m = UpdateDRACRequest{} }
func (m *UpdateDRACRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateDRACRequest) ProtoMessage()    {}
func (*UpdateDRACRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e7c5f5f1ce4adfa, []int{4}
}

func (m *UpdateDRACRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDRACRequest.Unmarshal(m, b)
}
func (m *UpdateDRACRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDRACRequest.Marshal(b, m, deterministic)
}
func (m *UpdateDRACRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDRACRequest.Merge(m, src)
}
func (m *UpdateDRACRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateDRACRequest.Size(m)
}
func (m *UpdateDRACRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDRACRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDRACRequest proto.InternalMessageInfo

func (m *UpdateDRACRequest) GetDrac() *DRAC {
	if m != nil {
		return m.Drac
	}
	return nil
}

func (m *UpdateDRACRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func init() {
	proto.RegisterType((*DRAC)(nil), "crimson.DRAC")
	proto.RegisterType((*CreateDRACRequest)(nil), "crimson.CreateDRACRequest")
	proto.RegisterType((*ListDRACsRequest)(nil), "crimson.ListDRACsRequest")
	proto.RegisterType((*ListDRACsResponse)(nil), "crimson.ListDRACsResponse")
	proto.RegisterType((*UpdateDRACRequest)(nil), "crimson.UpdateDRACRequest")
}

func init() {
	proto.RegisterFile("github.com/TriggerMail/luci-go/machine-db/api/crimson/v1/dracs.proto", fileDescriptor_0e7c5f5f1ce4adfa)
}

var fileDescriptor_0e7c5f5f1ce4adfa = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x55, 0x36, 0x69, 0xcb, 0x4e, 0x59, 0x89, 0x5a, 0x08, 0x59, 0x3d, 0x40, 0xc8, 0x5e, 0x72,
	0x21, 0x16, 0x0b, 0x42, 0x08, 0x4e, 0xab, 0x45, 0x9c, 0xe0, 0x62, 0x89, 0xf3, 0xca, 0x75, 0xbc,
	0xad, 0xd5, 0x3a, 0x0e, 0x9e, 0xa4, 0x7c, 0x17, 0x07, 0xfe, 0x0f, 0x79, 0x9c, 0x42, 0x25, 0x4e,
	0xdc, 0xe6, 0xbd, 0x79, 0x23, 0xcf, 0x7b, 0x1e, 0xf8, 0xb0, 0xf5, 0x8d, 0xde, 0x05, 0xef, 0xec,
	0xe8, 0x1a, 0x1f, 0xb6, 0xe2, 0x30, 0x6a, 0x2b, 0x9c, 0xd2, 0x3b, 0xdb, 0x99, 0x57, 0xed, 0x46,
	0xa8, 0xde, 0x0a, 0x1d, 0xac, 0x43, 0xdf, 0x89, 0xe3, 0x6b, 0xd1, 0x06, 0xa5, 0xb1, 0xe9, 0x83,
	0x1f, 0x3c, 0x5b, 0x4c, 0xfc, 0xba, 0xdc, 0x7a, 0xbf, 0x3d, 0x18, 0x41, 0xf4, 0x66, 0x7c, 0x10,
	0x0f, 0xd6, 0x1c, 0xda, 0x7b, 0xa7, 0x70, 0x9f, 0xa4, 0xd5, 0xaf, 0x0c, 0x8a, 0x4f, 0xf2, 0xf6,
	0x8e, 0x31, 0x28, 0x3a, 0xe5, 0x0c, 0xcf, 0xca, 0xac, 0xbe, 0x94, 0x54, 0x33, 0x0e, 0x8b, 0xe9,
	0x41, 0x7e, 0x41, 0xf4, 0x09, 0x46, 0xb5, 0xed, 0x8f, 0x6f, 0x79, 0x9e, 0xd4, 0xb1, 0x8e, 0xdc,
	0xf1, 0xa0, 0x3a, 0x5e, 0x94, 0x59, 0x9d, 0x4b, 0xaa, 0xd9, 0x0b, 0x58, 0x3a, 0xa5, 0xef, 0x55,
	0xdb, 0x06, 0x83, 0xc8, 0x67, 0x24, 0x07, 0xa7, 0xf4, 0x6d, 0x62, 0xd8, 0x33, 0x98, 0xe3, 0x0f,
	0x3b, 0xe8, 0x1d, 0x9f, 0x53, 0x6f, 0x42, 0xec, 0x39, 0x40, 0xaa, 0x7a, 0x1f, 0x06, 0xbe, 0x28,
	0xb3, 0x7a, 0x26, 0xcf, 0x98, 0xea, 0x1d, 0xac, 0xee, 0x82, 0x51, 0x83, 0x89, 0xcb, 0x4b, 0xf3,
	0x7d, 0x34, 0x38, 0xb0, 0x97, 0x50, 0xc4, 0x18, 0xc8, 0xc3, 0xf2, 0xe6, 0xaa, 0x99, 0x62, 0x68,
	0x48, 0x43, 0xad, 0xea, 0x67, 0x06, 0x4f, 0xbe, 0x58, 0x1c, 0x22, 0x85, 0xa7, 0xb9, 0xa7, 0x30,
	0x8b, 0x7e, 0x91, 0x67, 0x65, 0x5e, 0x5f, 0xca, 0x04, 0xd8, 0x1a, 0x1e, 0x4d, 0x76, 0x91, 0x5f,
	0x50, 0xe3, 0x0f, 0x8e, 0x13, 0xd1, 0x33, 0xf2, 0x3c, 0x4d, 0x10, 0x88, 0x6c, 0x74, 0x8d, 0xbc,
	0x28, 0xf3, 0x3a, 0x97, 0x09, 0xb0, 0x6b, 0xb8, 0x3a, 0xcb, 0xc0, 0xc4, 0x14, 0xe2, 0xcc, 0xe3,
	0xbf, 0x29, 0xa4, 0xc7, 0x92, 0x3b, 0x83, 0x7c, 0x9e, 0x1e, 0x3b, 0xe1, 0xea, 0x3d, 0xac, 0xce,
	0x56, 0xc6, 0xde, 0x77, 0x68, 0xd8, 0x35, 0xcc, 0xe8, 0xcb, 0x69, 0xe7, 0x7f, 0xcc, 0xa6, 0x5e,
	0x85, 0xb0, 0xfa, 0xd6, 0xb7, 0xff, 0x9d, 0x12, 0xfb, 0x08, 0xcb, 0x91, 0xe6, 0xe8, 0x54, 0xe8,
	0xf3, 0x97, 0x37, 0xeb, 0x26, 0x5d, 0x53, 0x73, 0xba, 0xa6, 0xe6, 0x73, 0xbc, 0xa6, 0xaf, 0x0a,
	0xf7, 0x12, 0x92, 0x3c, 0xd6, 0x9b, 0x39, 0xf5, 0xdf, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xb9,
	0x55, 0xce, 0x11, 0xc2, 0x02, 0x00, 0x00,
}
