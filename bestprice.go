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

	"gopkg.in/go-playground/colors.v1"
)

func bestPrice(nodes []Node, pod *Pod) (Node, error) {

	podColor, ok := pod.Metadata.Annotations["hightower.com/color"]
	if !ok {
		return Node{}, errors.New("couldn't get annotation from pod")
	}
	fmt.Printf("Processing pod with string color: %s\n", podColor)
	podColorParsed, _ := colors.ParseHEX(podColor)

	for _, n := range nodes {
		nodeColor, ok := n.Metadata.Annotations["hightower.com/color"]
		if !ok {
			return Node{}, errors.New("couldn't get annotation from node")
		}
		nodeColorParsed, _ := colors.ParseHEX(nodeColor)
		if nodeColorParsed.IsLight() == podColorParsed.IsLight() {
			fmt.Printf("Match found\n")
			return n, nil
		}
	}

	fmt.Printf("WARNING: Could not find a color-matching node, defaulting to first node\n")
	return nodes[0], nil
}
