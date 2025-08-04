package main

import "fmt"

func main() {
	day := 3
	grade := 'B'
	score := 85
	
	// TODO: Write a basic switch statement for day (int)
	// 1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday"
	// 6: "Saturday", 7: "Sunday", default: "Invalid day"
	
	// TODO: Write a switch statement with multiple values per case for grade (rune)
	// 'A', 'a': "Excellence"
	// 'B', 'b': "Good" 
	// 'C', 'c': "Average"
	// 'D', 'd', 'F', 'f': "Needs improvement"
	// default: "Invalid grade"
	
	// TODO: Write a switch statement with no expression (works like if-else if)
	// Check score ranges:
	// score >= 90: "Outstanding"
	// score >= 80: "Very Good"  
	// score >= 70: "Good"
	// score >= 60: "Satisfactory"
	// default: "Needs work"
	
	// TODO: Write a switch with fallthrough
	// For day value:
	// case 1: print "Start of work week" and fallthrough to next case
	// case 2, 3, 4: print "Weekday"
	// case 5: print "TGIF!"
	// case 6, 7: print "Weekend"
	
	// TODO: Write a switch with initialization
	// Initialize x := score / 10 in the switch
	// Use x for cases: 10,9: "A", 8: "B", 7: "C", 6: "D", default: "F"
}