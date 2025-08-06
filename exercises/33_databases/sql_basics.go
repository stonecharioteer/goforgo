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

// TODO: User struct for database operations
type User struct {
	/* define fields: ID, Name, Email, CreatedAt, IsActive */
}

func main() {
	fmt.Println("=== SQL Database Basics ===")
	
	// TODO: Create database connection
	db, err := /* open SQLite database connection to "users.db" */
	if /* check for error */ {
		/* log fatal error */
	}
	defer /* close database connection */
	
	// TODO: Initialize database schema
	if err := /* call initializeSchema with db */; err != nil {
		/* log fatal error */
	}
	
	// TODO: Demonstrate CRUD operations
	fmt.Println("1. Creating users...")
	if err := /* call createUsers with db */; err != nil {
		/* log error */
	}
	
	fmt.Println("\n2. Reading users...")
	if err := /* call readUsers with db */; err != nil {
		/* log error */
	}
	
	fmt.Println("\n3. Updating users...")
	if err := /* call updateUser with db and user ID 1 */; err != nil {
		/* log error */
	}
	
	fmt.Println("\n4. Reading users after update...")
	if err := /* call readUsers with db */; err != nil {
		/* log error */
	}
	
	fmt.Println("\n5. Deleting user...")
	if err := /* call deleteUser with db and user ID 2 */; err != nil {
		/* log error */
	}
	
	fmt.Println("\n6. Final user list...")
	if err := /* call readUsers with db */; err != nil {
		/* log error */
	}
	
	// TODO: Demonstrate transactions
	fmt.Println("\n7. Transaction demo...")
	if err := /* call demonstrateTransaction with db */; err != nil {
		/* log error */
	}
	
	fmt.Println("\nSQL operations completed successfully!")
}

// TODO: Initialize database schema
func initializeSchema(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		is_active BOOLEAN DEFAULT 1
	)`
	
	/* execute the create table query */
	if /* check for error */ {
		return /* wrap error with context */
	}
	
	fmt.Println("Database schema initialized")
	return nil
}

// TODO: Create sample users
func createUsers(db *sql.DB) error {
	users := []User{
		{Name: "Alice Johnson", Email: "alice@example.com", IsActive: true},
		{Name: "Bob Smith", Email: "bob@example.com", IsActive: true},
		{Name: "Charlie Brown", Email: "charlie@example.com", IsActive: false},
	}
	
	// TODO: Prepare insert statement
	stmt, err := /* prepare insert statement for users table */
	if /* check for error */ {
		return err
	}
	defer /* close statement */
	
	for _, user := range users {
		// TODO: Execute insert for each user
		result, err := /* execute statement with user data */
		if /* check for error */ {
			/* log error and continue */
			continue
		}
		
		// TODO: Get inserted ID
		id, err := /* get last insert ID from result */
		if /* check for error */ {
			/* log error */
		} else {
			/* log success with user name and ID */
		}
	}
	
	return nil
}

// TODO: Read all users from database
func readUsers(db *sql.DB) error {
	// TODO: Execute select query
	rows, err := /* query all users ordered by ID */
	if /* check for error */ {
		return err
	}
	defer /* close rows */
	
	fmt.Println("Current users:")
	fmt.Println("ID | Name           | Email                | Created At          | Active")
	fmt.Println("---|----------------|----------------------|---------------------|-------")
	
	// TODO: Iterate through results
	for /* iterate through rows */ {
		var user User
		var createdAtStr string
		
		// TODO: Scan row data into variables
		err := /* scan row into user struct fields and createdAtStr */
		if /* check for error */ {
			/* log error and continue */
			continue
		}
		
		// TODO: Parse created_at timestamp
		createdAt, err := /* parse createdAtStr as time */
		if err == nil {
			user.CreatedAt = createdAt
		}
		
		// TODO: Format and print user data
		activeStr := "No"
		if user.IsActive {
			activeStr = "Yes"
		}
		
		/* print formatted user data */
	}
	
	// TODO: Check for iteration errors
	if err := /* check rows.Err() */; err != nil {
		return err
	}
	
	return nil
}

// TODO: Update user information
func updateUser(db *sql.DB, userID int) error {
	// TODO: Update user's name and email
	query := `UPDATE users SET name = ?, email = ? WHERE id = ?`
	
	result, err := /* execute update query */
	if /* check for error */ {
		return err
	}
	
	// TODO: Check how many rows were affected
	rowsAffected, err := /* get rows affected from result */
	if /* check for error */ {
		return err
	}
	
	if rowsAffected == 0 {
		/* log no rows updated message */
	} else {
		/* log successful update message */
	}
	
	return nil
}

// TODO: Delete user from database
func deleteUser(db *sql.DB, userID int) error {
	// TODO: Delete user by ID
	query := `DELETE FROM users WHERE id = ?`
	
	result, err := /* execute delete query */
	if /* check for error */ {
		return err
	}
	
	// TODO: Check deletion result
	rowsAffected, err := /* get rows affected from result */
	if /* check for error */ {
		return err
	}
	
	if rowsAffected == 0 {
		/* log no rows deleted message */
	} else {
		/* log successful deletion message */
	}
	
	return nil
}

// TODO: Demonstrate database transactions
func demonstrateTransaction(db *sql.DB) error {
	// TODO: Begin transaction
	tx, err := /* begin database transaction */
	if /* check for error */ {
		return err
	}
	
	// TODO: Set up transaction cleanup
	defer func() {
		if r := recover(); r != nil {
			/* rollback transaction */
			/* log rollback due to panic */
		}
	}()
	
	// TODO: Insert multiple users in transaction
	insertQuery := `INSERT INTO users (name, email, is_active) VALUES (?, ?, ?)`
	
	// Insert first user
	_, err = /* execute insert in transaction */
	if /* check for error */ {
		/* rollback transaction */
		return /* wrap error */
	}
	
	// Insert second user
	_, err = /* execute insert in transaction */
	if /* check for error */ {
		/* rollback transaction */
		return /* wrap error */
	}
	
	// TODO: Simulate conditional logic (rollback scenario)
	rollbackDemo := false // Set to true to see rollback
	if rollbackDemo {
		/* rollback transaction */
		/* log rollback demo message */
		return nil
	}
	
	// TODO: Commit transaction
	if err := /* commit transaction */; err != nil {
		return /* wrap error */
	}
	
	/* log successful transaction completion */
	return nil
}

// TODO: Helper function to format time
func formatTime(t time.Time) string {
	return /* format time as "2006-01-02 15:04:05" */
}