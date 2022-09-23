package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var listOnly bool
var podsOnly bool

type NodeList struct {
	Items []Node `json:"items"`
}

type Node struct {
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	Name        string            `json:"name,omitempty"`
	Annotations map[string]string `json:"annotations"`
}

func main() {
	flag.BoolVar(&listOnly, "l", false, "List current annotations and exist")
	flag.BoolVar(&podsOnly, "p", false, "Annotate pods")
	flag.Parse()

	embeddings := []string{""}
	url := ""
	if podsOnly {
		embeddings = []string{"#feffef", "#604010", "#00ff00", "#333333", "#aaaaaa", "#927a0c"}
		url = "http://127.0.0.1:8011/api/v1/namespaces/embeddings/pods"
	} else {
		embeddings = []string{"#ffffff", "#000000"}
		url = "http://127.0.0.1:8011/api/v1/nodes"
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		fmt.Println("Invalid status code", resp.Status)
		os.Exit(1)
	}

	var nodes NodeList
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&nodes)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if listOnly {
		for _, node := range nodes.Items {
			embedding := node.Metadata.Annotations["hightower.com/embedding"]
			fmt.Printf("%s %s\n", node.Metadata.Name, embedding)
		}
		os.Exit(0)
	}

	rand.Seed(time.Now().Unix())
	for _, node := range nodes.Items {
		embedding := embeddings[rand.Intn(len(embeddings))]
		annotations := map[string]string{
			"hightower.com/embedding": embedding,
		}
		patch := Node{
			Metadata{
				Annotations: annotations,
			},
		}

		var b []byte
		body := bytes.NewBuffer(b)
		err := json.NewEncoder(body).Encode(patch)
		if err != nil {
			fmt.Println("failed encode patch")
			fmt.Println(err)
			os.Exit(1)
		}

		urlfull := url + "/" + node.Metadata.Name
		request, err := http.NewRequest("PATCH", urlfull, body)
		if err != nil {
			fmt.Println("failed patch request")
			fmt.Println(err)
			os.Exit(1)
		}

		request.Header.Set("Content-Type", "application/strategic-merge-patch+json")
		request.Header.Set("Accept", "application/json, */*")

		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if resp.StatusCode != 200 {
			fmt.Println(resp)
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%s %s\n", node.Metadata.Name, embedding)
	}
}
