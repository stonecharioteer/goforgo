package main

import "fmt"

// Define a struct called 'Person' with name (string) and age (int) fields
type Person struct {
	name string
	age  int
}

// Define a method called 'String' with a value receiver on Person
// It should return a string in the format: "Name: [name], Age: [age]"
func (p Person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.name, p.age)
}

// Define a method called 'HaveBirthday' with a pointer receiver on Person  
// It should increment the age by 1 and print "[name] is now [age] years old!"
func (p *Person) HaveBirthday() {
	p.age++
	fmt.Printf("%s is now %d years old!\n", p.name, p.age)
}

// Define a method called 'IsAdult' with a value receiver on Person
// It should return true if age >= 18, false otherwise
func (p Person) IsAdult() bool {
	return p.age >= 18
}

func main() {
	// Create a Person with name "Alice" and age 17
	alice := Person{name: "Alice", age: 17}
	
	// Print the person using the String method
	fmt.Println(alice.String())
	
	// Check if the person is an adult and print the result
	fmt.Println("Is Alice an adult?", alice.IsAdult())
	
	// Call HaveBirthday method
	alice.HaveBirthday()
	
	// Print the person again using the String method
	fmt.Println(alice.String())
	
	// Check if the person is an adult now and print the result
	fmt.Println("Is Alice an adult now?", alice.IsAdult())
	
	// Create a pointer to a Person with name "Bob" and age 25
	// Use: person := &Person{name: "Bob", age: 25}
	bob := &Person{name: "Bob", age: 25}
	
	// Call all methods on the pointer and observe that it works the same
	// Print the person, check if adult, have birthday, print again
	fmt.Println(bob.String())
	fmt.Println("Is Bob an adult?", bob.IsAdult())
	bob.HaveBirthday()
	fmt.Println(bob.String())
}