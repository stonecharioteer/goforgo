package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func demonstrateRuneVsByteGotchasFixed() {
	fmt.Println("=== Rune vs Byte - EXPLAINED ===")
	
	// Key concept: Go strings are UTF-8 encoded byte sequences
	text := "Hello 世界" // Contains Chinese characters (multi-byte UTF-8)
	
	fmt.Printf("String: %s\n", text)
	fmt.Printf("len(text): %d bytes\n", len(text))                    // Byte count (13 bytes)
	fmt.Printf("Character count: %d runes\n", len([]rune(text)))      // Rune count (8 characters)
	fmt.Printf("UTF-8 rune count: %d\n", utf8.RuneCountInString(text)) // More efficient way (8)
	
	// EXPLAINED: '世' and '界' are 3 bytes each in UTF-8
	
	fmt.Println("\n=== String Indexing - EXPLAINED ===")
	
	// WRONG: Byte-based iteration breaks multi-byte characters
	fmt.Println("Byte-based iteration (WRONG for Unicode):")
	for i := 0; i < len(text); i++ {
		b := text[i] // Returns byte, not rune
		if utf8.ValidString(string(b)) {
			fmt.Printf("text[%d] = %d, char: %c (valid)\n", i, b, b)
		} else {
			fmt.Printf("text[%d] = %d (invalid UTF-8 byte)\n", i, b)
		}
	}
	
	// CORRECT: Rune-based iteration
	fmt.Println("Rune-based iteration (CORRECT):")
	for i, r := range text {
		fmt.Printf("Byte pos %d: rune %c (U+%04X)\n", i, r, r)
	}
	
	fmt.Println("\n=== String Slicing - SAFE APPROACH ===")
	
	// UNSAFE: Byte-based slicing can break UTF-8
	if len(text) > 8 {
		unsafeSlice := text[:8] // Might cut multi-byte character
		fmt.Printf("Unsafe slice: %s\n", unsafeSlice)
		fmt.Printf("Is valid UTF-8: %t\n", utf8.ValidString(unsafeSlice))
	}
	
	// SAFE: Rune-aware slicing
	safeSlice := safeStringSlice(text, 6) // Get first 6 characters
	fmt.Printf("Safe slice: %s\n", safeSlice)
	fmt.Printf("Is valid UTF-8: %t\n", utf8.ValidString(safeSlice))
	
	fmt.Println("\n=== Rune vs Byte Literals - EXPLAINED ===")
	
	var byteValue byte = 'A'     // byte is alias for uint8 (0-255)
	var runeValue rune = 'A'     // rune is alias for int32 (Unicode code point)
	var runeUnicode rune = '世'   // Unicode code point (19990)
	
	fmt.Printf("Byte 'A': %d, type: %T\n", byteValue, byteValue)
	fmt.Printf("Rune 'A': %d, type: %T\n", runeValue, runeValue)
	fmt.Printf("Rune '世': %d, type: %T\n", runeUnicode, runeUnicode)
	
	// FIXED: Safe conversion checks
	if runeUnicode <= 255 {
		fmt.Printf("Safe conversion to byte: %d\n", byte(runeUnicode))
	} else {
		fmt.Printf("Cannot safely convert rune %d to byte (would overflow)\n", runeUnicode)
	}
}

func safeStringSlice(s string, maxRunes int) string {
	// FIXED: Safe string slicing that respects UTF-8 boundaries
	runes := []rune(s)
	if len(runes) <= maxRunes {
		return s
	}
	return string(runes[:maxRunes])
}

func stringManipulationFixed() {
	fmt.Println("\n=== String Manipulation - FIXED ===")
	
	text := "café" // Contains accented character (2-byte UTF-8)
	fmt.Printf("Original: %s (len: %d bytes, %d runes)\n", 
		text, len(text), utf8.RuneCountInString(text))
	
	// FIXED: Proper string reversal
	reversed := reverseStringCorrect(text)
	fmt.Printf("Correct reverse: %s\n", reversed)
	fmt.Printf("Is valid UTF-8: %t\n", utf8.ValidString(reversed))
	
	// FIXED: Proper first character capitalization
	capitalized := capitalizeFirstCorrect(text)
	fmt.Printf("Correct capitalize: %s\n", capitalized)
	
	// FIXED: Safe string truncation
	truncated := truncateStringCorrect(text, 3)
	fmt.Printf("Correct truncate: %s\n", truncated)
	fmt.Printf("Is valid UTF-8: %t\n", utf8.ValidString(truncated))
}

func reverseStringCorrect(s string) string {
	// FIXED: Reverse runes, not bytes
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func capitalizeFirstCorrect(s string) string {
	// FIXED: Handle Unicode characters properly
	if s == "" {
		return s
	}
	
	runes := []rune(s)
	if len(runes) > 0 {
		runes[0] = unicode.ToUpper(runes[0])
	}
	return string(runes)
}

func truncateStringCorrect(s string, maxRunes int) string {
	// FIXED: Truncate by rune count, not byte count
	runes := []rune(s)
	if len(runes) <= maxRunes {
		return s
	}
	return string(runes[:maxRunes])
}

func unicodeNormalizationExplained() {
	fmt.Println("\n=== Unicode Normalization - EXPLAINED ===")
	
	// These look identical but are encoded differently!
	str1 := "é" // Single codepoint: U+00E9 (LATIN SMALL LETTER E WITH ACUTE)
	str2 := "é" // Combining sequence: U+0065 (e) + U+0301 (combining acute accent)
	
	fmt.Printf("str1: %s (bytes: %d, runes: %d)\n", str1, len(str1), len([]rune(str1)))
	fmt.Printf("str2: %s (bytes: %d, runes: %d)\n", str2, len(str2), len([]rune(str2)))
	fmt.Printf("Byte-wise equal? %t\n", str1 == str2) // false!
	
	// Show the actual code points
	fmt.Println("str1 runes:")
	for i, r := range str1 {
		fmt.Printf("  [%d]: U+%04X\n", i, r)
	}
	
	fmt.Println("str2 runes:")
	for i, r := range str2 {
		fmt.Printf("  [%d]: U+%04X\n", i, r)
	}
	
	// SOLUTION: Use unicode normalization for comparison (not shown in this basic example)
	// In real applications, you'd use golang.org/x/text/unicode/norm package
}

func practicalExamples() {
	fmt.Println("\n=== Practical Examples ===")
	
	// Example 1: Counting characters for UI display
	username := "用户名123" // Mixed Chinese and ASCII
	fmt.Printf("Username: %s\n", username)
	fmt.Printf("Display width (characters): %d\n", len([]rune(username))) // 6 characters
	fmt.Printf("Storage size (bytes): %d\n", len(username))               // 11 bytes
	
	// Example 2: Safe substring extraction
	longText := "这是一个很长的中文字符串"
	preview := safeStringSlice(longText, 8) // First 8 characters
	fmt.Printf("Preview: %s...\n", preview)
	
	// Example 3: Character-based validation
	fmt.Printf("Contains only letters: %t\n", isAllLetters(username))
}

func isAllLetters(s string) bool {
	// CORRECT: Check each rune, not each byte
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func main() {
	demonstrateRuneVsByteGotchasFixed()
	stringManipulationFixed()
	unicodeNormalizationExplained()
	practicalExamples()
	
	fmt.Println("\n=== KEY TAKEAWAYS ===")
	fmt.Println("1. Go strings are UTF-8 encoded byte sequences")
	fmt.Println("2. len(string) returns byte count, not character count")
	fmt.Println("3. string[i] returns a byte, not a character")
	fmt.Println("4. Use []rune(string) for character-based operations")
	fmt.Println("5. Range over string iterates over runes automatically")
	fmt.Println("6. Always validate UTF-8 when manipulating byte slices")
	fmt.Println("7. Unicode normalization affects string equality")
}