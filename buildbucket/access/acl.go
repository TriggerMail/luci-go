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

package access

import (
	"net/http"

	"golang.org/x/net/context"

	"go.chromium.org/luci/common/proto/access"
	"go.chromium.org/luci/grpc/prpc"
)

// Permissions represents a set of permitted actions for a set of buckets.
//
// Since Action is encoded as bits, the Action here is actually the bitwise OR
// of all available Actions for a particular bucket.
type Permissions map[string]Action

// Can checks whether an Action is allowed for a given bucket.
func (p Permissions) Can(bucket string, action Action) bool {
	return (p[bucket] & action) == action
}

// NewClient returns a new AccessClient which can be used to talk to
// buildbucket's access API.
func NewClient(host string, client *http.Client) access.AccessClient {
	return access.NewAccessPRPCClient(&prpc.Client{
		Host: host,
		C:    client,
	})
}

// BucketPermissions retrieves permitted actions for a set of buckets, for the
// identity specified in the client.
func BucketPermissions(c context.Context, client access.AccessClient, buckets []string) (Permissions, error) {
	// Make permitted actions call.
	req := access.PermittedActionsRequest{
		ResourceKind: "bucket",
		ResourceIds:  buckets,
	}
	resp, err := client.PermittedActions(c, &req)
	if err != nil {
		return nil, err
	}

	// Parse proto into a convenient format.
	perms := make(map[string]Action, len(resp.Permitted))
	for bucket, actions := range resp.Permitted {
		newActions := Action(0)
		for _, action := range actions.Actions {
			newAction, err := ParseAction(action)
			if err != nil {
				return nil, err
			}
			newActions |= newAction
		}
		perms[bucket] = newActions
	}
	return perms, nil
}