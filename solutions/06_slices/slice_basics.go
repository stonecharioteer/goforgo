// slice_basics.go - SOLUTION
// Learn the fundamentals of slices in Go
// Slices are dynamic arrays with flexible size

package main

import "fmt"

func main() {
	// Create a slice from an array
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[1:4] // Contains elements 1-3
	
	// Create a slice using make()
	slice2 := make([]string, 3, 5) // length 3, capacity 5
	
	// Create a slice using slice literal
	slice3 := []int{10, 20, 30}
	
	// Create an empty slice
	var slice4 []int
	
	// Check if slice4 is nil
	if slice4 == nil {
		fmt.Println("slice4 is nil")
	}
	
	// Get length and capacity of slice2
	length := len(slice2)
	capacity := cap(slice2)
	
	// Modify elements in slice1
	// Change the second element to 99
	slice1[1] = 99
	
	// Print results
	fmt.Println("Original array:", arr)
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("slice3:", slice3)
	fmt.Printf("slice2 - length: %d, capacity: %d\n", length, capacity)
}