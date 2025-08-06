// GoForGo Exercise: Logrus Custom Formatters
// Learn how to create and use custom log formatters for different output formats and requirements

package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// TODO: Create a custom formatter that implements logrus.Formatter interface
// Requirements:
// - Format: [TIMESTAMP] [LEVEL] MESSAGE (key=value, key=value)
// - Timestamp format: "2006-01-02 15:04:05"
// - Level should be uppercase and padded to 5 characters
// - Fields should be sorted alphabetically
type CustomFormatter struct {
	// Your custom formatter struct here
}

// TODO: Implement the Format method for CustomFormatter
// The method signature should be: Format(entry *logrus.Entry) ([]byte, error)
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Your Format implementation here
	return nil, nil
}

// TODO: Create a colored console formatter
// Requirements:
// - Different colors for different log levels
// - Format: [LEVEL] MESSAGE fields...
// - Use ANSI color codes: Red for Error, Yellow for Warn, Green for Info, Blue for Debug
type ColoredFormatter struct {
	// Your colored formatter struct here
}

// TODO: Implement the Format method for ColoredFormatter
func (f *ColoredFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Your Format implementation here
	// Color codes: Red=31, Yellow=33, Green=32, Blue=34, Reset=0
	// Format: \033[<color>m<text>\033[0m
	return nil, nil
}

// TODO: Create a compact single-line formatter
// Requirements:
// - Format: LEVEL|MESSAGE|field1=val1,field2=val2
// - No timestamps, very compact for high-volume logging
type CompactFormatter struct {
	// Your compact formatter struct here
}

// TODO: Implement the Format method for CompactFormatter
func (f *CompactFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Your Format implementation here
	return nil, nil
}

func main() {
	// TODO: Create three different logger instances for testing formatters

	// Logger 1: Custom formatter
	// TODO: Create logger with CustomFormatter

	// Logger 2: Colored formatter  
	// TODO: Create logger with ColoredFormatter

	// Logger 3: Compact formatter
	// TODO: Create logger with CompactFormatter

	// TODO: Test all formatters with the same log messages:
	
	// Test message 1: Info level
	// Message: "User authentication successful"
	// Fields: userID=12345, method="password", ip="192.168.1.1"

	// Test message 2: Warning level
	// Message: "Rate limit warning"
	// Fields: endpoint="/api/users", requests=90, limit=100

	// Test message 3: Error level
	// Message: "Database connection error"
	// Fields: database="users", error="connection timeout", retries=3

	fmt.Println("=== Custom Formatter Test ===")
	// Test custom formatter here

	fmt.Println("\n=== Colored Formatter Test ===")
	// Test colored formatter here

	fmt.Println("\n=== Compact Formatter Test ===")
	// Test compact formatter here

	// TODO: Bonus: Show how to switch formatters at runtime
	// Create a single logger and change its formatter dynamically
}