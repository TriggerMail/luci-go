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

package skylarkproto

import (
	"reflect"
	"testing"

	"github.com/golang/protobuf/proto"

	"github.com/TriggerMail/luci-go/skylark/skylarkproto/testprotos"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetMessageType(t *testing.T) {
	t.Parallel()

	Convey("Works", t, func() {
		mt, err := GetMessageType(proto.MessageType("testprotos.Complex"))
		So(err, ShouldBeNil)
		So(mt.Name(), ShouldEqual, "testprotos.Complex")
		So(mt.Type(), ShouldEqual, reflect.TypeOf(&testprotos.Complex{}))

		// Getting same type again returns exact same object.
		mt2, _ := GetMessageType(proto.MessageType("testprotos.Complex"))
		So(mt2, ShouldEqual, mt)

		// Discovered all fields.
		So(mt.fieldNames, ShouldResemble, []string{
			"enum_val",
			"i64",
			"i64_rep",
			"inner_msg",
			"msg_val",
			"msg_val_rep",
			"simple",
		})

		// Types and value obtainers for fields are correct.
		msg := &testprotos.Complex{
			EnumVal: testprotos.Complex_ENUM_VAL_1,
			I64:     123,
			I64Rep:  []int64{1, 2, 3},
			MsgVal:  &testprotos.Complex_InnerMessage{I: 456},
		}
		val := reflect.ValueOf(msg).Elem()

		var desc fieldDesc

		// Copy-pasta below to avoid using reflection for testing reflection to reduce
		// chances of making identical self-canceling mistakes in tests and code under
		// test.

		desc = mt.fields["enum_val"]
		So(desc.typ, ShouldEqual, reflect.TypeOf(msg.EnumVal))
		So(desc.onProtoReflection(val).Interface().(testprotos.Complex_InnerEnum), ShouldEqual, msg.EnumVal)

		desc = mt.fields["i64"]
		So(desc.typ, ShouldEqual, reflect.TypeOf(msg.I64))
		So(desc.onProtoReflection(val).Interface().(int64), ShouldEqual, msg.I64)

		desc = mt.fields["i64_rep"]
		So(desc.typ, ShouldEqual, reflect.TypeOf(msg.I64Rep))
		So(desc.onProtoReflection(val).Interface().([]int64), ShouldResemble, msg.I64Rep)

		desc = mt.fields["msg_val"]
		So(desc.typ, ShouldEqual, reflect.TypeOf(msg.MsgVal))
		So(desc.onProtoReflection(val).Interface().(*testprotos.Complex_InnerMessage), ShouldEqual, msg.MsgVal)

		desc = mt.fields["msg_val_rep"]
		So(desc.typ, ShouldEqual, reflect.TypeOf(msg.MsgValRep))
		So(desc.onProtoReflection(val).Interface().([]*testprotos.Complex_InnerMessage), ShouldEqual, msg.MsgValRep)

		// Grabbing a oneof alternative "switches" the wrapper to point to it.
		desc = mt.fields["inner_msg"]
		So(desc.typ, ShouldEqual, reflect.TypeOf(&testprotos.Complex_InnerMessage{}))
		So(desc.onProtoReflection(val).Interface().(*testprotos.Complex_InnerMessage),
			ShouldResemble, (*testprotos.Complex_InnerMessage)(nil))
		So(msg.OneofVal, ShouldHaveSameTypeAs, &testprotos.Complex_InnerMsg{})
	})
}
