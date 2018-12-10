// Copyright 2018 The LUCI Authors.
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

package git

import (
	"context"
	"testing"

	"go.chromium.org/gae/impl/memory"
	"github.com/TriggerMail/luci-go/auth/identity"
	"github.com/TriggerMail/luci-go/milo/common"
	"github.com/TriggerMail/luci-go/server/auth"
	"github.com/TriggerMail/luci-go/server/auth/authtest"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTagError(t *testing.T) {
	t.Parallel()

	Convey("tagError Works", t, func() {
		c := memory.Use(context.Background())
		cUser := auth.WithState(c, &authtest.FakeState{Identity: "user:user@example.com"})
		cAnon := auth.WithState(c, &authtest.FakeState{Identity: identity.AnonymousIdentity})

		So(common.ErrorCodeIn(tagError(cAnon, errGRPCNotFound)), ShouldEqual, common.CodeUnauthorized)
		So(common.ErrorCodeIn(tagError(cUser, errGRPCNotFound)), ShouldEqual, common.CodeNotFound)
	})
}
