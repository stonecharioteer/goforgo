// GoForGo Solution: Elasticsearch Aggregations
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type AggregationClient struct {
	client *elasticsearch.Client
}

func NewAggregationClient() (*AggregationClient, error) {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		return nil, err
	}
	return &AggregationClient{client: client}, nil
}

func (a *AggregationClient) performAggregation(ctx context.Context, index string) (map[string]interface{}, error) {
	aggBody := map[string]interface{}{
		"size": 0,
		"aggs": map[string]interface{}{
			"authors": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "author.keyword",
				},
			},
		},
	}
	
	body, _ := json.Marshal(aggBody)
	
	res, err := a.client.Search(
		a.client.Search.WithContext(ctx),
		a.client.Search.WithIndex(index),
		a.client.Search.WithBody(strings.NewReader(string(body))),
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
	client, err := NewAggregationClient()
	if err != nil {
		log.Fatal(err)
	}
	
	results, err := client.performAggregation(context.Background(), "articles")
	if err != nil {
		log.Printf("Error performing aggregation: %v", err)
	} else {
		fmt.Printf("Aggregation results: %+v\n", results)
	}
	
	fmt.Println("Elasticsearch aggregations operations completed!")
}