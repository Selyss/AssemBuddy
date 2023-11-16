package assembuddy

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListOptions provides options for listing assemblies.
func ListArchSyscalls() {
}

func ListQueryMatches(query string) ([]string, error) {
	url := fmt.Sprintf("https://api.syscall.sh/v1/syscalls/%s", query)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var syscalls []Syscall

	if err := json.Unmarshal(body, &syscalls); err != nil {
		return nil, err
	}

	var syscallNames []string
	for _, syscall := range syscalls {
		syscallNames = append(syscallNames, syscall.Name)
	}

	return syscallNames, nil
}
