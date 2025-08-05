package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
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

	// Stats section with progress - use cached completed count
	progressBar := m.renderProgressBar(m.completedCount, m.getTotalCount(), 30)
	
	statsStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#10B981")).
		Bold(true)
	
	statsText := statsStyle.Render(fmt.Sprintf("📊 Progress: %s %d/%d exercises completed", progressBar, m.completedCount, m.getTotalCount()))

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
	progress := float64(m.completedCount) / float64(m.getTotalCount())
	progressPercent := int(progress * 100)

	// Use the existing renderProgressBar function with a reasonable width
	progressBar := m.renderProgressBar(m.completedCount, m.getTotalCount(), 30)
	progressText := fmt.Sprintf("%d/%d (%d%%)", m.completedCount, m.getTotalCount(), progressPercent)

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
		footer += "\n" + errorStyle.Render("⚠️  File watcher error: " + m.watcherErr.Error())
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

	// Render visible exercises
	for i := startIndex; i < endIndex; i++ {
		ex := m.exercises[i]
		
		// Status icon
		status := "❌"
		statusColor := "#EF4444" // Red for incomplete
		if ex.Completed {
			status = "✅"
			statusColor = "#10B981" // Green for complete
		}
		
		// Selection indicator and styling
		var prefix string
		var lineStyle lipgloss.Style
		
		if i == m.listSelectedIndex {
			prefix = " ► "
			// Selected item - bold and italicized with highlight color
			lineStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#7C3AED")).
				Bold(true).
				Italic(true).
				Padding(0, 1)
		} else {
			prefix = "   "
			// Normal item - regular styling
			lineStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#1F2937")).
				Padding(0, 1)
		}
		
		// Current exercise indicator
		currentMarker := ""
		if ex == m.currentExercise {
			currentMarker = " ← current"
		}
		
		// Extract topic from exercise name or category
		topic := m.getExerciseTopic(ex)
		
		// Simplified difficulty (no stars)
		difficulty := m.getSimpleDifficulty(ex.GetDifficultyString())
		difficultyStyle := m.getDifficultyStyle(ex.GetDifficultyString())
		
		// Exercise number (1-based)
		exerciseNum := fmt.Sprintf("%3d.", i+1)
		exerciseNumStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#6B7280"))
		
		// Topic with simple styling
		topicStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#6B7280"))
		
		// Compact format to fit on one line - adjust spacing
		exerciseLine := fmt.Sprintf("%s %s %s %-30s %-12s %s%s", 
			prefix,
			exerciseNumStyle.Render(exerciseNum),
			lipgloss.NewStyle().Foreground(lipgloss.Color(statusColor)).Render(status),
			ex.Info.Name, // Exercise name with padding
			topicStyle.Render(topic), // Topic with padding
			difficultyStyle.Render(difficulty),
			lipgloss.NewStyle().Foreground(lipgloss.Color("#6B7280")).Render(currentMarker))
		
		content.WriteString(lineStyle.Render(exerciseLine))
		content.WriteString("\n")
	}

	// Add spacing if list is shorter than available height
	for i := endIndex - startIndex; i < listHeight; i++ {
		content.WriteString("\n")
	}

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
	// Access the exercise's category and info through the model's exercises slice
	// Since we can't directly access ex.Info.Category due to import removal,
	// we'll extract topic from the exercise name patterns
	
	// Find the exercise in our slice to get its properties
	for _, exercise := range m.exercises {
		if exercise == ex {
			// Extract topic from category
			category := exercise.Info.Category
			switch {
			case strings.HasPrefix(category, "01_"):
				return "basics"
			case strings.HasPrefix(category, "02_"):
				return "variables"
			case strings.HasPrefix(category, "03_"):
				return "functions"
			case strings.HasPrefix(category, "04_"):
				return "control"
			case strings.HasPrefix(category, "05_"):
				return "arrays"
			case strings.HasPrefix(category, "06_"):
				return "slices"
			case strings.HasPrefix(category, "07_"):
				return "maps"
			case strings.HasPrefix(category, "08_"):
				return "structs"
			case strings.HasPrefix(category, "09_"):
				return "interfaces"
			case strings.HasPrefix(category, "10_"):
				return "errors"
			case strings.HasPrefix(category, "11_"):
				return "concurrency"
			case strings.HasPrefix(category, "12_"):
				return "generics"
			case strings.HasPrefix(category, "13_"):
				return "testing"
			case strings.HasPrefix(category, "14_"):
				return "stdlib"
			case strings.HasPrefix(category, "15_"):
				return "json"
			case strings.HasPrefix(category, "16_"):
				return "http"
			case strings.HasPrefix(category, "17_"):
				return "files"
			case strings.HasPrefix(category, "18_"):
				return "regex"
			case strings.HasPrefix(category, "19_"):
				return "reflection"
			case strings.HasPrefix(category, "20_"):
				return "advanced"
			case strings.HasPrefix(category, "21_"):
				return "crypto"
			case strings.HasPrefix(category, "22_"):
				return "networking"
			case strings.HasPrefix(category, "23_"):
				return "encoding"
			case strings.HasPrefix(category, "24_"):
				return "io"
			case strings.HasPrefix(category, "25_"):
				return "paths"
			case strings.HasPrefix(category, "26_"):
				return "os"
			case strings.HasPrefix(category, "27_"):
				return "math"
			case strings.HasPrefix(category, "28_"):
				return "sorting"
			case strings.HasPrefix(category, "29_"):
				return "data-struct"
			case strings.HasPrefix(category, "30_"):
				return "algorithms"
			case strings.HasPrefix(category, "31_"):
				return "web"
			default:
				return "misc"
			}
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