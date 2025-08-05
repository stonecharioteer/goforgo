// json_validation.go
// Learn JSON validation and schema enforcement

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
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
	
	// TODO: Test valid JSON
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
	
	// TODO: Unmarshal and validate
	err := /* unmarshal validJSON into validProduct */
	if err != nil {
		fmt.Printf("JSON parsing error: %v\n", err)
	} else {
		fmt.Printf("Parsed product: %+v\n", validProduct)
		
		// TODO: Validate the product
		if /* validate validProduct */ {
			fmt.Println("✓ Product is valid")
		} else {
			fmt.Println("✗ Product validation failed")
		}
	}
	
	fmt.Println("\n=== Invalid JSON Testing ===")
	
	// TODO: Test various invalid JSON scenarios
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
		err := /* unmarshal invalidJSON */
		if err != nil {
			fmt.Printf("JSON parsing error: %v\n", err)
			continue
		}
		
		// TODO: Validate and show specific errors
		errors := /* validate product and get errors */
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
	
	// TODO: Define expected JSON structure
	expectedSchema := map[string]interface{}{
		"id":          "number",
		"name":        "string",
		"price":       "number", 
		"category":    "string",
		"tags":        "array",
		"in_stock":    "boolean",
		"description": "string",
	}
	
	// TODO: Test JSON against schema
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
	schemaErrors := /* validate JSON against schema */
	
	if len(schemaErrors) > 0 {
		fmt.Println("✗ Schema validation errors:")
		for _, err := range schemaErrors {
			fmt.Printf("  - %s\n", err)
		}
	} else {
		fmt.Println("✓ JSON matches schema")
	}
	
	fmt.Println("\n=== Custom Validation Rules ===")
	
	// TODO: Implement custom validation logic
	customProduct := Product{
		ID:       1,
		Name:     "Custom Product",
		Price:    50.0,
		Category: "electronics",
		Tags:     []string{"custom", "test"},
		InStock:  true,
	}
	
	// TODO: Apply business logic validation
	businessErrors := /* apply custom business rules */
	
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

// TODO: Implement validation functions
func validateProduct(p Product) bool {
	// Basic validation logic
	return /* implement validation */
}

func validateProductWithErrors(p Product) []string {
	var errors []string
	
	// TODO: Check required fields and constraints
	// Check ID
	// Check Name
	// Check Price
	// Check Category
	// Check Tags
	
	return errors
}

func validateJSONSchema(jsonStr string, schema map[string]interface{}) []string {
	var errors []string
	
	// TODO: Parse JSON and check types against schema
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		errors = append(errors, fmt.Sprintf("Invalid JSON: %v", err))
		return errors
	}
	
	// TODO: Check each field against expected type
	for field, expectedType := range schema {
		// Validate field type
	}
	
	return errors
}

func validateBusinessRules(p Product) []string {
	var errors []string
	
	// TODO: Implement custom business logic
	// Example: Electronics must be priced above $10
	// Example: Books must have at least 2 tags
	// Example: Clothing must have description
	
	return errors
}

func getJSONType(value interface{}) string {
	// TODO: Return JSON type name for interface{} value
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