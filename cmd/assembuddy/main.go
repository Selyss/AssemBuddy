package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Selyss/AssemBuddy/internal/cli"
)

var Version = "dev"

func main() {
	root := cli.NewRootCommand(Version)
	if err := root.Execute(); err != nil {
		var exitErr cli.ExitError
		if errors.As(err, &exitErr) {
			if exitErr.Code != 0 && exitErr.Err != nil {
				fmt.Fprintln(os.Stderr, exitErr.Err)
			}
			os.Exit(exitErr.Code)
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(3)
	}
}
