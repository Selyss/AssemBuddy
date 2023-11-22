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

	listArch := parser.Flag("r", "list-arch", &argparse.Options{Help: "Get all supported architechture convensions"})

	prettyPrint := parser.Flag("p", "pretty-print", &argparse.Options{Help: "Pretty print JSON result"})

	err := parser.Parse(os.Args)
	if err != nil || (*query == "" && *arch == "") && !*listArch {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	opts.Syscall = *query
	opts.Arch = *arch
	opts.ListArch = *listArch
	opts.PrettyPrint = *prettyPrint

	return opts
}

func main() {
	opts := parseArgs()
	if opts.ListArch {
		_, err := assembuddy.ArchInfo()
		if err != nil {
			log.Fatalf("Error: %s", err)
		}
	}
	table, err := assembuddy.GetSyscallData(opts)
	if err != nil {
		log.Fatal(err)
	}
	assembuddy.RenderTable(opts, table)
}
