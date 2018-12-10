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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/maruel/subcommands"

	"github.com/TriggerMail/luci-go/auth"
	"github.com/TriggerMail/luci-go/client/archiver"
	"github.com/TriggerMail/luci-go/client/isolate"
	"github.com/TriggerMail/luci-go/common/data/text/units"
	"github.com/TriggerMail/luci-go/common/isolated"
	"github.com/TriggerMail/luci-go/common/isolatedclient"
)

const (
	// infraFailExit is the exit code used when the exparchive fails due to
	// infrastructure errors (for example, failed server requests).
	infraFailExit = 2
)

func cmdArchive(defaultAuthOpts auth.Options) *subcommands.Command {
	return &subcommands.Command{
		UsageLine: "archive <options>",
		ShortDesc: "parses a .isolate file to create a .isolated file, and uploads it and all referenced files to an isolate server",
		LongDesc:  "All the files listed in the .isolated file are put in the isolate server cache. Small files are combined together in a tar archive before uploading.",
		CommandRun: func() subcommands.CommandRun {
			c := archiveRun{}
			c.commonServerFlags.Init(defaultAuthOpts)
			c.isolateFlags.Init(&c.Flags)
			c.loggingFlags.Init(&c.Flags)
			c.Flags.BoolVar(&c.expArchive, "exparchive", true, "IGNORED (deprecated) Whether to use the new exparchive implementation, which tars small files before uploading them.")
			c.Flags.IntVar(&c.maxConcurrentChecks, "max-concurrent-checks", 1, "The maxiumum number of in-flight check requests.")
			c.Flags.IntVar(&c.maxConcurrentUploads, "max-concurrent-uploads", 8, "The maximum number of in-flight uploads.")
			c.Flags.StringVar(&c.dumpJSON, "dump-json", "",
				"Write isolated digests of archived trees to this file as JSON")
			return &c
		},
	}
}

type archiveRun struct {
	commonServerFlags
	isolateFlags
	loggingFlags         loggingFlags
	expArchive           bool // deprecated
	maxConcurrentChecks  int
	maxConcurrentUploads int
	dumpJSON             string
}

func (c *archiveRun) Parse(a subcommands.Application, args []string) error {
	if err := c.commonServerFlags.Parse(); err != nil {
		return err
	}
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	if err := c.isolateFlags.Parse(cwd, RequireIsolateFile&RequireIsolatedFile); err != nil {
		return err
	}
	if len(args) != 0 {
		return errors.New("position arguments not expected")
	}
	return nil
}

func (c *archiveRun) main(a subcommands.Application, args []string) error {
	start := time.Now()
	authCl, err := c.createAuthClient()
	if err != nil {
		return err
	}
	ctx := c.defaultFlags.MakeLoggingContext(os.Stderr)
	client := isolatedclient.New(nil, authCl, c.isolatedFlags.ServerURL, c.isolatedFlags.Namespace, nil, nil)

	al := archiveLogger{
		start: start,
		quiet: c.defaultFlags.Quiet,
	}

	return archive(ctx, client, &c.ArchiveOptions, c.dumpJSON, c.maxConcurrentChecks, c.maxConcurrentUploads, al)
}

// archiveLogger reports stats to stderr.
type archiveLogger struct {
	start time.Time
	quiet bool
}

// LogSummary logs (to eventlog and stderr) a high-level summary of archive operations(s).
func (al *archiveLogger) LogSummary(ctx context.Context, hits, misses int64, bytesHit, bytesPushed units.Size, digests []string) {
	end := time.Now()

	if !al.quiet {
		duration := end.Sub(al.start)
		fmt.Fprintf(os.Stderr, "Hits    : %5d (%s)\n", hits, bytesHit)
		fmt.Fprintf(os.Stderr, "Misses  : %5d (%s)\n", misses, bytesPushed)
		fmt.Fprintf(os.Stderr, "Duration: %s\n", units.Round(duration, time.Millisecond))
	}
}

// Print acts like fmt.Printf, but may prepend a prefix to format, depending on the value of al.quiet.
func (al *archiveLogger) Printf(format string, a ...interface{}) (n int, err error) {
	return al.Fprintf(os.Stdout, format, a...)
}

// Print acts like fmt.fprintf, but may prepend a prefix to format, depending on the value of al.quiet.
func (al *archiveLogger) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	prefix := "\n"
	if al.quiet {
		prefix = ""
	}
	args := []interface{}{prefix}
	args = append(args, a...)
	return fmt.Printf("%s"+format, args...)
}

// archive performs the archive operation for an isolate specified by opts.
// dumpJSON is the path to write a JSON summary of the uploaded isolate, in the same format as batch_archive.
func archive(ctx context.Context, client *isolatedclient.Client, opts *isolate.ArchiveOptions, dumpJSON string, concurrentChecks, concurrentUploads int, al archiveLogger) error {
	// Parse the incoming isolate file.
	deps, rootDir, isol, err := isolate.ProcessIsolate(opts)
	if err != nil {
		return fmt.Errorf("isolate %s: failed to process: %v", opts.Isolate, err)
	}
	log.Printf("Isolate %s referenced %d deps", opts.Isolate, len(deps))

	// Set up a checker and uploader.
	checker := archiver.NewChecker(ctx, client, concurrentChecks)
	uploader := archiver.NewUploader(ctx, client, concurrentUploads)
	archiver := archiver.NewTarringArchiver(checker, uploader)

	isolSummary, err := archiver.Archive(deps, rootDir, isol, opts.Blacklist, opts.Isolated)
	if err != nil {
		return fmt.Errorf("isolate %s: %v", opts.Isolate, err)
	}

	// Make sure that all pending items have been checked.
	if err := checker.Close(); err != nil {
		return err
	}

	// Make sure that all the uploads have completed successfully.
	if err := uploader.Close(); err != nil {
		return err
	}

	printSummary(al, isolSummary)
	if err := dumpSummaryJSON(dumpJSON, isolSummary); err != nil {
		return err
	}

	al.LogSummary(ctx, int64(checker.Hit.Count()), int64(checker.Miss.Count()), units.Size(checker.Hit.Bytes()), units.Size(checker.Miss.Bytes()), []string{string(isolSummary.Digest)})
	return nil
}

func (c *archiveRun) Run(a subcommands.Application, args []string, _ subcommands.Env) int {
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

// CancelOnCtrlC is a temporary copy of the CancelOnCtrlC in internal/common/concurrent.go
// This is needed until the old archive and batcharchive code (which uses Cancelers) is removed.
// It operates on a concrete Archiver to avoid the dependency on Canceler.
func CancelOnCtrlC(arch *archiver.Archiver) {
	interrupted := make(chan os.Signal, 1)
	signal.Notify(interrupted, os.Interrupt)
	go func() {
		defer signal.Stop(interrupted)
		select {
		case <-interrupted:
			arch.Cancel(errors.New("Ctrl-C"))
		case <-arch.Channel():
		}
	}()
}

func printSummary(al archiveLogger, summary archiver.IsolatedSummary) {
	al.Printf("%s\t%s\n", summary.Digest, summary.Name)
}

func dumpSummaryJSON(filename string, summaries ...archiver.IsolatedSummary) error {
	if len(filename) == 0 {
		return nil
	}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	m := map[string]isolated.HexDigest{}
	for _, summary := range summaries {
		m[summary.Name] = summary.Digest
	}
	return json.NewEncoder(f).Encode(m)
}
