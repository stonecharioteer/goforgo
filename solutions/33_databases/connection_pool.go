// connection_pool.go
// Learn database connection pooling and configuration

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"
	
	_ "github.com/mattn/go-sqlite3"
)

// Database stats structure
type DBStats struct {
	OpenConnections     int
	InUse               int
	Idle                int
	WaitCount           int64
	WaitDuration        time.Duration
	MaxIdleClosed       int64
	MaxLifetimeClosed   int64
	MaxIdleTimeClosed   int64
}

func main() {
	fmt.Println("=== Database Connection Pooling ===")
	
	// Create database with connection pool configuration
	db, err := createDatabaseWithPool()
	if err != nil {
		log.Fatal("Failed to create database with pool:", err)
	}
	defer db.Close()
	
	// Initialize schema
	if err := initSchema(db); err != nil {
		log.Fatal("Failed to initialize schema:", err)
	}
	
	// Demonstrate concurrent database access
	fmt.Println("1. Testing concurrent database access...")
	testConcurrentAccess(db)
	
	// Monitor connection pool stats
	fmt.Println("\n2. Connection pool statistics...")
	monitorConnectionPool(db)
	
	// Test connection timeout scenarios
	fmt.Println("\n3. Testing connection timeouts...")
	testConnectionTimeouts(db)
	
	// Demonstrate prepared statement pooling
	fmt.Println("\n4. Testing prepared statement pooling...")
	testPreparedStatementPool(db)
	
	fmt.Println("\nConnection pooling demo completed!")
}

// Create database with connection pool settings
func createDatabaseWithPool() (*sql.DB, error) {
	// Open database connection
	db, err := sql.Open("sqlite3", "pool_demo.db")
	if err != nil {
		return nil, err
	}
	
	// Configure connection pool settings
	db.SetMaxOpenConns(10)                 // Maximum 10 open connections
	db.SetMaxIdleConns(5)                  // Keep 5 idle connections
	db.SetConnMaxLifetime(30 * time.Minute) // Close connections after 30 minutes
	db.SetConnMaxIdleTime(5 * time.Minute)  // Close idle connections after 5 minutes
	
	// Verify database connectivity
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	
	fmt.Println("Database connection pool configured: MaxOpen=10, MaxIdle=5")
	return db, nil
}

// Initialize database schema
func initSchema(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		status TEXT DEFAULT 'pending',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	
	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}
	
	fmt.Println("Schema initialized successfully")
	return nil
}

// Test concurrent database access
func testConcurrentAccess(db *sql.DB) {
	const numWorkers = 20
	const tasksPerWorker = 10
	
	var wg sync.WaitGroup
	
	// Launch concurrent workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			
			// Each worker performs database operations
			for j := 0; j < tasksPerWorker; j++ {
				// Insert task
				err := insertTask(db, workerID, j)
				if err != nil {
					log.Printf("Worker %d failed to insert task %d: %v", workerID, j, err)
					continue
				}
				
				// Query tasks
				count, err := countTasks(db)
				if err != nil {
					log.Printf("Worker %d failed to count tasks: %v", workerID, err)
				} else if j%5 == 0 {
					fmt.Printf("Worker %d: Current task count: %d\n", workerID, count)
				}
				
				// Small delay to simulate work
				time.Sleep(10 * time.Millisecond)
			}
			
			fmt.Printf("Worker %d completed\n", workerID)
		}(i)
	}
	
	// Wait for all workers and print stats
	wg.Wait()
	printConnectionStats(db)
}

// Insert a task into database
func insertTask(db *sql.DB, workerID, taskNum int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	query := `INSERT INTO tasks (title, description) VALUES (?, ?)`
	title := fmt.Sprintf("Worker %d Task %d", workerID, taskNum)
	description := fmt.Sprintf("Task created by worker %d (task #%d)", workerID, taskNum)
	
	_, err := db.ExecContext(ctx, query, title, description)
	if err != nil {
		return fmt.Errorf("failed to insert task: %w", err)
	}
	
	return nil
}

// Count total tasks
func countTasks(db *sql.DB) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	var count int
	query := `SELECT COUNT(*) FROM tasks`
	
	err := db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count tasks: %w", err)
	}
	
	return count, nil
}

// Monitor connection pool statistics
func monitorConnectionPool(db *sql.DB) {
	// Monitor for 10 seconds, printing stats every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	
	timeout := time.NewTimer(10 * time.Second)
	defer timeout.Stop()
	
	for {
		select {
		case <-ticker.C:
			printConnectionStats(db)
		case <-timeout.C:
			fmt.Println("Monitoring completed")
			return
		}
	}
}

// Print connection pool statistics
func printConnectionStats(db *sql.DB) {
	stats := db.Stats()
	
	fmt.Printf("Connection Pool Stats: Open=%d, InUse=%d, Idle=%d, Wait=%d, WaitDuration=%v\n",
		stats.OpenConnections, stats.InUse, stats.Idle, stats.WaitCount, stats.WaitDuration)
}

// Test connection timeout scenarios
func testConnectionTimeouts(db *sql.DB) {
	// Create very short timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	
	// Try to execute query with short timeout (should fail)
	var count int
	err := db.QueryRowContext(ctx, `SELECT COUNT(*) FROM tasks`).Scan(&count)
	if err != nil {
		fmt.Printf("Expected timeout error: %v\n", err)
	} else {
		fmt.Printf("Unexpected success with very short timeout: count=%d\n", count)
	}
	
	// Test with reasonable timeout
	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()
	
	err = db.QueryRowContext(ctx2, `SELECT COUNT(*) FROM tasks`).Scan(&count)
	if err != nil {
		fmt.Printf("Unexpected error with reasonable timeout: %v\n", err)
	} else {
		fmt.Printf("Successful query with timeout: count=%d\n", count)
	}
}

// Test prepared statement pooling
func testPreparedStatementPool(db *sql.DB) {
	// Prepare statement for reuse
	stmt, err := db.Prepare(`INSERT INTO tasks (title, description) VALUES (?, ?)`)
	if err != nil {
		log.Printf("Failed to prepare statement: %v", err)
		return
	}
	defer stmt.Close()
	
	// Use prepared statement multiple times
	for i := 0; i < 5; i++ {
		title := fmt.Sprintf("Prepared Statement Task %d", i)
		description := fmt.Sprintf("Task #%d using prepared statement", i)
		
		_, err := stmt.Exec(title, description)
		if err != nil {
			log.Printf("Failed to execute prepared statement for task %d: %v", i, err)
		} else {
			fmt.Printf("Successfully executed prepared statement for task %d\n", i)
		}
	}
	
	fmt.Println("Prepared statement pooling demo completed")
}

// Cleanup database for demo
func cleanupDatabase(db *sql.DB) error {
	_, err := db.Exec(`DELETE FROM tasks`)
	if err != nil {
		return fmt.Errorf("failed to cleanup database: %w", err)
	}
	
	fmt.Println("Database cleaned up")
	return nil
}