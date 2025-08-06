// GoForGo Solution: Logrus Custom Formatters
// Complete implementation of custom log formatters for different output formats

package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// CustomFormatter implements logrus.Formatter interface
type CustomFormatter struct{}

// Format implements the logrus.Formatter interface
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b bytes.Buffer
	
	// Format timestamp
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	
	// Format level (uppercase, padded to 5 characters)
	level := strings.ToUpper(entry.Level.String())
	level = fmt.Sprintf("%-5s", level)
	
	// Start building the log line
	b.WriteString(fmt.Sprintf("[%s] [%s] %s", timestamp, level, entry.Message))
	
	// Add fields in alphabetical order
	if len(entry.Data) > 0 {
		b.WriteString(" (")
		
		// Sort field keys
		keys := make([]string, 0, len(entry.Data))
		for k := range entry.Data {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		
		// Add sorted fields
		for i, k := range keys {
			if i > 0 {
				b.WriteString(", ")
			}
			b.WriteString(fmt.Sprintf("%s=%v", k, entry.Data[k]))
		}
		b.WriteString(")")
	}
	
	b.WriteString("\n")
	return b.Bytes(), nil
}

// ColoredFormatter implements logrus.Formatter interface with colors
type ColoredFormatter struct{}

// Format implements the logrus.Formatter interface with ANSI colors
func (f *ColoredFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b bytes.Buffer
	
	// Color codes
	var color int
	switch entry.Level {
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		color = 31 // Red
	case logrus.WarnLevel:
		color = 33 // Yellow
	case logrus.InfoLevel:
		color = 32 // Green
	case logrus.DebugLevel, logrus.TraceLevel:
		color = 34 // Blue
	default:
		color = 0 // No color
	}
	
	// Format level with color
	level := strings.ToUpper(entry.Level.String())
	coloredLevel := fmt.Sprintf("\033[%dm[%s]\033[0m", color, level)
	
	// Build the log line
	b.WriteString(fmt.Sprintf("%s %s", coloredLevel, entry.Message))
	
	// Add fields
	if len(entry.Data) > 0 {
		b.WriteString(" ")
		keys := make([]string, 0, len(entry.Data))
		for k := range entry.Data {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		
		for i, k := range keys {
			if i > 0 {
				b.WriteString(" ")
			}
			b.WriteString(fmt.Sprintf("%s=%v", k, entry.Data[k]))
		}
	}
	
	b.WriteString("\n")
	return b.Bytes(), nil
}

// CompactFormatter implements logrus.Formatter interface for compact output
type CompactFormatter struct{}

// Format implements the logrus.Formatter interface in compact format
func (f *CompactFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b bytes.Buffer
	
	// Format: LEVEL|MESSAGE|field1=val1,field2=val2
	level := strings.ToUpper(entry.Level.String())
	b.WriteString(fmt.Sprintf("%s|%s", level, entry.Message))
	
	// Add fields as comma-separated key=value pairs
	if len(entry.Data) > 0 {
		b.WriteString("|")
		
		keys := make([]string, 0, len(entry.Data))
		for k := range entry.Data {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		
		for i, k := range keys {
			if i > 0 {
				b.WriteString(",")
			}
			b.WriteString(fmt.Sprintf("%s=%v", k, entry.Data[k]))
		}
	}
	
	b.WriteString("\n")
	return b.Bytes(), nil
}

func main() {
	// Create three different logger instances for testing formatters

	// Logger 1: Custom formatter
	customLogger := logrus.New()
	customLogger.SetFormatter(&CustomFormatter{})
	customLogger.SetLevel(logrus.InfoLevel)

	// Logger 2: Colored formatter
	coloredLogger := logrus.New()
	coloredLogger.SetFormatter(&ColoredFormatter{})
	coloredLogger.SetLevel(logrus.InfoLevel)

	// Logger 3: Compact formatter
	compactLogger := logrus.New()
	compactLogger.SetFormatter(&CompactFormatter{})
	compactLogger.SetLevel(logrus.InfoLevel)

	// Test all formatters with the same log messages

	fmt.Println("=== Custom Formatter Test ===")
	customLogger.WithFields(logrus.Fields{
		"userID": 12345,
		"method": "password",
		"ip":     "192.168.1.1",
	}).Info("User authentication successful")

	customLogger.WithFields(logrus.Fields{
		"endpoint": "/api/users",
		"requests": 90,
		"limit":    100,
	}).Warn("Rate limit warning")

	customLogger.WithFields(logrus.Fields{
		"database": "users",
		"error":    "connection timeout",
		"retries":  3,
	}).Error("Database connection error")

	fmt.Println("\n=== Colored Formatter Test ===")
	coloredLogger.WithFields(logrus.Fields{
		"userID": 12345,
		"method": "password",
		"ip":     "192.168.1.1",
	}).Info("User authentication successful")

	coloredLogger.WithFields(logrus.Fields{
		"endpoint": "/api/users",
		"requests": 90,
		"limit":    100,
	}).Warn("Rate limit warning")

	coloredLogger.WithFields(logrus.Fields{
		"database": "users",
		"error":    "connection timeout",
		"retries":  3,
	}).Error("Database connection error")

	fmt.Println("\n=== Compact Formatter Test ===")
	compactLogger.WithFields(logrus.Fields{
		"userID": 12345,
		"method": "password",
		"ip":     "192.168.1.1",
	}).Info("User authentication successful")

	compactLogger.WithFields(logrus.Fields{
		"endpoint": "/api/users",
		"requests": 90,
		"limit":    100,
	}).Warn("Rate limit warning")

	compactLogger.WithFields(logrus.Fields{
		"database": "users",
		"error":    "connection timeout",
		"retries":  3,
	}).Error("Database connection error")

	// Bonus: Show how to switch formatters at runtime
	fmt.Println("\n=== Runtime Formatter Switching ===")
	dynamicLogger := logrus.New()
	
	fmt.Println("Switching to JSON formatter:")
	dynamicLogger.SetFormatter(&logrus.JSONFormatter{})
	dynamicLogger.Info("This is JSON formatted")
	
	fmt.Println("Switching to Text formatter:")
	dynamicLogger.SetFormatter(&logrus.TextFormatter{})
	dynamicLogger.Info("This is text formatted")
	
	fmt.Println("Switching to Custom formatter:")
	dynamicLogger.SetFormatter(&CustomFormatter{})
	dynamicLogger.Info("This is custom formatted")

	fmt.Println("Custom formatters demonstration completed!")
}