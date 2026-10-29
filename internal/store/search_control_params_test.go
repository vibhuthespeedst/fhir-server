// Copyright (c) 2026, WSO2 LLC. (https://www.wso2.com).
//
// WSO2 LLC. licenses this file to you under the Apache License,
// Version 2.0 (the "License"); you may not use this file except
// in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied. See the License for the
// specific language governing permissions and limitations
// under the License.

package store

import "testing"

// _pretty only asks for formatted output, so it must not filter the search.
func TestApplyParam_PrettyIsNotACriterion(t *testing.T) {
	b := &queryBuilder{rt: "Patient"}
	b.applyParam("_pretty", "true")
	if b.where.Len() != 0 || b.err != nil {
		t.Errorf("_pretty changed the search: where = %q, err = %v", b.where.String(), b.err)
	}
}
