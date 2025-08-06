// GoForGo Solution: Gin Basic Routing
// Complete implementation of HTTP routes with Gin including path parameters and query strings

package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// User struct for JSON responses
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Global users slice to simulate a database
var users []User

func main() {
	// Create a Gin router with default middleware
	router := gin.Default()

	// Define a GET route for "/" that returns a welcome message
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Gin API!",
		})
	})

	// Define a GET route for "/users" that returns all users
	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})

	// Define a GET route for "/users/:id" that returns a specific user by ID
	router.GET("/users/:id", func(c *gin.Context) {
		// Extract the ID from the path parameter
		idParam := c.Param("id")
		
		// Convert string ID to integer
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid user ID",
			})
			return
		}

		// Find user with matching ID
		for _, user := range users {
			if user.ID == id {
				c.JSON(http.StatusOK, user)
				return
			}
		}

		// Return 404 error if user not found
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
	})

	// Define a GET route for "/search" that searches users by name query parameter
	router.GET("/search", func(c *gin.Context) {
		// Extract the "name" query parameter
		nameQuery := c.Query("name")
		
		if nameQuery == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Missing 'name' query parameter",
			})
			return
		}

		// Filter users whose name contains the search term (case-insensitive)
		var filteredUsers []User
		nameQuery = strings.ToLower(nameQuery)
		
		for _, user := range users {
			if strings.Contains(strings.ToLower(user.Name), nameQuery) {
				filteredUsers = append(filteredUsers, user)
			}
		}

		// Return filtered users array
		c.JSON(http.StatusOK, filteredUsers)
	})

	// Initialize sample data
	users = []User{
		{ID: 1, Name: "Alice Johnson", Email: "alice@example.com"},
		{ID: 2, Name: "Bob Smith", Email: "bob@example.com"},
		{ID: 3, Name: "Charlie Brown", Email: "charlie@example.com"},
	}

	// Start the server on port 8080
	router.Run(":8080")
}