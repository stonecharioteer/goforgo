package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset [exercise_name]",
	Short: "Reset exercise(s) to their original state",
	Long: `Reset exercises to their original state by copying from the source.

Examples:
  goforgo reset hello          # Reset specific exercise
  goforgo reset               # Reset all exercises
  goforgo reset --hard        # Reset and remove progress tracking

This is useful when you want to start over or fix corrupted exercise files.`,
	RunE: runReset,
}

var resetHard bool

func runReset(cmd *cobra.Command, args []string) error {
	currentDir, err := GetWorkingDirectory()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	exerciseDir := filepath.Join(currentDir, "exercises")
	
	// Check if exercises directory exists
	if _, err := os.Stat(exerciseDir); os.IsNotExist(err) {
		return fmt.Errorf("no exercises directory found in %s. Run 'goforgo init' first", currentDir)
	}

	if len(args) == 0 {
		// Reset all exercises
		fmt.Println("🔄 Resetting all exercises to original state...")
		return resetAllExercises(currentDir)
	} else {
		// Reset specific exercise
		exerciseName := args[0]
		fmt.Printf("🔄 Resetting exercise '%s' to original state...\n", exerciseName)
		return resetSpecificExercise(currentDir, exerciseName)
	}
}

func resetAllExercises(baseDir string) error {
	// Remove existing exercises directory
	exerciseDir := filepath.Join(baseDir, "exercises")
	if err := os.RemoveAll(exerciseDir); err != nil {
		return fmt.Errorf("failed to remove exercises directory: %w", err)
	}

	// Remove solutions directory
	solutionDir := filepath.Join(baseDir, "solutions")
	if err := os.RemoveAll(solutionDir); err != nil {
		return fmt.Errorf("failed to remove solutions directory: %w", err)
	}

	// Reset progress if --hard flag is used
	if resetHard {
		progressFile := filepath.Join(baseDir, ".goforgo-progress.toml")
		if err := os.Remove(progressFile); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to remove progress file: %w", err)
		}
		fmt.Println("🗑️  Removed progress tracking")
	}

	// Recreate from source (same logic as init command)
	if err := createExerciseStructure(baseDir); err != nil {
		return fmt.Errorf("failed to recreate exercise structure: %w", err)
	}

	fmt.Println("✅ All exercises reset successfully!")
	return nil
}

func resetSpecificExercise(baseDir, exerciseName string) error {
	// Find the exercise in source
	sourceExercises := ""
	
	// Get the binary's location to find source files relative to it
	execPath, err := os.Executable()
	if err != nil {
		execPath = "" // Fallback to current directory search
	}
	
	var possiblePaths []string
	if execPath != "" {
		// Binary-relative paths (for installed binary)
		execDir := filepath.Dir(execPath)
		possiblePaths = append(possiblePaths, 
			filepath.Join(execDir, "exercises"),     // Same dir as binary
			filepath.Join(execDir, "..", "exercises"), // Parent of binary dir
			filepath.Join(execDir, "..", "..", "exercises"), // Go up from bin/
		)
	}
	
	// Add current-directory relative paths (for development)
	possiblePaths = append(possiblePaths,
		"exercises",           // Current directory
		"../exercises",        // Parent directory 
		"../../exercises",     // Go up from bin/
	)
	
	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			sourceExercises = path
			break
		}
	}
	
	if sourceExercises == "" {
		return fmt.Errorf("could not find source exercises to reset from")
	}

	// Find the specific exercise files
	var exerciseFound bool
	err = filepath.Walk(sourceExercises, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if info.IsDir() {
			return nil
		}
		
		// Check if this file belongs to the exercise we're looking for
		filename := info.Name()
		if strings.HasPrefix(filename, exerciseName+".") {
			exerciseFound = true
			
			// Calculate destination path
			relPath, err := filepath.Rel(sourceExercises, srcPath)
			if err != nil {
				return err
			}
			destPath := filepath.Join(baseDir, "exercises", relPath)
			
			// Copy file
			if err := copyFile(srcPath, destPath); err != nil {
				return fmt.Errorf("failed to copy %s: %w", srcPath, err)
			}
			fmt.Printf("📝 Reset %s\n", relPath)
		}
		
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to reset exercise: %w", err)
	}

	if !exerciseFound {
		return fmt.Errorf("exercise '%s' not found", exerciseName)
	}

	fmt.Printf("✅ Exercise '%s' reset successfully!\n", exerciseName)
	return nil
}

func init() {
	resetCmd.Flags().BoolVar(&resetHard, "hard", false, "Also remove progress tracking")
	rootCmd.AddCommand(resetCmd)
}
