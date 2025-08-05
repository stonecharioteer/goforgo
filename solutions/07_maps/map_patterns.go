// map_patterns.go - SOLUTION
// Learn common map usage patterns in Go programming
// Practice real-world scenarios with maps

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Pattern 1 - Counting occurrences
	text := "hello world hello go programming go is awesome"
	wordCounts := make(map[string]int)
	
	// Split text into words and count each word
	words := strings.Fields(text)
	for _, word := range words {
		wordCounts[word]++
	}
	
	fmt.Println("Word counts:")
	for word, count := range wordCounts {
		fmt.Printf("  %s: %d\n", word, count)
	}

	// Pattern 2 - Grouping data
	students := []struct {
		Name  string
		Grade string
		Score int
	}{
		{"Alice", "A", 95},
		{"Bob", "B", 87},
		{"Charlie", "A", 92},
		{"Diana", "B", 89},
		{"Eve", "A", 98},
		{"Frank", "C", 75},
		{"Grace", "B", 85},
	}
	
	// Group students by grade
	gradeGroups := make(map[string][]string)
	for _, student := range students {
		gradeGroups[student.Grade] = append(gradeGroups[student.Grade], student.Name)
	}
	
	fmt.Println("\nStudents grouped by grade:")
	for grade, names := range gradeGroups {
		fmt.Printf("  Grade %s: %v\n", grade, names)
	}

	// Pattern 3 - Set operations using maps
	set1 := map[string]bool{
		"apple":  true,
		"banana": true,
		"cherry": true,
	}
	
	set2 := map[string]bool{
		"banana": true,
		"date":   true,
		"cherry": true,
	}
	
	// Find intersection (elements in both sets)
	intersection := make(map[string]bool)
	for key := range set1 {
		if set2[key] {
			intersection[key] = true
		}
	}
	
	// Find union (all elements from both sets)
	union := make(map[string]bool)
	for key := range set1 {
		union[key] = true
	}
	for key := range set2 {
		union[key] = true
	}
	
	// Find difference (elements in set1 but not in set2)
	difference := make(map[string]bool)
	for key := range set1 {
		if !set2[key] {
			difference[key] = true
		}
	}
	
	fmt.Println("\nSet operations:")
	fmt.Printf("Set1: %v\n", getKeys(set1))
	fmt.Printf("Set2: %v\n", getKeys(set2))
	fmt.Printf("Intersection: %v\n", getKeys(intersection))
	fmt.Printf("Union: %v\n", getKeys(union))
	fmt.Printf("Difference (set1 - set2): %v\n", getKeys(difference))

	// Pattern 4 - Cache/Memoization
	fmt.Println("\nFibonacci with memoization:")
	fibCache := make(map[int]int)
	
	// Calculate fibonacci numbers with caching
	for i := 0; i <= 10; i++ {
		result := fibonacciMemo(i, fibCache)
		fmt.Printf("fib(%d) = %d\n", i, result)
	}
	fmt.Printf("Cache after calculations: %v\n", fibCache)

	// Pattern 5 - Default values pattern
	fmt.Println("\nConfiguration with defaults:")
	
	userConfig := map[string]string{
		"theme": "dark",
		"lang":  "en",
	}
	
	// Get configuration value with default
	theme := getConfigValue(userConfig, "theme", "light")
	language := getConfigValue(userConfig, "lang", "en")
	timeout := getConfigValue(userConfig, "timeout", "30s")
	
	fmt.Printf("Theme: %s\n", theme)
	fmt.Printf("Language: %s\n", language)
	fmt.Printf("Timeout: %s (default used)\n", timeout)

	// Pattern 6 - Frequency analysis
	numbers := []int{1, 2, 3, 2, 1, 3, 1, 4, 5, 4, 1}
	fmt.Printf("\nNumbers: %v\n", numbers)
	
	// Find the most frequent number
	frequency := make(map[int]int)
	for _, num := range numbers {
		frequency[num]++
	}
	
	// Find most frequent
	var mostFrequent int
	var maxCount int
	for num, count := range frequency {
		if count > maxCount {
			mostFrequent = num
			maxCount = count
		}
	}
	
	fmt.Printf("Most frequent number: %d (appears %d times)\n", mostFrequent, maxCount)

	// Pattern 7 - Multi-level maps (nested maps)
	// Create a map of maps to store student grades by subject
	studentGrades := make(map[string]map[string]int)
	
	// Initialize and populate nested maps
	studentGrades["Alice"] = make(map[string]int)
	studentGrades["Bob"] = make(map[string]int)
	
	// Add grades for different subjects
	studentGrades["Alice"]["Math"] = 95
	studentGrades["Alice"]["Science"] = 92
	studentGrades["Alice"]["English"] = 88
	
	studentGrades["Bob"]["Math"] = 87
	studentGrades["Bob"]["Science"] = 90
	studentGrades["Bob"]["English"] = 85
	
	fmt.Println("\nStudent grades by subject:")
	for student, subjects := range studentGrades {
		fmt.Printf("%s:\n", student)
		for subject, grade := range subjects {
			fmt.Printf("  %s: %d\n", subject, grade)
		}
	}
}

// Helper function to get keys from a map[string]bool (set)
func getKeys(m map[string]bool) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// Fibonacci with memoization
func fibonacciMemo(n int, cache map[int]int) int {
	// Base cases
	if n <= 1 {
		return n
	}
	
	// Check cache first
	if value, exists := cache[n]; exists {
		return value
	}
	
	// Calculate and store in cache
	result := fibonacciMemo(n-1, cache) + fibonacciMemo(n-2, cache)
	cache[n] = result
	return result
}

// Get configuration value with default
func getConfigValue(config map[string]string, key, defaultValue string) string {
	if value, exists := config[key]; exists {
		return value
	}
	return defaultValue
}