package fetch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	*http.Response
}

// BodyEmpty error of body is empty
var BodyEmpty error = fmt.Errorf("the body of Response is empty")

// BodyIsEmpty return if body is empty or not.
func (r *Response) BodyIsEmpty() bool {
	return r.Body == nil
}

// Bytes return the Response in array of bytes.
func (r *Response) Bytes() ([]byte, error) {
	if r.BodyIsEmpty() {
		return nil, BodyEmpty
	}

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return bs, nil
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
	if r.BodyIsEmpty() {
		return BodyEmpty
	}

	return json.NewDecoder(r.Body).Decode(i)
}
