// map_iteration.go - SOLUTION
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
	
	// Iterate over map with both keys and values
	// Hint: for key, value := range map { ... }
	fmt.Println("Student grades:")
	for name, grade := range grades {
		fmt.Printf("  %s: %d\n", name, grade)
	}
	
	// Iterate over map keys only
	// Hint: for key := range map { ... }
	fmt.Println("\nStudent names (keys only):")
	for name := range grades {
		fmt.Printf("  %s\n", name)
	}
	
	// Iterate over map values only
	// Hint: for _, value := range map { ... }
	fmt.Println("\nGrade values only:")
	for _, grade := range grades {
		fmt.Printf("  %d\n", grade)
	}
	
	// Calculate the sum and average of grades
	var sum int
	var count int
	for _, grade := range grades {
		sum += grade
		count++
	}
	
	average := float64(sum) / float64(count)
	fmt.Printf("\nTotal sum: %d, Average: %.2f\n", sum, average)
	
	// Find the highest grade and student
	var maxGrade int
	var topStudent string
	for name, grade := range grades {
		if grade > maxGrade {
			maxGrade = grade
			topStudent = name
		}
	}
	
	fmt.Printf("Top student: %s with grade %d\n", topStudent, maxGrade)
	
	// Iterate in sorted order by keys
	// Note: Maps are unordered, so we need to sort keys separately
	var keys []string
	for product := range inventory {
		keys = append(keys, product)
	}
	
	// Sort the keys
	sort.Strings(keys)
	
	fmt.Println("\nInventory (sorted by product name):")
	for _, product := range keys {
		fmt.Printf("  %s: %d\n", product, inventory[product])
	}
	
	// Count items with specific criteria
	lowStock := 0
	threshold := 30
	for _, stock := range inventory {
		if stock <= threshold {
			lowStock++
		}
	}
	
	fmt.Printf("\nProducts with low stock (<= %d): %d\n", threshold, lowStock)
}