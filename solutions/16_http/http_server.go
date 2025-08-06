// http_server.go - SOLUTION
// Learn to build HTTP servers with routing and middleware

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

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
	{ID: 3, Name: "Charlie", Email: "charlie@example.com"},
}

var startTime = time.Now()

func main() {
	fmt.Println("=== HTTP Server with Routing ===")
	
	// Create HTTP server mux
	mux := http.NewServeMux()
	
	// Add logging middleware
	mux.Handle("/", loggingMiddleware(http.HandlerFunc(rootHandler)))
	
	// Add API routes
	mux.HandleFunc("/api/users", usersHandler)
	mux.HandleFunc("/api/users/", userHandler)
	mux.HandleFunc("/api/health", healthHandler)
	
	// Add static file serving
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	
	// Start server
	fmt.Println("Server starting on :8080")
	fmt.Println("Available endpoints:")
	fmt.Println("  GET  /api/users     - List all users")
	fmt.Println("  GET  /api/users/:id - Get user by ID")
	fmt.Println("  POST /api/users     - Create new user")
	fmt.Println("  GET  /api/health    - Health check")
	fmt.Println("  GET  /static/       - Static files")
	
	// Configure server with timeouts
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	
	// Start server with error handling
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

// Implement logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Log request details
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		
		// Call next handler
		next.ServeHTTP(w, r)
		
		// Log response time
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// Implement root handler
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Return welcome message
	response := map[string]string{
		"message": "Welcome to the Go HTTP Server API",
		"version": "1.0.0",
	}
	writeJSON(w, http.StatusOK, response)
}

// Implement users handler
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Return all users as JSON
		writeJSON(w, http.StatusOK, users)
	case http.MethodPost:
		// Create new user from JSON body
		var newUser User
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			writeError(w, http.StatusBadRequest, "Invalid JSON body")
			return
		}
		
		// Generate new ID
		newUser.ID = len(users) + 1
		users = append(users, newUser)
		
		writeJSON(w, http.StatusCreated, newUser)
	default:
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

// Implement single user handler
func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	
	// Extract user ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/users/")
	
	// Convert ID to integer
	id, err := strconv.Atoi(path)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	
	// Find user by ID
	for _, user := range users {
		if user.ID == id {
			writeJSON(w, http.StatusOK, user)
			return
		}
	}
	
	// Return not found if user doesn't exist
	writeError(w, http.StatusNotFound, "User not found")
}

// Implement health check handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	
	// Return health status
	health := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Unix(),
		"uptime":    time.Since(startTime).Seconds(),
	}
	
	writeJSON(w, http.StatusOK, health)
}

// Helper function to write JSON response
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Helper function to write error response
func writeError(w http.ResponseWriter, status int, message string) {
	errorResponse := map[string]string{"error": message}
	writeJSON(w, status, errorResponse)
}