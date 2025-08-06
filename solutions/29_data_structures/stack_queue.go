// stack_queue.go - SOLUTION
// Learn Stack and Queue data structure implementations

package main

import (
	"errors"
	"fmt"
	"strings"
)

// Stack implementation using slice
type Stack struct {
	items []int
}

// Implement Stack methods
func NewStack() *Stack {
	return &Stack{items: make([]int, 0)}
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	
	index := len(s.items) - 1
	value := s.items[index]
	s.items = s.items[:index]
	
	return value, nil
}

func (s *Stack) Peek() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	
	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) String() string {
	if len(s.items) == 0 {
		return "[]"
	}
	
	strs := make([]string, len(s.items))
	for i, v := range s.items {
		strs[i] = fmt.Sprintf("%d", v)
	}
	
	return "[" + strings.Join(strs, " ") + "] ← top"
}

// Queue implementation using slice
type Queue struct {
	items []int
}

// Implement Queue methods
func NewQueue() *Queue {
	return &Queue{items: make([]int, 0)}
}

func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, errors.New("queue is empty")
	}
	
	value := q.items[0]
	q.items = q.items[1:]
	
	return value, nil
}

func (q *Queue) Front() (int, error) {
	if len(q.items) == 0 {
		return 0, errors.New("queue is empty")
	}
	
	return q.items[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) String() string {
	if len(q.items) == 0 {
		return "[]"
	}
	
	strs := make([]string, len(q.items))
	for i, v := range q.items {
		strs[i] = fmt.Sprintf("%d", v)
	}
	
	return "front → [" + strings.Join(strs, " ") + "] ← rear"
}

// Circular Queue implementation with fixed size
type CircularQueue struct {
	items    []int
	front    int
	rear     int
	size     int
	capacity int
}

// Implement Circular Queue methods
func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		items:    make([]int, capacity),
		front:    0,
		rear:     -1,
		size:     0,
		capacity: capacity,
	}
}

func (cq *CircularQueue) Enqueue(value int) error {
	if cq.IsFull() {
		return errors.New("queue is full")
	}
	
	cq.rear = (cq.rear + 1) % cq.capacity
	cq.items[cq.rear] = value
	cq.size++
	
	return nil
}

func (cq *CircularQueue) Dequeue() (int, error) {
	if cq.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	
	value := cq.items[cq.front]
	cq.front = (cq.front + 1) % cq.capacity
	cq.size--
	
	return value, nil
}

func (cq *CircularQueue) IsFull() bool {
	return cq.size == cq.capacity
}

func (cq *CircularQueue) IsEmpty() bool {
	return cq.size == 0
}

func (cq *CircularQueue) Size() int {
	return cq.size
}

func (cq *CircularQueue) String() string {
	if cq.IsEmpty() {
		return "[]"
	}
	
	strs := make([]string, cq.size)
	for i := 0; i < cq.size; i++ {
		index := (cq.front + i) % cq.capacity
		strs[i] = fmt.Sprintf("%d", cq.items[index])
	}
	
	return "[" + strings.Join(strs, " ") + "]"
}

// Deque (Double-ended queue) implementation
type Deque struct {
	items []int
}

// Implement Deque methods
func NewDeque() *Deque {
	return &Deque{items: make([]int, 0)}
}

func (d *Deque) PushFront(value int) {
	d.items = append([]int{value}, d.items...)
}

func (d *Deque) PushBack(value int) {
	d.items = append(d.items, value)
}

func (d *Deque) PopFront() (int, error) {
	if len(d.items) == 0 {
		return 0, errors.New("deque is empty")
	}
	
	value := d.items[0]
	d.items = d.items[1:]
	
	return value, nil
}

func (d *Deque) PopBack() (int, error) {
	if len(d.items) == 0 {
		return 0, errors.New("deque is empty")
	}
	
	index := len(d.items) - 1
	value := d.items[index]
	d.items = d.items[:index]
	
	return value, nil
}

func (d *Deque) Front() (int, error) {
	if len(d.items) == 0 {
		return 0, errors.New("deque is empty")
	}
	
	return d.items[0], nil
}

func (d *Deque) Back() (int, error) {
	if len(d.items) == 0 {
		return 0, errors.New("deque is empty")
	}
	
	return d.items[len(d.items)-1], nil
}

func (d *Deque) IsEmpty() bool {
	return len(d.items) == 0
}

func (d *Deque) Size() int {
	return len(d.items)
}

func (d *Deque) String() string {
	if len(d.items) == 0 {
		return "[]"
	}
	
	strs := make([]string, len(d.items))
	for i, v := range d.items {
		strs[i] = fmt.Sprintf("%d", v)
	}
	
	return "[" + strings.Join(strs, " ") + "]"
}

func main() {
	fmt.Println("=== Data Structures: Stack and Queue ===")
	
	fmt.Println("\n=== Stack Operations ===")
	
	// Test stack operations
	stack := NewStack()
	
	fmt.Printf("Empty stack: %s (size: %d, empty: %t)\n", 
		stack.String(), stack.Size(), stack.IsEmpty())
	
	// Push elements
	elements := []int{10, 20, 30, 40, 50}
	fmt.Println("\nPushing elements:")
	for _, elem := range elements {
		stack.Push(elem)
		fmt.Printf("Pushed %d: %s\n", elem, stack.String())
	}
	
	// Peek at top element
	if top, err := stack.Peek(); err == nil {
		fmt.Printf("Top element: %d\n", top)
	}
	
	// Pop elements
	fmt.Println("\nPopping elements:")
	for !stack.IsEmpty() {
		if value, err := stack.Pop(); err == nil {
			fmt.Printf("Popped %d: %s\n", value, stack.String())
		}
	}
	
	// Try to pop from empty stack
	if _, err := stack.Pop(); err != nil {
		fmt.Printf("Pop from empty stack: %v\n", err)
	}
	
	fmt.Println("\n=== Queue Operations ===")
	
	// Test queue operations
	queue := NewQueue()
	
	fmt.Printf("Empty queue: %s (size: %d, empty: %t)\n", 
		queue.String(), queue.Size(), queue.IsEmpty())
	
	// Enqueue elements
	fmt.Println("\nEnqueuing elements:")
	for _, elem := range elements {
		queue.Enqueue(elem)
		fmt.Printf("Enqueued %d: %s\n", elem, queue.String())
	}
	
	// Check front element
	if front, err := queue.Front(); err == nil {
		fmt.Printf("Front element: %d\n", front)
	}
	
	// Dequeue elements
	fmt.Println("\nDequeuing elements:")
	for !queue.IsEmpty() {
		if value, err := queue.Dequeue(); err == nil {
			fmt.Printf("Dequeued %d: %s\n", value, queue.String())
		}
	}
	
	fmt.Println("\n=== Circular Queue Operations ===")
	
	// Test circular queue with capacity 5
	circularQueue := NewCircularQueue(5)
	
	fmt.Printf("Empty circular queue (capacity 5): %s\n", circularQueue.String())
	
	// Fill the circular queue
	fmt.Println("\nFilling circular queue:")
	for i := 1; i <= 5; i++ {
		if err := circularQueue.Enqueue(i); err == nil {
			fmt.Printf("Enqueued %d: %s (full: %t)\n", i, circularQueue.String(), circularQueue.IsFull())
		}
	}
	
	// Try to add to full queue
	if err := circularQueue.Enqueue(6); err != nil {
		fmt.Printf("Enqueue to full queue: %v\n", err)
	}
	
	// Dequeue some elements and add new ones
	fmt.Println("\nDequeue 2 elements and add 2 new ones:")
	circularQueue.Dequeue()
	circularQueue.Dequeue()
	
	circularQueue.Enqueue(6)
	circularQueue.Enqueue(7)
	
	fmt.Printf("After operations: %s\n", circularQueue.String())
	
	fmt.Println("\n=== Deque Operations ===")
	
	// Test deque operations
	deque := NewDeque()
	
	fmt.Printf("Empty deque: %s\n", deque.String())
	
	// Add elements to both ends
	fmt.Println("\nAdding elements to both ends:")
	deque.PushBack(10)
	fmt.Printf("PushBack(10): %s\n", deque.String())
	
	deque.PushFront(5)
	fmt.Printf("PushFront(5): %s\n", deque.String())
	
	deque.PushBack(20)
	fmt.Printf("PushBack(20): %s\n", deque.String())
	
	deque.PushFront(1)
	fmt.Printf("PushFront(1): %s\n", deque.String())
	
	// Check front and back elements
	if front, err := deque.Front(); err == nil {
		fmt.Printf("Front: %d\n", front)
	}
	if back, err := deque.Back(); err == nil {
		fmt.Printf("Back: %d\n", back)
	}
	
	// Remove elements from both ends
	fmt.Println("\nRemoving elements from both ends:")
	if value, err := deque.PopFront(); err == nil {
		fmt.Printf("PopFront(): %d, remaining: %s\n", value, deque.String())
	}
	
	if value, err := deque.PopBack(); err == nil {
		fmt.Printf("PopBack(): %d, remaining: %s\n", value, deque.String())
	}
	
	fmt.Println("\n=== Stack Applications: Balanced Parentheses ===")
	
	// Use stack to check balanced parentheses
	testStrings := []string{
		"()",
		"(())",
		"(()",
		"())",
		"((()))",
		"(()())",
		"(()()",
	}
	
	fmt.Println("Checking balanced parentheses:")
	for _, s := range testStrings {
		balanced := checkBalancedParentheses(s)
		status := "✓"
		if !balanced {
			status = "✗"
		}
		fmt.Printf("  %s %s: %t\n", status, s, balanced)
	}
	
	fmt.Println("\n=== Queue Applications: Level Order Traversal Simulation ===")
	
	// Simulate processing tasks in order
	tasks := []string{"Task1", "Task2", "Task3", "Task4", "Task5"}
	
	fmt.Println("Processing tasks in FIFO order:")
	simulateTaskProcessing(tasks)
}

// Implement utility functions

func checkBalancedParentheses(s string) bool {
	stack := NewStack()
	
	for _, char := range s {
		if char == '(' {
			stack.Push(1) // Use 1 to represent '('
		} else if char == ')' {
			if _, err := stack.Pop(); err != nil {
				return false // No matching '(' for this ')'
			}
		}
	}
	
	return stack.IsEmpty() // Balanced if no unmatched '(' remain
}

func simulateTaskProcessing(tasks []string) {
	queue := []string{} // Simple string queue for demonstration
	
	// Add all tasks to queue
	for _, task := range tasks {
		queue = append(queue, task)
		fmt.Printf("Added %s to queue\n", task)
	}
	
	// Process tasks in FIFO order
	for len(queue) > 0 {
		task := queue[0]
		queue = queue[1:]
		fmt.Printf("Processing %s...\n", task)
	}
	
	fmt.Println("All tasks processed!")
}