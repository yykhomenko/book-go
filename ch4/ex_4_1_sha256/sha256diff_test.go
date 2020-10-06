package ex_4_1_sha256

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_diff(t *testing.T) {
	a := sha256.Sum256([]byte("x"))
	b := sha256.Sum256([]byte("X"))
	c := diff(a, b)
	assert.Equal(t, 125, c)
}
