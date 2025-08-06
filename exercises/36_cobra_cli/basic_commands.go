// basic_commands.go
// Learn to build command-line applications with Cobra
//
// Cobra is a library for creating powerful modern CLI applications.
// It's used by many popular Go projects including Docker, Kubernetes,
// and GitHub CLI. This exercise covers basic command creation and structure.
//
// I AM NOT DONE YET!

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
	// TODO: Set the Use field to "version" 
	// Use: ???,
	
	// TODO: Set the Short field to a brief description
	// Short: ???,
	
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Print the version information
		// fmt.Println(???)
	},
}

var greetCmd = &cobra.Command{
	// TODO: Set the Use field to "greet [name]"
	// Use: ???,
	
	// TODO: Set the Short field to describe greeting functionality
	// Short: ???,
	
	// TODO: Set the Args field to require exactly one argument
	// Args: cobra.???,
	
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Get the name from args[0] and print a greeting
		// name := args[???]
		// fmt.Printf(???, ???)
	},
}

func init() {
	// TODO: Add the version command to the root command
	// rootCmd.AddCommand(???)
	
	// TODO: Add the greet command to the root command
	// rootCmd.AddCommand(???)
}

func main() {
	// TODO: Execute the root command and handle any errors
	// if err := rootCmd.Execute(); err != nil {
	//     fmt.Println(err)
	//     os.Exit(???)
	// }
}