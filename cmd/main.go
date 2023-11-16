package main

import (
	"fmt"
	"github.com/Selyss/chtsht/pkg/chtsht"
	"github.com/akamensky/argparse"
	// tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
)

func main() {

	parser := argparse.NewParser("cued", "Tool for querying programming keywords")
	topic := parser.String("p", "program", &argparse.Options{Required: false, Help: "Program to query", Default: ""})
	query := parser.String("q", "query", &argparse.Options{Required: false, Help: "Query", Default: ""})
	err := parser.Parse(os.Args)

	// TODO: refactor so we have one output no matter what and a set input place so i can wrap with a spinner
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	pager := os.Getenv("PAGER")
	if pager == "" {
		pager = "less"
	}

	// TODO: implement if theres no topic but there is a query
	// TODO: if theres lang but no topic look into lua/ and lua/:learn for general lang stuff

	// args
	if *topic != "" {
		// we got enough args to make it work
		chtsht.DisplayOutput(fmt.Sprintf("cht.sh/%s/%s", *topic, *query))
		return
	}

	// check config
	if config, err := chtsht.GetConfig(); config != nil {
		if err != nil {
			log.Fatalf("Error reading config, %s", err)
		}
		selection, err := chtsht.SelectFromList(config)
		if err != nil {
			log.Fatalf("Error while getting fzf selection: %s", err)
		}

		if selection == "asm" {
			chtsht.QueryASM()
			return
		}

		fmt.Print("Query: ")
		var query string
		fmt.Scanln(&query)
		url := fmt.Sprintf("cht.sh/%s/%s", selection, query)
		chtsht.DisplayOutput(url)
		return
	}

	// last resort
	if *topic == "" && *query == "" {
		// if theres still no topic and query
		opts, err := chtsht.ChtReadOptions() // TODO: add asm option
		if err != nil {
			log.Fatalf("Error reading default option file, %s", err)
		}

		selection, err := chtsht.SelectFromList(opts)
		if err != nil {
			log.Fatalf("Error while getting fzf selection: %s", err)
		}
		chtsht.DisplayOutput(fmt.Sprintf("cht.sh/%s", selection))
		return
	}

	chtsht.DisplayOutput(fmt.Sprintf("cht.sh/%s/%s", *topic, *query))
	return
}
