package fetch

import (
	"net/http"
	"time"
)

type Options struct {
	Header  http.Header
	Timeout time.Duration
	host    string
}
