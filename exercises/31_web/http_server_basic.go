// http_server_basic.go
// Learn basic HTTP server implementation

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	fmt.Println("=== Basic HTTP Server ===")
	
	// TODO: Create HTTP server with routes
	mux := http.NewServeMux()
	
	// TODO: Add route handlers
	mux.HandleFunc("/", /* create home handler */)
	mux.HandleFunc("/hello", /* create hello handler */)
	mux.HandleFunc("/api/time", /* create time API handler */)
	mux.HandleFunc("/api/echo", /* create echo handler */)
	mux.HandleFunc("/static/", /* create static file handler */)
	
	// TODO: Create server with configuration
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
	
	// TODO: Start server
	log.Fatal(server.ListenAndServe())
}

// TODO: Home page handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type
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
	/* write HTML response */
}

// TODO: Hello handler with query parameters
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Get query parameter "name"
	name := /* get "name" query parameter */
	if name == "" {
		name = "World"
	}
	
	// Get query parameter "count"
	countStr := /* get "count" query parameter */
	count := 1
	if countStr != "" {
		if c, err := strconv.Atoi(countStr); err == nil {
			count = c
		}
	}
	
	// Set content type and write response
	/* set content type to text/plain */
	
	for i := 0; i < count; i++ {
		/* write greeting message */
	}
}

// TODO: Time API handler (JSON response)
func timeHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if /* check if method is not GET */ {
		/* write method not allowed status */
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
	/* set content type to application/json */
	
	// Encode and send JSON response
	/* encode response as JSON and write to response */
}

// TODO: Echo handler (accepts POST data)
func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if /* check if method is not POST */ {
		/* write method not allowed status */
		return
	}
	
	// Parse form data
	/* parse form data */
	
	// Create echo response
	response := map[string]interface{}{
		"method":  r.Method,
		"url":     r.URL.String(),
		"headers": r.Header,
		"form":    r.Form,
		"echo_time": time.Now().Format(time.RFC3339),
	}
	
	// Set JSON content type and send response
	/* set content type to application/json */
	/* encode and send JSON response */
}

// TODO: Static file handler
func staticHandler(w http.ResponseWriter, r *http.Request) {
	// Extract file path from URL
	// In a real application, you'd serve actual files
	// For this example, we'll return mock content
	
	filePath := /* extract path after /static/ */
	
	// Mock file content based on extension
	var content string
	var contentType string
	
	if /* check if path ends with .css */ {
		contentType = "text/css"
		content = "body { font-family: Arial, sans-serif; }"
	} else if /* check if path ends with .js */ {
		contentType = "application/javascript"
		content = "console.log('Mock JavaScript file');"
	} else if /* check if path ends with .txt */ {
		contentType = "text/plain"
		content = "This is a mock text file."
	} else {
		/* write not found status */
		return
	}
	
	// Set content type and write mock content
	/* set content type */
	/* write content */
}