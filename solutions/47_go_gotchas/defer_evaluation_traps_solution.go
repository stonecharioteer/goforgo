package main

import (
	"fmt"
	"time"
)

func demonstrateDeferEvaluationTrapsFixed() {
	fmt.Println("=== Defer Evaluation - EXPLAINED ===")
	
	// Gotcha 1: Arguments are evaluated immediately, not when deferred function runs
	value := 1
	defer fmt.Printf("Deferred print: value = %d\n", value) // value=1 captured immediately
	
	value = 2
	fmt.Printf("Current value: %d\n", value)
	// Deferred function will print value=1, not 2!
	
	// FIXED: To capture the final value, use a closure
	value2 := 1
	defer func() {
		fmt.Printf("Closure deferred print: value2 = %d\n", value2) // Captures by reference
	}()
	value2 = 2
	
	// Gotcha 2: Loop with defer - EXPLAINED
	fmt.Println("\n=== Loop Defer - EXPLAINED ===")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("Deferred loop: i = %d\n", i) // Each i evaluated immediately
	}
	// These print in REVERSE order: 2, 1, 0 (LIFO - Last In, First Out)
	
	// Gotcha 3: Defer with function return values - EXPLAINED
	fmt.Printf("Function result: %d\n", deferredReturnFixed())
	
	// Gotcha 4: Defer with pointer vs value - EXPLAINED
	fmt.Println("\n=== Pointer vs Value Defer - EXPLAINED ===")
	num := 10
	defer printValue(num)     // Value 10 captured immediately
	defer printPointer(&num) // Pointer captured immediately, but *ptr changes
	
	num = 20
	fmt.Printf("Modified num: %d\n", num)
	// printPointer will show 20, printValue will show 10
	
	// Gotcha 5: Defer in loop with closure - FIXED
	fmt.Println("\n=== Defer with Closure - FIXED ===")
	for i := 0; i < 3; i++ {
		// WRONG (all print 3):
		// defer func() { fmt.Printf("Wrong: i = %d\n", i) }()
		
		// FIXED: Pass i as parameter
		defer func(index int) {
			fmt.Printf("Fixed closure defer: i = %d\n", index)
		}(i) // i evaluated immediately as argument
	}
	
	// Alternative fix: Create local copy
	for i := 0; i < 3; i++ {
		i := i // Create local copy (shadows loop variable)
		defer func() {
			fmt.Printf("Alternative fix: i = %d\n", i)
		}()
	}
	
	// Gotcha 6: Defer execution order with panic
	deferPanicExampleFixed()
}

func deferredReturnFixed() int {
	result := 1
	
	// Defers execute in LIFO order and can modify named return values
	defer func() {
		result = 10 // This executes second (after the next defer)
		fmt.Printf("Second defer sets result to: %d\n", result)
	}()
	
	defer func() {
		result = 20 // This executes first
		fmt.Printf("First defer sets result to: %d\n", result)
	}()
	
	return result // Returns 10 (modified by the second defer that runs last)
}

func namedReturnDeferExample() (result int) {
	result = 1
	
	defer func() {
		result = 100 // Modifies the named return value
	}()
	
	return result // Actually returns 100, not 1
}

func printValue(val int) {
	fmt.Printf("Deferred value: %d\n", val)
}

func printPointer(ptr *int) {
	fmt.Printf("Deferred pointer value: %d\n", *ptr)
}

func deferPanicExampleFixed() {
	fmt.Println("\n=== Defer with Panic - EXPLAINED ===")
	
	defer fmt.Println("Defer 1: Always runs (LIFO - runs last)")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Defer 2: Recovered from panic: %v\n", r)
		}
	}()
	defer fmt.Println("Defer 3: Runs before recovery (LIFO - runs first)")
	
	// Uncomment to see panic behavior
	// panic("demonstration panic")
	fmt.Println("No panic occurred - all defers still run at function end")
	
	// IMPORTANT: Defers run even during panic, in LIFO order
	// This is crucial for resource cleanup
}

func complexDeferExampleExplained() {
	fmt.Println("\n=== Complex Defer - EXPLAINED ===")
	
	counter := 0
	
	// Defers execute in LIFO (Last In, First Out) order
	defer func() { 
		counter++
		fmt.Printf("Defer 1 (runs LAST): counter = %d\n", counter) // counter will be 2
	}()
	
	defer func() { 
		counter++
		fmt.Printf("Defer 2 (runs MIDDLE): counter = %d\n", counter) // counter will be 1
	}()
	
	defer func(c int) { 
		fmt.Printf("Defer 3 (runs FIRST): captured counter = %d\n", c) // c=0 (captured immediately)
	}(counter) // counter=0 captured immediately
	
	counter = 5
	fmt.Printf("Final counter before defers: %d\n", counter)
	// Execution order: Defer 3 (c=0), Defer 2 (counter=1), Defer 1 (counter=2)
}

func deferWithResourceCleanupFixed() {
	fmt.Println("\n=== Defer Resource Cleanup - BEST PRACTICES ===")
	
	// Best practice: Check and defer immediately after resource acquisition
	file := openFile("test.txt")
	if file == nil {
		fmt.Println("Failed to open file")
		return
	}
	defer closeFile(file) // Safe because we know file is not nil
	
	// Alternative pattern for uncertain resource acquisition:
	file2 := openFileUncertain("test2.txt")
	defer func() {
		if file2 != nil {
			closeFile(file2) // Closure captures file2 by reference
		}
	}()
	
	// Simulate file operations
	writeToFile(file, "some data")
	if file2 != nil {
		writeToFile(file2, "more data")
	}
}

type File struct {
	name string
}

func openFile(name string) *File {
	fmt.Printf("Opening file: %s\n", name)
	return &File{name: name}
}

func openFileUncertain(name string) *File {
	// Simulate possible failure
	fmt.Printf("Attempting to open file: %s\n", name)
	return nil // Simulate failure
}

func closeFile(f *File) {
	if f != nil {
		fmt.Printf("Closing file: %s\n", f.name)
	}
}

func writeToFile(f *File, data string) {
	if f != nil {
		fmt.Printf("Writing '%s' to %s\n", data, f.name)
	}
}

func main() {
	fmt.Println("Starting main function...")
	
	demonstrateDeferEvaluationTrapsFixed()
	complexDeferExampleExplained()
	deferWithResourceCleanupFixed()
	
	// Demonstrate named return value modification
	fmt.Printf("Named return example: %d\n", namedReturnDeferExample())
	
	fmt.Println("\n=== End of main - all defers execute now in LIFO order ===")
	// All the deferred functions will execute in LIFO (Last In, First Out) order
	
	fmt.Println("\n=== KEY TAKEAWAYS ===")
	fmt.Println("1. Defer arguments evaluated immediately, not when function runs")
	fmt.Println("2. Defers execute in LIFO order (Last In, First Out)")
	fmt.Println("3. Defers can modify named return values")
	fmt.Println("4. Defers run even during panics (crucial for cleanup)")
	fmt.Println("5. Use closures to capture variables by reference")
	fmt.Println("6. Use parameters to capture variables by value")
}