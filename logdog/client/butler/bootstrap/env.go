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

// Package bootstrap handles Butler-side bootstrapping functionality.
package bootstrap

import (
	"github.com/TriggerMail/luci-go/common/system/environ"
	"github.com/TriggerMail/luci-go/logdog/client/butlerlib/bootstrap"
	"github.com/TriggerMail/luci-go/logdog/common/types"
)

// Environment is the set of configuration parameters for the bootstrap.
type Environment struct {
	// CoordinatorHost is the coordinator host name, or empty string if we are
	// not outputting to a Coordinator.
	CoordinatorHost string

	// Project is the project name. If not empty, this will be exported to
	// subprocesses.
	Project types.ProjectName
	// Prefix is the prefix name. If not empty, this will be exported to
	// subprocesses.
	Prefix types.StreamName
	// StreamServerURI is the streamserver URI. If not empty, this will be
	// exported to subprocesses.
	StreamServerURI string
}

// Augment augments the supplied base environment with LogDog Butler bootstrap
// parameters.
func (e *Environment) Augment(base environ.Env) {
	exportIf := func(envKey, v string) {
		if v != "" {
			base.Set(envKey, v)
		}
	}

	exportIf(bootstrap.EnvCoordinatorHost, e.CoordinatorHost)
	exportIf(bootstrap.EnvStreamPrefix, string(e.Prefix))
	exportIf(bootstrap.EnvStreamProject, string(e.Project))
	exportIf(bootstrap.EnvStreamServerPath, e.StreamServerURI)
}
