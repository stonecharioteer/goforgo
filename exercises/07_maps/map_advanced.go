// map_advanced.go
// Learn advanced map operations and patterns

package main

import "fmt"

func main() {
	// TODO: Create a map of slices
	// This represents students and their test scores
	studentScores := // Create map[string][]int
	
	// TODO: Add scores for students
	studentScores["Alice"] = []int{92, 88, 95}
	studentScores["Bob"] = []int{78, 82, 85}
	// Add scores for "Charlie": [90, 87, 93]
	
	// TODO: Create a nested map (map of maps)
	// This represents a gradebook: subject -> student -> grade
	gradebook := // Create map[string]map[string]int
	
	// TODO: Initialize nested maps
	gradebook["Math"] = make(map[string]int)
	gradebook["Science"] = make(map[string]int)
	
	// Add some grades
	gradebook["Math"]["Alice"] = 92
	gradebook["Math"]["Bob"] = 85
	gradebook["Science"]["Alice"] = 89
	// Add Science grade for Bob: 88
	
	// TODO: Access nested map values safely
	// Check if Math grades exist, then check if Alice has a Math grade
	mathGrades, mathExists := // Check if "Math" exists in gradebook
	if mathExists {
		aliceMath, aliceExists := // Check if "Alice" exists in Math grades
		if aliceExists {
			fmt.Printf("Alice's Math grade: %d\n", aliceMath)
		}
	}
	
	// TODO: Create a map with struct values
	type Person struct {
		Age  int
		City string
	}
	
	people := // Create map[string]Person
	
	// Add people
	people["John"] = Person{Age: 30, City: "New York"}
	// Add "Jane": Age 25, City "Los Angeles"
	
	// TODO: Modify struct in map (this creates a copy, doesn't modify original)
	john := people["John"]
	john.Age = 31
	people["John"] = john // Need to reassign
	
	// TODO: Use map as a set (values don't matter, only keys)
	uniqueWords := make(map[string]bool)
	words := []string{"hello", "world", "hello", "go", "world"}
	
	// Add words to set
	for _, word := range words {
		uniqueWords[word] = true
	}
	
	// TODO: Get unique words from the set
	var unique []string
	// Extract keys from the map
	
	// TODO: Merge two maps
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"b": 3, "c": 4} // Note: "b" exists in both
	
	merged := make(map[string]int)
	// Copy map1 to merged
	// Copy map2 to merged (this will overwrite "b")
	
	// Print results
	fmt.Println("Student scores:", studentScores)
	fmt.Println("Gradebook:", gradebook)
	fmt.Println("People:", people)
	fmt.Println("Unique words:", unique)
	fmt.Println("Merged maps:", merged)
}