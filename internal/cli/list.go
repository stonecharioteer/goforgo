package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stonecharioteer/goforgo/internal/exercise"
)

var (
	listAll      bool
	listCategory string
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all exercises with their completion status",
	Long: `List all exercises with their completion status.

By default, shows only incomplete exercises. Use flags to customize the view.

Examples:
  goforgo list                    # Show incomplete exercises
  goforgo list --all              # Show all exercises
  goforgo list --category basics  # Show exercises in 'basics' category`,
	RunE: listExercises,
}

func listExercises(cmd *cobra.Command, args []string) error {
	// Get working directory
	cwd, err := GetWorkingDirectory()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	// Initialize exercise manager
	em := exercise.NewExerciseManager(cwd)
	if err := em.LoadExercises(); err != nil {
		return err
	}

	exercises := em.GetExercises()
	if len(exercises) == 0 {
		fmt.Println("No exercises found. Run 'goforgo init' to set up exercises.")
		return nil
	}

	// Filter exercises based on flags
	var filteredExercises []*exercise.Exercise

	for _, ex := range exercises {
		// Apply filters
		if !listAll && ex.Completed {
			continue // Skip completed exercises unless --all is specified
		}

		if listCategory != "" && !strings.Contains(strings.ToLower(ex.Info.Category), strings.ToLower(listCategory)) {
			continue // Skip exercises not in the specified category
		}

		filteredExercises = append(filteredExercises, ex)
	}

	// Display header with progress using centralized counting
	completedCount, totalCount, percentage := em.GetProgressStats()
	fmt.Printf("GoForGo Exercises - Progress: %d/%d (%.1f%% complete)\n", 
		completedCount, totalCount, percentage)
	fmt.Println(strings.Repeat("═", 60))

	if len(filteredExercises) == 0 {
		if listAll {
			fmt.Println("🎉 All exercises completed!")
		} else {
			fmt.Println("🎉 No incomplete exercises found! All done!")
		}
		return nil
	}

	// Group exercises by category
	categories := make(map[string][]*exercise.Exercise)
	for _, ex := range filteredExercises {
		categories[ex.Info.Category] = append(categories[ex.Info.Category], ex)
	}

	// Display exercises grouped by category
	for category, categoryExercises := range categories {
		fmt.Printf("\n📁 %s\n", strings.ToTitle(strings.ReplaceAll(category, "_", " ")))
		fmt.Println(strings.Repeat("─", 40))

		for _, ex := range categoryExercises {
			status := "❌"
			if ex.Completed {
				status = "✅"
			}

			// Format difficulty stars
			difficultyStr := ex.GetDifficultyString()

			fmt.Printf("  %s %-20s %s\n", status, ex.Info.Name, difficultyStr)
			fmt.Printf("      %s\n", ex.Description.Title)
			
			if ex.Info.EstimatedTime != "" {
				fmt.Printf("      ⏱️  Estimated time: %s\n", ex.Info.EstimatedTime)
			}
			fmt.Println()
		}
	}

	// Show next steps
	if !listAll {
		nextEx := em.GetNextExercise()
		if nextEx != nil {
			fmt.Printf("🎯 Next exercise to work on: %s\n", nextEx.Info.Name)
			fmt.Printf("   Run 'goforgo run %s' or 'goforgo' to start working on it.\n", nextEx.Info.Name)
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
	
	// Add flags
	listCmd.Flags().BoolVar(&listAll, "all", false, "Show all exercises including completed ones")
	listCmd.Flags().StringVar(&listCategory, "category", "", "Filter exercises by category")
}