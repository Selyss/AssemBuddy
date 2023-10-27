package chtsht

import (
	"bufio"
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
	"os"
	"os/exec"
	"strings"
)

type errMsg error

type model struct {
	spinner  spinner.Model
	quitting bool
	err      error
}

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{spinner: s}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
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
	str := fmt.Sprintf("\n   %s Loading information...press q to quit\n", m.spinner.View())
	if m.quitting {
		return str + "\n"
	}
	return str
}

func SelectFromList(items []string) (string, error) {

	cmd := exec.Command("fzf")
	cmd.Stdin = strings.NewReader(strings.Join(items, "\n"))
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func DisplayOutput(url string) {
	p := tea.NewProgram(initialModel())

	cmd := exec.Command("curl", "-s", url)
	cmd.Stderr = os.Stderr

	lessCmd := exec.Command("less")
	lessCmd.Stdin, _ = cmd.StdoutPipe()
	lessCmd.Stdout = os.Stdout

	go func() {
		if _, err := p.Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error while querying: %s", err)
	}

	go func() {
		if err := lessCmd.Run(); err != nil {
			log.Fatalf("Error while piping into $PAGER: %s", err)
		}
		p.Quit()
	}()

	if err := cmd.Run(); err != nil {
		log.Fatalf("Error while querying: %s", err)
	}
}

func ChtReadOptions() ([]string, error) {
	readFile, err := os.Open("chtsht.txt") // FIXME:

	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	return fileLines, nil
}
