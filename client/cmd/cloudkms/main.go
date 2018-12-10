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

package main

import (
	"context"
	"os"

	"github.com/maruel/subcommands"

	cloudkms "google.golang.org/api/cloudkms/v1"

	"github.com/TriggerMail/luci-go/auth"
	"github.com/TriggerMail/luci-go/auth/client/authcli"
	"github.com/TriggerMail/luci-go/client/versioncli"
	"github.com/TriggerMail/luci-go/common/cli"
	"github.com/TriggerMail/luci-go/common/data/rand/mathrand"
	"github.com/TriggerMail/luci-go/common/logging/gologger"

	"github.com/TriggerMail/luci-go/hardcoded/chromeinfra"
)

// version must be updated whenever functional change (behavior, arguments,
// supported commands) is done.
const version = "0.1"

func GetApplication(defaultAuthOpts auth.Options) *cli.Application {
	defaultAuthOpts.Scopes = []string{auth.OAuthScopeEmail, cloudkms.CloudPlatformScope}
	return &cli.Application{
		Name:  "cloudkms",
		Title: "Client for interfacing with Google Cloud Key Management Service",
		Context: func(ctx context.Context) context.Context {
			return gologger.StdConfig.Use(ctx)
		},
		// Keep in alphabetical order of their name.
		Commands: []*subcommands.Command{
			subcommands.CmdHelp,
			cmdDecrypt(defaultAuthOpts),
			cmdEncrypt(defaultAuthOpts),
			authcli.SubcommandInfo(defaultAuthOpts, "whoami", false),
			authcli.SubcommandLogin(defaultAuthOpts, "login", false),
			authcli.SubcommandLogout(defaultAuthOpts, "logout", false),
			versioncli.CmdVersion(version),
		},
	}
}

func main() {
	mathrand.SeedRandomly()
	app := GetApplication(chromeinfra.DefaultAuthOptions())
	os.Exit(subcommands.Run(app, nil))
}
