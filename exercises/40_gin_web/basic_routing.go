// GoForGo Exercise: Gin Basic Routing
// Learn how to create HTTP routes with Gin web framework including path parameters and query strings

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: Define a User struct for JSON responses
// Fields: ID (int), Name (string), Email (string)
type User struct {
	// Your User struct here
}

// TODO: Create a global users slice to simulate a database
var users []User

func main() {
	// TODO: Create a Gin router with default middleware
	// Hint: Use gin.Default()

	// TODO: Define a GET route for "/" that returns a welcome message
	// Response: {"message": "Welcome to Gin API!"}

	// TODO: Define a GET route for "/users" that returns all users
	// Response: JSON array of users

	// TODO: Define a GET route for "/users/:id" that returns a specific user by ID
	// - Extract the ID from the path parameter
	// - Convert string ID to integer
	// - Find user with matching ID
	// - Return user if found, or 404 error if not found

	// TODO: Define a GET route for "/search" that searches users by name query parameter
	// - Extract the "name" query parameter
	// - Filter users whose name contains the search term (case-insensitive)
	// - Return filtered users array

	// TODO: Initialize sample data
	// Add 3 users: 
	// - ID: 1, Name: "Alice Johnson", Email: "alice@example.com"
	// - ID: 2, Name: "Bob Smith", Email: "bob@example.com"  
	// - ID: 3, Name: "Charlie Brown", Email: "charlie@example.com"

	// TODO: Start the server on port 8080
	// Hint: Use router.Run(":8080")
}