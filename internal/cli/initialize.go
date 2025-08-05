package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	
	"github.com/stonecharioteer/goforgo/internal/exercise"
)

func InitializeExercises(baseDir string) (int, error) {
	if err := createExerciseStructure(baseDir); err != nil {
		return 0, fmt.Errorf("failed to create exercise structure: %w", err)
	}
	
	// Count the exercises that were copied using the same logic as ExerciseManager
	exerciseCount, err := exercise.CountExercisesInDirectory(filepath.Join(baseDir, "exercises"))
	if err != nil {
		return 0, fmt.Errorf("failed to count exercises: %w", err)
	}
	
	return exerciseCount, nil
}

func createExerciseStructure(baseDir string) error {
	dirs := []string{
		"exercises",
		"solutions",
		"exercises/01_basics",
		"exercises/02_variables",
		"exercises/03_functions",
		"exercises/04_control_flow",
		"exercises/05_data_structures",
		"solutions/01_basics",
		"solutions/02_variables",
		"solutions/03_functions",
		"solutions/04_control_flow",
		"solutions/05_data_structures",
	}

	for _, dir := range dirs {
		fullPath := filepath.Join(baseDir, dir)
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", fullPath, err)
		}
	}

	configPath := filepath.Join(baseDir, ".goforgo.toml")
	configContent := `# GoForGo Configuration
version = "1.0"

[user]
auto_advance = true
show_hints = true
theme = "default"

[progress]
current_exercise = ""
completed_exercises = []
`

	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}

	return copyExerciseFiles(baseDir)
}

func copyExerciseFiles(baseDir string) error {
	sourceExercises := ""
	sourceSolutions := ""

	execPath, err := os.Executable()
	if err != nil {
		execPath = ""
	}

	var possiblePaths []string
	if execPath != "" {
		execDir := filepath.Dir(execPath)
		possiblePaths = append(possiblePaths,
			filepath.Join(execDir, "exercises"),
			filepath.Join(execDir, "..", "exercises"),
			filepath.Join(execDir, "..", "..", "exercises"),
		)
	}

	possiblePaths = append(possiblePaths,
		"exercises",
		"../exercises",
		"../../exercises",
	)

	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			sourceExercises = path
			sourceSolutions = strings.Replace(path, "exercises", "solutions", 1)
			break
		}
	}

	if sourceExercises == "" {
		fmt.Println("⚠️  No source exercises found, creating basic hello exercise")
		return createPlaceholderExercise(filepath.Join(baseDir, "exercises", "01_basics"))
	}

	fmt.Printf("📂 Copying exercises from %s\n", sourceExercises)

	err = filepath.Walk(sourceExercises, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(sourceExercises, srcPath)
		if err != nil {
			return err
		}
		destPath := filepath.Join(baseDir, "exercises", relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}

		return copyFile(srcPath, destPath)
	})

	if err != nil {
		return fmt.Errorf("failed to copy exercises: %w", err)
	}

	if _, err := os.Stat(sourceSolutions); err == nil {
		fmt.Printf("📂 Copying solutions from %s\n", sourceSolutions)
		return filepath.Walk(sourceSolutions, func(srcPath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			relPath, err := filepath.Rel(sourceSolutions, srcPath)
			if err != nil {
				return err
			}
			destPath := filepath.Join(baseDir, "solutions", relPath)

			if info.IsDir() {
				return os.MkdirAll(destPath, 0755)
			}

			return copyFile(srcPath, destPath)
		})
	}

	return nil
}

func copyFile(src, dst string) error {
	content, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, content, 0644)
}

func createPlaceholderExercise(dir string) error {
	exerciseContent := `package main

import "fmt"

// TODO: Fix this program to print "Hello, GoForGo!"
func main() {
	fmt.Println("Hello, World!")
}
`

	exercisePath := filepath.Join(dir, "hello.go")
	if err := os.WriteFile(exercisePath, []byte(exerciseContent), 0644); err != nil {
		return fmt.Errorf("failed to create placeholder exercise: %w", err)
	}

	metadataContent := `[exercise]
name = "hello"
category = "01_basics"
difficulty = 1
estimated_time = "2m"

[description]
title = "Hello GoForGo"
summary = "Your first Go program with GoForGo"
learning_objectives = [
  "Understand basic Go syntax",
  "Learn about the main function",
  "Practice string literals"
]

[validation]
mode = "run"
expected_output = "Hello, GoForGo!"
timeout = "10s"

[hints]
level_1 = "Look at what the TODO comment is asking you to print"
level_2 = "You need to change the string inside fmt.Println()"
level_3 = "Replace 'Hello, World!' with 'Hello, GoForGo!'"
`

	metadataPath := filepath.Join(dir, "hello.toml")
	return os.WriteFile(metadataPath, []byte(metadataContent), 0644)
}
