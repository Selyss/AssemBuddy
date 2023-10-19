package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Selyss/chtsht/pkg/chtsht"
)

func main() {
	// accept lang cli args
	flag.Usage = func() {
		fmt.Printf(`
    Usage:

        cued [OPTIONS|QUERY]

    Options:

        QUERY                   process QUERY and exit

        --help                  show this help page
        --shell [LANG]          shell mode (open LANG if specified)
        
        --config [PATH]         set language option path

        --standalone-install [DIR|help]
                                install cheat.sh in the standalone mode
                                (by default, into ~/.cheat.sh/)

    `)
	}

	config, err := chtsht.NewConfig()

}
