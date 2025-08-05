// generic_constraints.go
// Learn advanced generic constraints and custom type constraints

package main

import (
	"fmt"
	"strings"
)

// TODO: Define custom constraints using interfaces

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

// TODO: Generic functions with custom constraints

// Sum works with any numeric type
func Sum[T Numeric](values []T) T {
	var sum T
	// Calculate sum of all values
	return sum
}

// Min finds minimum value for ordered types
func Min[T Ordered](a, b T) T {
	// Return smaller value
}

// Sort sorts a slice of ordered values (simple bubble sort)
func Sort[T Ordered](slice []T) {
	// Implement bubble sort
	// Modify slice in place
}

// FormatSlice converts slice elements to strings
func FormatSlice[T Stringer](slice []T) []string {
	// Convert each element to string using String() method
}

// TODO: Generic data structures with constraints

// NumericStack is a stack that only accepts numeric types
type NumericStack[T Numeric] struct {
	items []T
}

func (s *NumericStack[T]) Push(item T) {
	// Add item to stack
}

func (s *NumericStack[T]) Pop() (T, bool) {
	// Remove and return top item
	// Return zero value and false if empty
}

func (s *NumericStack[T]) Sum() T {
	// Return sum of all items in stack
}

// TODO: Ordered map that maintains sorted keys
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
		// Insert key at correct position
	}
	om.values[key] = value
}

func (om *OrderedMap[K, V]) Get(key K) (V, bool) {
	// Get value by key
}

func (om *OrderedMap[K, V]) Keys() []K {
	// Return copy of sorted keys
}

func (om *OrderedMap[K, V]) Values() []V {
	// Return values in key order
}

// TODO: Type constraint with methods
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
}

// TODO: Complex constraint combinations
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
}

func (b *Buffer[T]) Write(item T) error {
	// Write item to buffer
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
	return results
}

// TODO: Custom types that implement constraints
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
	
	// TODO: Test Sum with different numeric types
	intSlice := []int{1, 2, 3, 4, 5}
	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	
	fmt.Printf("Sum of ints: %d\n", /* call Sum with intSlice */)
	fmt.Printf("Sum of floats: %.2f\n", /* call Sum with floatSlice */)
	
	// TODO: Test with custom numeric type
	temperatures := []Temperature{20.5, 25.0, 18.3, 22.7}
	fmt.Printf("Sum of temperatures: %s\n", /* call Sum and convert result to string */)
	
	fmt.Println("\n=== Ordered Constraints ===")
	
	// TODO: Test Min function
	fmt.Printf("Min(10, 5): %d\n", /* call Min with 10, 5 */)
	fmt.Printf("Min(3.14, 2.71): %.2f\n", /* call Min with 3.14, 2.71 */)
	fmt.Printf("Min(\"banana\", \"apple\"): %s\n", /* call Min with strings */)
	
	// TODO: Test Sort function
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Before sort: %v\n", numbers)
	/* call Sort with numbers */
	fmt.Printf("After sort: %v\n", numbers)
	
	words := []string{"zebra", "apple", "banana", "cherry"}
	fmt.Printf("Before sort: %v\n", words)
	/* call Sort with words */
	fmt.Printf("After sort: %v\n", words)
	
	fmt.Println("\n=== Stringer Constraints ===")
	
	// TODO: Test FormatSlice with types that implement String()
	temps := []Temperature{20.5, 25.0, 18.3}
	distances := []Distance{100, 250, 75}
	
	tempStrings := /* call FormatSlice with temps */
	distanceStrings := /* call FormatSlice with distances */
	
	fmt.Printf("Temperature strings: %v\n", tempStrings)
	fmt.Printf("Distance strings: %v\n", distanceStrings)
	
	fmt.Println("\n=== Generic Data Structures ===")
	
	// TODO: Test NumericStack
	stack := /* create NumericStack of int */
	
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
	
	// TODO: Test OrderedMap
	om := /* create OrderedMap with string keys and int values */
	
	// Add items in random order
	items := map[string]int{
		"zebra": 26,
		"apple": 1,
		"banana": 2,
		"cherry": 3,
		"date": 4,
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
	
	// TODO: Test Counter with Resettable interface
	counter1 := /* create Counter[int] with step=1 */
	counter2 := /* create Counter[float64] with step=0.5 */
	
	// Increment counters
	for i := 0; i < 3; i++ {
		counter1.Increment()
		counter2.Increment()
	}
	
	fmt.Printf("Counter1 value: %d\n", counter1.Value())
	fmt.Printf("Counter2 value: %.1f\n", counter2.Value())
	
	// TODO: Reset using generic function
	counters := []Resettable{&counter1, &counter2}
	/* call ResetAll with counters */
	
	fmt.Printf("After reset - Counter1: %d, Counter2: %.1f\n", 
		counter1.Value(), counter2.Value())
	
	fmt.Println("\n=== ReadWriter Interface ===")
	
	// TODO: Test Buffer with ReadWriter constraint
	buffer := /* create Buffer[string] */
	
	// Write some data
	testData := []string{"hello", "world", "go", "generics"}
	for _, data := range testData {
		buffer.Write(data)
	}
	
	fmt.Printf("Buffer data: %v\n", buffer.data)
	
	// TODO: Process using ReadWriter constraint
	results := /* call ProcessReadWriter with buffer */
	fmt.Printf("Processed results: %v\n", results)
	
	fmt.Println("\n=== Type Inference ===")
	
	// TODO: Demonstrate type inference - Go can infer types
	fmt.Printf("Inferred Min(42, 17): %d\n", Min(42, 17))
	fmt.Printf("Inferred Sum: %d\n", Sum([]int{1, 2, 3}))
	
	// TODO: Explicit type specification when needed
	var emptyInts []int
	fmt.Printf("Sum of empty slice: %d\n", Sum[int](emptyInts))
	
	fmt.Println("Generic constraints examples completed")
}