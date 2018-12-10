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

package cfgclient

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	"github.com/TriggerMail/luci-go/common/errors"
	"github.com/TriggerMail/luci-go/config"
	"github.com/TriggerMail/luci-go/config/server/cfgclient/backend"

	. "github.com/smartystreets/goconvey/convey"
)

// testingBackend is a Backend implementation that ignores Authority.
type testingBackend struct {
	serviceURL *url.URL
	err        error
	items      []*config.Config
	url        url.URL
	lastParams backend.Params
}

func (tb *testingBackend) ServiceURL(context.Context) url.URL { return *tb.serviceURL }

func (tb *testingBackend) Get(c context.Context, configSet config.Set, path string, p backend.Params) (*config.Config, error) {
	tb.lastParams = p
	if err := tb.err; err != nil {
		return nil, tb.err
	}
	if len(tb.items) == 0 {
		return nil, config.ErrNoConfig
	}
	return tb.cloneItems()[0], nil
}

func (tb *testingBackend) GetAll(c context.Context, t backend.GetAllTarget, path string, p backend.Params) (
	[]*config.Config, error) {

	tb.lastParams = p
	if err := tb.err; err != nil {
		return nil, tb.err
	}
	return tb.cloneItems(), nil
}

func (tb *testingBackend) GetConfigInterface(c context.Context, a backend.Authority) config.Interface {
	panic("not supported")
}

func (tb *testingBackend) cloneItems() []*config.Config {
	clones := make([]*config.Config, len(tb.items))
	for i, it := range tb.items {
		clone := *it
		clones[i] = &clone
	}
	return clones
}

type errorMultiResolver struct {
	getErr func(i int) error
	out    *[]string
}

func (er *errorMultiResolver) PrepareMulti(size int) {
	*er.out = make([]string, size)
}

func (er *errorMultiResolver) ResolveItemAt(i int, it *config.Config) error {
	if err := er.getErr(i); err != nil {
		return err
	}
	(*er.out)[i] = it.Content
	return nil
}

func TestConfig(t *testing.T) {
	t.Parallel()

	Convey(`A testing backend`, t, func() {
		c := context.Background()

		var err error
		tb := testingBackend{
			items: []*config.Config{
				{Meta: config.Meta{"projects/foo", "path", "####", "v1", "config_url"}, Content: "foo"},
				{Meta: config.Meta{"projects/bar", "path", "####", "v1", "config_url"}, Content: "bar"},
			},
		}
		transformItems := func(fn func(int, *config.Config)) {
			for i, itm := range tb.items {
				fn(i, itm)
			}
		}
		tb.serviceURL, err = url.Parse("http://example.com/config")
		if err != nil {
			panic(err)
		}
		c = backend.WithBackend(c, &tb)

		Convey(`Can resolve its service URL.`, func() {
			So(ServiceURL(c), ShouldResemble, *tb.serviceURL)
		})

		// Caching / type test cases.
		Convey(`Test byte resolver`, func() {
			transformItems(func(i int, ci *config.Config) {
				ci.Content = string([]byte{byte(i)})
			})

			Convey(`Single`, func() {
				var val []byte
				So(Get(c, AsService, "", "", Bytes(&val), nil), ShouldBeNil)
				So(val, ShouldResemble, []byte{0})
			})

			Convey(`Multi`, func() {
				var (
					val  [][]byte
					meta []*config.Meta
				)
				So(Projects(c, AsService, "", BytesSlice(&val), &meta), ShouldBeNil)
				So(val, ShouldResemble, [][]byte{{0}, {1}})
				So(meta, ShouldResemble, []*config.Meta{
					{"projects/foo", "path", "####", "v1", "config_url"},
					{"projects/bar", "path", "####", "v1", "config_url"},
				})
			})
		})

		Convey(`Test string resolver`, func() {
			transformItems(func(i int, ci *config.Config) {
				ci.Content = fmt.Sprintf("[%d]", i)
			})

			Convey(`Single`, func() {
				var val string
				So(Get(c, AsService, "", "", String(&val), nil), ShouldBeNil)
				So(val, ShouldEqual, "[0]")
			})

			Convey(`Multi`, func() {
				var (
					val  []string
					meta []*config.Meta
				)
				So(Projects(c, AsService, "", StringSlice(&val), &meta), ShouldBeNil)
				So(val, ShouldResemble, []string{"[0]", "[1]"})
				So(meta, ShouldResemble, []*config.Meta{
					{"projects/foo", "path", "####", "v1", "config_url"},
					{"projects/bar", "path", "####", "v1", "config_url"},
				})
			})
		})

		Convey(`Test nil resolver`, func() {
			transformItems(func(i int, ci *config.Config) {
				ci.Content = fmt.Sprintf("[%d]", i)
			})

			Convey(`Single`, func() {
				var meta config.Meta
				So(Get(c, AsService, "", "", nil, &meta), ShouldBeNil)
				So(meta, ShouldResemble, config.Meta{"projects/foo", "path", "####", "v1", "config_url"})
				So(tb.lastParams.Content, ShouldBeFalse)
			})

			Convey(`Multi`, func() {
				var meta []*config.Meta
				So(Projects(c, AsService, "", nil, &meta), ShouldBeNil)
				So(meta, ShouldResemble, []*config.Meta{
					{"projects/foo", "path", "####", "v1", "config_url"},
					{"projects/bar", "path", "####", "v1", "config_url"},
				})
				So(tb.lastParams.Content, ShouldBeFalse)
			})
		})

		Convey(`Projects with some errors returns a MultiError.`, func() {
			testErr := errors.New("test error")
			var (
				val  []string
				meta []*config.Meta
			)
			er := errorMultiResolver{
				out: &val,
				getErr: func(i int) error {
					if i == 1 {
						return testErr
					}
					return nil
				},
			}

			err := Projects(c, AsService, "", &er, &meta)

			// We have a MultiError with an error in position 1.
			So(err, ShouldResemble, errors.MultiError{nil, testErr})

			// One of our configs should have resolved.
			So(val, ShouldResemble, []string{"foo", ""})

			// Meta still works.
			So(meta, ShouldResemble, []*config.Meta{
				{"projects/foo", "path", "####", "v1", "config_url"},
				{"projects/bar", "path", "####", "v1", "config_url"},
			})
		})
	})
}
