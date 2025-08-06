// GoForGo Exercise: Logrus Structured Logging
// Learn how to use Logrus for structured logging with fields, contexts, and different log levels

package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	// TODO: Create a new logrus logger instance

	// TODO: Set the log level to Info (this will log Info, Warn, Error, Fatal, Panic)

	// TODO: Set the log format to JSON for structured logging

	// TODO: Log a simple info message
	// Message: "Application started"

	// TODO: Log an info message with fields
	// Message: "User logged in"
	// Fields: userID=12345, username="john_doe", ip="192.168.1.100"

	// TODO: Log a warning message with context
	// Message: "Rate limit approached"
	// Fields: endpoint="/api/users", requests=95, limit=100

	// TODO: Log an error message with error context
	// Create a mock error and log it with additional context
	// Message: "Database connection failed"
	// Fields: database="users_db", host="localhost", port=5432, and the error

	// TODO: Use WithFields to create a logger with persistent fields
	// Create a logger with service="user-service" and version="1.0.0"
	// Log multiple messages with this logger to see persistent fields

	// TODO: Create different log levels demonstration
	// - Debug: "Processing request details"
	// - Info: "Request completed successfully"
	// - Warn: "Deprecated API endpoint used"
	// - Error: "Validation failed"

	// TODO: Log with nested fields (using logrus.Fields)
	// Message: "Payment processed"
	// Fields: 
	//   - amount: 99.99
	//   - currency: "USD"
	//   - user: {id: 123, email: "user@example.com"}
	//   - transaction: {id: "txn_abc123", status: "completed"}
}