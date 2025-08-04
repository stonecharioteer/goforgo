package tui

import (
	"strings"

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
	
	// File watching
	watcher    *watcher.Watcher
	watcherErr error

	// UI state
	width  int
	height int
	ready  bool
	
	// Progress and statistics
	completedCount int
	totalCount     int
	
	// Messages and status
	statusMessage string
	showWelcome   bool
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

	// Count completed exercises
	completedCount := 0
	for _, ex := range exercises {
		if ex.Completed {
			completedCount++
		}
	}

	return &Model{
		exerciseManager: exerciseManager,
		currentExercise: currentEx,
		currentIndex:    currentIndex,
		exercises:       exercises,
		runner:          runner,
		completedCount:  completedCount,
		totalCount:      len(exercises),
		showWelcome:     true, // Show welcome on first run
	}
}

// Init initializes the model
func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		m.runCurrentExercise(),
		m.startFileWatcher(),
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
		return m, nil

	case exerciseRunningMsg:
		m.isRunning = true
		m.statusMessage = "Running exercise..."
		return m, nil

	case fileChangedMsg:
		var cmd tea.Cmd
		if !m.isRunning {
			cmd = m.runCurrentExercise()
		}
		// Continue listening for more file changes
		continueCmd := m.waitForFileChange(m.watcher)
		if cmd != nil {
			return m, tea.Batch(cmd, continueCmd)
		}
		return m, continueCmd

	case watcherErrorMsg:
		m.watcherErr = msg.err
		return m, nil

	case continueWatchingMsg:
		// Continue listening for file changes
		if m.watcher != nil {
			return m, m.waitForFileChange(m.watcher)
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
		// Toggle hint
		m.showingHint = !m.showingHint
		m.showingList = false
		return m, nil

	case "l":
		// Toggle exercise list
		m.showingList = !m.showingList
		m.showingHint = false
		return m, nil

	case "r":
		// Manually run exercise
		if !m.isRunning {
			return m, m.runCurrentExercise()
		}
		return m, nil

	case "enter", "esc":
		if m.showWelcome {
			m.showWelcome = false
			return m, nil
		}
		// Dismiss hint or list
		m.showingHint = false
		m.showingList = false
		return m, nil
	}

	return m, nil
}

// View renders the TUI
func (m *Model) View() string {
	if !m.ready {
		return "Initializing GoForGo..."
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

	// Watch the exercises directory
	exercisesDir := m.exerciseManager.ExercisesPath
	if err := w.Add(exercisesDir); err != nil {
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
	if !event.IsWrite() {
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