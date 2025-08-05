// json_basics.go
// Learn JSON encoding and decoding in Go

package main

import (
	"encoding/json"
	"fmt"
)

// TODO: Define structs for JSON operations
type Person struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Active   bool   `json:"active"`
	Hobbies  []string `json:"hobbies,omitempty"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

type Employee struct {
	ID      int     `json:"id"`
	Person  Person  `json:"person"`
	Address Address `json:"address"`
	Salary  float64 `json:"salary"`
}

func main() {
	fmt.Println("=== JSON Marshaling (Encoding) ===")
	
	// TODO: Create a Person and marshal to JSON
	person := Person{
		Name:    "Alice Johnson",
		Age:     28,
		Email:   "alice@example.com",
		Active:  true,
		Hobbies: []string{"reading", "hiking", "photography"},
	}
	
	// TODO: Marshal person to JSON
	personJSON, err := /* marshal person to JSON */
	if err != nil {
		fmt.Printf("Marshal error: %v\\n", err)
		return
	}
	
	fmt.Printf("Person JSON: %s\\n", personJSON)
	
	// TODO: Pretty print JSON
	prettyJSON, err := /* marshal person with indentation */
	if err != nil {
		fmt.Printf("Pretty marshal error: %v\\n", err)
	} else {
		fmt.Printf("Pretty JSON:\\n%s\\n", prettyJSON)
	}
	
	fmt.Println("\\n=== JSON Unmarshaling (Decoding) ===")
	
	// TODO: Unmarshal JSON string to struct
	jsonStr := `{
		"name": "Bob Smith",
		"age": 35,
		"email": "bob@example.com",
		"active": false,
		"hobbies": ["gaming", "cooking"]
	}`
	
	var newPerson Person
	err = /* unmarshal jsonStr into newPerson */
	if err != nil {
		fmt.Printf("Unmarshal error: %v\\n", err)
	} else {
		fmt.Printf("Unmarshaled person: %+v\\n", newPerson)
	}
	
	fmt.Println("\\n=== Working with Maps ===")
	
	// TODO: Marshal map to JSON
	data := map[string]interface{}{
		"name":    "Charlie",
		"age":     42,
		"married": true,
		"children": []string{"Emma", "Liam"},
		"address": map[string]string{
			"city":  "New York",
			"state": "NY",
		},
	}
	
	mapJSON, err := /* marshal data map to JSON */
	if err != nil {
		fmt.Printf("Map marshal error: %v\\n", err)
	} else {
		fmt.Printf("Map JSON: %s\\n", mapJSON)
	}
	
	// TODO: Unmarshal JSON to map
	var result map[string]interface{}
	err = /* unmarshal mapJSON into result */
	if err != nil {
		fmt.Printf("Map unmarshal error: %v\\n", err)
	} else {
		fmt.Printf("Unmarshaled map: %+v\\n", result)
		
		// Access nested values
		if addr, ok := /* get "address" from result and assert as map */; ok {
			if city, ok := /* get "city" from addr and assert as string */; ok {
				fmt.Printf("City: %s\\n", city)
			}
		}
	}
	
	fmt.Println("\\n=== Complex Nested Structures ===")
	
	// TODO: Create complex employee structure
	employee := Employee{
		ID: 1001,
		Person: Person{
			Name:    "Diana Prince",
			Age:     30,
			Email:   "diana@wayneenterprises.com",
			Active:  true,
			Hobbies: []string{"martial arts", "archaeology"},
		},
		Address: Address{
			Street:  "123 Justice League Ave",
			City:    "Metropolis",
			State:   "NY",
			ZipCode: "10001",
		},
		Salary: 95000.50,
	}
	
	// TODO: Marshal complex structure
	empJSON, err := /* marshal employee with pretty printing */
	if err != nil {
		fmt.Printf("Employee marshal error: %v\\n", err)
	} else {
		fmt.Printf("Employee JSON:\\n%s\\n", empJSON)
	}
	
	// TODO: Unmarshal back to struct
	var newEmployee Employee
	err = /* unmarshal empJSON into newEmployee */
	if err != nil {
		fmt.Printf("Employee unmarshal error: %v\\n", err)
	} else {
		fmt.Printf("Unmarshaled employee: %+v\\n", newEmployee)
		fmt.Printf("Employee name: %s\\n", newEmployee.Person.Name)
		fmt.Printf("Employee city: %s\\n", newEmployee.Address.City)
	}
	
	fmt.Println("\\n=== Handling Arrays/Slices ===")
	
	// TODO: Work with slice of structs
	people := []Person{
		{Name: "Alice", Age: 25, Email: "alice@test.com", Active: true},
		{Name: "Bob", Age: 30, Email: "bob@test.com", Active: false},
		{Name: "Charlie", Age: 35, Email: "charlie@test.com", Active: true},
	}
	
	// TODO: Marshal slice to JSON
	peopleJSON, err := /* marshal people slice to JSON */
	if err != nil {
		fmt.Printf("People marshal error: %v\\n", err)
	} else {
		fmt.Printf("People JSON: %s\\n", peopleJSON)
	}
	
	// TODO: Unmarshal JSON array back to slice
	var newPeople []Person
	err = /* unmarshal peopleJSON into newPeople */
	if err != nil {
		fmt.Printf("People unmarshal error: %v\\n", err)
	} else {
		fmt.Printf("Unmarshaled %d people:\\n", len(newPeople))
		for i, p := range newPeople {
			fmt.Printf("  %d: %s (age %d)\\n", i+1, p.Name, p.Age)
		}
	}
}