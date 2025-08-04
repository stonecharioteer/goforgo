package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
	
	// Global flags
	workingDirectory string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goforgo",
	Short: "Interactive Go tutorial inspired by Rustlings",
	Long: `GoForGo is an interactive CLI-based tutorial for learning Golang.

It provides 250+ exercises covering Go fundamentals through advanced topics
and popular libraries, with a beautiful terminal user interface that gives
real-time feedback as you code.

Similar to Rustlings for Rust, GoForGo helps you learn Go by fixing broken
code exercises, with automatic compilation and testing to guide your progress.`,
	Version: fmt.Sprintf("%s (commit: %s, built: %s)", version, commit, date),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringVarP(&workingDirectory, "directory", "d", "", "working directory for exercises (default: current directory)")
}

// GetWorkingDirectory returns the working directory, defaulting to current directory
func GetWorkingDirectory() (string, error) {
	if workingDirectory != "" {
		return workingDirectory, nil
	}
	return os.Getwd()
}