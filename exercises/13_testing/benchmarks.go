// benchmarks.go
// Learn how to write benchmarks and measure performance in Go

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// TODO: Functions to benchmark

// StringConcatenation using + operator
func StringConcatPlus(strs []string) string {
	var result string
	// Concatenate using + operator
	return result
}

// StringConcatenation using strings.Builder
func StringConcatBuilder(strs []string) string {
	var builder strings.Builder
	// Concatenate using strings.Builder
	return builder.String()
}

// StringConcatenation using strings.Join
func StringConcatJoin(strs []string) string {
	// Concatenate using strings.Join
}

// Fibonacci implementations for benchmarking
func FibonacciRecursive(n int) int {
	// Implement recursive fibonacci (inefficient)
}

func FibonacciIterative(n int) int {
	// Implement iterative fibonacci (efficient)
}

// Fibonacci with memoization
var fibCache = make(map[int]int)

func FibonacciMemoized(n int) int {
	// Check cache first, compute if not found
}

// Slice operations for benchmarking
func SliceAppendPrealloc(n int) []int {
	// Pre-allocate slice with known capacity
	slice := make([]int, 0, n)
	// Append n elements
	return slice
}

func SliceAppendGrow(n int) []int {
	// Start with empty slice and let it grow
	var slice []int
	// Append n elements
	return slice
}

// Search algorithms
func LinearSearch(slice []int, target int) int {
	// Implement linear search
	// Return index if found, -1 if not found
}

func BinarySearch(slice []int, target int) int {
	// Implement binary search (assume slice is sorted)
	// Return index if found, -1 if not found
}

// Map vs Slice lookup
func MapLookup(m map[string]int, keys []string) int {
	count := 0
	// Look up each key in map
	return count
}

func SliceLookup(slice []string, keys []string) int {
	count := 0
	// Look up each key in slice (linear search)
	return count
}

func main() {
	// This file contains functions to be benchmarked
	// The actual benchmarks will be in benchmarks_test.go
	
	fmt.Println("This file contains functions to be benchmarked.")
	fmt.Println("Run 'go test -bench=.' to execute the benchmarks.")
	
	// Demo the functions
	strs := []string{"hello", "world", "go", "benchmarks"}
	
	fmt.Printf("StringConcatPlus: %s\n", StringConcatPlus(strs))
	fmt.Printf("StringConcatBuilder: %s\n", StringConcatBuilder(strs))
	fmt.Printf("StringConcatJoin: %s\n", StringConcatJoin(strs))
	
	fmt.Printf("FibonacciRecursive(10): %d\n", FibonacciRecursive(10))
	fmt.Printf("FibonacciIterative(10): %d\n", FibonacciIterative(10))
	fmt.Printf("FibonacciMemoized(10): %d\n", FibonacciMemoized(10))
	
	slice1 := SliceAppendPrealloc(5)
	slice2 := SliceAppendGrow(5)
	fmt.Printf("SliceAppendPrealloc: %v\n", slice1)
	fmt.Printf("SliceAppendGrow: %v\n", slice2)
	
	testSlice := []int{1, 3, 5, 7, 9, 11, 13, 15}
	fmt.Printf("LinearSearch(7): %d\n", LinearSearch(testSlice, 7))
	fmt.Printf("BinarySearch(7): %d\n", BinarySearch(testSlice, 7))
}