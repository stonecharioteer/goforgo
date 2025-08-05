// multidimensional_arrays.go
// Learn how to work with 2D and higher dimensional arrays

package main

import "fmt"

func main() {
	// TODO: Declare a 3x3 matrix (2D array) of integers
	// Hint: var name [rows][cols]type
	var matrix // Complete this declaration

	// TODO: Initialize a 2x4 array with values
	// Hint: arrayName := [rows][cols]type{{row1}, {row2}}
	grades := // Complete this initialization

	// TODO: Access and modify elements in the matrix
	// Set matrix[1][1] to 5
	// Complete this assignment

	// TODO: Print the matrix using nested loops
	fmt.Println("Matrix:")
	// Write nested loops to print the matrix

	// TODO: Print the grades array using nested range loops
	fmt.Println("\nGrades:")
	// Write nested range loops

	// TODO: Calculate the sum of all elements in grades
	var totalSum int
	// Write nested loops to calculate sum
	fmt.Println("Total sum of grades:", totalSum)

	// TODO: Find the dimensions of the grades array
	rows := // Get number of rows
	cols := // Get number of columns (if rows > 0)
	fmt.Printf("Grades array dimensions: %dx%d\n", rows, cols)
}