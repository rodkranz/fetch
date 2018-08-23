package main

import (
	"log"

	"github.com/rodkranz/fetch"
)

const url = "https://api.github.com/users/rodkranz"

func main() {
	f := fetch.NewDefault()
	rsp, err := f.Get(url)
	if err != nil {
		log.Fatalf("could not fetch [%s] because: %s", url, err)
	}

	body, err := rsp.String()
	if err != nil {
		log.Fatalf("could not retrieve body because: %s", err)
	}

	log.Println(body)
}
