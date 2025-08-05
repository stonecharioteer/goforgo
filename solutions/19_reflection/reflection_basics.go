// reflection_basics.go - SOLUTION
// Learn reflection in Go - examining types and values at runtime

package main

import (
	"fmt"
	"reflect"
)

// Define structs for reflection examples
type Person struct {
	Name    string `json:"name" validate:"required"`
	Age     int    `json:"age" validate:"min=0,max=150"`
	Email   string `json:"email" validate:"email"`
	private string // unexported field
}

func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d}", p.Name, p.Age)
}

func (p Person) GetInfo() string {
	return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func main() {
	fmt.Println("=== Basic Type Information ===")
	
	// Get type information for different values
	values := []interface{}{
		42,
		"hello",
		3.14,
		true,
		[]int{1, 2, 3},
		map[string]int{"a": 1},
		Person{Name: "Alice", Age: 30},
	}
	
	for _, v := range values {
		// Get type and value using reflection
		t := reflect.TypeOf(v)
		val := reflect.ValueOf(v)
		
		fmt.Printf("Value: %v\n", v)
		fmt.Printf("  Type: %s\n", t.Name())
		fmt.Printf("  Kind: %s\n", val.Kind())
		fmt.Printf("  Value: %v\n", val.Interface())
		fmt.Printf("  Can set: %t\n\n", val.CanSet())
	}
	
	fmt.Println("=== Struct Field Inspection ===")
	
	person := Person{
		Name:    "Bob",
		Age:     25,
		Email:   "bob@example.com",
		private: "secret",
	}
	
	// Inspect struct fields
	t := reflect.TypeOf(person)
	v := reflect.ValueOf(person)
	
	fmt.Printf("Struct: %s\n", t.Name())
	fmt.Printf("Number of fields: %d\n\n", t.NumField())
	
	// Iterate through fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)
		
		fmt.Printf("Field %d:\n", i)
		fmt.Printf("  Name: %s\n", field.Name)
		fmt.Printf("  Type: %s\n", field.Type)
		fmt.Printf("  Value: %v\n", fieldValue.Interface())
		fmt.Printf("  Exported: %t\n", field.IsExported())
		
		// Get struct tags
		if tag := field.Tag.Get("json"); tag != "" {
			fmt.Printf("  JSON tag: %s\n", tag)
		}
		if tag := field.Tag.Get("validate"); tag != "" {
			fmt.Printf("  Validate tag: %s\n", tag)
		}
		fmt.Println()
	}
	
	fmt.Println("=== Method Inspection ===")
	
	// Inspect methods
	fmt.Printf("Number of methods: %d\n\n", t.NumMethod())
	
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf("Method %d: %s\n", i, method.Name)
		fmt.Printf("  Type: %s\n", method.Type)
		fmt.Printf("  Num inputs: %d\n", method.Type.NumIn())
		fmt.Printf("  Num outputs: %d\n\n", method.Type.NumOut())
	}
	
	fmt.Println("=== Calling Methods via Reflection ===")
	
	// Call String method
	stringMethod := v.MethodByName("String")
	if stringMethod.IsValid() {
		// Call method with no arguments
		results := stringMethod.Call([]reflect.Value{})
		if len(results) > 0 {
			fmt.Printf("String() result: %s\n", results[0].String())
		}
	}
	
	// Call GetInfo method
	getInfoMethod := v.MethodByName("GetInfo")
	if getInfoMethod.IsValid() {
		results := getInfoMethod.Call([]reflect.Value{})
		if len(results) > 0 {
			fmt.Printf("GetInfo() result: %s\n", results[0].String())
		}
	}
	
	fmt.Println("\n=== Modifying Values ===")
	
	// Modify struct fields using reflection
	// Need pointer to modify
	personPtr := &Person{Name: "Charlie", Age: 35}
	v = reflect.ValueOf(personPtr)
	v = v.Elem()
	
	if v.CanSet() {
		// Get and modify Name field
		nameField := v.FieldByName("Name")
		if nameField.IsValid() && nameField.CanSet() {
			fmt.Printf("Original name: %s\n", nameField.String())
			nameField.SetString("Modified Charlie")
			fmt.Printf("Modified name: %s\n", nameField.String())
		}
		
		// Get and modify Age field
		ageField := v.FieldByName("Age")
		if ageField.IsValid() && ageField.CanSet() {
			fmt.Printf("Original age: %d\n", ageField.Int())
			ageField.SetInt(40)
			fmt.Printf("Modified age: %d\n", ageField.Int())
		}
	}
	
	fmt.Printf("Final person: %+v\n", *personPtr)
	
	fmt.Println("\n=== Working with Slices ===")
	
	// Create and modify slice using reflection
	slice := []int{1, 2, 3, 4, 5}
	sliceValue := reflect.ValueOf(slice)
	
	fmt.Printf("Original slice: %v\n", slice)
	fmt.Printf("Slice length: %d\n", sliceValue.Len())
	fmt.Printf("Slice capacity: %d\n", sliceValue.Cap())
	
	// Access slice elements
	for i := 0; i < sliceValue.Len(); i++ {
		elem := sliceValue.Index(i)
		fmt.Printf("Element %d: %v\n", i, elem.Interface())
	}
	
	// Modify slice elements
	if sliceValue.Len() > 0 {
		firstElem := sliceValue.Index(0)
		if firstElem.CanSet() {
			firstElem.SetInt(100)
			fmt.Printf("Modified slice: %v\n", slice)
		}
	}
	
	fmt.Println("\n=== Type Creation ===")
	
	// Create new instance of type
	personType := reflect.TypeOf(Person{})
	newPersonValue := reflect.New(personType)
	newPerson := newPersonValue.Interface().(*Person)
	
	// Set fields on new instance
	newPersonReflect := reflect.ValueOf(newPerson)
	newPersonReflect = newPersonReflect.Elem()
	
	nameField := newPersonReflect.FieldByName("Name")
	if nameField.CanSet() {
		nameField.SetString("Reflection Created")
	}
	
	ageField := newPersonReflect.FieldByName("Age")
	if ageField.CanSet() {
		ageField.SetInt(99)
	}
	
	fmt.Printf("Created person: %+v\n", *newPerson)
	
	fmt.Println("\n=== Interface Conversion ===")
	
	// Check if value implements interface
	var stringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	
	personType = reflect.TypeOf(person)
	implementsStringer := personType.Implements(stringerType)
	
	fmt.Printf("Person implements fmt.Stringer: %t\n", implementsStringer)
	
	if implementsStringer {
		// Convert to Stringer and call String()
		stringer := reflect.ValueOf(person).Interface().(fmt.Stringer)
		fmt.Printf("Via Stringer interface: %s\n", stringer.String())
	}
}