package fetch

import (
	"context"
	"io"
	"net"
	"net/http"
	"time"

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

	return &Fetch{
		Client: &http.Client{
			Timeout:   opt.Timeout,
			Transport: opt.Transport,
		},
		Option: opt,
	}
}

// Fetch use http default but defined with a timeout.
type Fetch struct {
	*http.Client
	Option *Options
}

// IsJSON add Content-Type as JSON in header.
func (f *Fetch) IsJSON() *Fetch {
	if f.Option.Header == nil {
		f.Option.Header = http.Header{}
	}

	f.Option.Header.Set("Content-Type", "application/json")
	return f
}

// DoWithContext execute any kind of request passing context
func (f *Fetch) DoWithContext(ctx context.Context, req *http.Request) (*Response, error) {
	if f.Option.Header != nil {
		req.Header = f.Option.Header
	}

	return f.makeResponse(ctxhttp.Do(ctx, f.Client, req))
}

// Do execute any kind of request
func (f *Fetch) Do(req *http.Request) (*Response, error) {
	if f.Option.Header != nil {
		req.Header = f.Option.Header
	}

	return f.makeResponse(f.Client.Do(req))
}

// makeResponse format response from generic request
func (f Fetch) makeResponse(resp *http.Response, err error) (*Response, error) {
	if resp == nil {
		resp = &http.Response{
			StatusCode: http.StatusGatewayTimeout,
			Status:     http.StatusText(http.StatusGatewayTimeout),
		}
	}

	return &Response{Response: resp}, err
}

// Get do request with HTTP using HTTP Verb GET
func (f *Fetch) Get(url string) (*Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request GET: %s", err)
	}

	return f.Do(req)
}

// Post do request with HTTP using HTTP Verb GET
func (f *Fetch) Post(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request POST: %s", err)
	}

	return f.Do(req)
}

// Put do request with HTTP using HTTP Verb PUT
func (f *Fetch) Put(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request PUT: %s", err)
	}

	return f.Do(req)
}

// Delete do request with HTTP using HTTP Verb DELETE
func (f *Fetch) Delete(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodDelete, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request DELETE: %s", err)
	}
	return f.Do(req)
}

// Patch do request with HTTP using HTTP Verb PATCH
func (f *Fetch) Patch(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodPatch, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request PATCH: %s", err)
	}
	return f.Do(req)
}

// Options do request with HTTP using HTTP Verb OPTIONS
func (f *Fetch) Options(url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodOptions, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request OPTIONS: %s", err)
	}
	return f.Do(req)
}

// GetWithContext execute DoWithContext but define request to method GET
func (f *Fetch) GetWithContext(ctx context.Context, url string) (*Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request GET: %s", err)
	}

	return f.DoWithContext(ctx, req)
}

// PostWithContext execute DoWithContext but define request to method POST
func (f *Fetch) PostWithContext(ctx context.Context, url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request POST: %s", err)
	}

	return f.DoWithContext(ctx, req)
}

// PutWithContext execute DoWithContext but define request to method PUT
func (f *Fetch) PutWithContext(ctx context.Context, url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request PUT: %s", err)
	}
	return f.DoWithContext(ctx, req)
}

// DeleteWithContext execute DoWithContext but define request to method DELETE
func (f *Fetch) DeleteWithContext(ctx context.Context, url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodDelete, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request DELETE: %s", err)
	}

	return f.DoWithContext(ctx, req)
}

// PatchWithContext execute DoWithContext but define request to method PATCH
func (f *Fetch) PatchWithContext(ctx context.Context, url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodPatch, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request PATCH: %s", err)
	}
	return f.DoWithContext(ctx, req)
}

// OptionsWithContext execute DoWithContext but define request to method OPTIONS
func (f *Fetch) OptionsWithContext(ctx context.Context, url string, reader io.Reader) (*Response, error) {
	req, err := http.NewRequest(http.MethodOptions, url, reader)
	if err != nil {
		return newErrorResponse(http.StatusNoContent, "couldn't request OPTIONS: %s", err)
	}
	return f.DoWithContext(ctx, req)
}
