package fetch

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"strings"
	"fmt"
	"log"
	"time"
	"io"
	"context"
)

// NoBody is an io.ReadCloser with no bytes. Read always returns EOF
// and Close always returns nil. It can be used in an outgoing client
// request to explicitly signal that a request has zero bytes.
// An alternative, however, is to simply set Request.Body to nil.
var NoBody = noBody{}

type noBody struct{}

func (noBody) Read([]byte) (int, error)         { return 0, io.EOF }
func (noBody) Close() error                     { return nil }
func (noBody) WriteTo(io.Writer) (int64, error) { return 0, nil }

func TestDefaultOptions(t *testing.T) {
	var opt *Options

	opt = DefaultOptions()
	if opt.Timeout != DefaultTimeout {
		t.Errorf("Expected timeout [%v], but got [%v]", DefaultTimeout, opt.Timeout)
	}
	if opt.Host != "" {
		t.Errorf("Expected timeout [%v], but got [%v]", DefaultTimeout, opt.Timeout)
	}
	if opt.Header == nil {
		t.Error("Expected header already initialized, but got nil")
	}
}

func TestNewDefault(t *testing.T) {
	defaultFetch := NewDefault()
	if defaultFetch.Option == nil {
		t.Error("Expected Option not nil, but got nil")
	}
	if defaultFetch.Transport == nil {
		t.Error("Expected Transport is nil, but got nil")
	}
	if defaultFetch.Client == nil {
		t.Error("Expected Client not nil, but got nil")
	}
}

func TestNew(t *testing.T) {
	defaultFetch := New(nil)
	if defaultFetch.Option == nil {
		t.Error("Expected Option not nil, but got nil")
	}
	if defaultFetch.Transport == nil {
		t.Error("Expected Transport is nil, but got nil")
	}
	if defaultFetch.Client == nil {
		t.Error("Expected Client not nil, but got nil")
	}
}

func TestGetTransport(t *testing.T) {
	opt := Options{}
	getTransport(&opt)

	if opt.Timeout.Nanoseconds() == 0 {
		t.Error("Expected a Timeout defined, but got empty")
	}

	if opt.Transport == nil {
		t.Error("Expected a Transport defined, but got empty")
	}
}

func TestFetch_IsJSON(t *testing.T) {
	handlerContentTypeTest := func(writer http.ResponseWriter, request *http.Request) {
		contentType := request.Header.Get("Content-Type")
		if !strings.EqualFold(contentType, "application/json") {
			t.Errorf("Expected [Content-Type=application/json], but got [Content-Type=%s]", contentType)
		}
	}

	t.Run("Test-IsJSON", func(t *testing.T) {
		f := NewDefault().IsJSON()

		server := httptest.NewServer(http.HandlerFunc(handlerContentTypeTest))
		defer server.Close()

		f.Get(server.URL)
	})

	t.Run("Test-NoneHeaderIsJSON", func(t *testing.T) {
		f := New(&Options{}).IsJSON()

		server := httptest.NewServer(http.HandlerFunc(handlerContentTypeTest))
		defer server.Close()

		f.Get(server.URL)
	})
}

func TestMakeResponse(t *testing.T) {

	t.Run("test-MakeResponseOK", func(t *testing.T) {
		body := "Rodrigo Lopes not found"
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, body)
		}))
		defer ts.Close()

		res, err := NewDefault().Get(ts.URL)
		if err != nil {
			log.Fatal(err)
		}

		if bs, _ := res.String(); !strings.EqualFold(body, bs) {
			t.Errorf("Expected body [%s], but got [%s]", body, bs)
		}

		if http.StatusNotFound != res.StatusCode {
			t.Errorf("Expected body [%d], but got [%d]", http.StatusNotFound, res.StatusCode)
		}
	})

	t.Run("test-MakeResponseTimeout", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(100 * time.Millisecond)
		}))
		defer ts.Close()

		res, err := New(&Options{Timeout: time.Duration(10 * time.Millisecond)}).Get(ts.URL)
		if err == nil {
			t.Error("Expected timeout error, but got none error")
		}

		if http.StatusGatewayTimeout != res.StatusCode {
			t.Errorf("Expected body [%d], but got [%d]", http.StatusGatewayTimeout, res.StatusCode)
		}
	})
}

func serverHandlerMock(handlerFunc http.HandlerFunc) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(handlerFunc))
}

func TestFetch_Get(t *testing.T) {
	//handler := func(w http.ResponseWriter, r *http.Request) {
	//	w.WriteHeader(http.StatusNotFound)
	//	fmt.Fprint(w, "Hello World")
	//}

	//f := NewDefault()
	//t.Run("Test-Do-WrongRequest", func(t *testing.T) {
	//	rsp, err := f.Get("-")
	//	if _, ok := err.(*url.Error); !ok {
	//		t.Errorf("Expected error [*url.Error], but got [%T]", err)
	//	}
	//
	//	if http.StatusNoContent != rsp.StatusCode {
	//		t.Errorf("Expected status code [%d], but got [%d]", http.StatusNoContent, rsp.StatusCode)
	//	}
	//})

}

func TestFetch(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Hello World")
	}
	s := serverHandlerMock(handler)
	defer s.Close()

	type FetchWithoutContext func(string, io.Reader) (*Response, error)
	type FetchWithContext func(context.Context, string, io.Reader) (*Response, error)

	f := NewDefault()

	t.Run("Test-Method-GET", func(t *testing.T) {
		rsp, err := f.Get(s.URL)
		if err != nil {
			t.Errorf("Expected none error, but got [%s]", err)
		}

		if http.StatusNotFound != rsp.StatusCode {
			t.Errorf("Expected status code [%d], but got [%d]", http.StatusNoContent, rsp.StatusCode)
		}
	})
	t.Run("Test-Method-GET-With-Context", func(t *testing.T) {
		rsp, err := f.GetWithContext(context.Background(), s.URL)
		if err != nil {
			t.Errorf("Expected none error, but got [%s]", err)
		}

		if http.StatusNotFound != rsp.StatusCode {
			t.Errorf("Expected status code [%d], but got [%d]", http.StatusNoContent, rsp.StatusCode)
		}
	})

	testWithoutContext := []struct {
		Name   string
		Method FetchWithoutContext
	}{
		{Name: "DELETE", Method: f.Delete},
		{Name: "POST", Method: f.Post},
		{Name: "PUT", Method: f.Put},
		{Name: "OPTIONS", Method: f.Options},
		{Name: "PATCH", Method: f.Patch},
	}

	for _, test := range testWithoutContext {
		t.Run(fmt.Sprintf("Test-Method-%s", test.Name), func(t *testing.T) {
			rsp, err := test.Method(s.URL, NoBody)
			if err != nil {
				t.Errorf("Expected none error, but got [%s]", err)
			}

			if http.StatusNotFound != rsp.StatusCode {
				t.Errorf("Expected status code [%d], but got [%d]", http.StatusNoContent, rsp.StatusCode)
			}
		})
	}

	testWithContext := []struct {
		Name   string
		Method FetchWithContext
	}{
		{Name: "DELETE", Method: f.DeleteWithContext},
		{Name: "POST", Method: f.PostWithContext},
		{Name: "PUT", Method: f.PutWithContext},
		{Name: "OPTIONS", Method: f.OptionsWithContext},
		{Name: "PATCH", Method: f.PatchWithContext},
	}

	for _, test := range testWithContext {
		t.Run(fmt.Sprintf("Test-Method-With-Context-%s", test.Name), func(t *testing.T) {
			rsp, err := test.Method(context.Background(), s.URL, NoBody)
			if err != nil {
				t.Errorf("Expected none error, but got [%s]", err)
			}

			if http.StatusNotFound != rsp.StatusCode {
				t.Errorf("Expected status code [%d], but got [%d]", http.StatusNoContent, rsp.StatusCode)
			}
		})
	}

}
