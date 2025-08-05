// slice_tricks.go
// Learn advanced slice operations and common patterns

package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// TODO: Remove element at index 3 (value 4)
	// Use append to combine slices before and after the index
	index := 3
	// Hint: append(slice[:index], slice[index+1:]...)
	data = // Complete this removal operation
	
	fmt.Println("After removing index 3:", data)
	
	// TODO: Insert element 99 at index 2
	// Split the slice and insert the new element
	insertIndex := 2
	insertValue := 99
	// Hint: append(append(slice[:index], value), slice[index:]...)
	data = // Complete this insertion operation
	
	fmt.Println("After inserting 99 at index 2:", data)
	
	// TODO: Reverse a slice in place
	numbers := []int{1, 2, 3, 4, 5}
	// Use two pointers approach
	// Complete the reversal logic
	
	fmt.Println("Reversed numbers:", numbers)
	
	// TODO: Find and remove all occurrences of a value
	values := []int{1, 2, 3, 2, 4, 2, 5}
	target := 2
	
	// Filter out the target value
	var filtered []int
	// Write a loop to filter out target value
	
	fmt.Printf("After removing all %d: %v\n", target, filtered)
	
	// TODO: Check if slice contains a value
	haystack := []string{"apple", "banana", "cherry", "date"}
	needle := "cherry"
	
	var found bool
	// Write logic to check if needle exists in haystack
	
	if found {
		fmt.Printf("Found '%s' in slice\n", needle)
	} else {
		fmt.Printf("'%s' not found in slice\n", needle)
	}
	
	// TODO: Get unique elements from a slice
	duplicates := []int{1, 2, 2, 3, 1, 4, 3, 5}
	var unique []int
	
	// Use a map to track seen elements
	// Write logic to get unique elements
	
	fmt.Println("Unique elements:", unique)
}