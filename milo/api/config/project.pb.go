// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/milo/api/config/project.proto

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

// Project is a project definition for Milo.
type Project struct {
	// Consoles is a list of consoles to define under /console/
	Consoles []*Console `protobuf:"bytes,2,rep,name=consoles,proto3" json:"consoles,omitempty"`
	// Headers is a list of defined headers that may be referenced by a console.
	Headers []*Header `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty"`
	// LogoUrl is the URL to the logo for this project.
	// This field is optional. The logo URL must have a host of
	// storage.googleapis.com.
	LogoUrl              string   `protobuf:"bytes,4,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed3e109f2242818b, []int{0}
}

func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetConsoles() []*Console {
	if m != nil {
		return m.Consoles
	}
	return nil
}

func (m *Project) GetHeaders() []*Header {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Project) GetLogoUrl() string {
	if m != nil {
		return m.LogoUrl
	}
	return ""
}

// Link is a link to an internet resource, which will be rendered out as
// an anchor tag <a href="url" alt="alt">text</a>.
type Link struct {
	// Text is displayed as the text between the anchor tags.
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// Url is the URL to link to.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// Alt is the alt text displayed when hovering over the text.
	Alt                  string   `protobuf:"bytes,3,opt,name=alt,proto3" json:"alt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Link) Reset()         { *m = Link{} }
func (m *Link) String() string { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()    {}
func (*Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed3e109f2242818b, []int{1}
}

func (m *Link) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Link.Unmarshal(m, b)
}
func (m *Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Link.Marshal(b, m, deterministic)
}
func (m *Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Link.Merge(m, src)
}
func (m *Link) XXX_Size() int {
	return xxx_messageInfo_Link.Size(m)
}
func (m *Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Link proto.InternalMessageInfo

func (m *Link) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Link) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Link) GetAlt() string {
	if m != nil {
		return m.Alt
	}
	return ""
}

// Oncall contains information about who is currently scheduled as the
// oncall (Sheriff, trooper, etc) for certain rotations.
type Oncall struct {
	// Name is the name of the oncall rotation being displayed.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Url is an URL to a json endpoint with the following format:
	// {
	//   "updated_unix_timestamp": <int>,
	//   "emails": [
	//     "email@somewhere.com",
	//     "email@nowhere.com
	//   ]
	// }
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Oncall) Reset()         { *m = Oncall{} }
func (m *Oncall) String() string { return proto.CompactTextString(m) }
func (*Oncall) ProtoMessage()    {}
func (*Oncall) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed3e109f2242818b, []int{2}
}

func (m *Oncall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Oncall.Unmarshal(m, b)
}
func (m *Oncall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Oncall.Marshal(b, m, deterministic)
}
func (m *Oncall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Oncall.Merge(m, src)
}
func (m *Oncall) XXX_Size() int {
	return xxx_messageInfo_Oncall.Size(m)
}
func (m *Oncall) XXX_DiscardUnknown() {
	xxx_messageInfo_Oncall.DiscardUnknown(m)
}

var xxx_messageInfo_Oncall proto.InternalMessageInfo

func (m *Oncall) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Oncall) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// LinkGroup is a list of links, optionally given a name.
type LinkGroup struct {
	// Name is the name of this list of links. This is optional.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Links is a list of links to display.
	Links                []*Link  `protobuf:"bytes,2,rep,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinkGroup) Reset()         { *m = LinkGroup{} }
func (m *LinkGroup) String() string { return proto.CompactTextString(m) }
func (*LinkGroup) ProtoMessage()    {}
func (*LinkGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed3e109f2242818b, []int{3}
}

func (m *LinkGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinkGroup.Unmarshal(m, b)
}
func (m *LinkGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinkGroup.Marshal(b, m, deterministic)
}
func (m *LinkGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkGroup.Merge(m, src)
}
func (m *LinkGroup) XXX_Size() int {
	return xxx_messageInfo_LinkGroup.Size(m)
}
func (m *LinkGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkGroup.DiscardUnknown(m)
}

var xxx_messageInfo_LinkGroup proto.InternalMessageInfo

func (m *LinkGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LinkGroup) GetLinks() []*Link {
	if m != nil {
		return m.Links
	}
	return nil
}

// ConsoleSummaryGroup is a list of consoles to be displayed as console summaries
// (aka the little bubbles at the top of the console).  This can optionally
// have a group name if specified in the group_link.
// (e.g. "Tree closers", "Experimental", etc)
type ConsoleSummaryGroup struct {
	// Title is a name or label for this group of consoles.  This is optional.
	Title *Link `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// ConsoleIds is a list of console ids to display in this console group.
	// Each console id must be prepended with its related project (e.g.
	// chromium/main) because console ids are project-local.
	// Only consoles from the same project are supported.
	// TODO(hinoka): Allow cross-project consoles.
	ConsoleIds           []string `protobuf:"bytes,2,rep,name=console_ids,json=consoleIds,proto3" json:"console_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsoleSummaryGroup) Reset()         { *m = ConsoleSummaryGroup{} }
func (m *ConsoleSummaryGroup) String() string { return proto.CompactTextString(m) }
func (*ConsoleSummaryGroup) ProtoMessage()    {}
func (*ConsoleSummaryGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed3e109f2242818b, []int{4}
}

func (m *ConsoleSummaryGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsoleSummaryGroup.Unmarshal(m, b)
}
func (m *ConsoleSummaryGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsoleSummaryGroup.Marshal(b, m, deterministic)
}
func (m *ConsoleSummaryGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsoleSummaryGroup.Merge(m, src)
}
func (m *ConsoleSummaryGroup) XXX_Size() int {
	return xxx_messageInfo_ConsoleSummaryGroup.Size(m)
}
func (m *ConsoleSummaryGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsoleSummaryGroup.DiscardUnknown(m)
}

var xxx_messageInfo_ConsoleSummaryGroup proto.InternalMessageInfo

func (m *ConsoleSummaryGroup) GetTitle() *Link {
	if m != nil {
		return m.Title
	}
	return nil
}

func (m *ConsoleSummaryGroup) GetConsoleIds() []string {
	if m != nil {
		return m.ConsoleIds
	}
	return nil
}

// Header is a collection of links, rotation information, and console summaries
// that are displayed at the top of a console, below the tree status information.
// Links and oncall information is always laid out to the left, while
// console groups are laid out on the right.  Each oncall and links group
// take up a row.
type Header struct {
	// Oncalls are a reference to oncall rotations, which is a URL to a json
	// endpoint with the following format:
	// {
	//   "updated_unix_timestamp": <int>,
	//   "emails": [
	//     "email@somewhere.com",
	//     "email@nowhere.com
	//   ]
	// }
	Oncalls []*Oncall `protobuf:"bytes,1,rep,name=oncalls,proto3" json:"oncalls,omitempty"`
	// Links is a list of named groups of web links.
	Links []*LinkGroup `protobuf:"bytes,2,rep,name=links,proto3" json:"links,omitempty"`
	// ConsoleGroups are groups of console summaries, each optionally named.
	ConsoleGroups []*ConsoleSummaryGroup `protobuf:"bytes,3,rep,name=console_groups,json=consoleGroups,proto3" json:"console_groups,omitempty"`
	// TreeStatusHost is the hostname of the chromium-status instance where
	// the tree status of this console is hosted.  If provided, this will appear
	// as the bar at the very top of the page.
	TreeStatusHost string `protobuf:"bytes,4,opt,name=tree_status_host,json=treeStatusHost,proto3" json:"tree_status_host,omitempty"`
	// Id is a reference to the header.
	Id                   string   `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed3e109f2242818b, []int{5}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetOncalls() []*Oncall {
	if m != nil {
		return m.Oncalls
	}
	return nil
}

func (m *Header) GetLinks() []*LinkGroup {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *Header) GetConsoleGroups() []*ConsoleSummaryGroup {
	if m != nil {
		return m.ConsoleGroups
	}
	return nil
}

func (m *Header) GetTreeStatusHost() string {
	if m != nil {
		return m.TreeStatusHost
	}
	return ""
}

func (m *Header) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Console is a waterfall definition consisting of one or more builders.
type Console struct {
	// Id is the reference to the console, and will be the address to make the
	// console reachable from /console/<Project>/<ID>.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name is the longform name of the waterfall, and will be used to be
	// displayed in the title.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// RepoUrl is the URL of the git repository to display as the rows of the console.
	RepoUrl string `protobuf:"bytes,3,opt,name=repo_url,json=repoUrl,proto3" json:"repo_url,omitempty"`
	// Refs are the refs to pull commits from when displaying the console.
	//
	// Users can specify a regular expression to match several refs using
	// "regexp:" prefix, but the regular expression must have:
	//   * a literal prefix with at least two slashes present, e.g.
	//     "refs/release-\d+/foobar" is not allowed, because the literal prefix
	//     "refs/release-" only contains one slash, and
	//   * must not start with ^ or end with $ as they are added automatically.
	//
	// For best results, ensure each ref's has commit's **committer** timestamp
	// monotonically non-decreasing. Gerrit will take care of this if you require
	// each commmit to go through Gerrit by prohibiting "git push" on these refs.
	//
	// Eg. refs/heads/master, regexp:refs/branch-heads/\d+\.\d+
	Refs []string `protobuf:"bytes,14,rep,name=refs,proto3" json:"refs,omitempty"`
	// ExcludeRef is a ref, commits from which are ignored even when they are
	// reachable from the ref specified above. This must be specified as a single
	// fully-qualified ref, i.e. regexp syntax from above is not supported.
	//
	// Note: force pushes to this ref are not supported. Milo uses caching
	// assuming set of commits reachable from this ref may only grow, never lose
	// some commits.
	//
	// E.g. the config below allows to track commits from all release branches,
	// but ignore the commits from the master branch, from which these release
	// branches are branched off:
	//   ref: "regexp:refs/branch-heads/\d+\.\d+"
	//   exlude_ref: "refs/heads/master"
	ExcludeRef string `protobuf:"bytes,13,opt,name=exclude_ref,json=excludeRef,proto3" json:"exclude_ref,omitempty"`
	// ManifestName the name of the manifest the waterfall looks at.
	// This should always be "REVISION".
	// In the future, other manifest names can be supported.
	// TODO(hinoka,iannucci): crbug/832893 - Support custom manifest names, such as "UNPATCHED" / "PATCHED".
	ManifestName string `protobuf:"bytes,5,opt,name=manifest_name,json=manifestName,proto3" json:"manifest_name,omitempty"`
	// Builders is a list of builder configurations to display as the columns of the console.
	Builders []*Builder `protobuf:"bytes,6,rep,name=builders,proto3" json:"builders,omitempty"`
	// FaviconUrl is the URL to the favicon for this console page.
	// This field is optional. The favicon URL must have a host of
	// storage.googleapis.com.
	FaviconUrl string `protobuf:"bytes,7,opt,name=favicon_url,json=faviconUrl,proto3" json:"favicon_url,omitempty"`
	// Header is a collection of links, rotation information, and console summaries
	// displayed under the tree status but above the main console content.
	Header *Header `protobuf:"bytes,9,opt,name=header,proto3" json:"header,omitempty"`
	// HeaderId is a reference to a header.  Only one of Header or HeaderId should
	// be specified.
	HeaderId string `protobuf:"bytes,10,opt,name=header_id,json=headerId,proto3" json:"header_id,omitempty"`
	// If true, this console will not filter out builds marked as Experimental.
	// This field is optional. By default Consoles only show production builds.
	IncludeExperimentalBuilds bool `protobuf:"varint,11,opt,name=include_experimental_builds,json=includeExperimentalBuilds,proto3" json:"include_experimental_builds,omitempty"`
	// If true, only builders view will be available. Console view (i.e. git log
	// based view) will be disabled and users redirected to builder view.
	// Defaults to false.
	BuilderViewOnly      bool     `protobuf:"varint,12,opt,name=builder_view_only,json=builderViewOnly,proto3" json:"builder_view_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Console) Reset()         { *m = Console{} }
func (m *Console) String() string { return proto.CompactTextString(m) }
func (*Console) ProtoMessage()    {}
func (*Console) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed3e109f2242818b, []int{6}
}

func (m *Console) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Console.Unmarshal(m, b)
}
func (m *Console) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Console.Marshal(b, m, deterministic)
}
func (m *Console) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Console.Merge(m, src)
}
func (m *Console) XXX_Size() int {
	return xxx_messageInfo_Console.Size(m)
}
func (m *Console) XXX_DiscardUnknown() {
	xxx_messageInfo_Console.DiscardUnknown(m)
}

var xxx_messageInfo_Console proto.InternalMessageInfo

func (m *Console) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Console) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Console) GetRepoUrl() string {
	if m != nil {
		return m.RepoUrl
	}
	return ""
}

func (m *Console) GetRefs() []string {
	if m != nil {
		return m.Refs
	}
	return nil
}

func (m *Console) GetExcludeRef() string {
	if m != nil {
		return m.ExcludeRef
	}
	return ""
}

func (m *Console) GetManifestName() string {
	if m != nil {
		return m.ManifestName
	}
	return ""
}

func (m *Console) GetBuilders() []*Builder {
	if m != nil {
		return m.Builders
	}
	return nil
}

func (m *Console) GetFaviconUrl() string {
	if m != nil {
		return m.FaviconUrl
	}
	return ""
}

func (m *Console) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Console) GetHeaderId() string {
	if m != nil {
		return m.HeaderId
	}
	return ""
}

func (m *Console) GetIncludeExperimentalBuilds() bool {
	if m != nil {
		return m.IncludeExperimentalBuilds
	}
	return false
}

func (m *Console) GetBuilderViewOnly() bool {
	if m != nil {
		return m.BuilderViewOnly
	}
	return false
}

// Builder is a reference to a Milo builder.
type Builder struct {
	// Name is the BuilderID of the builders you wish to display for this column
	// in the console. e.g.
	//   * "buildbot/chromium.linux/Linux Tests"
	//   * "buildbucket/luci.chromium.try/linux_chromium_rel_ng"
	//
	// If multiple names are specified, the console will show the union of the
	// builders.
	Name []string `protobuf:"bytes,1,rep,name=name,proto3" json:"name,omitempty"`
	// Category describes the hierarchy of the builder on the header of the
	// console as a "|" delimited list.  Neighboring builders with common ancestors
	// will be have their headers merged.
	Category string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	// ShortName is the 1-3 character abbreviation of the builder.
	ShortName            string   `protobuf:"bytes,3,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Builder) Reset()         { *m = Builder{} }
func (m *Builder) String() string { return proto.CompactTextString(m) }
func (*Builder) ProtoMessage()    {}
func (*Builder) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed3e109f2242818b, []int{7}
}

func (m *Builder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Builder.Unmarshal(m, b)
}
func (m *Builder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Builder.Marshal(b, m, deterministic)
}
func (m *Builder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Builder.Merge(m, src)
}
func (m *Builder) XXX_Size() int {
	return xxx_messageInfo_Builder.Size(m)
}
func (m *Builder) XXX_DiscardUnknown() {
	xxx_messageInfo_Builder.DiscardUnknown(m)
}

var xxx_messageInfo_Builder proto.InternalMessageInfo

func (m *Builder) GetName() []string {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Builder) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Builder) GetShortName() string {
	if m != nil {
		return m.ShortName
	}
	return ""
}

func init() {
	proto.RegisterType((*Project)(nil), "milo.Project")
	proto.RegisterType((*Link)(nil), "milo.Link")
	proto.RegisterType((*Oncall)(nil), "milo.Oncall")
	proto.RegisterType((*LinkGroup)(nil), "milo.LinkGroup")
	proto.RegisterType((*ConsoleSummaryGroup)(nil), "milo.ConsoleSummaryGroup")
	proto.RegisterType((*Header)(nil), "milo.Header")
	proto.RegisterType((*Console)(nil), "milo.Console")
	proto.RegisterType((*Builder)(nil), "milo.Builder")
}

func init() {
	proto.RegisterFile("github.com/TriggerMail/luci-go/milo/api/config/project.proto", fileDescriptor_ed3e109f2242818b)
}

var fileDescriptor_ed3e109f2242818b = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x55, 0x9b, 0xb4, 0x49, 0xbe, 0xad, 0x5d, 0x31, 0x37, 0x2e, 0x13, 0xa2, 0x0a, 0x3f, 0xea,
	0xb8, 0x68, 0xa5, 0x71, 0x3f, 0xc1, 0x10, 0x62, 0x9b, 0x10, 0x43, 0x99, 0x86, 0x26, 0x6e, 0xa2,
	0x2c, 0x71, 0x5a, 0x33, 0x27, 0xae, 0x6c, 0x67, 0x5b, 0x2f, 0x78, 0x25, 0x5e, 0x86, 0x17, 0x42,
	0xfe, 0x49, 0xe9, 0x80, 0x3b, 0xfb, 0x9c, 0xef, 0xf7, 0x1c, 0x27, 0x70, 0xb8, 0xe0, 0xb3, 0x7c,
	0x29, 0x78, 0x45, 0x9b, 0x6a, 0xc6, 0xc5, 0x62, 0xce, 0x9a, 0x9c, 0xce, 0x2b, 0xca, 0xf8, 0x3c,
	0x5b, 0xd1, 0x79, 0xce, 0xeb, 0x92, 0x2e, 0xe6, 0x2b, 0xc1, 0xbf, 0x93, 0x5c, 0xcd, 0x56, 0x82,
	0x2b, 0x8e, 0x7c, 0x4d, 0xc7, 0x3f, 0x20, 0xf8, 0x62, 0x61, 0x74, 0x00, 0x61, 0xce, 0x6b, 0xc9,
	0x19, 0x91, 0xb8, 0x3b, 0xf1, 0xa6, 0x3b, 0x87, 0x83, 0x99, 0x8e, 0x99, 0xbd, 0xb7, 0x68, 0xb2,
	0xa1, 0xd1, 0x2b, 0x08, 0x96, 0x24, 0x2b, 0x88, 0x90, 0xd8, 0x33, 0x91, 0xbb, 0x36, 0xf2, 0xc4,
	0x80, 0x49, 0x4b, 0xa2, 0x31, 0x84, 0x8c, 0x2f, 0x78, 0xda, 0x08, 0x86, 0xfd, 0x49, 0x67, 0x1a,
	0x25, 0x81, 0xbe, 0x5f, 0x0a, 0x76, 0xe6, 0x87, 0x9d, 0x51, 0x37, 0x3e, 0x02, 0xff, 0x13, 0xad,
	0x6f, 0x10, 0x02, 0x5f, 0x91, 0x7b, 0x85, 0x3b, 0x26, 0xc8, 0x9c, 0xd1, 0x08, 0x3c, 0x9d, 0xd7,
	0x35, 0x90, 0x3e, 0x6a, 0x24, 0x63, 0x0a, 0x7b, 0x16, 0xc9, 0x98, 0x8a, 0x67, 0xd0, 0x3f, 0xaf,
	0xf3, 0x8c, 0x31, 0x5d, 0xa1, 0xce, 0x2a, 0xd2, 0x56, 0xd0, 0xe7, 0x7f, 0x2b, 0xc4, 0xef, 0x20,
	0xd2, 0xfd, 0x3e, 0x0a, 0xde, 0xac, 0xfe, 0x9b, 0x32, 0x81, 0x1e, 0xa3, 0xf5, 0x4d, 0xab, 0x00,
	0xd8, 0xbd, 0x74, 0x4e, 0x62, 0x89, 0xf8, 0x0a, 0x1e, 0x3b, 0x41, 0x2e, 0x9a, 0xaa, 0xca, 0xc4,
	0xda, 0x16, 0x9b, 0x40, 0x4f, 0x51, 0xc5, 0x6c, 0xb5, 0xbf, 0x12, 0x0d, 0x81, 0x9e, 0xc1, 0x8e,
	0x13, 0x30, 0xa5, 0x85, 0x6d, 0x10, 0x25, 0xe0, 0xa0, 0xd3, 0x42, 0xc6, 0xbf, 0x3a, 0xd0, 0xb7,
	0x0a, 0x6a, 0x81, 0xb9, 0xd9, 0x4b, 0xe2, 0xce, 0xb6, 0xc0, 0x76, 0xd9, 0xa4, 0x25, 0xd1, 0xcb,
	0x87, 0xe3, 0xee, 0xfd, 0xe9, 0x6a, 0xa6, 0x72, 0x33, 0xa3, 0xb7, 0x30, 0x6c, 0x5b, 0x2f, 0x34,
	0xde, 0xda, 0x36, 0x7e, 0x60, 0xf0, 0xf6, 0x3e, 0xc9, 0xc0, 0x25, 0x98, 0x9b, 0x44, 0x53, 0x18,
	0x29, 0x41, 0x48, 0x2a, 0x55, 0xa6, 0x1a, 0x99, 0x2e, 0xb9, 0x54, 0xce, 0xd1, 0xa1, 0xc6, 0x2f,
	0x0c, 0x7c, 0xc2, 0xa5, 0x42, 0x43, 0xe8, 0xd2, 0x02, 0xf7, 0x0c, 0xd7, 0xa5, 0x45, 0xfc, 0xd3,
	0x83, 0xc0, 0x35, 0x70, 0x5c, 0xa7, 0xe5, 0x36, 0x0e, 0x74, 0xb7, 0x1c, 0x18, 0x43, 0x28, 0xc8,
	0xca, 0xbe, 0x19, 0xeb, 0x74, 0xa0, 0xef, 0x97, 0xc2, 0x78, 0x2c, 0x48, 0x29, 0xf1, 0xd0, 0x48,
	0x67, 0xce, 0x5a, 0x55, 0x72, 0x9f, 0xb3, 0xa6, 0x20, 0xa9, 0x20, 0x25, 0x1e, 0x98, 0x0c, 0x70,
	0x50, 0x42, 0x4a, 0xf4, 0x1c, 0x06, 0x55, 0x56, 0xd3, 0x92, 0x48, 0x95, 0x9a, 0x66, 0x76, 0xb4,
	0xdd, 0x16, 0xfc, 0xac, 0x9b, 0x1e, 0x40, 0x78, 0xdd, 0x50, 0x66, 0x5e, 0x74, 0x7f, 0xfb, 0xed,
	0x1f, 0x5b, 0x34, 0xd9, 0xd0, 0xba, 0x61, 0x99, 0xdd, 0xd2, 0x9c, 0xd7, 0x66, 0xc4, 0xc0, 0x36,
	0x74, 0x90, 0x9e, 0xf2, 0x05, 0xf4, 0xed, 0xfb, 0xc7, 0x91, 0x79, 0x0a, 0x0f, 0xbf, 0x0d, 0xc7,
	0xa1, 0x7d, 0x88, 0xec, 0x29, 0xa5, 0x05, 0x06, 0x53, 0x24, 0xb4, 0xc0, 0x69, 0x81, 0x8e, 0x60,
	0x9f, 0xd6, 0x76, 0x29, 0x72, 0xbf, 0x22, 0x82, 0x56, 0xa4, 0x56, 0x19, 0x4b, 0xcd, 0x10, 0x12,
	0xef, 0x4c, 0x3a, 0xd3, 0x30, 0x19, 0xbb, 0x90, 0x0f, 0x5b, 0x11, 0x66, 0x5c, 0x89, 0x5e, 0xc3,
	0x23, 0x37, 0x6f, 0x7a, 0x4b, 0xc9, 0x5d, 0xca, 0x6b, 0xb6, 0xc6, 0xbb, 0x26, 0x6b, 0xcf, 0x11,
	0x5f, 0x29, 0xb9, 0x3b, 0xaf, 0xd9, 0xfa, 0xcc, 0x0f, 0xc3, 0x51, 0x74, 0xe6, 0x87, 0xfe, 0xa8,
	0x97, 0x78, 0x82, 0x94, 0xf1, 0x15, 0x04, 0x6e, 0xeb, 0xad, 0x2f, 0xc4, 0xdb, 0xf8, 0xf3, 0x04,
	0xc2, 0x3c, 0x53, 0x64, 0xc1, 0xc5, 0xda, 0xf9, 0xb6, 0xb9, 0xa3, 0xa7, 0x00, 0x72, 0xc9, 0x85,
	0x13, 0xda, 0xba, 0x17, 0x19, 0x44, 0xab, 0x7c, 0x1c, 0x7e, 0xeb, 0xdb, 0x5f, 0xd1, 0x75, 0xdf,
	0xfc, 0x83, 0xde, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x51, 0x1a, 0xf0, 0x82, 0xb9, 0x04, 0x00,
	0x00,
}
