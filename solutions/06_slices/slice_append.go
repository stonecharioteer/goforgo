// slice_append.go - SOLUTION
// Learn how to append elements to slices and understand slice growth

package main

import "fmt"

func main() {
	// Create an empty slice and append elements
	var numbers []int
	// Append the number 1
	numbers = append(numbers, 1)
	
	// Append multiple elements at once
	// Append 2, 3, 4 to numbers
	numbers = append(numbers, 2, 3, 4)
	
	// Append elements from another slice
	moreNumbers := []int{5, 6, 7}
	numbers = append(numbers, moreNumbers...)
	
	// Create a slice with initial capacity and observe growth
	fruits := make([]string, 0, 2) // length 0, capacity 2
	fmt.Printf("Initial - len: %d, cap: %d\n", len(fruits), cap(fruits))
	
	// Append fruits and observe capacity changes
	fruits = append(fruits, "apple")
	fmt.Printf("After 1 append - len: %d, cap: %d\n", len(fruits), cap(fruits))
	
	fruits = append(fruits, "banana")  
	fmt.Printf("After 2 appends - len: %d, cap: %d\n", len(fruits), cap(fruits))
	
	// This should trigger capacity growth
	fruits = append(fruits, "cherry")
	fmt.Printf("After 3 appends - len: %d, cap: %d\n", len(fruits), cap(fruits))
	
	// Append to nil slice (this works!)
	var colors []string
	colors = append(colors, "red")
	
	fmt.Println("Numbers:", numbers)
	fmt.Println("Fruits:", fruits)  
	fmt.Println("Colors:", colors)
}