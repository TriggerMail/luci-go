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

package isolatedclient

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	isolateservice "github.com/TriggerMail/luci-go/common/api/isolate/isolateservice/v1"
	"github.com/TriggerMail/luci-go/common/isolated"
	"github.com/TriggerMail/luci-go/common/lhttp"
	"github.com/TriggerMail/luci-go/common/retry"
	"github.com/TriggerMail/luci-go/common/runtime/tracer"
)

// DefaultNamespace is the namespace that should be used with the New function.
const DefaultNamespace = "default-gzip"

// Source is a generator method to return source data. A generated Source must
// be Closed before the generator is called again.
type Source func() (io.ReadCloser, error)

// NewBytesSource returns a Source implementation that reads from the supplied
// byte slice.
func NewBytesSource(d []byte) Source {
	return func() (io.ReadCloser, error) {
		return ioutil.NopCloser(bytes.NewReader(d)), nil
	}
}

// PushState is per-item state passed from IsolateServer.Contains() to
// IsolateServer.Push().
//
// Its content is implementation specific.
type PushState struct {
	status    isolateservice.HandlersEndpointsV1PreuploadStatus
	digest    isolated.HexDigest
	size      int64
	uploaded  bool
	finalized bool
}

// CloudStorage is the interface for clients to fetch from and push to GCS storage.
type CloudStorage interface {
	// Fetch is a handler for retrieving specified content from GCS and storing
	// the response in the provided destination buffer.
	Fetch(context.Context, *Client, isolateservice.HandlersEndpointsV1RetrievedContent, io.Writer) error
	// Push is a handler for pushing content from provided buffer to GCS.
	Push(context.Context, *Client, isolateservice.HandlersEndpointsV1PreuploadStatus, Source) error
}

// Client is a client to an isolated server.
type Client struct {
	// All the members are immutable.
	retryFactory retry.Factory
	url          string
	namespace    string

	authClient *http.Client // client that sends auth tokens
	anonClient *http.Client // client that does NOT send auth tokens
	gcsHandler CloudStorage // implements GCS fetch and push handlers
}

// New returns a new IsolateServer client.
//
// 'authClient' must implement authentication sufficient to talk to Isolate server
// (OAuth tokens with 'email' scope).
//
// 'anonClient' must be a functional http.Client.
//
// If either client is nil, it will use http.DefaultClient (which will not work
// on Classic AppEngine!).
//
// If you're unsure which namespace to use, use the DefaultNamespace constant.
//
// If gcs is nil, the defaultGCSHandler is used for fetching from and pushing to GCS.
func New(anonClient, authClient *http.Client, host, namespace string, rFn retry.Factory, gcs CloudStorage) *Client {
	if anonClient == nil {
		anonClient = http.DefaultClient
	}
	if authClient == nil {
		authClient = http.DefaultClient
	}
	if gcs == nil {
		gcs = defaultGCSHandler{}
	}
	i := &Client{
		retryFactory: rFn,
		url:          strings.TrimRight(host, "/"),
		namespace:    namespace,
		authClient:   authClient,
		anonClient:   anonClient,
		gcsHandler:   gcs,
	}
	tracer.NewPID(i, "isolatedclient:"+i.url)
	return i
}

// ServerCapabilities returns the server details.
func (i *Client) ServerCapabilities(c context.Context) (*isolateservice.HandlersEndpointsV1ServerDetails, error) {
	out := &isolateservice.HandlersEndpointsV1ServerDetails{}
	if err := i.postJSON(c, "/_ah/api/isolateservice/v1/server_details", nil, map[string]string{}, out); err != nil {
		return nil, err
	}
	return out, nil
}

// Contains looks up cache presence on the server of multiple items.
//
// The returned list is in the same order as 'items', with entries nil for
// items that were present.
func (i *Client) Contains(c context.Context, items []*isolateservice.HandlersEndpointsV1Digest) (out []*PushState, err error) {
	end := tracer.Span(i, "contains", tracer.Args{"number": len(items)})
	defer func() { end(tracer.Args{"err": err}) }()
	in := isolateservice.HandlersEndpointsV1DigestCollection{Items: items, Namespace: &isolateservice.HandlersEndpointsV1Namespace{}}
	in.Namespace.Namespace = i.namespace
	data := &isolateservice.HandlersEndpointsV1UrlCollection{}
	if err = i.postJSON(c, "/_ah/api/isolateservice/v1/preupload", nil, in, data); err != nil {
		return nil, err
	}
	out = make([]*PushState, len(items))
	for _, e := range data.Items {
		index := int(e.Index)
		out[index] = &PushState{
			status: *e,
			digest: isolated.HexDigest(items[index].Digest),
			size:   items[index].Size,
		}
	}
	return out, nil
}

// Push pushed a missing item, as reported by Contains(), to the server.
func (i *Client) Push(c context.Context, state *PushState, source Source) (err error) {
	// This push operation may be a retry after failed finalization call below,
	// no need to reupload contents in that case.
	if !state.uploaded {
		// PUT file to uploadURL.
		if err = i.doPush(c, state, source); err != nil {
			log.Printf("doPush(%s) failed: %s\n%#v", state.digest, err, state)
			return
		}
		state.uploaded = true
	}

	// Optionally notify the server that it's done.
	if state.status.GsUploadUrl != "" {
		end := tracer.Span(i, "finalize", nil)
		defer func() { end(tracer.Args{"err": err}) }()
		// TODO(vadimsh): Calculate MD5 or CRC32C sum while uploading a file and
		// send it to isolated server. That way isolate server can verify that
		// the data safely reached Google Storage (GS provides MD5 and CRC32C of
		// stored files).
		in := isolateservice.HandlersEndpointsV1FinalizeRequest{UploadTicket: state.status.UploadTicket}
		headers := map[string]string{"Cache-Control": "public, max-age=31536000"}
		if err = i.postJSON(c, "/_ah/api/isolateservice/v1/finalize_gs_upload", headers, in, nil); err != nil {
			log.Printf("Push(%s) (finalize) failed: %s\n%#v", state.digest, err, state)
			return
		}
	}
	state.finalized = true
	return
}

// Fetch downloads an item from the server.
func (i *Client) Fetch(c context.Context, digest isolated.HexDigest, dest io.Writer) error {
	// Perform initial request.
	url := i.url + "/_ah/api/isolateservice/v1/retrieve"
	in := &isolateservice.HandlersEndpointsV1RetrieveRequest{
		Digest: string(digest),
		Namespace: &isolateservice.HandlersEndpointsV1Namespace{
			DigestHash: "sha-1",
			Namespace:  i.namespace,
		},
		Offset: 0,
	}
	var out isolateservice.HandlersEndpointsV1RetrievedContent
	if _, err := lhttp.PostJSON(c, i.retryFactory, i.authClient, url, nil, in, &out); err != nil {
		return err
	}

	// Handle DB items.
	if out.Content != "" {
		decoded, err := base64.StdEncoding.DecodeString(out.Content)
		if err != nil {
			return err
		}
		decompressor, err := isolated.GetDecompressor(bytes.NewReader(decoded))
		if err != nil {
			return err
		}
		defer decompressor.Close()
		_, err = io.Copy(dest, decompressor)
		return err
	}

	// Handle GCS items.
	return i.gcsHandler.Fetch(c, i, out, dest)
}

// postJSON does authenticated POST request.
func (i *Client) postJSON(c context.Context, resource string, headers map[string]string, in, out interface{}) error {
	if len(resource) == 0 || resource[0] != '/' {
		return errors.New("resource must start with '/'")
	}
	_, err := lhttp.PostJSON(c, i.retryFactory, i.authClient, i.url+resource, headers, in, out)
	return err
}

func (i *Client) doPush(c context.Context, state *PushState, source Source) (err error) {
	useDB := state.status.GsUploadUrl == ""
	end := tracer.Span(i, "push", tracer.Args{"useDB": useDB, "size": state.size})
	defer func() { end(tracer.Args{"err": err}) }()
	if useDB {
		var src io.ReadCloser
		src, err = source()
		if err != nil {
			return err
		}
		defer src.Close()

		err = i.doPushDB(c, state, src)
	} else {
		err = i.gcsHandler.Push(c, i, state.status, source)
	}
	if err != nil {
		tracer.CounterAdd(i, "bytesUploaded", float64(state.size))
	}
	return err
}

func (i *Client) doPushDB(c context.Context, state *PushState, reader io.Reader) error {
	buf := bytes.Buffer{}
	compressor, err := isolated.GetCompressor(&buf)
	if err != nil {
		return err
	}
	if _, err := io.Copy(compressor, reader); err != nil {
		return err
	}
	if err := compressor.Close(); err != nil {
		return err
	}
	in := &isolateservice.HandlersEndpointsV1StorageRequest{
		UploadTicket: state.status.UploadTicket,
		Content:      base64.StdEncoding.EncodeToString(buf.Bytes()),
	}
	return i.postJSON(c, "/_ah/api/isolateservice/v1/store_inline", nil, in, nil)
}

// defaultGCSHandler implements the default Fetch and Push handlers for
// interacting with GCS.
type defaultGCSHandler struct{}

// Fetch uses the provided HandlersEndpointsV1RetrievedContent response to
// download content from GCS to the provided dest.
func (gcs defaultGCSHandler) Fetch(c context.Context, i *Client, content isolateservice.HandlersEndpointsV1RetrievedContent, dest io.Writer) error {
	rgen := func() (*http.Request, error) {
		return http.NewRequest("GET", content.Url, nil)
	}
	handler := func(resp *http.Response) error {
		defer resp.Body.Close()
		decompressor, err := isolated.GetDecompressor(resp.Body)
		if err != nil {
			return err
		}
		defer decompressor.Close()
		_, err = io.Copy(dest, decompressor)
		return err
	}
	_, err := lhttp.NewRequest(c, i.authClient, i.retryFactory, rgen, handler, nil)()
	return err
}

// Push uploads content from the provided source to the GCS path specified in
// the HandlersEndpointsV1PreuploadStatus response.
func (gcs defaultGCSHandler) Push(c context.Context, i *Client, status isolateservice.HandlersEndpointsV1PreuploadStatus, source Source) error {
	// GsUploadUrl is signed Google Storage URL that doesn't require additional
	// authentication. In fact, using authClient causes HTTP 403 because
	// authClient's tokens don't have Cloud Storage OAuth scope. Use anonymous
	// client instead.
	req := lhttp.NewRequest(c, i.anonClient, i.retryFactory, func() (*http.Request, error) {
		src, err := source()
		if err != nil {
			return nil, err
		}

		request, err := http.NewRequest("PUT", status.GsUploadUrl, nil)
		if err != nil {
			src.Close()
			return nil, err
		}
		request.Body = newCompressed(src)
		request.Header.Set("Content-Type", "application/octet-stream")
		return request, nil
	}, func(resp *http.Response) error {
		_, err4 := io.Copy(ioutil.Discard, resp.Body)
		err5 := resp.Body.Close()
		if err4 != nil {
			return err4
		}
		return err5
	}, nil)
	_, err := req()
	return err
}

// compressed is an io.ReadCloser that transparently compresses source data in
// a separate goroutine.
type compressed struct {
	pr  *io.PipeReader
	src io.ReadCloser
}

func (c *compressed) Read(data []byte) (int, error) {
	return c.pr.Read(data)
}

func (c *compressed) Close() error {
	err := c.pr.Close()
	if err1 := c.src.Close(); err == nil {
		err = err1
	}
	return err
}

// newCompressed creates a pipeline to compress a file into a ReadCloser via:
// src (file as ReadCloser) -> gzip compressor (via io.CopyBuffer) -> bufio.Writer
// -> io.Pipe Writer side -> io.Pipe Reader side
func newCompressed(src io.ReadCloser) io.ReadCloser {
	pr, pw := io.Pipe()
	go func() {
		// Memory is cheap, we never want this pipeline to stall.
		const outBufSize = 1024 * 1024
		// We write into a buffer so that the Reader can read larger chunks of
		// compressed data out of the compressor instead of being constrained to
		// compressed(4096 bytes) each read.
		bufWriter := bufio.NewWriterSize(pw, outBufSize)
		// The compressor itself is not thread safe.
		compressor, err := isolated.GetCompressor(bufWriter)
		if err != nil {
			pw.CloseWithError(err)
			return
		}
		// Make this 3x bigger than the output buffer since it is uncompressed.
		buf := make([]byte, outBufSize*3)
		if _, err := io.CopyBuffer(compressor, src, buf); err != nil {
			compressor.Close()
			pw.CloseWithError(err)
			return
		}
		// compressor needs to be closed first to flush the rest of the data
		// into the bufio.Writer
		if err := compressor.Close(); err != nil {
			pw.CloseWithError(err)
			return
		}
		pw.CloseWithError(bufWriter.Flush())
	}()

	return &compressed{pr, src}
}
