// worker_pools.go - SOLUTION
// Learn how to implement worker pools for concurrent task processing

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job represents a unit of work
type Job struct {
	ID       int
	Data     string
	Priority int
}

// Result represents the output of processing a job
type Result struct {
	JobID    int
	Output   string
	Duration time.Duration
	WorkerID int
	Error    error
}

// Simple worker pool implementation
type WorkerPool struct {
	numWorkers  int
	jobQueue    chan Job
	resultQueue chan Result
	ctx         context.Context
	cancel      context.CancelFunc
	wg          sync.WaitGroup
}

// Create new worker pool
func NewWorkerPool(numWorkers, queueSize int) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	
	return &WorkerPool{
		numWorkers:  numWorkers,
		jobQueue:    make(chan Job, queueSize),
		resultQueue: make(chan Result, queueSize),
		ctx:         ctx,
		cancel:      cancel,
	}
}

// Start the worker pool
func (wp *WorkerPool) Start() {
	fmt.Printf("Starting worker pool with %d workers\n", wp.numWorkers)
	
	for i := 1; i <= wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

// Worker function
func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	fmt.Printf("Worker %d started\n", id)
	
	for {
		select {
		case job, ok := <-wp.jobQueue:
			if !ok {
				fmt.Printf("Worker %d: Job queue closed\n", id)
				return
			}
			
			// Process the job
			result := wp.processJob(job, id)
			
			// Send result
			select {
			case wp.resultQueue <- result:
				// Result sent successfully
			case <-wp.ctx.Done():
				fmt.Printf("Worker %d: Context cancelled while sending result\n", id)
				return
			}
			
		case <-wp.ctx.Done():
			fmt.Printf("Worker %d: Context cancelled\n", id)
			return
		}
	}
}

// Process a single job
func (wp *WorkerPool) processJob(job Job, workerID int) Result {
	start := time.Now()
	
	fmt.Printf("Worker %d: Processing job %d (data: %s, priority: %d)\n", 
		workerID, job.ID, job.Data, job.Priority)
	
	// Simulate work with random duration
	workDuration := time.Duration(rand.Intn(1000)+200) * time.Millisecond
	
	// Context-aware work simulation
	select {
	case <-time.After(workDuration):
		// Work completed successfully
		output := fmt.Sprintf("Processed '%s' with priority %d", job.Data, job.Priority)
		return Result{
			JobID:    job.ID,
			Output:   output,
			Duration: time.Since(start),
			WorkerID: workerID,
			Error:    nil,
		}
	case <-wp.ctx.Done():
		// Work was cancelled
		return Result{
			JobID:    job.ID,
			Output:   "",
			Duration: time.Since(start),
			WorkerID: workerID,
			Error:    wp.ctx.Err(),
		}
	}
}

// Submit job to the pool
func (wp *WorkerPool) SubmitJob(job Job) error {
	select {
	case wp.jobQueue <- job:
		return nil
	case <-wp.ctx.Done():
		return wp.ctx.Err()
	default:
		return fmt.Errorf("job queue is full")
	}
}

// Get result from the pool
func (wp *WorkerPool) GetResult() (Result, bool) {
	select {
	case result := <-wp.resultQueue:
		return result, true
	case <-wp.ctx.Done():
		return Result{}, false
	}
}

// Stop the worker pool
func (wp *WorkerPool) Stop() {
	fmt.Println("Stopping worker pool...")
	
	// Close job queue to signal workers to finish current jobs
	close(wp.jobQueue)
	
	// Wait for workers to finish
	wp.wg.Wait()
	
	// Cancel context and close result queue
	wp.cancel()
	close(wp.resultQueue)
	
	fmt.Println("Worker pool stopped")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("=== Basic Worker Pool ===")
	
	// Create and start basic worker pool
	pool := NewWorkerPool(3, 10)
	pool.Start()
	
	// Submit jobs
	jobs := []Job{
		{ID: 1, Data: "process image", Priority: 1},
		{ID: 2, Data: "send email", Priority: 2},
		{ID: 3, Data: "backup data", Priority: 1},
		{ID: 4, Data: "generate report", Priority: 3},
		{ID: 5, Data: "update database", Priority: 2},
	}
	
	// Submit all jobs
	fmt.Println("Submitting jobs...")
	for _, job := range jobs {
		err := pool.SubmitJob(job)
		if err != nil {
			fmt.Printf("Failed to submit job %d: %v\n", job.ID, err)
		} else {
			fmt.Printf("Submitted job %d\n", job.ID)
		}
	}
	
	// Collect results
	fmt.Println("\nCollecting results...")
	for i := 0; i < len(jobs); i++ {
		result, ok := pool.GetResult()
		if !ok {
			fmt.Println("No more results available")
			break
		}
		
		if result.Error != nil {
			fmt.Printf("Job %d failed: %v (took %v)\n", 
				result.JobID, result.Error, result.Duration)
		} else {
			fmt.Printf("Job %d completed by worker %d: %s (took %v)\n", 
				result.JobID, result.WorkerID, result.Output, result.Duration)
		}
	}
	
	// Stop the pool
	pool.Stop()
	
	fmt.Println("\n=== Worker Pool with Timeout ===")
	
	// Create worker pool with context timeout
	timeoutPool := NewWorkerPool(2, 5)
	
	// Override context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	timeoutPool.ctx = ctx
	timeoutPool.cancel = cancel
	
	timeoutPool.Start()
	
	// Submit jobs that will be interrupted by timeout
	timeoutJobs := []Job{
		{ID: 21, Data: "quick task", Priority: 1},
		{ID: 22, Data: "slow task", Priority: 1},
		{ID: 23, Data: "medium task", Priority: 1},
	}
	
	fmt.Println("Submitting jobs with timeout...")
	for _, job := range timeoutJobs {
		err := timeoutPool.SubmitJob(job)
		if err != nil {
			fmt.Printf("Failed to submit timeout job %d: %v\n", job.ID, err)
		} else {
			fmt.Printf("Submitted timeout job %d\n", job.ID)
		}
	}
	
	// Try to collect results before timeout
	fmt.Println("\nCollecting results with timeout...")
	resultCount := 0
	for {
		result, ok := timeoutPool.GetResult()
		if !ok {
			fmt.Println("Result collection stopped (timeout or pool closed)")
			break
		}
		
		resultCount++
		if result.Error != nil {
			fmt.Printf("Timeout job %d failed: %v\n", result.JobID, result.Error)
		} else {
			fmt.Printf("Timeout job %d completed: %s\n", result.JobID, result.Output)
		}
		
		if resultCount >= len(timeoutJobs) {
			break
		}
	}
	
	timeoutPool.Stop()
	
	fmt.Println("\nWorker pool examples completed")
}