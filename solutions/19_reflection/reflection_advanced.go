// reflection_advanced.go - SOLUTION
// Learn advanced reflection techniques for dynamic programming

package main

import (
	"fmt"
	"reflect"
	"strings"
)

// Define interfaces and types for reflection
type Validator interface {
	Validate() error
}

type Serializer interface {
	Serialize() ([]byte, error)
	Deserialize([]byte) error
}

type User struct {
	ID       int    `json:"id" validate:"required,min=1"`
	Name     string `json:"name" validate:"required,min=2,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Age      int    `json:"age" validate:"min=0,max=150"`
	IsActive bool   `json:"is_active"`
	Tags     []string `json:"tags,omitempty"`
}

func (u User) String() string {
	return fmt.Sprintf("User{ID:%d, Name:%s}", u.ID, u.Name)
}

type Product struct {
	ID    int     `json:"id" validate:"required"`
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price" validate:"min=0"`
}

func main() {
	fmt.Println("=== Advanced Reflection Techniques ===")
	
	fmt.Println("\n=== Dynamic Function Calls ===")
	
	// Create a registry of functions
	funcRegistry := map[string]interface{}{
		"add":      add,
		"multiply": multiply,
		"greet":    greet,
		"format":   formatUser,
	}
	
	// Call functions dynamically
	testCalls := []struct {
		name string
		args []interface{}
	}{
		{"add", []interface{}{5, 3}},
		{"multiply", []interface{}{4, 7}},
		{"greet", []interface{}{"Alice"}},
		{"format", []interface{}{User{ID: 1, Name: "Bob", Email: "bob@test.com", Age: 30}}},
	}
	
	fmt.Println("Dynamic function calls:")
	for _, test := range testCalls {
		fmt.Printf("Calling %s with args %v:\n", test.name, test.args)
		
		// Get function from registry
		fn, exists := funcRegistry[test.name]
		if !exists {
			fmt.Printf("  Function %s not found\n", test.name)
			continue
		}
		
		// Call function using reflection
		result := callFunction(fn, test.args)
		fmt.Printf("  Result: %v\n", result)
	}
	
	fmt.Println("\n=== Struct Field Validation ===")
	
	// Implement validation using struct tags
	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Age: 25, IsActive: true},
		{ID: 0, Name: "B", Email: "invalid-email", Age: -5, IsActive: false}, // Invalid
		{ID: 2, Name: "Charlie", Email: "charlie@test.com", Age: 35, IsActive: true},
	}
	
	fmt.Println("Validating users:")
	for i, user := range users {
		fmt.Printf("User %d: %+v\n", i+1, user)
		
		// Validate user using reflection
		errors := validateStruct(user)
		if len(errors) == 0 {
			fmt.Println("  ✓ Valid")
		} else {
			fmt.Println("  ✗ Validation errors:")
			for _, err := range errors {
				fmt.Printf("    - %s\n", err)
			}
		}
		fmt.Println()
	}
	
	fmt.Println("=== Dynamic Struct Creation ===")
	
	// Create struct types dynamically
	fmt.Println("Creating dynamic struct types:")
	
	// Define field specifications
	fieldSpecs := []reflect.StructField{
		{Name: "ID", Type: reflect.TypeOf(0), Tag: `json:"id"`},
		{Name: "Title", Type: reflect.TypeOf(""), Tag: `json:"title"`},
		{Name: "Active", Type: reflect.TypeOf(false), Tag: `json:"active"`},
		{Name: "Score", Type: reflect.TypeOf(0.0), Tag: `json:"score"`},
	}
	
	// Create struct type dynamically
	dynamicType := reflect.StructOf(fieldSpecs)
	
	fmt.Printf("Created dynamic type: %s\n", dynamicType)
	
	// Create instance of dynamic type
	instance := reflect.New(dynamicType).Elem()
	
	// Set field values using reflection
	instance.Field(0).SetInt(100)           // ID
	instance.Field(1).SetString("Dynamic Title") // Title
	instance.Field(2).SetBool(true)         // Active
	instance.Field(3).SetFloat(95.5)        // Score
	
	fmt.Printf("Dynamic instance: %v\n", instance.Interface())
	
	fmt.Println("\n=== Interface Implementation Detection ===")
	
	// Check interface implementations
	types := []interface{}{
		User{},
		Product{},
		"string",
		42,
		[]int{1, 2, 3},
	}
	
	// Define interfaces to check
	validatorType := reflect.TypeOf((*Validator)(nil)).Elem()
	serializerType := reflect.TypeOf((*Serializer)(nil)).Elem()
	stringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	
	fmt.Println("Checking interface implementations:")
	for _, obj := range types {
		objType := reflect.TypeOf(obj)
		fmt.Printf("Type: %s\n", objType)
		
		// Check interface implementations
		fmt.Printf("  Implements Validator: %t\n", objType.Implements(validatorType))
		fmt.Printf("  Implements Serializer: %t\n", objType.Implements(serializerType))
		fmt.Printf("  Implements fmt.Stringer: %t\n", objType.Implements(stringerType))
		fmt.Println()
	}
	
	fmt.Println("=== Method Invocation ===")
	
	// Invoke methods dynamically
	user := User{ID: 1, Name: "John", Email: "john@example.com", Age: 30}
	userValue := reflect.ValueOf(&user)
	
	fmt.Printf("User: %+v\n", user)
	fmt.Println("Available methods:")
	
	// List all methods
	userType := userValue.Type()
	for i := 0; i < userType.NumMethod(); i++ {
		method := userType.Method(i)
		fmt.Printf("  %s: %s\n", method.Name, method.Type)
	}
	
	// Call method if it exists
	method := reflect.ValueOf(user).MethodByName("String")
	if method.IsValid() {
		// Call String method
		results := method.Call([]reflect.Value{})
		if len(results) > 0 {
			fmt.Printf("String() result: %s\n", results[0].String())
		}
	}
	
	fmt.Println("\n=== Deep Value Comparison ===")
	
	// Compare values deeply using reflection
	user1 := User{ID: 1, Name: "Alice", Email: "alice@test.com", Age: 25}
	user2 := User{ID: 1, Name: "Alice", Email: "alice@test.com", Age: 25}
	user3 := User{ID: 2, Name: "Bob", Email: "bob@test.com", Age: 30}
	
	fmt.Println("Deep value comparison:")
	fmt.Printf("user1 == user2: %t\n", deepEqual(user1, user2))
	fmt.Printf("user1 == user3: %t\n", deepEqual(user1, user3))
	
	// Show field-by-field comparison
	fmt.Println("\nField-by-field comparison (user1 vs user3):")
	differences := compareFields(user1, user3)
	for fieldName, values := range differences {
		fmt.Printf("  %s: %v != %v\n", fieldName, values[0], values[1])
	}
}

// Implement helper functions for reflection demos

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func formatUser(u User) string {
	return fmt.Sprintf("User{ID:%d, Name:%s, Age:%d}", u.ID, u.Name, u.Age)
}

// Implement callFunction using reflection
func callFunction(fn interface{}, args []interface{}) []interface{} {
	fnValue := reflect.ValueOf(fn)
	fnType := fnValue.Type()
	
	// Prepare arguments
	callArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		callArgs[i] = reflect.ValueOf(arg)
	}
	
	// Call function
	results := fnValue.Call(callArgs)
	
	// Convert results to interface{} slice
	resultInterfaces := make([]interface{}, len(results))
	for i, result := range results {
		resultInterfaces[i] = result.Interface()
	}
	
	return resultInterfaces
}

// Implement validateStruct using reflection and tags
func validateStruct(obj interface{}) []string {
	var errors []string
	
	value := reflect.ValueOf(obj)
	objType := reflect.TypeOf(obj)
	
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		fieldValue := value.Field(i)
		validateTag := field.Tag.Get("validate")
		
		if validateTag == "" {
			continue
		}
		
		rules := strings.Split(validateTag, ",")
		for _, rule := range rules {
			rule = strings.TrimSpace(rule)
			
			switch {
			case rule == "required":
				if isZeroValue(fieldValue) {
					errors = append(errors, fmt.Sprintf("Field %s is required", field.Name))
				}
			case strings.HasPrefix(rule, "min="):
				minVal := strings.TrimPrefix(rule, "min=")
				if fieldValue.Kind() == reflect.Int && fieldValue.Int() < parseInt(minVal) {
					errors = append(errors, fmt.Sprintf("Field %s must be >= %s", field.Name, minVal))
				}
				if fieldValue.Kind() == reflect.String && len(fieldValue.String()) < parseInt(minVal) {
					errors = append(errors, fmt.Sprintf("Field %s must have length >= %s", field.Name, minVal))
				}
			case strings.HasPrefix(rule, "max="):
				maxVal := strings.TrimPrefix(rule, "max=")
				if fieldValue.Kind() == reflect.Int && fieldValue.Int() > parseInt(maxVal) {
					errors = append(errors, fmt.Sprintf("Field %s must be <= %s", field.Name, maxVal))
				}
				if fieldValue.Kind() == reflect.String && len(fieldValue.String()) > parseInt(maxVal) {
					errors = append(errors, fmt.Sprintf("Field %s must have length <= %s", field.Name, maxVal))
				}
			case rule == "email":
				if fieldValue.Kind() == reflect.String && !strings.Contains(fieldValue.String(), "@") {
					errors = append(errors, fmt.Sprintf("Field %s must be a valid email", field.Name))
				}
			}
		}
	}
	
	return errors
}

// Implement deepEqual using reflection
func deepEqual(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

// Implement compareFields to show differences
func compareFields(a, b interface{}) map[string][2]interface{} {
	differences := make(map[string][2]interface{})
	
	valueA := reflect.ValueOf(a)
	valueB := reflect.ValueOf(b)
	typeA := reflect.TypeOf(a)
	
	if typeA != reflect.TypeOf(b) {
		return differences
	}
	
	for i := 0; i < typeA.NumField(); i++ {
		field := typeA.Field(i)
		fieldA := valueA.Field(i)
		fieldB := valueB.Field(i)
		
		if !reflect.DeepEqual(fieldA.Interface(), fieldB.Interface()) {
			differences[field.Name] = [2]interface{}{fieldA.Interface(), fieldB.Interface()}
		}
	}
	
	return differences
}

// Helper functions
func isZeroValue(v reflect.Value) bool {
	zero := reflect.Zero(v.Type())
	return reflect.DeepEqual(v.Interface(), zero.Interface())
}

func parseInt(s string) int64 {
	// Simple string to int conversion for demo
	switch s {
	case "0":
		return 0
	case "1":
		return 1
	case "2":
		return 2
	case "50":
		return 50
	case "150":
		return 150
	default:
		return 0
	}
}