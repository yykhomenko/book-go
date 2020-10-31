package reader

import (
	"io"
)

type strReader struct {
	s string
}

func (r *strReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF
	}
	return
}

func NewReader(s string) io.Reader {
	return &strReader{s}
}
