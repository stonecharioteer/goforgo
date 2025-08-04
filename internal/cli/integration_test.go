package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWorkingDirectoryFlag(t *testing.T) {
	// Test default behavior (no flag)
	workingDirectory = "" // Reset global variable
	cwd, err := GetWorkingDirectory()
	if err != nil {
		t.Fatalf("GetWorkingDirectory failed: %v", err)
	}
	
	expectedCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd failed: %v", err)
	}
	
	if cwd != expectedCwd {
		t.Errorf("Expected current directory %s, got %s", expectedCwd, cwd)
	}
	
	// Test with custom directory
	testDir := "/tmp/test-goforgo"
	workingDirectory = testDir
	
	customDir, err := GetWorkingDirectory()
	if err != nil {
		t.Fatalf("GetWorkingDirectory with custom dir failed: %v", err)
	}
	
	if customDir != testDir {
		t.Errorf("Expected custom directory %s, got %s", testDir, customDir)
	}
	
	// Clean up
	workingDirectory = ""
}

func TestInitCommandWithDirectory(t *testing.T) {
	tempDir := t.TempDir()
	workingDirectory = tempDir
	
	// Run init command
	err := runInit(nil, []string{})
	if err != nil {
		t.Fatalf("runInit failed: %v", err)
	}
	
	// Check if exercises directory was created
	exerciseDir := filepath.Join(tempDir, "exercises")
	if _, err := os.Stat(exerciseDir); os.IsNotExist(err) {
		t.Error("Exercises directory was not created")
	}
	
	// Check if hello exercise was created
	helloFile := filepath.Join(exerciseDir, "01_basics", "hello.go")
	helloToml := filepath.Join(exerciseDir, "01_basics", "hello.toml")
	
	if _, err := os.Stat(helloFile); os.IsNotExist(err) {
		t.Error("Hello exercise Go file was not created")
	}
	
	if _, err := os.Stat(helloToml); os.IsNotExist(err) {
		t.Error("Hello exercise TOML file was not created")
	}
	
	// Clean up
	workingDirectory = ""
}