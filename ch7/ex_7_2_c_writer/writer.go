package writer

import "io"

type CWriter struct {
	w io.Writer
	c int64
}

func (c *CWriter) Write(p []byte) (n int, err error) {
	n, err = c.w.Write(p)
	c.c += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &CWriter{w, 0}
	return c, &c.c
}
