package cmd

import (
	"fmt"

	// "os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	choices        []string
	cursor         int
	selected       map[int]struct{}
	command        string
	successMessage string
	quitFlag       bool
	errorMessage   string
	done           bool
	progressString []string
	input          string
}

type logMsg struct {
	msg    string
	remove bool
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
		Program.Send(logMsg{msg: fmt.Sprintf("Installing dependencies for %s...\n", option), remove: true})
		switch m.command {
		case "mern":
			if err := mernFunc(m.cursor, "myProject"); err != nil {
				return installError{errMsg: err.Error()}
			}
		case "nodeBackend":
			if err := nodeBackendFunction(m.cursor); err != nil {
				return installError{errMsg: err.Error()}
			}
		case "reactNative":
			// Program.Send(logMsg{msg: fmt.Sprint("Need to install react-native-cli (y/n): ", option), remove: true})
			// fmt.Print("Need to install react-native-cli (y/n): ")
			if err := reactNativeFunc(m.cursor, "App"); err != nil {
				return installError{errMsg: err.Error()}
			}
		case "react":
			if err := reactFunc(m.cursor, "myproject"); err != nil {
				return installError{errMsg: err.Error()}
			}
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
				fmt.Println("\033[A\033[K")
				fmt.Print("\033[A\033[K")
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
		default:
			// Append typed characters to the input string
			if m.done {
				m.input += msg.String()
				// Program.Send(logMsg{msg: m.progressString[len(m.progressString)-1] + m.input, remove: true})
				fmt.Print(m.input)
				return m, nil
			}

		}

	case installDoneMsg:
		m.successMessage = "\n\nDependencies installed!\n\n"
		m.quitFlag = true
		return m, nil
	case installError:
		m.errorMessage = "\n\n" + "ERROR: " + msg.errMsg + "\n\n"
		// "Error Occured while generating!"
		m.quitFlag = true
		return m, nil
	case logMsg:
		if msg.remove {
			if len(m.progressString) > 0 {
				m.progressString = m.progressString[:len(m.progressString)-1]
			}
			m.progressString = append(m.progressString, lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Render(msg.msg))
		} else {
			m.progressString = append(m.progressString, lipgloss.NewStyle().Foreground(lipgloss.Color("3")).Render(msg.msg))
		}
		return m, nil

	}
	if m.quitFlag {
		fmt.Print("\033[A\033[K")
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

	for _, v := range m.progressString {
		s += (v + "\n")
	}
	if m.successMessage != "" {
		s += successStyles.Render(m.successMessage) + "\n"
	} else if m.errorMessage != "" {
		s += ErrorStyles.Render(m.errorMessage) + "\n"
	}
	return s
}

func initialModel(opts []string, _command string) model {
	return model{
		choices:        opts,
		selected:       make(map[int]struct{}),
		command:        _command,
		done:           false,
		progressString: []string{},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
