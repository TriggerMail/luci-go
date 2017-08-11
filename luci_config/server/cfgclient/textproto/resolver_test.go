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

package textproto

import (
	"testing"

	configPB "go.chromium.org/luci/common/proto/config"
	"go.chromium.org/luci/luci_config/server/cfgclient"
	"go.chromium.org/luci/luci_config/server/cfgclient/backend"
	"go.chromium.org/luci/luci_config/server/cfgclient/backend/format"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func tpb(msg proto.Message) string { return proto.MarshalTextString(msg) }

// testingBackend is a backend.B implementation that ignores Authority.
type testingBackend struct {
	backend.B

	items []*backend.Item
}

// Get retrieves a single configuration.
func (tb *testingBackend) Get(c context.Context, configSet, path string, p backend.Params) (*backend.Item, error) {
	if len(tb.items) == 0 {
		return nil, cfgclient.ErrNoConfig
	}
	return tb.cloneItems()[0], nil
}

// GetAll retrieves all configurations of a given type.
func (tb *testingBackend) GetAll(c context.Context, t backend.GetAllTarget, path string, p backend.Params) (
	[]*backend.Item, error) {

	return tb.cloneItems(), nil
}

func (tb *testingBackend) cloneItems() []*backend.Item {
	clones := make([]*backend.Item, len(tb.items))
	for i, it := range tb.items {
		clone := *it
		clones[i] = &clone
	}
	return clones
}

func TestResolver(t *testing.T) {
	Convey(`A testing environment`, t, func() {
		c := context.Background()

		var be backend.B
		be = &testingBackend{
			items: []*backend.Item{
				{Meta: backend.Meta{"projects/foo", "path", "####", "v1", "config_url"},
					Content: tpb(&configPB.Project{Id: proto.String("foo")})},
				{Meta: backend.Meta{"projects/bar", "path", "####", "v1", "config_url"},
					Content: tpb(&configPB.Project{Id: proto.String("bar")})},
			},
		}

		Convey(`Without a formatter backend, succeeds`, func() {
			c = backend.WithBackend(c, be)

			Convey(`Single`, func() {
				var val configPB.Project
				So(cfgclient.Get(c, cfgclient.AsService, "", "", Message(&val), nil), ShouldBeNil)
				So(val, ShouldResemble, configPB.Project{Id: proto.String("foo")})
			})

			Convey(`Multi`, func() {
				var (
					val  []*configPB.Project
					meta []*cfgclient.Meta
				)
				So(cfgclient.Projects(c, cfgclient.AsService, "", Slice(&val), &meta), ShouldBeNil)
				So(val, ShouldResemble, []*configPB.Project{
					{Id: proto.String("foo")},
					{Id: proto.String("bar")},
				})
				So(meta, ShouldResemble, []*cfgclient.Meta{
					{"projects/foo", "path", "####", "v1", "config_url"},
					{"projects/bar", "path", "####", "v1", "config_url"},
				})
			})
		})

		Convey(`With a formatter backend`, func() {
			format.ClearRegistry()

			be := &format.Backend{
				B: be,
			}
			c = backend.WithBackend(c, be)

			Convey(`If the Formatter is not registered, fails.`, func() {
				Convey(`Single`, func() {
					var val configPB.Project
					So(cfgclient.Get(c, cfgclient.AsService, "", "", Message(&val), nil), ShouldErrLike, "unknown formatter")
				})

				Convey(`Multi`, func() {
					var val []*configPB.Project
					So(cfgclient.Projects(c, cfgclient.AsService, "", Slice(&val), nil), ShouldErrLike, "unknown formatter")
				})
			})

			Convey(`If the Formatter is registered, succeeds`, func() {
				registerFormat()

				Convey(`Single`, func() {
					var val configPB.Project
					So(cfgclient.Get(c, cfgclient.AsService, "", "", Message(&val), nil), ShouldBeNil)
					So(val, ShouldResemble, configPB.Project{Id: proto.String("foo")})
				})

				Convey(`Multi`, func() {
					var (
						val  []*configPB.Project
						meta []*cfgclient.Meta
					)
					So(cfgclient.Projects(c, cfgclient.AsService, "", Slice(&val), &meta), ShouldBeNil)
					So(val, ShouldResemble, []*configPB.Project{
						{Id: proto.String("foo")},
						{Id: proto.String("bar")},
					})
					So(meta, ShouldResemble, []*cfgclient.Meta{
						{"projects/foo", "path", "####", "v1", "config_url"},
						{"projects/bar", "path", "####", "v1", "config_url"},
					})
				})
			})
		})
	})
}
