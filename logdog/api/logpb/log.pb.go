// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/logdog/api/logpb/log.proto

package logpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// A log stream type.
type StreamType int32

const (
	StreamType_TEXT     StreamType = 0
	StreamType_BINARY   StreamType = 1
	StreamType_DATAGRAM StreamType = 2
)

var StreamType_name = map[int32]string{
	0: "TEXT",
	1: "BINARY",
	2: "DATAGRAM",
}

var StreamType_value = map[string]int32{
	"TEXT":     0,
	"BINARY":   1,
	"DATAGRAM": 2,
}

func (x StreamType) String() string {
	return proto.EnumName(StreamType_name, int32(x))
}

func (StreamType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_30887c96a468dac0, []int{0}
}

// *
// Log stream descriptor data. This is the full set of information that
// describes a logging stream.
type LogStreamDescriptor struct {
	//
	// The stream's prefix (required).
	//
	// Logs originating from the same Butler instance will share a Prefix.
	//
	// A valid prefix value is a StreamName described in:
	// https://go.chromium.org/luci/common/logdog/types
	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	//
	// The log stream's name (required).
	//
	// This is used to uniquely identify a log stream within the scope of its
	// prefix.
	//
	// A valid name value is a StreamName described in:
	// https://go.chromium.org/luci/common/logdog/types
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The log stream's content type (required).
	StreamType StreamType `protobuf:"varint,3,opt,name=stream_type,json=streamType,proto3,enum=logpb.StreamType" json:"stream_type,omitempty"`
	//
	// The stream's content type (required).
	//
	// This must be an HTTP Content-Type value. It is made available to LogDog
	// clients when querying stream metadata. It will also be applied to archived
	// binary log data.
	ContentType string `protobuf:"bytes,4,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	//
	// The log stream's base timestamp (required).
	//
	// This notes the start time of the log stream. All LogEntries express their
	// timestamp as microsecond offsets from this field.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	//
	// Tag is an arbitrary key/value tag associated with this log stream.
	//
	// LogDog clients can query for log streams based on tag values.
	Tags map[string]string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//
	// If set, the stream will be joined together during archival to recreate the
	// original stream and made available at <prefix>/+/<name>.ext.
	BinaryFileExt        string   `protobuf:"bytes,7,opt,name=binary_file_ext,json=binaryFileExt,proto3" json:"binary_file_ext,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogStreamDescriptor) Reset()         { *m = LogStreamDescriptor{} }
func (m *LogStreamDescriptor) String() string { return proto.CompactTextString(m) }
func (*LogStreamDescriptor) ProtoMessage()    {}
func (*LogStreamDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_30887c96a468dac0, []int{0}
}

func (m *LogStreamDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogStreamDescriptor.Unmarshal(m, b)
}
func (m *LogStreamDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogStreamDescriptor.Marshal(b, m, deterministic)
}
func (m *LogStreamDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogStreamDescriptor.Merge(m, src)
}
func (m *LogStreamDescriptor) XXX_Size() int {
	return xxx_messageInfo_LogStreamDescriptor.Size(m)
}
func (m *LogStreamDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_LogStreamDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_LogStreamDescriptor proto.InternalMessageInfo

func (m *LogStreamDescriptor) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *LogStreamDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogStreamDescriptor) GetStreamType() StreamType {
	if m != nil {
		return m.StreamType
	}
	return StreamType_TEXT
}

func (m *LogStreamDescriptor) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *LogStreamDescriptor) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *LogStreamDescriptor) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *LogStreamDescriptor) GetBinaryFileExt() string {
	if m != nil {
		return m.BinaryFileExt
	}
	return ""
}

// Text stream content.
type Text struct {
	Lines                []*Text_Line `protobuf:"bytes,1,rep,name=lines,proto3" json:"lines,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Text) Reset()         { *m = Text{} }
func (m *Text) String() string { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()    {}
func (*Text) Descriptor() ([]byte, []int) {
	return fileDescriptor_30887c96a468dac0, []int{1}
}

func (m *Text) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Text.Unmarshal(m, b)
}
func (m *Text) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Text.Marshal(b, m, deterministic)
}
func (m *Text) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Text.Merge(m, src)
}
func (m *Text) XXX_Size() int {
	return xxx_messageInfo_Text.Size(m)
}
func (m *Text) XXX_DiscardUnknown() {
	xxx_messageInfo_Text.DiscardUnknown(m)
}

var xxx_messageInfo_Text proto.InternalMessageInfo

func (m *Text) GetLines() []*Text_Line {
	if m != nil {
		return m.Lines
	}
	return nil
}

// Contiguous text lines and their delimiters.
type Text_Line struct {
	// The line's text content, not including its delimiter.
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	//
	// The line's delimiter string.
	//
	// If this is an empty string, this line is continued in the next sequential
	// line, and the line's sequence number does not advance.
	Delimiter            string   `protobuf:"bytes,2,opt,name=delimiter,proto3" json:"delimiter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Text_Line) Reset()         { *m = Text_Line{} }
func (m *Text_Line) String() string { return proto.CompactTextString(m) }
func (*Text_Line) ProtoMessage()    {}
func (*Text_Line) Descriptor() ([]byte, []int) {
	return fileDescriptor_30887c96a468dac0, []int{1, 0}
}

func (m *Text_Line) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Text_Line.Unmarshal(m, b)
}
func (m *Text_Line) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Text_Line.Marshal(b, m, deterministic)
}
func (m *Text_Line) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Text_Line.Merge(m, src)
}
func (m *Text_Line) XXX_Size() int {
	return xxx_messageInfo_Text_Line.Size(m)
}
func (m *Text_Line) XXX_DiscardUnknown() {
	xxx_messageInfo_Text_Line.DiscardUnknown(m)
}

var xxx_messageInfo_Text_Line proto.InternalMessageInfo

func (m *Text_Line) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Text_Line) GetDelimiter() string {
	if m != nil {
		return m.Delimiter
	}
	return ""
}

// Binary stream content.
type Binary struct {
	// The binary stream's data.
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Binary) Reset()         { *m = Binary{} }
func (m *Binary) String() string { return proto.CompactTextString(m) }
func (*Binary) ProtoMessage()    {}
func (*Binary) Descriptor() ([]byte, []int) {
	return fileDescriptor_30887c96a468dac0, []int{2}
}

func (m *Binary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Binary.Unmarshal(m, b)
}
func (m *Binary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Binary.Marshal(b, m, deterministic)
}
func (m *Binary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Binary.Merge(m, src)
}
func (m *Binary) XXX_Size() int {
	return xxx_messageInfo_Binary.Size(m)
}
func (m *Binary) XXX_DiscardUnknown() {
	xxx_messageInfo_Binary.DiscardUnknown(m)
}

var xxx_messageInfo_Binary proto.InternalMessageInfo

func (m *Binary) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Datagram stream content type.
type Datagram struct {
	// This datagram data.
	Data                 []byte            `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Partial              *Datagram_Partial `protobuf:"bytes,2,opt,name=partial,proto3" json:"partial,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Datagram) Reset()         { *m = Datagram{} }
func (m *Datagram) String() string { return proto.CompactTextString(m) }
func (*Datagram) ProtoMessage()    {}
func (*Datagram) Descriptor() ([]byte, []int) {
	return fileDescriptor_30887c96a468dac0, []int{3}
}

func (m *Datagram) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Datagram.Unmarshal(m, b)
}
func (m *Datagram) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Datagram.Marshal(b, m, deterministic)
}
func (m *Datagram) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Datagram.Merge(m, src)
}
func (m *Datagram) XXX_Size() int {
	return xxx_messageInfo_Datagram.Size(m)
}
func (m *Datagram) XXX_DiscardUnknown() {
	xxx_messageInfo_Datagram.DiscardUnknown(m)
}

var xxx_messageInfo_Datagram proto.InternalMessageInfo

func (m *Datagram) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Datagram) GetPartial() *Datagram_Partial {
	if m != nil {
		return m.Partial
	}
	return nil
}

//
// If this is not a partial datagram, this field will include reassembly and
// state details for the full datagram.
type Datagram_Partial struct {
	//
	// The index, starting with zero, of this datagram fragment in the full
	// datagram.
	Index uint32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	// The size of the full datagram
	Size uint64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	// If true, this is the last partial datagram in the overall datagram.
	Last                 bool     `protobuf:"varint,3,opt,name=last,proto3" json:"last,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Datagram_Partial) Reset()         { *m = Datagram_Partial{} }
func (m *Datagram_Partial) String() string { return proto.CompactTextString(m) }
func (*Datagram_Partial) ProtoMessage()    {}
func (*Datagram_Partial) Descriptor() ([]byte, []int) {
	return fileDescriptor_30887c96a468dac0, []int{3, 0}
}

func (m *Datagram_Partial) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Datagram_Partial.Unmarshal(m, b)
}
func (m *Datagram_Partial) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Datagram_Partial.Marshal(b, m, deterministic)
}
func (m *Datagram_Partial) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Datagram_Partial.Merge(m, src)
}
func (m *Datagram_Partial) XXX_Size() int {
	return xxx_messageInfo_Datagram_Partial.Size(m)
}
func (m *Datagram_Partial) XXX_DiscardUnknown() {
	xxx_messageInfo_Datagram_Partial.DiscardUnknown(m)
}

var xxx_messageInfo_Datagram_Partial proto.InternalMessageInfo

func (m *Datagram_Partial) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Datagram_Partial) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Datagram_Partial) GetLast() bool {
	if m != nil {
		return m.Last
	}
	return false
}

// *
// An individual log entry.
//
// This contains the superset of transmissible log data. Its content fields
// should be interpreted in the context of the log stream's content type.
type LogEntry struct {
	//
	// The stream time offset for this content.
	//
	// This offset is added to the log stream's base "timestamp" to resolve the
	// timestamp for this specific Content.
	TimeOffset *duration.Duration `protobuf:"bytes,1,opt,name=time_offset,json=timeOffset,proto3" json:"time_offset,omitempty"`
	//
	// The message index within the Prefix (required).
	//
	// This is value is unique to this LogEntry across the entire set of entries
	// sharing the stream's Prefix. It is used to designate unambiguous log
	// ordering.
	PrefixIndex uint64 `protobuf:"varint,2,opt,name=prefix_index,json=prefixIndex,proto3" json:"prefix_index,omitempty"`
	//
	// The message index within its Stream (required).
	//
	// This value is unique across all entries sharing the same Prefix and Stream
	// Name. It is used to designate unambiguous log ordering within the stream.
	StreamIndex uint64 `protobuf:"varint,3,opt,name=stream_index,json=streamIndex,proto3" json:"stream_index,omitempty"`
	//
	// The sequence number of the first content entry in this LogEntry.
	//
	// Text: This is the line index of the first included line. Line indices begin
	//     at zero.
	// Binary: This is the byte offset of the first byte in the included data.
	// Datagram: This is the index of the datagram. The first datagram has index
	//     zero.
	Sequence uint64 `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
	//
	// The content of the message. The field that is populated here must
	// match the log's `stream_type`.
	//
	// Types that are valid to be assigned to Content:
	//	*LogEntry_Text
	//	*LogEntry_Binary
	//	*LogEntry_Datagram
	Content              isLogEntry_Content `protobuf_oneof:"content"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LogEntry) Reset()         { *m = LogEntry{} }
func (m *LogEntry) String() string { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()    {}
func (*LogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_30887c96a468dac0, []int{4}
}

func (m *LogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntry.Unmarshal(m, b)
}
func (m *LogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntry.Marshal(b, m, deterministic)
}
func (m *LogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntry.Merge(m, src)
}
func (m *LogEntry) XXX_Size() int {
	return xxx_messageInfo_LogEntry.Size(m)
}
func (m *LogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntry proto.InternalMessageInfo

func (m *LogEntry) GetTimeOffset() *duration.Duration {
	if m != nil {
		return m.TimeOffset
	}
	return nil
}

func (m *LogEntry) GetPrefixIndex() uint64 {
	if m != nil {
		return m.PrefixIndex
	}
	return 0
}

func (m *LogEntry) GetStreamIndex() uint64 {
	if m != nil {
		return m.StreamIndex
	}
	return 0
}

func (m *LogEntry) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

type isLogEntry_Content interface {
	isLogEntry_Content()
}

type LogEntry_Text struct {
	Text *Text `protobuf:"bytes,10,opt,name=text,proto3,oneof"`
}

type LogEntry_Binary struct {
	Binary *Binary `protobuf:"bytes,11,opt,name=binary,proto3,oneof"`
}

type LogEntry_Datagram struct {
	Datagram *Datagram `protobuf:"bytes,12,opt,name=datagram,proto3,oneof"`
}

func (*LogEntry_Text) isLogEntry_Content() {}

func (*LogEntry_Binary) isLogEntry_Content() {}

func (*LogEntry_Datagram) isLogEntry_Content() {}

func (m *LogEntry) GetContent() isLogEntry_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *LogEntry) GetText() *Text {
	if x, ok := m.GetContent().(*LogEntry_Text); ok {
		return x.Text
	}
	return nil
}

func (m *LogEntry) GetBinary() *Binary {
	if x, ok := m.GetContent().(*LogEntry_Binary); ok {
		return x.Binary
	}
	return nil
}

func (m *LogEntry) GetDatagram() *Datagram {
	if x, ok := m.GetContent().(*LogEntry_Datagram); ok {
		return x.Datagram
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LogEntry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LogEntry_OneofMarshaler, _LogEntry_OneofUnmarshaler, _LogEntry_OneofSizer, []interface{}{
		(*LogEntry_Text)(nil),
		(*LogEntry_Binary)(nil),
		(*LogEntry_Datagram)(nil),
	}
}

func _LogEntry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LogEntry)
	// content
	switch x := m.Content.(type) {
	case *LogEntry_Text:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Text); err != nil {
			return err
		}
	case *LogEntry_Binary:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Binary); err != nil {
			return err
		}
	case *LogEntry_Datagram:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Datagram); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LogEntry.Content has unexpected type %T", x)
	}
	return nil
}

func _LogEntry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LogEntry)
	switch tag {
	case 10: // content.text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Text)
		err := b.DecodeMessage(msg)
		m.Content = &LogEntry_Text{msg}
		return true, err
	case 11: // content.binary
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Binary)
		err := b.DecodeMessage(msg)
		m.Content = &LogEntry_Binary{msg}
		return true, err
	case 12: // content.datagram
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Datagram)
		err := b.DecodeMessage(msg)
		m.Content = &LogEntry_Datagram{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LogEntry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LogEntry)
	// content
	switch x := m.Content.(type) {
	case *LogEntry_Text:
		s := proto.Size(x.Text)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LogEntry_Binary:
		s := proto.Size(x.Binary)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LogEntry_Datagram:
		s := proto.Size(x.Datagram)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// *
// LogIndex is an index into an at-rest log storage.
//
// The log stream and log index are generated by the Archivist during archival.
//
// An archived log stream is a series of contiguous LogEntry frames. The index
// maps a log's logical logation in its stream, prefix, and timeline to its
// frame's binary offset in the archived log stream blob.
type LogIndex struct {
	//
	// The LogStreamDescriptor for this log stream (required).
	//
	// The index stores the stream's LogStreamDescriptor so that a client can
	// know the full set of log metadata by downloading its index.
	Desc *LogStreamDescriptor `protobuf:"bytes,1,opt,name=desc,proto3" json:"desc,omitempty"`
	//
	// A series of ascending-ordered Entry messages representing snapshots of an
	// archived log stream.
	//
	// Within this set of Entry messages, the "offset", "prefix_index",
	// "stream_index", and "time_offset" fields will be ascending.
	//
	// The frequency of Entry messages is not defined; it is up to the Archivist
	// process to choose a frequency.
	Entries []*LogIndex_Entry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
	// *
	// The last prefix index in the log stream.
	//
	// This is optional. If zero, there is either no information about the last
	// prefix index, or there are zero entries in the prefix.
	LastPrefixIndex uint64 `protobuf:"varint,3,opt,name=last_prefix_index,json=lastPrefixIndex,proto3" json:"last_prefix_index,omitempty"`
	// *
	// The last stream index in the log stream.
	//
	// This is optional. If zero, there is either no information about the last
	// stream index, or there are zero entries in the stream.
	LastStreamIndex uint64 `protobuf:"varint,4,opt,name=last_stream_index,json=lastStreamIndex,proto3" json:"last_stream_index,omitempty"`
	// *
	// The number of log entries in the stream.
	//
	// This is optional. If zero, there is either no information about the number
	// of log entries, or there are zero entries in the stream.
	LogEntryCount        uint64   `protobuf:"varint,5,opt,name=log_entry_count,json=logEntryCount,proto3" json:"log_entry_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogIndex) Reset()         { *m = LogIndex{} }
func (m *LogIndex) String() string { return proto.CompactTextString(m) }
func (*LogIndex) ProtoMessage()    {}
func (*LogIndex) Descriptor() ([]byte, []int) {
	return fileDescriptor_30887c96a468dac0, []int{5}
}

func (m *LogIndex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogIndex.Unmarshal(m, b)
}
func (m *LogIndex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogIndex.Marshal(b, m, deterministic)
}
func (m *LogIndex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogIndex.Merge(m, src)
}
func (m *LogIndex) XXX_Size() int {
	return xxx_messageInfo_LogIndex.Size(m)
}
func (m *LogIndex) XXX_DiscardUnknown() {
	xxx_messageInfo_LogIndex.DiscardUnknown(m)
}

var xxx_messageInfo_LogIndex proto.InternalMessageInfo

func (m *LogIndex) GetDesc() *LogStreamDescriptor {
	if m != nil {
		return m.Desc
	}
	return nil
}

func (m *LogIndex) GetEntries() []*LogIndex_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *LogIndex) GetLastPrefixIndex() uint64 {
	if m != nil {
		return m.LastPrefixIndex
	}
	return 0
}

func (m *LogIndex) GetLastStreamIndex() uint64 {
	if m != nil {
		return m.LastStreamIndex
	}
	return 0
}

func (m *LogIndex) GetLogEntryCount() uint64 {
	if m != nil {
		return m.LogEntryCount
	}
	return 0
}

//
// Entry is a single index entry.
//
// The index is composed of a series of entries, each corresponding to a
// sequential snapshot of of the log stream.
type LogIndex_Entry struct {
	//
	// The byte offset in the emitted log stream of the RecordIO entry for the
	// LogEntry corresponding to this Entry.
	Offset uint64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	//
	// The sequence number of the first content entry.
	//
	// Text: This is the line index of the first included line. Line indices
	//     begin at zero.
	// Binary: This is the byte offset of the first byte in the included data.
	// Datagram: This is the index of the datagram. The first datagram has index
	//     zero.
	Sequence uint64 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	//
	// The log index that this entry describes (required).
	//
	// This is used by clients to identify a specific LogEntry within a set of
	// streams sharing a Prefix.
	PrefixIndex uint64 `protobuf:"varint,3,opt,name=prefix_index,json=prefixIndex,proto3" json:"prefix_index,omitempty"`
	//
	// The time offset of this log entry (required).
	//
	// This is used by clients to identify a specific LogEntry within a log
	// stream.
	StreamIndex uint64 `protobuf:"varint,4,opt,name=stream_index,json=streamIndex,proto3" json:"stream_index,omitempty"`
	//
	// The time offset of this log entry, in microseconds.
	//
	// This is added to the descriptor's "timestamp" field to identify the
	// specific timestamp of this log. It is used by clients to identify a
	// specific LogEntry by time.
	TimeOffset           *duration.Duration `protobuf:"bytes,5,opt,name=time_offset,json=timeOffset,proto3" json:"time_offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LogIndex_Entry) Reset()         { *m = LogIndex_Entry{} }
func (m *LogIndex_Entry) String() string { return proto.CompactTextString(m) }
func (*LogIndex_Entry) ProtoMessage()    {}
func (*LogIndex_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_30887c96a468dac0, []int{5, 0}
}

func (m *LogIndex_Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogIndex_Entry.Unmarshal(m, b)
}
func (m *LogIndex_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogIndex_Entry.Marshal(b, m, deterministic)
}
func (m *LogIndex_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogIndex_Entry.Merge(m, src)
}
func (m *LogIndex_Entry) XXX_Size() int {
	return xxx_messageInfo_LogIndex_Entry.Size(m)
}
func (m *LogIndex_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_LogIndex_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_LogIndex_Entry proto.InternalMessageInfo

func (m *LogIndex_Entry) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *LogIndex_Entry) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *LogIndex_Entry) GetPrefixIndex() uint64 {
	if m != nil {
		return m.PrefixIndex
	}
	return 0
}

func (m *LogIndex_Entry) GetStreamIndex() uint64 {
	if m != nil {
		return m.StreamIndex
	}
	return 0
}

func (m *LogIndex_Entry) GetTimeOffset() *duration.Duration {
	if m != nil {
		return m.TimeOffset
	}
	return nil
}

func init() {
	proto.RegisterEnum("logpb.StreamType", StreamType_name, StreamType_value)
	proto.RegisterType((*LogStreamDescriptor)(nil), "logpb.LogStreamDescriptor")
	proto.RegisterMapType((map[string]string)(nil), "logpb.LogStreamDescriptor.TagsEntry")
	proto.RegisterType((*Text)(nil), "logpb.Text")
	proto.RegisterType((*Text_Line)(nil), "logpb.Text.Line")
	proto.RegisterType((*Binary)(nil), "logpb.Binary")
	proto.RegisterType((*Datagram)(nil), "logpb.Datagram")
	proto.RegisterType((*Datagram_Partial)(nil), "logpb.Datagram.Partial")
	proto.RegisterType((*LogEntry)(nil), "logpb.LogEntry")
	proto.RegisterType((*LogIndex)(nil), "logpb.LogIndex")
	proto.RegisterType((*LogIndex_Entry)(nil), "logpb.LogIndex.Entry")
}

func init() {
	proto.RegisterFile("github.com/TriggerMail/luci-go/logdog/api/logpb/log.proto", fileDescriptor_30887c96a468dac0)
}

var fileDescriptor_30887c96a468dac0 = []byte{
	// 785 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xae, 0x1d, 0x27, 0x71, 0x8e, 0x53, 0x9a, 0x1d, 0x60, 0x31, 0x16, 0x82, 0x34, 0x42, 0x25,
	0x5a, 0x09, 0x07, 0xc2, 0x05, 0x55, 0xef, 0x52, 0x52, 0xb6, 0x8b, 0x0a, 0xac, 0x66, 0x73, 0x01,
	0x57, 0xd6, 0x24, 0x99, 0x98, 0x01, 0xdb, 0x63, 0xec, 0x09, 0x4a, 0x78, 0x14, 0x5e, 0x01, 0x89,
	0x37, 0xe0, 0x91, 0x78, 0x07, 0x34, 0x67, 0x26, 0x3f, 0x6d, 0x17, 0xa4, 0xbd, 0x89, 0xce, 0x9c,
	0xf3, 0x9d, 0x33, 0x9f, 0xbf, 0xf3, 0x4d, 0x60, 0x94, 0xca, 0x78, 0xf1, 0x53, 0x25, 0x73, 0xb1,
	0xce, 0x63, 0x59, 0xa5, 0xa3, 0x6c, 0xbd, 0x10, 0xa3, 0x4c, 0xa6, 0x4b, 0x99, 0x8e, 0x58, 0x89,
	0x61, 0x39, 0xd7, 0xbf, 0x71, 0x59, 0x49, 0x25, 0x49, 0x13, 0x13, 0xd1, 0x47, 0xa9, 0x94, 0x69,
	0xc6, 0x47, 0x98, 0x9c, 0xaf, 0x57, 0x23, 0x25, 0x72, 0x5e, 0x2b, 0x96, 0x97, 0x06, 0x17, 0x7d,
	0xf8, 0x10, 0xb0, 0x5c, 0x57, 0x4c, 0x09, 0x59, 0x98, 0xfa, 0xe0, 0x1f, 0x17, 0xde, 0xbe, 0x93,
	0xe9, 0x2b, 0x55, 0x71, 0x96, 0x4f, 0x79, 0xbd, 0xa8, 0x44, 0xa9, 0x64, 0x45, 0x9e, 0x42, 0xab,
	0xac, 0xf8, 0x4a, 0x6c, 0x42, 0xa7, 0xef, 0x0c, 0x3b, 0xd4, 0x9e, 0x08, 0x01, 0xaf, 0x60, 0x39,
	0x0f, 0x5d, 0xcc, 0x62, 0x4c, 0xc6, 0x10, 0xd4, 0xd8, 0x9f, 0xa8, 0x6d, 0xc9, 0xc3, 0x46, 0xdf,
	0x19, 0xbe, 0x35, 0x7e, 0x12, 0x23, 0xc3, 0xd8, 0x4c, 0x9e, 0x6d, 0x4b, 0x4e, 0xa1, 0xde, 0xc7,
	0xe4, 0x1c, 0xba, 0x0b, 0x59, 0x28, 0x5e, 0x28, 0xd3, 0xe4, 0xe1, 0xbc, 0xc0, 0xe6, 0x10, 0x72,
	0x09, 0x9d, 0xfd, 0xd7, 0x84, 0xcd, 0xbe, 0x33, 0x0c, 0xc6, 0x51, 0x6c, 0x3e, 0x27, 0xde, 0x7d,
	0x4e, 0x3c, 0xdb, 0x21, 0xe8, 0x01, 0x4c, 0x2e, 0xc1, 0x53, 0x2c, 0xad, 0xc3, 0x56, 0xbf, 0x31,
	0x0c, 0xc6, 0x1f, 0x5b, 0x26, 0xaf, 0xf9, 0xcc, 0x78, 0xc6, 0xd2, 0xfa, 0xa6, 0x50, 0xd5, 0x96,
	0x62, 0x07, 0xb9, 0x80, 0xb3, 0xb9, 0x28, 0x58, 0xb5, 0x4d, 0x56, 0x22, 0xe3, 0x09, 0xdf, 0xa8,
	0xb0, 0x8d, 0xcc, 0x4e, 0x4d, 0xfa, 0x6b, 0x91, 0xf1, 0x9b, 0x8d, 0x8a, 0xbe, 0x84, 0xce, 0xbe,
	0x95, 0xf4, 0xa0, 0xf1, 0x0b, 0xdf, 0x5a, 0xa1, 0x74, 0x48, 0xde, 0x81, 0xe6, 0x6f, 0x2c, 0x5b,
	0xef, 0x64, 0x32, 0x87, 0x2b, 0xf7, 0xd2, 0x19, 0xfc, 0x0c, 0xde, 0x8c, 0x6f, 0x14, 0xb9, 0x80,
	0x66, 0x26, 0x0a, 0x5e, 0x87, 0x0e, 0x72, 0xec, 0x59, 0x8e, 0xba, 0x16, 0xdf, 0x89, 0x82, 0x53,
	0x53, 0x8e, 0xae, 0xc0, 0xd3, 0xc7, 0xc3, 0x44, 0x7d, 0x4b, 0xd7, 0x4e, 0x24, 0x1f, 0x40, 0x67,
	0xc9, 0x33, 0x91, 0x0b, 0xc5, 0x2b, 0x7b, 0xd7, 0x21, 0x31, 0x18, 0x40, 0xeb, 0x1a, 0x59, 0xeb,
	0xad, 0x2d, 0x99, 0x62, 0x08, 0xe9, 0x52, 0x8c, 0xbf, 0xf1, 0x7c, 0xa7, 0xe7, 0x0e, 0xfe, 0x70,
	0xc0, 0x9f, 0x32, 0xc5, 0xd2, 0x8a, 0xe5, 0x7b, 0x98, 0x73, 0x80, 0x91, 0xcf, 0xa1, 0x5d, 0xb2,
	0x4a, 0x09, 0x96, 0x61, 0x77, 0x30, 0x7e, 0xcf, 0x52, 0xdd, 0x75, 0xc5, 0x2f, 0x4d, 0x99, 0xee,
	0x70, 0xd1, 0x73, 0x68, 0xdb, 0x9c, 0xa6, 0x2d, 0x8a, 0x25, 0x37, 0x2e, 0x3a, 0xa5, 0xe6, 0xa0,
	0xef, 0xa9, 0xc5, 0xef, 0x46, 0x1d, 0x8f, 0x62, 0xac, 0x73, 0x19, 0xab, 0x15, 0xba, 0xc7, 0xa7,
	0x18, 0x0f, 0xfe, 0x72, 0xc1, 0xbf, 0x93, 0xa9, 0x51, 0xf9, 0x0a, 0x02, 0xbd, 0xe1, 0x44, 0xae,
	0x56, 0x35, 0x57, 0x38, 0x30, 0x18, 0xbf, 0xff, 0xc8, 0x10, 0x53, 0xeb, 0x6f, 0x0a, 0x1a, 0xfd,
	0x3d, 0x82, 0xb5, 0xdb, 0x8c, 0x7f, 0x13, 0xc3, 0xc6, 0x5c, 0x1c, 0x98, 0xdc, 0x0b, 0xe4, 0x74,
	0x0e, 0x5d, 0x6b, 0x62, 0x03, 0x69, 0x18, 0x88, 0xc9, 0x19, 0x48, 0x04, 0x7e, 0xcd, 0x7f, 0x5d,
	0xf3, 0x62, 0x61, 0xfc, 0xea, 0xd1, 0xfd, 0x99, 0x9c, 0x83, 0xa7, 0xb4, 0x5b, 0x00, 0x69, 0x05,
	0x47, 0xeb, 0xbc, 0x3d, 0xa1, 0x58, 0x22, 0x9f, 0x40, 0xcb, 0x98, 0x28, 0x0c, 0x10, 0x74, 0x6a,
	0x41, 0x66, 0x47, 0xb7, 0x27, 0xd4, 0x96, 0xc9, 0xa7, 0xe0, 0x2f, 0xad, 0xb8, 0x61, 0x17, 0xa1,
	0x67, 0x0f, 0x34, 0xbf, 0x3d, 0xa1, 0x7b, 0xc8, 0x75, 0x07, 0xda, 0xf6, 0xd9, 0x0c, 0xfe, 0x6c,
	0xa0, 0x60, 0x86, 0x6e, 0x0c, 0xde, 0x92, 0xd7, 0x0b, 0xab, 0x54, 0xf4, 0xdf, 0xaf, 0x80, 0x22,
	0x8e, 0x8c, 0xa0, 0xcd, 0x0b, 0x55, 0x09, 0x5e, 0x87, 0x2e, 0x9a, 0xf2, 0xdd, 0x43, 0x0b, 0x4e,
	0x8c, 0xcd, 0x4b, 0xd9, 0xa1, 0xc8, 0x33, 0x78, 0xa2, 0xd7, 0x94, 0xdc, 0x93, 0xd6, 0xe8, 0x76,
	0xa6, 0x0b, 0x2f, 0x8f, 0xe4, 0xdd, 0x61, 0xef, 0x69, 0xec, 0x1d, 0xb0, 0xaf, 0x8e, 0x74, 0xbe,
	0x80, 0xb3, 0x4c, 0xa6, 0x89, 0xbe, 0x66, 0x9b, 0x2c, 0xe4, 0xba, 0x50, 0xf8, 0xfc, 0x3d, 0x7a,
	0x9a, 0x59, 0x33, 0x7c, 0xa5, 0x93, 0xd1, 0xdf, 0x0e, 0x34, 0x8d, 0x37, 0x9e, 0x42, 0xeb, 0xc8,
	0x16, 0x1e, 0xb5, 0xa7, 0x7b, 0x1b, 0x73, 0x1f, 0x6d, 0xac, 0xfb, 0x1a, 0xe2, 0xff, 0xeb, 0x09,
	0xef, 0xb1, 0x27, 0x1e, 0xb8, 0xb2, 0xf9, 0x06, 0xae, 0x7c, 0xf6, 0x19, 0xc0, 0xe1, 0xdf, 0x91,
	0xf8, 0xe0, 0xcd, 0x6e, 0x7e, 0x98, 0xf5, 0x4e, 0x08, 0x40, 0xeb, 0xfa, 0xc5, 0x77, 0x13, 0xfa,
	0x63, 0xcf, 0x21, 0x5d, 0xf0, 0xa7, 0x93, 0xd9, 0xe4, 0x39, 0x9d, 0x7c, 0xdb, 0x73, 0xe7, 0x2d,
	0x1c, 0xf8, 0xc5, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xc2, 0x12, 0xd9, 0x2f, 0x06, 0x00,
	0x00,
}
