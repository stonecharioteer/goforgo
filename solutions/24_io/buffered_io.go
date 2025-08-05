// buffered_io.go - SOLUTION
// Learn buffered I/O operations for performance

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== Buffered I/O Operations ===")
	
	// Create test data
	testData := []string{
		"Line 1: Hello, buffered I/O!",
		"Line 2: This is a test.",
		"Line 3: Buffering improves performance.",
		"Line 4: Especially for many small operations.",
		"Line 5: End of test data.",
	}
	
	filename := "buffered_test.txt"
	
	fmt.Println("\n=== Buffered Writing ===")
	
	// Create file for writing
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	
	// Create buffered writer
	writer := bufio.NewWriter(file)
	
	fmt.Println("Writing data with buffered writer...")
	for i, line := range testData {
		// Write line to buffered writer
		writer.WriteString(line + "\n")
		fmt.Printf("Wrote line %d (buffered)\n", i+1)
	}
	
	// Flush buffer to ensure data is written
	writer.Flush()
	file.Close()
	
	fmt.Println("Data written and flushed to file")
	
	fmt.Println("\n=== Buffered Reading ===")
	
	// Open file for reading
	file, err = os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()
	
	// Create buffered reader
	reader := bufio.NewReader(file)
	
	fmt.Println("Reading data with buffered reader:")
	lineNum := 1
	
	for {
		// Read line from buffered reader
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading line: %v\n", err)
			break
		}
		
		fmt.Printf("Line %d: %s", lineNum, line)
		lineNum++
	}
	
	fmt.Println("\n=== Scanner Usage ===")
	
	// Reset file position
	file.Seek(0, 0)
	
	// Create scanner
	scanner := bufio.NewScanner(file)
	
	fmt.Println("Reading with scanner:")
	lineNum = 1
	
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Scanned line %d: %s\n", lineNum, line)
		lineNum++
	}
	
	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v\n", err)
	}
	
	fmt.Println("\n=== String I/O ===")
	
	// Create string reader
	stringData := "Hello\nWorld\nFrom\nString\nReader"
	stringReader := strings.NewReader(stringData)
	
	// Create buffered reader from string
	bufferedStringReader := bufio.NewReader(stringReader)
	
	fmt.Println("Reading from string with buffered reader:")
	for {
		line, err := bufferedStringReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Printf("String line: %s", line)
	}
	
	fmt.Println("\n=== Custom Buffer Sizes ===")
	
	// Test different buffer sizes
	bufferSizes := []int{16, 64, 256, 1024}
	
	for _, size := range bufferSizes {
		fmt.Printf("\nTesting buffer size: %d bytes\n", size)
		
		// Create writer with custom buffer size
		file, _ = os.Create("buffer_test.txt")
		customWriter := bufio.NewWriterSize(file, size)
		
		// Write some data
		testString := strings.Repeat("A", 100) // 100 'A' characters
		customWriter.WriteString(testString)
		
		// Check available buffer space
		available := customWriter.Available()
		fmt.Printf("Available buffer space: %d bytes\n", available)
		
		customWriter.Flush()
		file.Close()
	}
	
	fmt.Println("\n=== Peek and Discard Operations ===")
	
	// Create reader for peeking
	peekData := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	peekReader := bufio.NewReader(strings.NewReader(peekData))
	
	// Peek at data without consuming it
	peeked, err := peekReader.Peek(5)
	if err != nil {
		fmt.Printf("Peek error: %v\n", err)
	} else {
		fmt.Printf("Peeked data: %s\n", string(peeked))
	}
	
	// Read data normally
	normalRead := make([]byte, 3)
	n, err := peekReader.Read(normalRead)
	if err != nil {
		fmt.Printf("Read error: %v\n", err)
	} else {
		fmt.Printf("Normal read (%d bytes): %s\n", n, string(normalRead))
	}
	
	// Discard some data
	discarded, err := peekReader.Discard(2)
	if err != nil {
		fmt.Printf("Discard error: %v\n", err)
	} else {
		fmt.Printf("Discarded %d bytes\n", discarded)
	}
	
	// Read remaining data
	remaining, err := io.ReadAll(peekReader)
	if err != nil && err != io.EOF {
		fmt.Printf("Final read error: %v\n", err)
	} else {
		fmt.Printf("Remaining data: %s\n", string(remaining))
	}
	
	fmt.Println("\n=== Cleanup ===")
	
	// Remove test files
	os.Remove("buffered_test.txt")
	os.Remove("buffer_test.txt")
	
	fmt.Println("Test files cleaned up")
}