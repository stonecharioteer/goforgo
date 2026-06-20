package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/stonecharioteer/goforgo/internal/exercise"
	"github.com/stonecharioteer/goforgo/internal/runner"
	"github.com/stonecharioteer/goforgo/internal/watcher"
)

// ViewMode represents the current view state of the TUI.
type ViewMode int

const (
	ViewSplash ViewMode = iota
	ViewWelcome
	ViewMain
	ViewList
	ViewHint
	ViewOutput
)

// Model represents the TUI application state
type Model struct {
	// Exercise management
	exerciseManager *exercise.ExerciseManager
	currentExercise *exercise.Exercise
	currentIndex    int
	exercises       []*exercise.Exercise

	// Execution and validation
	runner           *runner.Runner
	lastResult       *runner.Result
	isRunning        bool
	currentHintLevel int // Track current hint level (0=none, 1=level1, 2=level1+2, 3=all)

	// File watching
	watcher          *watcher.Watcher
	watcherErr       error
	watcherListening bool

	// UI state
	viewMode ViewMode
	width    int
	height   int
	ready    bool

	// List view state for scrollable exercise list
	listSelectedIndex int // Currently selected item in list
	listScrollOffset  int // Scroll offset for list view
	listViewHeight    int // Available height for list items

	// Filter state
	filterMode bool   // Whether we're in filter mode
	filterText string // Current filter text

	// Output view state
	outputScrollPos  int // Current scroll position in output view
	outputViewHeight int // Available height for output content

	// Progress and statistics
	// Counts are now calculated dynamically via exerciseManager methods

	// Vim-style key sequence state
	pendingKey   string // Buffered key for multi-key sequences (e.g., "g", "z")
	pendingCount int    // Numeric prefix for {count}j/{count}k motions

	// Auto-advance mode
	autoAdvance    bool // When true, auto-advance to next exercise on success
	showingSuccess bool // True during the success crossfade screen

	// Skip TODO check mode
	skipTodoCheck bool // When true, TODO comments do not block exercise completion

	// Messages and status
	statusMessage string
	updateNotice  string
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
		viewMode:        ViewSplash,
		splashFrame:     0,
	}
}

// SetUpdateNotice sets a startup update notice shown in the UI footer.
func (m *Model) SetUpdateNotice(notice string) {
	m.updateNotice = notice
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

		// Auto-advance: show success screen then move to next exercise
		if msg.result.Success && m.autoAdvance && m.currentIndex < len(m.exercises)-1 {
			m.showingSuccess = true
			return m, tea.Batch(m.autoAdvanceTick(), m.armFileWatcher())
		}

		return m, m.armFileWatcher()

	case exerciseRunningMsg:
		m.isRunning = true
		m.statusMessage = "Running exercise..."
		return m, nil

	case fileChangedMsg:
		m.watcherListening = false
		if !m.isRunning {
			return m, tea.Batch(m.runCurrentExercise(), m.armFileWatcher())
		}
		return m, m.armFileWatcher()

	case watcherErrorMsg:
		m.watcherListening = false
		m.watcherErr = msg.err
		return m, nil

	case continueWatchingMsg:
		m.watcherListening = false
		// Continue listening for file changes
		return m, m.armFileWatcher()

	case autoAdvanceMsg:
		if m.showingSuccess {
			m.showingSuccess = false
			return m, m.nextExercise()
		}
		return m, nil

	case splashTickMsg:
		if m.viewMode == ViewSplash {
			m.splashFrame++
			if m.splashFrame >= splashFrameCount {
				m.viewMode = ViewWelcome
				return m, nil
			}
			return m, m.splashTick()
		}
		return m, nil

	case syncResultMsg:
		m.exercises = m.exerciseManager.GetExercises()
		m.statusMessage = fmt.Sprintf("Synced: %d/%d complete", msg.completed, msg.total)
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
		if m.viewMode == ViewHint || m.viewMode == ViewList {
			m.viewMode = ViewMain
			return m, nil
		}
		return m, m.nextExercise()

	case "p":
		// Previous exercise
		if m.viewMode == ViewHint || m.viewMode == ViewList {
			m.viewMode = ViewMain
			return m, nil
		}
		return m, m.previousExercise()

	case "h":
		// Show next hint level or hide if at max
		if m.viewMode != ViewHint {
			m.currentHintLevel = 1
			m.viewMode = ViewHint
		} else {
			maxLevel := m.getMaxHintLevel()
			if m.currentHintLevel < maxLevel {
				m.currentHintLevel++
			} else {
				m.viewMode = ViewMain
				m.currentHintLevel = 0
			}
		}
		return m, nil

	case "l":
		// Toggle exercise list
		if m.viewMode != ViewList {
			m.viewMode = ViewList
			m.listSelectedIndex = m.currentIndex
			m.listScrollOffset = 0
			m.listViewHeight = max(m.height-listReservedHeight, minListHeight)
			m.ensureSelectedVisible()
		} else {
			m.viewMode = ViewMain
		}
		return m, nil

	case "a":
		// Toggle auto-advance mode
		if m.viewMode == ViewMain {
			m.autoAdvance = !m.autoAdvance
			if m.autoAdvance {
				m.statusMessage = "Auto-advance: ON"
			} else {
				m.statusMessage = "Auto-advance: OFF"
			}
			return m, nil
		}
		return m, nil

	case "t":
		// Toggle skip-TODO-check mode
		if m.viewMode == ViewMain {
			m.skipTodoCheck = !m.skipTodoCheck
			m.runner.SkipTodoCheck = m.skipTodoCheck
			if m.skipTodoCheck {
				m.statusMessage = "Skip TODO check: ON"
			} else {
				m.statusMessage = "Skip TODO check: OFF"
			}
			return m, m.runCurrentExercise()
		}
		return m, nil

	case "r":
		if m.viewMode == ViewList {
			// Sync all exercises in list view
			m.statusMessage = "Syncing..."
			return m, m.syncExercises()
		}
		// Manually run exercise
		if !m.isRunning {
			return m, m.runCurrentExercise()
		}
		return m, nil

	case "s":
		// Show output view
		if m.viewMode == ViewMain && m.lastResult != nil {
			m.viewMode = ViewOutput
			m.outputScrollPos = 0
			m.outputViewHeight = max(m.height-listReservedHeight, minListHeight)
			return m, nil
		}
		return m, nil

	case "up", "k":
		if m.viewMode == ViewList && !m.filterMode {
			count := m.consumeCount(1)
			return m, m.moveListSelection(-count)
		}
		if m.viewMode == ViewOutput {
			return m, m.scrollOutput(-1)
		}
		return m, nil

	case "down", "j":
		if m.viewMode == ViewList && !m.filterMode {
			count := m.consumeCount(1)
			return m, m.moveListSelection(count)
		}
		if m.viewMode == ViewOutput {
			return m, m.scrollOutput(1)
		}
		return m, nil

	case "ctrl+u":
		// Half-page up (vim style)
		if m.viewMode == ViewList && !m.filterMode {
			return m, m.moveListSelection(-m.listViewHeight / 2)
		}
		if m.viewMode == ViewOutput {
			return m, m.scrollOutput(-m.outputViewHeight / 2)
		}
		return m, nil

	case "ctrl+d":
		// Half-page down (vim style)
		if m.viewMode == ViewList && !m.filterMode {
			return m, m.moveListSelection(m.listViewHeight / 2)
		}
		if m.viewMode == ViewOutput {
			return m, m.scrollOutput(m.outputViewHeight / 2)
		}
		return m, nil

	case "page_up":
		if m.viewMode == ViewList && !m.filterMode {
			return m, m.moveListSelection(-m.listViewHeight)
		}
		if m.viewMode == ViewOutput {
			return m, m.scrollOutput(-m.outputViewHeight)
		}
		return m, nil

	case "page_down":
		if m.viewMode == ViewList && !m.filterMode {
			return m, m.moveListSelection(m.listViewHeight)
		}
		if m.viewMode == ViewOutput {
			return m, m.scrollOutput(m.outputViewHeight)
		}
		return m, nil

	case "home":
		if m.viewMode == ViewList && !m.filterMode {
			m.listSelectedIndex = 0
			m.ensureSelectedVisible()
			return m, nil
		}
		if m.viewMode == ViewOutput {
			m.outputScrollPos = 0
			return m, nil
		}
		return m, nil

	case "end":
		if m.viewMode == ViewList && !m.filterMode {
			filteredExercises := m.getFilteredExercises()
			m.listSelectedIndex = len(filteredExercises) - 1
			m.ensureSelectedVisible()
			return m, nil
		}
		if m.viewMode == ViewOutput {
			return m, m.scrollToBottom()
		}
		return m, nil

	case "G":
		// Go to bottom (vim style), or {count}G to go to line
		if m.viewMode == ViewList && !m.filterMode {
			filteredExercises := m.getFilteredExercises()
			if m.pendingCount > 0 {
				target := m.pendingCount - 1 // 1-indexed to 0-indexed
				m.pendingCount = 0
				if target >= len(filteredExercises) {
					target = len(filteredExercises) - 1
				}
				m.listSelectedIndex = target
			} else {
				m.listSelectedIndex = len(filteredExercises) - 1
			}
			m.ensureSelectedVisible()
			return m, nil
		}
		return m, nil

	case "g":
		// First 'g' in 'gg' sequence - go to top
		if m.viewMode == ViewList && !m.filterMode {
			if m.pendingKey == "g" {
				m.pendingKey = ""
				m.pendingCount = 0
				m.listSelectedIndex = 0
				m.ensureSelectedVisible()
				return m, nil
			}
			m.pendingKey = "g"
			return m, nil
		}
		return m, nil

	case "H":
		// Move to top of visible screen
		if m.viewMode == ViewList && !m.filterMode {
			m.listSelectedIndex = m.listScrollOffset
			return m, nil
		}
		return m, nil

	case "M":
		// Move to middle of visible screen
		if m.viewMode == ViewList && !m.filterMode {
			filteredExercises := m.getFilteredExercises()
			mid := m.listScrollOffset + m.listViewHeight/2
			if mid >= len(filteredExercises) {
				mid = len(filteredExercises) - 1
			}
			m.listSelectedIndex = mid
			return m, nil
		}
		return m, nil

	case "L":
		// Move to bottom of visible screen
		if m.viewMode == ViewList && !m.filterMode {
			filteredExercises := m.getFilteredExercises()
			bottom := m.listScrollOffset + m.listViewHeight - 1
			if bottom >= len(filteredExercises) {
				bottom = len(filteredExercises) - 1
			}
			m.listSelectedIndex = bottom
			return m, nil
		}
		return m, nil

	case "backspace":
		// Handle backspace in filter mode
		if m.filterMode && len(m.filterText) > 0 {
			m.filterText = m.filterText[:len(m.filterText)-1]
			return m, nil
		}
		return m, nil

	case "/":
		// Enter filter mode when in list view
		if m.viewMode == ViewList && !m.filterMode {
			m.filterMode = true
			m.filterText = ""
			return m, nil
		}
		return m, nil

	default:
		key := msg.String()

		// Handle text input in filter mode
		if m.filterMode && len(key) == 1 {
			// Only allow alphanumeric characters, underscore, and space
			if (key >= "a" && key <= "z") || (key >= "A" && key <= "Z") ||
				(key >= "0" && key <= "9") || key == "_" || key == " " {
				m.filterText += key
				return m, nil
			}
		}

		// Handle numeric prefix for vim-style {count} motions in list/output views
		if !m.filterMode && (m.viewMode == ViewList || m.viewMode == ViewOutput) {
			if key >= "0" && key <= "9" && (m.pendingCount > 0 || key != "0") {
				digit := int(key[0] - '0')
				m.pendingCount = m.pendingCount*10 + digit
				return m, nil
			}
		}

		// Clear pending key if unrecognized sequence
		if m.pendingKey != "" {
			m.pendingKey = ""
			m.pendingCount = 0
		}

		return m, nil

	case "enter", "esc":
		if m.viewMode == ViewSplash {
			m.viewMode = ViewWelcome
			return m, nil
		}
		if m.viewMode == ViewWelcome {
			m.viewMode = ViewMain
			return m, nil
		}
		if m.viewMode == ViewList && msg.String() == "enter" {
			if m.filterMode {
				m.filterMode = false
				m.listSelectedIndex = 0
				m.listScrollOffset = 0
				return m, nil
			}
			// Select the highlighted exercise
			filteredExercises := m.getFilteredExercises()
			if m.listSelectedIndex >= 0 && m.listSelectedIndex < len(filteredExercises) {
				selectedExercise := filteredExercises[m.listSelectedIndex]
				for i, ex := range m.exercises {
					if ex == selectedExercise {
						m.currentIndex = i
						m.currentExercise = ex
						break
					}
				}
				m.currentHintLevel = 0
				m.viewMode = ViewMain
				m.filterMode = false
				m.filterText = ""
				return m, m.runCurrentExercise()
			}
		}
		if msg.String() == "esc" {
			if m.filterMode {
				m.filterMode = false
				m.filterText = ""
				return m, nil
			} else if m.viewMode == ViewOutput {
				m.viewMode = ViewMain
				m.outputScrollPos = 0
				return m, nil
			} else if m.viewMode == ViewList && m.filterText != "" {
				m.filterText = ""
				m.listSelectedIndex = 0
				m.listScrollOffset = 0
				return m, nil
			}
		}
		// Dismiss any overlay view
		m.viewMode = ViewMain
		m.filterMode = false
		m.filterText = ""
		m.currentHintLevel = 0
		return m, nil
	}

	return m, nil
}

// View renders the TUI
func (m *Model) View() string {
	if !m.ready {
		return "Initializing GoForGo..."
	}

	switch m.viewMode {
	case ViewSplash:
		return m.renderSplash()
	case ViewWelcome:
		return m.renderWelcome()
	case ViewList:
		return m.renderExerciseList()
	case ViewHint:
		return m.renderHint()
	case ViewOutput:
		return m.renderOutput()
	default:
		return m.renderMain()
	}
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

type autoAdvanceMsg struct{}

type syncResultMsg struct {
	completed int
	total     int
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
		m.currentHintLevel = 0 // Reset hint level for new exercise
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
		m.currentHintLevel = 0 // Reset hint level for new exercise
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
	return m.armFileWatcher()
}

func (m *Model) armFileWatcher() tea.Cmd {
	if m.watcher == nil {
		// Keep command semantics predictable for update paths before Init() wires the watcher.
		return func() tea.Msg { return nil }
	}
	if m.watcherListening {
		return nil
	}
	m.watcherListening = true
	return m.waitForFileChange(m.watcher)
}

func (m *Model) waitForFileChange(w *watcher.Watcher) tea.Cmd {
	if w == nil {
		return nil
	}

	return func() tea.Msg {
		select {
		case event, ok := <-w.Events():
			if !ok {
				return watcherErrorMsg{err: fmt.Errorf("watcher events channel closed")}
			}
			if m.shouldProcessFileEvent(event) {
				return fileChangedMsg{path: event.Name}
			}
			// Event not relevant, continue listening
			return continueWatchingMsg{}
		case err, ok := <-w.Errors():
			if !ok {
				return watcherErrorMsg{err: fmt.Errorf("watcher errors channel closed")}
			}
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

// consumeCount returns the pending count (or the default if no count was entered) and resets it.
func (m *Model) consumeCount(defaultVal int) int {
	if m.pendingCount > 0 {
		count := m.pendingCount
		m.pendingCount = 0
		return count
	}
	return defaultVal
}

// Styles — GitHub Dark theme palette for readability on dark terminals.
var (
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#bc8cff")).
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(lipgloss.Color("#bc8cff"))

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#e6edf3"))

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#3fb950")).
			Bold(true)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f85149")).
			Bold(true)

	hintStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#d29922")).
			Italic(true)

	codeStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#161b22")).
			Foreground(lipgloss.Color("#e6edf3")).
			Padding(0, 1)

	progressBarStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#bc8cff"))

	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#8b949e")).
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

// autoAdvanceTick creates a command to auto-advance after a delay
func (m *Model) autoAdvanceTick() tea.Cmd {
	return tea.Tick(time.Second*1, func(time.Time) tea.Msg {
		return autoAdvanceMsg{}
	})
}

// splashTick creates a command for splash screen animation
func (m *Model) splashTick() tea.Cmd {
	return tea.Tick(time.Millisecond*splashTickMs, func(time.Time) tea.Msg {
		return splashTickMsg{}
	})
}

// moveListSelection moves the selection in the list view
func (m *Model) moveListSelection(delta int) tea.Cmd {
	filteredExercises := m.getFilteredExercises()
	newIndex := m.listSelectedIndex + delta

	// Clamp to valid range (no wrapping)
	if newIndex < 0 {
		newIndex = 0
	} else if newIndex >= len(filteredExercises) {
		newIndex = len(filteredExercises) - 1
	}

	m.listSelectedIndex = newIndex
	m.ensureSelectedVisible()

	return nil
}

// ensureSelectedVisible adjusts scroll offset to keep selected item visible
func (m *Model) ensureSelectedVisible() {
	filteredExercises := m.getFilteredExercises()
	if m.listSelectedIndex < m.listScrollOffset {
		// Selected item is above visible area
		m.listScrollOffset = m.listSelectedIndex
	} else if m.listSelectedIndex >= m.listScrollOffset+m.listViewHeight {
		// Selected item is below visible area
		m.listScrollOffset = m.listSelectedIndex - m.listViewHeight + 1
	}

	// Ensure scroll offset doesn't go negative or exceed filtered exercise count
	if m.listScrollOffset < 0 {
		m.listScrollOffset = 0
	}
	if len(filteredExercises) > 0 && m.listScrollOffset >= len(filteredExercises) {
		m.listScrollOffset = len(filteredExercises) - 1
	}
}

// getFilteredExercises returns exercises filtered by the current filter text
func (m *Model) getFilteredExercises() []*exercise.Exercise {
	if m.filterText == "" {
		return m.exercises
	}

	var filtered []*exercise.Exercise
	filterLower := strings.ToLower(m.filterText)

	for _, ex := range m.exercises {
		// Check exercise name
		if strings.Contains(strings.ToLower(ex.Info.Name), filterLower) {
			filtered = append(filtered, ex)
			continue
		}

		// Check exercise title
		if strings.Contains(strings.ToLower(ex.Description.Title), filterLower) {
			filtered = append(filtered, ex)
			continue
		}

		// Check category
		if strings.Contains(strings.ToLower(ex.Info.Category), filterLower) {
			filtered = append(filtered, ex)
			continue
		}

		// Check difficulty
		difficulty := ex.GetDifficultyString()
		if strings.Contains(strings.ToLower(difficulty), filterLower) {
			filtered = append(filtered, ex)
			continue
		}
	}

	return filtered
}

// scrollOutput scrolls the output view by the given delta
func (m *Model) scrollOutput(delta int) tea.Cmd {
	if m.lastResult == nil {
		return nil
	}

	// Split output into lines for scrolling
	outputLines := strings.Split(m.lastResult.Output, "\n")
	maxScroll := max(0, len(outputLines)-m.outputViewHeight)

	m.outputScrollPos += delta

	// Clamp scroll position
	if m.outputScrollPos < 0 {
		m.outputScrollPos = 0
	} else if m.outputScrollPos > maxScroll {
		m.outputScrollPos = maxScroll
	}

	return nil
}

// syncExercises validates all exercises and updates progress
func (m *Model) syncExercises() tea.Cmd {
	return func() tea.Msg {
		exercises := m.exerciseManager.GetExercises()
		completed := 0

		for _, ex := range exercises {
			result, err := m.runner.RunExercise(ex)
			if err != nil {
				continue
			}
			if result.Success {
				if !ex.Completed {
					m.exerciseManager.MarkExerciseCompleted(ex.Info.Name)
				}
				completed++
			} else {
				if ex.Completed {
					m.exerciseManager.UnmarkExerciseCompleted(ex.Info.Name)
				}
			}
		}

		return syncResultMsg{completed: completed, total: len(exercises)}
	}
}

// scrollToBottom scrolls the output view to the bottom
func (m *Model) scrollToBottom() tea.Cmd {
	if m.lastResult == nil {
		return nil
	}

	outputLines := strings.Split(m.lastResult.Output, "\n")
	maxScroll := max(0, len(outputLines)-m.outputViewHeight)
	m.outputScrollPos = maxScroll

	return nil
}
