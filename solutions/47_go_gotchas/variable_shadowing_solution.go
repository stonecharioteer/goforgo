package main

import "fmt"

func demonstrateVariableShadowing() {
	// This function demonstrates variable shadowing gotchas
	// that are unique to Go and surprise developers from other languages
	
	name := "outer scope"
	count := 42
	
	// First gotcha: Variable shadowing with if statement
	if name := "if scope"; len(name) > 0 {
		fmt.Printf("Inside if: %s\n", name)
		// The outer 'name' variable is shadowed here - it still exists but is inaccessible
	}
	
	// Second gotcha: Range loop variable shadowing
	numbers := []int{1, 2, 3}
	for i, count := range numbers {
		// 'count' is shadowed here - not the same as outer count
		// FIXED: Use different variable name to access outer count
		fmt.Printf("Index %d, value %d, outer count %d\n", i, count, 42) // Hard-coded for clarity
	}
	// Better fix: rename the range variable to avoid confusion
	for i, val := range numbers {
		fmt.Printf("Index %d, value %d, outer count %d\n", i, val, count)
	}
	
	// Third gotcha: Short variable declaration in nested scope
	if true {
		name, err := getName()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("Inner name: %s\n", name)
		// This 'name' shadows the outer 'name'
	}
	
	// Fourth gotcha: Multiple assignment with shadowing
	value := 10
	{
		// This creates a new 'value' variable, doesn't modify the outer one
		value, message := getValue()
		fmt.Printf("Inner block - value: %d, message: %s\n", value, message)
		// The outer 'value' remains 10
	}
	
	// Final values show that outer variables weren't modified due to shadowing
	fmt.Printf("Final values - name: %s, count: %d, value: %d\n", name, count, value)
}

func getName() (string, error) {
	return "inner function", nil
}

func getValue() (int, string) {
	return 99, "from getValue"
}

func main() {
	demonstrateVariableShadowing()
	
	// Additional shadowing gotcha with error handling - FIXED
	var result string
	if data, err := processData(); err != nil {
		fmt.Printf("Error processing: %v\n", err)
	} else {
		result = data // FIXED: This works because 'data' is in scope
	}
	
	// Alternative fix using separate declaration and assignment:
	// data, err := processData()
	// if err != nil {
	//     fmt.Printf("Error processing: %v\n", err)
	// } else {
	//     result = data
	// }
	
	fmt.Printf("Result: %s\n", result)
}

func processData() (string, error) {
	return "processed data", nil
}