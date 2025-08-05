// struct_basics.go
// Learn the fundamentals of structs in Go
// Structs are custom types that group related data together

package main

import "fmt"

// TODO: Define a Person struct with Name (string) and Age (int) fields
type Person struct {
	// Define fields here
}

// TODO: Define a Book struct with Title, Author (strings) and Pages (int)
type Book // Complete the struct definition

// TODO: Define an Address struct for the next example
type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

// TODO: Define a Student struct that embeds Address
type Student struct {
	Name    string
	ID      int
	// Embed Address struct here
}

func main() {
	// TODO: Create a Person using struct literal with field names
	person1 := // Create Person with Name: "Alice", Age: 30
	
	// TODO: Create a Person using struct literal without field names (positional)
	person2 := // Create Person("Bob", 25)
	
	// TODO: Create a zero-value Person and set fields individually
	var person3 Person
	// Set Name to "Charlie" and Age to 35
	
	// TODO: Access struct fields
	fmt.Printf("Person1: Name=%s, Age=%d\n", /* access fields */)
	
	// TODO: Create a Book
	book := // Create Book with Title: "The Go Programming Language", Author: "Kernighan & Ritchie", Pages: 380
	
	// TODO: Modify a struct field
	// Increase book pages by 20
	
	// TODO: Create a Student with embedded Address
	student := // Create Student with Name: "Diana", ID: 12345, and Address fields
	
	// TODO: Access embedded struct fields directly
	fmt.Printf("Student lives in: %s, %s\n", /* access city and state */)
	
	// TODO: Access embedded struct as a whole
	fmt.Printf("Full address: %+v\n", /* access the Address struct */)
	
	// TODO: Compare structs
	person4 := Person{Name: "Alice", Age: 30}
	if /* compare person1 and person4 */ {
		fmt.Println("person1 and person4 are equal")
	}
	
	// TODO: Copy a struct (creates a new copy, not a reference)
	personCopy := person1
	personCopy.Age = 31
	
	fmt.Printf("Original: %+v\n", person1)
	fmt.Printf("Copy: %+v\n", personCopy)
	
	// Print all results
	fmt.Printf("Book: %+v\n", book)
	fmt.Printf("Student: %+v\n", student)
}