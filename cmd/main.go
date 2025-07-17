package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Selyss/AssemBuddy/pkg/assembuddy"
	"github.com/akamensky/argparse"
)

func parseArgs() *assembuddy.CLIOptions {
	opts := &assembuddy.CLIOptions{}

	parser := argparse.NewParser("AssemBuddy", "Tool for querying assembly keywords")
	query := parser.String("q", "query", &argparse.Options{Help: "Search query"})
	arch := parser.String("a", "architecture", &argparse.Options{Help: "Architecture for queries"})

	prettyPrint := parser.Flag("p", "pretty-print", &argparse.Options{Help: "Pretty print JSON result"})

	err := parser.Parse(os.Args)
	if err != nil || (*query == "" && *arch == "") {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	opts.Syscall = *query
	opts.Arch = *arch
	opts.PrettyPrint = *prettyPrint

	return opts
}

func main() {
	opts := parseArgs()

	syscallData(opts)
}

func syscallData(opts *assembuddy.CLIOptions) {
	query, err := assembuddy.GetSyscallData(opts)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if opts.PrettyPrint {
		assembuddy.PrettyPrint(query)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	} else {
		table, err := assembuddy.FetchData(query)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		assembuddy.RenderTable(opts, table)
	}
}
