package chtsht

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Syscall struct {
	Name string `json:"name"`
}

func QueryASM() {
	archs := []string{"x64", "x86", "arm", "arm64"}

	arch, err := SelectFromList(archs)
	if err != nil {
		fmt.Printf("Error selecting a language: %v\n", err)
		os.Exit(1)
	}

	syscalls, err := getSyscalls(arch)
	if err != nil {
		fmt.Printf("Error fetching syscalls for %s: %v\n", arch, err)
		os.Exit(1)

	}

	selected, err := SelectFromList(syscalls)
	if err != nil {
		fmt.Printf("Error selecting a syscall: %v\n", err)
		os.Exit(1)
	}

	if selected == "" {
		fmt.Println("No syscall selected.")

		os.Exit(0)
	}

	fmt.Println("Selected syscall:", selected)

}

func getSyscalls(arch string) ([]string, error) {
	url := fmt.Sprintf("https://api.syscall.sh/v1/syscalls/%s", arch)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
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
