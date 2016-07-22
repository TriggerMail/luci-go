// Code generated by protoc-gen-go.
// source: github.com/luci/luci-go/common/api/logdog_coordinator/services/v1/tasks.proto
// DO NOT EDIT!

package logdog

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/luci/luci-go/common/proto/google"
import google_protobuf2 "github.com/luci/luci-go/common/proto/google"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ArchiveTask is a task queue task description for the archival of a single
// log stream.
type ArchiveTask struct {
	// The name of the project that this stream is bound to.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// The hash ID of the log stream to archive.
	Id string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	// The archival key of the log stream. If this key doesn't match the key in
	// the log stream state, the request is superfluous and should be deleted.
	Key []byte `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	// Don't waste time archiving the log stream until it is at least this old.
	//
	// This is in place to prevent overly-aggressive archivals from wasting time
	// trying, then failing, becuase the log stream data is still being collected
	// into intermediate storage.
	SettleDelay *google_protobuf.Duration `protobuf:"bytes,4,opt,name=settle_delay,json=settleDelay" json:"settle_delay,omitempty"`
	// The amount of time after the task was created that log stream completeness
	// will be used as a success criteria. If the task's age is older than this
	// value, completeness will not be enforced.
	//
	// The task's age can be calculated by subtracting its lease expiration time
	// (leaseTimestamp) from its enqueued timestamp (enqueueTimestamp).
	CompletePeriod *google_protobuf.Duration `protobuf:"bytes,5,opt,name=complete_period,json=completePeriod" json:"complete_period,omitempty"`
	// The time when this archive task was dispatched.
	//
	// This time is optional, and will be based on the Coordinator's clock. If not
	// zero, it can be used by the Archivist to avoid superfluous archival
	// processing by asserting that IF this time is close to the Archivist's local
	// clock by a specific threshold, it will punt the archival task.
	//
	// Because archival is dispatched by Tumble, the actual encoding of archival
	// parameters is oftentimes delayed such that the request is dispatched to
	// Pub/Sub before the datastore has been updated.
	DispatchedAt *google_protobuf2.Timestamp `protobuf:"bytes,6,opt,name=dispatched_at,json=dispatchedAt" json:"dispatched_at,omitempty"`
}

func (m *ArchiveTask) Reset()                    { *m = ArchiveTask{} }
func (m *ArchiveTask) String() string            { return proto.CompactTextString(m) }
func (*ArchiveTask) ProtoMessage()               {}
func (*ArchiveTask) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ArchiveTask) GetSettleDelay() *google_protobuf.Duration {
	if m != nil {
		return m.SettleDelay
	}
	return nil
}

func (m *ArchiveTask) GetCompletePeriod() *google_protobuf.Duration {
	if m != nil {
		return m.CompletePeriod
	}
	return nil
}

func (m *ArchiveTask) GetDispatchedAt() *google_protobuf2.Timestamp {
	if m != nil {
		return m.DispatchedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*ArchiveTask)(nil), "logdog.ArchiveTask")
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/common/api/logdog_coordinator/services/v1/tasks.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0x3b, 0x31,
	0x10, 0xc5, 0x69, 0xfb, 0xff, 0x57, 0x4c, 0x6b, 0x95, 0x9c, 0x62, 0x0f, 0x2a, 0x9e, 0xbc, 0xb8,
	0x41, 0xbd, 0x0a, 0x52, 0xe9, 0x55, 0x90, 0xa5, 0xf7, 0x25, 0x4d, 0xc6, 0x6d, 0xec, 0x6e, 0x27,
	0x24, 0xb3, 0x0b, 0xfd, 0x0e, 0x7e, 0x68, 0xd3, 0xac, 0x8b, 0xa0, 0x07, 0x2f, 0x61, 0xf2, 0xe6,
	0xf7, 0x1e, 0x8f, 0x61, 0x2f, 0xa5, 0xa5, 0x4d, 0xb3, 0xce, 0x34, 0xd6, 0xb2, 0x6a, 0xb4, 0x4d,
	0xcf, 0x6d, 0x89, 0x32, 0x0a, 0x35, 0xee, 0xa4, 0x72, 0x51, 0xc2, 0xd2, 0x60, 0x59, 0x68, 0x44,
	0x6f, 0xec, 0x4e, 0x11, 0x7a, 0x19, 0xc0, 0xb7, 0x56, 0x43, 0x90, 0xed, 0x9d, 0x24, 0x15, 0xb6,
	0x21, 0x73, 0x1e, 0x09, 0xf9, 0xb8, 0x63, 0xe7, 0x17, 0x25, 0x62, 0x59, 0x81, 0x4c, 0xea, 0xba,
	0x79, 0x93, 0xa6, 0xf1, 0x8a, 0x2c, 0xee, 0x3a, 0x6e, 0x7e, 0xf9, 0x73, 0x4f, 0xb6, 0x86, 0x40,
	0xaa, 0x76, 0x1d, 0x70, 0xfd, 0x31, 0x64, 0x93, 0x85, 0xd7, 0x1b, 0xdb, 0xc2, 0x2a, 0xe6, 0x73,
	0xc1, 0x8e, 0xe2, 0xe2, 0x1d, 0x34, 0x89, 0xc1, 0xd5, 0xe0, 0xe6, 0x38, 0xef, 0xbf, 0x7c, 0xc6,
	0x86, 0xd6, 0x88, 0x61, 0x12, 0xe3, 0xc4, 0xcf, 0xd8, 0x68, 0x0b, 0x7b, 0x31, 0x8a, 0xc2, 0x34,
	0x3f, 0x8c, 0xfc, 0x91, 0x4d, 0x03, 0x10, 0x55, 0x50, 0x18, 0xa8, 0xd4, 0x5e, 0xfc, 0x8b, 0xab,
	0xc9, 0xfd, 0x79, 0xd6, 0x75, 0xc8, 0xfa, 0x0e, 0xd9, 0xf2, 0xab, 0x63, 0x3e, 0xe9, 0xf0, 0xe5,
	0x81, 0xe6, 0xcf, 0xec, 0x34, 0xde, 0xc2, 0x55, 0x40, 0x50, 0x38, 0xf0, 0x16, 0x8d, 0xf8, 0xff,
	0x57, 0xc0, 0xac, 0x77, 0xbc, 0x26, 0x03, 0x7f, 0x62, 0x27, 0xc6, 0x06, 0xa7, 0x48, 0x6f, 0xc0,
	0x14, 0x8a, 0xc4, 0x38, 0x25, 0xcc, 0x7f, 0x25, 0xac, 0xfa, 0x33, 0xe4, 0xd3, 0x6f, 0xc3, 0x82,
	0xd6, 0xe3, 0x44, 0x3c, 0x7c, 0x06, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xad, 0x12, 0x41, 0xaf, 0x01,
	0x00, 0x00,
}
