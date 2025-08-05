// benchmarks.go - SOLUTION
// Learn how to write benchmarks and measure performance in Go

package main

import (
	"fmt"
	"strings"
)

// Functions to benchmark

// StringConcatenation using + operator
func StringConcatPlus(strs []string) string {
	var result string
	// Concatenate using + operator
	for _, s := range strs {
		result += s
	}
	return result
}

// StringConcatenation using strings.Builder
func StringConcatBuilder(strs []string) string {
	var builder strings.Builder
	// Concatenate using strings.Builder
	for _, s := range strs {
		builder.WriteString(s)
	}
	return builder.String()
}

// StringConcatenation using strings.Join
func StringConcatJoin(strs []string) string {
	// Concatenate using strings.Join
	return strings.Join(strs, "")
}

// Fibonacci implementations for benchmarking
func FibonacciRecursive(n int) int {
	// Implement recursive fibonacci (inefficient)
	if n <= 1 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}

func FibonacciIterative(n int) int {
	// Implement iterative fibonacci (efficient)
	if n <= 1 {
		return n
	}
	
	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}

// Fibonacci with memoization
var fibCache = make(map[int]int)

func FibonacciMemoized(n int) int {
	// Check cache first, compute if not found
	if n <= 1 {
		return n
	}
	
	if val, exists := fibCache[n]; exists {
		return val
	}
	
	result := FibonacciMemoized(n-1) + FibonacciMemoized(n-2)
	fibCache[n] = result
	return result
}

// Slice operations for benchmarking
func SliceAppendPrealloc(n int) []int {
	// Pre-allocate slice with known capacity
	slice := make([]int, 0, n)
	// Append n elements
	for i := 0; i < n; i++ {
		slice = append(slice, i)
	}
	return slice
}

func SliceAppendGrow(n int) []int {
	// Start with empty slice and let it grow
	var slice []int
	// Append n elements
	for i := 0; i < n; i++ {
		slice = append(slice, i)
	}
	return slice
}

// Search algorithms
func LinearSearch(slice []int, target int) int {
	// Implement linear search
	// Return index if found, -1 if not found
	for i, val := range slice {
		if val == target {
			return i
		}
	}
	return -1
}

func BinarySearch(slice []int, target int) int {
	// Implement binary search (assume slice is sorted)
	// Return index if found, -1 if not found
	left, right := 0, len(slice)-1
	
	for left <= right {
		mid := left + (right-left)/2
		
		if slice[mid] == target {
			return mid
		} else if slice[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// Map vs Slice lookup
func MapLookup(m map[string]int, keys []string) int {
	count := 0
	// Look up each key in map
	for _, key := range keys {
		if _, exists := m[key]; exists {
			count++
		}
	}
	return count
}

func SliceLookup(slice []string, keys []string) int {
	count := 0
	// Look up each key in slice (linear search)
	for _, key := range keys {
		for _, item := range slice {
			if item == key {
				count++
				break
			}
		}
	}
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