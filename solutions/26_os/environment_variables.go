// environment_variables.go - SOLUTION
// Learn environment variable management: reading, setting, parsing, and validation

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== Environment Variables ===")
	
	fmt.Println("\n=== Reading Environment Variables ===")
	
	// Read common environment variables
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
		// Get environment variable value
		value := os.Getenv(varName)
		if value != "" {
			// Truncate long values for display
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
	
	// Set custom environment variables
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
		// Set environment variable
		err := os.Setenv(key, value)
		if err != nil {
			fmt.Printf("❌ Failed to set %s: %v\n", key, err)
		} else {
			fmt.Printf("✅ Set %s = %s\n", key, value)
		}
	}
	
	fmt.Println("\n=== Environment Variable Parsing ===")
	
	// Parse different types of environment variables
	fmt.Println("Parsing environment variables:")
	
	// Parse boolean
	debugMode := parseBool(os.Getenv("DEBUG_MODE"))
	fmt.Printf("Debug Mode (bool): %t\n", debugMode)
	
	// Parse integer
	maxConnections := parseInt(os.Getenv("MAX_CONNECTIONS"), 10)
	fmt.Printf("Max Connections (int): %d\n", maxConnections)
	
	// Parse duration
	apiTimeout := parseDuration(os.Getenv("API_TIMEOUT"), 10*time.Second)
	fmt.Printf("API Timeout (duration): %v\n", apiTimeout)
	
	// Parse URL components
	dbURL := os.Getenv("DATABASE_URL")
	fmt.Printf("Database URL: %s\n", dbURL)
	// Parse URL to extract components
	if u, err := url.Parse(dbURL); err == nil {
		fmt.Printf("  Scheme: %s\n", u.Scheme)
		fmt.Printf("  Host: %s\n", u.Host)
		fmt.Printf("  Path: %s\n", u.Path)
		if u.User != nil {
			fmt.Printf("  User: %s\n", u.User.Username())
		}
	}
	
	fmt.Println("\n=== Environment Variable Validation ===")
	
	// Validate environment variables
	type EnvVar struct {
		Name         string
		Required     bool
		DefaultValue string
		Validator    func(string) error
	}
	
	// Define validation rules
	envVars := []EnvVar{
		{
			Name:     "APP_NAME",
			Required: true,
			Validator: func(value string) error {
				if len(value) < 2 {
					return errors.New("app name must be at least 2 characters")
				}
				return nil
			},
		},
		{
			Name:     "MAX_CONNECTIONS",
			Required: true,
			Validator: func(value string) error {
				if n, err := strconv.Atoi(value); err != nil {
					return errors.New("must be a valid integer")
				} else if n <= 0 {
					return errors.New("must be greater than 0")
				} else if n > 1000 {
					return errors.New("must be <= 1000")
				}
				return nil
			},
		},
		{
			Name:         "LOG_LEVEL",
			Required:     false,
			DefaultValue: "info",
			Validator: func(value string) error {
				validLevels := []string{"debug", "info", "warn", "error"}
				for _, level := range validLevels {
					if value == level {
						return nil
					}
				}
				return fmt.Errorf("must be one of: %s", strings.Join(validLevels, ", "))
			},
		},
		{
			Name:     "API_TIMEOUT",
			Required: true,
			Validator: func(value string) error {
				if _, err := time.ParseDuration(value); err != nil {
					return errors.New("must be a valid duration (e.g., 30s, 1m)")
				}
				return nil
			},
		},
	}
	
	fmt.Println("Validating environment variables:")
	for _, envVar := range envVars {
		// Get and validate environment variable
		value := os.Getenv(envVar.Name)
		
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
	
	// Configuration struct
	type Config struct {
		AppName        string        `env:"APP_NAME"`
		Version        string        `env:"APP_VERSION"`
		DebugMode      bool          `env:"DEBUG_MODE"`
		MaxConnections int           `env:"MAX_CONNECTIONS"`
		APITimeout     time.Duration `env:"API_TIMEOUT"`
		DatabaseURL    string        `env:"DATABASE_URL"`
	}
	
	// Load configuration from environment
	config := loadConfig()
	
	fmt.Printf("Loaded configuration:\n")
	fmt.Printf("  App Name: %s\n", config.AppName)
	fmt.Printf("  Version: %s\n", config.Version)
	fmt.Printf("  Debug Mode: %t\n", config.DebugMode)
	fmt.Printf("  Max Connections: %d\n", config.MaxConnections)
	fmt.Printf("  API Timeout: %v\n", config.APITimeout)
	fmt.Printf("  Database URL: %s\n", config.DatabaseURL)
	
	fmt.Println("\n=== Environment Profiles ===")
	
	// Handle different environments (dev, staging, prod)
	environment := getEnv("ENVIRONMENT", "development")
	fmt.Printf("Current environment: %s\n", environment)
	
	// Load environment-specific configuration
	switch environment {
	case "development":
		// Development-specific settings
		fmt.Println("Loading development configuration...")
		os.Setenv("LOG_LEVEL", "debug")
		os.Setenv("CACHE_TTL", "1m")
	case "staging":
		// Staging-specific settings
		fmt.Println("Loading staging configuration...")
		os.Setenv("LOG_LEVEL", "info")
		os.Setenv("CACHE_TTL", "5m")
	case "production":
		// Production-specific settings
		fmt.Println("Loading production configuration...")
		os.Setenv("LOG_LEVEL", "warn")
		os.Setenv("CACHE_TTL", "15m")
	default:
		fmt.Printf("Unknown environment '%s', using defaults\n", environment)
	}
	
	fmt.Println("\n=== Environment Variable Expansion ===")
	
	// Expand variables with references to other variables
	templateVars := map[string]string{
		"BASE_URL":    "https://api.example.com",
		"API_V1_URL":  "${BASE_URL}/v1",
		"API_V2_URL":  "${BASE_URL}/v2",
		"HEALTH_URL":  "${API_V1_URL}/health",
		"METRICS_URL": "${API_V2_URL}/metrics",
	}
	
	fmt.Println("Expanding environment variables:")
	for key, template := range templateVars {
		// Set template variable
		os.Setenv(key, template)
		
		// Expand variables in template
		expanded := expandEnvVars(template)
		fmt.Printf("  %s: %s → %s\n", key, template, expanded)
	}
	
	fmt.Println("\n=== Environment Variable Security ===")
	
	// Handle sensitive environment variables
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
			// Mask sensitive values in logs
			masked := maskSensitive(value)
			fmt.Printf("  %s: %s (masked)\n", varName, masked)
		} else {
			fmt.Printf("  %s: (not set)\n", varName)
		}
	}
	
	fmt.Println("\n=== Environment Variable Utilities ===")
	
	// List all environment variables
	fmt.Println("All environment variables:")
	allEnv := os.Environ()
	
	// Filter and sort environment variables
	var appVars []string
	for _, env := range allEnv {
		if strings.HasPrefix(env, "APP_") || strings.HasPrefix(env, "API_") || 
		   strings.HasPrefix(env, "DATABASE_") || strings.HasPrefix(env, "REDIS_") {
			appVars = append(appVars, env)
		}
	}
	
	sort.Strings(appVars)
	fmt.Printf("Found %d app-related variables:\n", len(appVars))
	for _, env := range appVars {
		fmt.Printf("  %s\n", env)
	}
	
	// Export environment to different formats
	fmt.Println("\nExporting app variables to different formats:")
	
	// Export as shell script
	fmt.Println("Shell script format:")
	for _, env := range appVars {
		fmt.Printf("export %s\n", env)
	}
	
	// Export as Docker ENV format
	fmt.Println("\nDocker ENV format:")
	for _, env := range appVars {
		fmt.Printf("ENV %s\n", env)
	}
	
	// Export as JSON
	fmt.Println("\nJSON format:")
	envMap := make(map[string]string)
	for _, env := range appVars {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) == 2 {
			envMap[parts[0]] = parts[1]
		}
	}
	if jsonData, err := json.MarshalIndent(envMap, "", "  "); err == nil {
		fmt.Println(string(jsonData))
	}
	
	fmt.Println("\n=== Cleanup ===")
	
	// Clean up test environment variables
	fmt.Println("Cleaning up test environment variables...")
	for key := range customVars {
		// Unset environment variable
		err := os.Unsetenv(key)
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

// Helper functions

func parseBool(value string) bool {
	if b, err := strconv.ParseBool(value); err == nil {
		return b
	}
	return false
}

func parseInt(value string, defaultValue int) int {
	if i, err := strconv.Atoi(value); err == nil {
		return i
	}
	return defaultValue
}

func parseDuration(value string, defaultValue time.Duration) time.Duration {
	if d, err := time.ParseDuration(value); err == nil {
		return d
	}
	return defaultValue
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func loadConfig() Config {
	return Config{
		AppName:        os.Getenv("APP_NAME"),
		Version:        os.Getenv("APP_VERSION"),
		DebugMode:      parseBool(os.Getenv("DEBUG_MODE")),
		MaxConnections: parseInt(os.Getenv("MAX_CONNECTIONS"), 10),
		APITimeout:     parseDuration(os.Getenv("API_TIMEOUT"), 30*time.Second),
		DatabaseURL:    os.Getenv("DATABASE_URL"),
	}
}

func expandEnvVars(template string) string {
	expanded := template
	
	// Simple variable expansion - replace ${VAR} with value
	for {
		start := strings.Index(expanded, "${")
		if start == -1 {
			break
		}
		
		end := strings.Index(expanded[start:], "}")
		if end == -1 {
			break
		}
		end += start
		
		varName := expanded[start+2 : end]
		varValue := os.Getenv(varName)
		
		expanded = expanded[:start] + varValue + expanded[end+1:]
	}
	
	return expanded
}

func maskSensitive(value string) string {
	if len(value) <= 8 {
		return strings.Repeat("*", len(value))
	}
	
	// Show first 4 and last 4 characters
	return value[:4] + strings.Repeat("*", len(value)-8) + value[len(value)-4:]
}