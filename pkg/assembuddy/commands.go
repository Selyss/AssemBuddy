package chtsht

import (
	"bufio"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
	"os/exec"
	"strings"
)

func SelectFromList(items []string) (string, error) {
	p := tea.NewProgram(initialModel())
	go DisplayLoadingSpinner(p)
	cmd := exec.Command("fzf")
	cmd.Stdin = strings.NewReader(strings.Join(items, "\n"))
	p.Quit()
	cmd.Stderr = os.Stderr
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func curlOutput(url string) error {
	cmd := exec.Command("curl", "-s", url)
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	return err
}

func DisplayOutput(url string) {
	p := tea.NewProgram(initialModel())
	go DisplayLoadingSpinner(p)

	err := curlOutput(url)

	p.Quit()
	cmd := exec.Command("less")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error running 'less': %s", err)
	}
}

func ChtReadOptions() ([]string, error) {
	readFile, err := os.Open("chtsht.txt") // TODO:

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
