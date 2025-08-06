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
	
	// TODO: Convert struct to map using reflection
	userMap := /* convert user struct to map[string]interface{} */
	fmt.Println("User as map:")
	for key, value := range userMap {
		fmt.Printf("  %s: %v (%T)\n", key, value, value)
	}
	
	productMap := /* convert product struct to map[string]interface{} */
	fmt.Println("\nProduct as map:")
	for key, value := range productMap {
		fmt.Printf("  %s: %v (%T)\n", key, value, value)
	}
	
	fmt.Println("\n=== JSON Tag Processing ===")
	
	// TODO: Extract JSON field names using reflection
	userJSONFields := /* extract JSON field names from user */
	fmt.Printf("User JSON fields: %v\n", userJSONFields)
	
	productJSONFields := /* extract JSON field names from product */
	fmt.Printf("Product JSON fields: %v\n", productJSONFields)
	
	fmt.Println("\n=== Validation Using Reflection ===")
	
	// TODO: Validate struct fields using validate tags
	fmt.Println("Validating user:")
	userErrors := /* validate user struct */
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
	invalidErrors := /* validate invalidUser struct */
	for field, err := range invalidErrors {
		fmt.Printf("  %s: %s\n", field, err)
	}
	
	fmt.Println("\n=== Dynamic Struct Creation ===")
	
	// TODO: Create struct type dynamically
	dynamicType := /* create dynamic struct type with fields: Name (string), Value (int) */
	
	// TODO: Create instance of dynamic type
	dynamicValue := /* create new instance of dynamic type */
	
	// TODO: Set field values using reflection
	/* set Name field to "Dynamic Field" */
	/* set Value field to 42 */
	
	fmt.Printf("Dynamic struct: %+v\n", dynamicValue.Interface())
	
	fmt.Println("\n=== Method Invocation ===")
	
	// TODO: Call methods using reflection
	methods := /* get methods of user */
	fmt.Printf("User has %d methods:\n", methods.Len())
	for i := 0; i < methods.Len(); i++ {
		method := methods.Index(i)
		fmt.Printf("  %s\n", method.Name())
	}
	
	// TODO: Call String method if it exists
	stringMethod := /* get String method from user */
	if stringMethod.IsValid() {
		// TODO: Call the method
		result := /* call String method with no arguments */
		if len(result) > 0 {
			fmt.Printf("User.String() result: %v\n", result[0].Interface())
		}
	}
	
	fmt.Println("\n=== Field Modification ===")
	
	// TODO: Modify struct fields using reflection
	userCopy := user
	/* get addressable value of userCopy */
	userValue := reflect.ValueOf(&userCopy).Elem()
	
	// TODO: Change Name field
	nameField := /* get Name field */
	if nameField.CanSet() {
		/* set Name to "Jane Doe" */
		fmt.Printf("Changed name to: %s\n", userCopy.Name)
	}
	
	// TODO: Change Age field
	ageField := /* get Age field */
	if ageField.CanSet() {
		/* set Age to 25 */
		fmt.Printf("Changed age to: %d\n", userCopy.Age)
	}
	
	fmt.Println("\n=== Type Analysis ===")
	
	// TODO: Analyze different types
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
		/* analyze each value using reflection */
		fmt.Printf("%d. %v\n", i+1, analyzeType(value))
	}
	
	fmt.Println("\n=== Custom Marshal/Unmarshal ===")
	
	// TODO: Create custom JSON-like serialization
	userSerialized := /* serialize user to custom format */
	fmt.Printf("Custom serialized user: %s\n", userSerialized)
	
	// TODO: Parse the serialized data back (simplified)
	fmt.Println("Serialized fields:")
	pairs := strings.Split(userSerialized, ",")
	for _, pair := range pairs {
		if kv := strings.Split(pair, ":"); len(kv) == 2 {
			fmt.Printf("  %s = %s\n", strings.TrimSpace(kv[0]), strings.TrimSpace(kv[1]))
		}
	}
	
	fmt.Println("\n=== Performance Comparison ===")
	
	// TODO: Compare reflection vs direct access performance
	iterations := 100000
	
	// Direct access
	start := /* get current time */
	for i := 0; i < iterations; i++ {
		_ = user.Name
		_ = user.Age
		_ = user.Email
	}
	directTime := /* calculate elapsed time */
	
	// Reflection access
	userVal := reflect.ValueOf(user)
	start = /* get current time */
	for i := 0; i < iterations; i++ {
		_ = userVal.FieldByName("Name").String()
		_ = int(userVal.FieldByName("Age").Int())
		_ = userVal.FieldByName("Email").String()
	}
	reflectionTime := /* calculate elapsed time */
	
	fmt.Printf("Direct access (%d iterations): %v\n", iterations, directTime)
	fmt.Printf("Reflection access (%d iterations): %v\n", iterations, reflectionTime)
	fmt.Printf("Reflection is %.2fx slower\n", float64(reflectionTime)/float64(directTime))
}

// TODO: Implement struct to map conversion
func structToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	
	/* get reflect value */
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)
	
	// TODO: Handle pointer types
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	
	// TODO: Iterate through fields
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		
		// Skip unexported fields
		if !field.CanInterface() {
			continue
		}
		
		/* add field name and value to result map */
		result[fieldType.Name] = field.Interface()
	}
	
	return result
}

// TODO: Implement JSON tag extraction
func extractJSONTags(obj interface{}) []string {
	var tags []string
	
	/* get type of object */
	typ := reflect.TypeOf(obj)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	
	// TODO: Extract JSON tags from struct fields
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		/* get json tag */
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

// TODO: Implement struct validation
func validateStruct(obj interface{}) map[string]string {
	errors := make(map[string]string)
	
	/* get value and type */
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)
	
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	
	// TODO: Validate each field
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		
		/* get validate tag */
		validateTag := fieldType.Tag.Get("validate")
		if validateTag == "" {
			continue
		}
		
		// Parse validation rules
		rules := strings.Split(validateTag, ",")
		for _, rule := range rules {
			rule = strings.TrimSpace(rule)
			
			/* validate field based on rule */
			if err := validateField(field, rule); err != "" {
				errors[fieldType.Name] = err
				break
			}
		}
	}
	
	return errors
}

// TODO: Implement field validation
func validateField(field reflect.Value, rule string) string {
	switch {
	case rule == "required":
		/* check if field is zero value */
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

// TODO: Implement type analysis
func analyzeType(value interface{}) string {
	/* get type and value using reflection */
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

// TODO: Implement custom serialization
func customSerialize(obj interface{}) string {
	/* get value and type */
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)
	
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	
	var parts []string
	
	// TODO: Serialize each field
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		
		if !field.CanInterface() {
			continue
		}
		
		/* create field:value pair */
		part := fmt.Sprintf("%s:%v", fieldType.Name, field.Interface())
		parts = append(parts, part)
	}
	
	return strings.Join(parts, ",")
}

// Add String method to User for method invocation example
func (u User) String() string {
	return fmt.Sprintf("User{ID: %d, Name: %s, Email: %s}", u.ID, u.Name, u.Email)
}