// basic_model.go - Solution
// Learn Bubble Tea fundamentals: Model, View, Update pattern

package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Define the model - this holds your application state
type model struct {
	counter  int
	quitting bool
}

// Initialize the model with default values
func initialModel() model {
	return model{
		counter:  0,
		quitting: false,
	}
}

// Update handles messages and returns updated model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		case "+":
			m.counter++
		case "-":
			m.counter--
		}
	}
	
	return m, nil
}

// View renders the current state as a string
func (m model) View() string {
	if m.quitting {
		return "Goodbye!\n"
	}
	
	s := fmt.Sprintf("Counter: %d\n\n", m.counter)
	s += "Press + to increment, - to decrement\n"
	s += "Press q to quit\n"
	return s
}

// Init returns initial command (can be nil)
func (m model) Init() tea.Cmd {
	return nil
}

func main() {
	p := tea.NewProgram(initialModel())
	
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}