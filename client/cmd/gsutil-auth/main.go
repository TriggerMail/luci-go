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

// Command gsutil-auth implements "gsutil -> LUCI auth" shim server.
//
// Use it as command wrapper:
// $ gsutil-auth gsutil ...
package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/TriggerMail/luci-go/auth"
	"github.com/TriggerMail/luci-go/auth/client/authcli"
	"github.com/TriggerMail/luci-go/auth/integration/gsutil"
	"github.com/TriggerMail/luci-go/common/errors"
	"github.com/TriggerMail/luci-go/common/logging"
	"github.com/TriggerMail/luci-go/common/logging/gologger"
	"github.com/TriggerMail/luci-go/common/system/exitcode"

	"github.com/TriggerMail/luci-go/hardcoded/chromeinfra"
)

var (
	flags    authcli.Flags
	lifetime time.Duration
)

func init() {
	defaults := chromeinfra.DefaultAuthOptions()
	defaults.Scopes = []string{
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/userinfo.email",
	}
	flags.RegisterScopesFlag = true
	flags.Register(flag.CommandLine, defaults)
	flag.DurationVar(
		&lifetime, "lifetime", time.Minute, "Minimum token lifetime.",
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: gsutil-auth command...\n")
		flag.PrintDefaults()
	}
}

func setBotoConfigEnv(c *exec.Cmd, botoCfg string) {
	pfx := "BOTO_CONFIG="
	if c.Env == nil {
		c.Env = os.Environ()
	}
	for i, l := range c.Env {
		if strings.HasPrefix(strings.ToUpper(l), pfx) {
			c.Env[i] = pfx + botoCfg
			return
		}
	}
	c.Env = append(c.Env, pfx+botoCfg)
}

func run() error {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		return errors.New("specify a command to run")
	}

	opts, err := flags.Options()
	if err != nil {
		return err
	}

	bin := args[0]
	if filepath.Base(bin) == bin {
		path, err := exec.LookPath(bin)
		if err != nil {
			return errors.Reason("can't find %q in PATH", bin).Err()
		}
		bin = path
	}
	cmd := &exec.Cmd{
		Path:   bin,
		Args:   args,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	ctx := gologger.StdConfig.Use(context.Background())

	auth := auth.NewAuthenticator(ctx, auth.SilentLogin, opts)
	source, err := auth.TokenSource()
	if err != nil {
		return errors.Annotate(err, "failed to get token source").Err()
	}

	// State dir is used to hold .boto and temporary credentials cache.
	stateDir, err := ioutil.TempDir("", "gsutil-auth")
	if err != nil {
		return errors.Annotate(err, "failed to create gsutil state dir").Err()
	}
	defer os.RemoveAll(stateDir)

	srv := &gsutil.Server{
		Source:   source,
		StateDir: stateDir,
	}
	botoCfg, err := srv.Start(ctx)
	if err != nil {
		return errors.Annotate(err, "failed to start the gsutil auth server").Err()
	}
	defer srv.Stop(ctx) // close the server no matter what

	setBotoConfigEnv(cmd, botoCfg)

	// Return the subprocess exit code, if available.
	logging.Debugf(ctx, "Running %q", cmd.Args)
	switch code, hasCode := exitcode.Get(cmd.Run()); {
	case hasCode:
		os.Exit(code)
	case err != nil:
		return errors.Annotate(err, "command failed to start").Err()
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
