package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akamensky/argparse"

	"os/exec"
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

	// Allow empty query search
	if *query != "" {
		url := fmt.Sprintf("cht.sh/%s/%s", *lang, *query)

		cmd := exec.Command("curl", "-s", url)
		cmd.Stderr = os.Stderr

		lessCmd := exec.Command("less")
		lessCmd.Stdin, _ = cmd.StdoutPipe()
		lessCmd.Stdout = os.Stdout

		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}

		if err := lessCmd.Run(); err != nil {
			log.Fatal(err)

		}

		cmd.Wait()
	} else {
		url := fmt.Sprintf("cht.sh/%s", *lang)

		cmd := exec.Command("curl", "-s", url)
		cmd.Stderr = os.Stderr

		lessCmd := exec.Command("less")
		lessCmd.Stdin, _ = cmd.StdoutPipe()
		lessCmd.Stdout = os.Stdout

		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}

		if err := lessCmd.Run(); err != nil {
			log.Fatal(err)

		}

		cmd.Wait()
	}
}
