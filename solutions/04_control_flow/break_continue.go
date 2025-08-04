package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// Use a for loop with break to find the first number greater than 5
	// Print "Found number greater than 5: [number]" and break
	fmt.Println("Finding first number > 5:")
	for _, num := range numbers {
		if num > 5 {
			fmt.Printf("Found number greater than 5: %d\n", num)
			break
		}
	}
	
	// Use a for loop with continue to print only even numbers from numbers slice
	// Skip odd numbers using continue
	fmt.Println("\nEven numbers only:")
	for _, num := range numbers {
		if num%2 == 1 { // odd number
			continue
		}
		fmt.Println(num)
	}
	
	// Use nested loops with labeled break
	// Outer loop: i from 1 to 3
	// Inner loop: j from 1 to 3
	// Print i*j, but break out of BOTH loops when i*j >= 4
	// Use label: OuterLoop:
	fmt.Println("\nNested loops with labeled break:")
OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			product := i * j
			fmt.Printf("%d * %d = %d\n", i, j, product)
			if product >= 4 {
				fmt.Println("Product >= 4, breaking out of both loops")
				break OuterLoop
			}
		}
	}
	
	// Use nested loops with labeled continue  
	// Outer loop: i from 1 to 3
	// Inner loop: j from 1 to 3
	// Skip printing when i == j (continue outer loop)
	// Otherwise print "i=%d, j=%d"
	// Use label: NextI:
	fmt.Println("\nNested loops with labeled continue:")
NextI:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == j {
				fmt.Printf("Skipping i=%d, j=%d (equal values)\n", i, j)
				continue NextI
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
	
	// Find and print all prime numbers from 2 to 20
	// Use nested loops: outer for candidate 2 to 20, inner to check divisors
	// Use continue to skip non-primes, break inner loop when divisor found
	fmt.Println("\nPrime numbers from 2 to 20:")
	for candidate := 2; candidate <= 20; candidate++ {
		isPrime := true
		for divisor := 2; divisor*divisor <= candidate; divisor++ {
			if candidate%divisor == 0 {
				isPrime = false
				break // Not prime, break inner loop
			}
		}
		if !isPrime {
			continue // Skip non-primes
		}
		fmt.Println(candidate)
	}
	
	// Implement a simple guessing game simulation
	// Secret number is 7
	// Loop through guesses: []int{3, 7, 5, 9, 7, 2}
	// Print each guess, break when correct guess found
	// Use continue to skip processing after printing wrong guesses
	fmt.Println("\nGuessing game simulation:")
	secretNumber := 7
	guesses := []int{3, 7, 5, 9, 7, 2}
	
	for i, guess := range guesses {
		fmt.Printf("Guess %d: %d", i+1, guess)
		if guess == secretNumber {
			fmt.Println(" - Correct! You win!")
			break
		}
		fmt.Println(" - Wrong, try again")
		continue // Skip any additional processing for wrong guesses
	}
}