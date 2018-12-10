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
	"fmt"
	"net/http"

	"github.com/TriggerMail/luci-go/server/router"
	"github.com/TriggerMail/luci-go/server/templates"

	"github.com/TriggerMail/luci-go/common/sync/parallel"
	"github.com/TriggerMail/luci-go/milo/buildsource/buildbot"
	"github.com/TriggerMail/luci-go/milo/buildsource/buildbucket"
	"github.com/TriggerMail/luci-go/milo/frontend/ui"
)

// openSearchXML is the template used to serve the OpenSearch Description Document.
// This needs to be a template because the URL template must be a fully qualified
// URL with the hostname.
// See http://www.opensearch.org/Specifications/OpenSearch/1.1#OpenSearch_description_document
var openSearchXML = `<?xml version="1.0" encoding="UTF-8"?>
<OpenSearchDescription xmlns="http://a9.com/-/spec/opensearch/1.1/">
  <ShortName>LUCI</ShortName>
  <Description>
    Layered Universal Continuous Integration - A cloud based CI solution.
  </Description>
  <Url type="text/html" template="https://%s/search/?q={searchTerms}" />
</OpenSearchDescription>`

func searchHandler(c *router.Context) {
	var buildbotService, buildbucketService *ui.CIService
	err := parallel.FanOutIn(func(ch chan<- func() error) {
		ch <- func() (err error) {
			buildbotService, err = buildbot.CIService(c.Context)
			return
		}
		ch <- func() (err error) {
			buildbucketService, err = buildbucket.CIService(c.Context)
			return
		}
	})

	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	services := make([]ui.CIService, 0, 2)
	if buildbucketService != nil {
		services = append(services, *buildbucketService)
	}
	if buildbotService != nil {
		services = append(services, *buildbotService)
	}
	templates.MustRender(c.Context, c.Writer, "pages/search.html", templates.Args{
		"search": &ui.Search{CIServices: services},
		"error":  errMsg,
	})
}

// searchXMLHandler returns the opensearch document for this domain.
func searchXMLHandler(c *router.Context) {
	c.Writer.Header().Set("Content-Type", "application/opensearchdescription+xml")
	c.Writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(c.Writer, openSearchXML, c.Request.URL.Host)
}
