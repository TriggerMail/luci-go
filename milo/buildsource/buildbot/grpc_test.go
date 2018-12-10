// Copyright 2016 The LUCI Authors.
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

package buildbot

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.chromium.org/gae/impl/memory"
	"go.chromium.org/gae/service/datastore"
	"github.com/TriggerMail/luci-go/common/clock/testclock"
	"github.com/TriggerMail/luci-go/milo/api/buildbot"
	milo "github.com/TriggerMail/luci-go/milo/api/proto"
	"github.com/TriggerMail/luci-go/milo/buildsource/buildbot/buildstore"
	"github.com/TriggerMail/luci-go/milo/common/model"
	"github.com/TriggerMail/luci-go/server/caching"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGRPC(t *testing.T) {
	t.Parallel()

	Convey(`A test environment`, t, func() {
		c := memory.Use(context.Background())
		c, _ = testclock.UseTime(c, testclock.TestTimeUTC)
		c = caching.WithRequestCache(c)

		name := "testmaster"
		bname := "testbuilder"
		master := &buildbot.Master{
			Name:     name,
			Builders: map[string]*buildbot.Builder{"fake": {}},
			Slaves: map[string]*buildbot.Slave{
				"foo": {
					RunningbuildsMap: map[string][]int{
						"fake": {1},
					},
				},
			},
		}

		So(buildstore.SaveMaster(c, master, false, nil), ShouldBeNil)
		importBuild(c, &buildbot.Build{
			Master:      name,
			Buildername: "fake",
			Number:      1,
		})
		datastore.GetTestable(c).Consistent(true)
		datastore.GetTestable(c).AutoIndex(true)
		svc := Service{}

		Convey(`Get finished builds`, func() {
			// Add in some builds.
			for i := 0; i < 5; i++ {
				importBuild(c, &buildbot.Build{
					Master:      name,
					Buildername: bname,
					Number:      i,
					Finished:    true,
				})
			}
			importBuild(c, &buildbot.Build{
				Master:      name,
				Buildername: bname,
				Number:      6,
				Finished:    false,
			})
			datastore.GetTestable(c).CatchupIndexes()

			r := &milo.BuildbotBuildsRequest{
				Master:  name,
				Builder: bname,
			}
			result, err := svc.GetBuildbotBuildsJSON(c, r)
			So(err, ShouldBeNil)
			So(len(result.Builds), ShouldEqual, 5)

			Convey(`Also get incomplete builds`, func() {
				r := &milo.BuildbotBuildsRequest{
					Master:         name,
					Builder:        bname,
					IncludeCurrent: true,
				}
				result, err := svc.GetBuildbotBuildsJSON(c, r)
				So(err, ShouldBeNil)
				So(len(result.Builds), ShouldEqual, 6)
			})

			Convey(`Bad request`, func() {
				_, err := svc.GetBuildbotBuildsJSON(c, &milo.BuildbotBuildsRequest{})
				So(err, ShouldResemble, status.Errorf(codes.InvalidArgument, "No master specified"))
				_, err = svc.GetBuildbotBuildsJSON(c, &milo.BuildbotBuildsRequest{Master: name})
				So(err, ShouldResemble, status.Errorf(codes.InvalidArgument, "No builder specified"))
			})
		})

		Convey(`Get buildbotMasterEntry`, func() {
			Convey(`Bad request`, func() {
				_, err := svc.GetCompressedMasterJSON(c, &milo.MasterRequest{})
				So(err, ShouldResemble, status.Errorf(codes.InvalidArgument, "No master specified"))
			})
			_, err := svc.GetCompressedMasterJSON(c, &milo.MasterRequest{Name: name})
			So(err, ShouldBeNil)
		})

		Convey(`Get Build`, func() {
			Convey(`Invalid input`, func() {
				_, err := svc.GetBuildbotBuildJSON(c, &milo.BuildbotBuildRequest{})
				So(err, ShouldResemble, status.Errorf(codes.InvalidArgument, "No master specified"))
				_, err = svc.GetBuildbotBuildJSON(c, &milo.BuildbotBuildRequest{
					Master: "foo",
				})
				So(err, ShouldResemble, status.Errorf(codes.InvalidArgument, "No builder specified"))
			})
			Convey(`Basic`, func() {
				_, err := svc.GetBuildbotBuildJSON(c, &milo.BuildbotBuildRequest{
					Master:   name,
					Builder:  "fake",
					BuildNum: 1,
				})
				So(err, ShouldBeNil)
			})
			Convey(`Basic Not found`, func() {
				_, err := svc.GetBuildbotBuildJSON(c, &milo.BuildbotBuildRequest{
					Master:   name,
					Builder:  "fake",
					BuildNum: 2,
				})
				So(err, ShouldResemble, status.Errorf(codes.NotFound, "Not found"))
			})
		})
	})
}

func importBuild(c context.Context, b *buildbot.Build) {
	// Make sure a BuilderSummary exists for this Build.
	builderID := fmt.Sprintf("buildbot/%s/%s", b.Master, b.Buildername)
	builder := model.BuilderSummary{BuilderID: builderID}
	datastore.Put(c, &builder)
	datastore.GetTestable(c).CatchupIndexes()

	_, err := buildstore.SaveBuild(c, b)
	So(err, ShouldBeNil)
}
