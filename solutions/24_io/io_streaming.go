// io_streaming.go
// Learn streaming I/O and data processing in Go

package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== Streaming I/O Operations ===")
	
	// Create sample data for streaming
	sampleData := "Line 1: Hello, World!\nLine 2: Go streaming I/O\nLine 3: Processing data\nLine 4: Final line"
	
	// Create string reader
	reader := strings.NewReader(sampleData)
	
	fmt.Println("=== Line-by-Line Processing ===")
	
	// Process lines using scanner
	scanner := bufio.NewScanner(reader)
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line %d: %s\n", lineNum, line)
		lineNum++
	}
	
	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %v\n", err)
	}
	
	fmt.Println("\n=== Streaming Copy Operations ===")
	
	// Create source and destination
	source := strings.NewReader(sampleData)
	var dest strings.Builder
	
	// Copy data from source to destination
	bytescopied, err := io.Copy(&dest, source)
	if err != nil {
		fmt.Printf("Copy error: %v\n", err)
		return
	}
	
	fmt.Printf("Copied %d bytes\n", bytescopied)
	fmt.Printf("Destination content: %s\n", dest.String())
	
	fmt.Println("\n=== Compressed Streaming ===")
	
	// Create file for compression test
	filename := "stream_test.txt.gz"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Create error: %v\n", err)
		return
	}
	defer file.Close()
	defer os.Remove(filename)
	
	// Create gzip writer
	gzipWriter := gzip.NewWriter(file)
	
	// Write compressed data
	data := "This is compressed streaming data!\nMultiple lines of text.\nCompressed efficiently."
	_, err = gzipWriter.Write([]byte(data))
	if err != nil {
		fmt.Printf("Gzip write error: %v\n", err)
		return
	}
	
	// Close gzip writer
	gzipWriter.Close()
	
	fmt.Println("Data written to compressed file")
	
	// Read compressed data back
	file, err = os.Open(filename)
	if err != nil {
		fmt.Printf("Open error: %v\n", err)
		return
	}
	defer file.Close()
	
	// Create gzip reader
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		fmt.Printf("Gzip reader error: %v\n", err)
		return
	}
	defer gzipReader.Close()
	
	// Read decompressed data
	decompressed, err := io.ReadAll(gzipReader)
	if err != nil {
		fmt.Printf("Read error: %v\n", err)
		return
	}
	
	fmt.Printf("Decompressed data: %s\n", string(decompressed))
	
	fmt.Println("\n=== Streaming Transformer ===")
	
	// Test the streaming transformer
	input := strings.NewReader("hello world\ngo programming\nstreaming io")
	var output strings.Builder
	
	err = streamTransformer(input, &output, strings.ToUpper)
	if err != nil {
		fmt.Printf("Transform error: %v\n", err)
	} else {
		fmt.Printf("Transformed output:\n%s", output.String())
	}
	
	fmt.Println("\n=== Chunked Reading ===")
	
	// Test chunked reader
	input = strings.NewReader("This is a test of chunked reading functionality in Go")
	chunkedReader := NewChunkedReader(input, 10)
	
	chunk := make([]byte, 20) // Try to read more than chunk size
	for {
		n, err := chunkedReader.Read(chunk)
		if n > 0 {
			fmt.Printf("Read chunk (%d bytes): %q\n", n, string(chunk[:n]))
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Chunked read error: %v\n", err)
			break
		}
	}
	
	fmt.Println("\n=== Progress Reading ===")
	
	// Test progress reader
	testData := strings.Repeat("Hello, World! ", 100)
	input = strings.NewReader(testData)
	progressReader := NewProgressReader(input, int64(len(testData)))
	
	buffer := make([]byte, 50)
	for {
		n, err := progressReader.Read(buffer)
		if err == io.EOF {
			fmt.Println("\nReading complete!")
			break
		}
		if err != nil {
			fmt.Printf("\nProgress read error: %v\n", err)
			break
		}
		// Simulate processing time
		// time.Sleep(100 * time.Millisecond)
	}
}

// Implement streaming transformer
func streamTransformer(input io.Reader, output io.Writer, transform func(string) string) error {
	// Create scanner for input
	scanner := bufio.NewScanner(input)
	
	// Create buffered writer for output
	writer := bufio.NewWriter(output)
	defer writer.Flush()
	
	// Process each line
	for scanner.Scan() {
		line := scanner.Text()
		transformed := transform(line)
		
		// Write transformed line
		_, err := writer.WriteString(transformed + "\n")
		if err != nil {
			return err
		}
	}
	
	return scanner.Err()
}

// Implement chunked reader
type ChunkedReader struct {
	reader io.Reader
	chunkSize int
}

// Implement NewChunkedReader
func NewChunkedReader(reader io.Reader, chunkSize int) *ChunkedReader {
	return &ChunkedReader{
		reader:    reader,
		chunkSize: chunkSize,
	}
}

// Implement Read method for ChunkedReader
func (cr *ChunkedReader) Read(p []byte) (n int, err error) {
	// Limit read size to chunk size
	if len(p) > cr.chunkSize {
		p = p[:cr.chunkSize]
	}
	
	return cr.reader.Read(p)
}

// Implement progress reader
type ProgressReader struct {
	reader io.Reader
	total  int64
	read   int64
}

// Implement NewProgressReader
func NewProgressReader(reader io.Reader, total int64) *ProgressReader {
	return &ProgressReader{
		reader: reader,
		total:  total,
	}
}

// Implement Read method for ProgressReader
func (pr *ProgressReader) Read(p []byte) (n int, err error) {
	n, err = pr.reader.Read(p)
	
	pr.read += int64(n)
	
	// Show progress
	if pr.total > 0 {
		progress := float64(pr.read) / float64(pr.total) * 100
		fmt.Printf("\rProgress: %.1f%% (%d/%d bytes)", progress, pr.read, pr.total)
	}
	
	return n, err
}