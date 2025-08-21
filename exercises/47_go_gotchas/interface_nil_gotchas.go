package main

import "fmt"

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
	fmt.Println("=== Interface Nil Gotchas ===")
	
	// Gotcha 1: Nil interface vs nil concrete value
	var writer Writer // nil interface
	fmt.Printf("writer == nil: %t\n", writer == nil)
	fmt.Printf("writer: %v\n", writer)
	
	// This will panic - interface is nil
	// writer.Write("test") // Uncomment to see panic
	
	// Gotcha 2: Interface with nil concrete value is NOT nil
	var fileWriter *FileWriter // nil pointer
	writer = fileWriter         // interface now holds nil *FileWriter
	
	fmt.Printf("fileWriter == nil: %t\n", fileWriter == nil)
	fmt.Printf("writer == nil: %t\n", writer == nil) // This is false!
	fmt.Printf("writer: %v\n", writer)
	
	// This won't panic - the interface is not nil, even though the concrete value is
	err := writer.Write("test data")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	
	// Gotcha 3: Function returning interface with nil concrete value
	writer = getNilWriter() // Returns interface containing nil pointer
	fmt.Printf("getNilWriter() == nil: %t\n", writer == nil) // false!
	
	// Gotcha 4: Checking if interface holds a nil concrete value
	if writer != nil {
		fmt.Println("Writer is not nil (interface-wise)")
		// How do you check if the concrete value inside is nil?
		// This is a common source of bugs
		
		// Wrong way to check:
		if writer == (*FileWriter)(nil) {
			fmt.Println("This check doesn't work as expected")
		}
	}
}

func getNilWriter() Writer {
	// This function has a bug - it returns a nil pointer wrapped in an interface
	var fw *FileWriter // nil pointer
	return fw          // interface{(*FileWriter)(nil)} - NOT nil interface!
	
	// The correct return would be:
	// return nil // This returns a truly nil interface
}

func checkWriterCorrectly(w Writer) {
	// How do you properly check if an interface contains a nil concrete value?
	// This function demonstrates the wrong and right ways
	
	if w == nil {
		fmt.Println("Writer interface is nil")
		return
	}
	
	// Wrong: This doesn't work for checking nil concrete values
	// if w == (*FileWriter)(nil) { ... }
	
	// One approach: Type assertion and nil check
	if fw, ok := w.(*FileWriter); ok {
		if fw == nil {
			fmt.Println("Writer contains nil *FileWriter")
			return
		}
		fmt.Println("Writer contains valid *FileWriter")
	}
	
	// Another approach: reflection (more general but more complex)
	// We'll explore this in the solution
}

func interfaceComparisonTrap() {
	fmt.Println("\n=== Interface Comparison Trap ===")
	
	// Gotcha 5: Interface comparison doesn't work as expected with nil values
	var w1 Writer = (*FileWriter)(nil)
	var w2 Writer = (*NetworkWriter)(nil)
	
	fmt.Printf("w1 == nil: %t\n", w1 == nil) // false
	fmt.Printf("w2 == nil: %t\n", w2 == nil) // false
	fmt.Printf("w1 == w2: %t\n", w1 == w2)   // false - different types!
	
	// This comparison might surprise you
	var w3 Writer = (*FileWriter)(nil)
	fmt.Printf("w1 == w3: %t\n", w1 == w3) // true - same type and value
	
	// But this creates subtle bugs in error handling:
	err := returnNilError()
	if err != nil {
		fmt.Printf("Got error (but error is actually nil!): %v\n", err)
		// This prints "Got error: <nil>" - confusing!
	}
}

func returnNilError() error {
	// This function has a bug - it returns a nil pointer wrapped as error interface
	var customErr *CustomError
	if someCondition() {
		customErr = &CustomError{message: "something went wrong"}
	}
	return customErr // Bug: returns interface{(*CustomError)(nil)} when customErr is nil
	
	// Fix would be:
	// if customErr == nil {
	//     return nil
	// }
	// return customErr
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
	interfaceComparisonTrap()
	
	fmt.Println("\n=== Key Takeaways ===")
	fmt.Println("1. nil interface != interface containing nil concrete value")
	fmt.Println("2. Interface with nil concrete value is not == nil")
	fmt.Println("3. This causes bugs in error handling and nil checks")
	fmt.Println("4. Always return nil (not nil concrete type) from interface-returning functions")
}