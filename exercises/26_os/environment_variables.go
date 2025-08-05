// environment_variables.go
// Learn environment variable management: reading, setting, parsing, and validation

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== Environment Variables ===")
	
	fmt.Println("\n=== Reading Environment Variables ===")
	
	// TODO: Read common environment variables
	commonVars := []string{
		"HOME",
		"USER", 
		"PATH",
		"PWD",
		"SHELL",
		"LANG",
		"TERM",
	}
	
	fmt.Println("Common environment variables:")
	for _, varName := range commonVars {
		// TODO: Get environment variable value
		value := /* get environment variable */
		if value != "" {
			// TODO: Truncate long values for display
			displayValue := value
			if len(displayValue) > 50 {
				displayValue = displayValue[:47] + "..."
			}
			fmt.Printf("  %s = %s\n", varName, displayValue)
		} else {
			fmt.Printf("  %s = (not set)\n", varName)
		}
	}
	
	fmt.Println("\n=== Setting Environment Variables ===")
	
	// TODO: Set custom environment variables
	customVars := map[string]string{
		"APP_NAME":        "GoForGo",
		"APP_VERSION":     "1.0.0",
		"DEBUG_MODE":      "true",
		"MAX_CONNECTIONS": "100",
		"API_TIMEOUT":     "30s",
		"DATABASE_URL":    "postgres://localhost:5432/myapp",
		"REDIS_URL":       "redis://localhost:6379",
	}
	
	fmt.Println("Setting custom environment variables:")
	for key, value := range customVars {
		// TODO: Set environment variable
		err := /* set environment variable */
		if err != nil {
			fmt.Printf("❌ Failed to set %s: %v\n", key, err)
		} else {
			fmt.Printf("✅ Set %s = %s\n", key, value)
		}
	}
	
	fmt.Println("\n=== Environment Variable Parsing ===")
	
	// TODO: Parse different types of environment variables
	fmt.Println("Parsing environment variables:")
	
	// TODO: Parse boolean
	debugMode := /* parse DEBUG_MODE as boolean */
	fmt.Printf("Debug Mode (bool): %t\n", debugMode)
	
	// TODO: Parse integer
	maxConnections := /* parse MAX_CONNECTIONS as integer */
	fmt.Printf("Max Connections (int): %d\n", maxConnections)
	
	// TODO: Parse duration
	apiTimeout := /* parse API_TIMEOUT as duration */
	fmt.Printf("API Timeout (duration): %v\n", apiTimeout)
	
	// TODO: Parse URL components
	dbURL := /* get DATABASE_URL */
	fmt.Printf("Database URL: %s\n", dbURL)
	// TODO: Parse URL to extract components
	/* parse database URL and display components */
	
	fmt.Println("\n=== Environment Variable Validation ===")
	
	// TODO: Validate environment variables
	type EnvVar struct {
		Name         string
		Required     bool
		DefaultValue string
		Validator    func(string) error
	}
	
	// TODO: Define validation rules
	envVars := []EnvVar{
		{
			Name:     "APP_NAME",
			Required: true,
			Validator: func(value string) error {
				// TODO: Validate app name
			},
		},
		{
			Name:     "MAX_CONNECTIONS",
			Required: true,
			Validator: func(value string) error {
				// TODO: Validate max connections
			},
		},
		{
			Name:         "LOG_LEVEL",
			Required:     false,
			DefaultValue: "info",
			Validator: func(value string) error {
				// TODO: Validate log level
			},
		},
		{
			Name:     "API_TIMEOUT",
			Required: true,
			Validator: func(value string) error {
				// TODO: Validate timeout duration
			},
		},
	}
	
	fmt.Println("Validating environment variables:")
	for _, envVar := range envVars {
		// TODO: Get and validate environment variable
		value := /* get environment variable with default */
		
		if envVar.Required && value == "" {
			fmt.Printf("❌ %s: Required but not set\n", envVar.Name)
			continue
		}
		
		if value == "" && envVar.DefaultValue != "" {
			value = envVar.DefaultValue
		}
		
		if envVar.Validator != nil {
			if err := envVar.Validator(value); err != nil {
				fmt.Printf("❌ %s: Invalid value '%s' - %v\n", envVar.Name, value, err)
			} else {
				fmt.Printf("✅ %s: Valid value '%s'\n", envVar.Name, value)
			}
		}
	}
	
	fmt.Println("\n=== Environment Configuration Management ===")
	
	// TODO: Configuration struct
	type Config struct {
		// TODO: Define configuration fields with tags
	}
	
	// TODO: Load configuration from environment
	config := /* load configuration from environment */
	
	fmt.Printf("Loaded configuration:\n")
	/* display configuration */
	
	fmt.Println("\n=== Environment Profiles ===")
	
	// TODO: Handle different environments (dev, staging, prod)
	environment := /* get current environment */
	fmt.Printf("Current environment: %s\n", environment)
	
	// TODO: Load environment-specific configuration
	switch environment {
	case "development":
		// TODO: Development-specific settings
		fmt.Println("Loading development configuration...")
		/* set development defaults */
	case "staging":
		// TODO: Staging-specific settings
		fmt.Println("Loading staging configuration...")
		/* set staging defaults */
	case "production":
		// TODO: Production-specific settings
		fmt.Println("Loading production configuration...")
		/* set production defaults */
	default:
		fmt.Printf("Unknown environment '%s', using defaults\n", environment)
	}
	
	fmt.Println("\n=== Environment Variable Expansion ===")
	
	// TODO: Expand variables with references to other variables
	templateVars := map[string]string{
		"BASE_URL":    "https://api.example.com",
		"API_V1_URL":  "${BASE_URL}/v1",
		"API_V2_URL":  "${BASE_URL}/v2",
		"HEALTH_URL":  "${API_V1_URL}/health",
		"METRICS_URL": "${API_V2_URL}/metrics",
	}
	
	fmt.Println("Expanding environment variables:")
	for key, template := range templateVars {
		// TODO: Set template variable
		os.Setenv(key, template)
		
		// TODO: Expand variables in template
		expanded := /* expand environment variables in template */
		fmt.Printf("  %s: %s → %s\n", key, template, expanded)
	}
	
	fmt.Println("\n=== Environment Variable Security ===")
	
	// TODO: Handle sensitive environment variables
	sensitiveVars := []string{
		"DATABASE_PASSWORD",
		"API_SECRET_KEY", 
		"JWT_SECRET",
		"PRIVATE_KEY",
		"ACCESS_TOKEN",
	}
	
	fmt.Println("Checking for sensitive variables:")
	for _, varName := range sensitiveVars {
		value := os.Getenv(varName)
		if value != "" {
			// TODO: Mask sensitive values in logs
			masked := /* mask sensitive value */
			fmt.Printf("  %s: %s (masked)\n", varName, masked)
		} else {
			fmt.Printf("  %s: (not set)\n", varName)
		}
	}
	
	fmt.Println("\n=== Environment Variable Utilities ===")
	
	// TODO: List all environment variables
	fmt.Println("All environment variables:")
	allEnv := /* get all environment variables */
	
	// TODO: Filter and sort environment variables
	var appVars []string
	for _, env := range allEnv {
		if /* check if app-related variable */ {
			appVars = append(appVars, env)
		}
	}
	
	fmt.Printf("Found %d app-related variables:\n", len(appVars))
	for _, env := range appVars {
		fmt.Printf("  %s\n", env)
	}
	
	// TODO: Export environment to different formats
	fmt.Println("\nExporting app variables to different formats:")
	
	// TODO: Export as shell script
	fmt.Println("Shell script format:")
	/* export as shell script */
	
	// TODO: Export as Docker ENV format
	fmt.Println("\nDocker ENV format:")
	/* export as Docker ENV */
	
	// TODO: Export as JSON
	fmt.Println("\nJSON format:")
	/* export as JSON */
	
	fmt.Println("\n=== Cleanup ===")
	
	// TODO: Clean up test environment variables
	fmt.Println("Cleaning up test environment variables...")
	for key := range customVars {
		// TODO: Unset environment variable
		err := /* unset environment variable */
		if err != nil {
			fmt.Printf("❌ Failed to unset %s: %v\n", key, err)
		} else {
			fmt.Printf("✅ Unset %s\n", key)
		}
	}
	
	fmt.Println("\n=== Best Practices ===")
	fmt.Println("Environment variable best practices:")
	fmt.Println("✅ Use descriptive, consistent naming (UPPER_SNAKE_CASE)")
	fmt.Println("✅ Provide sensible defaults for non-critical settings")
	fmt.Println("✅ Validate environment variables at startup")
	fmt.Println("✅ Never log sensitive environment variables")
	fmt.Println("✅ Use environment-specific configurations")
	fmt.Println("✅ Document all required environment variables")
	fmt.Println("✅ Use tools like direnv for local development")
	fmt.Println("✅ Consider using configuration files for complex settings")
	fmt.Println("✅ Use type-safe parsing for non-string values")
	fmt.Println("✅ Fail fast if required variables are missing")
}