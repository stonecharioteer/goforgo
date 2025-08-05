// map_performance.go - SOLUTION
// Learn about map performance, memory usage, and optimization techniques
// Understand when to use maps vs slices and other data structures

package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a map with initial capacity using make
	largeMap := make(map[string]int, 1000)
	
	// Benchmark map insertion performance
	fmt.Println("Benchmarking map insertion...")
	
	start := time.Now()
	// Insert 10000 key-value pairs
	for i := 0; i < 10000; i++ {
		largeMap[fmt.Sprintf("key%d", i)] = i * 2
	}
	insertTime := time.Since(start)
	fmt.Printf("Time to insert 10,000 items: %v\n", insertTime)
	
	// Benchmark map lookup performance
	fmt.Println("\nBenchmarking map lookup...")
	
	start = time.Now()
	// Perform 100,000 lookups
	var found int
	for i := 0; i < 100000; i++ {
		// Look up random keys and count how many are found
		key := fmt.Sprintf("key%d", i%10000)
		if _, exists := largeMap[key]; exists {
			found++
		}
	}
	lookupTime := time.Since(start)
	fmt.Printf("Time for 100,000 lookups: %v (found %d items)\n", lookupTime, found)
	
	// Compare map vs slice for membership testing
	fmt.Println("\nComparing map vs slice for membership testing...")
	
	// Create a slice with same data
	var slice []string
	for i := 0; i < 10000; i++ {
		slice = append(slice, fmt.Sprintf("key%d", i))
	}
	
	// Test finding an element in slice (linear search)
	target := "key5000"
	
	start = time.Now()
	var foundInSlice bool
	// Linear search in slice
	for _, item := range slice {
		if item == target {
			foundInSlice = true
			break
		}
	}
	sliceSearchTime := time.Since(start)
	
	// Test finding same element in map
	start = time.Now()
	_, foundInMap := largeMap[target]
	mapSearchTime := time.Since(start)
	
	fmt.Printf("Slice search time: %v (found: %t)\n", sliceSearchTime, foundInSlice)
	fmt.Printf("Map search time: %v (found: %t)\n", mapSearchTime, foundInMap)
	if mapSearchTime.Nanoseconds() > 0 {
		fmt.Printf("Map is %dx faster for lookup\n", sliceSearchTime.Nanoseconds()/mapSearchTime.Nanoseconds())
	}
	
	// Demonstrate map memory optimization
	fmt.Println("\nDemonstrating map memory patterns...")
	
	// Create map and fill it
	memoryMap := make(map[int]string)
	for i := 0; i < 1000; i++ {
		memoryMap[i] = fmt.Sprintf("value_%d", i)
	}
	fmt.Printf("Map with 1000 items created\n")
	
	// Delete most items
	for i := 0; i < 900; i++ {
		delete(memoryMap, i)
	}
	fmt.Printf("Deleted 900 items, %d items remaining\n", len(memoryMap))
	
	// Show that map retains memory even after deletions
	// Note: In real code, you might need to create a new map to free memory
	fmt.Println("Note: Maps don't automatically shrink memory after deletions")
	
	// Demonstrate map iteration performance
	fmt.Println("\nMap iteration performance...")
	
	iterMap := make(map[string]int)
	for i := 0; i < 10000; i++ {
		iterMap[fmt.Sprintf("item%d", i)] = i
	}
	
	start = time.Now()
	var sum int
	// Iterate through all key-value pairs and sum the values
	for _, value := range iterMap {
		sum += value
	}
	iterTime := time.Since(start)
	fmt.Printf("Time to iterate and sum 10,000 items: %v (sum: %d)\n", iterTime, sum)
	
	// Show key collision example (advanced)
	fmt.Println("\nDemonstrating string keys performance...")
	
	// Test with different key patterns
	shortKeys := make(map[string]int)
	longKeys := make(map[string]int)
	
	// Short keys
	start = time.Now()
	for i := 0; i < 1000; i++ {
		shortKeys[fmt.Sprintf("%d", i)] = i
	}
	shortKeyTime := time.Since(start)
	
	// Long keys
	start = time.Now()
	for i := 0; i < 1000; i++ {
		longKeys[fmt.Sprintf("very_long_key_name_with_lots_of_characters_%d", i)] = i
	}
	longKeyTime := time.Since(start)
	
	fmt.Printf("Short keys insertion time: %v\n", shortKeyTime)
	fmt.Printf("Long keys insertion time: %v\n", longKeyTime)
	
	// Best practices summary
	fmt.Println("\n=== Map Performance Best Practices ===")
	fmt.Println("1. Use make(map[K]V, capacity) when you know approximate size")
	fmt.Println("2. Maps are faster than slices for lookups (O(1) vs O(n))")
	fmt.Println("3. Maps don't shrink memory after deletions")
	fmt.Println("4. Shorter keys are generally faster to hash")
	fmt.Println("5. Use maps for: lookups, counting, grouping")
	fmt.Println("6. Use slices for: ordered data, iteration, memory efficiency")
}