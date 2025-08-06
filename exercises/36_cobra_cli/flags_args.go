// flags_args.go
// Learn Cobra flags and advanced argument handling
//
// Cobra provides powerful flag handling capabilities including persistent flags,
// required flags, default values, and flag binding. This exercise demonstrates
// various flag types and argument validation techniques.
//
// I AM NOT DONE YET!

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	// Global flags
	verbose bool
	output  string
)

var rootCmd = &cobra.Command{
	Use:   "flagapp",
	Short: "Demonstrate Cobra flags and arguments",
	Long:  `This application shows various flag types and argument handling in Cobra.`,
}

var processCmd = &cobra.Command{
	Use:   "process [file...]",
	Short: "Process files with various options",
	Long:  `Process one or more files with configurable options using flags.`,
	// TODO: Set Args to require at least one argument
	// Args: cobra.???,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Check if verbose flag is set and print debug info
		// if ??? {
		//     fmt.Printf("Verbose mode enabled\n")
		//     fmt.Printf("Output format: %s\n", ???)
		// }
		
		// TODO: Get the format flag value
		// format, _ := cmd.Flags().GetString(???)
		
		// TODO: Get the count flag value  
		// count, _ := cmd.Flags().GetInt(???)
		
		// TODO: Print processing information
		// fmt.Printf("Processing %d files in %s format\n", ???, ???)
		// fmt.Printf("Max items to process: %d\n", ???)
		
		// TODO: List all files to process
		// for i, file := range ??? {
		//     fmt.Printf("File %d: %s\n", i+1, ???)
		// }
	},
}

var searchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search with query and options",
	// TODO: Set Args to require exactly one argument
	// Args: cobra.???,
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		
		// TODO: Get the case-sensitive flag
		// caseSensitive, _ := cmd.Flags().GetBool(???)
		
		// TODO: Get the limit flag  
		// limit, _ := cmd.Flags().GetInt(???)
		
		// TODO: Get the exclude flag (string slice)
		// exclude, _ := cmd.Flags().GetStringSlice(???)
		
		// TODO: Print search parameters
		// fmt.Printf("Searching for: %s\n", ???)
		// fmt.Printf("Case sensitive: %t\n", ???)
		// fmt.Printf("Result limit: %d\n", ???)
		
		// TODO: Print excluded terms if any
		// if len(???) > 0 {
		//     fmt.Printf("Excluding: %s\n", strings.Join(???, ", "))
		// }
	},
}

func init() {
	// TODO: Add persistent flags to root command (available to all subcommands)
	// rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	// rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "text", "Output format (text, json, xml)")

	// TODO: Add the process command to root
	// rootCmd.AddCommand(???)
	
	// TODO: Add local flags to process command
	// processCmd.Flags().String("format", "default", "Processing format")
	// processCmd.Flags().IntP("count", "c", 10, "Maximum number of items to process")

	// TODO: Add the search command to root
	// rootCmd.AddCommand(???)
	
	// TODO: Add flags to search command
	// searchCmd.Flags().Bool("case-sensitive", false, "Perform case-sensitive search")
	// searchCmd.Flags().IntP("limit", "l", 100, "Maximum number of results")
	// searchCmd.Flags().StringSlice("exclude", []string{}, "Terms to exclude from search")
	
	// TODO: Mark the case-sensitive flag as required
	// searchCmd.MarkFlagRequired(???)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}