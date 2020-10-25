package ex_5_15_min_max

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_min(t *testing.T) {
	assert.Equal(t, 5, min(5))
	assert.Equal(t, -5, min(5, 1, -5, -2))
}

func Test_max(t *testing.T) {
	assert.Equal(t, 5, max(5))
	assert.Equal(t, 5, max(5, 1, -5, -2))
}
