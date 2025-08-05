// channels_basics.go - SOLUTION
// Learn the fundamentals of channels - Go's way of goroutine communication

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Unbuffered Channels ===")
	
	// Create unbuffered channel
	ch := make(chan int)
	
	// Send and receive in different goroutines
	go func() {
		// Send values to channel
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch) // Close channel when done sending
	}()
	
	// Receive values from channel
	// Use range to receive all values until channel is closed
	for value := range ch {
		fmt.Printf("Received: %d\n", value)
	}
	
	fmt.Println("\n=== Buffered Channels ===")
	
	// Create buffered channel
	buffered := make(chan int, 3)
	
	// Send values without blocking (up to capacity)
	buffered <- 10
	buffered <- 20
	buffered <- 30
	
	// Receive values
	fmt.Printf("Received: %d\n", <-buffered)
	fmt.Printf("Received: %d\n", <-buffered)
	fmt.Printf("Received: %d\n", <-buffered)
	
	fmt.Println("\n=== Channel Directions ===")
	
	// Function with send-only channel parameter
	sender := func(ch chan<- string) {
		// Send messages to channel
		ch <- "Hello"
		ch <- "World"
		close(ch)
	}
	
	// Function with receive-only channel parameter  
	receiver := func(ch <-chan string) {
		// Receive and print all messages
		for msg := range ch {
			fmt.Printf("Received: %s\n", msg)
		}
	}
	
	msgCh := make(chan string)
	go sender(msgCh)
	receiver(msgCh)
	
	fmt.Println("\n=== Channel Select ===")
	
	// Use select to handle multiple channels
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
	
	// Use select to receive from first available channel
	// Handle both channels and print which one responded first
	select {
	case msg1 := <-ch1:
		fmt.Printf("Received: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("Received: %s\n", msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout - no message received")
	}
	
	// Receive from the second channel (if it hasn't sent yet)
	select {
	case msg2 := <-ch2:
		fmt.Printf("Received: %s\n", msg2)
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout waiting for second message")
	}
}