// Code generated by protoc-gen-go.
// source: github.com/luci/luci-go/common/proto/logdog/svcconfig/storage.proto
// DO NOT EDIT!

package svcconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/luci/luci-go/common/proto/google"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

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
	MaxLogAge *google_protobuf.Duration `protobuf:"bytes,2,opt,name=max_log_age,json=maxLogAge" json:"max_log_age,omitempty"`
}

func (m *Storage) Reset()                    { *m = Storage{} }
func (m *Storage) String() string            { return proto.CompactTextString(m) }
func (*Storage) ProtoMessage()               {}
func (*Storage) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type isStorage_Type interface {
	isStorage_Type()
}

type Storage_Bigtable struct {
	Bigtable *Storage_BigTable `protobuf:"bytes,1,opt,name=bigtable,oneof"`
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

func (m *Storage) GetMaxLogAge() *google_protobuf.Duration {
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
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
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
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// The name of the BigTable instance.
	Instance string `protobuf:"bytes,2,opt,name=instance" json:"instance,omitempty"`
	// The name of the BigTable instance's log table.
	LogTableName string `protobuf:"bytes,3,opt,name=log_table_name,json=logTableName" json:"log_table_name,omitempty"`
}

func (m *Storage_BigTable) Reset()                    { *m = Storage_BigTable{} }
func (m *Storage_BigTable) String() string            { return proto.CompactTextString(m) }
func (*Storage_BigTable) ProtoMessage()               {}
func (*Storage_BigTable) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

func init() {
	proto.RegisterType((*Storage)(nil), "svcconfig.Storage")
	proto.RegisterType((*Storage_BigTable)(nil), "svcconfig.Storage.BigTable")
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/common/proto/logdog/svcconfig/storage.proto", fileDescriptor3)
}

var fileDescriptor3 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x3c, 0x8e, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0xc7, 0x09, 0xa0, 0x36, 0x71, 0x11, 0x43, 0xa6, 0x10, 0x24, 0x84, 0x10, 0x03, 0x0b, 0xb6,
	0x04, 0x13, 0x23, 0x85, 0x81, 0x01, 0x31, 0x84, 0xee, 0x91, 0xed, 0x3a, 0x87, 0x51, 0xe2, 0xab,
	0x52, 0x07, 0xc1, 0x53, 0xf3, 0x0a, 0xbd, 0xda, 0x4d, 0x16, 0xcb, 0x77, 0xf7, 0xfb, 0x7f, 0xb0,
	0x17, 0xb0, 0xfe, 0x6b, 0x50, 0x5c, 0x63, 0x27, 0xda, 0x41, 0xdb, 0xf0, 0xdc, 0x03, 0x0a, 0x5a,
	0x74, 0xe8, 0xc4, 0xa6, 0x47, 0x8f, 0xa2, 0x45, 0x58, 0x23, 0x88, 0xed, 0x8f, 0xd6, 0xe8, 0x1a,
	0x4b, 0x3f, 0x8f, 0xbd, 0x04, 0xc3, 0xc3, 0x39, 0xcf, 0xa6, 0x43, 0x79, 0x05, 0x88, 0xd0, 0x9a,
	0xa8, 0x53, 0x43, 0x23, 0xd6, 0x43, 0x2f, 0xbd, 0x45, 0x17, 0xd1, 0x9b, 0xff, 0x84, 0xcd, 0x3f,
	0xa3, 0x38, 0x7f, 0x62, 0xa9, 0xb2, 0xe0, 0xa5, 0x6a, 0x4d, 0x91, 0x5c, 0x27, 0x77, 0x8b, 0x87,
	0x4b, 0x3e, 0x39, 0xf1, 0x03, 0xc5, 0x97, 0x16, 0x56, 0x7b, 0xe4, 0xed, 0xa8, 0x9a, 0x70, 0x92,
	0x2e, 0x3a, 0xf9, 0x5b, 0x53, 0xaf, 0x9a, 0x98, 0xe2, 0x38, 0xa8, 0x2f, 0x78, 0x0c, 0xe7, 0x63,
	0x38, 0x7f, 0x3d, 0x84, 0x57, 0x19, 0xd1, 0xef, 0x08, 0xcf, 0x60, 0xca, 0x86, 0xa5, 0xa3, 0x65,
	0x5e, 0xb0, 0x39, 0xb1, 0xdf, 0x46, 0xfb, 0x50, 0x20, 0xab, 0xc6, 0x31, 0x2f, 0x59, 0x6a, 0xdd,
	0xd6, 0x4b, 0xa7, 0xa3, 0x7b, 0x56, 0x4d, 0x73, 0x7e, 0xcb, 0xce, 0xf7, 0xc1, 0xa1, 0x49, 0xed,
	0x64, 0x67, 0x8a, 0x93, 0x40, 0x9c, 0xd1, 0x36, 0xf8, 0x7e, 0xd0, 0x6e, 0x39, 0x63, 0xa7, 0xab,
	0xbf, 0x8d, 0x51, 0xb3, 0xd0, 0xe6, 0x71, 0x17, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x61, 0x2c, 0xe1,
	0x6a, 0x01, 0x00, 0x00,
}
