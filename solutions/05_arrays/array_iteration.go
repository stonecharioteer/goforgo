// array_iteration.go - SOLUTION
// Learn different ways to iterate over arrays in Go

package main

import "fmt"

func main() {
	scores := [5]int{85, 92, 78, 96, 88}
	names := [...]string{"Alice", "Bob", "Charlie", "Diana", "Eve"}

	// Use a traditional for loop to print all scores
	fmt.Println("Scores using traditional for loop:")
	for i := 0; i < len(scores); i++ {
		fmt.Printf("Score %d: %d\n", i+1, scores[i])
	}

	// Use range to iterate over the names array
	// Print both index and value
	fmt.Println("\nNames with indices using range:")
	for index, value := range names {
		fmt.Printf("Index %d: %s\n", index, value)
	}

	// Use range to iterate over scores array
	// but only use the values (ignore the index)
	fmt.Println("\nScores using range (values only):")
	for _, value := range scores {
		fmt.Printf("Score: %d\n", value)
	}

	// Find the maximum score using iteration
	maxScore := scores[0]
	for i := 1; i < len(scores); i++ {
		if scores[i] > maxScore {
			maxScore = scores[i]
		}
	}
	fmt.Println("\nMaximum score:", maxScore)

	// Calculate the average score
	var sum int
	for _, score := range scores {
		sum += score
	}
	average := float64(sum) / float64(len(scores))
	fmt.Printf("Average score: %.2f\n", average)
}