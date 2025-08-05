// search_algorithms.go - SOLUTION
// Implement various search algorithms

package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println("=== Search Algorithms ===")
	
	// Test data
	data := []int{2, 5, 8, 12, 16, 23, 38, 45, 67, 78, 89, 99}
	fmt.Printf("Sorted data: %v\n", data)
	
	// Test search algorithms
	searchValues := []int{5, 23, 50, 99, 1, 100}
	
	for _, target := range searchValues {
		fmt.Printf("\nSearching for %d:\n", target)
		
		// Linear search
		linearResult := linearSearch(data, target)
		fmt.Printf("  Linear search: index %d\n", linearResult)
		
		// Binary search
		binaryResult := binarySearch(data, target)
		fmt.Printf("  Binary search: index %d\n", binaryResult)
		
		// Interpolation search
		interpResult := interpolationSearch(data, target)
		fmt.Printf("  Interpolation search: index %d\n", interpResult)
	}
	
	// Performance comparison
	fmt.Println("\n=== Performance Comparison ===")
	
	// Generate large sorted array
	const size = 100000
	largeData := make([]int, size)
	for i := 0; i < size; i++ {
		largeData[i] = i * 2 // Even numbers
	}
	
	target := 50000
	
	// Time linear search
	start := time.Now()
	linearSearch(largeData, target)
	linearTime := time.Since(start)
	
	// Time binary search
	start = time.Now()
	binarySearch(largeData, target)
	binaryTime := time.Since(start)
	
	fmt.Printf("Linear search time: %v\n", linearTime)
	fmt.Printf("Binary search time: %v\n", binaryTime)
	if binaryTime > 0 {
		fmt.Printf("Binary search is %.2fx faster\n", float64(linearTime)/float64(binaryTime))
	}
}

// Linear search implementation
func linearSearch(arr []int, target int) int {
	// Search through array sequentially
	for i, value := range arr {
		if value == target {
			return i
		}
	}
	// Return -1 if not found
	return -1
}

// Binary search implementation
func binarySearch(arr []int, target int) int {
	// Use divide and conquer approach
	left, right := 0, len(arr)-1
	
	for left <= right {
		// Compare with middle element
		mid := left + (right-left)/2
		
		if arr[mid] == target {
			return mid
		}
		
		// Search left or right half based on comparison
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return -1
}

// Interpolation search implementation
func interpolationSearch(arr []int, target int) int {
	// Improvement over binary search for uniformly distributed data
	left, right := 0, len(arr)-1
	
	for left <= right && target >= arr[left] && target <= arr[right] {
		// If array has only one element
		if left == right {
			if arr[left] == target {
				return left
			}
			return -1
		}
		
		// Estimate position based on value
		// Works best with evenly distributed data
		pos := left + int(float64(right-left)*float64(target-arr[left])/float64(arr[right]-arr[left]))
		
		// Ensure pos is within bounds
		if pos < left {
			pos = left
		} else if pos > right {
			pos = right
		}
		
		if arr[pos] == target {
			return pos
		}
		
		if arr[pos] < target {
			left = pos + 1
		} else {
			right = pos - 1
		}
	}
	
	return -1
}

// Jump search implementation
func jumpSearch(arr []int, target int) int {
	// Jump through array in blocks
	n := len(arr)
	step := int(math.Sqrt(float64(n)))
	prev := 0
	
	// Find the block where element is present
	for arr[int(math.Min(float64(step), float64(n)))-1] < target {
		prev = step
		step += int(math.Sqrt(float64(n)))
		if prev >= n {
			return -1
		}
	}
	
	// Linear search within block
	for arr[prev] < target {
		prev++
		if prev == int(math.Min(float64(step), float64(n))) {
			return -1
		}
	}
	
	if arr[prev] == target {
		return prev
	}
	
	return -1
}

// Exponential search implementation
func exponentialSearch(arr []int, target int) int {
	// Find range where element might be present
	if arr[0] == target {
		return 0
	}
	
	n := len(arr)
	i := 1
	for i < n && arr[i] <= target {
		i = i * 2
	}
	
	// Use binary search in that range
	left := i / 2
	right := int(math.Min(float64(i), float64(n-1)))
	
	for left <= right {
		mid := left + (right-left)/2
		
		if arr[mid] == target {
			return mid
		}
		
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return -1
}