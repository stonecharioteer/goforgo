package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

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
		Foreground(lipgloss.Color("#7C3AED")).
		Bold(true).
		Align(lipgloss.Center)

	// Create colorful banner with gradient effect
	banner := logoStyle.Render(logo)

	// Subtitle with gradient
	subtitleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#A855F7")).
		Bold(true).
		Italic(true).
		Align(lipgloss.Center)

	subtitle := subtitleStyle.Render("🚀 Interactive Go Learning Platform 🚀")

	// Stats section with progress - use dynamic completed count
	progressBar := m.renderProgressBar(m.getCompletedCount(), m.getTotalCount(), 30)

	statsStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#10B981")).
		Bold(true)

	statsText := statsStyle.Render(fmt.Sprintf("📊 Progress: %s %d/%d exercises completed", progressBar, m.getCompletedCount(), m.getTotalCount()))

	// Feature highlights with emojis and colors
	featuresStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#3B82F6"))

	features := featuresStyle.Render(`✨ What makes GoForGo special:
   🔥 Real-time feedback as you code
   🧠 Progressive hints that guide your learning  
   📈 Smart progress tracking with auto-skip
   🎯 TODO-driven exercises for flexible learning
   ⚡ Instant file change detection`)

	// Learning topics with nice formatting
	topicsStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F59E0B"))

	topics := topicsStyle.Render(fmt.Sprintf(`📚 %d exercises covering:
   • Go fundamentals & syntax        • Error handling patterns
   • Variables & data types          • Concurrency & goroutines  
   • Functions & methods             • Channels & sync primitives
   • Structs & interfaces            • Popular libraries (Gin, GORM)
   • Control flow & loops            • Real-world projects`, m.getTotalCount()))

	// Keyboard shortcuts in a nice box
	shortcutsStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#EC4899")).
		Padding(0, 1).
		Foreground(lipgloss.Color("#EC4899"))

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
			BorderForeground(lipgloss.Color("#10B981")).
			Padding(0, 2).
			Foreground(lipgloss.Color("#10B981")).
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
		Foreground(lipgloss.Color("#FBBF24")).
		Bold(true).
		Blink(true)

	welcomeText += ctaStyle.Render("✨ Press Enter to begin your Go journey! ✨")

	// Add decorative border using text characters
	borderStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#7C3AED"))
	borderLine := borderStyle.Render(strings.Repeat("═", 80))

	welcomeText = fmt.Sprintf(`%s
%s
%s`, borderLine, welcomeText, borderLine)

	// Center and style the entire content
	// Account for border (2 chars) and padding (4 chars) = 6 chars total
	// Add extra margin for safety
	contentWidth := m.width - 10
	if contentWidth < 50 {
		contentWidth = 50 // Minimum readable width
	}
	if contentWidth > 90 {
		contentWidth = 90 // Maximum width to prevent overly wide content
	}

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
	borderStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#7C3AED"))
	borderLine := borderStyle.Render(strings.Repeat("═", 80))

	borderedContent := fmt.Sprintf(`%s
%s
%s`, borderLine, mainContent, borderLine)

	// Center and style the content consistently
	contentWidth := m.width - 10
	if contentWidth < 50 {
		contentWidth = 50 // Minimum readable width
	}
	if contentWidth > 90 {
		contentWidth = 90 // Maximum width to prevent overly wide content
	}

	style := lipgloss.NewStyle().
		Width(contentWidth).
		Align(lipgloss.Center).
		Padding(1, 0)

	return style.Render(borderedContent)
}

// renderHeader shows the progress bar and current status
func (m *Model) renderHeader() string {
	progress := float64(m.getCompletedCount()) / float64(m.getTotalCount())
	progressPercent := int(progress * 100)

	// Use the existing renderProgressBar function with a reasonable width
	progressBar := m.renderProgressBar(m.getCompletedCount(), m.getTotalCount(), 30)
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
	filePath := codeStyle.Render(ex.FilePath)

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
		footer += "\n" + errorStyle.Render("⚠️  File watcher error: "+m.watcherErr.Error())
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
	filledStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#10B981"))
	emptyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#6B7280"))

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
		"#5F27CD", // Purple
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
		Foreground(lipgloss.Color("#A855F7")).
		Italic(true).
		Align(lipgloss.Center)

	subtitle := subtitleStyle.Render(loadingText)

	// Create pulsing effect with different opacity
	content := fmt.Sprintf(`%s

%s

🚀 Interactive Go Tutorial Platform 🚀`, logo, subtitle)

	// Center and style the splash consistently with other views
	contentWidth := m.width - 10
	if contentWidth < 50 {
		contentWidth = 50
	}
	if contentWidth > 90 {
		contentWidth = 90
	}

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
	borderStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#F59E0B"))
	borderLine := borderStyle.Render(strings.Repeat("═", 80))

	borderedContent := fmt.Sprintf(`%s
%s
%s`, borderLine, content, borderLine)

	// Center and style consistently
	contentWidth := m.width - 10
	if contentWidth < 50 {
		contentWidth = 50
	}
	if contentWidth > 90 {
		contentWidth = 90
	}

	style := lipgloss.NewStyle().
		Width(contentWidth).
		Align(lipgloss.Center).
		Padding(1, 0)

	return style.Render(borderedContent)
}

// renderExerciseList shows a scrollable exercise list with proper navigation
func (m *Model) renderExerciseList() string {
	var content strings.Builder

	// Header with navigation instructions
	header := headerStyle.Render("📚 Exercise List")
	content.WriteString(header)
	content.WriteString("\n")

	// Navigation instructions
	navInstructions := statusStyle.Render("Use ↑↓/jk to navigate, Enter to select, Esc to return")
	content.WriteString(navInstructions)
	content.WriteString("\n\n")

	// Calculate list dimensions
	listHeight := m.listViewHeight
	totalExercises := len(m.exercises)
	startIndex := m.listScrollOffset
	endIndex := min(startIndex+listHeight, totalExercises)

	// Progress indicator
	progressText := fmt.Sprintf("Exercises %d-%d of %d", startIndex+1, endIndex, totalExercises)
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

	// Calculate column widths based on ALL exercises for consistent sizing
	for _, ex := range m.exercises {
		// Exercise number (use total count for max width)
		exerciseNum := fmt.Sprintf("%d", len(m.exercises))
		if len(exerciseNum) > maxWidths[1] {
			maxWidths[1] = len(exerciseNum)
		}

		// Exercise name with potential current marker
		exerciseName := ex.Info.Name
		if ex == m.currentExercise {
			exerciseName += " (current)"
		}
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
		ex := m.exercises[i]

		// Selection indicator
		selectionIndicator := " "
		if i == m.listSelectedIndex {
			selectionIndicator = "►"
		}

		// Exercise number
		exerciseNum := fmt.Sprintf("%d", i+1)

		// Exercise name with current marker
		exerciseName := ex.Info.Name
		if ex == m.currentExercise {
			exerciseName += " (current)"
		}

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

	// Add 10% padding to each column width
	for i := range maxWidths {
		maxWidths[i] = int(float64(maxWidths[i]) * 1.1)
		if maxWidths[i] < 3 { // Minimum width
			maxWidths[i] = 3
		}
	}

	// Create table with dynamic column widths
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#6B7280"))).
		Width(maxWidths[0] + maxWidths[1] + maxWidths[2] + maxWidths[3] + maxWidths[4] + maxWidths[5] + 12). // Account for borders and padding
		StyleFunc(func(row, col int) lipgloss.Style {
			// Header row styling
			if row == 0 {
				return lipgloss.NewStyle().
					Foreground(lipgloss.Color("#D1D5DB")).
					Bold(true).
					Align(lipgloss.Center).
					Width(maxWidths[col])
			}

			// Check if this row is selected
			actualIndex := startIndex + row - 1
			if actualIndex == m.listSelectedIndex {
				return lipgloss.NewStyle().
					Foreground(lipgloss.Color("#A855F7")).
					Bold(true).
					Width(maxWidths[col])
			}

			// Column-specific colors for regular rows
			baseStyle := lipgloss.NewStyle().Width(maxWidths[col])
			switch col {
			case 0: // Selection indicator
				return baseStyle.Foreground(lipgloss.Color("#A855F7")) // Purple
			case 1: // Exercise number
				return baseStyle.Foreground(lipgloss.Color("#6B7280")) // Gray
			case 2: // Exercise name
				return baseStyle.Foreground(lipgloss.Color("#E5E7EB")) // Light gray
			case 3: // Category
				return baseStyle.Foreground(lipgloss.Color("#60A5FA")) // Blue
			case 4: // Difficulty
				// Color based on difficulty level
				if row > 0 && row-1 < len(rows) {
					rowData := rows[row-1]
					if strings.Contains(rowData.difficulty, "Beginner") {
						return baseStyle.Foreground(lipgloss.Color("#10B981")) // Green
					} else if strings.Contains(rowData.difficulty, "Easy") {
						return baseStyle.Foreground(lipgloss.Color("#3B82F6")) // Blue
					} else if strings.Contains(rowData.difficulty, "Medium") {
						return baseStyle.Foreground(lipgloss.Color("#F59E0B")) // Orange
					} else if strings.Contains(rowData.difficulty, "Hard") {
						return baseStyle.Foreground(lipgloss.Color("#EF4444")) // Red
					} else {
						return baseStyle.Foreground(lipgloss.Color("#8B5CF6")) // Purple
					}
				}
				return baseStyle.Foreground(lipgloss.Color("#F59E0B")) // Default orange
			case 5: // Status
				// Color based on completion status
				if row > 0 && row-1 < len(rows) {
					rowData := rows[row-1]
					if rowData.status == "Complete" {
						return baseStyle.Foreground(lipgloss.Color("#10B981")) // Green
					} else {
						return baseStyle.Foreground(lipgloss.Color("#EF4444")) // Red
					}
				}
				return baseStyle.Foreground(lipgloss.Color("#EF4444")) // Default red
			default:
				return baseStyle.Foreground(lipgloss.Color("#E5E7EB")) // Light gray
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

	// End-of-list indicator when at bottom
	if endIndex >= totalExercises {
		content.WriteString("\n")
		endIndicator := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6B7280")).
			Italic(true).
			Render("── End of exercises ──")
		content.WriteString(endIndicator)
	}

	// Footer with additional controls
	content.WriteString("\n\n")
	footerText := "Navigation: ↑↓/jk=move  PgUp/PgDn=page  Home/End=jump  Enter=select  Esc=back"
	content.WriteString(statusStyle.Render(footerText))

	// Apply consistent border styling
	listContent := content.String()
	borderStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#7C3AED"))
	borderLine := borderStyle.Render(strings.Repeat("═", 80))

	borderedContent := fmt.Sprintf(`%s
%s
%s`, borderLine, listContent, borderLine)

	// Center and style consistently
	contentWidth := m.width - 10
	if contentWidth < 50 {
		contentWidth = 50
	}
	if contentWidth > 90 {
		contentWidth = 90
	}

	style := lipgloss.NewStyle().
		Width(contentWidth).
		Align(lipgloss.Center).
		Padding(1, 0)

	return style.Render(borderedContent)
}

// getDifficultyStyle returns appropriate styling for difficulty level
func (m *Model) getDifficultyStyle(difficulty string) lipgloss.Style {
	if strings.Contains(difficulty, "Beginner") {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#10B981")) // Green
	} else if strings.Contains(difficulty, "Easy") {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#3B82F6")) // Blue
	} else if strings.Contains(difficulty, "Medium") {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#F59E0B")) // Orange
	} else if strings.Contains(difficulty, "Hard") {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("#EF4444")) // Red
	}
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#6B7280")) // Gray
}

// getSimpleDifficulty extracts just the difficulty level without stars
func (m *Model) getSimpleDifficulty(difficulty string) string {
	if strings.Contains(difficulty, "Beginner") {
		return "Beginner"
	} else if strings.Contains(difficulty, "Easy") {
		return "Easy"
	} else if strings.Contains(difficulty, "Medium") {
		return "Medium"
	} else if strings.Contains(difficulty, "Hard") {
		return "Hard"
	}
	return "Unknown"
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

