// subcommands.go
// Learn to build nested command hierarchies with Cobra
//
// Cobra supports complex command hierarchies with multiple levels of nesting.
// This is useful for creating CLI tools with grouped functionality like
// kubectl, docker, or git which have commands like "kubectl get pods" or "docker image ls".
//
// I AM NOT DONE YET!

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
	// TODO: Set Use to "user"
	// Use: ???,
	
	// TODO: Set Short description
	// Short: ???,
	
	// TODO: Add a Run function that shows available user subcommands
	// Run: func(cmd *cobra.Command, args []string) {
	//     fmt.Println("User management commands. Use --help to see subcommands.")
	// },
}

var userListCmd = &cobra.Command{
	// TODO: Set Use to "list"
	// Use: ???,
	
	// TODO: Set Short description
	// Short: ???,
	
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Get the format flag and print user list
		// format, _ := cmd.Flags().GetString(???)
		// fmt.Printf("Listing users in %s format:\n", ???)
		// fmt.Println("1. alice@example.com")
		// fmt.Println("2. bob@example.com") 
	},
}

var userCreateCmd = &cobra.Command{
	// TODO: Set Use to "create <email>"  
	// Use: ???,
	
	// TODO: Set Short description
	// Short: ???,
	
	// TODO: Set Args to require exactly one argument
	// Args: cobra.???,
	
	Run: func(cmd *cobra.Command, args []string) {
		email := args[0]
		
		// TODO: Get the admin flag
		// isAdmin, _ := cmd.Flags().GetBool(???)
		
		// TODO: Print user creation message
		// fmt.Printf("Creating user: %s\n", ???)
		// if ??? {
		//     fmt.Println("User will have admin privileges")  
		// }
	},
}

// Database commands
var dbCmd = &cobra.Command{
	// TODO: Set Use to "db"
	// Use: ???,
	
	// TODO: Set Short description  
	// Short: ???,
	
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Database commands. Use --help to see subcommands.")
	},
}

var dbMigrateCmd = &cobra.Command{
	// TODO: Set Use to "migrate"
	// Use: ???,
	
	// TODO: Set Short description
	// Short: ???,
	
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Get the dry-run flag
		// dryRun, _ := cmd.Flags().GetBool(???)
		
		// TODO: Print migration message based on dry-run flag
		// if ??? {
		//     fmt.Println("Dry run: Would run database migrations")
		// } else {
		//     fmt.Println("Running database migrations...")
		// }
	},
}

var dbBackupCmd = &cobra.Command{
	// TODO: Set Use to "backup [filename]"
	// Use: ???,
	
	// TODO: Set Short description
	// Short: ???,
	
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Use default filename if none provided
		// filename := "backup.sql"
		// if len(???) > 0 {
		//     filename = args[???]
		// }
		
		// TODO: Print backup message
		// fmt.Printf("Creating database backup: %s\n", ???)
	},
}

func init() {
	// TODO: Add user command to root
	// rootCmd.AddCommand(???)
	
	// TODO: Add subcommands to user command
	// userCmd.AddCommand(???)
	// userCmd.AddCommand(???)
	
	// TODO: Add format flag to user list command
	// userListCmd.Flags().String("format", "table", "Output format (table, json)")
	
	// TODO: Add admin flag to user create command
	// userCreateCmd.Flags().Bool("admin", false, "Create user with admin privileges")

	// TODO: Add db command to root
	// rootCmd.AddCommand(???)
	
	// TODO: Add subcommands to db command  
	// dbCmd.AddCommand(???)
	// dbCmd.AddCommand(???)
	
	// TODO: Add dry-run flag to migrate command
	// dbMigrateCmd.Flags().Bool("dry-run", false, "Show what would be done without executing")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}