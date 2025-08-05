// testing_basics.go
// Learn the fundamentals of testing in Go

package main

import "fmt"

// TODO: Functions to test
func Add(a, b int) int {
	// Return sum of a and b
}

func Multiply(a, b int) int {
	// Return product of a and b
}

func IsEven(n int) bool {
	// Return true if n is even
}

func Factorial(n int) int {
	// Calculate factorial of n
	// Handle edge cases: n <= 0 should return 1
}

func ReverseString(s string) string {
	// Reverse the string
	// Return reversed string
}

func FindMax(numbers []int) (int, error) {
	// Find maximum number in slice
	// Return error if slice is empty
}

// TODO: This function has a bug - fix it in the test
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	// This file contains functions to be tested
	// The actual tests will be in testing_basics_test.go
	
	fmt.Println("This file contains functions to be tested.")
	fmt.Println("Run 'go test' to execute the tests.")
	
	// Demo the functions
	fmt.Printf("Add(5, 3) = %d\\n", Add(5, 3))
	fmt.Printf("Multiply(4, 6) = %d\\n", Multiply(4, 6))
	fmt.Printf("IsEven(8) = %t\\n", IsEven(8))
	fmt.Printf("Factorial(5) = %d\\n", Factorial(5))
	fmt.Printf("ReverseString(\\\"hello\\\") = %s\\n", ReverseString("hello"))
	
	max, err := FindMax([]int{3, 7, 2, 9, 1})
	if err != nil {
		fmt.Printf("Error: %v\\n", err)
	} else {
		fmt.Printf("FindMax([3,7,2,9,1]) = %d\\n", max)
	}
	
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\\n", err)
	} else {
		fmt.Printf("Divide(10, 2) = %.2f\\n", result)
	}
}