package main

// A simple program demonstrating the spinner component from the Bubbles
// component library.

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type errMsg error

type model struct {
	err      error
	resp     string
	spinner  spinner.Model
	quitting bool
}

func Action() tea.Msg {
	time.Sleep(time.Second * 2)
	resp := "Wow"
	return statusMsg(resp)
}

type statusMsg string

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{spinner: s}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(Action, m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case statusMsg:
		m.resp = string(msg)
		return m, tea.Quit
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("%sLoading", m.spinner.View())

	if m.resp != "" {
		return fmt.Sprintf("%s\n", m.resp)
	}

	if m.quitting {
		return ""
	}
	return str
}

func Spinner() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
