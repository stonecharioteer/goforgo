// sql_basics.go
// Learn database operations with Go's database/sql package

package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

// User struct for database operations
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	IsActive  bool      `json:"is_active"`
}

func main() {
	fmt.Println("=== SQL Database Basics ===")
	
	// Create database connection
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()
	
	// Initialize database schema
	if err := initializeSchema(db); err != nil {
		log.Fatal("Failed to initialize schema:", err)
	}
	
	// Demonstrate CRUD operations
	fmt.Println("1. Creating users...")
	if err := createUsers(db); err != nil {
		log.Printf("Error creating users: %v", err)
	}
	
	fmt.Println("\n2. Reading users...")
	if err := readUsers(db); err != nil {
		log.Printf("Error reading users: %v", err)
	}
	
	fmt.Println("\n3. Updating users...")
	if err := updateUser(db, 1); err != nil {
		log.Printf("Error updating user: %v", err)
	}
	
	fmt.Println("\n4. Reading users after update...")
	if err := readUsers(db); err != nil {
		log.Printf("Error reading users: %v", err)
	}
	
	fmt.Println("\n5. Deleting user...")
	if err := deleteUser(db, 2); err != nil {
		log.Printf("Error deleting user: %v", err)
	}
	
	fmt.Println("\n6. Final user list...")
	if err := readUsers(db); err != nil {
		log.Printf("Error reading users: %v", err)
	}
	
	// Demonstrate transactions
	fmt.Println("\n7. Transaction demo...")
	if err := demonstrateTransaction(db); err != nil {
		log.Printf("Error in transaction: %v", err)
	}
	
	fmt.Println("\nSQL operations completed successfully!")
}

// Initialize database schema
func initializeSchema(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		is_active BOOLEAN DEFAULT 1
	)`
	
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	
	fmt.Println("Database schema initialized")
	return nil
}

// Create sample users
func createUsers(db *sql.DB) error {
	users := []User{
		{Name: "Alice Johnson", Email: "alice@example.com", IsActive: true},
		{Name: "Bob Smith", Email: "bob@example.com", IsActive: true},
		{Name: "Charlie Brown", Email: "charlie@example.com", IsActive: false},
	}
	
	// Prepare insert statement
	stmt, err := db.Prepare(`INSERT INTO users (name, email, is_active) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	
	for _, user := range users {
		// Execute insert for each user
		result, err := stmt.Exec(user.Name, user.Email, user.IsActive)
		if err != nil {
			log.Printf("Failed to insert user %s: %v", user.Name, err)
			continue
		}
		
		// Get inserted ID
		id, err := result.LastInsertId()
		if err != nil {
			log.Printf("Failed to get insert ID for %s: %v", user.Name, err)
		} else {
			fmt.Printf("Created user: %s (ID: %d)\n", user.Name, id)
		}
	}
	
	return nil
}

// Read all users from database
func readUsers(db *sql.DB) error {
	// Execute select query
	rows, err := db.Query(`SELECT id, name, email, created_at, is_active FROM users ORDER BY id`)
	if err != nil {
		return err
	}
	defer rows.Close()
	
	fmt.Println("Current users:")
	fmt.Println("ID | Name           | Email                | Created At          | Active")
	fmt.Println("---|----------------|----------------------|---------------------|-------")
	
	// Iterate through results
	for rows.Next() {
		var user User
		var createdAtStr string
		
		// Scan row data into variables
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &createdAtStr, &user.IsActive)
		if err != nil {
			log.Printf("Failed to scan row: %v", err)
			continue
		}
		
		// Parse created_at timestamp
		createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err == nil {
			user.CreatedAt = createdAt
		}
		
		// Format and print user data
		activeStr := "No"
		if user.IsActive {
			activeStr = "Yes"
		}
		
		fmt.Printf("%2d | %-14s | %-20s | %-19s | %s\n",
			user.ID, user.Name, user.Email, formatTime(user.CreatedAt), activeStr)
	}
	
	// Check for iteration errors
	if err := rows.Err(); err != nil {
		return err
	}
	
	return nil
}

// Update user information
func updateUser(db *sql.DB, userID int) error {
	// Update user's name and email
	query := `UPDATE users SET name = ?, email = ? WHERE id = ?`
	
	result, err := db.Exec(query, "Alice Smith-Updated", "alice.updated@example.com", userID)
	if err != nil {
		return err
	}
	
	// Check how many rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		fmt.Printf("No user found with ID %d to update\n", userID)
	} else {
		fmt.Printf("Successfully updated user ID %d\n", userID)
	}
	
	return nil
}

// Delete user from database
func deleteUser(db *sql.DB, userID int) error {
	// Delete user by ID
	query := `DELETE FROM users WHERE id = ?`
	
	result, err := db.Exec(query, userID)
	if err != nil {
		return err
	}
	
	// Check deletion result
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		fmt.Printf("No user found with ID %d to delete\n", userID)
	} else {
		fmt.Printf("Successfully deleted user ID %d\n", userID)
	}
	
	return nil
}

// Demonstrate database transactions
func demonstrateTransaction(db *sql.DB) error {
	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	
	// Set up transaction cleanup
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Transaction rolled back due to panic: %v", r)
		}
	}()
	
	// Insert multiple users in transaction
	insertQuery := `INSERT INTO users (name, email, is_active) VALUES (?, ?, ?)`
	
	// Insert first user
	_, err = tx.Exec(insertQuery, "David Wilson", "david@example.com", true)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert first user in transaction: %w", err)
	}
	
	// Insert second user
	_, err = tx.Exec(insertQuery, "Emma Davis", "emma@example.com", true)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert second user in transaction: %w", err)
	}
	
	// Simulate conditional logic (rollback scenario)
	rollbackDemo := false // Set to true to see rollback
	if rollbackDemo {
		tx.Rollback()
		fmt.Println("Transaction rolled back (demo)")
		return nil
	}
	
	// Commit transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	
	fmt.Println("Transaction completed successfully - 2 users added")
	return nil
}

// Helper function to format time
func formatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}