// sorting_comparison.go - SOLUTION
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
	
	// Define test data sizes
	dataSizes := []int{100, 500, 1000, 5000, 10000}
	
	// Define sorting algorithms to compare
	algorithms := map[string]func([]int){
		"Built-in sort": func(arr []int) { sort.Ints(arr) },
		"Bubble Sort":   bubbleSort,
		"Quick Sort":    quickSort,
		"Merge Sort":    mergeSort,
		"Heap Sort":     heapSort,
	}
	
	fmt.Println("Performance Analysis:")
	fmt.Printf("%-15s", "Data Size")
	for name := range algorithms {
		fmt.Printf("%-15s", name)
	}
	fmt.Println()
	
	// Test each algorithm with different data sizes
	for _, size := range dataSizes {
		fmt.Printf("%-15d", size)
		
		for name, sortFunc := range algorithms {
			// Generate test data for this size
			testData := generateRandomData(size)
			
			// Time the sorting algorithm
			start := time.Now()
			sortFunc(testData)
			duration := time.Since(start)
			
			// Verify the data is sorted
			isCorrect := isSorted(testData)
			
			// Display timing with status indicator
			status := "✓"
			if !isCorrect {
				status = "✗"
			}
			
			fmt.Printf("%-15s", fmt.Sprintf("%s %v", status, duration))
		}
		fmt.Println()
	}
	
	fmt.Println("\n=== Best Case vs Worst Case Analysis ===")
	
	// Test algorithms on different data patterns
	testSize := 1000
	patterns := map[string][]int{
		"Random":          generateRandomData(testSize),
		"Already Sorted":  generateSortedData(testSize),
		"Reverse Sorted":  generateReverseSortedData(testSize),
		"Nearly Sorted":   generateNearlySortedData(testSize),
		"Many Duplicates": generateDataWithDuplicates(testSize),
		"Single Value":    generateSingleValueData(testSize),
	}
	
	fmt.Printf("%-15s", "Pattern")
	algorithmNames := []string{"Built-in", "Bubble", "Quick", "Merge", "Heap"}
	for _, name := range algorithmNames {
		fmt.Printf("%-12s", name)
	}
	fmt.Println()
	
	for patternName, data := range patterns {
		fmt.Printf("%-15s", patternName)
		
		// Test each algorithm on this pattern
		algorithmFuncs := []func([]int){
			func(arr []int) { sort.Ints(arr) },
			bubbleSort,
			quickSort,
			mergeSort,
			heapSort,
		}
		
		for _, sortFunc := range algorithmFuncs {
			// Copy data for each test
			testData := copySlice(data)
			
			// Time the algorithm
			start := time.Now()
			sortFunc(testData)
			duration := time.Since(start)
			
			fmt.Printf("%-12s", duration.String())
		}
		fmt.Println()
	}
	
	fmt.Println("\n=== Memory Usage Analysis ===")
	// Analyze in-place vs out-of-place sorting algorithms
	fmt.Println("In-place algorithms (O(1) extra space):")
	fmt.Println("  - Bubble Sort, Selection Sort, Insertion Sort")
	fmt.Println("  - Quick Sort (O(log n) stack space)")
	fmt.Println("  - Heap Sort")
	
	fmt.Println("\nOut-of-place algorithms (O(n) extra space):")
	fmt.Println("  - Merge Sort")
	fmt.Println("  - Counting Sort, Radix Sort")
	
	fmt.Println("\n=== Stability Analysis ===")
	// Test algorithm stability with custom data
	type Student struct {
		Name  string
		Grade int
	}
	
	students := []Student{
		{"Alice", 85}, {"Bob", 92}, {"Carol", 85}, {"David", 78}, {"Eve", 92},
	}
	
	fmt.Printf("Original order: %v\n", students)
	fmt.Println("Stable sorts preserve relative order of equal elements")
	fmt.Println("Unstable sorts may change relative order of equal elements")
}

// Implement Bubble Sort (in-place, stable, O(n²))
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
}

// Implement Quick Sort (in-place, unstable, O(n log n) avg, O(n²) worst)
func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSortHelper(arr, low, pi-1)
		quickSortHelper(arr, pi+1, high)
	}
}

// Implement Merge Sort (out-of-place, stable, O(n log n))
func mergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	temp := make([]int, len(arr))
	mergeSortHelper(arr, temp, 0, len(arr)-1)
}

// Helper for merge sort
func mergeSortHelper(arr []int, temp []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		mergeSortHelper(arr, temp, left, mid)
		mergeSortHelper(arr, temp, mid+1, right)
		merge(arr, temp, left, mid, right)
	}
}

// Merge function for merge sort
func merge(arr []int, temp []int, left, mid, right int) {
	i, j, k := left, mid+1, left
	
	// Copy data to temp arrays
	for idx := left; idx <= right; idx++ {
		temp[idx] = arr[idx]
	}
	
	// Merge temp arrays back into arr[left..right]
	for i <= mid && j <= right {
		if temp[i] <= temp[j] {
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = temp[j]
			j++
		}
		k++
	}
	
	// Copy remaining elements
	for i <= mid {
		arr[k] = temp[i]
		i++
		k++
	}
	
	for j <= right {
		arr[k] = temp[j]
		j++
		k++
	}
}

// Implement Heap Sort (in-place, unstable, O(n log n))
func heapSort(arr []int) {
	n := len(arr)
	
	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	
	// Extract elements from heap
	for i := n - 1; i > 0; i-- {
		swap(arr, 0, i)
		heapify(arr, i, 0)
	}
}

// Heapify for heap sort
func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	
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

func generateSingleValueData(size int) []int {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = 42 // All same value
	}
	return data
}

func copySlice(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
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