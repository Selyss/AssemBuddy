package assembuddy

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
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

var syscalls []Syscall
var currentArch, currentSyscall string

func init() {
	if err := json.Unmarshal(syscallsData, &syscalls); err != nil {
		panic("Failed to load syscall data: " + err.Error())
	}
}

func FetchData(endpointURL string) ([]Syscall, error) {
	return filterSyscalls(), nil
}

func PrettyPrint(endpointURL string) error {
	results := filterSyscalls()
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	fmt.Println(string(jsonData))
	return nil
}

func GetSyscallData(opts *CLIOptions) (string, error) {
	if opts.Arch != "" && opts.Arch != "x64" && opts.Arch != "x86" && opts.Arch != "arm" && opts.Arch != "arm64" {
		return "", errors.New("invalid architecture")
	}
	// store filter options
	currentArch = opts.Arch
	currentSyscall = opts.Syscall
	return "offline", nil
}

func filterSyscalls() []Syscall {
	var results []Syscall
	
	for _, sc := range syscalls {
		matchesArch := currentArch == "" || strings.EqualFold(sc.Arch, currentArch)
		matchesSyscall := currentSyscall == ""
		
		if currentSyscall != "" {
			// match name
			if strings.EqualFold(sc.Name, currentSyscall) {
				matchesSyscall = true
			} else {
				// match number
				if nr, err := strconv.Atoi(currentSyscall); err == nil && sc.Nr == nr {
					matchesSyscall = true
				}
			}
		}
		
		if matchesArch && matchesSyscall {
			results = append(results, sc)
		}
	}
	
	return results
}
