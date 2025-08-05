// testing_basics_test.go - SOLUTION
// Learn how to write tests in Go

package main

import (
	"fmt"
	"testing"
)

// Basic test function
func TestAdd(t *testing.T) {
	// Test the Add function
	result := Add(2, 3)
	expected := 5
	
	if result != expected {
		// Use t.Errorf to report test failure
		t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
	}
}

// Test with multiple cases
func TestMultiply(t *testing.T) {
	// Test cases
	testCases := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 6},
		{0, 5, 0},
		{-2, 3, -6},
		{4, -2, -8},
	}
	
	for _, tc := range testCases {
		result := Multiply(tc.a, tc.b)
		if result != tc.expected {
			// Report failure with test case details
			t.Errorf("Multiply(%d, %d) = %d; expected %d", tc.a, tc.b, result, tc.expected)
		}
	}
}

// Test boolean function
func TestIsEven(t *testing.T) {
	// Test even numbers
	evenNumbers := []int{2, 4, 6, 8, 0, -2}
	for _, num := range evenNumbers {
		if !IsEven(num) {
			t.Errorf("IsEven(%d) should be true", num)
		}
	}
	
	// Test odd numbers
	oddNumbers := []int{1, 3, 5, 7, -1, -3}
	for _, num := range oddNumbers {
		if IsEven(num) {
			t.Errorf("IsEven(%d) should be false", num)
		}
	}
}

// Test edge cases
func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},   // Edge case
		{1, 1},   // Edge case
		{3, 6},   // 3! = 3*2*1
		{5, 120}, // 5! = 5*4*3*2*1
	}
	
	for _, test := range tests {
		result := Factorial(test.input)
		if result != test.expected {
			t.Errorf("Factorial(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

// Test string function
func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},                   // Empty string
		{"a", "a"},                 // Single character
		{"Go", "oG"},               // Short string
		{"racecar", "racecar"},     // Palindrome
	}
	
	for _, test := range tests {
		result := ReverseString(test.input)
		if result != test.expected {
			t.Errorf("ReverseString(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}

// Test function that returns error
func TestFindMax(t *testing.T) {
	// Test normal case
	numbers := []int{3, 7, 2, 9, 1}
	max, err := FindMax(numbers)
	
	if err != nil {
		t.Errorf("FindMax(%v) returned unexpected error: %v", numbers, err)
	}
	
	expected := 9
	if max != expected {
		t.Errorf("FindMax(%v) = %d; expected %d", numbers, max, expected)
	}
	
	// Test empty slice
	emptySlice := []int{}
	_, err = FindMax(emptySlice)
	
	if err == nil {
		t.Error("FindMax([]) should return an error for empty slice")
	}
	
	// Test single element
	singleElement := []int{42}
	max, err = FindMax(singleElement)
	
	if err != nil {
		t.Errorf("FindMax(%v) returned unexpected error: %v", singleElement, err)
	}
	
	if max != 42 {
		t.Errorf("FindMax(%v) = %d; expected 42", singleElement, max)
	}
}

// Test error cases
func TestDivide(t *testing.T) {
	// Test normal division
	result, err := Divide(10, 2)
	if err != nil {
		t.Errorf("Divide(10, 2) returned unexpected error: %v", err)
	}
	
	expected := 5.0
	if result != expected {
		t.Errorf("Divide(10, 2) = %.2f; expected %.2f", result, expected)
	}
	
	// Test division by zero
	_, err = Divide(5, 0)
	if err == nil {
		t.Error("Divide(5, 0) should return an error")
	}
}

// Benchmark function
func BenchmarkAdd(b *testing.B) {
	// Run the Add function b.N times
	for i := 0; i < b.N; i++ {
		// Call Add(100, 200)
		Add(100, 200)
	}
}

// Benchmark with setup
func BenchmarkFactorial(b *testing.B) {
	// Reset timer to exclude setup time
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		// Call Factorial(10)
		Factorial(10)
	}
}

// Table-driven benchmark
func BenchmarkReverseString(b *testing.B) {
	testStrings := []string{
		"short",
		"medium length string",
		"this is a much longer string that will take more time to reverse",
	}
	
	for _, s := range testStrings {
		// Create sub-benchmark for each string length
		b.Run(fmt.Sprintf("len_%d", len(s)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Call ReverseString(s)
				ReverseString(s)
			}
		})
	}
}