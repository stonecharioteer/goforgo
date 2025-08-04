package tui

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/charmbracelet/bubbletea"
	"github.com/stonecharioteer/goforgo/internal/exercise"
	"github.com/stonecharioteer/goforgo/internal/runner"
)

// setupTestEnvironment creates a test environment with exercises
func setupTestEnvironment(t *testing.T) (string, *exercise.ExerciseManager, *runner.Runner) {
	t.Helper()

	// Create temporary directory
	tmpDir := t.TempDir()

	// Create exercise structure
	exerciseDir := filepath.Join(tmpDir, "exercises", "01_basics")
	if err := os.MkdirAll(exerciseDir, 0755); err != nil {
		t.Fatalf("Failed to create exercise directory: %v", err)
	}

	// Create test exercise file
	exerciseContent := `package main

import "fmt"

// TODO: Fix this to print "Hello, GoForGo!"
func main() {
	fmt.Println("Hello, World!")
}
`
	exercisePath := filepath.Join(exerciseDir, "hello.go")
	if err := os.WriteFile(exercisePath, []byte(exerciseContent), 0644); err != nil {
		t.Fatalf("Failed to create exercise file: %v", err)
	}

	// Create test exercise metadata
	metadataContent := `[exercise]
name = "hello"
category = "01_basics"
difficulty = 1

[description]
title = "Hello GoForGo"
summary = "Test exercise"
learning_objectives = ["Test Go syntax"]

[validation]
mode = "run"
timeout = "10s" 

[hints]
level_1 = "Test hint 1"
level_2 = "Test hint 2"
level_3 = "Test hint 3"
`
	metadataPath := filepath.Join(exerciseDir, "hello.toml")
	if err := os.WriteFile(metadataPath, []byte(metadataContent), 0644); err != nil {
		t.Fatalf("Failed to create metadata file: %v", err)
	}

	// Initialize managers
	em := exercise.NewExerciseManager(tmpDir)
	if err := em.LoadExercises(); err != nil {
		t.Fatalf("Failed to load exercises: %v", err)
	}

	r := runner.NewRunner(tmpDir)

	return tmpDir, em, r
}

func TestModelInitialization(t *testing.T) {
	_, em, r := setupTestEnvironment(t)

	model := NewModel(em, r)

	// Test initial state
	if model == nil {
		t.Fatal("NewModel returned nil")
	}

	if model.exerciseManager != em {
		t.Error("Exercise manager not set correctly")
	}

	if model.runner != r {
		t.Error("Runner not set correctly")
	}

	if model.currentExercise == nil {
		t.Error("Current exercise should be set")
	}

	if model.totalCount != 1 {
		t.Errorf("Expected total count 1, got %d", model.totalCount)
	}

	if !model.showWelcome {
		t.Error("Should show welcome screen initially")
	}
}

// TestModelUpdate tests the Update function with table-driven tests (following best practices)
func TestModelUpdate(t *testing.T) {
	_, em, r := setupTestEnvironment(t)

	tests := []struct {
		name        string
		msg         tea.Msg
		setupModel  func(*Model)
		wantState   func(*Model) bool
		wantCmd     bool
		description string
	}{
		{
			name: "quit on q",
			msg:  tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("q")},
			setupModel: func(m *Model) {
				m.ready = true
			},
			wantState: func(m *Model) bool { return true }, // Any state is fine for quit
			wantCmd:   true, // Should return tea.Quit command
		},
		{
			name: "quit on ctrl+c",
			msg:  tea.KeyMsg{Type: tea.KeyCtrlC},
			setupModel: func(m *Model) {
				m.ready = true
			},
			wantState: func(m *Model) bool { return true },
			wantCmd:   true,
		},
		{
			name: "toggle hint on h",
			msg:  tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("h")},
			setupModel: func(m *Model) {
				m.ready = true
				m.showingHint = false
			},
			wantState: func(m *Model) bool { return m.showingHint },
			wantCmd:   false,
		},
		{
			name: "toggle hint off on h when showing",
			msg:  tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("h")},
			setupModel: func(m *Model) {
				m.ready = true
				m.showingHint = true
			},
			wantState: func(m *Model) bool { return !m.showingHint },
			wantCmd:   false,
		},
		{
			name: "show list on l",
			msg:  tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("l")},
			setupModel: func(m *Model) {
				m.ready = true
				m.showingList = false
			},
			wantState: func(m *Model) bool { return m.showingList },
			wantCmd:   false,
		},
		{
			name: "dismiss welcome on enter",
			msg:  tea.KeyMsg{Type: tea.KeyEnter},
			setupModel: func(m *Model) {
				m.ready = true
				m.showWelcome = true
			},
			wantState: func(m *Model) bool { return !m.showWelcome },
			wantCmd:   false,
		},
		{
			name: "dismiss hint on enter",
			msg:  tea.KeyMsg{Type: tea.KeyEnter},
			setupModel: func(m *Model) {
				m.ready = true
				m.showingHint = true
				m.showWelcome = false
			},
			wantState: func(m *Model) bool { return !m.showingHint },
			wantCmd:   false,
		},
		{
			name: "dismiss list on esc",
			msg:  tea.KeyMsg{Type: tea.KeyEsc},
			setupModel: func(m *Model) {
				m.ready = true
				m.showingList = true
				m.showWelcome = false
			},
			wantState: func(m *Model) bool { return !m.showingList },
			wantCmd:   false,
		},
		{
			name: "manual run on r",
			msg:  tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("r")},
			setupModel: func(m *Model) {
				m.ready = true
				m.isRunning = false
			},
			wantState: func(m *Model) bool { return true }, // State changes happen in command
			wantCmd:   true, // Should trigger run command
		},
		{
			name: "no run when already running",
			msg:  tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("r")},
			setupModel: func(m *Model) {
				m.ready = true
				m.isRunning = true
			},
			wantState: func(m *Model) bool { return m.isRunning },
			wantCmd:   false,
		},
		{
			name: "window resize updates dimensions",
			msg:  tea.WindowSizeMsg{Width: 100, Height: 50},
			setupModel: func(m *Model) {
				m.ready = false
			},
			wantState: func(m *Model) bool {
				return m.width == 100 && m.height == 50 && m.ready
			},
			wantCmd: false,
		},
		{
			name: "exercise result updates state",
			msg: exerciseResultMsg{
				result: &runner.Result{Success: true, Output: "Test output"},
			},
			setupModel: func(m *Model) {
				m.ready = true
				m.isRunning = true
			},
			wantState: func(m *Model) bool {
				return !m.isRunning && m.lastResult != nil && m.lastResult.Success
			},
			wantCmd: false,
		},
		{
			name: "exercise running sets state",
			msg:  exerciseRunningMsg{},
			setupModel: func(m *Model) {
				m.ready = true
				m.isRunning = false
			},
			wantState: func(m *Model) bool {
				return m.isRunning && m.statusMessage == "Running exercise..."
			},
			wantCmd: false,
		},
		{
			name: "status message updates",
			msg:  statusMsg{message: "Test status"},
			setupModel: func(m *Model) {
				m.ready = true
			},
			wantState: func(m *Model) bool {
				return m.statusMessage == "Test status"
			},
			wantCmd: false,
		},
		{
			name: "watcher error sets error",
			msg:  watcherErrorMsg{err: &os.PathError{Op: "test", Path: "/test"}},
			setupModel: func(m *Model) {
				m.ready = true
			},
			wantState: func(m *Model) bool {
				return m.watcherErr != nil
			},
			wantCmd: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create fresh model for each test
			model := NewModel(em, r)
			if tt.setupModel != nil {
				tt.setupModel(model)
			}

			// Run Update
			updatedModel, cmd := model.Update(tt.msg)
			finalModel := updatedModel.(*Model)

			// Check state
			if !tt.wantState(finalModel) {
				t.Errorf("State check failed for test '%s'", tt.name)
			}

			// Check command presence
			hasCmd := cmd != nil
			if hasCmd != tt.wantCmd {
				t.Errorf("Expected command %v, got command present: %v", tt.wantCmd, hasCmd)
			}
		})
	}
}

func TestModelWindowResize(t *testing.T) {
	_, em, r := setupTestEnvironment(t)
	model := NewModel(em, r)

	// Test window resize
	resizeMsg := tea.WindowSizeMsg{Width: 100, Height: 50}
	updatedModel, _ := model.Update(resizeMsg)
	finalModel := updatedModel.(*Model)

	if finalModel.width != 100 {
		t.Errorf("Expected width 100, got %d", finalModel.width)
	}

	if finalModel.height != 50 {
		t.Errorf("Expected height 50, got %d", finalModel.height)
	}

	if !finalModel.ready {
		t.Error("Model should be ready after window resize")
	}
}

func TestModelExerciseNavigation(t *testing.T) {
	// Create test environment with multiple exercises
	tmpDir := t.TempDir()

	// Create multiple exercises
	exercises := []struct {
		name     string
		category string
	}{
		{"hello", "01_basics"},
		{"variables", "02_variables"},
		{"functions", "03_functions"},
	}

	for _, ex := range exercises {
		exerciseDir := filepath.Join(tmpDir, "exercises", ex.category)
		if err := os.MkdirAll(exerciseDir, 0755); err != nil {
			t.Fatalf("Failed to create exercise directory: %v", err)
		}

		// Create Go file
		exerciseContent := `package main
import "fmt"
func main() { fmt.Println("test") }
`
		exercisePath := filepath.Join(exerciseDir, ex.name+".go")
		if err := os.WriteFile(exercisePath, []byte(exerciseContent), 0644); err != nil {
			t.Fatalf("Failed to create exercise file: %v", err)
		}

		// Create metadata
		metadataContent := `[exercise]
name = "` + ex.name + `"
category = "` + ex.category + `"
difficulty = 1

[description]
title = "Test Exercise"
summary = "Test exercise"

[validation]
mode = "run"

[hints]
level_1 = "Test hint"
`
		metadataPath := filepath.Join(exerciseDir, ex.name+".toml")
		if err := os.WriteFile(metadataPath, []byte(metadataContent), 0644); err != nil {
			t.Fatalf("Failed to create metadata file: %v", err)
		}
	}

	em := exercise.NewExerciseManager(tmpDir)
	if err := em.LoadExercises(); err != nil {
		t.Fatalf("Failed to load exercises: %v", err)
	}

	r := runner.NewRunner(tmpDir)
	model := NewModel(em, r)
	model.ready = true

	// Test next exercise
	initialIndex := model.currentIndex
	keyMsg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("n")}
	updatedModel, _ := model.Update(keyMsg)
	finalModel := updatedModel.(*Model)

	if len(model.exercises) > 1 && finalModel.currentIndex <= initialIndex {
		t.Error("Next exercise navigation should increase current index")
	}
}

func TestModelExerciseResults(t *testing.T) {
	_, em, r := setupTestEnvironment(t)
	model := NewModel(em, r)
	model.ready = true

	// Test successful result
	successResult := &runner.Result{
		Success:  true,
		ExitCode: 0,
		Output:   "Success!",
	}

	resultMsg := exerciseResultMsg{result: successResult}
	updatedModel, _ := model.Update(resultMsg)
	finalModel := updatedModel.(*Model)

	if finalModel.lastResult != successResult {
		t.Error("Last result should be set")
	}

	if finalModel.isRunning {
		t.Error("Should not be running after result")
	}

	// Test failed result
	failResult := &runner.Result{
		Success:  false,
		ExitCode: 1,
		Output:   "Compilation failed",
		Error:    "syntax error",
	}

	failMsg := exerciseResultMsg{result: failResult}
	updatedModel2, _ := finalModel.Update(failMsg)
	finalModel2 := updatedModel2.(*Model)

	if finalModel2.lastResult != failResult {
		t.Error("Last result should be updated")
	}
}

// TestModelView tests View() function with minimal assertions (following best practices)
func TestModelView(t *testing.T) {
	_, em, r := setupTestEnvironment(t)

	tests := []struct {
		name         string
		setupModel   func(*Model)
		expectString string
	}{
		{
			name: "not ready shows initializing",
			setupModel: func(m *Model) {
				m.ready = false
			},
			expectString: "Initializing GoForGo...",
		},
		{
			name: "welcome screen shows welcome",
			setupModel: func(m *Model) {
				m.ready = true
				m.showWelcome = true
			},
			expectString: "Welcome to...",
		},
		{
			name: "hint view shows hint",
			setupModel: func(m *Model) {
				m.ready = true
				m.showWelcome = false
				m.showingHint = true
			},
			expectString: "Hint",
		},
		{
			name: "list view shows exercise list",
			setupModel: func(m *Model) {
				m.ready = true
				m.showWelcome = false
				m.showingList = true
			},
			expectString: "Exercise List",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			model := NewModel(em, r)
			model.width = 80
			model.height = 24
			tt.setupModel(model)

			view := model.View()
			
			if !strings.Contains(view, tt.expectString) {
				preview := view
				if len(view) > 100 {
					preview = view[:100] + "..."
				}
				t.Errorf("Expected view to contain '%s', got: %s", tt.expectString, preview)
			}
		})
	}
}

func TestModelStatusMessages(t *testing.T) {
	_, em, r := setupTestEnvironment(t)
	model := NewModel(em, r)

	// Test status message
	statusMessage := "Test status message"
	statusMsg := statusMsg{message: statusMessage}

	updatedModel, _ := model.Update(statusMsg)
	finalModel := updatedModel.(*Model)

	if finalModel.statusMessage != statusMessage {
		t.Errorf("Expected status message '%s', got '%s'", statusMessage, finalModel.statusMessage)
	}
}

func TestModelFileWatcherError(t *testing.T) {
	_, em, r := setupTestEnvironment(t)
	model := NewModel(em, r)

	// Test watcher error
	testErr := &os.PathError{Op: "watch", Path: "/test", Err: os.ErrNotExist}
	watcherMsg := watcherErrorMsg{err: testErr}

	updatedModel, _ := model.Update(watcherMsg)
	finalModel := updatedModel.(*Model)

	if finalModel.watcherErr != testErr {
		t.Error("Watcher error should be set")
	}
}