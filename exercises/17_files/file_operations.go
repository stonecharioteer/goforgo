// file_operations.go
// Learn file I/O operations in Go

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("=== Basic File Writing ===")
	
	// TODO: Write string to file
	filename := "test.txt"
	content := "Hello, Go file operations!\\nThis is line 2.\\nThis is line 3."
	
	err := /* write content to filename */
	if err != nil {
		fmt.Printf("Write error: %v\\n", err)
		return
	}
	fmt.Printf("Successfully wrote to %s\\n", filename)
	
	fmt.Println("\\n=== Basic File Reading ===")
	
	// TODO: Read entire file
	data, err := /* read all data from filename */
	if err != nil {
		fmt.Printf("Read error: %v\\n", err)
		return
	}
	
	fmt.Printf("File content:\\n%s\\n", data)
	
	fmt.Println("\\n=== File Info ===")
	
	// TODO: Get file information
	info, err := /* get file info for filename */
	if err != nil {
		fmt.Printf("Stat error: %v\\n", err)
	} else {
		fmt.Printf("Name: %s\\n", info.Name())
		fmt.Printf("Size: %d bytes\\n", info.Size())
		fmt.Printf("Mode: %s\\n", info.Mode())
		fmt.Printf("ModTime: %s\\n", info.ModTime())
		fmt.Printf("IsDir: %t\\n", info.IsDir())
	}
	
	fmt.Println("\\n=== Line-by-Line Reading ===")
	
	// TODO: Open file for reading
	file, err := /* open filename for reading */
	if err != nil {
		fmt.Printf("Open error: %v\\n", err)
		return
	}
	defer file.Close()
	
	// TODO: Create scanner to read line by line
	scanner := /* create new scanner for file */
	lineNum := 1
	
	fmt.Println("Reading line by line:")
	for /* scan next line */ {
		line := /* get scanned text */
		fmt.Printf("Line %d: %s\\n", lineNum, line)	
		lineNum++
	}
	
	// TODO: Check for scan errors
	if err := /* get scanner error */; err != nil {
		fmt.Printf("Scan error: %v\\n", err)
	}
	
	fmt.Println("\\n=== Buffered Writing ===")
	
	bufferedFile := "buffered.txt"
	
	// TODO: Create file for writing
	file, err = /* create bufferedFile */
	if err != nil {
		fmt.Printf("Create error: %v\\n", err)
		return
	}
	defer file.Close()
	
	// TODO: Create buffered writer
	writer := /* create buffered writer for file */
	defer writer.Flush() // Important: flush buffer
	
	// TODO: Write multiple lines using buffered writer
	lines := []string{
		"First buffered line",
		"Second buffered line", 
		"Third buffered line",
	}
	
	for _, line := range lines {
		_, err := /* write line + newline to writer */
		if err != nil {
			fmt.Printf("Write error: %v\\n", err)
			return
		}
	}
	
	fmt.Printf("Successfully wrote buffered content to %s\\n", bufferedFile)
	
	fmt.Println("\\n=== File Copying ===")
	
	srcFile := "test.txt"
	dstFile := "copy.txt"
	
	// TODO: Copy file
	err = copyFile(srcFile, dstFile)
	if err != nil {
		fmt.Printf("Copy error: %v\\n", err)
	} else {
		fmt.Printf("Successfully copied %s to %s\\n", srcFile, dstFile)
	}
	
	fmt.Println("\\n=== Directory Operations ===")
	
	dirName := "testdir"
	
	// TODO: Create directory
	err = /* create directory with 0755 permissions */
	if err != nil {
		fmt.Printf("Mkdir error: %v\\n", err)
	} else {
		fmt.Printf("Created directory: %s\\n", dirName)
	}
	
	// TODO: List directory contents
	entries, err := /* read directory entries from current directory */
	if err != nil {
		fmt.Printf("ReadDir error: %v\\n", err)
	} else {
		fmt.Println("Directory contents:")
		for _, entry := range entries {
			fmt.Printf("  %s (dir: %t)\\n", entry.Name(), entry.IsDir())
		}
	}
	
	fmt.Println("\\n=== Path Operations ===")
	
	// TODO: Work with file paths
	fullPath := "/home/user/documents/file.txt"
	
	fmt.Printf("Full path: %s\\n", fullPath)
	fmt.Printf("Dir: %s\\n", /* get directory part */)
	fmt.Printf("Base: %s\\n", /* get filename part */)
	fmt.Printf("Ext: %s\\n", /* get extension */)
	
	// TODO: Join paths
	joined := /* join "testdir", "subdir", "file.txt" */
	fmt.Printf("Joined path: %s\\n", joined)
	
	// TODO: Get absolute path
	abs, err := /* get absolute path of "." */
	if err != nil {
		fmt.Printf("Abs error: %v\\n", err)
	} else {
		fmt.Printf("Absolute path: %s\\n", abs)
	}
	
	fmt.Println("\\n=== Cleanup ===")
	
	// TODO: Remove created files and directories
	filesToRemove := []string{"test.txt", "buffered.txt", "copy.txt", "testdir"}
	
	for _, f := range filesToRemove {
		err := /* remove f (handles both files and directories) */
		if err != nil {
			fmt.Printf("Remove %s error: %v\\n", f, err)
		} else {
			fmt.Printf("Removed: %s\\n", f)
		}
	}
}

// TODO: Implement file copy function
func copyFile(src, dst string) error {
	// Open source file
	srcFile, err := /* open src for reading */
	if err != nil {
		return err
	}
	defer srcFile.Close()
	
	// Create destination file
	dstFile, err := /* create dst file */
	if err != nil {
		return err
	}
	defer dstFile.Close()
	
	// Copy contents
	_, err = /* copy from srcFile to dstFile */
	return err
}