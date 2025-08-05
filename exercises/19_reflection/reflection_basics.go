// reflection_basics.go
// Learn reflection in Go - examining types and values at runtime

package main

import (
	"fmt"
	"reflect"
)

// TODO: Define structs for reflection examples
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
	
	// TODO: Get type information for different values
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
		// TODO: Get type and value using reflection
		t := /* get type of v */
		val := /* get value of v */
		
		fmt.Printf("Value: %v\\n", v)
		fmt.Printf("  Type: %s\\n", /* get type name */)
		fmt.Printf("  Kind: %s\\n", /* get type kind */)
		fmt.Printf("  Value: %v\\n", /* get interface value */)
		fmt.Printf("  Can set: %t\\n\\n", /* check if settable */)
	}
	
	fmt.Println("=== Struct Field Inspection ===")
	
	person := Person{
		Name:    "Bob",
		Age:     25,
		Email:   "bob@example.com",
		private: "secret",
	}
	
	// TODO: Inspect struct fields
	t := /* get type of person */
	v := /* get value of person */
	
	fmt.Printf("Struct: %s\\n", t.Name())
	fmt.Printf("Number of fields: %d\\n\\n", /* get number of fields */)
	
	// TODO: Iterate through fields
	for i := 0; i < /* get number of fields */; i++ {
		field := /* get field at index i */
		fieldValue := /* get field value at index i */
		
		fmt.Printf("Field %d:\\n", i)
		fmt.Printf("  Name: %s\\n", field.Name)
		fmt.Printf("  Type: %s\\n", field.Type)
		fmt.Printf("  Value: %v\\n", /* get interface value */)
		fmt.Printf("  Exported: %t\\n", /* check if field is exported */)
		
		// TODO: Get struct tags
		if tag := /* get "json" tag */; tag != "" {
			fmt.Printf("  JSON tag: %s\\n", tag)
		}
		if tag := /* get "validate" tag */; tag != "" {
			fmt.Printf("  Validate tag: %s\\n", tag)
		}
		fmt.Println()
	}
	
	fmt.Println("=== Method Inspection ===")
	
	// TODO: Inspect methods
	fmt.Printf("Number of methods: %d\\n\\n", /* get number of methods */)
	
	for i := 0; i < /* get number of methods */; i++ {
		method := /* get method at index i */
		fmt.Printf("Method %d: %s\\n", i, method.Name)
		fmt.Printf("  Type: %s\\n", method.Type)
		fmt.Printf("  Num inputs: %d\\n", /* get number of input parameters */)
		fmt.Printf("  Num outputs: %d\\n\\n", /* get number of output parameters */)
	}
	
	fmt.Println("=== Calling Methods via Reflection ===")
	
	// TODO: Call String method
	stringMethod := /* get method by name "String" */
	if stringMethod.IsValid() {
		// TODO: Call method with no arguments
		results := /* call method with empty args */
		if len(results) > 0 {
			fmt.Printf("String() result: %s\\n", results[0].String())
		}
	}
	
	// TODO: Call GetInfo method
	getInfoMethod := /* get method by name "GetInfo" */
	if getInfoMethod.IsValid() {
		results := /* call method */
		if len(results) > 0 {
			fmt.Printf("GetInfo() result: %s\\n", results[0].String())
		}
	}
	
	fmt.Println("\\n=== Modifying Values ===")
	
	// TODO: Modify struct fields using reflection
	// Need pointer to modify
	personPtr := &Person{Name: "Charlie", Age: 35}
	v = /* get value of personPtr */
	v = /* get element that pointer points to */
	
	if /* check if v is settable */ {
		// TODO: Get and modify Name field
		nameField := /* get field by name "Name" */
		if nameField.IsValid() && nameField.CanSet() {
			fmt.Printf("Original name: %s\\n", nameField.String())
			/* set name field to "Modified Charlie" */
			fmt.Printf("Modified name: %s\\n", nameField.String())
		}
		
		// TODO: Get and modify Age field
		ageField := /* get field by name "Age" */
		if ageField.IsValid() && ageField.CanSet() {
			fmt.Printf("Original age: %d\\n", /* get int value */)
			/* set age field to 40 */
			fmt.Printf("Modified age: %d\\n", /* get int value */)
		}
	}
	
	fmt.Printf("Final person: %+v\\n", *personPtr)
	
	fmt.Println("\\n=== Working with Slices ===")
	
	// TODO: Create and modify slice using reflection
	slice := []int{1, 2, 3, 4, 5}
	sliceValue := /* get value of slice */
	
	fmt.Printf("Original slice: %v\\n", slice)
	fmt.Printf("Slice length: %d\\n", /* get slice length */)
	fmt.Printf("Slice capacity: %d\\n", /* get slice capacity */)
	
	// TODO: Access slice elements
	for i := 0; i < sliceValue.Len(); i++ {
		elem := /* get element at index i */
		fmt.Printf("Element %d: %v\\n", i, /* get interface value */)
	}
	
	// TODO: Modify slice elements
	if sliceValue.Len() > 0 {
		firstElem := /* get element at index 0 */
		if firstElem.CanSet() {
			/* set first element to 100 */
			fmt.Printf("Modified slice: %v\\n", slice)
		}
	}
	
	fmt.Println("\\n=== Type Creation ===")
	
	// TODO: Create new instance of type
	personType := /* get type of Person{} */
	newPersonValue := /* create new value of personType */
	newPerson := /* get interface value as *Person */
	
	// TODO: Set fields on new instance
	newPersonReflect := /* get value of newPerson */
	newPersonReflect = /* get element */
	
	nameField := /* get "Name" field */
	if nameField.CanSet() {
		/* set name to "Reflection Created" */
	}
	
	ageField := /* get "Age" field */
	if ageField.CanSet() {
		/* set age to 99 */
	}
	
	fmt.Printf("Created person: %+v\\n", *newPerson)
	
	fmt.Println("\\n=== Interface Conversion ===")
	
	// TODO: Check if value implements interface
	var stringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	
	personType = /* get type of person */
	implementsStringer := /* check if personType implements stringerType */
	
	fmt.Printf("Person implements fmt.Stringer: %t\\n", implementsStringer)
	
	if implementsStringer {
		// TODO: Convert to Stringer and call String()
		stringer := /* convert person to fmt.Stringer */
		fmt.Printf("Via Stringer interface: %s\\n", stringer.String())
	}
}