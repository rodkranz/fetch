[![Build Status](https://travis-ci.org/rodkranz/fetch.svg?branch=master)](https://travis-ci.org/rodkranz/fetch)
[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/rodkranz/fetch)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/rodkranz/fetch/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/rodkranz/fetch)](https://goreportcard.com/report/github.com/rodkranz/fetch)
 
# Fetch HTTP Client

Simple fetch made in Go to simplify the life of programmer.

## About
Go’s http package doesn’t specify request timeouts by default, allowing services to hijack your goroutines. **Always specify a custom http.Client** when connecting to outside services.


## Install

> Default 
```shell
go get github.com/rodkranz/fetch
```

> [Go DEP](https://github.com/golang/dep)
```shell
dep ensure --add github.com/rodkranz/fetch
```

## Import

```go
import (
  "github.com/rodkranz/fetch"
)
```

## Test 
To run the project test

```shell
go test -v --cover
```


## Example: 

#### Simple
    
```go
client := fetch.NewDefault()
response, err := client.Get("http://www.google.com/")
``` 

#### Custom Headers

```go
opt := fetch.Options{
    Header: http.Header{
        "Content-Type": []string{"application/json"},
        "User-Agent":   []string{"OLX-Group"},
    },
}

f := fetch.New(&opt)
rsp, err := f.GetWithContext(context.Background(), "http://www.google.com")
```

#### Simple JSON POST

```go
login := map[string]interface{}{
	"username": "rodkranz",
	"password": "loremIpsum",
}
response, err := fetch.NewDefault().
		IsJSON().
		Post("http://www.google.com/", fetch.NewReader(login))
```

  