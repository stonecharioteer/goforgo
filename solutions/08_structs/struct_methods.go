// struct_methods.go - SOLUTION
// Learn how to add methods to structs in Go

package main

import (
	"fmt"
	"math"
)

// Rectangle struct with Width and Height fields
type Rectangle struct {
	Width  float64
	Height float64
}

// Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle struct with Radius field
type Circle struct {
	Radius float64
}

// Area method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// BankAccount struct
type BankAccount struct {
	AccountNumber string
	Balance       float64
}

// Deposit method that increases the balance
// Uses a pointer receiver to modify the original struct
func (ba *BankAccount) Deposit(amount float64) {
	ba.Balance += amount
}

// Withdraw method that decreases the balance
// Returns true if successful, false if insufficient funds
func (ba *BankAccount) Withdraw(amount float64) bool {
	if ba.Balance >= amount {
		ba.Balance -= amount
		return true
	}
	return false
}

// GetBalance method (value receiver is fine here)
func (ba BankAccount) GetBalance() float64 {
	return ba.Balance
}

// String method to customize how BankAccount prints
// This implements the fmt.Stringer interface
func (ba BankAccount) String() string {
	return fmt.Sprintf("Account %s: $%.2f", ba.AccountNumber, ba.Balance)
}

func main() {
	// Create a Rectangle and call its methods
	rect := Rectangle{Width: 5, Height: 3}
	
	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area())
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter())
	
	// Create a Circle and call its method
	circle := Circle{Radius: 4}
	
	fmt.Printf("Circle area: %.2f\n", circle.Area())
	
	// Create a BankAccount and perform operations
	account := BankAccount{AccountNumber: "12345", Balance: 1000.0}
	
	fmt.Println("Initial account:", account) // This will use our String method
	
	// Make a deposit
	account.Deposit(250)
	
	fmt.Printf("After deposit: Balance = %.2f\n", account.GetBalance())
	
	// Try to withdraw money
	success := account.Withdraw(500)
	if success {
		fmt.Println("Withdrawal successful")
	} else {
		fmt.Println("Insufficient funds")
	}
	
	// Try to withdraw too much money
	success = account.Withdraw(2000)
	if !success {
		fmt.Println("Cannot withdraw $2000 - insufficient funds")
	}
	
	fmt.Println("Final account:", account)
}