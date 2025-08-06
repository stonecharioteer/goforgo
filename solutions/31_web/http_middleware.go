// http_middleware.go
// Learn HTTP middleware patterns and request processing

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== HTTP Middleware Patterns ===")
	
	// Create mux and add middleware chain
	mux := http.NewServeMux()
	
	// Add routes
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/protected", protectedHandler)
	mux.HandleFunc("/api/data", apiHandler)
	
	// Create middleware chain
	handler := loggingMiddleware(mux)
	handler = authMiddleware(handler)
	handler = corsMiddleware(handler)
	
	server := &http.Server{
		Addr:    ":8081",
		Handler: handler,
	}
	
	fmt.Println("Starting server on :8081")
	fmt.Println("Try these endpoints:")
	fmt.Println("  GET  /")
	fmt.Println("  GET  /protected (needs Authorization header)")
	fmt.Println("  GET  /api/data")
	
	log.Fatal(server.ListenAndServe())
}

// Logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Log request details
		fmt.Printf("[%s] %s %s from %s\n", 
			start.Format("15:04:05"), r.Method, r.URL.Path, r.RemoteAddr)
		
		// Call next handler
		next.ServeHTTP(w, r)
		
		// Log response time
		duration := time.Since(start)
		fmt.Printf("[%s] Completed in %v\n", 
			time.Now().Format("15:04:05"), duration)
	})
}

// Authentication middleware
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip auth for public routes
		if !strings.HasPrefix(r.URL.Path, "/protected") {
			next.ServeHTTP(w, r)
			return
		}
		
		// Check Authorization header
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Authorization required")
			return
		}
		
		// Simple token validation (in real app, use proper auth)
		if auth != "Bearer secret-token" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprint(w, "Invalid authorization token")
			return
		}
		
		// Auth passed, continue
		next.ServeHTTP(w, r)
	})
}

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		// Continue to next handler
		next.ServeHTTP(w, r)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `<!DOCTYPE html>
<html>
<head><title>Middleware Demo</title></head>
<body>
	<h1>HTTP Middleware Demo</h1>
	<p>This server demonstrates middleware patterns:</p>
	<ul>
		<li><strong>Logging:</strong> All requests are logged</li>
		<li><strong>Authentication:</strong> /protected routes need auth</li>
		<li><strong>CORS:</strong> Cross-origin requests are allowed</li>
	</ul>
	<p>Try: <a href="/protected">Protected Route</a> (will fail without auth)</p>
</body>
</html>`
	fmt.Fprint(w, html)
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"message": "You accessed a protected route!", "user": "authenticated"}`)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"data": [1, 2, 3], "timestamp": "`+time.Now().Format(time.RFC3339)+`"}`)
}