package main

import "fmt"

// Define a function called 'greetPerson' that takes a string parameter called 'name'
// The function should print "Hello, [name]!" where [name] is the parameter value
func greetPerson(name string) {
	fmt.Println("Hello, " + name + "!")
}

// Define a function called 'add' that takes two int parameters and prints their sum
// The function should print "The sum is: [result]"
func add(a int, b int) {
	result := a + b
	fmt.Println("The sum is:", result)
}

func main() {
	// Call greetPerson with the name "Alice"
	greetPerson("Alice")
	
	// Call greetPerson with the name "Bob"
	greetPerson("Bob")
	
	// Call add with the numbers 5 and 3
	add(5, 3)
	
	// Call add with the numbers 10 and 7
	add(10, 7)
}