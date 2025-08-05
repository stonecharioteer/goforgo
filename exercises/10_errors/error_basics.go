// error_basics.go
// Learn the fundamentals of error handling in Go
// Go uses explicit error values rather than exceptions

package main

import (
	"errors"
	"fmt"
	"strconv"
)

// TODO: Function that returns an error
func divide(a, b float64) (float64, error) {
	// Check for division by zero and return appropriate error
	if /* check for zero */ {
		// Return zero value and error
	}
	// Return result and nil error
}

// TODO: Function that uses errors.New()
func validateAge(age int) error {
	if /* age is negative */ {
		// Return error using errors.New()
	}
	if /* age is too high (> 150) */ {
		// Return error using errors.New()
	}
	// Return nil if valid
}

// TODO: Function that returns multiple types of errors
func parseAndValidateAge(ageStr string) (int, error) {
	// First, try to parse the string to int
	age, err := // Parse ageStr to int
	if err != nil {
		// Return parsing error
	}
	
	// Then validate the parsed age
	err = // Validate the age
	if err != nil {
		// Return validation error
	}
	
	// Return age and nil error if all good
}

// TODO: Function that demonstrates error propagation
func processUserAge(ageStr string) {
	age, err := parseAndValidateAge(ageStr)
	if err != nil {
		// Handle the error - print it
		return
	}
	
	fmt.Printf("Valid age: %d\n", age)
}

// TODO: Function with named return values for cleaner error handling
func safeDivide(a, b float64) (result float64, err error) {
	if b == 0 {
		// Set error and return (result will be zero value)
		err = errors.New("division by zero")
		return
	}
	// Set result and return (err will be nil)
	result = a / b
	return
}

// TODO: Function that checks specific error types
func handleStringConversion(s string) {
	num, err := strconv.Atoi(s)
	if err != nil {
		// Check if it's a specific error type
		if /* check if err is strconv.NumError */ {
			fmt.Printf("Number parsing error: %v\n", err)
		} else {
			fmt.Printf("Other error: %v\n", err)
		}
		return
	}
	
	fmt.Printf("Converted number: %d\n", num)
}

func main() {
	fmt.Println("=== Basic Error Handling ===")
	
	// TODO: Test divide function
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", result)
	}
	
	// TODO: Test divide by zero
	result, err = // Try to divide by zero
	if /* check for error */ {
		fmt.Printf("Division by zero error: %v\n", err)
	}
	
	fmt.Println("\n=== Error Validation ===")
	
	// TODO: Test age validation
	ages := []int{25, -5, 200, 0, 45}
	for _, age := range ages {
		err := // Validate each age
		if err != nil {
			fmt.Printf("Age %d is invalid: %v\n", age, err)
		} else {
			fmt.Printf("Age %d is valid\n", age)
		}
	}
	
	fmt.Println("\n=== Error Propagation ===")
	
	// TODO: Test parseAndValidateAge with different inputs
	testInputs := []string{\"25\", \"-5\", \"abc\", \"200\", \"30\"}\n	for _, input := range testInputs {\n		fmt.Printf(\"Processing '%s': \", input)\n		// Process each input\n	}\n	\n	fmt.Println(\"\\n=== Named Return Values ===\\n\")\n	\n	// TODO: Test safeDivide\n	result, err = safeDivide(15, 3)\n	if err != nil {\n		fmt.Printf(\"Error: %v\\n\", err)\n	} else {\n		fmt.Printf(\"15 / 3 = %.2f\\n\", result)\n	}\n	\n	// Test with zero\n	result, err = // Try safeDivide with zero\n	if err != nil {\n		fmt.Printf(\"Safe divide error: %v\\n\", err)\n	}\n	\n	fmt.Println(\"\\n=== Error Type Checking ===\\n\")\n	\n	// TODO: Test string conversion with different inputs\n	testStrings := []string{\"123\", \"abc\", \"456\", \"12.34\"}\n	for _, s := range testStrings {\n		fmt.Printf(\"Converting '%s': \", s)\n		// Handle each string conversion\n	}\n	\n	fmt.Println(\"\\n=== Error Patterns ===\\n\")\n	\n	// TODO: Demonstrate ignoring errors (not recommended but sometimes needed)\n	num, _ := strconv.Atoi(\"123\") // Ignore error with blank identifier\n	fmt.Printf(\"Ignoring error result: %d\\n\", num)\n	\n	// TODO: Demonstrate multiple error checks\n	s1, s2 := \"10\", \"5\"\n	n1, err1 := strconv.Atoi(s1)\n	n2, err2 := strconv.Atoi(s2)\n	\n	if err1 != nil || err2 != nil {\n		fmt.Printf(\"Parsing errors: err1=%v, err2=%v\\n\", err1, err2)\n		return\n	}\n	\n	// Both parsing succeeded\n	result, err = divide(float64(n1), float64(n2))\n	if err != nil {\n		fmt.Printf(\"Division error: %v\\n\", err)\n	} else {\n		fmt.Printf(\"%s / %s = %.2f\\n\", s1, s2, result)\n	}\n}"}