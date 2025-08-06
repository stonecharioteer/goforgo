// GoForGo Exercise: Logrus Log Levels
// Learn how to configure and use different log levels, filtering, and environment-based level setting

package main

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	// TODO: Create a new logrus logger instance

	// TODO: Create a function to set log level from environment variable
	// Check for LOG_LEVEL environment variable
	// Default to "info" if not set
	// Support levels: "trace", "debug", "info", "warn", "error", "fatal", "panic"
	// Hint: Use logrus.ParseLevel() and logger.SetLevel()

	// TODO: Set the formatter to text with timestamps
	// Use logrus.TextFormatter with TimestampFormat

	// TODO: Demonstrate all log levels with meaningful messages:

	// Trace level (most verbose)
	// Message: "Entering function processPayment"
	// Fields: function="processPayment", args=["userID", "amount"]

	// Debug level (debugging information)
	// Message: "Processing payment validation"
	// Fields: step="validation", userID=123, amount=50.00

	// Info level (general information)
	// Message: "Payment processing started"
	// Fields: userID=123, amount=50.00, currency="USD"

	// Warn level (warning conditions)
	// Message: "Payment amount exceeds daily limit"
	// Fields: amount=50.00, dailyLimit=100.00, remaining=50.00

	// Error level (error conditions)
	// Message: "Payment gateway returned error"
	// Fields: gateway="stripe", errorCode="card_declined", userID=123

	// TODO: Create a function that demonstrates conditional logging
	// Only log debug messages if debug mode is enabled
	// Check if log level is Debug or Trace before logging expensive operations

	// TODO: Demonstrate different loggers with different levels
	// Create separate loggers for different components:
	// - Database logger (level: Warn)
	// - API logger (level: Info) 
	// - Auth logger (level: Debug)
	// Each should log the same message but only some will appear based on level

	// TODO: Show log level hierarchy
	// Print current log level
	// Demonstrate that setting level to Warn will show Warn, Error, Fatal, Panic
	// But not Trace, Debug, Info
}