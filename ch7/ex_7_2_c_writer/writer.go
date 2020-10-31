package writer

import "io"

type CWriter struct {
	w io.Writer
	c int64
}

func (c *CWriter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	if err != nil {
		return 0, err
	}
	c.c += int64(n)
	return n, nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &CWriter{w: w}
	return c, &c.c
}
