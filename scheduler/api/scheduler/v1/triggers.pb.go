// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/scheduler/api/scheduler/v1/triggers.proto

package scheduler

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// NoopTrigger is used by Scheduler integration tests to represent test
// triggers.
type NoopTrigger struct {
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *NoopTrigger) Reset()                    { *m = NoopTrigger{} }
func (m *NoopTrigger) String() string            { return proto.CompactTextString(m) }
func (*NoopTrigger) ProtoMessage()               {}
func (*NoopTrigger) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *NoopTrigger) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// GitilesTrigger is emitted by sources that watch Gitiles and consumed by
// Buildbucket tasks.
//
// Such triggers are emitted whenever the repository state changes.
type GitilesTrigger struct {
	Repo     string `protobuf:"bytes,1,opt,name=repo" json:"repo,omitempty"`
	Ref      string `protobuf:"bytes,2,opt,name=ref" json:"ref,omitempty"`
	Revision string `protobuf:"bytes,3,opt,name=revision" json:"revision,omitempty"`
}

func (m *GitilesTrigger) Reset()                    { *m = GitilesTrigger{} }
func (m *GitilesTrigger) String() string            { return proto.CompactTextString(m) }
func (*GitilesTrigger) ProtoMessage()               {}
func (*GitilesTrigger) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *GitilesTrigger) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *GitilesTrigger) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

func (m *GitilesTrigger) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func init() {
	proto.RegisterType((*NoopTrigger)(nil), "scheduler.NoopTrigger")
	proto.RegisterType((*GitilesTrigger)(nil), "scheduler.GitilesTrigger")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/scheduler/api/scheduler/v1/triggers.proto", fileDescriptor1)
}

var fileDescriptor1 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0x31, 0xef, 0x82, 0x30,
	0x10, 0x47, 0xc3, 0x9f, 0x7f, 0x8c, 0x9c, 0x89, 0x31, 0x9d, 0x88, 0x93, 0x32, 0x39, 0xd1, 0x18,
	0x77, 0x57, 0x37, 0x07, 0xe2, 0x17, 0x40, 0x38, 0xcb, 0x25, 0xe0, 0x35, 0xd7, 0x96, 0xcf, 0x6f,
	0xa8, 0x91, 0xb0, 0xbd, 0xdf, 0xbb, 0x37, 0x1c, 0x5c, 0x0d, 0x97, 0x4d, 0x27, 0x3c, 0x50, 0x18,
	0x4a, 0x16, 0xa3, 0xfb, 0xd0, 0x90, 0x76, 0x4d, 0x87, 0x6d, 0xe8, 0x51, 0x74, 0x6d, 0x97, 0x6b,
	0x3c, 0x6b, 0x2f, 0x64, 0x0c, 0x8a, 0x2b, 0xad, 0xb0, 0x67, 0x95, 0xcd, 0xc7, 0xe2, 0x08, 0x9b,
	0x3b, 0xb3, 0x7d, 0x7c, 0x03, 0xa5, 0xe0, 0xbf, 0xad, 0x7d, 0x9d, 0x27, 0x87, 0xe4, 0x94, 0x55,
	0x91, 0x8b, 0x0a, 0xb6, 0x37, 0xf2, 0xd4, 0xa3, 0x5b, 0x54, 0x82, 0x96, 0x7f, 0xd5, 0xc4, 0x6a,
	0x07, 0xa9, 0xe0, 0x2b, 0xff, 0x8b, 0x6a, 0x42, 0xb5, 0x87, 0xb5, 0xe0, 0x48, 0x8e, 0xf8, 0x9d,
	0xa7, 0x51, 0xcf, 0xfb, 0xb9, 0x8a, 0x8f, 0x5c, 0x3e, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x39,
	0x03, 0x51, 0xca, 0x00, 0x00, 0x00,
}
