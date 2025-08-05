// channels_basics.go
// Learn the fundamentals of channels - Go's way of goroutine communication

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Unbuffered Channels ===")
	
	// TODO: Create unbuffered channel
	ch := // Create unbuffered int channel
	
	// TODO: Send and receive in different goroutines
	go func() {
		// Send values to channel
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch) // Close channel when done sending
	}()
	
	// TODO: Receive values from channel
	// Use range to receive all values until channel is closed
	
	fmt.Println("\\n=== Buffered Channels ===")
	
	// TODO: Create buffered channel
	buffered := // Create buffered channel with capacity 3
	
	// TODO: Send values without blocking (up to capacity)
	buffered <- 10
	buffered <- 20
	buffered <- 30
	
	// TODO: Receive values
	fmt.Printf("Received: %d\\n", /* receive from buffered */)
	fmt.Printf("Received: %d\\n", /* receive from buffered */)
	fmt.Printf("Received: %d\\n", /* receive from buffered */)
	
	fmt.Println("\\n=== Channel Directions ===")
	
	// TODO: Function with send-only channel parameter
	sender := func(ch chan<- string) {
		// Send messages to channel
		ch <- "Hello"
		ch <- "World"
		close(ch)
	}
	
	// TODO: Function with receive-only channel parameter  
	receiver := func(ch <-chan string) {
		// Receive and print all messages
		for msg := range ch {
			fmt.Printf("Received: %s\\n", msg)
		}
	}
	
	msgCh := make(chan string)
	go sender(msgCh)
	receiver(msgCh)
	
	fmt.Println("\\n=== Channel Select ===")
	
	// TODO: Use select to handle multiple channels
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "from ch1"
	}()
	
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "from ch2"
	}()
	
	// TODO: Use select to receive from first available channel
	// Handle both channels and print which one responded first
}