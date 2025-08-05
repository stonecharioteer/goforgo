// slice_sorting_custom.go - SOLUTION
// Learn custom sorting with slices using sort.Slice and sort.SliceStable
// Practice sorting complex data structures and custom comparison functions

package main

import (
	"fmt"
	"sort"
	"strings"
)

// Define a Person struct with Name, Age, and Salary fields
type Person struct {
	Name   string
	Age    int
	Salary int
}

func main() {
	// Create a slice of Person structs
	people := []Person{
		{"Alice", 28, 75000},
		{"Bob", 32, 85000},
		{"Charlie", 25, 65000},
		{"Diana", 29, 90000},
		{"Eve", 25, 70000},
	}

	fmt.Println("Original people:")
	printPeople(people)

	// Sort by age (ascending) using sort.Slice
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	
	fmt.Println("\nSorted by age (ascending):")
	printPeople(people)

	// Sort by salary (descending) using sort.Slice
	sort.Slice(people, func(i, j int) bool {
		return people[i].Salary > people[j].Salary
	})
	
	fmt.Println("\nSorted by salary (descending):")
	printPeople(people)

	// Sort by name (alphabetical) using sort.Slice
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	
	fmt.Println("\nSorted by name (alphabetical):")
	printPeople(people)

	// Multi-level sorting: first by age, then by salary if ages are equal
	sort.SliceStable(people, func(i, j int) bool {
		if people[i].Age == people[j].Age {
			return people[i].Salary > people[j].Salary // Higher salary first for same age
		}
		return people[i].Age < people[j].Age
	})
	
	fmt.Println("\nSorted by age, then salary (stable sort):")
	printPeople(people)

	// Custom sorting with complex conditions
	// Sort by: under 30s first, then by salary descending
	sort.Slice(people, func(i, j int) bool {
		under30i := people[i].Age < 30
		under30j := people[j].Age < 30
		
		if under30i && !under30j {
			return true // i is under 30, j is not
		}
		if !under30i && under30j {
			return false // j is under 30, i is not
		}
		// Both in same age category, sort by salary descending
		return people[i].Salary > people[j].Salary
	})
	
	fmt.Println("\nCustom sort (under 30s first, then by salary desc):")
	printPeople(people)

	// Sort a slice of strings by length, then alphabetically
	words := []string{"apple", "pie", "banana", "cat", "elephant", "dog", "a"}
	fmt.Printf("\nOriginal words: %v\n", words)
	
	// Sort by length first, then alphabetically for same length
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return strings.ToLower(words[i]) < strings.ToLower(words[j])
		}
		return len(words[i]) < len(words[j])
	})
	
	fmt.Printf("Sorted by length, then alphabetically: %v\n", words)

	// Sort a slice of integers by absolute value
	numbers := []int{-5, 3, -1, 8, -10, 2, -3}
	fmt.Printf("\nOriginal numbers: %v\n", numbers)
	
	// Sort by absolute value
	sort.Slice(numbers, func(i, j int) bool {
		return abs(numbers[i]) < abs(numbers[j])
	})
	
	fmt.Printf("Sorted by absolute value: %v\n", numbers)

	// Check if a slice is sorted with custom comparison
	ages := []int{25, 30, 35, 40, 45}
	isSorted := sort.SliceIsSorted(ages, func(i, j int) bool {
		return ages[i] <= ages[j]
	})
	fmt.Printf("\nAges %v is sorted: %t\n", ages, isSorted)
}

// Helper function to print people slice
func printPeople(people []Person) {
	for _, p := range people {
		fmt.Printf("  %s (Age: %d, Salary: $%d)\n", p.Name, p.Age, p.Salary)
	}
}

// Helper function to get absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}