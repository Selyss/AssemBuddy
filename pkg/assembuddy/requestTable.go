package assembuddy

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Syscall struct {
	Arch        string `json:"arch"`
	Name        string `json:"name"`
	ReturnValue string `json:"return"`
	Arg0        string `json:"arg0"`
	Arg1        string `json:"arg1"`
	Arg2        string `json:"arg2"`
	Arg3        string `json:"arg3"`
	Arg4        string `json:"arg4"`
	Arg5        string `json:"arg5"`
	Nr          int    `json:"nr"`
}

type CLIOptions struct {
	Syscall     string
	Arch        string
	ListArch    bool
	PrettyPrint bool
}

const (
	syscallEndpoint    = "https://api.syscall.sh/v1/syscalls"
	conventionEndpoint = "https://api.syscall.sh/v1/conventions"
)

func FetchData(endpointURL string) ([]Syscall, error) {
	response, err := http.Get(endpointURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var systemCalls []Syscall
	err = json.Unmarshal(body, &systemCalls)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return systemCalls, nil
}

func PrettyPrint(endpointURL string) error {
	response, err := http.Get(endpointURL)
	if err != nil {
		return fmt.Errorf("failed to fetch data: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	fmt.Println(string(body))

	return nil
}

func GetSyscallData(opts *CLIOptions) (string, error) {
	arch := opts.Arch
	url := syscallEndpoint
	// if arch is x64, x86, arm, or arm64, concat to endpointURL
	if arch == "x64" || arch == "x86" || arch == "arm" || arch == "arm64" {
		url += "/" + arch
		// if arch is not empty, return error
	} else if arch != "" {
		return "", errors.New("invalid architecture")
	}
	if opts.Syscall != "" {
		url += "/" + opts.Syscall
	}
	return url, nil
}
