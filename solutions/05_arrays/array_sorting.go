// array_sorting.go - SOLUTION
// Learn how to sort arrays using the sort package
// This exercise covers sorting different types of arrays

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Create an array of integers to sort
	numbers := [5]int{64, 34, 25, 12, 22}

	// Create an array of strings to sort
	names := [4]string{"Charlie", "Alice", "Bob", "David"}

	// Create an array of float64 to sort
	prices := [6]float64{19.99, 9.99, 29.99, 4.99, 39.99, 14.99}

	fmt.Println("Before sorting:")
	fmt.Println("Numbers:", numbers)
	fmt.Println("Names:", names)
	fmt.Println("Prices:", prices)

	// Convert arrays to slices and sort them
	// Hint: Use numbers[:] to convert array to slice, then sort.Ints()
	sort.Ints(numbers[:])

	// Sort the names array
	// Hint: Use sort.Strings()
	sort.Strings(names[:])

	// Sort the prices array
	// Hint: Use sort.Float64s()
	sort.Float64s(prices[:])

	fmt.Println("\nAfter sorting:")
	fmt.Println("Numbers:", numbers)
	fmt.Println("Names:", names)
	fmt.Println("Prices:", prices)

	// Check if an array is sorted
	// Hint: Use sort.IntsAreSorted() for integers
	isSorted := sort.IntsAreSorted(numbers[:])
	fmt.Println("Numbers array is sorted:", isSorted)

	// Custom sorting - sort in descending order
	descNumbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Before reverse sort:", descNumbers)
	
	// Sort in descending order using sort.Sort with sort.Reverse
	// Hint: sort.Sort(sort.Reverse(sort.IntSlice(descNumbers[:])))
	sort.Sort(sort.Reverse(sort.IntSlice(descNumbers[:])))
	
	fmt.Println("After reverse sort:", descNumbers)
}