// generic_data_structures.go - SOLUTION
// Learn to implement data structures using Go generics

package main

import (
	"fmt"
)

// Generic Stack implementation
type Stack[T any] struct {
	items []T
}

// Implement Stack methods
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	
	index := len(s.items) - 1
	value := s.items[index]
	s.items = s.items[:index]
	
	return value, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Generic Queue implementation
type Queue[T any] struct {
	items []T
}

// Implement Queue methods
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: make([]T, 0)}
}

func (q *Queue[T]) Enqueue(value T) {
	q.items = append(q.items, value)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	
	value := q.items[0]
	q.items = q.items[1:]
	
	return value, true
}

func (q *Queue[T]) Front() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	
	return q.items[0], true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}

// Generic Map implementation with custom key constraints
type Map[K comparable, V any] struct {
	data map[K]V
}

// Implement Map methods
func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{data: make(map[K]V)}
}

func (m *Map[K, V]) Put(key K, value V) {
	m.data[key] = value
}

func (m *Map[K, V]) Get(key K) (V, bool) {
	value, exists := m.data[key]
	return value, exists
}

func (m *Map[K, V]) Delete(key K) bool {
	if _, exists := m.data[key]; exists {
		delete(m.data, key)
		return true
	}
	return false
}

func (m *Map[K, V]) Contains(key K) bool {
	_, exists := m.data[key]
	return exists
}

func (m *Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(m.data))
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}

func (m *Map[K, V]) Values() []V {
	values := make([]V, 0, len(m.data))
	for _, v := range m.data {
		values = append(values, v)
	}
	return values
}

func (m *Map[K, V]) Size() int {
	return len(m.data)
}

// Generic Set implementation
type Set[T comparable] struct {
	data map[T]struct{}
}

// Implement Set methods
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]struct{})}
}

func (s *Set[T]) Add(value T) {
	s.data[value] = struct{}{}
}

func (s *Set[T]) Remove(value T) bool {
	if _, exists := s.data[value]; exists {
		delete(s.data, value)
		return true
	}
	return false
}

func (s *Set[T]) Contains(value T) bool {
	_, exists := s.data[value]
	return exists
}

func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.data))
	for v := range s.data {
		slice = append(slice, v)
	}
	return slice
}

func (s *Set[T]) Size() int {
	return len(s.data)
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	
	// Add all elements from current set
	for v := range s.data {
		result.Add(v)
	}
	
	// Add all elements from other set
	for v := range other.data {
		result.Add(v)
	}
	
	return result
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	
	// Add elements that exist in both sets
	for v := range s.data {
		if other.Contains(v) {
			result.Add(v)
		}
	}
	
	return result
}

func main() {
	fmt.Println("=== Generic Data Structures ===")
	
	fmt.Println("\n=== Generic Stack ===")
	
	// Test generic stack with integers
	intStack := NewStack[int]()
	
	fmt.Println("Testing integer stack:")
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	
	for !intStack.IsEmpty() {
		if value, ok := intStack.Pop(); ok {
			fmt.Printf("Popped: %d\n", value)
		}
	}
	
	// Test generic stack with strings
	strStack := NewStack[string]()
	
	fmt.Println("\nTesting string stack:")
	strStack.Push("hello")
	strStack.Push("world")
	strStack.Push("golang")
	
	for !strStack.IsEmpty() {
		if value, ok := strStack.Pop(); ok {
			fmt.Printf("Popped: %s\n", value)
		}
	}
	
	fmt.Println("\n=== Generic Queue ===")
	
	// Test generic queue with floats
	floatQueue := NewQueue[float64]()
	
	fmt.Println("Testing float queue:")
	floatQueue.Enqueue(1.1)
	floatQueue.Enqueue(2.2)
	floatQueue.Enqueue(3.3)
	
	for !floatQueue.IsEmpty() {
		if value, ok := floatQueue.Dequeue(); ok {
			fmt.Printf("Dequeued: %.1f\n", value)
		}
	}
	
	fmt.Println("\n=== Generic Map ===")
	
	// Test generic map with string keys and int values
	studentGrades := NewMap[string, int]()
	
	fmt.Println("Testing student grades map:")
	studentGrades.Put("Alice", 95)
	studentGrades.Put("Bob", 87)
	studentGrades.Put("Charlie", 92)
	
	students := []string{"Alice", "Bob", "Charlie", "Dave"}
	for _, student := range students {
		if grade, found := studentGrades.Get(student); found {
			fmt.Printf("%s: %d\n", student, grade)
		} else {
			fmt.Printf("%s: Not found\n", student)
		}
	}
	
	fmt.Printf("Map size: %d\n", studentGrades.Size())
	fmt.Printf("All keys: %v\n", studentGrades.Keys())
	
	fmt.Println("\n=== Generic Set ===")
	
	// Test generic set with integers
	intSet := NewSet[int]()
	
	fmt.Println("Testing integer set:")
	intSet.Add(1)
	intSet.Add(2)
	intSet.Add(3)
	intSet.Add(2) // Duplicate, should be ignored
	intSet.Add(1) // Duplicate, should be ignored
	
	fmt.Printf("Set contains: %v\n", intSet.ToSlice())
	fmt.Printf("Set size: %d\n", intSet.Size())
	fmt.Printf("Contains 2: %t\n", intSet.Contains(2))
	fmt.Printf("Contains 5: %t\n", intSet.Contains(5))
	
	// Test set operations
	otherSet := NewSet[int]()
	otherSet.Add(3)
	otherSet.Add(4)
	otherSet.Add(5)
	
	union := intSet.Union(otherSet)
	intersection := intSet.Intersection(otherSet)
	
	fmt.Printf("Original set: %v\n", intSet.ToSlice())
	fmt.Printf("Other set: %v\n", otherSet.ToSlice())
	fmt.Printf("Union: %v\n", union.ToSlice())
	fmt.Printf("Intersection: %v\n", intersection.ToSlice())
}