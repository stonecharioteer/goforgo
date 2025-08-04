package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
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
	// Global flags can be defined here
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goforgo.yaml)")
}