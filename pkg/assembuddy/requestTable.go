package assembuddy

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

//go:embed syscalls.json
var syscallsData []byte

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

func LoadSyscallData() ([]Syscall, error) {
	var syscalls []Syscall
	err := json.Unmarshal(syscallsData, &syscalls)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal embedded JSON: %w", err)
	}

	return syscalls, nil
}

func FetchData(query string) ([]Syscall, error) {
	allSyscalls, err := LoadSyscallData()
	if err != nil {
		return nil, err
	}

	if query == "" {
		return allSyscalls, nil
	}

	return FilterSyscalls(allSyscalls, query), nil
}

func FilterSyscalls(syscalls []Syscall, query string) []Syscall {
	queryLower := strings.ToLower(query)
	var filtered []Syscall

	for _, syscall := range syscalls {
		if strings.Contains(strings.ToLower(syscall.Name), queryLower) ||
			strings.Contains(strings.ToLower(syscall.Arch), queryLower) {
			filtered = append(filtered, syscall)
		}
	}

	return filtered
}

func PrettyPrint(query string) error {
	syscalls, err := FetchData(query)
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(syscalls, "", " ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	fmt.Println(string(jsonData))
	return nil
}

func GetSyscallData(opts *CLIOptions) (string, error) {
	arch := opts.Arch
	syscall := opts.Syscall

	// validate architecture
	if arch != "" && arch != "x64" && arch != "x86" && arch != "arm" && arch != "arm64" {
		return "", errors.New("invalid architecture")
	}

	query := ""

	if arch != "" {
		query = arch
	}
	if syscall != "" {
		if query != "" {
			query += " " + syscall
		} else {
			query = syscall
		}
	}

	return query, nil
}
