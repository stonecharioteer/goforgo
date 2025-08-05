// http_server.go
// Learn to build HTTP servers with routing and middleware

package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	fmt.Println("=== HTTP Server with Routing ===")
	
	// TODO: Create HTTP server mux
	mux := /* create new HTTP mux */
	
	// TODO: Add logging middleware
	mux.Handle("/", /* wrap with logging middleware */(/* create handler for root */))
	
	// TODO: Add API routes
	mux.HandleFunc("/api/users", /* create users handler */)
	mux.HandleFunc("/api/users/", /* create single user handler */)
	mux.HandleFunc("/api/health", /* create health check handler */)
	
	// TODO: Add static file serving
	fileServer := /* create file server for ./static */
	mux.Handle("/static/", /* strip prefix and serve files */)
	
	// TODO: Start server
	fmt.Println("Server starting on :8080")
	fmt.Println("Available endpoints:")
	fmt.Println("  GET  /api/users     - List all users")
	fmt.Println("  GET  /api/users/:id - Get user by ID")
	fmt.Println("  POST /api/users     - Create new user")
	fmt.Println("  GET  /api/health    - Health check")
	fmt.Println("  GET  /static/       - Static files")
	
	// TODO: Configure server with timeouts
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  /* set read timeout */,
		WriteTimeout: /* set write timeout */,
		IdleTimeout:  /* set idle timeout */,
	}
	
	// TODO: Start server with error handling
	err := /* start server */
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

// TODO: Implement logging middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// TODO: Log request details
		
		// TODO: Call next handler
		
		// TODO: Log response time
	})
}

// TODO: Implement root handler
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Return welcome message
}

// TODO: Implement users handler
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// TODO: Return all users as JSON
	case http.MethodPost:
		// TODO: Create new user from JSON body
	default:
		/* return method not allowed */
	}
}

// TODO: Implement single user handler
func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		/* return method not allowed */
		return
	}
	
	// TODO: Extract user ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/api/users/")
	
	// TODO: Convert ID to integer
	id, err := /* convert path to integer */
	if err != nil {
		/* return bad request */
		return
	}
	
	// TODO: Find user by ID
	for _, user := range users {
		if /* user ID matches */ {
			/* return user as JSON */
			return
		}
	}
	
	// TODO: Return not found if user doesn't exist
}

// TODO: Implement health check handler
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		/* return method not allowed */
		return
	}
	
	// TODO: Return health status
	health := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Unix(),
		"uptime":    /* calculate uptime */,
	}
	
	/* return health as JSON */
}

// TODO: Helper function to write JSON response
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	/* set content type header */
	/* set status code */
	/* encode data as JSON */
}

// TODO: Helper function to write error response
func writeError(w http.ResponseWriter, status int, message string) {
	/* create error response */
	/* write as JSON */
}