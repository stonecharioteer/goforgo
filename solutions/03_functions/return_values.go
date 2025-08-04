package main

import "fmt"

// Define a function called 'double' that takes an int parameter and returns an int
// The function should return the parameter multiplied by 2
func double(x int) int {
	return x * 2
}

// Define a function called 'getGreeting' that takes a string parameter 'name'
// and returns a string greeting like "Hello, [name]!"
func getGreeting(name string) string {
	return "Hello, " + name + "!"
}

// Define a function called 'divide' that takes two float64 parameters (a, b)
// and returns two values: the result (float64) and a success flag (bool)
// Return (0.0, false) if b is zero, otherwise return (a/b, true)
func divide(a, b float64) (float64, bool) {
	if b == 0 {
		return 0.0, false
	}
	return a / b, true
}

func main() {
	// Call double with 5 and print the result
	result := double(5)
	fmt.Println("Double of 5 is:", result)
	
	// Call getGreeting with "World" and print the result
	greeting := getGreeting("World")
	fmt.Println(greeting)
	
	// Call divide with 10.0 and 2.0, capture both return values and print them
	quotient, success := divide(10.0, 2.0)
	fmt.Println("10.0 / 2.0 =", quotient, "Success:", success)
	
	// Call divide with 10.0 and 0.0, capture both return values and print them
	quotient2, success2 := divide(10.0, 0.0)
	fmt.Println("10.0 / 0.0 =", quotient2, "Success:", success2)
}