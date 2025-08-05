// http_client.go - SOLUTION
// Learn HTTP client operations in Go

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	fmt.Println("=== Basic GET Request ===")
	
	// Make a GET request
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Printf("GET error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Read error: %v\n", err)
		return
	}
	
	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Response: %s\n", body)
	
	// Decode JSON response
	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Printf("JSON decode error: %v\n", err)
	} else {
		fmt.Printf("Post: %+v\n", post)
	}
	
	fmt.Println("\n=== POST Request with JSON ===")
	
	// Create new post data
	newPost := Post{
		UserID: 1,
		Title:  "My New Post",
		Body:   "This is the content of my new post.",
	}
	
	// Marshal to JSON
	jsonData, err := json.Marshal(newPost)
	if err != nil {
		fmt.Printf("JSON marshal error: %v\n", err)
		return
	}
	
	// Make POST request
	resp, err = http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("POST error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	// Read POST response
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Read error: %v\n", err)
		return
	}
	
	fmt.Printf("POST Status: %s\n", resp.Status)
	fmt.Printf("POST Response: %s\n", body)
	
	fmt.Println("\n=== Custom HTTP Client ===")
	
	// Create custom client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	
	// Create request with custom headers
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts", nil)
	if err != nil {
		fmt.Printf("Request creation error: %v\n", err)
		return
	}
	
	// Add custom headers
	req.Header.Set("User-Agent", "GoForGo-Client/1.0")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Custom-Header", "CustomValue")
	
	// Execute request with custom client
	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("Custom client error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	fmt.Printf("Custom client status: %s\n", resp.Status)
	
	// Read response headers
	fmt.Println("Response headers:")
	for name, values := range resp.Header {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", name, value)
		}
	}
	
	fmt.Println("\n=== Multiple Requests ===")
	
	// Make multiple concurrent requests
	urls := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
	}
	
	// Create channels for results
	results := make(chan string, len(urls))
	errors := make(chan error, len(urls))
	
	// Start concurrent requests
	for i, url := range urls {
		go func(id int, url string) {
			// Make request and send result to channel
			resp, err := http.Get(url)
			if err != nil {
				errors <- fmt.Errorf("request %d failed: %v", id+1, err)
				return
			}
			defer resp.Body.Close()
			
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				errors <- fmt.Errorf("read %d failed: %v", id+1, err)
				return
			}
			
			var post Post
			if err := json.Unmarshal(body, &post); err != nil {
				errors <- fmt.Errorf("decode %d failed: %v", id+1, err)
				return
			}
			
			results <- fmt.Sprintf("Post %d: %s", post.ID, post.Title)
		}(i, url)
	}
	
	// Collect results
	for i := 0; i < len(urls); i++ {
		select {
		case result := <-results:
			fmt.Printf("Success: %s\n", result)
		case err := <-errors:
			fmt.Printf("Error: %v\n", err)
		case <-time.After(5 * time.Second):
			fmt.Println("Timeout waiting for response")
		}
	}
}