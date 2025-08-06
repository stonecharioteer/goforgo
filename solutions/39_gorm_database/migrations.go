// GoForGo Solution: GORM Migrations
// Complete implementation of database schema migrations and data transformations

package main

import (
	"fmt"
	"log"
	"strings"
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

// Updated User model with additional fields (version 2)
type UserV2 struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"not null;size:100"` // Increased size
	Email     string    `gorm:"uniqueIndex;not null"`
	Age       int       `gorm:"default:0"`
	IsActive  bool      `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Product model for table creation migration
type Product struct {
	ID       uint    `gorm:"primaryKey;autoIncrement"`
	Name     string  `gorm:"not null;size:200"`
	Price    float64 `gorm:"type:decimal(10,2);not null"`
	Category string  `gorm:"size:100;index"` // Added index
	InStock  bool    `gorm:"default:true"`
}

func main() {
	// Connect to SQLite database
	db, err := gorm.Open(sqlite.Open("migrations.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("=== Step 1: Initial Migration ===")
	// Create initial User table
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to migrate User:", err)
	}
	fmt.Println("Created initial User table")

	// Insert sample users
	users := []User{
		{Name: "John", Email: "john@example.com", CreatedAt: time.Now()},
		{Name: "Jane", Email: "jane@example.com", CreatedAt: time.Now()},
	}
	
	result := db.Create(&users)
	if result.Error != nil {
		log.Fatal("Failed to create users:", result.Error)
	}
	fmt.Printf("Inserted %d users\n", len(users))

	fmt.Println("=== Step 2: Schema Migration ===")
	// Migrate to UserV2 schema (add new columns)
	err = db.AutoMigrate(&UserV2{})
	if err != nil {
		log.Fatal("Failed to migrate UserV2:", err)
	}
	fmt.Println("Migrated User table to V2 schema")

	// Update existing users with default values
	result = db.Model(&UserV2{}).Where("age = ?", 0).Updates(UserV2{
		Age:       25,
		IsActive:  true,
		UpdatedAt: time.Now(),
	})
	if result.Error != nil {
		log.Fatal("Failed to update users:", result.Error)
	}
	fmt.Printf("Updated %d users with default values\n", result.RowsAffected)

	fmt.Println("=== Step 3: Add New Table ===")
	// Create Product table using AutoMigrate
	err = db.AutoMigrate(&Product{})
	if err != nil {
		log.Fatal("Failed to migrate Product:", err)
	}
	fmt.Println("Created Product table")

	// Insert sample products
	products := []Product{
		{Name: "Laptop", Price: 999.99, Category: "Electronics", InStock: true},
		{Name: "Book", Price: 29.99, Category: "Education", InStock: false},
	}
	
	result = db.Create(&products)
	if result.Error != nil {
		log.Fatal("Failed to create products:", result.Error)
	}
	fmt.Printf("Inserted %d products\n", len(products))

	fmt.Println("=== Step 4: Column Modifications ===")
	// Change User.Name column size from 50 to 100 (already done in UserV2 struct)
	// SQLite doesn't support column modification directly, so we use the struct definition
	fmt.Println("Column modifications applied via struct definition")

	// Add index to Product.Category (already done in struct definition)
	if db.Migrator().HasIndex(&Product{}, "category") {
		fmt.Println("Category index exists")
	} else {
		err = db.Migrator().CreateIndex(&Product{}, "category")
		if err != nil {
			log.Printf("Note: Index creation might not be needed: %v", err)
		}
	}

	fmt.Println("=== Step 5: Data Migration ===")
	// Update all users with name starting with 'J' to set Age=30
	result = db.Model(&UserV2{}).Where("name LIKE ?", "J%").Updates(UserV2{
		Age:       30,
		UpdatedAt: time.Now(),
	})
	if result.Error != nil {
		log.Fatal("Failed to update users:", result.Error)
	}
	fmt.Printf("Updated %d users starting with 'J' to age 30\n", result.RowsAffected)

	fmt.Println("=== Step 6: Verify Migration ===")
	// Query and display all users with their updated schema
	var allUsers []UserV2
	result = db.Find(&allUsers)
	if result.Error != nil {
		log.Fatal("Failed to query users:", result.Error)
	}
	
	fmt.Println("\n--- Users Table ---")
	for _, user := range allUsers {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d, Active: %t\n",
			user.ID, user.Name, user.Email, user.Age, user.IsActive)
	}

	// Query and display all products
	var allProducts []Product
	result = db.Find(&allProducts)
	if result.Error != nil {
		log.Fatal("Failed to query products:", result.Error)
	}
	
	fmt.Println("\n--- Products Table ---")
	for _, product := range allProducts {
		fmt.Printf("ID: %d, Name: %s, Price: %.2f, Category: %s, InStock: %t\n",
			product.ID, product.Name, product.Price, product.Category, product.InStock)
	}

	// Show table schema information
	tables, err := db.Migrator().GetTables()
	if err != nil {
		log.Printf("Failed to get tables: %v", err)
	} else {
		fmt.Printf("\n--- Database Tables ---\n")
		fmt.Printf("Tables: %s\n", strings.Join(tables, ", "))
	}

	fmt.Println("\nGORM migrations operations completed successfully!")
}