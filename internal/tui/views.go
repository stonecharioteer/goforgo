package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/stonecharioteer/goforgo/internal/exercise"
)

// renderWelcome shows the welcome screen (like Rustlings)
func (m *Model) renderWelcome() string {
	logo := `       Welcome to...
                 _____      ______           _____      
                /  __ \     |  ___|          |  __ \     
                | |  \/  ___| |_ ___  _ __    | |  \/ ___ 
                | | __ / _ \  _/ _ \| '__|   | | __ / _ \
                | |_\ \ (_) | || (_) | |      | |_\ \ (_) |
                 \____/\___/|_| \___/|_|       \____/\___/`

	welcomeText := fmt.Sprintf(`%s

🎯 Interactive Go Learning Platform

Welcome to GoForGo! This tool will help you learn Go through interactive exercises.
You'll fix broken code, learn Go concepts, and build practical skills.

📚 %d exercises loaded, covering:
   • Go fundamentals and syntax
   • Data structures and algorithms  
   • Concurrency and channels
   • Popular libraries and frameworks
   • Real-world programming patterns

🎮 How it works:
   1. Edit exercise files in your favorite editor
   2. Save the file and watch real-time feedback
   3. Fix errors and complete exercises progressively
   4. Get hints when you're stuck

⌨️  Keyboard shortcuts:
   • Enter/Space: Start exercises
   • h: Show hints
   • l: List all exercises  
   • n/p: Next/previous exercise
   • r: Manually run exercise
   • q: Quit

`, titleStyle.Render(logo), m.totalCount)

	if m.currentExercise != nil {
		welcomeText += fmt.Sprintf(`🚀 Ready to start with: %s
   %s

`, successStyle.Render(m.currentExercise.Info.Name), m.currentExercise.Description.Title)
	}

	welcomeText += statusStyle.Render("Press Enter to begin your Go journey! 🎉")

	// Center the content
	width := max(m.width, 80)
	style := lipgloss.NewStyle().
		Width(width).
		Align(lipgloss.Center).
		Padding(2, 0)

	return style.Render(welcomeText)
}

// renderMain shows the main exercise interface (like Rustlings watch mode)
func (m *Model) renderMain() string {
	if m.currentExercise == nil {
		return m.renderCompleted()
	}

	var content strings.Builder

	// Header with progress
	header := m.renderHeader()
	content.WriteString(header)
	content.WriteString("\n\n")

	// Current exercise info
	exerciseInfo := m.renderExerciseInfo()
	content.WriteString(exerciseInfo)
	content.WriteString("\n\n")

	// Exercise status and results
	results := m.renderResults()
	content.WriteString(results)
	content.WriteString("\n\n")

	// Footer with shortcuts
	footer := m.renderFooter()
	content.WriteString(footer)

	return content.String()
}

// renderHeader shows the progress bar and current status
func (m *Model) renderHeader() string {
	progress := float64(m.completedCount) / float64(m.totalCount)
	progressPercent := int(progress * 100)

	// Create progress bar
	barWidth := max(m.width-20, 40)
	filled := int(float64(barWidth) * progress)
	bar := strings.Repeat("█", filled) + strings.Repeat("▒", barWidth-filled)

	progressText := fmt.Sprintf("Progress: %d/%d (%d%%)", m.completedCount, m.totalCount, progressPercent)

	header := fmt.Sprintf(`%s

%s %s`,
		headerStyle.Render("🚀 GoForGo - Interactive Go Tutorial"),
		progressBarStyle.Render(bar),
		progressBarStyle.Render(progressText))

	return header
}

// renderExerciseInfo shows current exercise details
func (m *Model) renderExerciseInfo() string {
	ex := m.currentExercise
	
	difficulty := ex.GetDifficultyString()
	filePath := codeStyle.Render(ex.FilePath)

	info := fmt.Sprintf(`📝 Current Exercise: %s
📁 File: %s
%s %s

📖 %s`,
		titleStyle.Render(ex.Description.Title),
		filePath,
		difficulty,
		statusStyle.Render(fmt.Sprintf("(Exercise %d of %d)", m.currentIndex+1, m.totalCount)),
		ex.Description.Summary)

	if len(ex.Description.LearningObjectives) > 0 {
		info += "\n\n🎯 Learning Objectives:"
		for _, objective := range ex.Description.LearningObjectives {
			info += fmt.Sprintf("\n   • %s", objective)
		}
	}

	return info
}

// renderResults shows compilation/execution results
func (m *Model) renderResults() string {
	if m.isRunning {
		return statusStyle.Render("🔄 Running exercise...")
	}

	if m.lastResult == nil {
		return statusStyle.Render("💡 Save the file to see results...")
	}

	var result strings.Builder

	if m.lastResult.Success {
		result.WriteString(successStyle.Render("✅ SUCCESS! Exercise completed!"))
		result.WriteString("\n\n")
		result.WriteString("🎉 Well done! You've mastered this concept.")
		
		if m.currentIndex < len(m.exercises)-1 {
			result.WriteString("\n")
			result.WriteString(statusStyle.Render("Press 'n' for the next exercise."))
		} else {
			result.WriteString("\n")
			result.WriteString(successStyle.Render("🏆 All exercises completed! You're a Go expert!"))
		}
	} else {
		result.WriteString(errorStyle.Render("❌ Not quite there yet..."))
		result.WriteString("\n\n")

		// Show compilation or execution errors
		if m.lastResult.Error != "" {
			result.WriteString(errorStyle.Render("Error: "))
			result.WriteString(m.lastResult.Error)
			result.WriteString("\n\n")
		}

		if m.lastResult.Output != "" {
			result.WriteString("🔨 Output:\n")
			result.WriteString(codeStyle.Render(m.lastResult.Output))
			result.WriteString("\n\n")
		}

		result.WriteString(hintStyle.Render("💡 Need help? Press 'h' for a hint!"))
		result.WriteString("\n")
		result.WriteString(statusStyle.Render("Keep editing and save to try again."))
	}

	if m.statusMessage != "" {
		result.WriteString("\n\n")
		result.WriteString(statusStyle.Render(m.statusMessage))
	}

	return result.String()
}

// renderFooter shows keyboard shortcuts
func (m *Model) renderFooter() string {
	shortcuts := []string{
		"[n] next",
		"[p] prev", 
		"[h] hint",
		"[l] list",
		"[r] run",
		"[q] quit",
	}

	footer := statusStyle.Render("⌨️  " + strings.Join(shortcuts, " • "))
	
	// Add file watching status
	if m.watcherErr == nil {
		footer += "\n" + statusStyle.Render("👁️  Watching for file changes...")
	} else {
		footer += "\n" + errorStyle.Render("⚠️  File watcher error: " + m.watcherErr.Error())
	}

	return footer
}

// renderHint shows the hint overlay
func (m *Model) renderHint() string {
	if m.currentExercise == nil {
		return "No exercise selected."
	}

	// Build progressive hints based on current hint level
	var hints []string
	maxLevel := m.getMaxHintLevel()
	
	if m.currentHintLevel >= 1 && m.currentExercise.Hints.Level1 != "" {
		hints = append(hints, fmt.Sprintf("💡 Hint 1:\n%s", m.currentExercise.Hints.Level1))
	}
	if m.currentHintLevel >= 2 && m.currentExercise.Hints.Level2 != "" {
		hints = append(hints, fmt.Sprintf("💡 Hint 2:\n%s", m.currentExercise.Hints.Level2))
	}
	if m.currentHintLevel >= 3 && m.currentExercise.Hints.Level3 != "" {
		hints = append(hints, fmt.Sprintf("💡 Hint 3:\n%s", m.currentExercise.Hints.Level3))
	}
	
	var hintText string
	if len(hints) == 0 {
		hintText = "No hints available for this exercise."
	} else {
		hintText = strings.Join(hints, "\n\n")
	}
	
	// Show progression info
	var progressInfo string
	if maxLevel > 1 {
		if m.currentHintLevel < maxLevel {
			progressInfo = fmt.Sprintf("Press 'h' again for more hints (%d/%d)", m.currentHintLevel, maxLevel)
		} else {
			progressInfo = fmt.Sprintf("All hints shown (%d/%d) - Press 'h' to hide", m.currentHintLevel, maxLevel)
		}
	} else {
		progressInfo = "Press 'h' to hide hint"
	}

	content := fmt.Sprintf(`%s

📝 Exercise: %s

%s

%s
%s`,
		headerStyle.Render("💡 Hints"),
		titleStyle.Render(m.currentExercise.Description.Title),
		hintStyle.Render(hintText),
		statusStyle.Render(progressInfo),
		statusStyle.Render("Press Enter or Esc to return"))

	// Center and add border
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#F59E0B")).
		Padding(2).
		Width(min(m.width-4, 80))

	return style.Render(content)
}

// renderExerciseList shows all exercises with status
func (m *Model) renderExerciseList() string {
	var content strings.Builder

	content.WriteString(headerStyle.Render("📚 Exercise List"))
	content.WriteString("\n\n")

	// Group exercises by category
	categories := make(map[string][]*exercise.Exercise)
	for _, ex := range m.exercises {
		categories[ex.Info.Category] = append(categories[ex.Info.Category], ex)
	}

	for category, exercises := range categories {
		// Category header
		categoryName := strings.ReplaceAll(category, "_", " ")
		categoryName = strings.Title(categoryName)
		content.WriteString(titleStyle.Render(fmt.Sprintf("📁 %s", categoryName)))
		content.WriteString("\n")

		// List exercises in category
		for _, ex := range exercises {
			status := "❌"
			if ex.Completed {
				status = "✅"
			}
			
			marker := "  "
			if ex == m.currentExercise {
				marker = "▶ "
			}

			content.WriteString(fmt.Sprintf("%s%s %-20s %s\n", 
				marker, status, ex.Info.Name, ex.GetDifficultyString()))
		}
		content.WriteString("\n")
	}

	content.WriteString(statusStyle.Render("Press Enter or Esc to return"))

	// Add scrolling if needed (simplified for now)
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#7C3AED")).
		Padding(1).
		Width(min(m.width-4, 80)).
		Height(min(m.height-4, 30))

	return style.Render(content.String())
}

// renderCompleted shows completion screen
func (m *Model) renderCompleted() string {
	completion := `🎉 Congratulations! 🎉

You've completed all GoForGo exercises!

🏆 You are now a Go expert! You've mastered:
   • Go syntax and fundamentals
   • Data structures and algorithms
   • Concurrency and channels
   • Popular libraries and frameworks
   • Real-world programming patterns

🚀 What's Next?
   • Build your own Go projects
   • Contribute to open source
   • Explore advanced Go topics
   • Share your knowledge with others

Thank you for using GoForGo! 🎊`

	style := lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Center).
		Padding(4, 0)

	return style.Render(successStyle.Render(completion))
}

// Helper functions
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}