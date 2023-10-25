package main

import (
	"bufio"
	"fmt"
	"github.com/Selyss/chtsht/pkg/chtsht"
	"github.com/akamensky/argparse"
	"io"
	"log"
	"os"

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
	if *topic == "" && *query != "" {
		log.Fatalf("Error, query but no language: %s", err)
	}

	// if there are args we wanna process them
	if *topic != "" {
		url := "cht.sh/%s"

		if *query != "" {
			url = fmt.Sprintf(url+"/%s", *topic, *query)
		} else {
			// TODO: if theres lang but no topic look into lua/ and lua/:learn for general lang stuff
			url = fmt.Sprintf(url, *topic)
		}

		cmd := exec.Command("curl", "-s", url)
		cmd.Stderr = os.Stderr

		lessCmd := exec.Command("less")
		lessCmd.Stdin, _ = cmd.StdoutPipe()
		lessCmd.Stdout = os.Stdout

		if err := cmd.Start(); err != nil {
			log.Fatalf("Error while querying: %s", err)
		}

		if err := lessCmd.Run(); err != nil {
			log.Fatalf("Error while piping into $PAGER: %s", err)
		}

		cmd.Wait()

		// regular fzf
	} else {
		// // get lang config
		config, err := chtsht.GetConfig()
		if err != nil {
			// log.Fatalf("Error while getting config: %s", err)
			//
			// get list of topic opts
			readFile, err := os.Open("../chtsht.txt")

			if err != nil {
				fmt.Println(err)
			}
			fileScanner := bufio.NewScanner(readFile)
			fileScanner.Split(bufio.ScanLines)
			var fileLines []string

			for fileScanner.Scan() {
				fileLines = append(fileLines, fileScanner.Text())
			}

			readFile.Close()

			for _, line := range fileLines {
				fmt.Println(line)
			}

			chtsht.SelectFromList(fileLines)

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
		cmd := exec.Command("curl", "-s", url)
		cmd.Stderr = os.Stderr

		lessCmd := exec.Command("less")
		lessCmd.Stdin, _ = cmd.StdoutPipe()
		lessCmd.Stdout = os.Stdout

		if err := cmd.Start(); err != nil {
			log.Fatalf("Error while querying: %s", err)
		}

		if err := lessCmd.Run(); err != nil {
			log.Fatalf("Error while piping into $PAGER: %s", err)
		}

		cmd.Wait()

		// if no config make a list with all options
	}
}
