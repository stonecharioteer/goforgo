package main

import "fmt"

// Define a struct called 'Counter' with a count field (int)
type Counter struct {
	count int
}

// Define a method called 'Increment' with a pointer receiver on Counter
// The method should increase the count by 1
func (c *Counter) Increment() {
	c.count++
}

// Define a method called 'Value' with a value receiver on Counter  
// The method should return the current count
func (c Counter) Value() int {
	return c.count
}

// Define a method called 'Reset' with a pointer receiver on Counter
// The method should set count back to 0
func (c *Counter) Reset() {
	c.count = 0
}

// Define a struct called 'BankAccount' with balance field (float64)
type BankAccount struct {
	balance float64
}

// Define a method called 'Deposit' with a pointer receiver on BankAccount
// The method should take an amount (float64) and add it to the balance
func (b *BankAccount) Deposit(amount float64) {
	b.balance += amount
}

// Define a method called 'Withdraw' with a pointer receiver on BankAccount  
// The method should take an amount (float64) and subtract it from the balance
// Only withdraw if there are sufficient funds (balance >= amount)
func (b *BankAccount) Withdraw(amount float64) {
	if b.balance >= amount {
		b.balance -= amount
		fmt.Printf("Withdrew %.2f\n", amount)
	} else {
		fmt.Printf("Insufficient funds to withdraw %.2f\n", amount)
	}
}

// Define a method called 'Balance' with a value receiver on BankAccount
// The method should return the current balance
func (b BankAccount) Balance() float64 {
	return b.balance
}

func main() {
	// Create a Counter with initial count of 0
	counter := Counter{count: 0}
	
	// Print the initial value
	fmt.Println("Initial counter value:", counter.Value())
	
	// Call Increment 3 times
	counter.Increment()
	counter.Increment()
	counter.Increment()
	
	// Print the value after increments
	fmt.Println("Counter after 3 increments:", counter.Value())
	
	// Call Reset
	counter.Reset()
	
	// Print the value after reset
	fmt.Println("Counter after reset:", counter.Value())
	
	// Create a BankAccount with initial balance of 100.0
	account := BankAccount{balance: 100.0}
	
	// Print the initial balance
	fmt.Printf("Initial balance: %.2f\n", account.Balance())
	
	// Deposit 50.0
	account.Deposit(50.0)
	
	// Print balance after deposit
	fmt.Printf("Balance after deposit: %.2f\n", account.Balance())
	
	// Withdraw 30.0  
	account.Withdraw(30.0)
	
	// Print balance after withdrawal
	fmt.Printf("Balance after withdrawal: %.2f\n", account.Balance())
	
	// Try to withdraw 200.0 (should fail)
	account.Withdraw(200.0)
	
	// Print final balance
	fmt.Printf("Final balance: %.2f\n", account.Balance())
}