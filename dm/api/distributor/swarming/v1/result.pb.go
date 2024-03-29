// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/dm/api/distributor/swarming/v1/result.proto

package swarmingV1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// This is the swarming-specific result for Executions run via swarming.
type Result struct {
	ExitCode int64 `protobuf:"varint,1,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	// The isolated hash of the output directory
	IsolatedOutdir *IsolatedRef `protobuf:"bytes,2,opt,name=isolated_outdir,json=isolatedOutdir,proto3" json:"isolated_outdir,omitempty"`
	// The pinned cipd packages that this task actually used.
	CipdPins *CipdSpec `protobuf:"bytes,3,opt,name=cipd_pins,json=cipdPins,proto3" json:"cipd_pins,omitempty"`
	// The captured snapshot dimensions that the bot actually had.
	SnapshotDimensions   map[string]string `protobuf:"bytes,4,rep,name=snapshot_dimensions,json=snapshotDimensions,proto3" json:"snapshot_dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_f78100aa064bdfa0, []int{0}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetExitCode() int64 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *Result) GetIsolatedOutdir() *IsolatedRef {
	if m != nil {
		return m.IsolatedOutdir
	}
	return nil
}

func (m *Result) GetCipdPins() *CipdSpec {
	if m != nil {
		return m.CipdPins
	}
	return nil
}

func (m *Result) GetSnapshotDimensions() map[string]string {
	if m != nil {
		return m.SnapshotDimensions
	}
	return nil
}

func init() {
	proto.RegisterType((*Result)(nil), "swarmingV1.Result")
	proto.RegisterMapType((map[string]string)(nil), "swarmingV1.Result.SnapshotDimensionsEntry")
}

func init() {
	proto.RegisterFile("github.com/TriggerMail/luci-go/dm/api/distributor/swarming/v1/result.proto", fileDescriptor_f78100aa064bdfa0)
}

var fileDescriptor_f78100aa064bdfa0 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xe9, 0xaa, 0x63, 0xcd, 0x40, 0x25, 0x0e, 0x56, 0xe6, 0x65, 0x78, 0x1a, 0x1e, 0x1a,
	0x36, 0x2f, 0xe2, 0x41, 0x06, 0x73, 0x07, 0x4f, 0x4a, 0x06, 0x5e, 0x3c, 0x94, 0xae, 0xc9, 0xb6,
	0x87, 0x6d, 0x5e, 0x48, 0xd2, 0xe9, 0xfe, 0x39, 0xff, 0x36, 0xe9, 0x8f, 0xe9, 0x40, 0xbc, 0xec,
	0xd6, 0xe6, 0x7d, 0x3e, 0x5f, 0xde, 0xf7, 0x91, 0xe9, 0x1a, 0xa3, 0x74, 0x63, 0x30, 0x87, 0x22,
	0x8f, 0xd0, 0xac, 0x59, 0x56, 0xa4, 0xc0, 0x44, 0xce, 0x12, 0x0d, 0x4c, 0x80, 0x75, 0x06, 0x96,
	0x85, 0x43, 0xc3, 0xec, 0x47, 0x62, 0x72, 0x50, 0x6b, 0xb6, 0x1d, 0x33, 0x23, 0x6d, 0x91, 0xb9,
	0x48, 0x1b, 0x74, 0x48, 0xc9, 0x7e, 0xf2, 0x3a, 0x1e, 0x3c, 0x1c, 0x93, 0x96, 0x82, 0x16, 0x75,
	0xd6, 0x60, 0x7e, 0x8c, 0x0f, 0x16, 0xb3, 0xc4, 0xc9, 0xd8, 0xc8, 0x55, 0x1d, 0x73, 0xfd, 0xd5,
	0x22, 0x6d, 0x5e, 0xed, 0x48, 0xaf, 0x48, 0x20, 0x3f, 0xc1, 0xc5, 0x29, 0x0a, 0x19, 0x7a, 0x43,
	0x6f, 0xe4, 0xf3, 0x4e, 0xf9, 0x30, 0x43, 0x21, 0xe9, 0x94, 0x9c, 0x37, 0xb2, 0x88, 0xb1, 0x70,
	0x02, 0x4c, 0xd8, 0x1a, 0x7a, 0xa3, 0xee, 0xa4, 0x1f, 0xfd, 0x96, 0x8a, 0x9e, 0x1a, 0x84, 0xcb,
	0x15, 0x3f, 0xdb, 0xf3, 0xcf, 0x15, 0x4e, 0xc7, 0x24, 0x28, 0xd7, 0x8f, 0x35, 0x28, 0x1b, 0xfa,
	0x95, 0xdb, 0x3b, 0x74, 0x67, 0xa0, 0xc5, 0x42, 0xcb, 0x94, 0x77, 0x4a, 0xec, 0x05, 0x94, 0xa5,
	0x6f, 0xe4, 0xd2, 0xaa, 0x44, 0xdb, 0x0d, 0xba, 0x58, 0x40, 0x2e, 0x95, 0x05, 0x54, 0x36, 0x3c,
	0x19, 0xfa, 0xa3, 0xee, 0xe4, 0xe6, 0x50, 0xae, 0x2b, 0x44, 0x8b, 0x86, 0x7e, 0xfc, 0x81, 0xe7,
	0xca, 0x99, 0x1d, 0xa7, 0xf6, 0xcf, 0x60, 0x30, 0x27, 0xfd, 0x7f, 0x70, 0x7a, 0x41, 0xfc, 0x77,
	0xb9, 0xab, 0x6e, 0x10, 0xf0, 0xf2, 0x93, 0xf6, 0xc8, 0xe9, 0x36, 0xc9, 0x0a, 0x59, 0x95, 0x0e,
	0x78, 0xfd, 0x73, 0xdf, 0xba, 0xf3, 0x96, 0xed, 0xea, 0x8e, 0xb7, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x09, 0x97, 0x83, 0xdb, 0x1e, 0x02, 0x00, 0x00,
}
