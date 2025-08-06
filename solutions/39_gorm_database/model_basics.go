// GoForGo Solution: GORM Model Basics
// Complete implementation of GORM models with proper field tags and database operations

package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User struct with GORM model tags
type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"not null;size:100"`
	Email     string `gorm:"uniqueIndex;not null"`
	Age       int    `gorm:"default:0"`
	CreatedAt gorm.DeletedAt
	UpdatedAt gorm.DeletedAt
	Profile   Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Profile struct that belongs to User
type Profile struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	UserID   uint   `gorm:"not null"`
	Bio      string `gorm:"type:text"`
	Website  string `gorm:"size:255"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func main() {
	// Connect to SQLite database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the User and Profile models
	err = db.AutoMigrate(&User{}, &Profile{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Create a new user
	user := User{
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   30,
	}
	
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal("Failed to create user:", result.Error)
	}
	fmt.Printf("Created user with ID: %d\n", user.ID)

	// Create a profile for this user
	profile := Profile{
		UserID:  user.ID,
		Bio:     "Software Developer",
		Website: "https://johndoe.dev",
	}
	
	result = db.Create(&profile)
	if result.Error != nil {
		log.Fatal("Failed to create profile:", result.Error)
	}
	fmt.Printf("Created profile with ID: %d\n", profile.ID)

	// Query the user by email and include their profile
	var foundUser User
	result = db.Preload("Profile").Where("email = ?", "john@example.com").First(&foundUser)
	if result.Error != nil {
		log.Fatal("Failed to find user:", result.Error)
	}
	fmt.Printf("Found user: %+v\n", foundUser)
	fmt.Printf("User profile: %+v\n", foundUser.Profile)

	// Update the user's age to 31
	result = db.Model(&foundUser).Update("age", 31)
	if result.Error != nil {
		log.Fatal("Failed to update user:", result.Error)
	}
	fmt.Printf("Updated user age to: %d\n", foundUser.Age)

	// Delete the profile (hard delete for this example)
	result = db.Unscoped().Delete(&profile)
	if result.Error != nil {
		log.Fatal("Failed to delete profile:", result.Error)
	}
	fmt.Println("Deleted profile successfully")

	// Print the final user data to verify operations
	var finalUser User
	db.First(&finalUser, foundUser.ID)
	fmt.Printf("Final user data: ID=%d, Name=%s, Email=%s, Age=%d\n", 
		finalUser.ID, finalUser.Name, finalUser.Email, finalUser.Age)
	
	fmt.Println("GORM model operations completed successfully!")
}