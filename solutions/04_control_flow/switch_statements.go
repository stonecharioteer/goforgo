package main

import "fmt"

func main() {
	day := 3
	grade := 'B'
	score := 85
	
	// Write a basic switch statement for day (int)
	// 1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday"
	// 6: "Saturday", 7: "Sunday", default: "Invalid day"
	fmt.Print("Day: ")
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}
	
	// Write a switch statement with multiple values per case for grade (rune)
	// 'A', 'a': "Excellence"
	// 'B', 'b': "Good" 
	// 'C', 'c': "Average"
	// 'D', 'd', 'F', 'f': "Needs improvement"
	// default: "Invalid grade"
	fmt.Print("Grade performance: ")
	switch grade {
	case 'A', 'a':
		fmt.Println("Excellence")
	case 'B', 'b':
		fmt.Println("Good")
	case 'C', 'c':
		fmt.Println("Average")
	case 'D', 'd', 'F', 'f':
		fmt.Println("Needs improvement")
	default:
		fmt.Println("Invalid grade")
	}
	
	// Write a switch statement with no expression (works like if-else if)
	// Check score ranges:
	// score >= 90: "Outstanding"
	// score >= 80: "Very Good"  
	// score >= 70: "Good"
	// score >= 60: "Satisfactory"
	// default: "Needs work"
	fmt.Print("Score evaluation: ")
	switch {
	case score >= 90:
		fmt.Println("Outstanding")
	case score >= 80:
		fmt.Println("Very Good")
	case score >= 70:
		fmt.Println("Good")
	case score >= 60:
		fmt.Println("Satisfactory")
	default:
		fmt.Println("Needs work")
	}
	
	// Write a switch with fallthrough
	// For day value:
	// case 1: print "Start of work week" and fallthrough to next case
	// case 2, 3, 4: print "Weekday"
	// case 5: print "TGIF!"
	// case 6, 7: print "Weekend"
	fmt.Print("Day type: ")
	switch day {
	case 1:
		fmt.Print("Start of work week, ")
		fallthrough
	case 2, 3, 4:
		fmt.Println("Weekday")
	case 5:
		fmt.Println("TGIF!")
	case 6, 7:
		fmt.Println("Weekend")
	}
	
	// Write a switch with initialization
	// Initialize x := score / 10 in the switch
	// Use x for cases: 10,9: "A", 8: "B", 7: "C", 6: "D", default: "F"
	fmt.Print("Letter grade: ")
	switch x := score / 10; x {
	case 10, 9:
		fmt.Println("A")
	case 8:
		fmt.Println("B")
	case 7:
		fmt.Println("C")
	case 6:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}
}