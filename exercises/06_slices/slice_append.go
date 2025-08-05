// slice_append.go
// Learn how to append elements to slices and understand slice growth

package main

import "fmt"

func main() {
	// TODO: Create an empty slice and append elements
	var numbers []int
	// Append the number 1
	numbers = // Complete this append operation
	
	// TODO: Append multiple elements at once
	// Append 2, 3, 4 to numbers
	numbers = // Complete this append operation
	
	// TODO: Append elements from another slice
	moreNumbers := []int{5, 6, 7}
	// Hint: Use ... to expand the slice
	numbers = // Complete this append operation
	
	// TODO: Create a slice with initial capacity and observe growth
	fruits := make([]string, 0, 2) // length 0, capacity 2
	fmt.Printf("Initial - len: %d, cap: %d\n", len(fruits), cap(fruits))
	
	// TODO: Append fruits and observe capacity changes
	fruits = append(fruits, "apple")
	fmt.Printf("After 1 append - len: %d, cap: %d\n", len(fruits), cap(fruits))
	
	fruits = append(fruits, "banana")  
	fmt.Printf("After 2 appends - len: %d, cap: %d\n", len(fruits), cap(fruits))
	
	// TODO: This should trigger capacity growth
	fruits = append(fruits, "cherry")
	fmt.Printf("After 3 appends - len: %d, cap: %d\n", len(fruits), cap(fruits))
	
	// TODO: Append to nil slice (this works!)
	var colors []string
	colors = // Append "red" to the nil slice
	
	fmt.Println("Numbers:", numbers)
	fmt.Println("Fruits:", fruits)  
	fmt.Println("Colors:", colors)
}