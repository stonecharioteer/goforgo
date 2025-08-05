// io_interfaces.go - SOLUTION
// Learn Go's io interfaces: Reader, Writer, Closer, Seeker

package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Custom Reader implementation
type CustomReader struct {
	data []byte
	pos  int
}

// Implement io.Reader interface
func (cr *CustomReader) Read(p []byte) (n int, err error) {
	if cr.pos >= len(cr.data) {
		return 0, io.EOF
	}
	
	n = copy(p, cr.data[cr.pos:])
	cr.pos += n
	
	if cr.pos >= len(cr.data) {
		err = io.EOF
	}
	
	return n, err
}

// Custom Writer implementation  
type CustomWriter struct {
	data []byte
}

// Implement io.Writer interface
func (cw *CustomWriter) Write(p []byte) (n int, err error) {
	cw.data = append(cw.data, p...)
	return len(p), nil
}

func (cw *CustomWriter) String() string {
	return string(cw.data)
}

// Custom ReadWriter implementation
type CustomReadWriter struct {
	buffer  []byte
	readPos int
}

// Implement io.Reader
func (crw *CustomReadWriter) Read(p []byte) (n int, err error) {
	if crw.readPos >= len(crw.buffer) {
		return 0, io.EOF
	}
	
	n = copy(p, crw.buffer[crw.readPos:])
	crw.readPos += n
	
	if crw.readPos >= len(crw.buffer) {
		err = io.EOF
	}
	
	return n, err
}

// Implement io.Writer  
func (crw *CustomReadWriter) Write(p []byte) (n int, err error) {
	crw.buffer = append(crw.buffer, p...)
	return len(p), nil
}

// Custom Seeker implementation
type SeekableBuffer struct {
	data []byte
	pos  int64
}

// Implement io.Reader
func (sb *SeekableBuffer) Read(p []byte) (n int, err error) {
	if sb.pos >= int64(len(sb.data)) {
		return 0, io.EOF
	}
	
	n = copy(p, sb.data[sb.pos:])
	sb.pos += int64(n)
	
	if sb.pos >= int64(len(sb.data)) {
		err = io.EOF
	}
	
	return n, err
}

// Implement io.Writer
func (sb *SeekableBuffer) Write(p []byte) (n int, err error) {
	// Expand buffer if necessary
	needed := int(sb.pos) + len(p)
	if needed > len(sb.data) {
		newData := make([]byte, needed)
		copy(newData, sb.data)
		sb.data = newData
	}
	
	n = copy(sb.data[sb.pos:], p)
	sb.pos += int64(n)
	
	return n, nil
}

// Implement io.Seeker
func (sb *SeekableBuffer) Seek(offset int64, whence int) (int64, error) {
	var newPos int64
	
	switch whence {
	case io.SeekStart:
		newPos = offset
	case io.SeekCurrent:
		newPos = sb.pos + offset
	case io.SeekEnd:
		newPos = int64(len(sb.data)) + offset
	default:
		return 0, fmt.Errorf("invalid whence value")
	}
	
	if newPos < 0 {
		return 0, fmt.Errorf("negative position")
	}
	
	sb.pos = newPos
	return newPos, nil
}

func main() {
	fmt.Println("=== Go I/O Interfaces ===")
	
	fmt.Println("\n=== Custom Reader ===")
	
	// Test custom reader
	testData := []byte("Hello, World! This is test data for our custom reader.")
	reader := &CustomReader{data: testData, pos: 0}
	
	fmt.Printf("Original data: %s\n", testData)
	fmt.Println("Reading in chunks:")
	
	// Read data in chunks
	buffer := make([]byte, 10)
	chunkNum := 1
	
	for {
		// Read next chunk
		n, err := reader.Read(buffer)
		if n > 0 {
			fmt.Printf("Chunk %d: %q (%d bytes)\n", chunkNum, buffer[:n], n)
			chunkNum++
		}
		
		if err == io.EOF {
			fmt.Println("Reached end of data")
			break
		}
		if err != nil {
			fmt.Printf("Read error: %v\n", err)
			break
		}
	}
	
	fmt.Println("\n=== Custom Writer ===")
	
	// Test custom writer
	writer := &CustomWriter{}
	
	testStrings := []string{
		"First line\n",
		"Second line\n", 
		"Third line\n",
	}
	
	fmt.Println("Writing strings:")
	for _, s := range testStrings {
		// Write string to custom writer
		n, err := writer.Write([]byte(s))
		if err != nil {
			fmt.Printf("Write error: %v\n", err)
		} else {
			fmt.Printf("Wrote %d bytes: %q\n", n, s)
		}
	}
	
	fmt.Printf("Writer contents: %s", writer.String())
	
	fmt.Println("\n=== io.Copy ===")
	
	// Use io.Copy with custom types
	sourceData := "Data to copy from reader to writer using io.Copy"
	sourceReader := &CustomReader{data: []byte(sourceData), pos: 0}
	destWriter := &CustomWriter{}
	
	fmt.Printf("Source data: %s\n", sourceData)
	
	// Copy data using io.Copy
	bytesCopied, err := io.Copy(destWriter, sourceReader)
	if err != nil {
		fmt.Printf("Copy error: %v\n", err)
	} else {
		fmt.Printf("Copied %d bytes\n", bytesCopied)
		fmt.Printf("Destination: %s\n", destWriter.String())
	}
	
	fmt.Println("\n=== ReadWriter Interface ===")
	
	// Test ReadWriter
	readWriter := &CustomReadWriter{}
	
	// Write some data
	writeData := "Hello from ReadWriter!"
	n, err := readWriter.Write([]byte(writeData))
	if err != nil {
		fmt.Printf("Write error: %v\n", err)
	} else {
		fmt.Printf("Wrote %d bytes: %s\n", n, writeData)
	}
	
	// Read the data back
	readBuffer := make([]byte, len(writeData))
	n, err = readWriter.Read(readBuffer)
	if err != nil && err != io.EOF {
		fmt.Printf("Read error: %v\n", err)
	} else {
		fmt.Printf("Read %d bytes: %s\n", n, string(readBuffer[:n]))
	}
	
	fmt.Println("\n=== Seeker Interface ===")
	
	// Test seeker
	testContent := "0123456789ABCDEFGHIJ"
	seekBuffer := &SeekableBuffer{
		data: []byte(testContent),
		pos:  0,
	}
	
	fmt.Printf("Buffer content: %s\n", testContent)
	fmt.Println("Testing seek operations:")
	
	// Test different seek operations
	seekTests := []struct {
		offset int64
		whence int
		desc   string
	}{
		{5, io.SeekStart, "Seek to position 5"},
		{3, io.SeekCurrent, "Seek 3 bytes forward from current"},
		{-2, io.SeekEnd, "Seek 2 bytes back from end"},
		{0, io.SeekStart, "Seek to beginning"},
	}
	
	for _, test := range seekTests {
		// Perform seek operation
		newPos, err := seekBuffer.Seek(test.offset, test.whence)
		if err != nil {
			fmt.Printf("%s: Error - %v\n", test.desc, err)
		} else {
			fmt.Printf("%s: New position %d\n", test.desc, newPos)
			
			// Read a few bytes from current position
			readBuf := make([]byte, 3)
			n, err := seekBuffer.Read(readBuf)
			if err != nil && err != io.EOF {
				fmt.Printf("  Read error: %v\n", err)
			} else if n > 0 {
				fmt.Printf("  Read: %q\n", readBuf[:n])
			}
		}
	}
	
	fmt.Println("\n=== Working with Standard Library ===")
	
	// Use custom types with standard library functions
	
	// Test with strings.Reader
	testText := "Testing with strings.Reader"
	strReader := strings.NewReader(testText)
	
	fmt.Printf("strings.Reader content: %s\n", testText)
	
	// Read all content using io.ReadAll
	allContent, err := io.ReadAll(strReader)
	if err != nil {
		fmt.Printf("ReadAll error: %v\n", err)
	} else {
		fmt.Printf("Read all: %s\n", allContent)
	}
	
	// Test with bytes.Buffer
	fmt.Println("\nTesting with bytes.Buffer:")
	buffer := &bytes.Buffer{}
	
	// Write to buffer
	buffer.WriteString("Buffer test data")
	fmt.Printf("Buffer after write: %s\n", buffer.String())
	
	// Read from buffer
	readData := make([]byte, 6)
	n, err = buffer.Read(readData)
	if err != nil {
		fmt.Printf("Buffer read error: %v\n", err)
	} else {
		fmt.Printf("Read from buffer: %q (%d bytes)\n", readData[:n], n)
		fmt.Printf("Buffer remaining: %s\n", buffer.String())
	}
	
	fmt.Println("\n=== Chaining I/O Operations ===")
	
	// Chain multiple I/O operations
	originalData := "Original data for chaining operations"
	
	fmt.Printf("Original: %s\n", originalData)
	
	// Create chain: Reader -> Writer1 -> Reader -> Writer2
	reader1 := &CustomReader{data: []byte(originalData), pos: 0}
	writer1 := &CustomWriter{}
	
	// Copy from reader1 to writer1
	io.Copy(writer1, reader1)
	
	fmt.Printf("After first copy: %s\n", writer1.String())
	
	// Create reader from writer1 content
	reader2 := &CustomReader{data: []byte(writer1.String()), pos: 0}
	writer2 := &CustomWriter{}
	
	// Copy from reader2 to writer2
	io.Copy(writer2, reader2)
	
	fmt.Printf("After second copy: %s\n", writer2.String())
	
	// Verify data integrity
	finalData := writer2.String()
	dataMatches := originalData == finalData
	fmt.Printf("Data integrity check: %t\n", dataMatches)
}