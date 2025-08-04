package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stonecharioteer/goforgo/internal/exercise"
	"github.com/stonecharioteer/goforgo/internal/runner"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [exercise_name]",
	Short: "Run a specific exercise",
	Long: `Run a specific exercise by name.

If no exercise name is provided, runs the next incomplete exercise.

Examples:
  goforgo run hello          # Run the 'hello' exercise
  goforgo run                # Run the next incomplete exercise`,
	RunE: runExercise,
}

func runExercise(cmd *cobra.Command, args []string) error {
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	// Initialize exercise manager
	em := exercise.NewExerciseManager(cwd)
	if err := em.LoadExercises(); err != nil {
		return err
	}

	// Determine which exercise to run
	var ex *exercise.Exercise
	if len(args) > 0 {
		// Run specific exercise
		exerciseName := args[0]
		ex, err = em.GetExerciseByName(exerciseName)
		if err != nil {
			return err
		}
	} else {
		// Run next incomplete exercise
		ex = em.GetNextExercise()
		if ex == nil {
			fmt.Println("🎉 Congratulations! You've completed all exercises!")
			return nil
		}
	}

	// Display exercise information
	fmt.Printf("🚀 Running exercise: %s\n", ex.String())
	fmt.Printf("📁 File: %s\n", ex.FilePath)
	fmt.Printf("⭐ Difficulty: %s\n", ex.GetDifficultyString())
	fmt.Printf("📖 Description: %s\n\n", ex.Description.Summary)

	// Create runner and execute
	r := runner.NewRunner(cwd)
	success, feedback, err := r.ValidateExercise(ex)
	if err != nil {
		return fmt.Errorf("failed to run exercise: %w", err)
	}

	// Display results
	fmt.Println(feedback)

	if success {
		fmt.Printf("🎯 Exercise '%s' completed! 🎉\n", ex.Info.Name)
		
		// Suggest next steps
		nextEx := em.GetNextExercise()
		if nextEx != nil {
			fmt.Printf("\n💡 Next exercise: %s (%s)\n", nextEx.Info.Name, nextEx.Description.Title)
			fmt.Println("   Run 'goforgo' to enter watch mode or 'goforgo run' for the next exercise.")
		} else {
			fmt.Println("\n🏆 All exercises completed! You're a Go expert now!")
		}
	} else {
		fmt.Printf("💡 Hint: %s\n", ex.GetHint())
		fmt.Printf("\n🔧 Edit the file and run 'goforgo run %s' again, or use 'goforgo' for watch mode.\n", ex.Info.Name)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(runCmd)
}