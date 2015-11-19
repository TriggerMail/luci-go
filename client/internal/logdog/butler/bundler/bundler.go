// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package bundler

import (
	"container/heap"
	"fmt"
	"sync"
	"time"

	"github.com/luci/luci-go/client/logdog/butlerlib/streamproto"
	"github.com/luci/luci-go/common/clock"
	"github.com/luci/luci-go/common/logdog/protocol"
	"github.com/luci/luci-go/common/logdog/types"
	"github.com/luci/luci-go/common/proto/google"
)

// Config is the Bundler configuration.
type Config struct {
	// Clock is the clock instance that will be used for Bundler and stream
	// timing.
	Clock clock.Clock

	// Source is the bundle source string to use. This can be empty if there is no
	// source information to include.
	Source string

	// MaxBufferedBytes is the maximum number of bytes to buffer in memory per
	// stream.
	MaxBufferedBytes int64

	// MaxBundleSize is the maximum bundle size in bytes that may be generated.
	//
	// If this value is zero, no size constraint will be applied to generated
	// bundles.
	MaxBundleSize int

	// MaxBufferDelay is the maximum amount of time we're willing to buffer
	// bundled data. Other factors can cause the bundle to be sent before this,
	// but it is an upper bound.
	MaxBufferDelay time.Duration
}

type bundlerStream interface {
	isDrained() bool
	name() string
	expireTime() (time.Time, bool)
	nextBundleEntry(*builder, bool) bool
}

// Bundler is the main Bundler instance. It exposes goroutine-safe endpoints for
// stream registration and bundle consumption.
type Bundler interface {
	Register(streamproto.Properties) (Stream, error)

	// Next causes the Bundler to pack a new ButlerLogBundle message for
	// transmission.
	//
	// The maximum marshalled byte size of the bundle's protobuf is supplied; the
	// generated bundle will not exceed this size.
	Next() *protocol.ButlerLogBundle

	// CloseAndFlush flushes the Bundler, blocking until all registered Streams
	// have closed and all buffered data has been output.
	CloseAndFlush()
}

type bundlerImpl struct {
	*Config

	// finishedC is closed when makeBundles goroutine has terminated.
	finishedC chan struct{}
	bundleC   chan *protocol.ButlerLogBundle

	// streamsLock is a lock around the `streams` map and its contents.
	streamsLock sync.Mutex
	// streamsCond is a Cond bound to streamsLock, used to signal Next() when new
	// data is available.
	streamsCond *timeoutCond
	// streams is the set of currently-registered Streams.
	streams map[string]bundlerStream
	// flushing is true if we're blocking on CloseAndFlush().
	flushing bool

	// prefixCounter is a global counter for Prefix-wide streams.
	prefixCounter counter
}

// New instantiates a new Bundler instance.
func New(c Config) Bundler {
	b := bundlerImpl{
		Config:    &c,
		finishedC: make(chan struct{}),
		bundleC:   make(chan *protocol.ButlerLogBundle),
		streams:   map[string]bundlerStream{},
	}
	b.streamsCond = newTimeoutCond(b.getClock(), &b.streamsLock)

	go b.makeBundles()
	return &b
}

func (b *bundlerImpl) Register(p streamproto.Properties) (Stream, error) {
	if s := b.streams[p.Name]; s != nil {
		return nil, fmt.Errorf("a Stream is already registered for %q", p.Name)
	}

	// Our Properties must validate.
	if err := p.Validate(); err != nil {
		return nil, err
	}

	// Construct a parser for this stream.
	c := streamConfig{
		name: p.Name,
		template: protocol.ButlerLogBundle_Entry{
			Desc: &p.LogStreamDescriptor,
		},
		maximumBufferDuration: b.MaxBufferDelay,
		maximumBufferedBytes:  b.MaxBufferedBytes,
		onAppend: func(appended bool) {
			if appended {
				b.signalStreamUpdate()
			}
		},
	}

	err := error(nil)
	c.parser, err = newParser(&p, &b.prefixCounter)
	if err != nil {
		return nil, fmt.Errorf("failed to create stream parser: %s", err)
	}

	// Generate a secret for this Stream instance.
	c.template.Secret, err = types.NewStreamSecret()
	if err != nil {
		return nil, fmt.Errorf("failed to generate stream secret: %s", err)
	}

	// Create a new stream. This will kick off its processing goroutine, which
	// will not stop until it is closed.
	s := newStream(c)

	// Register the stream.
	b.streamsLock.Lock()
	defer b.streamsLock.Unlock()

	b.registerStreamLocked(s)
	return s, nil
}

func (b *bundlerImpl) CloseAndFlush() {
	// Mark that we're flushing. This will cause us to perform more aggressive
	// bundling in Next().
	b.startFlushing()
	<-b.finishedC
}

func (b *bundlerImpl) Next() *protocol.ButlerLogBundle {
	return <-b.bundleC
}

func (b *bundlerImpl) startFlushing() {
	b.streamsLock.Lock()
	defer b.streamsLock.Unlock()

	if !b.flushing {
		b.flushing = true
		b.signalStreamUpdateLocked()
	}
}

// makeBundles is run in its own goroutine. It runs continuously, responding
// to Stream constraints and availability and sending ButlerLogBundles through
// bundleC when available.
//
// makeBundles will terminate when closeC is closed and all streams are drained.
func (b *bundlerImpl) makeBundles() {
	defer close(b.finishedC)
	defer close(b.bundleC)

	b.streamsLock.Lock()
	defer b.streamsLock.Unlock()

	bb := (*builder)(nil)
	defer func() {
		if bb != nil && bb.hasContent() {
			b.bundleC <- bb.bundle()
		}
	}()

	for {
		bb = &builder{
			size: b.MaxBundleSize,
			template: protocol.ButlerLogBundle{
				Source:    b.Source,
				Timestamp: google.NewTimestamp(b.getClock().Now()),
			},
		}
		oldestContentTime := time.Time{}

		for {
			state := b.getStreamStateLocked()

			// Attempt to create more bundles.
			sendNow := b.bundleRoundLocked(bb, state)

			// Prune any drained streams.
			state.forEachStream(func(s bundlerStream) bool {
				if s.isDrained() {
					state.removeStream(s.name())
					b.unregisterStreamLocked(s)
				}

				return true
			})

			if b.flushing && len(state.streams) == 0 {
				// We're flushing, and there are no more streams, so we're completely
				// finished.
				//
				// If we have any content in our builder, it will be exported via defer.
				return
			}

			// If we have content, consider emitting this bundle.
			if bb.hasContent() {
				if b.MaxBufferDelay == 0 || sendNow || bb.ready() {
					break
				}
			}

			// Mark the first time this round where we actually saw data.
			if oldestContentTime.IsZero() && bb.hasContent() {
				oldestContentTime = state.now
			}

			// We will yield our stream lock and sleep, waiting for either:
			// 1) The earliest expiration time.
			// 2) A streams channel signal.
			//
			// We use a Cond here because we want Streams to be able to be added
			// while we're waiting for stream data.
			nextExpire, has := state.nextExpire()

			// If we have an oldest content time, that also means that we have
			// content. Factor this constraint in.
			if !oldestContentTime.IsZero() {
				roundExpire := oldestContentTime.Add(b.MaxBufferDelay)
				if !roundExpire.After(state.now) {
					break
				}

				if !has || roundExpire.Before(nextExpire) {
					nextExpire = roundExpire
					has = true
				}
			}

			if has {
				b.streamsCond.waitTimeout(nextExpire.Sub(state.now))
			} else {
				// No data, no expire constraints. Wait indefinitely for something
				// to change.
				b.streamsCond.Wait()
			}
		}

		// If our bundler has contents, send them.
		if bb.hasContent() {
			b.bundleC <- bb.bundle()
			bb = nil
		}
	}
}

// Implements a single bundle building round. This incrementally adds data from
// the stream state to the supplied builder.
//
// This method will block until a suitable bundle is available. Availability
// is subject both to time and data constraints:
//   - If buffered data, which is timestampped at ingest, has exceeded its
//     buffer duration threshold, a Bundle will be cut immediately.
//   - If no data is set to expire, the Bundler may wait for more data to
//     produce a more optimally-packed bundle.
//
// At a high level, Next operates as follows:
//   1) Freeze all stream state.
//   2) Scan streams for data that has exceeded its threshold; if data is found:
//     - Aggressively pack expired data into a Bundle.
//     - Optimally pack the remainder of the Bundle with any available data.
//     - Return the Bundle.
//
//   3) Examine the remaining data sizes, waiting for either:
//     - Enough stream data to fill our Bundle.
//     - Our timeout, if the Bundler is not closed.
//   4) Pack a Bundle with the remaining data optimally, emphasizing streams
//      with older data.
//
// Returns true if bundle some data was added that should be sent immediately.
func (b *bundlerImpl) bundleRoundLocked(bb *builder, state *streamState) bool {
	sendNow := false

	// First pass: non-blocking data that has exceeded its storage threshold.
	for bb.remaining() > 0 {
		s := state.next()
		if s == nil {
			break
		}

		if et, has := s.expireTime(); !has || et.After(state.now) {
			// This stream (and all other streams, since we're sorted) expires in
			// the future, so we're done with the first pass.
			break
		}

		// Pull bundles from this stream.
		if modified := s.nextBundleEntry(bb, true); modified {
			state.streamUpdated(s.name())

			// We have at least one time-sensitive bundle, so send this round.
			sendNow = true
		}
	}

	// Second pass: bundle any available data.
	state.forEachStream(func(s bundlerStream) bool {
		if bb.remaining() == 0 {
			return false
		}

		if modified := s.nextBundleEntry(bb, b.flushing); modified {
			state.streamUpdated(s.name())
		}
		return true
	})

	return sendNow
}

func (b *bundlerImpl) getStreamStateLocked() *streamState {
	// Lock and collect each stream.
	state := &streamState{
		streams: make([]bundlerStream, 0, len(b.streams)),
		now:     b.getClock().Now(),
	}

	for _, s := range b.streams {
		state.streams = append(state.streams, s)
	}
	heap.Init(state)

	return state
}

func (b *bundlerImpl) registerStreamLocked(s bundlerStream) {
	b.streams[s.name()] = s
	b.signalStreamUpdateLocked()
}

func (b *bundlerImpl) unregisterStreamLocked(s bundlerStream) {
	delete(b.streams, s.name())
}

func (b *bundlerImpl) signalStreamUpdate() {
	b.streamsLock.Lock()
	defer b.streamsLock.Unlock()

	b.signalStreamUpdateLocked()
}

func (b *bundlerImpl) signalStreamUpdateLocked() {
	b.streamsCond.Broadcast()
}

// nextPrefixIndex is a goroutine-safe method that returns the next prefix index
// for the given Bundler.
func (b *bundlerImpl) nextPrefixIndex() uint64 {
	return uint64(b.prefixCounter.next())
}

func (b *bundlerImpl) getClock() clock.Clock {
	c := b.Clock
	if c != nil {
		return c
	}
	return clock.GetSystemClock()
}

// streamState is a snapshot of the current stream registration. All operations
// performed on the state require streamLock to be held.
//
// streamState implements heap.Interface for its streams array. Streams without
// data times (nil) are considered to be greater than those with times.
type streamState struct {
	streams []bundlerStream
	now     time.Time
}

var _ heap.Interface = (*streamState)(nil)

func (s *streamState) next() bundlerStream {
	if len(s.streams) == 0 {
		return nil
	}
	return s.streams[0]
}

func (s *streamState) nextExpire() (time.Time, bool) {
	if next := s.next(); next != nil {
		if ts, ok := next.expireTime(); ok {
			return ts, true
		}
	}
	return time.Time{}, false
}

func (s *streamState) streamUpdated(name string) {
	if si, idx := s.streamIndex(name); si != nil {
		heap.Fix(s, idx)
	}
}

func (s *streamState) forEachStream(f func(bundlerStream) bool) {
	// Clone our streams, since the callback may mutate their order.
	streams := make([]bundlerStream, len(s.streams))
	for i, s := range s.streams {
		streams[i] = s
	}

	for _, s := range streams {
		if !f(s) {
			break
		}
	}
}

// removeStream removes a stream from the stream state.
func (s *streamState) removeStream(name string) bundlerStream {
	if si, idx := s.streamIndex(name); si != nil {
		heap.Remove(s, idx)
		return si
	}
	return nil
}

func (s *streamState) streamIndex(name string) (bundlerStream, int) {
	for i, si := range s.streams {
		if si.name() == name {
			return si, i
		}
	}
	return nil, -1
}

func (s *streamState) Len() int {
	return len(s.streams)
}

func (s *streamState) Less(i, j int) bool {
	si, sj := s.streams[i], s.streams[j]

	if it, ok := si.expireTime(); ok {
		if jt, ok := sj.expireTime(); ok {
			return it.Before(jt)
		}

		// i has data, but j does not, so i < j.
		return true
	}

	// i has no data, so i us greater than all other streams.
	return false
}

func (s *streamState) Swap(i, j int) {
	s.streams[i], s.streams[j] = s.streams[j], s.streams[i]
}

func (s *streamState) Push(x interface{}) {
	s.streams = append(s.streams, x.(bundlerStream))
}

func (s *streamState) Pop() interface{} {
	last := s.streams[len(s.streams)-1]
	s.streams = s.streams[:len(s.streams)-1]
	return last
}
