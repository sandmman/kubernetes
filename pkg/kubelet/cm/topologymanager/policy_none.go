/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package topologymanager

import "context"

type nonePolicy struct{}

var _ Policy = &nonePolicy{}

// PolicyNone policy name.
const PolicyNone string = "none"

// NewNonePolicy returns none policy.
func NewNonePolicy() Policy {
	return &nonePolicy{}
}

func (p *nonePolicy) Name() string {
	return PolicyNone
}

func (p *nonePolicy) canAdmitPodResult(hint *TopologyHint) bool {
	return true
}

func (p *nonePolicy) Merge(_ context.Context, providersHints []map[string][]TopologyHint) (TopologyHint, bool) {
	return TopologyHint{}, p.canAdmitPodResult(nil)
}
