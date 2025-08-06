// json_validation.go - SOLUTION
// Learn JSON validation and schema enforcement

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Product struct {
	ID          int      `json:"id" validate:"required,min=1"`
	Name        string   `json:"name" validate:"required,min=3,max=100"`
	Price       float64  `json:"price" validate:"required,min=0"`
	Category    string   `json:"category" validate:"required,oneof=electronics books clothing"`
	Tags        []string `json:"tags" validate:"dive,min=1"`
	InStock     bool     `json:"in_stock"`
	Description *string  `json:"description,omitempty" validate:"omitempty,max=500"`
}

func main() {
	fmt.Println("=== JSON Validation ===")
	
	// Test valid JSON
	validJSON := `{
		"id": 1,
		"name": "Laptop Computer",
		"price": 999.99,
		"category": "electronics",
		"tags": ["computer", "laptop", "portable"],
		"in_stock": true,
		"description": "High-performance laptop for professionals"
	}`
	
	fmt.Println("Testing valid JSON:")
	var validProduct Product
	
	// Unmarshal and validate
	err := json.Unmarshal([]byte(validJSON), &validProduct)
	if err != nil {
		fmt.Printf("JSON parsing error: %v\n", err)
	} else {
		fmt.Printf("Parsed product: %+v\n", validProduct)
		
		// Validate the product
		if validateProduct(validProduct) {
			fmt.Println("✓ Product is valid")
		} else {
			fmt.Println("✗ Product validation failed")
		}
	}
	
	fmt.Println("\n=== Invalid JSON Testing ===")
	
	// Test various invalid JSON scenarios
	invalidJSONs := map[string]string{
		"missing_required": `{
			"name": "Test Product",
			"price": 50.0,
			"category": "books",
			"tags": ["test"],
			"in_stock": true
		}`,
		"invalid_category": `{
			"id": 2,
			"name": "Invalid Product", 
			"price": 25.0,
			"category": "invalid_category",
			"tags": ["test"],
			"in_stock": false
		}`,
		"negative_price": `{
			"id": 3,
			"name": "Negative Price Product",
			"price": -10.0,
			"category": "electronics",
			"tags": ["test"],
			"in_stock": true
		}`,
		"empty_name": `{
			"id": 4,
			"name": "",
			"price": 100.0,
			"category": "clothing",
			"tags": ["test"],
			"in_stock": true
		}`,
	}
	
	for testName, invalidJSON := range invalidJSONs {
		fmt.Printf("\nTesting %s:\n", testName)
		
		var product Product
		err := json.Unmarshal([]byte(invalidJSON), &product)
		if err != nil {
			fmt.Printf("JSON parsing error: %v\n", err)
			continue
		}
		
		// Validate and show specific errors
		errors := validateProductWithErrors(product)
		if len(errors) > 0 {
			fmt.Printf("✗ Validation errors:\n")
			for _, errMsg := range errors {
				fmt.Printf("  - %s\n", errMsg)
			}
		} else {
			fmt.Println("✓ Product is valid")
		}
	}
	
	fmt.Println("\n=== JSON Schema Validation ===")
	
	// Define expected JSON structure
	expectedSchema := map[string]interface{}{
		"id":          "number",
		"name":        "string",
		"price":       "number", 
		"category":    "string",
		"tags":        "array",
		"in_stock":    "boolean",
		"description": "string",
	}
	
	// Test JSON against schema
	testJSON := `{
		"id": "not-a-number",
		"name": 123,
		"price": "invalid-price",
		"category": true,
		"tags": "not-an-array",
		"in_stock": "yes",
		"description": null
	}`
	
	fmt.Println("Testing JSON against schema:")
	schemaErrors := validateJSONSchema(testJSON, expectedSchema)
	
	if len(schemaErrors) > 0 {
		fmt.Println("✗ Schema validation errors:")
		for _, err := range schemaErrors {
			fmt.Printf("  - %s\n", err)
		}
	} else {
		fmt.Println("✓ JSON matches schema")
	}
	
	fmt.Println("\n=== Custom Validation Rules ===")
	
	// Implement custom validation logic
	customProduct := Product{
		ID:       1,
		Name:     "Custom Product",
		Price:    50.0,
		Category: "electronics",
		Tags:     []string{"custom", "test"},
		InStock:  true,
	}
	
	// Apply business logic validation
	businessErrors := validateBusinessRules(customProduct)
	
	fmt.Printf("Product: %+v\n", customProduct)
	if len(businessErrors) > 0 {
		fmt.Println("✗ Business rule violations:")
		for _, err := range businessErrors {
			fmt.Printf("  - %s\n", err)
		}
	} else {
		fmt.Println("✓ Product passes all business rules")
	}
}

// Implement validation functions
func validateProduct(p Product) bool {
	errors := validateProductWithErrors(p)
	return len(errors) == 0
}

func validateProductWithErrors(p Product) []string {
	var errors []string
	
	// Check required fields and constraints
	if p.ID <= 0 {
		errors = append(errors, "ID must be greater than 0")
	}
	
	if len(strings.TrimSpace(p.Name)) < 3 {
		errors = append(errors, "Name must be at least 3 characters")
	}
	
	if len(p.Name) > 100 {
		errors = append(errors, "Name must be at most 100 characters")
	}
	
	if p.Price < 0 {
		errors = append(errors, "Price must be non-negative")
	}
	
	validCategories := []string{"electronics", "books", "clothing"}
	validCategory := false
	for _, cat := range validCategories {
		if p.Category == cat {
			validCategory = true
			break
		}
	}
	if !validCategory {
		errors = append(errors, "Category must be one of: electronics, books, clothing")
	}
	
	if len(p.Tags) == 0 {
		errors = append(errors, "At least one tag is required")
	}
	
	for _, tag := range p.Tags {
		if len(strings.TrimSpace(tag)) == 0 {
			errors = append(errors, "Tags cannot be empty")
			break
		}
	}
	
	if p.Description != nil && len(*p.Description) > 500 {
		errors = append(errors, "Description must be at most 500 characters")
	}
	
	return errors
}

func validateJSONSchema(jsonStr string, schema map[string]interface{}) []string {
	var errors []string
	
	// Parse JSON and check types against schema
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		errors = append(errors, fmt.Sprintf("Invalid JSON: %v", err))
		return errors
	}
	
	// Check each field against expected type
	for field, expectedType := range schema {
		if value, exists := data[field]; exists {
			actualType := getJSONType(value)
			if actualType != expectedType {
				errors = append(errors, fmt.Sprintf("Field '%s' expected %s, got %s", field, expectedType, actualType))
			}
		}
	}
	
	return errors
}

func validateBusinessRules(p Product) []string {
	var errors []string
	
	// Implement custom business logic
	if p.Category == "electronics" && p.Price < 10.0 {
		errors = append(errors, "Electronics must be priced above $10")
	}
	
	if p.Category == "books" && len(p.Tags) < 2 {
		errors = append(errors, "Books must have at least 2 tags")
	}
	
	if p.Category == "clothing" && p.Description == nil {
		errors = append(errors, "Clothing items must have a description")
	}
	
	return errors
}

func getJSONType(value interface{}) string {
	// Return JSON type name for interface{} value
	switch value.(type) {
	case bool:
		return "boolean"
	case float64:
		return "number"
	case string:
		return "string"
	case []interface{}:
		return "array"
	case map[string]interface{}:
		return "object"
	case nil:
		return "null"
	default:
		return "unknown"
	}
}