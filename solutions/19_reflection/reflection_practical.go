// reflection_practical.go
// Learn practical reflection applications in Go

package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Sample structs for demonstration
type User struct {
	ID       int       `json:"id" validate:"required,min=1"`
	Name     string    `json:"name" validate:"required,min=2"`
	Email    string    `json:"email" validate:"required,email"`
	Age      int       `json:"age" validate:"min=13,max=120"`
	IsActive bool      `json:"is_active"`
	Created  time.Time `json:"created_at"`
	Tags     []string  `json:"tags"`
}

type Product struct {
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required,min=0"`
	Category    string  `json:"category"`
	InStock     bool    `json:"in_stock"`
	Description *string `json:"description,omitempty"`
}

func main() {
	fmt.Println("=== Practical Reflection Applications ===")
	
	// Sample data
	user := User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john@example.com",
		Age:      30,
		IsActive: true,
		Created:  time.Now(),
		Tags:     []string{"admin", "developer"},
	}
	
	product := Product{
		Name:     "Laptop",
		Price:    999.99,
		Category: "Electronics",
		InStock:  true,
	}
	
	fmt.Println("\n=== Struct to Map Conversion ===")
	
	// Convert struct to map using reflection
	userMap := structToMap(user)
	fmt.Println("User as map:")
	for key, value := range userMap {
		fmt.Printf("  %s: %v (%T)\n", key, value, value)
	}
	
	productMap := structToMap(product)
	fmt.Println("\nProduct as map:")
	for key, value := range productMap {
		fmt.Printf("  %s: %v (%T)\n", key, value, value)
	}
	
	fmt.Println("\n=== JSON Tag Processing ===")
	
	// Extract JSON field names using reflection
	userJSONFields := extractJSONTags(user)
	fmt.Printf("User JSON fields: %v\n", userJSONFields)
	
	productJSONFields := extractJSONTags(product)
	fmt.Printf("Product JSON fields: %v\n", productJSONFields)
	
	fmt.Println("\n=== Validation Using Reflection ===")
	
	// Validate struct fields using validate tags
	fmt.Println("Validating user:")
	userErrors := validateStruct(user)
	if len(userErrors) == 0 {
		fmt.Println("  User is valid!")
	} else {
		for field, err := range userErrors {
			fmt.Printf("  %s: %s\n", field, err)
		}
	}
	
	// Test with invalid user
	invalidUser := User{
		ID:    0,    // Invalid: min=1
		Name:  "A",  // Invalid: min=2
		Email: "invalid-email", // Invalid: not email format
		Age:   150, // Invalid: max=120
	}
	
	fmt.Println("\nValidating invalid user:")
	invalidErrors := validateStruct(invalidUser)
	for field, err := range invalidErrors {
		fmt.Printf("  %s: %s\n", field, err)
	}
	
	fmt.Println("\n=== Dynamic Struct Creation ===")
	
	// Create struct type dynamically
	dynamicType := reflect.StructOf([]reflect.StructField{
		{
			Name: "Name",
			Type: reflect.TypeOf(""),
		},
		{
			Name: "Value",
			Type: reflect.TypeOf(0),
		},
	})
	
	// Create instance of dynamic type
	dynamicValue := reflect.New(dynamicType).Elem()
	
	// Set field values using reflection
	dynamicValue.FieldByName("Name").SetString("Dynamic Field")
	dynamicValue.FieldByName("Value").SetInt(42)
	
	fmt.Printf("Dynamic struct: %+v\n", dynamicValue.Interface())
	
	fmt.Println("\n=== Method Invocation ===")
	
	// Call methods using reflection
	userType := reflect.TypeOf(user)
	fmt.Printf("User has %d methods:\n", userType.NumMethod())
	for i := 0; i < userType.NumMethod(); i++ {
		method := userType.Method(i)
		fmt.Printf("  %s\n", method.Name)
	}
	
	// Call String method if it exists
	userValue := reflect.ValueOf(user)
	stringMethod := userValue.MethodByName("String")
	if stringMethod.IsValid() {
		// Call the method
		result := stringMethod.Call(nil)
		if len(result) > 0 {
			fmt.Printf("User.String() result: %v\n", result[0].Interface())
		}
	}
	
	fmt.Println("\n=== Field Modification ===")
	
	// Modify struct fields using reflection
	userCopy := user
	userValue = reflect.ValueOf(&userCopy).Elem()
	
	// Change Name field
	nameField := userValue.FieldByName("Name")
	if nameField.CanSet() {
		nameField.SetString("Jane Doe")
		fmt.Printf("Changed name to: %s\n", userCopy.Name)
	}
	
	// Change Age field
	ageField := userValue.FieldByName("Age")
	if ageField.CanSet() {
		ageField.SetInt(25)
		fmt.Printf("Changed age to: %d\n", userCopy.Age)
	}
	
	fmt.Println("\n=== Type Analysis ===")
	
	// Analyze different types
	values := []interface{}{
		42,
		"hello",
		[]int{1, 2, 3},
		map[string]int{"a": 1},
		user,
		&user,
		func(int) string { return "" },
		make(chan int),
	}
	
	for i, value := range values {
		fmt.Printf("%d. %v\n", i+1, analyzeType(value))
	}
	
	fmt.Println("\n=== Custom Marshal/Unmarshal ===")
	
	// Create custom JSON-like serialization
	userSerialized := customSerialize(user)
	fmt.Printf("Custom serialized user: %s\n", userSerialized)
	
	// Parse the serialized data back (simplified)
	fmt.Println("Serialized fields:")
	pairs := strings.Split(userSerialized, ",")
	for _, pair := range pairs {
		if kv := strings.Split(pair, ":"); len(kv) == 2 {
			fmt.Printf("  %s = %s\n", strings.TrimSpace(kv[0]), strings.TrimSpace(kv[1]))
		}
	}
	
	fmt.Println("\n=== Performance Comparison ===")
	
	// Compare reflection vs direct access performance
	iterations := 100000
	
	// Direct access
	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = user.Name
		_ = user.Age
		_ = user.Email
	}
	directTime := time.Since(start)
	
	// Reflection access
	userVal := reflect.ValueOf(user)
	start = time.Now()
	for i := 0; i < iterations; i++ {
		_ = userVal.FieldByName("Name").String()
		_ = int(userVal.FieldByName("Age").Int())
		_ = userVal.FieldByName("Email").String()
	}
	reflectionTime := time.Since(start)
	
	fmt.Printf("Direct access (%d iterations): %v\n", iterations, directTime)
	fmt.Printf("Reflection access (%d iterations): %v\n", iterations, reflectionTime)
	fmt.Printf("Reflection is %.2fx slower\n", float64(reflectionTime)/float64(directTime))
}

// Implement struct to map conversion
func structToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)
	
	// Handle pointer types
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	
	// Iterate through fields
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		
		// Skip unexported fields
		if !field.CanInterface() {
			continue
		}
		
		result[fieldType.Name] = field.Interface()
	}
	
	return result
}

// Implement JSON tag extraction
func extractJSONTags(obj interface{}) []string {
	var tags []string
	
	typ := reflect.TypeOf(obj)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	
	// Extract JSON tags from struct fields
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := field.Tag.Get("json")
		
		// Parse tag (remove omitempty, etc.)
		if jsonTag != "" {
			tagName := strings.Split(jsonTag, ",")[0]
			if tagName != "-" {
				tags = append(tags, tagName)
			}
		}
	}
	
	return tags
}

// Implement struct validation
func validateStruct(obj interface{}) map[string]string {
	errors := make(map[string]string)
	
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)
	
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	
	// Validate each field
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		
		validateTag := fieldType.Tag.Get("validate")
		if validateTag == "" {
			continue
		}
		
		// Parse validation rules
		rules := strings.Split(validateTag, ",")
		for _, rule := range rules {
			rule = strings.TrimSpace(rule)
			
			if err := validateField(field, rule); err != "" {
				errors[fieldType.Name] = err
				break
			}
		}
	}
	
	return errors
}

// Implement field validation
func validateField(field reflect.Value, rule string) string {
	switch {
	case rule == "required":
		if field.IsZero() {
			return "field is required"
		}
		
	case strings.HasPrefix(rule, "min="):
		minStr := strings.TrimPrefix(rule, "min=")
		min, _ := strconv.Atoi(minStr)
		
		switch field.Kind() {
		case reflect.String:
			if len(field.String()) < min {
				return fmt.Sprintf("minimum length is %d", min)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if field.Int() < int64(min) {
				return fmt.Sprintf("minimum value is %d", min)
			}
		}
		
	case strings.HasPrefix(rule, "max="):
		maxStr := strings.TrimPrefix(rule, "max=")
		max, _ := strconv.Atoi(maxStr)
		
		switch field.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if field.Int() > int64(max) {
				return fmt.Sprintf("maximum value is %d", max)
			}
		}
		
	case rule == "email":
		email := field.String()
		if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
			return "invalid email format"
		}
	}
	
	return ""
}

// Implement type analysis
func analyzeType(value interface{}) string {
	typ := reflect.TypeOf(value)
	val := reflect.ValueOf(value)
	
	analysis := fmt.Sprintf("Type: %v, Kind: %v", typ, val.Kind())
	
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		analysis += fmt.Sprintf(", Length: %d", val.Len())
	case reflect.Map:
		analysis += fmt.Sprintf(", Keys: %d", val.Len())
	case reflect.Struct:
		analysis += fmt.Sprintf(", Fields: %d", val.NumField())
	case reflect.Ptr:
		analysis += fmt.Sprintf(", Points to: %v", typ.Elem())
	case reflect.Func:
		funcType := typ
		analysis += fmt.Sprintf(", In: %d, Out: %d", funcType.NumIn(), funcType.NumOut())
	}
	
	return analysis
}

// Implement custom serialization
func customSerialize(obj interface{}) string {
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)
	
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	
	var parts []string
	
	// Serialize each field
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		
		if !field.CanInterface() {
			continue
		}
		
		part := fmt.Sprintf("%s:%v", fieldType.Name, field.Interface())
		parts = append(parts, part)
	}
	
	return strings.Join(parts, ",")
}

// Add String method to User for method invocation example
func (u User) String() string {
	return fmt.Sprintf("User{ID: %d, Name: %s, Email: %s}", u.ID, u.Name, u.Email)
}