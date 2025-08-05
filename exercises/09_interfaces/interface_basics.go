// interface_basics.go
// Learn the fundamentals of interfaces in Go
// Interfaces define behavior through method signatures

package main

import (
	"fmt"
	"math"
)

// TODO: Define a Shape interface with Area() method
type Shape interface {
	// Define Area method signature
}

// TODO: Define a Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// TODO: Implement Shape interface for Rectangle
func (r Rectangle) Area() float64 {
	// Calculate rectangle area
}

// TODO: Define a Circle struct  
type Circle struct {
	Radius float64
}

// TODO: Implement Shape interface for Circle
func (c Circle) Area() float64 {
	// Calculate circle area
}

// TODO: Define a Writer interface (similar to io.Writer)
type Writer interface {
	// Define Write method that takes []byte and returns (int, error)
}

// TODO: Define a FileWriter struct that implements Writer
type FileWriter struct {
	Filename string
}

func (fw FileWriter) Write(data []byte) (int, error) {
	// Simulate writing to file
	fmt.Printf("Writing %d bytes to file: %s\n", len(data), fw.Filename)
	fmt.Printf("Content: %s\n", string(data))
	return len(data), nil
}

// TODO: Define a ConsoleWriter struct that implements Writer
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	// Write to console
	fmt.Printf("Console: %s", string(data))
	return len(data), nil
}

// TODO: Function that accepts any Shape
func printShapeInfo(s Shape) {
	// Print the area of any shape
	fmt.Printf("Shape area: %.2f\n", /* call Area method */)
}

// TODO: Function that accepts any Writer
func writeMessage(w Writer, message string) {
	// Use the writer to write a message
	// Convert message to []byte and call Write
}

func main() {
	// TODO: Create shapes and use them through the interface
	rect := // Create Rectangle with Width: 5, Height: 3
	circle := // Create Circle with Radius: 4
	
	// TODO: Call printShapeInfo with different shapes
	fmt.Println("Using shapes through interface:")
	// Call printShapeInfo with rect
	// Call printShapeInfo with circle
	
	// TODO: Store shapes in a slice of Shape interface
	shapes := []Shape{/* add rect and circle */}
	
	var totalArea float64
	for _, shape := range shapes {
		// Add each shape's area to totalArea
	}
	fmt.Printf("Total area of all shapes: %.2f\n", totalArea)
	
	// TODO: Use Writer interface with different implementations
	fileWriter := // Create FileWriter with Filename: "output.txt"
	consoleWriter := // Create ConsoleWriter
	
	fmt.Println("\nUsing writers through interface:")
	// Write "Hello, File!" using fileWriter
	// Write "Hello, Console!" using consoleWriter
	
	// TODO: Store writers in a slice and use them
	writers := []Writer{/* add fileWriter and consoleWriter */}
	
	for i, writer := range writers {
		message := fmt.Sprintf("Message %d from writer\n", i+1)
		// Write message using current writer
	}
	
	// TODO: Demonstrate interface assignment
	var shape Shape
	shape = rect // Rectangle implements Shape
	fmt.Printf("Shape (Rectangle) area: %.2f\n", shape.Area())
	
	shape = circle // Circle also implements Shape  
	fmt.Printf("Shape (Circle) area: %.2f\n", shape.Area())
}