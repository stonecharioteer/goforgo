// array_sorting.go
// Learn how to sort arrays using the sort package
// This exercise covers sorting different types of arrays

package main

import (
	"fmt"
	// TODO: Import the sort package
)

func main() {
	// TODO: Create an array of integers to sort
	numbers := // Complete this with [5]int{64, 34, 25, 12, 22}

	// TODO: Create an array of strings to sort
	names := // Complete this with [4]string{"Charlie", "Alice", "Bob", "David"}

	// TODO: Create an array of float64 to sort
	prices := // Complete this with [6]float64{19.99, 9.99, 29.99, 4.99, 39.99, 14.99}

	fmt.Println("Before sorting:")
	fmt.Println("Numbers:", numbers)
	fmt.Println("Names:", names)
	fmt.Println("Prices:", prices)

	// TODO: Convert arrays to slices and sort them
	// Hint: Use numbers[:] to convert array to slice, then sort.Ints()
	// Sort the numbers array

	// TODO: Sort the names array
	// Hint: Use sort.Strings()

	// TODO: Sort the prices array
	// Hint: Use sort.Float64s()

	fmt.Println("\nAfter sorting:")
	fmt.Println("Numbers:", numbers)
	fmt.Println("Names:", names)
	fmt.Println("Prices:", prices)

	// TODO: Check if an array is sorted
	// Hint: Use sort.IntsAreSorted() for integers
	isSorted := // Check if numbers slice is sorted
	fmt.Println("Numbers array is sorted:", isSorted)

	// TODO: Custom sorting - sort in descending order
	descNumbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Before reverse sort:", descNumbers)
	
	// TODO: Sort in descending order using sort.Sort with sort.Reverse
	// Hint: sort.Sort(sort.Reverse(sort.IntSlice(descNumbers[:])))
	
	fmt.Println("After reverse sort:", descNumbers)
}