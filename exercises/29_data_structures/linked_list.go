// linked_list.go
// Implement a singly linked list data structure

package main

import "fmt"

// TODO: Define Node structure
type Node struct {
	Data int
	Next *Node
}

// TODO: Define LinkedList structure
type LinkedList struct {
	Head *Node
	Size int
}

// TODO: Create new linked list
func NewLinkedList() *LinkedList {
	// Return new empty linked list
}

// TODO: Add element to front
func (ll *LinkedList) AddFront(data int) {
	// Create new node
	// Set as new head
	// Update size
}

// TODO: Add element to back
func (ll *LinkedList) AddBack(data int) {
	// Create new node
	// If empty, set as head
	// Otherwise traverse to end and add
	// Update size
}

// TODO: Remove from front
func (ll *LinkedList) RemoveFront() (int, bool) {
	// Check if empty
	// Get head data
	// Update head to next
	// Update size
	// Return data and success
}

// TODO: Remove from back
func (ll *LinkedList) RemoveBack() (int, bool) {
	// Check if empty
	// Handle single element case
	// Traverse to second-to-last node
	// Remove last node
	// Update size
	// Return data and success
}

// TODO: Find element
func (ll *LinkedList) Find(data int) bool {
	// Traverse list looking for data
	// Return true if found
}

// TODO: Remove specific element
func (ll *LinkedList) Remove(data int) bool {
	// Handle empty list
	// Handle head removal
	// Traverse to find element
	// Remove element if found
	// Update size
	// Return success
}

// TODO: Get element at index
func (ll *LinkedList) Get(index int) (int, bool) {
	// Check bounds
	// Traverse to index
	// Return data
}

// TODO: Insert at index
func (ll *LinkedList) Insert(index int, data int) bool {
	// Check bounds
	// Handle insertion at head
	// Traverse to position
	// Insert new node
	// Update size
}

// TODO: Convert to slice
func (ll *LinkedList) ToSlice() []int {
	// Create slice
	// Traverse and add elements
	// Return slice
}

// TODO: Print list
func (ll *LinkedList) Print() {
	// Traverse and print elements
	// Format: [1] -> [2] -> [3] -> nil
}

// TODO: Reverse the list
func (ll *LinkedList) Reverse() {
	// Use three pointers: prev, current, next
	// Reverse links while traversing
}

func main() {
	fmt.Println("=== Linked List Implementation ===")
	
	// TODO: Create new linked list
	ll := /* create new linked list */
	
	fmt.Println("Adding elements to front and back...")
	
	// TODO: Add elements
	// Add 1, 2, 3 to front
	// Add 4, 5, 6 to back
	
	fmt.Printf("List size: %d\n", ll.Size)
	ll.Print()
	
	fmt.Println("\n=== Finding Elements ===")
	
	// TODO: Search for elements
	searchValues := []int{1, 3, 5, 7}
	for _, val := range searchValues {
		found := /* search for val */
		fmt.Printf("Find %d: %t\n", val, found)
	}
	
	fmt.Println("\n=== Getting Elements by Index ===")
	
	// TODO: Get elements by index
	for i := 0; i < ll.Size+2; i++ {
		if val, ok := /* get element at index i */; ok {
			fmt.Printf("Index %d: %d\n", i, val)
		} else {
			fmt.Printf("Index %d: out of bounds\n", i)
		}
	}
	
	fmt.Println("\n=== Removing Elements ===")
	
	// TODO: Remove from front and back
	if val, ok := /* remove from front */; ok {
		fmt.Printf("Removed from front: %d\n", val)
	}
	
	if val, ok := /* remove from back */; ok {
		fmt.Printf("Removed from back: %d\n", val)
	}
	
	ll.Print()
	
	// TODO: Remove specific element
	fmt.Println("Removing element 3...")
	success := /* remove element 3 */
	fmt.Printf("Remove success: %t\n", success)
	ll.Print()
	
	fmt.Println("\n=== Inserting at Index ===")
	
	// TODO: Insert at various positions
	insertions := []struct{ index, value int }{
		{0, 100},  // Insert at beginning
		{2, 200},  // Insert in middle
		{ll.Size, 300}, // Insert at end
	}
	
	for _, ins := range insertions {
		success := /* insert ins.value at ins.index */
		fmt.Printf("Insert %d at index %d: %t\n", ins.value, ins.index, success)
		ll.Print()
	}
	
	fmt.Println("\n=== Reversing List ===")
	
	fmt.Println("Before reverse:")
	ll.Print()
	
	/* reverse the list */
	
	fmt.Println("After reverse:")
	ll.Print()
	
	fmt.Println("\n=== Converting to Slice ===")
	
	slice := /* convert to slice */
	fmt.Printf("As slice: %v\n", slice)
}