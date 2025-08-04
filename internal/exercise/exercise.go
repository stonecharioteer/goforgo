package exercise

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

// Exercise represents a single GoForGo exercise
type Exercise struct {
	// File paths
	FilePath     string `toml:"-"`         // Path to the .go file
	MetadataPath string `toml:"-"`         // Path to the .toml file
	SolutionPath string `toml:"-"`         // Path to the solution file

	// Metadata from TOML file
	Info        ExerciseInfo        `toml:"exercise"`
	Description ExerciseDescription `toml:"description"`
	Validation  ExerciseValidation  `toml:"validation"`
	Hints       ExerciseHints       `toml:"hints"`

	// Runtime state
	Completed   bool      `toml:"-"`
	LastAttempt time.Time `toml:"-"`
	Attempts    int       `toml:"-"`
}

// ExerciseInfo contains basic exercise information
type ExerciseInfo struct {
	Name          string `toml:"name"`
	Category      string `toml:"category"`
	Difficulty    int    `toml:"difficulty"`    // 1-5 scale
	EstimatedTime string `toml:"estimated_time"` // e.g., "5m", "15m"
	GoVersion     string `toml:"go_version,omitempty"` // Minimum Go version required
}

// ExerciseDescription contains learning content
type ExerciseDescription struct {
	Title              string   `toml:"title"`
	Summary            string   `toml:"summary"`
	LearningObjectives []string `toml:"learning_objectives"`
}

// ExerciseValidation contains validation configuration
type ExerciseValidation struct {
	Mode          string   `toml:"mode"`           // "build", "test", "run"
	Timeout       string   `toml:"timeout"`        // e.g., "30s"
	RequiredFiles []string `toml:"required_files,omitempty"`
}

// ExerciseHints contains progressive hints
type ExerciseHints struct {
	Level1 string `toml:"level_1"`
	Level2 string `toml:"level_2,omitempty"`
	Level3 string `toml:"level_3,omitempty"`
}

// ExerciseManager manages loading and organizing exercises
type ExerciseManager struct {
	ExercisesPath string
	SolutionsPath string
	exercises     []*Exercise
}

// NewExerciseManager creates a new exercise manager
func NewExerciseManager(basePath string) *ExerciseManager {
	return &ExerciseManager{
		ExercisesPath: filepath.Join(basePath, "exercises"),
		SolutionsPath: filepath.Join(basePath, "solutions"),
		exercises:     make([]*Exercise, 0),
	}
}

// LoadExercises discovers and loads all exercises from the exercises directory
func (em *ExerciseManager) LoadExercises() error {
	if _, err := os.Stat(em.ExercisesPath); os.IsNotExist(err) {
		return fmt.Errorf("exercises directory not found at %s. Run 'goforgo init' first", em.ExercisesPath)
	}

	// Walk through the exercises directory
	err := filepath.Walk(em.ExercisesPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Look for .toml metadata files
		if !info.IsDir() && strings.HasSuffix(path, ".toml") {
			exercise, err := em.loadExercise(path)
			if err != nil {
				return fmt.Errorf("failed to load exercise from %s: %w", path, err)
			}
			em.exercises = append(em.exercises, exercise)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to scan exercises directory: %w", err)
	}

	if len(em.exercises) == 0 {
		return fmt.Errorf("no exercises found in %s. Run 'goforgo init' to set up exercises", em.ExercisesPath)
	}

	fmt.Printf("📚 Loaded %d exercises\n", len(em.exercises))
	return nil
}

// loadExercise loads a single exercise from its metadata file
func (em *ExerciseManager) loadExercise(metadataPath string) (*Exercise, error) {
	exercise := &Exercise{
		MetadataPath: metadataPath,
	}

	// Parse TOML metadata
	if _, err := toml.DecodeFile(metadataPath, exercise); err != nil {
		return nil, fmt.Errorf("failed to parse TOML metadata: %w", err)
	}

	// Determine file paths
	dir := filepath.Dir(metadataPath)
	baseName := strings.TrimSuffix(filepath.Base(metadataPath), ".toml")

	exercise.FilePath = filepath.Join(dir, baseName+".go")

	// Check if the Go file exists
	if _, err := os.Stat(exercise.FilePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("Go file not found at %s", exercise.FilePath)
	}

	// Determine solution path
	relDir, err := filepath.Rel(em.ExercisesPath, dir)
	if err != nil {
		return nil, fmt.Errorf("failed to determine relative path: %w", err)
	}
	exercise.SolutionPath = filepath.Join(em.SolutionsPath, relDir, baseName+".go")

	return exercise, nil
}

// GetExercises returns all loaded exercises
func (em *ExerciseManager) GetExercises() []*Exercise {
	return em.exercises
}

// GetExerciseByName finds an exercise by its name
func (em *ExerciseManager) GetExerciseByName(name string) (*Exercise, error) {
	for _, exercise := range em.exercises {
		if exercise.Info.Name == name {
			return exercise, nil
		}
	}
	return nil, fmt.Errorf("exercise '%s' not found", name)
}

// GetNextExercise returns the next incomplete exercise
func (em *ExerciseManager) GetNextExercise() *Exercise {
	for _, exercise := range em.exercises {
		if !exercise.Completed {
			return exercise
		}
	}
	return nil // All exercises completed
}

// String returns a human-readable representation of the exercise
func (e *Exercise) String() string {
	status := "❌"
	if e.Completed {
		status = "✅"
	}
	return fmt.Sprintf("%s %s/%s: %s", status, e.Info.Category, e.Info.Name, e.Description.Title)
}

// GetDifficultyString returns a visual representation of difficulty
func (e *Exercise) GetDifficultyString() string {
	switch e.Info.Difficulty {
	case 1:
		return "⭐ Beginner"
	case 2:
		return "⭐⭐ Easy"
	case 3:
		return "⭐⭐⭐ Medium"
	case 4:
		return "⭐⭐⭐⭐ Hard"
	case 5:
		return "⭐⭐⭐⭐⭐ Expert"
	default:
		return "❓ Unknown"
	}
}

// GetHint returns the appropriate hint based on attempt count
func (e *Exercise) GetHint() string {
	switch {
	case e.Attempts <= 2 && e.Hints.Level1 != "":
		return e.Hints.Level1
	case e.Attempts <= 5 && e.Hints.Level2 != "":
		return e.Hints.Level2
	case e.Hints.Level3 != "":
		return e.Hints.Level3
	case e.Hints.Level1 != "":
		return e.Hints.Level1
	default:
		return "No hints available for this exercise."
	}
}