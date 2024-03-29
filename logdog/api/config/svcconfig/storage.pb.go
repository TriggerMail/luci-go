// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/logdog/api/config/svcconfig/storage.proto

package svcconfig

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// Storage is the in-transit storage configuration.
type Storage struct {
	// Type is the transport configuration that is being used.
	//
	// Types that are valid to be assigned to Type:
	//	*Storage_Bigtable
	Type isStorage_Type `protobuf_oneof:"Type"`
	// The maximum lifetime of a log's intermediate storage entries. The Storage
	// instance is free to begin deleting log entries if they are older than this.
	//
	// It is recommended that this be set to 4*[terminal archival threshold],
	// where the terminal archival threshold is the amount of time that the
	// Coordinator will wait on a log stream that has not been terminated before
	// constructing an archive.
	//
	// Waiting at least the archival threshold ensures that the log entries are
	// available for streams that expire. Waiting longer than the threshold is
	// good because a user may be streaming logs from intermediate storage as they
	// become archived. Waiting a *lot* longer is useful to prevent data loss in
	// the event of issues with the archival process.
	MaxLogAge            *duration.Duration `protobuf:"bytes,2,opt,name=max_log_age,json=maxLogAge,proto3" json:"max_log_age,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Storage) Reset()         { *m = Storage{} }
func (m *Storage) String() string { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()    {}
func (*Storage) Descriptor() ([]byte, []int) {
	return fileDescriptor_955b461662b6fa9d, []int{0}
}

func (m *Storage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storage.Unmarshal(m, b)
}
func (m *Storage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storage.Marshal(b, m, deterministic)
}
func (m *Storage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storage.Merge(m, src)
}
func (m *Storage) XXX_Size() int {
	return xxx_messageInfo_Storage.Size(m)
}
func (m *Storage) XXX_DiscardUnknown() {
	xxx_messageInfo_Storage.DiscardUnknown(m)
}

var xxx_messageInfo_Storage proto.InternalMessageInfo

type isStorage_Type interface {
	isStorage_Type()
}

type Storage_Bigtable struct {
	Bigtable *Storage_BigTable `protobuf:"bytes,1,opt,name=bigtable,proto3,oneof"`
}

func (*Storage_Bigtable) isStorage_Type() {}

func (m *Storage) GetType() isStorage_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Storage) GetBigtable() *Storage_BigTable {
	if x, ok := m.GetType().(*Storage_Bigtable); ok {
		return x.Bigtable
	}
	return nil
}

func (m *Storage) GetMaxLogAge() *duration.Duration {
	if m != nil {
		return m.MaxLogAge
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Storage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Storage_OneofMarshaler, _Storage_OneofUnmarshaler, _Storage_OneofSizer, []interface{}{
		(*Storage_Bigtable)(nil),
	}
}

func _Storage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Storage)
	// Type
	switch x := m.Type.(type) {
	case *Storage_Bigtable:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Bigtable); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Storage.Type has unexpected type %T", x)
	}
	return nil
}

func _Storage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Storage)
	switch tag {
	case 1: // Type.bigtable
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Storage_BigTable)
		err := b.DecodeMessage(msg)
		m.Type = &Storage_Bigtable{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Storage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Storage)
	// Type
	switch x := m.Type.(type) {
	case *Storage_Bigtable:
		s := proto.Size(x.Bigtable)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// BigTable is the set of BigTable configuration parameters.
type Storage_BigTable struct {
	// The name of the BigTable instance project.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// The name of the BigTable instance.
	Instance string `protobuf:"bytes,2,opt,name=instance,proto3" json:"instance,omitempty"`
	// The name of the BigTable instance's log table.
	LogTableName         string   `protobuf:"bytes,3,opt,name=log_table_name,json=logTableName,proto3" json:"log_table_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Storage_BigTable) Reset()         { *m = Storage_BigTable{} }
func (m *Storage_BigTable) String() string { return proto.CompactTextString(m) }
func (*Storage_BigTable) ProtoMessage()    {}
func (*Storage_BigTable) Descriptor() ([]byte, []int) {
	return fileDescriptor_955b461662b6fa9d, []int{0, 0}
}

func (m *Storage_BigTable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Storage_BigTable.Unmarshal(m, b)
}
func (m *Storage_BigTable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Storage_BigTable.Marshal(b, m, deterministic)
}
func (m *Storage_BigTable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Storage_BigTable.Merge(m, src)
}
func (m *Storage_BigTable) XXX_Size() int {
	return xxx_messageInfo_Storage_BigTable.Size(m)
}
func (m *Storage_BigTable) XXX_DiscardUnknown() {
	xxx_messageInfo_Storage_BigTable.DiscardUnknown(m)
}

var xxx_messageInfo_Storage_BigTable proto.InternalMessageInfo

func (m *Storage_BigTable) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *Storage_BigTable) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *Storage_BigTable) GetLogTableName() string {
	if m != nil {
		return m.LogTableName
	}
	return ""
}

func init() {
	proto.RegisterType((*Storage)(nil), "svcconfig.Storage")
	proto.RegisterType((*Storage_BigTable)(nil), "svcconfig.Storage.BigTable")
}

func init() {
	proto.RegisterFile("github.com/TriggerMail/luci-go/logdog/api/config/svcconfig/storage.proto", fileDescriptor_955b461662b6fa9d)
}

var fileDescriptor_955b461662b6fa9d = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x29, 0xa0, 0x36, 0x71, 0x11, 0x83, 0xa7, 0x10, 0x24, 0x84, 0x10, 0x03, 0x93, 0x2d,
	0xc1, 0xd4, 0x05, 0x89, 0x8a, 0x81, 0x01, 0x31, 0x84, 0xee, 0xd1, 0xc5, 0x75, 0x0e, 0x23, 0x3b,
	0x17, 0xb9, 0x0e, 0x2a, 0x4f, 0xcd, 0x2b, 0xa0, 0xda, 0x4d, 0xb6, 0x3b, 0xdd, 0x77, 0xff, 0xf7,
	0xb3, 0x67, 0x24, 0xa1, 0xbe, 0x3c, 0x39, 0x33, 0x38, 0x41, 0x1e, 0xa5, 0x1d, 0x94, 0x91, 0x96,
	0x70, 0x4b, 0x28, 0xa1, 0x37, 0x52, 0x51, 0xd7, 0x1a, 0x94, 0xbb, 0x1f, 0x35, 0x4e, 0x81, 0x3c,
	0xa0, 0x16, 0xbd, 0xa7, 0x40, 0x3c, 0x9f, 0x0e, 0xe5, 0x0d, 0x12, 0xa1, 0xd5, 0x32, 0x1e, 0x9a,
	0xa1, 0x95, 0xdb, 0xc1, 0x43, 0x30, 0xd4, 0x25, 0xf4, 0xee, 0x6f, 0xc6, 0x16, 0x9f, 0xe9, 0x99,
	0xaf, 0x58, 0xd6, 0x18, 0x0c, 0xd0, 0x58, 0x5d, 0xcc, 0x6e, 0x67, 0x0f, 0xcb, 0xc7, 0x6b, 0x31,
	0x25, 0x89, 0x23, 0x25, 0xd6, 0x06, 0x37, 0x07, 0xe4, 0xed, 0xa4, 0x9a, 0x70, 0xbe, 0x62, 0x4b,
	0x07, 0xfb, 0xda, 0x12, 0xd6, 0x80, 0xba, 0x38, 0x8d, 0xdf, 0x57, 0x22, 0xc9, 0xc5, 0x28, 0x17,
	0xaf, 0x47, 0x79, 0x95, 0x3b, 0xd8, 0xbf, 0x13, 0xbe, 0xa0, 0x2e, 0x5b, 0x96, 0x8d, 0x91, 0xbc,
	0x60, 0x8b, 0xde, 0xd3, 0xb7, 0x56, 0x21, 0x16, 0xc8, 0xab, 0x71, 0xe5, 0x25, 0xcb, 0x4c, 0xb7,
	0x0b, 0xd0, 0xa9, 0x94, 0x9e, 0x57, 0xd3, 0xce, 0xef, 0xd9, 0xe5, 0x41, 0x1c, 0x9b, 0xd4, 0x1d,
	0x38, 0x5d, 0x9c, 0x45, 0xe2, 0xc2, 0x52, 0xca, 0xfd, 0x00, 0xa7, 0xd7, 0x73, 0x76, 0xbe, 0xf9,
	0xed, 0x75, 0x33, 0x8f, 0x6d, 0x9e, 0xfe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x2c, 0xc6, 0xfe,
	0x65, 0x01, 0x00, 0x00,
}
