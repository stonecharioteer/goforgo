// pattern_matching.go - SOLUTION
// Learn string pattern matching algorithms and text processing techniques

package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== String Pattern Matching Algorithms ===")
	
	// Test data for pattern matching
	text := "ABABDABACDABABCABCABCABCABC"
	patterns := []string{"ABC", "ABAB", "CAB", "XYZ", "ABABCAB"}
	
	fmt.Printf("Text: %s\n", text)
	fmt.Printf("Text length: %d\n", len(text))
	
	fmt.Println("\n=== Pattern Matching Results ===")
	
	// Test different pattern matching algorithms
	for _, pattern := range patterns {
		fmt.Printf("\nSearching for pattern: '%s'\n", pattern)
		
		// Naive/Brute Force approach
		naiveResults := naivePatternMatch(text, pattern)
		fmt.Printf("  Naive Algorithm: %v\n", naiveResults)
		
		// KMP (Knuth-Morris-Pratt) algorithm
		kmpResults := kmpPatternMatch(text, pattern)
		fmt.Printf("  KMP Algorithm: %v\n", kmpResults)
		
		// Boyer-Moore algorithm (simplified)
		bmResults := boyerMoorePatternMatch(text, pattern)
		fmt.Printf("  Boyer-Moore: %v\n", bmResults)
		
		// Rabin-Karp algorithm (rolling hash)
		rkResults := rabinKarpPatternMatch(text, pattern)
		fmt.Printf("  Rabin-Karp: %v\n", rkResults)
		
		// Built-in Go string functions for comparison
		builtinResults := findAllOccurrences(text, pattern)
		fmt.Printf("  Built-in (Go): %v\n", builtinResults)
	}
	
	fmt.Println("\n=== Algorithm Performance Comparison ===")
	
	// Performance testing with larger text
	largeText := strings.Repeat("ABCDEFGHIJK", 1000) + "PATTERN" + strings.Repeat("LMNOPQRSTU", 1000)
	testPattern := "PATTERN"
	
	fmt.Printf("Testing performance with text length: %d\n", len(largeText))
	fmt.Printf("Searching for pattern: '%s'\n", testPattern)
	
	// Time each algorithm
	algorithms := map[string]func(string, string) []int{
		"Naive":       naivePatternMatch,
		"KMP":         kmpPatternMatch,
		"Boyer-Moore": boyerMoorePatternMatch,
		"Rabin-Karp":  rabinKarpPatternMatch,
	}
	
	for name, alg := range algorithms {
		start := time.Now()
		result := alg(largeText, testPattern)
		duration := time.Since(start)
		fmt.Printf("  %s: found at positions %v (%v)\n", name, result, duration)
	}
	
	fmt.Println("\n=== Wildcard Pattern Matching ===")
	
	// Implement wildcard pattern matching with * and ?
	wildcardTests := []struct {
		text    string
		pattern string
	}{
		{"hello", "h*o"},
		{"hello", "h?llo"},
		{"abc", "a*c"},
		{"abcdef", "*def"},
		{"test", "t??t"},
		{"programming", "prog*ing"},
		{"golang", "go*"},
	}
	
	fmt.Println("Wildcard matching (* = any sequence, ? = any single char):")
	for _, test := range wildcardTests {
		matches := wildcardMatch(test.text, test.pattern)
		fmt.Printf("  '%s' matches '%s': %t\n", test.text, test.pattern, matches)
	}
	
	fmt.Println("\n=== Longest Common Subsequence ===")
	
	// Find LCS between two strings
	str1 := "ABCDGH"
	str2 := "AEDFHR"
	
	lcs, lcsLength := longestCommonSubsequence(str1, str2)
	
	fmt.Printf("String 1: %s\n", str1)
	fmt.Printf("String 2: %s\n", str2)
	fmt.Printf("LCS: %s (length: %d)\n", lcs, lcsLength)
	
	fmt.Println("\n=== Edit Distance (Levenshtein) ===")
	
	// Calculate edit distance between strings
	editTests := [][]string{
		{"kitten", "sitting"},
		{"saturday", "sunday"},
		{"hello", "hallo"},
		{"algorithm", "altruistic"},
	}
	
	fmt.Println("Edit distance (minimum operations to transform):")
	for _, test := range editTests {
		distance := editDistance(test[0], test[1])
		fmt.Printf("  '%s' -> '%s': %d operations\n", test[0], test[1], distance)
	}
	
	fmt.Println("\n=== String Similarity Metrics ===")
	
	// Implement different similarity measures
	word1, word2 := "programming", "programs"
	
	jaccardSim := jaccardSimilarity(word1, word2)
	cosineSim := cosineSimilarity(word1, word2)
	
	fmt.Printf("Comparing '%s' and '%s':\n", word1, word2)
	fmt.Printf("  Jaccard Similarity: %.3f\n", jaccardSim)
	fmt.Printf("  Cosine Similarity: %.3f\n", cosineSim)
	if len(word1) == len(word2) {
		hammingDist := hammingDistance(word1, word2)
		fmt.Printf("  Hamming Distance: %d\n", hammingDist)
	}
	
	fmt.Println("\n=== Algorithm Complexity Summary ===")
	fmt.Println("Time complexities (n = text length, m = pattern length):")
	fmt.Println("  Naive: O(n*m)")
	fmt.Println("  KMP: O(n+m)")
	fmt.Println("  Boyer-Moore: O(n*m) worst, O(n/m) best")
	fmt.Println("  Rabin-Karp: O(n+m) average, O(n*m) worst")
	fmt.Println("  LCS: O(n*m)")
	fmt.Println("  Edit Distance: O(n*m)")
}

// Implement Naive Pattern Matching
func naivePatternMatch(text, pattern string) []int {
	var positions []int
	n, m := len(text), len(pattern)
	
	for i := 0; i <= n-m; i++ {
		j := 0
		for j < m && text[i+j] == pattern[j] {
			j++
		}
		if j == m {
			positions = append(positions, i)
		}
	}
	
	return positions
}

// Implement KMP Pattern Matching
func kmpPatternMatch(text, pattern string) []int {
	var positions []int
	n, m := len(text), len(pattern)
	
	if m == 0 {
		return positions
	}
	
	// Build failure function
	failure := buildKMPTable(pattern)
	
	i, j := 0, 0
	for i < n {
		if pattern[j] == text[i] {
			i++
			j++
		}
		
		if j == m {
			positions = append(positions, i-j)
			j = failure[j-1]
		} else if i < n && pattern[j] != text[i] {
			if j != 0 {
				j = failure[j-1]
			} else {
				i++
			}
		}
	}
	
	return positions
}

// Build KMP failure function
func buildKMPTable(pattern string) []int {
	m := len(pattern)
	failure := make([]int, m)
	
	length := 0
	i := 1
	
	for i < m {
		if pattern[i] == pattern[length] {
			length++
			failure[i] = length
			i++
		} else {
			if length != 0 {
				length = failure[length-1]
			} else {
				failure[i] = 0
				i++
			}
		}
	}
	
	return failure
}

// Implement Boyer-Moore Pattern Matching (simplified)
func boyerMoorePatternMatch(text, pattern string) []int {
	var positions []int
	n, m := len(text), len(pattern)
	
	if m == 0 {
		return positions
	}
	
	// Build bad character table
	badChar := buildBadCharTable(pattern)
	
	i := 0
	for i <= n-m {
		j := m - 1
		
		// Check pattern from right to left
		for j >= 0 && pattern[j] == text[i+j] {
			j--
		}
		
		if j < 0 {
			// Pattern found
			positions = append(positions, i)
			// Skip based on bad character rule
			if i+m < n {
				if shift, exists := badChar[rune(text[i+m])]; exists {
					i += max(1, m-shift)
				} else {
					i += m
				}
			} else {
				i++
			}
		} else {
			// Skip based on bad character rule
			if shift, exists := badChar[rune(text[i+j])]; exists {
				i += max(1, j-shift)
			} else {
				i += j + 1
			}
		}
	}
	
	return positions
}

// Build bad character table for Boyer-Moore
func buildBadCharTable(pattern string) map[rune]int {
	table := make(map[rune]int)
	
	for i, char := range pattern {
		table[char] = i
	}
	
	return table
}

// Implement Rabin-Karp Pattern Matching
func rabinKarpPatternMatch(text, pattern string) []int {
	var positions []int
	n, m := len(text), len(pattern)
	
	if m == 0 || m > n {
		return positions
	}
	
	const prime = 101
	patternHash := polynomialHash(pattern)
	textHash := polynomialHash(text[:m])
	
	// Check first window
	if textHash == patternHash && text[:m] == pattern {
		positions = append(positions, 0)
	}
	
	// Rolling hash for remaining windows
	for i := 1; i <= n-m; i++ {
		// Remove old character and add new character
		textHash = updateRollingHash(textHash, rune(text[i-1]), rune(text[i+m-1]), m)
		
		if textHash == patternHash && text[i:i+m] == pattern {
			positions = append(positions, i)
		}
	}
	
	return positions
}

// Polynomial rolling hash function
func polynomialHash(s string) int {
	hash := 0
	prime := 31
	
	for _, char := range s {
		hash = hash*prime + int(char)
	}
	
	return hash
}

// Update rolling hash
func updateRollingHash(oldHash int, oldChar, newChar rune, patternLen int) int {
	prime := 31
	// Remove old character
	oldHash -= int(oldChar) * int(math.Pow(float64(prime), float64(patternLen-1)))
	// Shift and add new character
	oldHash = oldHash*prime + int(newChar)
	
	return oldHash
}

// Implement wildcard pattern matching
func wildcardMatch(text, pattern string) bool {
	n, m := len(text), len(pattern)
	
	// DP table
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	
	// Empty pattern matches empty text
	dp[0][0] = true
	
	// Handle patterns with '*' at the beginning
	for j := 1; j <= m; j++ {
		if pattern[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		}
	}
	
	// Fill DP table
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if pattern[j-1] == '*' {
				// '*' can match empty string or any character
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if pattern[j-1] == '?' || text[i-1] == pattern[j-1] {
				// '?' matches any single character, or exact match
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	
	return dp[n][m]
}

// Find Longest Common Subsequence
func longestCommonSubsequence(str1, str2 string) (string, int) {
	n, m := len(str1), len(str2)
	
	// DP table
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	
	// Fill DP table
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	
	// Reconstruct LCS
	lcsLength := dp[n][m]
	lcs := make([]byte, lcsLength)
	i, j := n, m
	index := lcsLength - 1
	
	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			lcs[index] = str1[i-1]
			i--
			j--
			index--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}
	
	return string(lcs), lcsLength
}

// Calculate Edit Distance (Levenshtein)
func editDistance(str1, str2 string) int {
	n, m := len(str1), len(str2)
	
	// DP table
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	
	// Initialize base cases
	for i := 0; i <= n; i++ {
		dp[i][0] = i // Delete all characters from str1
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j // Insert all characters from str2
	}
	
	// Fill DP table
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] // No operation needed
			} else {
				dp[i][j] = 1 + min(min(
					dp[i-1][j],   // Delete
					dp[i][j-1]),  // Insert
					dp[i-1][j-1]) // Substitute
			}
		}
	}
	
	return dp[n][m]
}

// Calculate Jaccard Similarity
func jaccardSimilarity(str1, str2 string) float64 {
	set1 := make(map[rune]bool)
	set2 := make(map[rune]bool)
	
	for _, char := range str1 {
		set1[char] = true
	}
	for _, char := range str2 {
		set2[char] = true
	}
	
	intersection := 0
	union := make(map[rune]bool)
	
	for char := range set1 {
		union[char] = true
		if set2[char] {
			intersection++
		}
	}
	
	for char := range set2 {
		union[char] = true
	}
	
	if len(union) == 0 {
		return 0
	}
	
	return float64(intersection) / float64(len(union))
}

// Calculate Cosine Similarity
func cosineSimilarity(str1, str2 string) float64 {
	// Create character frequency vectors
	freq1 := make(map[rune]int)
	freq2 := make(map[rune]int)
	
	for _, char := range str1 {
		freq1[char]++
	}
	for _, char := range str2 {
		freq2[char]++
	}
	
	// Calculate dot product and magnitudes
	dotProduct := 0
	magnitude1 := 0
	magnitude2 := 0
	
	allChars := make(map[rune]bool)
	for char := range freq1 {
		allChars[char] = true
	}
	for char := range freq2 {
		allChars[char] = true
	}
	
	for char := range allChars {
		f1 := freq1[char]
		f2 := freq2[char]
		
		dotProduct += f1 * f2
		magnitude1 += f1 * f1
		magnitude2 += f2 * f2
	}
	
	if magnitude1 == 0 || magnitude2 == 0 {
		return 0
	}
	
	return float64(dotProduct) / (math.Sqrt(float64(magnitude1)) * math.Sqrt(float64(magnitude2)))
}

// Calculate Hamming Distance
func hammingDistance(str1, str2 string) int {
	if len(str1) != len(str2) {
		return -1 // Invalid for different lengths
	}
	
	distance := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			distance++
		}
	}
	
	return distance
}

// Utility functions

func findAllOccurrences(text, pattern string) []int {
	var positions []int
	start := 0
	for {
		pos := strings.Index(text[start:], pattern)
		if pos == -1 {
			break
		}
		positions = append(positions, start+pos)
		start += pos + 1
	}
	return positions
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}