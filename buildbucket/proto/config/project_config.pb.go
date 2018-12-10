// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/buildbucket/proto/config/project_config.proto

package configpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// Toggle is a boolean with an extra state UNSET.
// When protobuf messages are merged, UNSET does not overwrite an existing
// value.
// TODO(nodir): replace with Trinary in ../common.proto.
type Toggle int32

const (
	Toggle_UNSET Toggle = 0
	Toggle_YES   Toggle = 1
	Toggle_NO    Toggle = 2
)

var Toggle_name = map[int32]string{
	0: "UNSET",
	1: "YES",
	2: "NO",
}

var Toggle_value = map[string]int32{
	"UNSET": 0,
	"YES":   1,
	"NO":    2,
}

func (x Toggle) String() string {
	return proto.EnumName(Toggle_name, int32(x))
}

func (Toggle) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_90cd2cb373fa8af0, []int{0}
}

// A buiildbucket user role.
// Defines what a user can do.
//
// The order of enum member tags is important.
// A role with a higher tag number can perform any action that a role with a
// lower tag number can perform.
type Acl_Role int32

const (
	// Can do read-only operations, such as search for builds.
	Acl_READER Acl_Role = 0
	// Same as READER + can schedule and cancel builds.
	Acl_SCHEDULER Acl_Role = 1
	// Can do all write operations.
	Acl_WRITER Acl_Role = 2
)

var Acl_Role_name = map[int32]string{
	0: "READER",
	1: "SCHEDULER",
	2: "WRITER",
}

var Acl_Role_value = map[string]int32{
	"READER":    0,
	"SCHEDULER": 1,
	"WRITER":    2,
}

func (x Acl_Role) String() string {
	return proto.EnumName(Acl_Role_name, int32(x))
}

func (Acl_Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_90cd2cb373fa8af0, []int{0, 0}
}

// A single access control rule.
type Acl struct {
	// Role denotes a list of actions that an identity can perform.
	Role Acl_Role `protobuf:"varint,1,opt,name=role,proto3,enum=buildbucket.Acl_Role" json:"role,omitempty"`
	// Name of the group defined in the auth service.
	Group string `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	// An email address or a full identity string "kind:name". See auth service
	// on kinds of identities. Anonymous users are "anonymous:anonymous".
	// Either identity or group must be present, not both.
	Identity             string   `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Acl) Reset()         { *m = Acl{} }
func (m *Acl) String() string { return proto.CompactTextString(m) }
func (*Acl) ProtoMessage()    {}
func (*Acl) Descriptor() ([]byte, []int) {
	return fileDescriptor_90cd2cb373fa8af0, []int{0}
}

func (m *Acl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Acl.Unmarshal(m, b)
}
func (m *Acl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Acl.Marshal(b, m, deterministic)
}
func (m *Acl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Acl.Merge(m, src)
}
func (m *Acl) XXX_Size() int {
	return xxx_messageInfo_Acl.Size(m)
}
func (m *Acl) XXX_DiscardUnknown() {
	xxx_messageInfo_Acl.DiscardUnknown(m)
}

var xxx_messageInfo_Acl proto.InternalMessageInfo

func (m *Acl) GetRole() Acl_Role {
	if m != nil {
		return m.Role
	}
	return Acl_READER
}

func (m *Acl) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *Acl) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

// A set of Acl messages. Can be referenced in a bucket by name.
type AclSet struct {
	// A name of the ACL set. Required. Must match regex '^[a-z0-9_]+$'.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// List of access control rules.
	// The order does not matter.
	Acls                 []*Acl   `protobuf:"bytes,2,rep,name=acls,proto3" json:"acls,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AclSet) Reset()         { *m = AclSet{} }
func (m *AclSet) String() string { return proto.CompactTextString(m) }
func (*AclSet) ProtoMessage()    {}
func (*AclSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_90cd2cb373fa8af0, []int{1}
}

func (m *AclSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AclSet.Unmarshal(m, b)
}
func (m *AclSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AclSet.Marshal(b, m, deterministic)
}
func (m *AclSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AclSet.Merge(m, src)
}
func (m *AclSet) XXX_Size() int {
	return xxx_messageInfo_AclSet.Size(m)
}
func (m *AclSet) XXX_DiscardUnknown() {
	xxx_messageInfo_AclSet.DiscardUnknown(m)
}

var xxx_messageInfo_AclSet proto.InternalMessageInfo

func (m *AclSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AclSet) GetAcls() []*Acl {
	if m != nil {
		return m.Acls
	}
	return nil
}

// Defines a swarmbucket builder or a builder mixin. A builder has a name, a
// category and specifies what should happen if a build is scheduled to that
// builder.
//
// SECURITY WARNING: if adding more fields to this message, keep in mind that
// a user that has permissions to schedule a build to the bucket, can override
// this config.
//
// Next tag: 21.
type Builder struct {
	// Name of the builder or builder mixin.
	//
	// If a builder name, will be propagated to "builder" build tag and
	// "buildername" recipe property.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Names of mixins to apply to this builder definition.
	//
	// FLATTENING
	//
	// Final builder/mixin values are computed as follows:
	// - start with an empty builder definition.
	// - if this is a builder, apply values in a bucket's builder_defaults,
	//   flattened in advance.
	// - apply each mixin, flattened in advance, in the same order.
	// - apply values in this builder/mixin.
	//
	// EXAMPLE
	//
	//   A definition
	//
	//     builder_mixins {
	//       name: "foo"
	//       dimensions: "os:Linux"
	//       dimensions: "cpu:x86"
	//       recipe {
	//         repository: "https://example.com"
	//         name: "x"
	//       }
	//     }
	//     builder_mixins {
	//       name: "bar"
	//       dimensions: "cores:8"
	//       dimensions: "cpu:x86-64"
	//     }
	//     bucket {
	//       name: "luci.x.try"
	//       swarming {
	//         builders {
	//           name: "release"
	//           mixins: "foo"
	//           mixins: "bar"
	//           recipe {
	//             name: "y"
	//           }
	//         }
	//       }
	//     }
	//
	//   is equivalent to
	//
	//     bucket {
	//      name: "luci.x.try"
	//      swarming {
	//         builders {
	//           name: "release"
	//           dimensions: "os:Linux"
	//           dimensions: "cpu:x86-64"
	//           dimensions: "cores:8"
	//           recipe {
	//             repository: "https://example.com"
	//             name: "y"
	//           }
	//         }
	//       }
	//     }
	//
	// A NOTE ON DIAMOND MERGES
	//
	// Given
	//   B mixes in A and overrides some values defined in A
	//   C mixes in A
	//   D mixes in B and C
	// B's overrides won't affect D because D mixes in C after B.
	//
	//   builder_mixins {
	//     name: "A"
	//     dimensions: "dim:a"
	//   }
	//   builder_mixins {
	//     name: "B"
	//     mixins: "A"
	//     dimensions: "dim:b"
	//   }
	//   builder_mixins {
	//     name: "C"
	//     mixins: "A"
	//   }
	//   ...
	//   builders {
	//     name: "D"
	//     mixins: "B"
	//     mixins: "C"
	//   }
	//
	// D's dim will be "a", not "b" because it is "a" in C which is applied after
	// B.
	//
	// OTHER
	//
	// Circular references are prohibited.
	Mixins []string `protobuf:"bytes,10,rep,name=mixins,proto3" json:"mixins,omitempty"`
	// Builder category. Will be used for visual grouping, for example in Code Review.
	Category string `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
	// Will be become to swarming task tags.
	// Each tag will end up in "swarming_tag" buildbucket tag, for example
	// "swarming_tag:builder:release"
	SwarmingTags []string `protobuf:"bytes,2,rep,name=swarming_tags,json=swarmingTags,proto3" json:"swarming_tags,omitempty"`
	// A requirement for a bot to execute the build.
	//
	// Supports 3 forms:
	// - "<key>:" - exclude the defaults for the key.
	//   Mutually exclusive with other forms.
	// - "<key>:<value>" - require a bot with this dimension.
	//   This is a shortcut for "0:<key>:<value>", see below.
	// - "<expiration_secs>:<key>:<value>" - wait for up to expiration_secs.
	//   for a bot with the dimension.
	//   Supports mutliple values for different keys and expiration_secs.
	//   expiration_secs must be a multiple of 60.
	//
	// When merging a set of dimensions S1 into S2, all dimensions in S1 with a
	// key K replace all dimensions in S2 with K. This logic is used when applying
	// builder mixins and dimensions specified in a build request.
	//
	// If this builder is defined in a bucket, dimension "pool" is defaulted
	// to the name of the bucket. See Bucket message below.
	Dimensions []string `protobuf:"bytes,3,rep,name=dimensions,proto3" json:"dimensions,omitempty"`
	// Specifies that a recipe to run.
	Recipe *Builder_Recipe `protobuf:"bytes,4,opt,name=recipe,proto3" json:"recipe,omitempty"`
	// Swarming task priority.
	Priority uint32 `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
	// Maximum build execution time. Not to be confused with pending time.
	ExecutionTimeoutSecs uint32 `protobuf:"varint,7,opt,name=execution_timeout_secs,json=executionTimeoutSecs,proto3" json:"execution_timeout_secs,omitempty"`
	// Maximum build pending time.
	ExpirationSecs uint32 `protobuf:"varint,20,opt,name=expiration_secs,json=expirationSecs,proto3" json:"expiration_secs,omitempty"`
	// Caches that should be present on the bot.
	Caches []*Builder_CacheEntry `protobuf:"bytes,9,rep,name=caches,proto3" json:"caches,omitempty"`
	// If YES, generate monotonically increasing contiguous numbers for each
	// build, unique within the builder.
	// Note: this limits the build creation rate in this builder to 5 per second.
	BuildNumbers Toggle `protobuf:"varint,16,opt,name=build_numbers,json=buildNumbers,proto3,enum=buildbucket.Toggle" json:"build_numbers,omitempty"`
	// Email of a service account to run the build as or literal 'bot' string to
	// use Swarming bot's account (if available). Passed directly to Swarming.
	// Subject to Swarming's ACLs.
	ServiceAccount string `protobuf:"bytes,12,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
	// If YES, each builder will get extra dimension "builder:<builder name>"
	// added. Default is UNSET.
	//
	// For example, this config
	//
	//   builder {
	//     name: "linux-compiler"
	//     dimension: "builder:linux-compiler"
	//   }
	//
	// is equivalent to this:
	//
	//   builders {
	//     name: "linux-compiler"
	//     auto_builder_dimension: YES
	//   }
	//
	// We've considered providing interpolation like this
	//   builder_defaults {
	//     dimensions: "builder:${BUILDER}"
	//   }
	// (see also http://docs.buildbot.net/0.8.9/manual/cfg-properties.html#interpolate)
	// but are currently against complicating config with this.
	AutoBuilderDimension Toggle `protobuf:"varint,17,opt,name=auto_builder_dimension,json=autoBuilderDimension,proto3,enum=buildbucket.Toggle" json:"auto_builder_dimension,omitempty"`
	// If YES, by default a new build in this builder will be marked as
	// experimental.
	// This is useful for inherently experimental builders that use production
	// recipes.
	// See also luci_migration_host field.
	Experimental Toggle `protobuf:"varint,18,opt,name=experimental,proto3,enum=buildbucket.Toggle" json:"experimental,omitempty"`
	// If not empty and not "-", and a build request is not marked as
	// experimental/prod explicitly and has "mastername" property, buildbucket
	// will contact this instance of luci-migration app to determine whether the
	// builder is experimental.
	// On success, this takes precedence over Builder.experimental proto field
	// (above).
	//
	// Special value "-" means no luci_migration_host specified.
	// Useful in a builder that wants to override luci_migration_host
	// specified in builder_defaults or mixin.
	LuciMigrationHost    string   `protobuf:"bytes,19,opt,name=luci_migration_host,json=luciMigrationHost,proto3" json:"luci_migration_host,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Builder) Reset()         { *m = Builder{} }
func (m *Builder) String() string { return proto.CompactTextString(m) }
func (*Builder) ProtoMessage()    {}
func (*Builder) Descriptor() ([]byte, []int) {
	return fileDescriptor_90cd2cb373fa8af0, []int{2}
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

func (m *Builder) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Builder) GetMixins() []string {
	if m != nil {
		return m.Mixins
	}
	return nil
}

func (m *Builder) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Builder) GetSwarmingTags() []string {
	if m != nil {
		return m.SwarmingTags
	}
	return nil
}

func (m *Builder) GetDimensions() []string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *Builder) GetRecipe() *Builder_Recipe {
	if m != nil {
		return m.Recipe
	}
	return nil
}

func (m *Builder) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Builder) GetExecutionTimeoutSecs() uint32 {
	if m != nil {
		return m.ExecutionTimeoutSecs
	}
	return 0
}

func (m *Builder) GetExpirationSecs() uint32 {
	if m != nil {
		return m.ExpirationSecs
	}
	return 0
}

func (m *Builder) GetCaches() []*Builder_CacheEntry {
	if m != nil {
		return m.Caches
	}
	return nil
}

func (m *Builder) GetBuildNumbers() Toggle {
	if m != nil {
		return m.BuildNumbers
	}
	return Toggle_UNSET
}

func (m *Builder) GetServiceAccount() string {
	if m != nil {
		return m.ServiceAccount
	}
	return ""
}

func (m *Builder) GetAutoBuilderDimension() Toggle {
	if m != nil {
		return m.AutoBuilderDimension
	}
	return Toggle_UNSET
}

func (m *Builder) GetExperimental() Toggle {
	if m != nil {
		return m.Experimental
	}
	return Toggle_UNSET
}

func (m *Builder) GetLuciMigrationHost() string {
	if m != nil {
		return m.LuciMigrationHost
	}
	return ""
}

// Describes a cache directory persisted on a bot.
//
// If a build requested a cache, a cache directory is available on build
// startup. If the cache was present on the bot, the directory contains files
// from the previous run.
// The build can read/write to the cache directory while it runs.
// After build completes, the cache directory is persisted.
// Next time another build requests the same cache and runs on the same bot,
// if the cache wasn't evicted, the files will still be there.
//
// One bot can keep multiple caches at the same time and one build can request
// multiple different caches.
// A cache is identified by its name and mapped to a path.
// In recipes-based builds, the path is relative to api.paths['cache'] dir.
// For example, a cache {"name": "foo", "path": "bar"} maps foo to bar and
// the cache dir is available at
//   my_cache = api.path['cache'].join('bar')
//
// If the bot is running out of space, caches are evicted in LRU manner.
//
// Renaming a cache is equivalent to clearing it from the builder perspective.
// The files will still be there, but eventually will be purged by GC.
//
// Builder cache.
//
// Buildbucket implicitly declares cache
//   {"name": "<hash(bucket/builder)>", "path": "builder"}.
// This means that any LUCI builder has a "personal disk space" on the bot.
// Builder cache is often a good start before customizing caching.
// In recipes, it is available at api.path['cache'].join('builder').
//
// In order to share the builder cache among multiple builders, it can be
// overridden:
//
//   builders {
//     name: "a"
//     caches {
//       path: "builder"
//       name: "my_shared_cache"
//     }
//   }
//   builders {
//     name: "b"
//     caches {
//       path: "builder"
//       name: "my_shared_cache"
//     }
//   }
//
// Builders "a" and "b" share their builder cache. If an "a" build ran on a bot
// and left some files in the builder cache and then a "b" build runs on the
// same bot, the same files will be available in the builder cache.
type Builder_CacheEntry struct {
	// Identifier of the cache. Required. Length is limited to 128.
	//
	// If the pool of swarming bots is shared among multiple LUCI projects and
	// projects use same cache name, the cache will be shared across projects.
	// To avoid affecting and being affected by other projects, prefix the cache
	// name with something project-specific, e.g. "v8-".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Relative path where the cache in mapped into. Required.
	// Must use POSIX format (forward slashes).
	// In most cases, it does not need slashes at all.
	// Must be unique in the given builder.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Number of seconds to wait for a bot with a warm cache to pick up the
	// task, before falling back to a bot with a cold (non-existent) cache.
	//
	// The default is 0, which means that no preference will be chosen for a bot
	// with this or without this cache, and a bot without this cache may be
	// chosen instead.
	//
	// If no bot has this cache warm, the task will skip this wait and will
	// immediately fallback to a cold cache request.
	//
	// The value must be multiples of 60 seconds.
	WaitForWarmCacheSecs int32    `protobuf:"varint,3,opt,name=wait_for_warm_cache_secs,json=waitForWarmCacheSecs,proto3" json:"wait_for_warm_cache_secs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Builder_CacheEntry) Reset()         { *m = Builder_CacheEntry{} }
func (m *Builder_CacheEntry) String() string { return proto.CompactTextString(m) }
func (*Builder_CacheEntry) ProtoMessage()    {}
func (*Builder_CacheEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_90cd2cb373fa8af0, []int{2, 0}
}

func (m *Builder_CacheEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Builder_CacheEntry.Unmarshal(m, b)
}
func (m *Builder_CacheEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Builder_CacheEntry.Marshal(b, m, deterministic)
}
func (m *Builder_CacheEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Builder_CacheEntry.Merge(m, src)
}
func (m *Builder_CacheEntry) XXX_Size() int {
	return xxx_messageInfo_Builder_CacheEntry.Size(m)
}
func (m *Builder_CacheEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_Builder_CacheEntry.DiscardUnknown(m)
}

var xxx_messageInfo_Builder_CacheEntry proto.InternalMessageInfo

func (m *Builder_CacheEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Builder_CacheEntry) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Builder_CacheEntry) GetWaitForWarmCacheSecs() int32 {
	if m != nil {
		return m.WaitForWarmCacheSecs
	}
	return 0
}

// Specifies a recipe to run.
type Builder_Recipe struct {
	// Repository URL of the recipe package.
	// Value "-" is treated as unset.
	Repository string `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	// Name of the recipe to run.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// If set, will use CIPD to fetch the recipes, rather than using git. This
	// obsoletes the `repository` parameter.
	//
	// Typically the package will look like:
	//
	//   infra/recipe_bundles/chromium.googlesource.com/chromium/tools/build
	//
	// Recipes bundled from internal repositories are typically under
	// `infra_internal/recipe_bundles/...`.
	//
	// But if you're building your own recipe bundles, they could be located
	// elsewhere.
	CipdPackage string `protobuf:"bytes,6,opt,name=cipd_package,json=cipdPackage,proto3" json:"cipd_package,omitempty"`
	// The CIPD version to fetch. This can be a lower-cased git ref (like
	// `refs/heads/master` or `head`), or it can be a cipd tag (like
	// `git_revision:dead...beef`).
	//
	// The default is `head`, which corresponds to the git repo's HEAD ref. This
	// is typically (but not always) a symbolic ref for `refs/heads/master`.
	CipdVersion string `protobuf:"bytes,5,opt,name=cipd_version,json=cipdVersion,proto3" json:"cipd_version,omitempty"`
	// Colon-separated build properties to set.
	// A property can be overriden by "properties" build parameter.
	//
	// Use this field for string properties and use properties_j for other
	// types.
	Properties []string `protobuf:"bytes,3,rep,name=properties,proto3" json:"properties,omitempty"`
	// Same as properties, but the value must valid JSON. For example
	//   properties_j: "a:1"
	// means property a is a number 1, not string "1".
	//
	// If null, it means no property must be defined. In particular, it removes
	// a default value for the property, if any.
	//
	// Fields properties and properties_j can be used together, but cannot both
	// specify values for same property.
	PropertiesJ          []string `protobuf:"bytes,4,rep,name=properties_j,json=propertiesJ,proto3" json:"properties_j,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Builder_Recipe) Reset()         { *m = Builder_Recipe{} }
func (m *Builder_Recipe) String() string { return proto.CompactTextString(m) }
func (*Builder_Recipe) ProtoMessage()    {}
func (*Builder_Recipe) Descriptor() ([]byte, []int) {
	return fileDescriptor_90cd2cb373fa8af0, []int{2, 1}
}

func (m *Builder_Recipe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Builder_Recipe.Unmarshal(m, b)
}
func (m *Builder_Recipe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Builder_Recipe.Marshal(b, m, deterministic)
}
func (m *Builder_Recipe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Builder_Recipe.Merge(m, src)
}
func (m *Builder_Recipe) XXX_Size() int {
	return xxx_messageInfo_Builder_Recipe.Size(m)
}
func (m *Builder_Recipe) XXX_DiscardUnknown() {
	xxx_messageInfo_Builder_Recipe.DiscardUnknown(m)
}

var xxx_messageInfo_Builder_Recipe proto.InternalMessageInfo

func (m *Builder_Recipe) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *Builder_Recipe) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Builder_Recipe) GetCipdPackage() string {
	if m != nil {
		return m.CipdPackage
	}
	return ""
}

func (m *Builder_Recipe) GetCipdVersion() string {
	if m != nil {
		return m.CipdVersion
	}
	return ""
}

func (m *Builder_Recipe) GetProperties() []string {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Builder_Recipe) GetPropertiesJ() []string {
	if m != nil {
		return m.PropertiesJ
	}
	return nil
}

// Configuration of buildbucket-swarming integration for one bucket.
type Swarming struct {
	// Hostname of the swarming instance, e.g. "chromium-swarm.appspot.com".
	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// DEPRECATED, IGNORED.
	// Used to generate a URL for Build, may contain parameters
	// {swarming_hostname}, {task_id}, {bucket} and {builder}. Defaults to:
	// https://{swarming_hostname}/user/task/{task_id}
	UrlFormat string `protobuf:"bytes,2,opt,name=url_format,json=urlFormat,proto3" json:"url_format,omitempty"`
	// Defines default values for builders.
	BuilderDefaults *Builder `protobuf:"bytes,3,opt,name=builder_defaults,json=builderDefaults,proto3" json:"builder_defaults,omitempty"`
	// Configuration for each builder.
	// Swarming tasks are created only for builds for builders that are not
	// explicitly specified.
	Builders []*Builder `protobuf:"bytes,4,rep,name=builders,proto3" json:"builders,omitempty"`
	// Percentage of builds that should use a canary swarming task template.
	// A value from 0 to 100.
	// If omitted, default percentage is used.
	TaskTemplateCanaryPercentage *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=task_template_canary_percentage,json=taskTemplateCanaryPercentage,proto3" json:"task_template_canary_percentage,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}              `json:"-"`
	XXX_unrecognized             []byte                `json:"-"`
	XXX_sizecache                int32                 `json:"-"`
}

func (m *Swarming) Reset()         { *m = Swarming{} }
func (m *Swarming) String() string { return proto.CompactTextString(m) }
func (*Swarming) ProtoMessage()    {}
func (*Swarming) Descriptor() ([]byte, []int) {
	return fileDescriptor_90cd2cb373fa8af0, []int{3}
}

func (m *Swarming) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Swarming.Unmarshal(m, b)
}
func (m *Swarming) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Swarming.Marshal(b, m, deterministic)
}
func (m *Swarming) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Swarming.Merge(m, src)
}
func (m *Swarming) XXX_Size() int {
	return xxx_messageInfo_Swarming.Size(m)
}
func (m *Swarming) XXX_DiscardUnknown() {
	xxx_messageInfo_Swarming.DiscardUnknown(m)
}

var xxx_messageInfo_Swarming proto.InternalMessageInfo

func (m *Swarming) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Swarming) GetUrlFormat() string {
	if m != nil {
		return m.UrlFormat
	}
	return ""
}

func (m *Swarming) GetBuilderDefaults() *Builder {
	if m != nil {
		return m.BuilderDefaults
	}
	return nil
}

func (m *Swarming) GetBuilders() []*Builder {
	if m != nil {
		return m.Builders
	}
	return nil
}

func (m *Swarming) GetTaskTemplateCanaryPercentage() *wrappers.UInt32Value {
	if m != nil {
		return m.TaskTemplateCanaryPercentage
	}
	return nil
}

// Defines one bucket in buildbucket.cfg
type Bucket struct {
	// Name of the bucket. Names are unique within one instance of buildbucket.
	// If another project already uses this name, a config will be rejected.
	// Name reservation is first-come first-serve.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// List of access control rules for the bucket.
	// The order does not matter.
	Acls []*Acl `protobuf:"bytes,2,rep,name=acls,proto3" json:"acls,omitempty"`
	// A list of ACL set names. Each ACL in each referenced ACL set will be
	// included in this bucket.
	// The order does not matter.
	AclSets []string `protobuf:"bytes,4,rep,name=acl_sets,json=aclSets,proto3" json:"acl_sets,omitempty"`
	// Buildbucket-swarming integration.
	Swarming             *Swarming `protobuf:"bytes,3,opt,name=swarming,proto3" json:"swarming,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Bucket) Reset()         { *m = Bucket{} }
func (m *Bucket) String() string { return proto.CompactTextString(m) }
func (*Bucket) ProtoMessage()    {}
func (*Bucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_90cd2cb373fa8af0, []int{4}
}

func (m *Bucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bucket.Unmarshal(m, b)
}
func (m *Bucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bucket.Marshal(b, m, deterministic)
}
func (m *Bucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bucket.Merge(m, src)
}
func (m *Bucket) XXX_Size() int {
	return xxx_messageInfo_Bucket.Size(m)
}
func (m *Bucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Bucket.DiscardUnknown(m)
}

var xxx_messageInfo_Bucket proto.InternalMessageInfo

func (m *Bucket) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Bucket) GetAcls() []*Acl {
	if m != nil {
		return m.Acls
	}
	return nil
}

func (m *Bucket) GetAclSets() []string {
	if m != nil {
		return m.AclSets
	}
	return nil
}

func (m *Bucket) GetSwarming() *Swarming {
	if m != nil {
		return m.Swarming
	}
	return nil
}

// Schema of buildbucket.cfg file, a project config.
type BuildbucketCfg struct {
	// All buckets defined for this project.
	Buckets []*Bucket `protobuf:"bytes,1,rep,name=buckets,proto3" json:"buckets,omitempty"`
	// A list of ACL sets. Names must be unique.
	AclSets []*AclSet `protobuf:"bytes,2,rep,name=acl_sets,json=aclSets,proto3" json:"acl_sets,omitempty"`
	// A list of builder mixin definitions.
	// A mixin can be referenced in any builder defined within the BuildbucketCfg.
	// See also Buider.mixins field.
	BuilderMixins        []*Builder `protobuf:"bytes,3,rep,name=builder_mixins,json=builderMixins,proto3" json:"builder_mixins,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BuildbucketCfg) Reset()         { *m = BuildbucketCfg{} }
func (m *BuildbucketCfg) String() string { return proto.CompactTextString(m) }
func (*BuildbucketCfg) ProtoMessage()    {}
func (*BuildbucketCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_90cd2cb373fa8af0, []int{5}
}

func (m *BuildbucketCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildbucketCfg.Unmarshal(m, b)
}
func (m *BuildbucketCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildbucketCfg.Marshal(b, m, deterministic)
}
func (m *BuildbucketCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildbucketCfg.Merge(m, src)
}
func (m *BuildbucketCfg) XXX_Size() int {
	return xxx_messageInfo_BuildbucketCfg.Size(m)
}
func (m *BuildbucketCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildbucketCfg.DiscardUnknown(m)
}

var xxx_messageInfo_BuildbucketCfg proto.InternalMessageInfo

func (m *BuildbucketCfg) GetBuckets() []*Bucket {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func (m *BuildbucketCfg) GetAclSets() []*AclSet {
	if m != nil {
		return m.AclSets
	}
	return nil
}

func (m *BuildbucketCfg) GetBuilderMixins() []*Builder {
	if m != nil {
		return m.BuilderMixins
	}
	return nil
}

func init() {
	proto.RegisterEnum("buildbucket.Toggle", Toggle_name, Toggle_value)
	proto.RegisterEnum("buildbucket.Acl_Role", Acl_Role_name, Acl_Role_value)
	proto.RegisterType((*Acl)(nil), "buildbucket.Acl")
	proto.RegisterType((*AclSet)(nil), "buildbucket.AclSet")
	proto.RegisterType((*Builder)(nil), "buildbucket.Builder")
	proto.RegisterType((*Builder_CacheEntry)(nil), "buildbucket.Builder.CacheEntry")
	proto.RegisterType((*Builder_Recipe)(nil), "buildbucket.Builder.Recipe")
	proto.RegisterType((*Swarming)(nil), "buildbucket.Swarming")
	proto.RegisterType((*Bucket)(nil), "buildbucket.Bucket")
	proto.RegisterType((*BuildbucketCfg)(nil), "buildbucket.BuildbucketCfg")
}

func init() {
	proto.RegisterFile("github.com/TriggerMail/luci-go/buildbucket/proto/config/project_config.proto", fileDescriptor_90cd2cb373fa8af0)
}

var fileDescriptor_90cd2cb373fa8af0 = []byte{
	// 1000 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xed, 0x6e, 0xdb, 0x36,
	0x14, 0xad, 0x6c, 0xc7, 0x1f, 0xd7, 0x1f, 0x55, 0x19, 0x2f, 0xd0, 0xb2, 0xae, 0xf5, 0xbc, 0x02,
	0xf3, 0x06, 0x54, 0xd9, 0x9c, 0xa1, 0x2d, 0xd0, 0x1f, 0x43, 0x9c, 0xb8, 0x68, 0x82, 0x35, 0x2b,
	0x64, 0xa7, 0xc5, 0xf6, 0x47, 0xa0, 0x19, 0x46, 0x61, 0x23, 0x89, 0x02, 0x49, 0x35, 0xc9, 0x5b,
	0x6c, 0x0f, 0xb0, 0x37, 0xd8, 0x6b, 0xec, 0x4d, 0xf6, 0x20, 0x03, 0x49, 0x59, 0x76, 0x8a, 0xb4,
	0xc0, 0xf6, 0xcb, 0xbe, 0xe7, 0x1e, 0x92, 0x87, 0xe7, 0xde, 0x4b, 0xc1, 0x24, 0xe2, 0x3e, 0x39,
	0x17, 0x3c, 0x61, 0x79, 0xe2, 0x73, 0x11, 0xed, 0xc4, 0x39, 0x61, 0x3b, 0x8b, 0x9c, 0xc5, 0xa7,
	0x8b, 0x9c, 0x5c, 0x50, 0xb5, 0x93, 0x09, 0xae, 0xf8, 0x0e, 0xe1, 0xe9, 0x19, 0x8b, 0x74, 0xf0,
	0x8e, 0x12, 0x15, 0xda, 0xd0, 0x37, 0x39, 0xd4, 0x5e, 0xa3, 0x6f, 0x3f, 0x88, 0x38, 0x8f, 0x62,
	0x6a, 0x97, 0x2d, 0xf2, 0xb3, 0x9d, 0x4b, 0x81, 0xb3, 0x8c, 0x0a, 0x69, 0xc9, 0xc3, 0x3f, 0x1c,
	0xa8, 0xee, 0x91, 0x18, 0x7d, 0x0b, 0x35, 0xc1, 0x63, 0xea, 0x39, 0x03, 0x67, 0xd4, 0x1b, 0x7f,
	0xe6, 0xaf, 0xed, 0xe1, 0xef, 0x91, 0xd8, 0x0f, 0x78, 0x4c, 0x03, 0x43, 0x41, 0x7d, 0xd8, 0x88,
	0x04, 0xcf, 0x33, 0xaf, 0x32, 0x70, 0x46, 0xad, 0xc0, 0x06, 0x68, 0x1b, 0x9a, 0xec, 0x94, 0xa6,
	0x8a, 0xa9, 0x6b, 0xaf, 0x6a, 0x12, 0x65, 0x3c, 0x7c, 0x0c, 0x35, 0xbd, 0x1e, 0x01, 0xd4, 0x83,
	0xe9, 0xde, 0xc1, 0x34, 0x70, 0xef, 0xa0, 0x2e, 0xb4, 0x66, 0xfb, 0x2f, 0xa7, 0x07, 0x27, 0x3f,
	0x4f, 0x03, 0xd7, 0xd1, 0xa9, 0xb7, 0xc1, 0xe1, 0x7c, 0x1a, 0xb8, 0x95, 0xe1, 0x04, 0xea, 0x7b,
	0x24, 0x9e, 0x51, 0x85, 0x10, 0xd4, 0x52, 0x9c, 0x58, 0x55, 0xad, 0xc0, 0xfc, 0x47, 0x8f, 0xa0,
	0x86, 0x49, 0x2c, 0xbd, 0xca, 0xa0, 0x3a, 0x6a, 0x8f, 0xdd, 0x0f, 0x95, 0x06, 0x26, 0x3b, 0xfc,
	0xa7, 0x01, 0x8d, 0x89, 0xce, 0x50, 0x71, 0xeb, 0x2e, 0x5b, 0x50, 0x4f, 0xd8, 0x15, 0x4b, 0xa5,
	0x07, 0x83, 0xea, 0xa8, 0x15, 0x14, 0x91, 0xbe, 0x06, 0xc1, 0x8a, 0x46, 0x5c, 0x5c, 0x7b, 0x75,
	0x7b, 0x8d, 0x65, 0x8c, 0xbe, 0x86, 0xae, 0xbc, 0xc4, 0x22, 0x61, 0x69, 0x14, 0x2a, 0x1c, 0x59,
	0x09, 0xad, 0xa0, 0xb3, 0x04, 0xe7, 0x38, 0x92, 0xe8, 0x01, 0xc0, 0x29, 0x4b, 0x68, 0x2a, 0x19,
	0x4f, 0xa5, 0x57, 0x35, 0x8c, 0x35, 0x04, 0xed, 0x42, 0x5d, 0x50, 0xc2, 0x32, 0xea, 0xd5, 0x06,
	0xce, 0xa8, 0x3d, 0xfe, 0xe2, 0xc6, 0x05, 0x0a, 0xc9, 0x7e, 0x60, 0x28, 0x41, 0x41, 0xd5, 0xaa,
	0x32, 0xc1, 0xb8, 0xd0, 0xe6, 0x6e, 0x0c, 0x9c, 0x51, 0x37, 0x28, 0x63, 0xf4, 0x23, 0x6c, 0xd1,
	0x2b, 0x4a, 0x72, 0xc5, 0x78, 0x1a, 0x2a, 0x96, 0x50, 0x9e, 0xab, 0x50, 0x52, 0x22, 0xbd, 0x86,
	0x61, 0xf6, 0xcb, 0xec, 0xdc, 0x26, 0x67, 0x94, 0x48, 0xf4, 0x0d, 0xdc, 0xa5, 0x57, 0x19, 0x13,
	0xd8, 0x2c, 0x33, 0xf4, 0xbe, 0xa1, 0xf7, 0x56, 0xb0, 0x21, 0x3e, 0x85, 0x3a, 0xc1, 0xe4, 0x9c,
	0x4a, 0xaf, 0x65, 0x0c, 0x7f, 0x78, 0xab, 0xde, 0x7d, 0x4d, 0x99, 0xa6, 0x4a, 0x5c, 0x07, 0x05,
	0x1d, 0x3d, 0x83, 0xae, 0x61, 0x86, 0x69, 0x9e, 0x2c, 0xa8, 0x90, 0x9e, 0x6b, 0x5a, 0x6b, 0xf3,
	0xc6, 0xfa, 0x39, 0x8f, 0xa2, 0x98, 0x06, 0x1d, 0x83, 0x1d, 0x5b, 0xa2, 0xd6, 0x26, 0xa9, 0x78,
	0xcf, 0x08, 0x0d, 0x31, 0x21, 0x3c, 0x4f, 0x95, 0xd7, 0x31, 0xa5, 0xe8, 0x15, 0xf0, 0x9e, 0x45,
	0xd1, 0x21, 0x6c, 0xe1, 0x5c, 0xf1, 0x70, 0x61, 0x55, 0x84, 0xa5, 0xcd, 0xde, 0xbd, 0x8f, 0x9f,
	0xd5, 0xd7, 0x4b, 0x0a, 0xdd, 0x07, 0xcb, 0x05, 0xe8, 0x29, 0x74, 0xe8, 0x55, 0x46, 0x85, 0x06,
	0x14, 0x8e, 0x3d, 0xf4, 0x09, 0xb1, 0xeb, 0x44, 0xe4, 0xc3, 0xa6, 0x9e, 0xd1, 0x30, 0x61, 0x51,
	0x61, 0xe6, 0x39, 0x97, 0xca, 0xdb, 0x34, 0x82, 0xef, 0xe9, 0xd4, 0xab, 0x65, 0xe6, 0x25, 0x97,
	0x6a, 0x3b, 0x06, 0x58, 0x99, 0x75, 0x6b, 0x6b, 0x22, 0xa8, 0x65, 0x58, 0x9d, 0x17, 0xe3, 0x65,
	0xfe, 0xa3, 0x27, 0xe0, 0x5d, 0x62, 0xa6, 0xc2, 0x33, 0x2e, 0x42, 0xdd, 0x6d, 0xa1, 0x31, 0xd9,
	0xd6, 0x4d, 0x4f, 0xdb, 0x46, 0xd0, 0xd7, 0xf9, 0x17, 0x5c, 0xbc, 0xc5, 0x22, 0x31, 0x07, 0xe8,
	0xea, 0x6d, 0xff, 0xed, 0x40, 0xdd, 0xf6, 0x92, 0x6e, 0x4c, 0x41, 0x33, 0x2e, 0x99, 0xd2, 0xbd,
	0x6d, 0x0f, 0x5c, 0x43, 0x4a, 0x29, 0x95, 0x35, 0x29, 0x5f, 0x41, 0x87, 0xb0, 0xec, 0x34, 0xcc,
	0x30, 0xb9, 0xc0, 0x11, 0x2d, 0x26, 0xa2, 0xad, 0xb1, 0xd7, 0x16, 0x2a, 0x29, 0xef, 0xa9, 0x30,
	0xce, 0x6f, 0xac, 0x28, 0x6f, 0x2c, 0xa4, 0x4f, 0xce, 0x04, 0xcf, 0xa8, 0x50, 0x8c, 0x96, 0x23,
	0xb1, 0x42, 0xf4, 0x16, 0xab, 0x28, 0x7c, 0xe7, 0xd5, 0x0c, 0xa3, 0xbd, 0xc2, 0x8e, 0x8e, 0x6a,
	0xcd, 0xa6, 0xdb, 0x3a, 0xaa, 0x35, 0xdb, 0x6e, 0xe7, 0xa8, 0xd6, 0xec, 0xba, 0xbd, 0xa3, 0x5a,
	0xf3, 0xae, 0xeb, 0x0e, 0xff, 0xac, 0x40, 0x73, 0x56, 0x8c, 0x9f, 0x9e, 0x12, 0xed, 0xfd, 0x9a,
	0xa1, 0x65, 0x8c, 0xbe, 0x04, 0xc8, 0x45, 0xac, 0xfd, 0x4b, 0xb0, 0x2a, 0xee, 0xd8, 0xca, 0x45,
	0xfc, 0xc2, 0x00, 0xe8, 0x27, 0x70, 0xcb, 0x26, 0xa2, 0x67, 0x38, 0x8f, 0x95, 0xf5, 0xb5, 0x3d,
	0xee, 0xdf, 0xd6, 0xef, 0xc1, 0xdd, 0x82, 0x7d, 0x50, 0x90, 0xd1, 0xf7, 0xd0, 0x2c, 0x20, 0x69,
	0xf4, 0x7f, 0x6c, 0x61, 0xc9, 0x42, 0x04, 0x1e, 0x2a, 0x2c, 0x2f, 0x42, 0x45, 0x93, 0x2c, 0xc6,
	0x8a, 0x86, 0x04, 0xa7, 0x58, 0x5c, 0x87, 0x19, 0x15, 0x44, 0xb7, 0x56, 0x44, 0x8d, 0x97, 0xed,
	0xf1, 0x7d, 0xdf, 0xbe, 0xe1, 0xfe, 0xf2, 0x0d, 0xf7, 0x4f, 0x0e, 0x53, 0xb5, 0x3b, 0x7e, 0x83,
	0xe3, 0x9c, 0x06, 0xf7, 0xf5, 0x26, 0xf3, 0x62, 0x8f, 0x7d, 0xb3, 0xc5, 0xeb, 0x72, 0x87, 0xe1,
	0xef, 0x0e, 0xd4, 0x27, 0x46, 0xc1, 0xff, 0x7f, 0x4b, 0xd1, 0xe7, 0xd0, 0xc4, 0x24, 0x0e, 0x25,
	0x55, 0xb2, 0xa8, 0x4d, 0x03, 0x9b, 0xf7, 0x59, 0xa2, 0x1f, 0xa0, 0xb9, 0x7c, 0xfd, 0x0a, 0xbf,
	0x6e, 0x7e, 0x3a, 0x96, 0xb5, 0x09, 0x4a, 0xda, 0xf0, 0x2f, 0x07, 0x7a, 0x93, 0x15, 0x65, 0xff,
	0x2c, 0x42, 0x8f, 0xa1, 0x61, 0x03, 0xe9, 0x39, 0x46, 0xc9, 0xe6, 0x07, 0xde, 0xe9, 0x9f, 0x60,
	0xc9, 0x41, 0xfe, 0x9a, 0x9e, 0xca, 0x2d, 0x7c, 0xfb, 0xf1, 0x58, 0x89, 0x7c, 0x0e, 0xbd, 0x65,
	0x71, 0x8b, 0x37, 0xbf, 0xfa, 0x89, 0x0a, 0x75, 0x0b, 0xee, 0x2b, 0x43, 0xfd, 0xee, 0x11, 0xd4,
	0xed, 0xdc, 0xa3, 0x16, 0x6c, 0x9c, 0x1c, 0xcf, 0xa6, 0x73, 0xf7, 0x0e, 0x6a, 0x40, 0xf5, 0xd7,
	0xe9, 0xcc, 0x75, 0x50, 0x1d, 0x2a, 0xc7, 0xbf, 0xb8, 0x95, 0xc9, 0xb3, 0xdf, 0x9e, 0xfc, 0xa7,
	0x2f, 0xf7, 0x73, 0xfb, 0x93, 0x2d, 0x16, 0x75, 0x03, 0xef, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0xe6, 0x99, 0x53, 0x79, 0xfa, 0x07, 0x00, 0x00,
}
