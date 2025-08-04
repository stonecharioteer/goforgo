package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/stonecharioteer/goforgo/internal/tui"
)

var (
	manualRun bool
)

// watchCmd represents the watch command (also the default when no subcommand is given)
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Start interactive watch mode (default)",
	Long: `Start interactive watch mode with file watching and real-time feedback.

This is the default mode when you run 'goforgo' without any subcommands.
It will:
  1. Watch for changes in exercise files
  2. Automatically compile and test your code
  3. Provide real-time feedback and hints
  4. Guide you through exercises progressively

Use keyboard shortcuts:
  r - manually run current exercise
  n - move to next exercise
  h - show hint
  q - quit`,
	RunE: startWatchMode,
}

func startWatchMode(cmd *cobra.Command, args []string) error {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	// Check if exercises exist
	exerciseDir := filepath.Join(cwd, "exercises")
	if _, err := os.Stat(exerciseDir); os.IsNotExist(err) {
		fmt.Println("❌ No exercises directory found.")
		fmt.Println("💡 Run 'goforgo init' to initialize exercises in this directory.")
		return nil
	}

	// Start the TUI interface (like Rustlings)
	return tui.CheckAndInitializeTUI(cwd)
}


func init() {
	rootCmd.AddCommand(watchCmd)
	
	// Add flags
	watchCmd.Flags().BoolVar(&manualRun, "manual", false, "Disable file watching, require manual execution")
	
	// Make watch the default command when no subcommand is specified
	rootCmd.RunE = startWatchMode
}