// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/scheduler/appengine/internal/triggers.proto

package internal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import scheduler "go.chromium.org/luci/scheduler/api/scheduler/v1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Trigger can be emitted by triggering tasks (such as Gitiles tasks) and
// consumed by triggered tasks (such as Buildbucket tasks).
type Trigger struct {
	// Unique in time identifier of the trigger.
	//
	// It is used to deduplicate and hence provide idempotency for adding
	// a trigger. Must be provided by whoever emits the trigger.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// ID of a job that emitted this trigger.
	//
	// Set by the engine, can't be overridden.
	JobId string `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	// ID of an invocation that emitted this trigger.
	//
	// Set by the engine, can't be overridden.
	InvocationId int64 `protobuf:"varint,3,opt,name=invocation_id,json=invocationId" json:"invocation_id,omitempty"`
	// Timestamp when the trigger was created.
	//
	// Set by the engine, can't be overridden.
	Created *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=created" json:"created,omitempty"`
	// User friendly name for this trigger that shows up in UI.
	//
	// Can be provided by whoever emits the trigger. Doesn't have to be unique.
	Title string `protobuf:"bytes,5,opt,name=title" json:"title,omitempty"`
	// Optional HTTP link to display in UI.
	//
	// Can be provided by whoever emits the trigger. Doesn't have to be unique.
	Url string `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
	// Actual trigger data that depends on type of the trigger.
	//
	// Types that are valid to be assigned to Payload:
	//	*Trigger_Noop
	//	*Trigger_Gitiles
	Payload isTrigger_Payload `protobuf_oneof:"payload"`
}

func (m *Trigger) Reset()                    { *m = Trigger{} }
func (m *Trigger) String() string            { return proto.CompactTextString(m) }
func (*Trigger) ProtoMessage()               {}
func (*Trigger) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type isTrigger_Payload interface {
	isTrigger_Payload()
}

type Trigger_Noop struct {
	Noop *scheduler.NoopTrigger `protobuf:"bytes,50,opt,name=noop,oneof"`
}
type Trigger_Gitiles struct {
	Gitiles *scheduler.GitilesTrigger `protobuf:"bytes,51,opt,name=gitiles,oneof"`
}

func (*Trigger_Noop) isTrigger_Payload()    {}
func (*Trigger_Gitiles) isTrigger_Payload() {}

func (m *Trigger) GetPayload() isTrigger_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Trigger) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Trigger) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *Trigger) GetInvocationId() int64 {
	if m != nil {
		return m.InvocationId
	}
	return 0
}

func (m *Trigger) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Trigger) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Trigger) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Trigger) GetNoop() *scheduler.NoopTrigger {
	if x, ok := m.GetPayload().(*Trigger_Noop); ok {
		return x.Noop
	}
	return nil
}

func (m *Trigger) GetGitiles() *scheduler.GitilesTrigger {
	if x, ok := m.GetPayload().(*Trigger_Gitiles); ok {
		return x.Gitiles
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Trigger) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Trigger_OneofMarshaler, _Trigger_OneofUnmarshaler, _Trigger_OneofSizer, []interface{}{
		(*Trigger_Noop)(nil),
		(*Trigger_Gitiles)(nil),
	}
}

func _Trigger_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Trigger)
	// payload
	switch x := m.Payload.(type) {
	case *Trigger_Noop:
		b.EncodeVarint(50<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Noop); err != nil {
			return err
		}
	case *Trigger_Gitiles:
		b.EncodeVarint(51<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Gitiles); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Trigger.Payload has unexpected type %T", x)
	}
	return nil
}

func _Trigger_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Trigger)
	switch tag {
	case 50: // payload.noop
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(scheduler.NoopTrigger)
		err := b.DecodeMessage(msg)
		m.Payload = &Trigger_Noop{msg}
		return true, err
	case 51: // payload.gitiles
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(scheduler.GitilesTrigger)
		err := b.DecodeMessage(msg)
		m.Payload = &Trigger_Gitiles{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Trigger_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Trigger)
	// payload
	switch x := m.Payload.(type) {
	case *Trigger_Noop:
		s := proto.Size(x.Noop)
		n += proto.SizeVarint(50<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Trigger_Gitiles:
		s := proto.Size(x.Gitiles)
		n += proto.SizeVarint(51<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// TriggerList is what we store in datastore entities.
type TriggerList struct {
	Triggers []*Trigger `protobuf:"bytes,1,rep,name=triggers" json:"triggers,omitempty"`
}

func (m *TriggerList) Reset()                    { *m = TriggerList{} }
func (m *TriggerList) String() string            { return proto.CompactTextString(m) }
func (*TriggerList) ProtoMessage()               {}
func (*TriggerList) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *TriggerList) GetTriggers() []*Trigger {
	if m != nil {
		return m.Triggers
	}
	return nil
}

func init() {
	proto.RegisterType((*Trigger)(nil), "internal.triggers.Trigger")
	proto.RegisterType((*TriggerList)(nil), "internal.triggers.TriggerList")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/scheduler/appengine/internal/triggers.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4f, 0xe3, 0x30,
	0x10, 0x85, 0x37, 0x49, 0xdb, 0xb4, 0xee, 0xee, 0x6a, 0xd7, 0xda, 0x5d, 0x79, 0x7b, 0x21, 0x2a,
	0x97, 0x1e, 0x90, 0x23, 0x5a, 0xe0, 0x88, 0x50, 0x25, 0x04, 0x95, 0x10, 0x87, 0xa8, 0x27, 0x2e,
	0x28, 0x89, 0x8d, 0xeb, 0xca, 0xc9, 0x44, 0x8e, 0x53, 0x89, 0xdf, 0xc3, 0x1f, 0x45, 0x75, 0xea,
	0xb6, 0x82, 0x03, 0xb7, 0xcc, 0x7b, 0xdf, 0xcb, 0xe8, 0x8d, 0xd1, 0x8d, 0x00, 0x9a, 0xaf, 0x34,
	0x14, 0xb2, 0x29, 0x28, 0x68, 0x11, 0xab, 0x26, 0x97, 0x71, 0x9d, 0xaf, 0x38, 0x6b, 0x14, 0xd7,
	0x71, 0x5a, 0x55, 0xbc, 0x14, 0xb2, 0xe4, 0xb1, 0x2c, 0x0d, 0xd7, 0x65, 0xaa, 0x62, 0xa3, 0xa5,
	0x10, 0x5c, 0xd7, 0xb4, 0xd2, 0x60, 0x00, 0xff, 0x76, 0x06, 0x75, 0xc6, 0xe8, 0x44, 0x00, 0x08,
	0xc5, 0x63, 0x0b, 0x64, 0xcd, 0x4b, 0x6c, 0x64, 0xc1, 0x6b, 0x93, 0x16, 0x55, 0x9b, 0x19, 0x5d,
	0x7f, 0xb9, 0xf5, 0x78, 0xda, 0x9c, 0x7f, 0xd8, 0x39, 0x7e, 0xf3, 0x51, 0xb8, 0x6c, 0x25, 0xfc,
	0x13, 0xf9, 0x92, 0x11, 0x2f, 0xf2, 0x26, 0x83, 0xc4, 0x97, 0x0c, 0xff, 0x45, 0xbd, 0x35, 0x64,
	0xcf, 0x92, 0x11, 0xdf, 0x6a, 0xdd, 0x35, 0x64, 0x0b, 0x86, 0x4f, 0xd1, 0x0f, 0x59, 0x6e, 0x20,
	0x4f, 0x8d, 0x84, 0x72, 0xeb, 0x06, 0x91, 0x37, 0x09, 0x92, 0xef, 0x07, 0x71, 0xc1, 0xf0, 0x05,
	0x0a, 0x73, 0xcd, 0x53, 0xc3, 0x19, 0xe9, 0x44, 0xde, 0x64, 0x38, 0x1d, 0xd1, 0xb6, 0x0a, 0x75,
	0x55, 0xe8, 0xd2, 0x55, 0x49, 0x1c, 0x8a, 0xff, 0xa0, 0xae, 0x91, 0x46, 0x71, 0xd2, 0x6d, 0x17,
	0xda, 0x01, 0xff, 0x42, 0x41, 0xa3, 0x15, 0xe9, 0x59, 0x6d, 0xfb, 0x89, 0xcf, 0x50, 0xa7, 0x04,
	0xa8, 0xc8, 0xd4, 0xfe, 0xfa, 0x1f, 0xdd, 0x37, 0xa4, 0x8f, 0x00, 0xd5, 0xae, 0xcf, 0xfd, 0xb7,
	0xc4, 0x52, 0xf8, 0x12, 0x85, 0x42, 0x1a, 0xa9, 0x78, 0x4d, 0x66, 0x36, 0xf0, 0xff, 0x28, 0x70,
	0xd7, 0x3a, 0x87, 0x8c, 0x63, 0xe7, 0x03, 0x14, 0x56, 0xe9, 0xab, 0x82, 0x94, 0x8d, 0x6f, 0xd1,
	0x70, 0x07, 0x3c, 0xc8, 0xda, 0xe0, 0x2b, 0xd4, 0x77, 0x67, 0x24, 0x5e, 0x14, 0xd8, 0x76, 0x9f,
	0xde, 0x8e, 0xee, 0x12, 0xc9, 0x9e, 0x9d, 0xa3, 0xa7, 0xbe, 0xc3, 0xb2, 0x9e, 0xbd, 0xc3, 0xec,
	0x3d, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xbe, 0xaa, 0x3e, 0x37, 0x02, 0x00, 0x00,
}
