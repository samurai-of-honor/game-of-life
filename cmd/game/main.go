package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	p := tea.NewProgram(InitModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}

type Model struct {
	Choices []string // items on list
	Cursor  int      // which item our cursor is pointing at
}

func InitModel() *Model {
	return &Model{
		Choices: []string{
			"From file",
			"Input manually",
			"Random universe",
		},
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit

		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}

		case "down":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}

		case "enter":

		}
	}
	return m, nil
}

func (m *Model) View() string {
	s := "Choose source for universe:\n\n"

	for i, choice := range m.Choices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}

		// Render the row
		s += fmt.Sprintf("[%s] %s\n", cursor, choice)
	}

	s += "\nPress q to quit.\n"
	return s
}
