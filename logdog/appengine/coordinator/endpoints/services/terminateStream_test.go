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

package services

import (
	"context"
	"errors"
	"testing"
	"time"

	"go.chromium.org/gae/filter/featureBreaker"
	ds "go.chromium.org/gae/service/datastore"

	"github.com/TriggerMail/luci-go/logdog/api/config/svcconfig"
	"github.com/TriggerMail/luci-go/logdog/api/endpoints/coordinator/services/v1"
	"github.com/TriggerMail/luci-go/logdog/appengine/coordinator"
	ct "github.com/TriggerMail/luci-go/logdog/appengine/coordinator/coordinatorTest"
	"github.com/TriggerMail/luci-go/logdog/appengine/coordinator/endpoints"
	"github.com/TriggerMail/luci-go/logdog/appengine/coordinator/mutations"

	"github.com/TriggerMail/luci-go/common/clock"
	"github.com/TriggerMail/luci-go/common/proto/google"
	"github.com/TriggerMail/luci-go/tumble"

	. "github.com/smartystreets/goconvey/convey"
	. "github.com/TriggerMail/luci-go/common/testing/assertions"
)

func TestTerminateStream(t *testing.T) {
	t.Parallel()

	Convey(`With a testing configuration`, t, func() {
		c, env := ct.Install(true)

		// Set our archival delays. The project delay is smaller than the service
		// delay, so it should be used.
		env.ModServiceConfig(c, func(cfg *svcconfig.Config) {
			coord := cfg.Coordinator
			coord.ArchiveTopic = "projects/test/topics/archive"
			coord.ArchiveSettleDelay = google.NewDuration(10 * time.Second)
			coord.ArchiveDelayMax = google.NewDuration(24 * time.Hour)
		})
		env.ModProjectConfig(c, "proj-foo", func(pcfg *svcconfig.ProjectConfig) {
			pcfg.MaxStreamAge = google.NewDuration(time.Hour)
		})

		svr := New()

		tls := ct.MakeStream(c, "proj-foo", "testing/+/foo/bar")

		req := logdog.TerminateStreamRequest{
			Project:       string(tls.Project),
			Id:            string(tls.Stream.ID),
			Secret:        tls.Prefix.Secret,
			TerminalIndex: 1337,
		}

		Convey(`Returns Forbidden error if not a service.`, func() {
			_, err := svr.TerminateStream(c, &req)
			So(err, ShouldBeRPCPermissionDenied)
		})

		Convey(`When logged in as a service`, func() {
			env.JoinGroup("services")

			Convey(`A non-terminal registered stream, "testing/+/foo/bar"`, func() {
				tls.WithProjectNamespace(c, func(c context.Context) {
					So(tls.Put(c), ShouldBeNil)

					// Create an archival request for Tumble so we can ensure that it is
					// replaced on termination. This is normally done by RegisterStream.
					areq := mutations.CreateArchiveTask{
						ID:         tls.Stream.ID,
						Expiration: env.Clock.Now().Add(time.Hour),
					}
					arParent, arName := ds.KeyForObj(c, tls.Stream), areq.TaskName(c)
					err := tumble.PutNamedMutations(c, arParent, map[string]tumble.Mutation{
						arName: &areq,
					})
					if err != nil {
						panic(err)
					}
				})
				ds.GetTestable(c).CatchupIndexes()

				Convey(`Can be marked terminal and schedules an archival mutation.`, func() {
					_, err := svr.TerminateStream(c, &req)
					So(err, ShouldBeRPCOK)
					ds.GetTestable(c).CatchupIndexes()

					// Reload the state and confirm.
					tls.WithProjectNamespace(c, func(c context.Context) {
						So(ds.Get(c, tls.State), ShouldBeNil)
					})
					So(tls.State.TerminalIndex, ShouldEqual, 1337)
					So(tls.State.Terminated(), ShouldBeTrue)
					So(tls.State.ArchivalState(), ShouldEqual, coordinator.NotArchived)

					Convey(`Replaces the pessimistic archival mutation with an optimistic one.`, func() {
						// We have replaced the pessimistic archival mutation with an
						// optimistic one. Assert that this happened by advancing time by
						// the optimistic period and confirming the published archival
						// request.
						env.IterateTumbleAll(c)
						So(env.ArchivalPublisher.Hashes(), ShouldBeEmpty)

						// Add our settle delay, confirm that archival is scheduled.
						env.Clock.Add(10 * time.Second)
						env.Clock.SetTimerCallback(func(d time.Duration, tmr clock.Timer) {
							env.Clock.Add(3 * time.Second)
						})
						env.IterateTumbleAll(c)
						// TODO(hinoka): Fix me.  This racily fails on bots for some reason,
						// likely because the hack to increment clock by 3s on every timer call
						// causes the time to go beyond the 9min settle delay, so the pending
						// archival request ends up getting processed.
						SkipSo(env.ArchivalPublisher.Hashes(), ShouldResemble, []string{string(tls.Stream.ID)})

						// Add our pessimistic delay, confirm that no additional tasks
						// are scheduled (because pessimistic was replaced).
						env.ArchivalPublisher.Clear()
						env.Clock.Add(time.Hour)
						env.IterateTumbleAll(c)
						SkipSo(env.ArchivalPublisher.Hashes(), ShouldBeEmpty)
					})

					Convey(`Can be marked terminal again (idempotent).`, func() {
						_, err := svr.TerminateStream(c, &req)
						So(err, ShouldBeRPCOK)

						// Reload state and confirm.
						So(tls.Get(c), ShouldBeNil)

						So(tls.State.Terminated(), ShouldBeTrue)
						So(tls.State.TerminalIndex, ShouldEqual, 1337)
						So(tls.State.ArchivalState(), ShouldEqual, coordinator.NotArchived)
					})

					Convey(`Will reject attempts to change the terminal index.`, func() {
						req.TerminalIndex = 1338
						_, err := svr.TerminateStream(c, &req)
						So(err, ShouldBeRPCFailedPrecondition, "Log stream is incompatibly terminated.")

						// Reload state and confirm.
						So(tls.Get(c), ShouldBeNil)

						So(tls.State.TerminalIndex, ShouldEqual, 1337)
						So(tls.State.Terminated(), ShouldBeTrue)
						So(tls.State.ArchivalState(), ShouldEqual, coordinator.NotArchived)
					})

					Convey(`Will reject attempts to clear the terminal index.`, func() {
						req.TerminalIndex = -1
						_, err := svr.TerminateStream(c, &req)
						So(err, ShouldBeRPCInvalidArgument, "Negative terminal index.")

						// Reload state and confirm.
						So(tls.Get(c), ShouldBeNil)

						So(tls.State.TerminalIndex, ShouldEqual, 1337)
						So(tls.State.Terminated(), ShouldBeTrue)
						So(tls.State.ArchivalState(), ShouldEqual, coordinator.NotArchived)
					})
				})

				Convey(`Will return an internal server error if Put() fails.`, func() {
					c, fb := featureBreaker.FilterRDS(c, nil)
					fb.BreakFeatures(errors.New("test error"), "PutMulti")
					_, err := svr.TerminateStream(c, &req)
					So(err, ShouldBeRPCInternal)
				})

				Convey(`Will return an internal server error if Get() fails.`, func() {
					c, fb := featureBreaker.FilterRDS(c, nil)
					fb.BreakFeatures(errors.New("test error"), "GetMulti")
					_, err := svr.TerminateStream(c, &req)
					So(err, ShouldBeRPCInternal)
				})

				Convey(`Will return a bad request error if the secret doesn't match.`, func() {
					req.Secret[0] ^= 0xFF
					_, err := svr.TerminateStream(c, &req)
					So(err, ShouldBeRPCInvalidArgument, "Request secret doesn't match the stream secret.")
				})
			})

			Convey(`Will not try and terminate a stream with an invalid path.`, func() {
				req.Id = "!!!invalid path!!!"
				_, err := svr.TerminateStream(c, &req)
				So(err, ShouldBeRPCInvalidArgument, "Invalid ID")
			})

			Convey(`Will fail if the stream is not registered.`, func() {
				_, err := svr.TerminateStream(c, &req)
				So(err, ShouldBeRPCNotFound, "is not registered")
			})
		})

		Convey(`Will choose the correct archival delay`, func() {
			getParams := func() *coordinator.ArchivalParams {
				svc := endpoints.GetServices(c)
				cfg, err := svc.Config(c)
				if err != nil {
					panic(err)
				}

				pcfg, err := svc.ProjectConfig(c, "proj-foo")
				if err != nil {
					panic(err)
				}

				return standardArchivalParams(cfg, pcfg)
			}

			Convey(`When there is no project config delay.`, func() {
				env.ModProjectConfig(c, "proj-foo", func(pcfg *svcconfig.ProjectConfig) {
					pcfg.MaxStreamAge = nil
				})

				So(getParams(), ShouldResemble, &coordinator.ArchivalParams{
					SettleDelay:    10 * time.Second,
					CompletePeriod: 24 * time.Hour,
				})
			})

			Convey(`When there is no service or project config delay.`, func() {
				env.ModServiceConfig(c, func(cfg *svcconfig.Config) {
					cfg.Coordinator.ArchiveDelayMax = nil
				})
				env.ModProjectConfig(c, "proj-foo", func(pcfg *svcconfig.ProjectConfig) {
					pcfg.MaxStreamAge = nil
				})

				So(getParams(), ShouldResemble, &coordinator.ArchivalParams{
					SettleDelay:    10 * time.Second,
					CompletePeriod: 0,
				})
			})
		})
	})
}
