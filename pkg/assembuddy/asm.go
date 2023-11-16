package assembuddy

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Syscall struct {
	Arch   string `json:"arch"`
	Nr     int    `json:"nr"`
	Name   string `json:"name"`
	Refs   string `json:"refs"`
	Return string `json:"return"`
	Arg0   string `json:"arg0"`
	Arg1   string `json:"arg1"`
	Arg2   string `json:"arg2"`
	Arg3   string `json:"arg3"`
	Arg4   string `json:"arg4"`
	Arg5   string `json:"arg5"`
}

type CLIOptions struct {
	Syscall          string
	Arch             string
	ListQueryMatches bool
	ListArchQueries  bool
}

func QueryASM(opts *CLIOptions) {
	arch := opts.Arch
	syscall := opts.Syscall

	var selectedName string

	if syscall != "" {
		selectedName = syscall
	}

	if selectedName == "" {
		log.Fatalf("No syscall selected.")
	}

	selectedSyscall, err := getSyscallDetails(arch, selectedName)
	if err != nil {
		log.Fatalf("Error fetching syscall details: %s", err)
	}

	syscallJSON, err := json.Marshal(selectedSyscall)
	if err != nil {
		log.Fatalf("Error marshaling syscall data: %s", err)
	}

	DisplaySyscall(syscallJSON)
}

func getSyscalls(arch string) ([]string, error) {
	url := fmt.Sprintf("https://api.syscall.sh/v1/syscalls/%s", arch)
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

func getSyscallDetails(arch, name string) (*Syscall, error) {
	url := fmt.Sprintf("https://api.syscall.sh/v1/syscalls/%s", arch)
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

	var selectedSyscall *Syscall
	for _, syscall := range syscalls {
		if syscall.Name == name {
			selectedSyscall = &syscall
			break
		}
	}

	if selectedSyscall == nil {
		return nil, fmt.Errorf("Syscall not found: %s", name)
	}

	return selectedSyscall, nil
}
