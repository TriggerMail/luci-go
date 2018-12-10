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

// Package lucictx implements a Go client for the protocol defined here:
//   https://github.com/luci/luci-py/blob/master/client/LUCI_CONTEXT.md
//
// It differs from the python client in a couple ways:
//   * The initial LUCI_CONTEXT value is captured once at application start.
//   * Writes are cached into the golang context.Context, not a global variable.
//   * The LUCI_CONTEXT environment variable is not changed automatically when
//     using the Set function. To pass the new context on to a child process,
//     you must use the Export function to dump the current context state to
//     disk and call exported.SetInCmd(cmd) to configure new command's
//     environment.
package lucictx

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"sync"

	"github.com/TriggerMail/luci-go/common/errors"
)

// EnvKey is the environment variable key for the LUCI_CONTEXT file.
const EnvKey = "LUCI_CONTEXT"

// lctx is wrapper around top-level JSON dict of a LUCI_CONTEXT file.
//
// Note that we must use '*json.RawMessage' as dict value type since only
// pointer *json.RawMessage type implements json.Marshaler and json.Unmarshaler
// interfaces. Without '*' the JSON library treats json.RawMessage as []byte and
// marshals it as base64 blob.
type lctx struct {
	sections map[string]*json.RawMessage // readonly! lives outside the lock

	lock sync.Mutex
	path string // non-empty if exists as file on disk
	refs int    // number of open references to the dropped file
}

func alloc(size int) *lctx {
	return &lctx{sections: make(map[string]*json.RawMessage, size)}
}

func (l *lctx) clone() *lctx {
	ret := alloc(len(l.sections))
	for k, v := range l.sections {
		ret.sections[k] = v
	}
	return ret
}

var lctxKey = "Holds the current lctx"

// This is the LUCI_CONTEXT loaded from the environment once when the process
// starts.
var externalContext = extractFromEnv(os.Stderr)

func extractFromEnv(out io.Writer) *lctx {
	path := os.Getenv(EnvKey)
	if path == "" {
		return &lctx{}
	}
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(out, "Could not open LUCI_CONTEXT file %q: %s", path, err)
		return &lctx{}
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	dec.UseNumber()
	tmp := map[string]interface{}{}
	if err := dec.Decode(&tmp); err != nil {
		fmt.Fprintf(out, "Could not decode LUCI_CONTEXT file %q: %s", path, err)
		return &lctx{}
	}

	ret := alloc(len(tmp))
	for k, v := range tmp {
		if reflect.TypeOf(v).Kind() != reflect.Map {
			fmt.Fprintf(out, "Could not reencode LUCI_CONTEXT file %q, section %q: Not a map.", path, k)
			continue
		}
		item, _ := json.Marshal(v)
		// This section just came from json.Unmarshal, so we know that json.Marshal
		// will work on it.
		raw := json.RawMessage(item)
		ret.sections[k] = &raw
	}

	ret.path = path // reuse existing external file in Export()
	ret.refs = 1    // never decremented, ensuring we don't delete the external file
	return ret
}

// Note: it never returns nil.
func getCurrent(ctx context.Context) *lctx {
	if l := ctx.Value(&lctxKey); l != nil {
		return l.(*lctx)
	}
	return externalContext
}

// Get retrieves the current section from the current LUCI_CONTEXT, and
// deserializes it into out. Out may be any target for json.Unmarshal. If the
// section exists, it deserializes it into the provided out object. If not, then
// out is unmodified.
func Get(ctx context.Context, section string, out interface{}) error {
	_, err := Lookup(ctx, section, out)
	return err
}

// Lookup retrieves the current section from the current LUCI_CONTEXT, and
// deserializes it into out. Out may be any target for json.Unmarshal. It
// returns a deserialization error (if any), and a boolean indicating if the
// section was actually found.
func Lookup(ctx context.Context, section string, out interface{}) (bool, error) {
	data, _ := getCurrent(ctx).sections[section]
	if data == nil {
		return false, nil
	}
	return true, json.Unmarshal(*data, out)
}

// Set writes the json serialization of `in` as the given section into the
// LUCI_CONTEXT, returning the new ctx object containing it. This ctx can be
// passed to Export to serialize it to disk.
//
// If in is nil, it will clear that section of the LUCI_CONTEXT.
//
// The returned context is always safe to use, even if this returns an error.
// This only returns an error if `in` cannot be marshalled to a JSON Object.
func Set(ctx context.Context, section string, in interface{}) (context.Context, error) {
	var err error
	var data json.RawMessage
	if in != nil {
		if data, err = json.Marshal(in); err != nil {
			return ctx, err
		}
		if data[0] != '{' {
			return ctx, errors.New("LUCI_CONTEXT sections must always be JSON Objects")
		}
	}
	cur := getCurrent(ctx)
	if _, alreadyHas := cur.sections[section]; data == nil && !alreadyHas {
		// Removing a section which is already missing is a no-op
		return ctx, nil
	}
	newLctx := cur.clone()
	if data == nil {
		delete(newLctx.sections, section)
	} else {
		newLctx.sections[section] = &data
	}
	return context.WithValue(ctx, &lctxKey, newLctx), nil
}

// Export takes the current LUCI_CONTEXT information from ctx, writes it to
// a file in os.TempDir and returns a wrapping Exported object. This exported
// value must then be installed into the environment of any subcommands (see
// the methods on Exported).
//
// It is required that the caller of this function invoke Close() on the
// returned Exported object, or they will leak temporary files.
//
// Internally this function reuses existing files, when possible, so if you
// anticipate calling a lot of subcommands with exported LUCI_CONTEXT, you can
// export it in advance (thus grabbing a reference to the exported file). Then
// subsequent Export() calls with this context will be extremely cheap, since
// they will just reuse the existing file. Don't forget to release it with
// Close() when done.
func Export(ctx context.Context) (Exported, error) {
	return getCurrent(ctx).export()
}

func (ctx *lctx) export() (Exported, error) {
	if len(ctx.sections) == 0 {
		return &nullExport{}, nil
	}

	ctx.lock.Lock()
	defer ctx.lock.Unlock()

	if ctx.refs == 0 {
		if ctx.path != "" {
			panic("lctx.path is supposed to be empty here")
		}
		path, err := dropToDisk(ctx.sections)
		if err != nil {
			return nil, err
		}
		ctx.path = path
	}

	ctx.refs++
	return &liveExport{
		path: ctx.path,
		closer: func() {
			ctx.lock.Lock()
			defer ctx.lock.Unlock()
			if ctx.refs == 0 {
				panic("lctx.refs can't be zero here")
			}
			ctx.refs--
			if ctx.refs == 0 {
				removeFromDisk(ctx.path)
				ctx.path = ""
			}
		},
	}, nil
}

func dropToDisk(sections map[string]*json.RawMessage) (string, error) {
	// Note: this makes a file in 0600 mode. This is what we want, the context
	// may have secrets.
	f, err := ioutil.TempFile("", "luci_context.")
	if err != nil {
		return "", errors.Annotate(err, "creating luci_context file").Err()
	}

	err = json.NewEncoder(f).Encode(sections)
	f.Close() // intentionally do this even on error.
	if err != nil {
		removeFromDisk(f.Name())
		return "", errors.Annotate(err, "writing luci_context").Err()
	}

	return f.Name(), nil
}

func removeFromDisk(path string) {
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Could not remove LUCI_CONTEXT file %q: %s", path, err)
	}
}
