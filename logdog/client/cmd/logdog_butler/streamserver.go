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
	"flag"
	"fmt"
	"strings"

	"github.com/TriggerMail/luci-go/common/errors"
	"github.com/TriggerMail/luci-go/logdog/client/butler/streamserver"
)

type streamServerURI string

var commonStreamServerExamples = []string{
	"tcp4:[addr_v4][:port]",
	"tcp6:[addr_v6][:port]",
}

func exampleStreamServerURIs() string {
	examples := make([]string, 0, len(commonStreamServerExamples)+len(platformStreamServerExamples))
	for _, ex := range commonStreamServerExamples {
		examples = append(examples, fmt.Sprintf(`"%s"`, ex))
	}
	for _, ex := range platformStreamServerExamples {
		examples = append(examples, fmt.Sprintf(`"%s"`, ex))
	}

	return strings.Join(examples, ", ")
}

// (*streamServerURI) must implement flag.Value.
var _ = flag.Value((*streamServerURI)(nil))

// String implements flag.Value.
func (u *streamServerURI) String() string {
	return string(*u)
}

// Set implements flag.Value.
func (u *streamServerURI) Set(v string) error {
	uri := streamServerURI(v)
	if _, err := uri.resolve(context.Background()); err != nil {
		return err
	}
	*u = uri
	return nil
}

func (u streamServerURI) resolve(ctx context.Context) (streamserver.StreamServer, error) {
	// Split URI into typ[:spec]
	parts := strings.SplitN(string(u), ":", 2)
	typ, spec := parts[0], ""
	if len(parts) >= 2 {
		spec = parts[1]
	}

	// Platform-specific.
	//
	// This will return a non-nil error if the string itself is invalid.
	// Otherwise, it will return a StreamServer if the combination successfully
	// resolved to one.
	//
	// If an error is returned, it should be properly annotated.
	switch s, err := resolvePlatform(ctx, typ, spec); {
	case err != nil:
		return nil, err
	case s != nil:
		return s, nil
	}

	// Common implementations.
	switch typ {
	case "tcp4":
		return streamserver.NewTCP4Server(ctx, spec)

	case "tcp6":
		return streamserver.NewTCP6Server(ctx, spec)

	default:
		return nil, errors.Reason("unknown stream server type: %q", typ).Err()
	}
}
