// form_handling.go
// Learn Bubble Tea form input and text field handling
//
// Forms are essential for gathering user input in TUI applications.
// This exercise demonstrates text input fields, form validation,
// and multi-field navigation using Bubble Tea.
//
// I AM NOT DONE YET!

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
	// TODO: Add slice of fields to represent form fields
	// fields ???
	
	// TODO: Add current field index
	// currentField ???
	
	// TODO: Add submitted flag
	// submitted ???
	
	// TODO: Add quitting flag  
	// quitting ???
}

func initialModel() model {
	// TODO: Create initial model with form fields
	// return model{
	//     fields: []field{
	//         {label: "Name", value: "", active: true, required: true},
	//         {label: "Email", value: "", active: false, required: true},
	//         {label: "Age", value: "", active: false, required: false},
	//         {label: "City", value: "", active: false, required: false},
	//     },
	//     currentField: ???,
	//     submitted: ???,
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
		// TODO: Handle quit
		// case ???:
		//     m.quitting = true
		//     return m, tea.Quit
		
		// TODO: Handle tab - move to next field
		// case "tab":
		//     m.fields[m.currentField].active = false
		//     m.currentField = (m.currentField + 1) % len(m.fields)
		//     m.fields[m.currentField].active = true
		
		// TODO: Handle shift+tab - move to previous field  
		// case "shift+tab":
		//     m.fields[m.currentField].active = false
		//     if m.currentField == 0 {
		//         m.currentField = len(m.fields) - 1
		//     } else {
		//         m.currentField--
		//     }
		//     m.fields[m.currentField].active = true
		
		// TODO: Handle enter - submit form
		// case "enter":
		//     if m.validateForm() {
		//         m.submitted = true
		//         return m, tea.Quit
		//     }
		
		// TODO: Handle backspace - remove character
		// case "backspace":
		//     if len(m.fields[m.currentField].value) > 0 {
		//         m.fields[m.currentField].value = m.fields[m.currentField].value[:len(m.fields[m.currentField].value)-1]
		//     }
		
		// TODO: Handle regular character input
		// default:
		//     if len(msg.String()) == 1 {
		//         m.fields[m.currentField].value += msg.String()
		//     }
		}
	}
	
	return m, nil
}

// TODO: Implement form validation
// func (m model) validateForm() bool {
//     for _, field := range m.fields {
//         if field.required && strings.TrimSpace(field.value) == "" {
//             return false
//         }
//     }
//     return true
// }

func (m model) View() string {
	// TODO: Handle different states
	// if m.quitting {
	//     if m.submitted {
	//         s := "Form submitted successfully!\n\n"
	//         for _, field := range m.fields {
	//             if field.value != "" {
	//                 s += fmt.Sprintf("%s: %s\n", field.label, field.value)
	//             }
	//         }
	//         return s + "\nGoodbye!\n"
	//     }
	//     return "Form cancelled. Goodbye!\n"
	// }
	
	// TODO: Render form
	// s := "User Registration Form\n"
	// s += "=====================\n\n"
	
	// TODO: Render each field
	// for i, field := range m.fields {
	//     // TODO: Show field with appropriate styling
	//     prefix := "  "
	//     if field.active {
	//         prefix = "> "
	//     }
		
	//     required := ""
	//     if field.required {
	//         required = " *"
	//     }
		
	//     s += fmt.Sprintf("%s%s%s: %s\n", prefix, field.label, required, field.value)
	// }
	
	// TODO: Add instructions
	// s += "\n* Required fields\n"
	// s += "Tab/Shift+Tab: Navigate fields\n"
	// s += "Enter: Submit form\n" 
	// s += "Ctrl+C: Quit\n"
	
	// TODO: Show validation message if needed
	// if !m.validateForm() {
	//     s += "\n❌ Please fill all required fields\n"
	// }
	
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