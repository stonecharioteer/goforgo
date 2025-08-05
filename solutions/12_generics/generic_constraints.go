// generic_constraints.go - SOLUTION
// Learn advanced generic constraints and custom type constraints

package main

import (
	"fmt"
	"sort"
)

// Define custom constraints using interfaces

// Numeric constraint for number types
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Ordered constraint for types that can be compared
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 | ~string
}

// Stringer constraint for types that can be converted to string
type Stringer interface {
	String() string
}

// Combined constraint example
type StringLike interface {
	~string | fmt.Stringer
}

// Generic functions with custom constraints

// Sum works with any numeric type
func Sum[T Numeric](values []T) T {
	var sum T
	// Calculate sum of all values
	for _, v := range values {
		sum += v
	}
	return sum
}

// Min finds minimum value for ordered types
func Min[T Ordered](a, b T) T {
	// Return smaller value
	if a < b {
		return a
	}
	return b
}

// Sort sorts a slice of ordered values (simple bubble sort)
func Sort[T Ordered](slice []T) {
	// Implement bubble sort
	// Modify slice in place
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

// FormatSlice converts slice elements to strings
func FormatSlice[T Stringer](slice []T) []string {
	// Convert each element to string using String() method
	result := make([]string, len(slice))
	for i, v := range slice {
		result[i] = v.String()
	}
	return result
}

// Generic data structures with constraints

// NumericStack is a stack that only accepts numeric types
type NumericStack[T Numeric] struct {
	items []T
}

func (s *NumericStack[T]) Push(item T) {
	// Add item to stack
	s.items = append(s.items, item)
}

func (s *NumericStack[T]) Pop() (T, bool) {
	// Remove and return top item
	// Return zero value and false if empty
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

func (s *NumericStack[T]) Sum() T {
	// Return sum of all items in stack
	var sum T
	for _, item := range s.items {
		sum += item
	}
	return sum
}

// Ordered map that maintains sorted keys
type OrderedMap[K Ordered, V any] struct {
	keys   []K
	values map[K]V
}

func NewOrderedMap[K Ordered, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		keys:   make([]K, 0),
		values: make(map[K]V),
	}
}

func (om *OrderedMap[K, V]) Set(key K, value V) {
	// If key doesn't exist, add it in sorted order
	if _, exists := om.values[key]; !exists {
		// Find insertion point to maintain sorted order
		insertIndex := sort.Search(len(om.keys), func(i int) bool {
			return om.keys[i] >= key
		})
		
		// Insert key at correct position
		om.keys = append(om.keys, key)
		copy(om.keys[insertIndex+1:], om.keys[insertIndex:])
		om.keys[insertIndex] = key
	}
	om.values[key] = value
}

func (om *OrderedMap[K, V]) Get(key K) (V, bool) {
	// Get value by key
	value, exists := om.values[key]
	return value, exists
}

func (om *OrderedMap[K, V]) Keys() []K {
	// Return copy of sorted keys
	result := make([]K, len(om.keys))
	copy(result, om.keys)
	return result
}

func (om *OrderedMap[K, V]) Values() []V {
	// Return values in key order
	result := make([]V, len(om.keys))
	for i, key := range om.keys {
		result[i] = om.values[key]
	}
	return result
}

// Type constraint with methods
type Resettable interface {
	Reset()
}

type Counter[T Numeric] struct {
	count T
	step  T
}

func (c *Counter[T]) Increment() {
	c.count += c.step
}

func (c *Counter[T]) Value() T {
	return c.count
}

func (c *Counter[T]) Reset() {
	var zero T
	c.count = zero
}

// Generic function that works with resettable types
func ResetAll[T Resettable](items []T) {
	// Reset all items in the slice
	for _, item := range items {
		item.Reset()
	}
}

// Complex constraint combinations
type ReadWriter[T any] interface {
	Read() (T, error)
	Write(T) error
}

type Buffer[T any] struct {
	data []T
	pos  int
}

func (b *Buffer[T]) Read() (T, error) {
	// Read next item from buffer
	// Return zero value and error if at end
	if b.pos >= len(b.data) {
		var zero T
		return zero, fmt.Errorf("end of buffer")
	}
	
	item := b.data[b.pos]
	b.pos++
	return item, nil
}

func (b *Buffer[T]) Write(item T) error {
	// Write item to buffer
	b.data = append(b.data, item)
	return nil
}

func (b *Buffer[T]) Reset() {
	b.data = b.data[:0]
	b.pos = 0
}

// Generic function using complex constraint
func ProcessReadWriter[T any, RW ReadWriter[T]](rw RW) []T {
	var results []T
	// Read all available items
	for {
		item, err := rw.Read()
		if err != nil {
			break
		}
		results = append(results, item)
	}
	return results
}

// Custom types that implement constraints
type Temperature float64

func (t Temperature) String() string {
	return fmt.Sprintf("%.1f°C", float64(t))
}

type Distance int

func (d Distance) String() string {
	return fmt.Sprintf("%dm", int(d))
}

func main() {
	fmt.Println("=== Numeric Constraints ===")
	
	// Test Sum with different numeric types
	intSlice := []int{1, 2, 3, 4, 5}
	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	
	fmt.Printf("Sum of ints: %d\n", Sum(intSlice))
	fmt.Printf("Sum of floats: %.2f\n", Sum(floatSlice))
	
	// Test with custom numeric type
	temperatures := []Temperature{20.5, 25.0, 18.3, 22.7}
	fmt.Printf("Sum of temperatures: %s\n", Sum(temperatures).String())
	
	fmt.Println("\n=== Ordered Constraints ===")
	
	// Test Min function
	fmt.Printf("Min(10, 5): %d\n", Min(10, 5))
	fmt.Printf("Min(3.14, 2.71): %.2f\n", Min(3.14, 2.71))
	fmt.Printf("Min(\"banana\", \"apple\"): %s\n", Min("banana", "apple"))
	
	// Test Sort function
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Before sort: %v\n", numbers)
	Sort(numbers)
	fmt.Printf("After sort: %v\n", numbers)
	
	words := []string{"zebra", "apple", "banana", "cherry"}
	fmt.Printf("Before sort: %v\n", words)
	Sort(words)
	fmt.Printf("After sort: %v\n", words)
	
	fmt.Println("\n=== Stringer Constraints ===")
	
	// Test FormatSlice with types that implement String()
	temps := []Temperature{20.5, 25.0, 18.3}
	distances := []Distance{100, 250, 75}
	
	tempStrings := FormatSlice(temps)
	distanceStrings := FormatSlice(distances)
	
	fmt.Printf("Temperature strings: %v\n", tempStrings)
	fmt.Printf("Distance strings: %v\n", distanceStrings)
	
	fmt.Println("\n=== Generic Data Structures ===")
	
	// Test NumericStack
	stack := &NumericStack[int]{}
	
	// Push some values
	for _, val := range []int{10, 20, 30, 40} {
		stack.Push(val)
		fmt.Printf("Pushed %d, stack sum: %d\n", val, stack.Sum())
	}
	
	// Pop values
	for i := 0; i < 2; i++ {
		if val, ok := stack.Pop(); ok {
			fmt.Printf("Popped %d, stack sum: %d\n", val, stack.Sum())
		}
	}
	
	fmt.Println("\n=== Ordered Map ===")
	
	// Test OrderedMap
	om := NewOrderedMap[string, int]()
	
	// Add items in random order
	items := map[string]int{
		"zebra":  26,
		"apple":  1,
		"banana": 2,
		"cherry": 3,
		"date":   4,
	}
	
	for key, value := range items {
		om.Set(key, value)
		fmt.Printf("Added %s: %d\n", key, value)
	}
	
	fmt.Printf("Ordered keys: %v\n", om.Keys())
	fmt.Printf("Ordered values: %v\n", om.Values())
	
	// Test retrieval
	if value, ok := om.Get("banana"); ok {
		fmt.Printf("Retrieved banana: %d\n", value)
	}
	
	fmt.Println("\n=== Resettable Constraint ===")
	
	// Test Counter with Resettable interface
	counter1 := &Counter[int]{step: 1}
	counter2 := &Counter[float64]{step: 0.5}
	
	// Increment counters
	for i := 0; i < 3; i++ {
		counter1.Increment()
		counter2.Increment()
	}
	
	fmt.Printf("Counter1 value: %d\n", counter1.Value())
	fmt.Printf("Counter2 value: %.1f\n", counter2.Value())
	
	// Reset using generic function
	counters := []Resettable{counter1, counter2}
	ResetAll(counters)
	
	fmt.Printf("After reset - Counter1: %d, Counter2: %.1f\n", 
		counter1.Value(), counter2.Value())
	
	fmt.Println("\n=== ReadWriter Interface ===")
	
	// Test Buffer with ReadWriter constraint
	buffer := &Buffer[string]{}
	
	// Write some data
	testData := []string{"hello", "world", "go", "generics"}
	for _, data := range testData {
		buffer.Write(data)
	}
	
	fmt.Printf("Buffer data: %v\n", buffer.data)
	
	// Process using ReadWriter constraint
	results := ProcessReadWriter(buffer)
	fmt.Printf("Processed results: %v\n", results)
	
	fmt.Println("\n=== Type Inference ===")
	
	// Demonstrate type inference - Go can infer types
	fmt.Printf("Inferred Min(42, 17): %d\n", Min(42, 17))
	fmt.Printf("Inferred Sum: %d\n", Sum([]int{1, 2, 3}))
	
	// Explicit type specification when needed
	var emptyInts []int
	fmt.Printf("Sum of empty slice: %d\n", Sum[int](emptyInts))
	
	fmt.Println("Generic constraints examples completed")
}