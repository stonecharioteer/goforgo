package main

import (
	"fmt"
	"unicode/utf8"
)

func demonstrateRuneVsByteGotchas() {
	fmt.Println("=== Rune vs Byte Gotchas ===")
	
	// Gotcha 1: String length confusion
	text := "Hello 世界" // Contains Chinese characters (multi-byte UTF-8)
	
	fmt.Printf("String: %s\n", text)
	fmt.Printf("len(text): %d\n", len(text)) // This returns byte count, not character count!
	
	// What's the actual character count?
	fmt.Printf("Character count: %d\n", len([]rune(text)))
	fmt.Printf("UTF-8 rune count: %d\n", utf8.RuneCountInString(text))
	
	// Gotcha 2: String indexing returns bytes, not characters
	fmt.Println("\n=== String Indexing Gotcha ===")
	
	// This will cause issues with multi-byte characters
	for i := 0; i < len(text); i++ {
		// text[i] returns a byte, not a rune!
		fmt.Printf("text[%d] = %d (byte), char: %c\n", i, text[i], text[i])
		// This will print garbage for multi-byte UTF-8 sequences
	}
	
	// Gotcha 3: String slicing can break UTF-8 sequences
	fmt.Println("\n=== String Slicing Gotcha ===")
	
	// Dangerous: This might slice in the middle of a UTF-8 sequence
	if len(text) > 8 {
		broken := text[:8] // This might cut a multi-byte character in half!
		fmt.Printf("Broken slice: %s\n", broken)
		fmt.Printf("Is valid UTF-8: %t\n", utf8.ValidString(broken))
	}
	
	// Gotcha 4: Rune literal vs byte literal confusion
	fmt.Println("\n=== Rune vs Byte Literals ===")
	
	var byteValue byte = 'A'     // Single quotes for byte (ASCII only)
	var runeValue rune = 'A'     // Single quotes for rune (Unicode code point)
	var runeUnicode rune = '世'   // Unicode rune
	
	fmt.Printf("Byte 'A': %d, type: %T\n", byteValue, byteValue)
	fmt.Printf("Rune 'A': %d, type: %T\n", runeValue, runeValue)
	fmt.Printf("Rune '世': %d, type: %T\n", runeUnicode, runeUnicode)
	
	// What happens when you convert between them?
	// This conversion has bugs - fix them
	wrongConversion := byte(runeUnicode) // This will overflow!
	fmt.Printf("Wrong conversion: %d\n", wrongConversion)
	
	// Gotcha 5: Range over string vs range over []byte
	fmt.Println("\n=== Range Iteration Differences ===")
	
	fmt.Println("Range over string (iterates over runes):")
	for i, r := range text {
		fmt.Printf("Index %d: rune %c (U+%04X)\n", i, r, r)
	}
	
	fmt.Println("Range over []byte (iterates over bytes):")
	for i, b := range []byte(text) {
		fmt.Printf("Index %d: byte %d\n", i, b)
	}
	
	// Notice how the indices are different!
}

func stringManipulationProblems() {
	fmt.Println("\n=== String Manipulation Problems ===")
	
	text := "café" // Contains accented character (2-byte UTF-8)
	fmt.Printf("Original: %s (len: %d bytes, %d runes)\n", 
		text, len(text), utf8.RuneCountInString(text))
	
	// Problem 1: Trying to reverse a string byte by byte
	reversed := reverseStringWrong(text)
	fmt.Printf("Wrong reverse: %s\n", reversed)
	fmt.Printf("Is valid UTF-8: %t\n", utf8.ValidString(reversed))
	
	// Problem 2: Trying to capitalize first character
	capitalized := capitalizeFirstWrong(text)
	fmt.Printf("Wrong capitalize: %s\n", capitalized)
	
	// Problem 3: Truncating string safely
	truncated := truncateStringWrong(text, 3)
	fmt.Printf("Wrong truncate: %s\n", truncated)
	fmt.Printf("Is valid UTF-8: %t\n", utf8.ValidString(truncated))
}

func reverseStringWrong(s string) string {
	// This function has a bug - it reverses bytes, not runes
	bytes := []byte(s)
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}

func capitalizeFirstWrong(s string) string {
	// This function has a bug - it assumes first character is ASCII
	if len(s) == 0 {
		return s
	}
	
	// This breaks for non-ASCII first characters
	bytes := []byte(s)
	if bytes[0] >= 'a' && bytes[0] <= 'z' {
		bytes[0] = bytes[0] - 'a' + 'A'
	}
	return string(bytes)
}

func truncateStringWrong(s string, maxLen int) string {
	// This function has a bug - it truncates by bytes, potentially breaking UTF-8
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] // This can break UTF-8 sequences!
}

func unicodeNormalizationIssues() {
	fmt.Println("\n=== Unicode Normalization Issues ===")
	
	// These look the same but are different!
	str1 := "é" // Single codepoint (U+00E9)
	str2 := "é" // Combining sequence (U+0065 U+0301)
	
	fmt.Printf("str1: %s (len: %d, runes: %d)\n", str1, len(str1), len([]rune(str1)))
	fmt.Printf("str2: %s (len: %d, runes: %d)\n", str2, len(str2), len([]rune(str2)))
	fmt.Printf("Equal? %t\n", str1 == str2)
	
	// They look the same but len() returns different values!
	// This is a common source of bugs in internationalized applications
}

func main() {
	demonstrateRuneVsByteGotchas()
	stringManipulationProblems()
	unicodeNormalizationIssues()
	
	fmt.Println("\n=== Fix the bugs in the string manipulation functions! ===")
	fmt.Println("Remember: Go strings are UTF-8 encoded, not ASCII!")
}