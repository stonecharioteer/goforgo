// slice_basics.go
// Learn the fundamentals of slices in Go
// Slices are dynamic arrays with flexible size

package main

import "fmt"

func main() {
	// TODO: Create a slice from an array
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := // Create a slice from arr containing elements 1-3
	
	// TODO: Create a slice using make()
	// Hint: make([]type, length, capacity)
	slice2 := // Create a slice of 3 strings with capacity 5
	
	// TODO: Create a slice using slice literal
	// Hint: []type{values}
	slice3 := // Create a slice with numbers 10, 20, 30
	
	// TODO: Create an empty slice
	var slice4 // Complete this declaration
	
	// TODO: Check if slice4 is nil
	if /* condition */ {
		fmt.Println("slice4 is nil")
	}
	
	// TODO: Get length and capacity of slice2
	length := // Get length of slice2
	capacity := // Get capacity of slice2
	
	// TODO: Modify elements in slice1
	// Change the second element to 99
	// Complete this assignment
	
	// Print results
	fmt.Println("Original array:", arr)
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("slice3:", slice3)
	fmt.Printf("slice2 - length: %d, capacity: %d\n", length, capacity)
}