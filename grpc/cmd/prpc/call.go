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

package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/maruel/subcommands"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/TriggerMail/luci-go/auth"
	"github.com/TriggerMail/luci-go/common/cli"
	"github.com/TriggerMail/luci-go/grpc/prpc"
)

const (
	cmdCallUsage = `call [flags] <server> <service>.<method>

  server: host ("example.com") or port for localhost (":8080").
  service: full name of a service, e.g. "pkg.service"
  method: name of the method.
`

	cmdCallDesc = "calls a service method."
)

func cmdCall(defaultAuthOpts auth.Options) *subcommands.Command {
	return &subcommands.Command{
		UsageLine: cmdCallUsage,
		ShortDesc: cmdCallDesc,
		LongDesc: `Calls a service method.
The input message is read from stdin (defaulting to JSONPB)`,
		CommandRun: func() subcommands.CommandRun {
			c := &callRun{
				format: formatFlagJSONPB,
			}
			c.registerBaseFlags(defaultAuthOpts)
			c.Flags.Var(&c.format, "format", fmt.Sprintf(
				`Message format. Valid values: %s. Indicates both input and output format. The default is json.`,
				formatFlagMap.Choices()))
			return c
		},
	}
}

// callRun implements "call" subcommand.
type callRun struct {
	cmdRun
	format  formatFlag
	message string
}

func (r *callRun) Run(a subcommands.Application, args []string, env subcommands.Env) int {
	if len(args) < 2 {
		return r.argErr(cmdCallDesc, cmdCallUsage, "")
	}
	host, target := args[0], args[1]
	args = args[2:]

	req := request{
		format:       r.format,
		message:      os.Stdin,
		messageFlags: args,
	}

	var err error
	req.service, req.method, err = splitServiceAndMethod(target)
	if err != nil {
		return r.argErr(cmdCallDesc, cmdCallUsage, "%s", err)
	}

	ctx := cli.GetContext(a, r, env)
	client, err := r.authenticatedClient(ctx, host)
	if err != nil {
		return ecAuthenticatedClientError
	}
	return r.done(call(ctx, client, &req, os.Stdout))
}

func splitServiceAndMethod(fullName string) (service string, method string, err error) {
	lastDot := strings.LastIndex(fullName, ".")
	if lastDot < 0 {
		return "", "", fmt.Errorf("invalid full method name %q. It must contain a '.'", fullName)
	}
	service = fullName[:lastDot]
	method = fullName[lastDot+1:]
	return
}

// request is an RPC request.
type request struct {
	service      string
	method       string
	message      io.Reader
	messageFlags []string
	format       formatFlag
}

// call makes an RPC and writes response to out.
func call(c context.Context, client *prpc.Client, req *request, out io.Writer) error {
	var inf, outf prpc.Format
	var message []byte
	switch req.format {

	default:
		var buf bytes.Buffer
		if _, err := buf.ReadFrom(req.message); err != nil {
			return err
		}
		message = buf.Bytes()
		inf = req.format.Format()
		outf = inf
	}

	// Send the request.
	var hmd, tmd metadata.MD
	res, err := client.CallRaw(c, req.service, req.method, message, inf, outf, prpc.Header(&hmd), prpc.Trailer(&tmd))
	if err != nil {
		return &exitCode{err, int(grpc.Code(err))}
	}

	// Read response.
	if _, err := out.Write(res); err != nil {
		return fmt.Errorf("failed to write response: %s", err)
	}
	return err
}
