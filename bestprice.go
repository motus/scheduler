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
	"math/rand"

	"gopkg.in/go-playground/colors.v1"
)

func bestPrice(nodes []Node, pod *Pod) (Node, error) {

	podEmbedding, ok := pod.Metadata.Annotations["hightower.com/embedding"]
	if !ok {
		return Node{}, errors.New("couldn't get annotation from pod")
	}
	fmt.Printf("Processing pod with string embedding: %s\n", podEmbedding)
	podEmbeddingParsed, _ := colors.ParseHEX(podEmbedding)

	// put nodes in random order, as they are currently assigned on first embedding match.
	// TODO: sort them by CPU availablity
	rand.Shuffle(len(nodes), func(i, j int) {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	})
	for _, n := range nodes {
		nodeEmbedding, ok := n.Metadata.Annotations["hightower.com/embedding"]
		if !ok {
			return Node{}, errors.New("couldn't get annotation from node")
		}
		nodeEmbeddingParsed, _ := colors.ParseHEX(nodeEmbedding)
		if nodeEmbeddingParsed.IsLight() == podEmbeddingParsed.IsLight() {
			fmt.Printf("Match found\n")
			return n, nil
		}
	}

	fmt.Printf("WARNING: Could not find a embedding-matching node, defaulting to first node\n")
	return nodes[0], nil
}
