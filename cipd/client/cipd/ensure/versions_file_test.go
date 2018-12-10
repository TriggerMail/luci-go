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

package ensure

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/TriggerMail/luci-go/cipd/common"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/TriggerMail/luci-go/common/testing/assertions"
)

func TestVersionsFile(t *testing.T) {
	t.Parallel()

	const (
		iid1 = "11111joOfFfFcq7fHCKAIrU34oeFAT174Bf8eHMajMUC"
		iid2 = "22222joOfFfFcq7fHCKAIrU34oeFAT174Bf8eHMajMUC"
	)

	Convey("Setter/getter", t, func() {
		v := VersionsFile{}

		_, err := v.ResolveVersion("pkg", "ver")
		So(err, ShouldErrLike, "not in the versions file")

		So(v.AddVersion("pkg", "ver", iid1), ShouldBeNil)
		So(v.AddVersion("pkg", iid1, iid1), ShouldBeNil) // noop
		So(v, ShouldHaveLength, 1)

		pin, err := v.ResolveVersion("pkg", "ver")
		So(err, ShouldBeNil)
		So(pin, ShouldResemble, common.Pin{
			PackageName: "pkg",
			InstanceID:  iid1,
		})

		pin, err = v.ResolveVersion("other-pkg", iid1)
		So(err, ShouldBeNil)
		So(pin, ShouldResemble, common.Pin{
			PackageName: "other-pkg",
			InstanceID:  iid1,
		})
	})

	Convey("AddVersion errors", t, func() {
		v := VersionsFile{}

		So(v.AddVersion("???", "ver", iid1), ShouldErrLike, "invalid package name")
		So(v.AddVersion("pkg", "???", iid1), ShouldErrLike, "bad version")
		So(v.AddVersion("pkg", "ver", "not-id"), ShouldErrLike, "not a valid package instance ID")
		So(v.AddVersion("pkg", iid1, iid2), ShouldErrLike, "should resolve into that ID")
	})

	Convey("Equal", t, func() {
		v1 := VersionsFile{
			{"pkg1", "ver1"}: iid1,
			{"pkg1", "ver2"}: iid2,
		}
		v2 := VersionsFile{
			{"pkg1", "ver1"}: iid1,
			{"pkg1", "ver2"}: iid1,
		}
		v3 := VersionsFile{
			{"pkg1", "ver1"}: iid1,
		}

		So(v1.Equal(v1), ShouldBeTrue)
		So(v1.Equal(v2), ShouldBeFalse)
		So(v1.Equal(v3), ShouldBeFalse)
	})

	Convey("Serialization and successful parsing", t, func() {

		testVersion := VersionsFile{
			{"pkg1", "ver1"}:      iid1,
			{"pkg1", "ver2"}:      iid1,
			{"pkg2", "tag:works"}: iid2,
		}

		expectedSerialization := fmt.Sprintf(`# This file is auto-generated by 'cipd ensure-file-resolve'.
# Do not modify manually. All changes will be overwritten.

pkg1
	ver1
	%s

pkg1
	ver2
	%s

pkg2
	tag:works
	%s
`, iid1, iid1, iid2)

		Convey("Serialization", func() {
			b := bytes.Buffer{}
			So(testVersion.Serialize(&b), ShouldBeNil)
			So(b.String(), ShouldEqual, expectedSerialization)
		})

		Convey("Parsing success", func() {
			v, err := ParseVersionsFile(strings.NewReader(expectedSerialization))
			So(err, ShouldBeNil)
			So(v, ShouldResemble, testVersion)
		})

		Convey("Parsing empty", func() {
			v, err := ParseVersionsFile(strings.NewReader(""))
			So(err, ShouldBeNil)
			So(v, ShouldResemble, VersionsFile{})
		})

		Convey("Parsing one", func() {
			v, err := ParseVersionsFile(strings.NewReader(fmt.Sprintf("pkg\nver\n%s", iid1)))
			So(err, ShouldBeNil)
			So(v, ShouldResemble, VersionsFile{
				{"pkg", "ver"}: iid1,
			})
		})

		Convey("Many new lines", func() {
			v, err := ParseVersionsFile(strings.NewReader(
				fmt.Sprintf("pkg\nver1\n%s\n\n\npkg\nver2\n%s", iid1, iid2)))
			So(err, ShouldBeNil)
			So(v, ShouldResemble, VersionsFile{
				{"pkg", "ver1"}: iid1,
				{"pkg", "ver2"}: iid2,
			})
		})
	})

	Convey("Parsing errors", t, func() {
		p := func(text string, args ...interface{}) error {
			v, err := ParseVersionsFile(strings.NewReader(fmt.Sprintf(text, args...)))
			So(err, ShouldNotBeNil)
			So(v, ShouldBeNil)
			return err
		}

		Convey("Bad format of identifiers", func() {
			So(p("???\nver\n%s", iid1), ShouldErrLike,
				"failed to parse versions file (line 1): invalid package name")

			So(p("pkg\n???\n%s", iid1), ShouldErrLike,
				"failed to parse versions file (line 2): bad version")

			So(p("pkg\nver\nnotid"), ShouldErrLike,
				"failed to parse versions file (line 3): not a valid package instance ID")
		})

		Convey("Unexpected empty line", func() {
			So(p("pkg\n\nver\n%s\n", iid1), ShouldErrLike,
				"failed to parse versions file (line 2): expecting a version name, not a new line")

			So(p("pkg\nver\n\n%s\n", iid1), ShouldErrLike,
				"failed to parse versions file (line 3): expecting an instance ID, not a new line")
		})

		Convey("Unexpected EOF", func() {
			So(p("pkg\n"), ShouldErrLike,
				"failed to parse versions file (line 1): unexpected EOF, expecting a package version")

			So(p("pkg\nver\n"), ShouldErrLike,
				"failed to parse versions file (line 2): unexpected EOF, expecting an instance ID")
		})

		Convey("Unexpected line after the triple", func() {
			So(p("pkg\nver\n%s\nsomething", iid1), ShouldErrLike,
				"failed to parse versions file (line 4): expecting an empty line between each version definition triple")
		})
	})
}
