package main

import "fmt"

func main() {
	age := 20
	temperature := 25
	score := 85
	
	// Write an if statement to check if age is >= 18
	// If true, print "You are an adult"
	if age >= 18 {
		fmt.Println("You are an adult")
	}
	
	// Write an if-else statement to check if temperature > 30
	// If true, print "It's hot outside"
	// If false, print "It's not hot outside"
	if temperature > 30 {
		fmt.Println("It's hot outside")
	} else {
		fmt.Println("It's not hot outside")
	}
	
	// Write an if-else if-else chain for score:
	// score >= 90: print "Grade: A"  
	// score >= 80: print "Grade: B"
	// score >= 70: print "Grade: C"
	// score >= 60: print "Grade: D"
	// otherwise: print "Grade: F"
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else if score >= 60 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}
	
	// Write an if statement with initialization
	// Initialize x := 10 + 5 in the if condition
	// Check if x > 12, if true print "x is greater than 12"
	if x := 10 + 5; x > 12 {
		fmt.Println("x is greater than 12")
	}
	
	// Write a nested if statement
	// First check if age >= 16, if true:
	//   Then check if age >= 21, if true print "Can drink alcohol"
	//   Otherwise print "Can drive but not drink alcohol"
	// If the first condition is false, print "Too young to drive"
	if age >= 16 {
		if age >= 21 {
			fmt.Println("Can drink alcohol")
		} else {
			fmt.Println("Can drive but not drink alcohol")
		}
	} else {
		fmt.Println("Too young to drive")
	}
}