// GoForGo Solution: GORM Associations
// Complete implementation of GORM associations: has one, has many, belongs to, and many to many

package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User model (has one Profile, has many Posts)
type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"not null;size:100"`
	Email    string `gorm:"uniqueIndex;not null"`
	Profile  Profile
	Posts    []Post
}

// Profile model (belongs to User)
type Profile struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	UserID   uint   `gorm:"not null"`
	Bio      string `gorm:"type:text"`
	Website  string `gorm:"size:255"`
}

// Post model (belongs to User, has many Comments, many to many Tags)
type Post struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	UserID   uint   `gorm:"not null"`
	Title    string `gorm:"not null;size:200"`
	Content  string `gorm:"type:text"`
	Comments []Comment
	Tags     []Tag `gorm:"many2many:post_tags;"`
}

// Comment model (belongs to Post)
type Comment struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	PostID  uint   `gorm:"not null"`
	Author  string `gorm:"not null;size:100"`
	Content string `gorm:"type:text;not null"`
}

// Tag model (many to many with Posts)
type Tag struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"uniqueIndex;not null;size:50"`
	Posts []Post `gorm:"many2many:post_tags;"`
}

func main() {
	// Connect to SQLite database
	db, err := gorm.Open(sqlite.Open("associations.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate all models
	err = db.AutoMigrate(&User{}, &Profile{}, &Post{}, &Comment{}, &Tag{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Create a user with profile
	user := User{
		Name:  "Alice",
		Email: "alice@example.com",
		Profile: Profile{
			Bio:     "Content Creator",
			Website: "https://alice.blog",
		},
	}
	
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal("Failed to create user:", result.Error)
	}
	fmt.Printf("Created user with ID: %d\n", user.ID)

	// Create posts for the user
	posts := []Post{
		{
			UserID:  user.ID,
			Title:   "First Post",
			Content: "Hello World",
		},
		{
			UserID:  user.ID,
			Title:   "Second Post",
			Content: "Learning GORM",
		},
	}
	
	result = db.Create(&posts)
	if result.Error != nil {
		log.Fatal("Failed to create posts:", result.Error)
	}
	fmt.Printf("Created %d posts\n", len(posts))

	// Create tags and associate them with posts
	tags := []Tag{
		{Name: "golang"},
		{Name: "tutorial"},
	}
	
	result = db.Create(&tags)
	if result.Error != nil {
		log.Fatal("Failed to create tags:", result.Error)
	}
	
	// Associate tags with posts
	for i := range posts {
		result = db.Model(&posts[i]).Association("Tags").Append(&tags)
		if result != nil {
			log.Fatal("Failed to associate tags:", result)
		}
	}
	fmt.Printf("Associated tags with posts\n")

	// Create comments for posts
	comments := []Comment{
		{PostID: posts[0].ID, Author: "Bob", Content: "Great post!"},
		{PostID: posts[0].ID, Author: "Charlie", Content: "Very helpful"},
		{PostID: posts[1].ID, Author: "Dave", Content: "Keep it up!"},
	}
	
	result = db.Create(&comments)
	if result.Error != nil {
		log.Fatal("Failed to create comments:", result.Error)
	}
	fmt.Printf("Created %d comments\n", len(comments))

	// Query user with all associations preloaded
	var foundUser User
	result = db.Preload("Profile").
		Preload("Posts.Comments").
		Preload("Posts.Tags").
		First(&foundUser, user.ID)
	if result.Error != nil {
		log.Fatal("Failed to find user:", result.Error)
	}

	// Print the complete user data with all associations
	fmt.Printf("\n=== User Data ===\n")
	fmt.Printf("User: %s (%s)\n", foundUser.Name, foundUser.Email)
	fmt.Printf("Profile: %s - %s\n", foundUser.Profile.Bio, foundUser.Profile.Website)
	
	fmt.Printf("\n=== Posts ===\n")
	for _, post := range foundUser.Posts {
		fmt.Printf("Post: %s\n", post.Title)
		fmt.Printf("  Content: %s\n", post.Content)
		
		fmt.Printf("  Tags: ")
		for i, tag := range post.Tags {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(tag.Name)
		}
		fmt.Println()
		
		fmt.Printf("  Comments:\n")
		for _, comment := range post.Comments {
			fmt.Printf("    %s: %s\n", comment.Author, comment.Content)
		}
		fmt.Println()
	}

	fmt.Println("GORM associations operations completed successfully!")
}