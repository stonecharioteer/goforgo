// file_watching.go
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
	
	// TODO: Create test directory and files
	testDir := "test_watch"
	testFile := filepath.Join(testDir, "watched_file.txt")
	
	// TODO: Create directory if it doesn't exist
	err := /* create directory with permissions */
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}
	
	fmt.Printf("Created watch directory: %s\n", testDir)
	
	// TODO: Create initial file
	initialContent := "Initial content\nCreated at: " + time.Now().Format(time.RFC3339)
	err = /* write initial content to file */
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	
	fmt.Printf("Created file: %s\n", testFile)
	
	fmt.Println("\n=== File Stat Monitoring ===")
	
	// TODO: Get initial file stats
	initialStat, err := /* get file stats */
	if err != nil {
		fmt.Printf("Error getting file stats: %v\n", err)
		return
	}
	
	fmt.Printf("Initial file stats:\n")
	fmt.Printf("  Size: %d bytes\n", /* get file size */)
	fmt.Printf("  Mode: %s\n", /* get file mode */)
	fmt.Printf("  ModTime: %s\n", /* get modification time */)
	
	// TODO: Monitor file changes with polling
	fmt.Println("\nStarting file monitoring (polling method)...")
	fmt.Println("Modify the file in another terminal to see changes")
	
	go func() {
		// TODO: Simulate file modifications
		time.Sleep(2 * time.Second)
		
		// TODO: Append to file
		appendContent := "\nModified at: " + time.Now().Format(time.RFC3339)
		f, err := /* open file for append */
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			return
		}
		defer f.Close()
		
		/* write append content */
		fmt.Println("File modified programmatically")
		
		// TODO: Create another file
		time.Sleep(2 * time.Second)
		newFile := filepath.Join(testDir, "new_file.txt")
		/* create new file with content */
		fmt.Printf("Created new file: %s\n", newFile)
		
		// TODO: Delete a file
		time.Sleep(2 * time.Second)
		/* remove new file */
		fmt.Printf("Deleted file: %s\n", newFile)
	}()
	
	// TODO: Poll for changes
	lastModTime := /* get initial modification time */
	pollCount := 0
	maxPolls := 20
	
	for pollCount < maxPolls {
		time.Sleep(500 * time.Millisecond)
		pollCount++
		
		// TODO: Check if file still exists
		currentStat, err := /* get current file stats */
		if err != nil {
			if /* check if file not found */ {
				fmt.Printf("Poll %d: File deleted!\n", pollCount)
				break
			}
			fmt.Printf("Poll %d: Error checking file: %v\n", pollCount, err)
			continue
		}
		
		// TODO: Check if file was modified
		currentModTime := /* get current modification time */
		if /* compare modification times */ {
			fmt.Printf("Poll %d: File modified!\n", pollCount)
			fmt.Printf("  New size: %d bytes\n", /* get new file size */)
			fmt.Printf("  New ModTime: %s\n", currentModTime.Format(time.RFC3339))
			
			/* update last modification time */
		}
	}
	
	fmt.Println("\n=== Directory Monitoring ===")
	
	// TODO: Monitor directory for new/deleted files
	initialFiles, err := /* read directory contents */
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}
	
	fmt.Printf("Initial directory contents (%d items):\n", len(initialFiles))
	for _, file := range initialFiles {
		/* print file info */
	}
	
	// TODO: Monitor for directory changes
	fmt.Println("\nMonitoring directory for changes...")
	
	go func() {
		// TODO: Create multiple files
		for i := 1; i <= 3; i++ {
			time.Sleep(1 * time.Second)
			filename := filepath.Join(testDir, fmt.Sprintf("file_%d.txt", i))
			/* create file with content */
			fmt.Printf("Created: %s\n", filename)
		}
		
		// TODO: Remove some files
		time.Sleep(2 * time.Second)
		/* remove first created file */
		fmt.Println("Removed: file_1.txt")
	}()
	
	// TODO: Poll directory changes
	lastFileCount := len(initialFiles)
	dirPollCount := 0
	maxDirPolls := 15
	
	for dirPollCount < maxDirPolls {
		time.Sleep(500 * time.Millisecond)
		dirPollCount++
		
		// TODO: Read current directory contents
		currentFiles, err := /* read directory */
		if err != nil {
			fmt.Printf("Error reading directory: %v\n", err)
			continue
		}
		
		currentFileCount := len(currentFiles)
		if /* file count changed */ {
			fmt.Printf("Directory poll %d: File count changed from %d to %d\n", 
				dirPollCount, lastFileCount, currentFileCount)
			
			/* update last file count */
		}
	}
	
	fmt.Println("\n=== File Permission Monitoring ===")
	
	// TODO: Monitor file permission changes
	permFile := filepath.Join(testDir, "perm_test.txt")
	/* create permission test file */
	
	// TODO: Get initial permissions
	initialPerm, err := /* get file permissions */
	if err != nil {
		fmt.Printf("Error getting permissions: %v\n", err)
	} else {
		fmt.Printf("Initial permissions: %s\n", initialPerm)
	}
	
	// TODO: Change permissions
	newPerm := /* set new permission mode */
	err = /* change file permissions */
	if err != nil {
		fmt.Printf("Error changing permissions: %v\n", err)
	} else {
		fmt.Printf("Changed permissions to: %s\n", newPerm)
	}
	
	// TODO: Verify permission change
	currentPerm, err := /* get updated permissions */
	if err != nil {
		fmt.Printf("Error checking new permissions: %v\n", err)
	} else {
		fmt.Printf("Current permissions: %s\n", currentPerm)
	}
	
	fmt.Println("\n=== Cleanup ===")
	
	// TODO: Clean up test directory
	err = /* remove entire test directory */
	if err != nil {
		fmt.Printf("Error cleaning up: %v\n", err)
	} else {
		fmt.Printf("Cleaned up test directory: %s\n", testDir)
	}
	
	fmt.Println("File monitoring demo completed!")
}