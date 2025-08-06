// middleware_usage.go
// Learn HTTP middleware with gorilla/mux for cross-cutting concerns
//
// Middleware in gorilla/mux allows you to process requests before they reach
// their final handler. Common uses include logging, authentication, CORS,
// rate limiting, and request validation.
//
// I AM NOT DONE YET!

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Apply middleware to all routes
	// TODO: Use router.Use() to apply the loggingMiddleware to all routes
	// router.Use(???)

	// Routes
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/protected", protectedHandler)

	// Subrouter with additional middleware
	// TODO: Create a subrouter for "/api" path prefix
	// apiRouter := router.PathPrefix(???).Subrouter()
	
	// TODO: Apply authMiddleware to the API subrouter
	// apiRouter.Use(???)
	
	apiRouter.HandleFunc("/users", getUsersHandler)
	apiRouter.HandleFunc("/orders", getOrdersHandler)

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Logging middleware - logs all incoming requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// TODO: Log the request method, URL path, and timestamp
		// Format: "[timestamp] method path"
		// log.Printf(???, ???, ???, ???)
		
		// Call the next handler
		next.ServeHTTP(w, r)
		
		// Log the duration
		duration := time.Since(start)
		log.Printf("Request completed in %v", duration)
	})
}

// Authentication middleware - checks for auth header
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Get the "Authorization" header from the request
		// authHeader := r.Header.Get(???)
		
		// TODO: Check if the auth header is missing or doesn't equal "Bearer secret-token"
		// If invalid, write a 401 status and return
		// if authHeader != ??? {
		//     http.Error(w, ???, http.StatusUnauthorized)
		//     return
		// }
		
		// If valid, continue to the next handler
		next.ServeHTTP(w, r)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome! This request was logged.\n")
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a protected endpoint with logging.\n")
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Write a response indicating successful API access
	// Example: "API Users: [user1, user2, user3]"
	fmt.Fprintf(w, ???)
}

func getOrdersHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Write a response indicating successful API access
	// Example: "API Orders: [order1, order2, order3]"
	fmt.Fprintf(w, ???)
}