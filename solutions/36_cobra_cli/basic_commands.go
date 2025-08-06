// basic_commands.go - Solution
// Learn to build command-line applications with Cobra

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "A simple CLI application built with Cobra",
	Long: `This is a sample application demonstrating basic Cobra usage.
It shows how to create commands, handle arguments, and provide help text.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to MyApp! Use --help to see available commands.")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("MyApp v1.0.0")
	},
}

var greetCmd = &cobra.Command{
	Use:   "greet [name]",
	Short: "Greet someone by name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		fmt.Printf("Hello, %s!\n", name)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(greetCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}