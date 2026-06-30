package tui

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

// Layout and animation constants.
const (
	contentPadding      = 10
	minContentWidth     = 50
	maxContentWidth     = 90
	borderCharWidth     = 80
	progressBarWidth    = 30
	splashFrameCount    = 8
	splashTickMs        = 250
	listReservedHeight  = 8
	minListHeight       = 10
	columnPaddingFactor = 1.1
	minColumnWidth      = 3
)

// getContentWidth returns the usable content width for the current terminal size.
func (m *Model) getContentWidth() int {
	w := m.width - contentPadding
	if w < minContentWidth {
		w = minContentWidth
	}
	if w > maxContentWidth {
		w = maxContentWidth
	}
	return w
}

// renderWelcome shows the welcome screen (like Rustlings)
func (m *Model) renderWelcome() string {
	// Beautiful GoForGo text logo
	logo := `
   ██████╗  ██████╗ ███████╗ ██████╗ ██████╗  ██████╗  ██████╗ 
  ██╔════╝ ██╔═══██╗██╔════╝██╔═══██╗██╔══██╗██╔════╝ ██╔═══██╗
  ██║  ███╗██║   ██║█████╗  ██║   ██║██████╔╝██║  ███╗██║   ██║
  ██║   ██║██║   ██║██╔══╝  ██║   ██║██╔══██╗██║   ██║██║   ██║
  ╚██████╔╝╚██████╔╝██║     ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝
   ╚═════╝  ╚═════╝ ╚═╝      ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝`

	// Gradient colors for the logo
	logoStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#bc8cff")).
		Bold(true).
		Align(lipgloss.Center)

	// Create colorful banner with gradient effect
	banner := logoStyle.Render(logo)

	// Subtitle with gradient
	subtitleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#d2a8ff")).
		Bold(true).
		Italic(true).
		Align(lipgloss.Center)

	subtitle := subtitleStyle.Render("🚀 Interactive Go Learning Platform 🚀")

	// Stats section with progress - use dynamic completed count
	progressBar := m.renderProgressBar(m.getCompletedCount(), m.getTotalCount(), progressBarWidth)

	statsStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#3fb950")).
		Bold(true)

	statsText := statsStyle.Render(fmt.Sprintf("📊 Progress: %s %d/%d exercises completed", progressBar, m.getCompletedCount(), m.getTotalCount()))

	// Feature highlights with emojis and colors
	featuresStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#58a6ff"))

	features := featuresStyle.Render(`✨ What makes GoForGo special:
   🔥 Real-time feedback as you code
   🧠 Progressive hints that guide your learning  
   📈 Smart progress tracking with auto-skip
   🎯 TODO-driven exercises for flexible learning
   ⚡ Instant file change detection`)

	// Learning topics with nice formatting
	topicsStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#d29922"))

	topics := topicsStyle.Render(fmt.Sprintf(`📚 %d exercises covering:
   • Go fundamentals & syntax        • Error handling patterns
   • Variables & data types          • Concurrency & goroutines  
   • Functions & methods             • Channels & sync primitives
   • Structs & interfaces            • Popular libraries (Gin, GORM)
   • Control flow & loops            • Real-world projects`, m.getTotalCount()))

	// Keyboard shortcuts in a nice box
	shortcutsStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#f778ba")).
		Padding(0, 1).
		Foreground(lipgloss.Color("#f778ba"))

	shortcuts := shortcutsStyle.Render(`⌨️  Keyboard Shortcuts:
 Enter/Space  Start your Go journey  |  h  Progressive hints
 n / p        Next/Previous exercise |  l  List all exercises  
 r            Run current exercise   |  q  Quit GoForGo`)

	welcomeText := fmt.Sprintf(`%s

%s

%s

%s

%s

%s

`, banner, subtitle, statsText, features, topics, shortcuts)

	// Next exercise info with attractive styling
	if m.currentExercise != nil {
		nextStyle := lipgloss.NewStyle().
			Border(lipgloss.DoubleBorder()).
			BorderForeground(lipgloss.Color("#3fb950")).
			Padding(0, 2).
			Foreground(lipgloss.Color("#3fb950")).
			Bold(true)

		nextExercise := nextStyle.Render(fmt.Sprintf(`🎯 Next Exercise: %s
📖 %s
⭐ Difficulty: %s`,
			m.currentExercise.Info.Name,
			m.currentExercise.Description.Title,
			m.currentExercise.GetDifficultyString()))

		welcomeText += fmt.Sprintf(`%s

`, nextExercise)
	}

	// Call to action with pulsing effect styling
	ctaStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#e3b341")).
		Bold(true).
		Blink(true)

	welcomeText += ctaStyle.Render("✨ Press Enter to begin your Go journey! ✨")

	// Add decorative border using text characters
	borderStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#bc8cff"))
	borderLine := borderStyle.Render(strings.Repeat("═", borderCharWidth))

	welcomeText = fmt.Sprintf(`%s
%s
%s`, borderLine, welcomeText, borderLine)

	// Center and style the entire content
	// Account for border (2 chars) and padding (4 chars) = 6 chars total
	// Add extra margin for safety
	contentWidth := m.getContentWidth()

	// Use a simpler approach without overall border to avoid width issues
	style := lipgloss.NewStyle().
		Width(contentWidth).
		Align(lipgloss.Center).
		Padding(1, 0)

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

	// Apply consistent border styling like welcome screen
	mainContent := content.String()

	// Add decorative border using text characters
	borderStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#bc8cff"))
	borderLine := borderStyle.Render(strings.Repeat("═", borderCharWidth))

	borderedContent := fmt.Sprintf(`%s
%s
%s`, borderLine, mainContent, borderLine)

	// Style the content left-aligned
	contentWidth := m.getContentWidth()

	style := lipgloss.NewStyle().
		Width(contentWidth).
		Align(lipgloss.Left).
		Padding(1, 2)

	return style.Render(borderedContent)
}

// renderHeader shows the progress bar and current status
func (m *Model) renderHeader() string {
	progress := float64(m.getCompletedCount()) / float64(m.getTotalCount())
	progressPercent := int(progress * 100)

	// Use the existing renderProgressBar function with a reasonable width
	progressBar := m.renderProgressBar(m.getCompletedCount(), m.getTotalCount(), progressBarWidth)
	progressText := fmt.Sprintf("%d/%d (%d%%)", m.getCompletedCount(), m.getTotalCount(), progressPercent)

	header := fmt.Sprintf(`%s

📊 Progress: %s %s`,
		headerStyle.Render("🚀 GoForGo - Interactive Go Tutorial"),
		progressBar,
		progressBarStyle.Render(progressText))

	return header
}

// renderExerciseInfo shows current exercise details
func (m *Model) renderExerciseInfo() string {
	ex := m.currentExercise

	difficulty := ex.GetDifficultyString()
	base := filepath.Base(ex.FilePath)
	category := filepath.Base(filepath.Dir(ex.FilePath))
	categoryStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#161b22")).
		Foreground(lipgloss.Color("#58a6ff")).
		Bold(true).
		Padding(0, 0, 0, 1)
	baseStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#161b22")).
		Foreground(lipgloss.Color("#f0883e")).
		Bold(true).
		Padding(0, 1, 0, 0)
	filePath := categoryStyle.Render(category+"/") + baseStyle.Render(base)

	info := fmt.Sprintf(`📝 Current Exercise: %s
📁 File: %s
%s %s

📖 %s`,
		titleStyle.Render(ex.Description.Title),
		filePath,
		difficulty,
		statusStyle.Render(fmt.Sprintf("(Exercise %d of %d)", m.currentIndex+1, m.getTotalCount())),
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
	// Auto-advance crossfade success screen
	if m.showingSuccess {
		crossfadeStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#3fb950")).
			Bold(true).
			Align(lipgloss.Center)

		return crossfadeStyle.Render("✅ Exercise completed! Moving on...")
	}

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
			if m.autoAdvance {
				result.WriteString(statusStyle.Render("Auto-advancing to the next exercise..."))
			} else {
				result.WriteString(statusStyle.Render("Press 'n' for the next exercise."))
			}
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
	autoAdvanceLabel := "[a] auto-adv"
	if m.autoAdvance {
		autoAdvanceLabel = "[a] auto-adv:ON"
	}

	skipTodoLabel := "[t] skip-todo"
	if m.skipTodoCheck {
		skipTodoLabel = "[t] skip-todo:ON"
	}

	shortcuts := []string{
		"[n] next",
		"[p] prev",
		"[h] hint",
		"[l] list",
		"[r] run",
		"[s] output",
		autoAdvanceLabel,
		skipTodoLabel,
		"[q] quit",
	}

	footer := statusStyle.Render("⌨️  " + strings.Join(shortcuts, " • "))

	// Add file watching status
	if m.watcherErr == nil {
		footer += "\n" + statusStyle.Render("👁️  Watching for file changes...")
	} else {
		footer += "\n" + errorStyle.Render("⚠️  File watcher error: "+m.watcherErr.Error())
	}

	if m.updateNotice != "" {
		footer += "\n" + hintStyle.Render("🔔 "+m.updateNotice)
	}

	return footer
}

// renderProgressBar creates a visual progress bar
func (m *Model) renderProgressBar(completed, total, width int) string {
	if total == 0 {
		return strings.Repeat("─", width)
	}

	progress := float64(completed) / float64(total)
	filledWidth := int(progress * float64(width))
	emptyWidth := width - filledWidth

	// Use different characters for visual appeal
	filled := strings.Repeat("█", filledWidth)
	empty := strings.Repeat("░", emptyWidth)

	// Color the progress bar
	filledStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#3fb950"))
	emptyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#8b949e"))

	return fmt.Sprintf("[%s%s]", filledStyle.Render(filled), emptyStyle.Render(empty))
}

// renderSplash shows an animated splash screen
func (m *Model) renderSplash() string {
	// Animated GoForGo logo with different frames
	frames := []string{
		`
   ██████╗  ██████╗ ███████╗ ██████╗ ██████╗  ██████╗  ██████╗ 
  ██╔════╝ ██╔═══██╗██╔════╝██╔═══██╗██╔══██╗██╔════╝ ██╔═══██╗
  ██║  ███╗██║   ██║█████╗  ██║   ██║██████╔╝██║  ███╗██║   ██║
  ██║   ██║██║   ██║██╔══╝  ██║   ██║██╔══██╗██║   ██║██║   ██║
  ╚██████╔╝╚██████╔╝██║     ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝
   ╚═════╝  ╚═════╝ ╚═╝      ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝`,

		`
   ██████▓  ██████▓ ███████▓ ██████▓ ██████▓  ██████▓  ██████▓ 
  ██▔════▓ ██▔═══██▓██▔════▓██▔═══██▓██▔══██▓██▔════▓ ██▔═══██▓
  ██║  ███▓██║   ██▓█████▓  ██║   ██▓██████▔▓██║  ███▓██║   ██▓
  ██║   ██▓██║   ██▓██▔══▓  ██║   ██▓██▔══██▓██║   ██▓██║   ██▓
  ╚██████▔▓╚██████▔▓██║     ╚██████▔▓██║  ██▓╚██████▔▓╚██████▔▓
   ╚═════▓  ╚═════▓ ╚═╝      ╚═════▓ ╚═╝  ╚═▓ ╚═════▓  ╚═════▓`,

		`
   ██████▒  ██████▒ ███████▒ ██████▒ ██████▒  ██████▒  ██████▒ 
  ██▔════▒ ██▔═══██▒██▔════▒██▔═══██▒██▔══██▒██▔════▒ ██▔═══██▒
  ██║  ███▒██║   ██▒█████▒  ██║   ██▒██████▔▒██║  ███▒██║   ██▒
  ██║   ██▒██║   ██▒██▔══▒  ██║   ██▒██▔══██▒██║   ██▒██║   ██▒
  ╚██████▔▒╚██████▔▒██║     ╚██████▔▒██║  ██▒╚██████▔▒╚██████▔▒
   ╚═════▒  ╚═════▒ ╚═╝      ╚═════▒ ╚═╝  ╚═▒ ╚═════▒  ╚═════▒`,
	}

	// Color gradients for animation
	colors := []string{
		"#FF6B6B", // Red
		"#4ECDC4", // Teal
		"#45B7D1", // Blue
		"#96CEB4", // Green
		"#FECA57", // Yellow
		"#FF9FF3", // Pink
		"#54A0FF", // Light Blue
		"#bc8cff", // Purple
	}

	// Select frame and color based on animation frame
	frameIndex := m.splashFrame % len(frames)
	colorIndex := m.splashFrame % len(colors)

	logoStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors[colorIndex])).
		Bold(true).
		Align(lipgloss.Center)

	logo := logoStyle.Render(frames[frameIndex])

	// Animated subtitle with loading dots
	dots := strings.Repeat(".", (m.splashFrame%4)+1)
	loadingText := fmt.Sprintf("Loading your Go learning experience%s", dots)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#d2a8ff")).
		Italic(true).
		Align(lipgloss.Center)

	subtitle := subtitleStyle.Render(loadingText)

	// Create pulsing effect with different opacity
	content := fmt.Sprintf(`%s

%s

🚀 Interactive Go Tutorial Platform 🚀`, logo, subtitle)

	// Center and style the splash consistently with other views
	contentWidth := m.getContentWidth()

	style := lipgloss.NewStyle().
		Width(contentWidth).
		Align(lipgloss.Center).
		Padding(1, 0)

	return style.Render(content)
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

	// Apply consistent border styling
	borderStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#d29922"))
	borderLine := borderStyle.Render(strings.Repeat("═", borderCharWidth))

	borderedContent := fmt.Sprintf(`%s
%s
%s`, borderLine, content, borderLine)

	// Style the content left-aligned
	contentWidth := m.getContentWidth()

	style := lipgloss.NewStyle().
		Width(contentWidth).
		Align(lipgloss.Left).
		Padding(1, 2)

	return style.Render(borderedContent)
}

// renderExerciseList shows a scrollable exercise list with proper navigation
func (m *Model) renderExerciseList() string {
	var content strings.Builder

	// Header with navigation instructions
	header := headerStyle.Render("📚 Exercise List")
	content.WriteString(header)
	content.WriteString("\n")

	// Filter state display
	if m.filterMode {
		filterPrompt := fmt.Sprintf("Filter: %s_", m.filterText)
		content.WriteString(hintStyle.Render(filterPrompt))
		content.WriteString("\n")
		content.WriteString(statusStyle.Render("Type to filter, Enter to apply, Esc to cancel"))
	} else if m.filterText != "" {
		filterStatus := fmt.Sprintf("Filter active: '%s' (press / to modify)", m.filterText)
		content.WriteString(hintStyle.Render(filterStatus))
		content.WriteString("\n")
		content.WriteString(statusStyle.Render("Use ↑↓/jk to navigate, Enter to select, Esc to return, / to filter"))
	} else {
		content.WriteString(statusStyle.Render("Use ↑↓/jk to navigate, Enter to select, Esc to return, / to filter"))
	}
	content.WriteString("\n\n")

	// Get filtered exercises
	filteredExercises := m.getFilteredExercises()

	// Calculate list dimensions
	listHeight := m.listViewHeight
	totalExercises := len(filteredExercises)
	startIndex := m.listScrollOffset
	endIndex := min(startIndex+listHeight, totalExercises)

	// Progress indicator
	var progressText string
	if m.filterText != "" {
		progressText = fmt.Sprintf("Showing %d-%d of %d filtered exercises (total: %d)",
			startIndex+1, endIndex, totalExercises, len(m.exercises))
	} else {
		progressText = fmt.Sprintf("Exercises %d-%d of %d", startIndex+1, endIndex, totalExercises)
	}
	content.WriteString(statusStyle.Render(progressText))
	content.WriteString("\n\n")

	// Calculate max widths for dynamic sizing based on ALL exercises, not just visible ones
	maxWidths := []int{1, 1, 12, 8, 9, 6} // Min widths: selector, #, EXERCISE NAME, CATEGORY, DIFFICULTY, STATUS
	headers := []string{" ", "#", "EXERCISE NAME", "CATEGORY", "DIFFICULTY", "STATUS"}

	// Update max widths based on headers
	for i, header := range headers {
		if len(header) > maxWidths[i] {
			maxWidths[i] = len(header)
		}
	}

	// Calculate column widths based on ALL filtered exercises for consistent sizing
	for i, ex := range filteredExercises {
		// Exercise number (use position in filtered list)
		exerciseNum := fmt.Sprintf("%d", i+1)
		if len(exerciseNum) > maxWidths[1] {
			maxWidths[1] = len(exerciseNum)
		}

		// Exercise name (no current marker needed, will use color)
		exerciseName := ex.Info.Name
		if len(exerciseName) > maxWidths[2] {
			maxWidths[2] = len(exerciseName)
		}

		// Category
		topic := m.getExerciseTopic(ex)
		if len(topic) > maxWidths[3] {
			maxWidths[3] = len(topic)
		}

		// Difficulty (check all possible difficulty strings)
		var difficulty string
		switch ex.Info.Difficulty {
		case 1:
			difficulty = "Beginner"
		case 2:
			difficulty = "Easy"
		case 3:
			difficulty = "Medium"
		case 4:
			difficulty = "Hard"
		case 5:
			difficulty = "Expert"
		default:
			difficulty = "Unknown"
		}
		if len(difficulty) > maxWidths[4] {
			maxWidths[4] = len(difficulty)
		}

		// Status - "Incomplete" is longer than "Complete"
		if len("Incomplete") > maxWidths[5] {
			maxWidths[5] = len("Incomplete")
		}
	}

	// Selection indicator width (1 for arrow)
	if maxWidths[0] < 1 {
		maxWidths[0] = 1
	}

	// Prepare visible row data for rendering
	type rowData struct {
		selection  string
		number     string
		name       string
		category   string
		difficulty string
		status     string
	}

	var rows []rowData
	for i := startIndex; i < endIndex; i++ {
		ex := filteredExercises[i]

		// Selection indicator
		selectionIndicator := " "
		if i == m.listSelectedIndex {
			selectionIndicator = "►"
		}

		// Exercise number (position in filtered list)
		exerciseNum := fmt.Sprintf("%d", i+1)

		// Exercise name (no current marker text, will use color highlighting)
		exerciseName := ex.Info.Name

		// Category
		topic := m.getExerciseTopic(ex)

		// Difficulty
		var difficulty string
		switch ex.Info.Difficulty {
		case 1:
			difficulty = "Beginner"
		case 2:
			difficulty = "Easy"
		case 3:
			difficulty = "Medium"
		case 4:
			difficulty = "Hard"
		case 5:
			difficulty = "Expert"
		default:
			difficulty = "Unknown"
		}

		// Status
		status := "Incomplete"
		if ex.Completed {
			status = "Complete"
		}

		row := rowData{
			selection:  selectionIndicator,
			number:     exerciseNum,
			name:       exerciseName,
			category:   topic,
			difficulty: difficulty,
			status:     status,
		}
		rows = append(rows, row)
	}

	// Add padding to each column width
	for i := range maxWidths {
		maxWidths[i] = int(float64(maxWidths[i]) * columnPaddingFactor)
		if maxWidths[i] < minColumnWidth {
			maxWidths[i] = minColumnWidth
		}
	}

	// Cap total table width to terminal width.
	// NormalBorder adds: 1 left + 5 between columns + 1 right = 7 chars for borders.
	// Each column gets 1 char padding on each side = 6 * 2 = 12 chars.
	// Total overhead = 19 chars.
	borderOverhead := len(maxWidths) + 1 + len(maxWidths)*2 // borders between + edges + padding
	availableForCols := m.width - borderOverhead
	if availableForCols < 40 {
		availableForCols = 40
	}

	totalColWidth := 0
	for _, w := range maxWidths {
		totalColWidth += w
	}

	// Shrink the two widest flexible columns (name and category) if we exceed available space.
	if totalColWidth > availableForCols {
		excess := totalColWidth - availableForCols
		// Shrink name (col 2) first, then category (col 3)
		nameMax := maxWidths[2] - excess
		if nameMax < 15 {
			nameMax = 15
		}
		shrunk := maxWidths[2] - nameMax
		maxWidths[2] = nameMax
		excess -= shrunk
		if excess > 0 {
			catMax := maxWidths[3] - excess
			if catMax < 10 {
				catMax = 10
			}
			maxWidths[3] = catMax
		}
	}

	tableWidth := 0
	for _, w := range maxWidths {
		tableWidth += w
	}
	tableWidth += borderOverhead

	// Create table with dynamic column widths
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#8b949e"))).
		Width(tableWidth).
		StyleFunc(func(row, col int) lipgloss.Style {
			// Header row styling
			if row == 0 {
				return lipgloss.NewStyle().
					Foreground(lipgloss.Color("#c9d1d9")).
					Bold(true).
					Align(lipgloss.Center).
					Width(maxWidths[col])
			}

			// Check if this row is selected
			actualIndex := startIndex + row - 1
			if actualIndex == m.listSelectedIndex {
				return lipgloss.NewStyle().
					Foreground(lipgloss.Color("#d2a8ff")).
					Bold(true).
					Width(maxWidths[col])
			}

			// Check if this row is the current exercise (highlight in green)
			if actualIndex < len(filteredExercises) && filteredExercises[actualIndex] == m.currentExercise {
				return lipgloss.NewStyle().
					Foreground(lipgloss.Color("#3fb950")).
					Bold(true).
					Width(maxWidths[col])
			}

			// Column-specific colors for regular rows
			baseStyle := lipgloss.NewStyle().Width(maxWidths[col])
			switch col {
			case 0: // Selection indicator
				return baseStyle.Foreground(lipgloss.Color("#d2a8ff")) // Purple
			case 1: // Exercise number
				return baseStyle.Foreground(lipgloss.Color("#8b949e")) // Gray
			case 2: // Exercise name
				return baseStyle.Foreground(lipgloss.Color("#c9d1d9")) // Light gray
			case 3: // Category
				return baseStyle.Foreground(lipgloss.Color("#58a6ff")) // Blue
			case 4: // Difficulty
				// Color based on difficulty level
				if row > 0 && row-1 < len(rows) {
					rowData := rows[row-1]
					if strings.Contains(rowData.difficulty, "Beginner") {
						return baseStyle.Foreground(lipgloss.Color("#3fb950")) // Green
					} else if strings.Contains(rowData.difficulty, "Easy") {
						return baseStyle.Foreground(lipgloss.Color("#58a6ff")) // Blue
					} else if strings.Contains(rowData.difficulty, "Medium") {
						return baseStyle.Foreground(lipgloss.Color("#d29922")) // Orange
					} else if strings.Contains(rowData.difficulty, "Hard") {
						return baseStyle.Foreground(lipgloss.Color("#f85149")) // Red
					} else {
						return baseStyle.Foreground(lipgloss.Color("#d2a8ff")) // Purple
					}
				}
				return baseStyle.Foreground(lipgloss.Color("#d29922")) // Default orange
			case 5: // Status
				// Color based on completion status
				if row > 0 && row-1 < len(rows) {
					rowData := rows[row-1]
					if rowData.status == "Complete" {
						return baseStyle.Foreground(lipgloss.Color("#3fb950")) // Green
					} else {
						return baseStyle.Foreground(lipgloss.Color("#f85149")) // Red
					}
				}
				return baseStyle.Foreground(lipgloss.Color("#f85149")) // Default red
			default:
				return baseStyle.Foreground(lipgloss.Color("#c9d1d9")) // Light gray
			}
		})

	// Add headers
	t.Row(headers...)

	// Add exercise rows
	for _, row := range rows {
		t.Row(row.selection, row.number, row.name, row.category, row.difficulty, row.status)
	}

	// Render the table
	content.WriteString(t.Render())
	content.WriteString("\n")

	// Add minimal spacing
	content.WriteString("\n")

	// Handle case where no exercises match filter
	if totalExercises == 0 && m.filterText != "" {
		content.WriteString("\n")
		noResultsMsg := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f85149")).
			Italic(true).
			Render(fmt.Sprintf("No exercises match filter '%s'", m.filterText))
		content.WriteString(noResultsMsg)
		content.WriteString("\n")
		clearFilterMsg := statusStyle.Render("Press Esc to clear filter or / to modify")
		content.WriteString(clearFilterMsg)
	} else if endIndex >= totalExercises && totalExercises > 0 {
		content.WriteString("\n")
		endIndicator := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8b949e")).
			Italic(true).
			Render("── End of exercises ──")
		content.WriteString(endIndicator)
	}

	// Footer with additional controls
	content.WriteString("\n\n")
	var footerText string
	if m.filterMode {
		footerText = "Filter mode: Type to search  Enter=apply  Esc=cancel  Backspace=delete"
	} else if m.filterText != "" {
		footerText = "Navigation: ↑↓/jk=move  Enter=select  /=filter  Esc=clear filter or back"
	} else {
		footerText = "Navigation: {n}j/k=move  gg/G=top/end  H/M/L=screen  Ctrl+u/d=halfpage  /=filter  r=sync  Esc=back"
	}
	content.WriteString(statusStyle.Render(footerText))

	// Apply consistent border styling using table width for alignment
	listContent := content.String()
	borderWidth := tableWidth
	if borderWidth < borderCharWidth {
		borderWidth = borderCharWidth
	}
	listBorderStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#bc8cff"))
	borderLine := listBorderStyle.Render(strings.Repeat("═", borderWidth))

	borderedContent := fmt.Sprintf(`%s
%s
%s`, borderLine, listContent, borderLine)

	// Use the full terminal width for the list view so the table isn't constrained
	listWidth := m.width
	if listWidth < minContentWidth {
		listWidth = minContentWidth
	}

	style := lipgloss.NewStyle().
		Width(listWidth).
		Padding(1, 0)

	return style.Render(borderedContent)
}

// getExerciseTopic extracts a topic tag from the exercise
func (m *Model) getExerciseTopic(ex interface{}) string {
	// Find the exercise in our slice to get its properties
	for _, exercise := range m.exercises {
		if exercise == ex {
			// Extract topic from category using dynamic string splitting
			category := exercise.Info.Category
			parts := strings.SplitN(category, "_", 2)
			if len(parts) >= 2 {
				return parts[1]
			}
			// Fallback: return the whole category if no underscore found
			return category
		}
	}
	return "unknown"
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

// renderOutput shows a scrollable view of the exercise output
func (m *Model) renderOutput() string {
	var content strings.Builder

	// Header
	header := headerStyle.Render("📋 Exercise Output")
	content.WriteString(header)
	content.WriteString("\n")

	// Exercise info
	if m.currentExercise != nil {
		exerciseInfo := fmt.Sprintf("Exercise: %s", m.currentExercise.Info.Name)
		content.WriteString(statusStyle.Render(exerciseInfo))
		content.WriteString("\n\n")
	}

	// Check if we have output to display
	if m.lastResult == nil {
		content.WriteString(statusStyle.Render("No output available. Run an exercise first (press 'r')."))
		content.WriteString("\n\n")
		content.WriteString(statusStyle.Render("Press Esc to return"))
	} else {
		// Show result status
		if m.lastResult.Success {
			content.WriteString(successStyle.Render("✅ Exercise completed successfully"))
		} else {
			content.WriteString(errorStyle.Render("❌ Exercise failed"))
		}
		content.WriteString("\n\n")

		// Show error if present
		if m.lastResult.Error != "" {
			content.WriteString(errorStyle.Render("Error:"))
			content.WriteString("\n")
			content.WriteString(codeStyle.Render(m.lastResult.Error))
			content.WriteString("\n\n")
		}

		// Show scrollable output
		if m.lastResult.Output != "" {
			outputLines := strings.Split(m.lastResult.Output, "\n")
			totalLines := len(outputLines)

			// Calculate visible range
			startLine := m.outputScrollPos
			endLine := min(startLine+m.outputViewHeight, totalLines)

			// Show scroll position info
			scrollInfo := fmt.Sprintf("Output (lines %d-%d of %d)",
				startLine+1, endLine, totalLines)
			content.WriteString(statusStyle.Render(scrollInfo))
			content.WriteString("\n\n")

			// Show visible output lines
			if startLine < totalLines {
				visibleLines := outputLines[startLine:endLine]
				outputContent := strings.Join(visibleLines, "\n")

				// Style the output in a code block
				outputStyle := lipgloss.NewStyle().
					Background(lipgloss.Color("#161b22")).
					Foreground(lipgloss.Color("#e6edf3")).
					Padding(1, 2).
					Width(m.width - 20) // Leave some margin

				content.WriteString(outputStyle.Render(outputContent))
			}
		} else {
			content.WriteString(statusStyle.Render("No output produced by the exercise."))
		}

		content.WriteString("\n\n")

		// Scroll indicators
		if m.lastResult.Output != "" {
			outputLines := strings.Split(m.lastResult.Output, "\n")
			totalLines := len(outputLines)
			maxScroll := max(0, totalLines-m.outputViewHeight)

			if m.outputScrollPos > 0 {
				content.WriteString(statusStyle.Render("↑ More content above (scroll up)"))
				content.WriteString("\n")
			}
			if m.outputScrollPos < maxScroll {
				content.WriteString(statusStyle.Render("↓ More content below (scroll down)"))
				content.WriteString("\n")
			}
		}

		// Footer with controls
		footerText := "Navigation: ↑↓/jk=scroll  PgUp/PgDn=page  Home/End=jump  Esc=back"
		content.WriteString(statusStyle.Render(footerText))
	}

	// Apply consistent border styling
	outputContent := content.String()
	borderStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#3fb950"))
	borderLine := borderStyle.Render(strings.Repeat("═", borderCharWidth))

	borderedContent := fmt.Sprintf(`%s
%s
%s`, borderLine, outputContent, borderLine)

	// Style the content left-aligned
	contentWidth := m.getContentWidth()

	style := lipgloss.NewStyle().
		Width(contentWidth).
		Align(lipgloss.Left).
		Padding(1, 2)

	return style.Render(borderedContent)
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
