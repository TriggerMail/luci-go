// Code generated by protoc-gen-go.
// source: github.com/luci/luci-go/dm/api/distributor/jobsim/task.proto
// DO NOT EDIT!

package jobsim

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/luci/luci-go/common/proto/google"
import google_protobuf1 "github.com/luci/luci-go/common/proto/google"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Phrase is a task to do. It consists of zero or more stages, followed by
// an optional ReturnStage.
type Phrase struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Seed should be selected via a fair dice roll (using a d1.84e19).
	Seed        int64        `protobuf:"varint,2,opt,name=seed" json:"seed,omitempty"`
	Stages      []*Stage     `protobuf:"bytes,3,rep,name=stages" json:"stages,omitempty"`
	ReturnStage *ReturnStage `protobuf:"bytes,4,opt,name=return_stage,json=returnStage" json:"return_stage,omitempty"`
}

func (m *Phrase) Reset()                    { *m = Phrase{} }
func (m *Phrase) String() string            { return proto.CompactTextString(m) }
func (*Phrase) ProtoMessage()               {}
func (*Phrase) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Phrase) GetStages() []*Stage {
	if m != nil {
		return m.Stages
	}
	return nil
}

func (m *Phrase) GetReturnStage() *ReturnStage {
	if m != nil {
		return m.ReturnStage
	}
	return nil
}

// ReturnStage indicates that the Phrase should return the numerical value
// 'retval' as the result of the current Phrase. If expiration is provided,
// it will set that as the expiration timestamp for the provided retval.
type ReturnStage struct {
	Retval     int64                      `protobuf:"varint,1,opt,name=retval" json:"retval,omitempty"`
	Expiration *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=expiration" json:"expiration,omitempty"`
}

func (m *ReturnStage) Reset()                    { *m = ReturnStage{} }
func (m *ReturnStage) String() string            { return proto.CompactTextString(m) }
func (*ReturnStage) ProtoMessage()               {}
func (*ReturnStage) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *ReturnStage) GetExpiration() *google_protobuf.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

// Stage is the union of the following stage types:
//   * FailureStage
//   * StallStage
//   * DepsStage
type Stage struct {
	// Types that are valid to be assigned to StageType:
	//	*Stage_Failure
	//	*Stage_Stall
	//	*Stage_Deps
	StageType isStage_StageType `protobuf_oneof:"stage_type"`
}

func (m *Stage) Reset()                    { *m = Stage{} }
func (m *Stage) String() string            { return proto.CompactTextString(m) }
func (*Stage) ProtoMessage()               {}
func (*Stage) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type isStage_StageType interface {
	isStage_StageType()
}

type Stage_Failure struct {
	Failure *FailureStage `protobuf:"bytes,1,opt,name=failure,oneof"`
}
type Stage_Stall struct {
	Stall *StallStage `protobuf:"bytes,2,opt,name=stall,oneof"`
}
type Stage_Deps struct {
	Deps *DepsStage `protobuf:"bytes,3,opt,name=deps,oneof"`
}

func (*Stage_Failure) isStage_StageType() {}
func (*Stage_Stall) isStage_StageType()   {}
func (*Stage_Deps) isStage_StageType()    {}

func (m *Stage) GetStageType() isStage_StageType {
	if m != nil {
		return m.StageType
	}
	return nil
}

func (m *Stage) GetFailure() *FailureStage {
	if x, ok := m.GetStageType().(*Stage_Failure); ok {
		return x.Failure
	}
	return nil
}

func (m *Stage) GetStall() *StallStage {
	if x, ok := m.GetStageType().(*Stage_Stall); ok {
		return x.Stall
	}
	return nil
}

func (m *Stage) GetDeps() *DepsStage {
	if x, ok := m.GetStageType().(*Stage_Deps); ok {
		return x.Deps
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Stage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Stage_OneofMarshaler, _Stage_OneofUnmarshaler, _Stage_OneofSizer, []interface{}{
		(*Stage_Failure)(nil),
		(*Stage_Stall)(nil),
		(*Stage_Deps)(nil),
	}
}

func _Stage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Stage)
	// stage_type
	switch x := m.StageType.(type) {
	case *Stage_Failure:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Failure); err != nil {
			return err
		}
	case *Stage_Stall:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Stall); err != nil {
			return err
		}
	case *Stage_Deps:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Deps); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Stage.StageType has unexpected type %T", x)
	}
	return nil
}

func _Stage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Stage)
	switch tag {
	case 1: // stage_type.failure
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FailureStage)
		err := b.DecodeMessage(msg)
		m.StageType = &Stage_Failure{msg}
		return true, err
	case 2: // stage_type.stall
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StallStage)
		err := b.DecodeMessage(msg)
		m.StageType = &Stage_Stall{msg}
		return true, err
	case 3: // stage_type.deps
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DepsStage)
		err := b.DecodeMessage(msg)
		m.StageType = &Stage_Deps{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Stage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Stage)
	// stage_type
	switch x := m.StageType.(type) {
	case *Stage_Failure:
		s := proto.Size(x.Failure)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Stage_Stall:
		s := proto.Size(x.Stall)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Stage_Deps:
		s := proto.Size(x.Deps)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// FailureStage is the /chance/ to fail with a certain liklihood. The chance
// is calculated using the current Phrase's 'seed' with the "math/rand" package,
// The seed is either 0 (unspecified), or the value of the 'seed' property for
// the currently running phrase.
//
// 0 is a 0-percent chance of failure.
// 1 is a 100-percent chance of failure.
type FailureStage struct {
	Chance float32 `protobuf:"fixed32,1,opt,name=chance" json:"chance,omitempty"`
}

func (m *FailureStage) Reset()                    { *m = FailureStage{} }
func (m *FailureStage) String() string            { return proto.CompactTextString(m) }
func (*FailureStage) ProtoMessage()               {}
func (*FailureStage) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

// StallStage delays the phrase for the provided Duration. This could be used
// to simulate long-running tasks (like builds).
type StallStage struct {
	Delay *google_protobuf1.Duration `protobuf:"bytes,1,opt,name=delay" json:"delay,omitempty"`
}

func (m *StallStage) Reset()                    { *m = StallStage{} }
func (m *StallStage) String() string            { return proto.CompactTextString(m) }
func (*StallStage) ProtoMessage()               {}
func (*StallStage) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *StallStage) GetDelay() *google_protobuf1.Duration {
	if m != nil {
		return m.Delay
	}
	return nil
}

// DepsStage represents the opportunity to depend on 1 or more dependencies
// simultaneously.
type DepsStage struct {
	Deps []*Dependency `protobuf:"bytes,1,rep,name=deps" json:"deps,omitempty"`
}

func (m *DepsStage) Reset()                    { *m = DepsStage{} }
func (m *DepsStage) String() string            { return proto.CompactTextString(m) }
func (*DepsStage) ProtoMessage()               {}
func (*DepsStage) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *DepsStage) GetDeps() []*Dependency {
	if m != nil {
		return m.Deps
	}
	return nil
}

// Dependency represents a nested Phrase that this Phrase depends on.
type Dependency struct {
	// shards append [1], [2], [3], etc. to the "name"s of the dependencies, making
	// them unique quests.
	Shards uint64 `protobuf:"varint,1,opt,name=shards" json:"shards,omitempty"`
	// Types that are valid to be assigned to AttemptStrategy:
	//	*Dependency_Attempts
	//	*Dependency_Retries
	AttemptStrategy isDependency_AttemptStrategy `protobuf_oneof:"attempt_strategy"`
	// MixSeed will blend the current seed with the seed in the phrase seed,
	// when depending on it.
	//
	//   mix_seed phrase.seed==0 -> dep uses "random" seed
	//   mix_seed phrase.seed!=0 -> dep uses blend(current seed, phrase.seed)
	//  !mix_seed phrase.seed==0 -> dep uses current seed
	//  !mix_seed phrase.seed!=0 -> dep uses phrase.seed
	MixSeed bool    `protobuf:"varint,4,opt,name=mix_seed,json=mixSeed" json:"mix_seed,omitempty"`
	Phrase  *Phrase `protobuf:"bytes,5,opt,name=phrase" json:"phrase,omitempty"`
}

func (m *Dependency) Reset()                    { *m = Dependency{} }
func (m *Dependency) String() string            { return proto.CompactTextString(m) }
func (*Dependency) ProtoMessage()               {}
func (*Dependency) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

type isDependency_AttemptStrategy interface {
	isDependency_AttemptStrategy()
}

type Dependency_Attempts struct {
	Attempts *SparseRange `protobuf:"bytes,2,opt,name=attempts,oneof"`
}
type Dependency_Retries struct {
	Retries uint32 `protobuf:"varint,3,opt,name=retries,oneof"`
}

func (*Dependency_Attempts) isDependency_AttemptStrategy() {}
func (*Dependency_Retries) isDependency_AttemptStrategy()  {}

func (m *Dependency) GetAttemptStrategy() isDependency_AttemptStrategy {
	if m != nil {
		return m.AttemptStrategy
	}
	return nil
}

func (m *Dependency) GetAttempts() *SparseRange {
	if x, ok := m.GetAttemptStrategy().(*Dependency_Attempts); ok {
		return x.Attempts
	}
	return nil
}

func (m *Dependency) GetRetries() uint32 {
	if x, ok := m.GetAttemptStrategy().(*Dependency_Retries); ok {
		return x.Retries
	}
	return 0
}

func (m *Dependency) GetPhrase() *Phrase {
	if m != nil {
		return m.Phrase
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Dependency) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Dependency_OneofMarshaler, _Dependency_OneofUnmarshaler, _Dependency_OneofSizer, []interface{}{
		(*Dependency_Attempts)(nil),
		(*Dependency_Retries)(nil),
	}
}

func _Dependency_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Dependency)
	// attempt_strategy
	switch x := m.AttemptStrategy.(type) {
	case *Dependency_Attempts:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Attempts); err != nil {
			return err
		}
	case *Dependency_Retries:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Retries))
	case nil:
	default:
		return fmt.Errorf("Dependency.AttemptStrategy has unexpected type %T", x)
	}
	return nil
}

func _Dependency_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Dependency)
	switch tag {
	case 2: // attempt_strategy.attempts
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SparseRange)
		err := b.DecodeMessage(msg)
		m.AttemptStrategy = &Dependency_Attempts{msg}
		return true, err
	case 3: // attempt_strategy.retries
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.AttemptStrategy = &Dependency_Retries{uint32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Dependency_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Dependency)
	// attempt_strategy
	switch x := m.AttemptStrategy.(type) {
	case *Dependency_Attempts:
		s := proto.Size(x.Attempts)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Dependency_Retries:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Retries))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// SparseRange allows the expression of mixed partial ranges like [1,3-10,19,21].
type SparseRange struct {
	Items []*RangeItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *SparseRange) Reset()                    { *m = SparseRange{} }
func (m *SparseRange) String() string            { return proto.CompactTextString(m) }
func (*SparseRange) ProtoMessage()               {}
func (*SparseRange) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *SparseRange) GetItems() []*RangeItem {
	if m != nil {
		return m.Items
	}
	return nil
}

// RangeItem is either a single number or a Range.
type RangeItem struct {
	// Types that are valid to be assigned to RangeItem:
	//	*RangeItem_Single
	//	*RangeItem_Range
	RangeItem isRangeItem_RangeItem `protobuf_oneof:"range_item"`
}

func (m *RangeItem) Reset()                    { *m = RangeItem{} }
func (m *RangeItem) String() string            { return proto.CompactTextString(m) }
func (*RangeItem) ProtoMessage()               {}
func (*RangeItem) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

type isRangeItem_RangeItem interface {
	isRangeItem_RangeItem()
}

type RangeItem_Single struct {
	Single uint32 `protobuf:"varint,1,opt,name=single,oneof"`
}
type RangeItem_Range struct {
	Range *Range `protobuf:"bytes,2,opt,name=range,oneof"`
}

func (*RangeItem_Single) isRangeItem_RangeItem() {}
func (*RangeItem_Range) isRangeItem_RangeItem()  {}

func (m *RangeItem) GetRangeItem() isRangeItem_RangeItem {
	if m != nil {
		return m.RangeItem
	}
	return nil
}

func (m *RangeItem) GetSingle() uint32 {
	if x, ok := m.GetRangeItem().(*RangeItem_Single); ok {
		return x.Single
	}
	return 0
}

func (m *RangeItem) GetRange() *Range {
	if x, ok := m.GetRangeItem().(*RangeItem_Range); ok {
		return x.Range
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RangeItem) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RangeItem_OneofMarshaler, _RangeItem_OneofUnmarshaler, _RangeItem_OneofSizer, []interface{}{
		(*RangeItem_Single)(nil),
		(*RangeItem_Range)(nil),
	}
}

func _RangeItem_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RangeItem)
	// range_item
	switch x := m.RangeItem.(type) {
	case *RangeItem_Single:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Single))
	case *RangeItem_Range:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Range); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RangeItem.RangeItem has unexpected type %T", x)
	}
	return nil
}

func _RangeItem_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RangeItem)
	switch tag {
	case 1: // range_item.single
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.RangeItem = &RangeItem_Single{uint32(x)}
		return true, err
	case 2: // range_item.range
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Range)
		err := b.DecodeMessage(msg)
		m.RangeItem = &RangeItem_Range{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RangeItem_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RangeItem)
	// range_item
	switch x := m.RangeItem.(type) {
	case *RangeItem_Single:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Single))
	case *RangeItem_Range:
		s := proto.Size(x.Range)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Range represents a single low-high pair (e.g. [3-40])
type Range struct {
	Low  uint32 `protobuf:"varint,1,opt,name=low" json:"low,omitempty"`
	High uint32 `protobuf:"varint,2,opt,name=high" json:"high,omitempty"`
}

func (m *Range) Reset()                    { *m = Range{} }
func (m *Range) String() string            { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()               {}
func (*Range) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func init() {
	proto.RegisterType((*Phrase)(nil), "jobsim.Phrase")
	proto.RegisterType((*ReturnStage)(nil), "jobsim.ReturnStage")
	proto.RegisterType((*Stage)(nil), "jobsim.Stage")
	proto.RegisterType((*FailureStage)(nil), "jobsim.FailureStage")
	proto.RegisterType((*StallStage)(nil), "jobsim.StallStage")
	proto.RegisterType((*DepsStage)(nil), "jobsim.DepsStage")
	proto.RegisterType((*Dependency)(nil), "jobsim.Dependency")
	proto.RegisterType((*SparseRange)(nil), "jobsim.SparseRange")
	proto.RegisterType((*RangeItem)(nil), "jobsim.RangeItem")
	proto.RegisterType((*Range)(nil), "jobsim.Range")
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/dm/api/distributor/jobsim/task.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x53, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xfe, 0xb9, 0x89, 0xdd, 0x76, 0xdc, 0xfc, 0x54, 0x16, 0x84, 0xdc, 0x1c, 0xa0, 0xb2, 0x44,
	0xa9, 0x90, 0x6a, 0x43, 0x2a, 0xf5, 0x80, 0xe0, 0x52, 0x55, 0x08, 0x6e, 0x68, 0xcb, 0x89, 0x4b,
	0xb4, 0x89, 0xb7, 0xf6, 0x82, 0xff, 0x69, 0xbd, 0x86, 0xe6, 0x29, 0x78, 0x01, 0x9e, 0x88, 0xa7,
	0x62, 0x77, 0x76, 0x9d, 0x18, 0xb8, 0x44, 0x3b, 0xdf, 0x7c, 0x99, 0xf9, 0xe6, 0x9b, 0x31, 0xbc,
	0xc9, 0x85, 0x2a, 0xfa, 0x55, 0xb2, 0x6e, 0xaa, 0xb4, 0xec, 0xd7, 0x02, 0x7f, 0x2e, 0xf2, 0x26,
	0xcd, 0xaa, 0x94, 0xb5, 0x22, 0xcd, 0x44, 0xa7, 0xa4, 0x58, 0xf5, 0xaa, 0x91, 0xe9, 0x97, 0x66,
	0xd5, 0x89, 0x2a, 0x55, 0xac, 0xfb, 0x9a, 0xb4, 0xb2, 0x51, 0x0d, 0x09, 0x2c, 0x34, 0x7f, 0x9a,
	0x37, 0x4d, 0x5e, 0xf2, 0x14, 0xd1, 0x55, 0x7f, 0x97, 0x2a, 0x51, 0xf1, 0x4e, 0xb1, 0xaa, 0xb5,
	0xc4, 0xf9, 0x93, 0xbf, 0x09, 0x59, 0x2f, 0x99, 0x12, 0x4d, 0x6d, 0xf3, 0xf1, 0x0f, 0x0f, 0x82,
	0x8f, 0x85, 0x64, 0x1d, 0x27, 0x04, 0xa6, 0x35, 0xab, 0x78, 0xe4, 0x9d, 0x7a, 0xe7, 0x87, 0x14,
	0xdf, 0x06, 0xeb, 0x38, 0xcf, 0xa2, 0x3d, 0x8d, 0x4d, 0x28, 0xbe, 0xc9, 0x33, 0x08, 0x74, 0x87,
	0x9c, 0x77, 0xd1, 0xe4, 0x74, 0x72, 0x1e, 0x2e, 0x66, 0x89, 0x15, 0x93, 0xdc, 0x1a, 0x94, 0xba,
	0x24, 0xb9, 0x82, 0x23, 0xc9, 0x55, 0x2f, 0xeb, 0x25, 0x02, 0xd1, 0x54, 0x97, 0x08, 0x17, 0x0f,
	0x07, 0x32, 0xc5, 0x9c, 0xfd, 0x4b, 0x28, 0x77, 0x41, 0xcc, 0x20, 0x1c, 0xe5, 0xc8, 0x63, 0x08,
	0x74, 0xf6, 0x1b, 0x2b, 0x51, 0xd7, 0x84, 0xba, 0x88, 0xbc, 0x06, 0xe0, 0xf7, 0xad, 0xb0, 0xc3,
	0xa0, 0xbe, 0x70, 0x31, 0x4f, 0xec, 0xb4, 0xc9, 0x30, 0x6d, 0xf2, 0x69, 0xb0, 0x83, 0x8e, 0xd8,
	0xf1, 0x4f, 0x0f, 0x7c, 0x5b, 0xfd, 0x25, 0xec, 0xdf, 0x31, 0x51, 0xf6, 0xd2, 0x8e, 0x1d, 0x2e,
	0x1e, 0x0d, 0xfa, 0xde, 0x59, 0x18, 0x69, 0xef, 0xff, 0xa3, 0x03, 0x8d, 0xbc, 0x00, 0x5f, 0x17,
	0x2c, 0x4b, 0xd7, 0x92, 0x8c, 0x86, 0x2f, 0xcb, 0x81, 0x6d, 0x29, 0xe4, 0x39, 0x4c, 0x33, 0xde,
	0x1a, 0x9f, 0x0c, 0xf5, 0xc1, 0x40, 0xbd, 0xd1, 0xd8, 0xc0, 0x44, 0xc2, 0xf5, 0x11, 0x00, 0x9a,
	0xb4, 0x54, 0x9b, 0x96, 0xc7, 0x67, 0x70, 0x34, 0xee, 0x6e, 0x2c, 0x58, 0x17, 0xac, 0x5e, 0x5b,
	0x8d, 0x7b, 0xd4, 0x45, 0xf1, 0x5b, 0x80, 0x5d, 0x57, 0x92, 0x82, 0x9f, 0xf1, 0x92, 0x6d, 0xdc,
	0x20, 0x27, 0xff, 0x78, 0x71, 0xe3, 0x36, 0x4f, 0x2d, 0x2f, 0xbe, 0x84, 0xc3, 0xad, 0x12, 0x72,
	0xe6, 0xa4, 0x7a, 0xb8, 0x52, 0x32, 0x92, 0xca, 0xeb, 0x8c, 0xd7, 0xeb, 0x8d, 0x55, 0x1a, 0xff,
	0xf2, 0x00, 0x76, 0xa0, 0x91, 0xd6, 0x15, 0x4c, 0x66, 0x1d, 0x76, 0x9d, 0x52, 0x17, 0x91, 0x57,
	0x70, 0xc0, 0x94, 0xe2, 0x55, 0xab, 0x3a, 0x67, 0xd4, 0x76, 0xf1, 0xb7, 0x2d, 0x93, 0x1d, 0xa7,
	0xac, 0xc6, 0xf9, 0xb7, 0x34, 0x32, 0x87, 0x7d, 0xbd, 0x5a, 0x29, 0xb8, 0xf5, 0x6b, 0x66, 0x4c,
	0x77, 0x00, 0x39, 0x81, 0x83, 0x4a, 0xdc, 0x2f, 0xf1, 0x14, 0xcd, 0x1d, 0x1d, 0xd0, 0x7d, 0x1d,
	0xdf, 0x9a, 0x6b, 0x3c, 0x83, 0xa0, 0xc5, 0xfb, 0x8d, 0x7c, 0xec, 0xf3, 0xff, 0xd0, 0xc7, 0x5e,
	0x35, 0x75, 0xd9, 0x6b, 0x02, 0xc7, 0xae, 0x95, 0xbe, 0x47, 0x6d, 0x04, 0xcf, 0x37, 0xf1, 0x15,
	0x84, 0x23, 0x35, 0x7a, 0x5d, 0xbe, 0xd0, 0x8c, 0xc1, 0x84, 0xed, 0xbe, 0x30, 0xfb, 0x41, 0x67,
	0xa8, 0xcd, 0xc7, 0x9f, 0xe1, 0x70, 0x8b, 0x91, 0x48, 0x5b, 0x20, 0x6a, 0xed, 0x34, 0x5a, 0x60,
	0x64, 0xbb, 0x58, 0x7f, 0x28, 0xbe, 0x34, 0x34, 0xe7, 0xc0, 0xec, 0x8f, 0x7a, 0xe6, 0x4a, 0x30,
	0x6b, 0x96, 0x8f, 0x8f, 0xa5, 0x29, 0x1e, 0x5f, 0x80, 0x6f, 0xd5, 0x1c, 0xc3, 0xa4, 0x6c, 0xbe,
	0xdb, 0xa2, 0xd4, 0x3c, 0xcd, 0xc7, 0x58, 0x88, 0xbc, 0xc0, 0x72, 0x33, 0x8a, 0xef, 0x55, 0x80,
	0xeb, 0xbd, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x78, 0x97, 0x9a, 0xd3, 0x4f, 0x04, 0x00, 0x00,
}
