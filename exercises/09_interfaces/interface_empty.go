// interface_empty.go
// Learn about the empty interface and type assertions

package main

import "fmt"

// TODO: Function that accepts any type using empty interface
func printAnything(value interface{}) {
	// Print the value and its type
	fmt.Printf("Value: %v, Type: %T\n", value, value)
}

// TODO: Function that handles different types using type assertion
func handleValue(value interface{}) {
	// TODO: Use type assertion to check if value is a string
	if str, ok := /* type assert to string */; ok {
		fmt.Printf("It's a string: %s (length: %d)\n", str, len(str))
		return
	}
	
	// TODO: Use type assertion to check if value is an int
	if num, ok := /* type assert to int */; ok {
		fmt.Printf("It's an int: %d (doubled: %d)\n", num, num*2)
		return
	}
	
	// TODO: Use type assertion to check if value is a bool
	if b, ok := /* type assert to bool */; ok {
		fmt.Printf("It's a bool: %t (negated: %t)\n", b, !b)
		return
	}
	
	fmt.Printf("Unknown type: %T\n", value)
}

// TODO: Function using type switch
func processValue(value interface{}) {
	// Use type switch to handle different types
	switch v := value.(type) {
	case string:
		// Handle string case
		fmt.Printf("String processing: '%s' -> '%s'\n", v, /* convert to uppercase */)
	case int:
		// Handle int case  
		fmt.Printf("Int processing: %d -> %d\n", v, /* calculate square */)
	case float64:
		// Handle float64 case
		fmt.Printf("Float processing: %.2f -> %.2f\n", v, /* calculate square root */)
	case []int:
		// Handle slice of ints
		var sum int
		// Calculate sum of slice
		fmt.Printf("Slice processing: %v -> sum: %d\n", v, sum)
	case bool:
		// Handle bool case
		fmt.Printf("Bool processing: %t -> %s\n", v, /* convert to "yes"/"no" */)
	default:
		// Handle unknown types
		fmt.Printf("Unknown type processing: %T = %v\n", v, v)
	}
}

// TODO: Function that stores mixed types in slice
func demonstrateSliceOfInterface() {
	// Create slice of empty interface to store different types
	var mixed []interface{}
	
	// TODO: Add different types to the slice
	// Add: 42, "hello", 3.14, true, []int{1,2,3}
	
	fmt.Println("Mixed slice contents:")
	for i, item := range mixed {
		fmt.Printf("Index %d: %v (%T)\n", i, item, item)
	}
}

// TODO: Function that uses interface{} in map
func demonstrateMapOfInterface() {
	// Create map with string keys and interface{} values
	data := make(map[string]interface{})
	
	// TODO: Add different types of data
	// Add: "name" -> "Alice", "age" -> 30, "height" -> 5.8, "married" -> true
	
	fmt.Println("Map with mixed value types:")
	for key, value := range data {
		fmt.Printf("%s: %v (%T)\n", key, value, value)
	}
	
	// TODO: Safely extract values with type assertion
	if name, ok := /* get and assert "name" as string */; ok {
		fmt.Printf("Name is: %s\n", name)
	}
	
	if age, ok := /* get and assert "age" as int */; ok {
		fmt.Printf("Age is: %d\n", age)
	}
}

func main() {
	fmt.Println("=== Empty Interface Demo ===")
	
	// TODO: Test printAnything with different types
	// Test with: 42, "hello world", 3.14159, true, []int{1,2,3}
	
	fmt.Println("\n=== Type Assertion Demo ===")
	
	// TODO: Test handleValue with different types
	// Test with: "Go programming", 25, false, 3.14
	
	fmt.Println("\n=== Type Switch Demo ===")
	
	// TODO: Test processValue with different types
	// Test with: "hello", 5, 16.0, []int{1,2,3,4,5}, true, struct{}{}
	
	fmt.Println("\n=== Slice of Interface Demo ===")
	demonstrateSliceOfInterface()
	
	fmt.Println("\n=== Map of Interface Demo ===")
	demonstrateMapOfInterface()
	
	// TODO: Demonstrate potential panic with wrong type assertion
	fmt.Println("\n=== Safe vs Unsafe Type Assertion ===")
	var value interface{} = "hello"
	
	// Safe type assertion (with ok)
	if str, ok := value.(string); ok {
		fmt.Printf("Safe assertion: %s\n", str)
	}
	
	// Unsafe type assertion (would panic if wrong type)
	str := value.(string)
	fmt.Printf("Unsafe assertion: %s\n", str)
	
	// TODO: This would panic - uncomment to test
	// num := value.(int) // This would panic!
}