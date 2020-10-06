package ex_4_4_rotate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	a := []int{0, 1, 2, 3, 4, 5}
	b := []int{2, 3, 4, 5, 0, 1}
	rotate(a, 2)
	assert.Equal(t, b, a)
}
