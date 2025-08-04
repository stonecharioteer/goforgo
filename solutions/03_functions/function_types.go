package main

import "fmt"

// Define a function called 'add' that takes two integers and returns their sum
func add(a, b int) int {
	return a + b
}

// Define a function called 'multiply' that takes two integers and returns their product
func multiply(a, b int) int {
	return a * b
}

// Define a function called 'operate' that takes three parameters:
// - a: int
// - b: int  
// - operation: a function that takes two ints and returns an int
// The function should call the operation function with a and b and return the result
func operate(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func main() {
	// Create a variable called 'mathFunc' with the type that matches add and multiply
	// Assign the add function to it
	var mathFunc func(int, int) int = add
	
	// Call mathFunc with 5 and 3, print the result
	result1 := mathFunc(5, 3)
	fmt.Println("mathFunc(5, 3) with add:", result1)
	
	// Assign the multiply function to mathFunc
	mathFunc = multiply
	
	// Call mathFunc with 5 and 3, print the result
	result2 := mathFunc(5, 3)
	fmt.Println("mathFunc(5, 3) with multiply:", result2)
	
	// Call operate with 10, 4, and the add function, print the result
	result3 := operate(10, 4, add)
	fmt.Println("operate(10, 4, add):", result3)
	
	// Call operate with 10, 4, and the multiply function, print the result
	result4 := operate(10, 4, multiply)
	fmt.Println("operate(10, 4, multiply):", result4)
}