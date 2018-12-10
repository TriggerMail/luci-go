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

package butler

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sort"
	"sync"
	"time"

	"github.com/TriggerMail/luci-go/common/clock"
	"github.com/TriggerMail/luci-go/common/data/stringset"
	"github.com/TriggerMail/luci-go/common/iotools"
	log "github.com/TriggerMail/luci-go/common/logging"
	"github.com/TriggerMail/luci-go/common/proto/google"
	"github.com/TriggerMail/luci-go/common/runtime/paniccatcher"
	"github.com/TriggerMail/luci-go/common/sync/parallel"
	"github.com/TriggerMail/luci-go/logdog/client/butler/bundler"
	"github.com/TriggerMail/luci-go/logdog/client/butler/output"
	"github.com/TriggerMail/luci-go/logdog/client/butler/streamserver"
	"github.com/TriggerMail/luci-go/logdog/client/butlerlib/streamproto"
	"github.com/TriggerMail/luci-go/logdog/common/types"
)

const (
	// DefaultMaxBufferAge is the default amount of time that a log entry may
	// be buffered before being dispatched.
	DefaultMaxBufferAge = time.Duration(5 * time.Second)

	// DefaultOutputWorkers is the default number of output workers to use.
	DefaultOutputWorkers = 16

	// streamBufferSize is the maximum amount of stream data to buffer in memory.
	streamBufferSize = 1024 * 1024 * 5

	// keepAliveChannelSize is the size of the I/O keep-alive channel buffer.
	keepAliveChannelSize = 128
)

// ioKeepAliveMessage is a message that will be sent when I/O keep-alive is
// configured (see keepAliveMonitor).
var ioKeepAliveMessage = []byte("LogDog Butler: I/O Keep-alive.\n")

// Config is the set of Butler configuration parameters.
type Config struct {
	// Output is the output instance to use for log dispatch.
	Output output.Output
	// OutputWorkers is the number of simultaneous goroutines that will be used
	// to output Butler log data. If zero, DefaultOutputWorkers will be used.
	OutputWorkers int

	// Project is the project that the log stream will be bound to.
	Project types.ProjectName
	// Prefix is the log stream common prefix value.
	Prefix types.StreamName

	// GlobalTags are a set of global log stream tags to apply to individual
	// streams on registration. Individual stream tags will override tags with
	// the same key.
	GlobalTags streamproto.TagMap

	// BufferLogs, if true, instructs the butler to buffer collected log data
	// before sending it to Output.
	BufferLogs bool
	// If buffering logs, this is the maximum amount of time that a log will
	// be buffered before being marked for dispatch. If this is zero,
	// DefaultMaxBufferAge will be used.
	MaxBufferAge time.Duration

	// TeeStdout, if not nil, is the Writer that will be used for streams
	// requesting STDOUT tee.
	TeeStdout io.Writer
	// TeeStderr, if not nil, is the Writer that will be used for streams
	// requesting STDERR tee.
	TeeStderr io.Writer

	// IOKeepAliveInterval configures the Butler to perform I/O pings.
	//
	// The Butler will monitor I/O events on all of its subordinate streams. If
	// a stream receives I/O, and no streams marked "IOKeepAlive" have written
	// I/O in this period, an I/O keep-alive event will be written to all
	// "IOKeepAlive" streams.
	IOKeepAliveInterval time.Duration
	// IOKeepAliveWriter is an io.Writer to send keep-alive updates through. This
	// must be set for I/O keep-alive to be active.
	IOKeepAliveWriter io.Writer
}

// Validate validates that the configuration is sufficient to instantiate a
// Butler instance.
func (c *Config) Validate() error {
	if c.Output == nil {
		return errors.New("butler: an Output must be supplied")
	}
	if err := c.Project.Validate(); err != nil {
		return fmt.Errorf("invalid project: %v", err)
	}
	if err := c.Prefix.Validate(); err != nil {
		return fmt.Errorf("invalid prefix: %v", err)
	}
	return nil
}

// Butler is the Butler application structure. The Butler runs until closed.
// During operation, it acts as a service manager and data router, routing:
// - Messages from Streams to the attached Output.
// - Streams from a StreamServer to the Stream list (AddStream).
type Butler struct {
	c   *Config
	ctx context.Context

	// bundler is the Bundler instance.
	bundler *bundler.Bundler
	// bundlerDrainedC is a signal channel that is closed when the Bundler has
	// been drained.
	bundlerDrainedC chan struct{}

	// activateC is closed when Activate() is called.
	activateC chan struct{}
	// activateOnce ensures we close activeC exactly once.
	activateOnce sync.Once

	// streamsFinishedC is a signal channel that will be closed when the stream
	// monitor finishes its managed stream set.
	streamsFinishedC chan struct{}

	// WaitGroup to manage running StreamServers.
	streamServerWG sync.WaitGroup
	// Channel to signal StreamServers to stop.
	streamServerStopC chan struct{}

	streamC chan *stream
	// streamSeen tracks if a stream has been seen. This is used to prevent
	// duplicate stream names from being created.
	streamSeen stringset.Set
	// streamSeenLock is a lock to protect streamSeen.
	streamSeenLock sync.Mutex

	// keepAliveC, if not nil, will receive I/O keep-alive events.
	//
	// If the boolean is true, it is an IOKeepAlive stream signalling that I/O
	// data has been written to it.
	//
	// If the boolean is false, it is a non-IOKeepAlive stream signalling that
	// I/O data has been written to it.
	keepAliveC chan bool
	// keepAliveFinishedC is closed when the I/O Keep Alive monitor is finished.
	keepAliveFinishedC chan struct{}

	// shutdownMu is a mutex to protect shutdown parameters.
	shutdownMu sync.Mutex
	// isShutdown is true if the Butler been shut down.
	isShutdown bool
	// runErr is the error returned by Run.
	runErr error
	// streamStopC is a stop signal channel for stream. This will cause streams
	// to prematurely terminate (before EOF) on shutdown.
	streamStopC chan struct{}
}

// New instantiates a new Butler instance and starts its processing.
func New(ctx context.Context, config Config) (*Butler, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	if config.OutputWorkers <= 0 {
		config.OutputWorkers = DefaultOutputWorkers
	}

	bc := bundler.Config{
		Clock:            clock.Get(ctx),
		Project:          config.Project,
		Prefix:           config.Prefix,
		MaxBufferedBytes: streamBufferSize,
		MaxBundleSize:    config.Output.MaxSize(),
	}
	if config.BufferLogs {
		bc.MaxBufferDelay = config.MaxBufferAge
		if bc.MaxBufferDelay <= 0 {
			bc.MaxBufferDelay = DefaultMaxBufferAge
		}
	}
	lb := bundler.New(bc)

	b := &Butler{
		c:   &config,
		ctx: ctx,

		bundler:         lb,
		bundlerDrainedC: make(chan struct{}),

		streamsFinishedC: make(chan struct{}),

		activateC:         make(chan struct{}),
		streamC:           make(chan *stream),
		streamSeen:        stringset.New(0),
		streamServerStopC: make(chan struct{}),
		streamStopC:       make(chan struct{}),
	}

	// Load bundles from our Bundler into the queue.
	go func() {
		defer close(b.bundlerDrainedC)
		parallel.Ignore(parallel.Run(config.OutputWorkers, func(workC chan<- func() error) {
			// Read bundles until the bundler is drained.
			for {
				bundle := b.bundler.Next()
				if bundle == nil {
					return
				}

				workC <- func() error {
					b.c.Output.SendBundle(bundle)
					return nil
				}
			}
		}))
	}()

	// Run our primary stream monitor until the Butler instance is activated.
	go func() {
		defer close(b.streamsFinishedC)
		b.runStreams(b.activateC)
	}()

	// Shutdown our Butler if our Context is cancelled.
	go func() {
		select {
		case <-b.streamStopC:
			break
		case <-b.ctx.Done():
			log.WithError(b.ctx.Err()).Warningf(b.ctx, "Butler Context was cancelled. Initiating shutdown.")
			b.shutdown(b.ctx.Err())
		}
	}()

	// If we're doing I/O keep-alive, kick off our I/O keep-alive reactor.
	if b.c.IOKeepAliveInterval > 0 && b.c.IOKeepAliveWriter != nil {
		b.keepAliveC = make(chan bool, keepAliveChannelSize)
		b.keepAliveFinishedC = make(chan struct{})
		go func() {
			defer close(b.keepAliveFinishedC)
			b.keepAliveMonitor(b.c.IOKeepAliveInterval, b.c.IOKeepAliveWriter)
		}()
	}

	return b, nil
}

// Wait blocks until the Butler instance has completed, returning with the
// Butler's return code.
func (b *Butler) Wait() error {
	// Run until our stream monitor shuts down, meaning all streams have finished.
	//
	// A race can exist here when a stream server may add new streams after we've
	// drained our streams, but before we've shut them down. Since "all streams
	// are done" is the edge that we use to begin shutdown, we can't simply tell
	// stream servers to stop in advance.
	//
	// We manage this race as follows:
	// 1) Wait until our stream monitor finishes. This will happen when there's a
	//    point that no streams are running.
	// 2) Start a new stream monitor to handle (3).
	// 3) Initiate stream server shutdown, wait until they have all finished.
	// 4) Wait until the stream monitor in (2) has finished.
	log.Debugf(b.ctx, "Waiting for Butler primary stream monitor to finish...")
	<-b.streamsFinishedC
	log.Debugf(b.ctx, "Butler streams have finished.")

	log.Debugf(b.ctx, "Shutting down stream servers, starting residual stream monitor.")

	auxStreamsFinishedC := make(chan struct{})
	auxActivateC := make(chan struct{})
	go func() {
		defer close(auxStreamsFinishedC)
		b.runStreams(auxActivateC)
	}()

	close(b.streamServerStopC)
	b.streamServerWG.Wait()
	log.Debugf(b.ctx, "Stream servers have shut down.")

	log.Debugf(b.ctx, "Waiting for residual streams to finish...")
	close(b.streamC)
	close(auxActivateC)
	<-auxStreamsFinishedC
	log.Debugf(b.ctx, "Residual streams have finished.")

	log.Debugf(b.ctx, "Waiting for bundler to flush.")
	b.bundler.CloseAndFlush()
	log.Debugf(b.ctx, "Bundler has flushed.")

	if b.keepAliveC != nil {
		log.Debugf(b.ctx, "Closing keep-alive monitor.")
		close(b.keepAliveC)
		<-b.keepAliveFinishedC
		log.Debugf(b.ctx, "Keep-alive monitor has finished.")
	}

	log.Debugf(b.ctx, "Waiting for output queue to shut down.")
	<-b.bundlerDrainedC
	log.Debugf(b.ctx, "Output queue has shut down.")

	log.Fields{
		"stats": b.c.Output.Stats(),
	}.Infof(b.ctx, "Message output has closed")
	return b.getRunErr()
}

// Streams returns a sorted list of stream names that have been registered to
// the Butler.
func (b *Butler) Streams() []types.StreamName {
	var streams types.StreamNameSlice
	func() {
		b.streamSeenLock.Lock()
		defer b.streamSeenLock.Unlock()

		streams = make([]types.StreamName, 0, b.streamSeen.Len())
		b.streamSeen.Iter(func(s string) bool {
			streams = append(streams, types.StreamName(s))
			return true
		})
	}()

	sort.Sort(streams)
	return ([]types.StreamName)(streams)
}

// AddStreamServer adds a StreamServer to the Butler. This is goroutine-safe
// and may be called anytime before or during Butler execution.
//
// After this call completes, the Butler assumes ownership of the StreamServer.
func (b *Butler) AddStreamServer(streamServer streamserver.StreamServer) {
	ctx := log.SetField(b.ctx, "streamServer", streamServer)

	log.Debugf(ctx, "Adding stream server.")

	// Pull streams from the streamserver and add them to the Butler.
	streamServerFinishedC := make(chan struct{})
	go func() {
		defer close(streamServerFinishedC)

		defer paniccatcher.Catch(func(p *paniccatcher.Panic) {
			log.Fields{
				"panic.error": p.Reason,
			}.Errorf(b.ctx, "Panic while running StreamServer:\n%s", p.Stack)
			b.shutdown(fmt.Errorf("butler: panic while running StreamServer: %v", p.Reason))
		})

		for {
			rc, config := streamServer.Next()
			if rc == nil {
				log.Debugf(ctx, "StreamServer returned nil stream; terminating.")
				return
			}

			// Add this Stream to the Butler.
			//
			// We run this in a function so we can ensure cleanup on failure.
			if err := b.AddStream(rc, config); err != nil {
				log.Fields{
					log.ErrorKey: err,
				}.Errorf(ctx, "Failed to add stream.")

				if err := rc.Close(); err != nil {
					log.Fields{
						log.ErrorKey: err,
					}.Warningf(ctx, "Failed to close stream.")
				}
			}
		}
	}()

	// Monitor the StreamServer's close signal channel; terminate our server when
	// it's set.
	b.streamServerWG.Add(1)
	go func() {
		defer b.streamServerWG.Done()

		<-b.streamServerStopC
		log.Debugf(ctx, "Stop signal received for StreamServer.")
		streamServer.Close()
		<-streamServerFinishedC
	}()
}

// AddStream adds a Stream to the Butler. This is goroutine-safe.
//
// If no error is returned, the Butler assumes ownership of the supplied stream.
// The stream will be closed when processing is finished.
//
// If an error is occurred, the caller is still the owner of the stream and
// is responsible for closing it.
func (b *Butler) AddStream(rc io.ReadCloser, p *streamproto.Properties) error {
	p = p.Clone()
	if p.Timestamp == nil || google.TimeFromProto(p.Timestamp).IsZero() {
		p.Timestamp = google.NewTimestamp(clock.Now(b.ctx))
	}
	if err := p.Validate(); err != nil {
		return err
	}

	// Build per-stream tag map.
	if l := len(b.c.GlobalTags); l > 0 {
		if p.Tags == nil {
			p.Tags = make(map[string]string, l)
		}

		for k, v := range b.c.GlobalTags {
			// Add only global flags that aren't already present (overridden) in
			// stream tags.
			if _, ok := p.Tags[k]; !ok {
				p.Tags[k] = v
			}
		}
	}

	if p.Timeout > 0 {
		if rts, ok := rc.(iotools.ReadTimeoutSetter); ok {
			if err := rts.SetReadTimeout(p.Timeout); err != nil {
				log.Fields{
					log.ErrorKey: err,
					"timeout":    p.Timeout,
				}.Warningf(b.ctx, "Failed to set stream timeout.")
			}
		} else {
			log.Fields{
				"connection":     rc,
				"connectionType": fmt.Sprintf("%T", rc),
			}.Warningf(b.ctx, "Don't know how to set timeout for type, so ignoring Timeout parameter.")
		}
	}

	// If this stream is configured to tee, set that up.
	reader := io.Reader(rc)
	isKeepAlive := false
	switch p.Tee {
	case streamproto.TeeNone:
		break

	case streamproto.TeeStdout:
		if b.c.TeeStdout == nil {
			return errors.New("butler: cannot tee through STDOUT; no STDOUT is configured")
		}
		reader = io.TeeReader(rc, b.c.TeeStdout)
		isKeepAlive = true

	case streamproto.TeeStderr:
		if b.c.TeeStderr == nil {
			return errors.New("butler: cannot tee through STDERR; no STDERR is configured")
		}
		reader = io.TeeReader(rc, b.c.TeeStderr)
		isKeepAlive = true

	default:
		return fmt.Errorf("invalid tee value: %v", p.Tee)
	}

	if err := b.registerStream(p.Name); err != nil {
		return err
	}

	// Build our stream struct.
	streamCtx := log.SetField(b.ctx, "stream", p.Name)
	s := stream{
		Context:     streamCtx,
		r:           reader,
		c:           rc,
		isKeepAlive: isKeepAlive,
	}

	// Register this stream with our Bundler. It will take ownership of "p", so
	// we should not use it after this point.
	var err error
	if s.bs, err = b.bundler.Register(p); err != nil {
		return err
	}
	p = nil

	b.streamC <- &s
	return nil
}

func (b *Butler) runStreams(activateC chan struct{}) {
	streamFinishedC := make(chan *stream)
	streamC := b.streamC

	activeCount := 0
	for {
		select {
		case s, ok := <-streamC:
			if !ok {
				// Our streamC has been closed. At this point, we wait for current
				// streams to finish.
				streamC = nil
				continue
			}

			// Monitor goroutine to respond to shutdown signal and clean up stream.
			log.Debugf(s, "Adding stream.")
			activeCount++

			closeOnce := sync.Once{}
			closeStream := func() {
				closeOnce.Do(s.closeStream)
			}

			go func() {
				defer func() {
					// Report that our stream has finished.
					streamFinishedC <- s
				}()
				defer closeStream()

				// Read the stream continuously until we're finished or interrupted.
				finishedC := make(chan struct{})
				go func(s *stream) {
					defer close(finishedC)

					for s.readChunk() {
						// If our keep-alive monitor is enabled, send a keep-alive event.
						if b.keepAliveC != nil {
							b.keepAliveC <- s.isKeepAlive
						}
					}
				}(s)

				// Stop processing when either the stream is finished or we are instructed
				// to close the stream via 'streamStopC'.
				select {
				case <-finishedC:
					// The stream has finished on its own.
					break

				case <-b.streamStopC:
					log.Debugf(s, "Received stop signal.")
					closeStream()
				}

				// Wait for our stream to finish.
				<-finishedC
			}()

		case <-streamFinishedC:
			// A stream has reported that it finished.
			//
			// If this is the last active stream and we've been activated, exit the
			// monitor.
			activeCount--
			if activeCount == 0 && activateC == nil {
				return
			}

		case <-activateC:
			// If we're not managing any streams, then we're done.
			if activeCount == 0 {
				return
			}

			// Record that we've been activated. Clearing the channel stops the select
			// from hammering the "activateC" case in the period after activation but
			// before our streams are finished.
			activateC = nil
		}
	}
}

// registerStream registers awareness of the named Stream with the Butler. An
// error will be returned if the Stream has ever been registered.
func (b *Butler) registerStream(name string) error {
	b.streamSeenLock.Lock()
	defer b.streamSeenLock.Unlock()

	if added := b.streamSeen.Add(name); added {
		return nil
	}
	return fmt.Errorf("a stream has already been registered with name %q", name)
}

func (b *Butler) keepAliveMonitor(interval time.Duration, w io.Writer) {
	t := clock.NewTimer(b.ctx)
	defer t.Stop()

	// Begin our keep-alive countdown.
	t.Reset(interval)
	timerRunning := true

	for {
		select {
		case event, ok := <-b.keepAliveC:
			if !ok {
				// Our keep-alive monitor is finished.
				return
			}
			if event {
				// If "event", we have been signalled by an tee stream that I/O has
				// arrived. Clear our timer, since teeing data inherently fulfills the
				// same role as the keep-alive event.
				t.Stop()
			} else if !timerRunning {
				// If !event, data has come in on a non-tee stream. Make sure our timer
				// is running to represent that data was received.
				t.Reset(interval)
				timerRunning = true
			}

		case <-t.GetC():
			// Keep-alive timer has timed out, write keep-alive.
			timerRunning = false

			// Send keep-alive through our keep-alive Writer.
			if w != nil {
				if _, err := w.Write(ioKeepAliveMessage); err != nil {
					log.WithError(err).Errorf(b.ctx, "Failed to send I/O keep-alive message.")
				}
			}
		}
	}
}

// Activate notifies the Butler that its current stream load is sufficient.
// This enables it to exit Run when it reaches a stream count of zero. Prior
// to activation, the Butler would block in Run regardless of stream count.
func (b *Butler) Activate() {
	b.activateOnce.Do(func() {
		close(b.activateC)
	})
}

// shutdown is a goroutine-safe method instructing the Butler to terminate
// with the supplied error code. It may be called more than once, although
// the first supplied error message will be the one returned by Run.
func (b *Butler) shutdown(err error) {
	log.Fields{
		log.ErrorKey: err,
	}.Debugf(b.ctx, "Received 'shutdown()' command; shutting down streams.")

	func() {
		b.shutdownMu.Lock()
		defer b.shutdownMu.Unlock()

		if b.isShutdown {
			// Already shut down.
			return
		}

		// Signal our streams to shutdown prematurely.
		close(b.streamStopC)

		b.runErr = err
		b.isShutdown = true
	}()

	// Activate the Butler, if it hasn't already been activated. The Butler will
	// block pending stream draining, but we've instructed our streams to
	// shutdown prematurely, so this should be reasonably quick.
	b.Activate()
}

// Returns the configured Butler error.
func (b *Butler) getRunErr() error {
	b.shutdownMu.Lock()
	defer b.shutdownMu.Unlock()
	return b.runErr
}
