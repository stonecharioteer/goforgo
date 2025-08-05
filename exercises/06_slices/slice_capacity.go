// slice_capacity.go
// Understand slice capacity, memory allocation, and growth behavior
// Learn how Go manages slice memory behind the scenes

package main

import "fmt"

func main() {
	// TODO: Create a slice with make() specifying length and capacity
	// make([]type, length, capacity)
	slice1 := // Create a slice with length 5 and capacity 10

	fmt.Printf("slice1 - len: %d, cap: %d, slice: %v\n", len(slice1), cap(slice1), slice1)

	// TODO: Fill the slice with some values
	for i := range slice1 {
		slice1[i] = i * 2
	}
	fmt.Printf("After filling: %v\n", slice1)

	// TODO: Create a slice from an array and observe capacity
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := // Create slice from array[2:6]
	
	fmt.Printf("slice2 - len: %d, cap: %d, slice: %v\n", len(slice2), cap(slice2), slice2)

	// TODO: Modify slice2 and see how it affects the original array
	slice2[0] = 999
	fmt.Printf("Array after modifying slice2: %v\n", array)
	fmt.Printf("slice2 after modification: %v\n", slice2)

	// TODO: Demonstrate slice growth by appending beyond capacity
	smallSlice := make([]int, 2, 3) // length 2, capacity 3
	fmt.Printf("Initial smallSlice - len: %d, cap: %d\n", len(smallSlice), cap(smallSlice))

	// Fill initial elements
	smallSlice[0] = 10
	smallSlice[1] = 20

	// TODO: Append elements and watch capacity changes
	fmt.Println("Appending elements and observing capacity:")
	
	// Append one element (should still fit in capacity)
	smallSlice = append(smallSlice, 30)
	fmt.Printf("After append(30) - len: %d, cap: %d, slice: %v\n", len(smallSlice), cap(smallSlice), smallSlice)

	// TODO: Append another element (should trigger capacity increase)
	smallSlice = append(smallSlice, 40)
	fmt.Printf("After append(40) - len: %d, cap: %d, slice: %v\n", len(smallSlice), cap(smallSlice), smallSlice)

	// TODO: Append multiple elements at once
	smallSlice = append(smallSlice, 50, 60, 70)
	fmt.Printf("After append(50,60,70) - len: %d, cap: %d, slice: %v\n", len(smallSlice), cap(smallSlice), smallSlice)

	// TODO: Create two slices from the same underlying array
	base := []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice3 := // Create slice from base[1:4]
	slice4 := // Create slice from base[3:6]
	
	fmt.Printf("base: %v\n", base)
	fmt.Printf("slice3 [1:4]: %v (len: %d, cap: %d)\n", slice3, len(slice3), cap(slice3))
	fmt.Printf("slice4 [3:6]: %v (len: %d, cap: %d)\n", slice4, len(slice4), cap(slice4))

	// TODO: Modify overlapping element and observe both slices
	slice3[2] = 999 // This should affect both slices since they share base[3]
	fmt.Printf("After modifying slice3[2] to 999:\n")
	fmt.Printf("base: %v\n", base)
	fmt.Printf("slice3: %v\n", slice3)
	fmt.Printf("slice4: %v\n", slice4)

	// TODO: Demonstrate full slice expression [low:high:max]
	fullSlice := // Create slice with base[1:4:5] to limit capacity
	fmt.Printf("Full slice expression [1:4:5] - len: %d, cap: %d, slice: %v\n", 
		len(fullSlice), cap(fullSlice), fullSlice)

	// TODO: Show memory allocation with a large append operation
	fmt.Println("\nDemonstrating memory reallocation:")
	growSlice := []int{1}
	for i := 0; i < 5; i++ {
		fmt.Printf("Before append - len: %d, cap: %d\n", len(growSlice), cap(growSlice))
		growSlice = append(growSlice, i)
		fmt.Printf("After append  - len: %d, cap: %d\n", len(growSlice), cap(growSlice))
		fmt.Println()
	}
}