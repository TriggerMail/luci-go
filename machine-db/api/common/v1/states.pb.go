// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/common/v1/states.proto

package common

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

// State is an enumeration of possible states a resource may be in.
// When adding a state here, choose a name that results in no common prefixes
// between any two states, then update prefix-matching in prefixes.go.
type State int32

const (
	// Resource state unspecified.
	State_STATE_UNSPECIFIED State = 0
	// Resource is not allocated.
	State_FREE State = 1
	// Resource is allocated for future use.
	State_PRERELEASE State = 2
	// Resource is allocated and currently used in production.
	State_SERVING State = 3
	// Resource is allocated and currently used for testing.
	State_TEST State = 4
	// Resource is undergoing repairs.
	State_REPAIR State = 5
	// Resource is allocated but unused.
	State_DECOMMISSIONED State = 6
)

var State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "FREE",
	2: "PRERELEASE",
	3: "SERVING",
	4: "TEST",
	5: "REPAIR",
	6: "DECOMMISSIONED",
}

var State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"FREE":              1,
	"PRERELEASE":        2,
	"SERVING":           3,
	"TEST":              4,
	"REPAIR":            5,
	"DECOMMISSIONED":    6,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f170d7ecb0c39b98, []int{0}
}

func init() {
	proto.RegisterEnum("common.State", State_name, State_value)
}

func init() {
	proto.RegisterFile("github.com/TriggerMail/luci-go/machine-db/api/common/v1/states.proto", fileDescriptor_f170d7ecb0c39b98)
}

var fileDescriptor_f170d7ecb0c39b98 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x24, 0xcd, 0x5d, 0x4b, 0xc3, 0x30,
	0x14, 0x80, 0x61, 0x3f, 0xb6, 0x28, 0x47, 0x18, 0xf1, 0x80, 0x7f, 0x42, 0xb0, 0x41, 0xbc, 0xf3,
	0xae, 0xac, 0x67, 0x12, 0x70, 0x5d, 0xc9, 0x89, 0xde, 0x4a, 0x17, 0x47, 0x1b, 0x30, 0x3d, 0xa5,
	0x1f, 0xfe, 0x7e, 0xa9, 0xbb, 0x7d, 0x78, 0xe1, 0x85, 0xd7, 0x46, 0xb2, 0xd0, 0x0e, 0x92, 0xe2,
	0x9c, 0x32, 0x19, 0x1a, 0xf3, 0x33, 0x87, 0x68, 0x52, 0x1d, 0xda, 0xd8, 0x9d, 0x9e, 0xbe, 0x8f,
	0xa6, 0xee, 0xa3, 0x09, 0x92, 0x92, 0x74, 0xe6, 0xf7, 0xd9, 0x8c, 0x53, 0x3d, 0x9d, 0xc6, 0xac,
	0x1f, 0x64, 0x12, 0x54, 0x67, 0x7f, 0x14, 0x58, 0xf3, 0xe2, 0xf8, 0x00, 0xf7, 0xec, 0x73, 0x4f,
	0x5f, 0x1f, 0x25, 0x57, 0xb4, 0xb5, 0x3b, 0x4b, 0x85, 0xbe, 0xc0, 0x5b, 0x58, 0xed, 0x1c, 0x91,
	0xbe, 0xc4, 0x0d, 0x40, 0xe5, 0xc8, 0xd1, 0x3b, 0xe5, 0x4c, 0xfa, 0x0a, 0xef, 0xe0, 0x86, 0xc9,
	0x7d, 0xda, 0xf2, 0x4d, 0x5f, 0x2f, 0x99, 0x27, 0xf6, 0x7a, 0x85, 0x00, 0xca, 0x51, 0x95, 0x5b,
	0xa7, 0xd7, 0x88, 0xb0, 0x29, 0x68, 0x7b, 0xd8, 0xef, 0x2d, 0xb3, 0x3d, 0x94, 0x54, 0x68, 0x75,
	0x54, 0xff, 0xff, 0x97, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x56, 0x9a, 0xab, 0xbd, 0x00,
	0x00, 0x00,
}
