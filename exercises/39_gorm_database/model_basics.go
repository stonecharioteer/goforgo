// GoForGo Exercise: GORM Model Basics
// Learn how to define GORM models with proper field tags, relationships, and database operations

package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TODO: Define a User struct with GORM model tags
// - ID field as primary key (uint, auto-increment)
// - Name field as string, not null, max 100 characters
// - Email field as string, unique index, not null
// - Age field as int with default value 0
// - CreatedAt and UpdatedAt timestamps (gorm.Model embeds these)
type User struct {
	// Your User struct here
}

// TODO: Define a Profile struct that belongs to User
// - ID field as primary key
// - UserID as foreign key referencing User.ID
// - Bio field as text
// - Website field as string
type Profile struct {
	// Your Profile struct here
}

func main() {
	// Connect to SQLite database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// TODO: Auto-migrate the User and Profile models
	// Hint: Use db.AutoMigrate()

	// TODO: Create a new user with the following data:
	// Name: "John Doe", Email: "john@example.com", Age: 30

	// TODO: Create a profile for this user:
	// Bio: "Software Developer", Website: "https://johndoe.dev"

	// TODO: Query the user by email and include their profile
	// Hint: Use db.Preload() to include associated Profile

	// TODO: Update the user's age to 31

	// TODO: Delete the profile (soft delete if using gorm.Model)

	// TODO: Print the final user data to verify operations
	fmt.Println("GORM model operations completed successfully!")
}