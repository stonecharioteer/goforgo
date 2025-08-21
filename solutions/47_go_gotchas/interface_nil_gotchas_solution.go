package main

import (
	"fmt"
	"reflect"
)

// Interface definition
type Writer interface {
	Write(data string) error
}

// Concrete implementation
type FileWriter struct {
	filename string
}

func (fw *FileWriter) Write(data string) error {
	if fw == nil {
		return fmt.Errorf("FileWriter is nil")
	}
	fmt.Printf("Writing '%s' to %s\n", data, fw.filename)
	return nil
}

// Another implementation
type NetworkWriter struct {
	url string
}

func (nw *NetworkWriter) Write(data string) error {
	if nw == nil {
		return fmt.Errorf("NetworkWriter is nil")
	}
	fmt.Printf("Sending '%s' to %s\n", data, nw.url)
	return nil
}

func demonstrateInterfaceNilGotchas() {
	fmt.Println("=== Interface Nil Gotchas - EXPLAINED ===")
	
	// Gotcha 1: Nil interface vs nil concrete value
	var writer Writer // nil interface (both type and value are nil)
	fmt.Printf("writer == nil: %t\n", writer == nil)           // true
	fmt.Printf("writer: %v\n", writer)                         // <nil>
	fmt.Printf("writer type: %T\n", writer)                    // <nil>
	
	// This would panic because interface itself is nil
	// writer.Write("test") // runtime panic: invalid memory address
	
	// Gotcha 2: Interface with nil concrete value is NOT nil
	var fileWriter *FileWriter // nil pointer to FileWriter
	writer = fileWriter         // interface now holds (*FileWriter, nil)
	
	fmt.Printf("fileWriter == nil: %t\n", fileWriter == nil)   // true
	fmt.Printf("writer == nil: %t\n", writer == nil)           // FALSE! Interface is not nil
	fmt.Printf("writer: %v\n", writer)                         // <nil> (but interface is not nil)
	fmt.Printf("writer type: %T\n", writer)                    // *main.FileWriter
	
	// This works because interface has type info, even though concrete value is nil
	err := writer.Write("test data")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	
	// Gotcha 3: Function returning interface with nil concrete value
	writer = getCorrectNilWriter() // FIXED: Returns truly nil interface
	fmt.Printf("getCorrectNilWriter() == nil: %t\n", writer == nil) // true
	
	// Gotcha 4: FIXED - Properly checking if interface holds a nil concrete value
	writer = getNilWriter() // This still returns interface with nil concrete value
	checkWriterCorrectly(writer)
}

func getNilWriter() Writer {
	// This demonstrates the common mistake
	var fw *FileWriter // nil pointer
	return fw          // Returns interface{(*FileWriter)(nil)} - NOT nil interface!
}

func getCorrectNilWriter() Writer {
	// FIXED: Return truly nil interface when you want nil
	return nil // This returns interface{(nil, nil)}
	
	// Or with explicit check:
	var fw *FileWriter
	if fw == nil {
		return nil // Return nil interface, not interface containing nil pointer
	}
	return fw
}

func checkWriterCorrectly(w Writer) {
	// FIXED: Proper ways to check if interface contains nil concrete value
	
	if w == nil {
		fmt.Println("Writer interface is nil")
		return
	}
	
	// Method 1: Type assertion and nil check (type-specific)
	if fw, ok := w.(*FileWriter); ok {
		if fw == nil {
			fmt.Println("Writer contains nil *FileWriter")
			return
		}
		fmt.Println("Writer contains valid *FileWriter")
		return
	}
	
	if nw, ok := w.(*NetworkWriter); ok {
		if nw == nil {
			fmt.Println("Writer contains nil *NetworkWriter")
			return
		}
		fmt.Println("Writer contains valid *NetworkWriter")
		return
	}
	
	// Method 2: Reflection (general solution)
	v := reflect.ValueOf(w)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		fmt.Println("Writer contains nil pointer (detected via reflection)")
		return
	}
	
	fmt.Println("Writer contains valid concrete value")
}

func isInterfaceNil(i interface{}) bool {
	// Generic function to check if interface contains nil concrete value
	if i == nil {
		return true
	}
	
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}

func interfaceComparisonTrapFixed() {
	fmt.Println("\n=== Interface Comparison - EXPLAINED ===")
	
	// Interface contains (type, value) pair
	var w1 Writer = (*FileWriter)(nil)    // (*FileWriter, nil)
	var w2 Writer = (*NetworkWriter)(nil) // (*NetworkWriter, nil)
	
	fmt.Printf("w1 == nil: %t\n", w1 == nil) // false - interface has type info
	fmt.Printf("w2 == nil: %t\n", w2 == nil) // false - interface has type info
	fmt.Printf("w1 == w2: %t\n", w1 == w2)   // false - different types
	
	// Same type and value comparison
	var w3 Writer = (*FileWriter)(nil) // (*FileWriter, nil)
	fmt.Printf("w1 == w3: %t\n", w1 == w3) // true - same (type, value) pair
	
	// FIXED error handling example
	err := returnCorrectError()
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
	} else {
		fmt.Println("No error - correctly returns nil interface")
	}
	
	// Demonstrate the nil interface check
	nilWriter := getNilWriter()
	fmt.Printf("getNilWriter() contains nil concrete value: %t\n", isInterfaceNil(nilWriter))
}

func returnCorrectError() error {
	// FIXED: Proper error return pattern
	var customErr *CustomError
	if someCondition() {
		customErr = &CustomError{message: "something went wrong"}
	}
	
	// FIXED: Return nil interface when error is nil
	if customErr == nil {
		return nil // Return nil interface, not interface containing nil pointer
	}
	return customErr
}

type CustomError struct {
	message string
}

func (e *CustomError) Error() string {
	if e == nil {
		return "<nil>"
	}
	return e.message
}

func someCondition() bool {
	return false // Simulate condition that doesn't create error
}

func main() {
	demonstrateInterfaceNilGotchas()
	interfaceComparisonTrapFixed()
	
	fmt.Println("\n=== Key Takeaways - FIXED ===")
	fmt.Println("1. Interface stores (type, value) - nil interface has both nil")
	fmt.Println("2. Interface with nil concrete value has (type, nil) - NOT nil interface")
	fmt.Println("3. Always return explicit nil from interface-returning functions when you want nil")
	fmt.Println("4. Use type assertions or reflection to check nil concrete values")
	fmt.Println("5. This is Go-specific behavior - most languages don't have this distinction")
}