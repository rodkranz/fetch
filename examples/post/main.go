package main

import (
	"log"
	
	"github.com/rodkranz/fetch"
)

const targetURL = "http://website.com/"

type LoginForm struct {
	Username string
	Password string
}

func main() {
	login := LoginForm{
		Username: "username",
		Password: "password",
	}

	rsp , err := fetch.NewDefault().IsJSON().Post(targetURL, fetch.NewReader(login))
	if err != nil {
		log.Fatalf("could not login because: %s", err)
	}

	body := fetch.MustString(rsp.String())

	log.Println(body)
}
