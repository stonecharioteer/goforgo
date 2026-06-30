package cli

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	goforgo "github.com/stonecharioteer/goforgo"
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
	fmt.Println("📂 Extracting embedded exercises...")

	if err := copyEmbeddedDir("exercises", filepath.Join(baseDir, "exercises")); err != nil {
		return fmt.Errorf("failed to extract exercises: %w", err)
	}

	if err := copyEmbeddedDir("solutions", filepath.Join(baseDir, "solutions")); err != nil {
		return fmt.Errorf("failed to extract solutions: %w", err)
	}

	return nil
}

func copyEmbeddedDir(embedDir, destDir string) error {
	return fs.WalkDir(goforgo.Content, embedDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(embedDir, path)
		if err != nil {
			return err
		}
		destPath := filepath.Join(destDir, relPath)

		if d.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}

		content, err := goforgo.Content.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(destPath, content, 0644)
	})
}
