// advanced_sync.go
// Learn advanced synchronization primitives: Cond, Pool, Map, errgroup, semaphore, singleflight
//
// This exercise covers advanced synchronization primitives that are used in
// high-performance and complex concurrent systems. These patterns help manage
// resource pooling, conditional synchronization, and coordinated goroutine execution.
//
// I AM NOT DONE YET!

package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
	"golang.org/x/sync/singleflight"
)

// TODO: Producer-Consumer with sync.Cond
type Buffer struct {
	items []int
	cond  *sync.Cond
	// TODO: Add mutex for protecting the buffer
	// mu    ???
}

func NewBuffer() *Buffer {
	// TODO: Initialize buffer with condition variable
	// return &Buffer{
	//     items: make([]int, 0, 10),
	//     cond:  sync.NewCond(???),
	// }
}

func (b *Buffer) Put(item int) {
	// TODO: Lock, add item, signal waiting consumers
	// b.cond.L.Lock()
	// defer b.cond.L.Unlock()
	
	// TODO: Add item to buffer
	// b.items = append(b.items, item)
	// fmt.Printf("Produced: %d (buffer size: %d)\n", item, len(b.items))
	
	// TODO: Signal waiting consumers
	// b.cond.Signal()
}

func (b *Buffer) Get() int {
	// TODO: Lock, wait for items, consume item
	// b.cond.L.Lock()
	// defer b.cond.L.Unlock()
	
	// TODO: Wait while buffer is empty
	// for len(b.items) == 0 {
	//     b.cond.Wait()
	// }
	
	// TODO: Remove and return first item
	// item := b.items[0]
	// b.items = b.items[1:]
	// fmt.Printf("Consumed: %d (buffer size: %d)\n", item, len(b.items))
	// return item
}

// TODO: Object pooling with sync.Pool
type ExpensiveObject struct {
	ID   int
	Data []byte
}

var objectPool = sync.Pool{
	// TODO: New function to create objects when pool is empty
	// New: func() interface{} {
	//     fmt.Println("Creating new expensive object")
	//     return &ExpensiveObject{
	//         Data: make([]byte, 1024),
	//     }
	// },
}

func useExpensiveObject(id int) {
	// TODO: Get object from pool
	// obj := objectPool.Get().(*ExpensiveObject)
	// obj.ID = id
	
	// TODO: Use the object (simulate work)
	// fmt.Printf("Using object %d\n", obj.ID)
	// time.Sleep(100 * time.Millisecond)
	
	// TODO: Reset and return to pool
	// obj.ID = 0
	// objectPool.Put(obj)
}

// TODO: Concurrent map operations with sync.Map
func demonstrateSyncMap() {
	fmt.Println("\n=== sync.Map Example ===")
	
	var m sync.Map
	
	// TODO: Start multiple goroutines for concurrent map operations
	var wg sync.WaitGroup
	
	// Writer goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				key := fmt.Sprintf("key-%d-%d", id, j)
				value := fmt.Sprintf("value-%d-%d", id, j)
				
				// TODO: Store key-value pair
				// m.Store(key, value)
				fmt.Printf("Writer %d stored %s\n", id, key)
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}
	
	// Reader goroutines
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond) // Let writers start first
			
			// TODO: Range over all key-value pairs
			// m.Range(func(key, value interface{}) bool {
			//     fmt.Printf("Reader %d found %s = %s\n", id, key, value)
			//     return true
			// })
		}(i)
	}
	
	wg.Wait()
}

// TODO: Error group for coordinated goroutines
func demonstrateErrGroup() {
	fmt.Println("\n=== errgroup Example ===")
	
	// TODO: Create error group with context
	// g, ctx := errgroup.WithContext(context.Background())
	
	// URLs to fetch (simulate with delays)
	urls := []string{
		"https://api1.example.com",
		"https://api2.example.com", 
		"https://api3.example.com",
		"https://api4.example.com",
	}
	
	results := make([]string, len(urls))
	
	for i, url := range urls {
		// TODO: Capture loop variables
		// i, url := i, url
		
		// TODO: Add goroutine to error group
		// g.Go(func() error {
		//     // Simulate HTTP request
		//     select {
		//     case <-time.After(time.Duration(rand.Intn(300)+100) * time.Millisecond):
		//         if rand.Float32() < 0.2 { // 20% chance of error
		//             return fmt.Errorf("failed to fetch %s", url)
		//         }
		//         results[i] = fmt.Sprintf("Response from %s", url)
		//         fmt.Printf("Successfully fetched %s\n", url)
		//         return nil
		//     case <-ctx.Done():
		//         return ctx.Err()
		//     }
		// })
	}
	
	// TODO: Wait for all goroutines to complete
	// if err := g.Wait(); err != nil {
	//     fmt.Printf("Error occurred: %v\n", err)
	//     return
	// }
	
	fmt.Println("All requests completed successfully!")
	for i, result := range results {
		if result != "" {
			fmt.Printf("Result %d: %s\n", i, result)
		}
	}
}

// TODO: Semaphore for resource limiting
func demonstrateSemaphore() {
	fmt.Println("\n=== Semaphore Example ===")
	
	// TODO: Create weighted semaphore (limit to 3 concurrent operations)
	// sem := semaphore.NewWeighted(3)
	// ctx := context.Background()
	
	var wg sync.WaitGroup
	
	// Start 10 goroutines, but only 3 can run concurrently
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// TODO: Acquire semaphore (weight 1)
			// if err := sem.Acquire(ctx, 1); err != nil {
			//     fmt.Printf("Worker %d failed to acquire semaphore: %v\n", id, err)
			//     return
			// }
			// defer sem.Release(1)
			
			fmt.Printf("Worker %d acquired semaphore, starting work...\n", id)
			
			// TODO: Simulate work
			time.Sleep(time.Duration(rand.Intn(500)+200) * time.Millisecond)
			
			fmt.Printf("Worker %d finished work\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Println("All workers completed")
}

// TODO: SingleFlight for deduplicating expensive operations
var sf singleflight.Group

func expensiveOperation(key string) (string, error) {
	fmt.Printf("Performing expensive operation for key: %s\n", key)
	
	// TODO: Simulate expensive operation
	time.Sleep(1 * time.Second)
	
	// TODO: Simulate occasional errors
	if rand.Float32() < 0.1 { // 10% chance of error
		return "", fmt.Errorf("operation failed for key: %s", key)
	}
	
	return fmt.Sprintf("Result for %s", key), nil
}

func demonstrateSingleFlight() {
	fmt.Println("\n=== SingleFlight Example ===")
	
	var wg sync.WaitGroup
	
	// TODO: Multiple goroutines requesting the same expensive operation
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// TODO: Use singleflight to deduplicate calls
			// result, err, shared := sf.Do("expensive-key", func() (interface{}, error) {
			//     return expensiveOperation("expensive-key")
			// })
			
			// if err != nil {
			//     fmt.Printf("Goroutine %d got error: %v (shared: %t)\n", id, err, shared)
			// } else {
			//     fmt.Printf("Goroutine %d got result: %s (shared: %t)\n", id, result, shared)
			// }
		}(i)
	}
	
	// TODO: Add a second set of requests for different key
	time.Sleep(100 * time.Millisecond)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// TODO: Different key, should not be deduplicated with first set
			// result, err, shared := sf.Do("another-key", func() (interface{}, error) {
			//     return expensiveOperation("another-key")
			// })
			
			// if err != nil {
			//     fmt.Printf("Second group %d got error: %v (shared: %t)\n", id, err, shared)
			// } else {
			//     fmt.Printf("Second group %d got result: %s (shared: %t)\n", id, result, shared)
			// }
		}(i)
	}
	
	wg.Wait()
}

// TODO: Advanced pattern: Condition variable with timeout
func demonstrateCondWithTimeout() {
	fmt.Println("\n=== Condition Variable with Timeout ===")
	
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false
	
	// TODO: Consumer with timeout
	go func() {
		mu.Lock()
		defer mu.Unlock()
		
		// TODO: Wait with timeout using goroutine and channel
		timeout := make(chan struct{})
		go func() {
			time.Sleep(2 * time.Second)
			close(timeout)
		}()
		
		for !ready {
			// TODO: Create a goroutine to signal when condition is ready
			waitDone := make(chan struct{})
			go func() {
				cond.Wait()
				close(waitDone)
			}()
			
			// TODO: Race condition wait vs timeout
			select {
			case <-waitDone:
				// Condition was signaled
				if !ready {
					continue // Spurious wakeup, check condition again
				}
			case <-timeout:
				fmt.Println("Consumer timed out waiting for condition")
				return
			}
		}
		
		fmt.Println("Consumer: condition is ready!")
	}()
	
	// TODO: Producer delays signal
	time.Sleep(1 * time.Second)
	mu.Lock()
	ready = true
	cond.Signal()
	mu.Unlock()
	fmt.Println("Producer: signaled condition")
	
	time.Sleep(500 * time.Millisecond)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("=== Advanced Synchronization Primitives ===")
	
	fmt.Println("\n=== sync.Cond Producer-Consumer ===")
	
	// TODO: Demonstrate condition variables
	buffer := NewBuffer()
	
	// Start producer
	go func() {
		for i := 1; i <= 5; i++ {
			buffer.Put(i)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	
	// Start consumers
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				item := buffer.Get()
				fmt.Printf("Consumer %d got: %d\n", id, item)
			}
		}(i)
	}
	
	wg.Wait()
	
	fmt.Println("\n=== sync.Pool Object Pooling ===")
	
	// TODO: Demonstrate object pooling
	var poolWg sync.WaitGroup
	for i := 0; i < 10; i++ {
		poolWg.Add(1)
		go func(id int) {
			defer poolWg.Done()
			useExpensiveObject(id)
		}(i)
	}
	poolWg.Wait()
	
	// TODO: Force garbage collection to see pool behavior
	runtime.GC()
	fmt.Println("After GC - using pool again:")
	useExpensiveObject(999)
	
	// TODO: Call remaining demonstrations
	demonstrateSyncMap()
	demonstrateErrGroup()
	demonstrateSemaphore()
	demonstrateSingleFlight()
	demonstrateCondWithTimeout()
	
	fmt.Println("\nAdvanced synchronization examples completed!")
}