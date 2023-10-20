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
	topic := parser.String("p", "program", &argparse.Options{Required: false, Help: "Program to query", Default: ""})
	query := parser.String("q", "query", &argparse.Options{Required: false, Help: "Query", Default: ""})
	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	// use pager env var if possible
	pager := os.Getenv("PAGER")
	if pager == "" {
		pager = "less"
	}

	// TODO: impl later
	if *topic == "" {
		log.Fatal(err)
	}

	url := "cht.sh/%s"

	if *query != "" {
		url = fmt.Sprintf(url+"/%s", *topic, *query)
	} else {
		url = fmt.Sprintf(url, *topic)
	}

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
