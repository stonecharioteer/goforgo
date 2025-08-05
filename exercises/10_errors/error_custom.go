// error_custom.go
// Learn how to create custom error types and implement the error interface

package main

import (
	"fmt"
	"strconv"
	"time"
)

// TODO: Define a simple custom error type
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

// TODO: Implement the error interface
func (e ValidationError) Error() string {
	// Return formatted error message
}

// TODO: Define a more complex error type with multiple fields
type NetworkError struct {
	Operation string
	URL       string
	Timestamp time.Time
	Err       error // Wrapped error
}

func (e NetworkError) Error() string {
	// Return detailed error message with timestamp
}

// TODO: Implement error unwrapping
func (e NetworkError) Unwrap() error {
	// Return the wrapped error
}

// TODO: Define sentinel errors (package-level error variables)
var (
	ErrInsufficientFunds = // Create error for insufficient funds
	ErrAccountNotFound   = // Create error for account not found  
	ErrInvalidAmount     = // Create error for invalid amount
)

// TODO: Custom error with additional methods
type AccountError struct {
	AccountID string
	Balance   float64
	Operation string
	Code      int
}

func (e AccountError) Error() string {
	return fmt.Sprintf("account error [%d]: %s failed for account %s (balance: %.2f)", 
		e.Code, e.Operation, e.AccountID, e.Balance)
}

// TODO: Add method to check error severity
func (e AccountError) IsCritical() bool {
	// Return true for critical error codes (>= 500)
}

// TODO: Add method to get user-friendly message
func (e AccountError) UserMessage() string {
	switch e.Code {
	case 404:
		return "Account not found"
	case 400:
		return "Invalid request"
	case 500:
		return "Internal server error"
	default:
		return "An error occurred"
	}
}

// TODO: Functions that return custom errors
func validateAge(age int) error {
	if age < 0 {
		// Return ValidationError
	}
	if age > 150 {
		// Return ValidationError
	}
	return nil
}

func validateEmail(email string) error {
	if email == "" {
		// Return ValidationError
	}
	if !contains(email, "@") {
		// Return ValidationError  
	}
	return nil
}

func fetchUserData(url string) error {
	// Simulate network operation that might fail
	if url == "" {
		// Return NetworkError with wrapped error
	}
	if url == "timeout.com" {
		// Return NetworkError simulating timeout
	}
	return nil
}

func withdraw(accountID string, amount float64) error {
	// Simulate account operations
	balance := 100.0 // Mock balance
	
	if accountID == "" {
		return ErrAccountNotFound
	}
	if amount <= 0 {
		return ErrInvalidAmount
	}
	if amount > balance {
		return ErrInsufficientFunds
	}
	
	// Simulate system error
	if amount > 1000 {
		return AccountError{
			AccountID: accountID,
			Balance:   balance,
			Operation: "withdraw",
			Code:      500,
		}
	}
	
	return nil
}

// TODO: Helper function for string contains (since we're learning)
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("=== Custom Error Types ===")
	
	// TODO: Test validation errors
	testAges := []int{25, -5, 200, 0}
	for _, age := range testAges {
		err := validateAge(age)
		if err != nil {
			fmt.Printf("Age validation error: %v\n", err)
			// TODO: Type assert to ValidationError for additional info
			if valErr, ok := /* type assert err to ValidationError */; ok {
				fmt.Printf("  Field: %s, Value: %v\n", valErr.Field, valErr.Value)
			}
		} else {
			fmt.Printf("Age %d is valid\n", age)
		}
	}
	
	fmt.Println("\n=== Email Validation ===")
	
	testEmails := []string{"user@example.com", "", "invalid-email", "test@domain.org"}
	for _, email := range testEmails {
		err := validateEmail(email)
		if err != nil {
			fmt.Printf("Email validation error: %v\n", err)
		} else {
			fmt.Printf("Email '%s' is valid\n", email)
		}
	}
	
	fmt.Println("\n=== NetworkError with Wrapping ===")
	
	testURLs := []string{"https://api.example.com", "", "timeout.com", "https://valid.com"}
	for _, url := range testURLs {
		err := fetchUserData(url)
		if err != nil {
			fmt.Printf("Network error: %v\n", err)
			
			// TODO: Type assert to NetworkError
			if netErr, ok := /* type assert err to NetworkError */; ok {
				fmt.Printf("  Operation: %s, URL: %s, Time: %s\n", 
					netErr.Operation, netErr.URL, netErr.Timestamp.Format("15:04:05"))
				
				// TODO: Check for wrapped error
				if wrappedErr := /* unwrap the error */; wrappedErr != nil {
					fmt.Printf("  Wrapped error: %v\n", wrappedErr)
				}
			}
		} else {
			fmt.Printf("Successfully fetched data from %s\n", url)
		}
	}
	
	fmt.Println("\n=== Sentinel Errors ===")
	
	// TODO: Test account operations with sentinel errors
	testOperations := []struct {
		accountID string
		amount    float64
	}{
		{"ACC123", 50.0},
		{"", 25.0},
		{"ACC456", -10.0},
		{"ACC789", 150.0},
		{"ACC999", 2000.0},
	}
	
	for _, op := range testOperations {
		err := withdraw(op.accountID, op.amount)
		if err != nil {
			fmt.Printf("Withdraw error: %v\n", err)
			
			// TODO: Check for specific sentinel errors
			switch err {
			case ErrAccountNotFound:
				fmt.Println("  -> Please check the account ID")
			case ErrInvalidAmount:
				fmt.Println("  -> Amount must be positive")
			case ErrInsufficientFunds:
				fmt.Println("  -> Please add funds to your account")
			default:
				// TODO: Check for AccountError type
				if accErr, ok := /* type assert to AccountError */; ok {
					fmt.Printf("  -> User message: %s\n", accErr.UserMessage())
					fmt.Printf("  -> Critical: %t\n", accErr.IsCritical())
				}
			}
		} else {
			fmt.Printf("Successfully withdrew %.2f from %s\n", op.amount, op.accountID)
		}
	}
	
	fmt.Println("\n=== Error Type Checking ===")
	
	// TODO: Create different types of errors and demonstrate type checking
	errors := []error{
		ValidationError{Field: "name", Value: "", Message: "name is required"},
		NetworkError{Operation: "GET", URL: "api.test.com", Timestamp: time.Now()},
		ErrInsufficientFunds,
		AccountError{AccountID: "TEST", Balance: 50.0, Operation: "transfer", Code: 400},
		fmt.Errorf("generic error"),
	}
	
	for i, err := range errors {
		fmt.Printf("Error %d: %v\n", i+1, err)
		
		// TODO: Use type switch to handle different error types
		switch e := err.(type) {
		case ValidationError:
			fmt.Printf("  -> Validation error on field: %s\n", e.Field)
		case NetworkError:
			fmt.Printf("  -> Network error during: %s\n", e.Operation)
		case AccountError:
			fmt.Printf("  -> Account error (code %d): %s\n", e.Code, e.UserMessage())
		default:
			fmt.Printf("  -> Generic error type: %T\n", e)
		}
		fmt.Println()
	}
}