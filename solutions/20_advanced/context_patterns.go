// context_patterns.go
// Learn advanced context patterns and cancellation in Go

package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== Advanced Context Patterns ===")
	
	fmt.Println("\n=== Basic Context Usage ===")
	
	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	// Use context in a function
	result, err := doWork(ctx, "basic task", 2*time.Second)
	if err != nil {
		fmt.Printf("Basic task failed: %v\n", err)
	} else {
		fmt.Printf("Basic task completed: %s\n", result)
	}
	
	fmt.Println("\n=== Context with Values ===")
	
	// Create context with values
	userCtx := context.WithValue(context.Background(), "userID", "12345")
	requestCtx := context.WithValue(userCtx, "requestID", "req-abc-123")
	
	// Pass context through call chain
	processRequest(requestCtx)
	
	fmt.Println("\n=== Cancellation Propagation ===")
	
	// Create cancellable context
	parentCtx, parentCancel := context.WithCancel(context.Background())
	
	// Start multiple workers
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			worker(parentCtx, workerID)
		}(i + 1)
	}
	
	// Let workers run for a bit
	time.Sleep(2 * time.Second)
	fmt.Println("Cancelling all workers...")
	parentCancel()
	
	wg.Wait()
	fmt.Println("All workers stopped")
	
	fmt.Println("\n=== Timeout Patterns ===")
	
	// Test different timeout scenarios
	testTimeouts()
	
	fmt.Println("\n=== Pipeline with Context ===")
	
	// Create pipeline context
	pipelineCtx, pipelineCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer pipelineCancel()
	
	// Run data processing pipeline
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	results := processingPipeline(pipelineCtx, numbers)
	
	// Collect results
	fmt.Println("Pipeline results:")
	for result := range results {
		fmt.Printf("  Processed: %d\n", result)
	}
	
	fmt.Println("\n=== Context Best Practices ===")
	
	// Demonstrate context best practices
	demonstrateBestPractices()
	
	// HTTP server simulation
	simulateHTTPServer()
}

// Implement work function that respects context
func doWork(ctx context.Context, taskName string, duration time.Duration) (string, error) {
	fmt.Printf("Starting %s (duration: %v)\n", taskName, duration)
	
	// Create timer for work simulation
	timer := time.NewTimer(duration)
	defer timer.Stop()
	
	select {
	case <-timer.C:
		return fmt.Sprintf("%s completed successfully", taskName), nil
	case <-ctx.Done():
		return "", fmt.Errorf("%s cancelled: %w", taskName, ctx.Err())
	}
}

// Implement request processing with context values
func processRequest(ctx context.Context) {
	// Extract values from context
	userID := ctx.Value("userID").(string)
	requestID := ctx.Value("requestID").(string)
	
	fmt.Printf("Processing request %s for user %s\n", requestID, userID)
	
	// Pass context to other functions
	authenticateUser(ctx)
	fetchUserData(ctx)
	logActivity(ctx, "request processed")
}

func authenticateUser(ctx context.Context) {
	userID, _ := ctx.Value("userID").(string)
	if userID == "" {
		userID = "unknown"
	}
	fmt.Printf("  Authenticating user: %s\n", userID)
	
	// Simulate auth work
	time.Sleep(100 * time.Millisecond)
}

func fetchUserData(ctx context.Context) {
	userID, _ := ctx.Value("userID").(string)
	if userID == "" {
		userID = "unknown"
	}
	requestID, _ := ctx.Value("requestID").(string)
	if requestID == "" {
		requestID = "unknown"
	}
	fmt.Printf("  Fetching data for user %s (request: %s)\n", userID, requestID)
	
	// Simulate data fetch
	time.Sleep(200 * time.Millisecond)
}

func logActivity(ctx context.Context, activity string) {
	userID, _ := ctx.Value("userID").(string)
	if userID == "" {
		userID = "unknown"
	}
	requestID, _ := ctx.Value("requestID").(string)
	if requestID == "" {
		requestID = "unknown"
	}
	fmt.Printf("  LOG [%s][%s]: %s\n", requestID, userID, activity)
}

// Implement worker function
func worker(ctx context.Context, workerID int) {
	fmt.Printf("Worker %d starting\n", workerID)
	
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			fmt.Printf("Worker %d is working...\n", workerID)
		case <-ctx.Done():
			fmt.Printf("Worker %d stopping: %v\n", workerID, ctx.Err())
			return
		}
	}
}

// Implement timeout testing
func testTimeouts() {
	testCases := []struct {
		name    string
		timeout time.Duration
		work    time.Duration
	}{
		{"Quick task", 2 * time.Second, 1 * time.Second},
		{"Slow task", 1 * time.Second, 2 * time.Second},
		{"Just in time", 1 * time.Second, 1 * time.Second},
	}
	
	for _, tc := range testCases {
		fmt.Printf("Testing: %s (timeout: %v, work: %v)\n", tc.name, tc.timeout, tc.work)
		
		ctx, cancel := context.WithTimeout(context.Background(), tc.timeout)
		
		_, err := doWork(ctx, tc.name, tc.work)
		if err != nil {
			fmt.Printf("  Result: %v\n", err)
		} else {
			fmt.Printf("  Result: Success\n")
		}
		
		cancel() // Always clean up
		fmt.Println()
	}
}

// Implement processing pipeline with context
func processingPipeline(ctx context.Context, numbers []int) <-chan int {
	output := make(chan int)
	
	go func() {
		defer close(output)
		
		for _, num := range numbers {
			// Check for cancellation
			select {
			case <-ctx.Done():
				fmt.Printf("Pipeline cancelled: %v\n", ctx.Err())
				return
			default:
				// Continue processing
			}
			
			// Simulate processing time
			time.Sleep(time.Duration(rand.Intn(200)+100) * time.Millisecond)
			
			// Process the number (square it)
			processed := num * num
			
			// Send result or handle cancellation
			select {
			case output <- processed:
				// Successfully sent
			case <-ctx.Done():
				fmt.Printf("Pipeline cancelled during send: %v\n", ctx.Err())
				return
			}
		}
	}()
	
	return output
}

// Implement context best practices demonstration
func demonstrateBestPractices() {
	fmt.Println("Context Best Practices:")
	fmt.Println("1. Always pass context as first parameter")
	fmt.Println("2. Don't store contexts in structs")
	fmt.Println("3. Always call cancel() to avoid resource leaks")
	fmt.Println("4. Use context.WithValue() sparingly")
	fmt.Println("5. Check context.Done() in long-running operations")
	
	// Demonstrate deadline vs timeout
	fmt.Println("\n=== Deadline vs Timeout ===")
	
	// Using WithDeadline
	deadline := time.Now().Add(2 * time.Second)
	deadlineCtx, deadlineCancel := context.WithDeadline(context.Background(), deadline)
	defer deadlineCancel()
	
	fmt.Printf("Using WithDeadline (deadline: %v)\n", deadline.Format("15:04:05"))
	_, err := doWork(deadlineCtx, "deadline task", 3*time.Second)
	if err != nil {
		fmt.Printf("Deadline task result: %v\n", err)
	}
	
	// Using WithTimeout
	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer timeoutCancel()
	
	fmt.Printf("Using WithTimeout (timeout: 2s)\n")
	_, err = doWork(timeoutCtx, "timeout task", 3*time.Second)
	if err != nil {
		fmt.Printf("Timeout task result: %v\n", err)
	}
	
	// Demonstrate context composition
	fmt.Println("\n=== Context Composition ===")
	
	baseCtx, baseCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer baseCancel()
	
	userCtx := context.WithValue(baseCtx, "userID", "comp-user-123")
	
	childCtx, childCancel := context.WithTimeout(userCtx, 2*time.Second)
	defer childCancel()
	
	fmt.Println("Child context inherits from parent but has shorter timeout")
	processCompositeRequest(childCtx)
}

// Implement composite request processing
func processCompositeRequest(ctx context.Context) {
	userID, _ := ctx.Value("userID").(string)
	if userID == "" {
		userID = "unknown"
	}
	
	fmt.Printf("Processing composite request for user: %s\n", userID)
	
	// Check context deadline
	if deadline, ok := ctx.Deadline(); ok {
		remaining := time.Until(deadline)
		fmt.Printf("  Time remaining: %v\n", remaining)
	}
	
	// Simulate work that might be cancelled
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("  Composite request completed")
	case <-ctx.Done():
		fmt.Printf("  Composite request cancelled: %v\n", ctx.Err())
	}
}

// Implement HTTP-like server simulation with context
func simulateHTTPServer() {
	fmt.Println("\n=== HTTP Server Simulation ===")
	
	// Simulate handling multiple requests with context
	requests := []string{"req1", "req2", "req3"}
	
	var wg sync.WaitGroup
	for _, reqID := range requests {
		wg.Add(1)
		go func(requestID string) {
			defer wg.Done()
			
			// Create request context with timeout
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			
			// Add request ID to context
			requestCtx := context.WithValue(ctx, "requestID", requestID)
			
			handleHTTPRequest(requestCtx)
		}(reqID)
	}
	
	wg.Wait()
	fmt.Println("All requests processed")
}

// Implement HTTP request handler
func handleHTTPRequest(ctx context.Context) {
	requestID, _ := ctx.Value("requestID").(string)
	if requestID == "" {
		requestID = "unknown"
	}
	fmt.Printf("Handling HTTP request: %s\n", requestID)
	
	// Simulate request processing steps
	steps := []struct {
		name     string
		duration time.Duration
	}{
		{"Authentication", 100 * time.Millisecond},
		{"Database Query", 300 * time.Millisecond},
		{"Business Logic", 200 * time.Millisecond},
		{"Response Generation", 150 * time.Millisecond},
	}
	
	for _, step := range steps {
		select {
		case <-ctx.Done():
			fmt.Printf("  Request %s cancelled during %s: %v\n", 
				requestID, step.name, ctx.Err())
			return
		case <-time.After(step.duration):
			fmt.Printf("  %s: %s completed\n", requestID, step.name)
		}
	}
	
	fmt.Printf("  Request %s completed successfully\n", requestID)
}