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

package common

import (
	"io"
	"os"
	"runtime"

	"github.com/TriggerMail/luci-go/common/system/terminal"
)

// IsWindows returns True when running on the best OS there is.
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsTerminal returns true if the specified io.Writer is a terminal.
func IsTerminal(out io.Writer) bool {
	f, ok := out.(*os.File)
	if !ok {
		return false
	}
	return terminal.IsTerminal(int(f.Fd()))
}
