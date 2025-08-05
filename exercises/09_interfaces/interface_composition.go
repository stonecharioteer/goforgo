// interface_composition.go
// Learn how to compose interfaces and create interface hierarchies

package main

import (
	"fmt"
	"io"
	"strings"
)

// TODO: Define basic interfaces
type Reader interface {
	Read([]byte) (int, error)
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

// TODO: Compose interfaces to create more complex ones
type ReadWriter interface {
	// Embed Reader and Writer interfaces
}

type ReadWriteCloser interface {
	// Embed ReadWriter and Closer interfaces
}

// Alternative composition syntax
type ReadWriteCloser2 interface {
	// Embed all three interfaces directly
	Reader
	Writer  
	Closer
}

// TODO: Define a File struct that implements ReadWriteCloser
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
	
	// TODO: Implement read logic
	// Read from f.content starting at f.position
	// Update f.position and return bytes read
	
}

func (f *File) Write(p []byte) (int, error) {
	if f.closed {
		return 0, fmt.Errorf("file is closed")
	}
	
	// TODO: Implement write logic  
	// Append p to f.content and return bytes written
	
}

func (f *File) Close() error {
	if f.closed {
		return fmt.Errorf("file already closed")
	}
	
	// TODO: Close the file
	
	fmt.Printf("File %s closed\n", f.name)
	return nil
}

// TODO: Define more specific interfaces
type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}

type ReadSeeker interface {
	// Compose Reader and Seeker
}

type WriteSeeker interface {
	// Compose Writer and Seeker
}

type ReadWriteSeeker interface {
	// Compose Reader, Writer, and Seeker
}

// TODO: Add Seek method to File to implement Seeker
func (f *File) Seek(offset int64, whence int) (int64, error) {
	if f.closed {
		return 0, fmt.Errorf("file is closed")
	}
	
	// TODO: Implement seek logic (simple version)
	// whence: 0 = from start, 1 = from current, 2 = from end
	// Update f.position and return new position
	
}

// TODO: Function that works with any Reader
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

// TODO: Function that works with any Writer
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

// TODO: Function that uses composed interface
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

// TODO: Function that demonstrates interface upgrade/downgrade
func processFile(rwc ReadWriteCloser) {
	fmt.Printf("Processing ReadWriteCloser: %T\n", rwc)
	
	// Use as Writer
	message := "Hello from composed interface!\n"
	writeAll(rwc, []byte(message))
	
	// Check if it also implements Seeker (interface upgrade)
	if seeker, ok := /* check if rwc implements Seeker */; ok {
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
	
	// TODO: Create a File instance
	file := NewFile("test.txt", "Initial content")
	
	// TODO: Use File as different interface types
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
	
	// TODO: Use standard library types that implement our interfaces
	stringReader := strings.NewReader("Hello from strings.Reader!")
	
	fmt.Println("Reading from strings.Reader:")
	data, err = readAll(stringReader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Read: %s\n", data)
	}
	
	// TODO: Demonstrate interface composition with different implementations
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