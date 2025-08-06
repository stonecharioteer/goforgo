// middleware_usage.go - Solution
// Learn HTTP middleware with gorilla/mux for cross-cutting concerns

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
	router.Use(loggingMiddleware)

	// Routes
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/protected", protectedHandler)

	// Subrouter with additional middleware
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(authMiddleware)
	
	apiRouter.HandleFunc("/users", getUsersHandler)
	apiRouter.HandleFunc("/orders", getOrdersHandler)

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Logging middleware - logs all incoming requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%v] %s %s", start.Format("15:04:05"), r.Method, r.URL.Path)
		
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
		authHeader := r.Header.Get("Authorization")
		
		if authHeader != "Bearer secret-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		
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
	fmt.Fprintf(w, "API Users: [user1, user2, user3]\n")
}

func getOrdersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API Orders: [order1, order2, order3]\n")
}