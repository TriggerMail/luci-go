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

package bootstrap

import (
	"fmt"

	"github.com/TriggerMail/luci-go/common/errors"
	"github.com/TriggerMail/luci-go/common/system/environ"
	"github.com/TriggerMail/luci-go/logdog/client/butlerlib/streamclient"
	"github.com/TriggerMail/luci-go/logdog/common/types"
	"github.com/TriggerMail/luci-go/logdog/common/viewer"
)

// ErrNotBootstrapped is returned by Get when the current process is not
// bootstrapped.
var ErrNotBootstrapped = errors.New("not bootstrapped")

// Bootstrap contains information about the configured bootstrap environment.
//
// The bootstrap environment is loaded by probing the local application
// environment for variables emitted by a bootstrapping Butler.
type Bootstrap struct {
	// CoordinatorHost is the name of the upstream Coordinator host.
	//
	// This is just the host name ("example.appspot.com"), not a full URL.
	//
	// If this instance is not configured using a production Coordinator Output,
	// this will be empty.
	CoordinatorHost string

	// Project is the Butler instance project name.
	Project types.ProjectName
	// Prefix is the Butler instance prefix.
	Prefix types.StreamName

	// Client is the streamclient for this instance, or nil if the Butler has no
	// streamserver.
	Client streamclient.Client
}

func getFromEnv(env environ.Env, reg *streamclient.Registry) (*Bootstrap, error) {
	// Detect Butler by looking for EnvStreamPrefix in the envrironent.
	prefix, ok := env.Get(EnvStreamPrefix)
	if !ok {
		return nil, ErrNotBootstrapped
	}

	bs := &Bootstrap{
		CoordinatorHost: env.GetEmpty(EnvCoordinatorHost),
		Prefix:          types.StreamName(prefix),
		Project:         types.ProjectName(env.GetEmpty(EnvStreamProject)),
	}
	if err := bs.Prefix.Validate(); err != nil {
		return nil, fmt.Errorf("bootstrap: failed to validate prefix %q: %s", prefix, err)
	}
	if err := bs.Project.Validate(); err != nil {
		return nil, fmt.Errorf("bootstrap: failed to validate project %q: %s", bs.Project, err)
	}

	// If we have a stream server attached; instantiate a stream Client.
	if p, ok := env.Get(EnvStreamServerPath); ok {
		if err := bs.initializeClient(p, reg); err != nil {
			return nil, fmt.Errorf("bootstrap: failed to create stream client [%s]: %s", p, err)
		}
	}

	return bs, nil
}

func (bs *Bootstrap) initializeClient(v string, reg *streamclient.Registry) error {
	c, err := reg.NewClient(v)
	if err != nil {
		return errors.Annotate(err, "bootstrap: failed to create stream client [%s]", v).Err()
	}
	bs.Client = c
	return nil
}

// Get loads a Bootstrap instance from the environment. It will return an error
// if the bootstrap data is invalid, and will return ErrNotBootstrapped if the
// current process is not bootstrapped.
func Get() (*Bootstrap, error) {
	return getFromEnv(environ.System(), streamclient.GetDefaultRegistry())
}

// GetViewerURL returns a log stream viewer URL to the aggregate set of supplied
// stream paths.
//
// If both the Project and CoordinatorHost values are not populated, an error
// will be returned.
func (bs *Bootstrap) GetViewerURL(paths ...types.StreamPath) (string, error) {
	if bs.Project == "" {
		return "", errors.New("no project is configured")
	}
	if bs.CoordinatorHost == "" {
		return "", errors.New("no coordinator host is configured")
	}
	return viewer.GetURL(bs.CoordinatorHost, bs.Project, paths...), nil
}

// GetViewerURLForStreams returns a log stream viewer URL to the aggregate set
// of supplied streams.
//
// If the any of the Prefix, Project, or CoordinatorHost values is not
// populated, an error will be returned.
func (bs *Bootstrap) GetViewerURLForStreams(streams ...streamclient.Stream) (string, error) {
	if bs.Prefix == "" {
		return "", errors.New("no prefix is configured")
	}

	paths := make([]types.StreamPath, len(streams))
	for i, s := range streams {
		paths[i] = bs.Prefix.Join(types.StreamName(s.Properties().Name))
	}
	return bs.GetViewerURL(paths...)
}
