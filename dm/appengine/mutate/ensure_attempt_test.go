// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package mutate

import (
	"testing"

	"github.com/luci/gae/filter/featureBreaker"
	"github.com/luci/gae/impl/memory"
	ds "github.com/luci/gae/service/datastore"
	"github.com/luci/luci-go/dm/api/service/v1"
	"github.com/luci/luci-go/dm/appengine/model"

	"golang.org/x/net/context"

	. "github.com/luci/luci-go/common/testing/assertions"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEnsureAttempt(t *testing.T) {
	t.Parallel()

	Convey("EnsureAttempt", t, func() {
		c := memory.Use(context.Background())
		ea := &EnsureAttempt{dm.NewAttemptID("quest", 1)}

		Convey("Root", func() {
			So(ea.Root(c).String(), ShouldEqual, `dev~app::/Attempt,"quest|fffffffe"`)
		})

		Convey("RollForward", func() {
			a := &model.Attempt{ID: *ea.ID}

			Convey("Good", func() {
				So(ds.Get(c, a), ShouldEqual, ds.ErrNoSuchEntity)

				muts, err := ea.RollForward(c)
				So(err, ShouldBeNil)
				So(muts, ShouldHaveLength, 1)

				So(ds.Get(c, a), ShouldEqual, nil)
				So(a.State, ShouldEqual, dm.Attempt_SCHEDULING)

				Convey("replaying the mutation after the state has evolved is a noop", func() {
					So(a.ModifyState(c, dm.Attempt_EXECUTING), ShouldBeNil)
					So(ds.Put(c, a), ShouldBeNil)

					muts, err = ea.RollForward(c)
					So(err, ShouldBeNil)
					So(muts, ShouldBeEmpty)

					So(ds.Get(c, a), ShouldEqual, nil)
					So(a.State, ShouldEqual, dm.Attempt_EXECUTING)
				})
			})

			Convey("Bad", func() {
				c, fb := featureBreaker.FilterRDS(c, nil)
				fb.BreakFeatures(nil, "GetMulti")

				muts, err := ea.RollForward(c)
				So(err, ShouldErrLike, `feature "GetMulti" is broken`)
				So(muts, ShouldBeEmpty)

				fb.UnbreakFeatures("GetMulti")

				So(ds.Get(c, a), ShouldEqual, ds.ErrNoSuchEntity)
			})
		})
	})
}
