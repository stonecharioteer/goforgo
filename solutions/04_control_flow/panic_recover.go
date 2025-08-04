package main

import "fmt"

// Define a function called 'safeDivide' that takes two int parameters a and b
// Use defer and recover to catch panics
// If b is 0, panic with message "division by zero"
// In the deferred function, use recover() to catch the panic
// If panic occurred, print "Recovered from panic:", and the panic value
// Return the division result (a/b) if no panic, or 0 if panic recovered
func safeDivide(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			result = 0 // Set return value for panic case
		}
	}()
	
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

// Define a function called 'processArray' that takes []int parameter
// Use defer and recover to handle out-of-bounds access
// Try to access arr[10] (will panic if array is shorter)
// In deferred function, recover and print "Recovered from array access panic"
// Print the value at arr[10] if no panic occurs
func processArray(arr []int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from array access panic")
		}
	}()
	
	value := arr[10] // This will panic if arr has < 11 elements
	fmt.Println("Value at index 10:", value)
}

// Define a function called 'demonstratePanic'  
// Print "Before panic"
// Call panic("Something went wrong!")
// Print "After panic" (this should not execute)
func demonstratePanic() {
	fmt.Println("Before panic")
	panic("Something went wrong!")
	fmt.Println("After panic") // This will not execute
}

// Define a function called 'nestedPanic'
// Defer a recovery function that prints "Outer recovery:", recover()
// Call a function that:
//   - Defers a recovery function that prints "Inner recovery:", recover()
//   - Panics with "Inner panic"
// Note: only the innermost recover catches the panic
func nestedPanic() {
	defer func() {
		fmt.Println("Outer recovery:", recover()) // This will be nil
	}()
	
	func() {
		defer func() {
			fmt.Println("Inner recovery:", recover()) // This catches the panic
		}()
		panic("Inner panic")
	}()
}

func main() {
	// Test safeDivide with normal division: safeDivide(10, 2)
	fmt.Println("=== safeDivide normal case ===")
	result1 := safeDivide(10, 2)
	fmt.Println("Result:", result1)
	
	// Test safeDivide with division by zero: safeDivide(10, 0)
	fmt.Println("\n=== safeDivide panic case ===")
	result2 := safeDivide(10, 0)
	fmt.Println("Result:", result2)
	
	// Test processArray with short array: []int{1, 2, 3}
	fmt.Println("\n=== processArray with short array ===")
	shortArr := []int{1, 2, 3}
	processArray(shortArr)
	
	// Test processArray with long array: make([]int, 15) (fill with some values)
	fmt.Println("\n=== processArray with long array ===")
	longArr := make([]int, 15)
	for i := 0; i < 15; i++ {
		longArr[i] = i * 10
	}
	processArray(longArr)
	
	// Call demonstratePanic inside a function with defer/recover
	// Use an anonymous function with defer/recover to catch the panic
	fmt.Println("\n=== demonstratePanic with recovery ===")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Caught panic from demonstratePanic:", r)
			}
		}()
		demonstratePanic()
	}()
	
	// Call nestedPanic to see nested recovery behavior
	fmt.Println("\n=== nestedPanic example ===")
	nestedPanic()
	
	// Demonstrate that the program continues after recovered panics
	// Print "Program continues normally"
	fmt.Println("\nProgram continues normally")
}