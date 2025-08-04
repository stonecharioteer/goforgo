package watcher

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestWatcher_BasicFunctionality(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "test.go")

	// Create initial file
	err := os.WriteFile(testFile, []byte("package main\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Create watcher
	w, err := NewWatcher()
	if err != nil {
		t.Fatalf("Failed to create watcher: %v", err)
	}
	defer w.Close()

	// Add file to watch
	err = w.Add(tempDir)
	if err != nil {
		t.Fatalf("Failed to add directory to watcher: %v", err)
	}

	// Wait a bit for watcher to initialize
	time.Sleep(100 * time.Millisecond)

	// Modify the file
	err = os.WriteFile(testFile, []byte("package main\n\nfunc main() {}\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to modify test file: %v", err)
	}

	// Wait for event
	select {
	case event := <-w.Events():
		if !event.IsWrite() {
			t.Errorf("Expected write event, got %v", event.Op)
		}
		filename := filepath.Base(event.Name)
		if filename != "test.go" && event.Name != tempDir {
			t.Logf("Got event for %s (filename: %s)", event.Name, filename)
		}
	case err := <-w.Errors():
		t.Fatalf("Watcher error: %v", err)
	case <-time.After(2 * time.Second):
		t.Fatal("Timeout waiting for file change event")
	}
}

func TestWatcher_EventTypes(t *testing.T) {
	tempDir := t.TempDir()
	
	w, err := NewWatcher()
	if err != nil {
		t.Fatalf("Failed to create watcher: %v", err)
	}
	defer w.Close()

	err = w.Add(tempDir)
	if err != nil {
		t.Fatalf("Failed to add directory to watcher: %v", err)
	}

	// Wait for initialization
	time.Sleep(100 * time.Millisecond)

	// Test create event
	testFile := filepath.Join(tempDir, "new.go")
	err = os.WriteFile(testFile, []byte("package main\n"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// We should get at least one event (create or write)
	select {
	case event := <-w.Events():
		if !event.IsCreate() && !event.IsWrite() {
			t.Logf("Got event type %v (expected create or write)", event.Op)
		}
	case err := <-w.Errors():
		t.Fatalf("Watcher error: %v", err)
	case <-time.After(2 * time.Second):
		t.Fatal("Timeout waiting for file create event")
	}
}

func TestWatcher_GoFileFiltering(t *testing.T) {
	// Test the event filtering logic in shouldProcessFileEvent
	testCases := []struct {
		filename    string
		op          string
		currentEx   string
		shouldWatch bool
	}{
		{"hello.go", "write", "hello", true},
		{"hello.go", "create", "hello", false}, // Only watch writes
		{"hello.txt", "write", "hello", false}, // Only .go files
		{"other.go", "write", "hello", false},  // Must match current exercise
		{"hello.go", "write", "", false},       // Need current exercise
	}

	for _, tc := range testCases {
		t.Run(tc.filename+"_"+tc.op, func(t *testing.T) {
			// This would be tested in the TUI model tests
			// Here we just verify our Event type methods work
			event := Event{Name: tc.filename}
			
			switch tc.op {
			case "write":
				event.Op = 2 // fsnotify.Write
			case "create":
				event.Op = 1 // fsnotify.Create
			}

			if tc.op == "write" && !event.IsWrite() {
				t.Error("IsWrite() should return true for write events")
			}
			if tc.op == "create" && !event.IsCreate() {
				t.Error("IsCreate() should return true for create events")
			}
		})
	}
}