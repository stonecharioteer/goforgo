package main

import "fmt"

// Define a function called 'calculateArea' that takes width and height (both float64)
// Use named return values 'area' (float64) and 'perimeter' (float64)
// Calculate area = width * height and perimeter = 2 * (width + height)
// Use a naked return statement
func calculateArea(width, height float64) (area, perimeter float64) {
	area = width * height
	perimeter = 2 * (width + height)
	return // naked return
}

// Define a function called 'swapStrings' that takes two string parameters a and b
// Use named return values 'first' and 'second' (both strings)
// Return b as first and a as second (effectively swapping them)
// Use a naked return statement
func swapStrings(a, b string) (first, second string) {
	first = b
	second = a
	return // naked return
}

func main() {
	// Call calculateArea with width=5.0 and height=3.0
	// Print both the area and perimeter
	area, perimeter := calculateArea(5.0, 3.0)
	fmt.Println("Area:", area, "Perimeter:", perimeter)
	
	// Call swapStrings with "hello" and "world"
	// Print both returned values
	first, second := swapStrings("hello", "world")
	fmt.Println("First:", first, "Second:", second)
}