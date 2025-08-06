// file_watching.go - SOLUTION
// Learn file system monitoring and watching

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println("=== File System Monitoring ===")
	
	// Create test directory and files
	testDir := "test_watch"
	testFile := filepath.Join(testDir, "watched_file.txt")
	
	// Create directory if it doesn't exist
	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}
	
	fmt.Printf("Created watch directory: %s\n", testDir)
	
	// Create initial file
	initialContent := "Initial content\nCreated at: " + time.Now().Format(time.RFC3339)
	err = os.WriteFile(testFile, []byte(initialContent), 0644)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	
	fmt.Printf("Created file: %s\n", testFile)
	
	fmt.Println("\n=== File Stat Monitoring ===")
	
	// Get initial file stats
	initialStat, err := os.Stat(testFile)
	if err != nil {
		fmt.Printf("Error getting file stats: %v\n", err)
		return
	}
	
	fmt.Printf("Initial file stats:\n")
	fmt.Printf("  Size: %d bytes\n", initialStat.Size())
	fmt.Printf("  Mode: %s\n", initialStat.Mode())
	fmt.Printf("  ModTime: %s\n", initialStat.ModTime().Format(time.RFC3339))
	
	// Monitor file changes with polling
	fmt.Println("\nStarting file monitoring (polling method)...")
	fmt.Println("Modify the file in another terminal to see changes")
	
	go func() {
		// Simulate file modifications
		time.Sleep(2 * time.Second)
		
		// Append to file
		appendContent := "\nModified at: " + time.Now().Format(time.RFC3339)
		f, err := os.OpenFile(testFile, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			return
		}
		defer f.Close()
		
		f.WriteString(appendContent)
		fmt.Println("File modified programmatically")
		
		// Create another file
		time.Sleep(2 * time.Second)
		newFile := filepath.Join(testDir, "new_file.txt")
		os.WriteFile(newFile, []byte("New file content"), 0644)
		fmt.Printf("Created new file: %s\n", newFile)
		
		// Delete a file
		time.Sleep(2 * time.Second)
		os.Remove(newFile)
		fmt.Printf("Deleted file: %s\n", newFile)
	}()
	
	// Poll for changes
	lastModTime := initialStat.ModTime()
	pollCount := 0
	maxPolls := 20
	
	for pollCount < maxPolls {
		time.Sleep(500 * time.Millisecond)
		pollCount++
		
		// Check if file still exists
		currentStat, err := os.Stat(testFile)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("Poll %d: File deleted!\n", pollCount)
				break
			}
			fmt.Printf("Poll %d: Error checking file: %v\n", pollCount, err)
			continue
		}
		
		// Check if file was modified
		currentModTime := currentStat.ModTime()
		if currentModTime.After(lastModTime) {
			fmt.Printf("Poll %d: File modified!\n", pollCount)
			fmt.Printf("  New size: %d bytes\n", currentStat.Size())
			fmt.Printf("  New ModTime: %s\n", currentModTime.Format(time.RFC3339))
			
			lastModTime = currentModTime
		}
	}
	
	fmt.Println("\n=== Directory Monitoring ===")
	
	// Monitor directory for new/deleted files
	initialFiles, err := os.ReadDir(testDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}
	
	fmt.Printf("Initial directory contents (%d items):\n", len(initialFiles))
	for _, file := range initialFiles {
		fmt.Printf("  - %s (dir: %t)\n", file.Name(), file.IsDir())
	}
	
	// Monitor for directory changes
	fmt.Println("\nMonitoring directory for changes...")
	
	go func() {
		// Create multiple files
		for i := 1; i <= 3; i++ {
			time.Sleep(1 * time.Second)
			filename := filepath.Join(testDir, fmt.Sprintf("file_%d.txt", i))
			content := fmt.Sprintf("Content of file %d", i)
			os.WriteFile(filename, []byte(content), 0644)
			fmt.Printf("Created: %s\n", filename)
		}
		
		// Remove some files
		time.Sleep(2 * time.Second)
		os.Remove(filepath.Join(testDir, "file_1.txt"))
		fmt.Println("Removed: file_1.txt")
	}()
	
	// Poll directory changes
	lastFileCount := len(initialFiles)
	dirPollCount := 0
	maxDirPolls := 15
	
	for dirPollCount < maxDirPolls {
		time.Sleep(500 * time.Millisecond)
		dirPollCount++
		
		// Read current directory contents
		currentFiles, err := os.ReadDir(testDir)
		if err != nil {
			fmt.Printf("Error reading directory: %v\n", err)
			continue
		}
		
		currentFileCount := len(currentFiles)
		if currentFileCount != lastFileCount {
			fmt.Printf("Directory poll %d: File count changed from %d to %d\n", 
				dirPollCount, lastFileCount, currentFileCount)
			
			lastFileCount = currentFileCount
		}
	}
	
	fmt.Println("\n=== File Permission Monitoring ===")
	
	// Monitor file permission changes
	permFile := filepath.Join(testDir, "perm_test.txt")
	os.WriteFile(permFile, []byte("Permission test"), 0644)
	
	// Get initial permissions
	initialStat, err = os.Stat(permFile)
	if err != nil {
		fmt.Printf("Error getting permissions: %v\n", err)
	} else {
		fmt.Printf("Initial permissions: %s\n", initialStat.Mode())
	}
	
	// Change permissions
	newPerm := os.FileMode(0755)
	err = os.Chmod(permFile, newPerm)
	if err != nil {
		fmt.Printf("Error changing permissions: %v\n", err)
	} else {
		fmt.Printf("Changed permissions to: %s\n", newPerm)
	}
	
	// Verify permission change
	currentStat, err := os.Stat(permFile)
	if err != nil {
		fmt.Printf("Error checking new permissions: %v\n", err)
	} else {
		fmt.Printf("Current permissions: %s\n", currentStat.Mode())
	}
	
	fmt.Println("\n=== Cleanup ===")
	
	// Clean up test directory
	err = os.RemoveAll(testDir)
	if err != nil {
		fmt.Printf("Error cleaning up: %v\n", err)
	} else {
		fmt.Printf("Cleaned up test directory: %s\n", testDir)
	}
	
	fmt.Println("File monitoring demo completed!")
}