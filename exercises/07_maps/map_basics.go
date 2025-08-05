// map_basics.go
// Learn the fundamentals of maps in Go
// Maps are key-value pairs similar to hash tables or dictionaries

package main

import "fmt"

func main() {
	// TODO: Declare a map using make()
	// Hint: make(map[keyType]valueType)
	ages := // Create a map from string to int
	
	// TODO: Initialize a map with values using map literal
	// Hint: map[keyType]valueType{key: value, key: value}
	colors := // Create a map with color names as keys and hex codes as values
	
	// TODO: Add elements to the ages map
	ages["Alice"] = 25
	ages["Bob"] = 30
	// Add "Charlie" with age 35
	
	// TODO: Access a value from the map
	aliceAge := // Get Alice's age from the map
	
	// TODO: Check if a key exists using the comma ok idiom
	// Hint: value, ok := map[key]
	bobAge, exists := // Check if Bob exists in ages map
	
	if exists {
		fmt.Printf("Bob's age: %d\n", bobAge)
	}
	
	// TODO: Try to access a non-existent key
	// Hint: Non-existent keys return zero value
	unknown := // Try to get "Unknown" from ages map
	fmt.Printf("Unknown person's age (zero value): %d\n", unknown)
	
	// TODO: Delete a key from the map
	// Hint: delete(map, key)
	// Delete "Alice" from ages map
	
	// TODO: Get the length of the map
	length := // Get length of ages map
	
	// TODO: Create a nil map and try to assign (this will panic if uncommented)
	var nilMap map[string]int
	// nilMap["test"] = 1  // This would panic!
	
	// TODO: Check if a map is nil
	if nilMap == nil {
		fmt.Println("nilMap is nil - cannot assign to it")
	}
	
	// Print results
	fmt.Println("Ages map:", ages)
	fmt.Println("Colors map:", colors)
	fmt.Printf("Alice's age: %d\n", aliceAge)
	fmt.Printf("Map length: %d\n", length)
}