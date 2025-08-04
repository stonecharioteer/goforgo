package main

import (
	"fmt"
	"strconv"
)

func main() {
	// These variables have different types
	var number int = 42
	var pi float64 = 3.14159
	var message string = "123"
	
	floatNumber := float64(number)
	intPi := int(pi)
	convertedNumber, _ := strconv.Atoi(message)
	
	fmt.Printf("Float: %.2f, Int: %d, Converted: %d\n", floatNumber, intPi, convertedNumber)
}
