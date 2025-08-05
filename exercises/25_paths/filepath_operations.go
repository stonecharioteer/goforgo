// filepath_operations.go
// Learn file path manipulation and operations

package main

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	fmt.Println("=== File Path Operations ===")
	
	// TODO: Basic path operations
	samplePaths := []string{
		"/home/user/documents/file.txt",
		"C:\\Users\\User\\Documents\\file.txt",
		"./relative/path/file.go",
		"../parent/directory/script.sh",
		"/usr/local/bin/program",
		"simple-filename.txt",
	}
	
	fmt.Println("Analyzing sample paths:")
	for _, p := range samplePaths {
		fmt.Printf("\nPath: %s\n", p)
		
		// TODO: Extract path components
		dir := /* get directory part */
		base := /* get base filename */
		ext := /* get file extension */
		
		fmt.Printf("  Dir:  %s\n", dir)
		fmt.Printf("  Base: %s\n", base)
		fmt.Printf("  Ext:  %s\n", ext)
		
		// TODO: Check if path is absolute
		isAbs := /* check if path is absolute */
		fmt.Printf("  Absolute: %t\n", isAbs)
	}
	
	fmt.Println("\n=== Path Joining ===")
	
	// TODO: Join path components
	pathComponents := [][]string{
		{"home", "user", "documents", "file.txt"},
		{"usr", "local", "bin", "program"},
		{".", "src", "main.go"},
		{"..", "parent", "child", "grandchild"},
	}
	
	for _, components := range pathComponents {
		// TODO: Join components using filepath.Join
		joined := /* join path components */
		fmt.Printf("Components %v -> %s\n", components, joined)
		
		// TODO: Also show path.Join for comparison
		pathJoined := /* join using path.Join */
		fmt.Printf("  (path.Join: %s)\n", pathJoined)
	}
	
	fmt.Println("\n=== Path Cleaning ===")
	
	// TODO: Clean messy paths
	messyPaths := []string{
		"/home/user/../user/./documents//file.txt",
		"./src/../src/./main.go",
		"/usr/local/bin/../lib/../bin/program",
		"a//b/c/../../d",
	}
	
	for _, messy := range messyPaths {
		// TODO: Clean the path
		cleaned := /* clean the path */
		fmt.Printf("Messy:   %s\n", messy)
		fmt.Printf("Cleaned: %s\n\n", cleaned)
	}
	
	fmt.Println("=== Relative Paths ===")
	
	// TODO: Convert between absolute and relative paths
	basePath := "/home/user/projects"
	targetPaths := []string{
		"/home/user/projects/myapp/main.go",
		"/home/user/documents/readme.txt",
		"/etc/hosts",
		"/home/user/projects/lib/utils.go",
	}
	
	for _, target := range targetPaths {
		// TODO: Get relative path
		rel, err := /* get relative path from basePath to target */
		if err != nil {
			fmt.Printf("Error getting relative path: %v\n", err)
			continue
		}
		
		fmt.Printf("From: %s\n", basePath)
		fmt.Printf("To:   %s\n", target)
		fmt.Printf("Rel:  %s\n\n", rel)
	}
	
	fmt.Println("=== Path Matching ===")
	
	// TODO: Pattern matching with filepath.Match
	patterns := []string{
		"*.go",
		"test_*.txt",
		"dir/*/file.go",
		"**/*.json", // Note: ** not supported by filepath.Match
	}
	
	testFiles := []string{
		"main.go",
		"test_utils.txt",
		"readme.md",
		"dir/subdir/file.go",
		"config.json",
	}
	
	for _, pattern := range patterns {
		fmt.Printf("Pattern: %s\n", pattern)
		for _, file := range testFiles {
			// TODO: Check if file matches pattern
			matched, err := /* check if file matches pattern */
			if err != nil {
				fmt.Printf("  %s: error (%v)\n", file, err)
			} else {
				fmt.Printf("  %s: %t\n", file, matched)
			}
		}
		fmt.Println()
	}
	
	fmt.Println("=== Platform-Specific Operations ===")
	
	// TODO: Demonstrate cross-platform path handling
	fmt.Printf("Current OS: %s\n", runtime.GOOS)
	fmt.Printf("Path separator: %c\n", filepath.Separator)
	fmt.Printf("List separator: %c\n", filepath.ListSeparator)
	
	// TODO: Convert between slash types
	unixPath := "/home/user/file.txt"
	windowsPath := "C:\\Users\\User\\file.txt"
	
	fmt.Printf("\nUnix path: %s\n", unixPath)
	fmt.Printf("To slash: %s\n", /* convert to forward slashes */)
	fmt.Printf("From slash: %s\n", /* convert from forward slashes */)
	
	fmt.Printf("\nWindows path: %s\n", windowsPath)
	fmt.Printf("To slash: %s\n", filepath.ToSlash(windowsPath))
	fmt.Printf("From slash: %s\n", filepath.FromSlash("C:/Users/User/file.txt"))
	
	fmt.Println("\n=== Walking Directory Paths ===")
	
	// TODO: Simulate directory walking (without actual filesystem)
	simulatedPaths := []string{
		"project/",
		"project/main.go",
		"project/utils/",
		"project/utils/helper.go",
		"project/utils/math.go",
		"project/tests/",
		"project/tests/main_test.go",
		"project/docs/",
		"project/docs/readme.md",
	}
	
	fmt.Println("Simulated directory structure:")
	for _, p := range simulatedPaths {
		// TODO: Calculate depth based on separators
		depth := /* count path separators */
		indent := strings.Repeat("  ", depth)
		
		// TODO: Check if it's a directory (ends with /)
		if /* check if path ends with separator */ {
			fmt.Printf("%s📁 %s\n", indent, filepath.Base(p))
		} else {
			fmt.Printf("%s📄 %s\n", indent, filepath.Base(p))
		}
	}
	
	fmt.Println("\n=== Volume and Drive Operations ===")
	
	// TODO: Volume name operations (mainly for Windows)
	volumePaths := []string{
		"C:\\Windows\\System32",
		"/usr/local/bin",
		"D:\\Projects\\myapp",
		"//server/share/file.txt", // UNC path
	}
	
	for _, p := range volumePaths {
		// TODO: Get volume name
		volume := /* get volume name */
		fmt.Printf("Path: %s\n", p)
		fmt.Printf("  Volume: %s\n", volume)
	}
}