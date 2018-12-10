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

package services

import (
	"context"
	"crypto/subtle"

	ds "go.chromium.org/gae/service/datastore"

	"github.com/TriggerMail/luci-go/common/clock"
	log "github.com/TriggerMail/luci-go/common/logging"
	"github.com/TriggerMail/luci-go/common/proto/google"
	"github.com/TriggerMail/luci-go/grpc/grpcutil"
	"github.com/TriggerMail/luci-go/logdog/api/config/svcconfig"
	"github.com/TriggerMail/luci-go/logdog/api/endpoints/coordinator/services/v1"
	"github.com/TriggerMail/luci-go/logdog/appengine/coordinator"
	"github.com/TriggerMail/luci-go/logdog/appengine/coordinator/config"
	"github.com/TriggerMail/luci-go/logdog/appengine/coordinator/endpoints"
	"github.com/TriggerMail/luci-go/logdog/appengine/coordinator/mutations"
	"github.com/TriggerMail/luci-go/tumble"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
)

// TerminateStream is an idempotent stream state terminate operation.
func (s *server) TerminateStream(c context.Context, req *logdog.TerminateStreamRequest) (*empty.Empty, error) {
	log.Fields{
		"project":       req.Project,
		"id":            req.Id,
		"terminalIndex": req.TerminalIndex,
	}.Infof(c, "Request to terminate log stream.")

	if req.TerminalIndex < 0 {
		return nil, grpcutil.Errf(codes.InvalidArgument, "Negative terminal index.")
	}

	id := coordinator.HashID(req.Id)
	if err := id.Normalize(); err != nil {
		return nil, grpcutil.Errf(codes.InvalidArgument, "Invalid ID (%s): %s", id, err)
	}

	// Load our service and project configs.
	svc := endpoints.GetServices(c)
	cfg, err := svc.Config(c)
	if err != nil {
		log.WithError(err).Errorf(c, "Failed to load configuration.")
		return nil, grpcutil.Internal
	}

	pcfg, err := coordinator.CurrentProjectConfig(c)
	if err != nil {
		log.WithError(err).Errorf(c, "Failed to load current project configuration.")
		return nil, grpcutil.Internal
	}

	// Initialize our archival parameters.
	params := standardArchivalParams(cfg, pcfg)

	// Initialize our log stream state.
	lst := coordinator.NewLogStreamState(c, id)

	// Transactionally validate and update the terminal index.
	err = ds.RunInTransaction(c, func(c context.Context) error {
		if err := ds.Get(c, lst); err != nil {
			if err == ds.ErrNoSuchEntity {
				log.Debugf(c, "Log stream state not found.")
				return grpcutil.Errf(codes.NotFound, "Log stream %q is not registered", id)
			}

			log.WithError(err).Errorf(c, "Failed to load LogEntry.")
			return grpcutil.Internal
		}

		switch {
		case subtle.ConstantTimeCompare(lst.Secret, req.Secret) != 1:
			log.Errorf(c, "Secrets do not match.")
			return grpcutil.Errf(codes.InvalidArgument, "Request secret doesn't match the stream secret.")

		case lst.Terminated():
			// Succeed if this is non-conflicting (idempotent).
			if lst.TerminalIndex == req.TerminalIndex {
				log.Fields{
					"terminalIndex": lst.TerminalIndex,
				}.Infof(c, "Log stream is already terminated.")
				return nil
			}

			log.Fields{
				"terminalIndex": lst.TerminalIndex,
			}.Warningf(c, "Log stream is already incompatibly terminated.")
			return grpcutil.Errf(codes.FailedPrecondition, "Log stream is incompatibly terminated.")

		default:
			// Everything looks good, let's proceed...
			now := clock.Now(c).UTC()
			lst.Updated = now
			lst.TerminalIndex = req.TerminalIndex
			lst.TerminatedTime = now

			if err := ds.Put(c, lst); err != nil {
				log.Fields{
					log.ErrorKey: err,
				}.Errorf(c, "Failed to Put() LogStream.")
				return grpcutil.Internal
			}

			// Replace the pessimistic archive expiration mutation scheduled in
			// RegisterStream with an optimistic archival mutation.
			cat := mutations.CreateArchiveTask{
				ID: id,

				// Optimistic parameters.
				SettleDelay:    params.SettleDelay,
				CompletePeriod: params.CompletePeriod,

				// Schedule this mutation to execute after our settle delay.
				Expiration: now.Add(params.SettleDelay),
			}

			aeParent, aeName := ds.KeyForObj(c, lst), cat.TaskName(c)
			if err := tumble.PutNamedMutations(c, aeParent, map[string]tumble.Mutation{aeName: &cat}); err != nil {
				log.WithError(err).Errorf(c, "Failed to replace archive expiration mutation.")
				return grpcutil.Internal
			}

			log.Fields{
				"terminalIndex":  lst.TerminalIndex,
				"settleDelay":    cat.SettleDelay,
				"completePeriod": cat.CompletePeriod,
				"scheduledAt":    cat.Expiration,
			}.Debugf(c, "Terminal index was set, and archival mutation was scheduled.")
			return nil
		}
	}, nil)
	if err != nil {
		log.Fields{
			log.ErrorKey: err,
		}.Errorf(c, "Failed to update LogStream.")
		return nil, err
	}

	return &empty.Empty{}, nil
}

func standardArchivalParams(cfg *config.Config, pcfg *svcconfig.ProjectConfig) *coordinator.ArchivalParams {
	return &coordinator.ArchivalParams{
		SettleDelay:    google.DurationFromProto(cfg.Coordinator.ArchiveSettleDelay),
		CompletePeriod: endpoints.MinDuration(cfg.Coordinator.ArchiveDelayMax, pcfg.MaxStreamAge),
	}
}
