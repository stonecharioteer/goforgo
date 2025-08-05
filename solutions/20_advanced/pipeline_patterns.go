// pipeline_patterns.go - SOLUTION
// Learn advanced concurrency patterns: pipelines, fan-in, fan-out

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Basic pipeline stages

// Generator stage - produces values
func generator(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

// Square stage - squares input values
func square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case n, ok := <-in:
				if !ok {
					return
				}
				select {
				case out <- n * n:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

// Fan-out pattern - distribute work to multiple workers
func fanOut(ctx context.Context, in <-chan int, workers int) []<-chan int {
	channels := make([]<-chan int, workers)
	
	// Create worker channels
	for i := 0; i < workers; i++ {
		out := make(chan int)
		channels[i] = out
		
		// Start worker goroutine
		go func(ch chan int) {
			defer close(ch)
			for {
				select {
				case n, ok := <-in:
					if !ok {
						return
					}
					// Simulate work
					time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
					select {
					case ch <- n * n: // Square the number
					case <-ctx.Done():
						return
					}
				case <-ctx.Done():
					return
				}
			}
		}(out)
	}
	
	return channels
}

// Fan-in pattern - combine multiple channels into one
func fanIn(ctx context.Context, channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	
	// Start goroutine for each input channel
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for {
				select {
				case n, ok := <-c:
					if !ok {
						return
					}
					select {
					case out <- n:
					case <-ctx.Done():
						return
					}
				case <-ctx.Done():
					return
				}
			}
		}(ch)
	}
	
	// Close output channel when all input channels are done
	go func() {
		wg.Wait()
		close(out)
	}()
	
	return out
}

// Pipeline with processing stages
type ProcessingPipeline struct {
	input   chan int
	output  chan int
	ctx     context.Context
	cancel  context.CancelFunc
	wg      sync.WaitGroup
}

func NewProcessingPipeline() *ProcessingPipeline {
	ctx, cancel := context.WithCancel(context.Background())
	return &ProcessingPipeline{
		input:  make(chan int, 10),
		output: make(chan int, 10),
		ctx:    ctx,
		cancel: cancel,
	}
}

func (p *ProcessingPipeline) Start() {
	// Create processing stages
	
	// Stage 1: Multiply by 2
	stage1 := make(chan int, 10)
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		defer close(stage1)
		
		for {
			select {
			case n, ok := <-p.input:
				if !ok {
					return
				}
				select {
				case stage1 <- n * 2:
				case <-p.ctx.Done():
					return
				}
			case <-p.ctx.Done():
				return
			}
		}
	}()
	
	// Stage 2: Add 10
	stage2 := make(chan int, 10)
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		defer close(stage2)
		
		for {
			select {
			case n, ok := <-stage1:
				if !ok {
					return
				}
				select {
				case stage2 <- n + 10:
				case <-p.ctx.Done():
					return
				}
			case <-p.ctx.Done():
				return
			}
		}
	}()
	
	// Stage 3: Square the result
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		defer close(p.output)
		
		for {
			select {
			case n, ok := <-stage2:
				if !ok {
					return
				}
				select {
				case p.output <- n * n:
				case <-p.ctx.Done():
					return
				}
			case <-p.ctx.Done():
				return
			}
		}
	}()
}

func (p *ProcessingPipeline) Send(value int) error {
	select {
	case p.input <- value:
		return nil
	case <-p.ctx.Done():
		return p.ctx.Err()
	}
}

func (p *ProcessingPipeline) Receive() (int, bool) {
	select {
	case value, ok := <-p.output:
		return value, ok
	case <-p.ctx.Done():
		return 0, false
	}
}

func (p *ProcessingPipeline) Close() {
	close(p.input)
	p.wg.Wait()
	p.cancel()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("=== Basic Pipeline ===")
	
	// Create basic pipeline
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	// Create pipeline: generator -> square -> output
	numbers := generator(ctx, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	squared := square(ctx, numbers)
	
	// Consume results
	fmt.Println("Squared numbers:")
	for result := range squared {
		fmt.Printf("%d ", result)
	}
	fmt.Println()
	
	fmt.Println("\n=== Fan-Out/Fan-In Pattern ===")
	
	// Create fan-out/fan-in pipeline
	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()
	
	// Generate numbers
	input := generator(ctx2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	
	// Fan out to 3 workers
	workers := fanOut(ctx2, input, 3)
	
	// Fan in results
	results := fanIn(ctx2, workers...)
	
	// Collect results
	fmt.Println("Fan-out/Fan-in results:")
	var allResults []int
	for result := range results {
		allResults = append(allResults, result)
	}
	
	fmt.Printf("Received %d results: %v\n", len(allResults), allResults)
	
	fmt.Println("\n=== Processing Pipeline ===")
	
	// Test processing pipeline
	pipeline := NewProcessingPipeline()
	pipeline.Start()
	
	// Send data to pipeline
	go func() {
		for i := 1; i <= 5; i++ {
			err := pipeline.Send(i)
			if err != nil {
				fmt.Printf("Failed to send %d: %v\n", i, err)
				break
			}
			fmt.Printf("Sent: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
		pipeline.Close()
	}()
	
	// Receive results
	fmt.Println("Processing pipeline results:")
	for {
		result, ok := pipeline.Receive()
		if !ok {
			break
		}
		fmt.Printf("Received: %d\n", result)
	}
	
	fmt.Println("\nPipeline patterns completed")
}