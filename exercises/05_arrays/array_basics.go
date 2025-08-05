// array_basics.go
// Learn the fundamentals of arrays in Go
// Arrays have a fixed size and all elements must be of the same type

package main

import "fmt"

func main() {
	// TODO: Declare an array of 5 integers
	// Hint: var arrayName [size]type
	var numbers // Complete this declaration

	// TODO: Initialize an array with values using array literal
	// Hint: arrayName := [size]type{value1, value2, ...}
	colors := // Complete this initialization

	// TODO: Use the ... operator to let Go determine the array size
	// Hint: arrayName := [...]type{values}
	days := // Complete this initialization

	// TODO: Access the third element (index 2) of the numbers array
	// and assign it the value 42
	// Complete this assignment

	// TODO: Print the length of the colors array
	// Hint: use len() function
	fmt.Println("Length of colors array:") // Complete this line

	// Print all arrays to see the results
	fmt.Println("Numbers:", numbers)
	fmt.Println("Colors:", colors)
	fmt.Println("Days:", days)
}