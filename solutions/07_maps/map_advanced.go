// map_advanced.go - SOLUTION
// Learn advanced map operations and patterns

package main

import "fmt"

func main() {
	// Create a map of slices
	// This represents students and their test scores
	studentScores := make(map[string][]int)
	
	// Add scores for students
	studentScores["Alice"] = []int{92, 88, 95}
	studentScores["Bob"] = []int{78, 82, 85}
	studentScores["Charlie"] = []int{90, 87, 93}
	
	// Create a nested map (map of maps)
	// This represents a gradebook: subject -> student -> grade
	gradebook := make(map[string]map[string]int)
	
	// Initialize nested maps
	gradebook["Math"] = make(map[string]int)
	gradebook["Science"] = make(map[string]int)
	
	// Add some grades
	gradebook["Math"]["Alice"] = 92
	gradebook["Math"]["Bob"] = 85
	gradebook["Science"]["Alice"] = 89
	gradebook["Science"]["Bob"] = 88
	
	// Access nested map values safely
	// Check if Math grades exist, then check if Alice has a Math grade
	mathGrades, mathExists := gradebook["Math"]
	if mathExists {
		aliceMath, aliceExists := mathGrades["Alice"]
		if aliceExists {
			fmt.Printf("Alice's Math grade: %d\n", aliceMath)
		}
	}
	
	// Create a map with struct values
	type Person struct {
		Age  int
		City string
	}
	
	people := make(map[string]Person)
	
	// Add people
	people["John"] = Person{Age: 30, City: "New York"}
	people["Jane"] = Person{Age: 25, City: "Los Angeles"}
	
	// Modify struct in map (this creates a copy, doesn't modify original)
	john := people["John"]
	john.Age = 31
	people["John"] = john // Need to reassign
	
	// Use map as a set (values don't matter, only keys)
	uniqueWords := make(map[string]bool)
	words := []string{"hello", "world", "hello", "go", "world"}
	
	// Add words to set
	for _, word := range words {
		uniqueWords[word] = true
	}
	
	// Get unique words from the set
	var unique []string
	for word := range uniqueWords {
		unique = append(unique, word)
	}
	
	// Merge two maps
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"b": 3, "c": 4} // Note: "b" exists in both
	
	merged := make(map[string]int)
	// Copy map1 to merged
	for k, v := range map1 {
		merged[k] = v
	}
	// Copy map2 to merged (this will overwrite "b")
	for k, v := range map2 {
		merged[k] = v
	}
	
	// Print results
	fmt.Println("Student scores:", studentScores)
	fmt.Println("Gradebook:", gradebook)
	fmt.Println("People:", people)
	fmt.Println("Unique words:", unique)
	fmt.Println("Merged maps:", merged)
}