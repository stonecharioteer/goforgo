// GoForGo Solution: Logrus Log Levels
// Complete implementation of log levels, filtering, and environment-based configuration

package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func setLogLevelFromEnv(logger *logrus.Logger) {
	// Get log level from environment variable, default to "info"
	levelStr := os.Getenv("LOG_LEVEL")
	if levelStr == "" {
		levelStr = "info"
	}
	
	// Parse and set the log level
	level, err := logrus.ParseLevel(strings.ToLower(levelStr))
	if err != nil {
		logger.Warnf("Invalid log level '%s', defaulting to info", levelStr)
		level = logrus.InfoLevel
	}
	
	logger.SetLevel(level)
	logger.Infof("Log level set to: %s", level.String())
}

func conditionalDebugLog(logger *logrus.Logger, message string, fields logrus.Fields) {
	// Only log if debug level is enabled (Debug or Trace)
	if logger.IsLevelEnabled(logrus.DebugLevel) {
		logger.WithFields(fields).Debug(message)
	}
}

func main() {
	// Create a new logrus logger instance
	logger := logrus.New()

	// Set log level from environment variable
	setLogLevelFromEnv(logger)

	// Set the formatter to text with timestamps
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
		FullTimestamp:   true,
	})

	// Demonstrate all log levels with meaningful messages

	// Trace level (most verbose)
	logger.WithFields(logrus.Fields{
		"function": "processPayment",
		"args":     []string{"userID", "amount"},
	}).Trace("Entering function processPayment")

	// Debug level (debugging information)
	logger.WithFields(logrus.Fields{
		"step":   "validation",
		"userID": 123,
		"amount": 50.00,
	}).Debug("Processing payment validation")

	// Info level (general information)
	logger.WithFields(logrus.Fields{
		"userID":   123,
		"amount":   50.00,
		"currency": "USD",
	}).Info("Payment processing started")

	// Warn level (warning conditions)
	logger.WithFields(logrus.Fields{
		"amount":     50.00,
		"dailyLimit": 100.00,
		"remaining":  50.00,
	}).Warn("Payment amount exceeds daily limit")

	// Error level (error conditions)
	logger.WithFields(logrus.Fields{
		"gateway":   "stripe",
		"errorCode": "card_declined",
		"userID":    123,
	}).Error("Payment gateway returned error")

	// Demonstrate conditional logging
	conditionalDebugLog(logger, "Expensive debug operation completed", logrus.Fields{
		"operation": "database_query_analysis",
		"duration":  "2.5s",
		"rows":      1000,
	})

	// Demonstrate different loggers with different levels
	fmt.Println("\n=== Component-specific loggers ===")

	// Database logger (level: Warn)
	dbLogger := logrus.New()
	dbLogger.SetLevel(logrus.WarnLevel)
	dbLogger.SetFormatter(&logrus.TextFormatter{Prefix: "[DB] "})

	// API logger (level: Info)
	apiLogger := logrus.New()
	apiLogger.SetLevel(logrus.InfoLevel)
	apiLogger.SetFormatter(&logrus.TextFormatter{Prefix: "[API] "})

	// Auth logger (level: Debug)
	authLogger := logrus.New()
	authLogger.SetLevel(logrus.DebugLevel)
	authLogger.SetFormatter(&logrus.TextFormatter{Prefix: "[AUTH] "})

	// Each logs the same message but only some will appear based on level
	message := "Component operation completed"
	dbLogger.Debug(message)   // Won't appear (Warn level)
	dbLogger.Info(message)    // Won't appear (Warn level)
	dbLogger.Warn(message)    // Will appear

	apiLogger.Debug(message)  // Won't appear (Info level)
	apiLogger.Info(message)   // Will appear
	apiLogger.Warn(message)   // Will appear

	authLogger.Debug(message) // Will appear (Debug level)
	authLogger.Info(message)  // Will appear
	authLogger.Warn(message)  // Will appear

	// Show log level hierarchy
	fmt.Printf("\n=== Log Level Information ===\n")
	fmt.Printf("Current log level: %s\n", logger.GetLevel().String())
	fmt.Printf("Available levels (most to least verbose): %s\n", 
		strings.Join([]string{"trace", "debug", "info", "warn", "error", "fatal", "panic"}, " -> "))
	
	// Demonstrate level checking
	levels := []logrus.Level{
		logrus.TraceLevel, logrus.DebugLevel, logrus.InfoLevel,
		logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel,
	}
	
	fmt.Printf("Levels enabled at current setting (%s):\n", logger.GetLevel().String())
	for _, level := range levels {
		enabled := logger.IsLevelEnabled(level)
		status := "disabled"
		if enabled {
			status = "enabled"
		}
		fmt.Printf("  %s: %s\n", level.String(), status)
	}

	logger.Info("Log levels demonstration completed")
}