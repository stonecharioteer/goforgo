package main

import "fmt"

// Define a function called 'cleanup' that prints "Cleaning up resources"
func cleanup() {
	fmt.Println("Cleaning up resources")
}

// Define a function called 'openFile' that takes a filename string parameter
// Print "Opening file:", filename
// Defer a call to cleanup() 
// Print "File operations complete"
func openFile(filename string) {
	fmt.Println("Opening file:", filename)
	defer cleanup()
	fmt.Println("File operations complete")
}

// Define a function called 'deferOrder' that demonstrates defer execution order
// Print "Start"
// Defer printing "First defer"
// Defer printing "Second defer" 
// Defer printing "Third defer"
// Print "End"
// Note: defers execute in LIFO (Last In, First Out) order
func deferOrder() {
	fmt.Println("Start")
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")
	fmt.Println("End")
	// Execution order will be: Start, End, Third defer, Second defer, First defer
}

// Define a function called 'deferWithArgs' 
// Create variable x := 10
// Defer fmt.Println("Deferred x:", x)
// Change x to 20
// Print "Current x:", x
// Note: defer captures argument values at defer time, not execution time
func deferWithArgs() {
	x := 10
	defer fmt.Println("Deferred x:", x) // x is captured as 10 here
	x = 20
	fmt.Println("Current x:", x)
	// Deferred will print 10, not 20
}

// Define a function called 'deferInLoop'
// Use a for loop i from 1 to 3
// In each iteration, defer printing "Loop defer:", i
// Print "Loop iteration:", i
// Note: observe how defer behaves in loops
func deferInLoop() {
	for i := 1; i <= 3; i++ {
		defer fmt.Println("Loop defer:", i)
		fmt.Println("Loop iteration:", i)
	}
	// Defers will execute: 3, 2, 1 (LIFO order)
}

func main() {
	// Call openFile with "data.txt"
	fmt.Println("=== openFile example ===")
	openFile("data.txt")
	
	// Call deferOrder and observe the execution order
	fmt.Println("\n=== deferOrder example ===")
	deferOrder()
	
	// Call deferWithArgs and observe argument capture behavior
	fmt.Println("\n=== deferWithArgs example ===")
	deferWithArgs()
	
	// Call deferInLoop and observe defer in loops
	fmt.Println("\n=== deferInLoop example ===")
	deferInLoop()
	
	// Demonstrate defer with anonymous function
	// Create variable message := "Hello"
	// Defer an anonymous function that prints "Deferred message:", message
	// Change message to "Goodbye"  
	// Print "Current message:", message
	fmt.Println("\n=== defer with anonymous function ===")
	message := "Hello"
	defer func() {
		fmt.Println("Deferred message:", message) // This will capture "Goodbye" (closure)
	}()
	message = "Goodbye"
	fmt.Println("Current message:", message)
}