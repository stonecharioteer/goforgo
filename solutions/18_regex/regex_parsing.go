// regex_parsing.go - SOLUTION  
// Learn advanced regex parsing and text extraction

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("=== Advanced Regex Parsing ===")
	
	fmt.Println("\n=== Log File Parsing ===")
	
	// Parse structured log entries
	logEntries := []string{
		"2023-12-25 10:30:45 [INFO] User alice logged in from 192.168.1.100",
		"2023-12-25 10:31:02 [ERROR] Failed to connect to database: connection timeout", 
		"2023-12-25 10:31:15 [WARN] High memory usage detected: 85%",
		"2023-12-25 10:32:00 [INFO] User bob performed action: purchase_item",
		"2023-12-25 10:32:30 [DEBUG] Cache miss for key: user_preferences_123",
	}
	
	// Create regex for log parsing
	logPattern := regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) \[(\w+)\] (.+)`)
	
	fmt.Println("Parsing log entries:")
	for i, entry := range logEntries {
		// Extract log components
		matches := logPattern.FindStringSubmatch(entry)
		
		if len(matches) >= 4 {
			timestamp := matches[1]
			level := matches[2]
			message := matches[3]
			
			fmt.Printf("Entry %d:\n", i+1)
			fmt.Printf("  Timestamp: %s\n", timestamp)
			fmt.Printf("  Level: %s\n", level)
			fmt.Printf("  Message: %s\n", message)
		} else {
			fmt.Printf("Entry %d: Failed to parse\n", i+1)
		}
		fmt.Println()
	}
	
	fmt.Println("=== Web Data Extraction ===")
	
	// Extract data from HTML-like content
	htmlContent := `
	<div class="user-profile">
		<h2>John Doe</h2>
		<p>Email: john.doe@example.com</p>
		<p>Phone: +1-555-123-4567</p>
		<a href="https://github.com/johndoe">GitHub Profile</a>
		<span class="location">San Francisco, CA</span>
	</div>
	<div class="user-profile">
		<h2>Jane Smith</h2>
		<p>Email: jane.smith@company.org</p>
		<p>Phone: +1-555-987-6543</p>
		<a href="https://linkedin.com/in/janesmith">LinkedIn Profile</a>
		<span class="location">New York, NY</span>
	</div>
	`
	
	// Create patterns for different data types
	namePattern := regexp.MustCompile(`<h2>([^<]+)</h2>`)
	emailPattern := regexp.MustCompile(`([a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,})`)
	phonePattern := regexp.MustCompile(`(\+1-\d{3}-\d{3}-\d{4})`)
	urlPattern := regexp.MustCompile(`href="([^"]+)"`)
	locationPattern := regexp.MustCompile(`<span class="location">([^<]+)</span>`)
	
	fmt.Println("Extracting web data:")
	
	// Extract names
	names := namePattern.FindAllStringSubmatch(htmlContent, -1)
	namesList := make([]string, 0, len(names))
	for _, match := range names {
		namesList = append(namesList, match[1])
	}
	fmt.Printf("Names found: %v\n", namesList)
	
	// Extract emails  
	emails := emailPattern.FindAllString(htmlContent, -1)
	fmt.Printf("Emails found: %v\n", emails)
	
	// Extract phones
	phones := phonePattern.FindAllString(htmlContent, -1)
	fmt.Printf("Phones found: %v\n", phones)
	
	// Extract URLs
	urlMatches := urlPattern.FindAllStringSubmatch(htmlContent, -1)
	urls := make([]string, 0, len(urlMatches))
	for _, match := range urlMatches {
		urls = append(urls, match[1])
	}
	fmt.Printf("URLs found: %v\n", urls)
	
	// Extract locations
	locationMatches := locationPattern.FindAllStringSubmatch(htmlContent, -1)
	locations := make([]string, 0, len(locationMatches))
	for _, match := range locationMatches {
		locations = append(locations, match[1])
	}
	fmt.Printf("Locations found: %v\n", locations)
	
	fmt.Println("\n=== CSV Data Parsing ===")
	
	// Parse CSV with complex fields
	csvData := `"Product Name","Price","Description","Tags"
"Laptop Computer","$1,299.99","High-performance laptop with 16GB RAM","electronics,computer,portable"
"Book: \"Go Programming\"","$49.95","Learn Go programming language","books,programming,go"
"T-Shirt (Size: L)","$19.99","Comfortable cotton t-shirt","clothing,casual,cotton"`
	
	// Create CSV parsing regex
	csvPattern := regexp.MustCompile(`"([^"]*(?:""[^"]*)*)"|([^,]+)`)
	
	fmt.Println("Parsing CSV data:")
	lines := strings.Split(csvData, "\n")
	
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		
		// Extract CSV fields
		matches := csvPattern.FindAllStringSubmatch(line, -1)
		fields := make([]string, 0, len(matches))
		for _, match := range matches {
			if match[1] != "" {
				fields = append(fields, match[1])
			} else {
				fields = append(fields, match[2])
			}
		}
		
		fmt.Printf("Line %d: Found %d fields\n", i+1, len(fields))
		for j, field := range fields {
			// Clean field (remove quotes)
			cleaned := strings.Trim(field, `"`)
			cleaned = strings.ReplaceAll(cleaned, `""`, `"`)
			fmt.Printf("  Field %d: %s\n", j+1, cleaned)
		}
		fmt.Println()
	}
	
	fmt.Println("=== Configuration File Parsing ===")
	
	// Parse configuration files
	configData := `
	# Database configuration
	db.host = localhost
	db.port = 5432
	db.username = admin
	db.password = "secret123"
	
	# Server settings
	server.port = 8080
	server.debug = true
	server.allowed_hosts = ["localhost", "127.0.0.1", "example.com"]
	
	# Feature flags
	features.enable_logging = true
	features.max_connections = 100
	`
	
	// Create pattern for config parsing
	configPattern := regexp.MustCompile(`^\s*([a-zA-Z0-9_.]+)\s*=\s*(.+)\s*$`)
	
	fmt.Println("Parsing configuration:")
	configLines := strings.Split(configData, "\n")
	
	config := make(map[string]string)
	
	for _, line := range configLines {
		line = strings.TrimSpace(line)
		
		// Skip comments and empty lines
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		
		// Extract key-value pairs
		matches := configPattern.FindStringSubmatch(line)
		if len(matches) >= 3 {
			key := matches[1]
			value := matches[2]
			
			// Clean up value (remove quotes, trim)
			value = strings.Trim(value, `"`)
			value = strings.TrimSpace(value)
			
			config[key] = value
			fmt.Printf("Config: %s = %s\n", key, value)
		}
	}
	
	fmt.Printf("\nParsed %d configuration items\n", len(config))
	
	fmt.Println("\n=== Advanced Pattern Matching ===")
	
	// Match complex patterns
	textData := `
	Contact support at support@company.com or call 1-800-555-0123.
	Our office hours are Monday-Friday 9:00 AM - 5:00 PM PST.
	Visit our website at https://www.company.com for more information.
	Follow us on Twitter @CompanyName or LinkedIn /company/company-name.
	`
	
	// Create patterns for different contact methods
	patterns := map[string]*regexp.Regexp{
		"email":      regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`),
		"phone":      regexp.MustCompile(`1-800-\d{3}-\d{4}`),
		"url":        regexp.MustCompile(`https?://[^\s]+`),
		"time":       regexp.MustCompile(`\d{1,2}:\d{2} [AP]M`),
		"social":     regexp.MustCompile(`@\w+|/company/[\w-]+`),
	}
	
	fmt.Println("Advanced pattern matching:")
	for patternName, pattern := range patterns {
		if pattern == nil {
			fmt.Printf("%s: Pattern compilation failed\n", patternName)
			continue
		}
		
		// Find all matches for each pattern
		matches := pattern.FindAllString(textData, -1)
		fmt.Printf("%s: %v\n", patternName, matches)
	}
	
	fmt.Println("\n=== Text Transformation ===")
	
	// Transform text using regex replacements
	sourceText := "The user john.doe@email.com has ID 12345 and phone 555-0123."
	
	fmt.Printf("Original: %s\n", sourceText)
	
	// Mask sensitive data
	emailMask := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	phoneMask := regexp.MustCompile(`\d{3}-\d{4}`)
	idMask := regexp.MustCompile(`\bID \d+`)
	
	// Apply transformations
	masked := sourceText
	masked = emailMask.ReplaceAllString(masked, "***@***.***")
	masked = phoneMask.ReplaceAllString(masked, "***-****")
	masked = idMask.ReplaceAllString(masked, "ID ****")
	
	fmt.Printf("Masked: %s\n", masked)
}