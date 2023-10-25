package chtsht

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strings"
)

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
	cmd := exec.Command("curl", "-s", url)
	cmd.Stderr = os.Stderr

	lessCmd := exec.Command("less")
	lessCmd.Stdin, _ = cmd.StdoutPipe()
	lessCmd.Stdout = os.Stdout

	if err := cmd.Start(); err != nil {
		log.Fatalf("Error while querying: %s", err)
	}

	if err := lessCmd.Run(); err != nil {
		log.Fatalf("Error while piping into $PAGER: %s", err)
	}

	cmd.Wait()
	return
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
