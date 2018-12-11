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

package discovery

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"sync"

	"github.com/golang/protobuf/proto"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type entry struct {
	compressedBytes []byte

	init        sync.Once
	unmarshaled *descriptor.FileDescriptorSet
	err         error
}

var registry = struct {
	sync.RWMutex
	entries map[string]*entry
}{entries: map[string]*entry{}}

// RegisterDescriptorSetCompressed registers a descriptor set for a set of services.
// Called from code generated by github.com/TriggerMail/luci-go/grpc/cmd/cproto
//
// compressedDescriptorSet must be a valid descriptor.FileDescriptorSet message
// compressed wit gzip.
// It must contain descriptions for all the services, their message types
// and all transitive dependencies.
//
// This call is cheap.
func RegisterDescriptorSetCompressed(serviceNames []string, compressedDescriptorSet []byte) {
	registry.Lock()
	defer registry.Unlock()
	e := &entry{compressedBytes: compressedDescriptorSet}
	for _, s := range serviceNames {
		registry.entries[s] = e
	}
}

func getEntry(serviceName string) *entry {
	registry.RLock()
	defer registry.RUnlock()
	return registry.entries[serviceName]
}

// GetDescriptorSet returns a descriptor set that contains the request service,
// its message types and all transitive dependencies.
// Returns (nil, nil) if the service descriptor is unknown.
//
// Do NOT modify the returned descriptor.
func GetDescriptorSet(serviceName string) (*descriptor.FileDescriptorSet, error) {
	e := getEntry(serviceName)
	if e == nil {
		return nil, nil
	}
	e.init.Do(func() {
		var unGzip io.Reader
		unGzip, e.err = gzip.NewReader(bytes.NewBuffer(e.compressedBytes))
		if e.err != nil {
			return
		}

		var uncompressed []byte
		uncompressed, e.err = ioutil.ReadAll(unGzip)
		if e.err != nil {
			return
		}

		var unmarshaled descriptor.FileDescriptorSet
		e.err = proto.Unmarshal(uncompressed, &unmarshaled)
		if e.err != nil {
			return
		}
		e.unmarshaled = &unmarshaled
	})
	return e.unmarshaled, e.err
}
