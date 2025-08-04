package main

import "fmt"

// Define a struct called 'Engine' with horsepower (int) and fuelType (string) fields
type Engine struct {
	horsepower int
	fuelType   string
}

// Define methods on Engine:
// - Start() - prints "Engine started!"
func (e Engine) Start() {
	fmt.Println("Engine started!")
}

// - Stop() - prints "Engine stopped!"  
func (e Engine) Stop() {
	fmt.Println("Engine stopped!")
}

// - Info() - prints "Engine: [horsepower] HP, Fuel: [fuelType]"
func (e Engine) Info() {
	fmt.Printf("Engine: %d HP, Fuel: %s\n", e.horsepower, e.fuelType)
}

// Define a struct called 'Car' that embeds Engine and has brand (string) and model (string) fields
type Car struct {
	Engine
	brand string
	model string
}

// Define methods on Car:
// - Drive() - prints "[brand] [model] is driving!"
func (c Car) Drive() {
	fmt.Printf("%s %s is driving!\n", c.brand, c.model)
}

// - Park() - prints "[brand] [model] is parked!"
func (c Car) Park() {
	fmt.Printf("%s %s is parked!\n", c.brand, c.model)
}

// Define a struct called 'Motorcycle' that embeds Engine and has brand (string) field
type Motorcycle struct {
	Engine
	brand string
}

// Define a method on Motorcycle:
// - Ride() - prints "Riding the [brand] motorcycle!"
func (m Motorcycle) Ride() {
	fmt.Printf("Riding the %s motorcycle!\n", m.brand)
}

func main() {
	// Create a Car with:
	// - Engine: horsepower=200, fuelType="Gasoline"  
	// - brand="Toyota", model="Camry"
	car := Car{
		Engine: Engine{horsepower: 200, fuelType: "Gasoline"},
		brand:  "Toyota",
		model:  "Camry",
	}
	
	// Call Car methods: Drive() and Park()
	car.Drive()
	car.Park()
	
	// Call embedded Engine methods through the Car: Start(), Info(), Stop()
	car.Start()
	car.Info()
	car.Stop()
	
	fmt.Println() // blank line for separation
	
	// Create a Motorcycle with:
	// - Engine: horsepower=100, fuelType="Gasoline"
	// - brand="Harley-Davidson"
	motorcycle := Motorcycle{
		Engine: Engine{horsepower: 100, fuelType: "Gasoline"},
		brand:  "Harley-Davidson",
	}
	
	// Call Motorcycle method: Ride()
	motorcycle.Ride()
	
	// Call embedded Engine methods through the Motorcycle: Start(), Info(), Stop()
	motorcycle.Start()
	motorcycle.Info()
	motorcycle.Stop()
}