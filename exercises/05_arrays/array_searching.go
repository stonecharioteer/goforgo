// array_searching.go
// Learn different search algorithms with arrays
// Practice linear search, binary search, and finding min/max values

package main

import (
	"fmt"
	// TODO: Import sort package for binary search
)

func main() {
	// Sample data for searching
	numbers := [10]int{64, 34, 25, 12, 22, 11, 90, 88, 76, 50}
	sortedNumbers := [8]int{5, 12, 23, 34, 45, 67, 78, 89}
	
	fmt.Println("Numbers array:", numbers)
	fmt.Println("Sorted numbers:", sortedNumbers)

	// TODO: Implement linear search to find target value 22
	target := 22
	foundIndex := -1
	
	// Write a linear search loop to find target in numbers array
	// Set foundIndex to the index where target is found, or -1 if not found
	
	if foundIndex != -1 {
		fmt.Printf("Linear search: Found %d at index %d\n", target, foundIndex)
	} else {
		fmt.Printf("Linear search: %d not found\n", target)
	}

	// TODO: Implement binary search on sortedNumbers array
	// Hint: Use sort.SearchInts() or implement manually
	binaryTarget := 45
	binaryIndex := // Use sort.SearchInts(sortedNumbers[:], binaryTarget)
	
	if binaryIndex < len(sortedNumbers) && sortedNumbers[binaryIndex] == binaryTarget {
		fmt.Printf("Binary search: Found %d at index %d\n", binaryTarget, binaryIndex)
	} else {
		fmt.Printf("Binary search: %d not found\n", binaryTarget)
	}

	// TODO: Find minimum and maximum values in numbers array
	if len(numbers) > 0 {
		min := numbers[0]
		max := numbers[0]
		minIndex := 0
		maxIndex := 0
		
		// Write a loop to find min, max and their indices
		
		fmt.Printf("Minimum: %d at index %d\n", min, minIndex)
		fmt.Printf("Maximum: %d at index %d\n", max, maxIndex)
	}

	// TODO: Count occurrences of a value
	countTarget := 12
	count := 0
	
	// Write a loop to count how many times countTarget appears in numbers
	
	fmt.Printf("Number %d appears %d times\n", countTarget, count)

	// TODO: Find all indices where a value appears
	searchValue := 34
	var indices []int
	
	// Write a loop to collect all indices where searchValue appears
	
	fmt.Printf("Value %d found at indices: %v\n", searchValue, indices)
}