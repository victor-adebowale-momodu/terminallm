package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	counter int
}

func initialModel() model {
	return model{counter: 0}
}

// INIT
func (m model) Init() tea.Cmd {
	return nil
}

// UPDATE (logic)
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "+":
			m.counter++

		case "-":
			m.counter--
		}
	}

	return m, nil
}

// VIEW (UI)
func (m model) View() string {
	return fmt.Sprintf(
		"\nCounter: %d\n\nPress + or -\nPress q to quit\n",
		m.counter,
	)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
