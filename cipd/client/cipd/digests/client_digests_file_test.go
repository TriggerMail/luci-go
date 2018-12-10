// Copyright 2018 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package digests

import (
	"bytes"
	"strings"
	"testing"

	api "github.com/TriggerMail/luci-go/cipd/api/cipd/v1"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/TriggerMail/luci-go/common/testing/assertions"
)

func TestClientDigestsFile(t *testing.T) {
	t.Parallel()

	sha1 := func(char string) *api.ObjectRef {
		return &api.ObjectRef{
			HashAlgo:  api.HashAlgo_SHA1,
			HexDigest: strings.Repeat(char, 40),
		}
	}

	sha256 := func(char string) *api.ObjectRef {
		return &api.ObjectRef{
			HashAlgo:  api.HashAlgo_SHA256,
			HexDigest: strings.Repeat(char, 64),
		}
	}

	Convey("Works", t, func() {
		df := ClientDigestsFile{}

		df.AddClientRef("linux-amd64", sha1("a"))
		df.AddClientRef("linux-amd64", sha256("b"))
		df.AddClientRef("windows-amd64", sha1("c"))

		So(df.ClientRef("linux-amd64"), ShouldResemble, sha256("b"))
		So(df.ClientRef("windows-amd64"), ShouldResemble, sha1("c"))
		So(df.ClientRef("unknown"), ShouldBeNil)

		So(df.Contains("linux-amd64", sha1("a")), ShouldBeTrue)
		So(df.Contains("linux-amd64", sha256("b")), ShouldBeTrue)
		So(df.Contains("linux-amd64", sha1("c")), ShouldBeFalse)
		So(df.Contains("unknown", sha1("a")), ShouldBeFalse)
	})

	Convey("Errors in AddClientRef", t, func() {
		df := ClientDigestsFile{}

		So(df.AddClientRef("linux-amd64", &api.ObjectRef{
			HashAlgo:  12345,
			HexDigest: "aaaa",
		}), ShouldErrLike, "unsupported hash algorithm")

		df.AddClientRef("linux-amd64", sha1("a"))
		So(df.AddClientRef("linux-amd64", sha1("a")), ShouldErrLike, "has already been added")
	})

	Convey("Equal", t, func() {
		df1 := ClientDigestsFile{
			entries: []clientDigestEntry{
				{"a", sha1("a")}, {"b", sha1("b")},
			},
		}
		df2 := ClientDigestsFile{
			entries: []clientDigestEntry{
				{"b", sha1("b")}, {"a", sha1("a")},
			},
		}
		df3 := ClientDigestsFile{
			entries: []clientDigestEntry{
				{"a", sha1("a")},
			},
		}
		df4 := ClientDigestsFile{
			entries: []clientDigestEntry{
				{"a", sha1("b")},
			},
		}

		So(df1.Equal(&df1), ShouldBeTrue)
		So(df1.Equal(&df2), ShouldBeFalse)
		So(df1.Equal(&df3), ShouldBeFalse)
		So(df3.Equal(&df4), ShouldBeFalse)
	})

	Convey("Sort", t, func() {
		df := ClientDigestsFile{
			entries: []clientDigestEntry{
				{"windows-amd64", sha1("a")},
				{"linux-amd64", sha1("b")},
				{"linux-amd64", sha256("c")},
			},
		}
		df.Sort()
		So(df, ShouldResemble, ClientDigestsFile{
			entries: []clientDigestEntry{
				{"linux-amd64", sha256("c")},
				{"linux-amd64", sha1("b")},
				{"windows-amd64", sha1("a")},
			},
		})
	})

	Convey("Serialization", t, func() {
		df := ClientDigestsFile{
			entries: []clientDigestEntry{
				{"windows-amd64", sha1("a")},
				{"linux-amd64", sha1("b")},
				{"linux-amd64", sha256("c")},
			},
		}
		buf := bytes.Buffer{}
		So(df.Serialize(&buf, "tag:val", "cipd_client_version"), ShouldBeNil)
		So(buf.String(), ShouldEqual, strings.Join([]string{
			`# This file was generated by`,
			`#`,
			`#  cipd selfupdate-roll -version-file cipd_client_version \`,
			`#      -version tag:val`,
			`#`,
			`# Do not modify manually. All changes will be overwritten.`,
			`# Use 'cipd selfupdate-roll ...' to modify.`,
			``,
			`windows-amd64  sha1    aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa`,
			`linux-amd64    sha1    bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb`,
			`linux-amd64    sha256  cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc`,
			``,
		}, "\n"))
	})

	Convey("Parsing success", t, func() {
		df, err := ParseClientDigestsFile(strings.NewReader(strings.Join([]string{
			"# Comment",
			"linux-amd64	sha256	cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc",
			"",
			" # Another comment",
			"linux-amd64  unknown  bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
			"linux-amd64	sha1  bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
			"windows-amd64	sha1	aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		}, "\n")))
		So(err, ShouldBeNil)
		So(df, ShouldResemble, &ClientDigestsFile{
			entries: []clientDigestEntry{
				{"linux-amd64", sha256("c")},
				{"linux-amd64", sha1("b")},
				{"windows-amd64", sha1("a")},
			},
		})
	})

	Convey("Parsing errors", t, func() {
		call := func(lines ...string) error {
			df, err := ParseClientDigestsFile(strings.NewReader(strings.Join(lines, "\n")))
			So(err, ShouldNotBeNil)
			So(df, ShouldBeNil)
			return err
		}

		So(call("linux-amd64 sha1"), ShouldErrLike, "must have format")
		So(call("linux-amd64 sha256 aaaaaa"), ShouldErrLike, "invalid SHA256 digest")
		So(call(
			"linux-amd64 sha1 aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			"linux-amd64 sha1 bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
		), ShouldErrLike, "has already been added")
	})
}
