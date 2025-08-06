// GoForGo Exercise: GORM Associations
// Learn how to work with GORM associations: has one, has many, belongs to, and many to many

package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TODO: Define models with different types of associations:

// User model (should have one Profile, many Posts, and many to many Tags)
type User struct {
	// Your User struct here with associations
}

// Profile model (belongs to User)
type Profile struct {
	// Your Profile struct here
}

// Post model (belongs to User, has many Comments, many to many Tags)
type Post struct {
	// Your Post struct here
}

// Comment model (belongs to Post)
type Comment struct {
	// Your Comment struct here
}

// Tag model (many to many with Posts)
type Tag struct {
	// Your Tag struct here
}

func main() {
	// Connect to SQLite database
	db, err := gorm.Open(sqlite.Open("associations.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// TODO: Auto-migrate all models

	// TODO: Create a user with profile
	// User: Name="Alice", Email="alice@example.com"
	// Profile: Bio="Content Creator", Website="https://alice.blog"

	// TODO: Create posts for the user
	// Post 1: Title="First Post", Content="Hello World"
	// Post 2: Title="Second Post", Content="Learning GORM"

	// TODO: Create tags and associate them with posts
	// Tag 1: Name="golang"
	// Tag 2: Name="tutorial"
	// Associate both tags with both posts

	// TODO: Create comments for posts
	// Comment 1 on Post 1: Author="Bob", Content="Great post!"
	// Comment 2 on Post 1: Author="Charlie", Content="Very helpful"
	// Comment 1 on Post 2: Author="Dave", Content="Keep it up!"

	// TODO: Query user with all associations preloaded
	// Use nested preloading: User -> Profile, Posts -> Comments, Posts -> Tags

	// TODO: Print the complete user data with all associations

	fmt.Println("GORM associations operations completed successfully!")
}