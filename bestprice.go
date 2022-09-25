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
	"errors"
	"fmt"
	"math"
)

func bestPrice(nodes []Node, pod *Pod) (Node, error) {

	podEmbeddingStr, ok := pod.Metadata.Annotations["hightower.com/embedding"]
	if !ok {
		return Node{}, errors.New("couldn't get annotation from pod")
	}

	fmt.Printf("Processing pod with embedding string: %s\n", podEmbeddingStr)
	podEmbedding, error := parseEmbedding(podEmbeddingStr)
	if error != nil {
		return Node{}, error
	}

	bestNode := nodes[0]
	bestDistance := float32(math.Inf(1))
	for _, n := range nodes {

		nodeEmbeddingStr, ok := n.Metadata.Annotations["hightower.com/embedding"]
		if !ok {
			return Node{}, errors.New("couldn't get annotation from node")
		}

		nodeEmbedding, error := parseEmbedding(nodeEmbeddingStr)
		if error != nil {
			return Node{}, error
		}

		dist, error := euclideanDistance(podEmbedding, nodeEmbedding)
		if error != nil {
			return Node{}, error
		}

		if dist <= bestDistance {
			bestDistance = dist
			bestNode = n
		}
	}

	return bestNode, nil
}
