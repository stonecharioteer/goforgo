package main

import "fmt"

// Define a recursive function called 'factorial' that takes an int n
// and returns its factorial (n!)
// Base case: if n is 0 or 1, return 1
// Recursive case: return n * factorial(n-1)
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Define a recursive function called 'fibonacci' that takes an int n  
// and returns the nth Fibonacci number
// Base cases: if n is 0, return 0; if n is 1, return 1
// Recursive case: return fibonacci(n-1) + fibonacci(n-2)
func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Define a recursive function called 'countdown' that takes an int n
// It should print numbers from n down to 1, then print "Blast off!"
// Base case: if n is 0, print "Blast off!" and return
// Recursive case: print n, then call countdown(n-1)
func countdown(n int) {
	if n == 0 {
		fmt.Println("Blast off!")
		return
	}
	fmt.Println(n)
	countdown(n - 1)
}

func main() {
	// Calculate and print factorial of 5
	fact5 := factorial(5)
	fmt.Println("Factorial of 5:", fact5)
	
	// Calculate and print factorial of 0 (should be 1)
	fact0 := factorial(0)
	fmt.Println("Factorial of 0:", fact0)
	
	// Calculate and print the 7th Fibonacci number
	fib7 := fibonacci(7)
	fmt.Println("7th Fibonacci number:", fib7)
	
	// Calculate and print the 0th and 1st Fibonacci numbers
	fib0 := fibonacci(0)
	fib1 := fibonacci(1)
	fmt.Println("0th Fibonacci:", fib0, "1st Fibonacci:", fib1)
	
	// Call countdown with 5
	fmt.Println("Starting countdown:")
	countdown(5)
}