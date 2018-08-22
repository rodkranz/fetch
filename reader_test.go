package fetch

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

var readerTest = []struct {
	desc   string
	input  interface{}
	output string
}{
	{desc: "Empty", input: "", output: `""`},
	{desc: "Int", input: 2004, output: `2004`},
	{desc: "String", input: "Lorem Ipsum", output: `"Lorem Ipsum"`},
	{
		desc: "Struct",
		input: struct {
			Title string `json:"title"`
		}{
			Title: "Rodrigo",
		},
		output: `{"title":"Rodrigo"}`,
	},
	{
		desc:   "SliceString",
		input:  []string{"Rodrigo", "Lopes"},
		output: `["Rodrigo","Lopes"]`,
	},
	{
		desc:   "WrongFormat",
		input:  func() {},
		output: "error to read: func()",
	},
}

// TestNewStructIO func will be
// deprecated soon
func TestNewStructIO(t *testing.T) {
	for i, test := range readerTest {
		t.Run(fmt.Sprintf("%d-test-%s", i, test.desc), func(t *testing.T) {
			r := NewStructIO(test.input)

			bs, err := ioutil.ReadAll(r)
			if err != nil {
				t.Fatalf("Expected none error, but got [%v]", err)
			}

			output := string(bs)
			if !strings.EqualFold(test.output, output) {
				t.Errorf("Expected [%s] but got [%s]", test.output, output)
			}
		})
	}
}

func TestNewReader(t *testing.T) {
	for i, test := range readerTest {
		t.Run(fmt.Sprintf("%d-test-%s", i, test.desc), func(t *testing.T) {
			r := NewReader(test.input)

			bs, err := ioutil.ReadAll(r)
			if err != nil {
				t.Fatalf("Expected none error, but got [%v]", err)
			}

			output := string(bs)
			if !strings.EqualFold(test.output, output) {
				t.Errorf("Expected [%s] but got [%s]", test.output, output)
			}
		})
	}
}
