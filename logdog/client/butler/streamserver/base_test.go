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

package streamserver

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"testing"
	"time"

	"github.com/TriggerMail/luci-go/common/clock/clockflag"
	"github.com/TriggerMail/luci-go/common/clock/testclock"
	"github.com/TriggerMail/luci-go/logdog/client/butlerlib/streamclient"
	"github.com/TriggerMail/luci-go/logdog/client/butlerlib/streamproto"

	. "github.com/smartystreets/goconvey/convey"
)

type testAddr string

func (a testAddr) Network() string { return string(a) }
func (a testAddr) String() string  { return fmt.Sprintf("test(%s)", a.Network()) }

type testListener struct {
	err   error
	connC chan *testListenerConn
}

func newTestListener() *testListener {
	return &testListener{
		connC: make(chan *testListenerConn),
	}
}

func (l *testListener) Accept() (net.Conn, error) {
	if l.err != nil {
		return nil, l.err
	}
	conn, ok := <-l.connC
	if !ok {
		// Listener has been closed.
		return nil, errors.New("listener closed")
	}
	return conn, nil
}

func (l *testListener) Close() error {
	close(l.connC)
	return nil
}

func (l *testListener) Addr() net.Addr {
	return testAddr("test-listener")
}

func (l *testListener) connect(c *testListenerConn) {
	l.connC <- c
}

type testListenerConn struct {
	bytes.Buffer

	panicOnRead   bool
	readDeadline  time.Time
	writeDeadline time.Time
}

func (c *testListenerConn) Read(d []byte) (int, error) {
	if c.panicOnRead {
		panic("panic on read")
	}
	return c.Buffer.Read(d)
}

func (c *testListenerConn) Close() error         { return nil }
func (c *testListenerConn) LocalAddr() net.Addr  { return testAddr("local") }
func (c *testListenerConn) RemoteAddr() net.Addr { return testAddr("remote") }

func (c *testListenerConn) SetReadDeadline(t time.Time) error {
	c.readDeadline = t
	return nil
}

func (c *testListenerConn) SetWriteDeadline(t time.Time) error {
	c.writeDeadline = t
	return nil
}

func (c *testListenerConn) SetDeadline(t time.Time) error {
	c.readDeadline = t
	c.writeDeadline = t
	return nil
}

func TestListenerStreamServer(t *testing.T) {
	t.Parallel()

	Convey(`A stream server using a testing Listener`, t, func() {
		hb := handshakeBuilder{
			magic: streamproto.ProtocolFrameHeaderMagic,
		}

		var tl *testListener
		s := &listenerStreamServer{
			Context: context.Background(),
			gen: func() (net.Listener, string, error) {
				if tl != nil {
					panic("gen called more than once")
				}
				tl = newTestListener()
				return tl, "test", nil
			},
		}

		Convey(`Will panic if closed without listening.`, func() {
			So(func() { s.Close() }, ShouldPanic)
		})

		Convey(`Will fail to Listen if the Listener could not be created.`, func() {
			s.gen = func() (net.Listener, string, error) {
				return nil, "", errors.New("test error")
			}
			So(s.Listen(), ShouldNotBeNil)
		})

		Convey(`Can Listen for connections.`, func() {
			shouldClose := true
			So(s.Listen(), ShouldBeNil)
			defer func() {
				if shouldClose {
					s.Close()
				}
			}()

			tc := &testListenerConn{}

			Convey(`Can close, and will panic if double-closed.`, func() {
				s.Close()
				shouldClose = false

				So(func() { s.Close() }, ShouldPanic)
			})

			Convey(`Client with an invalid handshake magic number is rejected.`, func() {
				s.discardC = make(chan *streamClient)
				hb.magic = []byte(`NOT A HANDSHAKE MAGIC`)
				hb.writeTo(tc, "", nil)

				tl.connect(tc)
				So(<-s.discardC, ShouldNotBeNil)
			})

			Convey(`Client with invalid handshake JSON is rejected.`, func() {
				s.discardC = make(chan *streamClient)
				hb.writeTo(tc, "CLEARLY NOT JSON", nil)

				tl.connect(tc)
				So(<-s.discardC, ShouldNotBeNil)
			})

			Convey(`Client handshake panics are contained and rejected.`, func() {
				s.discardC = make(chan *streamClient)

				tc.panicOnRead = true
				hb.writeTo(tc, "", nil)

				tl.connect(tc)
				So(<-s.discardC, ShouldNotBeNil)
			})

			Convey(`Can receive stream data.`, func() {
				// Write our handshake and data to the stream.
				handshake := `{"name": "test", "contentType": "application/octet-stream"}`
				content := bytes.Repeat([]byte("THIS IS A TEST STREAM "), 100)
				hb.writeTo(tc, handshake, content)

				// Retrieve the ensuing stream.
				tl.connect(tc)
				stream, props := s.Next()
				So(stream, ShouldNotBeNil)
				defer stream.Close()
				So(props, ShouldNotBeNil)

				// Consume all of the data in the stream.
				recvData, _ := ioutil.ReadAll(stream)
				So(recvData, ShouldResemble, content)
			})

			Convey(`Will exit Next if closed.`, func() {
				streamC := make(chan *streamParams)
				defer close(streamC)

				// Get the stream.
				go func() {
					rc, props := s.Next()
					streamC <- &streamParams{rc, props}
				}()

				// Begin a client connection, but no handshake.
				tl.connect(tc)

				// Close the stream server.
				s.Close()
				shouldClose = false

				// Next must exit with nil.
				bundle := <-streamC
				So(bundle.rc, ShouldBeNil)
				So(bundle.properties, ShouldBeNil)
			})

			Convey(`Will refrain from outputting clients whose handshakes finish after the server is closed.`, func() {
				s.discardC = make(chan *streamClient, 1)

				handshake := `{"name": "test", "contentType": "application/octet-stream"}`
				content := bytes.Repeat([]byte("THIS IS A TEST STREAM "), 100)
				hb.writeTo(tc, handshake, content)

				tl.connect(tc)
				s.Close()
				shouldClose = false

				So(<-s.discardC, ShouldNotBeNil)
			})

		})
	})
}

func TestTCPServer(t *testing.T) {
	t.Parallel()

	Convey(`A TCP/IP4 client`, t, func() {
		ctx := context.Background()

		Convey(`When created and listening.`, func() {
			svr, err := NewTCP4Server(ctx, "")
			So(err, ShouldBeNil)

			So(svr.Listen(), ShouldBeNil)
			defer svr.Close()

			client, err := streamclient.New(svr.Address())
			So(err, ShouldBeNil)

			testClientServer(t, svr, client)
		})
	})
}

// testClientServer tests to ensure that a client can create streams with a
// server.
//
// svr must be in listening state when this is called.
func testClientServer(t *testing.T, svr StreamServer, client streamclient.Client) {
	flags := streamproto.Flags{
		Name:      "foo/bar",
		Timestamp: clockflag.Time(testclock.TestTimeLocal),
	}
	data := []byte("ohaithere")

	clientDoneC := make(chan error)
	go func() {
		var err error
		defer func() {
			clientDoneC <- err
		}()

		var stream streamclient.Stream
		if stream, err = client.NewStream(flags); err != nil {
			return
		}
		defer func() {
			if closeErr := stream.Close(); closeErr != nil && err == nil {
				err = closeErr
			}
		}()

		if _, err = stream.Write(data); err != nil {
			return
		}
	}()

	rc, props := svr.Next()
	defer rc.Close()

	So(props, ShouldResemble, flags.Properties())

	var buf bytes.Buffer
	_, err := buf.ReadFrom(rc)
	So(err, ShouldBeNil)
	So(buf.Bytes(), ShouldResemble, data)

	So(<-clientDoneC, ShouldBeNil)
}
