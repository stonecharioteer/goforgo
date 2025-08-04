package watcher

import (
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

// Event represents a file system event
type Event struct {
	Name string
	Op   fsnotify.Op
}

// IsWrite returns true if the event is a write operation
func (e Event) IsWrite() bool {
	return e.Op&fsnotify.Write == fsnotify.Write
}

// IsCreate returns true if the event is a create operation
func (e Event) IsCreate() bool {
	return e.Op&fsnotify.Create == fsnotify.Create
}

// IsRemove returns true if the event is a remove operation
func (e Event) IsRemove() bool {
	return e.Op&fsnotify.Remove == fsnotify.Remove
}

// IsRename returns true if the event is a rename operation
func (e Event) IsRename() bool {
	return e.Op&fsnotify.Rename == fsnotify.Rename
}

// IsChmod returns true if the event is a chmod operation
func (e Event) IsChmod() bool {
	return e.Op&fsnotify.Chmod == fsnotify.Chmod
}

// Watcher wraps fsnotify.Watcher with our custom Event type
type Watcher struct {
	watcher *fsnotify.Watcher
	events  chan Event
	errors  chan error
}

// NewWatcher creates a new file system watcher
func NewWatcher() (*Watcher, error) {
	fsWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	w := &Watcher{
		watcher: fsWatcher,
		events:  make(chan Event, 100), // Buffer events to prevent blocking
		errors:  make(chan error, 10),
	}

	// Start event processing goroutine
	go w.processEvents()

	return w, nil
}

// Add starts watching the specified file or directory
func (w *Watcher) Add(path string) error {
	// Resolve absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	return w.watcher.Add(absPath)
}

// Remove stops watching the specified file or directory
func (w *Watcher) Remove(path string) error {
	return w.watcher.Remove(path)
}

// Events returns the events channel
func (w *Watcher) Events() <-chan Event {
	return w.events
}

// Errors returns the errors channel
func (w *Watcher) Errors() <-chan error {
	return w.errors
}

// Close stops the watcher and closes all channels
func (w *Watcher) Close() error {
	if w.watcher != nil {
		err := w.watcher.Close()
		close(w.events)
		close(w.errors)
		return err
	}
	return nil
}

// processEvents converts fsnotify events to our custom Event type
func (w *Watcher) processEvents() {
	for {
		select {
		case event, ok := <-w.watcher.Events:
			if !ok {
				return // Watcher closed
			}
			w.events <- Event{
				Name: event.Name,
				Op:   event.Op,
			}
		case err, ok := <-w.watcher.Errors:
			if !ok {
				return // Watcher closed
			}
			w.errors <- err
		}
	}
}

// WatchRecursive adds a directory and all its subdirectories to the watcher
func (w *Watcher) WatchRecursive(root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if info.IsDir() {
			return w.Add(path)
		}
		
		return nil
	})
}