// GoForGo Solution: Elasticsearch Indexing
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type Document struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type ElasticsearchClient struct {
	client *elasticsearch.Client
}

func NewElasticsearchClient(addresses []string) (*ElasticsearchClient, error) {
	cfg := elasticsearch.Config{
		Addresses: addresses,
	}
	
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	
	return &ElasticsearchClient{client: client}, nil
}

func (es *ElasticsearchClient) indexDocument(ctx context.Context, index, docID string, doc Document) error {
	data, err := json.Marshal(doc)
	if err != nil {
		return err
	}
	
	res, err := es.client.Index(
		index,
		strings.NewReader(string(data)),
		es.client.Index.WithDocumentID(docID),
		es.client.Index.WithContext(ctx),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	
	log.Printf("Indexed document %s in index %s", docID, index)
	return nil
}

func main() {
	client, err := NewElasticsearchClient([]string{"http://localhost:9200"})
	if err != nil {
		log.Fatal(err)
	}
	
	doc := Document{
		Title:   "Go and Elasticsearch",
		Content: "Learning to integrate Go with Elasticsearch",
		Author:  "Go Developer",
	}
	
	if err := client.indexDocument(context.Background(), "articles", "1", doc); err != nil {
		log.Printf("Error indexing document: %v", err)
	}
	
	fmt.Println("Elasticsearch indexing operations completed!")
}