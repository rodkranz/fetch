package fetch

import (
	"encoding/json"
	"fmt"
	"strings"
)

// NewStructIO this function will create a new
// reader for your request.
// if input is json format will convert to json or send directly
//
// Deprecated: Use NewReader instead.
func NewStructIO(input interface{}) *strings.Reader {
	return NewReader(input)
}

// NewReader this function will return a new reader
// if the format is JSON Valid format it will be convert
// before send, if is not json will send as come.
func NewReader(input interface{}) *strings.Reader {
	bs, err := json.Marshal(input)
	if err != nil {
		return strings.NewReader(fmt.Sprintf("error to read: %T", input))
	}
	return strings.NewReader(string(bs))
}
