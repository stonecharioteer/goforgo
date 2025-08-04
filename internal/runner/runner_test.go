package runner

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stonecharioteer/goforgo/internal/exercise"
)

func TestRunner_ValidateExercise_ExpectedOutput(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()
	
	// Create a test exercise
	ex := &exercise.Exercise{
		FilePath: filepath.Join(tempDir, "hello.go"),
		Info: exercise.ExerciseInfo{
			Name:     "hello",
			Category: "test",
		},
		Validation: exercise.ExerciseValidation{
			Mode:           "run",
			ExpectedOutput: "Hello, GoForGo!",
			Timeout:        "10s",
		},
	}

	// Test case 1: Correct output
	t.Run("CorrectOutput", func(t *testing.T) {
		// Create Go file with correct output
		goContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, GoForGo!")
}
`
		err := os.WriteFile(ex.FilePath, []byte(goContent), 0644)
		if err != nil {
			t.Fatalf("Failed to write test file: %v", err)
		}

		runner := NewRunner(tempDir)
		success, feedback, err := runner.ValidateExercise(ex)

		if err != nil {
			t.Fatalf("ValidateExercise failed: %v", err)
		}
		if !success {
			t.Errorf("Expected success=true, got success=%v, feedback=%s", success, feedback)
		}
	})

	// Test case 2: Incorrect output
	t.Run("IncorrectOutput", func(t *testing.T) {
		// Create Go file with incorrect output
		goContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
`
		err := os.WriteFile(ex.FilePath, []byte(goContent), 0644)
		if err != nil {
			t.Fatalf("Failed to write test file: %v", err)
		}

		runner := NewRunner(tempDir)
		success, feedback, err := runner.ValidateExercise(ex)

		if err != nil {
			t.Fatalf("ValidateExercise failed: %v", err)
		}
		if success {
			t.Errorf("Expected success=false, got success=%v, feedback=%s", success, feedback)
		}
		if feedback == "" {
			t.Error("Expected non-empty feedback for failed validation")
		}
	})

	// Test case 3: Build failure
	t.Run("BuildFailure", func(t *testing.T) {
		// Create Go file with syntax error
		goContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!" // Missing closing parenthesis
}
`
		err := os.WriteFile(ex.FilePath, []byte(goContent), 0644)
		if err != nil {
			t.Fatalf("Failed to write test file: %v", err)
		}

		runner := NewRunner(tempDir)
		success, feedback, err := runner.ValidateExercise(ex)

		if err != nil {
			t.Fatalf("ValidateExercise failed: %v", err)
		}
		if success {
			t.Errorf("Expected success=false for build failure, got success=%v", success)
		}
		if feedback == "" {
			t.Error("Expected non-empty feedback for build failure")
		}
	})
}

func TestRunner_ValidateExercise_BuildMode(t *testing.T) {
	tempDir := t.TempDir()
	
	ex := &exercise.Exercise{
		FilePath: filepath.Join(tempDir, "test.go"),
		Info: exercise.ExerciseInfo{
			Name:     "test",
			Category: "test",
		},
		Validation: exercise.ExerciseValidation{
			Mode:    "build",
			Timeout: "10s",
		},
	}

	// Test successful build
	goContent := `package main

import "fmt"

func main() {
	fmt.Println("This should compile")
}
`
	err := os.WriteFile(ex.FilePath, []byte(goContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	runner := NewRunner(tempDir)
	success, _, err := runner.ValidateExercise(ex)

	if err != nil {
		t.Fatalf("ValidateExercise failed: %v", err)
	}
	if !success {
		t.Errorf("Expected success=true for valid build, got success=%v", success)
	}
}

func TestRunner_Timeout(t *testing.T) {
	tempDir := t.TempDir()
	
	ex := &exercise.Exercise{
		FilePath: filepath.Join(tempDir, "timeout.go"),
		Info: exercise.ExerciseInfo{
			Name:     "timeout",
			Category: "test",
		},
		Validation: exercise.ExerciseValidation{
			Mode:    "run",
			Timeout: "100ms", // Very short timeout
		},
	}

	// Create Go file that takes longer than timeout
	goContent := `package main

import "time"

func main() {
	time.Sleep(1 * time.Second) // Longer than 100ms timeout
	println("Done")
}
`
	err := os.WriteFile(ex.FilePath, []byte(goContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	runner := NewRunner(tempDir)
	start := time.Now()
	success, feedback, err := runner.ValidateExercise(ex)
	duration := time.Since(start)

	if err != nil {
		t.Fatalf("ValidateExercise failed: %v", err)
	}
	if success {
		t.Errorf("Expected success=false for timeout, got success=%v", success)
	}
	if duration > 2*time.Second {
		t.Errorf("Expected timeout to occur quickly, took %v", duration)
	}
	if feedback == "" {
		t.Error("Expected feedback for timeout")
	}
}