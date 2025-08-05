// error_wrapping.go
// Learn error wrapping and unwrapping patterns in Go 1.13+

package main

import (
	"errors"
	"fmt"
	"strconv"
)

// TODO: Custom error types for wrapping examples
type DatabaseError struct {
	Query string
	Err   error
}

func (e DatabaseError) Error() string {
	return fmt.Sprintf("database error executing query '%s': %v", e.Query, e.Err)
}

func (e DatabaseError) Unwrap() error {
	return e.Err
}

type ServiceError struct {
	Service string
	Err     error
}

func (e ServiceError) Error() string {
	return fmt.Sprintf("service '%s' error: %v", e.Service, e.Err)
}

func (e ServiceError) Unwrap() error {
	return e.Err
}

// TODO: Sentinel errors for demonstration
var (
	ErrConnectionFailed = errors.New("connection failed")
	ErrTimeout         = errors.New("operation timeout")
	ErrUnauthorized    = errors.New("unauthorized access")
)

// TODO: Low-level functions that return basic errors
func connectToDatabase() error {
	// Simulate connection failure
	return ErrConnectionFailed
}

func executeQuery(query string) error {
	if query == "" {
		return errors.New("empty query")
	}
	if query == "SELECT * FROM timeout_table" {
		return ErrTimeout
	}
	
	// Simulate connection error
	err := connectToDatabase()
	if err != nil {
		// TODO: Wrap the connection error with context
		return fmt.Errorf("failed to execute query '%s': %w", query, err)
	}
	
	return nil
}

// TODO: Mid-level functions that add more context
func getUserByID(userID string) error {
	if userID == "" {
		return errors.New("user ID cannot be empty")
	}
	
	query := fmt.Sprintf("SELECT * FROM users WHERE id = '%s'", userID)
	err := executeQuery(query)
	if err != nil {
		// TODO: Wrap with DatabaseError
		return DatabaseError{
			Query: query,
			Err:   err,
		}
	}
	
	return nil
}

func authenticateUser(username, password string) error {
	if username == "" || password == "" {
		return ErrUnauthorized
	}
	
	err := getUserByID(username)
	if err != nil {
		// TODO: Wrap authentication error
		return fmt.Errorf("authentication failed for user '%s': %w", username, err)
	}
	
	return nil
}

// TODO: High-level service functions
func loginUser(username, password string) error {
	err := authenticateUser(username, password)
	if err != nil {
		// TODO: Wrap with ServiceError
		return ServiceError{
			Service: "UserService",
			Err:     err,
		}
	}
	
	return nil
}

// TODO: Functions to demonstrate error unwrapping
func analyzeError(err error) {
	fmt.Printf("Analyzing error: %v\n", err)
	
	// TODO: Use errors.Is to check for specific errors
	if errors.Is(err, ErrConnectionFailed) {
		fmt.Println("  -> Root cause: Connection failed")
	}
	if errors.Is(err, ErrTimeout) {
		fmt.Println("  -> Root cause: Timeout occurred")  
	}
	if errors.Is(err, ErrUnauthorized) {
		fmt.Println("  -> Root cause: Unauthorized access")
	}
	
	// TODO: Use errors.As to extract specific error types
	var dbErr DatabaseError
	if errors.As(err, &dbErr) {
		fmt.Printf("  -> Database error found: query='%s'\n", dbErr.Query)
	}
	
	var serviceErr ServiceError
	if errors.As(err, &serviceErr) {
		fmt.Printf("  -> Service error found: service='%s'\n", serviceErr.Service)
	}
	
	// TODO: Manual unwrapping chain
	fmt.Println("  -> Error chain:")
	currentErr := err
	level := 0
	for currentErr != nil {
		fmt.Printf("    Level %d: %v (type: %T)\n", level, currentErr, currentErr)
		currentErr = errors.Unwrap(currentErr)
		level++
		if level > 10 { // Prevent infinite loops
			break
		}
	}
}

// TODO: Function that demonstrates error joining (Go 1.20+)
func validateUserInput(username, email, ageStr string) error {
	var errs []error
	
	// TODO: Collect multiple validation errors
	if username == "" {
		errs = append(errs, errors.New("username is required"))
	}
	if len(username) < 3 {
		errs = append(errs, errors.New("username must be at least 3 characters"))
	}
	
	if email == "" {
		errs = append(errs, errors.New("email is required"))
	}
	if !contains(email, "@") {
		errs = append(errs, errors.New("email must contain @ symbol"))
	}
	
	if ageStr == "" {
		errs = append(errs, errors.New("age is required"))
	} else {
		if age, err := strconv.Atoi(ageStr); err != nil {
			errs = append(errs, fmt.Errorf("age must be a number: %w", err))
		} else if age < 0 {
			errs = append(errs, errors.New("age cannot be negative"))
		}
	}
	
	// TODO: Join all errors if any exist
	if len(errs) > 0 {
		// For Go 1.20+: return errors.Join(errs...)
		// For earlier versions, create custom multi-error
		return createMultiError(errs)
	}
	
	return nil
}

// TODO: Custom multi-error type for pre-Go 1.20
type MultiError struct {
	Errors []error
}

func (e MultiError) Error() string {
	if len(e.Errors) == 0 {
		return "no errors"
	}
	if len(e.Errors) == 1 {
		return e.Errors[0].Error()
	}
	
	result := fmt.Sprintf("multiple errors (%d):", len(e.Errors))
	for i, err := range e.Errors {
		result += fmt.Sprintf("\n  %d: %v", i+1, err)
	}
	return result
}

func createMultiError(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	return MultiError{Errors: errs}
}

// Helper function
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("=== Basic Error Wrapping ===")
	
	// TODO: Test simple error wrapping
	queries := []string{"SELECT * FROM users", "", "SELECT * FROM timeout_table"}
	
	for _, query := range queries {
		fmt.Printf("Executing query: '%s'\n", query)
		err := executeQuery(query)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			analyzeError(err)
		} else {
			fmt.Println("Query executed successfully")
		}
		fmt.Println()
	}
	
	fmt.Println("=== Multi-Level Error Wrapping ===")
	
	// TODO: Test authentication flow with multiple wrapping levels
	testUsers := []struct {
		username, password string
	}{
		{"alice", "secret123"},
		{"", "password"},
		{"bob", ""},
		{"timeout_user", "pass"},
	}
	
	for _, user := range testUsers {
		fmt.Printf("Attempting login: username='%s', password='%s'\n", user.username, user.password)
		err := loginUser(user.username, user.password)
		if err != nil {
			analyzeError(err)
		} else {
			fmt.Println("Login successful!")
		}
		fmt.Println()
	}
	
	fmt.Println("=== Error Chain Analysis ===")
	
	// TODO: Create a deeply nested error
	originalErr := errors.New("original failure")
	level1 := fmt.Errorf("level 1 wrapper: %w", originalErr)
	level2 := fmt.Errorf("level 2 wrapper: %w", level1)
	level3 := fmt.Errorf("level 3 wrapper: %w", level2)
	
	fmt.Println("Deep error chain analysis:")
	analyzeError(level3)
	
	fmt.Println("\n=== Multiple Error Validation ===")
	
	// TODO: Test input validation with multiple errors
	testInputs := []struct {
		username, email, age string
	}{
		{"alice", "alice@example.com", "25"},
		{"", "", ""},
		{"ab", "invalid-email", "-5"},
		{"validuser", "user@domain.com", "not-a-number"},
	}
	
	for i, input := range testInputs {
		fmt.Printf("Validating input %d: username='%s', email='%s', age='%s'\n", 
			i+1, input.username, input.email, input.age)
		
		err := validateUserInput(input.username, input.email, input.age)
		if err != nil {
			fmt.Printf("Validation errors:\n%v\n", err)
			
			// TODO: Check if it's a MultiError and process individual errors
			if multiErr, ok := err.(MultiError); ok {
				fmt.Printf("Found %d individual errors:\n", len(multiErr.Errors))
				for j, individualErr := range multiErr.Errors {
					fmt.Printf("  %d: %v\n", j+1, individualErr)
				}
			}
		} else {
			fmt.Println("All input is valid!")
		}
		fmt.Println()
	}
	
	fmt.Println("=== Error Comparison ===")
	
	// TODO: Demonstrate the difference between wrapping and not wrapping
	originalError := errors.New("file not found")
	
	// Without wrapping (loses context)
	withoutWrap := fmt.Errorf("failed to read config: %v", originalError)
	
	// With wrapping (preserves error chain)
	withWrap := fmt.Errorf("failed to read config: %w", originalError)
	
	fmt.Println("Without wrapping:")
	fmt.Printf("  Error: %v\n", withoutWrap)
	fmt.Printf("  Is original error? %t\n", errors.Is(withoutWrap, originalError))
	
	fmt.Println("\nWith wrapping:")
	fmt.Printf("  Error: %v\n", withWrap)
	fmt.Printf("  Is original error? %t\n", errors.Is(withWrap, originalError))
	fmt.Printf("  Unwrapped: %v\n", errors.Unwrap(withWrap))
}