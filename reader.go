package fetch

import (
	"encoding/json"
	"fmt"
	"strings"
)

// NewStructIO deprecated it will be replaced by NewReader
func NewStructIO(i interface{}) (*strings.Reader) {
	return NewReader(i)
}

func NewReader(i interface{}) (*strings.Reader) {
	bs, err := json.Marshal(i)
	if err != nil {
		return strings.NewReader(fmt.Sprintf("error to read: %T", i))
	}
	return strings.NewReader(string(bs))
}
