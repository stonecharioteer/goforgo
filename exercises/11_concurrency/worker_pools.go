// worker_pools.go
// Learn how to implement worker pools for concurrent task processing

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// TODO: Job represents a unit of work
type Job struct {
	ID       int
	Data     string
	Priority int
}

// TODO: Result represents the output of processing a job
type Result struct {
	JobID    int
	Output   string
	Duration time.Duration
	WorkerID int
	Error    error
}

// TODO: Simple worker pool implementation
type WorkerPool struct {
	numWorkers int
	jobQueue   chan Job
	resultQueue chan Result
	ctx        context.Context
	cancel     context.CancelFunc
	wg         sync.WaitGroup
}

// TODO: Create new worker pool
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

// TODO: Start the worker pool
func (wp *WorkerPool) Start() {
	fmt.Printf("Starting worker pool with %d workers\n", wp.numWorkers)
	
	for i := 1; i <= wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

// TODO: Worker function
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
			
			// TODO: Process the job
			result := wp.processJob(job, id)
			
			// TODO: Send result
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

// TODO: Process a single job
func (wp *WorkerPool) processJob(job Job, workerID int) Result {
	start := time.Now()
	
	fmt.Printf("Worker %d: Processing job %d (data: %s, priority: %d)\n", 
		workerID, job.ID, job.Data, job.Priority)
	
	// TODO: Simulate work with random duration
	workDuration := time.Duration(rand.Intn(1000)+200) * time.Millisecond
	
	// TODO: Context-aware work simulation
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

// TODO: Submit job to the pool
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

// TODO: Get result from the pool
func (wp *WorkerPool) GetResult() (Result, bool) {
	select {
	case result := <-wp.resultQueue:
		return result, true
	case <-wp.ctx.Done():
		return Result{}, false
	}
}

// TODO: Stop the worker pool
func (wp *WorkerPool) Stop() {
	fmt.Println("Stopping worker pool...")
	
	// TODO: Close job queue to signal workers to finish current jobs
	close(wp.jobQueue)
	
	// TODO: Wait for workers to finish
	wp.wg.Wait()
	
	// TODO: Cancel context and close result queue
	wp.cancel()
	close(wp.resultQueue)
	
	fmt.Println("Worker pool stopped")
}

// TODO: Priority-based worker pool
type PriorityWorkerPool struct {
	highPriorityQueue chan Job
	lowPriorityQueue  chan Job
	resultQueue       chan Result
	numWorkers        int
	ctx               context.Context
	cancel            context.CancelFunc
	wg                sync.WaitGroup
}

// TODO: Create priority worker pool
func NewPriorityWorkerPool(numWorkers, queueSize int) *PriorityWorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	
	return &PriorityWorkerPool{
		highPriorityQueue: make(chan Job, queueSize),
		lowPriorityQueue:  make(chan Job, queueSize),
		resultQueue:       make(chan Result, queueSize*2),
		numWorkers:        numWorkers,
		ctx:               ctx,
		cancel:            cancel,
	}
}

// TODO: Start priority worker pool
func (pwp *PriorityWorkerPool) Start() {
	fmt.Printf("Starting priority worker pool with %d workers\n", pwp.numWorkers)
	
	for i := 1; i <= pwp.numWorkers; i++ {
		pwp.wg.Add(1)
		go pwp.priorityWorker(i)
	}
}

// TODO: Priority worker that prefers high priority jobs
func (pwp *PriorityWorkerPool) priorityWorker(id int) {
	defer pwp.wg.Done()
	fmt.Printf("Priority Worker %d started\n", id)
	
	for {
		select {
		// TODO: Check high priority jobs first
		case job := <-pwp.highPriorityQueue:
			result := pwp.processJob(job, id, "HIGH")
			select {
			case pwp.resultQueue <- result:
			case <-pwp.ctx.Done():
				return
			}
			
		// TODO: Check low priority jobs if no high priority jobs
		case job := <-pwp.lowPriorityQueue:
			result := pwp.processJob(job, id, "LOW")
			select {
			case pwp.resultQueue <- result:
			case <-pwp.ctx.Done():
				return
			}
			
		case <-pwp.ctx.Done():
			fmt.Printf("Priority Worker %d: Context cancelled\n", id)
			return
		}
	}
}

// TODO: Process job with priority info
func (pwp *PriorityWorkerPool) processJob(job Job, workerID int, priority string) Result {
	start := time.Now()
	
	fmt.Printf("Priority Worker %d: Processing %s priority job %d (data: %s)\n", 
		workerID, priority, job.ID, job.Data)
	
	// TODO: Different processing times based on priority
	var workDuration time.Duration
	if priority == "HIGH" {
		workDuration = time.Duration(rand.Intn(300)+100) * time.Millisecond
	} else {
		workDuration = time.Duration(rand.Intn(800)+400) * time.Millisecond
	}
	
	select {
	case <-time.After(workDuration):
		output := fmt.Sprintf("Processed %s priority job '%s'", priority, job.Data)
		return Result{
			JobID:    job.ID,
			Output:   output,
			Duration: time.Since(start),
			WorkerID: workerID,
			Error:    nil,
		}
	case <-pwp.ctx.Done():
		return Result{
			JobID:    job.ID,
			Output:   "",
			Duration: time.Since(start),
			WorkerID: workerID,
			Error:    pwp.ctx.Err(),
		}
	}
}

// TODO: Submit job based on priority
func (pwp *PriorityWorkerPool) SubmitJob(job Job) error {
	var targetQueue chan Job
	
	if job.Priority >= 5 {
		targetQueue = pwp.highPriorityQueue
	} else {
		targetQueue = pwp.lowPriorityQueue
	}
	
	select {
	case targetQueue <- job:
		return nil
	case <-pwp.ctx.Done():
		return pwp.ctx.Err()
	default:
		return fmt.Errorf("job queue is full")
	}
}

// TODO: Get result from priority pool
func (pwp *PriorityWorkerPool) GetResult() (Result, bool) {
	select {
	case result := <-pwp.resultQueue:
		return result, true
	case <-pwp.ctx.Done():
		return Result{}, false
	}
}

// TODO: Stop priority worker pool
func (pwp *PriorityWorkerPool) Stop() {
	fmt.Println("Stopping priority worker pool...")
	close(pwp.highPriorityQueue)
	close(pwp.lowPriorityQueue)
	pwp.wg.Wait()
	pwp.cancel()
	close(pwp.resultQueue)
	fmt.Println("Priority worker pool stopped")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("=== Basic Worker Pool ===")
	
	// TODO: Create and start basic worker pool
	pool := NewWorkerPool(3, 10)
	pool.Start()
	
	// TODO: Submit jobs
	jobs := []Job{
		{ID: 1, Data: "process image", Priority: 1},
		{ID: 2, Data: "send email", Priority: 2},
		{ID: 3, Data: "backup data", Priority: 1},
		{ID: 4, Data: "generate report", Priority: 3},
		{ID: 5, Data: "update database", Priority: 2},
	}
	
	// TODO: Submit all jobs
	fmt.Println("Submitting jobs...")
	for _, job := range jobs {
		err := pool.SubmitJob(job)
		if err != nil {
			fmt.Printf("Failed to submit job %d: %v\n", job.ID, err)
		} else {
			fmt.Printf("Submitted job %d\n", job.ID)
		}
	}
	
	// TODO: Collect results
	fmt.Println("\\nCollecting results...")
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
	
	// TODO: Stop the pool
	pool.Stop()
	
	fmt.Println("\\n=== Priority Worker Pool ===")
	
	// TODO: Create and start priority worker pool
	priorityPool := NewPriorityWorkerPool(2, 5)
	priorityPool.Start()
	
	// TODO: Submit jobs with different priorities
	priorityJobs := []Job{
		{ID: 11, Data: "critical alert", Priority: 9},
		{ID: 12, Data: "routine cleanup", Priority: 2},
		{ID: 13, Data: "urgent notification", Priority: 8},
		{ID: 14, Data: "background sync", Priority: 1},
		{ID: 15, Data: "important update", Priority: 7},
		{ID: 16, Data: "low priority task", Priority: 3},
	}
	
	fmt.Println("Submitting priority jobs...")
	for _, job := range priorityJobs {
		err := priorityPool.SubmitJob(job)
		if err != nil {
			fmt.Printf("Failed to submit priority job %d: %v\n", job.ID, err)
		} else {
			fmt.Printf("Submitted priority job %d (priority: %d)\n", job.ID, job.Priority)
		}
		
		// TODO: Stagger job submission to see priority handling
		time.Sleep(50 * time.Millisecond)
	}
	
	// TODO: Collect priority results
	fmt.Println("\\nCollecting priority results...")
	for i := 0; i < len(priorityJobs); i++ {
		result, ok := priorityPool.GetResult()
		if !ok {
			fmt.Println("No more priority results available")
			break
		}
		
		if result.Error != nil {
			fmt.Printf("Priority job %d failed: %v\n", result.JobID, result.Error)
		} else {
			fmt.Printf("Priority job %d completed by worker %d: %s (took %v)\n", 
				result.JobID, result.WorkerID, result.Output, result.Duration)
		}
	}
	
	// TODO: Stop priority pool
	priorityPool.Stop()
	
	fmt.Println("\\n=== Worker Pool with Timeout ===")
	
	// TODO: Create worker pool with context timeout
	timeoutPool := NewWorkerPool(2, 5)
	
	// TODO: Override context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	timeoutPool.ctx = ctx
	timeoutPool.cancel = cancel
	
	timeoutPool.Start()
	
	// TODO: Submit jobs that will be interrupted by timeout
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
	
	// TODO: Try to collect results before timeout
	fmt.Println("\\nCollecting results with timeout...")
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
	
	fmt.Println("\\nWorker pool examples completed")
}