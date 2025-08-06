// goroutine_debugging.go - Solution
// Learn goroutine debugging, performance analysis, race detection, and leak prevention

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
	_ "net/http/pprof"
	"net/http"
)

// Example of a race condition (intentional for demonstration)
type UnsafeCounter struct {
	value int64
}

func (c *UnsafeCounter) Increment() {
	c.value++ // Race condition!
}

func (c *UnsafeCounter) Get() int64 {
	return c.value // Race condition!
}

// Safe counter using atomic operations
type SafeCounter struct {
	value int64
}

func (c *SafeCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *SafeCounter) Get() int64 {
	return atomic.LoadInt64(&c.value)
}

// Fixed version that prevents leaks
func nonLeakyFunction(ctx context.Context) <-chan int {
	ch := make(chan int)
	
	go func() {
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

// Memory ordering variables
var (
	data  int32
	ready int32
)

func writer() {
	atomic.StoreInt32(&data, 42)
	atomic.StoreInt32(&ready, 1)
}

func reader() int32 {
	if atomic.LoadInt32(&ready) == 1 {
		return atomic.LoadInt32(&data)
	}
	return 0
}

// Benchmark functions
func sequentialWork(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i * i
	}
	return sum
}

func concurrentWork(n int, workers int) int {
	if workers <= 0 {
		workers = runtime.NumCPU()
	}
	
	workPerWorker := n / workers
	results := make(chan int, workers)
	
	for i := 0; i < workers; i++ {
		go func(start, end int) {
			sum := 0
			for j := start; j < end; j++ {
				sum += j * j
			}
			results <- sum
		}(i*workPerWorker, (i+1)*workPerWorker)
	}
	
	totalSum := 0
	for i := 0; i < workers; i++ {
		totalSum += <-results
	}
	
	return totalSum
}

// Goroutine pool
type GoroutinePool struct {
	workers int
	workCh  chan func()
	quit    chan struct{}
	wg      sync.WaitGroup
}

func NewGoroutinePool(workers int) *GoroutinePool {
	pool := &GoroutinePool{
		workers: workers,
		workCh:  make(chan func(), workers*2),
		quit:    make(chan struct{}),
	}
	
	for i := 0; i < workers; i++ {
		pool.wg.Add(1)
		go func() {
			defer pool.wg.Done()
			for {
				select {
				case work := <-pool.workCh:
					work()
				case <-pool.quit:
					return
				}
			}
		}()
	}
	
	return pool
}

func (p *GoroutinePool) Submit(work func()) {
	select {
	case p.workCh <- work:
	case <-p.quit:
		return
	}
}

func (p *GoroutinePool) Stop() {
	close(p.quit)
	p.wg.Wait()
}

// Monitor resources
func monitorResources(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	
	var lastGC debug.GCStats
	debug.ReadGCStats(&lastGC)
	
	for {
		select {
		case <-ticker.C:
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			
			var gc debug.GCStats
			debug.ReadGCStats(&gc)
			
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

// False sharing structs
type FalseSharingStruct struct {
	a int64
	b int64
}

type NoFalseSharingStruct struct {
	a   int64
	_   [7]int64 // Padding
	b   int64
}

func demonstrateFalseSharing(iterations int) {
	shared := &FalseSharingStruct{}
	var wg sync.WaitGroup
	
	start := time.Now()
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			atomic.AddInt64(&shared.a, 1)
		}
	}()
	
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			atomic.AddInt64(&shared.b, 1)
		}
	}()
	
	wg.Wait()
	falseSharingTime := time.Since(start)
	
	noShared := &NoFalseSharingStruct{}
	start = time.Now()
	wg.Add(2)
	
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			atomic.AddInt64(&noShared.a, 1)
		}
	}()
	
	go func() {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			atomic.AddInt64(&noShared.b, 1)
		}
	}()
	
	wg.Wait()
	noFalseSharingTime := time.Since(start)
	
	fmt.Printf("False sharing time: %v\n", falseSharingTime)
	fmt.Printf("No false sharing time: %v\n", noFalseSharingTime)
	if noFalseSharingTime > 0 {
		fmt.Printf("Performance improvement: %.2fx\n", float64(falseSharingTime)/float64(noFalseSharingTime))
	}
}

func panickyGoroutine(id int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Goroutine %d recovered from panic: %v\n", id, r)
		}
	}()
	
	if rand.Float32() < 0.3 {
		panic(fmt.Sprintf("Intentional panic in goroutine %d", id))
	}
	
	fmt.Printf("Goroutine %d completed successfully\n", id)
}

func benchmarkSync() {
	const iterations = 1000000
	
	// Benchmark mutex
	var mu sync.Mutex
	counter := 0
	
	start := time.Now()
	var wg sync.WaitGroup
	
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations/4; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	mutexTime := time.Since(start)
	
	// Benchmark atomic operations
	var atomicCounter int64
	
	start = time.Now()
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations/4; j++ {
				atomic.AddInt64(&atomicCounter, 1)
			}
		}()
	}
	wg.Wait()
	atomicTime := time.Since(start)
	
	fmt.Printf("Mutex time: %v (result: %d)\n", mutexTime, counter)
	fmt.Printf("Atomic time: %v (result: %d)\n", atomicTime, atomicCounter)
	if atomicTime > 0 {
		fmt.Printf("Atomic is %.2fx faster\n", float64(mutexTime)/float64(atomicTime))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("=== Goroutine Debugging and Performance Analysis ===")
	
	go func() {
		log.Println("pprof server starting on :6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	
	fmt.Println("\n=== Race Condition Demonstration ===")
	
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
	
	fmt.Println("\n=== Performance Comparison ===")
	
	n := 1000000
	
	start := time.Now()
	seqResult := sequentialWork(n)
	seqTime := time.Since(start)
	
	start = time.Now()
	concResult := concurrentWork(n, runtime.NumCPU())
	concTime := time.Since(start)
	
	fmt.Printf("Sequential: %v (result: %d)\n", seqTime, seqResult)
	fmt.Printf("Concurrent: %v (result: %d)\n", concTime, concResult)
	
	fmt.Println("\n=== False Sharing Demonstration ===")
	demonstrateFalseSharing(1000000)
	
	fmt.Println("\n=== Synchronization Benchmarks ===")
	benchmarkSync()
	
	fmt.Printf("\nRuntime info: %d goroutines\n", runtime.NumGoroutine())
	fmt.Println("Goroutine debugging completed!")
}