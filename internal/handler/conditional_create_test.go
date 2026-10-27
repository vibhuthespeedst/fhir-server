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

package handler_test

import (
	"context"
	"net/http"
	"testing"
)

func TestCreate_IfNoneExist(t *testing.T) {
	var gotQuery string
	ms := &mockStore{}
	ms.conditionalMatchFn = func(_ context.Context, _, rawQuery string) (string, int, error) {
		gotQuery = rawQuery
		return "", 0, nil
	}
	ms.createFn = func(_ context.Context, _ string, body map[string]any) (map[string]any, error) {
		return body, nil
	}
	h := newRouter(ms)
	patient := map[string]any{"resourceType": "Patient"}

	w := doWithHeaders(t, h, http.MethodPost, "/fhir/r4/Patient", patient,
		map[string]string{"If-None-Exist": "http://localhost/fhir/r4/Patient?identifier=urn:x|1"})
	if w.Code != http.StatusCreated || gotQuery != "identifier=urn:x|1" {
		t.Errorf("full URL: got %d, matched on %q", w.Code, gotQuery)
	}

	w = doWithHeaders(t, h, http.MethodPost, "/fhir/r4/Patient", patient,
		map[string]string{"If-None-Exist": "identifier"})
	if w.Code != http.StatusBadRequest {
		t.Errorf("no value: want 400, got %d", w.Code)
	}
}
