// directory_operations.go - SOLUTION
// Learn directory operations: walking, globbing, watching, and directory management

package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	fmt.Println("=== Directory Operations ===")
	
	fmt.Println("\n=== Directory Creation and Management ===")
	
	// Create a test directory structure
	baseDir := "test_directory_ops"
	
	// Create base directory
	err := os.MkdirAll(baseDir, 0755)
	if err != nil {
		fmt.Printf("Error creating base directory: %v\n", err)
		return
	}
	
	// Create nested directory structure
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
		fullPath := filepath.Join(baseDir, dir)
		err := os.MkdirAll(fullPath, 0755)
		if err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			continue
		}
		fmt.Printf("  ✅ Created: %s\n", dir)
	}
	
	// Create some test files
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
		fullPath := filepath.Join(baseDir, file)
		
		// Create file with some content
		content := fmt.Sprintf("// %s\n// Generated test file\npackage main\n", file)
		err := os.WriteFile(fullPath, []byte(content), 0644)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", file, err)
			continue
		}
		fmt.Printf("  ✅ Created: %s\n", file)
	}
	
	fmt.Println("\n=== Directory Walking ===")
	
	// Walk directory tree and collect information
	var totalFiles, totalDirs int
	var totalSize int64
	
	fmt.Printf("Walking directory tree: %s\n", baseDir)
	
	err = filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if info.IsDir() {
			totalDirs++
		} else {
			totalFiles++
			totalSize += info.Size()
		}
		
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		return
	}
	
	fmt.Printf("Summary:\n")
	fmt.Printf("  Total directories: %d\n", totalDirs)
	fmt.Printf("  Total files: %d\n", totalFiles)
	fmt.Printf("  Total size: %d bytes\n", totalSize)
	
	fmt.Println("\n=== File Pattern Matching (Glob) ===")
	
	// Find files by patterns
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
		fullPattern := filepath.Join(baseDir, pattern)
		
		// Find files matching pattern
		matches, err := filepath.Glob(fullPattern)
		if err != nil {
			fmt.Printf("  Error with pattern %s: %v\n", pattern, err)
			continue
		}
		
		// For ** patterns, we need to use Walk since filepath.Glob doesn't support **
		if strings.Contains(pattern, "**") {
			matches = []string{}
			filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return nil
				}
				
				rel, _ := filepath.Rel(baseDir, path)
				if matchPattern(rel, pattern) {
					matches = append(matches, path)
				}
				return nil
			})
		}
		
		fmt.Printf("  Pattern '%s': %d matches\n", pattern, len(matches))
		for _, match := range matches {
			// Make path relative to base directory for cleaner output
			relPath, _ := filepath.Rel(baseDir, match)
			fmt.Printf("    - %s\n", relPath)
		}
	}
	
	fmt.Println("\n=== Directory Information ===")
	
	// Get directory information
	dirsToInspect := []string{
		"src",
		"docs", 
		"tests",
		"config",
		"assets",
	}
	
	fmt.Println("Directory information:")
	for _, dir := range dirsToInspect {
		fullPath := filepath.Join(baseDir, dir)
		
		// Get directory info
		info, err := os.Stat(fullPath)
		if err != nil {
			fmt.Printf("  Error getting info for %s: %v\n", dir, err)
			continue
		}
		
		fmt.Printf("  %s:\n", dir)
		fmt.Printf("    Mode: %s\n", info.Mode())
		fmt.Printf("    Modified: %s\n", info.ModTime().Format("2006-01-02 15:04:05"))
		fmt.Printf("    Is Directory: %t\n", info.IsDir())
		
		// Count items in directory
		entries, err := os.ReadDir(fullPath)
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
	
	// Implement directory utilities
	fmt.Println("Directory utilities:")
	
	// Find empty directories
	emptyDirs := findEmptyDirectories(baseDir)
	fmt.Printf("Empty directories: %d\n", len(emptyDirs))
	for _, dir := range emptyDirs {
		relPath, _ := filepath.Rel(baseDir, dir)
		fmt.Printf("  - %s\n", relPath)
	}
	
	// Find largest files
	largestFiles := findLargestFiles(baseDir)
	fmt.Printf("Largest files (top 5):\n")
	for i, file := range largestFiles {
		if i >= 5 {
			break
		}
		info, _ := os.Stat(file)
		relPath, _ := filepath.Rel(baseDir, file)
		fmt.Printf("  %d. %s (%d bytes)\n", i+1, relPath, info.Size())
	}
	
	// Group files by extension
	filesByExt := groupFilesByExtension(baseDir)
	fmt.Println("Files by extension:")
	for ext, files := range filesByExt {
		if ext == "" {
			ext = "(no extension)"
		}
		fmt.Printf("  %s: %d files\n", ext, len(files))
	}
	
	fmt.Println("\n=== Directory Tree Visualization ===")
	
	// Print directory tree
	fmt.Printf("Directory tree for %s:\n", baseDir)
	printDirectoryTree(baseDir, "", true)
	
	fmt.Println("\n=== Directory Comparison ===")
	
	// Create a second directory for comparison
	compareDir := "test_directory_compare"
	err = os.MkdirAll(compareDir, 0755)
	if err != nil {
		fmt.Printf("Error creating comparison directory: %v\n", err)
	} else {
		// Copy some files to comparison directory
		filesToCopy := []string{
			"README.md",
			"go.mod",
			"src/main/main.go",
		}
		
		for _, file := range filesToCopy {
			srcPath := filepath.Join(baseDir, file)
			dstPath := filepath.Join(compareDir, file)
			
			// Ensure destination directory exists
			dstDir := filepath.Dir(dstPath)
			os.MkdirAll(dstDir, 0755)
			
			// Copy file
			err := copyFile(srcPath, dstPath)
			if err != nil {
				fmt.Printf("Error copying %s: %v\n", file, err)
			}
		}
		
		// Compare directories
		fmt.Printf("Comparing %s and %s:\n", baseDir, compareDir)
		compareDirectories(baseDir, compareDir)
	}
	
	fmt.Println("\n=== Cleanup ===")
	
	// Clean up test directories
	fmt.Println("Cleaning up test directories...")
	
	for _, dir := range []string{baseDir, compareDir} {
		if _, err := os.Stat(dir); err == nil {
			err = os.RemoveAll(dir)
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

// Helper functions

func matchPattern(path, pattern string) bool {
	// Simple pattern matching for ** (recursive)
	if strings.Contains(pattern, "**") {
		parts := strings.Split(pattern, "**")
		if len(parts) != 2 {
			return false
		}
		
		prefix := strings.TrimSuffix(parts[0], "/")
		suffix := strings.TrimPrefix(parts[1], "/")
		
		if prefix != "" && !strings.HasPrefix(path, prefix) {
			return false
		}
		
		if suffix != "" {
			matched, _ := filepath.Match(suffix, filepath.Base(path))
			return matched
		}
		
		return true
	}
	
	matched, _ := filepath.Match(pattern, path)
	return matched
}

func findEmptyDirectories(root string) []string {
	var emptyDirs []string
	
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || !info.IsDir() {
			return nil
		}
		
		entries, err := os.ReadDir(path)
		if err != nil {
			return nil
		}
		
		if len(entries) == 0 {
			emptyDirs = append(emptyDirs, path)
		}
		
		return nil
	})
	
	return emptyDirs
}

func findLargestFiles(root string) []string {
	type fileInfo struct {
		path string
		size int64
	}
	
	var files []fileInfo
	
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		
		files = append(files, fileInfo{path: path, size: info.Size()})
		return nil
	})
	
	// Sort by size (descending)
	sort.Slice(files, func(i, j int) bool {
		return files[i].size > files[j].size
	})
	
	var paths []string
	for _, f := range files {
		paths = append(paths, f.path)
	}
	
	return paths
}

func groupFilesByExtension(root string) map[string][]string {
	filesByExt := make(map[string][]string)
	
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		
		ext := filepath.Ext(path)
		filesByExt[ext] = append(filesByExt[ext], path)
		return nil
	})
	
	return filesByExt
}

func printDirectoryTree(root, prefix string, isLast bool) {
	info, err := os.Stat(root)
	if err != nil {
		return
	}
	
	connector := "├── "
	if isLast {
		connector = "└── "
	}
	
	if prefix == "" {
		fmt.Printf("%s\n", filepath.Base(root))
	} else {
		fmt.Printf("%s%s%s\n", prefix, connector, info.Name())
	}
	
	if !info.IsDir() {
		return
	}
	
	entries, err := os.ReadDir(root)
	if err != nil {
		return
	}
	
	newPrefix := prefix
	if prefix != "" {
		if isLast {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
	}
	
	for i, entry := range entries {
		isLastEntry := i == len(entries)-1
		childPath := filepath.Join(root, entry.Name())
		printDirectoryTree(childPath, newPrefix, isLastEntry)
	}
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	
	_, err = io.Copy(dstFile, srcFile)
	return err
}

func compareDirectories(dir1, dir2 string) {
	files1 := getFileList(dir1)
	files2 := getFileList(dir2)
	
	// Files only in dir1
	for file := range files1 {
		if _, exists := files2[file]; !exists {
			fmt.Printf("  Only in %s: %s\n", dir1, file)
		}
	}
	
	// Files only in dir2
	for file := range files2 {
		if _, exists := files1[file]; !exists {
			fmt.Printf("  Only in %s: %s\n", dir2, file)
		}
	}
	
	// Common files
	var common int
	for file := range files1 {
		if _, exists := files2[file]; exists {
			common++
		}
	}
	
	fmt.Printf("  Common files: %d\n", common)
}

func getFileList(root string) map[string]bool {
	files := make(map[string]bool)
	
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		
		rel, _ := filepath.Rel(root, path)
		files[rel] = true
		return nil
	})
	
	return files
}