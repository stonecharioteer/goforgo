package main

import "fmt"

// TODO: Define a function called 'cleanup' that prints "Cleaning up resources"

// TODO: Define a function called 'openFile' that takes a filename string parameter
// Print "Opening file:", filename
// Defer a call to cleanup() 
// Print "File operations complete"

// TODO: Define a function called 'deferOrder' that demonstrates defer execution order
// Print "Start"
// Defer printing "First defer"
// Defer printing "Second defer" 
// Defer printing "Third defer"
// Print "End"
// Note: defers execute in LIFO (Last In, First Out) order

// TODO: Define a function called 'deferWithArgs' 
// Create variable x := 10
// Defer fmt.Println("Deferred x:", x)
// Change x to 20
// Print "Current x:", x
// Note: defer captures argument values at defer time, not execution time

// TODO: Define a function called 'deferInLoop'
// Use a for loop i from 1 to 3
// In each iteration, defer printing "Loop defer:", i
// Print "Loop iteration:", i
// Note: observe how defer behaves in loops

func main() {
	// TODO: Call openFile with "data.txt"
	
	// TODO: Call deferOrder and observe the execution order
	
	// TODO: Call deferWithArgs and observe argument capture behavior
	
	// TODO: Call deferInLoop and observe defer in loops
	
	// TODO: Demonstrate defer with anonymous function
	// Create variable message := "Hello"
	// Defer an anonymous function that prints "Deferred message:", message
	// Change message to "Goodbye"  
	// Print "Current message:", message
}