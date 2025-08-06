// GoForGo Solution: Spark DataFrames
// Complete implementation of advanced DataFrame operations

package main

import (
	"context"
	"fmt"
	"log"
	"sort"
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

type GroupedDataFrame struct {
	SourceDF    *DataFrame
	GroupColumn string
	Groups      map[interface{}][]int // value -> row indices
}

func NewSparkSession(masterURL, appName string) (*SparkSession, error) {
	log.Printf("Created Spark session: %s with master: %s", appName, masterURL)
	return &SparkSession{
		URL:     masterURL,
		AppName: appName,
		Context: context.Background(),
	}, nil
}

// Select returns a new DataFrame with only the specified columns
func (df *DataFrame) Select(columns ...string) *DataFrame {
	// Find column indices
	columnIndices := make([]int, 0, len(columns))
	for _, col := range columns {
		for i, schemaCol := range df.Schema {
			if schemaCol == col {
				columnIndices = append(columnIndices, i)
				break
			}
		}
	}
	
	// Create new data with selected columns
	newData := make([][]interface{}, len(df.Data))
	for i, row := range df.Data {
		newRow := make([]interface{}, len(columnIndices))
		for j, colIndex := range columnIndices {
			newRow[j] = row[colIndex]
		}
		newData[i] = newRow
	}
	
	return &DataFrame{
		Name:   fmt.Sprintf("%s_selected", df.Name),
		Schema: columns,
		Data:   newData,
	}
}

// WithColumn adds a computed column to the DataFrame
func (df *DataFrame) WithColumn(newColName string, computation func(row []interface{}) interface{}) *DataFrame {
	newSchema := make([]string, len(df.Schema)+1)
	copy(newSchema, df.Schema)
	newSchema[len(df.Schema)] = newColName
	
	newData := make([][]interface{}, len(df.Data))
	for i, row := range df.Data {
		newRow := make([]interface{}, len(row)+1)
		copy(newRow, row)
		newRow[len(row)] = computation(row)
		newData[i] = newRow
	}
	
	return &DataFrame{
		Name:   fmt.Sprintf("%s_with_%s", df.Name, newColName),
		Schema: newSchema,
		Data:   newData,
	}
}

// GroupBy groups data by a column
func (df *DataFrame) GroupBy(column string) *GroupedDataFrame {
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
		return nil
	}
	
	// Group rows by column value
	groups := make(map[interface{}][]int)
	for i, row := range df.Data {
		value := row[colIndex]
		groups[value] = append(groups[value], i)
	}
	
	return &GroupedDataFrame{
		SourceDF:    df,
		GroupColumn: column,
		Groups:      groups,
	}
}

// Count returns count of rows in each group
func (gdf *GroupedDataFrame) Count() *DataFrame {
	data := make([][]interface{}, 0, len(gdf.Groups))
	for value, indices := range gdf.Groups {
		data = append(data, []interface{}{value, len(indices)})
	}
	
	return &DataFrame{
		Name:   fmt.Sprintf("%s_grouped_count", gdf.SourceDF.Name),
		Schema: []string{gdf.GroupColumn, "count"},
		Data:   data,
	}
}

// Join performs inner join on the specified column
func (df *DataFrame) Join(other *DataFrame, joinColumn string) *DataFrame {
	// Find join column indices
	leftIndex := -1
	rightIndex := -1
	
	for i, col := range df.Schema {
		if col == joinColumn {
			leftIndex = i
			break
		}
	}
	
	for i, col := range other.Schema {
		if col == joinColumn {
			rightIndex = i
			break
		}
	}
	
	if leftIndex == -1 || rightIndex == -1 {
		log.Printf("Join column '%s' not found in one of the DataFrames", joinColumn)
		return df
	}
	
	// Create index for right DataFrame
	rightIndex_map := make(map[interface{}][]int)
	for i, row := range other.Data {
		key := row[rightIndex]
		rightIndex_map[key] = append(rightIndex_map[key], i)
	}
	
	// Perform join
	var joinedData [][]interface{}
	for _, leftRow := range df.Data {
		leftKey := leftRow[leftIndex]
		if rightIndices, exists := rightIndex_map[leftKey]; exists {
			for _, rightIdx := range rightIndices {
				rightRow := other.Data[rightIdx]
				
				// Combine rows (excluding duplicate join column from right)
				joinedRow := make([]interface{}, len(leftRow)+len(rightRow)-1)
				copy(joinedRow, leftRow)
				
				copyIndex := len(leftRow)
				for i, val := range rightRow {
					if i != rightIndex {
						joinedRow[copyIndex] = val
						copyIndex++
					}
				}
				
				joinedData = append(joinedData, joinedRow)
			}
		}
	}
	
	// Create new schema
	newSchema := make([]string, len(df.Schema)+len(other.Schema)-1)
	copy(newSchema, df.Schema)
	
	copyIndex := len(df.Schema)
	for i, col := range other.Schema {
		if i != rightIndex {
			newSchema[copyIndex] = col
			copyIndex++
		}
	}
	
	return &DataFrame{
		Name:   fmt.Sprintf("%s_join_%s", df.Name, other.Name),
		Schema: newSchema,
		Data:   joinedData,
	}
}

// OrderBy sorts DataFrame by the specified column
func (df *DataFrame) OrderBy(column string, ascending bool) *DataFrame {
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
	
	// Create copy of data for sorting
	sortedData := make([][]interface{}, len(df.Data))
	copy(sortedData, df.Data)
	
	// Sort the data
	sort.Slice(sortedData, func(i, j int) bool {
		val1 := sortedData[i][colIndex]
		val2 := sortedData[j][colIndex]
		
		// Handle different types (simplified for integers)
		if v1, ok1 := val1.(int); ok1 {
			if v2, ok2 := val2.(int); ok2 {
				if ascending {
					return v1 < v2
				}
				return v1 > v2
			}
		}
		
		// Fallback to string comparison
		s1 := fmt.Sprintf("%v", val1)
		s2 := fmt.Sprintf("%v", val2)
		if ascending {
			return s1 < s2
		}
		return s1 > s2
	})
	
	return &DataFrame{
		Name:   fmt.Sprintf("%s_sorted_%s", df.Name, column),
		Schema: df.Schema,
		Data:   sortedData,
	}
}

func (df *DataFrame) Show(numRows int) {
	fmt.Printf("\nDataFrame: %s\n", df.Name)
	fmt.Println("Schema:", df.Schema)
	
	rowsToShow := min(numRows, len(df.Data))
	for i := 0; i < rowsToShow; i++ {
		fmt.Println(df.Data[i])
	}
	
	if numRows < len(df.Data) {
		fmt.Printf("... showing %d of %d rows\n", numRows, len(df.Data))
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
	// Create Spark session
	spark, _ := NewSparkSession("local[*]", "GoSpark-DataFrames")

	// Create sample employees DataFrame
	employees := &DataFrame{
		Name:   "employees",
		Schema: []string{"id", "name", "department", "salary"},
		Data: [][]interface{}{
			{1, "Alice", "Engineering", 75000},
			{2, "Bob", "Marketing", 60000},
			{3, "Charlie", "Engineering", 80000},
			{4, "Diana", "Sales", 65000},
		},
	}

	// Create departments DataFrame
	departments := &DataFrame{
		Name:   "departments",
		Schema: []string{"dept_name", "location"},
		Data: [][]interface{}{
			{"Engineering", "San Francisco"},
			{"Marketing", "New York"},
			{"Sales", "Chicago"},
		},
	}

	fmt.Println("=== Original Employees DataFrame ===")
	employees.Show(10)

	// Demonstrate Select operation
	fmt.Println("=== Select name and salary columns ===")
	selectedDF := employees.Select("name", "salary")
	selectedDF.Show(10)

	// Demonstrate WithColumn operation
	fmt.Println("=== Add salary_category column ===")
	salaryWithCategory := employees.WithColumn("salary_category", func(row []interface{}) interface{} {
		salary := row[3].(int) // salary is at index 3
		if salary >= 70000 {
			return "High"
		} else if salary >= 60000 {
			return "Medium"
		}
		return "Low"
	})
	salaryWithCategory.Show(10)

	// Demonstrate GroupBy and aggregation
	fmt.Println("=== Group by department and count ===")
	groupedDF := employees.GroupBy("department")
	countDF := groupedDF.Count()
	countDF.Show(10)

	// Demonstrate Join operation
	fmt.Println("=== Join employees with departments ===")
	// First, create a modified departments DataFrame to match join column name
	deptForJoin := &DataFrame{
		Name:   "departments_renamed",
		Schema: []string{"department", "location"},
		Data: [][]interface{}{
			{"Engineering", "San Francisco"},
			{"Marketing", "New York"},
			{"Sales", "Chicago"},
		},
	}
	
	joinedDF := employees.Join(deptForJoin, "department")
	joinedDF.Show(10)

	// Demonstrate OrderBy operation
	fmt.Println("=== Sort employees by salary (descending) ===")
	sortedDF := employees.OrderBy("salary", false)
	sortedDF.Show(10)

	fmt.Println("Spark DataFrame operations completed!")
}