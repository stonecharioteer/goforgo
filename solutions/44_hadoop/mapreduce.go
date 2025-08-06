// GoForGo Solution: Hadoop MapReduce
package main

import (
	"fmt"
	"strings"
	"sync"
)

type MapReduceJob struct {
	mappers  int
	reducers int
}

func NewMapReduceJob() *MapReduceJob {
	return &MapReduceJob{mappers: 4, reducers: 2}
}

func (mr *MapReduceJob) wordCount(input []string) map[string]int {
	// Map phase: split into words
	wordChan := make(chan map[string]int, len(input))
	var wg sync.WaitGroup

	for _, line := range input {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			words := strings.Fields(text)
			wordMap := make(map[string]int)
			for _, word := range words {
				word = strings.ToLower(strings.Trim(word, ".,!?"))
				wordMap[word]++
			}
			wordChan <- wordMap
		}(line)
	}

	go func() {
		wg.Wait()
		close(wordChan)
	}()

	// Reduce phase: combine results
	result := make(map[string]int)
	for wordMap := range wordChan {
		for word, count := range wordMap {
			result[word] += count
		}
	}

	return result
}

func main() {
	mr := NewMapReduceJob()
	
	input := []string{
		"Hello world hello",
		"World of Go programming",
		"Go is great for hadoop integration",
	}

	results := mr.wordCount(input)
	fmt.Println("Word count results:", results)
	fmt.Println("Hadoop MapReduce operations completed!")
}