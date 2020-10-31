package reader

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLimitReader(t *testing.T) {
	r := strings.NewReader("Example")
	lr := LimitReader(r, 3)
	data := make([]byte, 10)
	n, err := lr.Read(data)
	assert.Error(t, err)
	assert.Equal(t, 3, n)
}
