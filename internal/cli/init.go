package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize GoForGo exercises in the current directory",
	Long: `Initialize GoForGo exercises in the current directory.

This command will create the exercises directory structure and copy all
exercise files to your local directory. You can then start learning Go
by running 'goforgo' to enter watch mode.

The exercises directory will contain:
  - exercises/: All exercise source files organized by topic
  - solutions/: Complete solutions (for reference only)
  - .goforgo.toml: Progress tracking and configuration`,
	RunE: runInit,
}

func runInit(cmd *cobra.Command, args []string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	exerciseDir := filepath.Join(currentDir, "exercises")
	
	// Check if exercises directory already exists
	if _, err := os.Stat(exerciseDir); err == nil {
		return fmt.Errorf("exercises directory already exists in %s\nIf you want to reinitialize, please remove the directory first", currentDir)
	}

	fmt.Printf("🚀 Initializing GoForGo exercises in %s\n", currentDir)

	// Create exercises directory structure
	if err := createExerciseStructure(currentDir); err != nil {
		return fmt.Errorf("failed to create exercise structure: %w", err)
	}

	fmt.Printf(`✅ GoForGo initialized successfully!

📂 Created directories:
  - exercises/     (250+ Go exercises organized by topic)
  - solutions/     (complete solutions for reference)

🎯 Next steps:
  1. Run 'goforgo' to start the interactive tutorial
  2. Edit exercises in your favorite editor
  3. Watch as GoForGo automatically compiles and tests your code

📖 Need help? Run 'goforgo --help' or visit the documentation.

Happy learning! 🎉
`)

	return nil
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

	// Create initial configuration file
	configPath := filepath.Join(baseDir, ".goforgo.toml")
	configContent := `# GoForGo Configuration
version = "1.0"

[user]
# Your learning preferences
auto_advance = true
show_hints = true
theme = "default"

[progress]
# This section is automatically managed
current_exercise = ""
completed_exercises = []
`

	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}

	// TODO: Copy actual exercise files here
	// For now, create a placeholder exercise
	return createPlaceholderExercise(filepath.Join(baseDir, "exercises", "01_basics"))
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

	// Create corresponding TOML metadata
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
timeout = "10s"

[hints]
level_1 = "Look at what the TODO comment is asking you to print"
level_2 = "You need to change the string inside fmt.Println()"
level_3 = "Replace 'Hello, World!' with 'Hello, GoForGo!'"
`

	metadataPath := filepath.Join(dir, "hello.toml")
	return os.WriteFile(metadataPath, []byte(metadataContent), 0644)
}

func init() {
	rootCmd.AddCommand(initCmd)
}