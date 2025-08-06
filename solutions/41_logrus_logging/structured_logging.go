// GoForGo Solution: Logrus Structured Logging
// Complete implementation of structured logging with fields, contexts, and different log levels

package main

import (
	"errors"

	"github.com/sirupsen/logrus"
)

func main() {
	// Create a new logrus logger instance
	logger := logrus.New()

	// Set the log level to Info
	logger.SetLevel(logrus.InfoLevel)

	// Set the log format to JSON for structured logging
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Log a simple info message
	logger.Info("Application started")

	// Log an info message with fields
	logger.WithFields(logrus.Fields{
		"userID":   12345,
		"username": "john_doe",
		"ip":       "192.168.1.100",
	}).Info("User logged in")

	// Log a warning message with context
	logger.WithFields(logrus.Fields{
		"endpoint": "/api/users",
		"requests": 95,
		"limit":    100,
	}).Warn("Rate limit approached")

	// Log an error message with error context
	mockError := errors.New("connection timeout")
	logger.WithFields(logrus.Fields{
		"database": "users_db",
		"host":     "localhost",
		"port":     5432,
		"error":    mockError.Error(),
	}).Error("Database connection failed")

	// Use WithFields to create a logger with persistent fields
	serviceLogger := logger.WithFields(logrus.Fields{
		"service": "user-service",
		"version": "1.0.0",
	})

	// Log multiple messages with persistent fields
	serviceLogger.Info("Service initialized")
	serviceLogger.Info("Health check endpoint registered")
	serviceLogger.Info("Service ready to accept requests")

	// Create different log levels demonstration
	logger.Debug("Processing request details") // This won't show due to Info level
	logger.Info("Request completed successfully")
	logger.Warn("Deprecated API endpoint used")
	logger.Error("Validation failed")

	// Log with nested fields
	logger.WithFields(logrus.Fields{
		"amount":   99.99,
		"currency": "USD",
		"user": logrus.Fields{
			"id":    123,
			"email": "user@example.com",
		},
		"transaction": logrus.Fields{
			"id":     "txn_abc123",
			"status": "completed",
		},
	}).Info("Payment processed")

	// Additional demonstration: Logging with method chaining
	logger.WithField("component", "authentication").
		WithField("action", "login_attempt").
		WithField("success", true).
		Info("Authentication completed")

	logger.Info("Structured logging demonstration completed")
}