// http_middleware.go
// Learn HTTP middleware and routing patterns in Go

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// Create a logging middleware function
// Middleware takes http.Handler and returns http.Handler
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request method and URL
		start := time.Now()
		fmt.Printf("[%s] %s %s", time.Now().Format("15:04:05"), r.Method, r.URL.Path)
		
		// Call the next handler
		next.ServeHTTP(w, r)
		
		// Log response time
		fmt.Printf(" - %v\n", time.Since(start))
	})
}

// Create authentication middleware
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}
		
		// Validate token (simple check for "Bearer valid-token")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}
		
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != "valid-token" {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		
		// Add user info to request context
		// For simplicity, we'll just proceed
		next.ServeHTTP(w, r)
	})
}

// Create CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		// Handle preflight OPTIONS requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

// Handler functions
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!\n")
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a protected resource!\n")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "API endpoint", "method": "%s"}`, r.Method)
}

// Create a router with path-based routing
func createRouter() *http.ServeMux {
	// Create new ServeMux
	mux := http.NewServeMux()
	
	// Register routes
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/api/", apiHandler)
	mux.HandleFunc("/protected", protectedHandler)
	
	return mux
}

// Chain multiple middleware functions
func chainMiddleware(handler http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	// Apply middleware in reverse order
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}
	return handler
}

func main() {
	fmt.Println("Setting up HTTP server with middleware...")
	
	// Create router
	router := createRouter()
	
	// Apply middleware to all routes
	// Chain: CORS -> Logging -> (Auth for protected routes only)
	
	// For public routes (all except /protected)
	publicHandler := chainMiddleware(router,
		corsMiddleware,
		loggingMiddleware,
	)
	
	// Create a separate handler for protected routes
	protectedMux := http.NewServeMux()
	protectedMux.HandleFunc("/protected", protectedHandler)
	
	protectedHandler := chainMiddleware(protectedMux,
		corsMiddleware,
		loggingMiddleware,
		authMiddleware,
	)
	
	// Create main handler that routes to appropriate handler
	mainHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/protected") {
			protectedHandler.ServeHTTP(w, r)
			return
		}
		publicHandler.ServeHTTP(w, r)
	})
	
	fmt.Println("Server starting on :8080")
	fmt.Println("Routes:")
	fmt.Println("  GET  /           - Home page (public)")
	fmt.Println("  GET  /api/       - API endpoint (public)")
	fmt.Println("  GET  /protected  - Protected resource (requires auth)")
	fmt.Println("")
	fmt.Println("To test protected endpoint:")
	fmt.Println("  curl -H 'Authorization: Bearer valid-token' http://localhost:8080/protected")
	fmt.Println("")
	fmt.Println("Press Ctrl+C to stop the server")
	
	// Start server with timeout configurations
	server := &http.Server{
		Addr:           ":8080",
		Handler:        mainHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}
	
	log.Fatal(server.ListenAndServe())
}