// struct_basics.go - SOLUTION
// Learn the fundamentals of structs in Go
// Structs are custom types that group related data together

package main

import "fmt"

// Define a Person struct with Name (string) and Age (int) fields
type Person struct {
	Name string
	Age  int
}

// Define a Book struct with Title, Author (strings) and Pages (int)
type Book struct {
	Title  string
	Author string
	Pages  int
}

// Define an Address struct for the next example
type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

// Define a Student struct that embeds Address
type Student struct {
	Name string
	ID   int
	Address // Embedded struct
}

func main() {
	// Create a Person using struct literal with field names
	person1 := Person{Name: "Alice", Age: 30}
	
	// Create a Person using struct literal without field names (positional)
	person2 := Person{"Bob", 25}
	
	// Create a zero-value Person and set fields individually
	var person3 Person
	person3.Name = "Charlie"
	person3.Age = 35
	
	// Access struct fields
	fmt.Printf("Person1: Name=%s, Age=%d\n", person1.Name, person1.Age)
	
	// Create a Book
	book := Book{
		Title:  "The Go Programming Language",
		Author: "Kernighan & Ritchie",
		Pages:  380,
	}
	
	// Modify a struct field
	book.Pages += 20
	
	// Create a Student with embedded Address
	student := Student{
		Name: "Diana",
		ID:   12345,
		Address: Address{
			Street: "123 Main St",
			City:   "Boston",
			State:  "MA",
			Zip:    "02101",
		},
	}
	
	// Access embedded struct fields directly
	fmt.Printf("Student lives in: %s, %s\n", student.City, student.State)
	
	// Access embedded struct as a whole
	fmt.Printf("Full address: %+v\n", student.Address)
	
	// Compare structs
	person4 := Person{Name: "Alice", Age: 30}
	if person1 == person4 {
		fmt.Println("person1 and person4 are equal")
	}
	
	// Copy a struct (creates a new copy, not a reference)
	personCopy := person1
	personCopy.Age = 31
	
	fmt.Printf("Original: %+v\n", person1)
	fmt.Printf("Copy: %+v\n", personCopy)
	
	// Print all results
	fmt.Printf("Person2: %+v\n", person2)
	fmt.Printf("Person3: %+v\n", person3)
	fmt.Printf("Book: %+v\n", book)
	fmt.Printf("Student: %+v\n", student)
}