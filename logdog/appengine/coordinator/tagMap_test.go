// Copyright 2015 The LUCI Authors.
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

package coordinator

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	ds "go.chromium.org/gae/service/datastore"
	"github.com/TriggerMail/luci-go/common/errors"
	. "github.com/TriggerMail/luci-go/common/testing/assertions"
)

func TestTagMap(t *testing.T) {
	t.Parallel()

	Convey(`Will return nil when encoding an empty TagMap.`, t, func() {
		tm := TagMap{}
		props, err := tm.toProperties()
		So(err, ShouldBeNil)
		So(props, ShouldBeNil)
	})

	Convey(`Will return nil when decoding an empty property set.`, t, func() {
		tm, err := tagMapFromProperties([]ds.Property(nil))
		So(err, ShouldBeNil)
		So(tm, ShouldBeNil)
	})

	Convey(`Will encode tags as both value and presence tags.`, t, func() {
		tm := TagMap{
			"foo":  "bar",
			"baz":  "qux",
			"quux": "",
		}

		props, err := tm.toProperties()
		So(err, ShouldBeNil)
		So(props, ShouldResemble, sps(
			encodeKey("baz"),
			encodeKey("baz=qux"),
			encodeKey("foo"),
			encodeKey("foo=bar"),
			encodeKey("quux"),
			encodeKey("quux="),
		))

		Convey(`Can decode tags.`, func() {
			dtm, err := tagMapFromProperties(props)
			So(err, ShouldBeNil)
			So(dtm, ShouldResemble, tm)
		})

		Convey(`Will decode mixed invalid and valid tags, returning errors for the invalid.`, func() {
			props = append(props, []ds.Property{
				ds.MkProperty(123), // Not a string.
				ds.MkProperty("not a valid encoded key"),
				ds.MkProperty(encodeKey("!!!invalid tag!!!=12")),
			}...)

			dtm, err := tagMapFromProperties(props)
			So(dtm, ShouldResemble, tm)

			So(err, ShouldHaveSameTypeAs, errors.MultiError{})
			me := err.(errors.MultiError)
			So(len(me), ShouldEqual, 9)

			// 0-5 are the valid encoded tags.
			So(me[6], ShouldErrLike, "property is not a string")
			So(me[7], ShouldErrLike, "failed to decode property")
			So(me[8], ShouldErrLike, "invalid tag")
		})
	})
}

func TestTagMapQuery(t *testing.T) {
	t.Parallel()

	Convey(`A Datastore query`, t, func() {
		q := ds.NewQuery("test")

		Convey(`Will request an encoded presence key if no value is specified.`, func() {
			q = AddLogStreamTagFilter(q, "FindThisKey", "")

			fq, err := q.Finalize()
			So(err, ShouldBeNil)
			So(fq.EqFilters(), ShouldResemble, map[string]ds.PropertySlice{
				"_Tags": sps(encodeKey("FindThisKey")),
			})
		})

		Convey(`Will request an encoded key/value.`, func() {
			q = AddLogStreamTagFilter(q, "FindThisKey", "DesiredValue")

			fq, err := q.Finalize()
			So(err, ShouldBeNil)
			So(fq.EqFilters(), ShouldResemble, map[string]ds.PropertySlice{
				"_Tags": sps(encodeKey("FindThisKey=DesiredValue")),
			})
		})
	})
}
