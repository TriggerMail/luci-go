// Copyright 2017 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package frontend

import (
	"net/http"

	"github.com/TriggerMail/luci-go/common/errors"
	"github.com/TriggerMail/luci-go/server/auth"
	"github.com/TriggerMail/luci-go/server/router"
	"github.com/TriggerMail/luci-go/server/templates"

	"github.com/TriggerMail/luci-go/milo/common"
)

// ErrorHandler renders an error page for the user.
func ErrorHandler(c *router.Context, err error) {
	// TODO(iannucci): tag/extract other information from error, like a link to the
	// 'container'; i.e. a build may link to its builder, a builder to its
	// master/bucket, etc.

	code := common.ErrorCodeIn(err)
	switch code {
	case common.CodeUnknown:
		errors.Log(c.Context, err)
	case common.CodeUnauthorized:
		loginURL, err := auth.LoginURL(c.Context, c.Request.URL.RequestURI())
		if err == nil {
			http.Redirect(c.Writer, c.Request, loginURL, http.StatusFound)
			return
		}
		errors.Log(
			c.Context, errors.Annotate(err, "Failed to retrieve login URL").Err())
	}

	status := code.HTTPStatus()
	c.Writer.WriteHeader(status)
	templates.MustRender(c.Context, c.Writer, "pages/error.html", templates.Args{
		"Code":    status,
		"Message": err.Error(),
	})
}
