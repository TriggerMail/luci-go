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

package gs

import (
	"context"
	"net/http"
	"time"

	"google.golang.org/api/googleapi"

	"github.com/TriggerMail/luci-go/common/errors"
	"github.com/TriggerMail/luci-go/common/logging"
	"github.com/TriggerMail/luci-go/common/retry"
	"github.com/TriggerMail/luci-go/common/retry/transient"
)

var statusCodeTagKey = errors.NewTagKey("Google Storage API Status Code")

// StatusCode returns HTTP status code embedded inside the annotated error.
//
// Returns http.StatusOK if err is nil and 0 if the error doesn't have a status
// code.
func StatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	if val, ok := errors.TagValueIn(statusCodeTagKey, err); ok {
		return val.(int)
	}
	return 0
}

// StatusCodeTag can be used to attach HTTP status code to the error.
//
// This code will be available via StatusCode(err) function.
func StatusCodeTag(code int) errors.TagValue {
	return errors.TagValue{Key: statusCodeTagKey, Value: code}
}

// withRetry executes a Google Storage API call, retrying on transient errors.
//
// If request reached GS, but the service replied with an error, the
// corresponding HTTP status code can be extracted from the error via
// StatusCode(err). The error is also tagged as transient based on the code:
// response with HTTP statuses >=500 and 429 are considered transient errors.
//
// If the request never reached GS, StatusCode(err) would return 0 and the error
// will be tagged as transient.
func withRetry(c context.Context, call func() error) error {
	return retry.Retry(c, transient.Only(retry.Default), func() error {
		err := call()
		if err == nil {
			return nil
		}
		apiErr, _ := err.(*googleapi.Error)
		if apiErr == nil {
			// RestartUploadError errors are fatal and should be passed unannotated.
			if _, ok := err.(*RestartUploadError); ok {
				return err
			}
			return errors.Annotate(err, "failed to call GS").Tag(transient.Tag).Err()
		}
		ann := errors.Annotate(err, "GS replied with HTTP code %d", apiErr.Code).
			InternalReason("full response body:\n%s", apiErr.Body).
			Tag(StatusCodeTag(apiErr.Code))
		// Retry only on 429 and 5xx responses, according to
		// https://cloud.google.com/storage/docs/exponential-backoff.
		if apiErr.Code == 429 || apiErr.Code >= 500 {
			ann.Tag(transient.Tag)
		}
		return ann.Err()
	}, func(err error, d time.Duration) {
		logging.WithError(err).Errorf(c, "Transient error when accessing GS. Retrying in %s...", d)
	})
}
