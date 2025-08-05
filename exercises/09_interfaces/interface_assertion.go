// interface_assertion.go
// Learn advanced type assertion patterns and interface checking

package main

import (
	"fmt"
	"strconv"
)

// TODO: Define interfaces for this exercise
type Stringer interface {
	String() string
}

type Counter interface {
	Count() int
}

type Resetter interface {
	Reset()
}

// TODO: Define a Person struct that implements Stringer
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	// Return formatted string representation
}

// TODO: Define a WordCounter that implements Stringer, Counter, and Resetter
type WordCounter struct {
	words int
}

func (wc *WordCounter) AddWords(text string) {
	// Simple word counting (split by spaces)
	if text == "" {
		return
	}
	// Count words by splitting on spaces and add to wc.words
}

func (wc WordCounter) String() string {
	// Return string representation
}

func (wc WordCounter) Count() int {
	// Return current count
}

func (wc *WordCounter) Reset() {
	// Reset counter to zero
}

// TODO: Function to check if a value implements Stringer
func checkStringer(value interface{}) {
	if stringer, ok := /* check if value implements Stringer */; ok {
		fmt.Printf("✓ Implements Stringer: %s\n", stringer.String())
	} else {
		fmt.Printf("✗ Does not implement Stringer: %v\n", value)
	}
}

// TODO: Function to check multiple interfaces
func checkInterfaces(value interface{}) {
	fmt.Printf("Checking interfaces for %T:\n", value)
	
	// Check Stringer
	if _, ok := /* check Stringer */; ok {
		fmt.Println("  ✓ Implements Stringer")
	} else {
		fmt.Println("  ✗ Does not implement Stringer")
	}
	
	// Check Counter
	if _, ok := /* check Counter */; ok {
		fmt.Println("  ✓ Implements Counter")
	} else {
		fmt.Println("  ✗ Does not implement Counter")
	}
	
	// Check Resetter
	if _, ok := /* check Resetter */; ok {
		fmt.Println("  ✓ Implements Resetter")
	} else {
		fmt.Println("  ✗ Does not implement Resetter")
	}
}

// TODO: Function that works with any type that can be converted to string
func convertToString(value interface{}) string {
	// Try different conversion methods in order of preference
	
	// First, check if it implements Stringer
	if stringer, ok := /* check Stringer */; ok {
		return stringer.String()
	}
	
	// Then check for basic types and convert them
	switch v := value.(type) {
	case string:
		return v
	case int:
		// Convert int to string
	case float64:
		// Convert float64 to string with 2 decimal places
	case bool:
		// Convert bool to string
	default:
		// Use fmt.Sprintf as fallback
		return fmt.Sprintf("%v", v)
	}
}

// TODO: Function that accepts interface and tries to call methods conditionally
func processData(data interface{}) {
	fmt.Printf("Processing %T: ", data)
	
	// Always print the value
	fmt.Printf("Value = %v", data)
	
	// If it's a Stringer, use its String method
	if stringer, ok := /* check Stringer */; ok {
		fmt.Printf(", String() = %s", stringer.String())
	}
	
	// If it's a Counter, show the count
	if counter, ok := /* check Counter */; ok {
		fmt.Printf(", Count() = %d", counter.Count())
	}
	
	fmt.Println()
}

func main() {
	// TODO: Create test objects
	person := // Create Person with Name: "Alice", Age: 30
	wc := // Create WordCounter
	wc.AddWords("hello world from Go")
	
	plainInt := 42
	plainString := "just a string"
	
	fmt.Println("=== Stringer Interface Check ===")
	// TODO: Test checkStringer with different values
	// Test with: person, &wc, plainInt, plainString
	
	fmt.Println("\n=== Multiple Interface Check ===")
	// TODO: Test checkInterfaces with different values
	// Test with: person, &wc, plainInt
	
	fmt.Println("\n=== String Conversion ===")
	values := []interface{}{/* add: person, &wc, 123, 3.14159, true, []int{1,2,3} */}
	
	for _, value := range values {
		str := convertToString(value)
		fmt.Printf("%T -> \"%s\"\n", value, str)
	}
	
	fmt.Println("\n=== Conditional Method Calls ===")
	// TODO: Test processData with different values
	// Test with: person, &wc, plainInt, plainString
	
	fmt.Println("\n=== Interface Nil Check ===")
	var nilStringer Stringer
	var nilPerson *Person
	
	// TODO: Check for nil interfaces
	if nilStringer == nil {
		fmt.Println("nilStringer is nil")
	}
	
	// This is trickier - interface with nil pointer
	nilStringer = nilPerson
	if nilStringer == nil {
		fmt.Println("nilStringer with nil pointer is nil")
	} else {
		fmt.Printf("nilStringer with nil pointer is not nil: %T\n", nilStringer)
		// TODO: Safely call method on potentially nil interface
		// This would panic: nilStringer.String()
	}
	
	fmt.Println("\n=== Counter Operations ===")
	if counter, ok := /* check if wc implements Counter */; ok {
		fmt.Printf("Initial count: %d\n", counter.Count())
		wc.AddWords("more words here")
		fmt.Printf("After adding words: %d\n", counter.Count())
		
		if resetter, ok := /* check if wc implements Resetter */; ok {
			resetter.Reset()
			fmt.Printf("After reset: %d\n", counter.Count())
		}
	}
}