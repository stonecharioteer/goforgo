package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	fruits := []string{"apple", "banana", "cherry"}
	message := "Hello"
	scores := map[string]int{"Alice": 95, "Bob": 87, "Charlie": 92}
	
	// Use range to iterate over numbers slice
	// Print both index and value: "Index 0: 10"
	fmt.Println("Numbers with index and value:")
	for index, value := range numbers {
		fmt.Printf("Index %d: %d\n", index, value)
	}
	
	// Use range to iterate over fruits slice, but only use the values (ignore index)
	// Print: "Fruit: apple"
	fmt.Println("\nFruits (values only):")
	for _, fruit := range fruits {
		fmt.Println("Fruit:", fruit)
	}
	
	// Use range to iterate over fruits slice, but only use the index (ignore value)
	// Print: "Index: 0"
	fmt.Println("\nFruits (index only):")
	for index := range fruits {
		fmt.Println("Index:", index)
	}
	
	// Use range to iterate over the message string
	// Print each character with its byte position: "Position 0: H"
	fmt.Println("\nString characters:")
	for pos, char := range message {
		fmt.Printf("Position %d: %c\n", pos, char)
	}
	
	// Use range to iterate over the scores map
	// Print: "Alice scored 95"
	fmt.Println("\nScores (key and value):")
	for name, score := range scores {
		fmt.Printf("%s scored %d\n", name, score)
	}
	
	// Use range to iterate over just the keys of scores map
	// Print: "Student: Alice"
	fmt.Println("\nStudents (keys only):")
	for name := range scores {
		fmt.Println("Student:", name)
	}
	
	// Use range to iterate over just the values of scores map  
	// Print: "Score: 95"
	fmt.Println("\nScores (values only):")
	for _, score := range scores {
		fmt.Println("Score:", score)
	}
	
	// Create a channel and use range to iterate over it
	// Create: ch := make(chan int, 3)
	// Send values: ch <- 1, ch <- 2, ch <- 3, then close(ch)
	// Use range to receive: "Received: 1"
	fmt.Println("\nChannel values:")
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	
	for value := range ch {
		fmt.Println("Received:", value)
	}
}