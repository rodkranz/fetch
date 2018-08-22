package fetch

import (
	"io"
	"net"
	"net/http"
	"time"
	"context"

	"golang.org/x/net/context/ctxhttp"
	)

// DefaultTimeout defined timeout default for any request
const DefaultTimeout = time.Duration(30 * time.Second)

// Options default for any request in client
type Options struct {
	Header    http.Header
	Timeout   time.Duration
	Host      string
	Transport *http.Transport
}

// DefaultOptions returns options with timeout defined
func DefaultOptions() *Options {
	return &Options{
		Timeout: DefaultTimeout,
		Header:  http.Header{},
	}
}

// NewDefault get fetcher with netTransport and timeout defined
func NewDefault() *Fetch {
	return New(DefaultOptions())
}

// getTransport make transport from options definitions
func getTransport(opt *Options) {
	if opt.Timeout.Nanoseconds() == 0 {
		opt.Timeout = DefaultTimeout
	}

	opt.Transport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: opt.Timeout,
		}).DialContext,
		TLSHandshakeTimeout: opt.Timeout,
	}
}

// New get new fetcher and you need to specify the netTransport.
func New(opt *Options) *Fetch {
	if opt == nil {
		opt = DefaultOptions()
	}

	if opt.Transport == nil {
		getTransport(opt)
	}

	client := &http.Client{
		Timeout: opt.Timeout,
	}

	return &Fetch{
		Client: client,
		Option: opt,
	}
}

// Fetch
type Fetch struct {
	*http.Client
	Option *Options
}

func (f *Fetch) IsJSON() (*Fetch) {
	if f.Option.Header == nil {
		f.Option.Header = http.Header{}
	}

	f.Option.Header.Set("Content-Type", "application/json")
	return f
}

func (f *Fetch) DoWithContext(ctx context.Context, req *http.Request) (*Response, error) {
	if f.Option.Header != nil {
		req.Header = f.Option.Header
	}

	return f.makeResponse(ctxhttp.Do(ctx, f.Client, req))
}

func (f *Fetch) Do(req *http.Request) (*Response, error) {
	if f.Option.Header != nil {
		req.Header = f.Option.Header
	}

	return f.makeResponse(f.Client.Do(req))
}

func (f Fetch) makeResponse(resp *http.Response, err error) (*Response, error) {
	if resp == nil {
		resp = &http.Response{
			StatusCode: http.StatusGatewayTimeout,
			Status:     http.StatusText(http.StatusGatewayTimeout),
		}
	}

	return &Response{Response: resp}, err
}

func (f *Fetch) GetWithContext(ctx context.Context, url string) (*Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request GET: %s", err)
	}

	return f.DoWithContext(ctx, req)
}

// Get do request and with httpVerb GET
func (f *Fetch) Get(url string) (*Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request GET: %s", err)
	}

	return f.Do(req)
}

func (f *Fetch) PostWithContext(ctx context.Context, url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request POST: %s", err)
	}

	return f.DoWithContext(ctx, req)
}

// Post do request and with httpVerb POST
func (f *Fetch) Post(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request POST: %s", err)
	}

	return f.Do(req)
}

func (f *Fetch) PutWithContext(ctx context.Context, url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request PUT: %s", err)
	}
	return f.DoWithContext(ctx, req)
}

// Put do request and with httpVerb PUT
func (f *Fetch) Put(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request PUT: %s", err)
	}

	return f.Do(req)
}

func (f *Fetch) DeleteWithContext(ctx context.Context, url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodDelete, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request DELETE: %s", err)
	}

	return f.DoWithContext(ctx, req)
}

// Delete do request and with httpVerb DELETE
func (f *Fetch) Delete(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodDelete, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request DELETE: %s", err)
	}
	return f.Do(req)
}

// Patch do request and with httpVerb PATCH
func (f *Fetch) Patch(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodPatch, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request PATCH: %s", err)
	}
	return f.Do(req)
}

// PatchWithContext do request and with httpVerb PATCH
func (f *Fetch) PatchWithContext(ctx context.Context, url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodPatch, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request PATCH: %s", err)
	}
	return f.DoWithContext(ctx, req)
}

// Options do request and with httpVerb DELETE
func (f *Fetch) Options(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodOptions, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request OPTIONS: %s", err)
	}
	return f.Do(req)
}

// OptionsWithContext do request and with httpVerb OPTIONS
func (f *Fetch) OptionsWithContext(ctx context.Context, url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodOptions, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request OPTIONS: %s", err)
	}
	return f.DoWithContext(ctx, req)
}
