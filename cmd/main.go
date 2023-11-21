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

	listArch := parser.Flag("r", "list-arch", &argparse.Options{Help: "Get all supported architechture convensions"})
	listQuery := parser.Flag("n", "list-name", &argparse.Options{Help: "Get all syscalls with given name"})

	prettyPrint := parser.Flag("p", "pretty-print", &argparse.Options{Help: "Pretty print JSON result"})

	err := parser.Parse(os.Args)
	if err != nil || (*query == "" && *arch == "") && (!*listArch && !*listQuery) {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	opts.Syscall = *query
	opts.Arch = *arch
	opts.ListQueryMatches = *listQuery
	opts.ListArchQueries = *listArch
	opts.PrettyPrint = *prettyPrint

	return opts
}

func main() {
	opts := parseArgs()
	if opts.ListArchQueries {
		_, err := assembuddy.ArchInfo()
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
	}
	if opts.ListQueryMatches {
		_, err := assembuddy.QueryInfo(opts.Arch)
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
	}
	table, err := assembuddy.GetSyscallData(opts.Arch, opts.Syscall, opts.PrettyPrint)
	if err != nil {
		log.Fatal(err)
	}
	assembuddy.RenderTable(opts.Arch, table)
}
