package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	fruits := []string{"apple", "banana", "cherry"}
	message := "Hello"
	scores := map[string]int{"Alice": 95, "Bob": 87, "Charlie": 92}
	
	// TODO: Use range to iterate over numbers slice
	// Print both index and value: "Index 0: 10"
	
	// TODO: Use range to iterate over fruits slice, but only use the values (ignore index)
	// Print: "Fruit: apple"
	
	// TODO: Use range to iterate over fruits slice, but only use the index (ignore value)
	// Print: "Index: 0"
	
	// TODO: Use range to iterate over the message string
	// Print each character with its byte position: "Position 0: H"
	
	// TODO: Use range to iterate over the scores map
	// Print: "Alice scored 95"
	
	// TODO: Use range to iterate over just the keys of scores map
	// Print: "Student: Alice"
	
	// TODO: Use range to iterate over just the values of scores map  
	// Print: "Score: 95"
	
	// TODO: Create a channel and use range to iterate over it
	// Create: ch := make(chan int, 3)
	// Send values: ch <- 1, ch <- 2, ch <- 3, then close(ch)
	// Use range to receive: "Received: 1"
}