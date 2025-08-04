package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// TODO: Use a for loop with break to find the first number greater than 5
	// Print "Found number greater than 5: [number]" and break
	
	// TODO: Use a for loop with continue to print only even numbers from numbers slice
	// Skip odd numbers using continue
	
	// TODO: Use nested loops with labeled break
	// Outer loop: i from 1 to 3
	// Inner loop: j from 1 to 3
	// Print i*j, but break out of BOTH loops when i*j >= 4
	// Use label: OuterLoop:
	
	// TODO: Use nested loops with labeled continue  
	// Outer loop: i from 1 to 3
	// Inner loop: j from 1 to 3
	// Skip printing when i == j (continue outer loop)
	// Otherwise print "i=%d, j=%d"
	// Use label: NextI:
	
	// TODO: Find and print all prime numbers from 2 to 20
	// Use nested loops: outer for candidate 2 to 20, inner to check divisors
	// Use continue to skip non-primes, break inner loop when divisor found
	
	// TODO: Implement a simple guessing game simulation
	// Secret number is 7
	// Loop through guesses: []int{3, 7, 5, 9, 7, 2}
	// Print each guess, break when correct guess found
	// Use continue to skip processing after printing wrong guesses
}