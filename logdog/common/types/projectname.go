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

package types

import (
	"flag"

	"github.com/TriggerMail/luci-go/config"
)

// ProjectName is a luci-config project name.
//
// A valid project name may only include:
//  - Lowercase letters [a-z]
//  - Numbers [0-9]
//  - Hyphen (-)
//  - Underscore (_)
//
// It also must begin with a letter.
//
// See:
// https://github.com/luci/luci-py/blob/8e594074929871a9761d27e814541bc0d7d84744/appengine/components/components/config/common.py#L41
type ProjectName string

var _ flag.Value = (*ProjectName)(nil)

// Validate returns an error if the supplied ProjectName is not a valid project
// name.
func (p ProjectName) Validate() error {
	return config.ValidateProjectName(string(p))
}

// String implements flag.Value.
func (p *ProjectName) String() string {
	return string(*p)
}

// Set implements flag.Value.
func (p *ProjectName) Set(v string) error {
	vp := ProjectName(v)
	if err := vp.Validate(); err != nil {
		return err
	}
	*p = vp
	return nil
}
