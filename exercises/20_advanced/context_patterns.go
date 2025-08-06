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
	
	// TODO: Create context with timeout
	ctx, cancel := /* create context with 3 second timeout */
	defer cancel()
	
	// TODO: Use context in a function
	result, err := doWork(ctx, "basic task", 2*time.Second)
	if err != nil {
		fmt.Printf("Basic task failed: %v\n", err)
	} else {
		fmt.Printf("Basic task completed: %s\n", result)
	}
	
	fmt.Println("\n=== Context with Values ===")
	
	// TODO: Create context with values
	userCtx := /* add user ID "12345" to context */
	requestCtx := /* add request ID "req-abc-123" to userCtx */
	
	// TODO: Pass context through call chain
	processRequest(requestCtx)
	
	fmt.Println("\n=== Cancellation Propagation ===")
	
	// TODO: Create cancellable context
	parentCtx, parentCancel := /* create cancellable context */
	
	// Start multiple workers
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			/* start worker with parentCtx */
		}(i + 1)
	}
	
	// Let workers run for a bit
	time.Sleep(2 * time.Second)
	fmt.Println("Cancelling all workers...")
	/* cancel parent context */
	
	wg.Wait()
	fmt.Println("All workers stopped")
	
	fmt.Println("\n=== Timeout Patterns ===")
	
	// TODO: Test different timeout scenarios
	testTimeouts()
	
	fmt.Println("\n=== Pipeline with Context ===")
	
	// TODO: Create pipeline context
	pipelineCtx, pipelineCancel := /* create context with 10 second timeout */
	defer pipelineCancel()
	
	// TODO: Run data processing pipeline
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	results := /* run pipeline with pipelineCtx and numbers */
	
	// Collect results
	fmt.Println("Pipeline results:")
	for result := range results {
		fmt.Printf("  Processed: %d\n", result)
	}
	
	fmt.Println("\n=== Context Best Practices ===")
	
	// TODO: Demonstrate context best practices
	demonstrateBestPractices()
}

// TODO: Implement work function that respects context
func doWork(ctx context.Context, taskName string, duration time.Duration) (string, error) {
	fmt.Printf("Starting %s (duration: %v)\n", taskName, duration)
	
	// TODO: Create timer for work simulation
	timer := time.NewTimer(duration)
	defer timer.Stop()
	
	select {
	case <-timer.C:
		/* return success result */
		return fmt.Sprintf("%s completed successfully", taskName), nil
	case <-ctx.Done():
		/* return context cancellation error */
		return "", fmt.Errorf("%s cancelled: %w", taskName, ctx.Err())
	}
}

// TODO: Implement request processing with context values
func processRequest(ctx context.Context) {
	// TODO: Extract values from context
	userID := /* get user ID from context */
	requestID := /* get request ID from context */
	
	fmt.Printf("Processing request %s for user %s\n", requestID, userID)
	
	// TODO: Pass context to other functions
	authenticateUser(ctx)
	fetchUserData(ctx)
	logActivity(ctx, "request processed")
}

func authenticateUser(ctx context.Context) {
	userID := /* get user ID from context, provide default "unknown" */
	fmt.Printf("  Authenticating user: %s\n", userID)
	
	// Simulate auth work
	time.Sleep(100 * time.Millisecond)
}

func fetchUserData(ctx context.Context) {
	userID := /* get user ID from context, provide default "unknown" */
	requestID := /* get request ID from context, provide default "unknown" */
	fmt.Printf("  Fetching data for user %s (request: %s)\n", userID, requestID)
	
	// Simulate data fetch
	time.Sleep(200 * time.Millisecond)
}

func logActivity(ctx context.Context, activity string) {
	userID := /* get user ID from context, provide default "unknown" */
	requestID := /* get request ID from context, provide default "unknown" */
	fmt.Printf("  LOG [%s][%s]: %s\n", requestID, userID, activity)
}

// TODO: Implement worker function
func worker(ctx context.Context, workerID int) {
	fmt.Printf("Worker %d starting\n", workerID)
	
	/* create ticker for 500ms intervals */
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

// TODO: Implement timeout testing
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
		
		/* create context with tc.timeout */
		ctx, cancel := context.WithTimeout(context.Background(), tc.timeout)
		
		/* call doWork with context */
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

// TODO: Implement processing pipeline with context
func processingPipeline(ctx context.Context, numbers []int) <-chan int {
	/* create output channel */
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

// TODO: Implement context best practices demonstration
func demonstrateBestPractices() {
	fmt.Println("Context Best Practices:")
	fmt.Println("1. Always pass context as first parameter")
	fmt.Println("2. Don't store contexts in structs")
	fmt.Println("3. Always call cancel() to avoid resource leaks")
	fmt.Println("4. Use context.WithValue() sparingly")
	fmt.Println("5. Check context.Done() in long-running operations")
	
	// TODO: Demonstrate deadline vs timeout
	fmt.Println("\n=== Deadline vs Timeout ===")
	
	// Using WithDeadline
	deadline := time.Now().Add(2 * time.Second)
	/* create context with deadline */
	deadlineCtx, deadlineCancel := context.WithDeadline(context.Background(), deadline)
	defer deadlineCancel()
	
	fmt.Printf("Using WithDeadline (deadline: %v)\n", deadline.Format("15:04:05"))
	_, err := doWork(deadlineCtx, "deadline task", 3*time.Second)
	if err != nil {
		fmt.Printf("Deadline task result: %v\n", err)
	}
	
	// Using WithTimeout
	/* create context with 2 second timeout */
	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer timeoutCancel()
	
	fmt.Printf("Using WithTimeout (timeout: 2s)\n")
	_, err = doWork(timeoutCtx, "timeout task", 3*time.Second)
	if err != nil {
		fmt.Printf("Timeout task result: %v\n", err)
	}
	
	// TODO: Demonstrate context composition
	fmt.Println("\n=== Context Composition ===")
	
	/* create base context with 5 second timeout */
	baseCtx, baseCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer baseCancel()
	
	/* add user value to base context */
	userCtx = context.WithValue(baseCtx, "userID", "comp-user-123")
	
	/* create child context with 2 second timeout */
	childCtx, childCancel := context.WithTimeout(userCtx, 2*time.Second)
	defer childCancel()
	
	fmt.Println("Child context inherits from parent but has shorter timeout")
	processCompositeRequest(childCtx)
}

// TODO: Implement composite request processing
func processCompositeRequest(ctx context.Context) {
	userID := /* get userID from context, default "unknown" */
	
	fmt.Printf("Processing composite request for user: %s\n", userID)
	
	// TODO: Check context deadline
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

// TODO: Implement HTTP-like server simulation with context
func simulateHTTPServer() {
	fmt.Println("\n=== HTTP Server Simulation ===")
	
	// TODO: Simulate handling multiple requests with context
	requests := []string{"req1", "req2", "req3"}
	
	var wg sync.WaitGroup
	for _, reqID := range requests {
		wg.Add(1)
		go func(requestID string) {
			defer wg.Done()
			
			// TODO: Create request context with timeout
			ctx, cancel := /* create context with 2 second timeout */
			defer cancel()
			
			// TODO: Add request ID to context
			requestCtx := /* add requestID to context */
			
			/* handle HTTP request with context */
		}(reqID)
	}
	
	wg.Wait()
	fmt.Println("All requests processed")
}

// TODO: Implement HTTP request handler
func handleHTTPRequest(ctx context.Context) {
	requestID := /* get requestID from context, default "unknown" */
	fmt.Printf("Handling HTTP request: %s\n", requestID)
	
	// TODO: Simulate request processing steps
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