// http_server_basic.go - SOLUTION
// Learn basic HTTP server implementation

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== Basic HTTP Server ===")
	
	// Create HTTP server with routes
	mux := http.NewServeMux()
	
	// Add route handlers
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/api/time", timeHandler)
	mux.HandleFunc("/api/echo", echoHandler)
	mux.HandleFunc("/static/", staticHandler)
	
	// Create server with configuration
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	
	fmt.Println("Starting server on :8080")
	fmt.Println("Routes available:")
	fmt.Println("  GET  /")
	fmt.Println("  GET  /hello")
	fmt.Println("  GET  /api/time")
	fmt.Println("  POST /api/echo")
	fmt.Println("  GET  /static/...")
	
	// Start server
	log.Fatal(server.ListenAndServe())
}

// Home page handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type
	w.Header().Set("Content-Type", "text/html")
	
	// Write HTML response
	html := `<!DOCTYPE html>
<html>
<head><title>Go HTTP Server</title></head>
<body>
	<h1>Welcome to Go HTTP Server!</h1>
	<p>Available endpoints:</p>
	<ul>
		<li><a href="/hello">Hello</a></li>
		<li><a href="/api/time">Current Time API</a></li>
	</ul>
</body>
</html>`
	w.Write([]byte(html))
}

// Hello handler with query parameters
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Get query parameter "name"
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	
	// Get query parameter "count"
	countStr := r.URL.Query().Get("count")
	count := 1
	if countStr != "" {
		if c, err := strconv.Atoi(countStr); err == nil {
			count = c
		}
	}
	
	// Set content type and write response
	w.Header().Set("Content-Type", "text/plain")
	
	for i := 0; i < count; i++ {
		fmt.Fprintf(w, "Hello, %s! (greeting %d)\n", name, i+1)
	}
}

// Time API handler (JSON response)
func timeHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Create response data
	response := map[string]interface{}{
		"timestamp": time.Now().Unix(),
		"datetime":  time.Now().Format(time.RFC3339),
		"timezone":  "UTC",
		"format":    "RFC3339",
	}
	
	// Set JSON content type
	w.Header().Set("Content-Type", "application/json")
	
	// Encode and send JSON response
	json.NewEncoder(w).Encode(response)
}

// Echo handler (accepts POST data)
func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Parse form data
	r.ParseForm()
	
	// Create echo response
	response := map[string]interface{}{
		"method":  r.Method,
		"url":     r.URL.String(),
		"headers": r.Header,
		"form":    r.Form,
		"echo_time": time.Now().Format(time.RFC3339),
	}
	
	// Set JSON content type and send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Static file handler
func staticHandler(w http.ResponseWriter, r *http.Request) {
	// Extract file path from URL
	// In a real application, you'd serve actual files
	// For this example, we'll return mock content
	
	filePath := strings.TrimPrefix(r.URL.Path, "/static/")
	
	// Mock file content based on extension
	var content string
	var contentType string
	
	if strings.HasSuffix(filePath, ".css") {
		contentType = "text/css"
		content = "body { font-family: Arial, sans-serif; }"
	} else if strings.HasSuffix(filePath, ".js") {
		contentType = "application/javascript"
		content = "console.log('Mock JavaScript file');"
	} else if strings.HasSuffix(filePath, ".txt") {
		contentType = "text/plain"
		content = "This is a mock text file."
	} else {
		http.NotFound(w, r)
		return
	}
	
	// Set content type and write mock content
	w.Header().Set("Content-Type", contentType)
	w.Write([]byte(content))
}