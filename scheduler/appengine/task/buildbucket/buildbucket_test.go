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

package buildbucket

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/protobuf/proto"

	"google.golang.org/api/pubsub/v1"

	"go.chromium.org/gae/impl/memory"
	"github.com/TriggerMail/luci-go/config/validation"
	api "github.com/TriggerMail/luci-go/scheduler/api/scheduler/v1"
	"github.com/TriggerMail/luci-go/scheduler/appengine/internal"
	"github.com/TriggerMail/luci-go/scheduler/appengine/messages"
	"github.com/TriggerMail/luci-go/scheduler/appengine/task"
	"github.com/TriggerMail/luci-go/scheduler/appengine/task/utils/tasktest"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/TriggerMail/luci-go/common/testing/assertions"
)

var _ task.Manager = (*TaskManager)(nil)

func TestValidateProtoMessage(t *testing.T) {
	t.Parallel()

	tm := TaskManager{}
	c := context.Background()

	Convey("ValidateProtoMessage works", t, func() {
		ctx := &validation.Context{Context: c}
		validate := func(msg proto.Message) error {
			tm.ValidateProtoMessage(ctx, msg)
			return ctx.Finalize()
		}

		Convey("ValidateProtoMessage passes good msg", func() {
			So(validate(&messages.BuildbucketTask{
				Server:     "blah.com",
				Bucket:     "bucket",
				Builder:    "builder",
				Tags:       []string{"a:b", "c:d"},
				Properties: []string{"a:b", "c:d"},
			}), ShouldBeNil)
		})

		Convey("ValidateProtoMessage passes good minimal msg", func() {
			So(validate(&messages.BuildbucketTask{
				Server:  "blah.com",
				Bucket:  "bucket",
				Builder: "builder",
			}), ShouldBeNil)
		})

		Convey("ValidateProtoMessage wrong type", func() {
			So(validate(&messages.NoopTask{}), ShouldErrLike, "wrong type")
		})

		Convey("ValidateProtoMessage empty", func() {
			So(validate(tm.ProtoMessageType()), ShouldErrLike, "expecting a non-empty BuildbucketTask")
		})

		Convey("ValidateProtoMessage validates URL", func() {
			call := func(url string) error {
				ctx = &validation.Context{Context: c}
				tm.ValidateProtoMessage(ctx, &messages.BuildbucketTask{
					Server:  url,
					Bucket:  "bucket",
					Builder: "builder",
				})
				return ctx.Finalize()
			}
			So(call(""), ShouldErrLike, "field 'server' is required")
			So(call("https://host/not-root"), ShouldErrLike, "field 'server' should be just a host, not a URL")
			So(call("%%%%"), ShouldErrLike, "field 'server' is not a valid hostname")
			So(call("blah.com/abc"), ShouldErrLike, "field 'server' is not a valid hostname")
		})

		Convey("ValidateProtoMessage needs bucket", func() {
			So(validate(&messages.BuildbucketTask{
				Server:  "blah.com",
				Builder: "builder",
			}), ShouldErrLike, "'bucket' field is required")
		})

		Convey("ValidateProtoMessage needs builder", func() {
			So(validate(&messages.BuildbucketTask{
				Server: "blah.com",
				Bucket: "bucket",
			}), ShouldErrLike, "'builder' field is required")
		})

		Convey("ValidateProtoMessage validates properties", func() {
			So(validate(&messages.BuildbucketTask{
				Server:     "blah.com",
				Bucket:     "bucket",
				Builder:    "builder",
				Properties: []string{"not_kv_pair"},
			}), ShouldErrLike, "bad property, not a 'key:value' pair")
		})

		Convey("ValidateProtoMessage validates tags", func() {
			So(validate(&messages.BuildbucketTask{
				Server:  "blah.com",
				Bucket:  "bucket",
				Builder: "builder",
				Tags:    []string{"not_kv_pair"},
			}), ShouldErrLike, "bad tag, not a 'key:value' pair")
		})

		Convey("ValidateProtoMessage forbids default tags overwrite", func() {
			So(validate(&messages.BuildbucketTask{
				Server:  "blah.com",
				Bucket:  "bucket",
				Builder: "builder",
				Tags:    []string{"scheduler_job_id:blah"},
			}), ShouldErrLike, "tag \"scheduler_job_id\" is reserved")
		})
	})
}

func fakeController(testSrvURL string) *tasktest.TestController {
	return &tasktest.TestController{
		TaskMessage: &messages.BuildbucketTask{
			Server:  testSrvURL,
			Bucket:  "test-bucket",
			Builder: "builder",
			Tags:    []string{"a:b", "c:d"},
		},
		Client:       http.DefaultClient,
		SaveCallback: func() error { return nil },
		PrepareTopicCallback: func(publisher string) (string, string, error) {
			if publisher != testSrvURL {
				panic(fmt.Sprintf("expecting %q, got %q", testSrvURL, publisher))
			}
			return "topic", "auth_token", nil
		},
	}
}

func TestFullFlow(t *testing.T) {
	t.Parallel()

	Convey("LaunchTask and HandleNotification work", t, func(ctx C) {
		mockRunning := true

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			resp := ""
			switch {
			case r.Method == "PUT" && r.URL.Path == "/_ah/api/buildbucket/v1/builds":
				// There's more stuff in actual response that we don't use.
				resp = `{
					"build": {
						"id": "9025781602559305888",
						"status": "STARTED",
						"url": "https://chromium-swarm-dev.appspot.com/user/task/2bdfb7404d18ac10"
					}
				}`
			case r.Method == "GET" && r.URL.Path == "/_ah/api/buildbucket/v1/builds/9025781602559305888":
				if mockRunning {
					resp = `{
						"build": {
							"id": "9025781602559305888",
							"status": "STARTED"
						}
					}`
				} else {
					resp = `{
						"build": {
							"id": "9025781602559305888",
							"status": "COMPLETED",
							"result": "SUCCESS"
						}
					}`
				}
			default:
				ctx.Printf("Unknown URL fetch - %s %s\n", r.Method, r.URL.Path)
				w.WriteHeader(400)
				return
			}
			w.WriteHeader(200)
			w.Write([]byte(resp))
		}))
		defer ts.Close()

		c := memory.Use(context.Background())
		mgr := TaskManager{}
		ctl := fakeController(ts.URL)

		// Launch.
		So(mgr.LaunchTask(c, ctl), ShouldBeNil)
		So(ctl.TaskState, ShouldResemble, task.State{
			Status:   task.StatusRunning,
			TaskData: []byte(`{"build_id":"9025781602559305888"}`),
			ViewURL:  "https://chromium-swarm-dev.appspot.com/user/task/2bdfb7404d18ac10",
		})

		// Added the timer.
		So(ctl.Timers, ShouldResemble, []tasktest.TimerSpec{
			{
				Delay: statusCheckTimerInterval,
				Name:  statusCheckTimerName,
			},
		})
		ctl.Timers = nil

		// The timer is called. Checks the state, reschedules itself.
		So(mgr.HandleTimer(c, ctl, statusCheckTimerName, nil), ShouldBeNil)
		So(ctl.Timers, ShouldResemble, []tasktest.TimerSpec{
			{
				Delay: statusCheckTimerInterval,
				Name:  statusCheckTimerName,
			},
		})

		// Process finish notification.
		mockRunning = false
		So(mgr.HandleNotification(c, ctl, &pubsub.PubsubMessage{}), ShouldBeNil)
		So(ctl.TaskState.Status, ShouldEqual, task.StatusSucceeded)
	})
}

func TestAbort(t *testing.T) {
	t.Parallel()

	Convey("LaunchTask and AbortTask work", t, func(ctx C) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			resp := ""
			switch {
			case r.Method == "PUT" && r.URL.Path == "/_ah/api/buildbucket/v1/builds":
				// There's more stuff in actual response that we don't use.
				resp = `{
					"build": {
						"id": "9025781602559305888",
						"status": "STARTED",
						"url": "https://chromium-swarm-dev.appspot.com/user/task/2bdfb7404d18ac10"
					}
				}`
			case r.Method == "POST" && r.URL.Path == "/_ah/api/buildbucket/v1/builds/9025781602559305888/cancel":
				resp = `{
					"build": {
						"id": "9025781602559305888",
						"status": "CANCELED",
						"url": "https://chromium-swarm-dev.appspot.com/user/task/2bdfb7404d18ac10"
					}
				}`
			default:
				ctx.Printf("Unknown URL fetch - %s %s\n", r.Method, r.URL.Path)
				w.WriteHeader(400)
				return
			}
			w.WriteHeader(200)
			w.Write([]byte(resp))
		}))
		defer ts.Close()

		c := memory.Use(context.Background())
		mgr := TaskManager{}
		ctl := fakeController(ts.URL)

		// Launch and kill.
		So(mgr.LaunchTask(c, ctl), ShouldBeNil)
		So(mgr.AbortTask(c, ctl), ShouldBeNil)
	})
}

func TestTriggeredFlow(t *testing.T) {
	t.Parallel()

	Convey("LaunchTask with GitilesTrigger works", t, func(ctx C) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			r.UserAgent()
			resp := ""
			switch {
			case r.Method == "PUT" && r.URL.Path == "/_ah/api/buildbucket/v1/builds":
				// There's more stuff in actual response that we don't use.
				resp = `{
					"build": {
						"id": "9025781602559305888",
						"status": "STARTED",
						"url": "https://chromium-swarm-dev.appspot.com/user/task/2bdfb7404d18ac10"
					}
				}`
			case r.Method == "GET" && r.URL.Path == "/_ah/api/buildbucket/v1/builds/9025781602559305888":
				resp = `{
						"build": {
							"id": "9025781602559305888",
							"status": "COMPLETED",
							"result": "SUCCESS"
						}
					}`
			default:
				ctx.Printf("Unknown URL fetch - %s %s\n", r.Method, r.URL.Path)
				w.WriteHeader(400)
				return
			}
			w.WriteHeader(200)
			w.Write([]byte(resp))
		}))
		defer ts.Close()

		c := memory.Use(context.Background())
		mgr := TaskManager{}
		ctl := fakeController(ts.URL)

		ctl.Req = task.Request{
			IncomingTriggers: []*internal.Trigger{
				{Id: "1", Payload: makePayload("https://r.googlesource.com/repo", "refs/heads/master", "baadcafe")},
				{Id: "2", Payload: makePayload("https://r.googlesource.com/repo", "refs/heads/master", "deadbeef")},
			},
		}

		// Launch with triggers,
		So(mgr.LaunchTask(c, ctl), ShouldBeNil)
		So(ctl.TaskState, ShouldResemble, task.State{
			Status:   task.StatusRunning,
			TaskData: []byte(`{"build_id":"9025781602559305888"}`),
			ViewURL:  "https://chromium-swarm-dev.appspot.com/user/task/2bdfb7404d18ac10",
		})
	})
}

func makePayload(repo, ref, rev string) *internal.Trigger_Gitiles {
	return &internal.Trigger_Gitiles{
		Gitiles: &api.GitilesTrigger{Repo: repo, Ref: ref, Revision: rev},
	}
}
