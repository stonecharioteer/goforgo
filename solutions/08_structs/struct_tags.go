// struct_tags.go - SOLUTION
// Learn how to use struct tags for metadata and serialization

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// User struct with JSON tags
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`                // exclude from JSON
	Age      int    `json:"age,omitempty"`    // omit if zero value
}

// Product struct with multiple tag types
type Product struct {
	SKU         string  `json:"sku" validate:"required"`
	Name        string  `json:"product_name" validate:"required,min=3"`
	Price       float64 `json:"price" validate:"required,min=0"`
	InStock     bool    `json:"in_stock"`
	Description string  `json:"description,omitempty"`
}

// Config struct with different tag formats
type Config struct {
	DBHost     string `json:"db_host" env:"DB_HOST" default:"localhost"`
	DBPort     int    `json:"db_port" env:"DB_PORT" default:"5432"`
	DBName     string `json:"db_name" env:"DB_NAME"`
	EnableSSL  bool   `json:"enable_ssl" env:"ENABLE_SSL"`
	MaxRetries int    `json:"max_retries,omitempty" env:"MAX_RETRIES"`
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
		
		// Print JSON tag if present
		if jsonTag, ok := field.Tag.Lookup("json"); ok {
			fmt.Printf("    json: %q\n", jsonTag)
		}
		
		// Print validate tag if present
		if validateTag, ok := field.Tag.Lookup("validate"); ok {
			fmt.Printf("    validate: %q\n", validateTag)
		}
		
		// Print env tag if present
		if envTag, ok := field.Tag.Lookup("env"); ok {
			fmt.Printf("    env: %q\n", envTag)
		}
		
		// Print default tag if present
		if defaultTag, ok := field.Tag.Lookup("default"); ok {
			fmt.Printf("    default: %q\n", defaultTag)
		}
	}
	fmt.Println()
}

func main() {
	// Create a User instance
	user := User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "secret123",
		Age:      0, // Zero value, will be omitted due to omitempty
	}
	
	// Marshal User to JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("Error marshaling user: %v\n", err)
		return
	}
	
	fmt.Printf("User JSON: %s\n", userJSON)
	
	// Unmarshal JSON back to User
	var newUser User
	jsonData := `{"id":2,"name":"Jane Smith","email":"jane@example.com","age":25}`
	
	err = json.Unmarshal([]byte(jsonData), &newUser)
	if err != nil {
		fmt.Printf("Error unmarshaling: %v\n", err)
		return
	}
	
	fmt.Printf("Unmarshaled user: %+v\n", newUser)
	
	// Create a Product and marshal it
	product := Product{
		SKU:         "ABC123",
		Name:        "Wireless Headphones",
		Price:       99.99,
		InStock:     true,
		Description: "High-quality wireless headphones with noise cancellation",
	}
	
	productJSON, err := json.Marshal(product)
	if err != nil {
		fmt.Printf("Error marshaling product: %v\n", err)
		return
	}
	
	fmt.Printf("Product JSON: %s\n", productJSON)
	
	// Demonstrate struct tag reflection
	fmt.Println("=== Struct Tag Analysis ===")
	printStructTags(user)
	printStructTags(product)
	printStructTags(Config{})
	
	// Show how omitempty works
	emptyUser := User{ID: 3, Name: "Empty User", Email: "empty@example.com"}
	emptyJSON, _ := json.Marshal(emptyUser)
	fmt.Printf("User with zero Age (omitempty): %s\n", emptyJSON)
	
	fullUser := User{ID: 4, Name: "Full User", Email: "full@example.com", Age: 30}
	fullJSON, _ := json.Marshal(fullUser)
	fmt.Printf("User with Age (included): %s\n", fullJSON)
}