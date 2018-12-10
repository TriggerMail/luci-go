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

package bq

import (
	"testing"

	"cloud.google.com/go/bigquery"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/struct"

	"github.com/TriggerMail/luci-go/common/bq/testdata"
	"github.com/TriggerMail/luci-go/common/clock/testclock"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMetric(t *testing.T) {
	t.Parallel()
	u := Uploader{}
	u.UploadsMetricName = "fakeCounter"
	Convey("Test metric creation", t, func() {
		Convey("Expect uploads metric was created", func() {
			_ = u.getCounter() // To actually create the metric
			So(u.uploads.Info().Name, ShouldEqual, "fakeCounter")
		})
	})
}

func TestSave(t *testing.T) {
	Convey("TestSave", t, func() {
		recentTime := testclock.TestRecentTimeUTC
		ts, err := ptypes.TimestampProto(recentTime)
		So(err, ShouldBeNil)

		r := &Row{
			Message: &testdata.TestMessage{
				Name:      "testname",
				Timestamp: ts,
				Nested: &testdata.NestedTestMessage{
					Name: "nestedname",
				},
				RepeatedNested: []*testdata.NestedTestMessage{
					{Name: "repeated_one"},
					{Name: "repeated_two"},
				},
				Foo: testdata.TestMessage_Y,
				FooRepeated: []testdata.TestMessage_FOO{
					testdata.TestMessage_Y,
					testdata.TestMessage_X,
				},
				Struct: &structpb.Struct{
					Fields: map[string]*structpb.Value{
						"num": {Kind: &structpb.Value_NumberValue{NumberValue: 1}},
						"str": {Kind: &structpb.Value_StringValue{StringValue: "a"}},
					},
				},

				Empty:   &empty.Empty{},
				Empties: []*empty.Empty{{}, {}},
			},
			InsertID: "testid",
		}
		row, id, err := r.Save()
		So(err, ShouldBeNil)
		So(id, ShouldEqual, "testid")
		So(row, ShouldResemble, map[string]bigquery.Value{
			"name":      "testname",
			"timestamp": recentTime,
			"nested":    map[string]bigquery.Value{"name": "nestedname"},
			"repeated_nested": []interface{}{
				map[string]bigquery.Value{"name": "repeated_one"},
				map[string]bigquery.Value{"name": "repeated_two"},
			},
			"foo":          "Y",
			"foo_repeated": []interface{}{"Y", "X"},
			"struct":       `{"num":1,"str":"a"}`,
		})
	})
}

func TestBatch(t *testing.T) {
	t.Parallel()

	Convey("Test batch", t, func() {
		rowLimit := 2
		rows := make([]*Row, 3)
		for i := 0; i < 3; i++ {
			rows[i] = &Row{}
		}

		want := [][]*Row{
			{{}, {}},
			{{}},
		}
		rowSets := batch(rows, rowLimit)
		So(rowSets, ShouldResemble, want)
	})
}
