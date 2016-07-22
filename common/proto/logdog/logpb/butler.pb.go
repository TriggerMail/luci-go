// Code generated by protoc-gen-go.
// source: github.com/luci/luci-go/common/proto/logdog/logpb/butler.proto
// DO NOT EDIT!

/*
Package logpb is a generated protocol buffer package.

It is generated from these files:
	github.com/luci/luci-go/common/proto/logdog/logpb/butler.proto
	github.com/luci/luci-go/common/proto/logdog/logpb/log.proto

It has these top-level messages:
	ButlerMetadata
	ButlerLogBundle
	LogStreamDescriptor
	Text
	Binary
	Datagram
	LogEntry
	LogIndex
*/
package logpb

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

//
// This enumerates the possible contents of published Butler data.
type ButlerMetadata_ContentType int32

const (
	// An invalid content type. Do not use.
	ButlerMetadata_Invalid ButlerMetadata_ContentType = 0
	// The published data is a ButlerLogBundle protobuf message.
	ButlerMetadata_ButlerLogBundle ButlerMetadata_ContentType = 1
)

var ButlerMetadata_ContentType_name = map[int32]string{
	0: "Invalid",
	1: "ButlerLogBundle",
}
var ButlerMetadata_ContentType_value = map[string]int32{
	"Invalid":         0,
	"ButlerLogBundle": 1,
}

func (x ButlerMetadata_ContentType) String() string {
	return proto.EnumName(ButlerMetadata_ContentType_name, int32(x))
}
func (ButlerMetadata_ContentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

// Compression scheme of attached data.
type ButlerMetadata_Compression int32

const (
	ButlerMetadata_NONE ButlerMetadata_Compression = 0
	ButlerMetadata_ZLIB ButlerMetadata_Compression = 1
)

var ButlerMetadata_Compression_name = map[int32]string{
	0: "NONE",
	1: "ZLIB",
}
var ButlerMetadata_Compression_value = map[string]int32{
	"NONE": 0,
	"ZLIB": 1,
}

func (x ButlerMetadata_Compression) String() string {
	return proto.EnumName(ButlerMetadata_Compression_name, int32(x))
}
func (ButlerMetadata_Compression) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 1}
}

//
// ButlerMetadata appears as a frame at the beginning of Butler published data
// to describe the remainder of the contents.
type ButlerMetadata struct {
	// This is the type of data in the subsequent frame.
	Type        ButlerMetadata_ContentType `protobuf:"varint,1,opt,name=type,enum=logpb.ButlerMetadata_ContentType" json:"type,omitempty"`
	Compression ButlerMetadata_Compression `protobuf:"varint,2,opt,name=compression,enum=logpb.ButlerMetadata_Compression" json:"compression,omitempty"`
	// The protobuf version string (see version.go).
	ProtoVersion string `protobuf:"bytes,3,opt,name=proto_version,json=protoVersion" json:"proto_version,omitempty"`
}

func (m *ButlerMetadata) Reset()                    { *m = ButlerMetadata{} }
func (m *ButlerMetadata) String() string            { return proto.CompactTextString(m) }
func (*ButlerMetadata) ProtoMessage()               {}
func (*ButlerMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

//
// A message containing log data in transit from the Butler.
//
// The Butler is capable of conserving bandwidth by bundling collected log
// messages together into this protocol buffer. Based on Butler bundling
// settings, this message can represent anything from a single LogRecord to
// multiple LogRecords belonging to several different streams.
//
// Entries in a Log Bundle are fully self-descriptive: no additional information
// is needed to fully associate the contained data with its proper place in
// the source log stream.
type ButlerLogBundle struct {
	//
	// (DEPRECATED) Stream source information. Now supplied during prefix
	// registration.
	DeprecatedSource string `protobuf:"bytes,1,opt,name=deprecated_source,json=deprecatedSource" json:"deprecated_source,omitempty"`
	// The timestamp when this bundle was generated.
	//
	// This field will be used for debugging and internal accounting.
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	// *
	// Each Entry is an individual set of log records for a given log stream.
	Entries []*ButlerLogBundle_Entry `protobuf:"bytes,3,rep,name=entries" json:"entries,omitempty"`
	// * Project specifies which luci-config project this stream belongs to.
	Project string `protobuf:"bytes,4,opt,name=project" json:"project,omitempty"`
	// *
	// The log stream prefix that is shared by all bundled streams.
	//
	// This prefix is valid within the supplied project scope.
	Prefix string `protobuf:"bytes,5,opt,name=prefix" json:"prefix,omitempty"`
	//
	// The log prefix's secret value (required).
	//
	// The secret is bound to all log streams that share the supplied Prefix, and
	// The Coordinator will record the secret associated with a given log Prefix,
	// but will not expose the secret to users.
	//
	// The Collector will check the secret prior to ingesting logs. If the
	// secret doesn't match the value recorded by the Coordinator, the log
	// will be discarded.
	//
	// This ensures that only the Butler instance that generated the log stream
	// can emit log data for that stream. It also ensures that only authenticated
	// users can write to a Prefix.
	Secret []byte `protobuf:"bytes,6,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (m *ButlerLogBundle) Reset()                    { *m = ButlerLogBundle{} }
func (m *ButlerLogBundle) String() string            { return proto.CompactTextString(m) }
func (*ButlerLogBundle) ProtoMessage()               {}
func (*ButlerLogBundle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ButlerLogBundle) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ButlerLogBundle) GetEntries() []*ButlerLogBundle_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

//
// A bundle Entry describes a set of LogEntry messages originating from the
// same log stream.
type ButlerLogBundle_Entry struct {
	//
	// The descriptor for this entry's log stream.
	//
	// Each LogEntry in the "logs" field is shares this common descriptor.
	Desc *LogStreamDescriptor `protobuf:"bytes,1,opt,name=desc" json:"desc,omitempty"`
	// (DEPRECATED) Per-entry secret replaced with Butler-wide secret.
	DeprecatedEntrySecret []byte `protobuf:"bytes,2,opt,name=deprecated_entry_secret,json=deprecatedEntrySecret,proto3" json:"deprecated_entry_secret,omitempty"`
	//
	// Whether this log entry terminates its stream.
	//
	// If present and "true", this field declares that this Entry is the last
	// such entry in the stream. This fact is recorded by the Collector and
	// registered with the Coordinator. The largest stream prefix in this Entry
	// will be bound the stream's LogEntry records to [0:largest_prefix]. Once
	// all messages in that range have been received, the log may be archived.
	//
	// Further log entries belonging to this stream with stream indices
	// exceeding the terminal log's index will be discarded.
	Terminal bool `protobuf:"varint,3,opt,name=terminal" json:"terminal,omitempty"`
	//
	// If terminal is true, this is the terminal stream index; that is, the last
	// message index in the stream.
	TerminalIndex uint64 `protobuf:"varint,4,opt,name=terminal_index,json=terminalIndex" json:"terminal_index,omitempty"`
	//
	// Log entries attached to this record. These MUST be sequential.
	//
	// This is the main log entry content.
	Logs []*LogEntry `protobuf:"bytes,5,rep,name=logs" json:"logs,omitempty"`
}

func (m *ButlerLogBundle_Entry) Reset()                    { *m = ButlerLogBundle_Entry{} }
func (m *ButlerLogBundle_Entry) String() string            { return proto.CompactTextString(m) }
func (*ButlerLogBundle_Entry) ProtoMessage()               {}
func (*ButlerLogBundle_Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *ButlerLogBundle_Entry) GetDesc() *LogStreamDescriptor {
	if m != nil {
		return m.Desc
	}
	return nil
}

func (m *ButlerLogBundle_Entry) GetLogs() []*LogEntry {
	if m != nil {
		return m.Logs
	}
	return nil
}

func init() {
	proto.RegisterType((*ButlerMetadata)(nil), "logpb.ButlerMetadata")
	proto.RegisterType((*ButlerLogBundle)(nil), "logpb.ButlerLogBundle")
	proto.RegisterType((*ButlerLogBundle_Entry)(nil), "logpb.ButlerLogBundle.Entry")
	proto.RegisterEnum("logpb.ButlerMetadata_ContentType", ButlerMetadata_ContentType_name, ButlerMetadata_ContentType_value)
	proto.RegisterEnum("logpb.ButlerMetadata_Compression", ButlerMetadata_Compression_name, ButlerMetadata_Compression_value)
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/common/proto/logdog/logpb/butler.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0xb5, 0xdb, 0xf4, 0x6b, 0xba, 0xdb, 0xad, 0x23, 0x6a, 0x28, 0x82, 0x6e, 0x45, 0x10, 0xc4,
	0x09, 0xac, 0x28, 0x82, 0xe0, 0x43, 0xd7, 0x7d, 0x28, 0xac, 0x2b, 0x4c, 0x17, 0x1f, 0x7c, 0x29,
	0xf9, 0xb8, 0x1b, 0x47, 0x92, 0x4c, 0x98, 0x4c, 0x96, 0xed, 0x0f, 0xf0, 0x1f, 0x0a, 0xfe, 0x1d,
	0x6f, 0x6e, 0x9a, 0xa6, 0x8b, 0x20, 0xf8, 0x92, 0xe6, 0x9e, 0x7b, 0xee, 0xbd, 0xe7, 0x9c, 0x86,
	0x7d, 0x8c, 0x95, 0xfd, 0x5e, 0x06, 0x22, 0xd4, 0xa9, 0x97, 0x94, 0xa1, 0xa2, 0xc7, 0xeb, 0x58,
	0x7b, 0x08, 0xa4, 0x3a, 0xf3, 0x72, 0xa3, 0xad, 0xf6, 0x12, 0x1d, 0x47, 0x3a, 0xae, 0x7e, 0xf2,
	0xc0, 0x0b, 0x4a, 0x9b, 0x80, 0x11, 0xd4, 0xe1, 0x3d, 0xc2, 0x66, 0x1f, 0xfe, 0x7f, 0x0d, 0x3e,
	0xeb, 0x1d, 0xb3, 0xa7, 0xb1, 0xd6, 0x71, 0x02, 0x35, 0x29, 0x28, 0xaf, 0x3d, 0xab, 0x52, 0x28,
	0xac, 0x9f, 0xe6, 0x35, 0x61, 0xfe, 0xf3, 0x80, 0x4d, 0x16, 0x74, 0xf5, 0x33, 0x58, 0x3f, 0xf2,
	0xad, 0xcf, 0xdf, 0x32, 0xc7, 0x6e, 0x72, 0x70, 0x3b, 0xcf, 0x3a, 0x2f, 0x27, 0xa7, 0x27, 0x82,
	0x76, 0x8a, 0xbb, 0x24, 0x71, 0xa6, 0x33, 0x0b, 0x99, 0xbd, 0x42, 0xa2, 0x24, 0x3a, 0x3f, 0x63,
	0x63, 0x54, 0x94, 0x1b, 0x28, 0x0a, 0xa5, 0x33, 0xf7, 0xe0, 0xdf, 0xd3, 0x3b, 0xa2, 0xdc, 0x9f,
	0xe2, 0xcf, 0xd9, 0x11, 0xe9, 0x5a, 0xdf, 0x80, 0xa1, 0x35, 0x5d, 0x5c, 0x33, 0x92, 0x87, 0x04,
	0x7e, 0xad, 0xb1, 0xb9, 0xc7, 0xc6, 0x7b, 0xe7, 0xf9, 0x98, 0x0d, 0x96, 0xd9, 0x8d, 0x9f, 0xa8,
	0x68, 0x7a, 0x8f, 0x3f, 0x60, 0xc7, 0xf5, 0xad, 0x0b, 0x1d, 0x2f, 0xca, 0x2c, 0x4a, 0x60, 0xda,
	0x99, 0x9f, 0x54, 0x03, 0xed, 0x91, 0x21, 0x73, 0x2e, 0xbf, 0x5c, 0x9e, 0x23, 0x1b, 0xdf, 0xbe,
	0x5d, 0x2c, 0x17, 0x48, 0xf9, 0xd5, 0xfd, 0x6b, 0x90, 0xbf, 0x62, 0xf7, 0x23, 0xc0, 0xa9, 0xd0,
	0xb7, 0x10, 0xad, 0x0b, 0x5d, 0x9a, 0xb0, 0x4e, 0x65, 0x24, 0xa7, 0x6d, 0x63, 0x45, 0x38, 0x7f,
	0xcf, 0x46, 0xbb, 0x6c, 0xc9, 0xfc, 0xf8, 0x74, 0x26, 0xea, 0xf4, 0x45, 0x93, 0xbe, 0xb8, 0x6a,
	0x18, 0xb2, 0x25, 0xf3, 0x77, 0x6c, 0x80, 0x56, 0x8c, 0x82, 0x02, 0xdd, 0x76, 0x71, 0xee, 0xc9,
	0x9d, 0xd0, 0x76, 0x7a, 0xc4, 0x39, 0xb2, 0x36, 0xb2, 0x21, 0x73, 0x97, 0x0d, 0x70, 0xf1, 0x0f,
	0x08, 0xad, 0xeb, 0x90, 0xa8, 0xa6, 0xe4, 0x8f, 0x58, 0x1f, 0xd5, 0x5d, 0xab, 0x5b, 0xb7, 0x47,
	0x8d, 0x6d, 0x55, 0xe1, 0x05, 0x84, 0x06, 0xac, 0xdb, 0x47, 0xfc, 0x50, 0x6e, 0xab, 0xd9, 0xef,
	0x0e, 0xeb, 0xd1, 0x72, 0x2e, 0x98, 0x13, 0x41, 0x11, 0x92, 0xcb, 0xca, 0x40, 0x2d, 0x04, 0x25,
	0xac, 0xac, 0x01, 0x3f, 0xfd, 0x84, 0x3d, 0xa3, 0x72, 0xab, 0x8d, 0x24, 0x1e, 0x6a, 0x7f, 0xbc,
	0x17, 0x51, 0xa5, 0x6c, 0xb3, 0xde, 0x9e, 0x38, 0xa0, 0x13, 0x0f, 0xdb, 0x36, 0x5d, 0x58, 0x51,
	0x93, 0xcf, 0xd8, 0xd0, 0x82, 0x49, 0x55, 0xe6, 0x27, 0xf4, 0x17, 0x0f, 0xe5, 0xae, 0xe6, 0x2f,
	0xd8, 0xa4, 0x79, 0x5f, 0xab, 0x2c, 0x82, 0x5b, 0xb2, 0xe7, 0xc8, 0xa3, 0x06, 0x5d, 0x56, 0x20,
	0x7e, 0x2a, 0x0e, 0xaa, 0x2b, 0xd0, 0x62, 0x95, 0xd9, 0x71, 0x2b, 0xb5, 0x8e, 0x89, 0x9a, 0x41,
	0x9f, 0xa2, 0x7f, 0xf3, 0x27, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x23, 0x24, 0x1b, 0x8c, 0x03, 0x00,
	0x00,
}
