package reader

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimitReader(t *testing.T) {
	r := strings.NewReader("Example")
	lr := LimitReader(r, 3)
	data := make([]byte, 10)
	n, err := lr.Read(data)
	assert.Equal(t, err, io.EOF)
	assert.Equal(t, 3, n)

	expected := make([]byte, 10)
	copy(expected, "Exa")
	assert.Equal(t, expected, data)
}
