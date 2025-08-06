// GoForGo Exercise: Apache Spark Basics
// Learn how to connect to Apache Spark from Go and perform basic operations

package main

import (
	"context"
	"fmt"
	"log"
)

// TODO: Define a SparkSession struct to represent our Spark connection
// Fields: 
// - URL (string) - Spark master URL
// - AppName (string) - Application name
// - Context (context.Context) - Go context for operations
type SparkSession struct {
	// Your SparkSession struct here
}

// TODO: Define a DataFrame struct to represent distributed data
// Fields:
// - Name (string) - DataFrame name/identifier  
// - Schema ([]string) - Column names
// - Data ([][]interface{}) - Sample data (simplified for learning)
type DataFrame struct {
	// Your DataFrame struct here
}

// TODO: Create a NewSparkSession function
// Parameters: masterURL, appName string
// Returns: *SparkSession, error
// Should initialize a SparkSession with the provided parameters
func NewSparkSession(masterURL, appName string) (*SparkSession, error) {
	// Your NewSparkSession implementation here
	return nil, nil
}

// TODO: Create a method to read CSV data (simulated)
// Method signature: (s *SparkSession) ReadCSV(filename string) (*DataFrame, error)
// For this exercise, simulate reading CSV by creating sample data
// Create a DataFrame with columns: "id", "name", "age", "city"
// Sample data: 
// - Row 1: 1, "Alice", 25, "New York"
// - Row 2: 2, "Bob", 30, "San Francisco"  
// - Row 3: 3, "Charlie", 35, "Chicago"
func (s *SparkSession) ReadCSV(filename string) (*DataFrame, error) {
	// Your ReadCSV implementation here
	return nil, nil
}

// TODO: Create a method to show DataFrame contents
// Method signature: (df *DataFrame) Show(numRows int)
// Print the DataFrame schema and data in a formatted table
func (df *DataFrame) Show(numRows int) {
	// Your Show implementation here
}

// TODO: Create a method to count DataFrame rows
// Method signature: (df *DataFrame) Count() int
// Return the total number of rows in the DataFrame
func (df *DataFrame) Count() int {
	// Your Count implementation here
	return 0
}

// TODO: Create a method to filter DataFrame
// Method signature: (df *DataFrame) Filter(column, operator, value string) *DataFrame
// Support operators: "==", ">", "<", "!="
// For simplicity, work with the "age" column and integer comparisons
// Return a new DataFrame with filtered data
func (df *DataFrame) Filter(column, operator, value string) *DataFrame {
	// Your Filter implementation here
	return nil
}

func main() {
	// TODO: Create a Spark session
	// Use master URL: "local[*]" (local mode with all CPU cores)
	// App name: "GoSpark-BasicExample"

	// TODO: Read a simulated CSV file
	// Filename: "users.csv"

	// TODO: Show the DataFrame contents
	// Display first 10 rows

	// TODO: Print the total row count

	// TODO: Filter the DataFrame
	// Show users with age greater than 25

	// TODO: Show the filtered results
	// Display first 5 rows of filtered data

	fmt.Println("Spark basics operations completed!")
}