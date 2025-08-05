// multidimensional_arrays.go - SOLUTION
// Learn how to work with 2D and higher dimensional arrays

package main

import "fmt"

func main() {
	// Declare a 3x3 matrix (2D array) of integers
	var matrix [3][3]int

	// Initialize a 2x4 array with values
	grades := [2][4]int{{85, 92, 78, 96}, {88, 91, 84, 87}}

	// Access and modify elements in the matrix
	// Set matrix[1][1] to 5
	matrix[1][1] = 5

	// Print the matrix using nested loops
	fmt.Println("Matrix:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%3d ", matrix[i][j])
		}
		fmt.Println()
	}

	// Print the grades array using nested range loops
	fmt.Println("\nGrades:")
	for i, row := range grades {
		fmt.Printf("Student %d: ", i+1)
		for _, grade := range row {
			fmt.Printf("%d ", grade)
		}
		fmt.Println()
	}

	// Calculate the sum of all elements in grades
	var totalSum int
	for i := 0; i < len(grades); i++ {
		for j := 0; j < len(grades[i]); j++ {
			totalSum += grades[i][j]
		}
	}
	fmt.Println("Total sum of grades:", totalSum)

	// Find the dimensions of the grades array
	rows := len(grades)
	var cols int
	if rows > 0 {
		cols = len(grades[0])
	}
	fmt.Printf("Grades array dimensions: %dx%d\n", rows, cols)
}