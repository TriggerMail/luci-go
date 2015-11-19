// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package bundler

import (
	"io"

	"github.com/luci/luci-go/common/logdog/protocol"
	"github.com/luci/luci-go/common/recordio"
)

// datagramParser is a parser implementation for the LogDog datagram stream
// type.
type datagramParser struct {
	baseParser

	// maxSize is the maximum allowed datagram size. Datagrams larger than this
	// will result in a processing error.
	maxSize int64

	// seq is the current datagram sequence number.
	seq int64
	// remaining is the amount of data remaining in a datagram that has previously
	// been emitted partially.
	//
	// This will be zero if we're not continuing a partial datagram.
	remaining int64
	// size is the size of the current partial datagram.
	//
	// This value is only valid if we're continuing a partial datagram (i.e., if
	// remaining is non-zero).
	size int64
}

var _ parser = (*datagramParser)(nil)

func (s *datagramParser) nextEntry(c *constraints) (*protocol.LogEntry, error) {
	// Use the current Buffer timestamp.
	ts, has := s.firstChunkTime()
	if !has {
		// No chunks, so no data.
		return nil, nil
	}

	// If remaining is zero, we don't have a buffered size header.
	//
	// Note that zero-size datagrams will store zero here on load; however, such
	// datagrams will never fail to emit a LogEntry, so s.remaining will have been
	// reset to zero by the next call.
	if s.remaining == 0 {
		bv := s.View()

		// Read the next datagram size header.
		rio := recordio.NewReader(bv, s.maxSize)
		size, _, err := rio.ReadFrame()
		if err != nil {
			switch err {
			case io.EOF:
				// Not enough data for a size header.
				return nil, nil

			case recordio.ErrFrameTooLarge:
				return nil, recordio.ErrFrameTooLarge
			}
			// Other errors should not be possible, since all operations are against
			// in-memory buffers.
			memoryCorruption(err)
		}

		s.size = size
		s.remaining = size

		// Don't need to read the size header again.
		s.Consume(bv.Consumed())
	}

	// If we read this, will it be partial?
	emitCount := s.remaining
	partial := false
	if emitCount > int64(c.limit) {
		partial = true
		emitCount = int64(c.limit)
	}

	bv := s.ViewLimit(s.remaining)
	if r := bv.Remaining(); r < emitCount {
		// Not enough buffered data to complete the datagram in one round.
		partial = true
		emitCount = r
	}
	if s.remaining > 0 && emitCount == 0 {
		// The datagram has data, but we can't emit any of it. No point in issuing
		// a zero-size partial datagram.
		return nil, nil
	}

	// We're not willing to emit a partial datagram unless we're allowed to
	// split.
	if partial && !c.allowSplit {
		return nil, nil
	}

	dg := protocol.Datagram{
		Partial: partial,
		Size:    uint64(s.size),
	}
	if emitCount > 0 {
		dg.Data = make([]byte, emitCount)
		bv.Read(dg.Data)
		s.Consume(emitCount)
	}

	le := s.baseLogEntry(ts)
	le.Sequence = uint64(s.seq)
	le.Content = &protocol.LogEntry_Datagram{Datagram: &dg}

	if !partial {
		s.seq++
		s.remaining = 0
	} else {
		s.remaining -= emitCount
	}
	return le, nil
}
