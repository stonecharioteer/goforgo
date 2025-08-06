// GoForGo Solution: Gin Middleware Chain
// Complete implementation of custom middleware for logging, authentication, and request processing

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Custom logging middleware function
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Record start time
		startTime := time.Now()
		
		// Process request
		c.Next()
		
		// Calculate response time
		duration := time.Since(startTime)
		
		// Log request details
		log.Printf("%s %s - %d - %v",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
		)
	}
}

// Authentication middleware function
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for Authorization header
		authHeader := c.GetHeader("Authorization")
		
		// Validate token
		if authHeader == "" || authHeader != "Bearer valid-token" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized: Invalid or missing token",
			})
			c.Abort() // Stop processing
			return
		}
		
		// Token is valid, continue to next handler
		c.Next()
	}
}

// Simple rate limiting middleware
var requestCounts = make(map[string]int)

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		
		// Check current request count for this IP
		if requestCounts[clientIP] >= 5 {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests. Rate limit exceeded.",
			})
			c.Abort()
			return
		}
		
		// Increment request count
		requestCounts[clientIP]++
		
		// Continue to next handler
		c.Next()
	}
}

// CORS middleware function
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add CORS headers
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Header("Access-Control-Max-Age", "86400")
		
		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		
		// Continue to next handler
		c.Next()
	}
}

func main() {
	// Create Gin router without default middleware
	router := gin.New()

	// Add global middleware to the router
	router.Use(LoggingMiddleware())
	router.Use(CORSMiddleware())
	router.Use(gin.Recovery()) // Add recovery middleware

	// Create a public route group for non-authenticated endpoints
	publicGroup := router.Group("/")
	publicGroup.Use(RateLimitMiddleware())
	{
		// Add a public route
		publicGroup.GET("/public", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Public endpoint",
			})
		})
	}

	// Create a protected route group for authenticated endpoints
	apiGroup := router.Group("/api")
	apiGroup.Use(AuthMiddleware())
	{
		// Add protected routes
		apiGroup.GET("/profile", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"user": "authenticated_user",
				"role": "admin",
			})
		})
		
		apiGroup.GET("/data", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"data":      []string{"item1", "item2", "item3"},
				"timestamp": time.Now().Format(time.RFC3339),
			})
		})
	}

	// Add error handling route for 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Route %s not found", c.Request.URL.Path),
		})
	})

	log.Println("Server starting on :8080")
	// Start the server on port 8080
	router.Run(":8080")
}