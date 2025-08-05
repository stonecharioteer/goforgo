// slice_tricks.go - SOLUTION
// Learn advanced slice operations and common patterns

package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// Remove element at index 3 (value 4)
	// Use append to combine slices before and after the index
	index := 3
	// Hint: append(slice[:index], slice[index+1:]...)
	data = append(data[:index], data[index+1:]...)
	
	fmt.Println("After removing index 3:", data)
	
	// Insert element 99 at index 2
	// Split the slice and insert the new element
	insertIndex := 2
	insertValue := 99
	// Hint: append(append(slice[:index], value), slice[index:]...)
	data = append(append(data[:insertIndex], insertValue), data[insertIndex:]...)
	
	fmt.Println("After inserting 99 at index 2:", data)
	
	// Reverse a slice in place
	numbers := []int{1, 2, 3, 4, 5}
	// Use two pointers approach
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	
	fmt.Println("Reversed numbers:", numbers)
	
	// Find and remove all occurrences of a value
	values := []int{1, 2, 3, 2, 4, 2, 5}
	target := 2
	
	// Filter out the target value
	var filtered []int
	for _, v := range values {
		if v != target {
			filtered = append(filtered, v)
		}
	}
	
	fmt.Printf("After removing all %d: %v\n", target, filtered)
	
	// Check if slice contains a value
	haystack := []string{"apple", "banana", "cherry", "date"}
	needle := "cherry"
	
	var found bool
	for _, item := range haystack {
		if item == needle {
			found = true
			break
		}
	}
	
	if found {
		fmt.Printf("Found '%s' in slice\n", needle)
	} else {
		fmt.Printf("'%s' not found in slice\n", needle)
	}
	
	// Get unique elements from a slice
	duplicates := []int{1, 2, 2, 3, 1, 4, 3, 5}
	var unique []int
	
	// Use a map to track seen elements
	seen := make(map[int]bool)
	for _, v := range duplicates {
		if !seen[v] {
			seen[v] = true
			unique = append(unique, v)
		}
	}
	
	fmt.Println("Unique elements:", unique)
}