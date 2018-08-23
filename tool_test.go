package fetch

import (
	"testing"
	"fmt"
	"strings"
	"bytes"
)

func TestMustString(t *testing.T) {
	tests := []struct {
		input  func() (string, error)
		output string
	}{
		{output: "", input: func() (string, error) {
			return "", nil
		}},
		{output: "Bananas", input: func() (string, error) {
			return "Bananas", nil
		}},
		{output: "bananas", input: func() (string, error) {
			return "bananas", ErrEmptyBody
		}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-must-string-%d", i), func(t *testing.T) {
			output := MustString(test.input())
			if !strings.EqualFold(output, test.output) {
				t.Errorf("Expected [%s] output, but got [%s]", test.output, output)
			}
		})
	}
}

func TestMustBytes(t *testing.T) {
	tests := []struct {
		input  func() ([]byte, error)
		output []byte
	}{
		{output: []byte(""), input: func() ([]byte, error) {
			return []byte(""), nil
		}},
		{output: []byte("bananas"), input: func() ([]byte, error) {
			return []byte("bananas"), nil
		}},
		{output: []byte("bananas"), input: func() ([]byte, error) {
			return []byte("bananas"), ErrEmptyBody
		}},
		{output: []byte(""), input: func() ([]byte, error) {
			return nil, nil
		}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-must-string-%d", i), func(t *testing.T) {
			output := MustBytes(test.input())
			fmt.Println(output, test.output)

			if !bytes.Equal(output, test.output) {
				t.Errorf("Expected [%s] output, but got [%s]", test.output, output)
			}
		})
	}
}
