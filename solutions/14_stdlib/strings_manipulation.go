// strings_manipulation.go - SOLUTION
// Learn the strings package for string manipulation

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== String Inspection ===")
	
	text := "Hello, Go Programming World!"
	
	// Basic string inspection
	fmt.Printf("Length: %d\n", len(text))
	fmt.Printf("Contains 'Go': %t\n", strings.Contains(text, "Go"))
	fmt.Printf("Contains 'Python': %t\n", strings.Contains(text, "Python"))
	fmt.Printf("Starts with 'Hello': %t\n", strings.HasPrefix(text, "Hello"))
	fmt.Printf("Ends with 'World!': %t\n", strings.HasSuffix(text, "World!"))
	
	fmt.Println("\n=== String Searching ===")
	
	// String searching
	fmt.Printf("Index of 'Go': %d\n", strings.Index(text, "Go"))
	fmt.Printf("Index of 'Python': %d\n", strings.Index(text, "Python"))
	fmt.Printf("Last index of 'o': %d\n", strings.LastIndex(text, "o"))
	fmt.Printf("Count of 'o': %d\n", strings.Count(text, "o"))
	
	fmt.Println("\n=== String Modification ===")
	
	// Case conversion
	fmt.Printf("Upper case: %s\n", strings.ToUpper(text))
	fmt.Printf("Lower case: %s\n", strings.ToLower(text))
	fmt.Printf("Title case: %s\n", strings.Title(text))
	
	// String replacement
	replaced := strings.Replace(text, "Go", "Golang", 1)
	fmt.Printf("Replaced: %s\n", replaced)
	
	replacedAll := strings.ReplaceAll(text, "o", "0")
	fmt.Printf("Replaced all 'o' with '0': %s\n", replacedAll)
	
	fmt.Println("\n=== String Splitting and Joining ===")
	
	// String splitting
	words := strings.Split(text, " ")
	fmt.Printf("Words: %v\n", words)
	fmt.Printf("Number of words: %d\n", len(words))
	
	csvData := "apple,banana,cherry,date"
	fruits := strings.Split(csvData, ",")
	fmt.Printf("Fruits: %v\n", fruits)
	
	// String joining
	joined := strings.Join(fruits, " | ")
	fmt.Printf("Joined fruits: %s\n", joined)
	
	numbers := []string{"1", "2", "3", "4", "5"}
	numberString := strings.Join(numbers, "-")
	fmt.Printf("Number string: %s\n", numberString)
	
	fmt.Println("\n=== String Trimming ===")
	
	messyText := "   \n\t  Hello World!  \t\n   "
	
	// Various trimming operations
	fmt.Printf("Original: '%s'\n", messyText)
	fmt.Printf("TrimSpace: '%s'\n", strings.TrimSpace(messyText))
	fmt.Printf("Trim '!': '%s'\n", strings.Trim("Hello World!", "!"))
	fmt.Printf("TrimPrefix 'Hello': '%s'\n", strings.TrimPrefix("Hello World!", "Hello"))
	fmt.Printf("TrimSuffix '!': '%s'\n", strings.TrimSuffix("Hello World!", "!"))
	
	fmt.Println("\n=== String Building ===")
	
	// Efficient string building with strings.Builder
	var builder strings.Builder
	
	// Write different types to builder
	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World")
	builder.WriteByte('!')
	builder.WriteRune('🌍')
	
	result := builder.String()
	fmt.Printf("Built string: %s\n", result)
	fmt.Printf("Builder length: %d\n", builder.Len())
	
	// Build string with loop
	var loopBuilder strings.Builder
	for i := 1; i <= 5; i++ {
		loopBuilder.WriteString(fmt.Sprintf("Step %d ", i))
	}
	fmt.Printf("Loop result: %s\n", loopBuilder.String())
	
	fmt.Println("\n=== String Comparison ===")
	
	str1 := "Hello"
	str2 := "hello"
	str3 := "Hello"
	
	// String comparison
	fmt.Printf("'%s' == '%s': %t\n", str1, str2, str1 == str2)
	fmt.Printf("'%s' == '%s': %t\n", str1, str3, str1 == str3)
	fmt.Printf("Compare '%s' and '%s': %d\n", str1, str2, strings.Compare(str1, str2))
	fmt.Printf("EqualFold '%s' and '%s': %t\n", str1, str2, strings.EqualFold(str1, str2))
	
	fmt.Println("\n=== String Fields and Cleaning ===")
	
	messyData := "  apple   banana    cherry   date  "
	
	// Extract fields (split by any whitespace)
	fields := strings.Fields(messyData)
	fmt.Printf("Fields: %v\n", fields)
	fmt.Printf("Number of fields: %d\n", len(fields))
	
	// Clean and process each field
	var cleanFields []string
	for _, field := range fields {
		// Trim whitespace and convert to title case
		clean := strings.Title(strings.TrimSpace(field))
		cleanFields = append(cleanFields, clean)
	}
	fmt.Printf("Clean fields: %v\n", cleanFields)
}