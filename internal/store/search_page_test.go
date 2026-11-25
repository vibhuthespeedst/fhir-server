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
	"context"
	"errors"
	"math"
	"testing"
)

func TestSearch_PageTooLargeIsInvalid(t *testing.T) {
	for _, page := range []int{math.MaxInt, math.MaxInt/2 + 1} {
		_, err := (&Store{}).Search(context.Background(), SearchParams{ResourceType: "Patient", Page: page, PageSize: 2})
		var invalid *InvalidParamError
		if !errors.As(err, &invalid) {
			t.Errorf("_page=%d: err = %v, want *InvalidParamError", page, err)
		}
	}
}
