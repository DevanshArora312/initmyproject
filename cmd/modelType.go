package cmd

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	done     bool
}

var Program *tea.Program

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("13")).
			Underline(true)

	normalTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("6"))

	selectedTextStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("10")).
				Underline(true)
)

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.done {
		return m, nil
	}
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
			m.done = true
			return m, installDependencies(m.choices[m.cursor])
		}

		return m, nil
	}
	return m, nil
}

type installDoneMsg struct{}

func installDependencies(option string) tea.Cmd {
	return func() tea.Msg {
		fmt.Printf("Installing dependencies for %s...\n", option)
		time.Sleep(2 * time.Second)
		fmt.Println("Dependencies installed!")
		Program.Quit()
		return installDoneMsg{}
	}
}

func (m model) View() string {
	s := titleStyle.Render("Select an option (use arrow keys and press enter):") + "\n\n"
	cursor := " "
	for i, choice := range m.choices {
		var itemStyle lipgloss.Style
		if m.cursor == i {
			cursor = ">"
			itemStyle = selectedTextStyle // Apply style for the selected item
		} else {
			itemStyle = normalTextStyle
			cursor = " "
		}
		s += itemStyle.Render(fmt.Sprintf("%s %s", cursor, choice)) + "\n"
	}

	s += "\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Render("Press q to quit.") + "\n"
	return s
}

func initialModel(opts []string) model {
	return model{
		choices:  opts,
		done:     false,
		selected: make(map[int]struct{}),
	}
}

// func installDependencies(index int) {
// 	switch index {
// 	case 0:
// 		fmt.Println("0")
// 	case 1:
// 		fmt.Println("1")
// 	case 2:
// 		fmt.Println("2")
// 	default:
// 		fmt.Println("3")
// 	}
// }
