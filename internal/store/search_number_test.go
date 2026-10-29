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

import (
	"errors"
	"testing"
)

// TestBuildNumericExists_InvalidValueSetsError pins that a number or quantity
// value that doesn't parse surfaces as an InvalidParamError (HTTP 400) instead
// of being searched as 0.
func TestBuildNumericExists_InvalidValueSetsError(t *testing.T) {
	for _, value := range []string{"abc", "gtabc"} {
		b := &queryBuilder{rt: "RiskAssessment"}
		b.buildNumberExists("probability", value)
		var invalid *InvalidParamError
		if !errors.As(b.err, &invalid) {
			t.Errorf("number %q: b.err = %v, want *InvalidParamError", value, b.err)
		}
	}
	for _, value := range []string{"abc|http://unitsofmeasure.org|mg", "leabc"} {
		b := &queryBuilder{rt: "Observation"}
		b.buildQuantityExists("value-quantity", value)
		var invalid *InvalidParamError
		if !errors.As(b.err, &invalid) {
			t.Errorf("quantity %q: b.err = %v, want *InvalidParamError", value, b.err)
		}
	}
}
