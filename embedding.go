// Copyright 2016 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"strconv"
	"strings"
)

func parseEmbedding(embeddingCsv string) ([]float32, error) {
	embeddingStr := strings.Split(embeddingCsv, ",")
	embedding := make([]float32, len(embeddingStr))
	for i, s := range embeddingStr {
		val, error := strconv.ParseFloat(strings.Trim(s, " \t"), 32)
		if error != nil {
			return nil, error
		}
		embedding[i] = float32(val)
	}
	return embedding, nil
}
