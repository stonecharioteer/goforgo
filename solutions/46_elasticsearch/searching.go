// GoForGo Solution: Elasticsearch Searching
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type SearchClient struct {
	client *elasticsearch.Client
}

func NewSearchClient() (*SearchClient, error) {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}
	return &SearchClient{client: client}, nil
}

func (s *SearchClient) search(ctx context.Context, index, query string) (map[string]interface{}, error) {
	searchBody := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"content": query,
			},
		},
	}
	
	body, _ := json.Marshal(searchBody)
	
	res, err := s.client.Search(
		s.client.Search.WithContext(ctx),
		s.client.Search.WithIndex(index),
		s.client.Search.WithBody(strings.NewReader(string(body))),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	
	var result map[string]interface{}
	json.NewDecoder(res.Body).Decode(&result)
	
	return result, nil
}

func main() {
	client, err := NewSearchClient()
	if err != nil {
		log.Fatal(err)
	}
	
	results, err := client.search(context.Background(), "articles", "Go")
	if err != nil {
		log.Printf("Error searching: %v", err)
	} else {
		fmt.Printf("Search results: %+v\n", results)
	}
	
	fmt.Println("Elasticsearch searching operations completed!")
}