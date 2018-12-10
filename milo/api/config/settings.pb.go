// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/milo/api/config/settings.proto

package config

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

// Settings represents the format for the global (service) config for Milo.
type Settings struct {
	Buildbot    *Settings_Buildbot    `protobuf:"bytes,1,opt,name=buildbot,proto3" json:"buildbot,omitempty"`
	Buildbucket *Settings_Buildbucket `protobuf:"bytes,2,opt,name=buildbucket,proto3" json:"buildbucket,omitempty"`
	Swarming    *Settings_Swarming    `protobuf:"bytes,3,opt,name=swarming,proto3" json:"swarming,omitempty"`
	// source_acls instructs Milo to provide Git/Gerrit data
	// (e.g., blamelist) to some of its users on entire subdomains or individual
	// repositories (Gerrit "projects").
	//
	// Multiple records are allowed, but each host and project must appear only in
	// one record.
	SourceAcls           []*Settings_SourceAcls `protobuf:"bytes,4,rep,name=source_acls,json=sourceAcls,proto3" json:"source_acls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Settings) Reset()         { *m = Settings{} }
func (m *Settings) String() string { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()    {}
func (*Settings) Descriptor() ([]byte, []int) {
	return fileDescriptor_98dd5cb9562385c0, []int{0}
}

func (m *Settings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings.Unmarshal(m, b)
}
func (m *Settings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings.Marshal(b, m, deterministic)
}
func (m *Settings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings.Merge(m, src)
}
func (m *Settings) XXX_Size() int {
	return xxx_messageInfo_Settings.Size(m)
}
func (m *Settings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings proto.InternalMessageInfo

func (m *Settings) GetBuildbot() *Settings_Buildbot {
	if m != nil {
		return m.Buildbot
	}
	return nil
}

func (m *Settings) GetBuildbucket() *Settings_Buildbucket {
	if m != nil {
		return m.Buildbucket
	}
	return nil
}

func (m *Settings) GetSwarming() *Settings_Swarming {
	if m != nil {
		return m.Swarming
	}
	return nil
}

func (m *Settings) GetSourceAcls() []*Settings_SourceAcls {
	if m != nil {
		return m.SourceAcls
	}
	return nil
}

type Settings_Buildbot struct {
	// internal_reader is the infra-auth group that is allowed to read internal
	// buildbot data.
	InternalReader string `protobuf:"bytes,1,opt,name=internal_reader,json=internalReader,proto3" json:"internal_reader,omitempty"`
	// public_subscription is the name of the pubsub topic where public builds come in
	// from
	PublicSubscription string `protobuf:"bytes,2,opt,name=public_subscription,json=publicSubscription,proto3" json:"public_subscription,omitempty"`
	// internal_subscription is the name of the pubsub topic where internal builds
	// come in from
	InternalSubscription string   `protobuf:"bytes,3,opt,name=internal_subscription,json=internalSubscription,proto3" json:"internal_subscription,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings_Buildbot) Reset()         { *m = Settings_Buildbot{} }
func (m *Settings_Buildbot) String() string { return proto.CompactTextString(m) }
func (*Settings_Buildbot) ProtoMessage()    {}
func (*Settings_Buildbot) Descriptor() ([]byte, []int) {
	return fileDescriptor_98dd5cb9562385c0, []int{0, 0}
}

func (m *Settings_Buildbot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings_Buildbot.Unmarshal(m, b)
}
func (m *Settings_Buildbot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings_Buildbot.Marshal(b, m, deterministic)
}
func (m *Settings_Buildbot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings_Buildbot.Merge(m, src)
}
func (m *Settings_Buildbot) XXX_Size() int {
	return xxx_messageInfo_Settings_Buildbot.Size(m)
}
func (m *Settings_Buildbot) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings_Buildbot.DiscardUnknown(m)
}

var xxx_messageInfo_Settings_Buildbot proto.InternalMessageInfo

func (m *Settings_Buildbot) GetInternalReader() string {
	if m != nil {
		return m.InternalReader
	}
	return ""
}

func (m *Settings_Buildbot) GetPublicSubscription() string {
	if m != nil {
		return m.PublicSubscription
	}
	return ""
}

func (m *Settings_Buildbot) GetInternalSubscription() string {
	if m != nil {
		return m.InternalSubscription
	}
	return ""
}

type Settings_Buildbucket struct {
	// name is the user friendly name of the Buildbucket instance we're pointing to.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// host is the hostname of the buildbucket instance we're pointing to (sans schema).
	Host string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	// project is the name of the Google Cloud project that the pubsub topic
	// belongs to.
	Project              string   `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings_Buildbucket) Reset()         { *m = Settings_Buildbucket{} }
func (m *Settings_Buildbucket) String() string { return proto.CompactTextString(m) }
func (*Settings_Buildbucket) ProtoMessage()    {}
func (*Settings_Buildbucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_98dd5cb9562385c0, []int{0, 1}
}

func (m *Settings_Buildbucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings_Buildbucket.Unmarshal(m, b)
}
func (m *Settings_Buildbucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings_Buildbucket.Marshal(b, m, deterministic)
}
func (m *Settings_Buildbucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings_Buildbucket.Merge(m, src)
}
func (m *Settings_Buildbucket) XXX_Size() int {
	return xxx_messageInfo_Settings_Buildbucket.Size(m)
}
func (m *Settings_Buildbucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings_Buildbucket.DiscardUnknown(m)
}

var xxx_messageInfo_Settings_Buildbucket proto.InternalMessageInfo

func (m *Settings_Buildbucket) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Settings_Buildbucket) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Settings_Buildbucket) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

type Settings_Swarming struct {
	// default_host is the hostname of the swarming host Milo defaults to, if
	// none is specified.  Default host is implicitly an allowed host.
	DefaultHost string `protobuf:"bytes,1,opt,name=default_host,json=defaultHost,proto3" json:"default_host,omitempty"`
	// allowed_hosts is a whitelist of hostnames of swarming instances
	// that Milo is allowed to talk to.  This is specified here for security
	// reasons, because Milo will hand out its oauth2 token to a swarming host.
	AllowedHosts         []string `protobuf:"bytes,2,rep,name=allowed_hosts,json=allowedHosts,proto3" json:"allowed_hosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings_Swarming) Reset()         { *m = Settings_Swarming{} }
func (m *Settings_Swarming) String() string { return proto.CompactTextString(m) }
func (*Settings_Swarming) ProtoMessage()    {}
func (*Settings_Swarming) Descriptor() ([]byte, []int) {
	return fileDescriptor_98dd5cb9562385c0, []int{0, 2}
}

func (m *Settings_Swarming) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings_Swarming.Unmarshal(m, b)
}
func (m *Settings_Swarming) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings_Swarming.Marshal(b, m, deterministic)
}
func (m *Settings_Swarming) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings_Swarming.Merge(m, src)
}
func (m *Settings_Swarming) XXX_Size() int {
	return xxx_messageInfo_Settings_Swarming.Size(m)
}
func (m *Settings_Swarming) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings_Swarming.DiscardUnknown(m)
}

var xxx_messageInfo_Settings_Swarming proto.InternalMessageInfo

func (m *Settings_Swarming) GetDefaultHost() string {
	if m != nil {
		return m.DefaultHost
	}
	return ""
}

func (m *Settings_Swarming) GetAllowedHosts() []string {
	if m != nil {
		return m.AllowedHosts
	}
	return nil
}

// SourceAcls grants read access on a set of Git/Gerrit hosts or projects.
type Settings_SourceAcls struct {
	// host grants read access on all project at this host.
	//
	// For more granularity, use the project field instead.
	//
	// For *.googlesource.com domains, host should not be a Gerrit host,
	// i.e.  it shouldn't be <subdomain>-review.googlesource.com.
	Hosts []string `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// project is a URL to a Git repository.
	//
	// Read access is granted on both git data and Gerrit CLs of this project.
	//
	// For *.googlesource.com Git repositories:
	//   URL Path should not start with '/a/' (forced authentication).
	//   URL Path should not end with '.git' (redundant).
	Projects []string `protobuf:"bytes,2,rep,name=projects,proto3" json:"projects,omitempty"`
	// readers are allowed to read git/gerrit data from targets.
	//
	// Three types of identity strings are supported:
	//  * Emails.                   For example: "someuser@example.com"
	//  * Chrome-infra-auth Groups. For example: "group:committers"
	//  * Auth service identities.  For example: "kind:name"
	//
	// Required.
	Readers              []string `protobuf:"bytes,3,rep,name=readers,proto3" json:"readers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Settings_SourceAcls) Reset()         { *m = Settings_SourceAcls{} }
func (m *Settings_SourceAcls) String() string { return proto.CompactTextString(m) }
func (*Settings_SourceAcls) ProtoMessage()    {}
func (*Settings_SourceAcls) Descriptor() ([]byte, []int) {
	return fileDescriptor_98dd5cb9562385c0, []int{0, 3}
}

func (m *Settings_SourceAcls) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings_SourceAcls.Unmarshal(m, b)
}
func (m *Settings_SourceAcls) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings_SourceAcls.Marshal(b, m, deterministic)
}
func (m *Settings_SourceAcls) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings_SourceAcls.Merge(m, src)
}
func (m *Settings_SourceAcls) XXX_Size() int {
	return xxx_messageInfo_Settings_SourceAcls.Size(m)
}
func (m *Settings_SourceAcls) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings_SourceAcls.DiscardUnknown(m)
}

var xxx_messageInfo_Settings_SourceAcls proto.InternalMessageInfo

func (m *Settings_SourceAcls) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *Settings_SourceAcls) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

func (m *Settings_SourceAcls) GetReaders() []string {
	if m != nil {
		return m.Readers
	}
	return nil
}

func init() {
	proto.RegisterType((*Settings)(nil), "milo.Settings")
	proto.RegisterType((*Settings_Buildbot)(nil), "milo.Settings.Buildbot")
	proto.RegisterType((*Settings_Buildbucket)(nil), "milo.Settings.Buildbucket")
	proto.RegisterType((*Settings_Swarming)(nil), "milo.Settings.Swarming")
	proto.RegisterType((*Settings_SourceAcls)(nil), "milo.Settings.SourceAcls")
}

func init() {
	proto.RegisterFile("github.com/TriggerMail/luci-go/milo/api/config/settings.proto", fileDescriptor_98dd5cb9562385c0)
}

var fileDescriptor_98dd5cb9562385c0 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xbf, 0x8e, 0xd4, 0x30,
	0x10, 0xc6, 0x95, 0xcb, 0xde, 0x91, 0x9d, 0x1c, 0x20, 0x99, 0x43, 0x84, 0x54, 0x0b, 0x14, 0x6c,
	0x95, 0x48, 0x97, 0x0e, 0xd1, 0xb0, 0x15, 0x1d, 0x92, 0xb7, 0x41, 0x34, 0x91, 0xe3, 0x78, 0xb3,
	0x06, 0x27, 0x8e, 0xfc, 0x47, 0xfb, 0x0c, 0x3c, 0x02, 0x6f, 0x8b, 0xe2, 0xd8, 0x21, 0x9c, 0xb6,
	0x9b, 0x99, 0xef, 0xf7, 0x7d, 0xb6, 0x47, 0x86, 0xaa, 0x93, 0x05, 0x3d, 0x2b, 0xd9, 0x73, 0xdb,
	0x17, 0x52, 0x75, 0xa5, 0xb0, 0x94, 0x97, 0x3d, 0x17, 0xb2, 0x24, 0x23, 0x2f, 0xa9, 0x1c, 0x4e,
	0xbc, 0x2b, 0x35, 0x33, 0x86, 0x0f, 0x9d, 0x2e, 0x46, 0x25, 0x8d, 0x44, 0x9b, 0x49, 0x7f, 0xff,
	0xfb, 0x16, 0x92, 0xa3, 0x17, 0x50, 0x05, 0x49, 0x63, 0xb9, 0x68, 0x1b, 0x69, 0xb2, 0x68, 0x17,
	0xed, 0xd3, 0xc7, 0x37, 0xc5, 0x44, 0x15, 0x81, 0x28, 0x0e, 0x5e, 0xc6, 0x0b, 0x88, 0x3e, 0x43,
	0x3a, 0xd7, 0x96, 0xfe, 0x62, 0x26, 0xbb, 0x71, 0xbe, 0xfc, 0xaa, 0xcf, 0x11, 0x78, 0x8d, 0x4f,
	0x47, 0xea, 0x0b, 0x51, 0x3d, 0x1f, 0xba, 0x2c, 0xbe, 0x7a, 0xe4, 0xd1, 0xcb, 0x78, 0x01, 0xd1,
	0x27, 0x48, 0xb5, 0xb4, 0x8a, 0xb2, 0x9a, 0x50, 0xa1, 0xb3, 0xcd, 0x2e, 0xde, 0xa7, 0x8f, 0x6f,
	0x9f, 0xfa, 0x1c, 0xf1, 0x85, 0x0a, 0x8d, 0x41, 0x2f, 0x75, 0xfe, 0x27, 0x82, 0x24, 0xbc, 0x02,
	0x7d, 0x84, 0x97, 0x7c, 0x30, 0x4c, 0x0d, 0x44, 0xd4, 0x8a, 0x91, 0x96, 0x29, 0xf7, 0xee, 0x2d,
	0x7e, 0x11, 0xc6, 0xd8, 0x4d, 0x51, 0x09, 0xaf, 0x46, 0xdb, 0x08, 0x4e, 0x6b, 0x6d, 0x1b, 0x4d,
	0x15, 0x1f, 0x0d, 0x97, 0x83, 0x7b, 0xec, 0x16, 0xa3, 0x59, 0x3a, 0xae, 0x14, 0x54, 0xc1, 0xeb,
	0x25, 0xf9, 0x3f, 0x4b, 0xec, 0x2c, 0x0f, 0x41, 0x5c, 0x9b, 0xf2, 0x6f, 0x90, 0xae, 0x16, 0x85,
	0x10, 0x6c, 0x06, 0xd2, 0x33, 0x7f, 0x25, 0x57, 0x4f, 0xb3, 0xb3, 0xd4, 0xc6, 0x9f, 0xec, 0x6a,
	0x94, 0xc1, 0xb3, 0x51, 0xc9, 0x9f, 0x8c, 0x1a, 0x9f, 0x1e, 0xda, 0x1c, 0x43, 0x12, 0xd6, 0x87,
	0xde, 0xc1, 0x7d, 0xcb, 0x4e, 0xc4, 0x0a, 0x53, 0xbb, 0x84, 0x39, 0x35, 0xf5, 0xb3, 0xaf, 0x53,
	0xd0, 0x07, 0x78, 0x4e, 0x84, 0x90, 0x17, 0xd6, 0x3a, 0x44, 0x67, 0x37, 0xbb, 0x78, 0xbf, 0xc5,
	0xf7, 0x7e, 0x38, 0x31, 0x3a, 0xff, 0x0e, 0xf0, 0x6f, 0xb5, 0xe8, 0x01, 0x6e, 0x67, 0x34, 0x72,
	0xe8, 0xdc, 0xa0, 0x1c, 0x12, 0x7f, 0x85, 0x90, 0xb1, 0xf4, 0xd3, 0x6d, 0xe7, 0x55, 0xeb, 0x2c,
	0x76, 0x52, 0x68, 0x0f, 0xc9, 0x8f, 0xbb, 0xf9, 0xab, 0x36, 0x77, 0xee, 0x8b, 0x56, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x90, 0x9c, 0x30, 0xe5, 0xd9, 0x02, 0x00, 0x00,
}
