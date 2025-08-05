// benchmarks_test.go - SOLUTION
// Benchmark tests for performance measurement

package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// String concatenation benchmarks
func BenchmarkStringConcatPlus(b *testing.B) {
	strs := []string{"hello", "world", "go", "programming", "language"}
	
	// Reset timer to exclude setup time
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		// Call StringConcatPlus
		StringConcatPlus(strs)
	}
}

func BenchmarkStringConcatBuilder(b *testing.B) {
	strs := []string{"hello", "world", "go", "programming", "language"}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		// Call StringConcatBuilder
		StringConcatBuilder(strs)
	}
}

func BenchmarkStringConcatJoin(b *testing.B) {
	strs := []string{"hello", "world", "go", "programming", "language"}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		// Call StringConcatJoin
		StringConcatJoin(strs)
	}
}

// Fibonacci benchmarks
func BenchmarkFibonacciRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Call FibonacciRecursive(20) - small number to avoid timeout
		FibonacciRecursive(20)
	}
}

func BenchmarkFibonacciIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Call FibonacciIterative(20)
		FibonacciIterative(20)
	}
}

func BenchmarkFibonacciMemoized(b *testing.B) {
	// Clear cache before benchmark
	fibCache = make(map[int]int)
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		// Call FibonacciMemoized(20)
		FibonacciMemoized(20)
	}
}

// Slice allocation benchmarks
func BenchmarkSliceAppendPrealloc(b *testing.B) {
	const size = 1000
	
	for i := 0; i < b.N; i++ {
		// Call SliceAppendPrealloc(size)
		SliceAppendPrealloc(size)
	}
}

func BenchmarkSliceAppendGrow(b *testing.B) {
	const size = 1000
	
	for i := 0; i < b.N; i++ {
		// Call SliceAppendGrow(size)
		SliceAppendGrow(size)
	}
}

// Search algorithm benchmarks
func BenchmarkLinearSearch(b *testing.B) {
	// Setup: create slice with 1000 elements
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i * 2
	}
	target := 500 // Middle element
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		// Call LinearSearch
		LinearSearch(slice, target)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	// Setup: create sorted slice with 1000 elements
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i * 2
	}
	target := 500 // Middle element
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		// Call BinarySearch
		BinarySearch(slice, target)
	}
}

// Map vs Slice lookup benchmarks
func BenchmarkMapLookup(b *testing.B) {
	// Setup: create map with 1000 entries
	m := make(map[string]int)
	keys := make([]string, 100)
	
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key_%d", i)
		m[key] = i
		if i < 100 {
			keys[i] = key
		}
	}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		// Call MapLookup
		MapLookup(m, keys)
	}
}

func BenchmarkSliceLookup(b *testing.B) {
	// Setup: create slice with 1000 entries
	slice := make([]string, 1000)
	keys := make([]string, 100)
	
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("key_%d", i)
		slice[i] = key
		if i < 100 {
			keys[i] = key
		}
	}
	
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		// Call SliceLookup
		SliceLookup(slice, keys)
	}
}

// Sub-benchmarks for different sizes
func BenchmarkStringConcat(b *testing.B) {
	sizes := []int{10, 100, 1000}
	
	for _, size := range sizes {
		// Create strings for this size
		strs := make([]string, size)
		for i := range strs {
			strs[i] = fmt.Sprintf("string_%d", i)
		}
		
		// Sub-benchmark for Plus method
		b.Run(fmt.Sprintf("Plus_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Call StringConcatPlus
				StringConcatPlus(strs)
			}
		})
		
		// Sub-benchmark for Builder method
		b.Run(fmt.Sprintf("Builder_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Call StringConcatBuilder
				StringConcatBuilder(strs)
			}
		})
		
		// Sub-benchmark for Join method
		b.Run(fmt.Sprintf("Join_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Call StringConcatJoin
				StringConcatJoin(strs)
			}
		})
	}
}

// Memory allocation benchmarks
func BenchmarkSliceGrowth(b *testing.B) {
	sizes := []int{100, 1000, 10000}
	
	for _, size := range sizes {
		// Benchmark without preallocation
		b.Run(fmt.Sprintf("Growing_%d", size), func(b *testing.B) {
			b.ReportAllocs() // Report memory allocations
			
			for i := 0; i < b.N; i++ {
				// Call SliceAppendGrow
				SliceAppendGrow(size)
			}
		})
		
		// Benchmark with preallocation
		b.Run(fmt.Sprintf("Prealloc_%d", size), func(b *testing.B) {
			b.ReportAllocs()
			
			for i := 0; i < b.N; i++ {
				// Call SliceAppendPrealloc
				SliceAppendPrealloc(size)
			}
		})
	}
}

// Setup and teardown benchmark
func BenchmarkWithSetup(b *testing.B) {
	// This benchmark demonstrates setup/teardown
	
	// Setup (not timed)
	rand.Seed(time.Now().UnixNano())
	data := make([]int, 1000)
	for i := range data {
		data[i] = rand.Intn(10000)
	}
	
	// Reset timer after setup
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		// The actual work being benchmarked
		target := rand.Intn(10000)
		// Call LinearSearch with data and target
		LinearSearch(data, target)
		
		// Stop timer if we need to do cleanup
		b.StopTimer()
		// Any cleanup work here (not timed)
		b.StartTimer()
	}
}

// Parallel benchmark
func BenchmarkParallelWork(b *testing.B) {
	// This benchmark runs work in parallel
	b.RunParallel(func(pb *testing.PB) {
		// Each goroutine runs this function
		strs := []string{"hello", "world", "parallel", "benchmark"}
		
		for pb.Next() {
			// Call StringConcatBuilder in parallel
			StringConcatBuilder(strs)
		}
	})
}