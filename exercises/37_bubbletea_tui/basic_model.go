// basic_model.go
// Learn Bubble Tea fundamentals: Model, View, Update pattern
//
// Bubble Tea is a functional framework for building rich terminal applications.
// It uses the Elm Architecture with Model (state), View (render), and Update (handle messages).
// This exercise covers creating a basic counter application.
//
// I AM NOT DONE YET!

package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Define the model - this holds your application state
type model struct {
	// TODO: Add a counter field to track the current value
	// counter ???
	
	// TODO: Add a field to track if we should quit
	// quitting ???
}

// Initialize the model with default values
func initialModel() model {
	// TODO: Return a model with counter starting at 0 and quitting false
	// return model{
	//     counter: ???,
	//     quitting: ???,
	// }
}

// Update handles messages and returns updated model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// TODO: Handle "q", "ctrl+c", "esc" keys to quit
		// case ???, ???, ???:
		//     m.quitting = true
		//     return m, tea.Quit
		
		// TODO: Handle "+" key to increment counter
		// case ???:
		//     m.counter++
		
		// TODO: Handle "-" key to decrement counter  
		// case ???:
		//     m.counter--
		}
	}
	
	return m, nil
}

// View renders the current state as a string
func (m model) View() string {
	// TODO: Create the view based on the current state
	// if m.quitting {
	//     return "Goodbye!\n"
	// }
	
	// TODO: Display the counter and instructions
	// s := fmt.Sprintf("Counter: %d\n\n", m.???)
	// s += "Press + to increment, - to decrement\n"  
	// s += "Press q to quit\n"
	// return s
}

// Init returns initial command (can be nil)
func (m model) Init() tea.Cmd {
	// TODO: Return nil since we don't need any initial commands
	// return ???
}

func main() {
	// TODO: Create a new Bubble Tea program with our initial model
	// p := tea.NewProgram(???)
	
	// TODO: Start the program and handle any errors
	// if _, err := p.Run(); err != nil {
	//     fmt.Printf("Error: %v", err)
	//     os.Exit(1)
	// }
}