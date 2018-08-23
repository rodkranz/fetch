package main

import (
	"net/http"
	"context"
	"fmt"
	"log"
	
	"github.com/rodkranz/fetch"
)

func main() {
	opt := fetch.Options{
		Header: http.Header{
			"Content-Type": []string{"application/json"},
			"User-Agent":   []string{"OLX-Group"},
		},
	}

	f := fetch.New(&opt)
	rsp, err := f.GetWithContext(context.Background(), "http://www.google.com")
	if err != nil {
		log.Fatalf("could not fetch data from target because: %s", err)
	}

	fmt.Println(rsp.StatusCode)
}
