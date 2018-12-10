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

package frontend

import (
	"context"
	"net/http"

	authServer "github.com/TriggerMail/luci-go/appengine/gaeauth/server"
	"github.com/TriggerMail/luci-go/appengine/gaemiddleware"
	"github.com/TriggerMail/luci-go/appengine/gaemiddleware/standard"
	"github.com/TriggerMail/luci-go/appengine/tq"
	"github.com/TriggerMail/luci-go/common/logging"
	"github.com/TriggerMail/luci-go/common/retry/transient"
	"github.com/TriggerMail/luci-go/config/appengine/gaeconfig"
	"github.com/TriggerMail/luci-go/config/impl/remote"
	"github.com/TriggerMail/luci-go/server/auth"
	"github.com/TriggerMail/luci-go/server/router"

	"github.com/TriggerMail/luci-go/luci_notify/config"
	"github.com/TriggerMail/luci-go/luci_notify/notify"
)

func init() {
	r := router.New()
	standard.InstallHandlers(r)

	basemw := standard.Base().Extend(auth.Authenticate(authServer.CookieAuth), withRemoteConfigService)

	taskDispatcher := tq.Dispatcher{BaseURL: "/internal/tasks/"}
	notify.InitDispatcher(&taskDispatcher)
	taskDispatcher.InstallRoutes(r, basemw)

	// Cron endpoint.
	r.GET("/internal/cron/update-config", basemw.Extend(gaemiddleware.RequireCron), config.UpdateHandler)

	// Pub/Sub endpoint.
	r.POST("/_ah/push-handlers/buildbucket", basemw, func(c *router.Context) {
		if err := notify.BuildbucketPubSubHandler(c, &taskDispatcher); err != nil {
			logging.Errorf(c.Context, "%s", err)
			if transient.Tag.In(err) {
				// Retry transient errors.
				c.Writer.WriteHeader(http.StatusInternalServerError)
			}
		}
	})

	http.Handle("/", r)
}

func withRemoteConfigService(c *router.Context, next router.Handler) {
	s, err := gaeconfig.FetchCachedSettings(c.Context)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		logging.WithError(err).Errorf(c.Context, "failure retrieving cached settings")
		return
	}

	rInterface := remote.New(s.ConfigServiceHost, false, func(c context.Context) (*http.Client, error) {
		t, err := auth.GetRPCTransport(c, auth.AsSelf)
		if err != nil {
			return nil, err
		}
		return &http.Client{Transport: t}, nil
	})
	// insert into context
	c.Context = config.WithConfigService(c.Context, rInterface)
	next(c)
}
