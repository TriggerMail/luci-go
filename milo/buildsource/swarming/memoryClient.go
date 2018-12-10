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

package swarming

import (
	"bytes"
	"fmt"

	"github.com/golang/protobuf/proto"
	miloProto "github.com/TriggerMail/luci-go/common/proto/milo"
	"github.com/TriggerMail/luci-go/logdog/client/butlerlib/streamclient"
	"github.com/TriggerMail/luci-go/logdog/client/butlerlib/streamproto"
	"github.com/TriggerMail/luci-go/milo/buildsource/rawpresentation"
)

// In-memory datastructure to hold a fake butler client.
type memoryStream struct {
	props *streamproto.Properties

	closed     bool
	buf        bytes.Buffer
	isDatagram bool
}

func (s *memoryStream) ToLogDogStream() (*rawpresentation.Stream, error) {
	result := &rawpresentation.Stream{
		Closed:     s.closed,
		IsDatagram: s.isDatagram,
		Path:       s.props.Name,
		Prefix:     s.props.Prefix,
	}

	if s.isDatagram {
		result.Data = &miloProto.Step{}
		// Assume this is a miloProto.Step.
		if err := proto.Unmarshal(s.buf.Bytes(), result.Data); err != nil {
			return nil, err
		}
	} else {
		result.Text = s.buf.String()
	}

	return result, nil
}

func (s *memoryStream) Write(b []byte) (int, error) {
	return s.buf.Write(b)
}

func (s *memoryStream) Close() error {
	s.closed = true
	return nil
}

func (s *memoryStream) WriteDatagram(b []byte) error {
	s.isDatagram = true

	s.buf.Reset()
	_, err := s.buf.Write(b)
	return err
}

func (s *memoryStream) Properties() *streamproto.Properties { return s.props.Clone() }

type memoryClient struct {
	stream map[string]*memoryStream
}

func (c *memoryClient) NewStream(f streamproto.Flags) (streamclient.Stream, error) {
	props := f.Properties()
	if _, ok := c.stream[props.Name]; ok {
		return nil, fmt.Errorf("duplicate stream, %q", props.Name)
	}
	s := memoryStream{
		props: props,
	}
	if c.stream == nil {
		c.stream = map[string]*memoryStream{}
	}
	c.stream[s.props.Name] = &s
	return &s, nil
}

func (c *memoryClient) addLogDogTextStream(s *rawpresentation.Streams, ls *miloProto.LogdogStream) error {
	var keys []string
	for k := range c.stream {
		keys = append(keys, k)
	}
	ms, ok := c.stream[ls.Name]
	if !ok {
		return fmt.Errorf("Could not find text stream %q\n%s", ls.Name, keys)
	}
	lds, err := ms.ToLogDogStream()
	if err != nil {
		return fmt.Errorf("Could not convert text stream %s\n%s\n%s", ls.Name, err, keys)
	}
	if lds.IsDatagram {
		return fmt.Errorf("Expected %s to be a text stream, got a data stream", ls.Name)
	}
	s.Streams[ls.Name] = lds
	return nil
}

// addToStreams adds the set of stream with a given base path to the logdog
// stream map.  A base path is assumed to have a stream named "annotations".
func (c *memoryClient) addToStreams(s *rawpresentation.Streams, anno *miloProto.Step) error {
	if lds := anno.StdoutStream; lds != nil {
		if err := c.addLogDogTextStream(s, lds); err != nil {
			return fmt.Errorf(
				"Encountered error while processing step streams for STDOUT: %s", err)
		}
	}
	if lds := anno.StderrStream; lds != nil {
		if err := c.addLogDogTextStream(s, lds); err != nil {
			return fmt.Errorf(
				"Encountered error while processing step streams for STDERR: %s", err)
		}
	}

	// Look for substream.
	for _, link := range anno.GetOtherLinks() {
		lds := link.GetLogdogStream()
		// Not a logdog stream.
		if lds == nil {
			continue
		}

		if err := c.addLogDogTextStream(s, lds); err != nil {
			return fmt.Errorf(
				"Encountered error while processing step streams for '%s'\n%s", lds.Name, err)
		}
	}

	// Now do substeps.
	for _, subStepEntry := range anno.Substep {
		substep := subStepEntry.GetStep()
		if substep == nil {
			continue
		}

		if err := c.addToStreams(s, substep); err != nil {
			return err
		}
	}

	return nil
}

func (c *memoryClient) ToLogDogStreams() (*rawpresentation.Streams, error) {
	result := &rawpresentation.Streams{}
	result.Streams = map[string]*rawpresentation.Stream{}

	// Register annotation stream.
	const annotationStreamName = "annotations"
	ms, ok := c.stream[annotationStreamName]
	if !ok {
		return nil, fmt.Errorf("Could not find stream %s", annotationStreamName)
	}
	ls, err := ms.ToLogDogStream()
	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal stream %s\n%s", annotationStreamName, err)
	}
	if !ls.IsDatagram {
		return nil, fmt.Errorf(
			"Annotation stream %s is not a datagram\nText: %s",
			annotationStreamName, ls.Text)
	}
	result.Streams[annotationStreamName] = ls

	// Register any referenced LogDog streams.
	if err := c.addToStreams(result, ls.Data); err != nil {
		return nil, err
	}
	result.MainStream = ls

	if len(c.stream) != len(result.Streams) {
		var mk, lk []string
		for k := range c.stream {
			mk = append(mk, k)
		}
		for k := range result.Streams {
			lk = append(lk, k)
		}
		return nil, fmt.Errorf(
			"Number of streams do not match %d vs %d\nMemory:%s\nResult:%s",
			len(c.stream), len(result.Streams), mk, lk)
	}

	return result, nil
}
