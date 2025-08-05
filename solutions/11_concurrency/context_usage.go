// context_usage.go - SOLUTION
// Learn how to use context for cancellation, timeouts, and passing values

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simulate a long-running operation
func simulateWork(ctx context.Context, name string, duration time.Duration) error {
	fmt.Printf("%s: Starting work (duration: %v)\n", name, duration)
	
	// Use context-aware sleep
	select {
	case <-time.After(duration):
		fmt.Printf("%s: Work completed successfully\n", name)
		return nil
	case <-ctx.Done():
		fmt.Printf("%s: Work cancelled: %v\n", name, ctx.Err())
		return ctx.Err()
	}
}

// Function that performs multiple operations with context
func processData(ctx context.Context, data []string) error {
	for i, item := range data {
		// Check if context is cancelled before each operation
		select {
		case <-ctx.Done():
			fmt.Printf("Processing cancelled at item %d: %v\n", i, ctx.Err())
			return ctx.Err()
		default:
			// Continue processing
		}
		
		fmt.Printf("Processing item %d: %s\n", i+1, item)
		
		// Simulate work with context-aware sleep
		sleepDuration := time.Duration(rand.Intn(500)+200) * time.Millisecond
		select {
		case <-time.After(sleepDuration):
			// Work completed
		case <-ctx.Done():
			fmt.Printf("Processing cancelled during item %d: %v\n", i+1, ctx.Err())
			return ctx.Err()
		}
	}
	
	fmt.Println("All data processed successfully")
	return nil
}

// Function that uses context values
func authenticatedOperation(ctx context.Context) error {
	// Extract user from context
	user, ok := ctx.Value("user").(string)
	if !ok {
		return fmt.Errorf("user not found in context")
	}
	
	// Extract request ID from context
	requestID, ok := ctx.Value("requestID").(string)
	if !ok {
		requestID = "unknown"
	}
	
	fmt.Printf("[%s] User '%s' performing authenticated operation\n", requestID, user)
	
	// Simulate operation with context
	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Printf("[%s] Operation completed for user '%s'\n", requestID, user)
		return nil
	case <-ctx.Done():
		fmt.Printf("[%s] Operation cancelled for user '%s': %v\n", requestID, user, ctx.Err())
		return ctx.Err()
	}
}

// Worker pool with context
func worker(ctx context.Context, id int, jobs <-chan string, results chan<- string) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d: Job channel closed\n", id)
				return
			}
			
			fmt.Printf("Worker %d: Processing job '%s'\n", id, job)
			
			// Simulate work with random duration
			workDuration := time.Duration(rand.Intn(1000)+500) * time.Millisecond
			
			select {
			case <-time.After(workDuration):
				result := fmt.Sprintf("Worker %d completed '%s'", id, job)
				select {
				case results <- result:
					fmt.Printf("Worker %d: Sent result\n", id)
				case <-ctx.Done():
					fmt.Printf("Worker %d: Cancelled while sending result\n", id)
					return
				}
			case <-ctx.Done():
				fmt.Printf("Worker %d: Job '%s' cancelled: %v\n", id, job, ctx.Err())
				return
			}
			
		case <-ctx.Done():
			fmt.Printf("Worker %d: Shutting down: %v\n", id, ctx.Err())
			return
		}
	}
}

func main() {
	fmt.Println("=== Context with Timeout ===")
	
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	// Start operations with different durations
	operations := []struct {
		name     string
		duration time.Duration
	}{
		{"FastOp", 500 * time.Millisecond},
		{"SlowOp", 3 * time.Second}, // This will timeout
		{"MediumOp", 1 * time.Second},
	}
	
	var wg sync.WaitGroup
	for _, op := range operations {
		wg.Add(1)
		go func(name string, duration time.Duration) {
			defer wg.Done()
			err := simulateWork(ctx, name, duration)
			if err != nil {
				fmt.Printf("%s failed: %v\n", name, err)
			}
		}(op.name, op.duration)
	}
	
	wg.Wait()
	fmt.Println()
	
	fmt.Println("=== Context with Cancellation ===")
	
	// Create cancellable context
	ctx2, cancel2 := context.WithCancel(context.Background())
	
	// Start data processing
	data := []string{"item1", "item2", "item3", "item4", "item5"}
	
	go func() {
		err := processData(ctx2, data)
		if err != nil {
			fmt.Printf("Data processing failed: %v\n", err)
		}
	}()
	
	// Cancel after 1 second
	time.Sleep(1 * time.Second)
	fmt.Println("Cancelling data processing...")
	cancel2()
	
	time.Sleep(500 * time.Millisecond) // Give time for cancellation to take effect
	fmt.Println()
	
	fmt.Println("=== Context with Values ===")
	
	// Create context with values
	baseCtx := context.Background()
	ctxWithUser := context.WithValue(baseCtx, "user", "alice")
	ctxWithRequestID := context.WithValue(ctxWithUser, "requestID", "req-12345")
	
	// Add timeout to the context chain
	ctx3, cancel3 := context.WithTimeout(ctxWithRequestID, 5*time.Second)
	defer cancel3()
	
	// Test authenticated operations
	operations2 := []string{"operation1", "operation2", "operation3"}
	
	for _, op := range operations2 {
		// Create context with operation-specific request ID
		opCtx := context.WithValue(ctx3, "requestID", fmt.Sprintf("req-%s", op))
		
		err := authenticatedOperation(opCtx)
		if err != nil {
			fmt.Printf("Operation %s failed: %v\n", op, err)
		}
		
		time.Sleep(50 * time.Millisecond) // Small delay between operations
	}
	
	fmt.Println("\n=== Worker Pool with Context ===")
	
	// Create context with timeout for worker pool
	ctx4, cancel4 := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel4()
	
	// Create channels
	jobs := make(chan string, 10)
	results := make(chan string, 10)
	
	// Start workers
	numWorkers := 3
	var workerWg sync.WaitGroup
	
	for i := 1; i <= numWorkers; i++ {
		workerWg.Add(1)
		go func(id int) {
			defer workerWg.Done()
			worker(ctx4, id, jobs, results)
		}(i)
	}
	
	// Send jobs
	jobList := []string{"job1", "job2", "job3", "job4", "job5", "job6"}
	go func() {
		for _, job := range jobList {
			select {
			case jobs <- job:
				fmt.Printf("Sent job: %s\n", job)
			case <-ctx4.Done():
				fmt.Printf("Job sending cancelled: %v\n", ctx4.Err())
				break
			}
			time.Sleep(200 * time.Millisecond) // Stagger job sending
		}
		close(jobs)
	}()
	
	// Collect results
	go func() {
		for {
			select {
			case result, ok := <-results:
				if !ok {
					fmt.Println("Results channel closed")
					return
				}
				fmt.Printf("Received result: %s\n", result)
			case <-ctx4.Done():
				fmt.Printf("Result collection cancelled: %v\n", ctx4.Err())
				return
			}
		}
	}()
	
	// Wait for workers to finish
	workerWg.Wait()
	close(results)
	
	fmt.Println("\n=== Context Best Practices ===")
	
	// Demonstrate context inheritance
	parentCtx, parentCancel := context.WithCancel(context.Background())
	
	// Child context with timeout
	childCtx, childCancel := context.WithTimeout(parentCtx, 2*time.Second)
	
	// Grandchild context with value
	grandchildCtx := context.WithValue(childCtx, "level", "grandchild")
	
	fmt.Println("Testing context hierarchy:")
	
	go func() {
		// Operation using grandchild context
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Grandchild operation completed")
		case <-grandchildCtx.Done():
			fmt.Printf("Grandchild operation cancelled: %v\n", grandchildCtx.Err())
		}
	}()
	
	// Cancel parent after 500ms - should cascade to children
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Cancelling parent context...")
	parentCancel()
	
	time.Sleep(100 * time.Millisecond) // Give time for cancellation
	
	// Clean up remaining contexts
	childCancel()
	
	fmt.Println("Context examples completed")
}