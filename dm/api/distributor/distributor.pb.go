// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/dm/api/distributor/distributor.proto

package distributor

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	jobsim "github.com/TriggerMail/luci-go/dm/api/distributor/jobsim"
	v1 "github.com/TriggerMail/luci-go/dm/api/distributor/swarming/v1"
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

type Alias struct {
	OtherConfig          string   `protobuf:"bytes,1,opt,name=other_config,json=otherConfig,proto3" json:"other_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Alias) Reset()         { *m = Alias{} }
func (m *Alias) String() string { return proto.CompactTextString(m) }
func (*Alias) ProtoMessage()    {}
func (*Alias) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae589283bac413cb, []int{0}
}

func (m *Alias) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Alias.Unmarshal(m, b)
}
func (m *Alias) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Alias.Marshal(b, m, deterministic)
}
func (m *Alias) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Alias.Merge(m, src)
}
func (m *Alias) XXX_Size() int {
	return xxx_messageInfo_Alias.Size(m)
}
func (m *Alias) XXX_DiscardUnknown() {
	xxx_messageInfo_Alias.DiscardUnknown(m)
}

var xxx_messageInfo_Alias proto.InternalMessageInfo

func (m *Alias) GetOtherConfig() string {
	if m != nil {
		return m.OtherConfig
	}
	return ""
}

type Distributor struct {
	// TODO(iannucci): Maybe something like Any or extensions would be a better
	// fit here? The ultimate goal is that users will be able to use the proto
	// text format for luci-config. I suspect that Any or extensions would lose
	// the ability to validate such text-formatted protobufs, but maybe that's
	// not the case.
	//
	// Types that are valid to be assigned to DistributorType:
	//	*Distributor_Alias
	//	*Distributor_SwarmingV1
	//	*Distributor_Jobsim
	DistributorType      isDistributor_DistributorType `protobuf_oneof:"distributor_type"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *Distributor) Reset()         { *m = Distributor{} }
func (m *Distributor) String() string { return proto.CompactTextString(m) }
func (*Distributor) ProtoMessage()    {}
func (*Distributor) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae589283bac413cb, []int{1}
}

func (m *Distributor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Distributor.Unmarshal(m, b)
}
func (m *Distributor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Distributor.Marshal(b, m, deterministic)
}
func (m *Distributor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Distributor.Merge(m, src)
}
func (m *Distributor) XXX_Size() int {
	return xxx_messageInfo_Distributor.Size(m)
}
func (m *Distributor) XXX_DiscardUnknown() {
	xxx_messageInfo_Distributor.DiscardUnknown(m)
}

var xxx_messageInfo_Distributor proto.InternalMessageInfo

type isDistributor_DistributorType interface {
	isDistributor_DistributorType()
}

type Distributor_Alias struct {
	Alias *Alias `protobuf:"bytes,1,opt,name=alias,proto3,oneof"`
}

type Distributor_SwarmingV1 struct {
	SwarmingV1 *v1.Config `protobuf:"bytes,4,opt,name=swarming_v1,json=swarmingV1,proto3,oneof"`
}

type Distributor_Jobsim struct {
	Jobsim *jobsim.Config `protobuf:"bytes,2048,opt,name=jobsim,proto3,oneof"`
}

func (*Distributor_Alias) isDistributor_DistributorType() {}

func (*Distributor_SwarmingV1) isDistributor_DistributorType() {}

func (*Distributor_Jobsim) isDistributor_DistributorType() {}

func (m *Distributor) GetDistributorType() isDistributor_DistributorType {
	if m != nil {
		return m.DistributorType
	}
	return nil
}

func (m *Distributor) GetAlias() *Alias {
	if x, ok := m.GetDistributorType().(*Distributor_Alias); ok {
		return x.Alias
	}
	return nil
}

func (m *Distributor) GetSwarmingV1() *v1.Config {
	if x, ok := m.GetDistributorType().(*Distributor_SwarmingV1); ok {
		return x.SwarmingV1
	}
	return nil
}

func (m *Distributor) GetJobsim() *jobsim.Config {
	if x, ok := m.GetDistributorType().(*Distributor_Jobsim); ok {
		return x.Jobsim
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Distributor) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Distributor_OneofMarshaler, _Distributor_OneofUnmarshaler, _Distributor_OneofSizer, []interface{}{
		(*Distributor_Alias)(nil),
		(*Distributor_SwarmingV1)(nil),
		(*Distributor_Jobsim)(nil),
	}
}

func _Distributor_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Distributor)
	// distributor_type
	switch x := m.DistributorType.(type) {
	case *Distributor_Alias:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Alias); err != nil {
			return err
		}
	case *Distributor_SwarmingV1:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SwarmingV1); err != nil {
			return err
		}
	case *Distributor_Jobsim:
		b.EncodeVarint(2048<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Jobsim); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Distributor.DistributorType has unexpected type %T", x)
	}
	return nil
}

func _Distributor_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Distributor)
	switch tag {
	case 1: // distributor_type.alias
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Alias)
		err := b.DecodeMessage(msg)
		m.DistributorType = &Distributor_Alias{msg}
		return true, err
	case 4: // distributor_type.swarming_v1
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(v1.Config)
		err := b.DecodeMessage(msg)
		m.DistributorType = &Distributor_SwarmingV1{msg}
		return true, err
	case 2048: // distributor_type.jobsim
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(jobsim.Config)
		err := b.DecodeMessage(msg)
		m.DistributorType = &Distributor_Jobsim{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Distributor_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Distributor)
	// distributor_type
	switch x := m.DistributorType.(type) {
	case *Distributor_Alias:
		s := proto.Size(x.Alias)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Distributor_SwarmingV1:
		s := proto.Size(x.SwarmingV1)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Distributor_Jobsim:
		s := proto.Size(x.Jobsim)
		n += 3 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Config struct {
	DistributorConfigs   map[string]*Distributor `protobuf:"bytes,1,rep,name=distributor_configs,json=distributorConfigs,proto3" json:"distributor_configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_ae589283bac413cb, []int{2}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetDistributorConfigs() map[string]*Distributor {
	if m != nil {
		return m.DistributorConfigs
	}
	return nil
}

func init() {
	proto.RegisterType((*Alias)(nil), "distributor.Alias")
	proto.RegisterType((*Distributor)(nil), "distributor.Distributor")
	proto.RegisterType((*Config)(nil), "distributor.Config")
	proto.RegisterMapType((map[string]*Distributor)(nil), "distributor.Config.DistributorConfigsEntry")
}

func init() {
	proto.RegisterFile("github.com/TriggerMail/luci-go/dm/api/distributor/distributor.proto", fileDescriptor_ae589283bac413cb)
}

var fileDescriptor_ae589283bac413cb = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xd7, 0xb5, 0x1b, 0xf3, 0x8d, 0x48, 0x89, 0x07, 0xcb, 0x4e, 0xba, 0xd3, 0x9c, 0x90,
	0xb0, 0x89, 0xe0, 0x9f, 0x8b, 0x7f, 0x61, 0xec, 0xd8, 0x83, 0x27, 0xa1, 0x74, 0x5b, 0xed, 0xa2,
	0xeb, 0x52, 0xd2, 0xb4, 0xd2, 0x9b, 0xdf, 0xcb, 0x0f, 0xe0, 0xd7, 0x92, 0x26, 0x9d, 0xcb, 0x10,
	0x61, 0xa7, 0x86, 0x27, 0xcf, 0xf3, 0xbc, 0xfd, 0xbd, 0x81, 0xab, 0x98, 0x93, 0xd9, 0x42, 0xf0,
	0x84, 0xe5, 0x09, 0xe1, 0x22, 0xa6, 0xcb, 0x7c, 0xc6, 0xe8, 0x3c, 0xa1, 0x61, 0xca, 0xe8, 0x9c,
	0x65, 0x52, 0xb0, 0x69, 0x2e, 0xb9, 0x30, 0xcf, 0x24, 0x15, 0x5c, 0x72, 0x8c, 0x0c, 0xa9, 0x7b,
	0xb3, 0x6b, 0xcf, 0x1b, 0x9f, 0x66, 0x2c, 0xa9, 0x3f, 0xba, 0xa9, 0x7b, 0xbb, 0x6b, 0x38, 0xfb,
	0x08, 0x45, 0xc2, 0x56, 0x31, 0x2d, 0x86, 0x74, 0xc6, 0x57, 0xaf, 0x2c, 0xd6, 0x0d, 0xbd, 0x01,
	0xb4, 0xee, 0x96, 0x2c, 0xcc, 0xf0, 0x09, 0xec, 0x73, 0xb9, 0x88, 0x44, 0xa0, 0xaf, 0x3d, 0xeb,
	0xd8, 0xea, 0xef, 0xf9, 0x48, 0x69, 0x0f, 0x4a, 0xea, 0x7d, 0x59, 0x80, 0x1e, 0x37, 0xa5, 0x78,
	0x00, 0xad, 0xb0, 0xca, 0x2a, 0x2f, 0x1a, 0x61, 0x62, 0xa2, 0xaa, 0xd6, 0x71, 0xc3, 0xd7, 0x16,
	0x7c, 0x01, 0x68, 0xfd, 0x0f, 0x41, 0x31, 0xf4, 0x9c, 0x3a, 0xb1, 0xd6, 0x9e, 0x87, 0x44, 0x0f,
	0x19, 0x37, 0x7c, 0xd8, 0x88, 0xf8, 0x14, 0xda, 0x1a, 0xd8, 0xfb, 0x74, 0x55, 0xe4, 0x80, 0xd4,
	0x0b, 0xf8, 0xb5, 0xd7, 0x86, 0x7b, 0x0c, 0xae, 0x31, 0x3f, 0x90, 0x65, 0x1a, 0x4d, 0x9c, 0x4e,
	0xd3, 0xb5, 0x27, 0x4e, 0xc7, 0x76, 0x9d, 0xde, 0xb7, 0x05, 0x6d, 0x1d, 0xc2, 0x2f, 0x70, 0x68,
	0x5a, 0x35, 0x71, 0x85, 0x61, 0xf7, 0xd1, 0xe8, 0x6c, 0x0b, 0x43, 0x27, 0x88, 0x81, 0xad, 0x95,
	0xec, 0x69, 0x25, 0x45, 0xe9, 0xe3, 0xf9, 0x9f, 0x8b, 0x6e, 0x00, 0x47, 0xff, 0xd8, 0xb1, 0x0b,
	0xf6, 0x7b, 0x54, 0xd6, 0xbb, 0xad, 0x8e, 0x98, 0x40, 0xab, 0x08, 0x97, 0x79, 0xe4, 0x35, 0x15,
	0x9e, 0xb7, 0x35, 0xdc, 0xa8, 0xf1, 0xb5, 0xed, 0xba, 0x79, 0x69, 0x4d, 0xdb, 0xea, 0xe9, 0xce,
	0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x1a, 0x96, 0xe8, 0x83, 0x02, 0x00, 0x00,
}
