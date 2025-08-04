package main

import "fmt"

// Define a function called 'makeCounter' that returns a function
// The returned function should have no parameters and return an int
// Use a closure to create a counter that increments each time it's called
// The counter should start at 0 and increment by 1 each call
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Define a function called 'makeMultiplier' that takes an int parameter 'factor'
// It should return a function that takes an int and returns an int
// The returned function should multiply its input by the factor
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	// Create a counter using makeCounter
	counter1 := makeCounter()
	
	// Call the counter 3 times and print each result
	fmt.Println("Counter1 call 1:", counter1())
	fmt.Println("Counter1 call 2:", counter1())
	fmt.Println("Counter1 call 3:", counter1())
	
	// Create another counter using makeCounter
	counter2 := makeCounter()
	
	// Call the second counter 2 times and print each result
	// Note: each counter should maintain its own state
	fmt.Println("Counter2 call 1:", counter2())
	fmt.Println("Counter2 call 2:", counter2())
	
	// Create a multiplier that multiplies by 3 using makeMultiplier
	multiply3 := makeMultiplier(3)
	
	// Use the multiplier to multiply 5 and 10, print the results
	fmt.Println("3 * 5 =", multiply3(5))
	fmt.Println("3 * 10 =", multiply3(10))
	
	// Create a multiplier that multiplies by 7 using makeMultiplier
	multiply7 := makeMultiplier(7)
	
	// Use the second multiplier to multiply 4 and 6, print the results
	fmt.Println("7 * 4 =", multiply7(4))
	fmt.Println("7 * 6 =", multiply7(6))
}