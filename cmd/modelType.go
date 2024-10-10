package cmd

import (
	"fmt"
	// "os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var InstallMinorError string

type model struct {
	choices        []string
	cursor         int
	selected       map[int]struct{}
	command        string
	successMessage string
	quitFlag       bool
	errorMessage   string
	done           bool
}

type installDoneMsg struct{}
type installError struct{ errMsg string }

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
	ErrorStyles = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("1"))
	successStyles = lipgloss.NewStyle().
			Bold(true).Underline(true).Foreground(lipgloss.Color("10"))
)

func (m model) installDependencies(option string) tea.Cmd {
	return func() tea.Msg {
		fmt.Printf("Installing dependencies for %s...\n", option)
		switch m.command {
		case "mern":
			fmt.Println("hello there")
		case "nodeBackend":
			if err := nodeBackendFunction(m.cursor); err != nil {
				return installError{errMsg: err.Error()}
			}
		case "reactNative":
			fmt.Println("hello there rn")
		case "react":
			fmt.Println("hello there react")

		}

		// Program.Quit()
		return installDoneMsg{}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			if m.done {
				for _, proc := range activeProcesses {
					if proc.Process != nil {
						err := proc.Process.Kill()
						if err != nil {
							fmt.Println("Failed to kill process:", err)
						} else {
							fmt.Println("Process killed successfully")
						}
					}
				}
			}
			return m, tea.Quit
		case "up":
			if m.cursor > 0 && !m.done {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.choices)-1 && !m.done {
				m.cursor++
			}
		case "enter":
			if m.done {
				return m, nil
			}
			m.done = true
			return m, m.installDependencies(m.choices[m.cursor])
		}
	case installDoneMsg:
		m.successMessage = "\n\nDependencies installed!"
		m.quitFlag = true
		return m, nil
	case installError:
		m.errorMessage = "\n\n" + "ERROR: " + msg.errMsg
		// "Error Occured while generating!"
		m.quitFlag = true
		return m, nil

	}
	if m.quitFlag {

		return m, tea.Quit
	}
	return m, nil
}

func (m model) View() string {
	s := titleStyle.Render("Select an option (use arrow keys and press enter):") + "\n\n"
	cursor := " "
	for i, choice := range m.choices {
		var itemStyle lipgloss.Style
		if m.cursor == i {
			cursor = ">"
			itemStyle = selectedTextStyle
		} else {
			itemStyle = normalTextStyle
			cursor = " "
		}
		s += itemStyle.Render(fmt.Sprintf("%s %s", cursor, choice)) + "\n"
	}

	s += "\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Render("Press q to quit.") + "\n"

	if m.successMessage != "" {
		s = successStyles.Render(m.successMessage) + "\n"
		s += InstallMinorError

	} else if m.errorMessage != "" {
		s = ErrorStyles.Render(m.errorMessage) + "\n"
	}
	return s
}

func initialModel(opts []string, _command string) model {
	return model{
		choices:  opts,
		selected: make(map[int]struct{}),
		command:  _command,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
