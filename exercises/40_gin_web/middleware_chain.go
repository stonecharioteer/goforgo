// GoForGo Exercise: Gin Middleware Chain
// Learn how to create and use custom middleware for logging, authentication, and request processing

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// TODO: Create a custom logging middleware function
// - Log the HTTP method, path, and response time
// - Use c.Next() to call the next handler
// - Calculate time before and after the request
func LoggingMiddleware() gin.HandlerFunc {
	// Your middleware implementation here
	return nil
}

// TODO: Create an authentication middleware function
// - Check for "Authorization" header
// - If header is missing or doesn't equal "Bearer valid-token", return 401
// - If valid, call c.Next() to continue to the next handler
func AuthMiddleware() gin.HandlerFunc {
	// Your middleware implementation here
	return nil
}

// TODO: Create a rate limiting middleware function (simple implementation)
// - Track request count per IP address using a map
// - Allow maximum 5 requests per IP
// - Return 429 (Too Many Requests) if limit exceeded
// Note: This is a simplified version for learning purposes
var requestCounts = make(map[string]int)

func RateLimitMiddleware() gin.HandlerFunc {
	// Your middleware implementation here
	return nil
}

// TODO: Create a CORS middleware function
// - Add CORS headers to allow cross-origin requests
// - Set Access-Control-Allow-Origin, Access-Control-Allow-Methods, etc.
func CORSMiddleware() gin.HandlerFunc {
	// Your middleware implementation here
	return nil
}

func main() {
	// Create Gin router without default middleware
	router := gin.New()

	// TODO: Add global middleware to the router
	// - Add logging middleware
	// - Add CORS middleware

	// TODO: Create a public route group for non-authenticated endpoints
	// - Mount on "/"
	// - Add rate limiting middleware to this group

	// TODO: Add a public route GET "/public" that returns {"message": "Public endpoint"}

	// TODO: Create a protected route group for authenticated endpoints
	// - Mount on "/api"
	// - Add authentication middleware to this group

	// TODO: Add protected routes:
	// - GET "/api/profile" returns {"user": "authenticated_user", "role": "admin"}
	// - GET "/api/data" returns {"data": ["item1", "item2", "item3"], "timestamp": current_time}

	// TODO: Add error handling route for 404
	// Use router.NoRoute() to handle undefined routes

	log.Println("Server starting on :8080")
	// TODO: Start the server on port 8080
}