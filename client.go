package fetch

import (
	"io"
	"net"
	"net/http"
	"time"
)

// DefaultTimeout It is default timeout of requests by default is 30seconds
var DefaultTimeout time.Duration

func init() {
	DefaultTimeout = time.Duration(30 * time.Second)
}

// NewDefault get fetcher with netTransport and timeout defined
func NewDefault() *Fetch {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: DefaultTimeout,
		}).Dial,
		TLSHandshakeTimeout: DefaultTimeout,
	}
	return New(netTransport)
}

// New get new fetcher and you need to specify the netTransport.
func New(netTransport *http.Transport) *Fetch {
	client := &http.Client{
		Transport: netTransport,
		Timeout:   DefaultTimeout,
	}

	return &Fetch{
		Client: client,
	}
}

// Fetch
type Fetch struct {
	*http.Client
	Header http.Header
}

func (f *Fetch) Do(request *http.Request) (*Response, error) {
	request.Header = f.Header
	resp, err := f.Client.Do(request)
	return &Response{resp}, err
}

// Get do request and with httpVerb GET
func (f *Fetch) Get(url string) (*Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return f.Do(req)
}

// Post do request and with httpVerb POST
func (f *Fetch) Post(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return nil, err
	}

	return f.Do(req)
}

// Put do request and with httpVerb PUT
func (f *Fetch) Put(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		return nil, err
	}

	return f.Do(req)
}

// Delete do request and with httpVerb DELETE
func (f *Fetch) Delete(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodDelete, url, reader)
	if err != nil {
		return nil, err
	}

	return f.Do(req)
}

// Delete do request and with httpVerb DELETE
func (f *Fetch) Patch(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodPatch, url, reader)
	if err != nil {
		return nil, err
	}

	return f.Do(req)
}

// Delete do request and with httpVerb DELETE
func (f *Fetch) Options(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodOptions, url, reader)
	if err != nil {
		return nil, err
	}

	return f.Do(req)
}
