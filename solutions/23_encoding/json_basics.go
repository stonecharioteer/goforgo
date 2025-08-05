// json_basics.go - SOLUTION
// Learn JSON encoding and decoding in Go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Person represents a person with JSON tags
type Person struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Email    string    `json:"email"`
	IsActive bool      `json:"is_active"`
	Birthday time.Time `json:"birthday"`
	Tags     []string  `json:"tags,omitempty"`
}

// Book with custom JSON field names
type Book struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Price       float64  `json:"price"`
	InStock     bool     `json:"in_stock"`
	Categories  []string `json:"categories"`
	PublishedAt string   `json:"published_at"`
}

func main() {
	fmt.Println("=== JSON Encoding/Decoding ===")
	
	// Create a person
	person := Person{
		Name:     "Alice Johnson",
		Age:      30,
		Email:    "alice@example.com",
		IsActive: true,
		Birthday: time.Date(1993, 5, 15, 0, 0, 0, 0, time.UTC),
		Tags:     []string{"developer", "golang", "json"},
	}
	
	fmt.Printf("Original person: %+v\n", person)
	
	// Marshal to JSON (encode)
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatalf("Failed to marshal person: %v", err)
	}
	
	fmt.Printf("JSON string: %s\n", string(jsonData))
	
	// Marshal with indentation (pretty print)
	prettyJSON, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal with indent: %v", err)
	}
	
	fmt.Println("\nPretty JSON:")
	fmt.Println(string(prettyJSON))
	
	// Unmarshal from JSON (decode)
	var decodedPerson Person
	err = json.Unmarshal(jsonData, &decodedPerson)
	if err != nil {
		log.Fatalf("Failed to unmarshal person: %v", err)
	}
	
	fmt.Printf("\nDecoded person: %+v\n", decodedPerson)
	
	// Working with raw JSON string
	jsonString := `{
		"id": 1,
		"title": "The Go Programming Language",
		"author": "Alan Donovan and Brian Kernighan",
		"price": 39.99,
		"in_stock": true,
		"categories": ["programming", "golang", "computer science"],
		"published_at": "2015-11-16"
	}`
	
	fmt.Println("\n=== Decoding Raw JSON ===")
	fmt.Printf("Raw JSON string:\n%s\n", jsonString)
	
	var book Book
	err = json.Unmarshal([]byte(jsonString), &book)
	if err != nil {
		log.Fatalf("Failed to unmarshal book: %v", err)
	}
	
	fmt.Printf("Decoded book: %+v\n", book)
	
	// Working with map[string]interface{} for dynamic JSON
	fmt.Println("\n=== Dynamic JSON with map[string]interface{} ===")
	
	var dynamicData map[string]interface{}
	err = json.Unmarshal([]byte(jsonString), &dynamicData)
	if err != nil {
		log.Fatalf("Failed to unmarshal to map: %v", err)
	}
	
	fmt.Printf("Dynamic data: %+v\n", dynamicData)
	
	// Access dynamic fields
	fmt.Printf("Title: %v\n", dynamicData["title"])
	fmt.Printf("Price: %v\n", dynamicData["price"])
	fmt.Printf("Categories: %v\n", dynamicData["categories"])
	
	// Working with slice of structs
	fmt.Println("\n=== Array of JSON Objects ===")
	
	books := []Book{
		{
			ID:          1,
			Title:       "The Go Programming Language",
			Author:      "Alan Donovan",
			Price:       39.99,
			InStock:     true,
			Categories:  []string{"programming", "golang"},
			PublishedAt: "2015-11-16",
		},
		{
			ID:          2,
			Title:       "Clean Code",
			Author:      "Robert Martin",
			Price:       35.99,
			InStock:     false,
			Categories:  []string{"programming", "best practices"},
			PublishedAt: "2008-08-11",
		},
	}
	
	// Marshal slice to JSON
	booksJSON, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal books: %v", err)
	}
	
	fmt.Printf("Books JSON:\n%s\n", string(booksJSON))
	
	// Unmarshal back to slice
	var decodedBooks []Book
	err = json.Unmarshal(booksJSON, &decodedBooks)
	if err != nil {
		log.Fatalf("Failed to unmarshal books: %v", err)
	}
	
	fmt.Printf("Decoded books count: %d\n", len(decodedBooks))
	for i, book := range decodedBooks {
		fmt.Printf("Book %d: %s by %s\n", i+1, book.Title, book.Author)
	}
	
	// Handling JSON with missing fields
	fmt.Println("\n=== Handling Missing Fields ===")
	
	incompleteJSON := `{
		"name": "Bob Smith",
		"age": 25
	}`
	
	var incompletePerson Person
	err = json.Unmarshal([]byte(incompleteJSON), &incompletePerson)
	if err != nil {
		log.Fatalf("Failed to unmarshal incomplete person: %v", err)
	}
	
	fmt.Printf("Incomplete person: %+v\n", incompletePerson)
	fmt.Printf("Email (missing): '%s'\n", incompletePerson.Email)
	fmt.Printf("IsActive (missing): %t\n", incompletePerson.IsActive)
	fmt.Printf("Tags (omitempty): %v\n", incompletePerson.Tags)
}