package main

import (
	"fmt"
	"time"
)

func main() {
	// Create two channels: ch1 := make(chan string), ch2 := make(chan string)
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	// Start a goroutine that sends "Hello from ch1" to ch1 after 1 second
	// Use: go func() { time.Sleep(time.Second); ch1 <- "Hello from ch1" }()
	go func() {
		time.Sleep(time.Second)
		ch1 <- "Hello from ch1"
	}()
	
	// Start a goroutine that sends "Hello from ch2" to ch2 after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Hello from ch2"
	}()
	
	// Use select to receive from either channel
	// Print which channel sent the message and what message was received
	fmt.Println("=== Select with multiple channels ===")
	select {
	case msg1 := <-ch1:
		fmt.Println("Received from ch1:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from ch2:", msg2)
	}
	
	// Demonstrate select with default case
	// Create a channel: ch3 := make(chan int)
	// Use select to try to receive from ch3 with a default case
	// Default case should print "No data available"
	fmt.Println("\n=== Select with default case ===")
	ch3 := make(chan int)
	select {
	case data := <-ch3:
		fmt.Println("Received from ch3:", data)
	default:
		fmt.Println("No data available")
	}
	
	// Demonstrate select with timeout using time.After
	// Create a channel: slow := make(chan string)
	// Start a goroutine that sends to slow after 3 seconds
	// Use select with slow and time.After(1*time.Second)
	// Print appropriate messages for both cases
	fmt.Println("\n=== Select with timeout ===")
	slow := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		slow <- "Slow message"
	}()
	
	select {
	case msg := <-slow:
		fmt.Println("Received slow message:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: no message received within 1 second")
	}
	
	// Demonstrate select with multiple ready channels
	// Create channels: fast1 := make(chan int, 1), fast2 := make(chan int, 1)
	// Send values immediately: fast1 <- 1, fast2 <- 2
	// Use select to receive from either (Go will choose randomly)
	// Run this in a loop 5 times to see random selection
	fmt.Println("\n=== Select with multiple ready channels (random selection) ===")
	for i := 0; i < 5; i++ {
		fast1 := make(chan int, 1)
		fast2 := make(chan int, 1)
		fast1 <- 1
		fast2 <- 2
		
		select {
		case val1 := <-fast1:
			fmt.Printf("Iteration %d: Received from fast1: %d\n", i+1, val1)
		case val2 := <-fast2:
			fmt.Printf("Iteration %d: Received from fast2: %d\n", i+1, val2)
		}
	}
	
	// Demonstrate select for sending (not just receiving)
	// Create: send1 := make(chan string, 1), send2 := make(chan string, 1)
	// Use select to try sending "Message1" to send1 or "Message2" to send2
	// Both channels are buffered, so one will be chosen
	// Then read from both channels and print what was sent where
	fmt.Println("\n=== Select for sending ===")
	send1 := make(chan string, 1)
	send2 := make(chan string, 1)
	
	select {
	case send1 <- "Message1":
		fmt.Println("Sent 'Message1' to send1")
	case send2 <- "Message2":
		fmt.Println("Sent 'Message2' to send2")
	}
	
	// Check what was actually sent
	select {
	case msg := <-send1:
		fmt.Println("Retrieved from send1:", msg)
	default:
		fmt.Println("send1 is empty")
	}
	
	select {
	case msg := <-send2:
		fmt.Println("Retrieved from send2:", msg)
	default:
		fmt.Println("send2 is empty")
	}
}