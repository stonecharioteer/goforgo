// array_iteration.go
// Learn different ways to iterate over arrays in Go

package main

import "fmt"

func main() {
	scores := [5]int{85, 92, 78, 96, 88}
	names := [...]string{"Alice", "Bob", "Charlie", "Diana", "Eve"}

	// TODO: Use a traditional for loop to print all scores
	// Hint: for i := 0; i < len(array); i++ { ... }
	fmt.Println("Scores using traditional for loop:")
	// Write your loop here

	// TODO: Use range to iterate over the names array
	// Print both index and value
	// Hint: for index, value := range array { ... }
	fmt.Println("\nNames with indices using range:")
	// Write your range loop here

	// TODO: Use range to iterate over scores array
	// but only use the values (ignore the index)
	// Hint: for _, value := range array { ... }
	fmt.Println("\nScores using range (values only):")
	// Write your range loop here

	// TODO: Find the maximum score using iteration
	maxScore := // Initialize with first element
	// Write a loop to find the maximum
	fmt.Println("\nMaximum score:", maxScore)

	// TODO: Calculate the average score
	var sum int
	// Write a loop to calculate sum
	average := // Calculate average
	fmt.Printf("Average score: %.2f\n", average)
}