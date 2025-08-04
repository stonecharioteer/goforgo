package main

import (
	"fmt"
	"reflect"
)

// Define a function called 'describeType' that takes interface{} parameter
// Use a type switch to handle different types:
// - string: print "It's a string with length:", len(value)
// - int: print "It's an integer with value:", value
// - bool: print "It's a boolean with value:", value
// - []int: print "It's an int slice with length:", len(value)
// - default: print "Unknown type"
func describeType(value interface{}) {
	switch value.(type) {
	case string:
		v := value.(string)
		fmt.Println("It's a string with length:", len(v))
	case int:
		v := value.(int)
		fmt.Println("It's an integer with value:", v)
	case bool:
		v := value.(bool)
		fmt.Println("It's a boolean with value:", v)
	case []int:
		v := value.([]int)
		fmt.Println("It's an int slice with length:", len(v))
	default:
		fmt.Println("Unknown type")
	}
}

// Define a function called 'processValue' that takes interface{} parameter
// Use type switch with variable assignment: switch v := value.(type)
// Handle these types:
// - string: if len(v) > 5, print "Long string:", v, else print "Short string:", v
// - int: if v > 0, print "Positive:", v, else print "Non-positive:", v
// - float64: print "Float with 2 decimal places:", formatted to 2 decimal places
// - default: print "Cannot process type:", reflect the type somehow
func processValue(value interface{}) {
	switch v := value.(type) {
	case string:
		if len(v) > 5 {
			fmt.Println("Long string:", v)
		} else {
			fmt.Println("Short string:", v)
		}
	case int:
		if v > 0 {
			fmt.Println("Positive:", v)
		} else {
			fmt.Println("Non-positive:", v)
		}
	case float64:
		fmt.Printf("Float with 2 decimal places: %.2f\n", v)
	default:
		fmt.Println("Cannot process type:", reflect.TypeOf(value))
	}
}

// Define a function called 'handleError' that takes interface{} parameter
// This simulates error handling with different error types
// Handle these cases:
// - string: print "String error:", value
// - error: print "Error type:", value.Error()
// - nil: print "No error"
// - default: print "Unknown error type"
func handleError(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Println("String error:", v)
	case error:
		fmt.Println("Error type:", v.Error())
	case nil:
		fmt.Println("No error")
	default:
		fmt.Println("Unknown error type")
	}
}

func main() {
	// Test describeType with different values:
	// - "hello world"
	// - 42
	// - true
	// - []int{1, 2, 3, 4, 5}
	// - 3.14
	fmt.Println("=== describeType tests ===")
	describeType("hello world")
	describeType(42)
	describeType(true)
	describeType([]int{1, 2, 3, 4, 5})
	describeType(3.14)
	
	// Test processValue with different values:
	// - "short"
	// - "this is a long string"
	// - -10
	// - 25
	// - 3.14159
	// - []string{"not", "handled"}
	fmt.Println("\n=== processValue tests ===")
	processValue("short")
	processValue("this is a long string")
	processValue(-10)
	processValue(25)
	processValue(3.14159)
	processValue([]string{"not", "handled"})
	
	// Test handleError with different values:
	// - "file not found"
	// - fmt.Errorf("custom error")
	// - nil
	// - 123 (unknown type)
	fmt.Println("\n=== handleError tests ===")
	handleError("file not found")
	handleError(fmt.Errorf("custom error"))
	handleError(nil)
	handleError(123)
	
	// Demonstrate type switch with multiple types in one case
	// Create a function inline that handles int, int32, int64 in one case
	// Test with: 10, int32(20), int64(30)
	fmt.Println("\n=== Multiple types in one case ===")
	handleInteger := func(value interface{}) {
		switch v := value.(type) {
		case int, int32, int64:
			fmt.Printf("Integer type value: %v (type: %T)\n", v, v)
		default:
			fmt.Println("Not an integer type")
		}
	}
	
	handleInteger(10)
	handleInteger(int32(20))
	handleInteger(int64(30))
	handleInteger("not an integer")
}