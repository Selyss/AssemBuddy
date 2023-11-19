package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Selyss/AssemBuddy/pkg/assembuddy"
	"github.com/akamensky/argparse"
)

type CLIOptions struct {
	Syscall          string
	Arch             string
	ListQueryMatches bool
	ListArchQueries  bool
	PrettyPrint      bool
}

func parseArgs() *CLIOptions {
	opts := &CLIOptions{}

	parser := argparse.NewParser("AssemBuddy", "Tool for querying assembly keywords")
	query := parser.String("q", "query", &argparse.Options{Help: "Search query"})
	arch := parser.String("a", "architecture", &argparse.Options{Help: "Architecture for queries"})

	listArchQueries := parser.Flag("r", "list-arch", &argparse.Options{Help: "Get all syscalls from given architechture"})
	listQueryMatches := parser.Flag("n", "list-name", &argparse.Options{Help: "Get all syscalls with given name"})

	// TODO: result is definitely printed but im not convinced its pretty.
	prettyPrint := parser.Flag("p", "pretty-print", &argparse.Options{Help: "Pretty print JSON result"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	opts.Syscall = *query
	opts.Arch = *arch
	opts.ListQueryMatches = *listQueryMatches
	opts.ListArchQueries = *listArchQueries
	opts.PrettyPrint = *prettyPrint

	return opts
}

func main() {
	opts := parseArgs()
	if opts.Syscall == "" && opts.Arch == "" {
		table, err := assembuddy.GetArchData(opts.Arch, opts.PrettyPrint)
		if err != nil {
			log.Fatal(err)
		}
		assembuddy.RenderArchTable(opts.Arch, table)
	}

	if opts.Syscall != "" {
		table, err := assembuddy.GetNameData(opts.Syscall, opts.PrettyPrint)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		assembuddy.RenderNameTable(table)
	}
	if opts.Arch != "" {
		table, err := assembuddy.GetArchData(opts.Arch, opts.PrettyPrint)
		if err != nil {
			log.Fatal(err)
		}
		assembuddy.RenderArchTable(opts.Arch, table)
	}
}
