// sorting_algorithms.go - SOLUTION
// Implement various sorting algorithms

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== Sorting Algorithms ===")
	
	// Test data
	original := []int{64, 34, 25, 12, 22, 11, 90, 5}
	fmt.Printf("Original array: %v\n", original)
	
	// Test each sorting algorithm
	algorithms := []struct {
		name string
		fn   func([]int)
	}{
		{"Bubble Sort", bubbleSort},
		{"Selection Sort", selectionSort},
		{"Insertion Sort", insertionSort},
		{"Quick Sort", func(arr []int) { quickSort(arr, 0, len(arr)-1) }},
		{"Merge Sort", func(arr []int) { mergeSort(arr, 0, len(arr)-1) }},
		{"Heap Sort", heapSort},
	}
	
	for _, alg := range algorithms {
		// Make a copy of original array
		arr := make([]int, len(original))
		copy(arr, original)
		
		// Time the sorting
		start := time.Now()
		alg.fn(arr)
		duration := time.Since(start)
		
		fmt.Printf("\n%s: %v (took %v)\n", alg.name, arr, duration)
		
		// Verify it's sorted
		if isSorted(arr) {
			fmt.Printf("✓ %s result is correctly sorted\n", alg.name)
		} else {
			fmt.Printf("❌ %s result is not sorted!\n", alg.name)
		}
	}
	
	// Performance comparison with larger array
	fmt.Println("\n=== Performance Comparison ===")
	
	const size = 1000
	largeArray := make([]int, size)
	for i := 0; i < size; i++ {
		largeArray[i] = rand.Intn(1000)
	}
	
	fmt.Printf("Testing with array of size %d\n", size)
	
	// Test only efficient algorithms on large data
	efficientAlgorithms := []struct {
		name string
		fn   func([]int)
	}{
		{"Quick Sort", func(arr []int) { quickSort(arr, 0, len(arr)-1) }},
		{"Merge Sort", func(arr []int) { mergeSort(arr, 0, len(arr)-1) }},
		{"Heap Sort", heapSort},
	}
	
	for _, alg := range efficientAlgorithms {
		arr := make([]int, len(largeArray))
		copy(arr, largeArray)
		
		start := time.Now()
		alg.fn(arr)
		duration := time.Since(start)
		
		fmt.Printf("%s: %v (sorted: %t)\n", alg.name, duration, isSorted(arr))
	}
}

// Bubble Sort - O(n²) time complexity
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// If no swapping occurred, array is sorted
		if !swapped {
			break
		}
	}
}

// Selection Sort - O(n²) time complexity
func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// Insertion Sort - O(n²) time complexity, but efficient for small arrays
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		
		// Move elements greater than key one position ahead
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// Quick Sort - O(n log n) average, O(n²) worst case
func quickSort(arr []int, low, high int) {
	if low < high {
		// Partition the array and get pivot index
		pi := partition(arr, low, high)
		
		// Recursively sort elements before and after partition
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high] // Choose last element as pivot
	i := low - 1       // Index of smaller element
	
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// Merge Sort - O(n log n) time complexity, stable sort
func mergeSort(arr []int, left, right int) {
	if left < right {
		mid := left + (right-left)/2
		
		// Sort first and second halves
		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)
		
		// Merge the sorted halves
		merge(arr, left, mid, right)
	}
}

func merge(arr []int, left, mid, right int) {
	// Create temporary arrays for left and right subarrays
	leftArr := make([]int, mid-left+1)
	rightArr := make([]int, right-mid)
	
	// Copy data to temporary arrays
	copy(leftArr, arr[left:mid+1])
	copy(rightArr, arr[mid+1:right+1])
	
	// Merge the temporary arrays back into arr[left..right]
	i, j, k := 0, 0, left
	
	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}
	
	// Copy remaining elements
	for i < len(leftArr) {
		arr[k] = leftArr[i]
		i++
		k++
	}
	
	for j < len(rightArr) {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

// Heap Sort - O(n log n) time complexity
func heapSort(arr []int) {
	n := len(arr)
	
	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	
	// Extract elements from heap one by one
	for i := n - 1; i > 0; i-- {
		// Move current root to end
		arr[0], arr[i] = arr[i], arr[0]
		
		// Call heapify on the reduced heap
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n, i int) {
	largest := i // Initialize largest as root
	left := 2*i + 1
	right := 2*i + 2
	
	// If left child is larger than root
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	
	// If right child is larger than largest so far
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	
	// If largest is not root
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		
		// Recursively heapify the affected sub-tree
		heapify(arr, n, largest)
	}
}

// Helper function to check if array is sorted
func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}