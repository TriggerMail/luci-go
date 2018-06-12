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

package model

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"

	"go.chromium.org/gae/service/datastore"
	"go.chromium.org/luci/appengine/gaetesting"
	"go.chromium.org/luci/common/proto/google"
	"go.chromium.org/luci/grpc/grpcutil"

	api "go.chromium.org/luci/cipd/api/cipd/v1"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestRegisterInstance(t *testing.T) {
	t.Parallel()

	Convey("With datastore", t, func() {
		ctx := gaetesting.TestingContext()
		ts := time.Unix(1525136124, 0).UTC()

		pkg := &Package{
			Name:         "a/b/c",
			RegisteredBy: "user:a@example.com",
			RegisteredTs: ts,
		}

		inst := &Instance{
			InstanceID:   strings.Repeat("a", 40),
			Package:      PackageKey(ctx, "a/b/c"),
			RegisteredBy: "user:a@example.com",
			RegisteredTs: ts,
		}

		Convey("To proto", func() {
			So(inst.Proto(), ShouldResemble, &api.Instance{
				Package: "a/b/c",
				Instance: &api.ObjectRef{
					HashAlgo:  api.HashAlgo_SHA1,
					HexDigest: inst.InstanceID,
				},
				RegisteredBy: "user:a@example.com",
				RegisteredTs: google.NewTimestamp(ts),
			})
		})

		Convey("New package and instance", func() {
			reg, out, err := RegisterInstance(ctx, inst, func(c context.Context, inst *Instance) error {
				inst.ProcessorsPending = []string{"a"}
				return nil
			})
			So(err, ShouldBeNil)
			So(reg, ShouldBeTrue)

			expected := &Instance{
				InstanceID:        inst.InstanceID,
				Package:           inst.Package,
				RegisteredBy:      inst.RegisteredBy,
				RegisteredTs:      inst.RegisteredTs,
				ProcessorsPending: []string{"a"},
			}
			So(out, ShouldResemble, expected)

			// Created instance and package entities.
			storedInst := &Instance{
				InstanceID: out.InstanceID,
				Package:    inst.Package,
			}
			storedPkg := &Package{Name: "a/b/c"}
			So(datastore.Get(ctx, storedInst, storedPkg), ShouldBeNil)

			So(storedInst, ShouldResemble, expected)
			So(storedPkg, ShouldResemble, pkg)
		})

		Convey("Existing package, new instance", func() {
			So(datastore.Put(ctx, pkg), ShouldBeNil)

			inst.RegisteredBy = "user:someoneelse@example.com"
			reg, out, err := RegisterInstance(ctx, inst, func(c context.Context, inst *Instance) error {
				inst.ProcessorsPending = []string{"a"}
				return nil
			})
			So(err, ShouldBeNil)
			So(reg, ShouldBeTrue)
			So(out, ShouldResemble, &Instance{
				InstanceID:        inst.InstanceID,
				Package:           inst.Package,
				RegisteredBy:      inst.RegisteredBy,
				RegisteredTs:      inst.RegisteredTs,
				ProcessorsPending: []string{"a"},
			})

			// Package entity wasn't touched.
			storedPkg := &Package{Name: "a/b/c"}
			So(datastore.Get(ctx, storedPkg), ShouldBeNil)
			So(storedPkg, ShouldResemble, pkg)
		})

		Convey("Existing instance", func() {
			So(datastore.Put(ctx, pkg, inst), ShouldBeNil)

			modified := *inst
			modified.RegisteredBy = "user:someoneelse@example.com"
			reg, out, err := RegisterInstance(ctx, &modified, func(c context.Context, inst *Instance) error {
				panic("must not be called")
			})
			So(err, ShouldBeNil)
			So(reg, ShouldBeFalse)
			So(out, ShouldResemble, inst) // the original one
		})
	})
}

func TesRefIIDConversion(t *testing.T) {
	t.Parallel()

	Convey("SHA1 works", t, func() {
		sha1 := strings.Repeat("a", 40)

		So(ObjectRefToInstanceID(&api.ObjectRef{
			HashAlgo:  api.HashAlgo_SHA1,
			HexDigest: sha1,
		}), ShouldEqual, sha1)

		So(InstanceIDToObjectRef(sha1), ShouldResemble, &api.ObjectRef{
			HashAlgo:  api.HashAlgo_SHA1,
			HexDigest: sha1,
		})
	})
}

func TestListInstances(t *testing.T) {
	t.Parallel()

	Convey("With datastore", t, func() {
		ts := time.Unix(1525136124, 0).UTC()
		ctx := gaetesting.TestingContext()
		datastore.GetTestable(ctx).AutoIndex(true)

		inst := func(i int) *Instance {
			return &Instance{
				InstanceID:   fmt.Sprintf("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa%d", i),
				Package:      PackageKey(ctx, "a/b"),
				RegisteredTs: ts.Add(time.Duration(i) * time.Minute),
			}
		}
		for i := 0; i < 4; i++ {
			So(datastore.Put(ctx, inst(i)), ShouldBeNil)
		}

		Convey("Full listing", func() {
			out, cur, err := ListInstances(ctx, "a/b", 0, nil)
			So(err, ShouldBeNil)
			So(cur, ShouldBeNil)
			So(out, ShouldResemble, []*Instance{inst(3), inst(2), inst(1), inst(0)})
		})

		Convey("Paginated listing", func() {
			out, cur, err := ListInstances(ctx, "a/b", 3, nil)
			So(err, ShouldBeNil)
			So(cur, ShouldNotBeNil)
			So(out, ShouldResemble, []*Instance{inst(3), inst(2), inst(1)})

			out, cur, err = ListInstances(ctx, "a/b", 3, cur)
			So(err, ShouldBeNil)
			So(cur, ShouldBeNil)
			So(out, ShouldResemble, []*Instance{inst(0)})
		})
	})
}

func TestCheckInstanceReady(t *testing.T) {
	t.Parallel()

	Convey("With datastore", t, func() {
		ctx := gaetesting.TestingContext()

		put := func(pkg, iid string, failedProcs, pendingProcs []string) {
			So(datastore.Put(ctx,
				&Package{Name: pkg},
				&Instance{
					InstanceID:        iid,
					Package:           PackageKey(ctx, pkg),
					ProcessorsFailure: failedProcs,
					ProcessorsPending: pendingProcs,
				}), ShouldBeNil)
		}

		iid := ObjectRefToInstanceID(&api.ObjectRef{
			HashAlgo:  api.HashAlgo_SHA1,
			HexDigest: strings.Repeat("a", 40),
		})
		inst := &Instance{
			InstanceID: iid,
			Package:    PackageKey(ctx, "pkg"),
		}

		Convey("Happy path", func() {
			put("pkg", iid, nil, nil)
			So(CheckInstanceReady(ctx, inst), ShouldBeNil)
		})

		Convey("No such instance", func() {
			put("pkg", "f"+iid[1:], nil, nil)
			err := CheckInstanceReady(ctx, inst)
			So(grpcutil.Code(err), ShouldEqual, codes.NotFound)
			So(err, ShouldErrLike, "no such instance")
		})

		Convey("No such package", func() {
			put("pkg2", iid, nil, nil)
			err := CheckInstanceReady(ctx, inst)
			So(grpcutil.Code(err), ShouldEqual, codes.NotFound)
			So(err, ShouldErrLike, "no such package")
		})

		Convey("Failed processors", func() {
			put("pkg", iid, []string{"f1", "f2"}, []string{"p1", "p2"})
			err := CheckInstanceReady(ctx, inst)
			So(grpcutil.Code(err), ShouldEqual, codes.Aborted)
			So(err, ShouldErrLike, "some processors failed to process this instance: f1, f2")
		})

		Convey("Pending processors", func() {
			put("pkg", iid, nil, []string{"p1", "p2"})
			err := CheckInstanceReady(ctx, inst)
			So(grpcutil.Code(err), ShouldEqual, codes.FailedPrecondition)
			So(err, ShouldErrLike, "the instance is not ready yet, pending processors: p1, p2")
		})
	})
}
