// interactive_lists.go - Solution
// Learn Bubble Tea list selection and navigation

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
	cursor   int
	selected int
	quitting bool
}

func initialModel() model {
	return model{
		cursor:   0,
		selected: -1,
		quitting: false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			m.quitting = true
			return m, tea.Quit
		
		case "up":
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(items) - 1
			}
		
		case "down":
			if m.cursor < len(items)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}
		
		case "enter":
			m.selected = m.cursor
			return m, tea.Quit
		}
	}
	
	return m, nil
}

func (m model) View() string {
	if m.quitting {
		if m.selected >= 0 {
			return fmt.Sprintf("You selected: %s\nGoodbye!\n", items[m.selected])
		}
		return "Goodbye!\n"
	}
	
	s := "Select a fruit:\n\n"
	
	for i, item := range items {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor points to current item
		}
		s += fmt.Sprintf("%s %s\n", cursor, item)
	}
	
	s += "\nUse ↑/↓ to navigate, Enter to select, q to quit\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}