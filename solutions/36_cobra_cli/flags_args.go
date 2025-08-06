// flags_args.go - Solution
// Learn Cobra flags and advanced argument handling

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
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Printf("Verbose mode enabled\n")
			fmt.Printf("Output format: %s\n", output)
		}
		
		format, _ := cmd.Flags().GetString("format")
		count, _ := cmd.Flags().GetInt("count")
		
		fmt.Printf("Processing %d files in %s format\n", len(args), format)
		fmt.Printf("Max items to process: %d\n", count)
		
		for i, file := range args {
			fmt.Printf("File %d: %s\n", i+1, file)
		}
	},
}

var searchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search with query and options",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		
		caseSensitive, _ := cmd.Flags().GetBool("case-sensitive")
		limit, _ := cmd.Flags().GetInt("limit")
		exclude, _ := cmd.Flags().GetStringSlice("exclude")
		
		fmt.Printf("Searching for: %s\n", query)
		fmt.Printf("Case sensitive: %t\n", caseSensitive)
		fmt.Printf("Result limit: %d\n", limit)
		
		if len(exclude) > 0 {
			fmt.Printf("Excluding: %s\n", strings.Join(exclude, ", "))
		}
	},
}

func init() {
	// Add persistent flags to root command (available to all subcommands)
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "text", "Output format (text, json, xml)")

	// Add the process command to root
	rootCmd.AddCommand(processCmd)
	
	// Add local flags to process command
	processCmd.Flags().String("format", "default", "Processing format")
	processCmd.Flags().IntP("count", "c", 10, "Maximum number of items to process")

	// Add the search command to root
	rootCmd.AddCommand(searchCmd)
	
	// Add flags to search command
	searchCmd.Flags().Bool("case-sensitive", false, "Perform case-sensitive search")
	searchCmd.Flags().IntP("limit", "l", 100, "Maximum number of results")
	searchCmd.Flags().StringSlice("exclude", []string{}, "Terms to exclude from search")
	
	// Mark the case-sensitive flag as required
	searchCmd.MarkFlagRequired("case-sensitive")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}