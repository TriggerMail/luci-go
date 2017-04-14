// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package frontend

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/luci/gae/impl/memory"
	"github.com/luci/luci-go/common/clock/testclock"
	"github.com/luci/luci-go/milo/appengine/buildbot"
	"github.com/luci/luci-go/milo/appengine/common"
	"github.com/luci/luci-go/milo/appengine/swarming"
	"github.com/luci/luci-go/server/auth"
	"github.com/luci/luci-go/server/auth/identity"
	"github.com/luci/luci-go/server/settings"
	"github.com/luci/luci-go/server/templates"

	. "github.com/smartystreets/goconvey/convey"
)

type testPackage struct {
	Data         func() []common.TestBundle
	DisplayName  string
	TemplateName string
}

var (
	allPackages = []testPackage{
		{buildbot.BuildTestData, "buildbot.build", "build.html"},
		{buildbot.BuilderTestData, "buildbot.builder", "builder.html"},
		{swarming.BuildTestData, "swarming.build", "build.html"},
		{swarming.LogTestData, "swarming.log", "log.html"},
		{frontpageTestData, "frontpage", "frontpage.html"},
	}
)

var generate = flag.Bool(
	"test.generate", false, "Generate expectations instead of running tests.")

func expectFileName(name string) string {
	name = strings.Replace(name, " ", "_", -1)
	name = strings.Replace(name, "/", "_", -1)
	name = strings.Replace(name, ":", "-", -1)
	return filepath.Join("expectations", name)
}

func load(name string) ([]byte, error) {
	filename := expectFileName(name)
	return ioutil.ReadFile(filename)
}

// mustWrite Writes a buffer into an expectation file.  Should always work or
// panic.  This is fine because this only runs when -generate is passed in,
// not during tests.
func mustWrite(name string, buf []byte) {
	filename := expectFileName(name)
	err := ioutil.WriteFile(filename, buf, 0644)
	if err != nil {
		panic(err)
	}
}

// fakeOAuthMethod implements Method.
type fakeOAuthMethod struct {
	clientID string
}

func (m fakeOAuthMethod) Authenticate(context.Context, *http.Request) (*auth.User, error) {
	return &auth.User{
		Identity: identity.Identity("user:abc@example.com"),
		Email:    "abc@example.com",
		ClientID: m.clientID,
	}, nil
}

func (m fakeOAuthMethod) LoginURL(c context.Context, target string) (string, error) {
	return "https://login.url/?target=" + target, nil
}

func (m fakeOAuthMethod) LogoutURL(c context.Context, target string) (string, error) {
	return "https://logout.url/?target=" + target, nil
}

type analyticsSettings struct {
	AnalyticsID string `json:"analytics_id"`
}

func TestPages(t *testing.T) {
	fixZeroDurationRE := regexp.MustCompile(`(Running for:|waiting) 0s?`)
	fixZeroDuration := func(text string) string {
		return fixZeroDurationRE.ReplaceAllLiteralString(text, "[ZERO DURATION]")
	}

	Convey("Testing basic rendering.", t, func() {
		c := context.Background()
		c = memory.Use(c)
		c = common.WithRequest(c, &http.Request{URL: &url.URL{Path: "/foobar"}})
		c, _ = testclock.UseTime(c, testclock.TestTimeUTC)
		a := auth.Authenticator{fakeOAuthMethod{"some_client_id"}}
		c = auth.SetAuthenticator(c, a)
		c = settings.Use(c, settings.New(&settings.MemoryStorage{Expiration: time.Second}))
		err := settings.Set(c, "analytics", &analyticsSettings{"UA-12345-01"}, "", "")
		So(err, ShouldBeNil)
		c = templates.Use(c, common.GetTemplateBundle())
		for _, p := range allPackages {
			Convey(fmt.Sprintf("Testing handler %q", p.DisplayName), func() {
				for _, b := range p.Data() {
					Convey(fmt.Sprintf("Testing: %q", b.Description), func() {
						args := b.Data
						// This is not a path, but a file key, should always be "/".
						tmplName := "pages/" + p.TemplateName
						buf, err := templates.Render(c, tmplName, args)
						So(err, ShouldBeNil)
						fname := fmt.Sprintf(
							"%s-%s.html", p.DisplayName, b.Description)
						if *generate {
							mustWrite(fname, buf)
						} else {
							localBuf, err := load(fname)
							So(err, ShouldBeNil)
							So(fixZeroDuration(string(buf)), ShouldEqual, fixZeroDuration(string(localBuf)))
						}
					})
				}
			})
		}
	})
}
