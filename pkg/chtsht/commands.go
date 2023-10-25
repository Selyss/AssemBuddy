package chtsht

import (
	"encoding/json"
	"io"
	"net/http"
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

func SelectFromCht() ([]string, error) {
	url := "https://cht.sh/:list"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var topicResp []string
	if err := json.Unmarshal(body, &topicResp); err != nil {
		return nil, err
	}
	var topicOptions []string
	for _, topic := range topicResp {
		topicOptions = append(topicOptions, topic)
	}

	return topicOptions, nil
}
