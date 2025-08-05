// struct_tags.go
// Learn how to use struct tags for metadata and serialization

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// TODO: Define a User struct with JSON tags
type User struct {
	// Add struct tags for JSON serialization
	ID       int    // JSON tag: "id"
	Name     string // JSON tag: "name"  
	Email    string // JSON tag: "email"
	Password string // JSON tag: "-" (exclude from JSON)
	Age      int    // JSON tag: "age,omitempty" (omit if zero value)
}

// TODO: Define a Product struct with multiple tag types
type Product struct {
	// Add both JSON and custom validation tags
	SKU         string  // JSON: "sku", validate: "required"
	Name        string  // JSON: "product_name", validate: "required,min=3"
	Price       float64 // JSON: "price", validate: "required,min=0"
	InStock     bool    // JSON: "in_stock"
	Description string  // JSON: "description,omitempty"
}

// TODO: Define a Config struct with different tag formats
type Config struct {
	// Database config with custom tags
	DBHost     string // JSON: "db_host", env: "DB_HOST", default: "localhost"
	DBPort     int    // JSON: "db_port", env: "DB_PORT", default: "5432"
	DBName     string // JSON: "db_name", env: "DB_NAME"
	EnableSSL  bool   // JSON: "enable_ssl", env: "ENABLE_SSL"
	MaxRetries int    // JSON: "max_retries,omitempty", env: "MAX_RETRIES"
}

// Function to print struct tags using reflection
func printStructTags(v interface{}) {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	
	fmt.Printf("Struct tags for %s:\n", t.Name())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("  %s:\n", field.Name)
		
		// TODO: Print JSON tag if present
		if jsonTag, ok := field.Tag.Lookup("json"); ok {
			fmt.Printf("    json: %q\n", jsonTag)
		}
		
		// TODO: Print validate tag if present
		if validateTag, ok := /* get "validate" tag */; ok {
			fmt.Printf("    validate: %q\n", validateTag)
		}
		
		// TODO: Print env tag if present  
		if envTag, ok := /* get "env" tag */; ok {
			fmt.Printf("    env: %q\n", envTag)
		}
		
		// TODO: Print default tag if present
		if defaultTag, ok := /* get "default" tag */; ok {
			fmt.Printf("    default: %q\n", defaultTag)
		}
	}
	fmt.Println()
}

func main() {
	// TODO: Create a User instance
	user := // Create User with ID: 1, Name: "John Doe", Email: "john@example.com", Password: "secret123", Age: 0
	
	// TODO: Marshal User to JSON
	userJSON, err := // Convert user to JSON
	if err != nil {
		fmt.Printf("Error marshaling user: %v\n", err)
		return
	}
	
	fmt.Printf("User JSON: %s\n", userJSON)
	
	// TODO: Unmarshal JSON back to User
	var newUser User
	jsonData := `{"id":2,"name":"Jane Smith","email":"jane@example.com","age":25}`
	
	err = // Unmarshal jsonData into newUser
	if err != nil {
		fmt.Printf("Error unmarshaling: %v\n", err)
		return
	}
	
	fmt.Printf("Unmarshaled user: %+v\n", newUser)
	
	// TODO: Create a Product and marshal it
	product := // Create Product with all fields
	
	productJSON, err := json.Marshal(product)
	if err != nil {
		fmt.Printf("Error marshaling product: %v\n", err)
		return
	}
	
	fmt.Printf("Product JSON: %s\n", productJSON)
	
	// TODO: Demonstrate struct tag reflection
	fmt.Println("=== Struct Tag Analysis ===")
	printStructTags(user)
	printStructTags(product)
	printStructTags(Config{})
	
	// TODO: Show how omitempty works
	emptyUser := User{ID: 3, Name: "Empty User", Email: "empty@example.com"}
	emptyJSON, _ := json.Marshal(emptyUser)
	fmt.Printf("User with zero Age (omitempty): %s\n", emptyJSON)
	
	fullUser := User{ID: 4, Name: "Full User", Email: "full@example.com", Age: 30}
	fullJSON, _ := json.Marshal(fullUser)
	fmt.Printf("User with Age (included): %s\n", fullJSON)
}