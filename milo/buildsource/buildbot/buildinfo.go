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

package buildbot

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"

	"github.com/TriggerMail/luci-go/common/data/stringset"
	"github.com/TriggerMail/luci-go/common/logging"
	miloProto "github.com/TriggerMail/luci-go/common/proto/milo"
	"github.com/TriggerMail/luci-go/grpc/grpcutil"
	"github.com/TriggerMail/luci-go/logdog/common/types"
	"github.com/TriggerMail/luci-go/milo/api/buildbot"
	milo "github.com/TriggerMail/luci-go/milo/api/proto"
	"github.com/TriggerMail/luci-go/milo/buildsource/buildbot/buildstore"
	"github.com/TriggerMail/luci-go/milo/buildsource/rawpresentation"
	"github.com/TriggerMail/luci-go/milo/common"
)

// GetBuildInfo resolves a Milo protobuf Step for a given BuildBot build.
//
// On failure, it returns a (potentially-wrapped) gRPC error.
//
// This:
//
//	1) Fetches the BuildBot build JSON from storage.
//	2) Resolves the LogDog annotation stream path from the BuildBot state.
//	3) Fetches the LogDog annotation stream and resolves it into a Step.
//	4) Merges some operational BuildBot build information into the Step.
func GetBuildInfo(c context.Context, req *milo.BuildInfoRequest_BuildBot,
	projectHint string) (*milo.BuildInfoResponse, error) {

	logging.Infof(c, "Loading build info for master %q, builder %q, build #%d",
		req.MasterName, req.BuilderName, req.BuildNumber)

	// Load the BuildBot build from storage.
	buildID := buildbot.BuildID{
		Master:  req.MasterName,
		Builder: req.BuilderName,
		Number:  int(req.BuildNumber),
	}
	build, err := buildstore.GetBuild(c, buildID)
	switch code := common.ErrorCodeIn(err); {
	case code == common.CodeUnauthorized:
		return nil, grpcutil.Unauthenticated

	case err != nil:
		logging.WithError(err).Errorf(c, "Failed to load build info.")
		return nil, grpcutil.Internal

	case build == nil:
		return nil, grpcutil.Errf(codes.NotFound, "Build #%d for master %q, builder %q was not found",
			req.BuildNumber, req.MasterName, req.BuilderName)
	}

	// Identify the LogDog annotation stream from the build.
	//
	// This will return a gRPC error on failure.
	addr, err := getLogDogAnnotationAddr(c, build)
	if err != nil {
		return nil, err
	}
	logging.Infof(c, "Resolved annotation stream: %s / %s", addr.Project, addr.Path)

	step, err := rawpresentation.ReadAnnotations(c, addr)
	if err != nil {
		logging.WithError(err).Errorf(c, "Failed to load annotation stream.")
		return nil, grpcutil.Errf(codes.Internal, "failed to load LogDog annotation stream from: %s", addr.Path)
	}

	// Merge the information together.
	if err := mergeBuildIntoAnnotation(c, step, build); err != nil {
		logging.WithError(err).Errorf(c, "Failed to merge annotation with build.")
		return nil, grpcutil.Errf(codes.Internal, "failed to merge annotation and build data")
	}

	prefix, name := addr.Path.Split()
	return &milo.BuildInfoResponse{
		Project: string(addr.Project),
		Step:    step,
		AnnotationStream: &miloProto.LogdogStream{
			Server: addr.Host,
			Prefix: string(prefix),
			Name:   string(name),
		},
	}, nil
}

// Resolve LogDog annotation stream for this build.
func getLogDogAnnotationAddr(c context.Context, build *buildbot.Build) (*types.StreamAddr, error) {
	if v, ok := build.PropertyValue("log_location").(string); ok && v != "" {
		return types.ParseURL(v)
	}
	return nil, grpcutil.Errf(codes.NotFound, "annotation stream not found")
}

// mergeBuildInfoIntoAnnotation merges BuildBot-specific build informtion into
// a LogDog annotation protobuf.
//
// This consists of augmenting the Step's properties with BuildBot's properties,
// favoring the Step's version of the properties if there are two with the same
// name.
func mergeBuildIntoAnnotation(c context.Context, step *miloProto.Step, build *buildbot.Build) error {
	allProps := stringset.New(len(step.Property) + len(build.Properties))
	for _, prop := range step.Property {
		allProps.Add(prop.Name)
	}
	for _, prop := range build.Properties {
		// Annotation protobuf overrides BuildBot properties.
		if allProps.Has(prop.Name) {
			continue
		}
		allProps.Add(prop.Name)

		step.Property = append(step.Property, &miloProto.Step_Property{
			Name:  prop.Name,
			Value: fmt.Sprintf("%v", prop.Value),
		})
	}

	return nil
}
