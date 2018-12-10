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

package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/maruel/subcommands"

	"github.com/TriggerMail/luci-go/auth"
	"github.com/TriggerMail/luci-go/client/downloader"
	"github.com/TriggerMail/luci-go/client/internal/common"
	"github.com/TriggerMail/luci-go/common/isolated"
	"github.com/TriggerMail/luci-go/common/isolatedclient"
)

func cmdDownload(authOpts auth.Options) *subcommands.Command {
	return &subcommands.Command{
		UsageLine: "download <options>...",
		ShortDesc: "downloads a file or a .isolated tree from an isolate server.",
		LongDesc: `Downloads one or multiple files, or a isolated tree from the isolate server.

Files are referenced by their hash`,
		CommandRun: func() subcommands.CommandRun {
			c := downloadRun{}
			c.commonFlags.Init(authOpts)
			// TODO(mknyszek): Add support for downloading individual files.
			c.Flags.StringVar(&c.outputDir, "output-dir", ".", "The directory where files will be downloaded to.")
			c.Flags.StringVar(&c.outputFiles, "output-files", ".", "File into which the full list of downloaded files is written to.")
			c.Flags.StringVar(&c.isolated, "isolated", "", "Hash of a .isolated tree to download.")
			return &c
		},
	}
}

type downloadRun struct {
	commonFlags
	outputDir   string
	outputFiles string
	isolated    string
}

func (c *downloadRun) Parse(a subcommands.Application, args []string) error {
	if err := c.commonFlags.Parse(); err != nil {
		return err
	}
	if len(args) != 0 {
		return errors.New("position arguments not expected")
	}
	if c.isolated == "" {
		return errors.New("isolated is required")
	}
	return nil
}

func (c *downloadRun) main(a subcommands.Application, args []string) error {
	// Prepare isolated client.
	authClient, err := c.createAuthClient()
	if err != nil {
		return err
	}
	client := isolatedclient.New(nil, authClient, c.isolatedFlags.ServerURL, c.isolatedFlags.Namespace, nil, nil)
	ctx := context.Background()
	dl := downloader.New(ctx, client, 8)
	common.CancelOnCtrlC(dl)
	files, err := dl.FetchIsolated(isolated.HexDigest(c.isolated), c.outputDir)
	if err != nil {
		return err
	}
	if c.outputFiles != "" {
		filesData := strings.Join(files, "\n")
		if len(files) > 0 {
			filesData += "\n"
		}
		return ioutil.WriteFile(c.outputFiles, []byte(filesData), 0664)
	}
	return nil
}

func (c *downloadRun) Run(a subcommands.Application, args []string, _ subcommands.Env) int {
	if err := c.Parse(a, args); err != nil {
		fmt.Fprintf(a.GetErr(), "%s: %s\n", a.GetName(), err)
		return 1
	}
	cl, err := c.defaultFlags.StartTracing()
	if err != nil {
		fmt.Fprintf(a.GetErr(), "%s: %s\n", a.GetName(), err)
		return 1
	}
	defer cl.Close()
	if err := c.main(a, args); err != nil {
		fmt.Fprintf(a.GetErr(), "%s: %s\n", a.GetName(), err)
		return 1
	}
	return 0
}
