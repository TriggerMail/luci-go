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

// Package frontend contains the Machine Database AppEngine front end.
package frontend

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/TriggerMail/luci-go/appengine/gaemiddleware/standard"
	"github.com/TriggerMail/luci-go/common/data/rand/mathrand"
	"github.com/TriggerMail/luci-go/grpc/discovery"
	"github.com/TriggerMail/luci-go/grpc/grpcmon"
	"github.com/TriggerMail/luci-go/grpc/prpc"
	"github.com/TriggerMail/luci-go/server/router"

	"github.com/TriggerMail/luci-go/machine-db/api/crimson/v1"
	"github.com/TriggerMail/luci-go/machine-db/appengine/config"
	"github.com/TriggerMail/luci-go/machine-db/appengine/database"
	"github.com/TriggerMail/luci-go/machine-db/appengine/rpc"
	"github.com/TriggerMail/luci-go/machine-db/appengine/ui"
)

func init() {
	mathrand.SeedRandomly()
	databaseMiddleware := standard.Base().Extend(database.WithMiddleware)

	srv := rpc.NewServer()

	r := router.New()
	standard.InstallHandlers(r)
	config.InstallHandlers(r, databaseMiddleware)
	ui.InstallHandlers(r, databaseMiddleware, srv, "templates")

	api := prpc.Server{
		// Install an interceptor capable of reporting tsmon metrics.
		UnaryServerInterceptor: grpcmon.NewUnaryServerInterceptor(nil),
	}
	crimson.RegisterCrimsonServer(&api, srv)
	discovery.Enable(&api)
	api.InstallHandlers(r, databaseMiddleware)

	http.DefaultServeMux.Handle("/", r)
}
