// map_iteration.go
// Learn different ways to iterate over maps in Go

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Sample data
	grades := map[string]int{
		"Alice":   92,
		"Bob":     85,
		"Charlie": 78,
		"Diana":   96,
		"Eve":     88,
	}
	
	inventory := map[string]int{
		"apples":  50,
		"bananas": 30,
		"oranges": 25,
	}
	
	// TODO: Iterate over map with both keys and values
	// Hint: for key, value := range map { ... }
	fmt.Println("Student grades:")
	// Write your iteration loop here
	
	// TODO: Iterate over map keys only
	// Hint: for key := range map { ... }
	fmt.Println("\nStudent names (keys only):")
	// Write your iteration loop here
	
	// TODO: Iterate over map values only
	// Hint: for _, value := range map { ... }
	fmt.Println("\nGrade values only:")
	// Write your iteration loop here
	
	// TODO: Calculate the sum and average of grades
	var sum int
	var count int
	// Write loop to calculate sum and count
	
	average := float64(sum) / float64(count)
	fmt.Printf("\nTotal sum: %d, Average: %.2f\n", sum, average)
	
	// TODO: Find the highest grade and student
	var maxGrade int
	var topStudent string
	// Write loop to find maximum grade and corresponding student
	
	fmt.Printf("Top student: %s with grade %d\n", topStudent, maxGrade)
	
	// TODO: Iterate in sorted order by keys
	// Note: Maps are unordered, so we need to sort keys separately
	var keys []string
	// Extract keys into a slice
	
	// Sort the keys
	sort.Strings(keys)
	
	fmt.Println("\nInventory (sorted by product name):")
	// Iterate using sorted keys
	
	// TODO: Count items with specific criteria
	lowStock := 0
	threshold := 30
	// Count items with stock <= threshold
	
	fmt.Printf("\nProducts with low stock (<= %d): %d\n", threshold, lowStock)
}