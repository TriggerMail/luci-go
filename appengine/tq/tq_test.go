// Copyright 2017 The LUCI Authors.
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

package tq

import (
	"bytes"
	"context"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/empty"

	"go.chromium.org/gae/impl/memory"
	"go.chromium.org/gae/service/taskqueue"

	"github.com/TriggerMail/luci-go/common/clock"
	"github.com/TriggerMail/luci-go/common/clock/testclock"
	"github.com/TriggerMail/luci-go/common/retry/transient"
	"github.com/TriggerMail/luci-go/server/router"

	. "github.com/smartystreets/goconvey/convey"
)

var epoch = time.Unix(1500000000, 0).UTC()

func TestDispatcher(t *testing.T) {
	t.Parallel()

	Convey("With dispatcher", t, func() {
		ctx := memory.Use(context.Background())
		ctx = clock.Set(ctx, testclock.New(epoch))
		taskqueue.GetTestable(ctx).CreateQueue("another-q")

		d := Dispatcher{}
		r := router.New()

		installRoutes := func() {
			d.InstallRoutes(r, router.NewMiddlewareChain(func(c *router.Context, next router.Handler) {
				c.Context = ctx
				next(c)
			}))
		}
		runTasks := func(ctx context.Context) []int {
			var codes []int
			for _, tasks := range taskqueue.GetTestable(ctx).GetScheduledTasks() {
				for _, task := range tasks {
					// Execute the task.
					req := httptest.NewRequest("POST", "http://example.com"+task.Path, bytes.NewReader(task.Payload))
					rw := httptest.NewRecorder()
					r.ServeHTTP(rw, req)
					codes = append(codes, rw.Code)
				}
			}
			return codes
		}

		Convey("Single task", func() {
			var calls []proto.Message
			handler := func(c context.Context, payload proto.Message) error {
				hdr, err := RequestHeaders(c)
				So(err, ShouldBeNil)
				So(hdr, ShouldResemble, &taskqueue.RequestHeaders{})
				calls = append(calls, payload)
				return nil
			}

			// Abuse some well-known proto type to simplify the test. It's doesn't
			// matter what proto type we use here as long as it is registered in
			// protobuf type registry.
			d.RegisterTask(&duration.Duration{}, handler, "", nil)
			installRoutes()

			err := d.AddTask(ctx, &Task{
				Payload:          &duration.Duration{Seconds: 123},
				DeduplicationKey: "abc",
				NamePrefix:       "prefix",
				Title:            "abc-def",
				Delay:            30 * time.Second,
			})
			So(err, ShouldBeNil)

			// Added the task.
			expectedPath := "/internal/tasks/default/abc-def"
			expectedName := "prefix-afc6f8271b8598ee04e359916e6c584a9bc3c520a11dd5244e3399346ac0d3a7"
			expectedBody := []byte(`{"type":"google.protobuf.Duration","body":"123s"}`)
			tasks := taskqueue.GetTestable(ctx).GetScheduledTasks()
			So(tasks, ShouldResemble, taskqueue.QueueData{
				"default": map[string]*taskqueue.Task{
					expectedName: {
						Path:    expectedPath,
						Payload: expectedBody,
						Name:    expectedName,
						Method:  "POST",
						ETA:     epoch.Add(30 * time.Second),
					},
				},
				"another-q": {},
			})

			// Readd a task with same dedup key. Should be silently ignored.
			err = d.AddTask(ctx, &Task{
				Payload:          &duration.Duration{Seconds: 123},
				DeduplicationKey: "abc",
				NamePrefix:       "prefix",
			})
			So(err, ShouldBeNil)

			// No new tasks.
			tasks = taskqueue.GetTestable(ctx).GetScheduledTasks()
			So(len(tasks["default"]), ShouldResemble, 1)

			Convey("Executed", func() {
				// Execute the task.
				So(runTasks(ctx), ShouldResemble, []int{200})
				So(calls, ShouldResemble, []proto.Message{
					&duration.Duration{Seconds: 123},
				})
			})

			Convey("Deleted", func() {
				So(d.DeleteTask(ctx, &Task{
					Payload:          &duration.Duration{Seconds: 123},
					DeduplicationKey: "abc",
					NamePrefix:       "prefix",
				}), ShouldBeNil)

				// Did not execute any tasks.
				So(runTasks(ctx), ShouldHaveLength, 0)
				So(calls, ShouldHaveLength, 0)
			})
		})

		Convey("Deleting unknown task returns nil", func() {
			handler := func(c context.Context, payload proto.Message) error { return nil }
			d.RegisterTask(&duration.Duration{}, handler, "default", nil)

			So(d.DeleteTask(ctx, &Task{
				Payload:          &duration.Duration{Seconds: 123},
				DeduplicationKey: "something",
			}), ShouldBeNil)
		})

		Convey("Many tasks", func() {
			handler := func(c context.Context, payload proto.Message) error { return nil }
			d.RegisterTask(&duration.Duration{}, handler, "default", nil)
			d.RegisterTask(&empty.Empty{}, handler, "another-q", nil)
			installRoutes()

			t := []*Task{}
			for i := 0; i < 200; i++ {
				var task *Task

				if i%2 == 0 {
					task = &Task{
						DeduplicationKey: fmt.Sprintf("%d", i/2),
						Payload:          &duration.Duration{},
						Delay:            time.Duration(i) * time.Second,
					}
				} else {
					task = &Task{
						DeduplicationKey: fmt.Sprintf("%d", (i-1)/2),
						Payload:          &empty.Empty{},
						Delay:            time.Duration(i) * time.Second,
					}
				}

				t = append(t, task)

				// Mix in some duplicates.
				if i > 0 && i%100 == 0 {
					t = append(t, task)
				}
			}
			err := d.AddTask(ctx, t...)
			So(err, ShouldBeNil)

			// Added all the tasks.
			allTasks := taskqueue.GetTestable(ctx).GetScheduledTasks()
			delaysDefault := map[time.Duration]struct{}{}
			for _, task := range allTasks["default"] {
				delaysDefault[task.ETA.Sub(epoch)/time.Second] = struct{}{}
			}
			delaysAnotherQ := map[time.Duration]struct{}{}
			for _, task := range allTasks["another-q"] {
				delaysAnotherQ[task.ETA.Sub(epoch)/time.Second] = struct{}{}
			}
			So(len(delaysDefault), ShouldEqual, 100)
			So(len(delaysAnotherQ), ShouldEqual, 100)

			// Delete the tasks.
			So(d.DeleteTask(ctx, t...), ShouldBeNil)
			So(runTasks(ctx), ShouldHaveLength, 0)
		})

		Convey("Execution errors", func() {
			var returnErr error
			panicNow := false
			handler := func(c context.Context, payload proto.Message) error {
				if panicNow {
					panic("must not be called")
				}
				return returnErr
			}

			d.RegisterTask(&duration.Duration{}, handler, "", nil)
			installRoutes()

			goodBody := `{"type":"google.protobuf.Duration","body":"123.000s"}`

			execute := func(body string) *httptest.ResponseRecorder {
				req := httptest.NewRequest(
					"POST",
					"http://example.com/internal/tasks/default/abc-def",
					bytes.NewReader([]byte(body)))
				rw := httptest.NewRecorder()
				r.ServeHTTP(rw, req)
				return rw
			}

			// Error conditions inside the task body.

			returnErr = nil
			rw := execute(goodBody)
			So(rw.Code, ShouldEqual, 200)
			So(rw.Body.String(), ShouldEqual, "OK\n")

			returnErr = fmt.Errorf("fatal err")
			rw = execute(goodBody)
			So(rw.Code, ShouldEqual, 202) // no retry!
			So(rw.Body.String(), ShouldEqual, "Fatal error: fatal err\n")

			returnErr = transient.Tag.Apply(fmt.Errorf("transient err"))
			rw = execute(goodBody)
			So(rw.Code, ShouldEqual, 500) // 500 for retry
			So(rw.Body.String(), ShouldEqual, "Transient error: transient err\n")

			// Error conditions when routing the task.

			panicNow = true

			rw = execute("not a json")
			So(rw.Code, ShouldEqual, 202) // no retry!
			So(rw.Body.String(), ShouldStartWith, "Bad payload, can't deserialize")

			rw = execute(`{"type":"google.protobuf.Duration"}`)
			So(rw.Code, ShouldEqual, 202) // no retry!
			So(rw.Body.String(), ShouldStartWith, "Bad payload, can't deserialize")

			rw = execute(`{"type":"google.protobuf.Duration","body":"blah"}`)
			So(rw.Code, ShouldEqual, 202) // no retry!
			So(rw.Body.String(), ShouldStartWith, "Bad payload, can't deserialize")

			rw = execute(`{"type":"unknown.proto.type","body":"{}"}`)
			So(rw.Code, ShouldEqual, 202) // no retry!
			So(rw.Body.String(), ShouldStartWith, "Bad payload, can't deserialize")
		})
	})
}
