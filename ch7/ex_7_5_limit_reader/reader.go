package reader

import "io"

type limitReader struct {
	r      io.Reader
	remain int64
}

func (r *limitReader) Read(p []byte) (n int, err error) {
	switch {
	case r.remain <= 0:
		n, err = 0, io.EOF
	case r.remain < int64(len(p)):
		n, err = r.r.Read(p[:r.remain])
		if err == nil {
			err = io.EOF
		}
	default:
		n, err = r.r.Read(p)
	}
	r.remain -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r, n}
}
