// struct_embedding.go - SOLUTION
// Learn advanced struct embedding patterns and composition

package main

import "fmt"

// Base Engine struct
type Engine struct {
	Horsepower int
	Type       string
}

// Start method for Engine
func (e Engine) Start() {
	fmt.Printf("Starting %s engine with %d HP\n", e.Type, e.Horsepower)
}

// Stop method for Engine
func (e Engine) Stop() {
	fmt.Printf("Stopping %s engine\n", e.Type)
}

// Car struct that embeds Engine
type Car struct {
	Make   string
	Model  string
	Year   int
	Engine // Embedded Engine
}

// Drive method for Car
func (c Car) Drive() {
	fmt.Printf("Driving %d %s %s\n", c.Year, c.Make, c.Model)
}

// Override the Start method for Car (method promotion)
func (c Car) Start() {
	fmt.Printf("Starting %d %s %s\n", c.Year, c.Make, c.Model)
	// Call the embedded Engine's Start method
	c.Engine.Start()
}

// GPS struct
type GPS struct {
	Brand   string
	HasMaps bool
}

func (g GPS) Navigate(destination string) {
	fmt.Printf("%s GPS navigating to %s\n", g.Brand, destination)
}

// Radio struct
type Radio struct {
	Brand        string
	HasBluetooth bool
}

func (r Radio) PlayMusic() {
	fmt.Printf("Playing music on %s radio\n", r.Brand)
}

// LuxuryCar that embeds multiple structs
type LuxuryCar struct {
	Car          // Embed Car
	GPS          // Embed GPS
	Radio        // Embed Radio
	LeatherSeats bool
}

// Structs with name conflicts
type Person struct {
	Name string
	Age  int
}

type Company struct {
	Name      string
	Employees int
}

// Employee that embeds both Person and Company
// This creates a name conflict for the Name field
type Employee struct {
	ID       int
	Position string
	Person   // Embed Person
	Company  // Embed Company
}

func main() {
	// Create a Car and test method promotion
	car := Car{
		Make:  "Toyota",
		Model: "Camry",
		Year:  2023,
		Engine: Engine{
			Horsepower: 200,
			Type:       "V6",
		},
	}
	
	// Call methods - both Car's and Engine's
	car.Start()                 // This calls Car's Start method
	car.Engine.Start()          // Call the embedded Engine's Start method explicitly
	car.Drive()
	car.Stop()                  // Call Engine's Stop method through promotion
	
	fmt.Println("---")
	
	// Create a LuxuryCar with all embedded structs
	luxuryCar := LuxuryCar{
		Car: Car{
			Make:  "BMW",
			Model: "X7",
			Year:  2024,
			Engine: Engine{
				Horsepower: 400,
				Type:       "V8",
			},
		},
		GPS: GPS{
			Brand:   "Garmin",
			HasMaps: true,
		},
		Radio: Radio{
			Brand:        "Bose",
			HasBluetooth: true,
		},
		LeatherSeats: true,
	}
	
	// Call methods from all embedded structs
	luxuryCar.Start()
	luxuryCar.Drive()
	luxuryCar.Navigate("Downtown")
	luxuryCar.PlayMusic()
	
	fmt.Println("---")
	
	// Handle name conflicts in Employee
	emp := Employee{
		ID:       1001,
		Position: "Software Engineer",
		Person: Person{
			Name: "John Doe",
			Age:  30,
		},
		Company: Company{
			Name:      "Tech Corp",
			Employees: 500,
		},
	}
	
	// Access fields with conflicts - must be explicit
	fmt.Printf("Employee: %s (ID: %d)\n", emp.Person.Name, emp.ID)
	fmt.Printf("Works at: %s\n", emp.Company.Name)
	fmt.Printf("Position: %s\n", emp.Position)
	fmt.Printf("Age: %d\n", emp.Person.Age)
	
	// Demonstrate struct as interface
	// Since Car embeds Engine, Car has all of Engine's methods
	var engine Engine = car.Engine // Extract the embedded struct
	engine.Start()
}