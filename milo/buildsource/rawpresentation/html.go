// Copyright 2016 The LUCI Authors.
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

package rawpresentation

import (
	"context"
	"net/http"

	"github.com/TriggerMail/luci-go/common/data/stringset"
	"github.com/TriggerMail/luci-go/common/errors"
	"github.com/TriggerMail/luci-go/grpc/prpc"
	"github.com/TriggerMail/luci-go/hardcoded/chromeinfra"
	logdog "github.com/TriggerMail/luci-go/logdog/api/endpoints/coordinator/logs/v1"
	"github.com/TriggerMail/luci-go/logdog/client/coordinator"
	"github.com/TriggerMail/luci-go/server/auth"
)

// acceptableLogdogHosts is the (hard-coded) list of accepted logdog hosts.
var acceptableLogdogHosts = stringset.NewFromSlice(
	chromeinfra.LogDogHost,
	chromeinfra.LogDogHostAppSpot,
	chromeinfra.LogDogDevHost,
)

func resolveHost(host string) (string, error) {
	if host == "" {
		host = DefaultLogDogHost
	}
	if acceptableLogdogHosts.Has(host) {
		return host, nil
	}
	return "", errors.Reason("host %q is not whitelisted", host).Err()
}

var fakeLogKey = "holds a logdog.LogsClient"

// InjectFakeLogdogClient adds the given logdog.LogsClient to the context.
//
// You can obtain a fake logs client from
//   go.chromium.org/luci/logdog/api/endpoints/coordinator/logs/v1/fakelogs
//
// Injecting a nil logs client will panic.
func InjectFakeLogdogClient(c context.Context, client logdog.LogsClient) context.Context {
	if client == nil {
		panic("injecting nil logs client")
	}
	return context.WithValue(c, &fakeLogKey, client)
}

// NewClient generates a new LogDog client that issues requests on behalf of the
// current user.
func NewClient(c context.Context, host string) (*coordinator.Client, error) {
	if client, _ := c.Value(&fakeLogKey).(logdog.LogsClient); client != nil {
		return &coordinator.Client{
			C:    client,
			Host: "example.com",
		}, nil
	}

	var err error
	if host, err = resolveHost(host); err != nil {
		return nil, err
	}

	// Initialize the LogDog client authentication.
	t, err := auth.GetRPCTransport(c, auth.AsUser)
	if err != nil {
		return nil, errors.New("failed to get transport for LogDog server")
	}

	// Setup our LogDog client.
	return coordinator.NewClient(&prpc.Client{
		C: &http.Client{
			Transport: t,
		},
		Host: host,
	}), nil
}
