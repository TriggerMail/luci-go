// Code generated by protoc-gen-go.
// source: github.com/luci/luci-go/client/cipd/internal/messages/messages.proto
// DO NOT EDIT!

/*
Package messages is a generated protocol buffer package.

It is generated from these files:
	github.com/luci/luci-go/client/cipd/internal/messages/messages.proto

It has these top-level messages:
	BlobWithSHA1
	TagCache
	InstanceCache
*/
package messages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/luci/luci-go/common/proto/google"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// BlobWithSHA1 is a wrapper around a binary blob with SHA1 hash to verify
// its integrity.
type BlobWithSHA1 struct {
	Blob []byte `protobuf:"bytes,1,opt,name=blob,proto3" json:"blob,omitempty"`
	Sha1 []byte `protobuf:"bytes,2,opt,name=sha1,proto3" json:"sha1,omitempty"`
}

func (m *BlobWithSHA1) Reset()                    { *m = BlobWithSHA1{} }
func (m *BlobWithSHA1) String() string            { return proto.CompactTextString(m) }
func (*BlobWithSHA1) ProtoMessage()               {}
func (*BlobWithSHA1) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// TagCache stores a mapping (package name, tag) -> instance ID to speed up
// subsequence ResolveVersion calls when tags are used.
type TagCache struct {
	// Capped list of entries, most recently resolved is last.
	Entries []*TagCache_Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *TagCache) Reset()                    { *m = TagCache{} }
func (m *TagCache) String() string            { return proto.CompactTextString(m) }
func (*TagCache) ProtoMessage()               {}
func (*TagCache) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TagCache) GetEntries() []*TagCache_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type TagCache_Entry struct {
	Package    string `protobuf:"bytes,1,opt,name=package" json:"package,omitempty"`
	Tag        string `protobuf:"bytes,2,opt,name=tag" json:"tag,omitempty"`
	InstanceId string `protobuf:"bytes,3,opt,name=instance_id,json=instanceId" json:"instance_id,omitempty"`
}

func (m *TagCache_Entry) Reset()                    { *m = TagCache_Entry{} }
func (m *TagCache_Entry) String() string            { return proto.CompactTextString(m) }
func (*TagCache_Entry) ProtoMessage()               {}
func (*TagCache_Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

// InstanceCache stores a list of instances in cache
// and their last access time.
type InstanceCache struct {
	// Entries is a map of {instance id -> information about instance}.
	Entries map[string]*InstanceCache_Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// LastSynced is timestamp when we synchronized Entries with actual
	// instance files.
	LastSynced *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=last_synced,json=lastSynced" json:"last_synced,omitempty"`
}

func (m *InstanceCache) Reset()                    { *m = InstanceCache{} }
func (m *InstanceCache) String() string            { return proto.CompactTextString(m) }
func (*InstanceCache) ProtoMessage()               {}
func (*InstanceCache) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *InstanceCache) GetEntries() map[string]*InstanceCache_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *InstanceCache) GetLastSynced() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastSynced
	}
	return nil
}

// Entry stores info about an instance.
type InstanceCache_Entry struct {
	// LastAccess is last time this instance was retrieved from or put to the
	// cache.
	LastAccess *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=last_access,json=lastAccess" json:"last_access,omitempty"`
}

func (m *InstanceCache_Entry) Reset()                    { *m = InstanceCache_Entry{} }
func (m *InstanceCache_Entry) String() string            { return proto.CompactTextString(m) }
func (*InstanceCache_Entry) ProtoMessage()               {}
func (*InstanceCache_Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *InstanceCache_Entry) GetLastAccess() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastAccess
	}
	return nil
}

func init() {
	proto.RegisterType((*BlobWithSHA1)(nil), "messages.BlobWithSHA1")
	proto.RegisterType((*TagCache)(nil), "messages.TagCache")
	proto.RegisterType((*TagCache_Entry)(nil), "messages.TagCache.Entry")
	proto.RegisterType((*InstanceCache)(nil), "messages.InstanceCache")
	proto.RegisterType((*InstanceCache_Entry)(nil), "messages.InstanceCache.Entry")
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/client/cipd/internal/messages/messages.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x90, 0xd1, 0x4a, 0xf3, 0x30,
	0x14, 0xc7, 0xe9, 0xf6, 0xed, 0xdb, 0x76, 0x3a, 0x41, 0x72, 0x55, 0x0a, 0x32, 0x19, 0x5e, 0xec,
	0xc6, 0x96, 0x6d, 0x20, 0xa2, 0x20, 0x4c, 0x27, 0xb8, 0xdb, 0x6e, 0x20, 0x5e, 0x8d, 0x34, 0x8d,
	0x5d, 0x58, 0xd6, 0x94, 0x25, 0x15, 0xfa, 0x1e, 0xbe, 0x86, 0xef, 0x68, 0x9a, 0x2e, 0xd3, 0x79,
	0x21, 0xde, 0x94, 0x73, 0x4e, 0x7e, 0xcd, 0xf9, 0xe5, 0x0f, 0xb3, 0x94, 0xa9, 0x75, 0x11, 0x07,
	0x44, 0x6c, 0x43, 0x5e, 0x10, 0x66, 0x3e, 0x97, 0xa9, 0x08, 0x09, 0x67, 0x34, 0x53, 0x21, 0x61,
	0x79, 0x12, 0xb2, 0x4c, 0xd1, 0x5d, 0x86, 0x79, 0xb8, 0xa5, 0x52, 0xe2, 0x94, 0xca, 0x43, 0x11,
	0xe4, 0x3b, 0xa1, 0x04, 0xea, 0xd8, 0xde, 0xef, 0xa7, 0x42, 0xa4, 0x9c, 0x86, 0x66, 0x1e, 0x17,
	0xaf, 0xa1, 0x62, 0xfa, 0x4c, 0xe1, 0x6d, 0x5e, 0xa3, 0x83, 0x2b, 0xe8, 0xdd, 0x73, 0x11, 0x3f,
	0xeb, 0xb5, 0x8b, 0xa7, 0xe9, 0x08, 0x21, 0xf8, 0x17, 0xeb, 0xde, 0x73, 0xce, 0x9d, 0x61, 0x2f,
	0x32, 0x75, 0x35, 0x93, 0x6b, 0x3c, 0xf2, 0x1a, 0xf5, 0xac, 0xaa, 0x07, 0xef, 0x0e, 0x74, 0x96,
	0x38, 0x7d, 0xc0, 0x64, 0x4d, 0xd1, 0x18, 0xda, 0x5a, 0x6e, 0xc7, 0xa8, 0xd4, 0xff, 0x35, 0x87,
	0xee, 0xd8, 0x0b, 0x0e, 0x46, 0x16, 0x0a, 0x1e, 0x35, 0x51, 0x46, 0x16, 0xf4, 0x97, 0xd0, 0x32,
	0x13, 0xe4, 0x41, 0x3b, 0xc7, 0x64, 0xa3, 0x61, 0xb3, 0xb4, 0x1b, 0xd9, 0x16, 0x9d, 0x42, 0x53,
	0xe1, 0xd4, 0xac, 0xed, 0x46, 0x55, 0x89, 0xfa, 0xe0, 0xb2, 0x4c, 0xeb, 0x67, 0x84, 0xae, 0x58,
	0xe2, 0x35, 0xcd, 0x09, 0xd8, 0xd1, 0x3c, 0x19, 0x7c, 0x34, 0xe0, 0x64, 0xbe, 0x6f, 0x6b, 0xb7,
	0xbb, 0x9f, 0x6e, 0x17, 0x5f, 0x6e, 0x47, 0xa4, 0x11, 0xd4, 0xd8, 0xb1, 0x27, 0xba, 0x05, 0x97,
	0x63, 0xa9, 0x56, 0xb2, 0xd4, 0x60, 0x62, 0x64, 0xdc, 0xb1, 0x1f, 0xd4, 0xb9, 0x06, 0x36, 0xd7,
	0x60, 0x69, 0x73, 0x8d, 0xa0, 0xc2, 0x17, 0x86, 0xf6, 0x67, 0xf6, 0x91, 0xf6, 0x16, 0x4c, 0x88,
	0x5e, 0xfe, 0xd7, 0x5b, 0xa6, 0x86, 0xf6, 0x5f, 0xa0, 0xf7, 0xdd, 0xad, 0xca, 0x65, 0x43, 0xcb,
	0x7d, 0x5a, 0x55, 0x89, 0x26, 0xd0, 0x7a, 0xc3, 0xbc, 0xa0, 0xfb, 0x8b, 0xcf, 0x7e, 0x7b, 0x62,
	0x19, 0xd5, 0xec, 0x4d, 0xe3, 0xda, 0x89, 0xff, 0x9b, 0xd5, 0x93, 0xcf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xf9, 0x18, 0xc7, 0x31, 0x78, 0x02, 0x00, 0x00,
}
