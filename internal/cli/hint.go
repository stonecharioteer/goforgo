package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stonecharioteer/goforgo/internal/exercise"
)

// hintCmd represents the hint command
var hintCmd = &cobra.Command{
	Use:   "hint [exercise_name]",
	Short: "Show a hint for an exercise",
	Long: `Show a hint for the specified exercise.

If no exercise name is provided, shows a hint for the next incomplete exercise.

Examples:
  goforgo hint hello         # Show hint for the 'hello' exercise
  goforgo hint               # Show hint for the next incomplete exercise`,
	RunE: showHint,
}

func showHint(cmd *cobra.Command, args []string) error {
	em, _, err := loadExerciseManager()
	if err != nil {
		return err
	}

	// Determine which exercise to show hint for
	var ex *exercise.Exercise
	if len(args) > 0 {
		// Show hint for specific exercise
		exerciseName := args[0]
		ex, err = em.GetExerciseByName(exerciseName)
		if err != nil {
			return err
		}
	} else {
		// Show hint for next incomplete exercise
		ex = em.GetNextExercise()
		if ex == nil {
			fmt.Println("🎉 Congratulations! You've completed all exercises!")
			return nil
		}
	}

	// Display exercise information and hint
	fmt.Printf("💡 Hint for exercise: %s\n", ex.String())
	fmt.Printf("📁 File: %s\n", ex.FilePath)
	fmt.Printf("⭐ Difficulty: %s\n\n", ex.GetDifficultyString())

	fmt.Printf("📖 Description: %s\n\n", ex.Description.Summary)

	// Show learning objectives if available
	if len(ex.Description.LearningObjectives) > 0 {
		fmt.Println("🎯 Learning Objectives:")
		for _, objective := range ex.Description.LearningObjectives {
			fmt.Printf("  • %s\n", objective)
		}
		fmt.Println()
	}

	// Show the hint
	fmt.Printf("💡 Hint: %s\n\n", ex.GetHint())

	fmt.Printf("🔧 Edit %s and run 'goforgo run %s' to test your solution.\n", ex.FilePath, ex.Info.Name)

	return nil
}

func init() {
	rootCmd.AddCommand(hintCmd)
}
