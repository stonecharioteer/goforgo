// json_basics.go
// Learn JSON encoding and decoding in Go

package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// TODO: Define Person struct with JSON tags
type Person struct {
	// Add fields: Name (string), Age (int), Email (string), IsActive (bool), Birthday (time.Time), Tags ([]string)
	// Use appropriate JSON tags: "name", "age", "email", "is_active", "birthday", "tags,omitempty"
}

// TODO: Define Book struct with custom JSON field names
type Book struct {
	// Add fields: ID (int), Title (string), Author (string), Price (float64), InStock (bool), Categories ([]string)
	// Use JSON tags: "id", "title", "author", "price", "in_stock", "categories"
}

func main() {
	fmt.Println("=== JSON Basics ===")
	
	// Demonstrate struct to JSON
	demonstrateMarshaling()
	
	// Demonstrate JSON to struct
	demonstrateUnmarshaling()
	
	// Demonstrate working with maps
	demonstrateMapJSON()
	
	// Demonstrate array/slice JSON
	demonstrateArrayJSON()
}

func demonstrateMarshaling() {
	fmt.Println("\n=== Marshaling (Go to JSON) ===")
	
	// TODO: Create a Person instance
	person := Person{
		// Fill in with sample data
		// Name: "Alice Johnson", Age: 30, Email: "alice@example.com", etc.
	}
	
	// TODO: Marshal to JSON
	jsonData, err := /* marshal person to JSON */
	if err != nil {
		fmt.Printf("Error marshaling: %v\n", err)
		return
	}
	
	fmt.Printf("Person as JSON: %s\n", jsonData)
	
	// TODO: Marshal with indentation for pretty printing
	prettyJSON, err := /* marshal person with indent using 2 spaces */
	if err != nil {
		fmt.Printf("Error marshaling with indent: %v\n", err)
		return
	}
	
	fmt.Printf("Pretty JSON:\n%s\n", prettyJSON)
	
	// TODO: Marshal a Book
	book := Book{
		// Fill in with sample data
	}
	
	bookJSON, err := /* marshal book to JSON */
	if err != nil {
		fmt.Printf("Error marshaling book: %v\n", err)
		return
	}
	
	fmt.Printf("Book as JSON: %s\n", bookJSON)
}

func demonstrateUnmarshaling() {
	fmt.Println("\n=== Unmarshaling (JSON to Go) ===")
	
	// Sample JSON data
	personJSON := `{
		"name": "Bob Smith",
		"age": 25,
		"email": "bob@example.com",
		"is_active": true,
		"birthday": "1998-05-15T00:00:00Z",
		"tags": ["developer", "golang", "backend"]
	}`
	
	// TODO: Unmarshal JSON to Person struct
	var person Person
	err := /* unmarshal personJSON into person */
	if err != nil {
		fmt.Printf("Error unmarshaling: %v\n", err)
		return
	}
	
	fmt.Printf("Unmarshaled person: %+v\n", person)
	fmt.Printf("Person name: %s\n", person.Name)
	fmt.Printf("Person age: %d\n", person.Age)
	fmt.Printf("Person tags: %v\n", person.Tags)
	
	// TODO: Handle malformed JSON
	badJSON := `{"name": "Invalid", "age": "not a number"}`
	var badPerson Person
	err = /* try to unmarshal badJSON */
	if err != nil {
		fmt.Printf("Expected error with bad JSON: %v\n", err)
	}
}

func demonstrateMapJSON() {
	fmt.Println("\n=== Working with Maps and JSON ===")
	
	// TODO: Create a map and marshal to JSON
	data := map[string]interface{}{
		"message":   "Hello, JSON!",
		"count":     42,
		"active":    true,
		"items":     []string{"apple", "banana", "orange"},
		"timestamp": time.Now().Unix(),
	}
	
	mapJSON, err := /* marshal data map to JSON with indent */
	if err != nil {
		fmt.Printf("Error marshaling map: %v\n", err)
		return
	}
	
	fmt.Printf("Map as JSON:\n%s\n", mapJSON)
	
	// TODO: Unmarshal JSON to map
	jsonString := `{
		"service": "user-api",
		"version": "1.2.3",
		"port": 8080,
		"features": ["auth", "logging", "metrics"],
		"config": {
			"debug": true,
			"timeout": 30
		}
	}`
	
	var result map[string]interface{}
	err = /* unmarshal jsonString into result */
	if err != nil {
		fmt.Printf("Error unmarshaling to map: %v\n", err)
		return
	}
	
	fmt.Printf("Unmarshaled map: %+v\n", result)
	
	// TODO: Access nested values
	if config, ok := result["config"].(map[string]interface{}); ok {
		if debug, ok := config["debug"].(bool); ok {
			fmt.Printf("Debug mode: %t\n", debug)
		}
	}
}

func demonstrateArrayJSON() {
	fmt.Println("\n=== Working with Arrays/Slices ===")
	
	// TODO: Create slice of structs
	people := []Person{
		// Add 2-3 sample Person instances
	}
	
	// TODO: Marshal slice to JSON
	arrayJSON, err := /* marshal people slice with indent */
	if err != nil {
		fmt.Printf("Error marshaling array: %v\n", err)
		return
	}
	
	fmt.Printf("People array as JSON:\n%s\n", arrayJSON)
	
	// TODO: Unmarshal JSON array back to slice
	var unmarshaledPeople []Person
	err = /* unmarshal arrayJSON back to slice */
	if err != nil {
		fmt.Printf("Error unmarshaling array: %v\n", err)
		return
	}
	
	fmt.Printf("Unmarshaled %d people:\n", len(unmarshaledPeople))
	for i, person := range unmarshaledPeople {
		fmt.Printf("  %d. %s (age %d)\n", i+1, person.Name, person.Age)
	}
}

// TODO: Bonus function - demonstrate handling unknown fields
func demonstrateUnknownFields() {
	fmt.Println("\n=== Handling Unknown Fields ===")
	
	// JSON with extra fields not in our struct
	jsonWithExtra := `{
		"name": "Charlie Brown",
		"age": 35,
		"email": "charlie@example.com",
		"is_active": true,
		"birthday": "1988-01-01T00:00:00Z",
		"unknown_field": "this will be ignored",
		"extra_data": {"nested": "value"}
	}`
	
	var person Person
	/* unmarshal JSON - unknown fields will be ignored */
	
	fmt.Printf("Person (unknown fields ignored): %+v\n", person)
}