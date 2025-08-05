// generic_basics.go
// Learn the fundamentals of generics in Go (Go 1.18+)

package main

import "fmt"

// TODO: Generic function with type parameter
func Max[T comparable](a, b T) T {
	// Return the maximum of a and b
	// Use > operator (works with comparable types)
}

// TODO: Generic slice function
func Contains[T comparable](slice []T, item T) bool {
	// Check if slice contains item
	// Use == operator to compare
}

// TODO: Generic data structure
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	// Add item to stack
}

func (s *Stack[T]) Pop() (T, bool) {
	// Remove and return top item
	// Return zero value and false if empty
}

func (s *Stack[T]) Peek() (T, bool) {
	// Return top item without removing
	// Return zero value and false if empty
}

func (s *Stack[T]) Size() int {
	// Return number of items
}

// TODO: Generic map function
func MapSlice[T, U any](slice []T, fn func(T) U) []U {
	// Apply function to each element and return new slice
}

// TODO: Generic filter function  
func Filter[T any](slice []T, predicate func(T) bool) []T {
	// Return new slice with elements that match predicate
}

func main() {
	fmt.Println("=== Generic Functions ===")
	
	// TODO: Test Max with different types
	fmt.Printf("Max(5, 10): %d\\n", /* call Max with ints */)
	fmt.Printf("Max(3.14, 2.71): %.2f\\n", /* call Max with floats */)
	fmt.Printf("Max(\"apple\", \"banana\"): %s\\n", /* call Max with strings */)
	
	fmt.Println("\\n=== Generic Slice Operations ===")
	
	// TODO: Test Contains with different types
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Contains 3: %t\\n", /* check if numbers contains 3 */)
	fmt.Printf("Contains 6: %t\\n", /* check if numbers contains 6 */)
	
	words := []string{"hello", "world", "go"}
	fmt.Printf("Contains \\\"go\\\": %t\\n", /* check if words contains "go" */)
	
	fmt.Println("\\n=== Generic Stack ===")
	
	// TODO: Test Stack with integers
	intStack := /* create Stack of int */
	
	// Push some values
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	
	fmt.Printf("Stack size: %d\\n", intStack.Size())
	
	// Peek at top
	if top, ok := intStack.Peek(); ok {
		fmt.Printf("Top item: %d\\n", top)
	}
	
	// Pop all items
	for intStack.Size() > 0 {
		if item, ok := intStack.Pop(); ok {
			fmt.Printf("Popped: %d\\n", item)
		}
	}
	
	// TODO: Test Stack with strings
	stringStack := /* create Stack of string */
	stringStack.Push("first")
	stringStack.Push("second")
	stringStack.Push("third")
	
	fmt.Printf("\\nString stack contents:\\n")
	for stringStack.Size() > 0 {
		if item, ok := stringStack.Pop(); ok {
			fmt.Printf("Popped: %s\\n", item)
		}
	}
	
	fmt.Println("\\n=== Generic Higher-Order Functions ===")
	
	// TODO: Test MapSlice
	numbers = []int{1, 2, 3, 4, 5}
	
	// Map numbers to their squares
	squares := /* map numbers to squares using MapSlice */
	fmt.Printf("Squares: %v\\n", squares)
	
	// Map numbers to strings
	numStrings := /* map numbers to strings using MapSlice */
	fmt.Printf("Number strings: %v\\n", numStrings)
	
	// TODO: Test Filter
	allNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// Filter even numbers
	evens := /* filter even numbers */
	fmt.Printf("Even numbers: %v\\n", evens)
	
	// Filter numbers greater than 5
	greaterThan5 := /* filter numbers > 5 */
	fmt.Printf("Numbers > 5: %v\\n", greaterThan5)
	
	// Filter words by length
	allWords := []string{"a", "go", "test", "generics", "programming"}
	longWords := /* filter words with length > 3 */
	fmt.Printf("Long words: %v\\n", longWords)
}