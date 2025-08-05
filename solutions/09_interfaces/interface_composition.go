// interface_composition.go - SOLUTION
// Learn how to compose interfaces and create interface hierarchies

package main

import (
	"fmt"
	"io"
	"strings"
)

// Define basic interfaces
type Reader interface {
	Read([]byte) (int, error)
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

// Compose interfaces to create more complex ones
type ReadWriter interface {
	// Embed Reader and Writer interfaces
	Reader
	Writer
}

type ReadWriteCloser interface {
	// Embed ReadWriter and Closer interfaces
	ReadWriter
	Closer
}

// Alternative composition syntax
type ReadWriteCloser2 interface {
	// Embed all three interfaces directly
	Reader
	Writer
	Closer
}

// Define a File struct that implements ReadWriteCloser
type File struct {
	name     string
	content  []byte
	position int
	closed   bool
}

func NewFile(name string, content string) *File {
	return &File{
		name:     name,
		content:  []byte(content),
		position: 0,
		closed:   false,
	}
}

func (f *File) Read(p []byte) (int, error) {
	if f.closed {
		return 0, fmt.Errorf("file is closed")
	}
	
	// Implement read logic
	// Read from f.content starting at f.position
	// Update f.position and return bytes read
	if f.position >= len(f.content) {
		return 0, io.EOF
	}
	
	n := copy(p, f.content[f.position:])
	f.position += n
	
	if f.position >= len(f.content) {
		return n, io.EOF
	}
	
	return n, nil
}

func (f *File) Write(p []byte) (int, error) {
	if f.closed {
		return 0, fmt.Errorf("file is closed")
	}
	
	// Implement write logic
	// Append p to f.content and return bytes written
	f.content = append(f.content, p...)
	return len(p), nil
}

func (f *File) Close() error {
	if f.closed {
		return fmt.Errorf("file already closed")
	}
	
	// Close the file
	f.closed = true
	fmt.Printf("File %s closed\n", f.name)
	return nil
}

// Define more specific interfaces
type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}

type ReadSeeker interface {
	// Compose Reader and Seeker
	Reader
	Seeker
}

type WriteSeeker interface {
	// Compose Writer and Seeker
	Writer
	Seeker
}

type ReadWriteSeeker interface {
	// Compose Reader, Writer, and Seeker
	Reader
	Writer
	Seeker
}

// Add Seek method to File to implement Seeker
func (f *File) Seek(offset int64, whence int) (int64, error) {
	if f.closed {
		return 0, fmt.Errorf("file is closed")
	}
	
	// Implement seek logic (simple version)
	// whence: 0 = from start, 1 = from current, 2 = from end
	// Update f.position and return new position
	var newPos int64
	
	switch whence {
	case 0: // From start
		newPos = offset
	case 1: // From current
		newPos = int64(f.position) + offset
	case 2: // From end
		newPos = int64(len(f.content)) + offset
	default:
		return 0, fmt.Errorf("invalid whence value: %d", whence)
	}
	
	if newPos < 0 {
		newPos = 0
	}
	if newPos > int64(len(f.content)) {
		newPos = int64(len(f.content))
	}
	
	f.position = int(newPos)
	return newPos, nil
}

// Function that works with any Reader
func readAll(r Reader) ([]byte, error) {
	var result []byte
	buffer := make([]byte, 32) // Small buffer for demo
	
	for {
		n, err := r.Read(buffer)
		if n > 0 {
			result = append(result, buffer[:n]...)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return result, err
		}
	}
	return result, nil
}

// Function that works with any Writer
func writeAll(w Writer, data []byte) error {
	for len(data) > 0 {
		n, err := w.Write(data)
		if err != nil {
			return err
		}
		data = data[n:]
	}
	return nil
}

// Function that uses composed interface
func copyData(src Reader, dst Writer) (int64, error) {
	var total int64
	buffer := make([]byte, 32)
	
	for {
		// Read from source
		n, readErr := src.Read(buffer)
		if n > 0 {
			// Write to destination
			_, writeErr := dst.Write(buffer[:n])
			if writeErr != nil {
				return total, writeErr
			}
			total += int64(n)
		}
		
		if readErr != nil {
			if readErr == io.EOF {
				break
			}
			return total, readErr
		}
	}
	return total, nil
}

// Function that demonstrates interface upgrade/downgrade
func processFile(rwc ReadWriteCloser) {
	fmt.Printf("Processing ReadWriteCloser: %T\n", rwc)
	
	// Use as Writer
	message := "Hello from composed interface!\n"
	writeAll(rwc, []byte(message))
	
	// Check if it also implements Seeker (interface upgrade)
	if seeker, ok := rwc.(Seeker); ok {
		fmt.Println("File also implements Seeker!")
		// Seek to beginning
		seeker.Seek(0, 0)
	}
	
	// Use as Reader
	data, err := readAll(rwc)
	if err != nil {
		fmt.Printf("Read error: %v\n", err)
	} else {
		fmt.Printf("Read data: %s", data)
	}
	
	// Use Closer
	rwc.Close()
}

func main() {
	fmt.Println("=== Interface Composition Demo ===")
	
	// Create a File instance
	file := NewFile("test.txt", "Initial content")
	
	// Use File as different interface types
	fmt.Println("Using as Reader:")
	data, err := readAll(file)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Read: %s\n", data)
	}
	
	// Reset file position for next operations
	file.position = 0
	
	fmt.Println("\nUsing as Writer:")
	err = writeAll(file, []byte(" + additional content"))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	
	fmt.Println("\nUsing as ReadWriteCloser:")
	file.position = 0 // Reset for reading
	processFile(file)
	
	fmt.Println("\n=== Interface Flexibility Demo ===")
	
	// Use standard library types that implement our interfaces
	stringReader := strings.NewReader("Hello from strings.Reader!")
	
	fmt.Println("Reading from strings.Reader:")
	data, err = readAll(stringReader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Read: %s\n", data)
	}
	
	// Demonstrate interface composition with different implementations
	var readers []Reader = []Reader{
		strings.NewReader("Reader 1 content"),
		strings.NewReader("Reader 2 content"),
		NewFile("mem1.txt", "File reader content"),
	}
	
	fmt.Println("\nReading from multiple Reader implementations:")
	for i, reader := range readers {
		data, err := readAll(reader)
		if err != nil {
			fmt.Printf("Reader %d error: %v\n", i+1, err)
		} else {
			fmt.Printf("Reader %d: %s\n", i+1, data)
		}
	}
}