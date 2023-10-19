package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"os"

	"os/exec"
	"strings"
)

func main() {

	parser := argparse.NewParser("cued", "Tool for querying programming keywords")
	lang := parser.String("l", "language", &argparse.Options{Required: false, Help: "Language to query"})
	query := parser.String("q", "query", &argparse.Options{Required: false, Help: "Query"})
	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	// use pager env var if possible
	pager := os.Getenv("PAGER")
	if pager == "" {
		pager = "less"
	}

	if strings.Contains(*lang, *query) {
		exec.Command(fmt.Sprintf("curl -s cht.sh/%s/%s | $PAGER", *lang, *query))
	} else {
		// Allow empty query search
		if *query != "" {
			exec.Command(fmt.Sprintf("curl -s cht.sh/%s~%s | $PAGER", *lang, *query))
		} else {
			exec.Command(fmt.Sprintf("curl -s cht.sh/%s | $PAGER", *lang))
		}

	}
}
