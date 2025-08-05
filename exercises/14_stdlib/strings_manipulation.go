// strings_manipulation.go
// Learn the strings package for string manipulation

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== String Inspection ===")
	
	text := "Hello, Go Programming World!"
	
	// TODO: Basic string inspection
	fmt.Printf("Length: %d\\n", /* get length of text */)
	fmt.Printf("Contains 'Go': %t\\n", /* check if text contains "Go" */)
	fmt.Printf("Contains 'Python': %t\\n", /* check if text contains "Python" */)
	fmt.Printf("Starts with 'Hello': %t\\n", /* check if text starts with "Hello" */)
	fmt.Printf("Ends with 'World!': %t\\n", /* check if text ends with "World!" */)
	
	fmt.Println("\\n=== String Searching ===")
	
	// TODO: String searching
	fmt.Printf("Index of 'Go': %d\\n", /* find index of "Go" in text */)
	fmt.Printf("Index of 'Python': %d\\n", /* find index of "Python" in text */)
	fmt.Printf("Last index of 'o': %d\\n", /* find last index of "o" in text */)
	fmt.Printf("Count of 'o': %d\\n", /* count occurrences of "o" in text */)
	
	fmt.Println("\\n=== String Modification ===")
	
	// TODO: Case conversion
	fmt.Printf("Upper case: %s\\n", /* convert text to uppercase */)
	fmt.Printf("Lower case: %s\\n", /* convert text to lowercase */)
	fmt.Printf("Title case: %s\\n", /* convert text to title case */)
	
	// TODO: String replacement
	replaced := /* replace "Go" with "Golang" in text */
	fmt.Printf("Replaced: %s\\n", replaced)
	
	replacedAll := /* replace all "o" with "0" in text */
	fmt.Printf("Replaced all 'o' with '0': %s\\n", replacedAll)
	
	fmt.Println("\\n=== String Splitting and Joining ===")
	
	// TODO: String splitting
	words := /* split text by spaces */
	fmt.Printf("Words: %v\\n", words)
	fmt.Printf("Number of words: %d\\n", len(words))
	
	csvData := "apple,banana,cherry,date"
	fruits := /* split csvData by commas */
	fmt.Printf("Fruits: %v\\n", fruits)
	
	// TODO: String joining
	joined := /* join fruits with " | " */
	fmt.Printf("Joined fruits: %s\\n", joined)
	
	numbers := []string{"1", "2", "3", "4", "5"}
	numberString := /* join numbers with "-" */
	fmt.Printf("Number string: %s\\n", numberString)
	
	fmt.Println("\\n=== String Trimming ===")
	
	messyText := "   \\n\\t  Hello World!  \\t\\n   "
	
	// TODO: Various trimming operations
	fmt.Printf("Original: '%s'\\n", messyText)
	fmt.Printf("TrimSpace: '%s'\\n", /* trim whitespace from messyText */)
	fmt.Printf("Trim '!': '%s'\\n", /* trim '!' from "Hello World!" */)
	fmt.Printf("TrimPrefix 'Hello': '%s'\\n", /* trim prefix "Hello" from "Hello World!" */)
	fmt.Printf("TrimSuffix '!': '%s'\\n", /* trim suffix "!" from "Hello World!" */)
	
	fmt.Println("\\n=== String Building ===")
	
	// TODO: Efficient string building with strings.Builder
	var builder strings.Builder
	
	// Write different types to builder
	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("World")
	builder.WriteByte('!')
	builder.WriteRune('🌍')
	
	result := /* get string from builder */
	fmt.Printf("Built string: %s\\n", result)
	fmt.Printf("Builder length: %d\\n", /* get builder length */)
	
	// TODO: Build string with loop
	var loopBuilder strings.Builder
	for i := 1; i <= 5; i++ {
		/* write "Step " + i + " " to loopBuilder */
	}
	fmt.Printf("Loop result: %s\\n", loopBuilder.String())
	
	fmt.Println("\\n=== String Comparison ===")
	
	str1 := "Hello"
	str2 := "hello"
	str3 := "Hello"
	
	// TODO: String comparison
	fmt.Printf("'%s' == '%s': %t\\n", str1, str2, str1 == str2)
	fmt.Printf("'%s' == '%s': %t\\n", str1, str3, str1 == str3)
	fmt.Printf("Compare '%s' and '%s': %d\\n", str1, str2, /* compare str1 and str2 */)
	fmt.Printf("EqualFold '%s' and '%s': %t\\n", str1, str2, /* case-insensitive compare */)
	
	fmt.Println("\\n=== String Fields and Cleaning ===")
	
	messyData := "  apple   banana    cherry   date  "
	
	// TODO: Extract fields (split by any whitespace)
	fields := /* get fields from messyData */
	fmt.Printf("Fields: %v\\n", fields)
	fmt.Printf("Number of fields: %d\\n", len(fields))
	
	// TODO: Clean and process each field
	var cleanFields []string
	for _, field := range fields {
		// Trim whitespace and convert to title case
		clean := /* trim space and convert to title case */
		cleanFields = append(cleanFields, clean)
	}
	fmt.Printf("Clean fields: %v\\n", cleanFields)
}