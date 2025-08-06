// path_manipulation.go - SOLUTION
// Learn advanced path manipulation: URL paths, cleaning, joining, and path utilities

package main

import (
	"fmt"
	"net/url"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("=== Path Manipulation ===")
	
	fmt.Println("\n=== Filepath Operations ===")
	
	// Test various filepath operations
	testPaths := []string{
		"/home/user/documents/file.txt",
		"../relative/../path/to/file.go",
		"C:\\Windows\\System32\\file.exe",
		"./current/./directory/file.md",
		"/complex//path///with////extra/slashes/file.json",
		"",
		".",
		"..",
		"/",
	}
	
	fmt.Println("Filepath analysis:")
	for _, p := range testPaths {
		originalPath := p
		if p == "" {
			p = "<empty>"
		}
		
		fmt.Printf("\nPath: %s\n", p)
		
		if originalPath != "" {
			// Get directory and filename
			dir := filepath.Dir(originalPath)
			base := filepath.Base(originalPath)
			ext := filepath.Ext(originalPath)
			
			fmt.Printf("  Dir: %s\n", dir)
			fmt.Printf("  Base: %s\n", base)
			fmt.Printf("  Ext: %s\n", ext)
			
			// Clean the path
			cleaned := filepath.Clean(originalPath)
			fmt.Printf("  Cleaned: %s\n", cleaned)
			
			// Check if path is absolute
			isAbs := filepath.IsAbs(originalPath)
			fmt.Printf("  Is Absolute: %t\n", isAbs)
		}
	}
	
	fmt.Println("\n=== Path Joining ===")
	
	// Test path joining
	pathParts := [][]string{
		{"home", "user", "documents"},
		{"", "absolute", "path"},
		{"relative", "..", "path"},
		{"path", "with", "", "empty", "parts"},
		{"C:", "Windows", "System32"},
		{"/", "root", "path"},
	}
	
	fmt.Println("Path joining examples:")
	for i, parts := range pathParts {
		// Join path parts
		joined := filepath.Join(parts...)
		fmt.Printf("  %d. %v → %s\n", i+1, parts, joined)
	}
	
	fmt.Println("\n=== Relative Path Calculations ===")
	
	// Calculate relative paths
	basePaths := []string{
		"/home/user/documents",
		"/var/www/html",
		"C:\\Program Files",
	}
	
	targetPaths := []string{
		"/home/user/downloads/file.txt",
		"/var/log/apache/access.log",
		"C:\\Users\\Public\\Documents\\file.doc",
	}
	
	fmt.Println("Relative path calculations:")
	for i, base := range basePaths {
		if i < len(targetPaths) {
			target := targetPaths[i]
			
			// Calculate relative path from base to target
			rel, err := filepath.Rel(base, target)
			if err != nil {
				fmt.Printf("  Error: %v\n", err)
			} else {
				fmt.Printf("  From: %s\n", base)
				fmt.Printf("  To: %s\n", target)
				fmt.Printf("  Relative: %s\n", rel)
				fmt.Println()
			}
		}
	}
	
	fmt.Println("\n=== URL Path Manipulation ===")
	
	// Work with URL paths
	testURLs := []string{
		"https://example.com/api/v1/users",
		"http://localhost:8080/path/to/resource?param=value",
		"ftp://files.example.com/downloads/../uploads/file.txt",
		"https://api.github.com/repos/golang/go/issues",
	}
	
	fmt.Println("URL path manipulation:")
	for _, urlStr := range testURLs {
		// Parse URL
		u, err := url.Parse(urlStr)
		if err != nil {
			fmt.Printf("Error parsing URL %s: %v\n", urlStr, err)
			continue
		}
		
		fmt.Printf("\nOriginal URL: %s\n", urlStr)
		fmt.Printf("  Scheme: %s\n", u.Scheme)
		fmt.Printf("  Host: %s\n", u.Host)
		fmt.Printf("  Path: %s\n", u.Path)
		
		// Clean and manipulate URL path
		cleanPath := path.Clean(u.Path)
		u.Path = cleanPath
		fmt.Printf("  Cleaned Path: %s\n", cleanPath)
		
		// Join additional path segments
		additionalPath := "extra/segment"
		newPath := path.Join(u.Path, additionalPath)
		u.Path = newPath
		fmt.Printf("  Extended Path: %s\n", newPath)
		fmt.Printf("  Final URL: %s\n", u.String())
	}
	
	fmt.Println("\n=== Path Pattern Matching ===")
	
	// Test path pattern matching
	patterns := []string{
		"*.txt",
		"test_*.go",
		// Note: ** is not supported by filepath.Match, using simpler patterns
		"*.json",
		"main.go",
		"[abc]*.log",
	}
	
	testFiles := []string{
		"document.txt",
		"test_helper.go",
		"config.json",
		"main.go",
		"a_file.log",
		"data.xml",
		"test_runner.go",
		"settings.json",
		"b_debug.log",
	}
	
	fmt.Println("Pattern matching:")
	for _, pattern := range patterns {
		fmt.Printf("\nPattern: %s\n", pattern)
		fmt.Printf("  Matches:\n")
		
		for _, file := range testFiles {
			// Check if file matches pattern
			matched, err := filepath.Match(pattern, file)
			if err != nil {
				fmt.Printf("    Error matching %s: %v\n", file, err)
			} else if matched {
				fmt.Printf("    ✅ %s\n", file)
			}
		}
	}
	
	fmt.Println("\n=== Path Utilities ===")
	
	// Custom path utilities
	fmt.Println("Custom path utilities:")
	
	// Extract filename without extension
	testFilenames := []string{
		"document.pdf",
		"archive.tar.gz",
		"script.sh",
		"no_extension",
		".hidden",
		"multiple.dots.in.name.txt",
	}
	
	fmt.Println("\nFilename without extension:")
	for _, filename := range testFilenames {
		// Remove extension from filename
		ext := filepath.Ext(filename)
		withoutExt := filename
		if ext != "" {
			withoutExt = filename[:len(filename)-len(ext)]
		}
		fmt.Printf("  %s → %s\n", filename, withoutExt)
	}
	
	// Path validation
	fmt.Println("\nPath validation:")
	potentiallyDangerousPaths := []string{
		"../../../etc/passwd",
		"/safe/path/file.txt",
		"..\\..\\windows\\system32\\file.exe",
		"./safe/relative/path",
		"/tmp/../usr/bin/evil",
		"normal/path/file.doc",
	}
	
	basePath := "/home/user/webapp"
	for _, p := range potentiallyDangerousPaths {
		// Check if path is safe (doesn't escape base directory)
		cleanedPath := filepath.Clean(p)
		absPath := filepath.Join(basePath, cleanedPath)
		relPath, err := filepath.Rel(basePath, absPath)
		
		isSafe := err == nil && !strings.HasPrefix(relPath, "..")
		status := "✅ Safe"
		if !isSafe {
			status = "❌ Dangerous"
		}
		fmt.Printf("  %s: %s\n", p, status)
	}
	
	// Path normalization
	fmt.Println("\nPath normalization:")
	unnormalizedPaths := []string{
		"./path/../to/./file",
		"/home//user///documents////file.txt",
		"C:\\Windows\\..\\Windows\\System32",
		"relative/./path/../to/file",
	}
	
	for _, p := range unnormalizedPaths {
		// Normalize path
		normalized := filepath.Clean(p)
		fmt.Printf("  %s → %s\n", p, normalized)
	}
	
	fmt.Println("\n=== Path Building Utilities ===")
	
	// Build paths programmatically
	fmt.Println("Building paths programmatically:")
	
	// Create a path builder
	type PathBuilder struct {
		segments []string
	}
	
	// PathBuilder methods
	func NewPathBuilder() *PathBuilder {
		return &PathBuilder{segments: make([]string, 0)}
	}
	
	func (pb *PathBuilder) Root(root string) *PathBuilder {
		pb.segments = []string{root}
		return pb
	}
	
	func (pb *PathBuilder) Add(segment string) *PathBuilder {
		pb.segments = append(pb.segments, segment)
		return pb
	}
	
	func (pb *PathBuilder) AddMany(segments ...string) *PathBuilder {
		pb.segments = append(pb.segments, segments...)
		return pb
	}
	
	func (pb *PathBuilder) Build() string {
		return filepath.Join(pb.segments...)
	}
	
	// Demonstrate path builder usage
	pathBuilder := NewPathBuilder()
	complexPath := pathBuilder.
		Root("/home").
		Add("user").
		AddMany("projects", "myapp", "src").
		Add("main.go").
		Build()
	
	fmt.Printf("Built path: %s\n", complexPath)
	
	fmt.Println("\n=== Best Practices ===")
	fmt.Println("✅ Always use filepath.Join() for cross-platform compatibility")
	fmt.Println("✅ Clean paths with filepath.Clean() to remove redundant elements")
	fmt.Println("✅ Validate paths to prevent directory traversal attacks")
	fmt.Println("✅ Use filepath.Abs() to convert relative paths to absolute")
	fmt.Println("✅ Handle both forward and backward slashes appropriately")
	fmt.Println("✅ Use path.Join() for URL paths, filepath.Join() for file system paths")
	fmt.Println("✅ Always check for errors when working with paths")
	fmt.Println("✅ Consider case sensitivity based on the target file system")
}