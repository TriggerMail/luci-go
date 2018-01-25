// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/scheduler/appengine/internal/timers.proto

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	go.chromium.org/luci/scheduler/appengine/internal/timers.proto
	go.chromium.org/luci/scheduler/appengine/internal/tq.proto
	go.chromium.org/luci/scheduler/appengine/internal/triggers.proto

It has these top-level messages:
	Timer
	TimerList
	ReadProjectConfigTask
	LaunchInvocationTask
	LaunchInvocationsBatchTask
	TriageJobStateTask
	InvocationFinishedTask
	FanOutTriggersTask
	EnqueueTriggersTask
	ScheduleTimersTask
	TimerTask
	Trigger
	TriggerList
*/
package internal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Timer can be emitted by any invocation if it wants to be poked later.
//
// Timers are scoped to single invocation and owned by it, so we don't include
// invocation reference here. It is always available from the context of calls.
type Timer struct {
	// Unique in time identifier of this timer, auto-generated.
	//
	// It is used to deduplicate and hence provide idempotency for adding
	// timers.
	//
	// Set by the engine, can't be overridden.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Timestamp when the timer was created.
	//
	// Set by the engine, can't be overridden.
	Created *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created" json:"created,omitempty"`
	// Target time when this timer activates.
	//
	// Should be provided by whoever emits the timer.
	Eta *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=eta" json:"eta,omitempty"`
	// User friendly name for this timer that shows up in UI.
	//
	// Can be provided by whoever emits the timer. Doesn't have to be unique.
	Title string `protobuf:"bytes,4,opt,name=title" json:"title,omitempty"`
	// Arbitrary optional payload passed verbatim to the invocation.
	Payload []byte `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Timer) Reset()                    { *m = Timer{} }
func (m *Timer) String() string            { return proto.CompactTextString(m) }
func (*Timer) ProtoMessage()               {}
func (*Timer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Timer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Timer) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Timer) GetEta() *google_protobuf.Timestamp {
	if m != nil {
		return m.Eta
	}
	return nil
}

func (m *Timer) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Timer) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// TimerList is what we store in datastore entities.
type TimerList struct {
	Timers []*Timer `protobuf:"bytes,1,rep,name=timers" json:"timers,omitempty"`
}

func (m *TimerList) Reset()                    { *m = TimerList{} }
func (m *TimerList) String() string            { return proto.CompactTextString(m) }
func (*TimerList) ProtoMessage()               {}
func (*TimerList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TimerList) GetTimers() []*Timer {
	if m != nil {
		return m.Timers
	}
	return nil
}

func init() {
	proto.RegisterType((*Timer)(nil), "internal.timers.Timer")
	proto.RegisterType((*TimerList)(nil), "internal.timers.TimerList")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/scheduler/appengine/internal/timers.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0x3d, 0x4f, 0xf3, 0x30,
	0x10, 0xc7, 0xe5, 0xe4, 0x49, 0xfb, 0xf4, 0x8a, 0x40, 0xb2, 0x10, 0xb2, 0xba, 0x10, 0x75, 0xca,
	0x80, 0x6c, 0xa9, 0xb0, 0x21, 0x31, 0x30, 0x33, 0x45, 0x9d, 0xd8, 0xdc, 0xf8, 0x70, 0x2d, 0x39,
	0xb1, 0xe5, 0x38, 0x03, 0x9f, 0x89, 0x2f, 0x89, 0x6a, 0xe3, 0x85, 0x85, 0xf1, 0xee, 0x7e, 0xf7,
	0x7f, 0x81, 0x17, 0xed, 0xf8, 0x70, 0x0e, 0x6e, 0x34, 0xcb, 0xc8, 0x5d, 0xd0, 0xc2, 0x2e, 0x83,
	0x11, 0xf3, 0x70, 0x46, 0xb5, 0x58, 0x0c, 0x42, 0x7a, 0x8f, 0x93, 0x36, 0x13, 0x0a, 0x33, 0x45,
	0x0c, 0x93, 0xb4, 0x22, 0x9a, 0x11, 0xc3, 0xcc, 0x7d, 0x70, 0xd1, 0xd1, 0x9b, 0xb2, 0xe6, 0x79,
	0xbd, 0xbb, 0xd7, 0xce, 0x69, 0x8b, 0x22, 0x9d, 0x4f, 0xcb, 0x47, 0xc2, 0xe7, 0x28, 0x47, 0x9f,
	0x3f, 0xf6, 0x5f, 0x04, 0x9a, 0xe3, 0x85, 0xa5, 0xd7, 0x50, 0x19, 0xc5, 0x48, 0x4b, 0xba, 0x4d,
	0x5f, 0x19, 0x45, 0x9f, 0x60, 0x3d, 0x04, 0x94, 0x11, 0x15, 0xab, 0x5a, 0xd2, 0x6d, 0x0f, 0x3b,
	0x9e, 0xc5, 0x78, 0x11, 0xe3, 0xc7, 0x22, 0xd6, 0x17, 0x94, 0x3e, 0x40, 0x8d, 0x51, 0xb2, 0xfa,
	0xcf, 0x8f, 0x0b, 0x46, 0x6f, 0xa1, 0x89, 0x26, 0x5a, 0x64, 0xff, 0x92, 0x6d, 0x1e, 0x28, 0x83,
	0xb5, 0x97, 0x9f, 0xd6, 0x49, 0xc5, 0x9a, 0x96, 0x74, 0x57, 0x7d, 0x19, 0xf7, 0xcf, 0xb0, 0x49,
	0x61, 0xdf, 0xcc, 0x1c, 0x29, 0x87, 0x55, 0x6e, 0xc9, 0x48, 0x5b, 0x77, 0xdb, 0xc3, 0x1d, 0xff,
	0xd5, 0x3e, 0xb9, 0x85, 0xfe, 0x87, 0x7a, 0x85, 0xf7, 0xff, 0x05, 0x38, 0xad, 0x52, 0xa2, 0xc7,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x1c, 0xc7, 0x63, 0x71, 0x01, 0x00, 0x00,
}
