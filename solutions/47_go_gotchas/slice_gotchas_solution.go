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
	
	// FIXED: Correct way to check if a slice is empty
	if len(nilSlice) == 0 { // Use len() to check emptiness, not comparison with nil
		fmt.Println("nil slice is empty")
	}
	
	// Gotcha 2: Slice capacity surprises
	smallSlice := []int{1, 2}
	fmt.Printf("Small slice: len=%d, cap=%d\n", len(smallSlice), cap(smallSlice))
	
	// When capacity is exceeded, Go creates a new underlying array with ~2x capacity
	smallSlice = append(smallSlice, 3)
	fmt.Printf("After append: len=%d, cap=%d\n", len(smallSlice), cap(smallSlice))
	
	// Gotcha 3: Unexpected slice behavior with underlying arrays
	original := []int{1, 2, 3, 4, 5}
	slice1 := original[1:3] // [2, 3] - shares underlying array
	slice2 := original[2:4] // [3, 4] - shares same underlying array
	
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Slice1: %v\n", slice1)
	fmt.Printf("Slice2: %v\n", slice2)
	
	// Modifying slice1 affects original and slice2 because they share underlying array
	slice1[1] = 99 // This modifies original[2], which is also slice2[0]
	fmt.Printf("After slice1[1] = 99:\n")
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Slice1: %v\n", slice1)
	fmt.Printf("Slice2: %v\n", slice2) // slice2[0] is now 99!
	
	// Gotcha 4: Slice append gotcha - when does it create a new underlying array?
	base := make([]int, 2, 4) // len=2, cap=4
	base[0] = 1
	base[1] = 2
	
	derived := append(base, 3) // This fits in capacity, modifies underlying array
	derived2 := append(base, 4) // This overwrites the 3 with 4!
	
	fmt.Printf("Base: %v\n", base)
	fmt.Printf("Derived: %v\n", derived)   // [1, 2, 4] - not [1, 2, 3]!
	fmt.Printf("Derived2: %v\n", derived2) // [1, 2, 4]
	
	// Both derived slices share the same underlying array with base
}

func correctSliceEmptyCheck(s []int) bool {
	// FIXED: Correct way to check if slice is empty
	// Use len(s) == 0 which works for both nil and empty slices
	return len(s) == 0
	
	// Alternative approaches:
	// return s == nil || len(s) == 0  // Explicit nil check
	// The len() approach is preferred because it treats nil and empty consistently
}

func sliceMemoryLeakFixed() {
	// Gotcha 5: Slice memory leak gotcha - FIXED
	largeSlice := make([]byte, 1000000) // 1MB slice
	
	// WRONG: smallSlice := largeSlice[:5]
	// This keeps the entire 1MB array in memory!
	
	// FIXED: Create a new slice with only the needed data
	smallSlice := make([]byte, 5)
	copy(smallSlice, largeSlice[:5])
	// Now largeSlice can be garbage collected
	
	fmt.Printf("Small slice len: %d, capacity: %d bytes\n", 
		len(smallSlice), cap(smallSlice))
	
	// Alternative fix using append:
	// smallSlice := append([]byte(nil), largeSlice[:5]...)
}

func main() {
	fmt.Println("=== Slice Gotchas Demo ===")
	demonstrateSliceGotchas()
	
	fmt.Println("\n=== Correct Empty Check ===")
	var nilSlice []int
	emptySlice := []int{}
	
	fmt.Printf("nil slice empty check: %t\n", correctSliceEmptyCheck(nilSlice))     // true
	fmt.Printf("empty slice empty check: %t\n", correctSliceEmptyCheck(emptySlice)) // true
	
	fmt.Println("\n=== Memory Leak Fixed ===")
	sliceMemoryLeakFixed()
}