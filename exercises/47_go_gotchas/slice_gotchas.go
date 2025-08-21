package main

import "fmt"

func demonstrateSliceGotchas() {
	// Gotcha 1: nil vs empty slice confusion
	var nilSlice []int
	emptySlice := []int{}
	emptySliceMake := make([]int, 0)
	
	fmt.Printf("nil slice: %v, len: %d, cap: %d, is nil: %t\n", 
		nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("empty slice literal: %v, len: %d, cap: %d, is nil: %t\n", 
		emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)
	fmt.Printf("empty slice make: %v, len: %d, cap: %d, is nil: %t\n", 
		emptySliceMake, len(emptySliceMake), cap(emptySliceMake), emptySliceMake == nil)
	
	// Which comparison is correct for checking if a slice is empty?
	// Fix this function to properly check for empty slices
	if len(nilSlice) == nil { // This line has a bug
		fmt.Println("nil slice is empty")
	}
	
	// Gotcha 2: Slice capacity surprises
	smallSlice := []int{1, 2}
	fmt.Printf("Small slice: len=%d, cap=%d\n", len(smallSlice), cap(smallSlice))
	
	// What happens when we append?
	smallSlice = append(smallSlice, 3)
	fmt.Printf("After append: len=%d, cap=%d\n", len(smallSlice), cap(smallSlice))
	
	// Gotcha 3: Unexpected slice behavior with underlying arrays
	original := []int{1, 2, 3, 4, 5}
	slice1 := original[1:3] // [2, 3]
	slice2 := original[2:4] // [3, 4]
	
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Slice1: %v\n", slice1)
	fmt.Printf("Slice2: %v\n", slice2)
	
	// Modifying slice1 - what happens to original and slice2?
	slice1[1] = 99
	fmt.Printf("After slice1[1] = 99:\n")
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Slice1: %v\n", slice1)
	fmt.Printf("Slice2: %v\n", slice2)
	
	// Gotcha 4: Slice append gotcha - when does it create a new underlying array?
	base := make([]int, 2, 4) // len=2, cap=4
	base[0] = 1
	base[1] = 2
	
	derived := append(base, 3) // This fits in capacity
	derived2 := append(base, 4) // This also fits in capacity
	
	fmt.Printf("Base: %v\n", base)
	fmt.Printf("Derived: %v\n", derived)
	fmt.Printf("Derived2: %v\n", derived2)
	
	// What do you expect to see here? Fix the logic if needed.
}

func incorrectSliceEmptyCheck(s []int) bool {
	// This function has multiple bugs in checking if slice is empty
	if s == nil {
		return true
	}
	if len(s) == 0 {
		return false // This logic is wrong - fix it
	}
	return true
}

func sliceMemoryLeak() {
	// Gotcha 5: Slice memory leak gotcha
	largeSlice := make([]byte, 1000000) // 1MB slice
	
	// We only want the first 5 bytes, but this creates a memory leak
	smallSlice := largeSlice[:5]
	
	// The smallSlice still references the entire 1MB array!
	// How would you fix this to avoid the memory leak?
	
	fmt.Printf("Small slice len: %d, but underlying array is still %d bytes\n", 
		len(smallSlice), cap(smallSlice))
	
	// TODO: Fix this by properly copying just the needed data
}

func main() {
	fmt.Println("=== Slice Gotchas Demo ===")
	demonstrateSliceGotchas()
	
	fmt.Println("\n=== Incorrect Empty Check ===")
	var nilSlice []int
	emptySlice := []int{}
	
	fmt.Printf("nil slice empty check: %t\n", incorrectSliceEmptyCheck(nilSlice))
	fmt.Printf("empty slice empty check: %t\n", incorrectSliceEmptyCheck(emptySlice))
	
	fmt.Println("\n=== Memory Leak Demo ===")
	sliceMemoryLeak()
}