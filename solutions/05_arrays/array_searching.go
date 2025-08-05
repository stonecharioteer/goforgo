// array_searching.go - SOLUTION
// Learn different search algorithms with arrays
// Practice linear search, binary search, and finding min/max values

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Sample data for searching
	numbers := [10]int{64, 34, 25, 12, 22, 11, 90, 88, 76, 50}
	sortedNumbers := [8]int{5, 12, 23, 34, 45, 67, 78, 89}
	
	fmt.Println("Numbers array:", numbers)
	fmt.Println("Sorted numbers:", sortedNumbers)

	// Implement linear search to find target value 22
	target := 22
	foundIndex := -1
	
	// Linear search loop to find target in numbers array
	for i, value := range numbers {
		if value == target {
			foundIndex = i
			break
		}
	}
	
	if foundIndex != -1 {
		fmt.Printf("Linear search: Found %d at index %d\n", target, foundIndex)
	} else {
		fmt.Printf("Linear search: %d not found\n", target)
	}

	// Implement binary search on sortedNumbers array
	binaryTarget := 45
	binaryIndex := sort.SearchInts(sortedNumbers[:], binaryTarget)
	
	if binaryIndex < len(sortedNumbers) && sortedNumbers[binaryIndex] == binaryTarget {
		fmt.Printf("Binary search: Found %d at index %d\n", binaryTarget, binaryIndex)
	} else {
		fmt.Printf("Binary search: %d not found\n", binaryTarget)
	}

	// Find minimum and maximum values in numbers array
	if len(numbers) > 0 {
		min := numbers[0]
		max := numbers[0]
		minIndex := 0
		maxIndex := 0
		
		// Loop to find min, max and their indices
		for i, value := range numbers {
			if value < min {
				min = value
				minIndex = i
			}
			if value > max {
				max = value
				maxIndex = i
			}
		}
		
		fmt.Printf("Minimum: %d at index %d\n", min, minIndex)
		fmt.Printf("Maximum: %d at index %d\n", max, maxIndex)
	}

	// Count occurrences of a value
	countTarget := 12
	count := 0
	
	// Loop to count how many times countTarget appears in numbers
	for _, value := range numbers {
		if value == countTarget {
			count++
		}
	}
	
	fmt.Printf("Number %d appears %d times\n", countTarget, count)

	// Find all indices where a value appears
	searchValue := 34
	var indices []int
	
	// Loop to collect all indices where searchValue appears
	for i, value := range numbers {
		if value == searchValue {
			indices = append(indices, i)
		}
	}
	
	fmt.Printf("Value %d found at indices: %v\n", searchValue, indices)
}