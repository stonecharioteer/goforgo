package main

import "fmt"

// Define a function called 'sum' that takes a variadic parameter of integers
// The function should return the sum of all the numbers passed to it
// Use the syntax: func sum(numbers ...int) int
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Define a function called 'printAll' that takes a variadic parameter of strings
// The function should print each string on a separate line with a number prefix
// Example output: "1: hello", "2: world"
func printAll(items ...string) {
	for i, item := range items {
		fmt.Printf("%d: %s\n", i+1, item)
	}
}

func main() {
	// Call sum with no arguments and print the result
	result0 := sum()
	fmt.Println("Sum of no numbers:", result0)
	
	// Call sum with 1, 2, 3 and print the result
	result1 := sum(1, 2, 3)
	fmt.Println("Sum of 1, 2, 3:", result1)
	
	// Call sum with 10, 20, 30, 40, 50 and print the result
	result2 := sum(10, 20, 30, 40, 50)
	fmt.Println("Sum of 10, 20, 30, 40, 50:", result2)
	
	// Create a slice of integers: []int{1, 2, 3, 4, 5}
	// Call sum with the slice using the spread operator (...) and print the result
	numbers := []int{1, 2, 3, 4, 5}
	result3 := sum(numbers...)
	fmt.Println("Sum of slice:", result3)
	
	// Call printAll with "apple", "banana", "cherry"
	printAll("apple", "banana", "cherry")
}