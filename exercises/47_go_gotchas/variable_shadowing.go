package main

import "fmt"

func demonstrateVariableShadowing() {
	// This function contains several variable shadowing gotchas
	// that are unique to Go and surprise developers from other languages
	
	name := "outer scope"
	count := 42
	
	// First gotcha: Variable shadowing with if statement
	if name := "if scope"; len(name) > 0 {
		fmt.Printf("Inside if: %s\n", name)
		// What happens to the outer 'name' variable here?
	}
	
	// Second gotcha: Range loop variable shadowing
	numbers := []int{1, 2, 3}
	for i, count := range numbers {
		// 'count' is shadowed here - not the same as outer count
		fmt.Printf("Index %d, value %d, outer count %d\n", i, count, count)
		// This line has a bug - fix it to show the outer count correctly
	}
	
	// Third gotcha: Short variable declaration in nested scope
	if true {
		name, err := getName()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("Inner name: %s\n", name)
	}
	
	// Fourth gotcha: Multiple assignment with shadowing
	value := 10
	{
		// This creates a new 'value' variable, doesn't modify the outer one
		value, message := getValue()
		fmt.Printf("Inner block - value: %d, message: %s\n", value, message)
	}
	
	// What are the final values of our variables?
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
	
	// Additional shadowing gotcha with error handling
	var result string
	if data, err := processData(); err != nil {
		fmt.Printf("Error processing: %v\n", err)
	} else {
		result = data // This line has a compilation error - fix it
	}
	
	fmt.Printf("Result: %s\n", result)
}

func processData() (string, error) {
	return "processed data", nil
}