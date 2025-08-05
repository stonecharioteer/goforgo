// map_basics.go - SOLUTION
// Learn the fundamentals of maps in Go
// Maps are key-value pairs similar to hash tables or dictionaries

package main

import "fmt"

func main() {
	// Declare a map using make()
	// Hint: make(map[keyType]valueType)
	ages := make(map[string]int)
	
	// Initialize a map with values using map literal
	// Hint: map[keyType]valueType{key: value, key: value}
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
		"white": "#FFFFFF",
		"black": "#000000",
	}
	
	// Add elements to the ages map
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 35
	
	// Access a value from the map
	aliceAge := ages["Alice"]
	
	// Check if a key exists using the comma ok idiom
	// Hint: value, ok := map[key]
	bobAge, exists := ages["Bob"]
	
	if exists {
		fmt.Printf("Bob's age: %d\n", bobAge)
	}
	
	// Try to access a non-existent key
	// Hint: Non-existent keys return zero value
	unknown := ages["Unknown"]
	fmt.Printf("Unknown person's age (zero value): %d\n", unknown)
	
	// Delete a key from the map
	// Hint: delete(map, key)
	delete(ages, "Alice")
	
	// Get the length of the map
	length := len(ages)
	
	// Create a nil map and try to assign (this will panic if uncommented)
	var nilMap map[string]int
	// nilMap["test"] = 1  // This would panic!
	
	// Check if a map is nil
	if nilMap == nil {
		fmt.Println("nilMap is nil - cannot assign to it")
	}
	
	// Print results
	fmt.Println("Ages map:", ages)
	fmt.Println("Colors map:", colors)
	fmt.Printf("Alice's age: %d\n", aliceAge)
	fmt.Printf("Map length: %d\n", length)
}