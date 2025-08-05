// interface_empty.go - SOLUTION
// Learn about the empty interface and type assertions

package main

import (
	"fmt"
	"math"
	"strings"
)

// Function that accepts any type using empty interface
func printAnything(value interface{}) {
	// Print the value and its type
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

// Function that handles different types using type assertion
func handleValue(value interface{}) {
	// Use type assertion to check if value is a string
	if str, ok := value.(string); ok {
		fmt.Printf("It's a string: %s (length: %d)\n", str, len(str))
		return
	}
	
	// Use type assertion to check if value is an int
	if num, ok := value.(int); ok {
		fmt.Printf("It's an int: %d (doubled: %d)\n", num, num*2)
		return
	}
	
	// Use type assertion to check if value is a bool
	if b, ok := value.(bool); ok {
		fmt.Printf("It's a bool: %t (negated: %t)\n", b, !b)
		return
	}
	
	fmt.Printf("Unknown type: %T\n", value)
}

// Function using type switch
func processValue(value interface{}) {
	// Use type switch to handle different types
	switch v := value.(type) {
	case string:
		// Handle string case
		fmt.Printf("String processing: '%s' -> '%s'\n", v, strings.ToUpper(v))
	case int:
		// Handle int case  
		fmt.Printf("Int processing: %d -> %d\n", v, v*v)
	case float64:
		// Handle float64 case
		fmt.Printf("Float processing: %.2f -> %.2f\n", v, math.Sqrt(v))
	case []int:
		// Handle slice of ints
		var sum int
		// Calculate sum of slice
		for _, num := range v {
			sum += num
		}
		fmt.Printf("Slice processing: %v -> sum: %d\n", v, sum)
	case bool:
		// Handle bool case
		var result string
		if v {
			result = "yes"
		} else {
			result = "no"
		}
		fmt.Printf("Bool processing: %t -> %s\n", v, result)
	default:
		// Handle unknown types
		fmt.Printf("Unknown type processing: %T = %v\n", v, v)
	}
}

// Function that stores mixed types in slice
func demonstrateSliceOfInterface() {
	// Create slice of empty interface to store different types
	var mixed []interface{}
	
	// Add different types to the slice
	mixed = append(mixed, 42)
	mixed = append(mixed, "hello")
	mixed = append(mixed, 3.14)
	mixed = append(mixed, true)
	mixed = append(mixed, []int{1, 2, 3})
	
	fmt.Println("Mixed slice contents:")
	for i, item := range mixed {
		fmt.Printf("Index %d: %v (%T)\n", i, item, item)
	}
}

// Function that uses interface{} in map
func demonstrateMapOfInterface() {
	// Create map with string keys and interface{} values
	data := make(map[string]interface{})
	
	// Add different types of data
	data["name"] = "Alice"
	data["age"] = 30
	data["height"] = 5.8
	data["married"] = true
	
	fmt.Println("Map with mixed value types:")
	for key, value := range data {
		fmt.Printf("%s: %v (%T)\n", key, value, value)
	}
	
	// Safely extract values with type assertion
	if name, ok := data["name"].(string); ok {
		fmt.Printf("Name is: %s\n", name)
	}
	
	if age, ok := data["age"].(int); ok {
		fmt.Printf("Age is: %d\n", age)
	}
}

func main() {
	fmt.Println("=== Empty Interface Demo ===")
	
	// Test printAnything with different types
	printAnything(42)
	printAnything("hello world")
	printAnything(3.14159)
	printAnything(true)
	printAnything([]int{1, 2, 3})
	
	fmt.Println("\n=== Type Assertion Demo ===")
	
	// Test handleValue with different types
	handleValue("Go programming")
	handleValue(25)
	handleValue(false)
	handleValue(3.14)
	
	fmt.Println("\n=== Type Switch Demo ===")
	
	// Test processValue with different types
	processValue("hello")
	processValue(5)
	processValue(16.0)
	processValue([]int{1, 2, 3, 4, 5})
	processValue(true)
	processValue(struct{}{})
	
	fmt.Println("\n=== Slice of Interface Demo ===")
	demonstrateSliceOfInterface()
	
	fmt.Println("\n=== Map of Interface Demo ===")
	demonstrateMapOfInterface()
	
	// Demonstrate potential panic with wrong type assertion
	fmt.Println("\n=== Safe vs Unsafe Type Assertion ===")
	var value interface{} = "hello"
	
	// Safe type assertion (with ok)
	if str, ok := value.(string); ok {
		fmt.Printf("Safe assertion: %s\n", str)
	}
	
	// Unsafe type assertion (would panic if wrong type)
	str := value.(string)
	fmt.Printf("Unsafe assertion: %s\n", str)
	
	// This would panic - uncomment to test
	// num := value.(int) // This would panic!
	fmt.Println("Note: Unsafe assertion with wrong type would cause panic")
}