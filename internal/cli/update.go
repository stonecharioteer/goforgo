package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stonecharioteer/goforgo/internal/exercise"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update exercise source files without overwriting completed work",
	Long: `Update exercise source files (like .toml configuration files) from the binary
without overwriting Go files for exercises you've already completed. This allows you
to get updates to exercise metadata while preserving your solutions.

This command will:
- Update all .toml configuration files
- Update .go files only for incomplete exercises
- Preserve your completed exercise solutions
- Keep your progress intact`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := GetWorkingDirectory()
		if err != nil {
			return fmt.Errorf("failed to get working directory: %w", err)
		}

		// Check if this looks like a goforgo directory
		if _, err := os.Stat(filepath.Join(cwd, ".goforgo.toml")); err != nil {
			return fmt.Errorf("this doesn't appear to be a GoForGo directory (no .goforgo.toml found)")
		}

		fmt.Println("🔄 Updating GoForGo exercise files...")

		// Load exercise manager to get completion status
		em := exercise.NewExerciseManager(cwd)
		if err := em.LoadExercises(); err != nil {
			return fmt.Errorf("failed to load exercises: %w", err)
		}

		// Get completed exercises
		completedExercises := em.GetCompletedExercises()
		fmt.Printf("📊 Found %d completed exercises to preserve\n", len(completedExercises))

		// Update exercise files selectively
		if err := updateExerciseFiles(cwd, completedExercises); err != nil {
			return fmt.Errorf("failed to update exercise files: %w", err)
		}

		fmt.Println("✅ GoForGo exercises updated successfully!")
		fmt.Println("💡 Your completed exercise solutions have been preserved.")
		return nil
	},
}

func updateExerciseFiles(baseDir string, completedExercises map[string]bool) error {
	sourceExercises := ""
	sourceSolutions := ""

	// Find source directories (same logic as initialize.go)
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
		return fmt.Errorf("no source exercises found in binary location")
	}

	fmt.Printf("📂 Updating from source: %s\n", sourceExercises)

	// Walk through source exercises and update selectively
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
			// Ensure directory exists
			return os.MkdirAll(destPath, 0755)
		}

		// Determine if we should update this file
		shouldUpdate := shouldUpdateFile(srcPath, destPath, relPath, completedExercises)
		
		if shouldUpdate {
			fmt.Printf("  📝 Updating: %s\n", relPath)
			return copyFile(srcPath, destPath)
		} else {
			fmt.Printf("  ⏭️  Preserving: %s (exercise completed)\n", relPath)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to update exercises: %w", err)
	}

	// Update solutions directory too (these are reference solutions)
	if _, err := os.Stat(sourceSolutions); err == nil {
		fmt.Printf("📂 Updating solutions from: %s\n", sourceSolutions)
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

			fmt.Printf("  📝 Updating solution: %s\n", relPath)
			return copyFile(srcPath, destPath)
		})
	}

	return nil
}

func shouldUpdateFile(srcPath, destPath, relPath string, completedExercises map[string]bool) bool {
	// Always update .toml files (configuration/metadata)
	if strings.HasSuffix(relPath, ".toml") {
		return true
	}

	// Always update non-.go files (README, scripts, etc.)
	if !strings.HasSuffix(relPath, ".go") {
		return true
	}

	// For .go files, extract exercise name from path and check completion
	exerciseName := extractExerciseName(relPath)
	if exerciseName == "" {
		// Can't determine exercise name, update it to be safe
		return true
	}

	// If the exercise is completed, don't update the .go file (preserve user's work)
	if completedExercises[exerciseName] {
		return false
	}

	// Exercise is not completed, safe to update
	return true
}

func extractExerciseName(relPath string) string {
	// Extract exercise name from path like "04_control_flow/range_loops.go"
	// Returns "range_loops"
	if strings.HasSuffix(relPath, ".go") {
		base := filepath.Base(relPath)
		return strings.TrimSuffix(base, ".go")
	}
	return ""
}


func init() {
	rootCmd.AddCommand(updateCmd)
}