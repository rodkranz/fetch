package fetch

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestResponse_BodyIsEmpty(t *testing.T) {
	t.Run("Test-NoBodyValue", func(t *testing.T) {
		res := Response{body: nil}
		if !res.BodyIsEmpty() {
			t.Fatalf("Expecetd no body, but got [%s]", res.body)
		}
	})

	t.Run("Test-BodyWithValue", func(t *testing.T) {
		body := []byte("Lorem Ipsum")
		res := Response{body: body}
		if res.BodyIsEmpty() {
			t.Fatalf("Expecetd body [%s], but got [%s]", body, res.body)
		}
	})
}

func TestResponse_Bytes(t *testing.T) {
	t.Run("Test-BytesNil", func(t *testing.T) {
		res := Response{body: nil}
		bs, err := res.Bytes()
		if err == nil {
			t.Errorf("Expecetd error [%s], but got nil", ErrEmptyBody)
		}
		if err != ErrEmptyBody {
			t.Errorf("Expecetd error [%s], but got [%s]", ErrEmptyBody, err)
		}
		if len(bs) > 0 {
			t.Errorf("Expecetd none result bytes, but got [%s]", bs)
		}
	})
	t.Run("Test-BytesCached", func(t *testing.T) {
		body := []byte("Lorem Ipsum")
		res := Response{body: body}
		bs, err := res.Bytes()
		if err != nil {
			t.Errorf("Expecetd nil error, but got [%s]", err)
		}
		if len(bs) <= 0 {
			t.Errorf("Expecetd result bytes, but got [%s]", bs)
		}
		if !bytes.EqualFold(body, bs) {
			t.Errorf("Expecetd [%s] as bytes, but got [%s] as bytes", body, bs)
		}
	})
	t.Run("Test-BytesBodyRequested", func(t *testing.T) {
		body := []byte("Lorem Ipsum")

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { _, _ = w.Write(body) }))
		defer ts.Close()

		res, err := http.Get(ts.URL)
		if err != nil {
			log.Fatal(err)
		}

		response := &Response{Response: res}

		bs, err := response.Bytes()
		if err != nil {
			t.Errorf("Expecetd nil error, but got [%s]", err)
		}
		if len(bs) <= 0 {
			t.Errorf("Expecetd result bytes, but got [%s]", bs)
		}
		if !bytes.EqualFold(body, bs) {
			t.Errorf("Expecetd [%s] as bytes, but got [%s] as bytes", body, bs)
		}
	})
}

func TestResponse_String(t *testing.T) {
	t.Run("Test-StringNil", func(t *testing.T) {
		res := Response{body: nil}
		output := res.String()
		if len(output) > 0 {
			t.Errorf("Expecetd none result bytes, but got [%s]", output)
		}
	})
	t.Run("Test-StringCached", func(t *testing.T) {
		body := "Lorem Ipsum"
		res := Response{body: []byte(body)}
		output := res.String()
		if len(output) <= 0 {
			t.Errorf("Expecetd result bytes, but got [%s]", output)
		}
		if !strings.EqualFold(body, output) {
			t.Errorf("Expecetd [%s] as bytes, but got [%s] as bytes", body, output)
		}
	})
}

func TestResponse_ToString(t *testing.T) {
	t.Run("Test-ToStringNil", func(t *testing.T) {
		res := Response{body: nil}
		output, err := res.ToString()
		if err == nil {
			t.Errorf("Expecetd error [%s], but got nil", ErrEmptyBody)
		}
		if err != ErrEmptyBody {
			t.Errorf("Expecetd error [%s], but got [%s]", ErrEmptyBody, err)
		}
		if len(output) > 0 {
			t.Errorf("Expecetd none result bytes, but got [%s]", output)
		}
	})
	t.Run("Test-ToStringCached", func(t *testing.T) {
		body := "Lorem Ipsum"
		res := Response{body: []byte(body)}
		output, err := res.ToString()
		if err != nil {
			t.Errorf("Expecetd nil error, but got [%s]", err)
		}
		if len(output) <= 0 {
			t.Errorf("Expecetd result bytes, but got [%s]", output)
		}
		if !strings.EqualFold(body, output) {
			t.Errorf("Expecetd [%s] as bytes, but got [%s] as bytes", body, output)
		}
	})
}

func TestResponse_Decode(t *testing.T) {
	t.Run("Test-DecodeError", func(t *testing.T) {
		res := Response{body: nil}
		err := res.Decode(&struct{}{})
		if err == nil {
			t.Errorf("Expecetd error [%s], but got nil", ErrEmptyBody)
		}
		if err != ErrEmptyBody {
			t.Errorf("Expecetd error [%s], but got [%s]", ErrEmptyBody, err)
		}
	})
	t.Run("Test-Decode", func(t *testing.T) {
		var name, age = "Rodrigo Lopes", 30

		body := fmt.Sprintf(`{"name": "%s", "age": %d}`, name, age)
		res := Response{body: []byte(body)}
		output := struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{}

		err := res.Decode(&output)
		if err != nil {
			t.Errorf("Expecetd nil error, but got [%s]", err)
		}
		if output.Name != name {
			t.Errorf("Expecetd [%s], but got [%s]", output.Name, name)
		}
		if output.Age != age {
			t.Errorf("Expecetd [%d], but got [%d]", output.Age, age)
		}
	})
}

func TestNewErrorResponse(t *testing.T) {
	status := http.StatusNotFound
	msg := "could not test this because, %s"
	errTest := fmt.Errorf("nobody test")

	rsp, err := newErrorResponse(status, msg, errTest)
	if rsp == nil {
		t.Fatalf("Expecetd response not nil, but got nil")
	}
	if err == nil {
		t.Fatalf("Expecetd error not nil, but got nil")
	}
	if status != rsp.StatusCode {
		t.Errorf("Expecetd status code [%d], but got status code [%d]", status, rsp.StatusCode)
	}
	if http.StatusText(status) != rsp.Status {
		t.Errorf("Expecetd status [%s], but got status [%s]", http.StatusText(status), rsp.Status)
	}
	if !strings.EqualFold(fmt.Sprintf(msg, errTest.Error()), err.Error()) {
		t.Errorf("Expecetd error [%s], but got error [%s]", fmt.Sprintf(msg, errTest.Error()), err.Error())
	}
}
