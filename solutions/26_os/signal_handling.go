// signal_handling.go - SOLUTION
// Learn signal handling for graceful shutdowns and process control

package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

func main() {
	fmt.Println("=== Signal Handling ===")
	
	fmt.Println("\n=== Basic Signal Handling ===")
	
	// Create channel to receive OS signals
	sigChan := make(chan os.Signal, 1)
	
	// Register for specific signals
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	fmt.Println("Basic signal handler started. Press Ctrl+C to test...")
	
	// Start a background goroutine
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Working... %d/10\n", i+1)
			time.Sleep(1 * time.Second)
		}
	}()
	
	// Wait for signal (with timeout for demo)
	select {
	case sig := <-sigChan:
		fmt.Printf("\nReceived signal: %v\n", sig)
		fmt.Println("Performing graceful shutdown...")
	case <-time.After(3 * time.Second):
		fmt.Println("\nNo signal received, continuing with demo...")
	}
	
	fmt.Println("\n=== Graceful Shutdown Pattern ===")
	
	// Implement graceful shutdown with context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	// Start multiple workers
	var wg sync.WaitGroup
	
	// Worker 1 - Database operations
	wg.Add(1)
	go func() {
		defer wg.Done()
		databaseWorker(ctx, "DatabaseWorker")
	}()
	
	// Worker 2 - HTTP server simulation
	wg.Add(1)
	go func() {
		defer wg.Done()
		httpServerWorker(ctx, "HTTPServer")
	}()
	
	// Worker 3 - Background task processor
	wg.Add(1)
	go func() {
		defer wg.Done()
		backgroundProcessor(ctx, "BackgroundProcessor")
	}()
	
	fmt.Println("All workers started. Press Ctrl+C for graceful shutdown...")
	
	// Wait for shutdown signal (with timeout for demo)
	shutdownSig := make(chan os.Signal, 1)
	signal.Notify(shutdownSig, syscall.SIGINT, syscall.SIGTERM)
	
	select {
	case <-shutdownSig:
		fmt.Println("\nShutdown signal received. Gracefully stopping workers...")
	case <-time.After(2 * time.Second):
		fmt.Println("\nDemo timeout reached, simulating shutdown...")
	}
	
	// Cancel context to signal workers to stop
	cancel()
	
	// Wait for all workers to finish with timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	
	select {
	case <-done:
		fmt.Println("All workers stopped gracefully")
	case <-time.After(5 * time.Second):
		fmt.Println("Timeout waiting for workers to stop")
	}
	
	fmt.Println("\n=== Signal Types and Handling ===")
	
	// Handle different signal types
	signalTypes := map[os.Signal]string{
		syscall.SIGINT:  "Interrupt from keyboard (Ctrl+C)",
		syscall.SIGTERM: "Termination request",
		syscall.SIGHUP:  "Hangup (traditionally reload config)",
		syscall.SIGUSR1: "User-defined signal 1",
		syscall.SIGUSR2: "User-defined signal 2",
		syscall.SIGQUIT: "Quit from keyboard (Ctrl+\\)",
	}
	
	fmt.Println("Handling different signal types:")
	for sig, description := range signalTypes {
		fmt.Printf("  %v: %s\n", sig, description)
	}
	
	// Create comprehensive signal handler
	allSignals := make(chan os.Signal, 1)
	signal.Notify(allSignals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, 
		syscall.SIGUSR1, syscall.SIGUSR2)
	
	fmt.Println("\nComprehensive signal handler active...")
	
	// Handle signals in a loop
	go func() {
		for {
			select {
			case sig := <-allSignals:
				switch sig {
				case syscall.SIGINT, syscall.SIGTERM:
					fmt.Printf("Received shutdown signal: %v\n", sig)
				case syscall.SIGHUP:
					fmt.Println("Received SIGHUP - reloading configuration")
				case syscall.SIGUSR1:
					fmt.Println("Received SIGUSR1 - custom action 1")
				case syscall.SIGUSR2:
					fmt.Println("Received SIGUSR2 - custom action 2")
				default:
					fmt.Printf("Received unknown signal: %v\n", sig)
				}
			}
		}
	}()
	
	fmt.Println("\n=== Process Management ===")
	
	// Demonstrate sending signals to other processes
	fmt.Println("Process management examples:")
	
	// Get current process info
	pid := os.Getpid()
	ppid := os.Getppid()
	
	fmt.Printf("Current Process ID: %d\n", pid)
	fmt.Printf("Parent Process ID: %d\n", ppid)
	
	// Start a child process for signal testing
	cmd := exec.Command("sleep", "10")
	if err := cmd.Start(); err == nil {
		childProcess := cmd.Process
		defer func() {
			if childProcess != nil {
				childProcess.Kill()
				childProcess.Wait()
			}
		}()
		
		fmt.Printf("Started child process with PID: %d\n", childProcess.Pid)
		
		// Send signals to child process
		time.Sleep(1 * time.Second)
		fmt.Println("Sending SIGUSR1 to child process...")
		childProcess.Signal(syscall.SIGUSR1)
		
		time.Sleep(1 * time.Second)
		fmt.Println("Sending SIGTERM to child process...")
		childProcess.Signal(syscall.SIGTERM)
		
		// Wait for child process to exit
		cmd.Wait()
		fmt.Println("Child process terminated")
	}
	
	fmt.Println("\n=== Signal Safety and Best Practices ===")
	
	// Demonstrate signal-safe operations
	fmt.Println("Signal safety examples:")
	
	// Use atomic operations for signal handlers
	var counter int64
	
	// Signal handler that modifies shared state safely
	sigHandler := make(chan os.Signal, 1)
	signal.Notify(sigHandler, syscall.SIGUSR1)
	
	go func() {
		for {
			select {
			case <-sigHandler:
				// Atomically increment counter
				newValue := atomic.AddInt64(&counter, 1)
				fmt.Printf("Signal received, counter: %d\n", newValue)
			}
		}
	}()
	
	// Send some test signals
	currentProcess, _ := os.FindProcess(os.Getpid())
	for i := 0; i < 3; i++ {
		time.Sleep(200 * time.Millisecond)
		currentProcess.Signal(syscall.SIGUSR1)
	}
	
	time.Sleep(500 * time.Millisecond) // Allow signal processing
	
	fmt.Println("\n=== Advanced Signal Patterns ===")
	
	// Implement signal-based configuration reload
	type Config struct {
		LogLevel string
		MaxConns int
	}
	
	var (
		currentConfig = &Config{LogLevel: "info", MaxConns: 100}
		configMutex   sync.RWMutex
	)
	
	// Configuration reload handler
	reloadSignal := make(chan os.Signal, 1)
	signal.Notify(reloadSignal, syscall.SIGHUP)
	
	go func() {
		for {
			select {
			case <-reloadSignal:
				fmt.Println("Reloading configuration...")
				
				// Reload configuration (simulated)
				configMutex.Lock()
				currentConfig.LogLevel = "debug"
				currentConfig.MaxConns = 200
				configMutex.Unlock()
				
				fmt.Println("Configuration reloaded successfully")
			}
		}
	}()
	
	// Implement signal-based log rotation
	rotateSignal := make(chan os.Signal, 1)
	signal.Notify(rotateSignal, syscall.SIGUSR2)
	
	go func() {
		for {
			select {
			case <-rotateSignal:
				fmt.Println("Rotating log files...")
				// Simulate log rotation
				time.Sleep(100 * time.Millisecond)
				fmt.Println("Log files rotated successfully")
			}
		}
	}()
	
	fmt.Println("Advanced signal handlers active:")
	fmt.Println("  - Send SIGHUP to reload configuration")
	fmt.Println("  - Send SIGUSR2 to rotate log files")
	
	// Simulate some signals
	time.Sleep(500 * time.Millisecond)
	currentProcess.Signal(syscall.SIGHUP)
	
	time.Sleep(500 * time.Millisecond)
	currentProcess.Signal(syscall.SIGUSR2)
	
	time.Sleep(500 * time.Millisecond)
	
	fmt.Println("\n=== Signal Handling in Services ===")
	
	// Service-like signal handling
	type Service struct {
		name     string
		running  bool
		shutdown chan struct{}
		done     chan struct{}
		mu       sync.Mutex
	}
	
	// Implement service methods
	func NewService(name string) *Service {
		return &Service{
			name:     name,
			shutdown: make(chan struct{}),
			done:     make(chan struct{}),
		}
	}
	
	func (s *Service) Start() error {
		s.mu.Lock()
		defer s.mu.Unlock()
		
		if s.running {
			return fmt.Errorf("service %s is already running", s.name)
		}
		
		s.running = true
		go s.run()
		fmt.Printf("Service %s started\n", s.name)
		return nil
	}
	
	func (s *Service) Stop() error {
		s.mu.Lock()
		defer s.mu.Unlock()
		
		if !s.running {
			return fmt.Errorf("service %s is not running", s.name)
		}
		
		close(s.shutdown)
		<-s.done
		s.running = false
		fmt.Printf("Service %s stopped\n", s.name)
		return nil
	}
	
	func (s *Service) run() {
		defer close(s.done)
		
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		
		for {
			select {
			case <-s.shutdown:
				return
			case <-ticker.C:
				// Service work simulation
			}
		}
	}
	
	// Create and start services
	services := []*Service{
		NewService("WebServer"),
		NewService("DatabasePool"),
		NewService("MessageQueue"),
	}
	
	fmt.Printf("Starting %d services...\n", len(services))
	for _, service := range services {
		service.Start()
	}
	
	// Set up service shutdown handler
	serviceShutdown := make(chan os.Signal, 1)
	signal.Notify(serviceShutdown, syscall.SIGINT, syscall.SIGTERM)
	
	// Wait for shutdown signal (with timeout for demo)
	select {
	case <-serviceShutdown:
		fmt.Println("Received shutdown signal")
	case <-time.After(2 * time.Second):
		fmt.Println("Demo timeout, simulating shutdown")
	}
	
	fmt.Println("Shutting down all services...")
	for _, service := range services {
		service.Stop()
	}
	
	fmt.Println("\n=== Best Practices Summary ===")
	fmt.Println("Signal handling best practices:")
	fmt.Println("✅ Always use signal.Notify() to register for signals")
	fmt.Println("✅ Use buffered channels to avoid blocking signal delivery")
	fmt.Println("✅ Handle SIGINT and SIGTERM for graceful shutdown")
	fmt.Println("✅ Use context.Context for coordinated cancellation")
	fmt.Println("✅ Implement timeouts for shutdown operations")
	fmt.Println("✅ Use atomic operations in signal handlers")
	fmt.Println("✅ Avoid complex operations in signal handlers")
	fmt.Println("✅ Test signal handling thoroughly")
	fmt.Println("✅ Document custom signal meanings")
	fmt.Println("✅ Use WaitGroup for coordinating goroutine shutdown")
}

// Worker functions

func databaseWorker(ctx context.Context, name string) {
	fmt.Printf("%s started\n", name)
	defer fmt.Printf("%s stopped\n", name)
	
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s received shutdown signal\n", name)
			// Simulate cleanup
			time.Sleep(200 * time.Millisecond)
			return
		case <-ticker.C:
			// Simulate database operations
		}
	}
}

func httpServerWorker(ctx context.Context, name string) {
	fmt.Printf("%s started\n", name)
	defer fmt.Printf("%s stopped\n", name)
	
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s received shutdown signal\n", name)
			// Simulate server shutdown
			time.Sleep(300 * time.Millisecond)
			return
		case <-time.After(600 * time.Millisecond):
			// Simulate HTTP request processing
		}
	}
}

func backgroundProcessor(ctx context.Context, name string) {
	fmt.Printf("%s started\n", name)
	defer fmt.Printf("%s stopped\n", name)
	
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s received shutdown signal\n", name)
			// Simulate task completion
			time.Sleep(100 * time.Millisecond)
			return
		case <-time.After(800 * time.Millisecond):
			// Simulate background task processing
		}
	}
}