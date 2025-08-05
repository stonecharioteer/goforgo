// slice_copy.go - SOLUTION
// Learn how to copy slices and understand slice references

package main

import "fmt"

func main() {
	original := []int{1, 2, 3, 4, 5}
	
	// Create a shallow copy by slicing
	// This creates a new slice header but shares the underlying array
	shallow := original[:]
	
	// Modify the shallow copy and observe the effect on original
	shallow[0] = 999
	fmt.Println("After modifying shallow copy:")
	fmt.Println("Original:", original)
	fmt.Println("Shallow:", shallow)
	
	// Reset original
	original[0] = 1
	
	// Create a deep copy using copy() function
	// First create a slice with same length
	deep := make([]int, len(original))
	
	// Use copy() to copy elements
	// Hint: copy(destination, source)
	copy(deep, original)
	
	// Modify the deep copy and observe no effect on original
	deep[0] = 888
	fmt.Println("\nAfter modifying deep copy:")
	fmt.Println("Original:", original)
	fmt.Println("Deep:", deep)
	
	// Copy partial slice
	source := []string{"a", "b", "c", "d", "e"}
	destination := make([]string, 3)
	
	// Copy only first 3 elements from source to destination
	copied := copy(destination, source)
	
	fmt.Printf("\nCopied %d elements: %v\n", copied, destination)
	
	// Copy with different sized slices
	small := []int{100, 200}
	large := make([]int, 5)
	
	// Copy small to large (copy() handles size differences)
	n := copy(large, small)
	
	fmt.Printf("Copied %d elements to large slice: %v\n", n, large)
}