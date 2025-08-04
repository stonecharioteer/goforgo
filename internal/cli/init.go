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
	currentDir, err := GetWorkingDirectory()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	exerciseDir := filepath.Join(currentDir, "exercises")

	if _, err := os.Stat(exerciseDir); err == nil {
		return fmt.Errorf("exercises directory already exists in %s\nIf you want to reinitialize, please remove the directory first", currentDir)
	}

	fmt.Printf("🚀 Initializing GoForGo exercises in %s\n", currentDir)

	if err := InitializeExercises(currentDir); err != nil {
		return err
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

func init() {
	rootCmd.AddCommand(initCmd)
}
