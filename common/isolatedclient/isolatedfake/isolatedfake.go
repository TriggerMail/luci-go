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

// Package isolatedfake implements an in-process fake Isolated server for
// integration testing.
package isolatedfake

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"

	isolateservice "github.com/TriggerMail/luci-go/common/api/isolate/isolateservice/v1"
	"github.com/TriggerMail/luci-go/common/isolated"
)

const contentType = "application/json; charset=utf-8"

type jsonAPI func(r *http.Request) interface{}

type failure interface {
	Fail(err error)
}

// handlerJSON converts a jsonAPI http handler to a proper http.Handler.
func handlerJSON(f failure, handler jsonAPI) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Printf("%s", r.URL)
		if r.Header.Get("Content-Type") != contentType {
			f.Fail(fmt.Errorf("invalid content type: %s", r.Header.Get("Content-Type")))
			return
		}
		defer r.Body.Close()
		out := handler(r)
		w.Header().Set("Content-Type", contentType)
		j := json.NewEncoder(w)
		if err := j.Encode(out); err != nil {
			f.Fail(err)
		}
	})
}

func gsURLWithDigestQuery(host, path, digest string) *url.URL {
	v := url.Values{}
	v.Add("digest", digest)
	fullPath := fmt.Sprintf("/fake/cloudstorage/%s", path)
	return &url.URL{Scheme: "http", Host: host, Path: fullPath, RawQuery: v.Encode()}
}

// IsolatedFake is a functional fake in-memory isolated server.
type IsolatedFake interface {
	http.Handler
	// Contents returns all the uncompressed data on the fake isolated server.
	Contents() map[isolated.HexDigest][]byte
	// Inject adds uncompressed data in the fake isolated server.
	Inject(data []byte) isolated.HexDigest
	Error() error
}

type isolatedFake struct {
	mux      *http.ServeMux
	lock     sync.Mutex
	err      error
	contents map[isolated.HexDigest][]byte
	staging  map[isolated.HexDigest][]byte // Uploaded to GCS but not yet finalized.
}

// New create a HTTP router that implements an isolated server.
func New() IsolatedFake {
	server := &isolatedFake{
		mux:      http.NewServeMux(),
		contents: map[isolated.HexDigest][]byte{},
		staging:  map[isolated.HexDigest][]byte{},
	}

	server.handleJSON("/_ah/api/isolateservice/v1/server_details", server.serverDetails)
	server.handleJSON("/_ah/api/isolateservice/v1/preupload", server.preupload)
	server.handleJSON("/_ah/api/isolateservice/v1/finalize_gs_upload", server.finalizeGSUpload)
	server.handleJSON("/_ah/api/isolateservice/v1/store_inline", server.storeInline)
	server.handleJSON("/_ah/api/isolateservice/v1/retrieve", server.retrieve)
	server.mux.HandleFunc("/fake/cloudstorage/upload", server.fakeCloudStorageUpload)
	server.mux.HandleFunc("/fake/cloudstorage/download", server.fakeCloudStorageDownload)

	// Fail on anything else.
	server.mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		server.Fail(fmt.Errorf("unknown endpoint %s", req.URL))
	})
	return server
}

// Private details.

func (server *isolatedFake) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.mux.ServeHTTP(w, r)
}

func (server *isolatedFake) Contents() map[isolated.HexDigest][]byte {
	server.lock.Lock()
	defer server.lock.Unlock()
	out := map[isolated.HexDigest][]byte{}
	for k, v := range server.contents {
		out[k] = v
	}
	return out
}

func (server *isolatedFake) Inject(data []byte) isolated.HexDigest {
	h := isolated.HashBytes(data)
	server.lock.Lock()
	defer server.lock.Unlock()
	server.contents[h] = data
	return h
}

func (server *isolatedFake) Fail(err error) {
	server.lock.Lock()
	defer server.lock.Unlock()
	server.failLocked(err)
}

func (server *isolatedFake) Error() error {
	server.lock.Lock()
	defer server.lock.Unlock()
	return server.err
}

func (server *isolatedFake) failLocked(err error) {
	if server.err == nil {
		server.err = err
	}
}

func (server *isolatedFake) handleJSON(path string, handler jsonAPI) {
	server.mux.Handle(path, handlerJSON(server, handler))
}

func (server *isolatedFake) serverDetails(r *http.Request) interface{} {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		server.Fail(err)
	}
	if string(content) != "{}" {
		server.Fail(fmt.Errorf("unexpected content %#v", string(content)))
	}
	return map[string]string{"server_version": "v1"}
}

func (server *isolatedFake) preupload(r *http.Request) interface{} {
	data := &isolateservice.HandlersEndpointsV1DigestCollection{}
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		server.Fail(err)
	}
	if data.Namespace == nil || data.Namespace.Namespace != "default-gzip" {
		server.Fail(fmt.Errorf("unexpected namespace %#v", data.Namespace.Namespace))
	}
	out := &isolateservice.HandlersEndpointsV1UrlCollection{}

	server.lock.Lock()
	defer server.lock.Unlock()
	for i, d := range data.Items {
		if _, ok := server.contents[isolated.HexDigest(d.Digest)]; !ok {
			// Simulate a write to Cloud Storage for larger writes.
			ticket := "ticket:" + string(d.Digest)
			s := &isolateservice.HandlersEndpointsV1PreuploadStatus{
				Index:        int64(i),
				UploadTicket: ticket,
			}
			if d.Size > 1024 {
				s.GsUploadUrl = gsURLWithDigestQuery(r.Host, "upload", d.Digest).String()
			}
			out.Items = append(out.Items, s)
		}
	}
	return out
}

func (server *isolatedFake) fakeCloudStorageUpload(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Header.Get("Content-Type") != "application/octet-stream" {
		w.WriteHeader(400)
		server.Fail(fmt.Errorf("invalid content type: %s", r.Header.Get("Content-Type")))
		return
	}
	if r.Method != "PUT" {
		w.WriteHeader(405)
		server.Fail(fmt.Errorf("invalid method: %s", r.Method))
		return
	}
	decompressor, err := isolated.GetDecompressor(r.Body)
	if err != nil {
		w.WriteHeader(500)
		server.Fail(err)
		return
	}
	defer decompressor.Close()
	raw, err := ioutil.ReadAll(decompressor)
	if err != nil {
		w.WriteHeader(500)
		server.Fail(err)
		return
	}
	digest := isolated.HexDigest(r.URL.Query().Get("digest"))
	if digest != isolated.HashBytes(raw) {
		w.WriteHeader(400)
		server.Fail(fmt.Errorf("invalid digest %#v", digest))
		return
	}

	server.lock.Lock()
	defer server.lock.Unlock()
	server.staging[digest] = raw
	w.WriteHeader(200)
}

func (server *isolatedFake) fakeCloudStorageDownload(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != "GET" {
		w.WriteHeader(405)
		server.Fail(fmt.Errorf("invalid method: %s", r.Method))
		return
	}
	digest := isolated.HexDigest(r.URL.Query().Get("digest"))
	data, ok := server.contents[digest]
	if !ok {
		w.WriteHeader(404)
		server.Fail(fmt.Errorf("file not found: %s", digest))
		return
	}
	var buf bytes.Buffer
	compressor, err := isolated.GetCompressor(&buf)
	if err != nil {
		w.WriteHeader(500)
		server.Fail(err)
		return
	}
	if _, err := io.CopyBuffer(compressor, bytes.NewReader(data), nil); err != nil {
		compressor.Close()
		w.WriteHeader(500)
		server.Fail(err)
		return
	}
	if err := compressor.Close(); err != nil {
		w.WriteHeader(500)
		server.Fail(err)
		return
	}
	w.Write(buf.Bytes())
}

func (server *isolatedFake) finalizeGSUpload(r *http.Request) interface{} {
	data := &isolateservice.HandlersEndpointsV1FinalizeRequest{}
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		server.Fail(err)
		return map[string]string{"err": err.Error()}
	}
	prefix := "ticket:"
	if !strings.HasPrefix(data.UploadTicket, prefix) {
		err := fmt.Errorf("unexpected ticket %#v", data.UploadTicket)
		server.Fail(err)
		return map[string]string{"err": err.Error()}
	}
	digest := isolated.HexDigest(data.UploadTicket[len(prefix):])
	if !digest.Validate() {
		err := fmt.Errorf("invalid digest %#v", digest)
		server.Fail(err)
		return map[string]string{"err": err.Error()}
	}

	server.lock.Lock()
	defer server.lock.Unlock()
	if _, ok := server.staging[digest]; !ok {
		err := fmt.Errorf("finalizing non uploaded file")
		server.failLocked(err)
		return map[string]string{"err": err.Error()}
	}
	server.contents[digest] = server.staging[digest]
	delete(server.staging, digest)
	return map[string]string{"ok": "true"}
}

func (server *isolatedFake) storeInline(r *http.Request) interface{} {
	data := &isolateservice.HandlersEndpointsV1StorageRequest{}
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		server.Fail(err)
		return map[string]string{"err": err.Error()}
	}

	prefix := "ticket:"
	if !strings.HasPrefix(data.UploadTicket, prefix) {
		err := fmt.Errorf("unexpected ticket %#v", data.UploadTicket)
		server.Fail(err)
		return map[string]string{"err": err.Error()}
	}

	digest := isolated.HexDigest(data.UploadTicket[len(prefix):])
	if !digest.Validate() {
		err := fmt.Errorf("invalid digest %#v", digest)
		server.Fail(err)
		return map[string]string{"err": err.Error()}
	}
	blob, err := base64.StdEncoding.DecodeString(data.Content)
	if err != nil {
		server.Fail(err)
		return map[string]string{"err": err.Error()}
	}
	decompressor, err := isolated.GetDecompressor(bytes.NewReader(blob))
	if err != nil {
		server.Fail(err)
		return map[string]string{"err": err.Error()}
	}
	defer decompressor.Close()
	raw, err := ioutil.ReadAll(decompressor)
	if err != nil {
		server.Fail(err)
		return map[string]string{"err": err.Error()}
	}
	if digest != isolated.HashBytes(raw) {
		err := fmt.Errorf("invalid digest %#v", digest)
		server.Fail(err)
		return map[string]string{"err": err.Error()}
	}

	server.lock.Lock()
	defer server.lock.Unlock()
	server.contents[digest] = raw
	//log.Printf("  storing %s = %d bytes", digest, len(raw))
	return map[string]string{"ok": "true"}
}

func (server *isolatedFake) retrieve(r *http.Request) interface{} {
	data := &isolateservice.HandlersEndpointsV1RetrieveRequest{}
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		server.Fail(err)
		return map[string]string{"err": err.Error()}
	}
	digest := isolated.HexDigest(data.Digest)
	rawContent, ok := server.contents[digest]
	if !ok {
		err := fmt.Errorf("no such digest %#v", digest)
		server.Fail(err)
		return map[string]string{"err": err.Error()}
	}
	if len(rawContent) > 1024 {
		return &isolateservice.HandlersEndpointsV1RetrievedContent{
			Url: gsURLWithDigestQuery(r.Host, "download", data.Digest).String(),
		}
	}

	// Since we decompress when we get the data, we need to recompress when
	// something is fetched.
	var buf bytes.Buffer
	compressor, err := isolated.GetCompressor(&buf)
	if err != nil {
		server.Fail(err)
		return map[string]string{"err": err.Error()}
	}
	if _, err := io.CopyBuffer(compressor, bytes.NewReader(rawContent), nil); err != nil {
		compressor.Close()
		server.Fail(err)
		return map[string]string{"err": err.Error()}
	}
	if err := compressor.Close(); err != nil {
		server.Fail(err)
		return map[string]string{"err": err.Error()}
	}
	return &isolateservice.HandlersEndpointsV1RetrievedContent{
		Content: base64.StdEncoding.EncodeToString(buf.Bytes()),
	}
}
