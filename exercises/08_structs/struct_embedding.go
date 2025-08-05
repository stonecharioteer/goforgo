// struct_embedding.go
// Learn advanced struct embedding patterns and composition

package main

import "fmt"

// TODO: Define a base Engine struct
type Engine struct {
	Horsepower int
	Type       string
}

// TODO: Add a Start method to Engine
func (e Engine) Start() {
	fmt.Printf("Starting %s engine with %d HP\n", e.Type, e.Horsepower)
}

// TODO: Add a Stop method to Engine
func (e Engine) Stop() {
	fmt.Printf("Stopping %s engine\n", e.Type)
}

// TODO: Define a Car struct that embeds Engine
type Car struct {
	Make  string
	Model string
	Year  int
	// Embed Engine here
}

// TODO: Add a Drive method to Car
func (c Car) Drive() {
	fmt.Printf("Driving %d %s %s\n", c.Year, c.Make, c.Model)
}

// TODO: Override the Start method for Car (method promotion)
func (c Car) Start() {
	fmt.Printf("Starting %d %s %s\n", c.Year, c.Make, c.Model)
	// Call the embedded Engine's Start method
}

// TODO: Define multiple embedded structs
type GPS struct {
	Brand    string
	HasMaps  bool
}

func (g GPS) Navigate(destination string) {
	fmt.Printf("%s GPS navigating to %s\n", g.Brand, destination)
}

type Radio struct {
	Brand     string
	HasBluetooth bool
}

func (r Radio) PlayMusic() {
	fmt.Printf("Playing music on %s radio\n", r.Brand)
}

// TODO: Define a LuxuryCar that embeds multiple structs
type LuxuryCar struct {
	// Embed Car
	// Embed GPS  
	// Embed Radio
	LeatherSeats bool
}

// TODO: Define structs with name conflicts
type Person struct {
	Name string
	Age  int
}

type Company struct {
	Name      string
	Employees int
}

// TODO: Define Employee that embeds both Person and Company
// This will create a name conflict for the Name field
type Employee struct {
	ID       int
	Position string
	// Embed Person
	// Embed Company
}

func main() {
	// TODO: Create a Car and test method promotion
	car := // Create Car with Make: "Toyota", Model: "Camry", Year: 2023, Engine with Horsepower: 200, Type: "V6"
	
	// TODO: Call methods - both Car's and Engine's
	car.Start() // This calls Car's Start method
	// Call the embedded Engine's Start method explicitly
	car.Drive()
	// Call Engine's Stop method through promotion
	
	fmt.Println("---")
	
	// TODO: Create a LuxuryCar with all embedded structs
	luxuryCar := // Create LuxuryCar with all fields populated
	
	// TODO: Call methods from all embedded structs
	luxuryCar.Start()
	luxuryCar.Drive()
	// Navigate to "Downtown"
	// Play music
	
	fmt.Println("---")
	
	// TODO: Handle name conflicts in Employee
	emp := // Create Employee with conflicting Name fields
	
	// TODO: Access fields with conflicts - must be explicit
	fmt.Printf("Employee: %s (ID: %d)\n", /* access Person's Name */, emp.ID)
	fmt.Printf("Works at: %s\n", /* access Company's Name */)
	fmt.Printf("Position: %s\n", emp.Position)
	fmt.Printf("Age: %d\n", /* access Person's Age */)
	
	// TODO: Demonstrate struct as interface
	// Since Car embeds Engine, Car has all of Engine's methods
	var engine Engine = car.Engine // Extract the embedded struct
	engine.Start()
}