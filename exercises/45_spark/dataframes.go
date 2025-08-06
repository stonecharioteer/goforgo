// GoForGo Exercise: Spark DataFrames
// Learn advanced DataFrame operations: joins, aggregations, and transformations

package main

import (
	"context"
	"fmt"
	"log"
)

// Reuse basic types from previous exercise
type SparkSession struct {
	URL     string
	AppName string
	Context context.Context
}

type DataFrame struct {
	Name   string
	Schema []string
	Data   [][]interface{}
}

func NewSparkSession(masterURL, appName string) (*SparkSession, error) {
	return &SparkSession{
		URL:     masterURL,
		AppName: appName,
		Context: context.Background(),
	}, nil
}

// TODO: Create a method to select specific columns from DataFrame
// Method signature: (df *DataFrame) Select(columns ...string) *DataFrame
// Return a new DataFrame with only the specified columns
func (df *DataFrame) Select(columns ...string) *DataFrame {
	// Your Select implementation here
	return nil
}

// TODO: Create a method to add a computed column
// Method signature: (df *DataFrame) WithColumn(newColName string, computation func(row []interface{}) interface{}) *DataFrame
// Add a new column based on computation function applied to each row
func (df *DataFrame) WithColumn(newColName string, computation func(row []interface{}) interface{}) *DataFrame {
	// Your WithColumn implementation here
	return nil
}

// TODO: Create a method to group data by a column
// Method signature: (df *DataFrame) GroupBy(column string) *GroupedDataFrame
// Return a GroupedDataFrame that can perform aggregations
type GroupedDataFrame struct {
	// Your GroupedDataFrame struct here
}

func (df *DataFrame) GroupBy(column string) *GroupedDataFrame {
	// Your GroupBy implementation here
	return nil
}

// TODO: Create aggregation methods on GroupedDataFrame
// Method signature: (gdf *GroupedDataFrame) Count() *DataFrame
// Count rows in each group
func (gdf *GroupedDataFrame) Count() *DataFrame {
	// Your Count implementation here
	return nil
}

// TODO: Create a method to join two DataFrames
// Method signature: (df *DataFrame) Join(other *DataFrame, joinColumn string) *DataFrame
// Perform inner join on the specified column
func (df *DataFrame) Join(other *DataFrame, joinColumn string) *DataFrame {
	// Your Join implementation here
	return nil
}

// TODO: Create a method to sort DataFrame
// Method signature: (df *DataFrame) OrderBy(column string, ascending bool) *DataFrame
// Sort DataFrame by the specified column
func (df *DataFrame) OrderBy(column string, ascending bool) *DataFrame {
	// Your OrderBy implementation here
	return nil
}

func (df *DataFrame) Show(numRows int) {
	fmt.Printf("\nDataFrame: %s\n", df.Name)
	// Simplified show for exercise
	for _, row := range df.Data[:min(numRows, len(df.Data))] {
		fmt.Println(row)
	}
	fmt.Println()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// TODO: Create Spark session
	spark, _ := NewSparkSession("local[*]", "GoSpark-DataFrames")

	// TODO: Create sample employees DataFrame
	// Schema: "id", "name", "department", "salary"
	// Data:
	// - 1, "Alice", "Engineering", 75000
	// - 2, "Bob", "Marketing", 60000
	// - 3, "Charlie", "Engineering", 80000
	// - 4, "Diana", "Sales", 65000
	employees := &DataFrame{
		Name:   "employees",
		Schema: []string{"id", "name", "department", "salary"},
		Data: [][]interface{}{
			// Add your data here
		},
	}

	// TODO: Create departments DataFrame
	// Schema: "dept_name", "location"
	// Data:
	// - "Engineering", "San Francisco"
	// - "Marketing", "New York"
	// - "Sales", "Chicago"
	departments := &DataFrame{
		// Your departments DataFrame here
	}

	// TODO: Demonstrate Select operation
	// Select only "name" and "salary" columns

	// TODO: Demonstrate WithColumn operation
	// Add "salary_category" column based on salary:
	// - "High" if salary >= 70000
	// - "Medium" if salary >= 60000
	// - "Low" otherwise

	// TODO: Demonstrate GroupBy and aggregation
	// Group by department and count employees

	// TODO: Demonstrate Join operation
	// Join employees with departments on department/dept_name

	// TODO: Demonstrate OrderBy operation
	// Sort employees by salary in descending order

	fmt.Println("Spark DataFrame operations completed!")
}