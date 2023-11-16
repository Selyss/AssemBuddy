package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	// tea "github.com/charmbracelet/bubbletea"
)

type CLIOptions struct {
	Query        string
	Architecture string
	ListQueries  bool
	AllStrings   bool
}

func parseArgs() *CLIOptions {
	options := &CLIOptions{}

	parser := argparse.NewParser("AssemBuddy", "Tool for querying assembly keywords")
	query := parser.String("q", "query", &argparse.Options{Help: "Search query"})
	architecture := parser.String("a", "architecture", &argparse.Options{Help: "Architecture for queries"})

	listQueries := parser.Flag("r", "list-arch", &argparse.Options{Help: "Get all syscalls from given architechture"})
	allStrings := parser.Flag("n", "list-name", &argparse.Options{Help: "Get all syscalls with given name"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	options.Query = *query
	options.Architecture = *architecture
	options.ListQueries = *listQueries
	options.AllStrings = *allStrings

	return options
}

func main() {
	options := parseArgs()
	fmt.Println(options)
}
