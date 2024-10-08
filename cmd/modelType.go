package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("13")).
			Underline(true)

	normalTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("7"))

	selectedTextStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("10")).
				Underline(true)
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := titleStyle.Render("Select an option (use arrow keys and press enter):") + "\n\n"

	for i, choice := range m.choices {
		var itemStyle lipgloss.Style
		if m.cursor == i {
			itemStyle = selectedTextStyle // Apply style for the selected item
		} else {
			itemStyle = normalTextStyle // Apply default style
		}
		s += itemStyle.Render(choice) + "\n"
	}

	s += "\n" + normalTextStyle.Render("Press q to quit.") + "\n"
	return s
}

func initialModel(opts []string) model {
	return model{
		choices:  opts,
		selected: make(map[int]struct{}),
	}
}
