package main

import (
	"github.com/Selyss/chtsht/pkg/chtsht"

	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.syscall.sh/v1/syscalls/x64")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	chtsht.DrawTable(string(body))

}
