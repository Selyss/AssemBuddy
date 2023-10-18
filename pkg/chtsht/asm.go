package chtsht

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type Syscall struct {
	Name string `json:"name"`
}

func QueryASM(name string) {
	arch := []string{"x64", "x86", "arm", "arm64"}

	selected, err := SelectFromList(languages)
	if err != nil {
		fmt.Printf("Error selecting a language: %v\n", err)
		os.Exit(1)
	}

	var query string
	fmt.Print("Enter Query: ")
	_, err = fmt.Scanln(&query)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	// TODO: put in $PAGER
	//
	// FIXME: do i want to make the architecture choice and then it fuzzy finds over the possible names? I think so.

	// cmd := exec.Command("tmux", "neww", "bash", "-c", fmt.Sprintf("echo \"curl cht.sh/%s/%s/\" & curl cht.sh/%s/%s & while [ : ]; do sleep 1; done", selected, query, selected, query))

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error running command: %v\n", err)
		os.Exit(1)
	}
}
