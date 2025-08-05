// goroutines_basics.go
// Learn the fundamentals of goroutines - Go's lightweight threads

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// TODO: Simple function to run as goroutine
func sayHello(name string, delay time.Duration) {
	// Sleep for the specified delay
	// Print a greeting message with the name
}

// TODO: Function that demonstrates goroutine execution order
func printNumbers(prefix string, start, end int) {
	for i := start; i <= end; i++ {
		fmt.Printf("%s: %d\n", prefix, i)
		// Add small delay to see concurrent execution
		time.Sleep(100 * time.Millisecond)
	}
}

// TODO: Function that shows goroutine with WaitGroup
func workerWithWaitGroup(id int, wg *sync.WaitGroup) {
	// Signal that this goroutine is done when function exits
	// defer wg.Done()
	
	fmt.Printf("Worker %d starting\n", id)
	
	// Simulate some work
	time.Sleep(time.Duration(id*200) * time.Millisecond)
	
	fmt.Printf("Worker %d finished\n", id)
}

// TODO: Function that returns a value through channel
func calculate(n int, result chan<- int) {
	// Calculate sum of numbers from 1 to n
	sum := 0
	// Calculate sum
	
	// Send result to channel
	// result <- sum
}

func main() {
	fmt.Println("=== Basic Goroutines ===")
	
	// TODO: Start goroutines without waiting
	fmt.Println("Starting goroutines...")
	// Start sayHello goroutines for "Alice", "Bob", "Charlie"
	// Use different delays: 1s, 500ms, 200ms
	
	// Main goroutine continues immediately
	fmt.Println("Main goroutine continues...")
	
	// TODO: Wait a bit to see some output
	time.Sleep(2 * time.Second)
	
	fmt.Println("\n=== Concurrent Number Printing ===")
	
	// TODO: Start concurrent number printing
	// Start printNumbers("Odd", 1, 5) as goroutine
	// Start printNumbers("Even", 2, 6) as goroutine
	
	// Wait for goroutines to complete
	time.Sleep(1 * time.Second)
	
	fmt.Println("\n=== Using WaitGroup ===")
	
	// TODO: Use sync.WaitGroup to wait for goroutines
	var wg sync.WaitGroup
	
	numWorkers := 3
	// Set the number of goroutines to wait for
	// wg.Add(numWorkers)
	
	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		// Start workerWithWaitGroup as goroutine
	}
	
	// Wait for all goroutines to complete
	// wg.Wait()
	fmt.Println("All workers completed")
	
	fmt.Println("\n=== Goroutines with Channels ===")
	
	// TODO: Use channel to collect results from goroutines
	resultChan := make(chan int, 3) // Buffered channel
	
	// Start calculation goroutines
	numbers := []int{10, 100, 1000}
	for _, n := range numbers {
		// Start calculate goroutine for each number
	}
	
	// Collect results
	for i := 0; i < len(numbers); i++ {
		// Receive from channel and print result
	}
	
	fmt.Println("\n=== Runtime Information ===")
	
	// TODO: Show runtime information
	fmt.Printf("Number of goroutines: %d\n", /* get number of goroutines */)
	fmt.Printf("Number of CPUs: %d\n", /* get number of CPUs */)
	fmt.Printf("GOMAXPROCS: %d\n", /* get GOMAXPROCS */)
	
	// TODO: Anonymous goroutine
	fmt.Println("\n=== Anonymous Goroutine ===")
	
	done := make(chan bool)
	
	// Start anonymous goroutine
	go func() {
		// Do some work
		fmt.Println("Anonymous goroutine working...")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Anonymous goroutine done")
		// Signal completion
		// done <- true
	}()
	
	// Wait for anonymous goroutine to complete
	// <-done
	
	fmt.Println("Program completed")
}