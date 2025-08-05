// slice_sorting_custom.go
// Learn custom sorting with slices using sort.Slice and sort.SliceStable
// Practice sorting complex data structures and custom comparison functions

package main

import (
	"fmt"
	// TODO: Import the sort package
)

// TODO: Define a Person struct with Name, Age, and Salary fields
type Person struct {
	// Complete the struct definition
}

func main() {
	// TODO: Create a slice of Person structs
	people := []Person{
		// Add 5 people with different names, ages, and salaries
	}

	fmt.Println("Original people:")
	printPeople(people)

	// TODO: Sort by age (ascending) using sort.Slice
	// Hint: sort.Slice(slice, func(i, j int) bool { return condition })
	
	fmt.Println("\nSorted by age (ascending):")
	printPeople(people)

	// TODO: Sort by salary (descending) using sort.Slice
	
	fmt.Println("\nSorted by salary (descending):")
	printPeople(people)

	// TODO: Sort by name (alphabetical) using sort.Slice
	
	fmt.Println("\nSorted by name (alphabetical):")
	printPeople(people)

	// TODO: Multi-level sorting: first by age, then by salary if ages are equal
	// Use sort.SliceStable for stable sorting
	
	fmt.Println("\nSorted by age, then salary (stable sort):")
	printPeople(people)

	// TODO: Custom sorting with complex conditions
	// Sort by: under 30s first, then by salary descending
	
	fmt.Println("\nCustom sort (under 30s first, then by salary desc):")
	printPeople(people)

	// TODO: Sort a slice of strings by length, then alphabetically
	words := []string{"apple", "pie", "banana", "cat", "elephant", "dog", "a"}
	fmt.Printf("\nOriginal words: %v\n", words)
	
	// Sort by length first, then alphabetically for same length
	
	fmt.Printf("Sorted by length, then alphabetically: %v\n", words)

	// TODO: Sort a slice of integers by absolute value
	numbers := []int{-5, 3, -1, 8, -10, 2, -3}
	fmt.Printf("\nOriginal numbers: %v\n", numbers)
	
	// Sort by absolute value
	
	fmt.Printf("Sorted by absolute value: %v\n", numbers)

	// TODO: Check if a slice is sorted with custom comparison
	ages := []int{25, 30, 35, 40, 45}
	isSorted := // Use sort.SliceIsSorted with custom function
	fmt.Printf("\nAges %v is sorted: %t\n", ages, isSorted)
}

// Helper function to print people slice
func printPeople(people []Person) {
	for _, p := range people {
		fmt.Printf("  %s (Age: %d, Salary: $%d)\n", p.Name, p.Age, p.Salary)
	}
}

// TODO: Helper function to get absolute value
func abs(x int) int {
	// Return absolute value of x
}