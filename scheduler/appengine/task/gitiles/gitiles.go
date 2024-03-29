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

package gitiles

import (
	"context"
	"fmt"
	"net/url"
	"sort"
	"time"

	"github.com/golang/protobuf/proto"

	"google.golang.org/api/pubsub/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/TriggerMail/luci-go/common/api/gitiles"
	"github.com/TriggerMail/luci-go/common/clock"
	"github.com/TriggerMail/luci-go/common/errors"
	"github.com/TriggerMail/luci-go/common/logging"
	"github.com/TriggerMail/luci-go/common/proto/git"
	gitilespb "github.com/TriggerMail/luci-go/common/proto/gitiles"
	"github.com/TriggerMail/luci-go/common/retry/transient"
	"github.com/TriggerMail/luci-go/common/sync/parallel"
	"github.com/TriggerMail/luci-go/config/validation"
	"github.com/TriggerMail/luci-go/server/auth"

	api "github.com/TriggerMail/luci-go/scheduler/api/scheduler/v1"
	"github.com/TriggerMail/luci-go/scheduler/appengine/internal"
	"github.com/TriggerMail/luci-go/scheduler/appengine/messages"
	"github.com/TriggerMail/luci-go/scheduler/appengine/task"
)

// gitilesRPCTimeout limits how long Gitiles RPCs are allowed to last.
const gitilesRPCTimeout = time.Minute

// defaultMaxTriggersPerInvocation limits number of triggers emitted per one
// invocation.
const defaultMaxTriggersPerInvocation = 100

// defaultMaxCommitsPerRefUpdate limits number of commits (and hence triggers)
// emitted when a ref changes.
// Must be smaller than defaultMaxTriggersPerInvocation, else these many
// triggers could be emitted.
const defaultMaxCommitsPerRefUpdate = 50

// TaskManager implements task.Manager interface for tasks defined with
// GitilesTask proto message.
type TaskManager struct {
	mockGitilesClient        gitilespb.GitilesClient // Used for testing only.
	maxTriggersPerInvocation int                     // Avoid choking on DS or runtime limits.
	maxCommitsPerRefUpdate   int                     // Failsafe when someone pushes too many commits at once.
}

// Name is part of Manager interface.
func (m TaskManager) Name() string {
	return "gitiles"
}

// ProtoMessageType is part of Manager interface.
func (m TaskManager) ProtoMessageType() proto.Message {
	return (*messages.GitilesTask)(nil)
}

// Traits is part of Manager interface.
func (m TaskManager) Traits() task.Traits {
	return task.Traits{
		Multistage: false, // we don't use task.StatusRunning state
	}
}

// ValidateProtoMessage is part of Manager interface.
func (m TaskManager) ValidateProtoMessage(c *validation.Context, msg proto.Message) {
	cfg, ok := msg.(*messages.GitilesTask)
	if !ok {
		c.Errorf("wrong type %T, expecting *messages.GitilesTask", msg)
		return
	}

	// Validate 'repo' field.
	c.Enter("repo")
	if cfg.Repo == "" {
		c.Errorf("field 'repository' is required")
	} else {
		u, err := url.Parse(cfg.Repo)
		if err != nil {
			c.Errorf("invalid URL %q: %s", cfg.Repo, err)
		} else if !u.IsAbs() {
			c.Errorf("not an absolute url: %q", cfg.Repo)
		}
	}
	c.Exit()

	c.Enter("refs")
	gitiles.ValidateRefSet(c, cfg.Refs)
	c.Exit()
}

// LaunchTask is part of Manager interface.
func (m TaskManager) LaunchTask(c context.Context, ctl task.Controller) error {
	cfg := ctl.Task().(*messages.GitilesTask)
	ctl.DebugLog("Repo: %s, Refs: %s", cfg.Repo, cfg.Refs)

	g, err := m.getGitilesClient(c, ctl, cfg)
	if err != nil {
		return err
	}
	// TODO(tandrii): use g.host, g.project for saving/loading state
	// instead of repoURL.
	repoURL, err := url.Parse(cfg.Repo)
	if err != nil {
		return err
	}

	refs, err := m.fetchRefsState(c, ctl, cfg, g, repoURL)
	if err != nil {
		ctl.DebugLog("Error fetching state of the world: %s", err)
		return err
	}

	refs.pruneKnown(ctl)
	leftToProcess, err := m.emitTriggersRefAtATime(c, ctl, g, cfg.Repo, refs)

	if err != nil {
		switch {
		case leftToProcess == 0:
			panic(err) // leftToProcess must include the one processing of which failed.
		case refs.changed > 0:
			// Even though we hit error, we had progress. So, ignore error as transient.
			ctl.DebugLog("ignoring error %s as transient", err)
		default:
			ctl.DebugLog("no progress made due to %s", err)
			return err
		}
	}

	switch {
	case leftToProcess > 0 && refs.changed == 0:
		panic(errors.New("no progress with no errors must not happen"))
	case refs.changed == 0:
		ctl.DebugLog("No changes detected")
		ctl.State().Status = task.StatusSucceeded
		return nil
	case leftToProcess > 0:
		ctl.DebugLog("%d changed refs processed, %d refs not yet examined", refs.changed, leftToProcess)
	default:
		ctl.DebugLog("All %d changed refs processed", refs.changed)
	}
	// Force save to ensure triggers are actually emitted.
	if err := ctl.Save(c); err != nil {
		// At this point, triggers have not been sent, so bail now and don't save
		// the refs' heads newest values.
		return err
	}
	if err := saveState(c, ctl.JobID(), repoURL, refs.known); err != nil {
		return err
	}
	ctl.DebugLog("Saved %d known refs", len(refs.known))
	ctl.State().Status = task.StatusSucceeded
	return nil
}

// AbortTask is part of Manager interface.
func (m TaskManager) AbortTask(c context.Context, ctl task.Controller) error {
	return nil
}

// HandleNotification is part of Manager interface.
func (m TaskManager) HandleNotification(c context.Context, ctl task.Controller, msg *pubsub.PubsubMessage) error {
	return errors.New("not implemented")
}

// HandleTimer is part of Manager interface.
func (m TaskManager) HandleTimer(c context.Context, ctl task.Controller, name string, payload []byte) error {
	return errors.New("not implemented")
}

func (m TaskManager) fetchRefsState(c context.Context, ctl task.Controller, cfg *messages.GitilesTask, g *gitilesClient, repoURL *url.URL) (*refsState, error) {
	refs := &refsState{}
	refs.watched = gitiles.NewRefSet(cfg.GetRefs())
	return refs, parallel.FanOutIn(func(work chan<- func() error) {
		work <- func() (loadErr error) {
			refs.known, loadErr = loadState(c, ctl.JobID(), repoURL)
			return
		}
		work <- func() (resolveErr error) {
			c, cancel := clock.WithTimeout(c, gitilesRPCTimeout)
			defer cancel()
			refs.current, resolveErr = refs.watched.Resolve(c, g, g.project)
			return
		}
	})
}

// emitTriggersRefAtATime processes refs one a time and emits triggers if ref
// changed. Limits number of triggers emitted and so may stop early.
//
// Returns how many refs were not examined.
func (m TaskManager) emitTriggersRefAtATime(c context.Context, ctl task.Controller, g *gitilesClient, repo string, refs *refsState) (int, error) {
	// Safeguard against too many changes such as the first run after config
	// change to watch many more refs than before.
	maxTriggersPerInvocation := m.maxTriggersPerInvocation
	if maxTriggersPerInvocation <= 0 {
		maxTriggersPerInvocation = defaultMaxTriggersPerInvocation
	}
	maxCommitsPerRefUpdate := m.maxCommitsPerRefUpdate
	if maxCommitsPerRefUpdate <= 0 {
		maxCommitsPerRefUpdate = defaultMaxCommitsPerRefUpdate
	}
	emittedTriggers := 0
	// Note, that refs.current contain only watched refs (see getRefsTips).
	// For determinism, sort refs by name.
	sortedRefs := refs.sortedCurrentRefNames()
	for i, ref := range sortedRefs {
		commits, err := refs.newCommits(c, ctl, g, ref, maxCommitsPerRefUpdate)
		if err != nil {
			// This ref counts as not yet examined.
			return len(sortedRefs) - i, err
		}
		for i := range commits {
			// commit[0] is latest, so emit triggers in reverse order of commits.
			commit := commits[len(commits)-i-1]
			ctl.EmitTrigger(c, &internal.Trigger{
				Id:           fmt.Sprintf("%s/+/%s@%s", repo, ref, commit.Id),
				Created:      commit.Committer.Time,
				OrderInBatch: int64(emittedTriggers),
				Title:        commit.Id,
				Url:          fmt.Sprintf("%s/+/%s", repo, commit.Id),
				Payload: &internal.Trigger_Gitiles{
					Gitiles: &api.GitilesTrigger{Repo: repo, Ref: ref, Revision: commit.Id},
				},
			})
			emittedTriggers++
		}
		// Stop early if next iteration can't emit maxCommitsPerRefUpdate triggers.
		// But do so only after first successful fetch to ensure progress if
		// misconfigured.
		if emittedTriggers+maxCommitsPerRefUpdate > maxTriggersPerInvocation {
			ctl.DebugLog("Emitted %d triggers, postponing the rest", emittedTriggers)
			return len(sortedRefs) - i - 1, nil
		}
	}
	return 0, nil
}

func (m TaskManager) getGitilesClient(c context.Context, ctl task.Controller, cfg *messages.GitilesTask) (*gitilesClient, error) {
	host, project, err := gitiles.ParseRepoURL(cfg.Repo)
	if err != nil {
		return nil, errors.Annotate(err, "invalid repo URL %q", cfg.Repo).Err()
	}
	r := &gitilesClient{host: host, project: project}

	if m.mockGitilesClient != nil {
		// Used for testing only.
		logging.Infof(c, "using mockGitilesClient")
		r.GitilesClient = m.mockGitilesClient
		return r, nil
	}

	httpClient, err := ctl.GetClient(c, auth.WithScopes(gitiles.OAuthScope))
	if err != nil {
		return nil, err
	}
	if r.GitilesClient, err = gitiles.NewRESTClient(httpClient, host, true); err != nil {
		return nil, err
	}
	return r, nil
}

// gitilesClient embeds GitilesClient with useful metadata.
type gitilesClient struct {
	gitilespb.GitilesClient
	host    string // Gitiles host
	project string // Gitiles project
}

type refsState struct {
	watched gitiles.RefSet
	known   map[string]string // HEADs we saw before
	current map[string]string // HEADs available now
	changed int
}

func (s *refsState) pruneKnown(ctl task.Controller) {
	for ref := range s.known {
		switch {
		case !s.watched.Has(ref):
			ctl.DebugLog("Ref %s is no longer watched", ref)
			delete(s.known, ref)
			s.changed++
		case s.current[ref] == "":
			ctl.DebugLog("Ref %s deleted", ref)
			delete(s.known, ref)
			s.changed++
		}
	}
}

func (s *refsState) sortedCurrentRefNames() []string {
	sortedRefs := make([]string, 0, len(s.current))
	for ref := range s.current {
		sortedRefs = append(sortedRefs, ref)
	}
	sort.Strings(sortedRefs)
	return sortedRefs
}

// newCommits finds new commits for a given ref.
//
// If ref is new, returns only ref's HEAD,
// For updated refs, at most maxCommits of gitiles.Log(new..old)
func (s *refsState) newCommits(c context.Context, ctl task.Controller, g *gitilesClient, ref string, maxCommits int) ([]*git.Commit, error) {
	newHead := s.current[ref]
	oldHead, existed := s.known[ref]
	switch {
	case !existed:
		ctl.DebugLog("Ref %s is new: %s", ref, newHead)
		maxCommits = 1
	case oldHead != newHead:
		ctl.DebugLog("Ref %s updated: %s => %s", ref, oldHead, newHead)
	default:
		return nil, nil // no change
	}

	c, cancel := clock.WithTimeout(c, gitilesRPCTimeout)
	defer cancel()

	commits, err := gitiles.PagingLog(c, g, gitilespb.LogRequest{
		Project:            g.project,
		Committish:         newHead,
		ExcludeAncestorsOf: oldHead, // empty if ref is new, but then maxCommits is 1.
		PageSize:           int32(maxCommits),
	}, maxCommits)
	switch status.Code(err) {
	case codes.OK:
		// Happy fast path.
		// len(commits) may be 0 if this ref had a force push reverting to some
		// older revision. TODO(tAndrii): consider emitting trigger with just
		// newHead commit if there is a compelling use case.
		s.known[ref] = newHead
		s.changed++
		return commits, nil
	case codes.NotFound:
		// handled below.
		break
	default:
		// Any other error is presumably transient, so we'll retry.
		ctl.DebugLog("Ref %s: failed to fetch log between old %s and new %s revs", ref, oldHead, newHead)
		return nil, transient.Tag.Apply(err)
	}
	// Either:
	//  (1) oldHead is no longer known in gitiles (force push),
	//  (2) newHead is no longer known in gitiles (eventual consistency,
	//     or concurrent force push executed just now, or ACLs change)
	//  (3) gitiles accidental 404, aka fluke.
	// In cases (2) and (3), retries should clear the problem, while (1) we
	// should handle now.
	if !existed {
		// There was no oldHead, so definitely not (1). Retry later.
		ctl.DebugLog("Ref %s: log of first rev %s not found", ref, newHead)
		return nil, transient.Tag.Apply(err)
	}
	ctl.DebugLog("Ref %s: log old..new is not found, investigating further...", ref)

	// Fetch log of newHead only.
	commits, newErr := gitiles.PagingLog(c, g, gitilespb.LogRequest{
		Project:    g.project,
		Committish: newHead,
	}, 1)
	if newErr != nil {
		ctl.DebugLog("Ref %s: failed to fetch even log of just new rev %s %s", ref, newHead, err)
		return nil, transient.Tag.Apply(newErr)
	}
	// Fetch log of oldHead only.
	_, errOld := gitiles.PagingLog(c, g, gitilespb.LogRequest{
		Project:    g.project,
		Committish: oldHead,
	}, 1)
	switch status.Code(errOld) {
	case codes.NotFound:
		// This is case (1). Since we've already fetched just 1 commit from
		// newHead, we are done.
		ctl.DebugLog("Ref %s: force push detected; emitting trigger for new head", ref)
		s.known[ref] = newHead
		s.changed++
		return commits, nil
	case codes.OK:
		ctl.DebugLog("Ref %s: weirdly, log(%s) and log(%s) work, but not log(%s..%s)",
			ref, oldHead, newHead, oldHead, newHead)
		return nil, transient.Tag.Apply(err)
	default:
		// Any other error is presumably transient, so we'll retry.
		ctl.DebugLog("Ref %s: failed to fetch log of just old rev %s: %s", ref, oldHead, errOld)
		return nil, transient.Tag.Apply(err)
	}
}
