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
	
	// TODO: Create mux and add middleware chain
	mux := http.NewServeMux()
	
	// Add routes
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/protected", protectedHandler)
	mux.HandleFunc("/api/data", apiHandler)
	
	// TODO: Create middleware chain
	handler := /* apply logging middleware */
	handler = /* apply authentication middleware */
	handler = /* apply CORS middleware */
	
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

// TODO: Logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// TODO: Log request details
		/* log method, URL, and remote address */
		
		// Call next handler
		/* call next handler */
		
		// TODO: Log response time
		duration := time.Since(start)
		/* log completion time */
	})
}

// TODO: Authentication middleware
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip auth for public routes
		if !strings.HasPrefix(r.URL.Path, "/protected") {
			/* call next handler for non-protected routes */
			return
		}
		
		// TODO: Check Authorization header
		auth := /* get Authorization header */
		if /* check if auth header is empty */ {
			/* write unauthorized status */
			/* write error message */
			return
		}
		
		// TODO: Simple token validation (in real app, use proper auth)
		if /* check if auth doesn't equal "Bearer secret-token" */ {
			/* write forbidden status */
			/* write error message */
			return
		}
		
		// Auth passed, continue
		/* call next handler */
	})
}

// TODO: CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Set CORS headers
		/* set Access-Control-Allow-Origin header to "*" */
		/* set Access-Control-Allow-Methods header to "GET, POST, PUT, DELETE, OPTIONS" */
		/* set Access-Control-Allow-Headers header to "Content-Type, Authorization" */
		
		// Handle preflight requests
		if /* check if method is OPTIONS */ {
			/* write OK status */
			return
		}
		
		// Continue to next handler
		/* call next handler */
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