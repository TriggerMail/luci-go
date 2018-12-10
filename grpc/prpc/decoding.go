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

package prpc

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	"google.golang.org/grpc/metadata"

	"github.com/TriggerMail/luci-go/common/clock"
)

// This file implements decoding of HTTP requests to RPC parameters.

const (
	// headerSuffixBinary is a suffix of an HTTP header that specifies that
	// the header value is encoded in std base64.
	// After decoding, a handler must process the header without the suffix.
	headerSuffixBinary = "-Bin"
	headerContentType  = "Content-Type"
)

// readMessage decodes a protobuf message from an HTTP request.
// Does not close the request body.
func readMessage(r *http.Request, msg proto.Message) *protocolError {
	format, err := FormatFromContentType(r.Header.Get(headerContentType))
	if err != nil {
		// Spec: http://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html#sec10.4.16
		return errorf(http.StatusUnsupportedMediaType, "Content-Type header: %s", err)
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errorf(http.StatusBadRequest, "could not read body: %s", err)
	}
	switch format {
	// Do not redefine "err" below.

	case FormatJSONPB:
		err = jsonpb.Unmarshal(bytes.NewBuffer(buf), msg)

	case FormatText:
		err = proto.UnmarshalText(string(buf), msg)

	case FormatBinary:
		err = proto.Unmarshal(buf, msg)

	default:
		panic(fmt.Errorf("impossible: invalid format %v", format))
	}
	if err != nil {
		return errorf(http.StatusBadRequest, "could not decode body: %s", err)
	}
	return nil
}

// parseHeader parses HTTP headers and derives a new context.
// Supports HeaderTimeout.
// Ignores "Accept" and "Content-Type" headers.
//
// If there are unrecognized HTTP headers, with or without headerSuffixBinary,
// they are added to a metadata.MD and a new context is derived.
// If c already has metadata, the latter is copied.
//
// In case of an error, returns c unmodified.
func parseHeader(c context.Context, header http.Header) (context.Context, error) {
	origC := c

	md, ok := metadata.FromIncomingContext(c)
	if ok {
		md = md.Copy()
	} else {
		md = metadata.MD{}
	}

	addedMeta := false
	for name, values := range header {
		if len(values) == 0 {
			continue
		}
		name = http.CanonicalHeaderKey(name)
		switch name {

		case HeaderTimeout:
			// Decode only first value, ignore the rest
			// to be consistent with http.Header.Get.
			timeout, err := DecodeTimeout(values[0])
			if err != nil {
				return origC, fmt.Errorf("%s header: %s", HeaderTimeout, err)
			}
			c, _ = clock.WithTimeout(c, timeout)

		case headerAccept, headerContentType:
		// readMessage and writeMessage handle these headers.

		default:
			addedMeta = true
			if !strings.HasSuffix(name, headerSuffixBinary) {
				md[name] = append(md[name], values...)
				break // switch name
			}
			trimmedName := strings.TrimSuffix(name, headerSuffixBinary)
			for _, v := range values {
				decoded, err := base64.StdEncoding.DecodeString(v)
				if err != nil {
					return origC, fmt.Errorf("%s header: %s", name, err)
				}
				md[trimmedName] = append(md[trimmedName], string(decoded))
			}
		}
	}
	if addedMeta {
		c = metadata.NewIncomingContext(c, md)
	}
	return c, nil
}
