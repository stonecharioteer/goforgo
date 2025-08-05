// regex_basics.go - SOLUTION
// Learn regular expressions in Go

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("=== Basic Pattern Matching ===")
	
	// Compile basic regex patterns
	emailPattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	phonePattern := regexp.MustCompile(`\d{3}-\d{3}-\d{4}`)
	
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
		isEmail := emailPattern.MatchString(s)
		isPhone := phonePattern.MatchString(s)
		fmt.Printf("'%s' - Email: %t, Phone: %t\n", s, isEmail, isPhone)
	}
	
	fmt.Println("\n=== Finding Matches ===")
	
	text := "Contact us at info@company.com or call 555-123-4567. You can also email support@help.org or call 123-456-7890."
	
	// Find all email matches
	emails := emailPattern.FindAllString(text, -1)
	fmt.Printf("Found emails: %v\n", emails)
	
	// Find all phone matches
	phones := phonePattern.FindAllString(text, -1)
	fmt.Printf("Found phones: %v\n", phones)
	
	fmt.Println("\n=== Submatches and Groups ===")
	
	// Pattern with capture groups
	namePattern := regexp.MustCompile(`Name: (\w+) (\w+)`)
	
	nameText := "Name: John Doe, Age: 30, Name: Jane Smith, Age: 25"
	
	// Find all submatches
	matches := namePattern.FindAllStringSubmatch(nameText, -1)
	
	fmt.Println("Name matches:")
	for i, match := range matches {
		fmt.Printf("Match %d: Full='%s', First='%s', Last='%s'\n", 
			i+1, match[0], match[1], match[2])
	}
	
	fmt.Println("\n=== String Replacement ===")
	
	// Replace patterns
	original := "The price is $10.99 and the tax is $1.50."
	
	// Replace all dollar amounts with "PRICE"
	pricePattern := regexp.MustCompile(`\$\d+\.\d{2}`)
	replaced := pricePattern.ReplaceAllString(original, "PRICE")
	
	fmt.Printf("Original: %s\n", original)
	fmt.Printf("Replaced: %s\n", replaced)
	
	// Replace with capture groups
	dateText := "Today is 2023-12-25 and tomorrow is 2023-12-26."
	datePattern := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	
	// Replace YYYY-MM-DD with MM/DD/YYYY
	formattedDates := datePattern.ReplaceAllString(dateText, "$2/$3/$1")
	
	fmt.Printf("Original dates: %s\n", dateText)
	fmt.Printf("Formatted dates: %s\n", formattedDates)
	
	fmt.Println("\n=== Splitting with Regex ===")
	
	// Split string using regex
	sentence := "apple,banana;cherry:grape|orange"
	
	separatorPattern := regexp.MustCompile(`[,;:|]`)
	fruits := separatorPattern.Split(sentence, -1)
	
	fmt.Printf("Original: %s\n", sentence)
	fmt.Printf("Split fruits: %v\n", fruits)
	
	fmt.Println("\n=== Validation Examples ===")
	
	// Create validation patterns
	patterns := map[string]*regexp.Regexp{
		"username": regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`),
		"password": regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d@$!%*?&]{8,}$`),
		"url":      regexp.MustCompile(`^https?://[\w.-]+\.[a-zA-Z]{2,}(/.*)?$`),
		"zipcode":  regexp.MustCompile(`^\d{5}(-\d{4})?$`),
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
		fmt.Printf("\n%s validation:\n", strings.Title(patternName))
		for _, test := range testData[patternName] {
			valid := pattern.MatchString(test)
			fmt.Printf("  '%s': %t\n", test, valid)
		}
	}
	
	fmt.Println("\n=== Advanced Features ===")
	
	// Case-insensitive matching
	text2 := "Hello WORLD and hello universe"
	casePattern := regexp.MustCompile(`(?i)hello`)
	
	matches2 := casePattern.FindAllString(text2, -1)
	fmt.Printf("Case-insensitive 'hello' matches: %v\n", matches2)
	
	// Non-greedy matching
	htmlText := "<div>content1</div><div>content2</div>"
	
	// Greedy: matches the whole string
	greedyPattern := regexp.MustCompile(`<div>.*</div>`)
	greedyMatch := greedyPattern.FindString(htmlText)
	
	// Non-greedy: matches each div separately  
	nonGreedyPattern := regexp.MustCompile(`<div>.*?</div>`)
	nonGreedyMatches := nonGreedyPattern.FindAllString(htmlText, -1)
	
	fmt.Printf("HTML text: %s\n", htmlText)
	fmt.Printf("Greedy match: %s\n", greedyMatch)
	fmt.Printf("Non-greedy matches: %v\n", nonGreedyMatches)
}