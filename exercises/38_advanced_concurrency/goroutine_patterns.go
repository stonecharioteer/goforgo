// goroutine_patterns.go
// Learn advanced goroutine communication patterns: Fan-in/Fan-out, Pipelines, Or-channels, Graceful shutdown
//
// This exercise covers advanced concurrency patterns that are commonly used in
// production Go systems for building scalable, robust, and maintainable concurrent programs.
// These patterns help organize complex concurrent workflows and handle cancellation gracefully.
//
// I AM NOT DONE YET!

package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// TODO: Fan-out pattern - distribute work to multiple workers
func fanOut(input <-chan int, workers int) []<-chan int {
	// TODO: Create slice of output channels
	// outputs := make([]<-chan int, workers)
	
	for i := 0; i < workers; i++ {
		// TODO: Create output channel for each worker
		// output := make(chan int)
		// outputs[i] = output
		
		// TODO: Start worker goroutine
		go func(out chan<- int) {
			// TODO: Close channel when done
			// defer close(out)
			
			// TODO: Read from input and process
			// for n := range input {
			//     // Simulate processing work
			//     processedValue := n * n
			//     time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			//     
			//     fmt.Printf("Worker processing %d -> %d\n", n, processedValue)
			//     out <- processedValue
			// }
		}(output)
	}
	
	// return outputs
}

// TODO: Fan-in pattern - merge multiple channels into one
func fanIn(inputs ...<-chan int) <-chan int {
	// TODO: Create output channel
	// output := make(chan int)
	
	var wg sync.WaitGroup
	
	// TODO: Start goroutine for each input channel
	for _, input := range inputs {
		wg.Add(1)
		go func(input <-chan int) {
			defer wg.Done()
			// TODO: Forward all values from input to output
			// for value := range input {
			//     output <- value
			// }
		}(input)
	}
	
	// TODO: Close output when all inputs are done
	go func() {
		wg.Wait()
		close(output)
	}()
	
	// return output
}

// TODO: Pipeline stage function type
type StageFunc func(<-chan interface{}) <-chan interface{}

// TODO: Create a pipeline with multiple stages
func pipeline(ctx context.Context, stages ...StageFunc) (chan<- interface{}, <-chan interface{}) {
	// TODO: Create input channel
	// input := make(chan interface{})
	
	// TODO: Chain stages together
	// current := input
	// for _, stage := range stages {
	//     current = stage(current)
	// }
	
	// return input, current
}

// TODO: Pipeline stage 1: Number generator
func numberGenerator(ctx context.Context) <-chan interface{} {
	// TODO: Create output channel
	// output := make(chan interface{})
	
	go func() {
		// TODO: Close channel when done
		// defer close(output)
		
		for i := 1; i <= 20; i++ {
			select {
			// TODO: Send number to output
			// case output <- i:
			//     fmt.Printf("Generated: %d\n", i)
			//     time.Sleep(50 * time.Millisecond)
			case <-ctx.Done():
				fmt.Println("Number generator cancelled")
				return
			}
		}
	}()
	
	// return output
}

// TODO: Pipeline stage 2: Square numbers
func squareStage(input <-chan interface{}) <-chan interface{} {
	// TODO: Create output channel
	// output := make(chan interface{})
	
	go func() {
		// TODO: Close channel when done
		// defer close(output)
		
		// TODO: Process each input
		// for value := range input {
		//     if num, ok := value.(int); ok {
		//         squared := num * num
		//         fmt.Printf("Squared: %d -> %d\n", num, squared)
		//         output <- squared
		//         time.Sleep(30 * time.Millisecond)
		//     }
		// }
	}()
	
	// return output
}

// TODO: Pipeline stage 3: Filter even numbers
func filterEvenStage(input <-chan interface{}) <-chan interface{} {
	// TODO: Create output channel
	// output := make(chan interface{})
	
	go func() {
		// TODO: Close channel when done
		// defer close(output)
		
		// TODO: Filter even numbers
		// for value := range input {
		//     if num, ok := value.(int); ok && num%2 == 0 {
		//         fmt.Printf("Filtered (even): %d\n", num)
		//         output <- num
		//     }
		// }
	}()
	
	// return output
}

// TODO: Or-channel pattern - return first available channel
func orChannel(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	
	// TODO: Create output channel
	// output := make(chan interface{})
	
	go func() {
		defer close(output)
		
		switch len(channels) {
		case 2:
			// TODO: Handle two channels
			select {
			// case val := <-channels[0]:
			//     output <- val
			// case val := <-channels[1]:
			//     output <- val
			}
		default:
			// TODO: Handle multiple channels recursively
			// select {
			// case val := <-channels[0]:
			//     output <- val
			// case val := <-orChannel(channels[1:]...):
			//     output <- val
			// }
		}
	}()
	
	// return output
}

// TODO: Timeout channel helper
func after(d time.Duration) <-chan interface{} {
	// TODO: Create channel that closes after duration
	// ch := make(chan interface{})
	go func() {
		defer close(ch)
		time.Sleep(d)
	}()
	// return ch
}

// TODO: Done channel helper  
func done() <-chan interface{} {
	// TODO: Create channel that closes immediately
	// ch := make(chan interface{})
	close(ch)
	// return ch
}

// TODO: Service interface for graceful shutdown
type Service interface {
	Start(ctx context.Context) error
	Stop() error
	Name() string
}

// TODO: Example service implementation
type WorkerService struct {
	name    string
	stopped chan struct{}
	mu      sync.RWMutex
	running bool
}

func NewWorkerService(name string) *WorkerService {
	return &WorkerService{
		name:    name,
		stopped: make(chan struct{}),
	}
}

func (s *WorkerService) Name() string {
	return s.name
}

func (s *WorkerService) Start(ctx context.Context) error {
	s.mu.Lock()
	if s.running {
		s.mu.Unlock()
		return fmt.Errorf("service %s already running", s.name)
	}
	s.running = true
	s.mu.Unlock()
	
	fmt.Printf("Service %s starting...\n", s.name)
	
	// TODO: Main service loop
	go func() {
		defer func() {
			close(s.stopped)
			s.mu.Lock()
			s.running = false
			s.mu.Unlock()
		}()
		
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		
		for {
			select {
			case <-ticker.C:
				// TODO: Simulate work
				fmt.Printf("Service %s working...\n", s.name)
			case <-ctx.Done():
				fmt.Printf("Service %s received shutdown signal\n", s.name)
				return
			}
		}
	}()
	
	return nil
}

func (s *WorkerService) Stop() error {
	fmt.Printf("Service %s stopping...\n", s.name)
	
	// TODO: Wait for service to stop
	<-s.stopped
	
	fmt.Printf("Service %s stopped\n", s.name)
	return nil
}

// TODO: Service manager for graceful shutdown
type ServiceManager struct {
	services []Service
	ctx      context.Context
	cancel   context.CancelFunc
}

func NewServiceManager() *ServiceManager {
	ctx, cancel := context.WithCancel(context.Background())
	return &ServiceManager{
		services: make([]Service, 0),
		ctx:      ctx,
		cancel:   cancel,
	}
}

func (sm *ServiceManager) AddService(service Service) {
	sm.services = append(sm.services, service)
}

func (sm *ServiceManager) Start() error {
	fmt.Println("Starting all services...")
	
	// TODO: Start all services
	for _, service := range sm.services {
		if err := service.Start(sm.ctx); err != nil {
			return fmt.Errorf("failed to start service %s: %v", service.Name(), err)
		}
	}
	
	fmt.Printf("Started %d services\n", len(sm.services))
	return nil
}

func (sm *ServiceManager) Stop() error {
	fmt.Println("Stopping all services...")
	
	// TODO: Cancel context to signal all services
	sm.cancel()
	
	// TODO: Stop services in reverse order
	for i := len(sm.services) - 1; i >= 0; i-- {
		service := sm.services[i]
		if err := service.Stop(); err != nil {
			fmt.Printf("Error stopping service %s: %v\n", service.Name(), err)
		}
	}
	
	fmt.Println("All services stopped")
	return nil
}

// TODO: Demonstrate rate limiting pattern
func rateLimiter(requests <-chan string, rate time.Duration) <-chan string {
	// TODO: Create output channel
	// output := make(chan string)
	
	go func() {
		defer close(output)
		ticker := time.NewTicker(rate)
		defer ticker.Stop()
		
		for req := range requests {
			// TODO: Wait for tick before processing
			// <-ticker.C
			// fmt.Printf("Processing request: %s\n", req)
			// output <- fmt.Sprintf("Processed: %s", req)
		}
	}()
	
	// return output
}

// TODO: Demonstrate circuit breaker pattern with channels
type CircuitBreaker struct {
	maxFailures int
	timeout     time.Duration
	failures    int
	lastFailure time.Time
	state       string // "closed", "open", "half-open"
	mu          sync.RWMutex
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures: maxFailures,
		timeout:     timeout,
		state:       "closed",
	}
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	
	// TODO: Check circuit breaker state
	switch cb.state {
	case "open":
		// TODO: Check if timeout has passed
		// if time.Since(cb.lastFailure) < cb.timeout {
		//     return fmt.Errorf("circuit breaker is open")
		// }
		// cb.state = "half-open"
	case "half-open":
		// Allow one request to test the service
	case "closed":
		// Normal operation
	}
	
	// TODO: Execute function
	// err := fn()
	// if err != nil {
	//     cb.failures++
	//     cb.lastFailure = time.Now()
	//     
	//     if cb.failures >= cb.maxFailures {
	//         cb.state = "open"
	//     }
	//     return err
	// }
	
	// TODO: Reset on success
	// cb.failures = 0
	// cb.state = "closed"
	// return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("=== Advanced Goroutine Patterns ===")
	
	fmt.Println("\n=== Fan-Out/Fan-In Pattern ===")
	
	// TODO: Create input channel and send data
	input := make(chan int, 5)
	go func() {
		defer close(input)
		for i := 1; i <= 10; i++ {
			input <- i
		}
	}()
	
	// TODO: Fan-out to multiple workers
	workers := fanOut(input, 3)
	
	// TODO: Fan-in to single output
	output := fanIn(workers...)
	
	// TODO: Collect results
	fmt.Println("Results from fan-out/fan-in:")
	for result := range output {
		fmt.Printf("Final result: %d\n", result)
	}
	
	fmt.Println("\n=== Pipeline Pattern ===")
	
	// TODO: Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// TODO: Create pipeline stages
	numbers := numberGenerator(ctx)
	squared := squareStage(numbers)
	filtered := filterEvenStage(squared)
	
	// TODO: Process pipeline results
	fmt.Println("Pipeline results (even squares):")
	for result := range filtered {
		if num, ok := result.(int); ok {
			fmt.Printf("Pipeline output: %d\n", num)
		}
	}
	
	fmt.Println("\n=== Or-Channel Pattern ===")
	
	// TODO: Create channels with different delays
	fast := make(chan interface{}, 1)
	slow := make(chan interface{}, 1)
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		fast <- "fast result"
	}()
	
	go func() {
		time.Sleep(500 * time.Millisecond)
		slow <- "slow result"
	}()
	
	// TODO: Use or-channel to get first result
	first := orChannel(fast, slow, after(1*time.Second))
	
	select {
	case result := <-first:
		fmt.Printf("First result: %v\n", result)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout waiting for result")
	}
	
	fmt.Println("\n=== Graceful Shutdown Pattern ===")
	
	// TODO: Create service manager
	manager := NewServiceManager()
	manager.AddService(NewWorkerService("DatabaseService"))
	manager.AddService(NewWorkerService("CacheService"))
	manager.AddService(NewWorkerService("APIService"))
	
	// TODO: Start all services
	if err := manager.Start(); err != nil {
		fmt.Printf("Failed to start services: %v\n", err)
		return
	}
	
	// TODO: Let services run
	time.Sleep(2 * time.Second)
	
	// TODO: Graceful shutdown
	if err := manager.Stop(); err != nil {
		fmt.Printf("Failed to stop services: %v\n", err)
	}
	
	fmt.Println("\n=== Rate Limiter Pattern ===")
	
	// TODO: Create request channel
	requests := make(chan string, 10)
	go func() {
		defer close(requests)
		for i := 1; i <= 5; i++ {
			requests <- fmt.Sprintf("Request-%d", i)
		}
	}()
	
	// TODO: Apply rate limiting (1 request per 200ms)
	limited := rateLimiter(requests, 200*time.Millisecond)
	
	// TODO: Process rate-limited requests
	fmt.Println("Rate-limited requests:")
	for response := range limited {
		fmt.Println(response)
	}
	
	fmt.Println("\n=== Circuit Breaker Pattern ===")
	
	// TODO: Create circuit breaker
	cb := NewCircuitBreaker(3, 2*time.Second)
	
	// TODO: Simulate service calls with some failures
	for i := 1; i <= 8; i++ {
		err := cb.Execute(func() error {
			// Simulate service call with 40% failure rate
			if rand.Float32() < 0.4 {
				return fmt.Errorf("service unavailable")
			}
			fmt.Printf("Service call %d succeeded\n", i)
			return nil
		})
		
		if err != nil {
			fmt.Printf("Service call %d failed: %v\n", i, err)
		}
		
		time.Sleep(100 * time.Millisecond)
	}
	
	fmt.Printf("\nRuntime info: %d goroutines\n", runtime.NumGoroutine())
	fmt.Println("Advanced goroutine patterns completed!")
}