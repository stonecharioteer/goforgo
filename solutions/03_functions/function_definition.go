package main

import "fmt"

// Define a function called 'greet' that takes no parameters and returns nothing
// The function should print "Hello from a function!" when called
func greet() {
	fmt.Println("Hello from a function!")
}

func main() {
	// Call the greet function here
	fmt.Println("This runs before the function call")
	greet()
	fmt.Println("This runs after the function call")
}