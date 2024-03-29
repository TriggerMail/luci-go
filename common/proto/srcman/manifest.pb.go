// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/common/proto/srcman/manifest.proto

package srcman

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

// A Manifest attempts to make an accurate accounting of source/data directories
// during the execution of a LUCI task.
//
// These directories are primarily in the form of e.g. git checkouts of
// source, but also include things like isolated hashes and CIPD package
// deployments. In the future, other deployment forms may be supported (like
// other SCMs).
//
// The purpose of this manifest is so that other parts of the LUCI stack (e.g.
// Milo) can work with the descriptions of this deployed data as a first-class
// citizen. Initially this Manifest will be used to allow Milo to display diffs
// between jobs, but it will also be useful for tools and humans to get a
// record of exactly what data went into this LUCI task.
//
// Source Manifests can be emitted from recipes using the
// 'recipe_engine/source_manifest' module.
type Manifest struct {
	// Version will increment on backwards-incompatible changes only. Backwards
	// compatible changes will not alter this version number.
	//
	// Currently, the only valid version number is 0.
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Map of local file system directory path (with forward slashes) to
	// a Directory message containing one or more deployments.
	//
	// The local path is relative to some job-specific root. This should be used
	// for informational/display/organization purposes. In particular, think VERY
	// CAREFULLY before you configure remote services/recipes to look for
	// particular filesystem layouts here. For example, if you want to look for
	// "the version of chromium/src checked out by the job", prefer to look for
	// a Directory which checks out "chromium/src", as opposed to assuming this
	// checkout lives in a top-level folder called "src". The reason for this is
	// that jobs SHOULD reserve the right to do their checkouts in any way they
	// please.
	//
	// If you feel like you need to make some service configuration which uses one
	// of these local filesystem paths as a key, please consult with the Chrome
	// Infrastructure team to see if there's a better alternative.
	//
	// Ex.
	//   "": {...}  // root directory
	//   "src/third_party/something": {...}
	Directories          map[string]*Manifest_Directory `protobuf:"bytes,2,rep,name=directories,proto3" json:"directories,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Manifest) Reset()         { *m = Manifest{} }
func (m *Manifest) String() string { return proto.CompactTextString(m) }
func (*Manifest) ProtoMessage()    {}
func (*Manifest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f003f10f6b5b7bd, []int{0}
}

func (m *Manifest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manifest.Unmarshal(m, b)
}
func (m *Manifest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manifest.Marshal(b, m, deterministic)
}
func (m *Manifest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manifest.Merge(m, src)
}
func (m *Manifest) XXX_Size() int {
	return xxx_messageInfo_Manifest.Size(m)
}
func (m *Manifest) XXX_DiscardUnknown() {
	xxx_messageInfo_Manifest.DiscardUnknown(m)
}

var xxx_messageInfo_Manifest proto.InternalMessageInfo

func (m *Manifest) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Manifest) GetDirectories() map[string]*Manifest_Directory {
	if m != nil {
		return m.Directories
	}
	return nil
}

type Manifest_GitCheckout struct {
	// The canonicalized URL of the original repo that is considered the “source
	// of truth” for the source code.
	//
	// Ex.
	//   https://chromium.googlesource.com/chromium/tools/build
	//   https://chromium.googlesource.com/infra/luci/recipes-py
	RepoUrl string `protobuf:"bytes,1,opt,name=repo_url,json=repoUrl,proto3" json:"repo_url,omitempty"`
	// If different from repo_url, this can be the URL of the repo that the source
	// was actually fetched from (i.e. a mirror).
	//
	// If this is empty, it's presumed to be equal to repo_url.
	//
	// Ex.
	//   https://github.com/luci/recipes-py
	FetchUrl string `protobuf:"bytes,2,opt,name=fetch_url,json=fetchUrl,proto3" json:"fetch_url,omitempty"`
	// The fully resolved revision (commit hash) of the source.
	//
	// This must always be a revision on the hosted repo (not any locally
	// generated commit).
	//
	// Ex.
	//   3617b0eea7ec74b8e731a23fed2f4070cbc284c4
	Revision string `protobuf:"bytes,3,opt,name=revision,proto3" json:"revision,omitempty"`
	// The ref that the task used to resolve/fetch the revision of the source
	// (if any).
	//
	// This must always be a ref on the hosted repo (not any local alias
	// like 'refs/remotes/...').
	//
	// This must always be an absolute ref (i.e. starts with 'refs/'). An
	// example of a non-absolute ref would be 'master'.
	//
	// Ex.
	//   refs/heads/master
	FetchRef string `protobuf:"bytes,4,opt,name=fetch_ref,json=fetchRef,proto3" json:"fetch_ref,omitempty"`
	// If the checkout had a CL associated with it (i.e. a gerrit commit), this
	// is the fully resolved revision (commit hash) of the CL. If there was no
	// CL, this is empty. Typically the checkout application (e.g. bot_update)
	// rebases this revision on top of the `revision` fetched above.
	//
	// If specified, this must always be a revision on the hosted repo (not any
	// locally generated commit).
	//
	// Ex.
	//   6b0b5c12443cfb93305f8d9e21f8d762c8dad9f0
	PatchRevision string `protobuf:"bytes,5,opt,name=patch_revision,json=patchRevision,proto3" json:"patch_revision,omitempty"`
	// If the checkout had a CL associated with it, this is the ref that the
	// task used to fetch patch_revision. If `patch_revision` is supplied, this
	// field is required. If there was no CL, this is empty.
	//
	// If specified, this must always be a ref on the hosted repo (not any local
	// alias like 'refs/remotes/...').
	//
	// This must always be an absolute ref (i.e. starts with 'refs/').
	//
	// Ex.
	//   refs/changes/04/511804/4
	PatchFetchRef        string   `protobuf:"bytes,6,opt,name=patch_fetch_ref,json=patchFetchRef,proto3" json:"patch_fetch_ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Manifest_GitCheckout) Reset()         { *m = Manifest_GitCheckout{} }
func (m *Manifest_GitCheckout) String() string { return proto.CompactTextString(m) }
func (*Manifest_GitCheckout) ProtoMessage()    {}
func (*Manifest_GitCheckout) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f003f10f6b5b7bd, []int{0, 0}
}

func (m *Manifest_GitCheckout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manifest_GitCheckout.Unmarshal(m, b)
}
func (m *Manifest_GitCheckout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manifest_GitCheckout.Marshal(b, m, deterministic)
}
func (m *Manifest_GitCheckout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manifest_GitCheckout.Merge(m, src)
}
func (m *Manifest_GitCheckout) XXX_Size() int {
	return xxx_messageInfo_Manifest_GitCheckout.Size(m)
}
func (m *Manifest_GitCheckout) XXX_DiscardUnknown() {
	xxx_messageInfo_Manifest_GitCheckout.DiscardUnknown(m)
}

var xxx_messageInfo_Manifest_GitCheckout proto.InternalMessageInfo

func (m *Manifest_GitCheckout) GetRepoUrl() string {
	if m != nil {
		return m.RepoUrl
	}
	return ""
}

func (m *Manifest_GitCheckout) GetFetchUrl() string {
	if m != nil {
		return m.FetchUrl
	}
	return ""
}

func (m *Manifest_GitCheckout) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *Manifest_GitCheckout) GetFetchRef() string {
	if m != nil {
		return m.FetchRef
	}
	return ""
}

func (m *Manifest_GitCheckout) GetPatchRevision() string {
	if m != nil {
		return m.PatchRevision
	}
	return ""
}

func (m *Manifest_GitCheckout) GetPatchFetchRef() string {
	if m != nil {
		return m.PatchFetchRef
	}
	return ""
}

type Manifest_CIPDPackage struct {
	// The package pattern that was given to the CIPD client (if known).
	//
	// Ex.
	//   infra/tools/luci/led/${platform}
	PackagePattern string `protobuf:"bytes,1,opt,name=package_pattern,json=packagePattern,proto3" json:"package_pattern,omitempty"`
	// The fully resolved instance ID of the deployed package.
	//
	// Ex.
	//   0cfafb3a705bd8f05f86c6444ff500397fbb711c
	InstanceId string `protobuf:"bytes,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	// The unresolved version ID of the deployed package (if known).
	//
	// Ex.
	//   git_revision:aaf3a2cfccc227b5141caa1b6b3502c9907d7420
	//   latest
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Manifest_CIPDPackage) Reset()         { *m = Manifest_CIPDPackage{} }
func (m *Manifest_CIPDPackage) String() string { return proto.CompactTextString(m) }
func (*Manifest_CIPDPackage) ProtoMessage()    {}
func (*Manifest_CIPDPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f003f10f6b5b7bd, []int{0, 1}
}

func (m *Manifest_CIPDPackage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manifest_CIPDPackage.Unmarshal(m, b)
}
func (m *Manifest_CIPDPackage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manifest_CIPDPackage.Marshal(b, m, deterministic)
}
func (m *Manifest_CIPDPackage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manifest_CIPDPackage.Merge(m, src)
}
func (m *Manifest_CIPDPackage) XXX_Size() int {
	return xxx_messageInfo_Manifest_CIPDPackage.Size(m)
}
func (m *Manifest_CIPDPackage) XXX_DiscardUnknown() {
	xxx_messageInfo_Manifest_CIPDPackage.DiscardUnknown(m)
}

var xxx_messageInfo_Manifest_CIPDPackage proto.InternalMessageInfo

func (m *Manifest_CIPDPackage) GetPackagePattern() string {
	if m != nil {
		return m.PackagePattern
	}
	return ""
}

func (m *Manifest_CIPDPackage) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *Manifest_CIPDPackage) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type Manifest_Isolated struct {
	// The namespace of the isolated document.
	//
	// Ex.
	//   default-gzip
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// The hash of the isolated document.
	//
	// Ex.
	//   62a7df62ea122380afb306bb4d9cdac1bc7e9a96
	Hash                 string   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Manifest_Isolated) Reset()         { *m = Manifest_Isolated{} }
func (m *Manifest_Isolated) String() string { return proto.CompactTextString(m) }
func (*Manifest_Isolated) ProtoMessage()    {}
func (*Manifest_Isolated) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f003f10f6b5b7bd, []int{0, 2}
}

func (m *Manifest_Isolated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manifest_Isolated.Unmarshal(m, b)
}
func (m *Manifest_Isolated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manifest_Isolated.Marshal(b, m, deterministic)
}
func (m *Manifest_Isolated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manifest_Isolated.Merge(m, src)
}
func (m *Manifest_Isolated) XXX_Size() int {
	return xxx_messageInfo_Manifest_Isolated.Size(m)
}
func (m *Manifest_Isolated) XXX_DiscardUnknown() {
	xxx_messageInfo_Manifest_Isolated.DiscardUnknown(m)
}

var xxx_messageInfo_Manifest_Isolated proto.InternalMessageInfo

func (m *Manifest_Isolated) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Manifest_Isolated) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// A Directory contains one or more descriptions of deployed artifacts. Note
// that due to the practical nature of jobs on bots, it may be the case that
// a given directory contains e.g. a git checkout and multiple cipd packages.
type Manifest_Directory struct {
	GitCheckout *Manifest_GitCheckout `protobuf:"bytes,1,opt,name=git_checkout,json=gitCheckout,proto3" json:"git_checkout,omitempty"`
	// The canonicalized hostname of the CIPD server which hosts the CIPD
	// packages (if any). If no CIPD packages are in this Directory, this must
	// be blank.
	//
	// Ex.
	//   chrome-infra-packages.appspot.com
	CipdServerHost string `protobuf:"bytes,2,opt,name=cipd_server_host,json=cipdServerHost,proto3" json:"cipd_server_host,omitempty"`
	// Maps CIPD package name to CIPDPackage.
	//
	// Ex.
	//   "some/package/name": {...}
	//   "other/package": {...}
	CipdPackage map[string]*Manifest_CIPDPackage `protobuf:"bytes,4,rep,name=cipd_package,json=cipdPackage,proto3" json:"cipd_package,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The canonicalized hostname of the isolated server which hosts the
	// isolated. If no Isolated objects are in this Directory, this must be
	// blank.
	//
	// Ex.
	//   isolateserver.appspot.com
	IsolatedServerHost string `protobuf:"bytes,5,opt,name=isolated_server_host,json=isolatedServerHost,proto3" json:"isolated_server_host,omitempty"`
	// A list of all isolateds which have been installed in this directory.
	Isolated             []*Manifest_Isolated `protobuf:"bytes,6,rep,name=isolated,proto3" json:"isolated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Manifest_Directory) Reset()         { *m = Manifest_Directory{} }
func (m *Manifest_Directory) String() string { return proto.CompactTextString(m) }
func (*Manifest_Directory) ProtoMessage()    {}
func (*Manifest_Directory) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f003f10f6b5b7bd, []int{0, 3}
}

func (m *Manifest_Directory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manifest_Directory.Unmarshal(m, b)
}
func (m *Manifest_Directory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manifest_Directory.Marshal(b, m, deterministic)
}
func (m *Manifest_Directory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manifest_Directory.Merge(m, src)
}
func (m *Manifest_Directory) XXX_Size() int {
	return xxx_messageInfo_Manifest_Directory.Size(m)
}
func (m *Manifest_Directory) XXX_DiscardUnknown() {
	xxx_messageInfo_Manifest_Directory.DiscardUnknown(m)
}

var xxx_messageInfo_Manifest_Directory proto.InternalMessageInfo

func (m *Manifest_Directory) GetGitCheckout() *Manifest_GitCheckout {
	if m != nil {
		return m.GitCheckout
	}
	return nil
}

func (m *Manifest_Directory) GetCipdServerHost() string {
	if m != nil {
		return m.CipdServerHost
	}
	return ""
}

func (m *Manifest_Directory) GetCipdPackage() map[string]*Manifest_CIPDPackage {
	if m != nil {
		return m.CipdPackage
	}
	return nil
}

func (m *Manifest_Directory) GetIsolatedServerHost() string {
	if m != nil {
		return m.IsolatedServerHost
	}
	return ""
}

func (m *Manifest_Directory) GetIsolated() []*Manifest_Isolated {
	if m != nil {
		return m.Isolated
	}
	return nil
}

// Links to an externally stored Manifest proto.
type ManifestLink struct {
	// The fully qualified url of the Manifest proto. It's expected that this is
	// a binary logdog stream consisting of exactly one Manifest proto. For now
	// this will always be the `logdog` uri scheme, though it's feasible to put
	// other uri schemes here later.
	//
	// Ex.
	//   logdog://logs.chromium.org/infra/build/12345/+/some/path
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// The hash of the Manifest's raw binary form (i.e. the bytes at the end of
	// `url`, without any interpretation or decoding). Milo will use this as an
	// optimization; Manifests will be interned once into Milo's datastore.
	// Future hashes which match will not be loaded from the url, but will be
	// assumed to be identical. If the sha256 doesn't match the data at the URL,
	// Milo may render this build with the wrong manifest.
	//
	// This is the raw sha256, so it must be exactly 32 bytes.
	Sha256               []byte   `protobuf:"bytes,2,opt,name=sha256,proto3" json:"sha256,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManifestLink) Reset()         { *m = ManifestLink{} }
func (m *ManifestLink) String() string { return proto.CompactTextString(m) }
func (*ManifestLink) ProtoMessage()    {}
func (*ManifestLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f003f10f6b5b7bd, []int{1}
}

func (m *ManifestLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManifestLink.Unmarshal(m, b)
}
func (m *ManifestLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManifestLink.Marshal(b, m, deterministic)
}
func (m *ManifestLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManifestLink.Merge(m, src)
}
func (m *ManifestLink) XXX_Size() int {
	return xxx_messageInfo_ManifestLink.Size(m)
}
func (m *ManifestLink) XXX_DiscardUnknown() {
	xxx_messageInfo_ManifestLink.DiscardUnknown(m)
}

var xxx_messageInfo_ManifestLink proto.InternalMessageInfo

func (m *ManifestLink) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ManifestLink) GetSha256() []byte {
	if m != nil {
		return m.Sha256
	}
	return nil
}

func init() {
	proto.RegisterType((*Manifest)(nil), "srcman.Manifest")
	proto.RegisterMapType((map[string]*Manifest_Directory)(nil), "srcman.Manifest.DirectoriesEntry")
	proto.RegisterType((*Manifest_GitCheckout)(nil), "srcman.Manifest.GitCheckout")
	proto.RegisterType((*Manifest_CIPDPackage)(nil), "srcman.Manifest.CIPDPackage")
	proto.RegisterType((*Manifest_Isolated)(nil), "srcman.Manifest.Isolated")
	proto.RegisterType((*Manifest_Directory)(nil), "srcman.Manifest.Directory")
	proto.RegisterMapType((map[string]*Manifest_CIPDPackage)(nil), "srcman.Manifest.Directory.CipdPackageEntry")
	proto.RegisterType((*ManifestLink)(nil), "srcman.ManifestLink")
}

func init() {
	proto.RegisterFile("github.com/TriggerMail/luci-go/common/proto/srcman/manifest.proto", fileDescriptor_9f003f10f6b5b7bd)
}

var fileDescriptor_9f003f10f6b5b7bd = []byte{
	// 546 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x5d, 0x8b, 0xd3, 0x4c,
	0x14, 0xa6, 0xdf, 0xed, 0x49, 0xdf, 0xbe, 0x65, 0x10, 0xc9, 0xc6, 0x05, 0xd7, 0x05, 0xb5, 0x20,
	0xa4, 0x4b, 0x65, 0x75, 0x11, 0xc1, 0x8b, 0xae, 0x1f, 0x05, 0x95, 0x12, 0xf1, 0x46, 0x84, 0x30,
	0x4e, 0xa6, 0xcd, 0xd0, 0x66, 0x26, 0xce, 0x4c, 0x0b, 0xfd, 0x11, 0xfe, 0x34, 0x6f, 0xfd, 0x3d,
	0x92, 0xc9, 0x8c, 0xc9, 0x7e, 0xdd, 0x9d, 0xf3, 0x9c, 0xe7, 0x3c, 0x67, 0xce, 0x47, 0x02, 0x2f,
	0xd7, 0x22, 0x24, 0xa9, 0x14, 0x19, 0xdb, 0x65, 0xa1, 0x90, 0xeb, 0xe9, 0x76, 0x47, 0xd8, 0x94,
	0x88, 0x2c, 0x13, 0x7c, 0x9a, 0x4b, 0xa1, 0xc5, 0x54, 0x49, 0x92, 0x61, 0x3e, 0xcd, 0x30, 0x67,
	0x2b, 0xaa, 0x74, 0x68, 0x50, 0xd4, 0x2d, 0xe1, 0xd3, 0x3f, 0x3d, 0xe8, 0x7f, 0xb2, 0x21, 0xe4,
	0x43, 0x6f, 0x4f, 0xa5, 0x62, 0x82, 0xfb, 0x8d, 0x93, 0xc6, 0xa4, 0x13, 0x39, 0x17, 0xcd, 0xc1,
	0x4b, 0x98, 0xa4, 0x44, 0x0b, 0xc9, 0xa8, 0xf2, 0x9b, 0x27, 0xad, 0x89, 0x37, 0x7b, 0x14, 0x96,
	0x22, 0xa1, 0x13, 0x08, 0x2f, 0x2b, 0xce, 0x5b, 0xae, 0xe5, 0x21, 0xaa, 0x67, 0x05, 0xbf, 0x1b,
	0xe0, 0xbd, 0x67, 0x7a, 0x9e, 0x52, 0xb2, 0x11, 0x3b, 0x8d, 0x8e, 0xa0, 0x2f, 0x69, 0x2e, 0xe2,
	0x9d, 0xdc, 0x9a, 0x7a, 0x83, 0xa8, 0x57, 0xf8, 0x5f, 0xe5, 0x16, 0x3d, 0x80, 0xc1, 0x8a, 0x6a,
	0x92, 0x9a, 0x58, 0xd3, 0xc4, 0xfa, 0x06, 0x28, 0x82, 0x41, 0x91, 0xb7, 0x67, 0xe6, 0x9d, 0xad,
	0x32, 0xe6, 0xfc, 0x2a, 0x51, 0xd2, 0x95, 0xdf, 0xae, 0x25, 0x46, 0x74, 0x85, 0x1e, 0xc3, 0x28,
	0xc7, 0x65, 0xd0, 0xa6, 0x77, 0x0c, 0xe3, 0x3f, 0x83, 0x46, 0x4e, 0xe3, 0x09, 0xfc, 0x5f, 0xd2,
	0x2a, 0xa5, 0x6e, 0x8d, 0xf7, 0xce, 0xca, 0x05, 0x3f, 0xc1, 0x9b, 0x2f, 0x96, 0x97, 0x4b, 0x4c,
	0x36, 0x78, 0x4d, 0xd1, 0xd3, 0x22, 0xcd, 0x98, 0x71, 0x8e, 0xb5, 0xa6, 0x92, 0xdb, 0xae, 0x46,
	0x16, 0x5e, 0x96, 0x28, 0x7a, 0x08, 0x1e, 0xe3, 0x4a, 0x63, 0x4e, 0x68, 0xcc, 0x12, 0xdb, 0x1e,
	0x38, 0x68, 0x91, 0xd4, 0xf7, 0x50, 0xf6, 0xe7, 0xdc, 0xe0, 0x35, 0xf4, 0x17, 0x4a, 0x6c, 0xb1,
	0xa6, 0x09, 0x3a, 0x86, 0x01, 0xc7, 0x19, 0x55, 0x39, 0x26, 0xd4, 0x56, 0xaa, 0x00, 0x84, 0xa0,
	0x9d, 0x62, 0x95, 0x5a, 0x75, 0x63, 0x07, 0xbf, 0x5a, 0x30, 0x70, 0x2b, 0x3a, 0xa0, 0x37, 0x30,
	0x5c, 0x33, 0x1d, 0x13, 0xbb, 0x0e, 0x23, 0xe1, 0xcd, 0x8e, 0x6f, 0x2c, 0xb5, 0xb6, 0xb2, 0xc8,
	0x5b, 0xd7, 0xf6, 0x37, 0x81, 0x31, 0x61, 0x79, 0x12, 0x2b, 0x2a, 0xf7, 0x54, 0xc6, 0xa9, 0x50,
	0xda, 0x96, 0x1b, 0x15, 0xf8, 0x17, 0x03, 0x7f, 0x10, 0x4a, 0xa3, 0xcf, 0x30, 0x34, 0x4c, 0x3b,
	0x08, 0xbf, 0x6d, 0xee, 0xe7, 0xd9, 0x9d, 0xf7, 0x73, 0x08, 0xe7, 0x2c, 0x4f, 0xec, 0x60, 0xed,
	0x25, 0x91, 0x0a, 0x41, 0x67, 0x70, 0x8f, 0xd9, 0x31, 0x5c, 0xa9, 0x5e, 0xae, 0x13, 0xb9, 0x58,
	0xed, 0x05, 0xe7, 0xd0, 0x77, 0xa8, 0xdf, 0x35, 0xd5, 0x8f, 0x6e, 0x54, 0x77, 0x93, 0x8d, 0xfe,
	0x51, 0x83, 0xef, 0x30, 0xbe, 0xfe, 0x12, 0x34, 0x86, 0xd6, 0x86, 0x1e, 0xec, 0xc4, 0x0b, 0x13,
	0xcd, 0xa0, 0xb3, 0xc7, 0xdb, 0x1d, 0x35, 0xdd, 0xdf, 0x36, 0xc2, 0xda, 0x99, 0x44, 0x25, 0xf5,
	0x55, 0xf3, 0xa2, 0x11, 0x7c, 0x83, 0xf1, 0xf5, 0x2f, 0xe6, 0x16, 0xf5, 0xb3, 0xab, 0xea, 0xc1,
	0xdd, 0x53, 0xab, 0x69, 0x9f, 0x5e, 0xc0, 0xd0, 0x11, 0x3e, 0x32, 0xbe, 0x29, 0x74, 0xab, 0xef,
	0xac, 0x30, 0xd1, 0x7d, 0xe8, 0xaa, 0x14, 0xcf, 0xce, 0x5f, 0x18, 0xe1, 0x61, 0x64, 0xbd, 0x1f,
	0x5d, 0xf3, 0x87, 0x78, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x0c, 0xc7, 0x13, 0x5c, 0x04,
	0x00, 0x00,
}
