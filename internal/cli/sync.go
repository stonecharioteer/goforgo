package cli

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/stonecharioteer/goforgo/internal/runner"
)

const barWidth = 30

var (
	syncBarFilled = lipgloss.NewStyle().Foreground(lipgloss.Color("#3fb950"))
	syncBarEmpty  = lipgloss.NewStyle().Foreground(lipgloss.Color("#484f58"))
	syncPass      = lipgloss.NewStyle().Foreground(lipgloss.Color("#3fb950")).Bold(true)
	syncFail      = lipgloss.NewStyle().Foreground(lipgloss.Color("#f85149"))
	syncDim       = lipgloss.NewStyle().Foreground(lipgloss.Color("#8b949e"))
	syncBold      = lipgloss.NewStyle().Foreground(lipgloss.Color("#e6edf3")).Bold(true)
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Re-validate all exercises and update progress",
	Long: `Validate every exercise against its current state and update the
progress file to match reality. Exercises that now pass get marked
complete; exercises that no longer pass get unmarked.

Examples:
  goforgo sync`,
	RunE: runSync,
}

func runSync(cmd *cobra.Command, args []string) error {
	em, cwd, err := loadExerciseManager()
	if err != nil {
		return err
	}

	r := runner.NewRunner(cwd)
	exercises := em.GetExercises()
	total := len(exercises)
	completed := 0

	for i, ex := range exercises {
		// Render progress bar
		progress := float64(i+1) / float64(total)
		filled := int(progress * barWidth)
		empty := barWidth - filled
		bar := syncBarFilled.Render(strings.Repeat("█", filled)) +
			syncBarEmpty.Render(strings.Repeat("░", empty))

		pct := fmt.Sprintf("%3d%%", int(progress*100))

		// Truncate exercise name to fit
		name := ex.Info.Name
		if len(name) > 28 {
			name = name[:25] + "..."
		}

		fmt.Printf("\r  %s %s  %s  %s",
			bar,
			syncDim.Render(pct),
			syncDim.Render(fmt.Sprintf("[%d/%d]", i+1, total)),
			syncDim.Render(name)+strings.Repeat(" ", 30-len(name)),
		)

		result, err := r.RunExercise(ex)
		if err != nil {
			continue
		}

		if result.Success {
			if !ex.Completed {
				if err := em.MarkExerciseCompleted(ex.Info.Name); err != nil {
					return fmt.Errorf("failed to mark exercise %q complete: %w", ex.Info.Name, err)
				}
			}
			completed++
		} else {
			if ex.Completed {
				if err := em.UnmarkExerciseCompleted(ex.Info.Name); err != nil {
					return fmt.Errorf("failed to unmark exercise %q complete: %w", ex.Info.Name, err)
				}
			}
		}
	}

	// Clear progress line and print summary
	fmt.Print("\r\033[K")

	fullBar := syncBarFilled.Render(strings.Repeat("█", barWidth))
	fmt.Printf("  %s %s\n\n", fullBar, syncDim.Render("100%"))
	fmt.Printf("  %s  %s passed  %s failed\n",
		syncBold.Render("Synced:"),
		syncPass.Render(fmt.Sprintf("%d", completed)),
		syncFail.Render(fmt.Sprintf("%d", total-completed)),
	)

	return nil
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
