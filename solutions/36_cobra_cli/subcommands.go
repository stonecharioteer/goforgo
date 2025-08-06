// subcommands.go - Solution
// Learn to build nested command hierarchies with Cobra

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tool",
	Short: "A CLI tool with nested subcommands",
	Long:  `Demonstrates nested command hierarchies in Cobra applications.`,
}

// User management commands
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User management commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("User management commands. Use --help to see subcommands.")
	},
}

var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users",
	Run: func(cmd *cobra.Command, args []string) {
		format, _ := cmd.Flags().GetString("format")
		fmt.Printf("Listing users in %s format:\n", format)
		fmt.Println("1. alice@example.com")
		fmt.Println("2. bob@example.com")
	},
}

var userCreateCmd = &cobra.Command{
	Use:   "create <email>",
	Short: "Create a new user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		email := args[0]
		isAdmin, _ := cmd.Flags().GetBool("admin")
		
		fmt.Printf("Creating user: %s\n", email)
		if isAdmin {
			fmt.Println("User will have admin privileges")
		}
	},
}

// Database commands
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Database management commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Database commands. Use --help to see subcommands.")
	},
}

var dbMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		
		if dryRun {
			fmt.Println("Dry run: Would run database migrations")
		} else {
			fmt.Println("Running database migrations...")
		}
	},
}

var dbBackupCmd = &cobra.Command{
	Use:   "backup [filename]",
	Short: "Create database backup",
	Run: func(cmd *cobra.Command, args []string) {
		filename := "backup.sql"
		if len(args) > 0 {
			filename = args[0]
		}
		
		fmt.Printf("Creating database backup: %s\n", filename)
	},
}

func init() {
	// Add user command to root
	rootCmd.AddCommand(userCmd)
	
	// Add subcommands to user command
	userCmd.AddCommand(userListCmd)
	userCmd.AddCommand(userCreateCmd)
	
	// Add format flag to user list command
	userListCmd.Flags().String("format", "table", "Output format (table, json)")
	
	// Add admin flag to user create command
	userCreateCmd.Flags().Bool("admin", false, "Create user with admin privileges")

	// Add db command to root
	rootCmd.AddCommand(dbCmd)
	
	// Add subcommands to db command
	dbCmd.AddCommand(dbMigrateCmd)
	dbCmd.AddCommand(dbBackupCmd)
	
	// Add dry-run flag to migrate command
	dbMigrateCmd.Flags().Bool("dry-run", false, "Show what would be done without executing")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}