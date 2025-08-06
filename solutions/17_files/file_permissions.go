// file_permissions.go
// Learn file permissions and metadata operations in Go

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

func main() {
	fmt.Println("=== File Permissions and Metadata ===")
	
	// Create test files for demonstration
	testFile := "test_permissions.txt"
	testDir := "test_directory"
	
	// Create test file with content
	err := os.WriteFile(testFile, []byte("Test content for permissions"), 0644)
	if err != nil {
		fmt.Printf("Error creating test file: %v\n", err)
		return
	}
	defer os.Remove(testFile)
	
	// Create test directory
	err = os.Mkdir(testDir, 0755)
	if err != nil {
		fmt.Printf("Error creating test directory: %v\n", err)
		return
	}
	defer os.RemoveAll(testDir)
	
	fmt.Println("\n=== File Information ===")
	
	// Get file information
	fileInfo, err := os.Stat(testFile)
	if err != nil {
		fmt.Printf("Error getting file info: %v\n", err)
		return
	}
	
	// Display basic file information
	fmt.Printf("Name: %s\n", fileInfo.Name())
	fmt.Printf("Size: %d bytes\n", fileInfo.Size())
	fmt.Printf("Mode: %s (%o)\n", fileInfo.Mode(), fileInfo.Mode())
	fmt.Printf("Modified: %s\n", fileInfo.ModTime())
	fmt.Printf("IsDir: %t\n", fileInfo.IsDir())
	
	fmt.Println("\n=== Permission Checks ===")
	
	// Check file permissions
	mode := fileInfo.Mode()
	fmt.Printf("Permissions: %s\n", mode.Perm())
	
	// Check specific permissions
	fmt.Printf("Owner can read: %t\n", mode&0400 != 0)
	fmt.Printf("Owner can write: %t\n", mode&0200 != 0)
	fmt.Printf("Owner can execute: %t\n", mode&0100 != 0)
	fmt.Printf("Group can read: %t\n", mode&0040 != 0)
	fmt.Printf("Others can read: %t\n", mode&0004 != 0)
	
	fmt.Println("\n=== Changing Permissions ===")
	
	// Change file permissions
	newMode := os.FileMode(0600)
	err = os.Chmod(testFile, newMode)
	if err != nil {
		fmt.Printf("Error changing permissions: %v\n", err)
	} else {
		fmt.Printf("Changed permissions to: %o\n", newMode)
	}
	
	// Verify the change
	fileInfo, _ = os.Stat(testFile)
	fmt.Printf("New permissions: %s (%o)\n", fileInfo.Mode().Perm(), fileInfo.Mode().Perm())
	
	fmt.Println("\n=== File System Information ===")
	
	// Get file system stats (Unix-specific)
	var stat syscall.Stat_t
	err = syscall.Stat(testFile, &stat)
	if err != nil {
		fmt.Printf("Error getting syscall stat: %v\n", err)
	} else {
		fmt.Printf("Inode: %d\n", stat.Ino)
		fmt.Printf("Device: %d\n", stat.Dev)
		fmt.Printf("Links: %d\n", stat.Nlink)
		fmt.Printf("UID: %d\n", stat.Uid)
		fmt.Printf("GID: %d\n", stat.Gid)
		fmt.Printf("Size: %d\n", stat.Size)
		fmt.Printf("Access time: %s\n", time.Unix(stat.Atim.Sec, stat.Atim.Nsec))
		fmt.Printf("Modify time: %s\n", time.Unix(stat.Mtim.Sec, stat.Mtim.Nsec))
		fmt.Printf("Change time: %s\n", time.Unix(stat.Ctim.Sec, stat.Ctim.Nsec))
	}
	
	fmt.Println("\n=== Directory Walking with Permissions ===")
	
	// Create nested directory structure
	nestedDir := filepath.Join(testDir, "subdir", "nested")
	os.MkdirAll(nestedDir, 0755)
	
	// Create files with different permissions
	files := map[string]os.FileMode{
		filepath.Join(testDir, "public.txt"):      0644,
		filepath.Join(testDir, "private.txt"):     0600,
		filepath.Join(testDir, "executable.sh"):   0755,
		filepath.Join(nestedDir, "secret.txt"):    0400, // read-only
	}
	
	// Create files with specific permissions
	for filePath, mode := range files {
		// Ensure directory exists
		os.MkdirAll(filepath.Dir(filePath), 0755)
		
		// Create file with content and specific mode
		err := os.WriteFile(filePath, []byte("content"), mode)
		if err != nil {
			fmt.Printf("Error creating %s: %v\n", filePath, err)
		}
	}
	
	// Walk directory and display file permissions
	fmt.Printf("Walking directory: %s\n", testDir)
	err = filepath.Walk(testDir, walkFunc)
	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		return
	}
	
	fmt.Println("\n=== Access Checks ===")
	
	// Check if files are accessible
	testFiles := []string{
		testFile,
		filepath.Join(testDir, "public.txt"),
		filepath.Join(testDir, "private.txt"),
		filepath.Join(testDir, "nonexistent.txt"),
	}
	
	for _, file := range testFiles {
		// Check different types of access
		fmt.Printf("File: %s\n", file)
		
		// Check if file exists
		_, err := os.Stat(file)
		fmt.Printf("  Exists: %t\n", err == nil)
		
		if err == nil {
			// Check read access
			f, err := os.Open(file)
			readable := err == nil
			if f != nil {
				f.Close()
			}
			fmt.Printf("  Readable: %t\n", readable)
			
			// Check write access by trying to open for writing
			f, err = os.OpenFile(file, os.O_WRONLY, 0)
			writable := err == nil
			if f != nil {
				f.Close()
			}
			fmt.Printf("  Writable: %t\n", writable)
		}
		fmt.Println()
	}
	
	fmt.Println("\n=== Temporary Files with Permissions ===")
	
	// Create temporary file with specific permissions
	tempFile, err := os.CreateTemp("", "perm_test_*.txt")
	if err != nil {
		fmt.Printf("Error creating temp file: %v\n", err)
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()
	
	// Change temp file permissions
	err = os.Chmod(tempFile.Name(), 0600)
	if err != nil {
		fmt.Printf("Error changing temp file permissions: %v\n", err)
	}
	
	// Write content and verify
	tempFile.WriteString("Temporary file with custom permissions")
	tempInfo, _ := tempFile.Stat()
	fmt.Printf("Temp file: %s\n", tempFile.Name())
	fmt.Printf("Temp file permissions: %s (%o)\n", tempInfo.Mode().Perm(), tempInfo.Mode().Perm())
	
	fmt.Println("\n=== Cleanup Complete ===")
}

// Helper function for walking directories
func walkFunc(path string, info fs.FileInfo, err error) error {
	if err != nil {
		return err
	}
	
	// Display file information during walk
	relPath, _ := filepath.Rel(".", path)
	fileType := "FILE"
	if info.IsDir() {
		fileType = "DIR "
	}
	
	fmt.Printf("  %s %s %s (%o)\n", 
		fileType,
		info.Mode().Perm(),
		relPath,
		info.Mode().Perm())
	
	return nil
}