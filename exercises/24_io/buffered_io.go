// buffered_io.go
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
	
	// TODO: Create test data
	testData := []string{
		"Line 1: Hello, buffered I/O!",
		"Line 2: This is a test.",
		"Line 3: Buffering improves performance.",
		"Line 4: Especially for many small operations.",
		"Line 5: End of test data.",
	}
	
	filename := "buffered_test.txt"
	
	fmt.Println("\n=== Buffered Writing ===")
	
	// TODO: Create file for writing
	file, err := /* create file */
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	
	// TODO: Create buffered writer
	writer := /* create buffered writer */
	
	fmt.Println("Writing data with buffered writer...")
	for i, line := range testData {
		// TODO: Write line to buffered writer
		/* write line + newline */
		fmt.Printf("Wrote line %d (buffered)\n", i+1)
	}
	
	// TODO: Flush buffer to ensure data is written
	/* flush the buffer */
	file.Close()
	
	fmt.Println("Data written and flushed to file")
	
	fmt.Println("\n=== Buffered Reading ===")
	
	// TODO: Open file for reading
	file, err = /* open file for reading */
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()
	
	// TODO: Create buffered reader
	reader := /* create buffered reader */
	
	fmt.Println("Reading data with buffered reader:")
	lineNum := 1
	
	for {
		// TODO: Read line from buffered reader
		line, err := /* read line */
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading line: %v\n", err)
			break
		}
		
		fmt.Printf("Line %d: %s\n", lineNum, line)
		lineNum++
	}
	
	fmt.Println("\n=== Scanner Usage ===")
	
	// TODO: Reset file position
	file.Seek(0, 0)
	
	// TODO: Create scanner
	scanner := /* create scanner */
	
	fmt.Println("Reading with scanner:")
	lineNum = 1
	
	for /* scan next line */ {
		line := /* get scanned text */
		fmt.Printf("Scanned line %d: %s\n", lineNum, line)
		lineNum++
	}
	
	// TODO: Check for scanner errors
	if err := /* get scanner error */; err != nil {
		fmt.Printf("Scanner error: %v\n", err)
	}
	
	fmt.Println("\n=== String I/O ===")
	
	// TODO: Create string reader
	stringData := "Hello\nWorld\nFrom\nString\nReader"
	stringReader := /* create strings.Reader */
	
	// TODO: Create buffered reader from string
	bufferedStringReader := /* create buffered reader */
	
	fmt.Println("Reading from string with buffered reader:")
	for {
		line, err := /* read line */
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}
		fmt.Printf("String line: %s\n", line)
	}
	
	fmt.Println("\n=== Custom Buffer Sizes ===")
	
	// TODO: Test different buffer sizes
	bufferSizes := []int{16, 64, 256, 1024}
	
	for _, size := range bufferSizes {
		fmt.Printf("\nTesting buffer size: %d bytes\n", size)
		
		// TODO: Create writer with custom buffer size
		file, _ = os.Create("buffer_test.txt")
		customWriter := /* create buffered writer with custom size */
		
		// Write some data
		testString := strings.Repeat("A", 100) // 100 'A' characters
		/* write test string */
		
		// TODO: Check available buffer space
		available := /* get available buffer space */
		fmt.Printf("Available buffer space: %d bytes\n", available)
		
		customWriter.Flush()
		file.Close()
	}
	
	fmt.Println("\n=== Peek and Discard Operations ===")
	
	// TODO: Create reader for peeking
	peekData := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	peekReader := /* create buffered reader from string */
	
	// TODO: Peek at data without consuming it
	peeked, err := /* peek at first 5 bytes */
	if err != nil {
		fmt.Printf("Peek error: %v\n", err)
	} else {
		fmt.Printf("Peeked data: %s\n", string(peeked))
	}
	
	// TODO: Read data normally
	normalRead := make([]byte, 3)
	n, err := /* read 3 bytes normally */
	if err != nil {
		fmt.Printf("Read error: %v\n", err)
	} else {
		fmt.Printf("Normal read (%d bytes): %s\n", n, string(normalRead))
	}
	
	// TODO: Discard some data
	discarded, err := /* discard 2 bytes */
	if err != nil {
		fmt.Printf("Discard error: %v\n", err)
	} else {
		fmt.Printf("Discarded %d bytes\n", discarded)
	}
	
	// TODO: Read remaining data
	remaining, err := /* read remaining data */
	if err != nil && err != io.EOF {
		fmt.Printf("Final read error: %v\n", err)
	} else {
		fmt.Printf("Remaining data: %s\n", string(remaining))
	}
	
	fmt.Println("\n=== Cleanup ===")
	
	// TODO: Remove test files
	/* remove buffered_test.txt */
	/* remove buffer_test.txt */
	
	fmt.Println("Test files cleaned up")
}