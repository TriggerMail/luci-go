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

package module

import (
	"net/http"

	"github.com/TriggerMail/luci-go/appengine/gaemiddleware/standard"
	"github.com/TriggerMail/luci-go/logdog/appengine/coordinator/endpoints"
	"github.com/TriggerMail/luci-go/server/router"
	"github.com/TriggerMail/luci-go/tumble"

	// Include mutations package so its Mutations will register with tumble via
	// init().
	_ "github.com/TriggerMail/luci-go/logdog/appengine/coordinator/mutations"
)

func init() {
	tmb := tumble.Service{}
	ps := endpoints.ProdService{}

	r := router.New()
	standard.InstallHandlers(r)

	base := standard.Base().Extend(ps.Base)
	tmb.InstallHandlers(r, base)

	http.Handle("/", r)
}
