// sorting_algorithms.go - SOLUTION
// Learn various sorting algorithm implementations

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== Sorting Algorithms Implementation ===")
	
	// Generate test data
	testSizes := []int{10, 50, 100}
	
	for _, size := range testSizes {
		fmt.Printf("\n=== Testing with %d elements ===\n", size)
		
		// Generate random data
		data := generateRandomData(size)
		fmt.Printf("Original data: %v", data[:min(10, len(data))])
		if len(data) > 10 {
			fmt.Printf("... (%d total elements)", len(data))
		}
		fmt.Println()
		
		// Test different sorting algorithms
		algorithms := []struct {
			name string
			fn   func([]int) []int
		}{
			{"Bubble Sort", bubbleSort},
			{"Selection Sort", selectionSort},
			{"Insertion Sort", insertionSort},
			{"Merge Sort", mergeSort},
			{"Quick Sort", quickSort},
			{"Heap Sort", heapSort},
		}
		
		fmt.Println("\nSorting results:")
		for _, alg := range algorithms {
			// Copy data for each algorithm
			dataCopy := makeCopy(data)
			
			// Time the sorting
			start := time.Now()
			sorted := alg.fn(dataCopy)
			duration := time.Since(start)
			
			// Verify sorting is correct
			isCorrect := isSorted(sorted)
			
			// Display first few elements
			preview := sorted[:min(10, len(sorted))]
			status := "✓"
			if !isCorrect {
				status = "✗"
			}
			
			fmt.Printf("  %s %s: %v", status, alg.name, preview)
			if len(sorted) > 10 {
				fmt.Printf("...")
			}
			fmt.Printf(" (%v)\n", duration)
		}
	}
	
	fmt.Println("\n=== Algorithm Comparison ===")
	
	// Compare algorithms on different data patterns
	patterns := map[string][]int{
		"Random":          generateRandomData(1000),
		"Already Sorted":  generateSortedData(1000),
		"Reverse Sorted":  generateReverseSortedData(1000),
		"Nearly Sorted":   generateNearlySortedData(1000),
		"Many Duplicates": generateDataWithDuplicates(1000),
	}
	
	fmt.Println("Performance comparison (1000 elements):")
	fmt.Printf("%-15s", "Pattern")
	algorithms := []string{"Bubble", "Selection", "Insertion", "Merge", "Quick", "Heap"}
	for _, alg := range algorithms {
		fmt.Printf("%-12s", alg)
	}
	fmt.Println()
	
	for patternName, data := range patterns {
		fmt.Printf("%-15s", patternName)
		
		// Test each algorithm on this pattern
		algFuncs := []func([]int) []int{bubbleSort, selectionSort, insertionSort, mergeSort, quickSort, heapSort}
		
		for _, algFunc := range algFuncs {
			// Time the algorithm
			dataCopy := makeCopy(data)
			start := time.Now()
			algFunc(dataCopy)
			duration := time.Since(start)
			
			fmt.Printf("%-12s", duration.String())
		}
		fmt.Println()
	}
}

// Implement Bubble Sort
func bubbleSort(arr []int) []int {
	result := makeCopy(arr)
	n := len(result)
	
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				swap(result, j, j+1)
			}
		}
	}
	
	return result
}

// Implement Selection Sort
func selectionSort(arr []int) []int {
	result := makeCopy(arr)
	n := len(result)
	
	for i := 0; i < n-1; i++ {
		// Find minimum element in remaining array
		minIdx := i
		for j := i + 1; j < n; j++ {
			if result[j] < result[minIdx] {
				minIdx = j
			}
		}
		
		// Swap minimum with current position
		swap(result, i, minIdx)
	}
	
	return result
}

// Implement Insertion Sort
func insertionSort(arr []int) []int {
	result := makeCopy(arr)
	
	for i := 1; i < len(result); i++ {
		key := result[i]
		j := i - 1
		
		// Shift elements and insert
		for j >= 0 && result[j] > key {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}
	
	return result
}

// Implement Merge Sort
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return makeCopy(arr)
	}
	
	// Divide array
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	
	// Merge sorted halves
	return merge(left, right)
}

// Helper function for merge sort
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	
	// Merge elements in sorted order
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	
	// Append remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	
	return result
}

// Implement Quick Sort
func quickSort(arr []int) []int {
	result := makeCopy(arr)
	quickSortHelper(result, 0, len(result)-1)
	return result
}

// Helper function for quick sort
func quickSortHelper(arr []int, low, high int) {
	if low < high {
		// Partition array and get pivot index
		pi := partition(arr, low, high)
		
		// Recursively sort elements before and after partition
		quickSortHelper(arr, low, pi-1)
		quickSortHelper(arr, pi+1, high)
	}
}

// Partition function for quick sort
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			swap(arr, i, j)
		}
	}
	
	swap(arr, i+1, high)
	return i + 1
}

// Implement Heap Sort
func heapSort(arr []int) []int {
	result := makeCopy(arr)
	n := len(result)
	
	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(result, n, i)
	}
	
	// Extract elements from heap
	for i := n - 1; i > 0; i-- {
		swap(result, 0, i)
		heapify(result, i, 0)
	}
	
	return result
}

// Heapify function for heap sort
func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	
	// Find largest among root, left child, right child
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	
	// If largest is not root, swap and continue heapifying
	if largest != i {
		swap(arr, i, largest)
		heapify(arr, n, largest)
	}
}

// Utility functions

func generateRandomData(size int) []int {
	rand.Seed(time.Now().UnixNano())
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(1000)
	}
	return data
}

func generateSortedData(size int) []int {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = i
	}
	return data
}

func generateReverseSortedData(size int) []int {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = size - i - 1
	}
	return data
}

func generateNearlySortedData(size int) []int {
	data := generateSortedData(size)
	// Shuffle 10% of elements
	swaps := size / 10
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < swaps; i++ {
		i1 := rand.Intn(size)
		i2 := rand.Intn(size)
		swap(data, i1, i2)
	}
	return data
}

func generateDataWithDuplicates(size int) []int {
	data := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	// Use only 10 different values
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(10)
	}
	return data
}

func makeCopy(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	return result
}

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}