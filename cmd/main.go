package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Selyss/AssemBuddy/pkg/assembuddy"
	"github.com/akamensky/argparse"
	// tea "github.com/charmbracelet/bubbletea"
)

type CLIOptions struct {
	Syscall          string
	Arch             string
	ListQueryMatches bool
	ListArchQueries  bool
}

func parseArgs() *CLIOptions {
	opts := &CLIOptions{}

	parser := argparse.NewParser("AssemBuddy", "Tool for querying assembly keywords")
	query := parser.String("q", "query", &argparse.Options{Help: "Search query"})
	arch := parser.String("a", "architecture", &argparse.Options{Help: "Architecture for queries"})

	listArchQueries := parser.Flag("r", "list-arch", &argparse.Options{Help: "Get all syscalls from given architechture"})
	listQueryMatches := parser.Flag("n", "list-name", &argparse.Options{Help: "Get all syscalls with given name"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	opts.Syscall = *query
	opts.Arch = *arch
	opts.ListQueryMatches = *listQueryMatches
	opts.ListArchQueries = *listArchQueries

	return opts
}

func main() {
	opts := parseArgs()
	if opts.Syscall != "" {
		table, err := assembuddy.GetNameData(opts.Syscall)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Println(table)
	}
	if opts.Arch != "" {
		table, err := assembuddy.GetArchData(opts.Arch)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(table)
	}
}
