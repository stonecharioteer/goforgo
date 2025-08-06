// goroutine_debugging.go
// Learn goroutine debugging, performance analysis, race detection, and leak prevention
//
// This exercise covers essential debugging and performance analysis techniques for
// concurrent Go programs. These skills are crucial for building robust, production-ready
// concurrent systems and identifying performance bottlenecks and correctness issues.
//
// I AM NOT DONE YET!

package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
	_ "net/http/pprof" // Import for side effects
	"net/http"
)

// TODO: Example of a race condition (intentional for demonstration)
type UnsafeCounter struct {
	value int64
}

func (c *UnsafeCounter) Increment() {
	// TODO: This creates a race condition - fix with atomic operations
	// c.value++ // Race condition!
	
	// TODO: Use atomic operations to fix race condition
	// atomic.AddInt64(&c.value, 1)
}

func (c *UnsafeCounter) Get() int64 {
	// TODO: This also creates a race condition
	// return c.value // Race condition!
	
	// TODO: Use atomic load to fix race condition
	// return atomic.LoadInt64(&c.value)
}

// TODO: Safe counter using atomic operations
type SafeCounter struct {
	value int64
}

func (c *SafeCounter) Increment() {
	// TODO: Use atomic increment
	// atomic.AddInt64(&c.value, 1)
}

func (c *SafeCounter) Get() int64 {
	// TODO: Use atomic load
	// return atomic.LoadInt64(&c.value)
}

// TODO: Goroutine leak example - channels that are never closed
func leakyFunction() <-chan int {
	// TODO: This creates a goroutine leak
	ch := make(chan int)
	
	go func() {
		// TODO: This goroutine will never exit!
		for i := 0; i < 1000000; i++ {
			// This select will block forever because nothing reads from ch
			select {
			case ch <- i:
			// TODO: Add context or timeout to prevent leak
			// case <-ctx.Done():
			//     return
			// case <-time.After(1 * time.Second):
			//     return
			}
		}
		// TODO: Channel is never closed, causing potential leak
		// close(ch)
	}()
	
	return ch
}

// TODO: Fixed version that prevents leaks
func nonLeakyFunction(ctx context.Context) <-chan int {
	ch := make(chan int)
	
	go func() {
		// TODO: Properly handle context cancellation
		defer close(ch)
		
		for i := 0; i < 1000000; i++ {
			select {
			case ch <- i:
			case <-ctx.Done():
				fmt.Println("Function cancelled, cleaning up goroutine")
				return
			}
		}
	}()
	
	return ch
}

// TODO: Memory ordering and happens-before relationships
var (
	data  int
	ready bool
)

func writer() {
	// TODO: This creates a race condition due to memory reordering
	// data = 42
	// ready = true
	
	// TODO: Use atomic operations to ensure proper ordering
	// atomic.StoreInt32((*int32)(&data), 42)
	// atomic.StoreBool((*bool)(&ready), true)
}

func reader() int {
	// TODO: This may see ready=true but data=0 due to reordering
	// if ready {
	//     return data
	// }
	
	// TODO: Use atomic operations for consistent reads
	// if atomic.LoadBool(&ready) {
	//     return int(atomic.LoadInt32((*int32)(&data)))
	// }
	return 0
}

// TODO: Benchmark for comparing concurrent vs sequential performance
func sequentialWork(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		// TODO: Simulate CPU-intensive work
		sum += i * i
	}
	return sum
}

func concurrentWork(n int, workers int) int {
	// TODO: Split work across goroutines
	if workers <= 0 {
		workers = runtime.NumCPU()
	}
	
	workPerWorker := n / workers
	results := make(chan int, workers)
	
	// TODO: Start worker goroutines
	for i := 0; i < workers; i++ {
		go func(start, end int) {
			sum := 0
			for j := start; j < end; j++ {
				sum += j * j
			}
			// TODO: Send result to channel
			// results <- sum
		}(i*workPerWorker, (i+1)*workPerWorker)
	}
	
	// TODO: Collect results
	totalSum := 0
	for i := 0; i < workers; i++ {
		totalSum += <-results
	}
	
	return totalSum
}

// TODO: Goroutine pool to prevent unlimited goroutine creation
type GoroutinePool struct {
	workers int
	workCh  chan func()
	quit    chan struct{}
	wg      sync.WaitGroup
}

func NewGoroutinePool(workers int) *GoroutinePool {
	pool := &GoroutinePool{
		workers: workers,
		workCh:  make(chan func(), workers*2), // Buffered to prevent blocking
		quit:    make(chan struct{}),
	}
	
	// TODO: Start worker goroutines
	for i := 0; i < workers; i++ {
		pool.wg.Add(1)
		go func() {
			// TODO: Worker loop
			defer pool.wg.Done()
			for {
				select {
				case work := <-pool.workCh:
					// TODO: Execute work
					// work()
				case <-pool.quit:
					return
				}
			}
		}()
	}
	
	return pool
}

func (p *GoroutinePool) Submit(work func()) {
	// TODO: Submit work to pool
	select {
	case p.workCh <- work:
	// Work submitted successfully
	case <-p.quit:
		// Pool is shutting down
		return
	}
}

func (p *GoroutinePool) Stop() {
	close(p.quit)
	p.wg.Wait()
}

// TODO: Monitor goroutine count and memory usage
func monitorResources(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	
	var lastGC debug.GCStats
	debug.ReadGCStats(&lastGC)
	
	for {
		select {
		case <-ticker.C:
			// TODO: Get runtime statistics
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			
			var gc debug.GCStats
			debug.ReadGCStats(&gc)
			
			// TODO: Print monitoring information
			fmt.Printf("Goroutines: %d, Memory: %d KB, GC Cycles: %d\n",
				runtime.NumGoroutine(),
				m.Alloc/1024,
				gc.NumGC-lastGC.NumGC)
			
			lastGC = gc
			
		case <-ctx.Done():
			return
		}
	}
}

// TODO: False sharing example (performance killer in concurrent code)
type FalseSharingStruct struct {
	a int64
	b int64 // These will likely be on the same cache line
}

// TODO: Fixed version with padding to avoid false sharing
type NoFalseSharingStruct struct {
	a   int64
	_   [7]int64 // Padding to force different cache lines
	b   int64
}

func demonstrateFalseSharing(iterations int) {
	// TODO: False sharing version
	shared := &FalseSharingStruct{}
	var wg sync.WaitGroup
	
	start := time.Now()
	
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			// TODO: Increment field a
			// atomic.AddInt64(&shared.a, 1)
		}
	}()
	
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			// TODO: Increment field b (false sharing with a)
			// atomic.AddInt64(&shared.b, 1)
		}
	}()
	
	wg.Wait()
	falseSharingTime := time.Since(start)
	
	// TODO: No false sharing version
	noShared := &NoFalseSharingStruct{}
	
	start = time.Now()
	wg.Add(2)
	
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			// TODO: Increment field a (no false sharing)
			// atomic.AddInt64(&noShared.a, 1)
		}
	}()
	
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			// TODO: Increment field b (no false sharing due to padding)
			// atomic.AddInt64(&noShared.b, 1)
		}
	}()
	
	wg.Wait()
	noFalseSharingTime := time.Since(start)
	
	fmt.Printf("False sharing time: %v\n", falseSharingTime)
	fmt.Printf("No false sharing time: %v\n", noFalseSharingTime)
	fmt.Printf("Performance improvement: %.2fx\n", float64(falseSharingTime)/float64(noFalseSharingTime))
}

// TODO: Demonstrate panic recovery in goroutines
func panickyGoroutine(id int) {
	defer func() {
		// TODO: Recover from panic
		// if r := recover(); r != nil {
		//     fmt.Printf("Goroutine %d recovered from panic: %v\n", id, r)
		//     
		//     // TODO: Print stack trace
		//     debug.PrintStack()
		// }
	}()
	
	// TODO: Randomly panic to demonstrate recovery
	if rand.Float32() < 0.3 { // 30% chance of panic
		panic(fmt.Sprintf("Intentional panic in goroutine %d", id))
	}
	
	fmt.Printf("Goroutine %d completed successfully\n", id)
}

// TODO: Benchmark different synchronization primitives
func benchmarkSync() {
	const iterations = 1000000
	
	// TODO: Benchmark mutex
	var mu sync.Mutex
	counter := 0
	
	start := time.Now()
	var wg sync.WaitGroup
	
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations/4; j++ {
				// TODO: Use mutex to protect counter
				// mu.Lock()
				// counter++
				// mu.Unlock()
			}
		}()
	}
	wg.Wait()
	mutexTime := time.Since(start)
	
	// TODO: Benchmark atomic operations
	var atomicCounter int64
	
	start = time.Now()
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations/4; j++ {
				// TODO: Use atomic increment
				// atomic.AddInt64(&atomicCounter, 1)
			}
		}()
	}
	wg.Wait()
	atomicTime := time.Since(start)
	
	fmt.Printf("Mutex time: %v (result: %d)\n", mutexTime, counter)
	fmt.Printf("Atomic time: %v (result: %d)\n", atomicTime, atomicCounter)
	fmt.Printf("Atomic is %.2fx faster\n", float64(mutexTime)/float64(atomicTime))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("=== Goroutine Debugging and Performance Analysis ===")
	
	// TODO: Start pprof HTTP server for profiling
	go func() {
		log.Println("pprof server starting on :6060")
		log.Println("Visit http://localhost:6060/debug/pprof/ for profiling")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	
	fmt.Println("\n=== Race Condition Demonstration ===")
	
	// TODO: Demonstrate unsafe counter (race condition)
	fmt.Println("Testing unsafe counter (race conditions):")
	unsafe := &UnsafeCounter{}
	var wg sync.WaitGroup
	
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				unsafe.Increment()
			}
		}()
	}
	wg.Wait()
	
	fmt.Printf("Unsafe counter result: %d (should be 100000)\n", unsafe.Get())
	
	// TODO: Demonstrate safe counter (no race conditions)
	fmt.Println("Testing safe counter (atomic operations):")
	safe := &SafeCounter{}
	
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				safe.Increment()
			}
		}()
	}
	wg.Wait()
	
	fmt.Printf("Safe counter result: %d (should be 100000)\n", safe.Get())
	
	fmt.Println("\n=== Goroutine Leak Prevention ===")
	
	// TODO: Demonstrate goroutine leak prevention
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	// TODO: Use non-leaky function with context
	ch := nonLeakyFunction(ctx)
	
	// TODO: Read limited number of values
	count := 0
	for value := range ch {
		if count < 5 {
			fmt.Printf("Received: %d\n", value)
		}
		count++
		if count >= 10 { // Stop early to test cancellation
			break
		}
	}
	
	fmt.Printf("Processed %d values before cancellation\n", count)
	
	fmt.Println("\n=== Performance Comparison ===")
	
	// TODO: Compare sequential vs concurrent performance
	n := 1000000
	
	start := time.Now()
	seqResult := sequentialWork(n)
	seqTime := time.Since(start)
	
	start = time.Now()
	concResult := concurrentWork(n, runtime.NumCPU())
	concTime := time.Since(start)
	
	fmt.Printf("Sequential: %v (result: %d)\n", seqTime, seqResult)
	fmt.Printf("Concurrent: %v (result: %d)\n", concTime, concResult)
	
	if concTime < seqTime {
		fmt.Printf("Concurrent is %.2fx faster\n", float64(seqTime)/float64(concTime))
	} else {
		fmt.Printf("Sequential is %.2fx faster\n", float64(concTime)/float64(seqTime))
	}
	
	fmt.Println("\n=== Goroutine Pool ===")
	
	// TODO: Demonstrate goroutine pool usage
	pool := NewGoroutinePool(5)
	
	// TODO: Submit work to pool
	for i := 0; i < 20; i++ {
		taskID := i
		pool.Submit(func() {
			fmt.Printf("Pool task %d executed\n", taskID)
			time.Sleep(100 * time.Millisecond)
		})
	}
	
	time.Sleep(1 * time.Second) // Let tasks complete
	pool.Stop()
	
	fmt.Println("\n=== Resource Monitoring ===")
	
	// TODO: Start resource monitoring
	monitorCtx, monitorCancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer monitorCancel()
	
	go monitorResources(monitorCtx, 500*time.Millisecond)
	
	// TODO: Create some load to monitor
	for i := 0; i < 50; i++ {
		go func(id int) {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			fmt.Printf("Worker %d done\n", id)
		}(i)
	}
	
	time.Sleep(3500 * time.Millisecond) // Let monitoring complete
	
	fmt.Println("\n=== False Sharing Demonstration ===")
	
	// TODO: Demonstrate false sharing performance impact
	demonstrateFalseSharing(10000000)
	
	fmt.Println("\n=== Panic Recovery in Goroutines ===")
	
	// TODO: Start goroutines that may panic
	for i := 0; i < 10; i++ {
		go panickyGoroutine(i)
	}
	
	time.Sleep(1 * time.Second) // Let all goroutines complete
	
	fmt.Println("\n=== Synchronization Benchmarks ===")
	
	// TODO: Benchmark different synchronization approaches
	benchmarkSync()
	
	fmt.Println("\n=== Runtime Information ===")
	
	// TODO: Print detailed runtime information
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("Go version: %s\n", runtime.Version())
	
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory allocated: %d KB\n", m.Alloc/1024)
	fmt.Printf("Total allocations: %d\n", m.TotalAlloc/1024)
	fmt.Printf("Heap objects: %d\n", m.HeapObjects)
	
	// TODO: Force garbage collection and show stats
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf("After GC - Memory allocated: %d KB\n", m.Alloc/1024)
	
	fmt.Println("\nGoroutine debugging and performance analysis completed!")
	fmt.Println("Note: Run with 'go run -race' to detect race conditions")
	fmt.Println("Use 'go tool pprof http://localhost:6060/debug/pprof/goroutine' for goroutine profiling")
}