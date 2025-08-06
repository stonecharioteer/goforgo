// io_interfaces.go
// Learn Go's io interfaces: Reader, Writer, Closer, Seeker

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// TODO: Custom Reader implementation
type CustomReader struct {
	data []byte
	pos  int
}

// TODO: Implement io.Reader interface
func (cr *CustomReader) Read(p []byte) (n int, err error) {
	// TODO: Implement Read method
	// Read data into p, update position, return bytes read and error
}

// TODO: Custom Writer implementation  
type CustomWriter struct {
	data []byte
}

// TODO: Implement io.Writer interface
func (cw *CustomWriter) Write(p []byte) (n int, err error) {
	// TODO: Implement Write method
	// Append p to internal data, return bytes written and error
}

func (cw *CustomWriter) String() string {
	return string(cw.data)
}

// TODO: Custom ReadWriter implementation
type CustomReadWriter struct {
	buffer []byte
	readPos int
}

// TODO: Implement io.Reader
func (crw *CustomReadWriter) Read(p []byte) (n int, err error) {
	// TODO: Read from buffer starting at readPos
}

// TODO: Implement io.Writer  
func (crw *CustomReadWriter) Write(p []byte) (n int, err error) {
	// TODO: Append to buffer
}

// TODO: Custom Seeker implementation
type SeekableBuffer struct {
	data []byte
	pos  int64
}

// TODO: Implement io.Reader
func (sb *SeekableBuffer) Read(p []byte) (n int, err error) {
	// TODO: Read from current position
}

// TODO: Implement io.Writer
func (sb *SeekableBuffer) Write(p []byte) (n int, err error) {
	// TODO: Write at current position, expand buffer if needed
}

// TODO: Implement io.Seeker
func (sb *SeekableBuffer) Seek(offset int64, whence int) (int64, error) {
	// TODO: Implement seeking with io.SeekStart, io.SeekCurrent, io.SeekEnd
}

func main() {
	fmt.Println("=== Go I/O Interfaces ===")
	
	fmt.Println("\n=== Custom Reader ===")
	
	// TODO: Test custom reader
	testData := []byte("Hello, World! This is test data for our custom reader.")
	reader := /* create custom reader with testData */
	
	fmt.Printf("Original data: %s\n", testData)
	fmt.Println("Reading in chunks:")
	
	// TODO: Read data in chunks
	buffer := make([]byte, 10)
	chunkNum := 1
	
	for {
		// TODO: Read next chunk
		n, err := /* read into buffer */
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
	
	// TODO: Test custom writer
	writer := /* create custom writer */
	
	testStrings := []string{
		"First line\n",
		"Second line\n", 
		"Third line\n",
	}
	
	fmt.Println("Writing strings:")
	for i, s := range testStrings {
		// TODO: Write string to custom writer
		n, err := /* write string to writer */
		if err != nil {
			fmt.Printf("Write error: %v\n", err)
		} else {
			fmt.Printf("Wrote %d bytes: %q\n", n, s)
		}
	}
	
	fmt.Printf("Writer contents: %s", writer.String())
	
	fmt.Println("\n=== io.Copy ===")
	
	// TODO: Use io.Copy with custom types
	sourceReader := /* create reader with source data */
	destWriter := /* create writer for destination */
	
	sourceData := "Data to copy from reader to writer using io.Copy"
	sourceReader = &CustomReader{data: []byte(sourceData), pos: 0}
	
	fmt.Printf("Source data: %s\n", sourceData)
	
	// TODO: Copy data using io.Copy
	bytesCopied, err := /* copy from reader to writer */
	if err != nil {
		fmt.Printf("Copy error: %v\n", err)
	} else {
		fmt.Printf("Copied %d bytes\n", bytesCopied)
		fmt.Printf("Destination: %s\n", destWriter.String())
	}
	
	fmt.Println("\n=== ReadWriter Interface ===")
	
	// TODO: Test ReadWriter
	readWriter := /* create custom ReadWriter */
	
	// TODO: Write some data
	writeData := "Hello from ReadWriter!"
	n, err := /* write data to readWriter */
	if err != nil {
		fmt.Printf("Write error: %v\n", err)
	} else {
		fmt.Printf("Wrote %d bytes: %s\n", n, writeData)
	}
	
	// TODO: Read the data back
	readBuffer := make([]byte, len(writeData))
	n, err = /* read from readWriter */
	if err != nil {
		fmt.Printf("Read error: %v\n", err)
	} else {
		fmt.Printf("Read %d bytes: %s\n", n, string(readBuffer[:n]))
	}
	
	fmt.Println("\n=== Seeker Interface ===")
	
	// TODO: Test seeker
	seekBuffer := /* create seekable buffer */
	testContent := "0123456789ABCDEFGHIJ"
	
	// TODO: Write initial content
	/* write testContent to seeker */
	
	fmt.Printf("Buffer content: %s\n", testContent)
	fmt.Println("Testing seek operations:")
	
	// TODO: Test different seek operations
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
		// TODO: Perform seek operation
		newPos, err := /* seek with test.offset and test.whence */
		if err != nil {
			fmt.Printf("%s: Error - %v\n", test.desc, err)
		} else {
			fmt.Printf("%s: New position %d\n", test.desc, newPos)
			
			// TODO: Read a few bytes from current position
			readBuf := make([]byte, 3)
			n, err := /* read into readBuf */
			if err != nil && err != io.EOF {
				fmt.Printf("  Read error: %v\n", err)
			} else if n > 0 {
				fmt.Printf("  Read: %q\n", readBuf[:n])
			}
		}
	}
	
	fmt.Println("\n=== Working with Standard Library ===")
	
	// TODO: Use custom types with standard library functions
	
	// Test with strings.Reader
	strReader := /* create strings.Reader with test data */
	testText := "Testing with strings.Reader"
	strReader = strings.NewReader(testText)
	
	fmt.Printf("strings.Reader content: %s\n", testText)
	
	// TODO: Read all content using io.ReadAll
	allContent, err := /* read all from strReader */
	if err != nil {
		fmt.Printf("ReadAll error: %v\n", err)
	} else {
		fmt.Printf("Read all: %s\n", allContent)
	}
	
	// Test with bytes.Buffer
	fmt.Println("\nTesting with bytes.Buffer:")
	buffer := /* create new bytes.Buffer */
	
	// TODO: Write to buffer
	/* write "Buffer test data" to buffer */
	fmt.Printf("Buffer after write: %s\n", buffer.String())
	
	// TODO: Read from buffer
	readData := make([]byte, 6)
	n, err = /* read from buffer */
	if err != nil {
		fmt.Printf("Buffer read error: %v\n", err)
	} else {
		fmt.Printf("Read from buffer: %q (%d bytes)\n", readData[:n], n)
		fmt.Printf("Buffer remaining: %s\n", buffer.String())
	}
	
	fmt.Println("\n=== Chaining I/O Operations ===")
	
	// TODO: Chain multiple I/O operations
	originalData := "Original data for chaining operations"
	
	fmt.Printf("Original: %s\n", originalData)
	
	// TODO: Create chain: Reader -> Writer1 -> Reader -> Writer2
	reader1 := /* create reader with originalData */
	writer1 := /* create writer */
	
	// TODO: Copy from reader1 to writer1
	/* copy from reader1 to writer1 */
	
	fmt.Printf("After first copy: %s\n", writer1.String())
	
	// TODO: Create reader from writer1 content
	reader2 := /* create reader from writer1 content */
	writer2 := /* create second writer */
	
	// TODO: Copy from reader2 to writer2
	/* copy from reader2 to writer2 */
	
	fmt.Printf("After second copy: %s\n", writer2.String())
	
	// TODO: Verify data integrity
	finalData := writer2.String()
	dataMatches := /* compare original and final data */
	fmt.Printf("Data integrity check: %t\n", dataMatches)
}