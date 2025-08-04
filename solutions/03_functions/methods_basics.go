package main

import "fmt"

// Define a struct called 'Rectangle' with width and height fields (both float64)
type Rectangle struct {
	width  float64
	height float64
}

// Define a method called 'Area' on Rectangle that returns the area (float64)
// Area = width * height
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Define a method called 'Perimeter' on Rectangle that returns the perimeter (float64)  
// Perimeter = 2 * (width + height)
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Define a struct called 'Circle' with radius field (float64)
type Circle struct {
	radius float64
}

// Define a method called 'Area' on Circle that returns the area (float64)
// Use 3.14159 as pi, Area = pi * radius * radius
func (c Circle) Area() float64 {
	const pi = 3.14159
	return pi * c.radius * c.radius
}

func main() {
	// Create a Rectangle with width=5.0 and height=3.0
	rect := Rectangle{width: 5.0, height: 3.0}
	
	// Call the Area method on the rectangle and print the result
	rectArea := rect.Area()
	fmt.Println("Rectangle area:", rectArea)
	
	// Call the Perimeter method on the rectangle and print the result
	rectPerimeter := rect.Perimeter()
	fmt.Println("Rectangle perimeter:", rectPerimeter)
	
	// Create a Circle with radius=4.0
	circle := Circle{radius: 4.0}
	
	// Call the Area method on the circle and print the result
	circleArea := circle.Area()
	fmt.Println("Circle area:", circleArea)
}