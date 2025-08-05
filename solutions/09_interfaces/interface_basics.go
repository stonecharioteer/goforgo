// interface_basics.go - SOLUTION
// Learn the fundamentals of interfaces in Go
// Interfaces define behavior through method signatures

package main

import (
	"fmt"
	"math"
)

// Shape interface with Area() method
type Shape interface {
	Area() float64
}

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Implement Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle struct
type Circle struct {
	Radius float64
}

// Implement Shape interface for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Writer interface (similar to io.Writer)
type Writer interface {
	Write([]byte) (int, error)
}

// FileWriter struct that implements Writer
type FileWriter struct {
	Filename string
}

func (fw FileWriter) Write(data []byte) (int, error) {
	// Simulate writing to file
	fmt.Printf("Writing %d bytes to file: %s\n", len(data), fw.Filename)
	fmt.Printf("Content: %s\n", string(data))
	return len(data), nil
}

// ConsoleWriter struct that implements Writer
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	// Write to console
	fmt.Printf("Console: %s", string(data))
	return len(data), nil
}

// Function that accepts any Shape
func printShapeInfo(s Shape) {
	fmt.Printf("Shape area: %.2f\n", s.Area())
}

// Function that accepts any Writer
func writeMessage(w Writer, message string) {
	// Convert message to []byte and call Write
	w.Write([]byte(message))
}

func main() {
	// Create shapes and use them through the interface
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}
	
	// Call printShapeInfo with different shapes
	fmt.Println("Using shapes through interface:")
	printShapeInfo(rect)
	printShapeInfo(circle)
	
	// Store shapes in a slice of Shape interface
	shapes := []Shape{rect, circle}
	
	var totalArea float64
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	fmt.Printf("Total area of all shapes: %.2f\n", totalArea)
	
	// Use Writer interface with different implementations
	fileWriter := FileWriter{Filename: "output.txt"}
	consoleWriter := ConsoleWriter{}
	
	fmt.Println("\nUsing writers through interface:")
	writeMessage(fileWriter, "Hello, File!")
	writeMessage(consoleWriter, "Hello, Console!")
	
	// Store writers in a slice and use them
	writers := []Writer{fileWriter, consoleWriter}
	
	for i, writer := range writers {
		message := fmt.Sprintf("Message %d from writer\n", i+1)
		writeMessage(writer, message)
	}
	
	// Demonstrate interface assignment
	var shape Shape
	shape = rect // Rectangle implements Shape
	fmt.Printf("Shape (Rectangle) area: %.2f\n", shape.Area())
	
	shape = circle // Circle also implements Shape
	fmt.Printf("Shape (Circle) area: %.2f\n", shape.Area())
}