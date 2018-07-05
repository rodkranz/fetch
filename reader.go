package fetch

import (
	"encoding/json"
	"io"
	"bytes"
	"fmt"
)

type StructIO struct {
	original []byte
	b        []byte
	reader   io.Reader
	writer   io.Writer

	err error
}

func NewStructIO(i interface{}) (*StructIO) {
	bs, err := json.Marshal(i)
	return &StructIO{b: bs, original: bs, err: err}
}

func (s *StructIO) Flush() {
	s.b = make([]byte, len(s.original))
	copy(s.b, s.original)
}

func (s *StructIO) Write(p []byte) (n int, err error) {
	if s.writer == nil {
		s.writer = bytes.NewBuffer(s.b)
	}

	return s.writer.Write(p)
}

func (s *StructIO) Read(p []byte) (n int, err error) {
	if s.reader == nil {
		s.reader = bytes.NewReader(s.b)
	}

	return s.reader.Read(p)
}

func (s *StructIO) Error() (string) {
	return fmt.Sprintf("error struct io: %s", s.err)
}

func (s *StructIO) HasError() (bool) {
	return s.err != nil
}
