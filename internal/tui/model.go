package tui

import (
	"strings"
	"time"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/stonecharioteer/goforgo/internal/exercise"
	"github.com/stonecharioteer/goforgo/internal/runner"
	"github.com/stonecharioteer/goforgo/internal/watcher"
)

// Model represents the TUI application state
type Model struct {
	// Exercise management
	exerciseManager *exercise.ExerciseManager
	currentExercise *exercise.Exercise
	currentIndex    int
	exercises       []*exercise.Exercise

	// Execution and validation
	runner        *runner.Runner
	lastResult    *runner.Result
	isRunning     bool
	showingHint   bool
	showingList   bool
	currentHintLevel int  // Track current hint level (0=none, 1=level1, 2=level1+2, 3=all)
	
	// File watching
	watcher    *watcher.Watcher
	watcherErr error

	// UI state
	width  int
	height int
	ready  bool
	
	// List view state for scrollable exercise list
	listSelectedIndex int // Currently selected item in list
	listScrollOffset  int // Scroll offset for list view
	listViewHeight    int // Available height for list items
	
	// Progress and statistics
	// Counts are now calculated dynamically via exerciseManager methods
	
	// Messages and status
	statusMessage string
	showSplash    bool
	showWelcome   bool
	splashFrame   int
}

// NewModel creates a new TUI model
func NewModel(exerciseManager *exercise.ExerciseManager, runner *runner.Runner) *Model {
	exercises := exerciseManager.GetExercises()
	currentEx := exerciseManager.GetNextExercise()
	currentIndex := 0
	
	// Find the index of the current exercise
	for i, ex := range exercises {
		if currentEx != nil && ex.Info.Name == currentEx.Info.Name {
			currentIndex = i
			break
		}
	}

	// Exercise counts are now handled dynamically by ExerciseManager

	return &Model{
		exerciseManager: exerciseManager,
		currentExercise: currentEx,
		currentIndex:    currentIndex,
		exercises:       exercises,
		runner:          runner,
		showSplash:      true,  // Show splash animation first
		showWelcome:     false, // Then show welcome
		splashFrame:     0,
	}
}

// Init initializes the model
func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		m.runCurrentExercise(),
		m.startFileWatcher(),
		m.splashTick(), // Start splash animation
	)
}

// Update handles messages and state changes
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true
		return m, nil

	case tea.KeyMsg:
		return m.handleKeyPress(msg)

	case exerciseResultMsg:
		m.lastResult = msg.result
		m.isRunning = false
		m.statusMessage = ""
		
		// Mark exercise as completed if successful and not already completed
		if msg.result.Success && m.currentExercise != nil && !m.currentExercise.Completed {
			if err := m.exerciseManager.MarkExerciseCompleted(m.currentExercise.Info.Name); err == nil {
				// Update local completion tracking
				m.currentExercise.Completed = true
				
				// Update exercises list with fresh completion status
				m.exercises = m.exerciseManager.GetExercises()
			}
		}
		
		return m, m.waitForFileChange(m.watcher)

	case exerciseRunningMsg:
		m.isRunning = true
		m.statusMessage = "Running exercise..."
		return m, nil

	case fileChangedMsg:
		if !m.isRunning {
			return m, m.runCurrentExercise()
		}
		return m, nil

	case watcherErrorMsg:
		m.watcherErr = msg.err
		return m, nil

	case continueWatchingMsg:
		// Continue listening for file changes
		if m.watcher != nil {
			return m, m.waitForFileChange(m.watcher)
		}
		return m, nil

	case splashTickMsg:
		if m.showSplash {
			m.splashFrame++
			if m.splashFrame >= 8 {
				// End splash after 8 frames (~2 seconds)
				m.showSplash = false
				m.showWelcome = true
				return m, nil
			}
			return m, m.splashTick()
		}
		return m, nil

	case statusMsg:
		m.statusMessage = msg.message
		return m, nil
	}

	return m, nil
}

// handleKeyPress processes keyboard input
func (m *Model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit

	case "n":
		// Next exercise
		if m.showingHint || m.showingList {
			m.showingHint = false
			m.showingList = false
			return m, nil
		}
		return m, m.nextExercise()

	case "p":
		// Previous exercise
		if m.showingHint || m.showingList {
			m.showingHint = false
			m.showingList = false
			return m, nil
		}
		return m, m.previousExercise()

	case "h":
		// Show next hint level or hide if at max
		if !m.showingHint {
			// Starting to show hints - show level 1
			m.currentHintLevel = 1
			m.showingHint = true
			m.showingList = false
		} else {
			// Already showing hints - advance to next level or hide
			maxLevel := m.getMaxHintLevel()
			if m.currentHintLevel < maxLevel {
				m.currentHintLevel++
			} else {
				// At max level, hide hints and reset
				m.showingHint = false
				m.currentHintLevel = 0
			}
		}
		return m, nil

	case "l":
		// Toggle exercise list
		if !m.showingList {
			// Initialize list view
			m.showingList = true
			m.showingHint = false
			m.listSelectedIndex = m.currentIndex // Start at current exercise
			m.listScrollOffset = 0
			m.listViewHeight = max(m.height-8, 10) // Reserve space for header/footer, min 10 lines
			m.ensureSelectedVisible()
		} else {
			m.showingList = false
		}
		return m, nil

	case "r":
		// Manually run exercise
		if !m.isRunning {
			return m, m.runCurrentExercise()
		}
		return m, nil

	case "up", "k":
		if m.showingList {
			return m, m.moveListSelection(-1)
		}
		return m, nil

	case "down", "j":
		if m.showingList {
			return m, m.moveListSelection(1)
		}
		return m, nil

	case "page_up":
		if m.showingList {
			return m, m.moveListSelection(-m.listViewHeight)
		}
		return m, nil

	case "page_down":
		if m.showingList {
			return m, m.moveListSelection(m.listViewHeight)
		}
		return m, nil

	case "home":
		if m.showingList {
			m.listSelectedIndex = 0
			m.ensureSelectedVisible()
			return m, nil
		}
		return m, nil

	case "end":
		if m.showingList {
			m.listSelectedIndex = len(m.exercises) - 1
			m.ensureSelectedVisible()
			return m, nil
		}
		return m, nil

	case "enter", "esc":
		if m.showSplash {
			// Skip splash animation
			m.showSplash = false
			m.showWelcome = true
			return m, nil
		}
		if m.showWelcome {
			m.showWelcome = false
			return m, nil
		}
		if m.showingList && msg.String() == "enter" {
			// Select the highlighted exercise
			if m.listSelectedIndex >= 0 && m.listSelectedIndex < len(m.exercises) {
				m.currentIndex = m.listSelectedIndex
				m.currentExercise = m.exercises[m.currentIndex]
				m.currentHintLevel = 0 // Reset hint level for new exercise
				m.showingList = false
				return m, m.runCurrentExercise()
			}
		}
		// Dismiss hint or list
		m.showingHint = false
		m.showingList = false
		m.currentHintLevel = 0  // Reset hint level when dismissing
		return m, nil
	}

	return m, nil
}

// View renders the TUI
func (m *Model) View() string {
	if !m.ready {
		return "Initializing GoForGo..."
	}

	if m.showSplash {
		return m.renderSplash()
	}

	if m.showWelcome {
		return m.renderWelcome()
	}

	if m.showingList {
		return m.renderExerciseList()
	}

	if m.showingHint {
		return m.renderHint()
	}

	return m.renderMain()
}

// Custom messages for the tea program
type exerciseResultMsg struct {
	result *runner.Result
}

type exerciseRunningMsg struct{}

type fileChangedMsg struct {
	path string
}

type watcherErrorMsg struct {
	err error
}

type continueWatchingMsg struct{}

type statusMsg struct {
	message string
}

type splashTickMsg struct{}

// Commands
func (m *Model) runCurrentExercise() tea.Cmd {
	if m.currentExercise == nil {
		return nil
	}

	return tea.Batch(
		func() tea.Msg { return exerciseRunningMsg{} },
		func() tea.Msg {
			result, _ := m.runner.RunExercise(m.currentExercise)
			return exerciseResultMsg{result: result}
		},
	)
}

func (m *Model) nextExercise() tea.Cmd {
	if m.currentIndex < len(m.exercises)-1 {
		m.currentIndex++
		m.currentExercise = m.exercises[m.currentIndex]
		m.currentHintLevel = 0  // Reset hint level for new exercise
		return m.runCurrentExercise()
	}
	return func() tea.Msg {
		return statusMsg{message: "You've reached the last exercise!"}
	}
}

func (m *Model) previousExercise() tea.Cmd {
	if m.currentIndex > 0 {
		m.currentIndex--
		m.currentExercise = m.exercises[m.currentIndex]
		m.currentHintLevel = 0  // Reset hint level for new exercise
		return m.runCurrentExercise()
	}
	return func() tea.Msg {
		return statusMsg{message: "You're at the first exercise!"}
	}
}

func (m *Model) startFileWatcher() tea.Cmd {
	w, err := watcher.NewWatcher()
	if err != nil {
		return func() tea.Msg {
			return watcherErrorMsg{err: err}
		}
	}

	m.watcher = w

	// Watch the exercises directory recursively
	exercisesDir := m.exerciseManager.ExercisesPath
	if err := w.WatchRecursive(exercisesDir); err != nil {
		return func() tea.Msg {
			return watcherErrorMsg{err: err}
		}
	}

	// Start watching for file changes
	return m.waitForFileChange(w)
}

func (m *Model) waitForFileChange(w *watcher.Watcher) tea.Cmd {
	return func() tea.Msg {
		select {
		case event := <-w.Events():
			if m.shouldProcessFileEvent(event) {
				return fileChangedMsg{path: event.Name}
			}
			// Event not relevant, continue listening
			return continueWatchingMsg{}
		case err := <-w.Errors():
			return watcherErrorMsg{err: err}
		}
	}
}

func (m *Model) shouldProcessFileEvent(event watcher.Event) bool {
	// Many editors use atomic writes (create, rename), so we watch for more than just Write events.
	isModification := event.IsWrite() || event.IsCreate() || event.IsRename()
	if !isModification {
		return false
	}

	if !strings.HasSuffix(event.Name, ".go") {
		return false
	}

	if m.currentExercise == nil {
		return false
	}

	// Check if it's the current exercise file
	return strings.Contains(event.Name, m.currentExercise.Info.Name)
}

// Styles
var (
	headerStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7C3AED")).
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderForeground(lipgloss.Color("#7C3AED"))

	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#1F2937"))

	successStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#10B981")).
		Bold(true)

	errorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#EF4444")).
		Bold(true)

	hintStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F59E0B")).
		Italic(true)

	codeStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("#F3F4F6")).
		Foreground(lipgloss.Color("#1F2937")).
		Padding(0, 1)

	progressBarStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#7C3AED"))

	statusStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6B7280")).
		Italic(true)
)

// getTotalCount returns the total number of exercises (dynamic)
func (m *Model) getTotalCount() int {
	return m.exerciseManager.GetTotalExerciseCount()
}

// getCompletedCount returns the number of completed exercises (dynamic)
func (m *Model) getCompletedCount() int {
	return m.exerciseManager.GetCompletedExerciseCount()
}

// getMaxHintLevel returns the maximum hint level available for the current exercise
func (m *Model) getMaxHintLevel() int {
	if m.currentExercise == nil {
		return 0
	}
	
	maxLevel := 0
	if m.currentExercise.Hints.Level1 != "" {
		maxLevel = 1
	}
	if m.currentExercise.Hints.Level2 != "" {
		maxLevel = 2
	}
	if m.currentExercise.Hints.Level3 != "" {
		maxLevel = 3
	}
	
	return maxLevel
}

// splashTick creates a command for splash screen animation
func (m *Model) splashTick() tea.Cmd {
	return tea.Tick(time.Millisecond*250, func(time.Time) tea.Msg {
		return splashTickMsg{}
	})
}

// moveListSelection moves the selection in the list view
func (m *Model) moveListSelection(delta int) tea.Cmd {
	newIndex := m.listSelectedIndex + delta
	
	// Clamp to valid range (no wrapping)
	if newIndex < 0 {
		newIndex = 0
	} else if newIndex >= len(m.exercises) {
		newIndex = len(m.exercises) - 1
	}
	
	m.listSelectedIndex = newIndex
	m.ensureSelectedVisible()
	
	return nil
}

// ensureSelectedVisible adjusts scroll offset to keep selected item visible
func (m *Model) ensureSelectedVisible() {
	if m.listSelectedIndex < m.listScrollOffset {
		// Selected item is above visible area
		m.listScrollOffset = m.listSelectedIndex
	} else if m.listSelectedIndex >= m.listScrollOffset+m.listViewHeight {
		// Selected item is below visible area
		m.listScrollOffset = m.listSelectedIndex - m.listViewHeight + 1
	}
	
	// Ensure scroll offset doesn't go negative
	if m.listScrollOffset < 0 {
		m.listScrollOffset = 0
	}
}