// Copyright 2017 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package frontend

import (
	"github.com/TriggerMail/luci-go/server/router"
	"github.com/TriggerMail/luci-go/server/templates"

	"github.com/TriggerMail/luci-go/milo/buildsource"
)

// BuilderHandler is responsible for taking a universal builder ID and rendering
// the builder page (defined in ./appengine/templates/pages/builder.html).
// We don't need to do an ACL check because this endpoint delegates all ACL checks
// authentication to Buildbucket with the RPC calls.
func BuilderHandler(c *router.Context, builderID buildsource.BuilderID) error {
	limit := 25
	if tLimit := GetLimit(c.Request, -1); tLimit >= 0 {
		limit = tLimit
	}
	builder, err := builderID.Get(c.Context, limit, c.Request.FormValue("cursor"))
	if err != nil {
		return err
	}
	templates.MustRender(c.Context, c.Writer, "pages/builder.html", templates.Args{
		"Builder": builder,
	})
	return nil
}
