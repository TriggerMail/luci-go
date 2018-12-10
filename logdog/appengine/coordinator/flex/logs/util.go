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

package logs

import (
	"github.com/TriggerMail/luci-go/common/proto/google"
	"github.com/TriggerMail/luci-go/logdog/api/endpoints/coordinator/logs/v1"
	"github.com/TriggerMail/luci-go/logdog/appengine/coordinator"
)

func buildLogStreamState(ls *coordinator.LogStream, lst *coordinator.LogStreamState) *logdog.LogStreamState {
	lss := logdog.LogStreamState{
		ProtoVersion:  ls.ProtoVersion,
		Created:       google.NewTimestamp(ls.Created),
		TerminalIndex: lst.TerminalIndex,
		Purged:        ls.Purged,
	}

	if ast := lst.ArchivalState(); ast.Archived() {
		lss.Archive = &logdog.LogStreamState_ArchiveInfo{
			IndexUrl:      lst.ArchiveIndexURL,
			StreamUrl:     lst.ArchiveStreamURL,
			DataUrl:       lst.ArchiveDataURL,
			Complete:      ast == coordinator.ArchivedComplete,
			LogEntryCount: lst.ArchiveLogEntryCount,
		}
	}

	return &lss
}
