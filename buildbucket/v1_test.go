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

package buildbucket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	"github.com/TriggerMail/luci-go/buildbucket/proto"
	v1 "github.com/TriggerMail/luci-go/common/api/buildbucket/buildbucket/v1"

	. "github.com/smartystreets/goconvey/convey"
)

func TestV1(t *testing.T) {
	t.Parallel()

	Convey("BuildToV2", t, func() {
		// Read test cases.
		inputBytes, err := ioutil.ReadFile("testdata/v1_builds.json")
		So(err, ShouldBeNil)
		var input []*v1.ApiCommonBuildMessage
		err = json.Unmarshal(inputBytes, &input)
		So(err, ShouldBeNil)

		expectedBytes, err := ioutil.ReadFile("testdata/v2_builds.json")
		So(err, ShouldBeNil)
		var expectedRawMsgs []json.RawMessage
		err = json.Unmarshal(expectedBytes, &expectedRawMsgs)
		So(err, ShouldBeNil)

		So(len(expectedRawMsgs), ShouldEqual, len(input))

		// Make a Convey for each test case.
		for i, msg := range input {
			Convey(fmt.Sprintf("#%d", i), func() {
				var expected buildbucketpb.Build
				err := jsonpb.Unmarshal(bytes.NewReader([]byte(expectedRawMsgs[i])), &expected)
				So(err, ShouldBeNil)

				actual, err := BuildToV2(msg)
				So(err, ShouldBeNil)

				// Clear empty message and/or repeated fields before comparison.
				clearEmptySubmessages(actual)
				clearEmptySubmessages(&expected)

				So(actual, ShouldResemble, &expected)
			})
		}
	})

	Convey("StatusToV2", t, func() {
		cases := map[buildbucketpb.Status]*v1.ApiCommonBuildMessage{
			0: {},

			buildbucketpb.Status_SCHEDULED: {
				Status: "SCHEDULED",
			},

			buildbucketpb.Status_STARTED: {
				Status: "STARTED",
			},

			buildbucketpb.Status_SUCCESS: {
				Status: "COMPLETED",
				Result: "SUCCESS",
			},

			buildbucketpb.Status_FAILURE: {
				Status:        "COMPLETED",
				Result:        "FAILURE",
				FailureReason: "BUILD_FAILURE",
			},

			buildbucketpb.Status_INFRA_FAILURE: {
				Status:        "COMPLETED",
				Result:        "FAILURE",
				FailureReason: "INFRA_FAILURE",
			},

			buildbucketpb.Status_CANCELED: {
				Status:            "COMPLETED",
				Result:            "CANCELED",
				CancelationReason: "CANCELED_EXPLICITLY",
			},
		}
		for expected, build := range cases {
			expected := expected
			build := build
			Convey(expected.String(), func() {
				actual, err := StatusToV2(build)
				So(err, ShouldBeNil)
				So(actual, ShouldEqual, expected)
			})
		}
	})
}

func clearEmptySubmessages(m proto.Message) {
	strct := reflect.ValueOf(m).Elem()
	for _, p := range proto.GetProperties(strct.Type()).Prop {
		f := strct.FieldByName(p.Name)
		ft := f.Type()
		if p.Repeated {
			if f.Len() == 0 {
				f.Set(reflect.Zero(ft))
			}
		} else if subMsg, ok := f.Interface().(proto.Message); ok && !f.IsNil() {
			// this is a submessage
			clearEmptySubmessages(subMsg)
			empty := reflect.New(ft.Elem()).Interface().(proto.Message)
			if proto.Equal(subMsg, empty) {
				// it is an empty message. Zero the field.
				f.Set(reflect.Zero(ft))
			}
		}
	}
}
