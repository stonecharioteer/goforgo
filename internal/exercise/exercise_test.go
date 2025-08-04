package exercise

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestExerciseManager_ProgressTracking(t *testing.T) {
	tempDir := t.TempDir()
	
	// Create test exercise structure
	exerciseDir := filepath.Join(tempDir, "exercises", "01_basics")
	err := os.MkdirAll(exerciseDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create exercise directory: %v", err)
	}

	// Create a test exercise
	exerciseContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
`
	exerciseFile := filepath.Join(exerciseDir, "hello.go")
	err = os.WriteFile(exerciseFile, []byte(exerciseContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write exercise file: %v", err)
	}

	tomlContent := `[exercise]
name = "hello"
category = "01_basics"
difficulty = 1

[description]
title = "Hello Test"
summary = "Test exercise"

[validation]
mode = "run"
timeout = "10s"

[hints]
level_1 = "Test hint"
`
	tomlFile := filepath.Join(exerciseDir, "hello.toml")
	err = os.WriteFile(tomlFile, []byte(tomlContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write TOML file: %v", err)
	}

	// Test ExerciseManager
	em := NewExerciseManager(tempDir)
	err = em.LoadExercises()
	if err != nil {
		t.Fatalf("Failed to load exercises: %v", err)
	}

	if len(em.exercises) != 1 {
		t.Fatalf("Expected 1 exercise, got %d", len(em.exercises))
	}

	exercise := em.exercises[0]
	if exercise.Info.Name != "hello" {
		t.Errorf("Expected exercise name 'hello', got '%s'", exercise.Info.Name)
	}
	if exercise.Completed {
		t.Error("Exercise should not be completed initially")
	}

	// Test marking exercise as completed
	err = em.MarkExerciseCompleted("hello")
	if err != nil {
		t.Fatalf("Failed to mark exercise as completed: %v", err)
	}

	if !exercise.Completed {
		t.Error("Exercise should be marked as completed")
	}

	// Test progress persistence
	progressFile := filepath.Join(tempDir, ".goforgo-progress.toml")
	if _, err := os.Stat(progressFile); os.IsNotExist(err) {
		t.Error("Progress file should have been created")
	}

	// Test loading progress in new manager
	em2 := NewExerciseManager(tempDir)
	err = em2.LoadExercises()
	if err != nil {
		t.Fatalf("Failed to load exercises in new manager: %v", err)
	}

	if !em2.exercises[0].Completed {
		t.Error("Exercise completion should persist across manager instances")
	}
}

func TestExerciseManager_GetNextExercise(t *testing.T) {
	tempDir := t.TempDir()
	
	// Create multiple test exercises
	exercises := []struct {
		name     string
		category string
	}{
		{"hello", "01_basics"},
		{"variables", "02_variables"},
		{"functions", "03_functions"},
	}

	for _, ex := range exercises {
		exerciseDir := filepath.Join(tempDir, "exercises", ex.category)
		err := os.MkdirAll(exerciseDir, 0755)
		if err != nil {
			t.Fatalf("Failed to create directory for %s: %v", ex.name, err)
		}

		// Create Go file
		goFile := filepath.Join(exerciseDir, ex.name+".go")
		err = os.WriteFile(goFile, []byte("package main\n\nfunc main() {}\n"), 0644)
		if err != nil {
			t.Fatalf("Failed to write Go file for %s: %v", ex.name, err)
		}

		// Create TOML file
		tomlContent := `[exercise]
name = "` + ex.name + `"
category = "` + ex.category + `"
difficulty = 1

[description]
title = "Test ` + ex.name + `"
summary = "Test exercise"

[validation]
mode = "build"
timeout = "10s"

[hints]
level_1 = "Test hint"
`
		tomlFile := filepath.Join(exerciseDir, ex.name+".toml")
		err = os.WriteFile(tomlFile, []byte(tomlContent), 0644)
		if err != nil {
			t.Fatalf("Failed to write TOML file for %s: %v", ex.name, err)
		}
	}

	em := NewExerciseManager(tempDir)
	err := em.LoadExercises()
	if err != nil {
		t.Fatalf("Failed to load exercises: %v", err)
	}

	if len(em.exercises) != 3 {
		t.Fatalf("Expected 3 exercises, got %d", len(em.exercises))
	}

	// Test getting first exercise
	nextEx := em.GetNextExercise()
	if nextEx == nil {
		t.Fatal("Expected first exercise, got nil")
	}

	// Mark first exercise as completed
	err = em.MarkExerciseCompleted(nextEx.Info.Name)
	if err != nil {
		t.Fatalf("Failed to mark exercise as completed: %v", err)
	}

	// Test getting second exercise
	nextEx2 := em.GetNextExercise()
	if nextEx2 == nil {
		t.Fatal("Expected second exercise, got nil")
	}
	if nextEx2.Info.Name == nextEx.Info.Name {
		t.Error("Next exercise should be different from completed one")
	}
}

func TestExercise_GetHint(t *testing.T) {
	ex := &Exercise{
		Hints: ExerciseHints{
			Level1: "First hint",
			Level2: "Second hint",
			Level3: "Third hint",
		},
	}

	// Test hint progression based on attempts
	testCases := []struct {
		attempts     int
		expectedHint string
	}{
		{1, "First hint"},
		{2, "First hint"},
		{3, "Second hint"},
		{5, "Second hint"},
		{6, "Third hint"},
		{10, "Third hint"},
	}

	for _, tc := range testCases {
		ex.Attempts = tc.attempts
		hint := ex.GetHint()
		if hint != tc.expectedHint {
			t.Errorf("For %d attempts, expected hint '%s', got '%s'", 
				tc.attempts, tc.expectedHint, hint)
		}
	}
}

func TestExercise_GetDifficultyString(t *testing.T) {
	testCases := []struct {
		difficulty int
		expected   string
	}{
		{1, "⭐ Beginner"},
		{2, "⭐⭐ Easy"},
		{3, "⭐⭐⭐ Medium"},
		{4, "⭐⭐⭐⭐ Hard"},
		{5, "⭐⭐⭐⭐⭐ Expert"},
		{0, "❓ Unknown"},
		{6, "❓ Unknown"},
	}

	for _, tc := range testCases {
		ex := &Exercise{
			Info: ExerciseInfo{
				Difficulty: tc.difficulty,
			},
		}
		result := ex.GetDifficultyString()
		if result != tc.expected {
			t.Errorf("For difficulty %d, expected '%s', got '%s'", 
				tc.difficulty, tc.expected, result)
		}
	}
}

func TestProgress_Persistence(t *testing.T) {
	tempDir := t.TempDir()
	
	em := NewExerciseManager(tempDir)
	
	// Test initial progress
	if len(em.progress.CompletedExercises) != 0 {
		t.Error("Initial progress should be empty")
	}

	// Manually add completed exercise
	em.progress.CompletedExercises["test"] = true
	em.progress.CurrentExercise = "next"
	em.progress.LastUpdated = time.Now()

	// Save progress
	err := em.saveProgress()
	if err != nil {
		t.Fatalf("Failed to save progress: %v", err)
	}

	// Create new manager and load progress
	em2 := NewExerciseManager(tempDir)
	if len(em2.progress.CompletedExercises) != 1 {
		t.Errorf("Expected 1 completed exercise, got %d", len(em2.progress.CompletedExercises))
	}
	if !em2.progress.CompletedExercises["test"] {
		t.Error("Test exercise should be marked as completed")
	}
	if em2.progress.CurrentExercise != "next" {
		t.Errorf("Expected current exercise 'next', got '%s'", em2.progress.CurrentExercise)
	}
}