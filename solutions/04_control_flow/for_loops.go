package main

import "fmt"

func main() {
	// Write a basic for loop that prints numbers from 1 to 5
	// Use: for i := 1; i <= 5; i++ { }
	fmt.Println("Numbers 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	
	// Write a for loop that prints even numbers from 2 to 10
	// Use: for i := 2; i <= 10; i += 2 { }
	fmt.Println("\nEven numbers 2 to 10:")
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}
	
	// Write a while-style for loop that counts down from 5 to 1
	// Initialize count := 5 before the loop
	// Use: for count > 0 { }
	// Remember to decrement count inside the loop
	fmt.Println("\nCountdown from 5:")
	count := 5
	for count > 0 {
		fmt.Println(count)
		count--
	}
	
	// Write an infinite for loop that breaks when i reaches 3
	// Use: for { } and break statement
	// Initialize i := 0 before the loop and increment inside
	fmt.Println("\nInfinite loop with break:")
	i := 0
	for {
		fmt.Println("i is", i)
		i++
		if i >= 3 {
			break
		}
	}
	
	// Write a for loop that prints numbers 1 to 10 but skips 5
	// Use continue statement when i == 5
	fmt.Println("\nNumbers 1 to 10, skipping 5:")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	
	// Write nested for loops to print a 3x3 multiplication table
	// Outer loop: i from 1 to 3
	// Inner loop: j from 1 to 3  
	// Print: fmt.Printf("%d x %d = %d\n", i, j, i*j)
	fmt.Println("\n3x3 multiplication table:")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
	}
}