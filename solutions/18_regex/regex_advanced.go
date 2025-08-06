// regex_advanced.go
// Learn advanced regex patterns and text processing in Go

package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== Advanced Regular Expressions ===")
	
	// Sample text for various operations
	sampleText := `
		Contact Information:
		John Doe - john.doe@company.com - (555) 123-4567
		Jane Smith - jane.smith@example.org - +1-555-987-6543
		Bob Johnson - bob.johnson@test.net - 555.456.7890
		
		Log entries:
		2023-08-06 14:30:25 ERROR: Failed to connect to database
		2023-08-06 14:35:12 WARN: High memory usage detected (85%)
		2023-08-06 14:40:01 INFO: User login successful
		2023-08-06 14:45:33 DEBUG: Cache miss for key 'user_123'
		
		URLs:
		https://www.example.com/api/v1/users?id=123&active=true
		http://localhost:8080/health
		ftp://files.company.com/documents/report.pdf
		
		Code snippets:
		func main() { fmt.Println("Hello, World!") }
		var user = User{Name: "John", Age: 30, Email: "john@example.com"}
		if err != nil { return fmt.Errorf("failed: %w", err) }
	`
	
	fmt.Println("=== Email Extraction and Validation ===")
	
	// Create regex for email validation
	emailRegex, err := regexp.Compile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	if err != nil {
		fmt.Printf("Email regex compile error: %v\n", err)
		return
	}
	
	// Find all emails
	emails := emailRegex.FindAllString(sampleText, -1)
	fmt.Printf("Found %d emails:\n", len(emails))
	for i, email := range emails {
		fmt.Printf("  %d. %s\n", i+1, email)
		
		// Validate each email
		isValid := emailRegex.MatchString(email)
		fmt.Printf("      Valid: %t\n", isValid)
	}
	
	fmt.Println("\n=== Phone Number Processing ===")
	
	// Create regex for various phone number formats
	phoneRegex, err := regexp.Compile(`(\+?1[-.\s]?)?\(?([0-9]{3})\)?[-.\s]?([0-9]{3})[-.\s]?([0-9]{4})`)
	if err != nil {
		fmt.Printf("Phone regex compile error: %v\n", err)
		return
	}
	
	// Find and normalize phone numbers
	phoneMatches := phoneRegex.FindAllStringSubmatch(sampleText, -1)
	fmt.Printf("Found %d phone numbers:\n", len(phoneMatches))
	for i, match := range phoneMatches {
		fullMatch := match[0]
		fmt.Printf("  %d. Original: %s\n", i+1, fullMatch)
		
		// Extract and normalize phone number
		normalized := normalizePhoneNumber(fullMatch)
		fmt.Printf("      Normalized: %s\n", normalized)
	}
	
	fmt.Println("\n=== Log Entry Parsing ===")
	
	// Create regex for log entry parsing with named groups
	logRegex, err := regexp.Compile(`(?P<timestamp>\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) (?P<level>\w+): (?P<message>.+)`)
	if err != nil {
		fmt.Printf("Log regex compile error: %v\n", err)
		return
	}
	
	// Parse log entries
	logLines := strings.Split(sampleText, "\n")
	fmt.Println("Parsed log entries:")
	for _, line := range logLines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		// Check if line matches log format
		if logRegex.MatchString(line) {
			// Extract named groups
			matches := logRegex.FindStringSubmatch(line)
			if len(matches) > 0 {
				// Create map of named groups
				result := make(map[string]string)
				names := logRegex.SubexpNames()
				for i, name := range names {
					if i != 0 && name != "" {
						result[name] = matches[i]
					}
				}
				
				fmt.Printf("  Timestamp: %s, Level: %s, Message: %s\n",
					result["timestamp"], result["level"], result["message"])
			}
		}
	}
	
	fmt.Println("\n=== URL Analysis ===")
	
	// Create comprehensive URL regex
	urlRegex, err := regexp.Compile(`(https?|ftp)://([^/\s]+)(/[^\s]*)?`)
	if err != nil {
		fmt.Printf("URL regex compile error: %v\n", err)
		return
	}
	
	// Find and analyze URLs
	urls := urlRegex.FindAllString(sampleText, -1)
	fmt.Printf("Found %d URLs:\n", len(urls))
	for i, url := range urls {
		fmt.Printf("  %d. %s\n", i+1, url)
		
		// Parse URL components
		matches := urlRegex.FindStringSubmatch(url)
		if len(matches) >= 4 {
			protocol := matches[1]
			domain := matches[2]
			path := matches[3]
			
			fmt.Printf("      Protocol: %s, Domain: %s, Path: %s\n", protocol, domain, path)
		}
	}
	
	fmt.Println("\n=== Code Pattern Matching ===")
	
	// Find function definitions
	funcRegex, err := regexp.Compile(`func\s+\w+\s*\([^)]*\)\s*\{[^}]*\}`)
	if err != nil {
		fmt.Printf("Function regex compile error: %v\n", err)
		return
	}
	
	functions := funcRegex.FindAllString(sampleText, -1)
	fmt.Printf("Found %d function definitions:\n", len(functions))
	for i, fn := range functions {
		fmt.Printf("  %d. %s\n", i+1, fn)
	}
	
	// Find variable declarations
	varRegex, err := regexp.Compile(`var\s+\w+\s*=\s*[^;\n]+`)
	if err != nil {
		fmt.Printf("Variable regex compile error: %v\n", err)
		return
	}
	
	variables := varRegex.FindAllString(sampleText, -1)
	fmt.Printf("Found %d variable declarations:\n", len(variables))
	for i, v := range variables {
		fmt.Printf("  %d. %s\n", i+1, v)
	}
	
	fmt.Println("\n=== Text Replacement and Cleaning ===")
	
	// Redact sensitive information
	redactedText := sampleText
	
	// Replace emails with [EMAIL]
	redactedText = emailRegex.ReplaceAllString(redactedText, "[EMAIL]")
	
	// Replace phone numbers with [PHONE]
	redactedText = phoneRegex.ReplaceAllString(redactedText, "[PHONE]")
	
	// Replace URLs with [URL]
	redactedText = urlRegex.ReplaceAllString(redactedText, "[URL]")
	
	fmt.Println("Redacted text:")
	fmt.Println(redactedText)
	
	fmt.Println("\n=== Advanced Pattern Matching ===")
	
	// Find words that start with capital letter and are followed by punctuation
	capitalWordRegex, err := regexp.Compile(`\b[A-Z][a-z]*\b(?=[.,:;!?])`)
	if err != nil {
		fmt.Printf("Capital word regex compile error: %v\n", err)
		return
	}
	
	capitalWords := capitalWordRegex.FindAllString(sampleText, -1)
	fmt.Printf("Capitalized words before punctuation: %v\n", capitalWords)
	
	// Find repeated words
	repeatedWordRegex, err := regexp.Compile(`\b(\w+)\s+\1\b`)
	if err != nil {
		fmt.Printf("Repeated word regex compile error: %v\n", err)
		return
	}
	
	testText := "This is is a test test with with repeated repeated words words."
	repeatedWords := repeatedWordRegex.FindAllString(testText, -1)
	fmt.Printf("Repeated words in test text: %v\n", repeatedWords)
	
	// Extract hashtags and mentions (like Twitter)
	socialText := "Check out this #golang tutorial! Thanks @john_doe for the help. #programming #learning"
	hashtagRegex, _ := regexp.Compile(`#\w+`)
	mentionRegex, _ := regexp.Compile(`@\w+`)
	
	hashtags := hashtagRegex.FindAllString(socialText, -1)
	mentions := mentionRegex.FindAllString(socialText, -1)
	
	fmt.Printf("Hashtags: %v\n", hashtags)
	fmt.Printf("Mentions: %v\n", mentions)
	
	fmt.Println("\n=== Performance Testing ===")
	
	// Test regex performance with large text
	largeText := strings.Repeat(sampleText, 1000)
	
	// Compiled regex (better performance)
	start := time.Now()
	matches := emailRegex.FindAllString(largeText, -1)
	elapsed := time.Since(start)
	
	fmt.Printf("Compiled regex found %d matches in %v\n", len(matches), elapsed)
	
	// Compare with string operations
	start = time.Now()
	count := strings.Count(largeText, "@")
	elapsed = time.Since(start)
	
	fmt.Printf("String count found %d '@' characters in %v\n", count, elapsed)
	
	fmt.Println("\n=== Custom Validation Functions ===")
	
	testInputs := []string{
		"john.doe@example.com",
		"invalid-email",
		"(555) 123-4567",
		"555-invalid",
		"https://www.example.com",
		"not-a-url",
	}
	
	for _, input := range testInputs {
		fmt.Printf("Input: %-25s Email: %-5t Phone: %-5t URL: %t\n",
			input,
			isValidEmail(input),
			isValidPhone(input),
			isValidURL(input))
	}
}

// Implement phone number normalization
func normalizePhoneNumber(phone string) string {
	// Remove all non-digit characters except +
	digitRegex, _ := regexp.Compile(`[^\d+]`)
	normalized := digitRegex.ReplaceAllString(phone, "")
	
	// Format as standard US number if 10 digits
	if len(normalized) == 10 {
		return fmt.Sprintf("(%s) %s-%s", normalized[:3], normalized[3:6], normalized[6:])
	} else if len(normalized) == 11 && strings.HasPrefix(normalized, "1") {
		return fmt.Sprintf("1-(%s) %s-%s", normalized[1:4], normalized[4:7], normalized[7:])
	}
	
	return normalized
}

// Implement email validation function
func isValidEmail(email string) bool {
	// Create and use email validation regex
	emailRegex, _ := regexp.Compile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// Implement phone validation function
func isValidPhone(phone string) bool {
	// Create and use phone validation regex
	phoneRegex, _ := regexp.Compile(`^(\+?1[-.\s]?)?\(?([0-9]{3})\)?[-.\s]?([0-9]{3})[-.\s]?([0-9]{4})$`)
	return phoneRegex.MatchString(phone)
}

// Implement URL validation function
func isValidURL(url string) bool {
	// Create and use URL validation regex
	urlRegex, _ := regexp.Compile(`^(https?|ftp)://[^\s/$.?#].[^\s]*$`)
	return urlRegex.MatchString(url)
}