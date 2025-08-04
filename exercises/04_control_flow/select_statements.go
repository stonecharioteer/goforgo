package main

import (
	"fmt"
	"time"
)

func main() {
	// TODO: Create two channels: ch1 := make(chan string), ch2 := make(chan string)
	
	// TODO: Start a goroutine that sends "Hello from ch1" to ch1 after 1 second
	// Use: go func() { time.Sleep(time.Second); ch1 <- "Hello from ch1" }()
	
	// TODO: Start a goroutine that sends "Hello from ch2" to ch2 after 2 seconds
	
	// TODO: Use select to receive from either channel
	// Print which channel sent the message and what message was received
	
	// TODO: Demonstrate select with default case
	// Create a channel: ch3 := make(chan int)
	// Use select to try to receive from ch3 with a default case
	// Default case should print "No data available"
	
	// TODO: Demonstrate select with timeout using time.After
	// Create a channel: slow := make(chan string)
	// Start a goroutine that sends to slow after 3 seconds
	// Use select with slow and time.After(1*time.Second)
	// Print appropriate messages for both cases
	
	// TODO: Demonstrate select with multiple ready channels
	// Create channels: fast1 := make(chan int, 1), fast2 := make(chan int, 1)
	// Send values immediately: fast1 <- 1, fast2 <- 2
	// Use select to receive from either (Go will choose randomly)
	// Run this in a loop 5 times to see random selection
	
	// TODO: Demonstrate select for sending (not just receiving)
	// Create: send1 := make(chan string, 1), send2 := make(chan string, 1)
	// Use select to try sending "Message1" to send1 or "Message2" to send2
	// Both channels are buffered, so one will be chosen
	// Then read from both channels and print what was sent where
}