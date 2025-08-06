// GoForGo Exercise: Hadoop HDFS Operations
// Learn how to interact with HDFS using Go WebHDFS REST API

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// TODO: Define HDFSClient struct
type HDFSClient struct {
	// Your HDFSClient struct here
}

// TODO: Implement HDFS operations: NewHDFSClient, CreateFile, ReadFile, ListDirectory, DeleteFile
func NewHDFSClient(nameNodeURL, username string) *HDFSClient {
	// Your implementation here
	return nil
}

func (h *HDFSClient) createFile(ctx context.Context, path string, data []byte) error {
	// Your implementation here
	return nil
}

func (h *HDFSClient) readFile(ctx context.Context, path string) ([]byte, error) {
	// Your implementation here
	return nil, nil
}

func main() {
	fmt.Println("Hadoop HDFS operations completed!")
}