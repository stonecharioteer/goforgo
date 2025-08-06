package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset your progress and all exercises to a clean state",
	Long:  `Reset your progress and all exercises to a clean state. This will delete all your work and re-initialize the exercises directory.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := GetWorkingDirectory()
		if err != nil {
			return fmt.Errorf("failed to get working directory: %w", err)
		}

		fmt.Println("🔥 Resetting GoForGo exercises...")

		// Remove old directories and files
		dirsToRemove := []string{"exercises", "solutions"}
		for _, dir := range dirsToRemove {
			if err := os.RemoveAll(filepath.Join(cwd, dir)); err != nil {
				return fmt.Errorf("failed to remove directory %s: %w", dir, err)
			}
		}

		progressFile := filepath.Join(cwd, ".goforgo-progress.toml")
		if _, err := os.Stat(progressFile); err == nil {
			if err := os.Remove(progressFile); err != nil {
				return fmt.Errorf("failed to remove progress file: %w", err)
			}
		}

		// Re-initialize exercises
		_, err = InitializeExercises(cwd)
		if err != nil {
			return err
		}

		fmt.Println("✅ GoForGo has been reset to a clean state.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}