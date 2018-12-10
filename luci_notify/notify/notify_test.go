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

package notify

import (
	"testing"

	"go.chromium.org/gae/service/datastore"

	"github.com/TriggerMail/luci-go/appengine/gaetesting"
	"github.com/TriggerMail/luci-go/appengine/tq"
	"github.com/TriggerMail/luci-go/buildbucket/proto"
	"github.com/TriggerMail/luci-go/common/clock"
	"github.com/TriggerMail/luci-go/common/clock/testclock"
	"github.com/TriggerMail/luci-go/common/logging"
	"github.com/TriggerMail/luci-go/common/logging/gologger"

	"github.com/TriggerMail/luci-go/luci_notify/config"
	"github.com/TriggerMail/luci-go/luci_notify/internal"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNotify(t *testing.T) {
	Convey("Notify", t, func() {
		c := gaetesting.TestingContextWithAppID("luci-notify")
		c = clock.Set(c, testclock.New(testclock.TestRecentTimeUTC))
		c = gologger.StdConfig.Use(c)
		c = logging.SetLevel(c, logging.Debug)

		build := &Build{
			Build: buildbucketpb.Build{
				Id: 54,
				Builder: &buildbucketpb.BuilderID{
					Project: "chromium",
					Bucket:  "ci",
					Builder: "linux-rel",
				},
				Status: buildbucketpb.Status_SUCCESS,
			},
		}

		// Put Project and EmailTemplate entities.
		project := &config.Project{Name: "chromium", Revision: "deadbeef"}
		templates := []*config.EmailTemplate{
			{
				ProjectKey:          datastore.KeyForObj(c, project),
				Name:                "default",
				SubjectTextTemplate: "Build {{.Build.Id}} completed",
				BodyHTMLTemplate:    "Build {{.Build.Id}} completed with status {{.Build.Status}}",
			},
			{
				ProjectKey:          datastore.KeyForObj(c, project),
				Name:                "non-default",
				SubjectTextTemplate: "Build {{.Build.Id}} completed from non-default template",
				BodyHTMLTemplate:    "Build {{.Build.Id}} completed with status {{.Build.Status}} from non-default template",
			},
		}
		So(datastore.Put(c, project, templates), ShouldBeNil)
		datastore.GetTestable(c).CatchupIndexes()

		Convey("createEmailTasks", func() {
			emailNotify := []EmailNotify{
				{
					Email: "jane@example.com",
				},
				{
					Email: "john@example.com",
				},
				{
					Email:    "don@example.com",
					Template: "non-default",
				},
			}
			tasks, err := createEmailTasks(c, emailNotify, &EmailTemplateInput{
				Build:     &build.Build,
				OldStatus: buildbucketpb.Status_SUCCESS,
			})
			So(err, ShouldBeNil)
			So(tasks, ShouldResemble, []*tq.Task{
				{
					DeduplicationKey: "54-default-jane@example.com",
					Payload: &internal.EmailTask{
						Recipients: []string{"jane@example.com"},
						Subject:    "Build 54 completed",
						Body:       "Build 54 completed with status SUCCESS",
					},
				},
				{
					DeduplicationKey: "54-default-john@example.com",
					Payload: &internal.EmailTask{
						Recipients: []string{"john@example.com"},
						Subject:    "Build 54 completed",
						Body:       "Build 54 completed with status SUCCESS",
					},
				},
				{
					DeduplicationKey: "54-non-default-don@example.com",
					Payload: &internal.EmailTask{
						Recipients: []string{"don@example.com"},
						Subject:    "Build 54 completed from non-default template",
						Body:       "Build 54 completed with status SUCCESS from non-default template",
					},
				},
			})
		})

		Convey("createEmailTasks with dup notifies", func() {
			emailNotify := []EmailNotify{
				{
					Email: "jane@example.com",
				},
				{
					Email: "jane@example.com",
				},
			}
			tasks, err := createEmailTasks(c, emailNotify, &EmailTemplateInput{
				Build:     &build.Build,
				OldStatus: buildbucketpb.Status_SUCCESS,
			})
			So(err, ShouldBeNil)
			So(tasks, ShouldHaveLength, 1)
		})
	})
}
