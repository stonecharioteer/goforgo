// sync_primitives.go
// Learn synchronization primitives: Mutex, RWMutex, Once, Cond

package main

import (
	"fmt"
	"sync"
	"time"
)

// TODO: Shared counter with mutex protection
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment() {
	// TODO: Lock, increment, unlock
}

func (c *SafeCounter) Value() int {
	// TODO: Lock, read value, unlock, return
}

// TODO: Cache with RWMutex
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]string)}
}

func (c *Cache) Get(key string) (string, bool) {
	// TODO: Use RLock for reading
}

func (c *Cache) Set(key, value string) {
	// TODO: Use Lock for writing
}

// TODO: Expensive initialization with sync.Once
var (
	instance *ExpensiveResource
	once     sync.Once
)

type ExpensiveResource struct {
	data string
}

func GetInstance() *ExpensiveResource {
	// TODO: Use once.Do to initialize instance only once
	return instance
}

func createInstance() {
	fmt.Println("Creating expensive resource...")
	time.Sleep(100 * time.Millisecond) // Simulate expensive operation
	instance = &ExpensiveResource{data: "initialized"}
}

func main() {
	fmt.Println("=== Mutex Example ===")
	
	counter := &SafeCounter{}
	var wg sync.WaitGroup
	
	// TODO: Start multiple goroutines incrementing counter
	numGoroutines := 10
	incrementsPerGoroutine := 100
	
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				counter.Increment()
			}
		}()
	}
	
	wg.Wait()
	expected := numGoroutines * incrementsPerGoroutine
	actual := counter.Value()
	fmt.Printf("Expected: %d, Actual: %d\\n", expected, actual)
	
	fmt.Println("\\n=== RWMutex Example ===")
	
	cache := NewCache()
	
	// TODO: Start readers and writers
	numReaders := 5
	numWriters := 2
	
	wg.Add(numReaders + numWriters)
	
	// Start writers
	for i := 0; i < numWriters; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				key := fmt.Sprintf("key%d_%d", id, j)
				value := fmt.Sprintf("value%d_%d", id, j)
				cache.Set(key, value)
				fmt.Printf("Writer %d: Set %s = %s\\n", id, key, value)
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}
	
	// Start readers
	for i := 0; i < numReaders; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				key := fmt.Sprintf("key%d_0", id%numWriters)
				if value, ok := cache.Get(key); ok {
					fmt.Printf("Reader %d: Got %s = %s\\n", id, key, value)
				} else {
					fmt.Printf("Reader %d: Key %s not found\\n", id, key)
				}
				time.Sleep(20 * time.Millisecond)
			}
		}(i)
	}
	
	wg.Wait()
	
	fmt.Println("\\n=== sync.Once Example ===")
	
	// TODO: Multiple goroutines trying to get instance
	numGetters := 5
	wg.Add(numGetters)
	
	for i := 0; i < numGetters; i++ {
		go func(id int) {
			defer wg.Done()
			instance := GetInstance()
			fmt.Printf("Goroutine %d got instance: %s\\n", id, instance.data)
		}(i)
	}
	
	wg.Wait()
}