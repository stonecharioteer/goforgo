package main

import (
	"fmt"
	"time"
)

func demonstrateDeferEvaluationTraps() {
	fmt.Println("=== Defer Evaluation Traps ===")
	
	// Gotcha 1: Arguments are evaluated immediately, not when deferred function runs
	value := 1
	defer fmt.Printf("Deferred print: value = %d\n", value) // value evaluated NOW
	
	value = 2
	fmt.Printf("Current value: %d\n", value)
	// When deferred function runs, what value will be printed?
	
	// Gotcha 2: Loop with defer - all defers use the same variable
	fmt.Println("\n=== Loop Defer Trap ===")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("Deferred loop: i = %d\n", i) // i evaluated immediately
	}
	// What order will these print in?
	
	// Gotcha 3: Defer with function return values
	fmt.Printf("Function result: %d\n", deferredReturnTrap())
	
	// Gotcha 4: Defer with pointer vs value
	fmt.Println("\n=== Pointer vs Value Defer ===")
	num := 10
	defer printValue(num)     // Value evaluated immediately
	defer printPointer(&num) // Pointer evaluated immediately, but pointed value can change
	
	num = 20
	fmt.Printf("Modified num: %d\n", num)
	
	// Gotcha 5: Defer in loop with closure - classic trap
	fmt.Println("\n=== Defer with Closure Trap ===")
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Printf("Closure defer: i = %d\n", i) // i is captured by reference!
		}()
	}
	
	// Gotcha 6: Defer execution order with panic
	deferPanicExample()
}

func deferredReturnTrap() int {
	result := 1
	
	defer func() {
		result = 10 // This modifies the return value
	}()
	
	defer func() {
		result = 20 // This also modifies the return value
	}()
	
	return result // What value is actually returned?
}

func printValue(val int) {
	fmt.Printf("Deferred value: %d\n", val)
}

func printPointer(ptr *int) {
	fmt.Printf("Deferred pointer value: %d\n", *ptr)
}

func deferPanicExample() {
	fmt.Println("\n=== Defer with Panic ===")
	
	defer fmt.Println("This defer runs even during panic")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()
	defer fmt.Println("This defer also runs before recovery")
	
	// Uncomment to see panic behavior
	// panic("demonstration panic")
	fmt.Println("No panic occurred")
}

func complexDeferExample() {
	fmt.Println("\n=== Complex Defer Example ===")
	
	counter := 0
	
	// Multiple defers with shared state
	defer func() { 
		counter++
		fmt.Printf("Defer 1: counter = %d\n", counter) 
	}()
	
	defer func() { 
		counter++
		fmt.Printf("Defer 2: counter = %d\n", counter) 
	}()
	
	defer func(c int) { 
		fmt.Printf("Defer 3: captured counter = %d\n", c) 
	}(counter) // counter value captured immediately
	
	counter = 5
	fmt.Printf("Final counter before defers: %d\n", counter)
	// What will the deferred functions print?
}

func deferWithResourceCleanup() {
	fmt.Println("\n=== Defer Resource Cleanup Pattern ===")
	
	// Simulate opening a file
	file := openFile("test.txt")
	defer func() {
		if file != nil {
			closeFile(file) // file reference evaluated immediately
		}
	}()
	
	// Simulate file operations
	writeToFile(file, "some data")
	
	// What happens if openFile returns nil?
	// Is the defer still safe?
}

type File struct {
	name string
}

func openFile(name string) *File {
	fmt.Printf("Opening file: %s\n", name)
	return &File{name: name}
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
	demonstrateDeferEvaluationTraps()
	complexDeferExample()
	deferWithResourceCleanup()
	
	fmt.Println("\n=== End of main - defers execute now ===")
	// All the deferred functions will execute in LIFO order
}