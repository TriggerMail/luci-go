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

// Package module is a simple AppEngine LUCI service. It supplies basic LUCI
// service frontend and backend functionality.
//
// No RPC requests should target this service; instead, they are redirected to
// the appropriate service via "dispatch.yaml".
package module

import (
	"net/http"

	adminPb "github.com/TriggerMail/luci-go/logdog/api/endpoints/coordinator/admin/v1"
	logsPb "github.com/TriggerMail/luci-go/logdog/api/endpoints/coordinator/logs/v1"
	registrationPb "github.com/TriggerMail/luci-go/logdog/api/endpoints/coordinator/registration/v1"
	servicesPb "github.com/TriggerMail/luci-go/logdog/api/endpoints/coordinator/services/v1"
	"github.com/TriggerMail/luci-go/logdog/appengine/coordinator/endpoints"
	"github.com/TriggerMail/luci-go/logdog/appengine/coordinator/endpoints/admin"

	"github.com/TriggerMail/luci-go/appengine/gaemiddleware/standard"
	"github.com/TriggerMail/luci-go/grpc/discovery"
	"github.com/TriggerMail/luci-go/grpc/grpcmon"
	"github.com/TriggerMail/luci-go/grpc/prpc"
	"github.com/TriggerMail/luci-go/server/router"
)

// Run installs and executes this site.
func init() {
	ps := endpoints.ProdService{}

	r := router.New()

	// Standard HTTP endpoints.
	standard.InstallHandlers(r)

	// Register all of the handlers that we want to show up in RPC explorer (via
	// pRPC discovery).
	//
	// Note that most of these services have dedicated service handlers, and any
	// RPCs sent to this module will automatically be routed to them via
	// "dispatch.yaml".
	svr := &prpc.Server{
		UnaryServerInterceptor: grpcmon.NewUnaryServerInterceptor(nil),
	}
	logsPb.RegisterLogsServer(svr, dummyLogsService)
	registrationPb.RegisterRegistrationServer(svr, dummyRegistrationService)
	servicesPb.RegisterServicesServer(svr, dummyServicesService)
	adminPb.RegisterAdminServer(svr, admin.New())
	discovery.Enable(svr)

	base := standard.Base().Extend(ps.Base)
	svr.InstallHandlers(r, base)

	// Redirect "/" to "/app/".
	r.GET("/", router.MiddlewareChain{}, func(c *router.Context) {
		http.Redirect(c.Writer, c.Request, "/app/", http.StatusFound)
	})
	// Redirect "/v/?s=..." to "/logs/..."
	r.GET("/v/", router.MiddlewareChain{}, func(c *router.Context) {
		path := "/logs/" + c.Request.URL.Query().Get("s")
		http.Redirect(c.Writer, c.Request, path, http.StatusFound)
	})

	http.Handle("/", r)
}
