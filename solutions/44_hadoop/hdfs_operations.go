// GoForGo Solution: Hadoop HDFS Operations
package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type HDFSClient struct {
	baseURL  string
	username string
	client   *http.Client
}

func NewHDFSClient(nameNodeURL, username string) *HDFSClient {
	return &HDFSClient{
		baseURL:  nameNodeURL,
		username: username,
		client:   &http.Client{},
	}
}

func (h *HDFSClient) createFile(ctx context.Context, path string, data []byte) error {
	// Create file using WebHDFS REST API
	createURL := fmt.Sprintf("%s/webhdfs/v1%s?op=CREATE&user.name=%s", h.baseURL, path, h.username)
	
	req, err := http.NewRequestWithContext(ctx, "PUT", createURL, bytes.NewReader(data))
	if err != nil {
		return err
	}
	
	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	
	log.Printf("Created HDFS file: %s", path)
	return nil
}

func (h *HDFSClient) readFile(ctx context.Context, path string) ([]byte, error) {
	readURL := fmt.Sprintf("%s/webhdfs/v1%s?op=OPEN&user.name=%s", h.baseURL, path, h.username)
	
	req, err := http.NewRequestWithContext(ctx, "GET", readURL, nil)
	if err != nil {
		return nil, err
	}
	
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	log.Printf("Read HDFS file: %s (%d bytes)", path, len(data))
	return data, nil
}

func main() {
	client := NewHDFSClient("http://localhost:9870", "hadoop")
	
	ctx := context.Background()
	testData := []byte("Hello HDFS from Go!")
	
	if err := client.createFile(ctx, "/test/sample.txt", testData); err != nil {
		log.Printf("Error creating file: %v", err)
	}
	
	if data, err := client.readFile(ctx, "/test/sample.txt"); err != nil {
		log.Printf("Error reading file: %v", err)
	} else {
		log.Printf("File content: %s", string(data))
	}
	
	fmt.Println("Hadoop HDFS operations completed!")
}