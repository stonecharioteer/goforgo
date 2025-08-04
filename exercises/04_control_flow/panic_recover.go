package main

import "fmt"

// TODO: Define a function called 'safeDivide' that takes two int parameters a and b
// Use defer and recover to catch panics
// If b is 0, panic with message "division by zero"
// In the deferred function, use recover() to catch the panic
// If panic occurred, print "Recovered from panic:", and the panic value
// Return the division result (a/b) if no panic, or 0 if panic recovered

// TODO: Define a function called 'processArray' that takes []int parameter
// Use defer and recover to handle out-of-bounds access
// Try to access arr[10] (will panic if array is shorter)
// In deferred function, recover and print "Recovered from array access panic"
// Print the value at arr[10] if no panic occurs

// TODO: Define a function called 'demonstratePanic'  
// Print "Before panic"
// Call panic("Something went wrong!")
// Print "After panic" (this should not execute)

// TODO: Define a function called 'nestedPanic'
// Defer a recovery function that prints "Outer recovery:", recover()
// Call a function that:
//   - Defers a recovery function that prints "Inner recovery:", recover()
//   - Panics with "Inner panic"
// Note: only the innermost recover catches the panic

func main() {
	// TODO: Test safeDivide with normal division: safeDivide(10, 2)
	
	// TODO: Test safeDivide with division by zero: safeDivide(10, 0)
	
	// TODO: Test processArray with short array: []int{1, 2, 3}
	
	// TODO: Test processArray with long array: make([]int, 15) (fill with some values)
	
	// TODO: Call demonstratePanic inside a function with defer/recover
	// Use an anonymous function with defer/recover to catch the panic
	
	// TODO: Call nestedPanic to see nested recovery behavior
	
	// TODO: Demonstrate that the program continues after recovered panics
	// Print "Program continues normally"
}