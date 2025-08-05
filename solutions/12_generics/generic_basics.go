// generic_basics.go - SOLUTION
// Learn the fundamentals of generics in Go (Go 1.18+)

package main

import "fmt"

// Generic function with type parameter
func Max[T comparable](a, b T) T {
	// Return the maximum of a and b
	// Use > operator (works with comparable types)
	if a > b {
		return a
	}
	return b
}

// Generic slice function
func Contains[T comparable](slice []T, item T) bool {
	// Check if slice contains item
	// Use == operator to compare
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Generic data structure
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	// Add item to stack
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	// Remove and return top item
	// Return zero value and false if empty
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
	// Return top item without removing
	// Return zero value and false if empty
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) Size() int {
	// Return number of items
	return len(s.items)
}

// Generic map function
func MapSlice[T, U any](slice []T, fn func(T) U) []U {
	// Apply function to each element and return new slice
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Generic filter function  
func Filter[T any](slice []T, predicate func(T) bool) []T {
	// Return new slice with elements that match predicate
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	fmt.Println("=== Generic Functions ===")
	
	// Test Max with different types
	fmt.Printf("Max(5, 10): %d\n", Max(5, 10))
	fmt.Printf("Max(3.14, 2.71): %.2f\n", Max(3.14, 2.71))
	fmt.Printf("Max(\"apple\", \"banana\"): %s\n", Max("apple", "banana"))
	
	fmt.Println("\n=== Generic Slice Operations ===")
	
	// Test Contains with different types
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Contains 3: %t\n", Contains(numbers, 3))
	fmt.Printf("Contains 6: %t\n", Contains(numbers, 6))
	
	words := []string{"hello", "world", "go"}
	fmt.Printf("Contains \"go\": %t\n", Contains(words, "go"))
	
	fmt.Println("\n=== Generic Stack ===")
	
	// Test Stack with integers
	intStack := &Stack[int]{}
	
	// Push some values
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	
	fmt.Printf("Stack size: %d\n", intStack.Size())
	
	// Peek at top
	if top, ok := intStack.Peek(); ok {
		fmt.Printf("Top item: %d\n", top)
	}
	
	// Pop all items
	for intStack.Size() > 0 {
		if item, ok := intStack.Pop(); ok {
			fmt.Printf("Popped: %d\n", item)
		}
	}
	
	// Test Stack with strings
	stringStack := &Stack[string]{}
	stringStack.Push("first")
	stringStack.Push("second")
	stringStack.Push("third")
	
	fmt.Printf("\nString stack contents:\n")
	for stringStack.Size() > 0 {
		if item, ok := stringStack.Pop(); ok {
			fmt.Printf("Popped: %s\n", item)
		}
	}
	
	fmt.Println("\n=== Generic Higher-Order Functions ===")
	
	// Test MapSlice
	numbers = []int{1, 2, 3, 4, 5}
	
	// Map numbers to their squares
	squares := MapSlice(numbers, func(n int) int { return n * n })
	fmt.Printf("Squares: %v\n", squares)
	
	// Map numbers to strings
	numStrings := MapSlice(numbers, func(n int) string { return fmt.Sprintf("num-%d", n) })
	fmt.Printf("Number strings: %v\n", numStrings)
	
	// Test Filter
	allNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	// Filter even numbers
	evens := Filter(allNumbers, func(n int) bool { return n%2 == 0 })
	fmt.Printf("Even numbers: %v\n", evens)
	
	// Filter numbers greater than 5
	greaterThan5 := Filter(allNumbers, func(n int) bool { return n > 5 })
	fmt.Printf("Numbers > 5: %v\n", greaterThan5)
	
	// Filter words by length
	allWords := []string{"a", "go", "test", "generics", "programming"}
	longWords := Filter(allWords, func(word string) bool { return len(word) > 3 })
	fmt.Printf("Long words: %v\n", longWords)
}