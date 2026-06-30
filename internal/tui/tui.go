package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbletea"
	"github.com/stonecharioteer/goforgo/internal/exercise"
	"github.com/stonecharioteer/goforgo/internal/runner"
)

// RunTUI starts the Bubble Tea TUI application
func RunTUI(exerciseManager *exercise.ExerciseManager, runner *runner.Runner) error {
	return RunTUIWithNotice(exerciseManager, runner, "")
}

// RunTUIWithNotice starts the Bubble Tea TUI application with an optional startup notice.
func RunTUIWithNotice(exerciseManager *exercise.ExerciseManager, runner *runner.Runner, notice string) error {
	// Create the model
	model := NewModel(exerciseManager, runner)
	if notice != "" {
		model.SetUpdateNotice(notice)
	}

	// Create the program with options
	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),       // Use alternate screen buffer
		tea.WithMouseCellMotion(), // Enable mouse support
	)

	// Handle cleanup on exit
	defer func() {
		if model.watcher != nil {
			_ = model.watcher.Close()
		}
	}()

	// Run the program
	finalModel, err := p.Run()
	if err != nil {
		return fmt.Errorf("failed to run TUI: %w", err)
	}

	// Check if we need to show any final messages
	if m, ok := finalModel.(*Model); ok {
		if m.getCompletedCount() == m.getTotalCount() {
			fmt.Println("🎉 Congratulations on completing all exercises!")
		}
	}

	return nil
}

// CheckAndInitializeTUI checks if we should run the TUI and initializes it
func CheckAndInitializeTUI(basePath string) error {
	return CheckAndInitializeTUIWithNotice(basePath, "")
}

// CheckAndInitializeTUIWithNotice checks if we should run the TUI and initializes it with an optional notice.
func CheckAndInitializeTUIWithNotice(basePath, notice string) error {
	// Check if we're in a terminal
	if !isTerminal() {
		return fmt.Errorf("GoForGo requires a terminal to run the interactive interface")
	}

	// Initialize exercise manager
	em := exercise.NewExerciseManager(basePath)
	if err := em.LoadExercises(); err != nil {
		return fmt.Errorf("failed to load exercises: %w", err)
	}

	// Initialize runner
	r := runner.NewRunner(basePath)

	// Start the TUI
	return RunTUIWithNotice(em, r, notice)
}

// isTerminal checks if we're running in a terminal
func isTerminal() bool {
	// Check if stdout is a terminal
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) != 0 {
		return true
	}
	return false
}

// ShowProgress displays a simple progress indicator for CLI mode
func ShowProgress(completed, total int, current string) {
	progress := float64(completed) / float64(total) * 100
	fmt.Printf("Progress: %d/%d (%.1f%%) - Current: %s\n", completed, total, progress, current)
}
