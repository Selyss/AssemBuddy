package assembuddy

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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

func fetchData(endpointURL string, prettyp bool) ([]Syscall, error) {
	response, err := http.Get(endpointURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if prettyp {

		fmt.Println(string(body))
		os.Exit(0)
	}

	var systemCalls []Syscall
	err = json.Unmarshal(body, &systemCalls)
	if err != nil {
		return nil, err
	}

	return systemCalls, nil
}

func GetSyscallData(arch string, name string, prettyp bool) ([]Syscall, error) {
	url := "https://api.syscall.sh/v1/syscalls/"
	// if arch is x64, x86, arm, or arm64, concat to endpointURL
	if arch == "x64" || arch == "x86" || arch == "arm" || arch == "arm64" {
		url += arch
		// if arch is not empty, return error
	} else if arch != "" {
		return nil, errors.New("invalid architecture")
	}
	url += name
	return fetchData(url, prettyp)
}

func ArchInfo() ([]Syscall, error) {
	url := "https://api.syscall.sh/v1/conventions"
	return fetchData(url, true)
}
