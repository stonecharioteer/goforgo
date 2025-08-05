// file_operations.go - SOLUTION
// Learn file I/O operations in Go

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("=== Basic File Writing ===")
	
	// Write string to file
	filename := "test.txt"
	content := "Hello, Go file operations!\nThis is line 2.\nThis is line 3."
	
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Write error: %v\n", err)
		return
	}
	fmt.Printf("Successfully wrote to %s\n", filename)
	
	fmt.Println("\n=== Basic File Reading ===")
	
	// Read entire file
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Read error: %v\n", err)
		return
	}
	
	fmt.Printf("File content:\n%s\n", data)
	
	fmt.Println("\n=== File Info ===")
	
	// Get file information
	info, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("Stat error: %v\n", err)
	} else {
		fmt.Printf("Name: %s\n", info.Name())
		fmt.Printf("Size: %d bytes\n", info.Size())
		fmt.Printf("Mode: %s\n", info.Mode())
		fmt.Printf("ModTime: %s\n", info.ModTime())
		fmt.Printf("IsDir: %t\n", info.IsDir())
	}
	
	fmt.Println("\n=== Line-by-Line Reading ===")
	
	// Open file for reading
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Open error: %v\n", err)
		return
	}
	defer file.Close()
	
	// Create scanner to read line by line
	scanner := bufio.NewScanner(file)
	lineNum := 1
	
	fmt.Println("Reading line by line:")
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line %d: %s\n", lineNum, line)	
		lineNum++
	}
	
	// Check for scan errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Scan error: %v\n", err)
	}
	
	fmt.Println("\n=== Buffered Writing ===")
	
	bufferedFile := "buffered.txt"
	
	// Create file for writing
	file, err = os.Create(bufferedFile)
	if err != nil {
		fmt.Printf("Create error: %v\n", err)
		return
	}
	defer file.Close()
	
	// Create buffered writer
	writer := bufio.NewWriter(file)
	defer writer.Flush() // Important: flush buffer
	
	// Write multiple lines using buffered writer
	lines := []string{
		"First buffered line",
		"Second buffered line", 
		"Third buffered line",
	}
	
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Printf("Write error: %v\n", err)
			return
		}
	}
	
	fmt.Printf("Successfully wrote buffered content to %s\n", bufferedFile)
	
	fmt.Println("\n=== File Copying ===")
	
	srcFile := "test.txt"
	dstFile := "copy.txt"
	
	// Copy file
	err = copyFile(srcFile, dstFile)
	if err != nil {
		fmt.Printf("Copy error: %v\n", err)
	} else {
		fmt.Printf("Successfully copied %s to %s\n", srcFile, dstFile)
	}
	
	fmt.Println("\n=== Directory Operations ===")
	
	dirName := "testdir"
	
	// Create directory
	err = os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Printf("Mkdir error: %v\n", err)
	} else {
		fmt.Printf("Created directory: %s\n", dirName)
	}
	
	// List directory contents
	entries, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("ReadDir error: %v\n", err)
	} else {
		fmt.Println("Directory contents:")
		for _, entry := range entries {
			fmt.Printf("  %s (dir: %t)\n", entry.Name(), entry.IsDir())
		}
	}
	
	fmt.Println("\n=== Path Operations ===")
	
	// Work with file paths
	fullPath := "/home/user/documents/file.txt"
	
	fmt.Printf("Full path: %s\n", fullPath)
	fmt.Printf("Dir: %s\n", filepath.Dir(fullPath))
	fmt.Printf("Base: %s\n", filepath.Base(fullPath))
	fmt.Printf("Ext: %s\n", filepath.Ext(fullPath))
	
	// Join paths
	joined := filepath.Join("testdir", "subdir", "file.txt")
	fmt.Printf("Joined path: %s\n", joined)
	
	// Get absolute path
	abs, err := filepath.Abs(".")
	if err != nil {
		fmt.Printf("Abs error: %v\n", err)
	} else {
		fmt.Printf("Absolute path: %s\n", abs)
	}
	
	fmt.Println("\n=== Cleanup ===")
	
	// Remove created files and directories
	filesToRemove := []string{"test.txt", "buffered.txt", "copy.txt", "testdir"}
	
	for _, f := range filesToRemove {
		err := os.RemoveAll(f)
		if err != nil {
			fmt.Printf("Remove %s error: %v\n", f, err)
		} else {
			fmt.Printf("Removed: %s\n", f)
		}
	}
}

// Implement file copy function
func copyFile(src, dst string) error {
	// Open source file
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	
	// Create destination file
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	
	// Copy contents
	_, err = io.Copy(dstFile, srcFile)
	return err
}