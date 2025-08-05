// linked_list.go - SOLUTION
// Implement a singly linked list data structure

package main

import "fmt"

// Node structure
type Node struct {
	Data int
	Next *Node
}

// LinkedList structure
type LinkedList struct {
	Head *Node
	Size int
}

// Create new linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{Head: nil, Size: 0}
}

// Add element to front
func (ll *LinkedList) AddFront(data int) {
	newNode := &Node{Data: data, Next: ll.Head}
	ll.Head = newNode
	ll.Size++
}

// Add element to back
func (ll *LinkedList) AddBack(data int) {
	newNode := &Node{Data: data, Next: nil}
	
	if ll.Head == nil {
		ll.Head = newNode
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	ll.Size++
}

// Remove from front
func (ll *LinkedList) RemoveFront() (int, bool) {
	if ll.Head == nil {
		return 0, false
	}
	
	data := ll.Head.Data
	ll.Head = ll.Head.Next
	ll.Size--
	return data, true
}

// Remove from back
func (ll *LinkedList) RemoveBack() (int, bool) {
	if ll.Head == nil {
		return 0, false
	}
	
	if ll.Head.Next == nil {
		data := ll.Head.Data
		ll.Head = nil
		ll.Size--
		return data, true
	}
	
	current := ll.Head
	for current.Next.Next != nil {
		current = current.Next
	}
	
	data := current.Next.Data
	current.Next = nil
	ll.Size--
	return data, true
}

// Find element
func (ll *LinkedList) Find(data int) bool {
	current := ll.Head
	for current != nil {
		if current.Data == data {
			return true
		}
		current = current.Next
	}
	return false
}

// Remove specific element
func (ll *LinkedList) Remove(data int) bool {
	if ll.Head == nil {
		return false
	}
	
	if ll.Head.Data == data {
		ll.Head = ll.Head.Next
		ll.Size--
		return true
	}
	
	current := ll.Head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			ll.Size--
			return true
		}
		current = current.Next
	}
	return false
}

// Get element at index
func (ll *LinkedList) Get(index int) (int, bool) {
	if index < 0 || index >= ll.Size {
		return 0, false
	}
	
	current := ll.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Data, true
}

// Insert at index
func (ll *LinkedList) Insert(index int, data int) bool {
	if index < 0 || index > ll.Size {
		return false
	}
	
	if index == 0 {
		ll.AddFront(data)
		return true
	}
	
	newNode := &Node{Data: data}
	current := ll.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	
	newNode.Next = current.Next
	current.Next = newNode
	ll.Size++
	return true
}

// Convert to slice
func (ll *LinkedList) ToSlice() []int {
	result := make([]int, 0, ll.Size)
	current := ll.Head
	for current != nil {
		result = append(result, current.Data)
		current = current.Next
	}
	return result
}

// Print list
func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Printf("[%d]", current.Data)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println(" -> nil")
}

// Reverse the list
func (ll *LinkedList) Reverse() {
	var prev *Node
	current := ll.Head
	
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	
	ll.Head = prev
}

func main() {
	fmt.Println("=== Linked List Implementation ===")
	
	// Create new linked list
	ll := NewLinkedList()
	
	fmt.Println("Adding elements to front and back...")
	
	// Add elements
	ll.AddFront(3)
	ll.AddFront(2)
	ll.AddFront(1)
	ll.AddBack(4)
	ll.AddBack(5)
	ll.AddBack(6)
	
	fmt.Printf("List size: %d\n", ll.Size)
	ll.Print()
	
	fmt.Println("\n=== Finding Elements ===")
	
	// Search for elements
	searchValues := []int{1, 3, 5, 7}
	for _, val := range searchValues {
		found := ll.Find(val)
		fmt.Printf("Find %d: %t\n", val, found)
	}
	
	fmt.Println("\n=== Getting Elements by Index ===")
	
	// Get elements by index
	for i := 0; i < ll.Size+2; i++ {
		if val, ok := ll.Get(i); ok {
			fmt.Printf("Index %d: %d\n", i, val)
		} else {
			fmt.Printf("Index %d: out of bounds\n", i)
		}
	}
	
	fmt.Println("\n=== Removing Elements ===")
	
	// Remove from front and back
	if val, ok := ll.RemoveFront(); ok {
		fmt.Printf("Removed from front: %d\n", val)
	}
	
	if val, ok := ll.RemoveBack(); ok {
		fmt.Printf("Removed from back: %d\n", val)
	}
	
	ll.Print()
	
	// Remove specific element
	fmt.Println("Removing element 3...")
	success := ll.Remove(3)
	fmt.Printf("Remove success: %t\n", success)
	ll.Print()
	
	fmt.Println("\n=== Inserting at Index ===")
	
	// Insert at various positions
	insertions := []struct{ index, value int }{
		{0, 100},  // Insert at beginning
		{2, 200},  // Insert in middle
		{ll.Size, 300}, // Insert at end
	}
	
	for _, ins := range insertions {
		success := ll.Insert(ins.index, ins.value)
		fmt.Printf("Insert %d at index %d: %t\n", ins.value, ins.index, success)
		ll.Print()
	}
	
	fmt.Println("\n=== Reversing List ===")
	
	fmt.Println("Before reverse:")
	ll.Print()
	
	ll.Reverse()
	
	fmt.Println("After reverse:")
	ll.Print()
	
	fmt.Println("\n=== Converting to Slice ===")
	
	slice := ll.ToSlice()
	fmt.Printf("As slice: %v\n", slice)
}