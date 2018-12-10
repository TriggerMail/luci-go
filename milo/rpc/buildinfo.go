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

package rpc

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"

	"github.com/TriggerMail/luci-go/common/errors"
	"github.com/TriggerMail/luci-go/common/logging"
	"github.com/TriggerMail/luci-go/config"
	"github.com/TriggerMail/luci-go/grpc/grpcutil"
	milo "github.com/TriggerMail/luci-go/milo/api/proto"
	"github.com/TriggerMail/luci-go/milo/buildsource/buildbot"
	"github.com/TriggerMail/luci-go/milo/buildsource/buildbucket"
	"github.com/TriggerMail/luci-go/milo/buildsource/swarming"
)

// BuildInfoService is a BuildInfoServer implementation.
type BuildInfoService struct {
	// Swarming is the BuildInfoProvider for the Swarming service.
	Swarming swarming.BuildInfoProvider
}

var _ milo.BuildInfoServer = (*BuildInfoService)(nil)

func (svc *BuildInfoService) getFromContextURI(c context.Context, id int64, projectHint string) (
	*milo.BuildInfoResponse, error) {
	bs, err := buildbucket.GetBuildSummary(c, id)
	if err != nil {
		return nil, errors.Annotate(err, "getting build summary").Err()
	}
	// Look for either:
	//    buildbot://<master>/build/<builder>/<number>
	//    swarming://<host>/task/<taskID>
	for _, uri := range bs.ContextURI {
		switch url, err := url.Parse(uri); {
		case err != nil:
			continue // Ignore invalid context URIs... not our problem.
		case url.Scheme == "buildbot" && strings.HasPrefix(url.Path, "/build/"):
			comp := strings.Split(url.Path, "/")
			if len(comp) != 4 {
				logging.Debugf(c, "invalid buildbot context uri: %s", uri)
				continue
			}
			number, err := strconv.ParseInt(comp[3], 10, 64)
			if err != nil {
				logging.Debugf(c, "invalid build number in: %s", uri)
			}
			req := &milo.BuildInfoRequest_BuildBot{
				MasterName:  url.Host,
				BuilderName: comp[2],
				BuildNumber: number,
			}
			return buildbot.GetBuildInfo(c, req, projectHint)
		case url.Scheme == "swarming" && strings.HasPrefix(url.Path, "/task/"):
			comp := strings.Split(url.Path, "/")
			if len(comp) != 3 {
				logging.Debugf(c, "invalid swarming context uri: %s", uri)
				continue
			}
			req := &milo.BuildInfoRequest_Swarming{
				Host: url.Host,
				Task: comp[2],
			}
			return svc.Swarming.GetBuildInfo(c, req, projectHint)
		}
	}
	logging.Debugf(c, "valid buildbot or swarming context not found in %s", bs.ContextURI)
	return nil, buildbucket.ErrNotFound
}

// Get implements milo.BuildInfoServer.
func (svc *BuildInfoService) Get(c context.Context, req *milo.BuildInfoRequest) (*milo.BuildInfoResponse, error) {
	projectHint := req.ProjectHint
	if projectHint != "" {
		if err := config.ValidateProjectName(projectHint); err != nil {
			return nil, grpcutil.Errf(codes.InvalidArgument, "invalid project hint: %s", err.Error())
		}
	}

	switch {
	case req.GetBuildbot() != nil:
		return buildbot.GetBuildInfo(c, req.GetBuildbot(), projectHint)

	case req.GetSwarming() != nil:
		return svc.Swarming.GetBuildInfo(c, req.GetSwarming(), projectHint)

	case req.GetBuildbucket() != nil:
		switch resp, err := svc.getFromContextURI(c, req.GetBuildbucket().GetId(), projectHint); err {
		case nil:
			return resp, nil
		case buildbucket.ErrNotFound:
			logging.WithError(err).Infof(c, "%d not found in context URI for build summary")
			// continue to fallback code.
		default:
			return nil, err
		}

		// Resolve the swarming host/task from buildbucket.
		host, taskID, err := buildbucket.GetSwarmingTaskID(c, strconv.FormatInt(req.GetBuildbucket().GetId(), 10))
		if err != nil {
			return nil, err
		}
		sReq := &milo.BuildInfoRequest_Swarming{
			Host: host,
			Task: taskID,
		}
		return svc.Swarming.GetBuildInfo(c, sReq, projectHint)

	default:
		return nil, grpcutil.Errf(codes.InvalidArgument, "must supply a build")
	}
}
