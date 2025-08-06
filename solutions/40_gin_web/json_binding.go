// GoForGo Solution: Gin JSON Binding
// Complete implementation of JSON binding with validation and different binding scenarios

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateUserRequest for POST /users
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"required,min=18,max=120"`
	Password string `json:"password" binding:"required,min=8"`
}

// UpdateUserRequest for PUT /users/:id (partial updates allowed)
type UpdateUserRequest struct {
	Name  *string `json:"name,omitempty" binding:"omitempty,min=2,max=50"`
	Email *string `json:"email,omitempty" binding:"omitempty,email"`
	Age   *int    `json:"age,omitempty" binding:"omitempty,min=18,max=120"`
}

// UserResponse for API responses  
type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// LoginRequest for login endpoint
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Global users slice and counter for IDs
var users []UserResponse
var nextID int = 1

func main() {
	router := gin.Default()

	// POST /users - Create a new user
	router.POST("/users", func(c *gin.Context) {
		var req CreateUserRequest
		
		// Bind JSON request body to struct
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid request data",
				"details": err.Error(),
			})
			return
		}

		// Create new user
		newUser := UserResponse{
			ID:    nextID,
			Name:  req.Name,
			Email: req.Email,
			Age:   req.Age,
		}
		nextID++
		
		users = append(users, newUser)
		
		// Return 201 with user data
		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
			"user":    newUser,
		})
	})

	// PUT /users/:id - Update existing user
	router.PUT("/users/:id", func(c *gin.Context) {
		// Extract user ID from path parameter
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid user ID",
			})
			return
		}

		// Find existing user by ID
		var userIndex = -1
		for i, user := range users {
			if user.ID == id {
				userIndex = i
				break
			}
		}
		
		if userIndex == -1 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
			return
		}

		// Bind JSON request body to struct
		var req UpdateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid request data",
				"details": err.Error(),
			})
			return
		}

		// Update only the fields that were provided
		if req.Name != nil {
			users[userIndex].Name = *req.Name
		}
		if req.Email != nil {
			users[userIndex].Email = *req.Email
		}
		if req.Age != nil {
			users[userIndex].Age = *req.Age
		}

		// Return updated user data
		c.JSON(http.StatusOK, gin.H{
			"message": "User updated successfully",
			"user":    users[userIndex],
		})
	})

	// POST /users/batch - Create multiple users
	router.POST("/users/batch", func(c *gin.Context) {
		var req []CreateUserRequest
		
		// Bind JSON array to slice
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid request data",
				"details": err.Error(),
			})
			return
		}

		// Create all users
		var createdUsers []UserResponse
		for _, userReq := range req {
			newUser := UserResponse{
				ID:    nextID,
				Name:  userReq.Name,
				Email: userReq.Email,
				Age:   userReq.Age,
			}
			nextID++
			
			users = append(users, newUser)
			createdUsers = append(createdUsers, newUser)
		}

		// Return array of created users
		c.JSON(http.StatusCreated, gin.H{
			"message": "Users created successfully",
			"users":   createdUsers,
			"count":   len(createdUsers),
		})
	})

	// GET /users - Return all users
	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"users": users,
			"count": len(users),
		})
	})

	// POST /login - Simulate login with JSON binding
	router.POST("/login", func(c *gin.Context) {
		var req LoginRequest
		
		// Bind JSON and validate required fields
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid login data",
				"details": err.Error(),
			})
			return
		}

		// Return success message (simplified authentication)
		c.JSON(http.StatusOK, gin.H{
			"message":  "Login successful",
			"username": req.Username,
		})
	})

	router.Run(":8080")
}