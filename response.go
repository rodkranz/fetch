package fetch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ErrEmptyBody returns when there is no body to read
var ErrEmptyBody = fmt.Errorf("the body of response is empty")

// Response helper work with response from http.Client
type Response struct {
	*http.Response
	body []byte
}

// BodyIsEmpty return if body is empty or not.
func (r *Response) BodyIsEmpty() bool {
	return r.body == nil || len(r.body) == 0
}

// Bytes return the Response in array of bytes.
func (r *Response) Bytes() (_ []byte, err error) {
	// if body is not empty return itself
	if !r.BodyIsEmpty() {
		return r.body, nil
	}

	// if Body is empty
	if r.Response == nil || r.Response.Body == nil {
		return nil, ErrEmptyBody
	}

	r.body, err = ioutil.ReadAll(r.Body)
	return r.body, err
}

// String return the Response in string format.
func (r *Response) String() (string, error) {
	bs, err := r.Bytes()
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

// Decode body result into interface object.
func (r *Response) Decode(i interface{}) error {
	body, err := r.Bytes()
	if err != nil {
		return err
	}

	return json.Unmarshal(body, i)
}

// newErrorResponse return response if has any kind of error
// could be from request or execution of Http.Client
func newErrorResponse(status int, msg string, err error) (*Response, error) {
	return &Response{
		Response: &http.Response{
			StatusCode: status,
			Status:     http.StatusText(status),
		},
	}, fmt.Errorf(msg, err)
}