package assembuddy

import (
	_ "embed"
	"encoding/json"
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

	return FilterSyscalls(allSyscalls, query), nil
}

func FetchDataWithArch(query string, arch string) ([]Syscall, error) {
	allSyscalls, err := LoadSyscallData()
	if err != nil {
		return nil, err
	}

	filtered := FilterByArch(allSyscalls, arch)

	if query != "" {
		filtered = FilterSyscalls(filtered, query)
	}

	return filtered, nil
}

func FilterSyscalls(syscalls []Syscall, query string) []Syscall {
	if query == "" {
		return syscalls
	}

	queryLower := strings.ToLower(query)
	var filtered []Syscall

	for _, syscall := range syscalls {
		if strings.Contains(strings.ToLower(syscall.Name), queryLower) {
			filtered = append(filtered, syscall)
		}
	}

	return filtered
}

func FilterByArch(syscalls []Syscall, arch string) []Syscall {
	if arch == "" {
		return syscalls
	}

	var filtered []Syscall
	for _, syscall := range syscalls {
		if strings.EqualFold(syscall.Arch, arch) {
			filtered = append(filtered, syscall)
		}
	}

	return filtered
}

func PrettyPrint(query string, arch string) error {
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

func GetSyscallData(opts *CLIOptions) (string, string, error) {
	arch := opts.Arch
	syscall := opts.Syscall

	// validate architecture
	if arch != "" {
		validArchs := []string{"x64", "x86", "arm", "arm64"}
		valid := false
		for _, validArch := range validArchs {
			if arch == validArch {
				valid = true
				break
			}
		}
		if !valid {
			return "", "", fmt.Errorf("invalid architecture: %s", arch)
		}
	}
	return syscall, arch, nil
}
