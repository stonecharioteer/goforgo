// GoForGo Solution: Apache Spark Basics
// Complete implementation of Spark connection and basic operations

package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// SparkSession represents our Spark connection
type SparkSession struct {
	URL     string
	AppName string
	Context context.Context
}

// DataFrame represents distributed data
type DataFrame struct {
	Name   string
	Schema []string
	Data   [][]interface{}
}

// NewSparkSession creates a new Spark session
func NewSparkSession(masterURL, appName string) (*SparkSession, error) {
	session := &SparkSession{
		URL:     masterURL,
		AppName: appName,
		Context: context.Background(),
	}
	
	log.Printf("Created Spark session: %s with master: %s", appName, masterURL)
	return session, nil
}

// ReadCSV simulates reading CSV data and creates a DataFrame
func (s *SparkSession) ReadCSV(filename string) (*DataFrame, error) {
	log.Printf("Reading CSV file: %s", filename)
	
	// Simulate CSV data
	df := &DataFrame{
		Name:   filename,
		Schema: []string{"id", "name", "age", "city"},
		Data: [][]interface{}{
			{1, "Alice", 25, "New York"},
			{2, "Bob", 30, "San Francisco"},
			{3, "Charlie", 35, "Chicago"},
		},
	}
	
	log.Printf("Successfully loaded DataFrame with %d rows and %d columns", len(df.Data), len(df.Schema))
	return df, nil
}

// Show displays DataFrame contents in a formatted table
func (df *DataFrame) Show(numRows int) {
	fmt.Printf("\nDataFrame: %s\n", df.Name)
	fmt.Println(strings.Repeat("-", 50))
	
	// Print header
	for i, col := range df.Schema {
		if i > 0 {
			fmt.Print(" | ")
		}
		fmt.Printf("%-12s", col)
	}
	fmt.Println()
	fmt.Println(strings.Repeat("-", 50))
	
	// Print data rows
	rowsToShow := numRows
	if rowsToShow > len(df.Data) {
		rowsToShow = len(df.Data)
	}
	
	for i := 0; i < rowsToShow; i++ {
		for j, value := range df.Data[i] {
			if j > 0 {
				fmt.Print(" | ")
			}
			fmt.Printf("%-12v", value)
		}
		fmt.Println()
	}
	
	if numRows < len(df.Data) {
		fmt.Printf("... showing %d of %d rows\n", numRows, len(df.Data))
	}
	fmt.Println()
}

// Count returns the total number of rows in the DataFrame
func (df *DataFrame) Count() int {
	return len(df.Data)
}

// Filter creates a new DataFrame with filtered data
func (df *DataFrame) Filter(column, operator, value string) *DataFrame {
	// Find column index
	colIndex := -1
	for i, col := range df.Schema {
		if col == column {
			colIndex = i
			break
		}
	}
	
	if colIndex == -1 {
		log.Printf("Column '%s' not found", column)
		return df
	}
	
	// Parse the comparison value
	compareValue, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Error parsing value '%s': %v", value, err)
		return df
	}
	
	// Filter data
	var filteredData [][]interface{}
	for _, row := range df.Data {
		rowValue, ok := row[colIndex].(int)
		if !ok {
			continue
		}
		
		match := false
		switch operator {
		case "==":
			match = rowValue == compareValue
		case ">":
			match = rowValue > compareValue
		case "<":
			match = rowValue < compareValue
		case "!=":
			match = rowValue != compareValue
		}
		
		if match {
			filteredData = append(filteredData, row)
		}
	}
	
	// Create filtered DataFrame
	filteredDF := &DataFrame{
		Name:   fmt.Sprintf("%s_filtered_%s%s%s", df.Name, column, operator, value),
		Schema: df.Schema,
		Data:   filteredData,
	}
	
	log.Printf("Filtered DataFrame: %d rows match condition '%s %s %s'", len(filteredData), column, operator, value)
	return filteredDF
}

func main() {
	// Create a Spark session
	spark, err := NewSparkSession("local[*]", "GoSpark-BasicExample")
	if err != nil {
		log.Fatal("Failed to create Spark session:", err)
	}
	
	// Read a simulated CSV file
	df, err := spark.ReadCSV("users.csv")
	if err != nil {
		log.Fatal("Failed to read CSV:", err)
	}
	
	// Show the DataFrame contents
	fmt.Println("=== Original DataFrame ===")
	df.Show(10)
	
	// Print the total row count
	fmt.Printf("Total rows: %d\n", df.Count())
	
	// Filter the DataFrame
	filteredDF := df.Filter("age", ">", "25")
	
	// Show the filtered results
	fmt.Println("=== Filtered DataFrame (age > 25) ===")
	filteredDF.Show(5)
	fmt.Printf("Filtered rows: %d\n", filteredDF.Count())
	
	fmt.Println("Spark basics operations completed!")
}