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

func TestIfNoneExistQuery(t *testing.T) {
	for _, in := range []string{
		"identifier=urn:x|1",
		"?identifier=urn:x|1",
		"DocumentReference?identifier=urn:x|1",
		"http://localhost:9090/fhir/r4/DocumentReference?identifier=urn:x|1",
		"http://x.com/y=z/DocumentReference?identifier=urn:x|1",
	} {
		if got, err := IfNoneExistQuery(in); err != nil || got != "identifier=urn:x|1" {
			t.Errorf("IfNoneExistQuery(%q) = %q, %v", in, got, err)
		}
	}

	// Returned as is: a "?" inside a value is not a prefix, _count next to a
	// criterion is fine, and _list filters by the List's entries.
	for _, in := range []string{"identifier=urn:x|a?b&_count=1", "_list=abc"} {
		if got, err := IfNoneExistQuery(in); err != nil || got != in {
			t.Errorf("IfNoneExistQuery(%q) = %q, %v", in, got, err)
		}
	}

	for _, in := range []string{"?", "identifier", "identifier=", "identifier=urn:x|1&status=", "identifier=%zz", "_count=1", "_sort=name&_count=1"} {
		if _, err := IfNoneExistQuery(in); err == nil {
			t.Errorf("IfNoneExistQuery(%q): want an error", in)
		}
	}
}
