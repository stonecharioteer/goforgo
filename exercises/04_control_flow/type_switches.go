package main

import "fmt"

// TODO: Define a function called 'describeType' that takes interface{} parameter
// Use a type switch to handle different types:
// - string: print "It's a string with length:", len(value)
// - int: print "It's an integer with value:", value
// - bool: print "It's a boolean with value:", value
// - []int: print "It's an int slice with length:", len(value)
// - default: print "Unknown type"

// TODO: Define a function called 'processValue' that takes interface{} parameter
// Use type switch with variable assignment: switch v := value.(type)
// Handle these types:
// - string: if len(v) > 5, print "Long string:", v, else print "Short string:", v
// - int: if v > 0, print "Positive:", v, else print "Non-positive:", v
// - float64: print "Float with 2 decimal places:", formatted to 2 decimal places
// - default: print "Cannot process type:", reflect the type somehow

// TODO: Define a function called 'handleError' that takes interface{} parameter
// This simulates error handling with different error types
// Handle these cases:
// - string: print "String error:", value
// - error: print "Error type:", value.Error()
// - nil: print "No error"
// - default: print "Unknown error type"

func main() {
	// TODO: Test describeType with different values:
	// - "hello world"
	// - 42
	// - true
	// - []int{1, 2, 3, 4, 5}
	// - 3.14
	
	// TODO: Test processValue with different values:
	// - "short"
	// - "this is a long string"
	// - -10
	// - 25
	// - 3.14159
	// - []string{"not", "handled"}
	
	// TODO: Test handleError with different values:
	// - "file not found"
	// - fmt.Errorf("custom error")
	// - nil
	// - 123 (unknown type)
	
	// TODO: Demonstrate type switch with multiple types in one case
	// Create a function inline that handles int, int32, int64 in one case
	// Test with: 10, int32(20), int64(30)
}