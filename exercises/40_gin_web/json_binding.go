// GoForGo Exercise: Gin JSON Binding
// Learn how to bind JSON request data to structs, validate input, and handle different binding scenarios

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: Define request/response structs with validation tags

// CreateUserRequest for POST /users
type CreateUserRequest struct {
	// Your struct fields with validation tags:
	// - Name: required, min 2 characters, max 50 characters
	// - Email: required, valid email format  
	// - Age: required, minimum 18, maximum 120
	// - Password: required, minimum 8 characters
}

// UpdateUserRequest for PUT /users/:id (partial updates allowed)
type UpdateUserRequest struct {
	// Your struct fields (all optional for partial updates):
	// - Name: optional, min 2 characters if provided
	// - Email: optional, valid email format if provided
	// - Age: optional, minimum 18 if provided
}

// UserResponse for API responses
type UserResponse struct {
	// Your response struct:
	// - ID, Name, Email, Age (no password in response)
}

// TODO: Define a global users slice and counter for IDs
var users []UserResponse
var nextID int = 1

func main() {
	router := gin.Default()

	// TODO: POST /users - Create a new user
	// - Bind JSON request body to CreateUserRequest struct
	// - Validate the input using ShouldBindJSON
	// - If validation fails, return 400 with error details
	// - If successful, create new user and return 201 with user data

	// TODO: PUT /users/:id - Update existing user
	// - Extract user ID from path parameter
	// - Bind JSON request body to UpdateUserRequest struct  
	// - Find existing user by ID
	// - Update only the fields that were provided in the request
	// - Return updated user data

	// TODO: POST /users/batch - Create multiple users
	// - Bind JSON array to []CreateUserRequest
	// - Validate each user in the array
	// - If any validation fails, return 400 with details
	// - Create all users and return array of created users

	// TODO: GET /users - Return all users

	// TODO: POST /login - Simulate login with JSON binding
	// - Create LoginRequest struct with Username and Password fields
	// - Bind JSON and validate required fields
	// - Return success message (don't implement real authentication)

	router.Run(":8080")
}