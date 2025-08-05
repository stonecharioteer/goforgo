// number_theory.go - SOLUTION
// Implement mathematical algorithms and number theory functions

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Number Theory Algorithms ===")
	
	// Test prime number functions
	testNumbers := []int{2, 3, 4, 5, 17, 25, 29, 100, 101}
	
	fmt.Println("Prime number testing:")
	for _, num := range testNumbers {
		fmt.Printf("%d is prime: %t\n", num, isPrime(num))
	}
	
	// Generate prime numbers
	fmt.Printf("\nFirst 20 primes: %v\n", generatePrimes(20))
	
	// Test GCD and LCM
	pairs := []struct{ a, b int }{
		{12, 18}, {15, 25}, {7, 13}, {100, 75},
	}
	
	fmt.Println("\nGCD and LCM calculations:")
	for _, pair := range pairs {
		gcdResult := gcd(pair.a, pair.b)
		lcmResult := lcm(pair.a, pair.b)
		fmt.Printf("GCD(%d, %d) = %d, LCM(%d, %d) = %d\n", 
			pair.a, pair.b, gcdResult, pair.a, pair.b, lcmResult)
	}
	
	// Test factorial
	fmt.Println("\nFactorial calculations:")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d! = %d\n", i, factorial(i))
	}
	
	// Test Fibonacci
	fmt.Printf("\nFirst 15 Fibonacci numbers: %v\n", fibonacci(15))
	
	// Test perfect numbers
	fmt.Println("\nPerfect numbers up to 1000:")
	for i := 1; i <= 1000; i++ {
		if isPerfect(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

// Check if number is prime
func isPrime(n int) bool {
	// Handle edge cases
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	
	// Check divisibility up to sqrt(n)
	limit := int(math.Sqrt(float64(n)))
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Generate first n prime numbers
func generatePrimes(n int) []int {
	// Use trial division
	primes := make([]int, 0, n)
	num := 2
	
	for len(primes) < n {
		if isPrime(num) {
			primes = append(primes, num)
		}
		num++
	}
	
	return primes
}

// Calculate greatest common divisor
func gcd(a, b int) int {
	// Use Euclidean algorithm
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Calculate least common multiple
func lcm(a, b int) int {
	// Use formula: LCM(a,b) = (a*b) / GCD(a,b)
	return (a * b) / gcd(a, b)
}

// Calculate factorial
func factorial(n int) int64 {
	// Handle base case
	if n < 0 {
		return 0
	}
	if n == 0 || n == 1 {
		return 1
	}
	
	// Calculate factorial iteratively
	result := int64(1)
	for i := 2; i <= n; i++ {
		result *= int64(i)
	}
	return result
}

// Generate Fibonacci sequence
func fibonacci(n int) []int {
	// Generate first n Fibonacci numbers
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	
	fib := make([]int, n)
	fib[0], fib[1] = 0, 1
	
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	
	return fib
}

// Check if number is perfect
func isPerfect(n int) bool {
	// A perfect number equals sum of its proper divisors
	if n <= 1 {
		return false
	}
	
	// Find all divisors and sum them
	sum := 1 // 1 is always a proper divisor
	limit := int(math.Sqrt(float64(n)))
	
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i { // Avoid counting square root twice
				sum += n / i
			}
		}
	}
	
	return sum == n
}