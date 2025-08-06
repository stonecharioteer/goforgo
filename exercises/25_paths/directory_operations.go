// directory_operations.go
// Learn directory operations: walking, globbing, watching, and directory management

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("=== Directory Operations ===")
	
	fmt.Println("\n=== Directory Creation and Management ===")
	
	// TODO: Create a test directory structure
	baseDir := "test_directory_ops"
	
	// TODO: Create base directory
	err := /* create directory */
	if err != nil {
		fmt.Printf("Error creating base directory: %v\n", err)
		return
	}
	
	// TODO: Create nested directory structure
	dirs := []string{
		"src/main",
		"src/lib",
		"docs/api",
		"docs/guides", 
		"tests/unit",
		"tests/integration",
		"config",
		"assets/images",
		"assets/styles",
	}
	
	fmt.Println("Creating directory structure:")
	for _, dir := range dirs {
		fullPath := /* join base directory with subdirectory */
		err := /* create directory with parents */
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			continue
		}
		fmt.Printf("  ✅ Created: %s\n", dir)
	}
	
	// TODO: Create some test files
	testFiles := []string{
		"src/main/main.go",
		"src/main/config.go",
		"src/lib/utils.go",
		"src/lib/helpers.go",
		"docs/api/README.md",
		"docs/guides/quickstart.md",
		"tests/unit/main_test.go",
		"tests/integration/api_test.go",
		"config/app.json",
		"config/database.yaml",
		"assets/images/logo.png",
		"assets/styles/main.css",
		"README.md",
		"go.mod",
		".gitignore",
	}
	
	fmt.Println("\nCreating test files:")
	for _, file := range testFiles {
		fullPath := /* join base directory with file path */
		
		// TODO: Create file with some content
		content := fmt.Sprintf("// %s\n// Generated test file\n", file)
		err := /* write file with content */
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", file, err)
			continue
		}
		fmt.Printf("  ✅ Created: %s\n", file)
	}
	
	fmt.Println("\n=== Directory Walking ===")
	
	// TODO: Walk directory tree and collect information
	var totalFiles, totalDirs int
	var totalSize int64
	
	fmt.Printf("Walking directory tree: %s\n", baseDir)
	
	err = /* walk directory tree */
	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		return
	}
	
	fmt.Printf("Summary:\n")
	fmt.Printf("  Total directories: %d\n", totalDirs)
	fmt.Printf("  Total files: %d\n", totalFiles)
	fmt.Printf("  Total size: %d bytes\n", totalSize)
	
	fmt.Println("\n=== File Pattern Matching (Glob) ===")
	
	// TODO: Find files by patterns
	patterns := []string{
		"*.go",
		"*.md", 
		"*.json",
		"**/test*.go",
		"src/**/*.go",
		"docs/**/*",
	}
	
	fmt.Println("Finding files by patterns:")
	for _, pattern := range patterns {
		fullPattern := /* join base directory with pattern */
		
		// TODO: Find files matching pattern
		matches, err := /* glob pattern */
		if err != nil {
			fmt.Printf("  Error with pattern %s: %v\n", pattern, err)
			continue
		}
		
		fmt.Printf("  Pattern '%s': %d matches\n", pattern, len(matches))
		for _, match := range matches {
			// TODO: Make path relative to base directory for cleaner output
			relPath, _ := /* make relative to base directory */
			fmt.Printf("    - %s\n", relPath)
		}
	}
	
	fmt.Println("\n=== Directory Information ===")
	
	// TODO: Get directory information
	dirsToInspect := []string{
		"src",
		"docs", 
		"tests",
		"config",
		"assets",
	}
	
	fmt.Println("Directory information:")
	for _, dir := range dirsToInspect {
		fullPath := /* join base directory with subdirectory */
		
		// TODO: Get directory info
		info, err := /* get file info */
		if err != nil {
			fmt.Printf("  Error getting info for %s: %v\n", dir, err)
			continue
		}
		
		fmt.Printf("  %s:\n", dir)
		fmt.Printf("    Mode: %s\n", info.Mode())
		fmt.Printf("    Modified: %s\n", info.ModTime().Format("2006-01-02 15:04:05"))
		fmt.Printf("    Is Directory: %t\n", info.IsDir())
		
		// TODO: Count items in directory
		entries, err := /* read directory entries */
		if err != nil {
			fmt.Printf("    Error reading directory: %v\n", err)
		} else {
			var fileCount, dirCount int
			for _, entry := range entries {
				if entry.IsDir() {
					dirCount++
				} else {
					fileCount++
				}
			}
			fmt.Printf("    Contents: %d files, %d directories\n", fileCount, dirCount)
		}
	}
	
	fmt.Println("\n=== Directory Utilities ===")
	
	// TODO: Implement directory utilities
	fmt.Println("Directory utilities:")
	
	// TODO: Find empty directories
	emptyDirs := /* find empty directories */
	fmt.Printf("Empty directories: %d\n", len(emptyDirs))
	for _, dir := range emptyDirs {
		relPath, _ := filepath.Rel(baseDir, dir)
		fmt.Printf("  - %s\n", relPath)
	}
	
	// TODO: Find largest files
	largestFiles := /* find files sorted by size */
	fmt.Printf("Largest files (top 5):\n")
	for i, file := range largestFiles {
		if i >= 5 {
			break
		}
		info, _ := os.Stat(file)
		relPath, _ := filepath.Rel(baseDir, file)
		fmt.Printf("  %d. %s (%d bytes)\n", i+1, relPath, info.Size())
	}
	
	// TODO: Group files by extension
	filesByExt := /* group files by extension */
	fmt.Println("Files by extension:")
	for ext, files := range filesByExt {
		if ext == "" {
			ext = "(no extension)"
		}
		fmt.Printf("  %s: %d files\n", ext, len(files))
	}
	
	fmt.Println("\n=== Directory Tree Visualization ===")
	
	// TODO: Print directory tree
	fmt.Printf("Directory tree for %s:\n", baseDir)
	/* print directory tree */
	
	fmt.Println("\n=== Directory Comparison ===")
	
	// TODO: Create a second directory for comparison
	compareDir := "test_directory_compare"
	err = /* create comparison directory */
	if err != nil {
		fmt.Printf("Error creating comparison directory: %v\n", err)
	} else {
		// TODO: Copy some files to comparison directory
		filesToCopy := []string{
			"README.md",
			"go.mod",
			"src/main/main.go",
		}
		
		for _, file := range filesToCopy {
			srcPath := filepath.Join(baseDir, file)
			dstPath := filepath.Join(compareDir, file)
			
			// TODO: Ensure destination directory exists
			dstDir := filepath.Dir(dstPath)
			os.MkdirAll(dstDir, 0755)
			
			// TODO: Copy file
			err := /* copy file from src to dst */
			if err != nil {
				fmt.Printf("Error copying %s: %v\n", file, err)
			}
		}
		
		// TODO: Compare directories
		fmt.Printf("Comparing %s and %s:\n", baseDir, compareDir)
		/* compare directories and show differences */
	}
	
	fmt.Println("\n=== Cleanup ===")
	
	// TODO: Clean up test directories
	fmt.Println("Cleaning up test directories...")
	
	for _, dir := range []string{baseDir, compareDir} {
		if _, err := os.Stat(dir); err == nil {
			err = /* remove directory and all contents */
			if err != nil {
				fmt.Printf("❌ Failed to remove %s: %v\n", dir, err)
			} else {
				fmt.Printf("✅ Removed %s\n", dir)
			}
		}
	}
	
	fmt.Println("\n=== Best Practices ===")
	fmt.Println("Directory operations best practices:")
	fmt.Println("✅ Always check for errors when creating/removing directories")
	fmt.Println("✅ Use filepath.Join() for cross-platform path construction")
	fmt.Println("✅ Use os.MkdirAll() to create parent directories automatically")
	fmt.Println("✅ Be careful with os.RemoveAll() - it's irreversible")
	fmt.Println("✅ Use filepath.Walk() for recursive directory traversal")
	fmt.Println("✅ Handle permissions appropriately (0755 for directories, 0644 for files)")
	fmt.Println("✅ Use defer for cleanup operations")
	fmt.Println("✅ Consider using temporary directories for testing")
	fmt.Println("✅ Always validate paths to prevent directory traversal attacks")
}