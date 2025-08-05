// struct_methods.go
// Learn how to add methods to structs in Go

package main

import (
	"fmt"
	"math"
)

// TODO: Define a Rectangle struct with Width and Height fields
type Rectangle struct {
	// Define fields here
}

// TODO: Define an Area method for Rectangle
// Hint: func (receiver ReceiverType) MethodName() ReturnType
func (r Rectangle) Area() float64 {
	// Calculate and return area
}

// TODO: Define a Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	// Calculate and return perimeter
}

// TODO: Define a Circle struct with Radius field
type Circle struct {
	// Define field here
}

// TODO: Define an Area method for Circle
func (c Circle) Area() float64 {
	// Calculate and return area (π * r²)
}

// TODO: Define a BankAccount struct
type BankAccount struct {
	AccountNumber string
	Balance       float64
}

// TODO: Define a Deposit method that increases the balance
// This should use a pointer receiver to modify the original struct
func (ba *BankAccount) Deposit(amount float64) {
	// Add amount to balance
}

// TODO: Define a Withdraw method that decreases the balance
// Return true if successful, false if insufficient funds
func (ba *BankAccount) Withdraw(amount float64) bool {
	// Check if sufficient funds, then withdraw
}

// TODO: Define a GetBalance method (value receiver is fine here)
func (ba BankAccount) GetBalance() float64 {
	// Return current balance
}

// TODO: Define a String method to customize how BankAccount prints
// This implements the fmt.Stringer interface
func (ba BankAccount) String() string {
	// Return formatted string representation
}

func main() {
	// TODO: Create a Rectangle and call its methods
	rect := // Create Rectangle with Width: 5, Height: 3
	
	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", /* call Area method */)
	fmt.Printf("Perimeter: %.2f\n", /* call Perimeter method */)
	
	// TODO: Create a Circle and call its method
	circle := // Create Circle with Radius: 4
	
	fmt.Printf("Circle area: %.2f\n", /* call Area method */)
	
	// TODO: Create a BankAccount and perform operations
	account := // Create BankAccount with AccountNumber: "12345", Balance: 1000.0
	
	fmt.Println("Initial account:", account) // This will use our String method
	
	// TODO: Make a deposit
	// Deposit $250
	
	fmt.Printf("After deposit: Balance = %.2f\n", /* get balance */)
	
	// TODO: Try to withdraw money
	success := // Try to withdraw $500
	if success {
		fmt.Println("Withdrawal successful")
	} else {
		fmt.Println("Insufficient funds")
	}
	
	// TODO: Try to withdraw too much money
	success = // Try to withdraw $2000
	if !success {
		fmt.Println("Cannot withdraw $2000 - insufficient funds")
	}
	
	fmt.Println("Final account:", account)
}