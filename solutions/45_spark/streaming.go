// GoForGo Solution: Spark Streaming
// Complete implementation of streaming data processing patterns

package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// StreamingContext manages streaming operations
type StreamingContext struct {
	BatchDuration time.Duration
	Context       context.Context
	Cancel        context.CancelFunc
	wg            sync.WaitGroup
	streams       []*DStream
}

// DStream represents a discretized stream
type DStream struct {
	Name        string
	BatchChan   chan []interface{}
	ProcessFunc func([]interface{}) []interface{}
	Context     context.Context
	wg          *sync.WaitGroup
}

// NewStreamingContext creates a new streaming context
func NewStreamingContext(batchDuration time.Duration) *StreamingContext {
	ctx, cancel := context.WithCancel(context.Background())
	ssc := &StreamingContext{
		BatchDuration: batchDuration,
		Context:       ctx,
		Cancel:        cancel,
		streams:       make([]*DStream, 0),
	}
	
	log.Printf("Created StreamingContext with batch duration: %v", batchDuration)
	return ssc
}

// SocketTextStream simulates receiving text data from a socket
func (ssc *StreamingContext) SocketTextStream(host string, port int) *DStream {
	log.Printf("Creating socket text stream: %s:%d", host, port)
	
	batchChan := make(chan []interface{}, 10)
	
	ds := &DStream{
		Name:        fmt.Sprintf("SocketStream_%s_%d", host, port),
		BatchChan:   batchChan,
		ProcessFunc: func(batch []interface{}) []interface{} { return batch }, // identity
		Context:     ssc.Context,
		wg:          &ssc.wg,
	}
	
	ssc.streams = append(ssc.streams, ds)
	
	// Generate sample log entries
	ssc.wg.Add(1)
	go func() {
		defer ssc.wg.Done()
		defer close(batchChan)
		
		ticker := time.NewTicker(ssc.BatchDuration)
		defer ticker.Stop()
		
		logTemplates := []string{
			"INFO: User login successful - user123",
			"ERROR: Database connection failed - timeout",
			"WARN: High memory usage detected - 85%",
			"ERROR: Failed to process request - invalid data",
			"INFO: Cache refreshed successfully",
			"ERROR: Authentication failed - invalid token",
			"DEBUG: Processing batch job #%d",
		}
		
		batchNum := 0
		for {
			select {
			case <-ds.Context.Done():
				return
			case <-ticker.C:
				// Generate a batch of log entries
				batch := make([]interface{}, 0, 5)
				for i := 0; i < 3+rand.Intn(3); i++ {
					template := logTemplates[rand.Intn(len(logTemplates))]
					if strings.Contains(template, "%d") {
						batch = append(batch, fmt.Sprintf(template, batchNum))
					} else {
						batch = append(batch, template)
					}
				}
				
				select {
				case batchChan <- batch:
					log.Printf("Generated batch %d with %d log entries", batchNum, len(batch))
				case <-ds.Context.Done():
					return
				}
				
				batchNum++
			}
		}
	}()
	
	return ds
}

// Map transforms each element in the stream
func (ds *DStream) Map(mapFunc func(interface{}) interface{}) *DStream {
	log.Printf("Adding Map transformation to %s", ds.Name)
	
	outputChan := make(chan []interface{}, 10)
	
	newDS := &DStream{
		Name:      fmt.Sprintf("%s_mapped", ds.Name),
		BatchChan: outputChan,
		Context:   ds.Context,
		wg:        ds.wg,
	}
	
	// Process batches from input stream
	ds.wg.Add(1)
	go func() {
		defer ds.wg.Done()
		defer close(outputChan)
		
		for {
			select {
			case <-ds.Context.Done():
				return
			case batch, ok := <-ds.BatchChan:
				if !ok {
					return
				}
				
				// Apply transformation to each element
				mappedBatch := make([]interface{}, len(batch))
				for i, element := range batch {
					mappedBatch[i] = mapFunc(element)
				}
				
				select {
				case outputChan <- mappedBatch:
				case <-ds.Context.Done():
					return
				}
			}
		}
	}()
	
	return newDS
}

// Filter elements in the stream based on predicate
func (ds *DStream) Filter(filterFunc func(interface{}) bool) *DStream {
	log.Printf("Adding Filter transformation to %s", ds.Name)
	
	outputChan := make(chan []interface{}, 10)
	
	newDS := &DStream{
		Name:      fmt.Sprintf("%s_filtered", ds.Name),
		BatchChan: outputChan,
		Context:   ds.Context,
		wg:        ds.wg,
	}
	
	// Filter batches from input stream
	ds.wg.Add(1)
	go func() {
		defer ds.wg.Done()
		defer close(outputChan)
		
		for {
			select {
			case <-ds.Context.Done():
				return
			case batch, ok := <-ds.BatchChan:
				if !ok {
					return
				}
				
				// Filter elements in batch
				filteredBatch := make([]interface{}, 0, len(batch))
				for _, element := range batch {
					if filterFunc(element) {
						filteredBatch = append(filteredBatch, element)
					}
				}
				
				select {
				case outputChan <- filteredBatch:
				case <-ds.Context.Done():
					return
				}
			}
		}
	}()
	
	return newDS
}

// Window applies window operations to accumulate data over time
func (ds *DStream) Window(windowDuration, slideDuration time.Duration) *DStream {
	log.Printf("Adding Window operation to %s (window: %v, slide: %v)", ds.Name, windowDuration, slideDuration)
	
	outputChan := make(chan []interface{}, 10)
	
	newDS := &DStream{
		Name:      fmt.Sprintf("%s_windowed", ds.Name),
		BatchChan: outputChan,
		Context:   ds.Context,
		wg:        ds.wg,
	}
	
	// Simplified window implementation
	ds.wg.Add(1)
	go func() {
		defer ds.wg.Done()
		defer close(outputChan)
		
		var window []interface{}
		slideTimer := time.NewTicker(slideDuration)
		defer slideTimer.Stop()
		
		for {
			select {
			case <-ds.Context.Done():
				return
			case batch, ok := <-ds.BatchChan:
				if !ok {
					return
				}
				
				// Add batch to window
				window = append(window, batch...)
				
			case <-slideTimer.C:
				if len(window) > 0 {
					// Send current window contents
					windowCopy := make([]interface{}, len(window))
					copy(windowCopy, window)
					
					select {
					case outputChan <- windowCopy:
						log.Printf("Emitted window with %d elements", len(windowCopy))
					case <-ds.Context.Done():
						return
					}
					
					// Clear window (simplified - should maintain proper window duration)
					window = window[:0]
				}
			}
		}
	}()
	
	return newDS
}

// Print outputs the first numElements of each batch to console
func (ds *DStream) Print(numElements int) {
	log.Printf("Adding Print action to %s", ds.Name)
	
	ds.wg.Add(1)
	go func() {
		defer ds.wg.Done()
		
		batchNum := 0
		for {
			select {
			case <-ds.Context.Done():
				return
			case batch, ok := <-ds.BatchChan:
				if !ok {
					return
				}
				
				fmt.Printf("\n=== Batch %d (%s) ===\n", batchNum, ds.Name)
				elementsToShow := numElements
				if elementsToShow > len(batch) {
					elementsToShow = len(batch)
				}
				
				for i := 0; i < elementsToShow; i++ {
					fmt.Printf("%v\n", batch[i])
				}
				
				if len(batch) > numElements {
					fmt.Printf("... and %d more elements\n", len(batch)-numElements)
				}
				
				fmt.Printf("Total elements in batch: %d\n", len(batch))
				batchNum++
			}
		}
	}()
}

// Start begins processing streaming data
func (ssc *StreamingContext) Start() {
	log.Println("Starting StreamingContext...")
}

// AwaitTermination waits for streaming to finish or timeout
func (ssc *StreamingContext) AwaitTermination(timeout time.Duration) error {
	log.Printf("Waiting for termination (timeout: %v)...", timeout)
	
	done := make(chan struct{})
	go func() {
		ssc.wg.Wait()
		close(done)
	}()
	
	select {
	case <-done:
		log.Println("All streaming operations completed")
		return nil
	case <-time.After(timeout):
		log.Println("Timeout reached, stopping...")
		ssc.Stop()
		return fmt.Errorf("timeout after %v", timeout)
	}
}

// Stop gracefully stops all streaming operations
func (ssc *StreamingContext) Stop() {
	log.Println("Stopping StreamingContext...")
	ssc.Cancel()
	ssc.wg.Wait()
	log.Println("StreamingContext stopped")
}

func main() {
	// Create a StreamingContext with 2-second batch intervals
	ssc := NewStreamingContext(2 * time.Second)
	
	// Create a socket text stream (simulated)
	logStream := ssc.SocketTextStream("localhost", 9999)
	
	// Apply transformations to the stream
	upperStream := logStream.Map(func(element interface{}) interface{} {
		return strings.ToUpper(element.(string))
	})
	
	errorStream := upperStream.Filter(func(element interface{}) bool {
		return strings.Contains(element.(string), "ERROR")
	})
	
	windowedStream := errorStream.Window(10*time.Second, 5*time.Second)
	
	// Add output action to print first 10 elements of each batch
	windowedStream.Print(10)
	
	// Start the streaming context
	ssc.Start()
	
	// Wait for 30 seconds, then stop
	if err := ssc.AwaitTermination(30 * time.Second); err != nil {
		log.Printf("Streaming terminated with error: %v", err)
	}
	
	fmt.Println("Spark streaming operations completed!")
}