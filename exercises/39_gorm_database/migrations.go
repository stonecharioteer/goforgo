// GoForGo Exercise: GORM Migrations
// Learn how to handle database schema migrations, column modifications, and data transformations

package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Initial User model (version 1)
type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"not null;size:50"`
	Email     string `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
}

// TODO: Define UserV2 model with additional fields for migration
// Add: Age (int), IsActive (bool, default true), UpdatedAt (time.Time)
type UserV2 struct {
	// Your updated User struct here
}

// TODO: Define a Product model to demonstrate table creation migration
// Fields: ID, Name, Price (decimal), Category, InStock (bool)
type Product struct {
	// Your Product struct here
}

func main() {
	// Connect to SQLite database
	db, err := gorm.Open(sqlite.Open("migrations.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("=== Step 1: Initial Migration ===")
	// TODO: Create initial User table
	// Use db.AutoMigrate()

	// TODO: Insert sample users
	// User 1: Name="John", Email="john@example.com"
	// User 2: Name="Jane", Email="jane@example.com"

	fmt.Println("=== Step 2: Schema Migration ===")
	// TODO: Migrate to UserV2 schema (add new columns)
	// This will add Age, IsActive, and UpdatedAt columns

	// TODO: Update existing users with default values
	// Set Age=25, IsActive=true for all existing users

	fmt.Println("=== Step 3: Add New Table ===")
	// TODO: Create Product table using AutoMigrate

	// TODO: Insert sample products
	// Product 1: Name="Laptop", Price=999.99, Category="Electronics", InStock=true
	// Product 2: Name="Book", Price=29.99, Category="Education", InStock=false

	fmt.Println("=== Step 4: Column Modifications ===")
	// TODO: Use Migrator to modify columns
	// - Change User.Name column size from 50 to 100
	// - Add index to Product.Category

	fmt.Println("=== Step 5: Data Migration ===")
	// TODO: Perform data transformation
	// - Update all users with name starting with 'J' to set Age=30

	fmt.Println("=== Step 6: Verify Migration ===")
	// TODO: Query and display all users with their updated schema
	// TODO: Query and display all products
	// TODO: Show table schema information using db.Migrator().GetTables()

	fmt.Println("GORM migrations operations completed successfully!")
}