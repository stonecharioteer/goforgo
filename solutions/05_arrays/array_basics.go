// array_basics.go - SOLUTION
// Learn the fundamentals of arrays in Go
// Arrays have a fixed size and all elements must be of the same type

package main

import "fmt"

func main() {
	// Declare an array of 5 integers
	var numbers [5]int // Complete this declaration

	// Initialize an array with values using array literal
	colors := [3]string{"red", "green", "blue"} // Complete this initialization

	// Use the ... operator to let Go determine the array size
	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"} // Complete this initialization

	// Access the third element (index 2) of the numbers array and assign it the value 42
	numbers[2] = 42 // Complete this assignment

	// Print the length of the colors array
	fmt.Println("Length of colors array:", len(colors)) // Complete this line

	// Print all arrays to see the results
	fmt.Println("Numbers:", numbers)
	fmt.Println("Colors:", colors)
	fmt.Println("Days:", days)
}