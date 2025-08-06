// sorting_comparison.go
// Compare performance characteristics of different sorting algorithms

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	fmt.Println("=== Sorting Algorithm Performance Comparison ===")
	
	// TODO: Define test data sizes
	dataSizes := /* create slice with different sizes to test */
	
	// TODO: Define sorting algorithms to compare
	algorithms := map[string]func([]int){
		"Built-in sort": /* use Go's built-in sort */,
		"Bubble Sort": /* implement bubble sort */,
		"Quick Sort": /* implement quick sort */,
		"Merge Sort": /* implement merge sort */,
		"Heap Sort": /* implement heap sort */,
	}
	
	fmt.Println("Performance Analysis:")
	fmt.Printf("%-15s", "Data Size")
	for name := range algorithms {
		fmt.Printf("%-15s", name)
	}
	fmt.Println()
	
	// TODO: Test each algorithm with different data sizes
	for _, size := range dataSizes {
		fmt.Printf("%-15d", size)
		
		for name, sortFunc := range algorithms {
			// TODO: Generate test data for this size
			testData := /* generate random data */
			
			// TODO: Time the sorting algorithm
			start := time.Now()
			/* call sorting algorithm */
			duration := time.Since(start)
			
			// TODO: Verify the data is sorted
			isCorrect := /* verify sorted */
			
			// TODO: Display timing with status indicator
			status := "✓"
			if !isCorrect {
				status = "✗"
			}
			
			fmt.Printf("%-15s", fmt.Sprintf("%s %v", status, duration))
		}
		fmt.Println()
	}
	
	fmt.Println("\n=== Best Case vs Worst Case Analysis ===")
	
	// TODO: Test algorithms on different data patterns
	testSize := 1000
	patterns := map[string][]int{
		"Random":           /* random data */,
		"Already Sorted":   /* sorted data */,
		"Reverse Sorted":   /* reverse sorted */,
		"Nearly Sorted":    /* 90% sorted */,
		"Many Duplicates":  /* lots of duplicates */,
		"Single Value":     /* all same value */,
	}
	
	fmt.Printf("%-15s", "Pattern")
	algorithmNames := []string{"Built-in", "Bubble", "Quick", "Merge", "Heap"}
	for _, name := range algorithmNames {
		fmt.Printf("%-12s", name)
	}
	fmt.Println()
	
	for patternName, data := range patterns {
		fmt.Printf("%-15s", patternName)
		
		// TODO: Test each algorithm on this pattern
		algorithmFuncs := []func([]int){
			/* Go's sort.Ints */,
			bubbleSort,
			quickSort,
			mergeSort,
			heapSort,
		}
		
		for _, sortFunc := range algorithmFuncs {
			// TODO: Copy data for each test
			testData := /* copy data */
			
			// TODO: Time the algorithm
			start := time.Now()
			/* call sorting function */
			duration := time.Since(start)
			
			fmt.Printf("%-12s", duration.String())
		}
		fmt.Println()
	}
	
	fmt.Println("\n=== Memory Usage Analysis ===")
	// TODO: Analyze in-place vs out-of-place sorting algorithms
	fmt.Println("In-place algorithms (O(1) extra space):")
	fmt.Println("  - Bubble Sort, Selection Sort, Insertion Sort")
	fmt.Println("  - Quick Sort (O(log n) stack space)")
	fmt.Println("  - Heap Sort")
	
	fmt.Println("\nOut-of-place algorithms (O(n) extra space):")
	fmt.Println("  - Merge Sort")
	fmt.Println("  - Counting Sort, Radix Sort")
	
	fmt.Println("\n=== Stability Analysis ===")
	// TODO: Test algorithm stability with custom data
	type Student struct {
		Name  string
		Grade int
	}
	
	students := []Student{
		/* TODO: create test data with duplicate grades */
	}
	
	fmt.Printf("Original order: %v\n", students)
	
	// TODO: Test stability of different algorithms
	// Stable sorts preserve relative order of equal elements
}

// TODO: Implement Bubble Sort (in-place, stable, O(n²))
func bubbleSort(arr []int) {
	// TODO: Implement bubble sort that sorts in-place
}

// TODO: Implement Quick Sort (in-place, unstable, O(n log n) avg, O(n²) worst)
func quickSort(arr []int) {
	// TODO: Implement quicksort with partitioning
}

// TODO: Implement Merge Sort (out-of-place, stable, O(n log n))
func mergeSort(arr []int) {
	// TODO: Implement merge sort - needs helper for recursion
}

// TODO: Helper for merge sort
func mergeSortHelper(arr []int, temp []int, left, right int) {
	// TODO: Recursive merge sort implementation
}

// TODO: Merge function for merge sort
func merge(arr []int, temp []int, left, mid, right int) {
	// TODO: Merge two sorted subarrays
}

// TODO: Implement Heap Sort (in-place, unstable, O(n log n))
func heapSort(arr []int) {
	// TODO: Build max heap then extract elements
}

// TODO: Heapify for heap sort
func heapify(arr []int, n, i int) {
	// TODO: Maintain heap property
}

// TODO: Utility functions

func generateRandomData(size int) []int {
	// TODO: Generate random integers
}

func generateSortedData(size int) []int {
	// TODO: Generate already sorted data
}

func generateReverseSortedData(size int) []int {
	// TODO: Generate reverse sorted data
}

func generateNearlySortedData(size int) []int {
	// TODO: Generate mostly sorted data
}

func generateDataWithDuplicates(size int) []int {
	// TODO: Generate data with many duplicate values
}

func generateSingleValueData(size int) []int {
	// TODO: Generate data with all same values
}

func copySlice(src []int) []int {
	// TODO: Create deep copy of slice
}

func isSorted(arr []int) bool {
	// TODO: Check if array is sorted in ascending order
}

func swap(arr []int, i, j int) {
	// TODO: Swap elements at positions i and j
}

func partition(arr []int, low, high int) int {
	// TODO: Partition array for quicksort
}