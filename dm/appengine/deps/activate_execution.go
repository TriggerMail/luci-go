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

package deps

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/TriggerMail/luci-go/common/logging"
	dm "github.com/TriggerMail/luci-go/dm/api/service/v1"
	"github.com/TriggerMail/luci-go/dm/appengine/mutate"
)

func (d *deps) ActivateExecution(c context.Context, req *dm.ActivateExecutionReq) (ret *empty.Empty, err error) {
	ret = &empty.Empty{}
	logging.Fields{"execution": req.Auth.Id}.Infof(c, "activating")
	err = tumbleNow(c, &mutate.ActivateExecution{
		Auth:   req.Auth,
		NewTok: req.ExecutionToken,
	})
	return
}
