package writer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingWriter(t *testing.T) {
	buf := &bytes.Buffer{}
	w, n := CountingWriter(buf)
	w.Write([]byte("Hello world!"))

	assert.Equal(t, 12, buf.Len())
	assert.Equal(t, 12, *n)
}
