package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/stonecharioteer/goforgo/internal/exercise"
	"github.com/stonecharioteer/goforgo/internal/validation"
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
	em, cwd, err := loadExerciseManager()
	if err != nil {
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

	// Create universal runner and execute
	r := validation.NewUniversalRunner(cwd)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	result, err := r.ValidateExercise(ctx, ex)
	if err != nil {
		return fmt.Errorf("failed to run exercise: %w", err)
	}

	// Display results
	feedback := r.FormatValidationResult(result)
	fmt.Println(feedback)

	// Show summary for universal validation
	if len(result.ServiceResults) > 0 || len(result.ValidationResults) > 0 {
		summary := r.GetValidationSummary(result)
		fmt.Printf("\n📊 Validation Summary:\n")
		fmt.Printf("   Duration: %v\n", summary["duration"])
		fmt.Printf("   Services: %d/%d successful\n", summary["successful_services"], summary["services_count"])
		fmt.Printf("   Rules: %d/%d successful\n", summary["successful_rules"], summary["rules_count"])
		if envVars, ok := summary["environment_vars"].(int); ok && envVars > 0 {
			fmt.Printf("   Environment Variables: %d injected\n", envVars)
		}
	}

	// Always cleanup resources when done
	defer func() {
		cleanupCtx, cleanupCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cleanupCancel()
		if err := r.Cleanup(cleanupCtx); err != nil {
			fmt.Printf("⚠️  Warning: Failed to cleanup resources: %v\n", err)
		}
	}()

	success := result.Success

	if success {
		// Mark the exercise as completed
		if err := em.MarkExerciseCompleted(ex.Info.Name); err != nil {
			fmt.Printf("⚠️  Warning: Failed to save progress: %v\n", err)
		}

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
		// Increment attempt count for better hint selection
		ex.Attempts++
		fmt.Printf("💡 Hint: %s\n", ex.GetHint())
		fmt.Printf("\n🔧 Edit the file and run 'goforgo run %s' again, or use 'goforgo' for watch mode.\n", ex.Info.Name)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(runCmd)
}
