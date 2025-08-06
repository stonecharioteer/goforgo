// interactive_lists.go
// Learn Bubble Tea list selection and navigation
//
// Bubble Tea makes it easy to create interactive lists with navigation,
// selection, and custom rendering. This exercise demonstrates building
// a navigable list of items with keyboard controls.
//
// I AM NOT DONE YET!

package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Define our list items
var items = []string{
	"🍎 Apple",
	"🍌 Banana", 
	"🍊 Orange",
	"🍇 Grapes",
	"🥝 Kiwi",
	"🍓 Strawberry",
	"🥭 Mango",
}

type model struct {
	// TODO: Add field to track current cursor position
	// cursor ???
	
	// TODO: Add field to track selected item index (-1 if none selected)
	// selected ???
	
	// TODO: Add field to track if we should quit
	// quitting ???
}

func initialModel() model {
	// TODO: Return initial model with cursor at 0, no selection, not quitting
	// return model{
	//     cursor: ???,
	//     selected: ???,
	//     quitting: ???,
	// }
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// TODO: Handle quit keys
		// case ???, ???:
		//     m.quitting = true
		//     return m, tea.Quit
		
		// TODO: Handle up arrow - move cursor up (with wraparound)
		// case ???:
		//     if m.cursor > 0 {
		//         m.cursor--
		//     } else {
		//         m.cursor = len(items) - 1
		//     }
		
		// TODO: Handle down arrow - move cursor down (with wraparound)
		// case ???:
		//     if m.cursor < len(items)-1 {
		//         m.cursor++
		//     } else {
		//         m.cursor = 0
		//     }
		
		// TODO: Handle enter - select current item
		// case ???:
		//     m.selected = m.cursor
		//     return m, tea.Quit
		}
	}
	
	return m, nil
}

func (m model) View() string {
	// TODO: Handle quit state
	// if m.quitting {
	//     if m.selected >= 0 {
	//         return fmt.Sprintf("You selected: %s\nGoodbye!\n", items[m.selected])
	//     }
	//     return "Goodbye!\n"
	// }
	
	// TODO: Build the list view
	// s := "Select a fruit:\n\n"
	
	// TODO: Iterate through items and render each one
	// for i, item := range items {
	//     cursor := " " // no cursor
	//     if ??? == i {
	//         cursor = ">" // cursor points to current item
	//     }
	//     s += fmt.Sprintf("%s %s\n", cursor, item)
	// }
	
	// TODO: Add instructions
	// s += "\nUse ↑/↓ to navigate, Enter to select, q to quit\n"
	// return s
}

func main() {
	// TODO: Create and run the program
	// p := tea.NewProgram(???)
	// if _, err := p.Run(); err != nil {
	//     fmt.Printf("Error: %v", err)
	//     os.Exit(1)
	// }
}