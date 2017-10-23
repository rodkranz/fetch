package fetch

import (
	"io"
	"net"
	"net/http"
	"time"
	"fmt"
)

func getTransport(opt *Options) *http.Transport {
	return &http.Transport{
		Dial: (&net.Dialer{
			Timeout: opt.Timeout,
		}).Dial,
		TLSHandshakeTimeout: opt.Timeout,
	}
}

// NewDefault get fetcher with netTransport and timeout defined
func NewDefault() *Fetch {
	return New(&Options{
		Timeout: time.Duration(30 * time.Second),
	})
}

// New get new fetcher and you need to specify the netTransport.
func New(opt *Options) *Fetch {
	if opt == nil {
		opt = &Options{
			Timeout: time.Duration(30 * time.Second),
		}
	}

	client := &http.Client{
		Transport: getTransport(opt),
		Timeout:   opt.Timeout,
	}

	return &Fetch{
		Client: client,
		opt:    opt,
	}
}

// Fetch
type Fetch struct {
	*http.Client
	opt *Options
}

func (f *Fetch) Do(req *http.Request) (*Response, error) {
	if f.opt != nil
	req.Header = f.opt.Header
	resp, err := f.Client.Do(req)
	return &Response{Response: resp}, err
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
