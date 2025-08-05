// regex_basics.go
// Learn regular expressions in Go

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("=== Basic Pattern Matching ===")
	
	// TODO: Compile basic regex patterns
	emailPattern := /* compile regex for email pattern */
	phonePattern := /* compile regex for phone pattern: \\d{3}-\\d{3}-\\d{4} */
	
	if emailPattern == nil || phonePattern == nil {
		fmt.Println("Failed to compile regex patterns")
		return
	}
	
	// Test strings
	testStrings := []string{
		"john@example.com",
		"invalid-email",
		"555-123-4567",
		"not-a-phone",
		"alice@test.org",
		"123-456-7890",
	}
	
	fmt.Println("Testing email and phone patterns:")
	for _, s := range testStrings {
		isEmail := /* check if s matches emailPattern */
		isPhone := /* check if s matches phonePattern */
		fmt.Printf("'%s' - Email: %t, Phone: %t\\n", s, isEmail, isPhone)
	}
	
	fmt.Println("\\n=== Finding Matches ===")
	
	text := "Contact us at info@company.com or call 555-123-4567. You can also email support@help.org or call 123-456-7890."
	
	// TODO: Find all email matches
	emails := /* find all email matches in text */
	fmt.Printf("Found emails: %v\\n", emails)
	
	// TODO: Find all phone matches
	phones := /* find all phone matches in text */
	fmt.Printf("Found phones: %v\\n", phones)
	
	fmt.Println("\\n=== Submatches and Groups ===")
	
	// TODO: Pattern with capture groups
	namePattern := /* compile regex: "Name: (\\w+) (\\w+)" */
	
	nameText := "Name: John Doe, Age: 30, Name: Jane Smith, Age: 25"
	
	// TODO: Find all submatches
	matches := /* find all submatches in nameText */
	
	fmt.Println("Name matches:")
	for i, match := range matches {
		fmt.Printf("Match %d: Full='%s', First='%s', Last='%s'\\n", 
			i+1, match[0], match[1], match[2])
	}
	
	fmt.Println("\\n=== String Replacement ===")
	
	// TODO: Replace patterns
	original := "The price is $10.99 and the tax is $1.50."
	
	// Replace all dollar amounts with "PRICE"
	pricePattern := /* compile regex: "\\$\\d+\\.\\d{2}" */
	replaced := /* replace all matches with "PRICE" */
	
	fmt.Printf("Original: %s\\n", original)
	fmt.Printf("Replaced: %s\\n", replaced)
	
	// TODO: Replace with capture groups
	dateText := "Today is 2023-12-25 and tomorrow is 2023-12-26."
	datePattern := /* compile regex: "(\\d{4})-(\\d{2})-(\\d{2})" */
	
	// Replace YYYY-MM-DD with MM/DD/YYYY
	formattedDates := /* replace with "$2/$3/$1" */
	
	fmt.Printf("Original dates: %s\\n", dateText)
	fmt.Printf("Formatted dates: %s\\n", formattedDates)
	
	fmt.Println("\\n=== Splitting with Regex ===")
	
	// TODO: Split string using regex
	sentence := "apple,banana;cherry:grape|orange"
	
	separatorPattern := /* compile regex: "[,;:|]" */
	fruits := /* split sentence using separatorPattern */
	
	fmt.Printf("Original: %s\\n", sentence)
	fmt.Printf("Split fruits: %v\\n", fruits)
	
	fmt.Println("\\n=== Validation Examples ===")
	
	// TODO: Create validation patterns
	patterns := map[string]*regexp.Regexp{
		"username": /* compile: "^[a-zA-Z0-9_]{3,20}$" */,
		"password": /* compile: "^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)[a-zA-Z\\d@$!%*?&]{8,}$" */,
		"url":      /* compile: "^https?://[\\w.-]+\\.[a-zA-Z]{2,}(/.*)?$" */,
		"zipcode":  /* compile: "^\\d{5}(-\\d{4})?$" */,
	}
	
	testData := map[string][]string{
		"username": {"john_doe", "a", "user123", "invalid-user", "toolongusername123456"},
		"password": {"password", "Password123", "weak", "StrongPass1!", "12345678"},
		"url":      {"https://example.com", "http://test.org/path", "invalid-url", "ftp://files.com"},
		"zipcode":  {"12345", "12345-6789", "1234", "12345-67890", "abcde"},
	}
	
	for patternName, pattern := range patterns {
		if pattern == nil {
			continue
		}
		fmt.Printf("\\n%s validation:\\n", strings.Title(patternName))
		for _, test := range testData[patternName] {
			valid := /* check if test matches pattern */
			fmt.Printf("  '%s': %t\\n", test, valid)
		}
	}
	
	fmt.Println("\\n=== Advanced Features ===")
	
	// TODO: Case-insensitive matching
	text2 := "Hello WORLD and hello universe"
	casePattern := /* compile case-insensitive regex: "(?i)hello" */
	
	matches2 := /* find all matches in text2 */
	fmt.Printf("Case-insensitive 'hello' matches: %v\\n", matches2)
	
	// TODO: Non-greedy matching
	htmlText := "<div>content1</div><div>content2</div>"
	
	// Greedy: matches the whole string
	greedyPattern := /* compile: "<div>.*</div>" */
	greedyMatch := /* find first match in htmlText */
	
	// Non-greedy: matches each div separately  
	nonGreedyPattern := /* compile: "<div>.*?</div>" */
	nonGreedyMatches := /* find all matches in htmlText */
	
	fmt.Printf("HTML text: %s\\n", htmlText)
	fmt.Printf("Greedy match: %s\\n", greedyMatch)
	fmt.Printf("Non-greedy matches: %v\\n", nonGreedyMatches)
}