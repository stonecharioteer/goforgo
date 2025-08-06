// form_handling.go - Solution
// Learn Bubble Tea form input and text field handling

package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type field struct {
	label    string
	value    string
	active   bool
	required bool
}

type model struct {
	fields       []field
	currentField int
	submitted    bool
	quitting     bool
}

func initialModel() model {
	return model{
		fields: []field{
			{label: "Name", value: "", active: true, required: true},
			{label: "Email", value: "", active: false, required: true},
			{label: "Age", value: "", active: false, required: false},
			{label: "City", value: "", active: false, required: false},
		},
		currentField: 0,
		submitted:    false,
		quitting:     false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		
		case "tab":
			m.fields[m.currentField].active = false
			m.currentField = (m.currentField + 1) % len(m.fields)
			m.fields[m.currentField].active = true
		
		case "shift+tab":
			m.fields[m.currentField].active = false
			if m.currentField == 0 {
				m.currentField = len(m.fields) - 1
			} else {
				m.currentField--
			}
			m.fields[m.currentField].active = true
		
		case "enter":
			if m.validateForm() {
				m.submitted = true
				return m, tea.Quit
			}
		
		case "backspace":
			if len(m.fields[m.currentField].value) > 0 {
				m.fields[m.currentField].value = m.fields[m.currentField].value[:len(m.fields[m.currentField].value)-1]
			}
		
		default:
			if len(msg.String()) == 1 {
				m.fields[m.currentField].value += msg.String()
			}
		}
	}
	
	return m, nil
}

func (m model) validateForm() bool {
	for _, field := range m.fields {
		if field.required && strings.TrimSpace(field.value) == "" {
			return false
		}
	}
	return true
}

func (m model) View() string {
	if m.quitting {
		if m.submitted {
			s := "Form submitted successfully!\n\n"
			for _, field := range m.fields {
				if field.value != "" {
					s += fmt.Sprintf("%s: %s\n", field.label, field.value)
				}
			}
			return s + "\nGoodbye!\n"
		}
		return "Form cancelled. Goodbye!\n"
	}
	
	s := "User Registration Form\n"
	s += "=====================\n\n"
	
	for _, field := range m.fields {
		prefix := "  "
		if field.active {
			prefix = "> "
		}
		
		required := ""
		if field.required {
			required = " *"
		}
		
		s += fmt.Sprintf("%s%s%s: %s\n", prefix, field.label, required, field.value)
	}
	
	s += "\n* Required fields\n"
	s += "Tab/Shift+Tab: Navigate fields\n"
	s += "Enter: Submit form\n"
	s += "Ctrl+C: Quit\n"
	
	if !m.validateForm() {
		s += "\n❌ Please fill all required fields\n"
	}
	
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}