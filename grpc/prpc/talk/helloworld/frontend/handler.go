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

package helloworld

import (
	"net/http"

	"github.com/TriggerMail/luci-go/appengine/gaemiddleware/standard"
	"github.com/TriggerMail/luci-go/server/router"
)

func init() {
	r := router.New()

	standard.InstallHandlers(r)
	InstallAPIRoutes(r, standard.Base())
	http.DefaultServeMux.Handle("/", r)
}
