// regexp_advanced.go - SOLUTION
// Learn advanced regular expression patterns, named groups, and replacements

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("=== Advanced Regular Expressions ===")
	
	fmt.Println("\n=== Named Capture Groups ===")
	
	// Create regex pattern with named groups for parsing dates
	datePattern := regexp.MustCompile(`^(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})$`)
	testDates := []string{
		"2023-12-25",
		"2024-01-01", 
		"invalid-date",
		"2023-02-29", // Invalid leap year date
	}
	
	fmt.Println("Parsing dates with named groups:")
	for _, date := range testDates {
		if matches := datePattern.FindStringSubmatch(date); matches != nil {
			names := datePattern.SubexpNames()
			for i, match := range matches[1:] { // Skip full match
				fmt.Printf("  %s: %s = %s\n", date, names[i+1], match)
			}
		} else {
			fmt.Printf("  %s: Invalid date format\n", date)
		}
	}
	
	fmt.Println("\n=== Email Validation and Parsing ===")
	
	// Create regex for email validation with named groups
	emailPattern := regexp.MustCompile(`^(?P<username>[a-zA-Z0-9._%+-]+)@(?P<domain>[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})$`)
	testEmails := []string{
		"user@example.com",
		"first.last+tag@subdomain.example.co.uk",
		"invalid.email",
		"test@localhost",
		"user.name@example-domain.com",
	}
	
	fmt.Println("Validating and parsing emails:")
	for _, email := range testEmails {
		if matches := emailPattern.FindStringSubmatch(email); matches != nil {
			names := emailPattern.SubexpNames()
			fmt.Printf("  %s: Valid\n", email)
			for i, match := range matches[1:] {
				fmt.Printf("    %s = %s\n", names[i+1], match)
			}
		} else {
			fmt.Printf("  %s: Invalid email format\n", email)
		}
	}
	
	fmt.Println("\n=== URL Parsing with Named Groups ===")
	
	// Create regex for URL parsing
	urlPattern := regexp.MustCompile(`^(?P<protocol>https?|ftp)://(?P<host>[^:/?#]+)(?::(?P<port>\d+))?(?P<path>/[^?#]*)?(?:\?(?P<query>[^#]*))?(?:#(?P<fragment>.*))?$`)
	testURLs := []string{
		"https://www.example.com:8080/path/to/resource?param=value#section",
		"http://localhost:3000/api/users",
		"ftp://files.example.com/download/file.txt",
		"invalid-url",
		"https://api.github.com/repos/golang/go/issues?state=open",
	}
	
	fmt.Println("Parsing URLs:")
	for _, url := range testURLs {
		if matches := urlPattern.FindStringSubmatch(url); matches != nil {
			names := urlPattern.SubexpNames()
			fmt.Printf("  %s:\n", url)
			for i, match := range matches[1:] {
				if match != "" {
					fmt.Printf("    %s = %s\n", names[i+1], match)
				}
			}
		} else {
			fmt.Printf("  %s: Invalid URL format\n", url)
		}
	}
	
	fmt.Println("\n=== Advanced Replacements ===")
	
	// Phone number formatting
	phonePattern := regexp.MustCompile(`\(?(\d{3})\)?[-.\s]*(\d{3})[-.\s]*(\d{4})`)
	phoneNumbers := []string{
		"1234567890",
		"123-456-7890", 
		"(123) 456-7890",
		"123.456.7890",
		"+1 (123) 456-7890",
	}
	
	fmt.Println("Formatting phone numbers:")
	for _, phone := range phoneNumbers {
		formatted := phonePattern.ReplaceAllString(phone, "($1) $2-$3")
		fmt.Printf("  %s → %s\n", phone, formatted)
	}
	
	fmt.Println("\n=== Text Processing with Replacements ===")
	
	text := `The quick brown fox jumps over the lazy dog. 
The dog was very lazy, but the fox was quick and brown.
Another fox appeared, and both foxes ran quickly.`
	
	// Replace words with replacements using capture groups
	animalPattern := regexp.MustCompile(`(?i)\b(fox|dog)(es)?\b`)
	processedText := animalPattern.ReplaceAllStringFunc(text, func(match string) string {
		lower := strings.ToLower(match)
		if strings.Contains(lower, "fox") {
			if strings.HasSuffix(lower, "es") {
				return "FOXES"
			}
			return "FOX"
		} else if strings.Contains(lower, "dog") {
			if strings.HasSuffix(lower, "s") {
				return "DOGS"
			}
			return "DOG"
		}
		return strings.ToUpper(match)
	})
	
	fmt.Println("Original text:")
	fmt.Println(text)
	fmt.Println("\nProcessed text:")
	fmt.Println(processedText)
	
	fmt.Println("\n=== Log Parsing ===")
	
	logEntries := []string{
		"2024-01-15 10:30:45 [INFO] User john.doe logged in from 192.168.1.100",
		"2024-01-15 10:31:02 [ERROR] Database connection failed: timeout after 30s",
		"2024-01-15 10:31:15 [WARN] High memory usage: 85% of 8GB used",
		"2024-01-15 10:32:01 [DEBUG] Processing request ID: req-12345-abcde",
		"Invalid log entry without proper format",
	}
	
	// Create regex for parsing log entries
	logPattern := regexp.MustCompile(`^(?P<timestamp>\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) \[(?P<level>\w+)\] (?P<message>.+)$`)
	
	fmt.Println("Parsing log entries:")
	for _, entry := range logEntries {
		if matches := logPattern.FindStringSubmatch(entry); matches != nil {
			names := logPattern.SubexpNames()
			fmt.Printf("  Log entry parsed:\n")
			for i, match := range matches[1:] {
				fmt.Printf("    %s: %s\n", names[i+1], match)
			}
		} else {
			fmt.Printf("  Invalid log format: %s\n", entry)
		}
	}
	
	fmt.Println("\n=== HTML Tag Extraction ===")
	
	htmlContent := `<div class="container">
		<h1 id="title">Welcome to Go</h1>
		<p class="description">This is a <strong>powerful</strong> language.</p>
		<a href="https://golang.org" target="_blank">Learn Go</a>
		<img src="logo.png" alt="Go Logo" width="100" height="50">
	</div>`
	
	// Extract different types of HTML tags
	tagPattern := regexp.MustCompile(`<[^>]+>`)
	linkPattern := regexp.MustCompile(`<a[^>]+href="([^"]+)"[^>]*>([^<]*)</a>`)
	imgPattern := regexp.MustCompile(`<img[^>]+src="([^"]+)"[^>]*(?:alt="([^"]*)")?[^>]*(?:width="(\d+)")?[^>]*(?:height="(\d+)")?[^>]*>`)
	
	fmt.Println("Extracting HTML elements:")
	
	// Find all HTML tags
	tags := tagPattern.FindAllString(htmlContent, -1)
	fmt.Printf("All tags found: %d\n", len(tags))
	for _, tag := range tags {
		fmt.Printf("  %s\n", tag)
	}
	
	// Extract links
	links := linkPattern.FindAllStringSubmatch(htmlContent, -1)
	fmt.Println("\nLinks found:")
	for _, link := range links {
		fmt.Printf("  URL: %s, Text: %s\n", link[1], link[2])
	}
	
	// Extract images
	images := imgPattern.FindAllStringSubmatch(htmlContent, -1)
	fmt.Println("\nImages found:")
	for _, img := range images {
		fmt.Printf("  Src: %s", img[1])
		if img[2] != "" {
			fmt.Printf(", Alt: %s", img[2])
		}
		if img[3] != "" && img[4] != "" {
			fmt.Printf(", Size: %sx%s", img[3], img[4])
		}
		fmt.Println()
	}
	
	fmt.Println("\n=== Data Validation ===")
	
	// Validate various data formats
	creditCardPattern := regexp.MustCompile(`^\d{4}-?\d{4}-?\d{4}-?\d{4}$`)
	ipAddressPattern := regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
	passwordPattern := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)
	
	testData := map[string][]string{
		"Credit Cards": {
			"4532-1234-5678-9012", // Visa
			"5555-5555-5555-4444", // MasterCard
			"invalid-card-number",
			"4111111111111111",     // Visa test number
		},
		"IP Addresses": {
			"192.168.1.1",
			"10.0.0.255",
			"256.1.1.1",    // Invalid
			"192.168.1",    // Incomplete
			"127.0.0.1",
		},
		"Passwords": {
			"StrongPass123!",
			"weakpass",
			"NoNumbers!",
			"nonumbers123",
			"ValidPass1@",
		},
	}
	
	fmt.Println("Data validation:")
	
	// Validate credit cards
	fmt.Println("\nCredit Card validation:")
	for _, card := range testData["Credit Cards"] {
		valid := creditCardPattern.MatchString(card)
		status := "❌ Invalid"
		if valid {
			status = "✅ Valid"
		}
		fmt.Printf("  %s: %s\n", card, status)
	}
	
	// Validate IP addresses
	fmt.Println("\nIP Address validation:")
	for _, ip := range testData["IP Addresses"] {
		valid := ipAddressPattern.MatchString(ip)
		status := "❌ Invalid"
		if valid {
			status = "✅ Valid"
		}
		fmt.Printf("  %s: %s\n", ip, status)
	}
	
	// Validate passwords (at least 8 chars, 1 upper, 1 lower, 1 digit, 1 special)
	fmt.Println("\nPassword strength validation:")
	for _, password := range testData["Passwords"] {
		valid := passwordPattern.MatchString(password)
		status := "❌ Weak"
		if valid {
			status = "✅ Strong"
		}
		fmt.Printf("  %s: %s\n", password, status)
	}
}