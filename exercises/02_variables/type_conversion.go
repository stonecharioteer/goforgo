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
	
	// TODO: Convert number to float64 and assign to floatNumber
	// TODO: Convert pi to int and assign to intPi
	// TODO: Convert message to an actual number using strconv.Atoi() and assign to convertedNumber
	
	fmt.Printf("Float: %.2f, Int: %d, Converted: %d\n", floatNumber, intPi, convertedNumber)
}
